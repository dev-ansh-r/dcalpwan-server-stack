package distribution

import (
	"context"
	"sync"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

// newSubscriptionMap creates a mapping between application identifiers and subscription sets.
// The timeout represents the period after which a set will shut down if empty. If the timeout
// is zero, the sets never timeout.
func newSubscriptionMap(ctx context.Context, rd RequestDecoupler, timeout time.Duration, setup func(*subscriptionSet, *ttnpb.ApplicationIdentifiers) error, opts ...io.SubscriptionOption) *subscriptionMap {
	return &subscriptionMap{
		ctx:     ctx,
		rd:      rd,
		timeout: timeout,
		setup:   setup,
		subOpts: opts,
	}
}

type subscriptionMap struct {
	ctx      context.Context
	rd       RequestDecoupler
	timeout  time.Duration
	setup    func(*subscriptionSet, *ttnpb.ApplicationIdentifiers) error
	decouple func(context.Context) context.Context
	sets     sync.Map
	subOpts  []io.SubscriptionOption
}

type subscriptionMapSet struct {
	set *subscriptionSet

	init    chan struct{}
	initErr error
}

var errSetNotFound = errors.DefineNotFound("set_not_found", "set not found")

// Load loads the subscription set associated with the application identifiers.
func (m *subscriptionMap) Load(ctx context.Context, ids *ttnpb.ApplicationIdentifiers) (*subscriptionSet, error) {
	uid := unique.ID(ctx, ids)
	existing, ok := m.sets.Load(uid)
	if !ok {
		return nil, errSetNotFound.New()
	}
	exists := existing.(*subscriptionMapSet)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-exists.init:
	}
	if exists.initErr != nil {
		return nil, exists.initErr
	}
	return exists.set, nil
}

// LoadOrCreate loads the subscription set associated with the application identifiers.
// If the subscription set does not exist, it is created.
func (m *subscriptionMap) LoadOrCreate(ctx context.Context, ids *ttnpb.ApplicationIdentifiers) (*subscriptionSet, error) {
	uid := unique.ID(ctx, ids)
	s := &subscriptionMapSet{
		init: make(chan struct{}),
	}
	if existing, loaded := m.sets.LoadOrStore(uid, s); loaded {
		exists := existing.(*subscriptionMapSet)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-exists.init:
		}
		if exists.initErr != nil {
			return nil, exists.initErr
		}
		return exists.set, nil
	}

	var err error
	defer func() {
		if err != nil {
			s.initErr = err
			m.sets.Delete(uid)
		}
		close(s.init)
	}()

	ctx = log.NewContextWithField(m.ctx, "application_uid", uid)
	ctx, err = unique.WithContext(ctx, uid)
	if err != nil {
		return nil, err
	}

	set := newSubscriptionSet(ctx, m.rd, m.timeout, m.subOpts...)
	if err = m.setup(set, ids); err != nil {
		set.Cancel(err)
		return nil, err
	}
	go func() {
		<-set.Context().Done()
		m.sets.Delete(uid)
	}()
	s.set = set

	return set, nil
}

func noSetup(*subscriptionSet, *ttnpb.ApplicationIdentifiers) error {
	return nil
}

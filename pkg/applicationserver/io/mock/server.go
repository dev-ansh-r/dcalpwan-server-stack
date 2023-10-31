package mock

import (
	"context"
	"sync"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/ratelimit"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

type server struct {
	*component.Component
	subscriptionsMu sync.RWMutex
	appSubs         map[string][]*io.Subscription
	bcastSubs       []*io.Subscription
	subscriptionsCh chan *io.Subscription
	downlinkQueueMu sync.RWMutex
	downlinkQueue   map[string][]*ttnpb.ApplicationDownlink
	subscribeError  error
}

// Server represents a mock io.Server.
type Server interface {
	io.Server

	SetSubscribeError(error)
	Subscriptions() <-chan *io.Subscription
}

// NewServer instantiates a new Server.
func NewServer(c *component.Component) Server {
	return &server{
		Component:       c,
		appSubs:         make(map[string][]*io.Subscription),
		subscriptionsCh: make(chan *io.Subscription, 10),
		downlinkQueue:   make(map[string][]*ttnpb.ApplicationDownlink),
	}
}

// FillContext implements io.Server.
func (s *server) FillContext(ctx context.Context) context.Context {
	return s.Component.FillContext(ctx)
}

func (s *server) Publish(ctx context.Context, up *ttnpb.ApplicationUp) error {
	s.subscriptionsMu.RLock()
	defer s.subscriptionsMu.RUnlock()
	for _, sub := range s.appSubs[unique.ID(ctx, up.EndDeviceIds.ApplicationIds)] {
		if err := sub.Publish(ctx, up); err != nil {
			return err
		}
	}
	for _, sub := range s.bcastSubs {
		if err := sub.Publish(ctx, up); err != nil {
			return err
		}
	}
	return nil
}

// Subscribe implements io.Server.
func (s *server) Subscribe(ctx context.Context, protocol string, ids *ttnpb.ApplicationIdentifiers, global bool) (*io.Subscription, error) {
	s.subscriptionsMu.RLock()
	err := s.subscribeError
	s.subscriptionsMu.RUnlock()
	if err != nil {
		return nil, err
	}
	sub := io.NewSubscription(ctx, protocol, ids)
	s.subscriptionsMu.Lock()
	if ids != nil {
		s.appSubs[unique.ID(ctx, ids)] = append(s.appSubs[unique.ID(ctx, ids)], sub)
	} else {
		s.bcastSubs = append(s.bcastSubs, sub)
	}
	s.subscriptionsMu.Unlock()
	select {
	case s.subscriptionsCh <- sub:
	default:
	}
	return sub, nil
}

// DownlinkQueuePush implements io.Server.
func (s *server) DownlinkQueuePush(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, items []*ttnpb.ApplicationDownlink) error {
	s.downlinkQueueMu.Lock()
	uid := unique.ID(ctx, ids)
	s.downlinkQueue[uid] = append(s.downlinkQueue[uid], io.CleanDownlinks(items)...)
	s.downlinkQueueMu.Unlock()
	return nil
}

// DownlinkQueueReplace implements io.Server.
func (s *server) DownlinkQueueReplace(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, items []*ttnpb.ApplicationDownlink) error {
	s.downlinkQueueMu.Lock()
	s.downlinkQueue[unique.ID(ctx, ids)] = io.CleanDownlinks(items)
	s.downlinkQueueMu.Unlock()
	return nil
}

// DownlinkQueueList implements io.Server.
func (s *server) DownlinkQueueList(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) ([]*ttnpb.ApplicationDownlink, error) {
	s.downlinkQueueMu.RLock()
	defer s.downlinkQueueMu.RUnlock()
	return s.downlinkQueue[unique.ID(ctx, ids)], nil
}

func (s *server) SetSubscribeError(err error) {
	s.subscriptionsMu.Lock()
	defer s.subscriptionsMu.Unlock()
	s.subscribeError = err
}

func (s *server) Subscriptions() <-chan *io.Subscription {
	return s.subscriptionsCh
}

func (s *server) RateLimiter() ratelimit.Interface {
	return &ratelimit.NoopRateLimiter{}
}

func (s *server) RangeUplinks(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, paths []string, f func(ctx context.Context, up *ttnpb.ApplicationUplink) bool) error {
	panic("unimplemented")
}

func (s *server) GetEndDevice(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, paths []string) (*ttnpb.EndDevice, error) {
	panic("unimplemented")
}

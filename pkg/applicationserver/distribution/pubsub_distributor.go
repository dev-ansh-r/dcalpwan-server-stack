package distribution

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// NewPubSubDistributor creates a Distributor on top of the provided PubSub.
// The underlying subscription sets can timeout if there are no active subscribers.
// A timeout of 0 means the underlying subscription sets never timeout.
func NewPubSubDistributor(ctx context.Context, rd RequestDecoupler, timeout time.Duration, pubsub PubSub, mapOpts []io.SubscriptionOption) Distributor {
	return &pubSubDistributor{
		pubsub:        pubsub,
		subscriptions: newSubscriptionMap(ctx, rd, timeout, subscribeSetToPubSub(pubsub), mapOpts...),
	}
}

type pubSubDistributor struct {
	pubsub        PubSub
	subscriptions *subscriptionMap
}

// Publish publishes traffic to the underlying Pub/Sub.
func (d *pubSubDistributor) Publish(ctx context.Context, up *ttnpb.ApplicationUp) error {
	return d.pubsub.Publish(ctx, up)
}

var errMissingIdentifiers = errors.DefineFailedPrecondition("missing_identifiers", "subscriptions without identifiers are not supported")

// Subscribe creates a subscription in the associated subscription set.
func (d *pubSubDistributor) Subscribe(ctx context.Context, protocol string, ids *ttnpb.ApplicationIdentifiers) (*io.Subscription, error) {
	if ids == nil {
		return nil, errMissingIdentifiers.New()
	}
	s, err := d.subscriptions.LoadOrCreate(ctx, ids)
	if err != nil {
		return nil, err
	}
	return s.Subscribe(ctx, protocol, ids)
}

func subscribeSetToPubSub(pubsub PubSub) func(*subscriptionSet, *ttnpb.ApplicationIdentifiers) error {
	return func(set *subscriptionSet, ids *ttnpb.ApplicationIdentifiers) error {
		go func() {
			ctx := set.Context()
			if err := pubsub.Subscribe(ctx, ids, set.Publish); err != nil {
				log.FromContext(ctx).WithError(err).Warn("Pub/Sub subscription failed")
				set.Cancel(err)
			}
		}()
		return nil
	}
}

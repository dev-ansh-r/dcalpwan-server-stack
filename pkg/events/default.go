package events

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type noopPubSub struct{}

func (noopPubSub) Publish(...Event) {}

func (noopPubSub) Subscribe(context.Context, []string, []*ttnpb.EntityIdentifiers, Handler) error {
	return nil
}

var defaultPubSub PubSub = &noopPubSub{}

// SetDefaultPubSub sets pubsub used by the package to ps.
func SetDefaultPubSub(ps PubSub) {
	defaultPubSub = ps
}

// DefaultPubSub returns the default PubSub.
func DefaultPubSub() PubSub {
	return defaultPubSub
}

// Subscribe subscribes on the default PubSub.
func Subscribe(ctx context.Context, names []string, ids []*ttnpb.EntityIdentifiers, hdl Handler) error {
	return defaultPubSub.Subscribe(ctx, names, ids, hdl)
}

// Publish emits events on the default event pubsub.
func Publish(evs ...Event) {
	if len(evs) == 0 {
		return
	}
	evs = append(evs[:0:0], evs...)
	for i, evt := range evs {
		evs[i] = local(evt).withCaller()
		publishes.WithLabelValues(evt.Name()).Inc()
	}
	defaultPubSub.Publish(evs...)
}

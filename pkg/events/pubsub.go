package events

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// PubSub interface combines the Publisher and Subscriber interfaces.
type PubSub interface {
	Publisher
	Subscriber
}

// Publisher interface lets you publish events.
type Publisher interface {
	// Publish emits an event on the default event pubsub.
	Publish(evs ...Event)
}

// Subscriber interface lets you subscribe to events.
type Subscriber interface {
	// Subscribe to events that match the names and identifiers.
	// The subscription continues until the context is canceled.
	// Events matching _any_ of the names or identifiers will get sent to the handler.
	// The handler must be non-blocking.
	Subscribe(ctx context.Context, names []string, identifiers []*ttnpb.EntityIdentifiers, hdl Handler) error
}

// Subscription is the interface for PubSub subscriptions.
type Subscription interface {
	// Match returns whether the event matches the subscription.
	Match(Event) bool
	// Notify notifies the subscription of a new matching event.
	Notify(Event)
}

// PublishFunc is a function which implements Publisher.
type PublishFunc func(...Event)

// Publish implements events.Publisher.
func (f PublishFunc) Publish(evs ...Event) { f(evs...) }

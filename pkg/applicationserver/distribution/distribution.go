package distribution

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// PubSub sends upstream traffic from publishers to subscribers.
type PubSub interface {
	// Publish publishes traffic to the subscribers.
	Publish(context.Context, *ttnpb.ApplicationUp) error
	// Subscribe to the traffic of a specific application.
	Subscribe(context.Context, *ttnpb.ApplicationIdentifiers, func(context.Context, *ttnpb.ApplicationUp) error) error
}

// Distributor sends upstream traffic from publishers to subscribers.
type Distributor interface {
	// Publish publishes traffic to the subscribers.
	Publish(context.Context, *ttnpb.ApplicationUp) error
	// Subscribe to the traffic of a specific application.
	Subscribe(context.Context, string, *ttnpb.ApplicationIdentifiers) (*io.Subscription, error)
}

// RequestDecoupler decouples the security information found in a context
// from the lifetime of the context.
type RequestDecoupler interface {
	FromRequestContext(ctx context.Context) context.Context
}

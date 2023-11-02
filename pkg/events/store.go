package events

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Store extends PubSub implementations with storage of historical events.
type Store interface {
	PubSub
	// FindRelated finds events with matching correlation IDs.
	FindRelated(ctx context.Context, correlationID string) ([]Event, error)
	// FetchHistory fetches the tail (optional) of historical events matching the
	// given names (optional) and identifiers (mandatory) after the given time (optional).
	FetchHistory(
		ctx context.Context, names []string, ids []*ttnpb.EntityIdentifiers, after *time.Time, tail int,
	) ([]Event, error)
	// SubscribeWithHistory is like FetchHistory, but after fetching historical events,
	// this continues sending live events until the context is done.
	SubscribeWithHistory(
		ctx context.Context, names []string, ids []*ttnpb.EntityIdentifiers, after *time.Time, tail int, hdl Handler,
	) error
}

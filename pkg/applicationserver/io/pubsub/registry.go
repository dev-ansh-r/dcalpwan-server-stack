package pubsub

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Registry is a registry for pub/sub integrations.
type Registry interface {
	// Get returns the pub/sub integration by its identifiers.
	Get(ctx context.Context, ids *ttnpb.ApplicationPubSubIdentifiers, paths []string) (*ttnpb.ApplicationPubSub, error)
	// Range ranges over the pub/sub integrations and calls the callback function, until false is returned.
	Range(ctx context.Context, paths []string, f func(context.Context, *ttnpb.ApplicationIdentifiers, *ttnpb.ApplicationPubSub) bool) error
	// List returns all pub/sub integrations of the application.
	List(ctx context.Context, ids *ttnpb.ApplicationIdentifiers, paths []string) ([]*ttnpb.ApplicationPubSub, error)
	// Set creates, updates or deletes the pub/sub integration by its identifiers.
	Set(ctx context.Context, ids *ttnpb.ApplicationPubSubIdentifiers, paths []string, f func(*ttnpb.ApplicationPubSub) (*ttnpb.ApplicationPubSub, []string, error)) (*ttnpb.ApplicationPubSub, error)
}

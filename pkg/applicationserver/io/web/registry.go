package web

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// WebhookRegistry is a store for webhooks.
type WebhookRegistry interface {
	// Get returns the webhook by its identifiers.
	Get(ctx context.Context, ids *ttnpb.ApplicationWebhookIdentifiers, paths []string) (*ttnpb.ApplicationWebhook, error)
	// List returns all webhooks of the application.
	List(ctx context.Context, ids *ttnpb.ApplicationIdentifiers, paths []string) ([]*ttnpb.ApplicationWebhook, error)
	// Set creates, updates or deletes the webhook by its identifiers.
	Set(ctx context.Context, ids *ttnpb.ApplicationWebhookIdentifiers, paths []string, f func(*ttnpb.ApplicationWebhook) (*ttnpb.ApplicationWebhook, []string, error)) (*ttnpb.ApplicationWebhook, error)
	// Range ranges over the webhooks and calls the callback function, until false is returned.
	Range(ctx context.Context, paths []string, f func(context.Context, *ttnpb.ApplicationIdentifiers, *ttnpb.ApplicationWebhook) bool) error
}

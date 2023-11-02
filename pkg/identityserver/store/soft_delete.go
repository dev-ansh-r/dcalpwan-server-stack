package store

import (
	"context"
	"time"
)

type deletedOptionsKeyType struct{}

var deletedOptionsKey deletedOptionsKeyType

// DeletedOptions stores the options for selecting deleted entities.
type DeletedOptions struct {
	IncludeDeleted bool
	OnlyDeleted    bool
	DeletedBefore  *time.Time
	DeletedAfter   *time.Time
}

// WithSoftDeleted returns a context that tells the store to include (only) deleted entities.
func WithSoftDeleted(ctx context.Context, onlyDeleted bool) context.Context {
	return context.WithValue(ctx, deletedOptionsKey, &DeletedOptions{
		IncludeDeleted: true,
		OnlyDeleted:    onlyDeleted,
	})
}

// WithSoftDeletedBetween returns a context that tells the store to include deleted entities
// between (exclusive) the given times.
func WithSoftDeletedBetween(ctx context.Context, deletedAfter, deletedBefore *time.Time) context.Context {
	return context.WithValue(ctx, deletedOptionsKey, &DeletedOptions{
		IncludeDeleted: true,
		OnlyDeleted:    deletedBefore != nil || deletedAfter != nil,
		DeletedBefore:  deletedBefore,
		DeletedAfter:   deletedAfter,
	})
}

// SoftDeletedFromContext returns the DeletedOptions from the context if present.
func SoftDeletedFromContext(ctx context.Context) *DeletedOptions {
	if opts, ok := ctx.Value(deletedOptionsKey).(*DeletedOptions); ok {
		return opts
	}
	return nil
}

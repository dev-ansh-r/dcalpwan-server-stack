// Package blocklist implements a list of forbidden IDs.
package blocklist

import (
	"context"
	"sync"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var errForbiddenID = errors.DefineInvalidArgument("forbidden_id", "the ID `{id}` can not be registered")

// New returns a new Blocklist for the given IDs.
func New(ids ...string) *Blocklist {
	b := &Blocklist{
		forbidden: make(map[string]struct{}),
	}
	b.Add(ids...)
	return b
}

// Blocklist is a list of forbidden IDs.
type Blocklist struct {
	mu        sync.RWMutex
	forbidden map[string]struct{}
}

// Add an ID to the blocklist.
func (b *Blocklist) Add(ids ...string) {
	b.mu.Lock()
	for _, id := range ids {
		b.forbidden[id] = struct{}{}
	}
	b.mu.Unlock()
}

// Contains returns whether the blocklist contains the given ID.
func (b *Blocklist) Contains(id string) bool {
	b.mu.RLock()
	_, found := b.forbidden[id]
	b.mu.RUnlock()
	return found
}

// Blocklists contains multiple blocklists.
type Blocklists []*Blocklist

// Contains returns whether any of the blocklists contains the given ID.
func (b Blocklists) Contains(id string) bool {
	for _, b := range b {
		if b.Contains(id) {
			return true
		}
	}
	return false
}

// Check the given ID on the builtin blocklist as well as the blocklists that may
// be in the context.
func Check(ctx context.Context, id string) error {
	if builtin.Contains(id) {
		return errForbiddenID.WithAttributes("id", id)
	}
	if FromContext(ctx).Contains(id) {
		return errForbiddenID.WithAttributes("id", id)
	}
	return nil
}

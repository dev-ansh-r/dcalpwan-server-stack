package experimental

import (
	"context"
	"sync"
)

// Registry is a registry of enabled experimental features.
type Registry struct {
	mu       sync.RWMutex
	features map[string]bool
}

// NewRegistry returns a new feature registry with the given features enabled.
func NewRegistry(enabledFeatures ...string) *Registry {
	r := &Registry{
		features: make(map[string]bool),
	}
	for _, enabledFeature := range enabledFeatures {
		r.features[enabledFeature] = true
	}
	return r
}

// EnableFeatures enables the given features.
func (r *Registry) EnableFeatures(features ...string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, feature := range features {
		r.features[feature] = true
	}
}

// DisableFeatures disables the given features.
func (r *Registry) DisableFeatures(features ...string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, feature := range features {
		r.features[feature] = false
	}
}

func (r *Registry) getFeature(feature string) (value, ok bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, ok = r.features[feature]
	return
}

func (r *Registry) allFeatures() map[string]bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	features := make(map[string]bool, len(r.features))
	for k, v := range r.features {
		features[k] = v
	}
	return features
}

type registryContextKeyType struct{}

var registryContextKey registryContextKeyType

// NewContextWithRegistry returns a new context derived from the parent, that contains the given feature registry.
func NewContextWithRegistry(parent context.Context, r *Registry) context.Context {
	return context.WithValue(parent, registryContextKey, r)
}

func registryFromContext(ctx context.Context) *Registry {
	r, ok := ctx.Value(registryContextKey).(*Registry)
	if !ok {
		return nil
	}
	return r
}

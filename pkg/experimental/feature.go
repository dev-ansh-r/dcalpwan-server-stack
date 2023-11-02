package experimental

import (
	"context"
	"fmt"
	"sync"
)

// Feature is an experimental feature that can be enabled or disabled.
type Feature struct {
	name         string
	defaultValue bool
}

var (
	definedFeatures   = make(map[string]*Feature)
	definedFeaturesMu sync.RWMutex
)

// DefineFeature defines an experimental feature.
func DefineFeature(name string, defaultValue bool) *Feature {
	definedFeaturesMu.Lock()
	defer definedFeaturesMu.Unlock()
	if _, exists := definedFeatures[name]; exists {
		panic(fmt.Errorf("experimental feature %q already defined", name))
	}
	f := &Feature{name: name, defaultValue: defaultValue}
	definedFeatures[name] = f
	return f
}

// GetValue gets the value of the feature flag.
// The value comes from the registry in the context, the global registry,
// or the default value.
func (f *Feature) GetValue(ctx context.Context) bool {
	r := registryFromContext(ctx)
	if r != nil {
		if v, ok := r.getFeature(f.name); ok {
			return v
		}
	}
	if v, ok := globalRegistry.getFeature(f.name); ok {
		return v
	}
	return f.defaultValue
}

// AllFeatures returns all features and their values.
// The values come from the registry in the context, the global registry,
// or the default value.
func AllFeatures(ctx context.Context) map[string]bool {
	definedFeaturesMu.RLock()
	defer definedFeaturesMu.RUnlock()

	features := make(map[string]bool, len(definedFeatures))
	for k, v := range definedFeatures {
		features[k] = v.defaultValue
	}

	globalFeatures := globalRegistry.allFeatures()
	for k, v := range globalFeatures {
		if _, isDefined := features[k]; isDefined {
			features[k] = v
		}
	}

	if r := registryFromContext(ctx); r != nil {
		contextFeatures := r.allFeatures()
		for k, v := range contextFeatures {
			if _, isDefined := features[k]; isDefined {
				features[k] = v
			}
		}
	}

	return features
}

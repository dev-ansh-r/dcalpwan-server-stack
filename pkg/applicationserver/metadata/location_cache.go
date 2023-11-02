package metadata

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// EndDeviceLocationCache is a cache for end device locations.
type EndDeviceLocationCache interface {
	// Get retrieves the end device locations and the remaining TTL for the entry.
	Get(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) (map[string]*ttnpb.Location, *time.Time, error)
	// Set sets the end device locations.
	Set(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, update map[string]*ttnpb.Location, ttl time.Duration) error
	// Delete removes the locations from the cache.
	Delete(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) error
}

type metricsEndDeviceLocationCache struct {
	inner EndDeviceLocationCache
}

// Get implements EndDeviceLocationCache.
func (c *metricsEndDeviceLocationCache) Get(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) (map[string]*ttnpb.Location, *time.Time, error) {
	m, storedAt, err := c.inner.Get(ctx, ids)
	if storedAt == nil {
		registerMetadataCacheMiss(ctx, locationLabel)
	} else {
		registerMetadataCacheHit(ctx, locationLabel)
	}
	return m, storedAt, err
}

// Set implements EndDeviceLocationCache.
func (c *metricsEndDeviceLocationCache) Set(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, update map[string]*ttnpb.Location, ttl time.Duration) error {
	return c.inner.Set(ctx, ids, update, ttl)
}

// Delete implements EndDeviceLocationCache.
func (c *metricsEndDeviceLocationCache) Delete(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) error {
	return c.inner.Delete(ctx, ids)
}

// NewMetricsEndDeviceLocationCache constructs an EndDeviceLocationCache that collects metrics.
func NewMetricsEndDeviceLocationCache(inner EndDeviceLocationCache) EndDeviceLocationCache {
	return &metricsEndDeviceLocationCache{
		inner: inner,
	}
}

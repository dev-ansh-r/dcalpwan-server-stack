package cryptoutil

import (
	"context"
	"fmt"
	"time"

	"github.com/bluele/gcache"
	"go.thethings.network/lorawan-stack/v3/pkg/crypto"
)

type cacheKeyServiceEntry struct {
	value any
	err   error
}

type cacheKeyService struct {
	crypto.KeyService
	cache gcache.Cache
}

// NewCacheKeyService returns a new crypto.KeyService that caches the results of Unwrap.
func NewCacheKeyService(inner crypto.KeyService, ttl time.Duration, size int) crypto.KeyService {
	builder := gcache.New(size).ARC()
	if ttl != 0 {
		builder = builder.Expiration(ttl)
	}
	return &cacheKeyService{
		KeyService: inner,
		cache:      builder.Build(),
	}
}

func (c *cacheKeyService) getOrLoad(
	ctx context.Context, cache crypto.CacheKey, key string, loaderFunc func() (any, error),
) (any, error) {
	cacheKey := fmt.Sprintf("%s:%s", cache, key)
	if val, err := c.cache.Get(cacheKey); err == nil {
		crypto.RegisterCacheHit(ctx, cache)
		entry := val.(*cacheKeyServiceEntry)
		return entry.value, entry.err
	}
	crypto.RegisterCacheMiss(ctx, cache)
	val, err := loaderFunc()
	c.cache.Set(cacheKey, &cacheKeyServiceEntry{val, err}) //nolint:errcheck
	return val, err
}

func (c *cacheKeyService) Unwrap(ctx context.Context, ciphertext []byte, kekLabel string) ([]byte, error) {
	res, err := c.getOrLoad(ctx, crypto.CacheUnwrap, fmt.Sprintf("%s:%X", kekLabel, ciphertext),
		func() (any, error) {
			return c.KeyService.Unwrap(ctx, ciphertext, kekLabel)
		},
	)
	return res.([]byte), err
}

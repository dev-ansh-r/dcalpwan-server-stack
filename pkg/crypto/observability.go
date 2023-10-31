package crypto

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

type cryptoMetrics struct {
	cacheHit  *metrics.ContextualCounterVec
	cacheMiss *metrics.ContextualCounterVec
}

func (m cryptoMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.cacheHit.Describe(ch)
	m.cacheMiss.Describe(ch)
}

func (m cryptoMetrics) Collect(ch chan<- prometheus.Metric) {
	m.cacheHit.Collect(ch)
	m.cacheMiss.Collect(ch)
}

const (
	subsystem = "crypto"
)

var cMetrics = &cryptoMetrics{
	cacheHit: metrics.NewContextualCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "cache_hit",
			Help:      "Number of cache hits",
		},
		[]string{"cache"},
	),
	cacheMiss: metrics.NewContextualCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "cache_miss",
			Help:      "Number of cache misses",
		},
		[]string{"cache"},
	),
}

func init() {
	metrics.MustRegister(cMetrics)
}

// CacheKey is a cache key.
type CacheKey string

// Cache keys.
const (
	CacheEncryptionKey     CacheKey = "key"
	CacheServerCertificate CacheKey = "server_certificate"
	CacheClientCertificate CacheKey = "client_certificate"
	CacheUnwrap            CacheKey = "unwrap"
)

// RegisterCacheHit registers a cache hit for the provided cache.
func RegisterCacheHit(ctx context.Context, cache CacheKey) {
	cMetrics.cacheHit.WithLabelValues(ctx, string(cache)).Inc()
}

// RegisterCacheMiss registers a cache miss for the provided cache.
func RegisterCacheMiss(ctx context.Context, cache CacheKey) {
	cMetrics.cacheMiss.WithLabelValues(ctx, string(cache)).Inc()
}

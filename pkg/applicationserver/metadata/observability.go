package metadata

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

const (
	subsystem     = "as_metadata"
	metadataLabel = "metadata"
	locationLabel = "location"
)

var metaMetrics = &metadataMetrics{
	cacheHits: metrics.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "cache_hits_total",
			Help:      "Total number of metadata cache hits",
		},
		[]string{metadataLabel},
	),
	cacheMisses: metrics.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "cache_misses_total",
			Help:      "Total number of metadata cache misses",
		},
		[]string{metadataLabel},
	),
	registryRetrievals: metrics.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "registry_retrievals_total",
			Help:      "Total number of metadata registry retrievals",
		},
		[]string{metadataLabel},
	),
	registryUpdates: metrics.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "registry_updates_total",
			Help:      "Total number of metadata registry updates",
		},
		[]string{metadataLabel},
	),
}

func init() {
	metrics.MustRegister(metaMetrics)
}

type metadataMetrics struct {
	cacheHits          *prometheus.CounterVec
	cacheMisses        *prometheus.CounterVec
	registryRetrievals *prometheus.CounterVec
	registryUpdates    *prometheus.CounterVec
}

func (m metadataMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.cacheHits.Describe(ch)
	m.cacheMisses.Describe(ch)
	m.registryRetrievals.Describe(ch)
	m.registryUpdates.Describe(ch)
}

func (m metadataMetrics) Collect(ch chan<- prometheus.Metric) {
	m.cacheHits.Collect(ch)
	m.cacheMisses.Collect(ch)
	m.registryRetrievals.Collect(ch)
	m.registryUpdates.Collect(ch)
}

func registerMetadataCacheHit(ctx context.Context, metadata string) {
	metaMetrics.cacheHits.WithLabelValues(metadata).Inc()
	metaMetrics.cacheMisses.WithLabelValues(metadata)
}

func registerMetadataCacheMiss(ctx context.Context, metadata string) {
	metaMetrics.cacheHits.WithLabelValues(metadata)
	metaMetrics.cacheMisses.WithLabelValues(metadata).Inc()
}

func registerMetadataRegistryRetrieval(ctx context.Context, metadata string) {
	metaMetrics.registryRetrievals.WithLabelValues(metadata).Inc()
}

func registerMetadataRegistryUpdate(ctx context.Context, metadata string) {
	metaMetrics.registryUpdates.WithLabelValues(metadata).Inc()
}

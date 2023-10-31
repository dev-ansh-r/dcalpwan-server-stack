package batch

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

const subsystem = "events_batch"

var (
	batchDuration = metrics.NewHistogram(
		prometheus.HistogramOpts{
			Subsystem: subsystem,
			Name:      "batch_duration_seconds",
			Help:      "Histogram of batch durations (seconds)",
			Buckets:   []float64{0.0128, 0.0256, 0.0512, 0.1024, 0.2048, 0.4096, 0.8192, 1.6384},
		},
	)
	batchSize = metrics.NewHistogram(
		prometheus.HistogramOpts{
			Subsystem: subsystem,
			Name:      "batch_size",
			Help:      "Histogram of batch sizes",
			Buckets:   []float64{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384},
		},
	)
)

func registerBatchFlush(d time.Duration, n int) {
	batchDuration.Observe(d.Seconds())
	batchSize.Observe(float64(n))
}

func init() {
	metrics.MustRegister(batchDuration, batchSize)
}

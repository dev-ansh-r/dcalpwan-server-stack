package fetch

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

const subsystem = "fetch"

var fetchLatency = metrics.NewHistogramVec(
	prometheus.HistogramOpts{
		Subsystem: subsystem,
		Name:      "fetch_latency_seconds",
		Help:      "Histogram of latency (seconds) of file fetches",
	},
	[]string{"backend", "base"},
)

func init() {
	metrics.MustRegister(fetchLatency)
}

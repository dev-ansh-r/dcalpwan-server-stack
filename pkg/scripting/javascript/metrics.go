package javascript

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

const subsystem = "javascript"

var (
	compilations = metrics.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "compilations_total",
			Help:      "JavaScript compilations",
		},
		[]string{"result"},
	)
	compilationsLatency = metrics.NewHistogram(
		prometheus.HistogramOpts{
			Subsystem: subsystem,
			Name:      "compilations_latency_seconds",
			Help:      "Histogram of latency (seconds) of JavaScript compilations",
		},
	)
	runs = metrics.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "runs_total",
			Help:      "JavaScript runs",
		},
		[]string{"result"},
	)
	runLatency = metrics.NewHistogram(
		prometheus.HistogramOpts{
			Subsystem: subsystem,
			Name:      "run_latency_seconds",
			Help:      "Histogram of latency (seconds) of JavaScript runs",
		},
	)
)

func init() {
	metrics.MustRegister(
		compilations,
		compilationsLatency,
		runs,
		runLatency,
	)
}

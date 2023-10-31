// Package observability implements a pkg/log.Handler that exports metrics for the logged messages.
package observability

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

type logMessageMetrics struct {
	logMessages *metrics.ContextualCounterVec
}

const (
	subsystem = "log"
	level     = "level"
	namespace = "namespace"
	errorName = "error_name"
)

func (m logMessageMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.logMessages.Describe(ch)
}

func (m logMessageMetrics) Collect(ch chan<- prometheus.Metric) {
	m.logMessages.Collect(ch)
}

var logMetrics = &logMessageMetrics{
	logMessages: metrics.NewContextualCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "messages_total",
			Help:      "Total number of logged messages",
		},
		[]string{level, namespace, errorName},
	),
}

func init() {
	metrics.MustRegister(logMetrics)
}

// observability is a log.Handler that tracks metrics for logged messages.
type observability struct{}

// New creates a new observability log middleware.
func New() log.Middleware {
	return &observability{}
}

// Wrap an existing log handler with observability.
func (o *observability) Wrap(next log.Handler) log.Handler {
	return log.HandlerFunc(func(entry log.Entry) error {
		namespace := "unknown"
		if ns, ok := entry.Fields().Fields()["namespace"]; ok {
			if ns, ok := ns.(string); ok {
				namespace = ns
			}
		}
		errorName := "nil"
		if err, ok := entry.Fields().Fields()["error"].(error); ok {
			errorName = "unknown"
			if ttnErr, ok := errors.From(err); ok {
				errorName = ttnErr.FullName()
			}
		}
		logMetrics.logMessages.WithLabelValues(context.Background(), entry.Level().String(), namespace, errorName).Inc()
		return next.HandleLog(entry)
	})
}

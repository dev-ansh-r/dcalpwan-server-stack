package packages

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

const (
	subsystem = "as_packages"
	unknown   = "unknown"
)

var packagesMetrics = &messageMetrics{
	messagesProcessed: metrics.NewContextualCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "processed_total",
			Help:      "Total number of processed messages",
		},
		[]string{"package"},
	),
	messagesFailed: metrics.NewContextualCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "failed_total",
			Help:      "Total number of failed messages",
		},
		[]string{"package", "error"},
	),
}

func init() {
	metrics.MustRegister(packagesMetrics)
}

type messageMetrics struct {
	messagesProcessed *metrics.ContextualCounterVec
	messagesFailed    *metrics.ContextualCounterVec
}

func (m messageMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.messagesProcessed.Describe(ch)
	m.messagesFailed.Describe(ch)
}

func (m messageMetrics) Collect(ch chan<- prometheus.Metric) {
	m.messagesProcessed.Collect(ch)
	m.messagesFailed.Collect(ch)
}

func registerMessageProcessed(ctx context.Context, name string) {
	packagesMetrics.messagesProcessed.WithLabelValues(ctx, name).Inc()
}

func registerMessageFailed(ctx context.Context, name string, err error) {
	errorLabel := unknown
	if ttnErr, ok := errors.From(err); ok {
		errorLabel = ttnErr.FullName()
	}
	packagesMetrics.messagesFailed.WithLabelValues(ctx, name, errorLabel).Inc()
}

package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// NewCounter returns a new Counter and sets its namespace.
func NewCounter(opts prometheus.CounterOpts) prometheus.Counter {
	opts.Namespace = Namespace
	return prometheus.NewCounter(opts)
}

// MustRegisterCounter is a convenience function for NewCounter and MustRegister.
func MustRegisterCounter(opts prometheus.CounterOpts) prometheus.Counter {
	metric := NewCounter(opts)
	MustRegister(metric)
	return metric
}

// NewCounterFunc returns a new CounterFunc and sets its namespace.
func NewCounterFunc(opts prometheus.CounterOpts, function func() float64) prometheus.CounterFunc {
	opts.Namespace = Namespace
	return prometheus.NewCounterFunc(opts, function)
}

// MustRegisterCounterFunc is a convenience function for NewCounterFunc and MustRegister.
func MustRegisterCounterFunc(opts prometheus.CounterOpts, function func() float64) prometheus.CounterFunc {
	metric := NewCounterFunc(opts, function)
	MustRegister(metric)
	return metric
}

// NewCounterVec returns a new CounterVec and sets its namespace.
func NewCounterVec(opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	opts.Namespace = Namespace
	return prometheus.NewCounterVec(opts, labelNames)
}

// MustRegisterCounterVec is a convenience function for NewCounterVec and MustRegister.
func MustRegisterCounterVec(opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	metric := NewCounterVec(opts, labelNames)
	MustRegister(metric)
	return metric
}

// ContextualCounterVec wraps a CounterVec in order to get labels from the context.
type ContextualCounterVec struct {
	*prometheus.CounterVec
}

// With is the equivalent of CounterVec.With, but with a context.
func (c ContextualCounterVec) With(ctx context.Context, labels prometheus.Labels) prometheus.Counter {
	if LabelsFromContext == nil {
		return c.CounterVec.With(labels)
	}
	return c.CounterVec.MustCurryWith(LabelsFromContext(ctx)).With(labels)
}

// WithLabelValues is the equivalent of CounterVec.WithLabelValues, but with a context.
func (c ContextualCounterVec) WithLabelValues(ctx context.Context, lvs ...string) prometheus.Counter {
	if len(ContextLabelNames) == 0 {
		return c.CounterVec.WithLabelValues(lvs...)
	}
	return c.CounterVec.MustCurryWith(LabelsFromContext(ctx)).WithLabelValues(lvs...)
}

// NewContextualCounterVec returns a new ContextualCounterVec and sets its namespace.
func NewContextualCounterVec(opts prometheus.CounterOpts, labelNames []string) *ContextualCounterVec {
	opts.Namespace = Namespace
	if len(ContextLabelNames) > 0 {
		labelNames = append(ContextLabelNames, labelNames...)
	}
	return &ContextualCounterVec{prometheus.NewCounterVec(opts, labelNames)}
}

// MustRegisterContextualCounterVec is a convenience function for NewContextualCounterVec and MustRegister.
func MustRegisterContextualCounterVec(opts prometheus.CounterOpts, labelNames []string) *ContextualCounterVec {
	metric := NewContextualCounterVec(opts, labelNames)
	MustRegister(metric)
	return metric
}

// Package tracer provides mechanisms to propagate tracer in context.
package tracer

import (
	"context"

	otrace "go.opentelemetry.io/otel/trace"
	"go.thethings.network/lorawan-stack/v3/pkg/telemetry/tracing"
)

type tracerKeyType struct{}

var tracerKey = &tracerKeyType{}

// NewContext returns a derived context with the tracer set.
func NewContext(ctx context.Context, t otrace.Tracer) context.Context {
	return context.WithValue(ctx, tracerKey, t)
}

// NewContextWithTracer returns a derived context with a new tracer set.
func NewContextWithTracer(ctx context.Context, name string, opts ...otrace.TracerOption) context.Context {
	t := tracing.FromContext(ctx).Tracer(name, opts...)
	return NewContext(ctx, t)
}

// FromContext returns the tracer that is attatched to the context
// or returns a new anonymous tracer if it does not exist.
func FromContext(ctx context.Context) otrace.Tracer {
	v := ctx.Value(tracerKey)
	if v == nil {
		return tracing.FromContext(ctx).Tracer("")
	}
	return v.(otrace.Tracer)
}

// StartFromContext returns a derived context with a new span started from the tracer in context.
func StartFromContext(ctx context.Context, name string, opts ...otrace.SpanStartOption) (context.Context, otrace.Span) {
	return FromContext(ctx).Start(ctx, name, opts...)
}

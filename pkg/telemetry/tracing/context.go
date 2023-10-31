package tracing

import (
	"context"

	otrace "go.opentelemetry.io/otel/trace"
)

type tracerProviderKeyType struct{}

var tracerProviderKey = &tracerProviderKeyType{}

// NewContextWithTracerProvider returns a derived context with the tracer provider set.
func NewContextWithTracerProvider(ctx context.Context, tp otrace.TracerProvider) context.Context {
	return context.WithValue(ctx, tracerProviderKey, tp)
}

// FromContext returns the tracer provider that is attached to the context
// or returns a noop tracer provider if it does not exist.
func FromContext(ctx context.Context) otrace.TracerProvider {
	v := ctx.Value(tracerProviderKey)
	if v == nil {
		return otrace.NewNoopTracerProvider()
	}
	return v.(otrace.TracerProvider)
}

// Package tracing provides tools for working with tracing.
package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	otrace "go.opentelemetry.io/otel/trace"
	"go.thethings.network/lorawan-stack/v3/pkg/version"
)

func initResource(ctx context.Context) (*resource.Resource, error) {
	rsc, err := resource.New(ctx,
		resource.WithContainer(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("ttn-lw-stack"),
			semconv.ServiceVersionKey.String(version.String()),
		),
	)
	if err != nil {
		return nil, err
	}

	// Fill empty values with defaults
	return resource.Merge(resource.Default(), rsc)
}

// Initialize initializes the tracing package and returns the tracer provider.
// If tracing is not enabled it returns a noop tracer provider instead.
func Initialize(ctx context.Context, config *Config) (otrace.TracerProvider, func(context.Context) error, error) {
	if !config.Enable {
		return otrace.NewNoopTracerProvider(), func(_ context.Context) error { return nil }, nil
	}

	exp, err := exporterFromConfig(ctx, config)
	if err != nil {
		return nil, nil, err
	}
	bsp := sdktrace.NewBatchSpanProcessor(exp)

	rsc, err := initResource(ctx)
	if err != nil {
		return nil, nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(rsc),
		sdktrace.WithSampler(sdktrace.ParentBased(
			sdktrace.TraceIDRatioBased(config.SampleProbability),
		)),
	)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	return tp, tp.Shutdown, nil
}

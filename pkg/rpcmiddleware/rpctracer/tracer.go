// Package rpctracer implements a gRPC tracing middleware.
package rpctracer

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/otel/trace"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/hooks"
	"go.thethings.network/lorawan-stack/v3/pkg/telemetry/tracing/tracer"
	"google.golang.org/grpc"
)

// TracerHook is the name of the namespace hook.
const TracerHook = "tracer"

// UnaryTracerHook adds the tracer to the context of the unary call.
func UnaryTracerHook(name string, opts ...trace.TracerOption) hooks.UnaryHandlerMiddleware {
	return func(h grpc.UnaryHandler) grpc.UnaryHandler {
		return func(ctx context.Context, req any) (any, error) {
			ctx = tracer.NewContextWithTracer(ctx, name, opts...)
			return h(ctx, req)
		}
	}
}

// StreamTracerHook adds the tracer to the context of the stream.
func StreamTracerHook(name string, opts ...trace.TracerOption) hooks.StreamHandlerMiddleware {
	return func(h grpc.StreamHandler) grpc.StreamHandler {
		return func(srv any, stream grpc.ServerStream) error {
			wrapped := grpc_middleware.WrapServerStream(stream)
			wrapped.WrappedContext = tracer.NewContextWithTracer(stream.Context(), name, opts...)
			return h(srv, wrapped)
		}
	}
}

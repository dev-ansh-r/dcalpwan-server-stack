package events

import (
	"context"
	"fmt"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const requestIDKey = "x-request-id"

func extractOrGenerateCorrelationID(ctx context.Context, fullMethod string) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	var id string
	if xRequestID := md.Get(requestIDKey); len(xRequestID) > 0 {
		id = xRequestID[len(xRequestID)-1]
	} else {
		id = NewCorrelationID()
	}
	ctx = ContextWithCorrelationID(ctx, fmt.Sprintf("rpc:%s:%s", fullMethod, id))
	return ctx
}

func methodSet(methods []string) map[string]struct{} {
	m := make(map[string]struct{}, len(methods))
	for _, path := range methods {
		m[path] = struct{}{}
	}
	return m
}

// UnaryServerInterceptor returns a new unary server interceptor
// that modifies the context to include a correlation ID.
func UnaryServerInterceptor(ignoredMethods []string) grpc.UnaryServerInterceptor {
	ignored := methodSet(ignoredMethods)
	return func(
		ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (any, error) {
		if _, ok := ignored[info.FullMethod]; ok {
			return handler(ctx, req)
		}
		ctx = extractOrGenerateCorrelationID(ctx, info.FullMethod)
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor
// that that modifies the context.
func StreamServerInterceptor(ignoredMethods []string) grpc.StreamServerInterceptor {
	ignored := methodSet(ignoredMethods)
	return func(
		srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler,
	) error {
		if _, ok := ignored[info.FullMethod]; ok {
			return handler(srv, stream)
		}
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = extractOrGenerateCorrelationID(stream.Context(), info.FullMethod)
		return handler(srv, wrapped)
	}
}

package rpclog

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/hooks"
	"google.golang.org/grpc"
)

// NamespaceHook is the name of the namespace hook.
const NamespaceHook = "namespace"

// UnaryNamespaceHook adds the component namespace to the context of the unary call.
func UnaryNamespaceHook(namespace string) hooks.UnaryHandlerMiddleware {
	return func(h grpc.UnaryHandler) grpc.UnaryHandler {
		return func(ctx context.Context, req any) (any, error) {
			ctx = log.NewContextWithField(ctx, "namespace", namespace)
			return h(ctx, req)
		}
	}
}

// StreamNamespaceHook adds the component namespace to the context of the stream.
func StreamNamespaceHook(namespace string) hooks.StreamHandlerMiddleware {
	return func(h grpc.StreamHandler) grpc.StreamHandler {
		return func(srv any, stream grpc.ServerStream) error {
			wrapped := grpc_middleware.WrapServerStream(stream)
			wrapped.WrappedContext = log.NewContextWithField(stream.Context(), "namespace", namespace)
			return h(srv, wrapped)
		}
	}
}

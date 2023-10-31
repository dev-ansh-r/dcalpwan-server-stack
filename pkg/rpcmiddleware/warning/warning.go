// Package warning communicates warnings over gRPC headers.
// The Add func is used by the server to add a warning header.
// The client interceptors log warnings to the logger in the context, or to the
// default logger.
//
// Note that headers are currently not supported by ServeHTTP of the gRPC server.
// This means that warnings may not be received by clients using the fallback server.
package warning

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const warning = "warning"

func logWarnings(ctx context.Context, md metadata.MD) {
	if warnings := md.Get(warning); len(warnings) > 0 {
		logger := log.FromContext(ctx)
		if logger == log.Noop {
			logHandler, err := log.NewZap("console")
			if err != nil {
				panic(err)
			}
			logger = log.NewLogger(logHandler)
		}
		for _, warning := range warnings {
			logger.Warn(warning)
		}
	}
}

// Add a warning to the response headers.
func Add(ctx context.Context, message string) {
	grpc.SetHeader(ctx, metadata.Pairs(warning, message)) // nolint:gas
}

// UnaryClientInterceptor is a unary client interceptor that logs warnings sent by the server.
func UnaryClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var md metadata.MD
	err := invoker(ctx, method, req, reply, cc, append(opts, grpc.Header(&md))...)
	logWarnings(ctx, md)
	return err
}

// StreamClientInterceptor is a streaming client interceptor that logs warnings sent by the server.
func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var md metadata.MD
	stream, err := streamer(ctx, desc, cc, method, append(opts, grpc.Header(&md))...)
	logWarnings(ctx, md)
	return stream, err
}

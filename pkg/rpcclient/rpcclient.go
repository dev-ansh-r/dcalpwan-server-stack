// Package rpcclient contains the default options for TTN gRPC clients.
package rpcclient

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware"
	_ "go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/discover" // Register service discovery resolvers
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/rpclog"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/warning"
	"go.thethings.network/lorawan-stack/v3/pkg/telemetry/tracing"
	"go.thethings.network/lorawan-stack/v3/pkg/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
)

// DefaultDialOptions for gRPC clients
func DefaultDialOptions(ctx context.Context) []grpc.DialOption {
	streamInterceptors := []grpc.StreamClientInterceptor{
		metrics.StreamClientInterceptor,
		otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(tracing.FromContext(ctx))),
		rpclog.StreamClientInterceptor(ctx), // Gets logger from global context
		warning.StreamClientInterceptor,
		errors.StreamClientInterceptor(),
	}

	unaryInterceptors := []grpc.UnaryClientInterceptor{
		metrics.UnaryClientInterceptor,
		otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(tracing.FromContext(ctx))),
		rpclog.UnaryClientInterceptor(ctx), // Gets logger from global context
		warning.UnaryClientInterceptor,
		errors.UnaryClientInterceptor(),
	}

	return []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(1024*1024*16),
			grpc.MaxCallRecvMsgSize(1024*1024*16),
		),
		grpc.WithStatsHandler(rpcmiddleware.StatsHandlers{new(ocgrpc.ClientHandler), metrics.StatsHandler}),
		grpc.WithUserAgent(fmt.Sprintf(
			"%s go/%s ttn/%s",
			filepath.Base(os.Args[0]),
			strings.TrimPrefix(runtime.Version(), "go"),
			version.String(),
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(streamInterceptors...)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryInterceptors...)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                5 * time.Minute,
			Timeout:             20 * time.Second,
			PermitWithoutStream: false,
		}),
	}
}

func init() {
	resolver.SetDefaultScheme("dns")
}

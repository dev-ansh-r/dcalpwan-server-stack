package rpcmiddleware

import (
	"context"

	"google.golang.org/grpc/stats"
)

// StatsHandlers is a slice of stats.Handler that implements stats.Handler.
// Calls are delegated to all handlers in order.
type StatsHandlers []stats.Handler

// TagRPC implements stats.Handler.
func (s StatsHandlers) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	for _, hdl := range s {
		ctx = hdl.TagRPC(ctx, info)
	}
	return ctx
}

// HandleRPC implements stats.Handler.
func (s StatsHandlers) HandleRPC(ctx context.Context, stats stats.RPCStats) {
	for _, hdl := range s {
		hdl.HandleRPC(ctx, stats)
	}
}

// TagConn implements stats.Handler.
func (s StatsHandlers) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	for _, hdl := range s {
		ctx = hdl.TagConn(ctx, info)
	}
	return ctx
}

// HandleConn implements stats.Handler.
func (s StatsHandlers) HandleConn(ctx context.Context, stats stats.ConnStats) {
	for _, hdl := range s {
		hdl.HandleConn(ctx, stats)
	}
}

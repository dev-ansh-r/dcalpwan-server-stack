package packetbrokeragent_test

import (
	"context"
	"net"
	"testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/packetbrokeragent/mock"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

func mustHavePeer(ctx context.Context, c *component.Component, role ttnpb.ClusterRole) {
	for i := 0; i < 20; i++ {
		time.Sleep(20 * time.Millisecond)
		if _, err := c.GetPeer(ctx, role, nil); err == nil {
			return
		}
	}
	panic("could not connect to peer")
}

func mustServe[S interface {
	Serve(net.Listener) error
	GracefulStop()
}](ctx context.Context, tb testing.TB, create func(testing.TB) S,
) (S, net.Addr) {
	tb.Helper()
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	server := create(tb)
	go server.Serve(lis) //nolint:errcheck
	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()
	return server, lis.Addr()
}

func mustServePBIAM(ctx context.Context, tb testing.TB) (*mock.PBIAM, net.Addr) {
	tb.Helper()
	return mustServe(ctx, tb, mock.NewPBIAM)
}

func mustServePBControlPane(ctx context.Context, tb testing.TB) (*mock.PBControlPlane, net.Addr) {
	tb.Helper()
	return mustServe(ctx, tb, mock.NewPBControlPlane)
}

func mustServePBDataPlane(ctx context.Context, tb testing.TB) (*mock.PBDataPlane, net.Addr) {
	tb.Helper()
	return mustServe(ctx, tb, mock.NewPBDataPlane)
}

func mustServePBMapper(ctx context.Context, tb testing.TB) (*mock.PBMapper, net.Addr) {
	tb.Helper()
	return mustServe(ctx, tb, mock.NewPBMapper)
}

func eui64Ptr(v types.EUI64) *types.EUI64 {
	return &v
}

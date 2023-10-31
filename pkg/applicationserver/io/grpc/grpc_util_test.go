package grpc_test

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

var testRights = []ttnpb.Right{
	ttnpb.Right_RIGHT_APPLICATION_INFO,
	ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ,
	ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE,
	ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_READ,
	ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE,
	ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_UP_WRITE,
}

type mockRegisterer struct {
	context.Context
	ttnpb.AppAsServer
}

func (m *mockRegisterer) Roles() []ttnpb.ClusterRole {
	return nil
}

func (m *mockRegisterer) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterAppAsServer(s, m.AppAsServer)
}

func (m *mockRegisterer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
	ttnpb.RegisterAppAsHandler(m.Context, s, conn)
}

func mustHavePeer(ctx context.Context, c *component.Component, role ttnpb.ClusterRole) {
	for i := 0; i < 20; i++ {
		time.Sleep(20 * time.Millisecond)
		if _, err := c.GetPeer(ctx, role, nil); err == nil {
			return
		}
	}
	panic("could not connect to peer")
}

type mockFetcher struct {
	calledWithIdentifiers *ttnpb.EndDeviceIdentifiers

	ids *ttnpb.EndDeviceIdentifiers
	err error
}

func (f *mockFetcher) Get(_ context.Context, ids *ttnpb.EndDeviceIdentifiers) (*ttnpb.EndDeviceIdentifiers, error) {
	f.calledWithIdentifiers = ids
	return f.ids, f.err
}

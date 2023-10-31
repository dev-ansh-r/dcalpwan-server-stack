package grpc_test

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

var testRights = []ttnpb.Right{ttnpb.Right_RIGHT_GATEWAY_INFO, ttnpb.Right_RIGHT_GATEWAY_LINK}

type mockRegisterer struct {
	context.Context
	ttnpb.GtwGsServer
}

func (m *mockRegisterer) Roles() []ttnpb.ClusterRole {
	return nil
}

func (m *mockRegisterer) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterGtwGsServer(s, m.GtwGsServer)
}

func (m *mockRegisterer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
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

package packages_test

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcserver"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
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

type mockPackageHandler struct {
	HandleUpFunc func(context.Context, *ttnpb.ApplicationPackageDefaultAssociation, *ttnpb.ApplicationPackageAssociation, *ttnpb.ApplicationUp) error
}

func (h *mockPackageHandler) HandleUp(ctx context.Context, defaultAssoc *ttnpb.ApplicationPackageDefaultAssociation, assoc *ttnpb.ApplicationPackageAssociation, up *ttnpb.ApplicationUp) error {
	if h.HandleUpFunc == nil {
		panic("HandleUp called but HandleUpFunc is nil")
	}
	return h.HandleUpFunc(ctx, defaultAssoc, assoc, up)
}

func (h *mockPackageHandler) Package() *ttnpb.ApplicationPackage {
	return &ttnpb.ApplicationPackage{
		Name:         "test-package",
		DefaultFPort: 123,
	}
}

type handleUpRequest struct {
	ctx   context.Context
	assoc *ttnpb.ApplicationPackageAssociation
	up    *ttnpb.ApplicationUp
}

func createMockPackageHandler(ch chan<- *handleUpRequest) *mockPackageHandler {
	return &mockPackageHandler{
		HandleUpFunc: func(ctx context.Context, defaultAssoc *ttnpb.ApplicationPackageDefaultAssociation, assoc *ttnpb.ApplicationPackageAssociation, up *ttnpb.ApplicationUp) error {
			ch <- &handleUpRequest{ctx, assoc, up}
			return nil
		},
	}
}

type grpcServiceRegistererWrapper struct {
	rpcserver.ServiceRegisterer
}

func (*grpcServiceRegistererWrapper) Roles() []ttnpb.ClusterRole {
	return nil
}

func (s *grpcServiceRegistererWrapper) RegisterServices(gs *grpc.Server) {
	s.ServiceRegisterer.RegisterServices(gs)
}

func (s *grpcServiceRegistererWrapper) RegisterHandlers(sm *runtime.ServeMux, conn *grpc.ClientConn) {
	s.ServiceRegisterer.RegisterHandlers(sm, conn)
}

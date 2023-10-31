package component

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

// NewConfigurationServer returns a new ConfigurationServer on top of the given component.
func NewConfigurationServer(c *Component) *ConfigurationServer {
	return &ConfigurationServer{component: c}
}

// ConfigurationServer implements the Configuration RPC service.
type ConfigurationServer struct {
	ttnpb.UnimplementedConfigurationServer

	component *Component
}

// Roles implements the rpcserver.Registerer interface. It just returns nil.
func (*ConfigurationServer) Roles() []ttnpb.ClusterRole { return nil }

// RegisterServices registers the Configuration service.
func (c *ConfigurationServer) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterConfigurationServer(s, c)
}

// RegisterHandlers registers the Configuration service handler.
func (c *ConfigurationServer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
	_ = ttnpb.RegisterConfigurationHandler(c.component.Context(), s, conn)
}

// ListFrequencyPlans implements the Configuration service's ListFrequencyPlans RPC.
func (c *ConfigurationServer) ListFrequencyPlans(
	ctx context.Context, req *ttnpb.ListFrequencyPlansRequest,
) (*ttnpb.ListFrequencyPlansResponse, error) {
	fps, err := c.component.FrequencyPlansStore(ctx)
	if err != nil {
		return nil, err
	}
	return frequencyplans.NewRPCServer(fps).ListFrequencyPlans(ctx, req)
}

// GetPhyVersions implements the Configuration service's GetPhyVersions RPC.
func (*ConfigurationServer) GetPhyVersions(
	ctx context.Context, req *ttnpb.GetPhyVersionsRequest,
) (*ttnpb.GetPhyVersionsResponse, error) {
	return band.GetPhyVersions(ctx, req)
}

// ListBands implements the Configuration service's ListBands RPC.
func (*ConfigurationServer) ListBands(
	ctx context.Context, req *ttnpb.ListBandsRequest,
) (*ttnpb.ListBandsResponse, error) {
	return band.ListBands(ctx, req)
}

package mock

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

// GatewayServer is a mock Gateway Server.
type GatewayServer struct {
	ttnpb.UnimplementedNsGsServer

	*component.Component
	Downlink chan *ttnpb.DownlinkMessage
}

// NewGatewayServer returns a new GatewayServer.
func NewGatewayServer(c *component.Component) (*GatewayServer, error) {
	gs := &GatewayServer{
		Component: c,
		Downlink:  make(chan *ttnpb.DownlinkMessage, 1),
	}
	c.RegisterGRPC(gs)
	return gs, nil
}

// Roles implements rpcserver.Registerer.
func (gs *GatewayServer) Roles() []ttnpb.ClusterRole {
	return []ttnpb.ClusterRole{ttnpb.ClusterRole_GATEWAY_SERVER}
}

// RegisterServices implements rpcserver.Registerer.
func (gs *GatewayServer) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterNsGsServer(s, gs)
}

// RegisterHandlers implements rpcserver.Registerer.
func (gs *GatewayServer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
}

// Publish publishes the given message to Packet Broker Agent in the cluster.
func (gs *GatewayServer) Publish(ctx context.Context, up *ttnpb.GatewayUplinkMessage) error {
	client := ttnpb.NewGsPbaClient(gs.LoopbackConn())
	_, err := client.PublishUplink(ctx, up, gs.WithClusterAuth())
	return err
}

// ScheduleDownlink implements ttnpb.NsGsServer.
func (gs *GatewayServer) ScheduleDownlink(ctx context.Context, req *ttnpb.DownlinkMessage) (*ttnpb.ScheduleDownlinkResponse, error) {
	select {
	case gs.Downlink <- req:
	default:
	}
	return &ttnpb.ScheduleDownlinkResponse{}, nil
}

// UpdateGateway updates the gateway in Packet Broker Agent in the cluster.
func (gs *GatewayServer) UpdateGateway(ctx context.Context, req *ttnpb.UpdatePacketBrokerGatewayRequest) (*ttnpb.UpdatePacketBrokerGatewayResponse, error) {
	client := ttnpb.NewGsPbaClient(gs.LoopbackConn())
	return client.UpdateGateway(ctx, req, gs.WithClusterAuth())
}

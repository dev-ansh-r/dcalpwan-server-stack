// Package deviceclaimingserver implements the Device Claiming Server component.
package deviceclaimingserver

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/deviceclaimingserver/enddevices"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

// DeviceClaimingServer is the Device Claiming Server.
type DeviceClaimingServer struct {
	*component.Component
	ctx context.Context

	config Config

	endDeviceClaimingUpstream *enddevices.Upstream

	gatewayClaimingServerUpstream ttnpb.GatewayClaimingServerServer

	grpc struct {
		endDeviceClaimingServer      *endDeviceClaimingServer
		endDeviceBatchClaimingServer *endDeviceBatchClaimingServer
		gatewayClaimingServer        *gatewayClaimingServer
	}
}

// New returns a new Device Claiming component.
func New(c *component.Component, conf *Config, opts ...Option) (*DeviceClaimingServer, error) {
	ctx := log.NewContextWithField(c.Context(), "namespace", "deviceclaimingserver")

	dcs := &DeviceClaimingServer{
		Component: c,
		ctx:       ctx,
		config:    *conf,
	}
	for _, opt := range opts {
		opt(dcs)
	}

	if dcs.endDeviceClaimingUpstream == nil {
		upstream, err := enddevices.NewUpstream(ctx, c, conf.EndDeviceClaimingServerConfig)
		if err != nil {
			return nil, err
		}
		dcs.endDeviceClaimingUpstream = upstream
	}
	dcs.grpc.endDeviceClaimingServer = &endDeviceClaimingServer{
		DCS: dcs,
	}

	dcs.grpc.endDeviceBatchClaimingServer = &endDeviceBatchClaimingServer{
		DCS: dcs,
	}

	dcs.gatewayClaimingServerUpstream = noopGCLS{}
	dcs.grpc.gatewayClaimingServer = &gatewayClaimingServer{
		DCS: dcs,
	}

	c.RegisterGRPC(dcs)
	return dcs, nil
}

// Option configures GatewayClaimingServer.
type Option func(*DeviceClaimingServer)

// WithEndDeviceClaimingUpstream configures the upstream for end device claiming.
func WithEndDeviceClaimingUpstream(upstream *enddevices.Upstream) Option {
	return func(dcs *DeviceClaimingServer) {
		dcs.endDeviceClaimingUpstream = upstream
	}
}

// Context returns the context of the Device Claiming Server.
func (dcs *DeviceClaimingServer) Context() context.Context {
	return dcs.ctx
}

// Roles returns the roles that the Device Claiming Server fulfills.
func (*DeviceClaimingServer) Roles() []ttnpb.ClusterRole {
	return []ttnpb.ClusterRole{ttnpb.ClusterRole_DEVICE_CLAIMING_SERVER}
}

// RegisterServices registers services provided by dcs at s.
func (dcs *DeviceClaimingServer) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterEndDeviceClaimingServerServer(s, dcs.grpc.endDeviceClaimingServer)
	ttnpb.RegisterEndDeviceBatchClaimingServerServer(s, dcs.grpc.endDeviceBatchClaimingServer)
	ttnpb.RegisterGatewayClaimingServerServer(s, dcs.grpc.gatewayClaimingServer)
}

// RegisterHandlers registers gRPC handlers.
//
//nolint:errcheck
func (dcs *DeviceClaimingServer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
	ttnpb.RegisterEndDeviceClaimingServerHandler(dcs.Context(), s, conn)
	ttnpb.RegisterEndDeviceBatchClaimingServerHandler(dcs.Context(), s, conn)
	ttnpb.RegisterGatewayClaimingServerHandler(dcs.Context(), s, conn)
}

package deviceclaimingserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// noopGCLS is a no-op GCLS.
type noopGCLS struct {
	ttnpb.UnimplementedGatewayClaimingServerServer
}

// Claim implements GatewayClaimingServer.
func (noopGCLS) Claim(
	_ context.Context,
	_ *ttnpb.ClaimGatewayRequest,
) (ids *ttnpb.GatewayIdentifiers, retErr error) {
	return nil, errMethodUnavailable.New()
}

// AuthorizeGateway implements GatewayClaimingServer.
func (noopGCLS) AuthorizeGateway(
	_ context.Context,
	_ *ttnpb.AuthorizeGatewayRequest,
) (*emptypb.Empty, error) {
	return nil, errMethodUnavailable.New()
}

// UnauthorizeGateway implements GatewayClaimingServer.
func (noopGCLS) UnauthorizeGateway(
	_ context.Context,
	_ *ttnpb.GatewayIdentifiers,
) (*emptypb.Empty, error) {
	return nil, errMethodUnavailable.New()
}

// gatewayClaimingServer is the front facing entity for gRPC requests.
type gatewayClaimingServer struct {
	ttnpb.UnimplementedGatewayClaimingServerServer

	DCS *DeviceClaimingServer
}

// Claim implements GatewayClaimingServer.
func (gcls gatewayClaimingServer) Claim(
	ctx context.Context,
	req *ttnpb.ClaimGatewayRequest,
) (ids *ttnpb.GatewayIdentifiers, retErr error) {
	return gcls.DCS.gatewayClaimingServerUpstream.Claim(ctx, req)
}

// AuthorizeGateway implements GatewayClaimingServer.
func (gcls gatewayClaimingServer) AuthorizeGateway(
	ctx context.Context,
	req *ttnpb.AuthorizeGatewayRequest,
) (*emptypb.Empty, error) {
	return gcls.DCS.gatewayClaimingServerUpstream.AuthorizeGateway(ctx, req)
}

// UnauthorizeGateway implements GatewayClaimingServer.
func (gcls gatewayClaimingServer) UnauthorizeGateway(
	ctx context.Context,
	gtwIDs *ttnpb.GatewayIdentifiers,
) (*emptypb.Empty, error) {
	return gcls.DCS.gatewayClaimingServerUpstream.UnauthorizeGateway(ctx, gtwIDs)
}

package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type nsJsServer struct {
	ttnpb.UnimplementedNsJsServer

	JS *JoinServer
}

// HandleJoin is called by the Network Server to join a device.
func (srv nsJsServer) HandleJoin(ctx context.Context, req *ttnpb.JoinRequest) (res *ttnpb.JoinResponse, err error) {
	return srv.JS.HandleJoin(ctx, req, ClusterAuthorizer(ctx))
}

// GetNwkSKeys returns the NwkSKeys associated with session keys identified by the supplied request.
func (srv nsJsServer) GetNwkSKeys(ctx context.Context, req *ttnpb.SessionKeyRequest) (*ttnpb.NwkSKeysResponse, error) {
	return srv.JS.GetNwkSKeys(ctx, req, ClusterAuthorizer(ctx))
}

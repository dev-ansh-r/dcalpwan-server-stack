package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type asJsServer struct {
	ttnpb.UnimplementedAsJsServer

	JS *JoinServer
}

// GetAppSKey returns the AppSKey associated with session keys identified by the supplied request.
func (srv asJsServer) GetAppSKey(ctx context.Context, req *ttnpb.SessionKeyRequest) (*ttnpb.AppSKeyResponse, error) {
	return srv.JS.GetAppSKey(ctx, req, ClusterAuthorizer(ctx))
}

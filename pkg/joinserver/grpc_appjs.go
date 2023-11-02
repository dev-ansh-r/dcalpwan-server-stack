package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type appJsServer struct {
	ttnpb.UnimplementedAppJsServer

	JS *JoinServer
}

// GetAppSKey returns the AppSKey associated with session keys identified by the supplied request.
func (srv appJsServer) GetAppSKey(ctx context.Context, req *ttnpb.SessionKeyRequest) (*ttnpb.AppSKeyResponse, error) {
	return srv.JS.GetAppSKey(ctx, req, ApplicationRightsAuthorizer(ctx))
}

package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type jsServer struct {
	ttnpb.UnimplementedJsServer

	JS *JoinServer
}

// GetJoinEUIPrefixes returns the JoinEUIPrefixes associated with the join server.
func (srv jsServer) GetJoinEUIPrefixes(ctx context.Context, _ *emptypb.Empty) (*ttnpb.JoinEUIPrefixes, error) {
	prefixes := make([]*ttnpb.JoinEUIPrefix, 0, len(srv.JS.euiPrefixes))
	for _, p := range srv.JS.euiPrefixes {
		prefixes = append(prefixes, &ttnpb.JoinEUIPrefix{
			JoinEui: p.EUI64.Bytes(),
			Length:  uint32(p.Length),
		})
	}
	return &ttnpb.JoinEUIPrefixes{
		Prefixes: prefixes,
	}, nil
}

// GetDefaultJoinEUI returns the default JoinEUI that is configured for this Join Server.
func (srv jsServer) GetDefaultJoinEUI(ctx context.Context, _ *emptypb.Empty) (*ttnpb.GetDefaultJoinEUIResponse, error) {
	return &ttnpb.GetDefaultJoinEUIResponse{
		JoinEui: srv.JS.defaultJoinEUI.Bytes(),
	}, nil
}

package cluster

import (
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

func NewTestPeer(conn *grpc.ClientConn) Peer {
	return &peer{
		name:   "name",
		roles:  []ttnpb.ClusterRole{ttnpb.ClusterRole_ACCESS},
		target: "target",
		conn:   conn,
	}
}

package cluster

import (
	"context"
	"encoding/hex"

	clusterauth "go.thethings.network/lorawan-stack/v3/pkg/auth/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
	"google.golang.org/grpc"
)

// HookName is the name of the hook used to verify the identity of incoming calls within a cluster.
var HookName = "cluster-hook"

func (c *cluster) TLS() bool { return c.tls }

func (c *cluster) WithVerifiedSource(ctx context.Context) context.Context {
	return clusterauth.VerifySource(ctx, c.keys)
}

func (c *cluster) Auth() grpc.CallOption {
	md := rpcmetadata.MD{
		ID:            c.self.name,
		AuthType:      clusterauth.AuthType,
		AuthValue:     hex.EncodeToString(c.keys[0]),
		AllowInsecure: !c.tls,
	}
	return grpc.PerRPCCredentials(md)
}

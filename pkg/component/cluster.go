package component

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcclient"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

func (c *Component) initCluster() (err error) {
	clusterOpts := []cluster.Option{
		cluster.WithConn(c.LoopbackConn()),
		cluster.WithDialOptions(rpcclient.DefaultDialOptions),
	}
	for _, svc := range c.grpcSubsystems {
		clusterOpts = append(clusterOpts, cluster.WithServices(svc))
	}
	if tlsConfig, err := c.GetTLSClientConfig(c.Context()); err == nil {
		clusterOpts = append(clusterOpts, cluster.WithTLSConfig(tlsConfig))
	}
	ctx := log.NewContextWithField(c.ctx, "namespace", "cluster")
	c.cluster, err = c.clusterNew(ctx, &c.config.ServiceBase.Cluster, clusterOpts...)
	if err != nil {
		return err
	}
	return nil
}

// ClusterTLS returns whether the cluster uses TLS for cluster connections.
func (c *Component) ClusterTLS() bool {
	return c.cluster.TLS()
}

// GetPeers returns cluster peers with the given role and the given tags.
// See package ../cluster for more information.
func (c *Component) GetPeers(ctx context.Context, role ttnpb.ClusterRole) ([]cluster.Peer, error) {
	return c.cluster.GetPeers(ctx, role)
}

// GetPeer returns a cluster peer with the given role and the given tags.
// See package ../cluster for more information.
func (c *Component) GetPeer(ctx context.Context, role ttnpb.ClusterRole, ids cluster.EntityIdentifiers) (cluster.Peer, error) {
	return c.cluster.GetPeer(ctx, role, ids)
}

// GetPeerConn returns a gRPC client connection to the cluster peer.
// See package ../cluster for more information.
func (c *Component) GetPeerConn(ctx context.Context, role ttnpb.ClusterRole, ids cluster.EntityIdentifiers) (*grpc.ClientConn, error) {
	return c.cluster.GetPeerConn(ctx, role, ids)
}

// ClaimIDs claims the identifiers in the cluster.
// See package ../cluster for more information.
func (c *Component) ClaimIDs(ctx context.Context, ids cluster.EntityIdentifiers) error {
	return c.cluster.ClaimIDs(ctx, ids)
}

// UnclaimIDs unclaims the identifiers in the cluster.
// See package ../cluster for more information.
func (c *Component) UnclaimIDs(ctx context.Context, ids cluster.EntityIdentifiers) error {
	return c.cluster.UnclaimIDs(ctx, ids)
}

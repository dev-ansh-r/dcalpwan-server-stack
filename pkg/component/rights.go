package component

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

func (c *Component) initRights() {
	fetcher := rights.NewAccessFetcher(func(ctx context.Context) *grpc.ClientConn {
		conn, err := c.GetPeerConn(ctx, ttnpb.ClusterRole_ACCESS, nil)
		if err != nil {
			return nil
		}
		return conn
	}, c.config.GRPC.AllowInsecureForCredentials)

	if c.config.Rights.TTL > 0 {
		fetcher = rights.NewInMemoryCache(fetcher, c.config.Rights.TTL, c.config.Rights.TTL)
	} else {
		c.Logger().Warn("No rights TTL configured")
	}

	c.rightsFetcher = fetcher
	c.AddContextFiller(func(ctx context.Context) context.Context {
		return rights.NewContextWithFetcher(ctx, fetcher)
	})
}

package component_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcserver"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestPeers(t *testing.T) {
	a := assertions.New(t)

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	ctx, cancel := context.WithCancel(test.Context())
	defer cancel()

	srv := rpcserver.New(ctx)
	go srv.Serve(lis)
	defer srv.Stop()

	var c *component.Component

	config := &component.Config{
		ServiceBase: config.ServiceBase{Cluster: cluster.Config{
			Name:          "test-cluster",
			NetworkServer: lis.Addr().String(),
			TLS:           false,
		}},
	}

	c, err = component.New(test.GetLogger(t), config)
	a.So(err, should.BeNil)
	err = c.Start()
	a.So(err, should.BeNil)

	unusedRoles := []ttnpb.ClusterRole{
		ttnpb.ClusterRole_APPLICATION_SERVER,
		ttnpb.ClusterRole_GATEWAY_SERVER,
		ttnpb.ClusterRole_JOIN_SERVER,
		ttnpb.ClusterRole_ACCESS,
		ttnpb.ClusterRole_ENTITY_REGISTRY,
	}

	var peer cluster.Peer
	for i := 0; i < 20; i++ {
		time.Sleep(20 * time.Millisecond) // Wait for peers to join cluster.
		peer, err = c.GetPeer(context.Background(), ttnpb.ClusterRole_NETWORK_SERVER, nil)
		if err == nil {
			break
		}
	}

	if !a.So(peer, should.NotBeNil) {
		t.FailNow()
	}

	conn, err := peer.Conn()
	a.So(err, should.BeNil)
	a.So(conn, should.NotBeNil)

	for _, role := range unusedRoles {
		peer, err := c.GetPeer(context.Background(), role, nil)
		a.So(peer, should.BeNil)
		a.So(err, should.NotBeNil)
	}

	peers, err := c.GetPeers(context.Background(), ttnpb.ClusterRole_NETWORK_SERVER)
	a.So(peers, should.HaveLength, 1)
	a.So(err, should.BeNil)

	for _, role := range unusedRoles {
		peers, err = c.GetPeers(context.Background(), role)
		a.So(peers, should.HaveLength, 0)
		a.So(err, should.BeNil)
	}
}

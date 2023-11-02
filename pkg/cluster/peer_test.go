package cluster_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
)

func TestPeer(t *testing.T) {
	a := assertions.New(t)

	conn := new(grpc.ClientConn)

	p := cluster.NewTestPeer(conn)

	a.So(p.HasRole(ttnpb.ClusterRole_APPLICATION_SERVER), should.BeFalse)
	a.So(p.HasRole(ttnpb.ClusterRole_ACCESS), should.BeTrue)

	cc, err := p.Conn()
	a.So(err, should.BeNil)
	a.So(cc, should.Equal, conn)
}

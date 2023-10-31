package test_test

import (
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var (
	_ cluster.Cluster = MockCluster{}
	_ cluster.Peer    = MockPeer{}
)

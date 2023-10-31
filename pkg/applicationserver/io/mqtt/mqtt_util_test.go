package mqtt_test

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var testRights = []ttnpb.Right{
	ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_READ,
	ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE,
}

func mustHavePeer(ctx context.Context, c *component.Component, role ttnpb.ClusterRole) {
	for i := 0; i < 20; i++ {
		time.Sleep(20 * time.Millisecond)
		if _, err := c.GetPeer(ctx, role, nil); err == nil {
			return
		}
	}
	panic("could not connect to peer")
}

package pubsub_test

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/pubsub"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"google.golang.org/grpc"
)

var (
	registeredApplicationID   = &ttnpb.ApplicationIdentifiers{ApplicationId: "foo-app"}
	unregisteredApplicationID = &ttnpb.ApplicationIdentifiers{ApplicationId: "no-app"}
	registeredApplicationUID  = unique.ID(test.Context(), registeredApplicationID)
	registeredApplicationKey  = "secret"
	registeredDeviceID        = &ttnpb.EndDeviceIdentifiers{
		ApplicationIds: registeredApplicationID,
		DeviceId:       "foo-device",
		DevAddr:        types.DevAddr{0x42, 0xff, 0xff, 0xff}.Bytes(),
	}
	unregisteredDeviceID = ttnpb.EndDeviceIdentifiers{
		ApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "bar-app",
		},
		DeviceId: "bar-device",
		DevAddr:  types.DevAddr{0x42, 0x42, 0x42, 0x42}.Bytes(),
	}
	registeredPubSubID = "foo-integration"

	timeout = (1 << 8) * test.Delay
)

type mockRegisterer struct {
	*pubsub.PubSub
}

func (m *mockRegisterer) Roles() []ttnpb.ClusterRole {
	return nil
}

func (m *mockRegisterer) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterApplicationPubSubRegistryServer(s, m.PubSub)
}

func (m *mockRegisterer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
	ttnpb.RegisterApplicationPubSubRegistryHandler(m.PubSub.Context(), s, conn)
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

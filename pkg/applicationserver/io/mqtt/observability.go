package mqtt

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var evtConnectFail = events.Define(
	"as.mqtt.connect.fail", "fail to connect to MQTT",
	events.WithVisibility(ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_READ),
	events.WithErrorDataType(),
)

func registerConnectFail(ctx context.Context, ids *ttnpb.ApplicationIdentifiers, err error) {
	events.Publish(evtConnectFail.NewWithIdentifiersAndData(ctx, ids, err))
}

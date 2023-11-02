package loracloudgeolocationv3

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var evtPackageFail = events.Define(
	"as.packages.loracloudglsv3.fail", "fail to process upstream message",
	events.WithVisibility(ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_READ),
	events.WithErrorDataType(),
	events.WithPropagateToParent(),
)

func registerPackageFail(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, err error) {
	events.Publish(evtPackageFail.NewWithIdentifiersAndData(ctx, ids, err))
}

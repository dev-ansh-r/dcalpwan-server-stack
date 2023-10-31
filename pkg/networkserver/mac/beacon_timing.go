package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// NOTE: This command is deprecated in LoRaWAN 1.0.3

func DeviceNeedsBeaconTimingReq(dev *ttnpb.EndDevice) bool {
	// TODO: Support BeaconTimingReq. (https://github.com/TheThingsNetwork/lorawan-stack/issues/2431)
	return !dev.GetMulticast() && dev.GetMacState().GetDeviceClass() == ttnpb.Class_CLASS_B && false
}

func HandleBeaconTimingReq(ctx context.Context, dev *ttnpb.EndDevice) (events.Builders, error) {
	_ = DeviceNeedsBeaconTimingReq(dev)
	// TODO: Support BeaconTimingReq. (https://github.com/TheThingsNetwork/lorawan-stack/issues/2431)
	return nil, nil
}

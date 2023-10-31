package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	EvtReceiveDeviceModeIndication = defineReceiveMACIndicationEvent(
		"device_mode", "device mode",
		events.WithDataType(&ttnpb.MACCommand_DeviceModeInd{}),
	)()
	EvtEnqueueDeviceModeConfirmation = defineEnqueueMACConfirmationEvent(
		"device_mode", "device mode",
		events.WithDataType(&ttnpb.MACCommand_DeviceModeConf{}),
	)()
)

func HandleDeviceModeInd(ctx context.Context, dev *ttnpb.EndDevice, pld *ttnpb.MACCommand_DeviceModeInd) (events.Builders, error) {
	if pld == nil {
		return nil, ErrNoPayload.New()
	}

	evs := events.Builders{
		EvtReceiveDeviceModeIndication.With(events.WithData(pld)),
	}
	switch {
	case pld.Class == ttnpb.Class_CLASS_C && dev.SupportsClassC && dev.MacState.DeviceClass != ttnpb.Class_CLASS_C:
		evs = append(evs, EvtClassCSwitch.With(events.WithData(dev.MacState.DeviceClass)))
		dev.MacState.DeviceClass = ttnpb.Class_CLASS_C

	case pld.Class == ttnpb.Class_CLASS_A && dev.MacState.DeviceClass != ttnpb.Class_CLASS_A:
		evs = append(evs, EvtClassASwitch.With(events.WithData(dev.MacState.DeviceClass)))
		dev.MacState.DeviceClass = ttnpb.Class_CLASS_A
	}
	conf := &ttnpb.MACCommand_DeviceModeConf{
		Class: dev.MacState.DeviceClass,
	}
	dev.MacState.QueuedResponses = append(dev.MacState.QueuedResponses, conf.MACCommand())
	return append(evs,
		EvtEnqueueDeviceModeConfirmation.With(events.WithData(conf)),
	), nil
}

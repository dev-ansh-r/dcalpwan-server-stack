package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	lorautil "go.thethings.network/lorawan-stack/v3/pkg/util/lora"
)

var (
	EvtReceiveDeviceTimeRequest = defineReceiveMACRequestEvent(
		"device_time", "device time",
	)()
	EvtEnqueueDeviceTimeAnswer = defineEnqueueMACAnswerEvent(
		"device_time", "device time",
		events.WithData(&ttnpb.MACCommand_DeviceTimeAns{}),
	)()
)

func HandleDeviceTimeReq(ctx context.Context, dev *ttnpb.EndDevice, msg *ttnpb.UplinkMessage) (events.Builders, error) {
	ans := &ttnpb.MACCommand_DeviceTimeAns{
		Time: lorautil.GetAdjustedReceivedAt(msg),
	}
	dev.MacState.QueuedResponses = append(dev.MacState.QueuedResponses, ans.MACCommand())
	return events.Builders{
		EvtReceiveDeviceTimeRequest,
		EvtEnqueueDeviceTimeAnswer.With(events.WithData(ans)),
	}, nil
}

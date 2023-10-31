package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/networkserver/internal"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	EvtReceiveLinkCheckRequest = defineReceiveMACRequestEvent(
		"link_check", "link check",
	)()
	EvtEnqueueLinkCheckAnswer = defineEnqueueMACAnswerEvent(
		"link_check", "link check",
		events.WithDataType(&ttnpb.MACCommand_LinkCheckAns{}),
	)()
)

func HandleLinkCheckReq(ctx context.Context, dev *ttnpb.EndDevice, msg *ttnpb.UplinkMessage) (events.Builders, error) {
	evs := events.Builders{
		EvtReceiveLinkCheckRequest,
	}

	var floor float32
	if dr, ok := msg.Settings.DataRate.Modulation.(*ttnpb.DataRate_Lora); ok {
		floor, ok = demodulationFloor[dr.Lora.SpreadingFactor][dr.Lora.Bandwidth]
		if !ok {
			return evs, internal.ErrInvalidDataRate.New()
		}
	}
	if len(msg.RxMetadata) == 0 {
		return evs, nil
	}

	gtwCount, maxSNR := internal.RXMetadataStats(ctx, msg.RxMetadata)
	ans := &ttnpb.MACCommand_LinkCheckAns{
		Margin:       uint32(uint8(maxSNR - floor)),
		GatewayCount: uint32(gtwCount),
	}
	dev.MacState.QueuedResponses = append(dev.MacState.QueuedResponses, ans.MACCommand())
	return append(evs,
		EvtEnqueueLinkCheckAnswer.With(events.WithData(ans)),
	), nil
}

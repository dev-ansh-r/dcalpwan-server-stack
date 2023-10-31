package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	EvtReceiveResetIndication = defineReceiveMACIndicationEvent(
		"reset", "device reset",
		events.WithDataType(&ttnpb.MACCommand_ResetInd{}),
	)()
	EvtEnqueueResetConfirmation = defineEnqueueMACConfirmationEvent(
		"reset", "device reset",
		events.WithDataType(&ttnpb.MACCommand_ResetConf{}),
	)()
)

func HandleResetInd(ctx context.Context, dev *ttnpb.EndDevice, pld *ttnpb.MACCommand_ResetInd, fps *frequencyplans.Store, defaults *ttnpb.MACSettings) (events.Builders, error) {
	if pld == nil {
		return nil, ErrNoPayload.New()
	}

	evs := events.Builders{
		EvtReceiveResetIndication.With(events.WithData(pld)),
	}
	if dev.SupportsJoin {
		return evs, nil
	}

	macState, err := NewState(dev, fps, defaults)
	if err != nil {
		return evs, err
	}
	dev.MacState = macState
	dev.MacState.LorawanVersion = dev.LorawanVersion

	conf := &ttnpb.MACCommand_ResetConf{
		MinorVersion: pld.MinorVersion,
	}
	dev.MacState.QueuedResponses = append(dev.MacState.QueuedResponses, conf.MACCommand())
	return append(evs,
		EvtEnqueueResetConfirmation.With(events.WithData(conf)),
	), nil
}

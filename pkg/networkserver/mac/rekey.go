package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/specification/macspec"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

var (
	EvtReceiveRekeyIndication = defineReceiveMACIndicationEvent(
		"rekey", "device rekey",
		events.WithDataType(&ttnpb.MACCommand_RekeyInd{}),
	)()
	EvtEnqueueRekeyConfirmation = defineEnqueueMACConfirmationEvent(
		"rekey", "device rekey",
		events.WithDataType(&ttnpb.MACCommand_RekeyConf{}),
	)()
)

func HandleRekeyInd(ctx context.Context, dev *ttnpb.EndDevice, pld *ttnpb.MACCommand_RekeyInd, devAddr types.DevAddr) (events.Builders, error) {
	if pld == nil {
		return nil, ErrNoPayload.New()
	}

	evs := events.Builders{
		EvtReceiveRekeyIndication.With(events.WithData(pld)),
	}
	if !dev.SupportsJoin || !macspec.UseRekeyInd(dev.LorawanVersion) {
		return evs, nil
	}
	if dev.PendingSession != nil &&
		dev.MacState.PendingJoinRequest != nil &&
		types.MustDevAddr(dev.PendingSession.DevAddr).OrZero().Equal(devAddr) {
		dev.Ids.DevAddr = dev.PendingSession.DevAddr
		dev.Session = dev.PendingSession
	}

	conf := &ttnpb.MACCommand_RekeyConf{}
	dev.MacState.LorawanVersion, conf.MinorVersion = macspec.NegotiatedVersion(dev.LorawanVersion, pld.MinorVersion)
	dev.MacState.PendingJoinRequest = nil
	dev.PendingMacState = nil
	dev.PendingSession = nil

	dev.MacState.QueuedResponses = append(dev.MacState.QueuedResponses, conf.MACCommand())
	return append(evs,
		EvtEnqueueRekeyConfirmation.With(events.WithData(conf)),
	), nil
}

package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	EvtEnqueueRxTimingSetupRequest = defineEnqueueMACRequestEvent(
		"rx_timing_setup", "Rx timing setup",
		events.WithDataType(&ttnpb.MACCommand_RxTimingSetupReq{}),
	)()
	EvtReceiveRxTimingSetupAnswer = defineReceiveMACAnswerEvent(
		"rx_timing_setup", "Rx timing setup",
	)()

	containsRxTimingSetup = containsMACCommandIdentifier(ttnpb.MACCommandIdentifier_CID_RX_TIMING_SETUP)
	consumeRxTimingSetup  = consumeMACCommandIdentifier(ttnpb.MACCommandIdentifier_CID_RX_TIMING_SETUP)
)

func DeviceNeedsRxTimingSetupReq(dev *ttnpb.EndDevice) bool {
	if dev.GetMulticast() || dev.GetMacState() == nil {
		return false
	}
	macState := dev.MacState
	if containsRxTimingSetup(macState.RecentMacCommandIdentifiers...) { // See STICKY.md.
		return false
	}
	currentParameters, desiredParameters := macState.CurrentParameters, macState.DesiredParameters
	return desiredParameters.Rx1Delay != currentParameters.Rx1Delay
}

func EnqueueRxTimingSetupReq(ctx context.Context, dev *ttnpb.EndDevice, maxDownLen, maxUpLen uint16) EnqueueState {
	if !DeviceNeedsRxTimingSetupReq(dev) {
		return EnqueueState{
			MaxDownLen: maxDownLen,
			MaxUpLen:   maxUpLen,
			Ok:         true,
		}
	}

	var st EnqueueState
	dev.MacState.PendingRequests, st = enqueueMACCommand(ttnpb.MACCommandIdentifier_CID_RX_TIMING_SETUP, maxDownLen, maxUpLen, func(nDown, nUp uint16) ([]*ttnpb.MACCommand, uint16, events.Builders, bool) {
		if nDown < 1 || nUp < 1 {
			return nil, 0, nil, false
		}
		req := &ttnpb.MACCommand_RxTimingSetupReq{
			Delay: dev.MacState.DesiredParameters.Rx1Delay,
		}
		log.FromContext(ctx).WithFields(log.Fields(
			"delay", req.Delay,
		)).Debug("Enqueued RxTimingSetupReq")
		return []*ttnpb.MACCommand{
				req.MACCommand(),
			},
			1,
			events.Builders{
				EvtEnqueueRxTimingSetupRequest.With(events.WithData(req)),
			},
			true
	}, dev.MacState.PendingRequests...)
	return st
}

func HandleRxTimingSetupAns(ctx context.Context, dev *ttnpb.EndDevice) (events.Builders, error) {
	var allowMissing bool // See STICKY.md.
	dev.MacState.RecentMacCommandIdentifiers, allowMissing = consumeRxTimingSetup(
		dev.MacState.RecentMacCommandIdentifiers...,
	)
	var err error
	dev.MacState.PendingRequests, err = handleMACResponse(
		ttnpb.MACCommandIdentifier_CID_RX_TIMING_SETUP,
		allowMissing,
		func(cmd *ttnpb.MACCommand) error {
			req := cmd.GetRxTimingSetupReq()

			dev.MacState.CurrentParameters.Rx1Delay = req.Delay
			return nil
		},
		dev.MacState.PendingRequests...,
	)
	return events.Builders{
		EvtReceiveRxTimingSetupAnswer,
	}, err
}

package mac

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var EvtEnqueueForceRejoinRequest = defineEnqueueMACRequestEvent("force_rejoin", "force rejoin")()

func EnqueueForceRejoinReq(ctx context.Context, dev *ttnpb.EndDevice, maxDownLen, maxUpLen uint16) EnqueueState {
	// TODO: Support rejoins. (https://github.com/TheThingsNetwork/lorawan-stack/issues/8)
	_ = EvtEnqueueForceRejoinRequest
	return EnqueueState{
		MaxDownLen: maxDownLen,
		MaxUpLen:   maxUpLen,
		Ok:         true,
	}
}

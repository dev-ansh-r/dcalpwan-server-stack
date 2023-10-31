package packetbrokeragent

import (
	"context"

	packetbroker "go.packetbroker.org/api/v3"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var errNoPHYPayload = errors.DefineFailedPrecondition("no_phy_payload", "no PHYPayload in message")

func (a *Agent) encryptUplink(ctx context.Context, msg *packetbroker.UplinkMessage) error {
	// TODO: Obtain KEK, encrypt PHYPayload and gateway metadata (https://github.com/TheThingsIndustries/lorawan-stack/issues/1919).
	if msg.PhyPayload.GetPlain() == nil {
		return errNoPHYPayload.New()
	}
	return nil
}

func (a *Agent) decryptUplink(ctx context.Context, msg *packetbroker.UplinkMessage) error {
	// TODO: Obtain KEK, decrypt PHYPayload and gateway metadata (https://github.com/TheThingsIndustries/lorawan-stack/issues/1919).
	if msg.PhyPayload.GetPlain() == nil {
		return errNoPHYPayload.New()
	}
	return nil
}

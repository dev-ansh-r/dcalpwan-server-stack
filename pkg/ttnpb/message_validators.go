package ttnpb

import (
	"context"

	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (p *MACPayload) ValidateContext(context.Context) error {
	if h := p.GetFHdr(); h == nil || types.MustDevAddr(h.DevAddr).OrZero().IsZero() {
		return errMissing("DevAddr")
	}
	return p.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (p *JoinRequestPayload) ValidateContext(context.Context) error {
	if devEUI := types.MustEUI64(p.GetDevEui()).OrZero(); devEUI.IsZero() {
		return errMissing("DevEUI")
	}
	return p.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GatewayUplinkMessage) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *DownlinkQueueRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

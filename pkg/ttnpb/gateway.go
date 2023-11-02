package ttnpb

import types "go.thethings.network/lorawan-stack/v3/pkg/types"

func (r *GetGatewayIdentifiersForEUIRequest) RateLimitKey() string {
	if r == nil || r.Eui == nil {
		return ""
	}
	return types.MustEUI64(r.Eui).OrZero().String()
}

package deviceclaimingserver_test

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MockClaimer is a mock Claimer.
type MockClaimer struct {
	JoinEUIs []types.EUI64

	ClaimFunc        func(context.Context, types.EUI64, types.EUI64, string) error
	BatchUnclaimFunc func(
		context.Context,
		[]*ttnpb.EndDeviceIdentifiers,
	) error
}

// SupportsJoinEUI returns whether the Join Server supports this JoinEUI.
func (m MockClaimer) SupportsJoinEUI(joinEUI types.EUI64) bool {
	for _, eui := range m.JoinEUIs {
		if eui.Equal(joinEUI) {
			return true
		}
	}
	return false
}

// Claim claims an End Device.
func (m MockClaimer) Claim(
	ctx context.Context, joinEUI, devEUI types.EUI64, claimAuthenticationCode string,
) error {
	return m.ClaimFunc(ctx, joinEUI, devEUI, claimAuthenticationCode)
}

// GetClaimStatus returns the claim status an End Device.
func (MockClaimer) GetClaimStatus(_ context.Context,
	ids *ttnpb.EndDeviceIdentifiers,
) (*ttnpb.GetClaimStatusResponse, error) {
	return &ttnpb.GetClaimStatusResponse{
		EndDeviceIds: ids,
	}, nil
}

// Unclaim releases the claim on an End Device.
func (MockClaimer) Unclaim(_ context.Context,
	_ *ttnpb.EndDeviceIdentifiers,
) (err error) {
	return nil
}

// Unclaim releases the claim on an End Device.
func (m MockClaimer) BatchUnclaim(
	ctx context.Context,
	ids []*ttnpb.EndDeviceIdentifiers,
) error {
	return m.BatchUnclaimFunc(ctx, ids)
}

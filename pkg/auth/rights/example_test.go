package rights_test

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func ExampleRequireApplication() {
	var ( // Assume these come from a hypothetical Set RPC call.
		ctx context.Context
		dev ttnpb.EndDevice
	)

	if err := rights.RequireApplication(ctx, dev.Ids.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE); err != nil {
		// return nil, err
	}
}

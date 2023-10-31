package registry

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestUniqueEUIViolationErrDevIDs(t *testing.T) {
	t.Parallel()
	a, ctx := test.New(t)
	devIDs := &ttnpb.EndDeviceIdentifiers{
		DeviceId: "foo-device",
		ApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "foo-app",
		},
	}
	joinEUI := types.EUI64{0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	devEUI := types.EUI64{0x00, 0x00, 0x00, 0x00, 0x00, 0x02}
	devEUIStr := unique.ID(ctx, devIDs)
	expected := errEndDeviceEUIsTaken.WithAttributes(
		"join_eui", joinEUI,
		"dev_eui", devEUI,
		"device_id", devIDs.DeviceId,
		"application_id", devIDs.ApplicationIds.ApplicationId,
	)
	actual := UniqueEUIViolationErr(ctx, joinEUI, devEUI, devEUIStr)
	a.So(actual, should.NotBeNil)
	a.So(actual.Error(), should.Equal, expected.Error())
}

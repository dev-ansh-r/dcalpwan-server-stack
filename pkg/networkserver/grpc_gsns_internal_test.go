package networkserver

import (
	"context"
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestAppendRecentUplink(t *testing.T) {
	t.Parallel()
	dataRate := (&ttnpb.LoRaDataRate{
		SpreadingFactor: uint32(7),
		Bandwidth:       125_000,
		CodingRate:      "4/5",
	}).DataRate()
	macSettings := &ttnpb.MACState_UplinkMessage_TxSettings{
		DataRate: dataRate,
	}
	settings := &ttnpb.TxSettings{
		DataRate: dataRate,
	}
	ups := [...]*ttnpb.MACState_UplinkMessage{
		{
			DeviceChannelIndex: 1,
			Settings:           macSettings,
		},
		{
			DeviceChannelIndex: 2,
			Settings:           macSettings,
		},
		{
			DeviceChannelIndex: 3,
			Settings:           macSettings,
		},
	}
	for _, tc := range []struct {
		Recent   []*ttnpb.MACState_UplinkMessage
		Up       *ttnpb.UplinkMessage
		Window   int
		Expected []*ttnpb.MACState_UplinkMessage
	}{
		{
			Up: &ttnpb.UplinkMessage{
				DeviceChannelIndex: 1,
				Settings:           settings,
			},
			Window:   1,
			Expected: ups[:1],
		},
		{
			Recent: ups[:1],
			Up: &ttnpb.UplinkMessage{
				DeviceChannelIndex: 2,
				Settings:           settings,
			},
			Window:   1,
			Expected: ups[1:2],
		},
		{
			Recent: ups[:2],
			Up: &ttnpb.UplinkMessage{
				DeviceChannelIndex: 3,
				Settings:           settings,
			},
			Window:   1,
			Expected: ups[2:3],
		},
		{
			Recent: ups[:1],
			Up: &ttnpb.UplinkMessage{
				DeviceChannelIndex: 2,
				Settings:           settings,
			},
			Window:   2,
			Expected: ups[:2],
		},
		{
			Recent: ups[:2],
			Up: &ttnpb.UplinkMessage{
				DeviceChannelIndex: 3,
				Settings:           settings,
			},
			Window:   2,
			Expected: ups[1:3],
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     fmt.Sprintf("recent_length:%d,window:%v", len(tc.Recent), tc.Window),
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				recent := ttnpb.CloneSlice(tc.Recent)
				up := ttnpb.Clone(tc.Up)
				ret := appendRecentUplink(recent, up, tc.Window)
				a.So(recent, should.Resemble, tc.Recent)
				a.So(up, should.Resemble, tc.Up)
				a.So(ret, should.Resemble, tc.Expected)
			},
		})
	}
}

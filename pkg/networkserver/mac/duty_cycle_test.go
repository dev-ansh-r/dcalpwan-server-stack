package mac_test

import (
	"context"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/mac"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestNeedsDutyCycleReq(t *testing.T) {
	for _, tc := range []struct {
		Name        string
		InputDevice *ttnpb.EndDevice
		Needs       bool
	}{
		{
			Name:        "no MAC state",
			InputDevice: &ttnpb.EndDevice{},
		},
		{
			Name: "current(max-duty-cycle:1024),desired(max-duty-cycle:1024)",
			InputDevice: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{
						MaxDutyCycle: ttnpb.AggregatedDutyCycle_DUTY_CYCLE_1024,
					},
					DesiredParameters: &ttnpb.MACParameters{
						MaxDutyCycle: ttnpb.AggregatedDutyCycle_DUTY_CYCLE_1024,
					},
				},
			},
		},
		{
			Name: "current(max-duty-cycle:1024),desired(max-duty-cycle:2048)",
			InputDevice: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{
						MaxDutyCycle: ttnpb.AggregatedDutyCycle_DUTY_CYCLE_1024,
					},
					DesiredParameters: &ttnpb.MACParameters{
						MaxDutyCycle: ttnpb.AggregatedDutyCycle_DUTY_CYCLE_2048,
					},
				},
			},
			Needs: true,
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := ttnpb.Clone(tc.InputDevice)
				res := DeviceNeedsDutyCycleReq(dev)
				if tc.Needs {
					a.So(res, should.BeTrue)
				} else {
					a.So(res, should.BeFalse)
				}
				a.So(dev, should.Resemble, tc.InputDevice)
			},
		})
	}
}

func TestHandleDutyCycleAns(t *testing.T) {
	for _, tc := range []struct {
		Name             string
		Device, Expected *ttnpb.EndDevice
		Events           events.Builders
		Error            error
	}{
		{
			Name: "no request",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Events: events.Builders{
				EvtReceiveDutyCycleAnswer,
			},
			Error: ErrRequestNotFound.WithAttributes("cid", ttnpb.MACCommandIdentifier_CID_DUTY_CYCLE),
		},
		{
			Name: "2048",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_DutyCycleReq{
							MaxDutyCycle: ttnpb.AggregatedDutyCycle_DUTY_CYCLE_2048,
						}).MACCommand(),
					},
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{},
					CurrentParameters: &ttnpb.MACParameters{
						MaxDutyCycle: ttnpb.AggregatedDutyCycle_DUTY_CYCLE_2048,
					},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Events: events.Builders{
				EvtReceiveDutyCycleAnswer,
			},
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := ttnpb.Clone(tc.Device)

				evs, err := HandleDutyCycleAns(ctx, dev)
				if tc.Error != nil && !a.So(err, should.EqualErrorOrDefinition, tc.Error) ||
					tc.Error == nil && !a.So(err, should.BeNil) {
					t.FailNow()
				}
				a.So(dev, should.Resemble, tc.Expected)
				a.So(evs, should.ResembleEventBuilders, tc.Events)
			},
		})
	}
}

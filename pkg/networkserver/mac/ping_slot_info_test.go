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

func TestHandlePingSlotInfoReq(t *testing.T) {
	for _, tc := range []struct {
		Name             string
		Device, Expected *ttnpb.EndDevice
		Payload          *ttnpb.MACCommand_PingSlotInfoReq
		Events           events.Builders
		Error            error
	}{
		{
			Name: "nil payload",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass: ttnpb.Class_CLASS_B,
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass: ttnpb.Class_CLASS_B,
				},
			},
			Error: ErrNoPayload,
		},
		{
			Name: "class B device",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass: ttnpb.Class_CLASS_B,
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass: ttnpb.Class_CLASS_B,
				},
			},
			Payload: &ttnpb.MACCommand_PingSlotInfoReq{
				Period: 42,
			},
			Events: events.Builders{
				EvtReceivePingSlotInfoRequest.With(events.WithData(&ttnpb.MACCommand_PingSlotInfoReq{
					Period: 42,
				})),
			},
		},
		{
			Name: "empty queue",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass: ttnpb.Class_CLASS_A,
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass:         ttnpb.Class_CLASS_A,
					PingSlotPeriodicity: &ttnpb.PingSlotPeriodValue{Value: 42},
					QueuedResponses: []*ttnpb.MACCommand{
						ttnpb.MACCommandIdentifier_CID_PING_SLOT_INFO.MACCommand(),
					},
				},
			},
			Payload: &ttnpb.MACCommand_PingSlotInfoReq{
				Period: 42,
			},
			Events: events.Builders{
				EvtReceivePingSlotInfoRequest.With(events.WithData(&ttnpb.MACCommand_PingSlotInfoReq{
					Period: 42,
				})),
				EvtEnqueuePingSlotInfoAnswer,
			},
		},
		{
			Name: "non-empty queue",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass: ttnpb.Class_CLASS_A,
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
					},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					DeviceClass:         ttnpb.Class_CLASS_A,
					PingSlotPeriodicity: &ttnpb.PingSlotPeriodValue{Value: 42},
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
						ttnpb.MACCommandIdentifier_CID_PING_SLOT_INFO.MACCommand(),
					},
				},
			},
			Payload: &ttnpb.MACCommand_PingSlotInfoReq{
				Period: 42,
			},
			Events: events.Builders{
				EvtReceivePingSlotInfoRequest.With(events.WithData(&ttnpb.MACCommand_PingSlotInfoReq{
					Period: 42,
				})),
				EvtEnqueuePingSlotInfoAnswer,
			},
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := ttnpb.Clone(tc.Device)

				evs, err := HandlePingSlotInfoReq(ctx, dev, tc.Payload)
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

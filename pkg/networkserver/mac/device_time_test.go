package mac_test

import (
	"context"
	"testing"
	"time"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/mac"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHandleDeviceTimeReq(t *testing.T) {
	recvAt := time.Unix(42, 42).UTC()
	mdRecvAt, gpsTime := recvAt.Add(-250), recvAt.Add(-500)
	for _, tc := range []struct {
		Name             string
		Device, Expected *ttnpb.EndDevice
		Message          *ttnpb.UplinkMessage
		Events           events.Builders
		Error            error
	}{
		{
			Name: "empty queue",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					QueuedResponses: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_DeviceTimeAns{
							Time: timestamppb.New(recvAt),
						}).MACCommand(),
					},
				},
			},
			Message: &ttnpb.UplinkMessage{
				ReceivedAt: timestamppb.New(recvAt),
			},
			Events: events.Builders{
				EvtReceiveDeviceTimeRequest,
				EvtEnqueueDeviceTimeAnswer.With(events.WithData(&ttnpb.MACCommand_DeviceTimeAns{
					Time: timestamppb.New(recvAt),
				})),
			},
		},
		{
			Name: "non-empty queue/odd",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
					},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
						(&ttnpb.MACCommand_DeviceTimeAns{
							Time: timestamppb.New(recvAt),
						}).MACCommand(),
					},
				},
			},
			Message: &ttnpb.UplinkMessage{
				ReceivedAt: timestamppb.New(recvAt),
			},
			Events: events.Builders{
				EvtReceiveDeviceTimeRequest,
				EvtEnqueueDeviceTimeAnswer.With(events.WithData(&ttnpb.MACCommand_DeviceTimeAns{
					Time: timestamppb.New(recvAt),
				})),
			},
		},
		{
			Name: "non-empty queue/even",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
					},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
						(&ttnpb.MACCommand_DeviceTimeAns{
							Time: timestamppb.New(recvAt),
						}).MACCommand(),
					},
				},
			},
			Message: &ttnpb.UplinkMessage{
				ReceivedAt: timestamppb.New(recvAt),
			},
			Events: events.Builders{
				EvtReceiveDeviceTimeRequest,
				EvtEnqueueDeviceTimeAnswer.With(events.WithData(&ttnpb.MACCommand_DeviceTimeAns{
					Time: timestamppb.New(recvAt),
				})),
			},
		},
		{
			Name: "time from rx_metadata",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					QueuedResponses: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_DeviceTimeAns{
							Time: timestamppb.New(mdRecvAt),
						}).MACCommand(),
					},
				},
			},
			Message: &ttnpb.UplinkMessage{
				ReceivedAt: timestamppb.New(recvAt),
				RxMetadata: []*ttnpb.RxMetadata{
					{
						ReceivedAt: timestamppb.New(mdRecvAt),
					},
					{
						ReceivedAt: timestamppb.New(mdRecvAt.Add(200)),
					},
				},
			},
			Events: events.Builders{
				EvtReceiveDeviceTimeRequest,
				EvtEnqueueDeviceTimeAnswer.With(events.WithData(&ttnpb.MACCommand_DeviceTimeAns{
					Time: timestamppb.New(mdRecvAt),
				})),
			},
		},
		{
			Name: "gps_time from rx_metadata",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					QueuedResponses: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_DeviceTimeAns{
							Time: timestamppb.New(gpsTime),
						}).MACCommand(),
					},
				},
			},
			Message: &ttnpb.UplinkMessage{
				ReceivedAt: timestamppb.New(recvAt),
				RxMetadata: []*ttnpb.RxMetadata{
					{
						GpsTime:    timestamppb.New(gpsTime),
						ReceivedAt: timestamppb.New(mdRecvAt),
					},
					{
						ReceivedAt: timestamppb.New(mdRecvAt.Add(200)),
					},
					{
						GpsTime: timestamppb.New(gpsTime.Add(150)),
					},
				},
			},
			Events: events.Builders{
				EvtReceiveDeviceTimeRequest,
				EvtEnqueueDeviceTimeAnswer.With(events.WithData(&ttnpb.MACCommand_DeviceTimeAns{
					Time: timestamppb.New(gpsTime),
				})),
			},
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := ttnpb.Clone(tc.Device)

				evs, err := HandleDeviceTimeReq(ctx, dev, tc.Message)
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

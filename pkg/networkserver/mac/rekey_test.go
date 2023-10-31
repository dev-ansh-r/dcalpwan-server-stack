package mac_test

import (
	"context"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/mac"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestHandleRekeyInd(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		Name             string
		Device, Expected *ttnpb.EndDevice
		Payload          *ttnpb.MACCommand_RekeyInd
		Events           events.Builders
		Error            error
	}{
		{
			Name: "nil payload",
			Device: &ttnpb.EndDevice{
				SupportsJoin: true,
			},
			Expected: &ttnpb.EndDevice{
				SupportsJoin: true,
			},
			Error: ErrNoPayload,
		},
		{
			Name: "empty queue/original",
			Device: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids:            &ttnpb.EndDeviceIdentifiers{},
				PendingSession: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				PendingMacState: &ttnpb.MACState{},
				MacState: &ttnpb.MACState{
					PendingJoinRequest: &ttnpb.MACState_JoinRequest{},
					QueuedResponses:    []*ttnpb.MACCommand{},
				},
			},
			Expected: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids: &ttnpb.EndDeviceIdentifiers{
					DevAddr: test.DefaultDevAddr.Bytes(),
				},
				Session: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				MacState: &ttnpb.MACState{
					LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
					QueuedResponses: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_RekeyConf{
							MinorVersion: 1,
						}).MACCommand(),
					},
				},
			},
			Payload: &ttnpb.MACCommand_RekeyInd{
				MinorVersion: 1,
			},
			Events: events.Builders{
				EvtReceiveRekeyIndication.With(events.WithData(&ttnpb.MACCommand_RekeyInd{
					MinorVersion: 1,
				})),
				EvtEnqueueRekeyConfirmation.With(events.WithData(&ttnpb.MACCommand_RekeyConf{
					MinorVersion: 1,
				})),
			},
		},
		{
			Name: "empty queue/retransmission/non-matching pending session",
			Device: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids: &ttnpb.EndDeviceIdentifiers{
					DevAddr: test.DefaultDevAddr.Bytes(),
				},
				Session: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				PendingSession: &ttnpb.Session{
					DevAddr:       types.DevAddr{0x23, 0x23, 0x11, 0x42}.Bytes(),
					LastFCntUp:    101,
					LastNFCntDown: 2,
				},
				PendingMacState: &ttnpb.MACState{},
				MacState: &ttnpb.MACState{
					PendingJoinRequest: &ttnpb.MACState_JoinRequest{},
					QueuedResponses:    []*ttnpb.MACCommand{},
				},
			},
			Expected: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids: &ttnpb.EndDeviceIdentifiers{
					DevAddr: test.DefaultDevAddr.Bytes(),
				},
				Session: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				MacState: &ttnpb.MACState{
					LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
					QueuedResponses: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_RekeyConf{
							MinorVersion: 1,
						}).MACCommand(),
					},
				},
			},
			Payload: &ttnpb.MACCommand_RekeyInd{
				MinorVersion: 1,
			},
			Events: events.Builders{
				EvtReceiveRekeyIndication.With(events.WithData(&ttnpb.MACCommand_RekeyInd{
					MinorVersion: 1,
				})),
				EvtEnqueueRekeyConfirmation.With(events.WithData(&ttnpb.MACCommand_RekeyConf{
					MinorVersion: 1,
				})),
			},
		},
		{
			Name: "empty queue/retransmission/no pending session",
			Device: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids: &ttnpb.EndDeviceIdentifiers{
					DevAddr: test.DefaultDevAddr.Bytes(),
				},
				Session: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				MacState: &ttnpb.MACState{},
			},
			Expected: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids: &ttnpb.EndDeviceIdentifiers{
					DevAddr: test.DefaultDevAddr.Bytes(),
				},
				Session: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				MacState: &ttnpb.MACState{
					LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
					QueuedResponses: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_RekeyConf{
							MinorVersion: 1,
						}).MACCommand(),
					},
				},
			},
			Payload: &ttnpb.MACCommand_RekeyInd{
				MinorVersion: 1,
			},
			Events: events.Builders{
				EvtReceiveRekeyIndication.With(events.WithData(&ttnpb.MACCommand_RekeyInd{
					MinorVersion: 1,
				})),
				EvtEnqueueRekeyConfirmation.With(events.WithData(&ttnpb.MACCommand_RekeyConf{
					MinorVersion: 1,
				})),
			},
		},
		{
			Name: "non-empty queue",
			Device: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids:            &ttnpb.EndDeviceIdentifiers{},
				PendingSession: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				PendingMacState: &ttnpb.MACState{},
				MacState: &ttnpb.MACState{
					PendingJoinRequest: &ttnpb.MACState_JoinRequest{},
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
					},
				},
			},
			Expected: &ttnpb.EndDevice{
				LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
				SupportsJoin:   true,
				Ids: &ttnpb.EndDeviceIdentifiers{
					DevAddr: test.DefaultDevAddr.Bytes(),
				},
				Session: &ttnpb.Session{
					DevAddr:       test.DefaultDevAddr.Bytes(),
					LastFCntUp:    42,
					LastNFCntDown: 43,
				},
				MacState: &ttnpb.MACState{
					LorawanVersion: ttnpb.MACVersion_MAC_V1_1,
					QueuedResponses: []*ttnpb.MACCommand{
						{},
						{},
						{},
						(&ttnpb.MACCommand_RekeyConf{
							MinorVersion: 1,
						}).MACCommand(),
					},
				},
			},
			Payload: &ttnpb.MACCommand_RekeyInd{
				MinorVersion: 1,
			},
			Events: events.Builders{
				EvtReceiveRekeyIndication.With(events.WithData(&ttnpb.MACCommand_RekeyInd{
					MinorVersion: 1,
				})),
				EvtEnqueueRekeyConfirmation.With(events.WithData(&ttnpb.MACCommand_RekeyConf{
					MinorVersion: 1,
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

				evs, err := HandleRekeyInd(ctx, dev, tc.Payload, test.DefaultDevAddr)
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

package alcsyncv1

import (
	"testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimeSynchronizationCommandCalculatesCorrection(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name     string
		Command  Command
		Expected Result
	}{
		{
			Name: "NegativeTimeCorrection",
			Command: &TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime.Add(10 * time.Second)),
					TokenReq:    1,
					AnsRequired: true,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
			Expected: &TimeSyncCommandResult{
				ans: &ttnpb.ALCSyncCommand_AppTimeAns{
					TimeCorrection: -10,
					TokenAns:       1,
				},
			},
		},
		{
			Name: "PositiveTimeCorrection",
			Command: &TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime.Add(-10 * time.Second)),
					TokenReq:    1,
					AnsRequired: true,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
			Expected: &TimeSyncCommandResult{
				ans: &ttnpb.ALCSyncCommand_AppTimeAns{
					TimeCorrection: 10,
					TokenAns:       1,
				},
			},
		},
		{
			Name: "NoTimeCorrection",
			Command: &TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime),
					TokenReq:    1,
					AnsRequired: true,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
			Expected: &TimeSyncCommandResult{
				ans: &ttnpb.ALCSyncCommand_AppTimeAns{
					TimeCorrection: 0,
					TokenAns:       1,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			a, _ := test.New(t)
			result, err := tc.Command.Execute()
			a.So(err, should.BeNil)
			a.So(result, should.NotBeNil)
			a.So(result, should.Resemble, tc.Expected)
		})
	}
}

func TestTimeSynchronizationCommandRespectsThreshold(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name    string
		Command TimeSyncCommand
	}{
		{
			Name: "NegativeTimeCorrection",
			Command: TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime.Add(2 * time.Second)),
					TokenReq:    1,
					AnsRequired: false,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
		},
		{
			Name: "PositiveTimeCorrection",
			Command: TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime.Add(-2 * time.Second)),
					TokenReq:    1,
					AnsRequired: false,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
		},
		{
			Name: "NoTimeCorrection",
			Command: TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime),
					TokenReq:    1,
					AnsRequired: false,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			a, _ := test.New(t)
			result, err := tc.Command.Execute()
			a.So(err, should.NotBeNil)
			a.So(errors.IsUnavailable(err), should.BeTrue)
			a.So(result, should.BeNil)
		})
	}
}

func TestTimeSynchronizationCommandRespectsAnsRequired(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name    string
		Command TimeSyncCommand
	}{
		{
			Name: "NegativeTimeCorrection",
			Command: TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime.Add(2 * time.Second)),
					TokenReq:    1,
					AnsRequired: true,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
		},
		{
			Name: "PositiveTimeCorrection",
			Command: TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime.Add(-2 * time.Second)),
					TokenReq:    1,
					AnsRequired: true,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  threeSecondsDuration,
			},
		},
		{
			Name: "NoTimeCorrection",
			Command: TimeSyncCommand{
				req: &ttnpb.ALCSyncCommand_AppTimeReq{
					DeviceTime:  timestamppb.New(receivedAtTime),
					TokenReq:    1,
					AnsRequired: true,
				},
				receivedAt: receivedAtTime,
				fPort:      202,
				threshold:  3,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			a, _ := test.New(t)
			result, err := tc.Command.Execute()
			a.So(err, should.BeNil)
			a.So(result, should.NotBeNil)
		})
	}
}

func TestTimeSyncCommandResultMarshalsBytesCorrectly(t *testing.T) {
	t.Parallel()
	a, _ := test.New(t)
	expected := []byte{0x01, 0x05, 0x00, 0x00, 0x00, 0x02}
	result := &TimeSyncCommandResult{
		ans: &ttnpb.ALCSyncCommand_AppTimeAns{
			TimeCorrection: 5,
			TokenAns:       2,
		},
	}
	actual, err := result.MarshalBinary()
	a.So(err, should.BeNil)
	a.So(actual, should.Resemble, expected)
}

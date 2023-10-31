package alcsyncv1

import (
	"encoding/binary"
	"math"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// TimeSyncCommand is the command for time synchronization.
type TimeSyncCommand struct {
	req        *ttnpb.ALCSyncCommand_AppTimeReq
	receivedAt time.Time
	threshold  time.Duration
	fPort      uint32
}

// Code implements commands.Command.
func (*TimeSyncCommand) Code() ttnpb.ALCSyncCommandIdentifier {
	return ttnpb.ALCSyncCommandIdentifier_ALCSYNC_CID_APP_TIME
}

// Execute implements commands.Command.
func (cmd *TimeSyncCommand) Execute() (Result, error) {
	deviceTime := cmd.req.DeviceTime.AsTime()
	difference := cmd.receivedAt.Sub(deviceTime)
	exceedsThreshold := math.Abs(difference.Seconds()) > cmd.threshold.Seconds()
	if !cmd.req.AnsRequired && !exceedsThreshold {
		return nil, errIgnoreDownlink.New()
	}

	result := &TimeSyncCommandResult{
		ans: &ttnpb.ALCSyncCommand_AppTimeAns{
			TimeCorrection: int32(difference.Seconds()),
			TokenAns:       cmd.req.TokenReq,
		},
	}
	return result, nil
}

// CommandReceivedEventBuilder implements commands.Command.
func (cmd *TimeSyncCommand) CommandReceivedEventBuilder() events.Builder {
	return EvtTimeCorrectionCmdReceived.With(events.WithData(cmd.req))
}

// Ensure that TimeSyncCommand implements commands.Command.
var _ Command = (*TimeSyncCommand)(nil)

// TimeSyncCommandResult is the result of the time synchronization command.
type TimeSyncCommandResult struct {
	ans *ttnpb.ALCSyncCommand_AppTimeAns
}

// AnswerEnqueuedEventBuilder implements commands.Command.
func (r *TimeSyncCommandResult) AnswerEnqueuedEventBuilder() events.Builder {
	return EvtTimeCorrectionAnsEnqueue.With(events.WithData(r.ans))
}

// MarshalBinary implements Result.
func (r *TimeSyncCommandResult) MarshalBinary() ([]byte, error) {
	// CID - byte 0.
	// DeviceTime - bytes [1,4].
	// Param - byte 5 (bits: RFU [7,4]; TokenAns [3,0]).

	cPayload := make([]byte, 6)
	cPayload[0] = 0x01
	binary.LittleEndian.PutUint32(cPayload[1:5], uint32(r.ans.TimeCorrection))
	cPayload[5] = uint8(r.ans.TokenAns) & 0x0F
	return cPayload, nil
}

// Ensure that TimeSyncCommandResult implements Result.
var _ Result = (*TimeSyncCommandResult)(nil)

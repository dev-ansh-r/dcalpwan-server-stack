package alcsyncv1

import (
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Result is the result of a command execution.
type Result interface {
	// MarshalBinary marshals the result into a binary representation.
	MarshalBinary() ([]byte, error)

	// AnswerEnqueuedEventBuilder returns the event that should be emitted when the answer is enqueued or executed.
	AnswerEnqueuedEventBuilder() events.Builder
}

// Command is the interface for commands.
type Command interface {
	// Code returns the command code.
	Code() ttnpb.ALCSyncCommandIdentifier

	// CommandReceivedEventBuilder returns the event that should be emitted when the command is successfully parsed.
	CommandReceivedEventBuilder() events.Builder

	// Execute runs the command logic.
	Execute() (Result, error)
}

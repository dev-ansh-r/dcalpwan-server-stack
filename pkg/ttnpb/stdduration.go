package ttnpb

import (
	"time"

	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

// StdDuration converts a protobuf duration to a standard library duration.
//
// ProtoDuration panics if the Duration is invalid.
func StdDuration(protoDuration *durationpb.Duration) *time.Duration {
	if protoDuration == nil {
		return nil
	}
	stdDuration := protoDuration.AsDuration()
	return &stdDuration
}

// StdDurationOrZero converts a protobuf duration to a standard library duration.
// If protoDuration is nil, it returns a zero duration.
func StdDurationOrZero(protoDuration *durationpb.Duration) time.Duration {
	stdDuration := StdDuration(protoDuration)
	if stdDuration == nil {
		return 0
	}
	return *stdDuration
}

// ProtoDuration converts a standard library duration to a protobuf duration.
func ProtoDuration(stdDuration *time.Duration) *durationpb.Duration {
	if stdDuration == nil {
		return nil
	}
	return durationpb.New(*stdDuration)
}

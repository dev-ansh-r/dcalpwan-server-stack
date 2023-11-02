package ttnpb

import (
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// StdTime converts a protobuf timestamp to a standard library time.
//
// ProtoTime panics if the time is invalid.
func StdTime(protoTime *timestamppb.Timestamp) *time.Time {
	if protoTime == nil {
		return nil
	}
	stdTime := protoTime.AsTime()
	return &stdTime
}

// StdTimeOrZero converts a protobuf time to a standard library time.
// If protoTime is nil, it returns a zero time.
func StdTimeOrZero(protoTime *timestamppb.Timestamp) time.Time {
	stdTime := StdTime(protoTime)
	if stdTime == nil {
		return time.Time{}
	}
	return *stdTime
}

// ProtoTime converts a standard library time to a protobuf timestamp.
//
// ProtoTime panics if the time is invalid.
func ProtoTime(stdTime *time.Time) *timestamppb.Timestamp {
	if stdTime == nil {
		return nil
	}
	return timestamppb.New(*stdTime)
}

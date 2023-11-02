package ttnpb

import "google.golang.org/protobuf/proto"

// Clone creates a deep copy of the given message.
func Clone[X proto.Message](in X) X {
	return proto.Clone(in).(X)
}

// CloneSlice creates a deep copy of the given slice of messages.
func CloneSlice[X proto.Message](in []X) []X {
	if in == nil {
		return nil
	}
	out := make([]X, len(in))
	for i, x := range in {
		out[i] = Clone(x)
	}
	return out
}

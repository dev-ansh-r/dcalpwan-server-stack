package ttnpb

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// MustMarshalAny converts the proto message to an Any, or panics.
func MustMarshalAny(pb proto.Message) *anypb.Any {
	any, err := anypb.New(pb)
	if err != nil {
		panic(err)
	}
	return any
}

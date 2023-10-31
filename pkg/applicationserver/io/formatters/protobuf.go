package formatters

import (
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/proto"
)

type protobuf struct{}

func (protobuf) FromUp(msg *ttnpb.ApplicationUp) ([]byte, error) {
	return proto.Marshal(msg)
}

func (protobuf) ToDownlinks(buf []byte) (*ttnpb.ApplicationDownlinks, error) {
	res := &ttnpb.ApplicationDownlinks{}
	if err := proto.Unmarshal(buf, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (protobuf) ToDownlinkQueueRequest(buf []byte) (*ttnpb.DownlinkQueueRequest, error) {
	res := &ttnpb.DownlinkQueueRequest{}
	if err := proto.Unmarshal(buf, res); err != nil {
		return nil, err
	}
	return res, nil
}

// Protobuf is a formatter that uses proto marshaling.
var Protobuf Formatter = &protobuf{}

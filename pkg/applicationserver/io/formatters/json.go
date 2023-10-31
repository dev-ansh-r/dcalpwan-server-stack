package formatters

import (
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type json struct{}

func (json) FromUp(msg *ttnpb.ApplicationUp) ([]byte, error) {
	return jsonpb.TTN().Marshal(msg)
}

func (json) ToDownlinks(data []byte) (*ttnpb.ApplicationDownlinks, error) {
	res := &ttnpb.ApplicationDownlinks{}
	if err := jsonpb.TTN().Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (json) ToDownlinkQueueRequest(data []byte) (*ttnpb.DownlinkQueueRequest, error) {
	res := ttnpb.DownlinkQueueRequest{}
	if err := jsonpb.TTN().Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// JSON is a formatter that uses JSON marshaling.
var JSON Formatter = &json{}

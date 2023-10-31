package mqtt

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/mqtt/topics"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/proto"
)

type protobuf struct {
	topics.Layout
}

func (protobuf) FromDownlink(down *ttnpb.DownlinkMessage, _ *ttnpb.GatewayIdentifiers) ([]byte, error) {
	gwDown := &ttnpb.GatewayDown{
		DownlinkMessage: down,
	}
	return proto.Marshal(gwDown)
}

func (protobuf) ToUplink(message []byte, _ *ttnpb.GatewayIdentifiers) (*ttnpb.UplinkMessage, error) {
	uplink := &ttnpb.UplinkMessage{}
	if err := proto.Unmarshal(message, uplink); err != nil {
		return nil, err
	}
	return uplink, nil
}

func (protobuf) ToStatus(message []byte, _ *ttnpb.GatewayIdentifiers) (*ttnpb.GatewayStatus, error) {
	status := &ttnpb.GatewayStatus{}
	if err := proto.Unmarshal(message, status); err != nil {
		return nil, err
	}
	return status, nil
}

func (protobuf) ToTxAck(message []byte, _ *ttnpb.GatewayIdentifiers) (*ttnpb.TxAcknowledgment, error) {
	ack := &ttnpb.TxAcknowledgment{}
	if err := proto.Unmarshal(message, ack); err != nil {
		return nil, err
	}
	return ack, nil
}

// NewProtobuf returns a format that uses Protocol Buffers marshaling and unmarshaling.
func NewProtobuf(ctx context.Context) Format {
	return &protobuf{
		Layout: topics.New(ctx),
	}
}

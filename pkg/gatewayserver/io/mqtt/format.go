package mqtt

import (
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/mqtt/topics"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Format formats topics, downlink, uplink and status messages.
type Format interface {
	topics.Layout

	FromDownlink(down *ttnpb.DownlinkMessage, ids *ttnpb.GatewayIdentifiers) ([]byte, error)
	ToUplink(message []byte, ids *ttnpb.GatewayIdentifiers) (*ttnpb.UplinkMessage, error)
	ToStatus(message []byte, ids *ttnpb.GatewayIdentifiers) (*ttnpb.GatewayStatus, error)
	ToTxAck(message []byte, ids *ttnpb.GatewayIdentifiers) (*ttnpb.TxAcknowledgment, error)
}

var errNotSupported = errors.DefineFailedPrecondition("not_supported", "not supported")

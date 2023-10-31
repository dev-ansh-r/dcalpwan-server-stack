package upstream

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// Handler represents the upstream handler that connects to an upstream host.
type Handler interface {
	// DevAddrPrefixes returns the DevAddr prefixes for this upstream handler.
	DevAddrPrefixes() []types.DevAddrPrefix
	// Setup performs all the preparation necessary to connect the handler to a particular upstream host.
	Setup(context.Context) error
	// ConnectGateway informs the upstream handler that a particular gateway is connected to the frontend.
	ConnectGateway(context.Context, *ttnpb.GatewayIdentifiers, *io.Connection) error
	// HandleUp handles ttnpb.GatewayUplinkMessage. It must not mutate the gateway uplink message.
	HandleUplink(context.Context, *ttnpb.GatewayIdentifiers, *ttnpb.EndDeviceIdentifiers, *ttnpb.GatewayUplinkMessage) error
	// HandleStatus handles ttnpb.GatewayStatus.
	HandleStatus(context.Context, *ttnpb.GatewayIdentifiers, *ttnpb.GatewayStatus) error
	// HandleTxAck handles ttnpb.TxAcknowledgment.
	HandleTxAck(context.Context, *ttnpb.GatewayIdentifiers, *ttnpb.TxAcknowledgment) error
}

package ws

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/scheduling"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// ServerInfo contains information on the WebSocket server.
type ServerInfo struct {
	Scheme  string
	Address string
}

// Endpoints contains the connection endpoints for the protocol
type Endpoints struct {
	ConnectionInfo string
	Traffic        string
}

// Formatter abstracts messages to/from websocket based gateways.
type Formatter interface {
	// Endpoints fetches the connection endpoints for the protocol.
	Endpoints() Endpoints
	// HandleConnectionInfo handles connection information requests from web socket based protocols.
	// This function returns a byte stream that contains connection information (ex: scheme, host, port etc) or an error if applicable.
	HandleConnectionInfo(ctx context.Context, raw []byte, server io.Server, serverInfo ServerInfo, receivedAt time.Time) []byte
	// HandleUp handles upstream messages from web socket based gateways.
	// This function optionally returns a byte stream to be sent as response to the upstream message.
	HandleUp(ctx context.Context, raw []byte, ids *ttnpb.GatewayIdentifiers, conn *io.Connection, receivedAt time.Time) ([]byte, error)
	// FromDownlink generates a downlink byte stream that can be sent over the WS connection.
	FromDownlink(ctx context.Context, down *ttnpb.DownlinkMessage, bandID string, dlTime time.Time) ([]byte, error)
	// TransferTime generates a spurious time transfer message for a particular server time.
	TransferTime(ctx context.Context, serverTime time.Time, gpsTime *time.Time, concentratorTime *scheduling.ConcentratorTime) ([]byte, error)
}

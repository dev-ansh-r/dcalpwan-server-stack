package packetbroker

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// GatewayIdentifiers are the proxy gateway identifier of gateways connected through Packet Broker.
var GatewayIdentifiers = &ttnpb.GatewayIdentifiers{GatewayId: "packetbroker"}

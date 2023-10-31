package shared

import (
	"go.thethings.network/lorawan-stack/v3/cmd/internal/shared"
	gs "go.thethings.network/lorawan-stack/v3/cmd/internal/shared/gatewayserver"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayconfigurationserver"
)

// DefaultGatewayConfigurationServerConfig is the default configuration for the Gateway Configuration Server.
var DefaultGatewayConfigurationServerConfig = gatewayconfigurationserver.Config{
	RequireAuth: true,
}

func init() {
	DefaultGatewayConfigurationServerConfig.TheThingsGateway.Default.UpdateChannel = "stable"
	DefaultGatewayConfigurationServerConfig.TheThingsGateway.Default.MQTTServer = "mqtts://" + gs.DefaultGatewayServerConfig.MQTTV2.PublicTLSAddress
	DefaultGatewayConfigurationServerConfig.TheThingsGateway.Default.FirmwareURL = "https://ttkg-fw.thethingsindustries.com/v1"
	DefaultGatewayConfigurationServerConfig.BasicStation.Default.LNSURI = "wss://" + shared.DefaultPublicHost + gs.DefaultGatewayServerConfig.BasicStation.ListenTLS
}

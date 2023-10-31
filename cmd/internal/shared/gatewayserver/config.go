package shared

import (
	"fmt"
	"time"

	"go.thethings.network/lorawan-stack/v3/cmd/internal/shared"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/udp"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/ws"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/upstream/packetbroker"
)

// DefaultGatewayServerConfig is the default configuration for the GatewayServer.
var DefaultGatewayServerConfig = gatewayserver.Config{
	RequireRegisteredGateways:         false,
	FetchGatewayInterval:              10 * time.Minute,
	FetchGatewayJitter:                0.2,
	UpdateGatewayLocationDebounceTime: time.Hour,
	UpdateConnectionStatsInterval:     time.Minute,
	UpdateConnectionStatsDebounceTime: 30 * time.Second,
	ConnectionStatsTTL:                12 * time.Hour,
	ConnectionStatsDisconnectTTL:      48 * time.Hour,
	UpdateVersionInfoDelay:            5 * time.Second,
	Forward: map[string][]string{
		"": {"00000000/0"},
	},
	PacketBroker: gatewayserver.PacketBrokerConfig{
		UpdateGatewayInterval: packetbroker.DefaultUpdateGatewayInterval,
		UpdateGatewayJitter:   packetbroker.DefaultUpdateGatewayJitter,
		OnlineTTLMargin:       packetbroker.DefaultOnlineTTLMargin,
	},
	UDP: gatewayserver.UDPConfig{
		Config: udp.DefaultConfig,
		Listeners: map[string]string{
			":1700": "",
		},
	},
	MQTTV2: config.MQTT{
		Listen:           ":1881",
		ListenTLS:        ":8881",
		PublicAddress:    fmt.Sprintf("%s:1881", shared.DefaultPublicHost),
		PublicTLSAddress: fmt.Sprintf("%s:8881", shared.DefaultPublicHost),
	},
	MQTT: config.MQTT{
		Listen:           ":1882",
		ListenTLS:        ":8882",
		PublicAddress:    fmt.Sprintf("%s:1882", shared.DefaultPublicHost),
		PublicTLSAddress: fmt.Sprintf("%s:8882", shared.DefaultPublicHost),
	},
	BasicStation: gatewayserver.BasicStationConfig{
		Config:                 ws.DefaultConfig,
		MaxValidRoundTripDelay: 10 * time.Second,
		Listen:                 ":1887",
		ListenTLS:              ":8887",
	},
}

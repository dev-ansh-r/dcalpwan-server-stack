package shared

import (
	"go.thethings.network/lorawan-stack/v3/pkg/packetbrokeragent"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// DefaultPacketBrokerAgentConfig is the default configuration for the Packet Broker Agent.
var DefaultPacketBrokerAgentConfig = packetbrokeragent.Config{
	IAMAddress:          "iam.packetbroker.net:443",
	ControlPlaneAddress: "cp.packetbroker.net:443",
	MapperAddress:       "mapper.packetbroker.net:443",
	AuthenticationMode:  "oauth2",
	OAuth2: packetbrokeragent.OAuth2Config{
		TokenURL: "https://iam.packetbroker.net/token",
	},
	Registration: packetbrokeragent.RegistrationConfig{
		Listed: true,
	},
	HomeNetwork: packetbrokeragent.HomeNetworkConfig{
		WorkerPool: packetbrokeragent.WorkerPoolConfig{
			Limit: 4096,
		},
		IncludeHops:     false,
		DevAddrPrefixes: []types.DevAddrPrefix{{}}, // Subscribe to all DevAddr prefixes.
	},
	Forwarder: packetbrokeragent.ForwarderConfig{
		WorkerPool: packetbrokeragent.WorkerPoolConfig{
			Limit: 1024,
		},
		IncludeGatewayEUI: true,
		IncludeGatewayID:  true,
		HashGatewayID:     false,
		GatewayOnlineTTL:  packetbrokeragent.DefaultGatewayOnlineTTL,
	},
}

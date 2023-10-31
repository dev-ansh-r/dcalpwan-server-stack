package joinserver

import (
	"go.thethings.network/lorawan-stack/v3/pkg/joinserver"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// DefaultJoinServerConfig is the default configuration for the JoinServer
var DefaultJoinServerConfig = joinserver.Config{
	JoinEUIPrefixes: []types.EUI64Prefix{
		{},
	},
	DevNonceLimit:   10,
	SessionKeyLimit: 10,
}

package deviceclaimingserver

import (
	"go.thethings.network/lorawan-stack/v3/pkg/deviceclaimingserver/enddevices"
)

// Config is the configuration for the Device Claiming Server.
type Config struct {
	EndDeviceClaimingServerConfig enddevices.Config `name:"edcs"`
}

package gatewayconfigurationserver

import (
	bscups "go.thethings.network/lorawan-stack/v3/pkg/basicstation/cups"
	gcsv2 "go.thethings.network/lorawan-stack/v3/pkg/gatewayconfigurationserver/v2"
)

// Config contains the Gateway Configuration Server configuration.
type Config struct {
	// BasicStation defines the configuration for the BasicStation CUPS.
	BasicStation bscups.ServerConfig `name:"basic-station" description:"BasicStation CUPS configuration."`
	// TheThingsGateway defines the configuration for The Things Gateway CUPS.
	TheThingsGateway gcsv2.TheThingsGatewayConfig `name:"the-things-gateway" description:"The Things Gateway CUPS configuration."`
	// RequreAuth defines if the HTTP endpoints should require authentication or not.
	RequireAuth bool `name:"require-auth" description:"Require authentication for the HTTP endpoints."`
}

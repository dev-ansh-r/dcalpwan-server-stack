package ws

import "time"

// Config defines the LoRa Basics Station configuration of the Gateway Server.
type Config struct {
	UseTrafficTLSAddress bool          `name:"use-traffic-tls-address" description:"Use WSS for the traffic address regardless of the TLS setting"`
	WSPingInterval       time.Duration `name:"ws-ping-interval" description:"Interval to send WS ping messages"`
	MissedPongThreshold  int           `name:"missed-pong-threshold" description:"Number of consecutive missed pongs before disconnection. This value is used only if the gateway sends at least one pong."`
	TimeSyncInterval     time.Duration `name:"time-sync-interval" description:"Interval to send time transfer messages"`
	AllowUnauthenticated bool          `name:"allow-unauthenticated" description:"Allow unauthenticated connections"`
}

// DefaultConfig contains the default configuration.
var DefaultConfig = Config{
	UseTrafficTLSAddress: false,
	WSPingInterval:       30 * time.Second,
	MissedPongThreshold:  2,
	TimeSyncInterval:     200 * time.Second,
	AllowUnauthenticated: false,
}

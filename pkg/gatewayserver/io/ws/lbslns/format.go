package lbslns

import (
	"fmt"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/ws"
)

var (
	errSessionStateNotFound = errors.DefineUnavailable("session_state_not_found", "session state not found")
	trafficEndPointPrefix   = "/traffic"
)

type lbsLNS struct {
	maxRoundTripDelay time.Duration
	tokens            io.DownlinkTokens
}

// NewFormatter returns a new LoRa Basic Station LNS formatter.
func NewFormatter(maxRoundTripDelay time.Duration) ws.Formatter {
	return &lbsLNS{
		maxRoundTripDelay: maxRoundTripDelay,
	}
}

func (f *lbsLNS) Endpoints() ws.Endpoints {
	return ws.Endpoints{
		ConnectionInfo: "/router-info",
		Traffic:        fmt.Sprintf("%s/{id}", trafficEndPointPrefix),
	}
}

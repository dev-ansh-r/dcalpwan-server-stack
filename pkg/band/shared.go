package band

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// eirpDelta is the delta between EIRP and ERP.
const eirpDelta = 2.15

// SharedParameters contains the properties which are shared by multiple bands within the
// same Regional Parameters version.
type SharedParameters struct {
	// ReceiveDelay1 is the default Rx1 window timing in seconds.
	ReceiveDelay1 time.Duration
	// ReceiveDelay2 is the default Rx2 window timing in seconds (ReceiveDelay1 + 1s).
	ReceiveDelay2 time.Duration
	// ReceiveDelay1 is the default join-accept window timing in seconds.
	JoinAcceptDelay1 time.Duration
	// ReceiveDelay2 is the default join-accept second window timing in seconds.
	JoinAcceptDelay2 time.Duration
	// MaxFCntGap is the maximum allowed frame counter gap between two uplink messages.
	MaxFCntGap uint
	// ADRAckLimit is the default ADR acknowledgement limit.
	ADRAckLimit ttnpb.ADRAckLimitExponent
	// ADRAckDelay is the default ADR acknowledgement delay.
	ADRAckDelay ttnpb.ADRAckDelayExponent
	// MinRetransmitTimeout is the minimum retransmit timeout.
	MinRetransmitTimeout time.Duration
	// MaxRetransmitTimeout is the maximum retransmit timeout.
	MaxRetransmitTimeout time.Duration

	// RelayForwardDelay is the default delay between the end of the uplink transmission and the start of the
	// relay transmission.
	RelayForwardDelay time.Duration
	// RelayReceiveDelay is the default RxR window timing in seconds.
	RelayReceiveDelay time.Duration
}

var (
	universalSharedParameters = SharedParameters{
		ReceiveDelay1:        time.Second,
		ReceiveDelay2:        2 * time.Second,
		JoinAcceptDelay1:     5 * time.Second,
		JoinAcceptDelay2:     6 * time.Second,
		MaxFCntGap:           16384,
		ADRAckLimit:          ttnpb.ADRAckLimitExponent_ADR_ACK_LIMIT_64,
		ADRAckDelay:          ttnpb.ADRAckDelayExponent_ADR_ACK_DELAY_32,
		MinRetransmitTimeout: time.Second,
		MaxRetransmitTimeout: 3 * time.Second,
	}
	relayAwareSharedParameters = func() SharedParameters {
		parameters := universalSharedParameters
		parameters.RelayForwardDelay = 50 * time.Millisecond
		parameters.RelayReceiveDelay = 18 * time.Second
		return parameters
	}()
)

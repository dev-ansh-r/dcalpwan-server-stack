package udp

import (
	"sync"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	encoding "go.thethings.network/lorawan-stack/v3/pkg/ttnpb/udp"
)

type rateLimitingFirewall struct {
	f Firewall

	m sync.Map // string to timestamps

	messages  int
	threshold time.Duration
}

// NewRateLimitingFirewall returns a Firewall with rate limiting capabilities.
func NewRateLimitingFirewall(firewall Firewall, messages int, threshold time.Duration) Firewall {
	return &rateLimitingFirewall{
		f:         firewall,
		messages:  messages,
		threshold: threshold,
	}
}

var errRateExceeded = errors.DefineResourceExhausted("rate_exceeded", "gateway traffic exceeded allowed rate")

func (f *rateLimitingFirewall) Filter(packet encoding.Packet) error {
	if packet.GatewayEUI == nil {
		return errNoEUI.New()
	}
	if packet.GatewayAddr == nil {
		return errNoAddress.New()
	}
	now := time.Now().UTC()
	eui := *packet.GatewayEUI
	ts := newTimestamps(f.messages)
	if val, ok := f.m.LoadOrStore(eui, ts); ok {
		ts = val.(*timestamps)
	}

	oldestTimestamp := ts.Append(now)
	if !oldestTimestamp.IsZero() && now.Sub(oldestTimestamp) < f.threshold {
		return errRateExceeded.New()
	}

	// Continue filtering
	if f.f != nil {
		return f.f.Filter(packet)
	}
	return nil
}

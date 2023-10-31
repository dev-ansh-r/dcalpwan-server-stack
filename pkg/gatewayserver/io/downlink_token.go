package io

import (
	"fmt"
	"regexp"
	"strconv"
	"sync/atomic"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

const downlinkTokenItems = 1 << 4

type downlinkToken struct {
	key  uint16
	msg  *ttnpb.DownlinkMessage
	time time.Time
}

// DownlinkTokens stores a set of downlink tokens and can be used to track roundtrip time.
// The number of downlink tokens stored is fixed to 16. New issued tokens with `Next` overwrite the oldest token.
type DownlinkTokens struct {
	last  uint32
	items [downlinkTokenItems]downlinkToken
}

// Next returns a new downlink token for a downlink message.
func (t *DownlinkTokens) Next(msg *ttnpb.DownlinkMessage, time time.Time) uint16 {
	key := uint16(atomic.AddUint32(&t.last, 1))
	pos := key % downlinkTokenItems

	t.items[pos] = downlinkToken{
		key:  key,
		msg:  msg,
		time: time,
	}
	return key
}

// Get returns the correlation IDs and time difference between the time given to `Next` and the given time by the token.
// If the token could not be found, this method returns false.
func (t DownlinkTokens) Get(token uint16, time time.Time) (*ttnpb.DownlinkMessage, time.Duration, bool) {
	pos := token % downlinkTokenItems
	item := t.items[pos]
	if item.key != token || item.msg == nil || item.time.IsZero() {
		return nil, 0, false
	}
	return ttnpb.Clone(item.msg), time.Sub(item.time), true
}

var parseTokenRegex = regexp.MustCompile(`^gs:down:token:(\d+)$`)

// FormatCorrelationID formats a correlation ID for a downlink token.
func (t DownlinkTokens) FormatCorrelationID(token uint16) string {
	return fmt.Sprintf("gs:down:token:%d", token)
}

// ParseTokenFromCorrelationIDs parses the correlation ID
func (t DownlinkTokens) ParseTokenFromCorrelationIDs(cids []string) (uint16, bool) {
	for _, cid := range cids {
		matches := parseTokenRegex.FindStringSubmatch(cid)
		if len(matches) != 2 || matches[1] == "" {
			continue
		}
		token, err := strconv.ParseUint(matches[1], 10, 16)
		if err != nil {
			return 0, false
		}
		return uint16(token), true
	}
	return 0, false
}

package networkserver_test

import (
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/internal/test/shared"
)

func init() {
	NewApplicationUplinkQueue = NewRedisApplicationUplinkQueue
	NewDeviceRegistry = NewRedisDeviceRegistry
	NewDownlinkTaskQueue = NewRedisDownlinkTaskQueue
	NewUplinkDeduplicator = NewRedisUplinkDeduplicator
	NewScheduledDownlinkMatcher = NewRedisScheduledDownlinkMatcher
}

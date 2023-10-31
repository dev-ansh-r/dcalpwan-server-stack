package redis_test

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/networkserver"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/internal/test/shared"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var _ networkserver.DeviceRegistry = &DeviceRegistry{}

func TestDeviceRegistry(t *testing.T) {
	_, ctx := test.New(t)
	reg, closeFn := NewRedisDeviceRegistry(ctx)
	defer closeFn()
	HandleDeviceRegistryTest(t, reg)
}

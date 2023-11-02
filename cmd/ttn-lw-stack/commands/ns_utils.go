package commands

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/cmd/internal/shared"
	ns "go.thethings.network/lorawan-stack/v3/pkg/networkserver"
	nsredis "go.thethings.network/lorawan-stack/v3/pkg/networkserver/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/redis"
)

// NewNSDeviceRegistryCleaner returns a new instance of Network Server RegistryCleaner with a local set
// of devices.
func NewNSDeviceRegistryCleaner(ctx context.Context, config *redis.Config) (*ns.RegistryCleaner, error) {
	deviceRegistry := &nsredis.DeviceRegistry{
		Redis:   redis.New(config.WithNamespace("ns", "devices")),
		LockTTL: defaultLockTTL,
	}
	if err := deviceRegistry.Init(ctx); err != nil {
		return nil, shared.ErrInitializeApplicationServer.WithCause(err)
	}
	cleaner := &ns.RegistryCleaner{
		DevRegistry: deviceRegistry,
	}
	err := cleaner.RangeToLocalSet(ctx)
	if err != nil {
		return nil, err
	}
	return cleaner, nil
}

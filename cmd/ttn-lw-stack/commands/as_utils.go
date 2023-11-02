package commands

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/cmd/internal/shared"
	as "go.thethings.network/lorawan-stack/v3/pkg/applicationserver"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages"
	packageredis "go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/pubsub"
	pubsubredis "go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/pubsub/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/web"
	webredis "go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/web/redis"
	asredis "go.thethings.network/lorawan-stack/v3/pkg/applicationserver/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/redis"
)

// NewPubSubCleaner returns a new instance of pubsub RegistryCleaner with a local set of applications.
func NewPubSubCleaner(ctx context.Context, config *redis.Config) (*pubsub.RegistryCleaner, error) {
	pubSubRegistry := &pubsubredis.PubSubRegistry{
		Redis:   redis.New(config.WithNamespace("as", "io", "pubsub")),
		LockTTL: defaultLockTTL,
	}
	if err := pubSubRegistry.Init(ctx); err != nil {
		return nil, shared.ErrInitializeApplicationServer.WithCause(err)
	}
	cleaner := &pubsub.RegistryCleaner{
		PubSubRegistry: pubSubRegistry,
	}
	err := cleaner.RangeToLocalSet(ctx)
	if err != nil {
		return nil, err
	}
	return cleaner, nil
}

// NewPackagesCleaner returns a new instance of packages RegistryCleaner with a local set
// of applications and devices.
func NewPackagesCleaner(ctx context.Context, config *redis.Config) (*packages.RegistryCleaner, error) {
	applicationPackagesRegistry, err := packageredis.NewApplicationPackagesRegistry(
		ctx, redis.New(config.WithNamespace("as", "io", "applicationpackages")), defaultLockTTL,
	)
	if err != nil {
		return nil, shared.ErrInitializeApplicationServer.WithCause(err)
	}
	cleaner := &packages.RegistryCleaner{
		ApplicationPackagesRegistry: applicationPackagesRegistry,
	}
	if err := cleaner.RangeToLocalSet(ctx); err != nil {
		return nil, err
	}
	return cleaner, nil
}

// NewASDeviceRegistryCleaner returns a new instance of device RegistryCleaner with a local set
// of devices.
func NewASDeviceRegistryCleaner(ctx context.Context, config *redis.Config) (*as.RegistryCleaner, error) {
	deviceRegistry := &asredis.DeviceRegistry{
		Redis:   redis.New(config.WithNamespace("as", "devices")),
		LockTTL: defaultLockTTL,
	}
	if err := deviceRegistry.Init(ctx); err != nil {
		return nil, shared.ErrInitializeApplicationServer.WithCause(err)
	}
	cleaner := &as.RegistryCleaner{
		DevRegistry: deviceRegistry,
	}
	err := cleaner.RangeToLocalSet(ctx)
	if err != nil {
		return nil, err
	}
	return cleaner, nil
}

// NewWebhookCleaner returns a new instance of webhook RegistryCleaner with a local set
// of applications.
func NewWebhookCleaner(ctx context.Context, config *redis.Config) (*web.RegistryCleaner, error) {
	webhookRegistry := &webredis.WebhookRegistry{
		Redis:   redis.New(config.WithNamespace("as", "io", "webhooks")),
		LockTTL: defaultLockTTL,
	}
	if err := webhookRegistry.Init(ctx); err != nil {
		return nil, shared.ErrInitializeApplicationServer.WithCause(err)
	}
	cleaner := &web.RegistryCleaner{
		WebRegistry: webhookRegistry,
	}
	err := cleaner.RangeToLocalSet(ctx)
	if err != nil {
		return nil, err
	}
	return cleaner, nil
}

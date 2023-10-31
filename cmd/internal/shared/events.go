package shared

import (
	"context"
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/events/basic"
	"go.thethings.network/lorawan-stack/v3/pkg/events/cloud"
	"go.thethings.network/lorawan-stack/v3/pkg/events/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/task"
	_ "gocloud.dev/pubsub/awssnssqs" // AWS backend for PubSub.
	_ "gocloud.dev/pubsub/gcppubsub" // GCP backend for PubSub.
)

// Component contains a minimal component.Component definition.
type Component interface {
	task.Starter
	FromRequestContext(context.Context) context.Context
}

// InitializeEvents initializes the event system.
func InitializeEvents(ctx context.Context, component Component, conf config.ServiceBase) error {
	switch conf.Events.Backend {
	case "internal":
		events.SetDefaultPubSub(basic.NewPubSub())
		return nil
	case "redis":
		events.SetDefaultPubSub(redis.NewPubSub(ctx, component, conf.Events.Redis, conf.Events.Batch))
		return nil
	case "cloud":
		ps, err := cloud.NewPubSub(ctx, component, conf.Events.Cloud.PublishURL, conf.Events.Cloud.SubscribeURL)
		if err != nil {
			return err
		}
		events.SetDefaultPubSub(ps)
		return nil
	default:
		return fmt.Errorf("unknown events backend: %s", conf.Events.Backend)
	}
}

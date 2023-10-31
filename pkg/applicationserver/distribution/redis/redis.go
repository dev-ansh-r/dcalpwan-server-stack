package redis

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	ttnredis "go.thethings.network/lorawan-stack/v3/pkg/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

// PubSub is a Redis-based upstream traffic Pub/Sub.
type PubSub struct {
	Redis *ttnredis.Client
}

func (ps PubSub) uidUplinkKey(uid string) string {
	return ps.Redis.Key("uid", uid, "uplinks")
}

// Publish publishes the uplink to Pub/Sub.
func (ps PubSub) Publish(ctx context.Context, up *ttnpb.ApplicationUp) error {
	msg, err := ttnredis.MarshalProto(up)
	if err != nil {
		return err
	}
	uid := unique.ID(ctx, up.EndDeviceIds.ApplicationIds)
	if err = ps.Redis.Publish(ctx, ps.uidUplinkKey(uid), msg).Err(); err != nil {
		return ttnredis.ConvertError(err)
	}
	return nil
}

var errChannelClosed = errors.DefineAborted("channel_closed", "channel closed")

// Subscribe subscribes to the traffic of the provided application and processes it using the handler.
func (ps PubSub) Subscribe(
	ctx context.Context, ids *ttnpb.ApplicationIdentifiers, handler func(context.Context, *ttnpb.ApplicationUp) error,
) error {
	uid := unique.ID(ctx, ids)
	sub := ps.Redis.Subscribe(ctx, ps.uidUplinkKey(uid))
	defer sub.Close()

	// sub.Receive(.*) will not respect the asynchronous context
	// cancelation, only a context deadline, if present.
	// As such, we use the buffered channel instead and do the
	// asynchronous select here. This allows us to close the
	// subscription on context cancellation.
	ch := sub.Channel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg, ok := <-ch:
			if !ok {
				return errChannelClosed.New()
			}

			up := &ttnpb.ApplicationUp{}
			if err := ttnredis.UnmarshalProto(msg.Payload, up); err != nil {
				return err
			}

			if err := handler(ctx, up); err != nil {
				return err
			}
		}
	}
}

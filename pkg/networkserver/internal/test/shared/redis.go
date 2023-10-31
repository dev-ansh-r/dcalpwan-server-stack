package test

import (
	"context"
	"time"

	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver"
	"go.thethings.network/lorawan-stack/v3/pkg/networkserver/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var redisNamespace = [...]string{
	"networkserver_test",
}

const redisConsumerGroup = "ns"

func testStreamBlockLimit() time.Duration {
	return (1 << 5) * test.Delay
}

func NewRedisApplicationUplinkQueue(ctx context.Context) (ApplicationUplinkQueue, func()) {
	tb := test.MustTBFromContext(ctx)
	cl, flush := test.NewRedis(ctx, append(redisNamespace[:], "application-uplinks")...)
	q := redis.NewApplicationUplinkQueue(cl, 100, redisConsumerGroup, 0)
	if err := q.Init(ctx); err != nil {
		tb.Fatalf("Failed to initialize Redis application uplink queue: %s", test.FormatError(err))
	}
	return q,
		func() {
			if err := q.Close(ctx); err != nil {
				tb.Errorf("Failed to close Redis application uplink queue: %s", err)
			}
			flush()
			cl.Close()
		}
}

func NewRedisDeviceRegistry(ctx context.Context) (DeviceRegistry, func()) {
	tb := test.MustTBFromContext(ctx)
	cl, flush := test.NewRedis(ctx, append(redisNamespace[:], "devices")...)
	reg := &redis.DeviceRegistry{
		Redis:   cl,
		LockTTL: test.Delay << 10,
	}
	if err := reg.Init(ctx); err != nil {
		tb.Fatalf("Failed to initialize Redis device registry: %s", test.FormatError(err))
	}
	return reg,
		func() {
			flush()
			cl.Close()
		}
}

func NewRedisDownlinkTaskQueue(ctx context.Context) (DownlinkTaskQueue, func()) {
	tb := test.MustTBFromContext(ctx)
	cl, flush := test.NewRedis(ctx, append(redisNamespace[:], "downlink-tasks")...)
	q := redis.NewDownlinkTaskQueue(cl, 10000, redisConsumerGroup, testStreamBlockLimit())
	if err := q.Init(ctx); err != nil {
		tb.Fatalf("Failed to initialize Redis downlink task queue: %s", test.FormatError(err))
	}
	return q,
		func() {
			if err := q.Close(ctx); err != nil {
				tb.Errorf("Failed to close Redis downlink task queue: %s", test.FormatError(err))
			}
			flush()
			cl.Close()
		}
}

func NewRedisUplinkDeduplicator(ctx context.Context) (UplinkDeduplicator, func()) {
	cl, flush := test.NewRedis(ctx, append(redisNamespace[:], "uplink-deduplication")...)
	return &redis.UplinkDeduplicator{
			Redis: cl,
		},
		func() {
			flush()
			cl.Close()
		}
}

func NewRedisScheduledDownlinkMatcher(ctx context.Context) (ScheduledDownlinkMatcher, func()) {
	cl, flush := test.NewRedis(ctx, append(redisNamespace[:], "scheduled-downlink-matcher")...)
	return &redis.ScheduledDownlinkMatcher{
			Redis: cl,
		},
		func() {
			flush()
			cl.Close()
		}
}

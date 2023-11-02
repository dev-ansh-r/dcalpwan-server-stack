package redis_test

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/events/internal/eventstest"
	"go.thethings.network/lorawan-stack/v3/pkg/events/redis"
	ttnredis "go.thethings.network/lorawan-stack/v3/pkg/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/task"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var redisConfig = func() ttnredis.Config {
	var err error
	config := ttnredis.Config{
		Address:       "localhost:6379",
		Database:      1,
		RootNamespace: []string{"test"},
	}
	if address := os.Getenv("REDIS_ADDRESS"); address != "" {
		config.Address = address
	}
	if db := os.Getenv("REDIS_DB"); db != "" {
		config.Database, err = strconv.Atoi(db)
		if err != nil {
			panic(err)
		}
	}
	if prefix := os.Getenv("REDIS_PREFIX"); prefix != "" {
		config.RootNamespace = []string{prefix}
	}
	return config
}()

type mockComponent struct {
	task.Starter
}

func (mockComponent) FromRequestContext(ctx context.Context) context.Context {
	return ctx
}

func Example() {
	// The task starter is used for automatic re-subscription on failure.
	taskStarter := task.StartTaskFunc(task.DefaultStartTask)

	redisPubSub := redis.NewPubSub(context.TODO(), mockComponent{taskStarter}, config.RedisEvents{
		// Config here...
	}, config.BatchEvents{
		// Batch config here...
	})

	// Replace the default pubsub so that we will now publish to Redis.
	events.SetDefaultPubSub(redisPubSub)
}

var timeout = (1 << 11) * test.Delay

func TestRedisPubSub(t *testing.T) { //nolint:paralleltest
	events.IncludeCaller = true
	taskStarter := task.StartTaskFunc(task.DefaultStartTask)

	test.RunTest(t, test.TestConfig{
		Timeout: timeout,
		Func: func(ctx context.Context, a *assertions.Assertion) {
			conf := config.RedisEvents{
				Config: redisConfig,
			}
			batchConfig := config.BatchEvents{}
			pubsub := redis.NewPubSub(ctx, mockComponent{taskStarter}, conf, batchConfig)
			defer pubsub.(*redis.PubSub).Close(ctx)

			time.Sleep(timeout / 10)

			eventstest.TestBackend(ctx, t, a, pubsub)
		},
	})
}

func TestRedisPubSubStore(t *testing.T) { //nolint:paralleltest
	events.IncludeCaller = true
	taskStarter := task.StartTaskFunc(task.DefaultStartTask)

	test.RunTest(t, test.TestConfig{
		Timeout: timeout,
		Func: func(ctx context.Context, a *assertions.Assertion) {
			conf := config.RedisEvents{
				Config: redisConfig,
			}
			conf.Store.Enable = true
			batchConf := config.BatchEvents{Enable: true}
			pubsub := redis.NewPubSub(ctx, mockComponent{taskStarter}, conf, batchConf)
			defer pubsub.(*redis.PubSubStore).Close(ctx)

			time.Sleep(timeout / 10)

			eventstest.TestBackend(ctx, t, a, pubsub)
		},
	})
}

var _ events.Store = (*redis.PubSubStore)(nil)

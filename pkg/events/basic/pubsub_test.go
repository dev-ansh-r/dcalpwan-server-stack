package basic_test

import (
	"context"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/events/basic"
	"go.thethings.network/lorawan-stack/v3/pkg/events/internal/eventstest"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

func Example() {
	basicPubSub := basic.NewPubSub()

	// Replace the default pubsub so that we will now publish to this pub sub by default.
	events.SetDefaultPubSub(basicPubSub)
}

var timeout = (1 << 10) * test.Delay

func TestPubSub(t *testing.T) { //nolint:paralleltest
	events.IncludeCaller = true

	test.RunTest(t, test.TestConfig{
		Timeout: timeout,
		Func: func(ctx context.Context, a *assertions.Assertion) {
			pubsub := basic.NewPubSub()
			eventstest.TestBackend(ctx, t, a, pubsub)
		},
	})
}

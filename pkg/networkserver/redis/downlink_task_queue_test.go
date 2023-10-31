package redis_test

import (
	"fmt"
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/networkserver"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/internal/test/shared"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var _ networkserver.DownlinkTaskQueue = &DownlinkTaskQueue{}

func TestDownlinkTaskQueue(t *testing.T) {
	for _, consumers := range []int{1, 2, 4, 8} {
		t.Run(fmt.Sprintf("Consumers=%d", consumers), func(t *testing.T) {
			_, ctx := test.New(t)
			consumerIDs := make([]string, 0, consumers)
			for i := 0; i < consumers; i++ {
				consumerIDs = append(consumerIDs, fmt.Sprintf("consumer-%d-%d", consumers, i))
			}
			q, closeFn := NewRedisDownlinkTaskQueue(ctx)
			defer closeFn()
			HandleDownlinkTaskQueueTest(t, q, consumerIDs)
		})
	}
}

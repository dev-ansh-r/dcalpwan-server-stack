package pubsub

import (
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/task"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

func init() {
	io.DialTaskBackoffConfig.IntervalFunc = task.MakeBackoffIntervalFunc(false, task.DefaultBackoffResetDuration, (1<<5)*test.Delay)
}

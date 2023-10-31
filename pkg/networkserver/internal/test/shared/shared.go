// Package test contains testing utilities usable by all subpackages of networkserver excluding itself.
package test

import (
	"time"

	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var (
	CacheTTL           = (1 << 6) * test.Delay
	DefaultMACSettings = test.Must(DefaultConfig.DefaultMACSettings.Parse())
)

type TaskPopFuncResponse struct {
	Time  time.Time
	Error error
}

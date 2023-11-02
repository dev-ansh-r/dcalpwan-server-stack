package lora_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/lora"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestAdjustedRSSI(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)
	a.So(lora.AdjustedRSSI(0, 0), should.AlmostEqual, float32(-10.0/3.0))
	a.So(lora.AdjustedRSSI(-46, 7), should.Equal, -47.0)
	a.So(lora.AdjustedRSSI(-46, 11), should.Equal, -46.0)
	a.So(lora.AdjustedRSSI(-46, -6), should.Equal, -52.0)
}

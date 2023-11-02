package store

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestDutyCycleFromFloat(t *testing.T) {
	for _, tc := range []struct {
		Float float64
		Enum  ttnpb.AggregatedDutyCycle
	}{
		{
			Float: 1.0,
			Enum:  ttnpb.AggregatedDutyCycle_DUTY_CYCLE_1,
		},
		{
			Float: 0.5,
			Enum:  ttnpb.AggregatedDutyCycle_DUTY_CYCLE_2,
		},
		{
			Float: 0.25,
			Enum:  ttnpb.AggregatedDutyCycle_DUTY_CYCLE_4,
		},
		{
			Float: 0.14,
			Enum:  ttnpb.AggregatedDutyCycle_DUTY_CYCLE_8,
		},
		{
			Float: 1 / (2 << 20),
			Enum:  ttnpb.AggregatedDutyCycle_DUTY_CYCLE_32768,
		},
	} {
		a := assertions.New(t)
		a.So(dutyCycleFromFloat(tc.Float), should.Equal, tc.Enum)
	}
}

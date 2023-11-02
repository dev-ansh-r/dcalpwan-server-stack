package ws

import (
	"testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func timePtr(t time.Time) *time.Time { return &t }

func TestTimePtrFromUpInfo(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		Name          string
		GPSTime       int64
		RxTime        float64
		ReferenceTime time.Time
		ExpectedTime  *time.Time
	}{
		{
			Name:          "NoTimestamps",
			ReferenceTime: time.Unix(0, 456),
		},
		{
			Name:   "OnlyRxTime",
			RxTime: 123.456,

			ExpectedTime: timePtr(time.Unix(123, 456000000).UTC()),
		},
		{
			Name:    "EqualGPSTimeAndRxTime",
			GPSTime: -315964676544000, // microseconds. The timestamp is negative as the UTC epoch precedes the GPS epoch.
			RxTime:  123.456,

			ExpectedTime: timePtr(time.Unix(123, 456000000).UTC()),
		},
		{
			Name:    "OnlyGPSTime",
			GPSTime: 123456000, // microseconds.

			ExpectedTime: timePtr(time.Unix(315964923, 456000000).UTC()),
		},
	} {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			a, _ := test.New(t)
			a.So(TimePtrFromUpInfo(tc.GPSTime, tc.RxTime), should.Resemble, tc.ExpectedTime)
		})
	}
}

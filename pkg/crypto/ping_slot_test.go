package crypto_test

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/crypto"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestComputePingOffset(t *testing.T) {
	for _, tc := range []struct {
		BeaconTime uint32
		DevAddr    types.DevAddr
		PingPeriod uint16

		ExpectedPingOffset uint16
		ErrorAssertion     func(t *testing.T, err error) bool
	}{
		{
			ErrorAssertion: func(t *testing.T, err error) bool {
				return assertions.New(t).So(errors.IsInvalidArgument(err), should.BeTrue)
			},
		},
		{
			PingPeriod: 31,
			ErrorAssertion: func(t *testing.T, err error) bool {
				return assertions.New(t).So(errors.IsInvalidArgument(err), should.BeTrue)
			},
		},
		{
			PingPeriod: 4097,
			ErrorAssertion: func(t *testing.T, err error) bool {
				return assertions.New(t).So(errors.IsInvalidArgument(err), should.BeTrue)
			},
		},
		{
			PingPeriod:         32,
			ExpectedPingOffset: 6,
			ErrorAssertion: func(t *testing.T, err error) bool {
				return assertions.New(t).So(err, should.BeNil)
			},
		},
		{
			BeaconTime:         0xff42,
			DevAddr:            types.DevAddr{0x00, 0x42, 0x00, 0xff},
			PingPeriod:         4096,
			ExpectedPingOffset: 3994,
			ErrorAssertion: func(t *testing.T, err error) bool {
				return assertions.New(t).So(err, should.BeNil)
			},
		},
	} {
		t.Run(fmt.Sprintf("beacon_time(%d)/dev_addr(%s)/ping_period(%d)", tc.BeaconTime, tc.DevAddr, tc.PingPeriod), func(t *testing.T) {
			a := assertions.New(t)
			p, err := ComputePingOffset(tc.BeaconTime, tc.DevAddr, tc.PingPeriod)
			a.So(p, should.Equal, tc.ExpectedPingOffset)
			a.So(tc.ErrorAssertion(t, err), should.BeTrue)
		})
	}
}

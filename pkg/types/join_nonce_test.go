package types_test

import (
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestJoinNonce(t *testing.T) {
	for _, tc := range []struct {
		JoinNonce JoinNonce
		IsZero    bool
		String    string
	}{
		{
			JoinNonce{0x00, 0x00, 0x00},
			true,
			"000000",
		},
		{
			JoinNonce{0x20, 0x00, 0x2f},
			false,
			"20002F",
		},
		{
			JoinNonce{0x40, 0x00, 0xef},
			false,
			"4000EF",
		},
	} {
		a := assertions.New(t)

		a.So(tc.JoinNonce.IsZero(), should.Equal, tc.IsZero)
		a.So(tc.JoinNonce.String(), should.Equal, tc.String)
	}
}

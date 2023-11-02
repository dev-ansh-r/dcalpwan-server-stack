package macspec_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/specification/macspec"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestMACVersion(t *testing.T) {
	a := assertions.New(t)
	for v := range ttnpb.MACVersion_name {
		if v == int32(ttnpb.MACVersion_MAC_UNKNOWN) {
			continue
		}
		a.So(func() { macspec.MACVersion(ttnpb.MACVersion(v)) }, assertions.ShouldNotPanic)
	}
}

func TestCompareMACVersion(t *testing.T) {
	for _, tc := range []struct {
		A, B     ttnpb.MACVersion
		Expected int
		Panics   bool
	}{
		{
			A:        ttnpb.MACVersion_MAC_V1_0,
			B:        ttnpb.MACVersion_MAC_V1_0_1,
			Expected: -1,
		},
		{
			A:        ttnpb.MACVersion_MAC_V1_1,
			B:        ttnpb.MACVersion_MAC_V1_0,
			Expected: 1,
		},
		{
			A:        ttnpb.MACVersion_MAC_V1_1,
			B:        ttnpb.MACVersion_MAC_V1_1,
			Expected: 0,
		},
		{
			A:        ttnpb.MACVersion_MAC_V1_0_2,
			B:        ttnpb.MACVersion_MAC_V1_1,
			Expected: -1,
		},
		{
			A:      ttnpb.MACVersion_MAC_UNKNOWN,
			B:      ttnpb.MACVersion_MAC_V1_1,
			Panics: true,
		},
		{
			A:      ttnpb.MACVersion_MAC_UNKNOWN,
			B:      ttnpb.MACVersion_MAC_UNKNOWN,
			Panics: true,
		},
		{
			A:      ttnpb.MACVersion_MAC_V1_0,
			B:      ttnpb.MACVersion_MAC_UNKNOWN,
			Panics: true,
		},
	} {
		a := assertions.New(t)

		if tc.Panics {
			a.So(func() { macspec.CompareMACVersion(tc.A, tc.B) }, should.Panic)
			return
		}

		a.So(macspec.CompareMACVersion(tc.A, tc.B), should.Equal, tc.Expected)
		if tc.A != tc.B {
			a.So(macspec.CompareMACVersion(tc.B, tc.A), should.Equal, -tc.Expected)
		}
	}
}

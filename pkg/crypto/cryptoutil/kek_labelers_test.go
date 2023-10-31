package cryptoutil_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/crypto/cryptoutil"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestComponentPrefixKEKLabeler(t *testing.T) {
	t.Parallel()
	for i, tc := range []struct {
		Separator     string
		ReplaceOldNew []string
		Addr          string
		Func          func(context.Context, ComponentPrefixKEKLabeler, string) string
		Expected      string
	}{
		{
			Addr: "",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, nil, addr)
			},
			Expected: "ns",
		},
		{
			Addr: "",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, &types.NetID{0x00, 0x00, 0x42}, addr)
			},
			Expected: "ns/000042",
		},
		{
			Addr: "localhost",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, nil, addr)
			},
			Expected: "ns/localhost",
		},
		{
			Addr: "localhost",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, &types.NetID{0x00, 0x00, 0x42}, addr)
			},
			Expected: "ns/000042/localhost",
		},
		{
			Addr: "localhost",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.AsKEKLabel(ctx, addr)
			},
			Expected: "as/localhost",
		},
		{
			Addr: "localhost:1234",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, &types.NetID{0x00, 0x00, 0x42}, addr)
			},
			Expected: "ns/000042/localhost",
		},
		{
			Addr: "http://localhost",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, &types.NetID{0x00, 0x00, 0x42}, addr)
			},
			Expected: "ns/000042/localhost",
		},
		{
			Addr: "http://localhost:1234",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, &types.NetID{0x00, 0x00, 0x42}, addr)
			},
			Expected: "ns/000042/localhost",
		},
		{
			ReplaceOldNew: []string{":", "_"},
			Addr:          "http://[::1]:1234",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, &types.NetID{0x00, 0x00, 0x42}, addr)
			},
			Expected: "ns/000042/__1",
		},
		{
			Separator: "_",
			Addr:      "http://localhost:1234",
			Func: func(ctx context.Context, labeler ComponentPrefixKEKLabeler, addr string) string {
				return labeler.NsKEKLabel(ctx, &types.NetID{0x00, 0x00, 0x42}, addr)
			},
			Expected: "ns_000042_localhost",
		},
	} {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			a := assertions.New(t)
			labeler := ComponentPrefixKEKLabeler{
				Separator:     tc.Separator,
				ReplaceOldNew: tc.ReplaceOldNew,
			}
			label := tc.Func(test.Context(), labeler, tc.Addr)
			a.So(label, should.Equal, tc.Expected)
		})
	}
}

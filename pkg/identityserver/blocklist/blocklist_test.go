package blocklist_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/blocklist"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestBlocklist(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)

	b := blocklist.New("foo", "bar")
	bs := blocklist.Blocklists{b}

	a.So(bs.Contains("foo"), should.BeTrue)
	a.So(bs.Contains("baz"), should.BeFalse)
}

func TestGlobalBlocklist(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)

	b := blocklist.New("foo", "bar")

	ctx := blocklist.NewContext(test.Context(), b)

	a.So(blocklist.Check(ctx, "root"), should.NotBeNil)
	a.So(blocklist.Check(ctx, "foo"), should.NotBeNil)
	a.So(blocklist.Check(ctx, "foobar"), should.BeNil)
}

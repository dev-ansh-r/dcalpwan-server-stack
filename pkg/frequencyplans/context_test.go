package frequencyplans_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestContext(t *testing.T) {
	a := assertions.New(t)

	ctx := test.Context()

	// No fallback.
	{
		_, ok := frequencyplans.FallbackIDFromContext(ctx)
		a.So(ok, should.BeFalse)
	}

	ctx = frequencyplans.WithFallbackID(ctx, "EU_863_870")

	// Have fallback.
	{
		id, ok := frequencyplans.FallbackIDFromContext(ctx)
		a.So(ok, should.BeTrue)
		a.So(id, should.Equal, "EU_863_870")
	}
}

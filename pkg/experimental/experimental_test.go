package experimental_test

import (
	"context"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/experimental"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestExperimentalFeatures(t *testing.T) {
	a := assertions.New(t)

	r := NewRegistry()

	ctx := NewContextWithRegistry(context.Background(), r)

	feature := DefineFeature("experimental.feature", false)
	a.So(feature.GetValue(ctx), should.BeFalse)
	a.So(AllFeatures(ctx), should.Resemble, map[string]bool{"experimental.feature": false})
	a.So(feature.GetValue(context.Background()), should.BeFalse)
	a.So(AllFeatures(context.Background()), should.Resemble, map[string]bool{"experimental.feature": false})

	r.EnableFeatures("experimental.feature")
	a.So(feature.GetValue(ctx), should.BeTrue)
	a.So(AllFeatures(ctx), should.Resemble, map[string]bool{"experimental.feature": true})
	a.So(feature.GetValue(context.Background()), should.BeFalse)
	a.So(AllFeatures(context.Background()), should.Resemble, map[string]bool{"experimental.feature": false})

	EnableFeatures("experimental.feature")
	r.DisableFeatures("experimental.feature")

	a.So(feature.GetValue(ctx), should.BeFalse)
	a.So(AllFeatures(ctx), should.Resemble, map[string]bool{"experimental.feature": false})
	a.So(feature.GetValue(context.Background()), should.BeTrue)
	a.So(AllFeatures(context.Background()), should.Resemble, map[string]bool{"experimental.feature": true})
}

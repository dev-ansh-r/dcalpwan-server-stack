package test_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestFrequencyPlans(t *testing.T) {
	a := assertions.New(t)

	fp := frequencyplans.NewStore(test.FrequencyPlansFetcher)

	euFP, err := fp.GetByID(test.EUFrequencyPlanID)
	a.So(err, should.BeNil)
	a.So(euFP.BandID, should.Equal, "EU_863_870")

	krFP, err := fp.GetByID(test.KRFrequencyPlanID)
	a.So(err, should.BeNil)
	a.So(krFP.UplinkChannels[0].Frequency, should.Equal, 922100000)

	_, err = fp.GetByID(test.ExampleFrequencyPlanID)
	a.So(err, should.BeNil)
}

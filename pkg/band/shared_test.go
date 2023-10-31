package band_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestRelaySharedParameters(t *testing.T) {
	t.Parallel()
	for name, versions := range band.All {
		name := name
		for version, phy := range versions {
			version, phy := version, phy
			t.Run(fmt.Sprintf("%v/%v", name, version), func(t *testing.T) {
				t.Parallel()
				a := assertions.New(t)
				var expectedForwardDelay, expectedReceiveDelay time.Duration
				switch {
				case version == ttnpb.PHYVersion_RP002_V1_0_4,
					name == band.ISM_2400,
					name == band.MA_869_870_DRAFT:
					expectedForwardDelay, expectedReceiveDelay = 50*time.Millisecond, 18*time.Second
				}
				a.So(phy.RelayForwardDelay, should.Equal, expectedForwardDelay)
				a.So(phy.RelayReceiveDelay, should.Equal, expectedReceiveDelay)
			})
		}
	}
}

package gatewayserver_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestConfig(t *testing.T) {
	a := assertions.New(t)

	{
		conf := gatewayserver.Config{
			Forward: map[string][]string{
				"":                {"00000000/0"},
				"packetbroker.io": {"00000000/3", "26000000/7"},
			},
		}
		forward, err := conf.ForwardDevAddrPrefixes()
		a.So(err, should.BeNil)
		a.So(forward, should.HaveEmptyDiff, map[string][]types.DevAddrPrefix{
			"": {{}},
			"packetbroker.io": {
				{DevAddr: types.DevAddr{}, Length: 3},
				{DevAddr: types.DevAddr{0x26, 0x0, 0x0, 0x0}, Length: 7},
			},
		})
	}

	{
		conf := gatewayserver.Config{
			Forward: map[string][]string{
				"packetbroker.io": {"00000000/3", "invalid"},
			},
		}
		_, err := conf.ForwardDevAddrPrefixes()
		a.So(err, should.NotBeNil)
	}
}

package topics_test

import (
	"testing"

	"github.com/TheThingsIndustries/mystique/pkg/topic"
	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/mqtt/topics"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

const gatewayIDV2 = "test"

func TestV2Topics(t *testing.T) {
	ctx := test.Context()
	v2 := topics.NewV2(ctx)
	uid := unique.ID(ctx, &ttnpb.GatewayIdentifiers{GatewayId: gatewayIDV2})
	for _, tc := range []struct {
		UID      string
		Func     func(string) []string
		Expected []string
		Is       func([]string) bool
		IsNot    []func([]string) bool
	}{
		{
			UID:      uid,
			Func:     v2.UplinkTopic,
			Expected: []string{uid, "up"},
			Is:       v2.IsUplinkTopic,
			IsNot:    []func([]string) bool{v2.IsStatusTopic, v2.IsTxAckTopic},
		},
		{
			UID:      uid,
			Func:     v2.StatusTopic,
			Expected: []string{uid, "status"},
			Is:       v2.IsStatusTopic,
			IsNot:    []func([]string) bool{v2.IsUplinkTopic, v2.IsTxAckTopic},
		},
	} {
		t.Run(topic.Join(tc.Expected), func(t *testing.T) {
			a := assertions.New(t)
			actual := tc.Func(tc.UID)
			a.So(actual, should.Resemble, tc.Expected)
			a.So(tc.Is(actual), should.BeTrue)
			for _, isNot := range tc.IsNot {
				a.So(isNot(actual), should.BeFalse)
			}
		})
	}
}

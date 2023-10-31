package topics_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/TheThingsIndustries/mystique/pkg/topic"
	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/mqtt/topics"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestV3AcceptedTopic(t *testing.T) {
	t.Parallel()
	uid := unique.ID(test.Context(), &ttnpb.ApplicationIdentifiers{ApplicationId: "foo-app"})
	for i, tc := range []struct {
		Requested,
		Accepted string
		OK bool
	}{
		{
			Requested: "v3",
		},
		{
			Requested: "+",
		},
		{
			Requested: "#",
			Accepted:  fmt.Sprintf("v3/%s/#", uid),
			OK:        true,
		},
		{
			Requested: "v3/#",
			Accepted:  fmt.Sprintf("v3/%s/#", uid),
			OK:        true,
		},
		{
			Requested: "v3/+/uplink",
			Accepted:  fmt.Sprintf("v3/%s/uplink", uid),
			OK:        true,
		},
		{
			Requested: fmt.Sprintf("v3/%s/uplink", uid),
			Accepted:  fmt.Sprintf("v3/%s/uplink", uid),
			OK:        true,
		},
	} {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			a := assertions.New(t)
			actual, ok := topics.Default.AcceptedTopic(uid, topic.Split(tc.Requested))
			if !a.So(ok, should.Equal, tc.OK) {
				t.FailNow()
			}
			a.So(topic.Join(actual), should.Equal, tc.Accepted)
		})
	}
}

func TestV3Topics(t *testing.T) {
	t.Parallel()
	appUID := unique.ID(test.Context(), &ttnpb.ApplicationIdentifiers{ApplicationId: "foo-app"})
	devID := "foo-device"

	for _, tc := range []struct {
		Fn       func(applicationUID, deviceUID string) []string
		Expected string
	}{
		{
			Fn:       topics.Default.UplinkMessageTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/up", appUID, devID),
		},
		{
			Fn:       topics.Default.UplinkNormalizedTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/up/normalized", appUID, devID),
		},
		{
			Fn:       topics.Default.JoinAcceptTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/join", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkAckTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/ack", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkNackTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/nack", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkSentTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/sent", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkFailedTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/failed", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkQueuedTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/queued", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkQueueInvalidatedTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/invalidated", appUID, devID),
		},
		{
			Fn:       topics.Default.LocationSolvedTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/location/solved", appUID, devID),
		},
		{
			Fn:       topics.Default.ServiceDataTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/service/data", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkPushTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/push", appUID, devID),
		},
		{
			Fn:       topics.Default.DownlinkReplaceTopic,
			Expected: fmt.Sprintf("v3/%s/devices/%s/down/replace", appUID, devID),
		},
	} {
		tc := tc
		t.Run(tc.Expected, func(t *testing.T) {
			t.Parallel()
			actual := strings.Join(tc.Fn(appUID, devID), "/")
			assertions.New(t).So(actual, should.Equal, tc.Expected)
		})
	}
}

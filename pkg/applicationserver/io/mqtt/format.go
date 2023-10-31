package mqtt

import (
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/formatters"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/mqtt/topics"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

// Format represents a topic layout and message formatter.
type Format interface {
	topics.Layout
	formatters.Formatter
}

// TopicParts generates the topic parts for the provided uplink.
func TopicParts(up *io.ContextualApplicationUp, layout topics.Layout) []string {
	var f func(string, string) []string
	switch up.Up.(type) {
	case *ttnpb.ApplicationUp_UplinkMessage:
		f = layout.UplinkMessageTopic
	case *ttnpb.ApplicationUp_UplinkNormalized:
		f = layout.UplinkNormalizedTopic
	case *ttnpb.ApplicationUp_JoinAccept:
		f = layout.JoinAcceptTopic
	case *ttnpb.ApplicationUp_DownlinkAck:
		f = layout.DownlinkAckTopic
	case *ttnpb.ApplicationUp_DownlinkNack:
		f = layout.DownlinkNackTopic
	case *ttnpb.ApplicationUp_DownlinkSent:
		f = layout.DownlinkSentTopic
	case *ttnpb.ApplicationUp_DownlinkFailed:
		f = layout.DownlinkFailedTopic
	case *ttnpb.ApplicationUp_DownlinkQueued:
		f = layout.DownlinkQueuedTopic
	case *ttnpb.ApplicationUp_DownlinkQueueInvalidated:
		f = layout.DownlinkQueueInvalidatedTopic
	case *ttnpb.ApplicationUp_LocationSolved:
		f = layout.LocationSolvedTopic
	case *ttnpb.ApplicationUp_ServiceData:
		f = layout.ServiceDataTopic
	default:
		panic("unreachable")
	}
	return f(unique.ID(up.Context, up.EndDeviceIds.ApplicationIds), up.EndDeviceIds.DeviceId)
}

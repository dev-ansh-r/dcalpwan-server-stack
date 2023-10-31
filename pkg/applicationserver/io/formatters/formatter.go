package formatters

import (
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Formatter formats upstream and downstream messages.
type Formatter interface {
	FromUp(*ttnpb.ApplicationUp) ([]byte, error)
	ToDownlinks([]byte) (*ttnpb.ApplicationDownlinks, error)
	ToDownlinkQueueRequest([]byte) (*ttnpb.DownlinkQueueRequest, error)
}

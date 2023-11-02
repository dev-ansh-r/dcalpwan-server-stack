package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// Rx2Parameters contains downlink datarate index and channel.
type Rx2Parameters struct {
	DataRateIndex ttnpb.DataRateIndex
	Frequency     uint64
}

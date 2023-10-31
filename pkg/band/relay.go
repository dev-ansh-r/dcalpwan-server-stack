package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// RelayWORChannel encapsulates a relay channel.
type RelayWORChannel struct {
	Frequency     uint64
	ACKFrequency  uint64
	DataRateIndex ttnpb.DataRateIndex
}

// RelayParameters contains relay parameters.
type RelayParameters struct {
	WORChannels []RelayWORChannel
}

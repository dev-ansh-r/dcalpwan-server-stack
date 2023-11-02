package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// Channel abstracts a band's channel properties.
type Channel struct {
	// Frequency indicates the frequency of the channel.
	Frequency uint64
	// MinDataRate indicates the index of the minimal data rates accepted on this channel.
	MinDataRate ttnpb.DataRateIndex
	// MinDataRate indicates the index of the maximal data rates accepted on this channel.
	MaxDataRate ttnpb.DataRateIndex
}

func channelIndexIdentity(idx uint8) (uint8, error) {
	return idx, nil
}

func channelIndexModulo(n uint8) func(uint8) (uint8, error) {
	return func(idx uint8) (uint8, error) {
		return idx % n, nil
	}
}

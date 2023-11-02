package band

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// ComputePeriodicFrequency computes the frequency at time t given the period p and offset offset.
// It panics if no frequencies are provided.
func ComputePeriodicFrequency(t time.Duration, p time.Duration, offset uint32, frequencies ...uint64) uint64 {
	switch n := len(frequencies); n {
	case 0:
		panic("no frequencies available")
	case 1:
		return frequencies[0]
	default:
		return frequencies[(int(offset)+int(t/p))%n]
	}
}

// Beacon parameters of a specific band.
type Beacon struct {
	DataRateIndex ttnpb.DataRateIndex
	CodingRate    string
	Frequencies   []uint64
}

var usAuBeaconFrequencies = func() []uint64 {
	freqs := make([]uint64, 8)
	for i := 0; i < 8; i++ {
		freqs[i] = 923300000 + uint64(i*600000)
	}
	return freqs
}()

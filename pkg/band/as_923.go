package band

import (
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type as923GroupOffset int64

const (
	// AS_923 is the ID of the Asian 923Mhz Group 1 band
	AS_923 = "AS_923"
	// AS_923_2 is the ID of the Asian 923Mhz Group 2 band
	AS_923_2 = "AS_923_2"
	// AS_923_3 is the ID of the Asian 923Mhz Group 3 band
	AS_923_3 = "AS_923_3"
	// AS_923_4 is the ID of the Asian 923Mhz Group 4 band
	AS_923_4 = "AS_923_4"

	as923Group1Offset as923GroupOffset = 0
	as923Group2Offset as923GroupOffset = -1.8 * 1e6
	as923Group3Offset as923GroupOffset = -6.6 * 1e6
	as923Group4Offset as923GroupOffset = -5.9 * 1e6
)

var (
	as923BeaconFrequencies = func(offset as923GroupOffset) []uint64 {
		return []uint64{uint64(923400000 + offset)}
	}

	as923DefaultChannels = func(offset as923GroupOffset) []Channel {
		return []Channel{
			{
				Frequency:   uint64(923200000 + offset),
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			},
			{
				Frequency:   uint64(923400000 + offset),
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			},
		}
	}

	as923DefaultRX2Frequency = func(offset as923GroupOffset) uint64 {
		return uint64(923200000 + offset)
	}

	as923SubBandParameters = func(offset as923GroupOffset) []SubBandParameters {
		var minFrequency, maxFrequency uint64
		switch offset {
		case as923Group1Offset:
			minFrequency = 923000000
			maxFrequency = 923500000
		case as923Group2Offset:
			minFrequency = 921400000
			maxFrequency = 922000000
		case as923Group3Offset:
			minFrequency = 916500000
			maxFrequency = 917000000
		case as923Group4Offset:
			minFrequency = 917300000
			maxFrequency = 917500000
		default:
			panic(fmt.Sprintf("unknown offset %v", offset))
		}
		return []SubBandParameters{
			{
				MinFrequency: minFrequency,
				MaxFrequency: maxFrequency,
				DutyCycle:    0.01,
				MaxEIRP:      16,
			},
		}
	}

	as923RelayParameters = func(offset as923GroupOffset) RelayParameters {
		return RelayParameters{
			WORChannels: []RelayWORChannel{
				{
					Frequency:     uint64(923600000 + offset),
					ACKFrequency:  uint64(923800000 + offset),
					DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
				},
			},
		}
	}
)

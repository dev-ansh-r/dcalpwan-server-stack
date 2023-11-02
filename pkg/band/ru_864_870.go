package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// RU_864_870 is the ID of the Russian frequency plan
	RU_864_870 = "RU_864_870"
)

var (
	ru864870BeaconFrequencies   = []uint64{869100000}
	ru864870PingSlotFrequencies = []uint64{868900000}

	ru864870DefaultChannels = []Channel{
		{
			Frequency:   868900000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   869100000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
	}

	ru864870DownlinkDRTable = [8][6]ttnpb.DataRateIndex{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{3, 2, 1, 0, 0, 0},
		{4, 3, 2, 1, 0, 0},
		{5, 4, 3, 2, 1, 0},
		{6, 5, 4, 3, 2, 1},
		{7, 6, 5, 4, 3, 2},
	}

	ru864870RelayParameters = RelayParameters{
		WORChannels: []RelayWORChannel{
			{
				Frequency:     866100000,
				ACKFrequency:  866300000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
			{
				Frequency:     866500000,
				ACKFrequency:  866900000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
		},
	}
)

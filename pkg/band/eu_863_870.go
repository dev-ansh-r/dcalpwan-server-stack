package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// EU_863_870 is the ID of the European 863-870Mhz band
	EU_863_870 = "EU_863_870"
)

var (
	eu863870BeaconFrequencies = []uint64{869525000}

	eu863870DefaultChannels = []Channel{
		{
			Frequency:   868100000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   868300000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   868500000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
	}

	eu863870DownlinkDRTable = [12][6]ttnpb.DataRateIndex{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{3, 2, 1, 0, 0, 0},
		{4, 3, 2, 1, 0, 0},
		{5, 4, 3, 2, 1, 0},
		{6, 5, 4, 3, 2, 1},
		{7, 6, 5, 4, 3, 2},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
	}

	eu863870RelayParameters = RelayParameters{
		WORChannels: []RelayWORChannel{
			{
				Frequency:     865100000,
				ACKFrequency:  865300000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
			{
				Frequency:     865500000,
				ACKFrequency:  865900000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
		},
	}
)

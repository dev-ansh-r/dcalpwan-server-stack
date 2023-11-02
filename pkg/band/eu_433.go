package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// EU_433 is the ID of the European 433Mhz band
	EU_433 = "EU_433"
)

var (
	eu433BeaconFrequencies = []uint64{434655000}

	eu433DefaultChannels = []Channel{
		{
			Frequency:   433175000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   433375000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   433575000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
	}

	eu433DownlinkDRTable = [8][6]ttnpb.DataRateIndex{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{3, 2, 1, 0, 0, 0},
		{4, 3, 2, 1, 0, 0},
		{5, 4, 3, 2, 1, 0},
		{6, 5, 4, 3, 2, 1},
		{7, 6, 5, 4, 3, 2},
	}
)

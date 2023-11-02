package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// ISM_2400 is the ID of the LoRa 2.4 GHz band.
	ISM_2400 = "ISM_2400"
)

var (
	ism2400BeaconFrequencies = []uint64{2424000000}

	ism2400DefaultChannels = []Channel{
		{
			Frequency:   2403000000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_7,
		},
		{
			Frequency:   2425000000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_7,
		},
		{
			Frequency:   2479000000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_7,
		},
	}

	ism2400DownlinkDRTable = [8][6]ttnpb.DataRateIndex{
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

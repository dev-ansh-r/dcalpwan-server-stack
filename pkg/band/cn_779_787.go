package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// CN_779_787 is the ID of the Chinese 779-787Mhz band
	CN_779_787 = "CN_779_787"
)

var (
	cn779787BeaconFrequencies = []uint64{785000000}

	cn779787DefaultChannels = []Channel{
		{
			Frequency:   779500000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   779700000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   779900000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   780500000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   780700000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   780900000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
	}

	cn779787DownlinkDRTable = [8][6]ttnpb.DataRateIndex{
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

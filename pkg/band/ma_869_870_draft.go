package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// MA_869_870_DRAFT is the ID of the draft Morocco 869-870Mhz band.
	MA_869_870_DRAFT = "MA_869_870_DRAFT"
)

var (
	ma869870DraftBeaconFrequencies = []uint64{869525000}

	ma869870DraftDefaultChannels = []Channel{
		{
			Frequency:   869100000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   869300000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   869700000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
	}

	ma869870DraftDownlinkDRTable = [8][6]ttnpb.DataRateIndex{
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

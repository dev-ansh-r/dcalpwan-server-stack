package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// KR_920_923 is the ID of the Korean frequency plan
	KR_920_923 = "KR_920_923"
)

var (
	kr920923BeaconFrequencies = []uint64{923100000}

	kr920923DefaultChannels = []Channel{
		{
			Frequency:   922100000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   922300000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   922500000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
	}

	kr920923DownlinkDRTable = [6][6]ttnpb.DataRateIndex{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{3, 2, 1, 0, 0, 0},
		{4, 3, 2, 1, 0, 0},
		{5, 4, 3, 2, 1, 0},
	}

	kr920923RelayParameters = RelayParameters{
		WORChannels: []RelayWORChannel{
			{
				Frequency:     922700000,
				ACKFrequency:  922900000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
			{
				Frequency:     923100000,
				ACKFrequency:  923100000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
		},
	}
)

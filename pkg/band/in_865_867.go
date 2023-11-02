package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// IN_865_867 is the ID of the Indian frequency plan
	IN_865_867 = "IN_865_867"
)

var (
	in865867BeaconFrequencies = []uint64{866500000}

	in865867DefaultChannels = []Channel{
		{
			Frequency:   865062500,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   865402500,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
		{
			Frequency:   865985000,
			MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
		},
	}

	in865867RelayParameters = RelayParameters{
		WORChannels: []RelayWORChannel{
			{
				Frequency:     866000000,
				ACKFrequency:  866200000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
			{
				Frequency:     866700000,
				ACKFrequency:  866900000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			},
		},
	}
)

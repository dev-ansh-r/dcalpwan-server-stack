package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// US_902_928 is the ID of the US frequency plan
	US_902_928 = "US_902_928"
)

var (
	us902928UplinkChannels = func(wideChannelDelta ttnpb.DataRateIndex) []Channel {
		uplinkChannels := make([]Channel, 0, 72)
		for i := 0; i < 64; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(902300000 + 200000*i),
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_3,
			})
		}
		for i := 0; i < 8; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(903000000 + 1600000*i),
				MinDataRate: ttnpb.DataRateIndex_DATA_RATE_4,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_4 + wideChannelDelta,
			})
		}
		return uplinkChannels
	}

	us902928DownlinkChannels = func() []Channel {
		downlinkChannels := make([]Channel, 0, 8)
		for i := 0; i < 8; i++ {
			downlinkChannels = append(downlinkChannels, Channel{
				Frequency:   uint64(923300000 + 600000*i),
				MinDataRate: ttnpb.DataRateIndex_DATA_RATE_8,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_13,
			})
		}
		return downlinkChannels
	}()

	us902928DownlinkDRTable = [7][4]ttnpb.DataRateIndex{
		{10, 9, 8, 8},
		{11, 10, 9, 8},
		{12, 11, 10, 9},
		{13, 12, 11, 10},
		{13, 13, 12, 11},
		{10, 9, 8, 8},
		{11, 10, 9, 8},
	}

	us902928RelayParameters = RelayParameters{
		WORChannels: []RelayWORChannel{
			{
				Frequency:     916700000,
				ACKFrequency:  918300000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_10,
			},
			{
				Frequency:     919900000,
				ACKFrequency:  921500000,
				DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_10,
			},
		},
	}
)

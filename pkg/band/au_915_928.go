package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// AU_915_928 is the ID of the Australian band
	AU_915_928 = "AU_915_928"
)

var (
	au915928DownlinkDRTableLegacy = [7][6]ttnpb.DataRateIndex{
		{10, 9, 8, 8},
		{11, 10, 9, 8},
		{12, 11, 10, 9},
		{13, 12, 11, 10},
		{13, 13, 12, 11},
	}

	au915928DownlinkDRTable = [8][6]ttnpb.DataRateIndex{
		{8, 8, 8, 8, 8, 8},
		{9, 8, 8, 8, 8, 8},
		{10, 9, 8, 8, 8, 8},
		{11, 10, 9, 8, 8, 8},
		{12, 11, 10, 9, 8, 8},
		{13, 12, 11, 10, 9, 8},
		{13, 13, 12, 11, 10, 9},
		{9, 8, 8, 8, 8, 8},
	}

	au915928UplinkChannels = func(commonDelta, wideChannelDelta ttnpb.DataRateIndex) []Channel {
		uplinkChannels := make([]Channel, 0, 72)
		for i := 0; i < 64; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(915200000 + 200000*i),
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5 + commonDelta,
			})
		}
		for i := 0; i < 8; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(915900000 + 1600000*i),
				MinDataRate: ttnpb.DataRateIndex_DATA_RATE_6 + commonDelta,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_6 + commonDelta + wideChannelDelta,
			})
		}
		return uplinkChannels
	}

	au915928DownlinkChannels = func() []Channel {
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

	au915928RelayParameters = RelayParameters{
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

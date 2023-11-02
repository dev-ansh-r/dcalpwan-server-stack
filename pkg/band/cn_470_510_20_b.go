package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// CN_470_510_20_B is the ID of the Chinese 470-510Mhz band which uses a 20 MHz antenna, type B
	CN_470_510_20_B = "CN_470_510_20_B"
)

var (
	cn47051020BUplinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		uplinkChannels := make([]Channel, 0, 64)
		// 20 MHz Type B Group 1
		for i := 0; i < 32; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(476900000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		// 20 MHz Type B Group 2
		for i := 0; i < 32; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(496900000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return uplinkChannels
	}

	cn47051020BDownlinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		downlinkChannels := make([]Channel, 0, 64)
		// 20 MHz Type B Group 1
		for i := 0; i < 32; i++ {
			downlinkChannels = append(downlinkChannels, Channel{
				Frequency:   uint64(476900000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		// 20 MHz Type B Group 2
		for i := 0; i < 32; i++ {
			downlinkChannels = append(downlinkChannels, Channel{
				Frequency:   uint64(496900000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return downlinkChannels
	}
)

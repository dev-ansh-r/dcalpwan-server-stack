package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// CN_470_510_20_A is the ID of the Chinese 470-510Mhz band which uses a 20 MHz antenna, type A
	CN_470_510_20_A = "CN_470_510_20_A"
)

var (
	cn47051020AUplinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		uplinkChannels := make([]Channel, 0, 64)
		// 20 MHz Type A Group 1
		for i := 0; i < 32; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(470300000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		// 20 MHz Type A Group 2
		for i := 0; i < 32; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(503500000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return uplinkChannels
	}

	cn47051020ADownlinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		downlinkChannels := make([]Channel, 0, 64)
		// 20 MHz Type A Group 1
		for i := 0; i < 32; i++ {
			downlinkChannels = append(downlinkChannels, Channel{
				Frequency:   uint64(483900000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		// 20 MHz Type A Group 2
		for i := 0; i < 32; i++ {
			downlinkChannels = append(downlinkChannels, Channel{
				Frequency:   uint64(490300000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return downlinkChannels
	}
)

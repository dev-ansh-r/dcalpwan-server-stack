package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// CN_470_510_26_A is the ID of the Chinese 470-510Mhz band which uses a 26 MHz antenna, type A
	CN_470_510_26_A = "CN_470_510_26_A"
)

var (
	cn47051026ABeaconFrequencies = []uint64{494900000}

	cn47051026AUplinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		uplinkChannels := make([]Channel, 0, 48)
		// 26 MHz Type A
		for i := 0; i < 48; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(470300000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return uplinkChannels
	}

	cn47051026ADownlinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		downlinkChannels := make([]Channel, 0, 24)
		// 26 MHz Type A
		for i := 0; i < 24; i++ {
			downlinkChannels = append(downlinkChannels, Channel{
				Frequency:   uint64(490100000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return downlinkChannels
	}
)

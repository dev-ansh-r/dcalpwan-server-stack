package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

const (
	// CN_470_510_26_B is the ID of the Chinese 470-510Mhz band which uses a 26 MHz antenna, type B
	CN_470_510_26_B = "CN_470_510_26_B"
)

var (
	cn47051026BBeaconFrequencies = []uint64{504900000}

	cn47051026BUplinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		uplinkChannels := make([]Channel, 0, 48)
		// 26 MHz Type B
		for i := 0; i < 48; i++ {
			uplinkChannels = append(uplinkChannels, Channel{
				Frequency:   uint64(480300000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return uplinkChannels
	}

	cn47051026BDownlinkChannels = func(minDataRateIndex ttnpb.DataRateIndex) []Channel {
		downlinkChannels := make([]Channel, 0, 24)
		// 26 MHz Type B
		for i := 0; i < 24; i++ {
			downlinkChannels = append(downlinkChannels, Channel{
				Frequency:   uint64(500100000 + 200000*i),
				MinDataRate: minDataRateIndex,
				MaxDataRate: ttnpb.DataRateIndex_DATA_RATE_5,
			})
		}
		return downlinkChannels
	}
)

package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// ISM_2400_Universal is the band definition for universal LoRa 2.4 GHz.
var ISM_2400_Universal = Band{
	ID: ISM_2400,

	MaxUplinkChannels: 16,
	UplinkChannels:    ism2400DefaultChannels,

	MaxDownlinkChannels: 16,
	DownlinkChannels:    ism2400DefaultChannels,

	SubBands: []SubBandParameters{
		{
			MinFrequency: 2400000000,
			MaxFrequency: 2500000000,
			DutyCycle:    1,
			MaxEIRP:      8.0 + eirpDelta,
		},
	},

	DataRates: map[ttnpb.DataRateIndex]DataRate{
		ttnpb.DataRateIndex_DATA_RATE_0: makeLoRaDataRate(12, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_1: makeLoRaDataRate(11, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(123)),
		ttnpb.DataRateIndex_DATA_RATE_2: makeLoRaDataRate(10, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_3: makeLoRaDataRate(9, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_4: makeLoRaDataRate(8, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_5: makeLoRaDataRate(7, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_6: makeLoRaDataRate(6, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_7: makeLoRaDataRate(5, 812000, Cr4_8LI, makeConstMaxMACPayloadSizeFunc(250)),
	},
	StrictCodingRate: true,

	DefaultMaxEIRP: 10,
	TxOffset: []float32{
		0,
		-2,
		-4,
		-6,
		-8,
		-10,
		-12,
		-14,
	},

	Rx1Channel: channelIndexIdentity,
	Rx1DataRate: func(idx ttnpb.DataRateIndex, offset ttnpb.DataRateOffset, _ bool) (ttnpb.DataRateIndex, error) {
		if idx > ttnpb.DataRateIndex_DATA_RATE_7 {
			return 0, errDataRateIndexTooHigh.WithAttributes("max", 7)
		}
		if offset > 5 {
			return 0, errDataRateOffsetTooHigh.WithAttributes("max", 5)
		}
		return ism2400DownlinkDRTable[idx][offset], nil
	},

	GenerateChMasks: generateChMask16,
	ParseChMask:     parseChMask16,

	FreqMultiplier:   200,
	ImplementsCFList: true,
	CFListType:       ttnpb.CFListType_FREQUENCIES,

	DefaultRx2Parameters: Rx2Parameters{ttnpb.DataRateIndex_DATA_RATE_0, 2423000000},

	Beacon: Beacon{
		DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
		CodingRate:    Cr4_8LI,
		Frequencies:   ism2400BeaconFrequencies,
	},
	PingSlotFrequencies: ism2400BeaconFrequencies,

	SharedParameters: relayAwareSharedParameters,
}

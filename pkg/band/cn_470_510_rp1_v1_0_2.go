package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// CN_470_510_RP1_v1_0_2 is the band definition for CN470-510 in the RP1 v1.0.2 specification.
var CN_470_510_RP1_v1_0_2 = Band{
	ID: CN_470_510,

	SupportsDynamicADR: true,

	MaxUplinkChannels: 96,
	UplinkChannels:    cn470510UplinkChannels,

	MaxDownlinkChannels: 48,
	DownlinkChannels:    cn470510DownlinkChannels,

	// See IEEE 11-11/0972r0
	SubBands: []SubBandParameters{
		{
			MinFrequency: 470000000,
			MaxFrequency: 510000000,
			DutyCycle:    1,
			MaxEIRP:      17.0 + eirpDelta,
		},
	},

	DataRates: map[ttnpb.DataRateIndex]DataRate{
		ttnpb.DataRateIndex_DATA_RATE_0: makeLoRaDataRate(12, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_1: makeLoRaDataRate(11, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_2: makeLoRaDataRate(10, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_3: makeLoRaDataRate(9, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(123)),
		ttnpb.DataRateIndex_DATA_RATE_4: makeLoRaDataRate(8, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_5: makeLoRaDataRate(7, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
	},
	MaxADRDataRateIndex: ttnpb.DataRateIndex_DATA_RATE_5,
	StrictCodingRate:    true,

	DefaultMaxEIRP: 17,
	TxOffset: []float32{
		0,
		-1,
		-3,
		-5,
		-7,
		-10,
		-12,
		-15,
	},

	Rx1Channel: channelIndexModulo(48),
	Rx1DataRate: func(idx ttnpb.DataRateIndex, offset ttnpb.DataRateOffset, _ bool) (ttnpb.DataRateIndex, error) {
		if idx > ttnpb.DataRateIndex_DATA_RATE_5 {
			return 0, errDataRateIndexTooHigh.WithAttributes("max", 5)
		}
		if offset > 5 {
			return 0, errDataRateOffsetTooHigh.WithAttributes("max", 5)
		}
		return cn470510DownlinkDRTable[idx][offset], nil
	},

	GenerateChMasks: generateChMask96,
	ParseChMask:     parseChMask96,

	DefaultRx2Parameters: Rx2Parameters{ttnpb.DataRateIndex_DATA_RATE_0, 505300000},

	Beacon: Beacon{
		DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_2,
		CodingRate:    Cr4_5,
		Frequencies:   cn470510BeaconFrequencies,
	},
	PingSlotFrequencies: cn470510BeaconFrequencies,

	FreqMultiplier:   100,
	ImplementsCFList: false,
	CFListType:       ttnpb.CFListType_CHANNEL_MASKS,

	SharedParameters: universalSharedParameters,
}

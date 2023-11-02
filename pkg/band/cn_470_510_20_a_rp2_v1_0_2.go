package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// CN_470_510_20_A_RP2_v1_0_2 is the band definition for CN470-510 20MHz antenna, type A in the RP002-1.0.2 specification.
var CN_470_510_20_A_RP2_v1_0_2 = Band{
	ID: CN_470_510_20_A,

	SupportsDynamicADR: true,

	MaxUplinkChannels: 64,
	UplinkChannels:    cn47051020AUplinkChannels(1),

	MaxDownlinkChannels: 64,
	DownlinkChannels:    cn47051020ADownlinkChannels(1),

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
		ttnpb.DataRateIndex_DATA_RATE_1: makeLoRaDataRate(11, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(31)),
		ttnpb.DataRateIndex_DATA_RATE_2: makeLoRaDataRate(10, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(94)),
		ttnpb.DataRateIndex_DATA_RATE_3: makeLoRaDataRate(9, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(172)),
		ttnpb.DataRateIndex_DATA_RATE_4: makeLoRaDataRate(8, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_5: makeLoRaDataRate(7, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_6: makeLoRaDataRate(7, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_7: makeFSKDataRate(50000, makeConstMaxMACPayloadSizeFunc(250)),
	},
	MaxADRDataRateIndex: ttnpb.DataRateIndex_DATA_RATE_5,
	StrictCodingRate:    true,

	DefaultMaxEIRP: 19.15,
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
		if idx > ttnpb.DataRateIndex_DATA_RATE_5 {
			return 0, errDataRateIndexTooHigh.WithAttributes("max", 5)
		}
		if offset > 5 {
			return 0, errDataRateOffsetTooHigh.WithAttributes("max", 5)
		}
		// Unchanged from the pre-RP2 CN470-510 band definition.
		return cn470510DownlinkDRTable[idx][offset], nil
	},

	GenerateChMasks: generateChMask64,
	ParseChMask:     parseChMask64,

	DefaultRx2Parameters: Rx2Parameters{ttnpb.DataRateIndex_DATA_RATE_0, 486900000},

	Beacon: Beacon{
		DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_2,
		CodingRate:    Cr4_5,
	},

	FreqMultiplier:   100,
	ImplementsCFList: true,
	CFListType:       ttnpb.CFListType_CHANNEL_MASKS,

	SharedParameters: universalSharedParameters,
}

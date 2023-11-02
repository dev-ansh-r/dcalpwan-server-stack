package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// AU_915_928_TS1_v1_0_1 is the band definition for AU915-928 in the TS1 v1.0.1 specification.
var AU_915_928_TS1_v1_0_1 = Band{
	ID: AU_915_928,

	SupportsDynamicADR: true,

	MaxUplinkChannels: 72,
	UplinkChannels:    au915928UplinkChannels(-2, 0),

	MaxDownlinkChannels: 8,
	DownlinkChannels:    au915928DownlinkChannels,

	// See Radiocommunications (Low Interference Potential Devices) Class Licence 2015
	SubBands: []SubBandParameters{
		{
			MinFrequency: 915000000,
			MaxFrequency: 928000000,
			DutyCycle:    1,
			MaxEIRP:      30,
		},
	},

	DataRates: map[ttnpb.DataRateIndex]DataRate{
		ttnpb.DataRateIndex_DATA_RATE_0: makeLoRaDataRate(10, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_1: makeLoRaDataRate(9, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(123)),
		ttnpb.DataRateIndex_DATA_RATE_2: makeLoRaDataRate(8, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_3: makeLoRaDataRate(7, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_4: makeLoRaDataRate(8, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),

		ttnpb.DataRateIndex_DATA_RATE_8:  makeLoRaDataRate(12, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(61)),
		ttnpb.DataRateIndex_DATA_RATE_9:  makeLoRaDataRate(11, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(137)),
		ttnpb.DataRateIndex_DATA_RATE_10: makeLoRaDataRate(10, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_11: makeLoRaDataRate(9, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_12: makeLoRaDataRate(8, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_13: makeLoRaDataRate(7, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
	},
	MaxADRDataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
	StrictCodingRate:    true,

	DefaultMaxEIRP: 30,
	TxOffset: []float32{
		0,
		-2,
		-4,
		-6,
		-8,
		-10,
		-12,
		-14,
		-16,
		-18,
		-20,
	},

	FreqMultiplier:   100,
	ImplementsCFList: false,
	CFListType:       ttnpb.CFListType_CHANNEL_MASKS,

	Rx1Channel: channelIndexModulo(8),
	Rx1DataRate: func(idx ttnpb.DataRateIndex, offset ttnpb.DataRateOffset, _ bool) (ttnpb.DataRateIndex, error) {
		if idx > ttnpb.DataRateIndex_DATA_RATE_4 {
			return 0, errDataRateIndexTooHigh.WithAttributes("max", 4)
		}
		if offset > 3 {
			return 0, errDataRateOffsetTooHigh.WithAttributes("max", 3)
		}
		return au915928DownlinkDRTableLegacy[idx][offset], nil
	},

	GenerateChMasks: makeGenerateChMask72(false, false),
	ParseChMask:     parseChMask72,

	DefaultRx2Parameters: Rx2Parameters{ttnpb.DataRateIndex_DATA_RATE_8, 923300000},

	Beacon: Beacon{
		DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_10,
		CodingRate:    Cr4_5,
		Frequencies:   usAuBeaconFrequencies,
	},
	PingSlotFrequencies: usAuBeaconFrequencies,

	TxParamSetupReqSupport: false,

	SharedParameters: universalSharedParameters,
}

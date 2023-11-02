package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// IN_865_867_RP1_V1_0_2_Rev_B is the band definition for IN865-867 in the RP1 v1.0.2 rev. B specification.
var IN_865_867_RP1_V1_0_2_Rev_B = Band{
	ID: IN_865_867,

	SupportsDynamicADR: true,

	MaxUplinkChannels: 16,
	UplinkChannels:    in865867DefaultChannels,

	MaxDownlinkChannels: 16,
	DownlinkChannels:    in865867DefaultChannels,

	SubBands: []SubBandParameters{
		{
			MinFrequency: 865000000,
			MaxFrequency: 867000000,
			DutyCycle:    1,
			MaxEIRP:      14.0 + eirpDelta,
		},
	},

	DataRates: map[ttnpb.DataRateIndex]DataRate{
		ttnpb.DataRateIndex_DATA_RATE_0: makeLoRaDataRate(12, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_1: makeLoRaDataRate(11, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_2: makeLoRaDataRate(10, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(59)),
		ttnpb.DataRateIndex_DATA_RATE_3: makeLoRaDataRate(9, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(123)),
		ttnpb.DataRateIndex_DATA_RATE_4: makeLoRaDataRate(8, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_5: makeLoRaDataRate(7, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),

		ttnpb.DataRateIndex_DATA_RATE_7: makeFSKDataRate(50000, makeConstMaxMACPayloadSizeFunc(250)),
	},
	MaxADRDataRateIndex: ttnpb.DataRateIndex_DATA_RATE_5,

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

	Rx1Channel: channelIndexIdentity,
	Rx1DataRate: func(idx ttnpb.DataRateIndex, offset ttnpb.DataRateOffset, _ bool) (ttnpb.DataRateIndex, error) {
		so := int8(offset)
		if so > 5 {
			so = 5 - so
		}
		si := int8(idx) - so

		switch {
		case si <= 0:
			return ttnpb.DataRateIndex_DATA_RATE_0, nil
		case si >= 5:
			return ttnpb.DataRateIndex_DATA_RATE_5, nil
		}
		return ttnpb.DataRateIndex(si), nil
	},

	GenerateChMasks: generateChMask16,
	ParseChMask:     parseChMask16,

	FreqMultiplier:   100,
	ImplementsCFList: true,
	CFListType:       ttnpb.CFListType_FREQUENCIES,

	DefaultRx2Parameters: Rx2Parameters{ttnpb.DataRateIndex_DATA_RATE_2, 866550000},

	Beacon: Beacon{
		DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_4,
		CodingRate:    Cr4_5,
		Frequencies:   in865867BeaconFrequencies,
	},
	PingSlotFrequencies: in865867BeaconFrequencies,

	SharedParameters: universalSharedParameters,
}

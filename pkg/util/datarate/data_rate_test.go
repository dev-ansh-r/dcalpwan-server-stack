package datarate_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/datarate"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestDataRate(t *testing.T) {
	a := assertions.New(t)

	table := map[string]datarate.DR{
		`"SF7BW125"`: {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lora{
					Lora: &ttnpb.LoRaDataRate{
						SpreadingFactor: 7,
						Bandwidth:       125000,
					},
				},
			},
		},
		`50000`: {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Fsk{
					Fsk: &ttnpb.FSKDataRate{
						BitRate: 50000,
					},
				},
			},
		},
		`"M0CW137"`: {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lrfhss{
					Lrfhss: &ttnpb.LRFHSSDataRate{
						ModulationType:        0,
						OperatingChannelWidth: 137000,
					},
				},
			},
		},
	}

	for s, dr := range table {
		enc, err := dr.MarshalJSON()
		a.So(err, should.BeNil)
		a.So(string(enc), should.Equal, s)

		var dec datarate.DR
		err = dec.UnmarshalJSON(enc)
		a.So(err, should.BeNil)
		a.So(dec, should.Resemble, dr)
	}

	var dr datarate.DR
	err := dr.UnmarshalJSON([]byte{})
	a.So(err, should.NotBeNil)
}

func TestValidLoRaDataRateParsing(t *testing.T) {
	a := assertions.New(t)

	table := map[string]datarate.DR{
		"SF6BW125": {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lora{
					Lora: &ttnpb.LoRaDataRate{
						SpreadingFactor: 6,
						Bandwidth:       125000,
					},
				},
			},
		},
		"SF9BW500": {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lora{
					Lora: &ttnpb.LoRaDataRate{
						SpreadingFactor: 9,
						Bandwidth:       500000,
					},
				},
			},
		},
		"SF5BW31.25": {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lora{
					Lora: &ttnpb.LoRaDataRate{
						SpreadingFactor: 5,
						Bandwidth:       31250,
					},
				},
			},
		},
	}
	for dr, expected := range table {
		actual, err := datarate.ParseLoRa(dr)
		a.So(err, should.BeNil)
		a.So(actual, should.Resemble, expected)
	}
}

func TestInvalidLoRaDataRateParsing(t *testing.T) {
	a := assertions.New(t)

	table := []string{
		"6BW125",
		"SF9B500",
	}
	for _, dr := range table {
		_, err := datarate.ParseLoRa(dr)
		a.So(err, should.NotBeNil)
	}
}

func TestStringer(t *testing.T) {
	a := assertions.New(t)

	table := map[datarate.DR]string{
		{
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lora{
					Lora: &ttnpb.LoRaDataRate{
						SpreadingFactor: 6,
						Bandwidth:       125000,
					},
				},
			},
		}: "SF6BW125",
		{
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lora{
					Lora: &ttnpb.LoRaDataRate{
						SpreadingFactor: 9,
						Bandwidth:       500000,
					},
				},
			},
		}: "SF9BW500",
		{
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lora{
					Lora: &ttnpb.LoRaDataRate{
						SpreadingFactor: 5,
						Bandwidth:       31250,
					},
				},
			},
		}: "SF5BW31.25",
		{
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Fsk{
					Fsk: &ttnpb.FSKDataRate{
						BitRate: 50000,
					},
				},
			},
		}: "50000",
		{
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lrfhss{
					Lrfhss: &ttnpb.LRFHSSDataRate{
						ModulationType:        0,
						OperatingChannelWidth: 137000,
					},
				},
			},
		}: "M0CW137",
	}

	for dr, expected := range table {
		a.So(dr.String(), should.Equal, expected)
	}
}

func TestLRFHSSDataRateParsing(t *testing.T) {
	a := assertions.New(t)

	table := map[string]datarate.DR{
		"M0CW137": {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lrfhss{
					Lrfhss: &ttnpb.LRFHSSDataRate{
						ModulationType:        0,
						OperatingChannelWidth: 137000,
					},
				},
			},
		},
		"M12CW1375": {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lrfhss{
					Lrfhss: &ttnpb.LRFHSSDataRate{
						ModulationType:        12,
						OperatingChannelWidth: 1375000,
					},
				},
			},
		},
		"M1CW1": {
			DataRate: &ttnpb.DataRate{
				Modulation: &ttnpb.DataRate_Lrfhss{
					Lrfhss: &ttnpb.LRFHSSDataRate{
						ModulationType:        1,
						OperatingChannelWidth: 1000,
					},
				},
			},
		},
	}
	for dr, expected := range table {
		actual, err := datarate.ParseLRFHSS(dr)
		a.So(err, should.BeNil)
		a.So(actual, should.Resemble, expected)
	}

	invalidTable := []string{
		"MCW124",
		"M0CWA124",
		"M0CW124A",
	}
	for _, dr := range invalidTable {
		_, err := datarate.ParseLRFHSS(dr)
		a.So(err, should.NotBeNil)
	}
}

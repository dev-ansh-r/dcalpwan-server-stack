package ttnpb_test

import (
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestDataRateIndex(t *testing.T) {
	a := assertions.New(t)
	a.So(DataRateIndex_DATA_RATE_4.String(), should.Equal, "DATA_RATE_4")

	b, err := DataRateIndex_DATA_RATE_4.MarshalText()
	a.So(err, should.BeNil)
	a.So(string(b), should.Resemble, "4")

	for _, str := range []string{"4", "DATA_RATE_4"} {
		var idx DataRateIndex
		err = idx.UnmarshalText([]byte(str))
		a.So(idx, should.Equal, DataRateIndex_DATA_RATE_4)
	}
}

func TestDeviceEIRP(t *testing.T) {
	a := assertions.New(t)
	a.So(DeviceEIRP_DEVICE_EIRP_10.String(), should.Equal, "DEVICE_EIRP_10")

	b, err := DeviceEIRP_DEVICE_EIRP_10.MarshalText()
	a.So(err, should.BeNil)
	a.So(b, should.Resemble, []byte("DEVICE_EIRP_10"))

	var v DeviceEIRP
	err = v.UnmarshalText([]byte("DEVICE_EIRP_10"))
	a.So(v, should.Equal, DeviceEIRP_DEVICE_EIRP_10)
	a.So(err, should.BeNil)
}

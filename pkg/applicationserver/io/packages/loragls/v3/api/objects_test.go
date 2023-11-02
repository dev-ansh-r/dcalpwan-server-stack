package api_test

import (
	"encoding/json"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages/loragls/v3/api"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestUplinkUnmarshal(t *testing.T) {
	a := assertions.New(t)
	up := &api.Uplink{}

	err := up.UnmarshalJSON([]byte(`["gtw1", 1, 123, 42.42, 64.64]`))
	if a.So(err, should.BeNil) {
		a.So(up.GatewayID, should.Equal, "gtw1")
		if a.So(up.AntennaID, should.NotBeNil) {
			a.So(*up.AntennaID, should.Equal, 1)
		}
		if a.So(up.TDOA, should.NotBeNil) {
			a.So(*up.TDOA, should.Equal, 123)
		}
		a.So(up.RSSI, should.Equal, 42.42)
		a.So(up.SNR, should.Equal, 64.64)
	}

	err = up.UnmarshalJSON([]byte(`[]`))
	a.So(err, should.NotBeNil)

	err = up.UnmarshalJSON([]byte(`["gtw1", "error0", "error1", "error2", "error3"]`))
	a.So(err, should.NotBeNil)
}

func TestUplinkMarshal(t *testing.T) {
	a := assertions.New(t)

	b, err := json.Marshal(&api.Uplink{
		GatewayID: "gtw1",
		AntennaID: uint32Ptr(0),
		TDOA:      uint64Ptr(123),
		RSSI:      456.5,
		SNR:       567.8,
	})
	if a.So(err, should.BeNil) {
		a.So(b, should.Resemble, []byte(`["gtw1",0,123,456.5,567.8]`))
	}

	b, err = json.Marshal(&api.Uplink{
		GatewayID: "gtw1",
		AntennaID: nil,
		TDOA:      nil,
		RSSI:      123.4,
		SNR:       456.7,
	})
	if a.So(err, should.BeNil) {
		a.So(b, should.Resemble, []byte(`["gtw1",null,null,123.4,456.7]`))
	}
}

func uint32Ptr(u uint32) *uint32 {
	return &u
}

func uint64Ptr(u uint64) *uint64 {
	return &u
}

func TestGatewayIDsConversion(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	expectedProto := &ttnpb.GatewayIdentifiers{
		GatewayId: "test-gateway",
	}
	expectedData := &api.GatewayIDs{
		GatewayID: "test-gateway",
	}

	actualData := &api.GatewayIDs{}
	if err := actualData.FromProto(expectedProto); err != nil {
		t.Fatalf("FromProto failed: %v", err)
	}

	actualProto := actualData.ToProto()
	a.So(actualProto, assertions.ShouldResemble, expectedProto)
	a.So(actualData, assertions.ShouldResemble, expectedData)
}

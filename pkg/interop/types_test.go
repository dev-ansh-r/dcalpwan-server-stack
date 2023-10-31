package interop_test

import (
	"encoding/json"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/interop"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

type testMessage struct {
	MACVersion interop.MACVersion
	Buffer     interop.Buffer `json:",omitempty"`
	Key        *interop.KeyEnvelope
}

func TestMarshalTypes(t *testing.T) { //nolint:paralleltest
	a := assertions.New(t)

	{
		msg := &testMessage{
			MACVersion: interop.MACVersion(ttnpb.MACVersion_MAC_V1_0_2),
			Buffer:     interop.Buffer([]byte{0x1, 0x2, 0x3}),
			Key: &interop.KeyEnvelope{
				KekLabel: "",
				Key:      types.AES128Key{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf}.Bytes(), //nolint:lll
			},
		}
		data, err := json.Marshal(msg)
		if a.So(err, should.BeNil) {
			a.So(string(data), should.Equal, `{"MACVersion":"1.0.2","Buffer":"010203","Key":{"KEKLabel":"","AESKey":"000102030405060708090A0B0C0D0E0F"}}`) //nolint:lll
		}
	}

	{
		msg := &testMessage{
			MACVersion: interop.MACVersion(ttnpb.MACVersion_MAC_V1_1),
			Key: &interop.KeyEnvelope{
				KekLabel:     "test",
				EncryptedKey: []byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf},
			},
		}
		data, err := json.Marshal(msg)
		if a.So(err, should.BeNil) {
			a.So(string(data), should.Equal, `{"MACVersion":"1.1","Key":{"KEKLabel":"test","AESKey":"000102030405060708090A0B0C0D0E0F"}}`) //nolint:lll
		}
	}
}

func TestUnmarshalTypes(t *testing.T) { //nolint:paralleltest
	a := assertions.New(t)

	{
		data := []byte(`{"MACVersion":"1.0.2","Buffer":"010203"}`)
		var msg testMessage
		err := json.Unmarshal(data, &msg)
		if a.So(err, should.BeNil) {
			a.So(msg, should.Resemble, testMessage{
				MACVersion: interop.MACVersion(ttnpb.MACVersion_MAC_V1_0_2),
				Buffer:     interop.Buffer([]byte{0x1, 0x2, 0x3}),
			})
		}
	}

	{
		data := []byte(`{"MACVersion":"1.1","Buffer":"0x010203"}`)
		var msg testMessage
		err := json.Unmarshal(data, &msg)
		if a.So(err, should.BeNil) {
			a.So(msg, should.Resemble, testMessage{
				MACVersion: interop.MACVersion(ttnpb.MACVersion_MAC_V1_1),
				Buffer:     interop.Buffer([]byte{0x1, 0x2, 0x3}),
			})
		}
	}

	{
		data := []byte(`{"MACVersion":"1.0"}`)
		var msg testMessage
		err := json.Unmarshal(data, &msg)
		if a.So(err, should.BeNil) {
			a.So(msg, should.Resemble, testMessage{
				MACVersion: interop.MACVersion(ttnpb.MACVersion_MAC_V1_0),
			})
		}
	}
}

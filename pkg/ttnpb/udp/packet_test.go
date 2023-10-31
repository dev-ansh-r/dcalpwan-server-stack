package udp

import (
	"bytes"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestPacket(t *testing.T) {
	eui := new(types.EUI64)
	data := new(Data)
	p := Packet{
		ProtocolVersion: Version1,
		Token:           [2]byte{},
		PacketType:      PushData,
		GatewayEUI:      eui,
		Data:            data,
	}

	res, err := p.MarshalBinary()
	if err != nil {
		t.Error("Failed to marshal packet:", err)
	}

	var p2 Packet
	err = p2.UnmarshalBinary(res)
	if err != nil {
		t.Error("Failed to unmarshal binary packet:", err)
	}

	p.BuildAck()
}

func TestFailedPackets(t *testing.T) {
	a := assertions.New(t)

	{
		p := new(Packet)
		b := []byte{}
		err := p.UnmarshalBinary(b)
		a.So(err, should.NotBeNil)
	}

	{
		p := new(Packet)
		b := []byte{0, 0, 0, 0}
		err := p.UnmarshalBinary(b)
		a.So(err, should.NotBeNil)
	}

	{
		p := new(Packet)
		b := bytes.Repeat([]byte{0x0}, 9)
		err := p.UnmarshalBinary(b)
		a.So(err, should.NotBeNil)
	}

	{
		p := new(Packet)
		b := []byte{1, 0, 0, byte(PushData), 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8}
		err := p.UnmarshalBinary(b)
		a.So(err, should.BeNil)
		a.So(p.Data, should.NotBeNil)
	}
}

package component_test

import (
	"net"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

const udpListenAddr = "0.0.0.0:8056"

func TestListenUDP(t *testing.T) {
	a := assertions.New(t)

	c := component.MustNew(test.GetLogger(t), &component.Config{})
	conn, err := c.ListenUDP(udpListenAddr)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	senderConn, err := net.Dial("udp", udpListenAddr)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	content := []byte{0xaa, 0xbb, 0xcc, 0x03}
	go func() {
		_, err := senderConn.Write(content)
		a.So(err, should.BeNil)
	}()
	receptionBuf := make([]byte, 256)
	_, err = conn.Read(receptionBuf)
	a.So(err, should.BeNil)
	for i := range content {
		a.So(receptionBuf[i], should.Equal, content[i])
	}
}

package loraclouddevicemanagementv1

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages/loradms/v1/api/objects"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestTLVEncoding(t *testing.T) {
	a := assertions.New(t)

	a.So(parseTLVPayload(objects.Hex{0xbb, 0xaa}, func(tag uint8, data []byte) error {
		return fmt.Errorf("foo")
	}), should.NotBeNil)

	a.So(parseTLVPayload(objects.Hex{0x01, 0x02, 0xbb, 0xaa}, func(tag uint8, data []byte) error {
		a.So(tag, should.Equal, 0x01)
		a.So(data, should.HaveLength, 2)
		a.So(data, should.Resemble, []byte{0xbb, 0xaa})
		return nil
	}), should.BeNil)

	a.So(parseTLVPayload(objects.Hex{0xff, 0x02, 0xff}, func(tag uint8, data []byte) error {
		t.Fatal("f should not be called")
		return nil
	}), should.NotBeNil)

	a.So(parseTLVPayload(objects.Hex{0xff, 0xff, 0xff}, func(tag uint8, data []byte) error {
		t.Fatal("f should not be called")
		return nil
	}), should.NotBeNil)
}

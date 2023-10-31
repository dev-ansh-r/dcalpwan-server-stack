package joinserver_test

import (
	"strconv"

	"go.thethings.network/lorawan-stack/v3/pkg/provisioning"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"google.golang.org/protobuf/types/known/structpb"
)

var Timeout = (1 << 8) * test.Delay

type byteToSerialNumber struct{}

func (p *byteToSerialNumber) UniqueID(entry *structpb.Struct) (string, error) {
	return strconv.Itoa(int(entry.Fields["serial_number"].GetNumberValue())), nil
}

func init() {
	provisioning.Register("mock", &byteToSerialNumber{})
	provisioning.Register("mock-provisioner-1", &byteToSerialNumber{})
	provisioning.Register("mock-provisioner-2", &byteToSerialNumber{})
	provisioning.Register("mock-provisioner-3", &byteToSerialNumber{})
}

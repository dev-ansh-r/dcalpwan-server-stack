package provisioning_test

import (
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/provisioning"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestMicrochip(t *testing.T) {
	a := assertions.New(t)

	provisioner := Get(Microchip)
	if !a.So(provisioner, should.NotBeNil) {
		t.FailNow()
	}

	entry := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"uniqueId": {
				Kind: &structpb.Value_StringValue{
					StringValue: "abcd",
				},
			},
		},
	}

	uniqueID, err := provisioner.UniqueID(entry)
	a.So(err, should.BeNil)
	a.So(uniqueID, should.Equal, "ABCD")
}

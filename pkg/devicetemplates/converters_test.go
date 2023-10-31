package devicetemplates_test

import (
	"context"
	"io"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/devicetemplates"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestConverters(t *testing.T) {
	a := assertions.New(t)

	a.So(GetConverter("test"), should.BeNil)

	converter := &mockConverter{
		EndDeviceTemplateFormat: ttnpb.EndDeviceTemplateFormat{
			Name:        "Foo",
			Description: "Bar",
		},
	}
	RegisterConverter("test", converter)
	a.So(GetConverter("test"), should.Equal, converter)
}

type mockConverter struct {
	ttnpb.EndDeviceTemplateFormat
}

func (c *mockConverter) Format() *ttnpb.EndDeviceTemplateFormat {
	return &c.EndDeviceTemplateFormat
}

func (*mockConverter) Convert(context.Context, io.Reader, func(*ttnpb.EndDeviceTemplate) error) error {
	return nil
}

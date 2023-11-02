package devicetemplates_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func validateTemplate(t *testing.T, tmpl *ttnpb.EndDeviceTemplate) {
	a := assertions.New(t)
	if !a.So(tmpl, should.NotBeNil) {
		t.FailNow()
	}

	dev := &ttnpb.EndDevice{}
	a.So(dev.SetFields(tmpl.EndDevice, tmpl.FieldMask.GetPaths()...), should.BeNil)
	a.So(dev, should.Resemble, tmpl.EndDevice)
}

func validateTemplates(t *testing.T, templates []*ttnpb.EndDeviceTemplate, count int) {
	a := assertions.New(t)

	if !a.So(len(templates), should.Equal, count) {
		t.FailNow()
	}

	for _, template := range templates {
		validateTemplate(t, template)
	}
}

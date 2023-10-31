package devicetemplateconverter_test

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/component"
	componenttest "go.thethings.network/lorawan-stack/v3/pkg/component/test"
	. "go.thethings.network/lorawan-stack/v3/pkg/devicetemplateconverter"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

func TestDeviceTemplateConverter(t *testing.T) {
	t.Parallel()
	ctx := log.NewContext(test.Context(), test.GetLogger(t))

	conf := &component.Config{}
	c := componenttest.NewComponent(t, conf)

	test.Must(New(c, &Config{}))
	componenttest.StartComponent(t, c)
	defer c.Close()

	mustHavePeer(ctx, c, ttnpb.ClusterRole_DEVICE_TEMPLATE_CONVERTER)
}

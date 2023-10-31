package devicetemplateconverter_test

import (
	"context"
	"io"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func mustHavePeer(ctx context.Context, c *component.Component, role ttnpb.ClusterRole) {
	for i := 0; i < 20; i++ {
		time.Sleep(20 * time.Millisecond)
		if _, err := c.GetPeer(ctx, role, nil); err == nil {
			return
		}
	}
	panic("could not connect to peer")
}

type mockConverter struct {
	ttnpb.EndDeviceTemplateFormat
	ConvertFunc func(context.Context, io.Reader, func(*ttnpb.EndDeviceTemplate) error) error
}

func (c *mockConverter) Format() *ttnpb.EndDeviceTemplateFormat {
	return &c.EndDeviceTemplateFormat
}

func (c *mockConverter) Convert(ctx context.Context, r io.Reader, f func(*ttnpb.EndDeviceTemplate) error) error {
	if c.ConvertFunc == nil {
		panic("Convert should not be called")
	}
	return c.ConvertFunc(ctx, r, f)
}

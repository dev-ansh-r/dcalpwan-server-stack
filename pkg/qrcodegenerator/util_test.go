package qrcodegenerator_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/qrcodegenerator/qrcode/enddevices"
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

type mock struct {
	ids *ttnpb.EndDeviceIdentifiers
}

func (mock) Validate() error { return nil }

func (m *mock) Encode(dev *ttnpb.EndDevice) error {
	*m = mock{
		ids: dev.Ids,
	}
	return nil
}

func (*mock) EndDeviceTemplate() *ttnpb.EndDeviceTemplate { return nil }

func (*mock) FormatID() string {
	return "test"
}

func (m mock) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf(
		"%s:%s",
		strings.ToUpper(hex.EncodeToString(m.ids.JoinEui)),
		strings.ToUpper(hex.EncodeToString(m.ids.DevEui))),
	), nil
}

func (*mock) UnmarshalText([]byte) error { return nil }

type mockFormat struct{}

func (mockFormat) Format() *ttnpb.QRCodeFormat {
	return &ttnpb.QRCodeFormat{
		Name:        "Test",
		Description: "Test",
		FieldMask:   ttnpb.FieldMask("ids"),
	}
}

func (mockFormat) New() enddevices.Data {
	return new(mock)
}

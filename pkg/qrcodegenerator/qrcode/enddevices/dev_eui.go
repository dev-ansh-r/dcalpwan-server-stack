package enddevices

import (
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

const formatIDDevEUI = "dev_eui"

type devEUIData struct {
	DevEUI types.EUI64
}

// MarshalText implements Data.
func (d *devEUIData) MarshalText() (text []byte, err error) {
	return d.DevEUI.MarshalText()
}

// UnmarshalText implements Data.
func (d *devEUIData) UnmarshalText(text []byte) error {
	return d.DevEUI.UnmarshalText(text)
}

// Validate implements Data.
func (d *devEUIData) Validate() error {
	return nil
}

// Encode implements Data.
func (d *devEUIData) Encode(dev *ttnpb.EndDevice) error {
	if dev.Ids.DevEui == nil {
		return errNoDevEUI.New()
	}
	*d = devEUIData{
		DevEUI: types.MustEUI64(dev.Ids.DevEui).OrZero(),
	}
	return nil
}

// EndDeviceTemplate implements Data.
func (d *devEUIData) EndDeviceTemplate() *ttnpb.EndDeviceTemplate {
	return &ttnpb.EndDeviceTemplate{
		EndDevice: &ttnpb.EndDevice{
			Ids: &ttnpb.EndDeviceIdentifiers{
				DevEui: d.DevEUI.Bytes(),
			},
		},
		FieldMask: ttnpb.FieldMask("ids.dev_eui"),
	}
}

// FormatID implements Data.
func (d *devEUIData) FormatID() string {
	return formatIDDevEUI
}

type devEUIFormat struct{}

// Format implements Format.
func (*devEUIFormat) Format() *ttnpb.QRCodeFormat {
	return &ttnpb.QRCodeFormat{
		Name:        "DevEUI",
		Description: "LoRaWAN® DevEUI (hex encoded)",
		FieldMask:   ttnpb.FieldMask("ids.dev_eui"),
	}
}

// New implements Format.
func (*devEUIFormat) New() Data {
	return new(devEUIData)
}

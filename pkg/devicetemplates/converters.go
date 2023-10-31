package devicetemplates

import (
	"context"
	"io"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Converter converts a binary file in end device templates.
type Converter interface {
	Format() *ttnpb.EndDeviceTemplateFormat
	Convert(context.Context, io.Reader, func(*ttnpb.EndDeviceTemplate) error) error
}

var converters = map[string]Converter{}

// GetConverter returns the converter by ID.
func GetConverter(id string) Converter {
	return converters[id]
}

// RegisterConverter registers the given converter.
// Existing registrations with the same ID will be overwritten.
// This function is not goroutine-safe.
func RegisterConverter(id string, c Converter) {
	converters[id] = c
}

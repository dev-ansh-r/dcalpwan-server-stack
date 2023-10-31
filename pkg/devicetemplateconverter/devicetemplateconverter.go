// Package devicetemplateconverter provides device template services.
package devicetemplateconverter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/devicetemplates"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

// DeviceTemplateConverter implements the Device Template Converter component.
//
// The Device Template Converter exposes the EndDeviceTemplateConverter service.
type DeviceTemplateConverter struct {
	*component.Component
	ctx context.Context

	converters map[string]devicetemplates.Converter

	grpc struct {
		endDeviceTemplateConverter *endDeviceTemplateConverterServer
	}
}

var errNotFound = errors.DefineNotFound("converter", "converter `{id}` not found")

// New returns a new *DeviceTemplateConverter.
func New(c *component.Component, conf *Config) (*DeviceTemplateConverter, error) {
	// Always enable the TTS device template converters.
	conf.Enabled = append(conf.Enabled,
		devicetemplates.TTSJSON,
		devicetemplates.TTSCSV,
	)
	converters := make(map[string]devicetemplates.Converter, len(conf.Enabled))
	for _, id := range conf.Enabled {
		converter := devicetemplates.GetConverter(id)
		if converter == nil {
			return nil, errNotFound.WithAttributes("id", id)
		}
		converters[id] = converter
	}

	dtc := &DeviceTemplateConverter{
		Component:  c,
		ctx:        log.NewContextWithField(c.Context(), "namespace", "devicetemplateconverter"),
		converters: converters,
	}
	dtc.grpc.endDeviceTemplateConverter = &endDeviceTemplateConverterServer{DTC: dtc}

	c.RegisterGRPC(dtc)
	return dtc, nil
}

// Context returns the context of the Device Template Converter.
func (dtc *DeviceTemplateConverter) Context() context.Context {
	return dtc.ctx
}

// Roles returns the roles that the Device Template Converter fulfills.
func (*DeviceTemplateConverter) Roles() []ttnpb.ClusterRole {
	return []ttnpb.ClusterRole{ttnpb.ClusterRole_DEVICE_TEMPLATE_CONVERTER}
}

// RegisterServices registers services provided by dtc at s.
func (dtc *DeviceTemplateConverter) RegisterServices(s *grpc.Server) {
	ttnpb.RegisterEndDeviceTemplateConverterServer(s, dtc.grpc.endDeviceTemplateConverter)
}

// RegisterHandlers registers gRPC handlers.
func (dtc *DeviceTemplateConverter) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {
	ttnpb.RegisterEndDeviceTemplateConverterHandler(dtc.Context(), s, conn) //nolint:errcheck
}

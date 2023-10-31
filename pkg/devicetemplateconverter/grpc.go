package devicetemplateconverter

import (
	"bytes"
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/devicetemplateconverter/profilefetcher"
	"go.thethings.network/lorawan-stack/v3/pkg/devicetemplates"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type endDeviceTemplateConverterServer struct {
	ttnpb.UnimplementedEndDeviceTemplateConverterServer

	DTC *DeviceTemplateConverter
}

// ListFormats implements ttnpb.DeviceTemplateServiceServer.
func (s *endDeviceTemplateConverterServer) ListFormats(
	context.Context,
	*emptypb.Empty,
) (*ttnpb.EndDeviceTemplateFormats, error) {
	formats := make(map[string]*ttnpb.EndDeviceTemplateFormat, len(s.DTC.converters))
	for id, converter := range s.DTC.converters {
		formats[id] = converter.Format()
	}
	return &ttnpb.EndDeviceTemplateFormats{
		Formats: formats,
	}, nil
}

// Convert implements ttnpb.DeviceTemplateServiceServer.
func (s *endDeviceTemplateConverterServer) Convert(
	req *ttnpb.ConvertEndDeviceTemplateRequest,
	res ttnpb.EndDeviceTemplateConverter_ConvertServer,
) error {
	converter, ok := s.DTC.converters[req.FormatId]
	if !ok {
		return errNotFound.WithAttributes("id", req.FormatId)
	}

	ctx := res.Context()
	ctx = devicetemplates.NewContextWithProfileIDs(ctx, req.GetEndDeviceVersionIds())
	ctx = profilefetcher.NewContextWithFetcher(ctx, profilefetcher.NewTemplateFetcher(s.DTC.Component))

	return converter.Convert(
		ctx,
		bytes.NewReader(req.Data),
		res.Send,
	)
}

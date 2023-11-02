// Package mockdr contains the mock of a Device Repository Server.
package mockdr

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

// New returns a mock of the Device Repository client.
func New() *MockDR {
	return &MockDR{}
}

// MockDR contains the response that it will provide when a client communicates with the server address.
type MockDR struct {
	Err                     error
	ListEndDeviceBrandsResp *ttnpb.ListEndDeviceBrandsResponse
	EndDeviceBrand          *ttnpb.EndDeviceBrand
	ListEndDeviceModelsResp *ttnpb.ListEndDeviceModelsResponse
	EndDeviceModel          *ttnpb.EndDeviceModel
	EndDeviceTemplate       *ttnpb.EndDeviceTemplate
	MessagePayloadDecoder   *ttnpb.MessagePayloadDecoder
	MessagePayloadEncoder   *ttnpb.MessagePayloadEncoder
}

/// Methods of DR

// ListBrands mock method.
func (mdr *MockDR) ListBrands(
	context.Context,
	*ttnpb.ListEndDeviceBrandsRequest,
	...grpc.CallOption,
) (*ttnpb.ListEndDeviceBrandsResponse, error) {
	return mdr.ListEndDeviceBrandsResp, mdr.Err
}

// GetBrand mock method.
func (mdr *MockDR) GetBrand(
	context.Context,
	*ttnpb.GetEndDeviceBrandRequest,
	...grpc.CallOption,
) (*ttnpb.EndDeviceBrand, error) {
	return mdr.EndDeviceBrand, mdr.Err
}

// ListModels mock method.
func (mdr *MockDR) ListModels(
	context.Context,
	*ttnpb.ListEndDeviceModelsRequest,
	...grpc.CallOption,
) (*ttnpb.ListEndDeviceModelsResponse, error) {
	return mdr.ListEndDeviceModelsResp, mdr.Err
}

// GetModel mock method.
func (mdr *MockDR) GetModel(
	context.Context,
	*ttnpb.GetEndDeviceModelRequest,
	...grpc.CallOption,
) (*ttnpb.EndDeviceModel, error) {
	return mdr.EndDeviceModel, mdr.Err
}

// GetTemplate mock method.
func (mdr *MockDR) GetTemplate(
	context.Context,
	*ttnpb.GetTemplateRequest,
	...grpc.CallOption,
) (*ttnpb.EndDeviceTemplate, error) {
	return mdr.EndDeviceTemplate, mdr.Err
}

// GetUplinkDecoder mock method.
func (mdr *MockDR) GetUplinkDecoder(
	context.Context,
	*ttnpb.GetPayloadFormatterRequest,
	...grpc.CallOption,
) (*ttnpb.MessagePayloadDecoder, error) {
	return mdr.MessagePayloadDecoder, mdr.Err
}

// GetDownlinkDecoder mock method.
func (mdr *MockDR) GetDownlinkDecoder(
	context.Context,
	*ttnpb.GetPayloadFormatterRequest,
	...grpc.CallOption,
) (*ttnpb.MessagePayloadDecoder, error) {
	return mdr.MessagePayloadDecoder, mdr.Err
}

// GetDownlinkEncoder mock method.
func (mdr *MockDR) GetDownlinkEncoder(
	context.Context,
	*ttnpb.GetPayloadFormatterRequest,
	...grpc.CallOption,
) (*ttnpb.MessagePayloadEncoder, error) {
	return mdr.MessagePayloadEncoder, mdr.Err
}



// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: ttn/lorawan/v3/networkserver.proto

package ttnpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Ns_GenerateDevAddr_FullMethodName          = "/ttn.lorawan.v3.Ns/GenerateDevAddr"
	Ns_GetDefaultMACSettings_FullMethodName    = "/ttn.lorawan.v3.Ns/GetDefaultMACSettings"
	Ns_GetNetID_FullMethodName                 = "/ttn.lorawan.v3.Ns/GetNetID"
	Ns_GetDeviceAddressPrefixes_FullMethodName = "/ttn.lorawan.v3.Ns/GetDeviceAddressPrefixes"
)

// NsClient is the client API for Ns service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NsClient interface {
	// GenerateDevAddr requests a device address assignment from the Network Server.
	GenerateDevAddr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GenerateDevAddrResponse, error)
	// GetDefaultMACSettings retrieves the default MAC settings for a frequency plan.
	GetDefaultMACSettings(ctx context.Context, in *GetDefaultMACSettingsRequest, opts ...grpc.CallOption) (*MACSettings, error)
	GetNetID(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNetIDResponse, error)
	GetDeviceAddressPrefixes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDeviceAdressPrefixesResponse, error)
}

type nsClient struct {
	cc grpc.ClientConnInterface
}

func NewNsClient(cc grpc.ClientConnInterface) NsClient {
	return &nsClient{cc}
}

func (c *nsClient) GenerateDevAddr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GenerateDevAddrResponse, error) {
	out := new(GenerateDevAddrResponse)
	err := c.cc.Invoke(ctx, Ns_GenerateDevAddr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsClient) GetDefaultMACSettings(ctx context.Context, in *GetDefaultMACSettingsRequest, opts ...grpc.CallOption) (*MACSettings, error) {
	out := new(MACSettings)
	err := c.cc.Invoke(ctx, Ns_GetDefaultMACSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsClient) GetNetID(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNetIDResponse, error) {
	out := new(GetNetIDResponse)
	err := c.cc.Invoke(ctx, Ns_GetNetID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsClient) GetDeviceAddressPrefixes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDeviceAdressPrefixesResponse, error) {
	out := new(GetDeviceAdressPrefixesResponse)
	err := c.cc.Invoke(ctx, Ns_GetDeviceAddressPrefixes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsServer is the server API for Ns service.
// All implementations must embed UnimplementedNsServer
// for forward compatibility
type NsServer interface {
	// GenerateDevAddr requests a device address assignment from the Network Server.
	GenerateDevAddr(context.Context, *emptypb.Empty) (*GenerateDevAddrResponse, error)
	// GetDefaultMACSettings retrieves the default MAC settings for a frequency plan.
	GetDefaultMACSettings(context.Context, *GetDefaultMACSettingsRequest) (*MACSettings, error)
	GetNetID(context.Context, *emptypb.Empty) (*GetNetIDResponse, error)
	GetDeviceAddressPrefixes(context.Context, *emptypb.Empty) (*GetDeviceAdressPrefixesResponse, error)
	mustEmbedUnimplementedNsServer()
}

// UnimplementedNsServer must be embedded to have forward compatible implementations.
type UnimplementedNsServer struct {
}

func (UnimplementedNsServer) GenerateDevAddr(context.Context, *emptypb.Empty) (*GenerateDevAddrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDevAddr not implemented")
}
func (UnimplementedNsServer) GetDefaultMACSettings(context.Context, *GetDefaultMACSettingsRequest) (*MACSettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultMACSettings not implemented")
}
func (UnimplementedNsServer) GetNetID(context.Context, *emptypb.Empty) (*GetNetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetID not implemented")
}
func (UnimplementedNsServer) GetDeviceAddressPrefixes(context.Context, *emptypb.Empty) (*GetDeviceAdressPrefixesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceAddressPrefixes not implemented")
}
func (UnimplementedNsServer) mustEmbedUnimplementedNsServer() {}

// UnsafeNsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NsServer will
// result in compilation errors.
type UnsafeNsServer interface {
	mustEmbedUnimplementedNsServer()
}

func RegisterNsServer(s grpc.ServiceRegistrar, srv NsServer) {
	s.RegisterService(&Ns_ServiceDesc, srv)
}

func _Ns_GenerateDevAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsServer).GenerateDevAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ns_GenerateDevAddr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsServer).GenerateDevAddr(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ns_GetDefaultMACSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultMACSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsServer).GetDefaultMACSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ns_GetDefaultMACSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsServer).GetDefaultMACSettings(ctx, req.(*GetDefaultMACSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ns_GetNetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsServer).GetNetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ns_GetNetID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsServer).GetNetID(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ns_GetDeviceAddressPrefixes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsServer).GetDeviceAddressPrefixes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ns_GetDeviceAddressPrefixes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsServer).GetDeviceAddressPrefixes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Ns_ServiceDesc is the grpc.ServiceDesc for Ns service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ns_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Ns",
	HandlerType: (*NsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateDevAddr",
			Handler:    _Ns_GenerateDevAddr_Handler,
		},
		{
			MethodName: "GetDefaultMACSettings",
			Handler:    _Ns_GetDefaultMACSettings_Handler,
		},
		{
			MethodName: "GetNetID",
			Handler:    _Ns_GetNetID_Handler,
		},
		{
			MethodName: "GetDeviceAddressPrefixes",
			Handler:    _Ns_GetDeviceAddressPrefixes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/networkserver.proto",
}

const (
	AsNs_DownlinkQueueReplace_FullMethodName = "/ttn.lorawan.v3.AsNs/DownlinkQueueReplace"
	AsNs_DownlinkQueuePush_FullMethodName    = "/ttn.lorawan.v3.AsNs/DownlinkQueuePush"
	AsNs_DownlinkQueueList_FullMethodName    = "/ttn.lorawan.v3.AsNs/DownlinkQueueList"
)

// AsNsClient is the client API for AsNs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AsNsClient interface {
	// Replace the entire downlink queue with the specified messages.
	// This can also be used to empty the queue by specifying no messages.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueueReplace(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Push downlink messages to the end of the downlink queue.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueuePush(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List the items currently in the downlink queue.
	DownlinkQueueList(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*ApplicationDownlinks, error)
}

type asNsClient struct {
	cc grpc.ClientConnInterface
}

func NewAsNsClient(cc grpc.ClientConnInterface) AsNsClient {
	return &asNsClient{cc}
}

func (c *asNsClient) DownlinkQueueReplace(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AsNs_DownlinkQueueReplace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asNsClient) DownlinkQueuePush(ctx context.Context, in *DownlinkQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AsNs_DownlinkQueuePush_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asNsClient) DownlinkQueueList(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*ApplicationDownlinks, error) {
	out := new(ApplicationDownlinks)
	err := c.cc.Invoke(ctx, AsNs_DownlinkQueueList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsNsServer is the server API for AsNs service.
// All implementations must embed UnimplementedAsNsServer
// for forward compatibility
type AsNsServer interface {
	// Replace the entire downlink queue with the specified messages.
	// This can also be used to empty the queue by specifying no messages.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueueReplace(context.Context, *DownlinkQueueRequest) (*emptypb.Empty, error)
	// Push downlink messages to the end of the downlink queue.
	// Note that this will trigger an immediate downlink if a downlink slot is available.
	DownlinkQueuePush(context.Context, *DownlinkQueueRequest) (*emptypb.Empty, error)
	// List the items currently in the downlink queue.
	DownlinkQueueList(context.Context, *EndDeviceIdentifiers) (*ApplicationDownlinks, error)
	mustEmbedUnimplementedAsNsServer()
}

// UnimplementedAsNsServer must be embedded to have forward compatible implementations.
type UnimplementedAsNsServer struct {
}

func (UnimplementedAsNsServer) DownlinkQueueReplace(context.Context, *DownlinkQueueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownlinkQueueReplace not implemented")
}
func (UnimplementedAsNsServer) DownlinkQueuePush(context.Context, *DownlinkQueueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownlinkQueuePush not implemented")
}
func (UnimplementedAsNsServer) DownlinkQueueList(context.Context, *EndDeviceIdentifiers) (*ApplicationDownlinks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownlinkQueueList not implemented")
}
func (UnimplementedAsNsServer) mustEmbedUnimplementedAsNsServer() {}

// UnsafeAsNsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AsNsServer will
// result in compilation errors.
type UnsafeAsNsServer interface {
	mustEmbedUnimplementedAsNsServer()
}

func RegisterAsNsServer(s grpc.ServiceRegistrar, srv AsNsServer) {
	s.RegisterService(&AsNs_ServiceDesc, srv)
}

func _AsNs_DownlinkQueueReplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueueReplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsNs_DownlinkQueueReplace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueueReplace(ctx, req.(*DownlinkQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsNs_DownlinkQueuePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueuePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsNs_DownlinkQueuePush_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueuePush(ctx, req.(*DownlinkQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsNs_DownlinkQueueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsNsServer).DownlinkQueueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsNs_DownlinkQueueList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsNsServer).DownlinkQueueList(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// AsNs_ServiceDesc is the grpc.ServiceDesc for AsNs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AsNs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.AsNs",
	HandlerType: (*AsNsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownlinkQueueReplace",
			Handler:    _AsNs_DownlinkQueueReplace_Handler,
		},
		{
			MethodName: "DownlinkQueuePush",
			Handler:    _AsNs_DownlinkQueuePush_Handler,
		},
		{
			MethodName: "DownlinkQueueList",
			Handler:    _AsNs_DownlinkQueueList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/networkserver.proto",
}

const (
	GsNs_HandleUplink_FullMethodName           = "/ttn.lorawan.v3.GsNs/HandleUplink"
	GsNs_ReportTxAcknowledgment_FullMethodName = "/ttn.lorawan.v3.GsNs/ReportTxAcknowledgment"
)

// GsNsClient is the client API for GsNs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GsNsClient interface {
	// Called by the Gateway Server when an uplink message arrives.
	HandleUplink(ctx context.Context, in *UplinkMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Called by the Gateway Server when a Tx acknowledgment arrives.
	ReportTxAcknowledgment(ctx context.Context, in *GatewayTxAcknowledgment, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gsNsClient struct {
	cc grpc.ClientConnInterface
}

func NewGsNsClient(cc grpc.ClientConnInterface) GsNsClient {
	return &gsNsClient{cc}
}

func (c *gsNsClient) HandleUplink(ctx context.Context, in *UplinkMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GsNs_HandleUplink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gsNsClient) ReportTxAcknowledgment(ctx context.Context, in *GatewayTxAcknowledgment, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GsNs_ReportTxAcknowledgment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GsNsServer is the server API for GsNs service.
// All implementations must embed UnimplementedGsNsServer
// for forward compatibility
type GsNsServer interface {
	// Called by the Gateway Server when an uplink message arrives.
	HandleUplink(context.Context, *UplinkMessage) (*emptypb.Empty, error)
	// Called by the Gateway Server when a Tx acknowledgment arrives.
	ReportTxAcknowledgment(context.Context, *GatewayTxAcknowledgment) (*emptypb.Empty, error)
	mustEmbedUnimplementedGsNsServer()
}

// UnimplementedGsNsServer must be embedded to have forward compatible implementations.
type UnimplementedGsNsServer struct {
}

func (UnimplementedGsNsServer) HandleUplink(context.Context, *UplinkMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUplink not implemented")
}
func (UnimplementedGsNsServer) ReportTxAcknowledgment(context.Context, *GatewayTxAcknowledgment) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTxAcknowledgment not implemented")
}
func (UnimplementedGsNsServer) mustEmbedUnimplementedGsNsServer() {}

// UnsafeGsNsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GsNsServer will
// result in compilation errors.
type UnsafeGsNsServer interface {
	mustEmbedUnimplementedGsNsServer()
}

func RegisterGsNsServer(s grpc.ServiceRegistrar, srv GsNsServer) {
	s.RegisterService(&GsNs_ServiceDesc, srv)
}

func _GsNs_HandleUplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UplinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsNsServer).HandleUplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GsNs_HandleUplink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsNsServer).HandleUplink(ctx, req.(*UplinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GsNs_ReportTxAcknowledgment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayTxAcknowledgment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsNsServer).ReportTxAcknowledgment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GsNs_ReportTxAcknowledgment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsNsServer).ReportTxAcknowledgment(ctx, req.(*GatewayTxAcknowledgment))
	}
	return interceptor(ctx, in, info, handler)
}

// GsNs_ServiceDesc is the grpc.ServiceDesc for GsNs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GsNs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GsNs",
	HandlerType: (*GsNsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleUplink",
			Handler:    _GsNs_HandleUplink_Handler,
		},
		{
			MethodName: "ReportTxAcknowledgment",
			Handler:    _GsNs_ReportTxAcknowledgment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/networkserver.proto",
}

const (
	NsEndDeviceRegistry_Get_FullMethodName                  = "/ttn.lorawan.v3.NsEndDeviceRegistry/Get"
	NsEndDeviceRegistry_Set_FullMethodName                  = "/ttn.lorawan.v3.NsEndDeviceRegistry/Set"
	NsEndDeviceRegistry_ResetFactoryDefaults_FullMethodName = "/ttn.lorawan.v3.NsEndDeviceRegistry/ResetFactoryDefaults"
	NsEndDeviceRegistry_Delete_FullMethodName               = "/ttn.lorawan.v3.NsEndDeviceRegistry/Delete"
)

// NsEndDeviceRegistryClient is the client API for NsEndDeviceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NsEndDeviceRegistryClient interface {
	// Get returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Set creates or updates the device.
	Set(ctx context.Context, in *SetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// ResetFactoryDefaults resets device state to factory defaults.
	ResetFactoryDefaults(ctx context.Context, in *ResetAndGetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error)
	// Delete deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nsEndDeviceRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewNsEndDeviceRegistryClient(cc grpc.ClientConnInterface) NsEndDeviceRegistryClient {
	return &nsEndDeviceRegistryClient{cc}
}

func (c *nsEndDeviceRegistryClient) Get(ctx context.Context, in *GetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, NsEndDeviceRegistry_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) Set(ctx context.Context, in *SetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, NsEndDeviceRegistry_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) ResetFactoryDefaults(ctx context.Context, in *ResetAndGetEndDeviceRequest, opts ...grpc.CallOption) (*EndDevice, error) {
	out := new(EndDevice)
	err := c.cc.Invoke(ctx, NsEndDeviceRegistry_ResetFactoryDefaults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsEndDeviceRegistryClient) Delete(ctx context.Context, in *EndDeviceIdentifiers, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NsEndDeviceRegistry_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsEndDeviceRegistryServer is the server API for NsEndDeviceRegistry service.
// All implementations must embed UnimplementedNsEndDeviceRegistryServer
// for forward compatibility
type NsEndDeviceRegistryServer interface {
	// Get returns the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error)
	// Set creates or updates the device.
	Set(context.Context, *SetEndDeviceRequest) (*EndDevice, error)
	// ResetFactoryDefaults resets device state to factory defaults.
	ResetFactoryDefaults(context.Context, *ResetAndGetEndDeviceRequest) (*EndDevice, error)
	// Delete deletes the device that matches the given identifiers.
	// If there are multiple matches, an error will be returned.
	Delete(context.Context, *EndDeviceIdentifiers) (*emptypb.Empty, error)
	mustEmbedUnimplementedNsEndDeviceRegistryServer()
}

// UnimplementedNsEndDeviceRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedNsEndDeviceRegistryServer struct {
}

func (UnimplementedNsEndDeviceRegistryServer) Get(context.Context, *GetEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNsEndDeviceRegistryServer) Set(context.Context, *SetEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedNsEndDeviceRegistryServer) ResetFactoryDefaults(context.Context, *ResetAndGetEndDeviceRequest) (*EndDevice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetFactoryDefaults not implemented")
}
func (UnimplementedNsEndDeviceRegistryServer) Delete(context.Context, *EndDeviceIdentifiers) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNsEndDeviceRegistryServer) mustEmbedUnimplementedNsEndDeviceRegistryServer() {}

// UnsafeNsEndDeviceRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NsEndDeviceRegistryServer will
// result in compilation errors.
type UnsafeNsEndDeviceRegistryServer interface {
	mustEmbedUnimplementedNsEndDeviceRegistryServer()
}

func RegisterNsEndDeviceRegistryServer(s grpc.ServiceRegistrar, srv NsEndDeviceRegistryServer) {
	s.RegisterService(&NsEndDeviceRegistry_ServiceDesc, srv)
}

func _NsEndDeviceRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsEndDeviceRegistry_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Get(ctx, req.(*GetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsEndDeviceRegistry_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Set(ctx, req.(*SetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_ResetFactoryDefaults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetAndGetEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).ResetFactoryDefaults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsEndDeviceRegistry_ResetFactoryDefaults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).ResetFactoryDefaults(ctx, req.(*ResetAndGetEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsEndDeviceRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDeviceIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsEndDeviceRegistry_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceRegistryServer).Delete(ctx, req.(*EndDeviceIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// NsEndDeviceRegistry_ServiceDesc is the grpc.ServiceDesc for NsEndDeviceRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NsEndDeviceRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsEndDeviceRegistry",
	HandlerType: (*NsEndDeviceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NsEndDeviceRegistry_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _NsEndDeviceRegistry_Set_Handler,
		},
		{
			MethodName: "ResetFactoryDefaults",
			Handler:    _NsEndDeviceRegistry_ResetFactoryDefaults_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NsEndDeviceRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/networkserver.proto",
}

const (
	NsEndDeviceBatchRegistry_Delete_FullMethodName = "/ttn.lorawan.v3.NsEndDeviceBatchRegistry/Delete"
)

// NsEndDeviceBatchRegistryClient is the client API for NsEndDeviceBatchRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NsEndDeviceBatchRegistryClient interface {
	// Delete a list of devices within the same application.
	// This operation is atomic; either all devices are deleted or none.
	// Devices not found are skipped and no error is returned.
	Delete(ctx context.Context, in *BatchDeleteEndDevicesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nsEndDeviceBatchRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewNsEndDeviceBatchRegistryClient(cc grpc.ClientConnInterface) NsEndDeviceBatchRegistryClient {
	return &nsEndDeviceBatchRegistryClient{cc}
}

func (c *nsEndDeviceBatchRegistryClient) Delete(ctx context.Context, in *BatchDeleteEndDevicesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NsEndDeviceBatchRegistry_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsEndDeviceBatchRegistryServer is the server API for NsEndDeviceBatchRegistry service.
// All implementations must embed UnimplementedNsEndDeviceBatchRegistryServer
// for forward compatibility
type NsEndDeviceBatchRegistryServer interface {
	// Delete a list of devices within the same application.
	// This operation is atomic; either all devices are deleted or none.
	// Devices not found are skipped and no error is returned.
	Delete(context.Context, *BatchDeleteEndDevicesRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNsEndDeviceBatchRegistryServer()
}

// UnimplementedNsEndDeviceBatchRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedNsEndDeviceBatchRegistryServer struct {
}

func (UnimplementedNsEndDeviceBatchRegistryServer) Delete(context.Context, *BatchDeleteEndDevicesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNsEndDeviceBatchRegistryServer) mustEmbedUnimplementedNsEndDeviceBatchRegistryServer() {
}

// UnsafeNsEndDeviceBatchRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NsEndDeviceBatchRegistryServer will
// result in compilation errors.
type UnsafeNsEndDeviceBatchRegistryServer interface {
	mustEmbedUnimplementedNsEndDeviceBatchRegistryServer()
}

func RegisterNsEndDeviceBatchRegistryServer(s grpc.ServiceRegistrar, srv NsEndDeviceBatchRegistryServer) {
	s.RegisterService(&NsEndDeviceBatchRegistry_ServiceDesc, srv)
}

func _NsEndDeviceBatchRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteEndDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsEndDeviceBatchRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsEndDeviceBatchRegistry_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsEndDeviceBatchRegistryServer).Delete(ctx, req.(*BatchDeleteEndDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NsEndDeviceBatchRegistry_ServiceDesc is the grpc.ServiceDesc for NsEndDeviceBatchRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NsEndDeviceBatchRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsEndDeviceBatchRegistry",
	HandlerType: (*NsEndDeviceBatchRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _NsEndDeviceBatchRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/networkserver.proto",
}

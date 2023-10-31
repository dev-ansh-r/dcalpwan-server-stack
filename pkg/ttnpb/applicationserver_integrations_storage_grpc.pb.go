

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: ttn/lorawan/v3/applicationserver_integrations_storage.proto

package ttnpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApplicationUpStorage_GetStoredApplicationUp_FullMethodName      = "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUp"
	ApplicationUpStorage_GetStoredApplicationUpCount_FullMethodName = "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUpCount"
)

// ApplicationUpStorageClient is the client API for ApplicationUpStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationUpStorageClient interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error)
	// Returns how many application messages have been stored in the database for an application or end device.
	GetStoredApplicationUpCount(ctx context.Context, in *GetStoredApplicationUpCountRequest, opts ...grpc.CallOption) (*GetStoredApplicationUpCountResponse, error)
}

type applicationUpStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationUpStorageClient(cc grpc.ClientConnInterface) ApplicationUpStorageClient {
	return &applicationUpStorageClient{cc}
}

func (c *applicationUpStorageClient) GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApplicationUpStorage_ServiceDesc.Streams[0], ApplicationUpStorage_GetStoredApplicationUp_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationUpStorageGetStoredApplicationUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationUpStorage_GetStoredApplicationUpClient interface {
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type applicationUpStorageGetStoredApplicationUpClient struct {
	grpc.ClientStream
}

func (x *applicationUpStorageGetStoredApplicationUpClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *applicationUpStorageClient) GetStoredApplicationUpCount(ctx context.Context, in *GetStoredApplicationUpCountRequest, opts ...grpc.CallOption) (*GetStoredApplicationUpCountResponse, error) {
	out := new(GetStoredApplicationUpCountResponse)
	err := c.cc.Invoke(ctx, ApplicationUpStorage_GetStoredApplicationUpCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationUpStorageServer is the server API for ApplicationUpStorage service.
// All implementations must embed UnimplementedApplicationUpStorageServer
// for forward compatibility
type ApplicationUpStorageServer interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(*GetStoredApplicationUpRequest, ApplicationUpStorage_GetStoredApplicationUpServer) error
	// Returns how many application messages have been stored in the database for an application or end device.
	GetStoredApplicationUpCount(context.Context, *GetStoredApplicationUpCountRequest) (*GetStoredApplicationUpCountResponse, error)
	mustEmbedUnimplementedApplicationUpStorageServer()
}

// UnimplementedApplicationUpStorageServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationUpStorageServer struct {
}

func (UnimplementedApplicationUpStorageServer) GetStoredApplicationUp(*GetStoredApplicationUpRequest, ApplicationUpStorage_GetStoredApplicationUpServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStoredApplicationUp not implemented")
}
func (UnimplementedApplicationUpStorageServer) GetStoredApplicationUpCount(context.Context, *GetStoredApplicationUpCountRequest) (*GetStoredApplicationUpCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoredApplicationUpCount not implemented")
}
func (UnimplementedApplicationUpStorageServer) mustEmbedUnimplementedApplicationUpStorageServer() {}

// UnsafeApplicationUpStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationUpStorageServer will
// result in compilation errors.
type UnsafeApplicationUpStorageServer interface {
	mustEmbedUnimplementedApplicationUpStorageServer()
}

func RegisterApplicationUpStorageServer(s grpc.ServiceRegistrar, srv ApplicationUpStorageServer) {
	s.RegisterService(&ApplicationUpStorage_ServiceDesc, srv)
}

func _ApplicationUpStorage_GetStoredApplicationUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStoredApplicationUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationUpStorageServer).GetStoredApplicationUp(m, &applicationUpStorageGetStoredApplicationUpServer{stream})
}

type ApplicationUpStorage_GetStoredApplicationUpServer interface {
	Send(*ApplicationUp) error
	grpc.ServerStream
}

type applicationUpStorageGetStoredApplicationUpServer struct {
	grpc.ServerStream
}

func (x *applicationUpStorageGetStoredApplicationUpServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

func _ApplicationUpStorage_GetStoredApplicationUpCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoredApplicationUpCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationUpStorageServer).GetStoredApplicationUpCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationUpStorage_GetStoredApplicationUpCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationUpStorageServer).GetStoredApplicationUpCount(ctx, req.(*GetStoredApplicationUpCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApplicationUpStorage_ServiceDesc is the grpc.ServiceDesc for ApplicationUpStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationUpStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationUpStorage",
	HandlerType: (*ApplicationUpStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStoredApplicationUpCount",
			Handler:    _ApplicationUpStorage_GetStoredApplicationUpCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStoredApplicationUp",
			Handler:       _ApplicationUpStorage_GetStoredApplicationUp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ttn/lorawan/v3/applicationserver_integrations_storage.proto",
}

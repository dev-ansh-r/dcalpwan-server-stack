

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: ttn/lorawan/v3/client_services.proto

package ttnpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ttn_lorawan_v3_client_services_proto protoreflect.FileDescriptor

var file_ttn_lorawan_v3_client_services_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76,
	0x33, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcc, 0x07,
	0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0xcf, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x87, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x80, 0x01, 0x3a, 0x01, 0x2a, 0x5a, 0x4b, 0x3a, 0x01, 0x2a, 0x22, 0x46, 0x2f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x68, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xd3, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x8d, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x86, 0x01, 0x5a, 0x30, 0x12, 0x2e,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5a, 0x48,
	0x12, 0x46, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x08, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x71, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x24, 0x3a, 0x01, 0x2a, 0x1a, 0x1f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x61, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x21, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x2a, 0x14, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x1c, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x66, 0x0a, 0x05, 0x50, 0x75, 0x72, 0x67, 0x65, 0x12, 0x21, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x2a, 0x1a, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x32, 0x99, 0x08, 0x0a,
	0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x6c, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x16,
	0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0xb4, 0x02, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x2c, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xc2, 0x01,
	0x5a, 0x53, 0x12, 0x51, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x5a, 0x6b, 0x12, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x38, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x1a, 0x2d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x99, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2e, 0x2e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x35, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0xb3, 0x02, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2f, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0xd3, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xcc, 0x01, 0x5a, 0x58, 0x2a, 0x56,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x5a, 0x70, 0x2a, 0x6e, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76,
	0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x74, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_ttn_lorawan_v3_client_services_proto_goTypes = []interface{}{
	(*CreateClientRequest)(nil),             // 0: ttn.lorawan.v3.CreateClientRequest
	(*GetClientRequest)(nil),                // 1: ttn.lorawan.v3.GetClientRequest
	(*ListClientsRequest)(nil),              // 2: ttn.lorawan.v3.ListClientsRequest
	(*UpdateClientRequest)(nil),             // 3: ttn.lorawan.v3.UpdateClientRequest
	(*ClientIdentifiers)(nil),               // 4: ttn.lorawan.v3.ClientIdentifiers
	(*GetClientCollaboratorRequest)(nil),    // 5: ttn.lorawan.v3.GetClientCollaboratorRequest
	(*SetClientCollaboratorRequest)(nil),    // 6: ttn.lorawan.v3.SetClientCollaboratorRequest
	(*ListClientCollaboratorsRequest)(nil),  // 7: ttn.lorawan.v3.ListClientCollaboratorsRequest
	(*DeleteClientCollaboratorRequest)(nil), // 8: ttn.lorawan.v3.DeleteClientCollaboratorRequest
	(*Client)(nil),                          // 9: ttn.lorawan.v3.Client
	(*Clients)(nil),                         // 10: ttn.lorawan.v3.Clients
	(*emptypb.Empty)(nil),                   // 11: google.protobuf.Empty
	(*Rights)(nil),                          // 12: ttn.lorawan.v3.Rights
	(*GetCollaboratorResponse)(nil),         // 13: ttn.lorawan.v3.GetCollaboratorResponse
	(*Collaborators)(nil),                   // 14: ttn.lorawan.v3.Collaborators
}
var file_ttn_lorawan_v3_client_services_proto_depIdxs = []int32{
	0,  // 0: ttn.lorawan.v3.ClientRegistry.Create:input_type -> ttn.lorawan.v3.CreateClientRequest
	1,  // 1: ttn.lorawan.v3.ClientRegistry.Get:input_type -> ttn.lorawan.v3.GetClientRequest
	2,  // 2: ttn.lorawan.v3.ClientRegistry.List:input_type -> ttn.lorawan.v3.ListClientsRequest
	3,  // 3: ttn.lorawan.v3.ClientRegistry.Update:input_type -> ttn.lorawan.v3.UpdateClientRequest
	4,  // 4: ttn.lorawan.v3.ClientRegistry.Delete:input_type -> ttn.lorawan.v3.ClientIdentifiers
	4,  // 5: ttn.lorawan.v3.ClientRegistry.Restore:input_type -> ttn.lorawan.v3.ClientIdentifiers
	4,  // 6: ttn.lorawan.v3.ClientRegistry.Purge:input_type -> ttn.lorawan.v3.ClientIdentifiers
	4,  // 7: ttn.lorawan.v3.ClientAccess.ListRights:input_type -> ttn.lorawan.v3.ClientIdentifiers
	5,  // 8: ttn.lorawan.v3.ClientAccess.GetCollaborator:input_type -> ttn.lorawan.v3.GetClientCollaboratorRequest
	6,  // 9: ttn.lorawan.v3.ClientAccess.SetCollaborator:input_type -> ttn.lorawan.v3.SetClientCollaboratorRequest
	7,  // 10: ttn.lorawan.v3.ClientAccess.ListCollaborators:input_type -> ttn.lorawan.v3.ListClientCollaboratorsRequest
	8,  // 11: ttn.lorawan.v3.ClientAccess.DeleteCollaborator:input_type -> ttn.lorawan.v3.DeleteClientCollaboratorRequest
	9,  // 12: ttn.lorawan.v3.ClientRegistry.Create:output_type -> ttn.lorawan.v3.Client
	9,  // 13: ttn.lorawan.v3.ClientRegistry.Get:output_type -> ttn.lorawan.v3.Client
	10, // 14: ttn.lorawan.v3.ClientRegistry.List:output_type -> ttn.lorawan.v3.Clients
	9,  // 15: ttn.lorawan.v3.ClientRegistry.Update:output_type -> ttn.lorawan.v3.Client
	11, // 16: ttn.lorawan.v3.ClientRegistry.Delete:output_type -> google.protobuf.Empty
	11, // 17: ttn.lorawan.v3.ClientRegistry.Restore:output_type -> google.protobuf.Empty
	11, // 18: ttn.lorawan.v3.ClientRegistry.Purge:output_type -> google.protobuf.Empty
	12, // 19: ttn.lorawan.v3.ClientAccess.ListRights:output_type -> ttn.lorawan.v3.Rights
	13, // 20: ttn.lorawan.v3.ClientAccess.GetCollaborator:output_type -> ttn.lorawan.v3.GetCollaboratorResponse
	11, // 21: ttn.lorawan.v3.ClientAccess.SetCollaborator:output_type -> google.protobuf.Empty
	14, // 22: ttn.lorawan.v3.ClientAccess.ListCollaborators:output_type -> ttn.lorawan.v3.Collaborators
	11, // 23: ttn.lorawan.v3.ClientAccess.DeleteCollaborator:output_type -> google.protobuf.Empty
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ttn_lorawan_v3_client_services_proto_init() }
func file_ttn_lorawan_v3_client_services_proto_init() {
	if File_ttn_lorawan_v3_client_services_proto != nil {
		return
	}
	file_ttn_lorawan_v3_client_proto_init()
	file_ttn_lorawan_v3_identifiers_proto_init()
	file_ttn_lorawan_v3_rights_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ttn_lorawan_v3_client_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_ttn_lorawan_v3_client_services_proto_goTypes,
		DependencyIndexes: file_ttn_lorawan_v3_client_services_proto_depIdxs,
	}.Build()
	File_ttn_lorawan_v3_client_services_proto = out.File
	file_ttn_lorawan_v3_client_services_proto_rawDesc = nil
	file_ttn_lorawan_v3_client_services_proto_goTypes = nil
	file_ttn_lorawan_v3_client_services_proto_depIdxs = nil
}

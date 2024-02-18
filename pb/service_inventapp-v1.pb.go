// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: service_inventapp-v1.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	account "github.com/nicodanke/inventapp_v2/pb/requests/v1/account"
	login "github.com/nicodanke/inventapp_v2/pb/requests/v1/login"
	role "github.com/nicodanke/inventapp_v2/pb/requests/v1/role"
	user "github.com/nicodanke/inventapp_v2/pb/requests/v1/user"
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

var File_service_inventapp_v1_proto protoreflect.FileDescriptor

var file_service_inventapp_v1_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x61, 0x70, 0x70, 0x2d, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x0e, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x41, 0x70, 0x70, 0x56, 0x31, 0x12, 0xc9, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x92,
	0x41, 0x4c, 0x12, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x3e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x28, 0x50, 0x41, 0x53, 0x45, 0x54, 0x4f, 0x29, 0x20, 0x74,
	0x6f, 0x20, 0x6d, 0x61, 0x6b, 0x65, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0xec, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70,
	0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6a, 0x92, 0x41, 0x50, 0x12, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x3e, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65,
	0x77, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x62, 0x61,
	0x73, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0xdf, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f,
	0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x92, 0x41, 0x3e, 0x12, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x2c, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x20, 0x74, 0x6f, 0x20, 0x6d, 0x61, 0x6b, 0x65, 0x20, 0x61, 0x20, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6e, 0x20,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a,
	0x32, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0xc1, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x2e, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x54, 0x92, 0x41, 0x40, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x2f, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x20, 0x74, 0x6f, 0x20,
	0x67, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x20, 0x41, 0x75, 0x74, 0x68, 0x20, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0xc2, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61,
	0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x38,
	0x12, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x29, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x20, 0x41, 0x75, 0x74, 0x68, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01,
	0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0xb1, 0x01, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x59, 0x92, 0x41, 0x40, 0x12, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x1a, 0x2b, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x2e, 0x20, 0x41, 0x75, 0x74,
	0x68, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0xc1, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x2e, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54,
	0x92, 0x41, 0x40, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x1a, 0x2f, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74,
	0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x20, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x20, 0x41, 0x75, 0x74, 0x68, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0xc2, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x30, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f,
	0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70,
	0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x38, 0x12, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x52, 0x6f, 0x6c, 0x65, 0x1a, 0x29, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65,
	0x77, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x20, 0x41, 0x75, 0x74, 0x68, 0x20, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0xb1, 0x01, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x30, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x59, 0x92, 0x41, 0x40, 0x12, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x72, 0x6f, 0x6c, 0x65, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x1a, 0x2b, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x72, 0x6f,
	0x6c, 0x65, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x2e, 0x20, 0x41, 0x75, 0x74, 0x68, 0x20, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x98, 0x01,
	0x92, 0x41, 0x6f, 0x12, 0x6d, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70,
	0x20, 0x41, 0x50, 0x49, 0x22, 0x57, 0x0a, 0x13, 0x4e, 0x69, 0x63, 0x6f, 0x6c, 0xc3, 0xa1, 0x73,
	0x20, 0x44, 0x61, 0x6e, 0x6b, 0x69, 0x65, 0x77, 0x69, 0x63, 0x7a, 0x12, 0x29, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x69, 0x63, 0x6f, 0x64, 0x61, 0x6e, 0x6b, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x61, 0x70, 0x70, 0x5f, 0x76, 0x32, 0x1a, 0x15, 0x6e, 0x69, 0x63, 0x6f, 0x64, 0x61, 0x6e, 0x6b,
	0x69, 0x40, 0x68, 0x6f, 0x74, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31,
	0x2e, 0x30, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x69, 0x63, 0x6f, 0x64, 0x61, 0x6e, 0x6b, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61,
	0x70, 0x70, 0x5f, 0x76, 0x32, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_inventapp_v1_proto_goTypes = []interface{}{
	(*login.LoginRequest)(nil),            // 0: inventapp_v2.requests.v1.login.LoginRequest
	(*account.CreateAccountRequest)(nil),  // 1: inventapp_v2.requests.v1.account.CreateAccountRequest
	(*account.UpdateAccountRequest)(nil),  // 2: inventapp_v2.requests.v1.account.UpdateAccountRequest
	(*user.GetUsersRequest)(nil),          // 3: inventapp_v2.requests.v1.user.GetUsersRequest
	(*user.CreateUserRequest)(nil),        // 4: inventapp_v2.requests.v1.user.CreateUserRequest
	(*user.DeleteUserRequest)(nil),        // 5: inventapp_v2.requests.v1.user.DeleteUserRequest
	(*role.GetRolesRequest)(nil),          // 6: inventapp_v2.requests.v1.role.GetRolesRequest
	(*role.CreateRoleRequest)(nil),        // 7: inventapp_v2.requests.v1.role.CreateRoleRequest
	(*role.DeleteRoleRequest)(nil),        // 8: inventapp_v2.requests.v1.role.DeleteRoleRequest
	(*login.LoginResponse)(nil),           // 9: inventapp_v2.requests.v1.login.LoginResponse
	(*account.CreateAccountResponse)(nil), // 10: inventapp_v2.requests.v1.account.CreateAccountResponse
	(*account.UpdateAccountResponse)(nil), // 11: inventapp_v2.requests.v1.account.UpdateAccountResponse
	(*user.GetUsersResponse)(nil),         // 12: inventapp_v2.requests.v1.user.GetUsersResponse
	(*user.CreateUserResponse)(nil),       // 13: inventapp_v2.requests.v1.user.CreateUserResponse
	(*emptypb.Empty)(nil),                 // 14: google.protobuf.Empty
	(*role.GetRolesResponse)(nil),         // 15: inventapp_v2.requests.v1.role.GetRolesResponse
	(*role.CreateRoleResponse)(nil),       // 16: inventapp_v2.requests.v1.role.CreateRoleResponse
}
var file_service_inventapp_v1_proto_depIdxs = []int32{
	0,  // 0: pb.InventAppV1.Login:input_type -> inventapp_v2.requests.v1.login.LoginRequest
	1,  // 1: pb.InventAppV1.CreateAccount:input_type -> inventapp_v2.requests.v1.account.CreateAccountRequest
	2,  // 2: pb.InventAppV1.UpdateAccount:input_type -> inventapp_v2.requests.v1.account.UpdateAccountRequest
	3,  // 3: pb.InventAppV1.GetUsers:input_type -> inventapp_v2.requests.v1.user.GetUsersRequest
	4,  // 4: pb.InventAppV1.CreateUser:input_type -> inventapp_v2.requests.v1.user.CreateUserRequest
	5,  // 5: pb.InventAppV1.DeleteUser:input_type -> inventapp_v2.requests.v1.user.DeleteUserRequest
	6,  // 6: pb.InventAppV1.GetRoles:input_type -> inventapp_v2.requests.v1.role.GetRolesRequest
	7,  // 7: pb.InventAppV1.CreateRole:input_type -> inventapp_v2.requests.v1.role.CreateRoleRequest
	8,  // 8: pb.InventAppV1.DeleteRole:input_type -> inventapp_v2.requests.v1.role.DeleteRoleRequest
	9,  // 9: pb.InventAppV1.Login:output_type -> inventapp_v2.requests.v1.login.LoginResponse
	10, // 10: pb.InventAppV1.CreateAccount:output_type -> inventapp_v2.requests.v1.account.CreateAccountResponse
	11, // 11: pb.InventAppV1.UpdateAccount:output_type -> inventapp_v2.requests.v1.account.UpdateAccountResponse
	12, // 12: pb.InventAppV1.GetUsers:output_type -> inventapp_v2.requests.v1.user.GetUsersResponse
	13, // 13: pb.InventAppV1.CreateUser:output_type -> inventapp_v2.requests.v1.user.CreateUserResponse
	14, // 14: pb.InventAppV1.DeleteUser:output_type -> google.protobuf.Empty
	15, // 15: pb.InventAppV1.GetRoles:output_type -> inventapp_v2.requests.v1.role.GetRolesResponse
	16, // 16: pb.InventAppV1.CreateRole:output_type -> inventapp_v2.requests.v1.role.CreateRoleResponse
	14, // 17: pb.InventAppV1.DeleteRole:output_type -> google.protobuf.Empty
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_inventapp_v1_proto_init() }
func file_service_inventapp_v1_proto_init() {
	if File_service_inventapp_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_inventapp_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_inventapp_v1_proto_goTypes,
		DependencyIndexes: file_service_inventapp_v1_proto_depIdxs,
	}.Build()
	File_service_inventapp_v1_proto = out.File
	file_service_inventapp_v1_proto_rawDesc = nil
	file_service_inventapp_v1_proto_goTypes = nil
	file_service_inventapp_v1_proto_depIdxs = nil
}

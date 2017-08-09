// Code generated by protoc-gen-go.
// source: health.proto
// DO NOT EDIT!

/*
Package grpc_health_v1 is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	HealthCheckRequest
	HealthCheckResponse
*/
package grpc_health_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=grpc.health.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "grpc.health.v1.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "grpc.health.v1.HealthCheckResponse")
	proto.RegisterEnum("grpc.health.v1.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/grpc.health.v1.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HealthServer).Check(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0x2f, 0x2a, 0x48, 0xd6, 0x83,
	0x0a, 0x95, 0x19, 0x26, 0xe6, 0x14, 0x64, 0x24, 0x2a, 0xe9, 0x71, 0x09, 0x79, 0x80, 0x45, 0x9c,
	0x33, 0x52, 0x93, 0xb3, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x8b,
	0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0xa5,
	0x85, 0x8c, 0x5c, 0xc2, 0x28, 0x1a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xfc, 0xb8, 0xd8,
	0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xc1, 0x1a, 0xf8, 0x8c, 0xcc, 0xf4, 0xb0, 0xd8, 0xa6, 0x87,
	0x45, 0xa7, 0x5e, 0x30, 0xc8, 0xe4, 0xbc, 0xf4, 0x60, 0xb0, 0xee, 0x20, 0xa8, 0x29, 0x4a, 0x56,
	0x5c, 0xbc, 0x28, 0x12, 0x42, 0xdc, 0x5c, 0xec, 0xa1, 0x7e, 0xde, 0x7e, 0xfe, 0xe1, 0x7e, 0x02,
	0x0c, 0x20, 0x4e, 0xb0, 0x6b, 0x50, 0x98, 0xa7, 0x9f, 0xbb, 0x00, 0xa3, 0x10, 0x3f, 0x17, 0xb7,
	0x9f, 0x7f, 0x48, 0x3c, 0x4c, 0x80, 0xc9, 0x28, 0x85, 0x8b, 0x0d, 0x62, 0x91, 0x50, 0x14, 0x17,
	0x2b, 0xd8, 0x32, 0x21, 0x75, 0xc2, 0xce, 0x01, 0xfb, 0x5c, 0x4a, 0x83, 0x58, 0x77, 0x27, 0xb1,
	0x81, 0x43, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x3f, 0xd0, 0xe1, 0x65, 0x01, 0x00,
	0x00,
}

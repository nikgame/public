// Code generated by protoc-gen-go.
// source: PlatformProto.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	PlatformProto.proto

It has these top-level messages:
	DeviceInfo
	LoginRequest
	LoginReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceInfo struct {
	GameID     string `protobuf:"bytes,1,opt,name=GameID,json=gameID" json:"GameID,omitempty"`
	Channel    string `protobuf:"bytes,2,opt,name=Channel,json=channel" json:"Channel,omitempty"`
	DeviceType string `protobuf:"bytes,3,opt,name=DeviceType,json=deviceType" json:"DeviceType,omitempty"`
	IMEI       string `protobuf:"bytes,4,opt,name=IMEI,json=iMEI" json:"IMEI,omitempty"`
	MAC        string `protobuf:"bytes,5,opt,name=MAC,json=mAC" json:"MAC,omitempty"`
	IPv4       string `protobuf:"bytes,6,opt,name=IPv4,json=iPv4" json:"IPv4,omitempty"`
	IPv6       string `protobuf:"bytes,7,opt,name=IPv6,json=iPv6" json:"IPv6,omitempty"`
	OS         string `protobuf:"bytes,8,opt,name=OS,json=oS" json:"OS,omitempty"`
	Latitude   string `protobuf:"bytes,9,opt,name=Latitude,json=latitude" json:"Latitude,omitempty"`
	Longitude  string `protobuf:"bytes,10,opt,name=Longitude,json=longitude" json:"Longitude,omitempty"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto1.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoginRequest struct {
	UUID       string      `protobuf:"bytes,1,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	Token      string      `protobuf:"bytes,2,opt,name=Token,json=token" json:"Token,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,3,opt,name=DeviceInfo,json=deviceInfo" json:"DeviceInfo,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto1.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

type LoginReply struct {
	Result bool   `protobuf:"varint,1,opt,name=Result,json=result" json:"Result,omitempty"`
	UUID   string `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	Token  string `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token,omitempty"`
	Extra  string `protobuf:"bytes,4,opt,name=extra" json:"extra,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto1.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto1.RegisterType((*DeviceInfo)(nil), "proto.DeviceInfo")
	proto1.RegisterType((*LoginRequest)(nil), "proto.LoginRequest")
	proto1.RegisterType((*LoginReply)(nil), "proto.LoginReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for PlatformLogin service

type PlatformLoginClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type platformLoginClient struct {
	cc *grpc.ClientConn
}

func NewPlatformLoginClient(cc *grpc.ClientConn) PlatformLoginClient {
	return &platformLoginClient{cc}
}

func (c *platformLoginClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/proto.PlatformLogin/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlatformLogin service

type PlatformLoginServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterPlatformLoginServer(s *grpc.Server, srv PlatformLoginServer) {
	s.RegisterService(&_PlatformLogin_serviceDesc, srv)
}

func _PlatformLogin_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformLoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlatformLogin/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformLoginServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlatformLogin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PlatformLogin",
	HandlerType: (*PlatformLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _PlatformLogin_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x4b, 0x4f, 0xc2, 0x40,
	0x10, 0x96, 0xbe, 0x68, 0x47, 0x34, 0xb2, 0x18, 0xb3, 0x21, 0xc6, 0x98, 0x9e, 0x3c, 0x91, 0x88,
	0x86, 0x3b, 0x0f, 0x63, 0x9a, 0x40, 0x24, 0x05, 0x7e, 0x40, 0x85, 0x05, 0x1b, 0xda, 0x5d, 0x2c,
	0x5b, 0x22, 0x7f, 0xdc, 0xb3, 0xed, 0x74, 0x5b, 0x31, 0xf1, 0x02, 0xdf, 0x63, 0x9a, 0xf9, 0xe6,
	0x5b, 0x68, 0x4d, 0xa3, 0x40, 0xae, 0x45, 0x12, 0x4f, 0x13, 0x21, 0x45, 0x67, 0x97, 0xff, 0x12,
	0x13, 0xff, 0xdc, 0xef, 0x1a, 0xc0, 0x88, 0x1d, 0xc2, 0x25, 0xf3, 0xf8, 0x5a, 0x90, 0x1b, 0xb0,
	0x5e, 0x83, 0x98, 0x79, 0x23, 0x5a, 0xbb, 0xaf, 0x3d, 0x38, 0xbe, 0xb5, 0x41, 0x46, 0x28, 0xd4,
	0x87, 0x1f, 0x01, 0xe7, 0x2c, 0xa2, 0x1a, 0x1a, 0xf5, 0x65, 0x41, 0xc9, 0x5d, 0xf9, 0xfd, 0xfc,
	0xb8, 0x63, 0x54, 0x47, 0x13, 0x56, 0x95, 0x42, 0x08, 0x18, 0xde, 0xe4, 0xc5, 0xa3, 0x06, 0x3a,
	0x46, 0x98, 0x61, 0x72, 0x05, 0xfa, 0xa4, 0x3f, 0xa4, 0x26, 0x4a, 0x7a, 0xdc, 0x1f, 0xe2, 0xd4,
	0xf4, 0xf0, 0x4c, 0x2d, 0x35, 0x95, 0x61, 0xa5, 0xf5, 0x68, 0xbd, 0xd2, 0x7a, 0xe4, 0x12, 0xb4,
	0xb7, 0x19, 0xb5, 0x51, 0xd1, 0xc4, 0x8c, 0xb4, 0xc1, 0x1e, 0x07, 0x32, 0x94, 0xe9, 0x8a, 0x51,
	0x07, 0x55, 0x3b, 0x52, 0x9c, 0xdc, 0x82, 0x33, 0x16, 0x7c, 0x53, 0x98, 0x80, 0xa6, 0x13, 0x95,
	0x82, 0xbb, 0x85, 0xc6, 0x58, 0x6c, 0x42, 0xee, 0xb3, 0xcf, 0x94, 0xed, 0x65, 0xbe, 0x6d, 0xb1,
	0xa8, 0xee, 0x36, 0xd2, 0x0c, 0x93, 0x6b, 0x30, 0xe7, 0x62, 0xcb, 0xb8, 0xba, 0xd9, 0x94, 0x39,
	0x21, 0x8f, 0xa7, 0x8d, 0xe1, 0xc5, 0xe7, 0xdd, 0x66, 0xd1, 0x6a, 0xe7, 0xd7, 0x28, 0x4b, 0xc8,
	0xb1, 0xbb, 0x02, 0x50, 0xcb, 0x76, 0xd1, 0x31, 0x2f, 0xd9, 0x67, 0xfb, 0x34, 0x92, 0xb8, 0xcc,
	0xf6, 0xad, 0x04, 0x59, 0x15, 0x41, 0xfb, 0x2f, 0x82, 0x7e, 0x1a, 0x21, 0x53, 0xd9, 0x97, 0x4c,
	0x02, 0xd5, 0x6a, 0x41, 0xba, 0x03, 0xb8, 0x28, 0x5f, 0x1a, 0xb7, 0x65, 0x49, 0xcd, 0x02, 0xb4,
	0x54, 0xbc, 0xd3, 0x8b, 0xdb, 0xcd, 0xbf, 0x62, 0x96, 0xcc, 0x3d, 0x1b, 0x34, 0x00, 0x96, 0x22,
	0xee, 0xf0, 0x6d, 0xfe, 0xf0, 0xef, 0x16, 0x4e, 0x3c, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30,
	0x8f, 0x69, 0x87, 0x42, 0x02, 0x00, 0x00,
}
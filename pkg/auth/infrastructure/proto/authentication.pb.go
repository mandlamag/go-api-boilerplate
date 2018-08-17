// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

package proto

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GetTokenRequest is passed when dispatching
type GetTokenRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_157936222af7c30e, []int{0}
}
func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (dst *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(dst, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// GetTokenResponse is empty.
type GetTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenResponse) Reset()         { *m = GetTokenResponse{} }
func (m *GetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokenResponse) ProtoMessage()    {}
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_157936222af7c30e, []int{1}
}
func (m *GetTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResponse.Unmarshal(m, b)
}
func (m *GetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResponse.Marshal(b, m, deterministic)
}
func (dst *GetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResponse.Merge(dst, src)
}
func (m *GetTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokenResponse.Size(m)
}
func (m *GetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResponse proto.InternalMessageInfo

func (m *GetTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// RefreshTokenRequest is passed when dispatching
type RefreshTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenRequest) Reset()         { *m = RefreshTokenRequest{} }
func (m *RefreshTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenRequest) ProtoMessage()    {}
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_157936222af7c30e, []int{2}
}
func (m *RefreshTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenRequest.Unmarshal(m, b)
}
func (m *RefreshTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenRequest.Marshal(b, m, deterministic)
}
func (dst *RefreshTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenRequest.Merge(dst, src)
}
func (m *RefreshTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenRequest.Size(m)
}
func (m *RefreshTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenRequest proto.InternalMessageInfo

func (m *RefreshTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// RefreshTokenResponse is empty.
type RefreshTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenResponse) Reset()         { *m = RefreshTokenResponse{} }
func (m *RefreshTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenResponse) ProtoMessage()    {}
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_157936222af7c30e, []int{3}
}
func (m *RefreshTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenResponse.Unmarshal(m, b)
}
func (m *RefreshTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenResponse.Marshal(b, m, deterministic)
}
func (dst *RefreshTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenResponse.Merge(dst, src)
}
func (m *RefreshTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenResponse.Size(m)
}
func (m *RefreshTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenResponse proto.InternalMessageInfo

func (m *RefreshTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTokenRequest)(nil), "proto.GetTokenRequest")
	proto.RegisterType((*GetTokenResponse)(nil), "proto.GetTokenResponse")
	proto.RegisterType((*RefreshTokenRequest)(nil), "proto.RefreshTokenRequest")
	proto.RegisterType((*RefreshTokenResponse)(nil), "proto.RefreshTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.Authentication/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.Authentication/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Authentication/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Authentication/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _Authentication_GetToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Authentication_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

func init() {
	proto.RegisterFile("authentication.proto", fileDescriptor_authentication_157936222af7c30e)
}

var fileDescriptor_authentication_157936222af7c30e = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2c, 0x2d, 0xc9,
	0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x53, 0x4a, 0xea, 0x5c, 0xfc, 0xee, 0xa9, 0x25, 0x21, 0xf9, 0xd9, 0xa9, 0x79,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x92, 0x06, 0x97, 0x00, 0x42, 0x61, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x65, 0x09, 0x48, 0x00, 0xa6, 0x12, 0xcc, 0x51, 0xd2, 0xe6,
	0x12, 0x0e, 0x4a, 0x4d, 0x2b, 0x4a, 0x2d, 0xce, 0x40, 0x37, 0x16, 0x8b, 0x62, 0x1d, 0x2e, 0x11,
	0x54, 0xc5, 0xf8, 0x8c, 0x36, 0x9a, 0xc6, 0xc8, 0xc5, 0xe7, 0x88, 0xe2, 0x1b, 0x21, 0x6b, 0x2e,
	0x0e, 0x98, 0xbb, 0x84, 0xc4, 0x20, 0x7e, 0xd3, 0x43, 0xf3, 0x91, 0x94, 0x38, 0x86, 0x38, 0xd4,
	0x16, 0x77, 0x2e, 0x1e, 0x64, 0xdb, 0x85, 0xa4, 0xa0, 0x0a, 0xb1, 0xb8, 0x5f, 0x4a, 0x1a, 0xab,
	0x1c, 0xc4, 0x20, 0x27, 0x19, 0x2e, 0xde, 0xcc, 0x7c, 0xbd, 0xf4, 0xa2, 0x82, 0x64, 0x88, 0x2a,
	0x27, 0xce, 0xd0, 0xe2, 0xd4, 0xa2, 0x00, 0x10, 0x33, 0x80, 0x31, 0x89, 0x0d, 0x2c, 0x66, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x30, 0xb5, 0xd3, 0x8a, 0x01, 0x00, 0x00,
}
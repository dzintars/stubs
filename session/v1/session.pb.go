// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/v1/session.proto

package v1 // import "github.com/oswee/stubs/session/v1"

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

type CreateSessionRequest struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionRequest) Reset()         { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()    {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_c923fe806570fec5, []int{0}
}
func (m *CreateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionRequest.Unmarshal(m, b)
}
func (m *CreateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionRequest.Merge(dst, src)
}
func (m *CreateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionRequest.Size(m)
}
func (m *CreateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionRequest proto.InternalMessageInfo

func (m *CreateSessionRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type GetSessionRequest struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionRequest) Reset()         { *m = GetSessionRequest{} }
func (m *GetSessionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionRequest) ProtoMessage()    {}
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_c923fe806570fec5, []int{1}
}
func (m *GetSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionRequest.Unmarshal(m, b)
}
func (m *GetSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionRequest.Marshal(b, m, deterministic)
}
func (dst *GetSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionRequest.Merge(dst, src)
}
func (m *GetSessionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionRequest.Size(m)
}
func (m *GetSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionRequest proto.InternalMessageInfo

func (m *GetSessionRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type Session struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserSessionId        string   `protobuf:"bytes,2,opt,name=user_session_id,json=userSessionId,proto3" json:"user_session_id,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PermissionId         int64    `protobuf:"varint,4,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	CreateTime           string   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_c923fe806570fec5, []int{2}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Session) GetUserSessionId() string {
	if m != nil {
		return m.UserSessionId
	}
	return ""
}

func (m *Session) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Session) GetPermissionId() int64 {
	if m != nil {
		return m.PermissionId
	}
	return 0
}

func (m *Session) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSessionRequest)(nil), "oswee.session.v1.CreateSessionRequest")
	proto.RegisterType((*GetSessionRequest)(nil), "oswee.session.v1.GetSessionRequest")
	proto.RegisterType((*Session)(nil), "oswee.session.v1.Session")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error)
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/oswee.session.v1.SessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/oswee.session.v1.SessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*Session, error)
	GetSession(context.Context, *GetSessionRequest) (*Session, error)
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.session.v1.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.session.v1.SessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oswee.session.v1.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _SessionService_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session/v1/session.proto",
}

func init() { proto.RegisterFile("session/v1/session.proto", fileDescriptor_session_c923fe806570fec5) }

var fileDescriptor_session_c923fe806570fec5 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x03, 0xd5, 0x12, 0xa7, 0x52, 0x75, 0x62, 0x22, 0x7a, 0xb1, 0x96, 0xa4, 0xe9, 0x09,
	0xd2, 0xf6, 0x0d, 0xf4, 0xa0, 0x44, 0x4f, 0xd4, 0x93, 0x17, 0x22, 0xec, 0x44, 0xe7, 0x80, 0x54,
	0x76, 0xc1, 0xd7, 0xf1, 0x19, 0x7c, 0x42, 0xd3, 0xed, 0x22, 0x51, 0xd4, 0x53, 0x6f, 0xf0, 0xf3,
	0x7f, 0x33, 0xe1, 0xcb, 0x80, 0x27, 0x49, 0x4a, 0x2e, 0x5e, 0xc2, 0x7a, 0x16, 0x9a, 0xc7, 0x60,
	0x55, 0x16, 0xaa, 0xc0, 0xc3, 0x42, 0xbe, 0x11, 0x05, 0x4d, 0x58, 0xcf, 0xc6, 0xb7, 0x70, 0x7c,
	0x55, 0xd2, 0xa3, 0xa2, 0xe5, 0x26, 0x8b, 0xe9, 0xb5, 0x22, 0xa9, 0x70, 0x01, 0x8e, 0x69, 0x79,
	0xd6, 0xc8, 0x9a, 0x0e, 0xe6, 0xa7, 0xc1, 0x4f, 0x36, 0x68, 0x90, 0xa6, 0x39, 0xbe, 0x81, 0xa3,
	0x6b, 0x52, 0xdb, 0x98, 0xf4, 0x6e, 0x81, 0x63, 0x42, 0x1c, 0x82, 0xcd, 0x42, 0xb3, 0xbd, 0xd8,
	0x66, 0x81, 0x13, 0x38, 0xa8, 0x24, 0x95, 0x89, 0xe9, 0x26, 0x2c, 0x3c, 0x7b, 0x64, 0x4d, 0xf7,
	0x62, 0x77, 0x1d, 0x1b, 0x2a, 0x12, 0x78, 0x02, 0x8e, 0xee, 0xb1, 0xf0, 0x7a, 0x1a, 0xee, 0xaf,
	0x5f, 0x23, 0x81, 0x3e, 0xb8, 0x2b, 0x2a, 0x73, 0xfe, 0xc2, 0x77, 0xf4, 0xe7, 0xfd, 0x36, 0x8c,
	0x04, 0x9e, 0xc3, 0x20, 0xd3, 0x62, 0x12, 0xc5, 0x39, 0x79, 0xbb, 0x7a, 0x03, 0x6c, 0xa2, 0x7b,
	0xce, 0x69, 0xfe, 0x61, 0xc1, 0xd0, 0x2c, 0x5b, 0x52, 0x59, 0x73, 0x46, 0x18, 0x83, 0xfb, 0x4d,
	0x26, 0x4e, 0xba, 0xbf, 0xfa, 0x9b, 0xed, 0xb3, 0xbf, 0x95, 0xe0, 0x1d, 0x40, 0xeb, 0x14, 0xfd,
	0x6e, 0xb1, 0x63, 0xfc, 0x9f, 0x69, 0x97, 0xfe, 0xc3, 0xc5, 0x13, 0xab, 0xe7, 0x2a, 0x0d, 0xb2,
	0x22, 0x0f, 0x75, 0x2d, 0x94, 0xaa, 0x4a, 0x65, 0xd8, 0xde, 0x4c, 0xda, 0xd7, 0xc7, 0xb2, 0xf8,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x59, 0x6c, 0x99, 0x30, 0x48, 0x02, 0x00, 0x00,
}

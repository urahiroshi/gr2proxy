// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/helloworld/helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6850e8e9b6bacb47, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6850e8e9b6bacb47, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GoodbyeRequest struct {
	Hoge                 string   `protobuf:"bytes,1,opt,name=hoge,proto3" json:"hoge,omitempty"`
	Fuga                 string   `protobuf:"bytes,2,opt,name=fuga,proto3" json:"fuga,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodbyeRequest) Reset()         { *m = GoodbyeRequest{} }
func (m *GoodbyeRequest) String() string { return proto.CompactTextString(m) }
func (*GoodbyeRequest) ProtoMessage()    {}
func (*GoodbyeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6850e8e9b6bacb47, []int{2}
}

func (m *GoodbyeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodbyeRequest.Unmarshal(m, b)
}
func (m *GoodbyeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodbyeRequest.Marshal(b, m, deterministic)
}
func (m *GoodbyeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodbyeRequest.Merge(m, src)
}
func (m *GoodbyeRequest) XXX_Size() int {
	return xxx_messageInfo_GoodbyeRequest.Size(m)
}
func (m *GoodbyeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodbyeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GoodbyeRequest proto.InternalMessageInfo

func (m *GoodbyeRequest) GetHoge() string {
	if m != nil {
		return m.Hoge
	}
	return ""
}

func (m *GoodbyeRequest) GetFuga() string {
	if m != nil {
		return m.Fuga
	}
	return ""
}

type GoodbyeReply struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodbyeReply) Reset()         { *m = GoodbyeReply{} }
func (m *GoodbyeReply) String() string { return proto.CompactTextString(m) }
func (*GoodbyeReply) ProtoMessage()    {}
func (*GoodbyeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6850e8e9b6bacb47, []int{3}
}

func (m *GoodbyeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodbyeReply.Unmarshal(m, b)
}
func (m *GoodbyeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodbyeReply.Marshal(b, m, deterministic)
}
func (m *GoodbyeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodbyeReply.Merge(m, src)
}
func (m *GoodbyeReply) XXX_Size() int {
	return xxx_messageInfo_GoodbyeReply.Size(m)
}
func (m *GoodbyeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodbyeReply.DiscardUnknown(m)
}

var xxx_messageInfo_GoodbyeReply proto.InternalMessageInfo

func (m *GoodbyeReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*GoodbyeRequest)(nil), "helloworld.GoodbyeRequest")
	proto.RegisterType((*GoodbyeReply)(nil), "helloworld.GoodbyeReply")
}

func init() { proto.RegisterFile("test/helloworld/helloworld.proto", fileDescriptor_6850e8e9b6bacb47) }

var fileDescriptor_6850e8e9b6bacb47 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x75, 0x45, 0xad, 0x0e, 0x45, 0x61, 0x10, 0x09, 0xeb, 0xa5, 0x04, 0x11, 0x4f, 0xab, 0xe8,
	0xc5, 0x93, 0x87, 0x22, 0xd4, 0x63, 0x69, 0x0f, 0x9e, 0x53, 0x3b, 0x6e, 0x85, 0xac, 0x13, 0x93,
	0xac, 0x9a, 0xaf, 0xf0, 0x97, 0x25, 0x69, 0xb7, 0x46, 0xd8, 0xdb, 0x7b, 0x6f, 0xde, 0xcc, 0xdb,
	0xb7, 0x81, 0x91, 0x27, 0xe7, 0xaf, 0x57, 0xa4, 0x35, 0x7f, 0xb1, 0xd5, 0xcb, 0x0c, 0x56, 0xc6,
	0xb2, 0x67, 0x84, 0x3f, 0x45, 0x4a, 0x18, 0x3e, 0x45, 0x36, 0xa3, 0x8f, 0x96, 0x9c, 0x47, 0x84,
	0xbd, 0x77, 0xd5, 0x90, 0x28, 0x46, 0xc5, 0xd5, 0xd1, 0x2c, 0x61, 0x79, 0x09, 0xb0, 0xf1, 0x18,
	0x1d, 0x50, 0xc0, 0xa0, 0x21, 0xe7, 0x54, 0xdd, 0x99, 0x3a, 0x2a, 0xef, 0xe1, 0x78, 0xc2, 0xbc,
	0x5c, 0x04, 0xca, 0xae, 0xad, 0x78, 0x6b, 0x4c, 0x38, 0x6a, 0xaf, 0x6d, 0xad, 0xc4, 0xee, 0x5a,
	0x8b, 0x58, 0x5e, 0xc0, 0x70, 0xbb, 0x19, 0x33, 0x4e, 0x61, 0xff, 0x53, 0xe9, 0xb6, 0x5b, 0x5c,
	0x93, 0xdb, 0x9f, 0x02, 0x06, 0x13, 0x4b, 0xe4, 0xc9, 0xe2, 0x03, 0x1c, 0xce, 0x55, 0x48, 0x9f,
	0x85, 0xa2, 0xca, 0x2a, 0xe6, 0x6d, 0xca, 0xb3, 0x9e, 0x89, 0xd1, 0x41, 0xee, 0xe0, 0x23, 0xc0,
	0x5c, 0x85, 0x4d, 0x28, 0x96, 0xb9, 0xef, 0x7f, 0x87, 0x52, 0xf4, 0xce, 0xd2, 0x95, 0xf1, 0x0d,
	0x9c, 0xbf, 0x71, 0x55, 0x5b, 0xf3, 0x52, 0xd1, 0xb7, 0x6a, 0x8c, 0x26, 0x97, 0xb9, 0xc7, 0x27,
	0x29, 0xf2, 0x39, 0xe2, 0x69, 0xfc, 0xf3, 0xd3, 0x62, 0x71, 0x90, 0x9e, 0xe0, 0xee, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x65, 0x7b, 0x31, 0x65, 0xa6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayGoodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayGoodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeReply, error) {
	out := new(GoodbyeReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayGoodbye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayGoodbye(context.Context, *GoodbyeRequest) (*GoodbyeReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) SayGoodbye(ctx context.Context, req *GoodbyeRequest) (*GoodbyeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayGoodbye not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayGoodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodbyeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayGoodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayGoodbye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayGoodbye(ctx, req.(*GoodbyeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayGoodbye",
			Handler:    _Greeter_SayGoodbye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/helloworld/helloworld.proto",
}
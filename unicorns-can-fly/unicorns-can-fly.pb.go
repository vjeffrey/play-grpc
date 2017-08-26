// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unicorns-can-fly.proto

/*
Package unicorns is a generated protocol buffer package.

It is generated from these files:
	unicorns-can-fly.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package unicorns

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "unicorns.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "unicorns.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Friend service

type FriendClient interface {
	// says hi to his friend
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// Sends another greeting
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type friendClient struct {
	cc *grpc.ClientConn
}

func NewFriendClient(cc *grpc.ClientConn) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/unicorns.Friend/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/unicorns.Friend/SayHelloAgain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Friend service

type FriendServer interface {
	// says hi to his friend
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// Sends another greeting
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterFriendServer(s *grpc.Server, srv FriendServer) {
	s.RegisterService(&_Friend_serviceDesc, srv)
}

func _Friend_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unicorns.Friend/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unicorns.Friend/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Friend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "unicorns.Friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Friend_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Friend_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unicorns-can-fly.proto",
}

func init() { proto.RegisterFile("unicorns-can-fly.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xcd, 0xcb, 0x4c,
	0xce, 0x2f, 0xca, 0x2b, 0xd6, 0x4d, 0x4e, 0xcc, 0xd3, 0x4d, 0xcb, 0xa9, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x2b, 0x29, 0x71, 0xf1, 0x78, 0xa4, 0xe6, 0xe4, 0xe4, 0x07,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x6a, 0x5c, 0x5c, 0x50, 0x35, 0x05, 0x39, 0x95,
	0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x45, 0x30, 0xae, 0x51, 0x2b,
	0x23, 0x17, 0x9b, 0x5b, 0x51, 0x66, 0x6a, 0x5e, 0x8a, 0x90, 0x15, 0x17, 0x47, 0x70, 0x62, 0x25,
	0x58, 0x97, 0x90, 0x98, 0x1e, 0xcc, 0x36, 0x3d, 0x64, 0xab, 0xa4, 0x44, 0x30, 0xc4, 0x0b, 0x72,
	0x2a, 0x95, 0x18, 0x84, 0xec, 0xb9, 0x78, 0x61, 0x7a, 0x1d, 0xd3, 0x13, 0x33, 0xf3, 0x48, 0x35,
	0xc0, 0x89, 0x29, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x51, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x69, 0xe4, 0xb0, 0xd0, 0x02, 0x01, 0x00, 0x00,
}

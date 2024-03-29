// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-grpc-example/9-grpc_proto_validators/proto/simple.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type InnerMessage struct {
	// some_integer can only be in range (1, 100).
	SomeInteger int32 `protobuf:"varint,1,opt,name=some_integer,json=someInteger,proto3" json:"some_integer,omitempty"`
	// some_float can only be in range (0;1).
	SomeFloat            float64  `protobuf:"fixed64,2,opt,name=some_float,json=someFloat,proto3" json:"some_float,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerMessage) Reset()         { *m = InnerMessage{} }
func (m *InnerMessage) String() string { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()    {}
func (*InnerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e73575a25b94d6c6, []int{0}
}

func (m *InnerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerMessage.Unmarshal(m, b)
}
func (m *InnerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerMessage.Marshal(b, m, deterministic)
}
func (m *InnerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerMessage.Merge(m, src)
}
func (m *InnerMessage) XXX_Size() int {
	return xxx_messageInfo_InnerMessage.Size(m)
}
func (m *InnerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InnerMessage proto.InternalMessageInfo

func (m *InnerMessage) GetSomeInteger() int32 {
	if m != nil {
		return m.SomeInteger
	}
	return 0
}

func (m *InnerMessage) GetSomeFloat() float64 {
	if m != nil {
		return m.SomeFloat
	}
	return 0
}

type OuterMessage struct {
	// important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
	ImportantString string `protobuf:"bytes,1,opt,name=important_string,json=importantString,proto3" json:"important_string,omitempty"`
	// proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
	Inner                *InnerMessage `protobuf:"bytes,2,opt,name=inner,proto3" json:"inner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OuterMessage) Reset()         { *m = OuterMessage{} }
func (m *OuterMessage) String() string { return proto.CompactTextString(m) }
func (*OuterMessage) ProtoMessage()    {}
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e73575a25b94d6c6, []int{1}
}

func (m *OuterMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OuterMessage.Unmarshal(m, b)
}
func (m *OuterMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OuterMessage.Marshal(b, m, deterministic)
}
func (m *OuterMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OuterMessage.Merge(m, src)
}
func (m *OuterMessage) XXX_Size() int {
	return xxx_messageInfo_OuterMessage.Size(m)
}
func (m *OuterMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OuterMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OuterMessage proto.InternalMessageInfo

func (m *OuterMessage) GetImportantString() string {
	if m != nil {
		return m.ImportantString
	}
	return ""
}

func (m *OuterMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

func init() {
	proto.RegisterType((*InnerMessage)(nil), "proto.InnerMessage")
	proto.RegisterType((*OuterMessage)(nil), "proto.OuterMessage")
}

func init() {
	proto.RegisterFile("go-grpc-example/9-grpc_proto_validators/proto/simple.proto", fileDescriptor_e73575a25b94d6c6)
}

var fileDescriptor_e73575a25b94d6c6 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x97, 0x61, 0x87, 0xcb, 0x8a, 0x8e, 0x08, 0x32, 0x76, 0xe9, 0x18, 0x1e, 0x06, 0xda,
	0x16, 0x3b, 0x26, 0x28, 0x0c, 0xa1, 0x07, 0xa1, 0x07, 0x11, 0xbb, 0xa3, 0x68, 0xc9, 0xb6, 0x18,
	0x83, 0x6d, 0x53, 0x92, 0xd4, 0x89, 0x22, 0xf8, 0x4d, 0x0b, 0xfd, 0x04, 0x7e, 0x04, 0x69, 0xaa,
	0xb3, 0x07, 0x73, 0x79, 0x79, 0xef, 0xff, 0x7b, 0x79, 0xef, 0x1f, 0x78, 0x41, 0xb9, 0x4d, 0x45,
	0xb6, 0xb2, 0xc9, 0x2b, 0x4e, 0xb2, 0x98, 0xb8, 0xe7, 0x3a, 0x8d, 0x32, 0xc1, 0x15, 0x8f, 0x5e,
	0x70, 0xcc, 0xd6, 0x58, 0x71, 0x21, 0x5d, 0x5d, 0x70, 0x25, 0xab, 0x20, 0x47, 0x27, 0xc8, 0xd0,
	0x61, 0x78, 0x46, 0x99, 0x7a, 0xca, 0x97, 0xce, 0x8a, 0x27, 0x6e, 0xb2, 0x61, 0xea, 0x99, 0x6f,
	0x5c, 0xca, 0x6d, 0x2d, 0xda, 0x8d, 0x17, 0xb6, 0xd7, 0xba, 0x7d, 0x2c, 0xa0, 0x19, 0xa4, 0x29,
	0x11, 0xd7, 0x44, 0x4a, 0x4c, 0x09, 0x3a, 0x86, 0xa6, 0xe4, 0x09, 0x89, 0x58, 0xaa, 0x08, 0x25,
	0x62, 0x00, 0x46, 0x60, 0x62, 0xf8, 0xbb, 0x65, 0x61, 0xed, 0xf4, 0x5b, 0x83, 0x75, 0xd8, 0xab,
	0xd4, 0xa0, 0x16, 0xd1, 0x0c, 0x42, 0x0d, 0x3f, 0xc6, 0x1c, 0xab, 0x41, 0x7b, 0x04, 0x26, 0xc0,
	0x3f, 0x2c, 0x0b, 0x0b, 0x05, 0xad, 0x9f, 0x73, 0x5b, 0x87, 0xaf, 0xcb, 0xb0, 0x5b, 0x91, 0x57,
	0x15, 0x38, 0xfe, 0x04, 0xd0, 0xbc, 0xc9, 0xd5, 0xdf, 0xd0, 0x39, 0xec, 0xb3, 0x24, 0xe3, 0x42,
	0xe1, 0x54, 0x45, 0x52, 0x09, 0x96, 0x52, 0x3d, 0xb8, 0xeb, 0xa3, 0xb2, 0xb0, 0xf6, 0xa0, 0xf9,
	0x70, 0x87, 0xed, 0xb7, 0xfb, 0x77, 0xef, 0x64, 0xf6, 0x71, 0x14, 0xee, 0x6f, 0xd9, 0x85, 0x46,
	0xd1, 0x29, 0x34, 0x58, 0xe5, 0x41, 0x6f, 0xd0, 0xf3, 0x0e, 0x6a, 0x6b, 0x4e, 0xd3, 0x97, 0xdf,
	0x29, 0x0b, 0xab, 0x3d, 0x02, 0x61, 0x4d, 0x7a, 0x73, 0xd8, 0x59, 0xe8, 0x5f, 0x44, 0x53, 0x68,
	0x84, 0x3c, 0x57, 0x04, 0xfd, 0xd7, 0x36, 0xfc, 0x2d, 0x36, 0xd7, 0x1d, 0xb7, 0x96, 0x1d, 0x5d,
	0x9d, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x39, 0x9e, 0x62, 0xb9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleClient interface {
	Route(ctx context.Context, in *InnerMessage, opts ...grpc.CallOption) (*OuterMessage, error)
}

type simpleClient struct {
	cc *grpc.ClientConn
}

func NewSimpleClient(cc *grpc.ClientConn) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) Route(ctx context.Context, in *InnerMessage, opts ...grpc.CallOption) (*OuterMessage, error) {
	out := new(OuterMessage)
	err := c.cc.Invoke(ctx, "/proto.Simple/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
type SimpleServer interface {
	Route(context.Context, *InnerMessage) (*OuterMessage, error)
}

// UnimplementedSimpleServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleServer struct {
}

func (*UnimplementedSimpleServer) Route(ctx context.Context, req *InnerMessage) (*OuterMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}

func RegisterSimpleServer(s *grpc.Server, srv SimpleServer) {
	s.RegisterService(&_Simple_serviceDesc, srv)
}

func _Simple_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InnerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Simple/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).Route(ctx, req.(*InnerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Simple_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _Simple_Route_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-grpc-example/9-grpc_proto_validators/proto/simple.proto",
}

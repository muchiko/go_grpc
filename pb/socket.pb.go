// Code generated by protoc-gen-go. DO NOT EDIT.
// source: socket.proto

package socket

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

type Request struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b39cc5e3943e1cc, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Payload struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b39cc5e3943e1cc, []int{1}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "socket.Request")
	proto.RegisterType((*Payload)(nil), "socket.Payload")
}

func init() { proto.RegisterFile("socket.proto", fileDescriptor_6b39cc5e3943e1cc) }

var fileDescriptor_6b39cc5e3943e1cc = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0x4f, 0xce,
	0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x94, 0xb9, 0xd8,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13,
	0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x90, 0xa2, 0x80, 0xc4, 0xca,
	0x9c, 0xfc, 0xc4, 0x14, 0xdc, 0x8a, 0x8c, 0x5c, 0xb8, 0x78, 0x83, 0xc1, 0x66, 0x06, 0xa7, 0x16,
	0x95, 0x65, 0x26, 0xa7, 0x0a, 0x19, 0x73, 0x71, 0x86, 0x14, 0x25, 0xe6, 0x15, 0x17, 0xe4, 0x17,
	0x95, 0x08, 0xf1, 0xeb, 0x41, 0xad, 0x87, 0xda, 0x26, 0x05, 0x17, 0x80, 0x9a, 0xac, 0xc4, 0xa0,
	0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x05,
	0xec, 0xd0, 0xae, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SocketServiceClient is the client API for SocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SocketServiceClient interface {
	Transport(ctx context.Context, opts ...grpc.CallOption) (SocketService_TransportClient, error)
}

type socketServiceClient struct {
	cc *grpc.ClientConn
}

func NewSocketServiceClient(cc *grpc.ClientConn) SocketServiceClient {
	return &socketServiceClient{cc}
}

func (c *socketServiceClient) Transport(ctx context.Context, opts ...grpc.CallOption) (SocketService_TransportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SocketService_serviceDesc.Streams[0], "/socket.SocketService/Transport", opts...)
	if err != nil {
		return nil, err
	}
	x := &socketServiceTransportClient{stream}
	return x, nil
}

type SocketService_TransportClient interface {
	Send(*Request) error
	Recv() (*Payload, error)
	grpc.ClientStream
}

type socketServiceTransportClient struct {
	grpc.ClientStream
}

func (x *socketServiceTransportClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *socketServiceTransportClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SocketServiceServer is the server API for SocketService service.
type SocketServiceServer interface {
	Transport(SocketService_TransportServer) error
}

// UnimplementedSocketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSocketServiceServer struct {
}

func (*UnimplementedSocketServiceServer) Transport(srv SocketService_TransportServer) error {
	return status.Errorf(codes.Unimplemented, "method Transport not implemented")
}

func RegisterSocketServiceServer(s *grpc.Server, srv SocketServiceServer) {
	s.RegisterService(&_SocketService_serviceDesc, srv)
}

func _SocketService_Transport_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SocketServiceServer).Transport(&socketServiceTransportServer{stream})
}

type SocketService_TransportServer interface {
	Send(*Payload) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type socketServiceTransportServer struct {
	grpc.ServerStream
}

func (x *socketServiceTransportServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *socketServiceTransportServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SocketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "socket.SocketService",
	HandlerType: (*SocketServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transport",
			Handler:       _SocketService_Transport_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "socket.proto",
}

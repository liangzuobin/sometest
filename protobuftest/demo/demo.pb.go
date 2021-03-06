// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo/demo.proto

package demo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type JustKey struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JustKey) Reset()         { *m = JustKey{} }
func (m *JustKey) String() string { return proto.CompactTextString(m) }
func (*JustKey) ProtoMessage()    {}
func (*JustKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ceb39d273bd4a9, []int{0}
}

func (m *JustKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JustKey.Unmarshal(m, b)
}
func (m *JustKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JustKey.Marshal(b, m, deterministic)
}
func (m *JustKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JustKey.Merge(m, src)
}
func (m *JustKey) XXX_Size() int {
	return xxx_messageInfo_JustKey.Size(m)
}
func (m *JustKey) XXX_DiscardUnknown() {
	xxx_messageInfo_JustKey.DiscardUnknown(m)
}

var xxx_messageInfo_JustKey proto.InternalMessageInfo

func (m *JustKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type JustBytes struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JustBytes) Reset()         { *m = JustBytes{} }
func (m *JustBytes) String() string { return proto.CompactTextString(m) }
func (*JustBytes) ProtoMessage()    {}
func (*JustBytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ceb39d273bd4a9, []int{1}
}

func (m *JustBytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JustBytes.Unmarshal(m, b)
}
func (m *JustBytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JustBytes.Marshal(b, m, deterministic)
}
func (m *JustBytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JustBytes.Merge(m, src)
}
func (m *JustBytes) XXX_Size() int {
	return xxx_messageInfo_JustBytes.Size(m)
}
func (m *JustBytes) XXX_DiscardUnknown() {
	xxx_messageInfo_JustBytes.DiscardUnknown(m)
}

var xxx_messageInfo_JustBytes proto.InternalMessageInfo

func (m *JustBytes) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type StreamReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamReq) Reset()         { *m = StreamReq{} }
func (m *StreamReq) String() string { return proto.CompactTextString(m) }
func (*StreamReq) ProtoMessage()    {}
func (*StreamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ceb39d273bd4a9, []int{2}
}

func (m *StreamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamReq.Unmarshal(m, b)
}
func (m *StreamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamReq.Marshal(b, m, deterministic)
}
func (m *StreamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamReq.Merge(m, src)
}
func (m *StreamReq) XXX_Size() int {
	return xxx_messageInfo_StreamReq.Size(m)
}
func (m *StreamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamReq.DiscardUnknown(m)
}

var xxx_messageInfo_StreamReq proto.InternalMessageInfo

func (m *StreamReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StreamReq) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*JustKey)(nil), "demo.JustKey")
	proto.RegisterType((*JustBytes)(nil), "demo.JustBytes")
	proto.RegisterType((*StreamReq)(nil), "demo.StreamReq")
}

func init() { proto.RegisterFile("demo/demo.proto", fileDescriptor_18ceb39d273bd4a9) }

var fileDescriptor_18ceb39d273bd4a9 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0xcd, 0xcd,
	0xd7, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x34, 0x17,
	0xbb, 0x57, 0x69, 0x71, 0x89, 0x77, 0x6a, 0xa5, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0xa9, 0xa4, 0xc8, 0xc5, 0x09, 0x92, 0x74, 0xaa, 0x2c,
	0x49, 0x2d, 0x16, 0x12, 0xe1, 0x62, 0x4d, 0x02, 0x31, 0xc0, 0x0a, 0x78, 0x82, 0x20, 0x1c, 0x25,
	0x63, 0x2e, 0xce, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0xdc, 0xa0, 0xd4, 0x42, 0x4c, 0x13, 0x40, 0x9a,
	0xca, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0x20, 0x9a, 0xc0, 0x1c, 0xa3, 0x76, 0x46, 0x2e, 0xf6,
	0xe0, 0x92, 0xfc, 0xa2, 0xc4, 0xf4, 0x54, 0x21, 0x55, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x5e,
	0x3d, 0xb0, 0xd3, 0xa0, 0x6e, 0x91, 0xe2, 0x47, 0x70, 0x21, 0xb6, 0x6b, 0x72, 0x31, 0x07, 0x94,
	0x96, 0x08, 0x41, 0xc5, 0xe1, 0x56, 0x62, 0x28, 0xd4, 0x60, 0x14, 0xd2, 0xe2, 0x62, 0x09, 0x28,
	0x2d, 0x31, 0x22, 0x46, 0x6d, 0x12, 0x1b, 0x38, 0x2c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x2a, 0x01, 0x5a, 0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	Get(ctx context.Context, in *JustKey, opts ...grpc.CallOption) (*JustBytes, error)
	Put(ctx context.Context, opts ...grpc.CallOption) (Storage_PutClient, error)
	Put2(ctx context.Context, opts ...grpc.CallOption) (Storage_Put2Client, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Get(ctx context.Context, in *JustKey, opts ...grpc.CallOption) (*JustBytes, error) {
	out := new(JustBytes)
	err := c.cc.Invoke(ctx, "/demo.Storage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Put(ctx context.Context, opts ...grpc.CallOption) (Storage_PutClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Storage_serviceDesc.Streams[0], "/demo.Storage/Put", opts...)
	if err != nil {
		return nil, err
	}
	x := &storagePutClient{stream}
	return x, nil
}

type Storage_PutClient interface {
	Send(*StreamReq) error
	CloseAndRecv() (*JustBytes, error)
	grpc.ClientStream
}

type storagePutClient struct {
	grpc.ClientStream
}

func (x *storagePutClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storagePutClient) CloseAndRecv() (*JustBytes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(JustBytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Put2(ctx context.Context, opts ...grpc.CallOption) (Storage_Put2Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Storage_serviceDesc.Streams[1], "/demo.Storage/Put2", opts...)
	if err != nil {
		return nil, err
	}
	x := &storagePut2Client{stream}
	return x, nil
}

type Storage_Put2Client interface {
	Send(*StreamReq) error
	CloseAndRecv() (*JustBytes, error)
	grpc.ClientStream
}

type storagePut2Client struct {
	grpc.ClientStream
}

func (x *storagePut2Client) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storagePut2Client) CloseAndRecv() (*JustBytes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(JustBytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	Get(context.Context, *JustKey) (*JustBytes, error)
	Put(Storage_PutServer) error
	Put2(Storage_Put2Server) error
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JustKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Storage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*JustKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Put_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Put(&storagePutServer{stream})
}

type Storage_PutServer interface {
	SendAndClose(*JustBytes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type storagePutServer struct {
	grpc.ServerStream
}

func (x *storagePutServer) SendAndClose(m *JustBytes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storagePutServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_Put2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Put2(&storagePut2Server{stream})
}

type Storage_Put2Server interface {
	SendAndClose(*JustBytes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type storagePut2Server struct {
	grpc.ServerStream
}

func (x *storagePut2Server) SendAndClose(m *JustBytes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storagePut2Server) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Put",
			Handler:       _Storage_Put_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Put2",
			Handler:       _Storage_Put2_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "demo/demo.proto",
}

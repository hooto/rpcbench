// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package rpcbench

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

type BenchCompressType int32

const (
	BenchCompressType_None BenchCompressType = 0
)

var BenchCompressType_name = map[int32]string{
	0: "None",
}

var BenchCompressType_value = map[string]int32{
	"None": 0,
}

func (x BenchCompressType) String() string {
	return proto.EnumName(BenchCompressType_name, int32(x))
}

func (BenchCompressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

type BenchPayloadType int32

const (
	BenchPayloadType_Text BenchPayloadType = 0
)

var BenchPayloadType_name = map[int32]string{
	0: "Text",
}

var BenchPayloadType_value = map[string]int32{
	"Text": 0,
}

func (x BenchPayloadType) String() string {
	return proto.EnumName(BenchPayloadType_name, int32(x))
}

func (BenchPayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

type BenchPayloadSpec struct {
	Type                 BenchPayloadType  `protobuf:"varint,1,opt,name=type,proto3,enum=rpcbench.BenchPayloadType" json:"type,omitempty"`
	Size                 int32             `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	CompressType         BenchCompressType `protobuf:"varint,3,opt,name=compress_type,json=compressType,proto3,enum=rpcbench.BenchCompressType" json:"compress_type,omitempty"`
	WorkTime             int32             `protobuf:"varint,6,opt,name=work_time,json=workTime,proto3" json:"work_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BenchPayloadSpec) Reset()         { *m = BenchPayloadSpec{} }
func (m *BenchPayloadSpec) String() string { return proto.CompactTextString(m) }
func (*BenchPayloadSpec) ProtoMessage()    {}
func (*BenchPayloadSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *BenchPayloadSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BenchPayloadSpec.Unmarshal(m, b)
}
func (m *BenchPayloadSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BenchPayloadSpec.Marshal(b, m, deterministic)
}
func (m *BenchPayloadSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BenchPayloadSpec.Merge(m, src)
}
func (m *BenchPayloadSpec) XXX_Size() int {
	return xxx_messageInfo_BenchPayloadSpec.Size(m)
}
func (m *BenchPayloadSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BenchPayloadSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BenchPayloadSpec proto.InternalMessageInfo

func (m *BenchPayloadSpec) GetType() BenchPayloadType {
	if m != nil {
		return m.Type
	}
	return BenchPayloadType_Text
}

func (m *BenchPayloadSpec) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BenchPayloadSpec) GetCompressType() BenchCompressType {
	if m != nil {
		return m.CompressType
	}
	return BenchCompressType_None
}

func (m *BenchPayloadSpec) GetWorkTime() int32 {
	if m != nil {
		return m.WorkTime
	}
	return 0
}

type BenchPayload struct {
	Spec                 *BenchPayloadSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Body                 []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BenchPayload) Reset()         { *m = BenchPayload{} }
func (m *BenchPayload) String() string { return proto.CompactTextString(m) }
func (*BenchPayload) ProtoMessage()    {}
func (*BenchPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *BenchPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BenchPayload.Unmarshal(m, b)
}
func (m *BenchPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BenchPayload.Marshal(b, m, deterministic)
}
func (m *BenchPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BenchPayload.Merge(m, src)
}
func (m *BenchPayload) XXX_Size() int {
	return xxx_messageInfo_BenchPayload.Size(m)
}
func (m *BenchPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BenchPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BenchPayload proto.InternalMessageInfo

func (m *BenchPayload) GetSpec() *BenchPayloadSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *BenchPayload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type BenchRequest struct {
	Id                   uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload              *BenchPayload     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	ReplySpec            *BenchPayloadSpec `protobuf:"bytes,3,opt,name=reply_spec,json=replySpec,proto3" json:"reply_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BenchRequest) Reset()         { *m = BenchRequest{} }
func (m *BenchRequest) String() string { return proto.CompactTextString(m) }
func (*BenchRequest) ProtoMessage()    {}
func (*BenchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *BenchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BenchRequest.Unmarshal(m, b)
}
func (m *BenchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BenchRequest.Marshal(b, m, deterministic)
}
func (m *BenchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BenchRequest.Merge(m, src)
}
func (m *BenchRequest) XXX_Size() int {
	return xxx_messageInfo_BenchRequest.Size(m)
}
func (m *BenchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BenchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BenchRequest proto.InternalMessageInfo

func (m *BenchRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BenchRequest) GetPayload() *BenchPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *BenchRequest) GetReplySpec() *BenchPayloadSpec {
	if m != nil {
		return m.ReplySpec
	}
	return nil
}

type BenchReply struct {
	Id                   uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload              *BenchPayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BenchReply) Reset()         { *m = BenchReply{} }
func (m *BenchReply) String() string { return proto.CompactTextString(m) }
func (*BenchReply) ProtoMessage()    {}
func (*BenchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *BenchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BenchReply.Unmarshal(m, b)
}
func (m *BenchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BenchReply.Marshal(b, m, deterministic)
}
func (m *BenchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BenchReply.Merge(m, src)
}
func (m *BenchReply) XXX_Size() int {
	return xxx_messageInfo_BenchReply.Size(m)
}
func (m *BenchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BenchReply.DiscardUnknown(m)
}

var xxx_messageInfo_BenchReply proto.InternalMessageInfo

func (m *BenchReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BenchReply) GetPayload() *BenchPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("rpcbench.BenchCompressType", BenchCompressType_name, BenchCompressType_value)
	proto.RegisterEnum("rpcbench.BenchPayloadType", BenchPayloadType_name, BenchPayloadType_value)
	proto.RegisterType((*BenchPayloadSpec)(nil), "rpcbench.BenchPayloadSpec")
	proto.RegisterType((*BenchPayload)(nil), "rpcbench.BenchPayload")
	proto.RegisterType((*BenchRequest)(nil), "rpcbench.BenchRequest")
	proto.RegisterType((*BenchReply)(nil), "rpcbench.BenchReply")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4f, 0xfa, 0x30,
	0x14, 0xc7, 0x29, 0xec, 0xc7, 0x8f, 0x3d, 0x90, 0x60, 0x63, 0x0c, 0x01, 0x4d, 0xc8, 0x4e, 0x84,
	0xc3, 0x62, 0xf0, 0x64, 0xbc, 0x18, 0xb8, 0x78, 0x22, 0xa6, 0xe2, 0x99, 0x8c, 0xed, 0x25, 0x2e,
	0x6c, 0xb4, 0xb6, 0x33, 0x5a, 0x4f, 0xde, 0xfd, 0x7b, 0xfc, 0xff, 0x4c, 0x9f, 0x23, 0x41, 0xd0,
	0x78, 0xf0, 0xf6, 0x5d, 0xfb, 0xdd, 0xa7, 0x9f, 0x36, 0x0f, 0x9a, 0x85, 0x55, 0x68, 0x42, 0xa5,
	0x65, 0x21, 0x79, 0x43, 0xab, 0x78, 0x89, 0xeb, 0xf8, 0x3e, 0x78, 0x67, 0xd0, 0x99, 0xb8, 0x74,
	0x13, 0xd9, 0x4c, 0x46, 0xc9, 0xad, 0xc2, 0x98, 0x87, 0xe0, 0xb9, 0x76, 0x97, 0x0d, 0xd8, 0xb0,
	0x3d, 0xee, 0x85, 0x9b, 0x76, 0xb8, 0xdd, 0x9c, 0x5b, 0x85, 0x82, 0x7a, 0x9c, 0x83, 0x67, 0xd2,
	0x17, 0xec, 0x56, 0x07, 0x6c, 0xf8, 0x4f, 0x50, 0xe6, 0x57, 0x70, 0x10, 0xcb, 0x5c, 0x69, 0x34,
	0x66, 0x41, 0xb0, 0x1a, 0xc1, 0xfa, 0x3b, 0xb0, 0x69, 0xd9, 0x21, 0x5a, 0x2b, 0xde, 0xfa, 0xe2,
	0x7d, 0xf0, 0x9f, 0xa4, 0x5e, 0x2d, 0x8a, 0x34, 0xc7, 0x6e, 0x9d, 0xd0, 0x0d, 0xb7, 0x30, 0x4f,
	0x73, 0x0c, 0x04, 0xb4, 0xb6, 0x65, 0x9c, 0xb2, 0x51, 0x18, 0x93, 0x72, 0xf3, 0x27, 0x65, 0x77,
	0x39, 0x41, 0x3d, 0xa7, 0xbc, 0x94, 0x89, 0x25, 0xe5, 0x96, 0xa0, 0x1c, 0xbc, 0xb1, 0x12, 0x2a,
	0xf0, 0xe1, 0x11, 0x4d, 0xc1, 0xdb, 0x50, 0x4d, 0x13, 0x42, 0x7a, 0xa2, 0x9a, 0x26, 0xfc, 0x0c,
	0xfe, 0xab, 0x4f, 0x12, 0xfd, 0xd7, 0x1c, 0x1f, 0x7f, 0x7f, 0x8e, 0xd8, 0xd4, 0xf8, 0x05, 0x80,
	0x46, 0x95, 0xd9, 0x05, 0xc9, 0xd5, 0x7e, 0x95, 0xf3, 0xa9, 0xed, 0x62, 0x30, 0x03, 0x28, 0x65,
	0x54, 0x66, 0xff, 0xae, 0x32, 0x3a, 0x85, 0xc3, 0xbd, 0x17, 0xe7, 0x0d, 0xf0, 0x66, 0x72, 0x8d,
	0x9d, 0xca, 0xe8, 0xe4, 0xeb, 0x1c, 0x6c, 0x76, 0xe7, 0xf8, 0x5c, 0x74, 0x2a, 0xe3, 0x6b, 0xf0,
	0x69, 0x37, 0x8f, 0xf4, 0x8a, 0x5f, 0x82, 0x7f, 0xb7, 0x8e, 0xb4, 0x9d, 0x46, 0x59, 0xc6, 0x77,
	0xcf, 0x2d, 0xdf, 0xae, 0x77, 0xb4, 0xb7, 0xae, 0x32, 0x1b, 0x54, 0x26, 0xb5, 0x57, 0xc6, 0x96,
	0x75, 0x1a, 0xc3, 0xf3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x19, 0x65, 0xf8, 0x95, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BenchmarkClient is the client API for Benchmark service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BenchmarkClient interface {
	UnaryCall(ctx context.Context, in *BenchRequest, opts ...grpc.CallOption) (*BenchReply, error)
}

type benchmarkClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkClient(cc grpc.ClientConnInterface) BenchmarkClient {
	return &benchmarkClient{cc}
}

func (c *benchmarkClient) UnaryCall(ctx context.Context, in *BenchRequest, opts ...grpc.CallOption) (*BenchReply, error) {
	out := new(BenchReply)
	err := c.cc.Invoke(ctx, "/rpcbench.Benchmark/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BenchmarkServer is the server API for Benchmark service.
type BenchmarkServer interface {
	UnaryCall(context.Context, *BenchRequest) (*BenchReply, error)
}

// UnimplementedBenchmarkServer can be embedded to have forward compatible implementations.
type UnimplementedBenchmarkServer struct {
}

func (*UnimplementedBenchmarkServer) UnaryCall(ctx context.Context, req *BenchRequest) (*BenchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}

func RegisterBenchmarkServer(s *grpc.Server, srv BenchmarkServer) {
	s.RegisterService(&_Benchmark_serviceDesc, srv)
}

func _Benchmark_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BenchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchmarkServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcbench.Benchmark/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchmarkServer).UnaryCall(ctx, req.(*BenchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Benchmark_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcbench.Benchmark",
	HandlerType: (*BenchmarkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _Benchmark_UnaryCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}

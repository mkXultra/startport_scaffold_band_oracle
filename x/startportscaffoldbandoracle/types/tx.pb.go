// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: startportscaffoldbandoracle/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("startportscaffoldbandoracle/tx.proto", fileDescriptor_6b29ab0b04aa190f)
}

var fileDescriptor_6b29ab0b04aa190f = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x2e, 0x49, 0x2c,
	0x2a, 0x29, 0xc8, 0x2f, 0x2a, 0x29, 0x4e, 0x4e, 0x4c, 0x4b, 0xcb, 0xcf, 0x49, 0x49, 0x4a, 0xcc,
	0x4b, 0xc9, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x72, 0xc8, 0xcd, 0x8e, 0x28, 0xcd, 0x29, 0x29, 0x4a, 0xd4, 0xc3, 0xa3, 0x1c, 0x9f, 0x9c,
	0x11, 0x2b, 0x17, 0xb3, 0x6f, 0x71, 0xba, 0x53, 0xc1, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x44, 0x85, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3,
	0x6c, 0xd3, 0x87, 0x9b, 0x18, 0x0f, 0x33, 0x32, 0x1e, 0x64, 0x66, 0x3c, 0xd4, 0x7d, 0x15, 0xfa,
	0x78, 0x5d, 0x5f, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x98, 0xcd, 0x99, 0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mkXultra.startportscaffoldbandoracle.startportscaffoldbandoracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "startportscaffoldbandoracle/tx.proto",
}

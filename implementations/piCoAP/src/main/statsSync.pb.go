// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statsSync.proto

package main

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50bc6960f130f63, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Device struct {
	Building             int32    `protobuf:"varint,1,opt,name=building,proto3" json:"building,omitempty"`
	Room                 int32    `protobuf:"varint,2,opt,name=room,proto3" json:"room,omitempty"`
	LedOn                bool     `protobuf:"varint,3,opt,name=ledOn,proto3" json:"ledOn,omitempty"`
	Label                string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	OnTime               int32    `protobuf:"varint,5,opt,name=onTime,proto3" json:"onTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50bc6960f130f63, []int{1}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetBuilding() int32 {
	if m != nil {
		return m.Building
	}
	return 0
}

func (m *Device) GetRoom() int32 {
	if m != nil {
		return m.Room
	}
	return 0
}

func (m *Device) GetLedOn() bool {
	if m != nil {
		return m.LedOn
	}
	return false
}

func (m *Device) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Device) GetOnTime() int32 {
	if m != nil {
		return m.OnTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "syncProto.Empty")
	proto.RegisterType((*Device)(nil), "syncProto.Device")
}

func init() { proto.RegisterFile("statsSync.proto", fileDescriptor_a50bc6960f130f63) }

var fileDescriptor_a50bc6960f130f63 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2e, 0x49, 0x2c,
	0x29, 0x0e, 0xae, 0xcc, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x06, 0xb2,
	0x03, 0x40, 0x4c, 0x25, 0x76, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0xa5, 0x1a, 0x2e, 0x36,
	0x97, 0xd4, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x29, 0x2e, 0x8e, 0xa4, 0xd2, 0xcc, 0x9c, 0x94, 0xcc,
	0xbc, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x38, 0x5f, 0x48, 0x88, 0x8b, 0xa5, 0x28,
	0x3f, 0x3f, 0x57, 0x82, 0x09, 0x2c, 0x0e, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0xe6, 0xa4, 0xa6, 0xf8,
	0xe7, 0x49, 0x30, 0x03, 0x05, 0x39, 0x82, 0x20, 0x1c, 0xb0, 0x68, 0x62, 0x52, 0x6a, 0x8e, 0x04,
	0x0b, 0x50, 0x94, 0x33, 0x08, 0xc2, 0x11, 0x12, 0xe3, 0x62, 0xcb, 0xcf, 0x0b, 0xc9, 0xcc, 0x4d,
	0x95, 0x60, 0x05, 0x9b, 0x00, 0xe5, 0x19, 0x6d, 0x60, 0xe4, 0xe2, 0x01, 0x39, 0x30, 0xa3, 0x28,
	0x3f, 0x2f, 0xb3, 0x2a, 0xb5, 0x48, 0xc8, 0x80, 0x8b, 0x05, 0xc4, 0x17, 0x12, 0xd4, 0x83, 0xbb,
	0x55, 0x0f, 0xe2, 0x3e, 0x29, 0x4c, 0x21, 0x25, 0x06, 0x03, 0x46, 0x21, 0x33, 0x2e, 0x5e, 0x9f,
	0xcc, 0xf4, 0x8c, 0x92, 0xe0, 0xf2, 0xcc, 0x92, 0xe4, 0x0c, 0xa0, 0x11, 0x58, 0xb4, 0x0a, 0x20,
	0x09, 0x41, 0xbc, 0xcd, 0x00, 0xd4, 0xc7, 0xed, 0x93, 0x59, 0x5c, 0x02, 0x51, 0x51, 0x2c, 0x84,
	0xa1, 0x04, 0x87, 0x7d, 0x49, 0x6c, 0xe0, 0xb0, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x75,
	0xb5, 0x92, 0x7e, 0x5e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynchronizerClient interface {
	Sync(ctx context.Context, in *Device, opts ...grpc.CallOption) (Synchronizer_SyncClient, error)
	LightSwitcher(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error)
	ListDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Synchronizer_ListDevicesClient, error)
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

func (c *synchronizerClient) Sync(ctx context.Context, in *Device, opts ...grpc.CallOption) (Synchronizer_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synchronizer_serviceDesc.Streams[0], "/syncProto.Synchronizer/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &synchronizerSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synchronizer_SyncClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type synchronizerSyncClient struct {
	grpc.ClientStream
}

func (x *synchronizerSyncClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *synchronizerClient) LightSwitcher(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/syncProto.Synchronizer/LightSwitcher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizerClient) ListDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Synchronizer_ListDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synchronizer_serviceDesc.Streams[1], "/syncProto.Synchronizer/ListDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &synchronizerListDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synchronizer_ListDevicesClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type synchronizerListDevicesClient struct {
	grpc.ClientStream
}

func (x *synchronizerListDevicesClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SynchronizerServer is the server API for Synchronizer service.
type SynchronizerServer interface {
	Sync(*Device, Synchronizer_SyncServer) error
	LightSwitcher(context.Context, *Device) (*Empty, error)
	ListDevices(*Empty, Synchronizer_ListDevicesServer) error
}

// UnimplementedSynchronizerServer can be embedded to have forward compatible implementations.
type UnimplementedSynchronizerServer struct {
}

func (*UnimplementedSynchronizerServer) Sync(req *Device, srv Synchronizer_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (*UnimplementedSynchronizerServer) LightSwitcher(ctx context.Context, req *Device) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LightSwitcher not implemented")
}
func (*UnimplementedSynchronizerServer) ListDevices(req *Empty, srv Synchronizer_ListDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

func _Synchronizer_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Device)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SynchronizerServer).Sync(m, &synchronizerSyncServer{stream})
}

type Synchronizer_SyncServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type synchronizerSyncServer struct {
	grpc.ServerStream
}

func (x *synchronizerSyncServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

func _Synchronizer_LightSwitcher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizerServer).LightSwitcher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncProto.Synchronizer/LightSwitcher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizerServer).LightSwitcher(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronizer_ListDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SynchronizerServer).ListDevices(m, &synchronizerListDevicesServer{stream})
}

type Synchronizer_ListDevicesServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type synchronizerListDevicesServer struct {
	grpc.ServerStream
}

func (x *synchronizerListDevicesServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "syncProto.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LightSwitcher",
			Handler:    _Synchronizer_LightSwitcher_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _Synchronizer_Sync_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListDevices",
			Handler:       _Synchronizer_ListDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "statsSync.proto",
}

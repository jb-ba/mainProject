// Code generated by protoc-gen-go. DO NOT EDIT.
// source: implServer.proto

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ButtonMessage struct {
	Milli                *int32   `protobuf:"varint,1,req,name=milli" json:"milli,omitempty"`
	Sec                  *int32   `protobuf:"varint,2,req,name=sec" json:"sec,omitempty"`
	Label                *string  `protobuf:"bytes,3,req,name=label" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ButtonMessage) Reset()         { *m = ButtonMessage{} }
func (m *ButtonMessage) String() string { return proto.CompactTextString(m) }
func (*ButtonMessage) ProtoMessage()    {}
func (*ButtonMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_implServer_f0f6fdf609b6e264, []int{0}
}
func (m *ButtonMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButtonMessage.Unmarshal(m, b)
}
func (m *ButtonMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButtonMessage.Marshal(b, m, deterministic)
}
func (dst *ButtonMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButtonMessage.Merge(dst, src)
}
func (m *ButtonMessage) XXX_Size() int {
	return xxx_messageInfo_ButtonMessage.Size(m)
}
func (m *ButtonMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ButtonMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ButtonMessage proto.InternalMessageInfo

func (m *ButtonMessage) GetMilli() int32 {
	if m != nil && m.Milli != nil {
		return *m.Milli
	}
	return 0
}

func (m *ButtonMessage) GetSec() int32 {
	if m != nil && m.Sec != nil {
		return *m.Sec
	}
	return 0
}

func (m *ButtonMessage) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func init() {
	proto.RegisterType((*ButtonMessage)(nil), "main.ButtonMessage")
}

func init() { proto.RegisterFile("implServer.proto", fileDescriptor_implServer_f0f6fdf609b6e264) }

var fileDescriptor_implServer_f0f6fdf609b6e264 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcc, 0x2d, 0xc8,
	0x09, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d,
	0xcc, 0xcc, 0x53, 0xf2, 0xe5, 0xe2, 0x75, 0x2a, 0x2d, 0x29, 0xc9, 0xcf, 0xf3, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1, 0x62, 0xcd, 0xcd, 0xcc, 0xc9, 0xc9, 0x94, 0x60, 0x54, 0x60,
	0xd2, 0x60, 0x0d, 0x82, 0x70, 0x84, 0x04, 0xb8, 0x98, 0x8b, 0x53, 0x93, 0x25, 0x98, 0xc0, 0x62,
	0x20, 0x26, 0x48, 0x5d, 0x4e, 0x62, 0x52, 0x6a, 0x8e, 0x04, 0x33, 0x50, 0x8c, 0x33, 0x08, 0xc2,
	0x01, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x35, 0x69, 0x6d, 0x67, 0x00, 0x00, 0x00,
}

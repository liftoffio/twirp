// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package twirptest is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Hat
	Size
	Req
	Resp
*/
package twirptest

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

type Hat struct {
	Size  int32  `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color" json:"color,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Hat) Reset()                    { *m = Hat{} }
func (m *Hat) String() string            { return proto.CompactTextString(m) }
func (*Hat) ProtoMessage()               {}
func (*Hat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hat) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Hat) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Hat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Size struct {
	Inches int32 `protobuf:"varint,1,opt,name=inches" json:"inches,omitempty"`
}

func (m *Size) Reset()                    { *m = Size{} }
func (m *Size) String() string            { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()               {}
func (*Size) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Size) GetInches() int32 {
	if m != nil {
		return m.Inches
	}
	return 0
}

type Req struct {
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Resp struct {
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Hat)(nil), "twirp.internal.twirptest.Hat")
	proto.RegisterType((*Size)(nil), "twirp.internal.twirptest.Size")
	proto.RegisterType((*Req)(nil), "twirp.internal.twirptest.Req")
	proto.RegisterType((*Resp)(nil), "twirp.internal.twirptest.Resp")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x65, 0x9b, 0x0f, 0xd3, 0x09, 0x5e, 0x16, 0x91, 0x50, 0x30, 0x94, 0x9c, 0x72, 0x0a, 0xa5,
	0xfe, 0x03, 0xeb, 0x21, 0x45, 0x04, 0x49, 0x15, 0xc1, 0xdb, 0x34, 0x1d, 0xe8, 0x62, 0xb2, 0x9b,
	0xee, 0x8e, 0x16, 0xfc, 0x73, 0xfe, 0x35, 0x49, 0x2c, 0xde, 0xda, 0x4b, 0xbd, 0xbd, 0x8f, 0x7d,
	0x6f, 0xdf, 0x61, 0xe0, 0xd2, 0x91, 0xfd, 0x54, 0x35, 0x15, 0x9d, 0x35, 0x6c, 0x64, 0xc2, 0x7b,
	0x65, 0xbb, 0x42, 0x69, 0x26, 0xab, 0xb1, 0x29, 0x06, 0xca, 0xe4, 0x38, 0x5b, 0x80, 0x57, 0x22,
	0x4b, 0x09, 0xbe, 0x53, 0x5f, 0x94, 0x88, 0xa9, 0xc8, 0x83, 0x6a, 0xc0, 0xf2, 0x0a, 0x82, 0xda,
	0x34, 0xc6, 0x26, 0xa3, 0xa9, 0xc8, 0xc7, 0xd5, 0x2f, 0xe9, 0x5f, 0x6a, 0x6c, 0x29, 0xf1, 0x06,
	0x71, 0xc0, 0x59, 0x0a, 0xfe, 0xaa, 0x4f, 0x5c, 0x43, 0xa8, 0x74, 0xbd, 0x25, 0x77, 0xe8, 0x39,
	0xb0, 0x2c, 0x00, 0xaf, 0xa2, 0x5d, 0x16, 0x82, 0x5f, 0x91, 0xeb, 0xe6, 0xaf, 0x10, 0x97, 0xb8,
	0x26, 0xbb, 0x41, 0xb7, 0x25, 0x2b, 0x4b, 0xb8, 0x78, 0xc4, 0x77, 0xea, 0x67, 0xa4, 0xc5, 0xb1,
	0xa1, 0x45, 0xff, 0xc1, 0xe4, 0xe6, 0xb8, 0x5f, 0x22, 0xcf, 0xbf, 0x47, 0x10, 0xad, 0xd8, 0x12,
	0xb6, 0x64, 0xe5, 0x12, 0xa2, 0x67, 0x8b, 0xda, 0x61, 0xcd, 0xf2, 0x44, 0xae, 0xa2, 0xdd, 0x24,
	0x3d, 0x65, 0xbb, 0x4e, 0x2e, 0x21, 0x7c, 0xe9, 0x1a, 0x83, 0x9b, 0x33, 0x8b, 0x72, 0x21, 0x1f,
	0x20, 0xba, 0x37, 0x7b, 0xfd, 0x0f, 0x65, 0x33, 0x21, 0x9f, 0x20, 0x5e, 0x98, 0xb6, 0xfd, 0xd0,
	0xaa, 0x46, 0xa6, 0xb3, 0xc7, 0xcd, 0xc4, 0x5d, 0xfc, 0x36, 0xfe, 0x53, 0xd7, 0xe1, 0x70, 0x3c,
	0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x64, 0xbb, 0x2f, 0x4d, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Rig.proto

package linuxCommand

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Responses struct {
	StdoutResponse       string   `protobuf:"bytes,1,opt,name=stdoutResponse,proto3" json:"stdoutResponse,omitempty"`
	StderrResponse       string   `protobuf:"bytes,2,opt,name=stderrResponse,proto3" json:"stderrResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Responses) Reset()         { *m = Responses{} }
func (m *Responses) String() string { return proto.CompactTextString(m) }
func (*Responses) ProtoMessage()    {}
func (*Responses) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a397db3166ae542, []int{0}
}

func (m *Responses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Responses.Unmarshal(m, b)
}
func (m *Responses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Responses.Marshal(b, m, deterministic)
}
func (m *Responses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Responses.Merge(m, src)
}
func (m *Responses) XXX_Size() int {
	return xxx_messageInfo_Responses.Size(m)
}
func (m *Responses) XXX_DiscardUnknown() {
	xxx_messageInfo_Responses.DiscardUnknown(m)
}

var xxx_messageInfo_Responses proto.InternalMessageInfo

func (m *Responses) GetStdoutResponse() string {
	if m != nil {
		return m.StdoutResponse
	}
	return ""
}

func (m *Responses) GetStderrResponse() string {
	if m != nil {
		return m.StderrResponse
	}
	return ""
}

type Rig struct {
	Nbr                  string   `protobuf:"bytes,1,opt,name=Nbr,proto3" json:"Nbr,omitempty"`
	Man                  bool     `protobuf:"varint,2,opt,name=Man,proto3" json:"Man,omitempty"`
	Woman                bool     `protobuf:"varint,3,opt,name=Woman,proto3" json:"Woman,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rig) Reset()         { *m = Rig{} }
func (m *Rig) String() string { return proto.CompactTextString(m) }
func (*Rig) ProtoMessage()    {}
func (*Rig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a397db3166ae542, []int{1}
}

func (m *Rig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rig.Unmarshal(m, b)
}
func (m *Rig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rig.Marshal(b, m, deterministic)
}
func (m *Rig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rig.Merge(m, src)
}
func (m *Rig) XXX_Size() int {
	return xxx_messageInfo_Rig.Size(m)
}
func (m *Rig) XXX_DiscardUnknown() {
	xxx_messageInfo_Rig.DiscardUnknown(m)
}

var xxx_messageInfo_Rig proto.InternalMessageInfo

func (m *Rig) GetNbr() string {
	if m != nil {
		return m.Nbr
	}
	return ""
}

func (m *Rig) GetMan() bool {
	if m != nil {
		return m.Man
	}
	return false
}

func (m *Rig) GetWoman() bool {
	if m != nil {
		return m.Woman
	}
	return false
}

func init() {
	proto.RegisterType((*Responses)(nil), "linuxCommand.Responses")
	proto.RegisterType((*Rig)(nil), "linuxCommand.Rig")
}

func init() {
	proto.RegisterFile("proto/Rig.proto", fileDescriptor_7a397db3166ae542)
}

var fileDescriptor_7a397db3166ae542 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x0f, 0xca, 0x4c, 0xd7, 0x03, 0xb3, 0x84, 0x78, 0x72, 0x32, 0xf3, 0x4a, 0x2b, 0x9c,
	0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x94, 0xa2, 0xb9, 0x38, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x8b, 0x85, 0xd4, 0xb8, 0xf8, 0x8a, 0x4b, 0x52, 0xf2, 0x4b, 0x4b, 0x60, 0x42, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x68, 0xa2, 0x50, 0x75, 0xa9, 0x45, 0x45, 0x70, 0x75, 0x4c,
	0x70, 0x75, 0x48, 0xa2, 0x4a, 0xf6, 0x5c, 0xcc, 0x41, 0x99, 0xe9, 0x42, 0x02, 0x5c, 0xcc, 0x7e,
	0x49, 0x45, 0x50, 0xb3, 0x40, 0x4c, 0x90, 0x88, 0x6f, 0x62, 0x1e, 0x58, 0x17, 0x47, 0x10, 0x88,
	0x29, 0x24, 0xc2, 0xc5, 0x1a, 0x9e, 0x9f, 0x9b, 0x98, 0x27, 0xc1, 0x0c, 0x16, 0x83, 0x70, 0x8c,
	0x5c, 0xb8, 0xb8, 0x82, 0x32, 0xd3, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xcc, 0xb8,
	0xd8, 0xdc, 0x53, 0x4b, 0x40, 0x26, 0x0a, 0xea, 0x21, 0x7b, 0x42, 0x2f, 0x28, 0x33, 0x5d, 0x4a,
	0x1c, 0x4d, 0x08, 0xe6, 0x29, 0x25, 0x06, 0x27, 0x81, 0x28, 0x3e, 0x3d, 0x6b, 0x64, 0xd9, 0x24,
	0x36, 0x70, 0x50, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xbb, 0x84, 0xd6, 0x1d, 0x01,
	0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/toilet.proto

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

type Toilet struct {
	S                    bool     `protobuf:"varint,1,opt,name=S,proto3" json:"S,omitempty"`
	SS                   bool     `protobuf:"varint,2,opt,name=SS,proto3" json:"SS,omitempty"`
	K                    bool     `protobuf:"varint,3,opt,name=K,proto3" json:"K,omitempty"`
	W                    bool     `protobuf:"varint,4,opt,name=W,proto3" json:"W,omitempty"`
	O                    bool     `protobuf:"varint,5,opt,name=O,proto3" json:"O,omitempty"`
	F                    string   `protobuf:"bytes,6,opt,name=F,proto3" json:"F,omitempty"`
	FF                   string   `protobuf:"bytes,7,opt,name=FF,proto3" json:"FF,omitempty"`
	FFF                  string   `protobuf:"bytes,8,opt,name=FFF,proto3" json:"FFF,omitempty"`
	E                    string   `protobuf:"bytes,9,opt,name=E,proto3" json:"E,omitempty"`
	F3                   []string `protobuf:"bytes,10,rep,name=F3,proto3" json:"F3,omitempty"`
	Message              string   `protobuf:"bytes,11,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Toilet) Reset()         { *m = Toilet{} }
func (m *Toilet) String() string { return proto.CompactTextString(m) }
func (*Toilet) ProtoMessage()    {}
func (*Toilet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f7adcbf5a8a08a, []int{0}
}

func (m *Toilet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Toilet.Unmarshal(m, b)
}
func (m *Toilet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Toilet.Marshal(b, m, deterministic)
}
func (m *Toilet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Toilet.Merge(m, src)
}
func (m *Toilet) XXX_Size() int {
	return xxx_messageInfo_Toilet.Size(m)
}
func (m *Toilet) XXX_DiscardUnknown() {
	xxx_messageInfo_Toilet.DiscardUnknown(m)
}

var xxx_messageInfo_Toilet proto.InternalMessageInfo

func (m *Toilet) GetS() bool {
	if m != nil {
		return m.S
	}
	return false
}

func (m *Toilet) GetSS() bool {
	if m != nil {
		return m.SS
	}
	return false
}

func (m *Toilet) GetK() bool {
	if m != nil {
		return m.K
	}
	return false
}

func (m *Toilet) GetW() bool {
	if m != nil {
		return m.W
	}
	return false
}

func (m *Toilet) GetO() bool {
	if m != nil {
		return m.O
	}
	return false
}

func (m *Toilet) GetF() string {
	if m != nil {
		return m.F
	}
	return ""
}

func (m *Toilet) GetFF() string {
	if m != nil {
		return m.FF
	}
	return ""
}

func (m *Toilet) GetFFF() string {
	if m != nil {
		return m.FFF
	}
	return ""
}

func (m *Toilet) GetE() string {
	if m != nil {
		return m.E
	}
	return ""
}

func (m *Toilet) GetF3() []string {
	if m != nil {
		return m.F3
	}
	return nil
}

func (m *Toilet) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Toilet)(nil), "linuxCommand.Toilet")
}

func init() {
	proto.RegisterFile("proto/toilet.proto", fileDescriptor_c2f7adcbf5a8a08a)
}

var fileDescriptor_c2f7adcbf5a8a08a = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcd, 0x4a, 0xf3, 0x50,
	0x10, 0x86, 0xbf, 0x93, 0x7c, 0xa6, 0xcd, 0x58, 0x63, 0x3d, 0x16, 0x1c, 0xb2, 0x0a, 0x5d, 0x75,
	0x15, 0xc1, 0x2c, 0x75, 0xe1, 0x0f, 0x19, 0x17, 0x45, 0x0a, 0x39, 0x42, 0xc1, 0x5d, 0xd5, 0x83,
	0x04, 0xf2, 0x47, 0x13, 0x45, 0xef, 0xc3, 0x8b, 0xf1, 0xf2, 0xe4, 0x9c, 0xb1, 0x92, 0xee, 0xa4,
	0xbb, 0x3c, 0xef, 0x3c, 0x79, 0x99, 0x4c, 0x40, 0x36, 0xeb, 0xba, 0xab, 0x4f, 0xbb, 0x3a, 0x2f,
	0x74, 0x17, 0x5b, 0x90, 0xa3, 0x22, 0xaf, 0x5e, 0xdf, 0x6f, 0xea, 0xb2, 0x5c, 0x55, 0xcf, 0xe1,
	0x11, 0x1b, 0x69, 0xd9, 0x74, 0x1f, 0x2c, 0x84, 0x13, 0x8e, 0x32, 0xdd, 0x36, 0x75, 0xd5, 0x6a,
	0x4e, 0xa7, 0x5f, 0x02, 0xbc, 0x7b, 0xdb, 0x23, 0x47, 0x20, 0x14, 0x8a, 0x48, 0xcc, 0x86, 0x99,
	0x50, 0x32, 0x00, 0x47, 0x29, 0x74, 0x2c, 0x3a, 0x4a, 0x99, 0xe9, 0x1c, 0x5d, 0x9e, 0xce, 0x0d,
	0x2d, 0xf1, 0x3f, 0xd3, 0xd2, 0xd0, 0x02, 0xf7, 0x98, 0x16, 0x86, 0x08, 0xbd, 0x48, 0xcc, 0xfc,
	0x4c, 0x90, 0xe9, 0x21, 0xc2, 0x81, 0x45, 0x87, 0x48, 0x8e, 0xc1, 0x25, 0x22, 0x1c, 0xda, 0xc0,
	0x3c, 0x1a, 0x3f, 0x45, 0x9f, 0xfd, 0xd4, 0xfa, 0x09, 0x42, 0xe4, 0x5a, 0x3f, 0x91, 0x08, 0x83,
	0x3b, 0xdd, 0xb6, 0xab, 0x17, 0x8d, 0xfb, 0xd6, 0xd9, 0xe0, 0xd9, 0xa7, 0x03, 0x07, 0xbc, 0xba,
	0xd2, 0xeb, 0xb7, 0xfc, 0x49, 0xcb, 0x0b, 0xf0, 0x6f, 0x75, 0xf7, 0xf3, 0x39, 0x93, 0xb8, 0x7f,
	0x91, 0x98, 0xd3, 0xf0, 0x64, 0x3b, 0xdd, 0x5c, 0xa3, 0x9d, 0xfe, 0x93, 0x97, 0x10, 0xfc, 0xbe,
	0x4d, 0x94, 0x17, 0x5a, 0x1e, 0x6f, 0xcb, 0xf6, 0x9a, 0x7f, 0x6d, 0x48, 0x77, 0x6a, 0xb8, 0x82,
	0xc3, 0xde, 0x0e, 0xbb, 0x54, 0x5c, 0x8f, 0x1f, 0x82, 0xf8, 0xbc, 0x3f, 0x7d, 0xf4, 0xec, 0xaf,
	0x4e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x60, 0x1f, 0xbb, 0xf5, 0x37, 0x02, 0x00, 0x00,
}

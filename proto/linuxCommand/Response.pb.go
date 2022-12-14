// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Response.proto

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
	StdoutResponse       string   `protobuf:"bytes,1,opt,name=StdoutResponse,proto3" json:"StdoutResponse,omitempty"`
	StderrResponse       string   `protobuf:"bytes,2,opt,name=StderrResponse,proto3" json:"StderrResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Responses) Reset()         { *m = Responses{} }
func (m *Responses) String() string { return proto.CompactTextString(m) }
func (*Responses) ProtoMessage()    {}
func (*Responses) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a746925d5b434a1, []int{0}
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

func init() {
	proto.RegisterType((*Responses)(nil), "linuxCommand.Responses")
}

func init() {
	proto.RegisterFile("proto/Response.proto", fileDescriptor_7a746925d5b434a1)
}

var fileDescriptor_7a746925d5b434a1 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x0f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0x73, 0x85, 0x78, 0x72,
	0x32, 0xf3, 0x4a, 0x2b, 0x9c, 0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x94, 0xa2, 0xb9, 0x38, 0x61,
	0xf2, 0xc5, 0x42, 0x6a, 0x5c, 0x7c, 0xc1, 0x25, 0x29, 0xf9, 0xa5, 0x25, 0x30, 0x21, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x34, 0x51, 0xa8, 0xba, 0xd4, 0xa2, 0x22, 0xb8, 0x3a, 0x26, 0xb8,
	0x3a, 0x24, 0x51, 0x27, 0x81, 0x28, 0x3e, 0x3d, 0x6b, 0x64, 0xeb, 0x92, 0xd8, 0xc0, 0x6e, 0x30,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30, 0x18, 0xc6, 0xec, 0x9b, 0x00, 0x00, 0x00,
}

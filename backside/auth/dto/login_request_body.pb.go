// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/proto/login_request_body.proto

package dto

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

type LoginRequestBody struct {
	IdToken              string   `protobuf:"bytes,1,opt,name=idToken,proto3" json:"idToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequestBody) Reset()         { *m = LoginRequestBody{} }
func (m *LoginRequestBody) String() string { return proto.CompactTextString(m) }
func (*LoginRequestBody) ProtoMessage()    {}
func (*LoginRequestBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbcb750f17043ea, []int{0}
}

func (m *LoginRequestBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequestBody.Unmarshal(m, b)
}
func (m *LoginRequestBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequestBody.Marshal(b, m, deterministic)
}
func (m *LoginRequestBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequestBody.Merge(m, src)
}
func (m *LoginRequestBody) XXX_Size() int {
	return xxx_messageInfo_LoginRequestBody.Size(m)
}
func (m *LoginRequestBody) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequestBody.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequestBody proto.InternalMessageInfo

func (m *LoginRequestBody) GetIdToken() string {
	if m != nil {
		return m.IdToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequestBody)(nil), "main.LoginRequestBody")
}

func init() {
	proto.RegisterFile("auth/proto/login_request_body.proto", fileDescriptor_2fbcb750f17043ea)
}

var fileDescriptor_2fbcb750f17043ea = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0xc9, 0x4f, 0xcf, 0xcc, 0x8b, 0x2f, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x89, 0x4f, 0xca, 0x4f, 0xa9, 0xd4, 0x03, 0x4b, 0x08, 0xb1, 0xe4, 0x26,
	0x66, 0xe6, 0x29, 0xe9, 0x70, 0x09, 0xf8, 0x80, 0x54, 0x04, 0x41, 0x14, 0x38, 0xe5, 0xa7, 0x54,
	0x0a, 0x49, 0x70, 0xb1, 0x67, 0xa6, 0x84, 0xe4, 0x67, 0xa7, 0xe6, 0x49, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0xc1, 0xb8, 0x4e, 0xec, 0x51, 0xac, 0x7a, 0xfa, 0x29, 0x25, 0xf9, 0x49, 0x6c, 0x60,
	0x33, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0xea, 0x3b, 0xf2, 0x6a, 0x00, 0x00, 0x00,
}
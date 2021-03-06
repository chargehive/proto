// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chargehive/chtype/authentication.proto

package chtype

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AuthenticationChain struct {
	Parent               *AuthenticationChain `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	VerificationHash     string               `protobuf:"bytes,2,opt,name=verification_hash,json=verificationHash,proto3" json:"verification_hash,omitempty"`
	HandlerType          HandlerType          `protobuf:"varint,3,opt,name=handler_type,json=handlerType,proto3,enum=chargehive.chtype.HandlerType" json:"handler_type,omitempty"`
	Handler              string               `protobuf:"bytes,4,opt,name=handler,proto3" json:"handler,omitempty"`
	ProjectId            string               `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	UserId               string               `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserAgent            string               `protobuf:"bytes,7,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	IpAddress            string               `protobuf:"bytes,8,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthenticationChain) Reset()         { *m = AuthenticationChain{} }
func (m *AuthenticationChain) String() string { return proto.CompactTextString(m) }
func (*AuthenticationChain) ProtoMessage()    {}
func (*AuthenticationChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf019ef63386500e, []int{0}
}
func (m *AuthenticationChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationChain.Unmarshal(m, b)
}
func (m *AuthenticationChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationChain.Marshal(b, m, deterministic)
}
func (m *AuthenticationChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationChain.Merge(m, src)
}
func (m *AuthenticationChain) XXX_Size() int {
	return xxx_messageInfo_AuthenticationChain.Size(m)
}
func (m *AuthenticationChain) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationChain.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationChain proto.InternalMessageInfo

func (m *AuthenticationChain) GetParent() *AuthenticationChain {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *AuthenticationChain) GetVerificationHash() string {
	if m != nil {
		return m.VerificationHash
	}
	return ""
}

func (m *AuthenticationChain) GetHandlerType() HandlerType {
	if m != nil {
		return m.HandlerType
	}
	return HANDLER_TYPE_INVALID
}

func (m *AuthenticationChain) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *AuthenticationChain) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *AuthenticationChain) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthenticationChain) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *AuthenticationChain) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthenticationChain)(nil), "chargehive.chtype.AuthenticationChain")
	golang_proto.RegisterType((*AuthenticationChain)(nil), "chargehive.chtype.AuthenticationChain")
}

func init() {
	proto.RegisterFile("chargehive/chtype/authentication.proto", fileDescriptor_cf019ef63386500e)
}
func init() {
	golang_proto.RegisterFile("chargehive/chtype/authentication.proto", fileDescriptor_cf019ef63386500e)
}

var fileDescriptor_cf019ef63386500e = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x5f, 0xf2, 0xde, 0x4b, 0xed, 0x54, 0xc4, 0xc6, 0x85, 0xa1, 0x68, 0x29, 0x2e, 0x4a,
	0x41, 0x4c, 0xa0, 0xba, 0x16, 0x62, 0x37, 0xed, 0xc2, 0x4d, 0x71, 0x25, 0x85, 0x70, 0x9b, 0xdc,
	0x66, 0x46, 0xda, 0x4c, 0x98, 0x4c, 0x0a, 0xfd, 0x77, 0xae, 0x5d, 0xb9, 0xf1, 0x07, 0xa8, 0x7f,
	0x44, 0x66, 0x26, 0xa5, 0x95, 0x74, 0x97, 0x73, 0xbe, 0x7b, 0xee, 0x0d, 0x73, 0x48, 0x3f, 0xa6,
	0x20, 0x52, 0xa4, 0x6c, 0x8d, 0x41, 0x4c, 0xe5, 0x26, 0xc7, 0x00, 0x4a, 0x49, 0x31, 0x93, 0x2c,
	0x06, 0xc9, 0x78, 0xe6, 0xe7, 0x82, 0x4b, 0xee, 0xb6, 0x77, 0x73, 0xbe, 0x99, 0xeb, 0x5c, 0xd4,
	0xa3, 0x98, 0x95, 0x2b, 0x13, 0xe8, 0xdc, 0xa4, 0x4c, 0xd2, 0x72, 0xee, 0xc7, 0x7c, 0x15, 0xa4,
	0x3c, 0xe5, 0x81, 0xb6, 0xe7, 0xe5, 0x42, 0x2b, 0x2d, 0xf4, 0x97, 0x19, 0xbf, 0xfa, 0xb0, 0xc9,
	0x59, 0xf8, 0xeb, 0xf0, 0x88, 0x02, 0xcb, 0xdc, 0x7b, 0xe2, 0xe4, 0x20, 0x30, 0x93, 0x9e, 0xd5,
	0xb3, 0x06, 0xad, 0x61, 0xdf, 0xaf, 0xfd, 0x88, 0x7f, 0x20, 0x37, 0xad, 0x52, 0xee, 0x35, 0x69,
	0xaf, 0x51, 0xb0, 0x45, 0x05, 0x23, 0x0a, 0x05, 0xf5, 0xec, 0x9e, 0x35, 0x68, 0x4e, 0x4f, 0xf7,
	0xc1, 0x18, 0x0a, 0xea, 0x86, 0xe4, 0x98, 0x42, 0x96, 0x2c, 0x51, 0x44, 0x6a, 0xb1, 0xf7, 0xb7,
	0x67, 0x0d, 0x4e, 0x86, 0xdd, 0x03, 0x27, 0xc7, 0x66, 0xec, 0x69, 0x93, 0xe3, 0xb4, 0x45, 0x77,
	0xc2, 0xf5, 0x48, 0xa3, 0x92, 0xde, 0x3f, 0x7d, 0x65, 0x2b, 0xdd, 0x4b, 0x42, 0x72, 0xc1, 0x5f,
	0x30, 0x96, 0x11, 0x4b, 0xbc, 0xff, 0x1a, 0x36, 0x2b, 0x67, 0x92, 0xb8, 0xe7, 0xa4, 0x51, 0x16,
	0x28, 0x14, 0x73, 0x34, 0x73, 0x94, 0x9c, 0x24, 0x2a, 0xa7, 0x01, 0xa4, 0xea, 0x15, 0x1a, 0x26,
	0xa7, 0x9c, 0x50, 0x19, 0x0a, 0xb3, 0x3c, 0x82, 0x24, 0x11, 0x58, 0x14, 0xde, 0x91, 0xc1, 0x2c,
	0x0f, 0x8d, 0xf1, 0xc0, 0xdf, 0x3f, 0xbb, 0x7f, 0x5e, 0xbf, 0xbb, 0xd6, 0xf3, 0xdd, 0x5e, 0x21,
	0x7b, 0xcd, 0x6d, 0x9b, 0x58, 0x42, 0x96, 0x06, 0xb5, 0x46, 0xdf, 0xec, 0xf6, 0x48, 0x7b, 0x63,
	0xb6, 0xc6, 0xd9, 0x48, 0x7b, 0x5f, 0x76, 0xa7, 0xe6, 0xcd, 0x1e, 0x51, 0x42, 0x02, 0x12, 0xe6,
	0x8e, 0x5e, 0x78, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa9, 0x9f, 0xe0, 0x59, 0x02, 0x00,
	0x00,
}

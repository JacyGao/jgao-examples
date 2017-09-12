// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

/*
Package gateway is a generated protocol buffer package.

It is generated from these files:
	gateway.proto

It has these top-level messages:
	UserReq
	UserRes
	User
*/
package gateway

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

type UserReq struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserReq) Reset()                    { *m = UserReq{} }
func (m *UserReq) String() string            { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()               {}
func (*UserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserRes struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *UserRes) Reset()                    { *m = UserRes{} }
func (m *UserRes) String() string            { return proto.CompactTextString(m) }
func (*UserRes) ProtoMessage()               {}
func (*UserRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRes) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type User struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc" json:"desc,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*UserReq)(nil), "gateway.UserReq")
	proto.RegisterType((*UserRes)(nil), "gateway.UserRes")
	proto.RegisterType((*User)(nil), "gateway.User")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x24, 0xb9,
	0xd8, 0x43, 0x8b, 0x53, 0x8b, 0x82, 0x52, 0x0b, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0x14, 0x61, 0x52, 0xc5, 0x42, 0x62, 0x5c,
	0x6c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x60, 0x69, 0x8e, 0x20, 0x28, 0x4f, 0x49, 0x8b, 0x8b,
	0x05, 0xa4, 0x04, 0x5d, 0xab, 0x90, 0x10, 0x17, 0x4b, 0x4a, 0x6a, 0x71, 0xb2, 0x04, 0x13, 0x58,
	0x04, 0xcc, 0x4e, 0x62, 0x03, 0xdb, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x70, 0xc7, 0xa4,
	0x45, 0x8a, 0x00, 0x00, 0x00,
}

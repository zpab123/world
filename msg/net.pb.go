// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net.proto

package msg

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

// 客户端->服务器 握手消息
type HandshakeReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Acceptor             uint32   `protobuf:"varint,2,opt,name=acceptor,proto3" json:"acceptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeReq) Reset()         { *m = HandshakeReq{} }
func (m *HandshakeReq) String() string { return proto.CompactTextString(m) }
func (*HandshakeReq) ProtoMessage()    {}
func (*HandshakeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{0}
}

func (m *HandshakeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeReq.Unmarshal(m, b)
}
func (m *HandshakeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeReq.Marshal(b, m, deterministic)
}
func (m *HandshakeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeReq.Merge(m, src)
}
func (m *HandshakeReq) XXX_Size() int {
	return xxx_messageInfo_HandshakeReq.Size(m)
}
func (m *HandshakeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeReq.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeReq proto.InternalMessageInfo

func (m *HandshakeReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HandshakeReq) GetAcceptor() uint32 {
	if m != nil {
		return m.Acceptor
	}
	return 0
}

// 服务器->客户端 握手结果
type HandshakeRes struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Heartbeat            uint32   `protobuf:"varint,2,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeRes) Reset()         { *m = HandshakeRes{} }
func (m *HandshakeRes) String() string { return proto.CompactTextString(m) }
func (*HandshakeRes) ProtoMessage()    {}
func (*HandshakeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{1}
}

func (m *HandshakeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeRes.Unmarshal(m, b)
}
func (m *HandshakeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeRes.Marshal(b, m, deterministic)
}
func (m *HandshakeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeRes.Merge(m, src)
}
func (m *HandshakeRes) XXX_Size() int {
	return xxx_messageInfo_HandshakeRes.Size(m)
}
func (m *HandshakeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeRes.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeRes proto.InternalMessageInfo

func (m *HandshakeRes) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HandshakeRes) GetHeartbeat() uint32 {
	if m != nil {
		return m.Heartbeat
	}
	return 0
}

func init() {
	proto.RegisterType((*HandshakeReq)(nil), "msg.HandshakeReq")
	proto.RegisterType((*HandshakeRes)(nil), "msg.HandshakeRes")
}

func init() { proto.RegisterFile("net.proto", fileDescriptor_a5b10ce944527a32) }

var fileDescriptor_a5b10ce944527a32 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4b, 0x2d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xb2, 0xe1, 0xe2, 0xf1, 0x48,
	0xcc, 0x4b, 0x29, 0xce, 0x48, 0xcc, 0x4e, 0x0d, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0xce, 0x4e,
	0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0xa4, 0xb8, 0x38, 0x12, 0x93,
	0x93, 0x53, 0x0b, 0x4a, 0xf2, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78, 0x83, 0xe0, 0x7c, 0x25,
	0x07, 0x14, 0xdd, 0xc5, 0x42, 0x42, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x60, 0xed, 0xbc, 0x41,
	0x60, 0xb6, 0x90, 0x0c, 0x17, 0x67, 0x46, 0x6a, 0x62, 0x51, 0x49, 0x52, 0x6a, 0x62, 0x09, 0xd4,
	0x00, 0x84, 0x40, 0x12, 0x1b, 0xd8, 0x2d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x98,
	0xa0, 0xf2, 0x98, 0x00, 0x00, 0x00,
}

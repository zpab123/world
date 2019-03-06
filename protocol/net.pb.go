// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net.proto

package protocol

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	proto.RegisterType((*HandshakeReq)(nil), "protocol.HandshakeReq")
	proto.RegisterType((*HandshakeRes)(nil), "protocol.HandshakeRes")
}

func init() { proto.RegisterFile("net.proto", fileDescriptor_a5b10ce944527a32) }

var fileDescriptor_a5b10ce944527a32 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4b, 0x2d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x36, 0x5c, 0x3c,
	0x1e, 0x89, 0x79, 0x29, 0xc5, 0x19, 0x89, 0xd9, 0xa9, 0x41, 0xa9, 0x85, 0x42, 0x02, 0x5c, 0xcc,
	0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x14, 0x17, 0x47,
	0x62, 0x72, 0x72, 0x6a, 0x41, 0x49, 0x7e, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x9c,
	0xaf, 0xe4, 0x80, 0xa2, 0xbb, 0x58, 0x48, 0x88, 0x8b, 0x25, 0x39, 0x3f, 0x25, 0x15, 0xac, 0x9d,
	0x37, 0x08, 0xcc, 0x16, 0x92, 0xe1, 0xe2, 0xcc, 0x48, 0x4d, 0x2c, 0x2a, 0x49, 0x4a, 0x4d, 0x2c,
	0x81, 0x1a, 0x80, 0x10, 0x48, 0x62, 0x03, 0xbb, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3a,
	0x94, 0x28, 0xc9, 0x9d, 0x00, 0x00, 0x00,
}
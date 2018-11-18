// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_send_ack.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	message_send_ack.proto

It has these top-level messages:
	MessageSendACK
*/
package pb

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

// 消息发送回执
type MessageSendACK struct {
	SendSequence int64 `protobuf:"varint,1,opt,name=send_sequence,json=sendSequence" json:"send_sequence,omitempty"`
	Code         int32 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
}

func (m *MessageSendACK) Reset()                    { *m = MessageSendACK{} }
func (m *MessageSendACK) String() string            { return proto.CompactTextString(m) }
func (*MessageSendACK) ProtoMessage()               {}
func (*MessageSendACK) Descriptor() ([]byte, []int) { return messageSendACKFileDescriptor, []int{0} }

func (m *MessageSendACK) GetSendSequence() int64 {
	if m != nil {
		return m.SendSequence
	}
	return 0
}

func (m *MessageSendACK) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageSendACK)(nil), "pb.MessageSendACK")
}

func init() { proto.RegisterFile("message_send_ack.proto", messageSendACKFileDescriptor) }

var messageSendACKFileDescriptor = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x8d, 0x2f, 0x4e, 0xcd, 0x4b, 0x89, 0x4f, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe4, 0xe2, 0xf3, 0x85, 0xc8, 0x06, 0xa7, 0xe6,
	0xa5, 0x38, 0x3a, 0x7b, 0x0b, 0x29, 0x73, 0xf1, 0x82, 0xd5, 0x15, 0xa7, 0x16, 0x96, 0xa6, 0xe6,
	0x25, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xf1, 0x80, 0x04, 0x83, 0xa1, 0x62, 0x42,
	0x42, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60, 0x76,
	0x12, 0x1b, 0xd8, 0x54, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x20, 0x51, 0xf7, 0x6f,
	0x00, 0x00, 0x00,
}
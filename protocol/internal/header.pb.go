// Code generated by protoc-gen-go. DO NOT EDIT.
// source: header.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	header.proto
	payload.proto

It has these top-level messages:
	Header
	PayloadBandwidth
	PayloadInspect
	PayloadHello
	PayloadConnect
	PayloadDisconnect
	PayloadNode
	PayloadHeartbeat
	PayloadTunnelEvent
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Header struct {
	// id is the header identifier. Example: Hello message (represented as int).
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// modules contains the ids of the other modules present in the package.
	Modules []string `protobuf:"bytes,2,rep,name=modules" json:"modules,omitempty"`
	// sentAt tells when the message was send.
	SentAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=sentAt" json:"sentAt,omitempty"`
	// protocolVersion used by the packet.
	ProtocolVersion string `protobuf:"bytes,4,opt,name=protocolVersion" json:"protocolVersion,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Header) GetModules() []string {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *Header) GetSentAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.SentAt
	}
	return nil
}

func (m *Header) GetProtocolVersion() string {
	if m != nil {
		return m.ProtocolVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*Header)(nil), "internal.Header")
}

func init() { proto.RegisterFile("header.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8d, 0xb1, 0x0e, 0xc2, 0x20,
	0x14, 0x45, 0x43, 0xab, 0xd5, 0xa2, 0xd1, 0x84, 0x89, 0x74, 0x91, 0x38, 0x31, 0xd1, 0xa4, 0x7e,
	0x81, 0x9b, 0x33, 0x31, 0xee, 0xad, 0x3c, 0x2b, 0x09, 0x85, 0x06, 0xe8, 0x7f, 0xf8, 0xc9, 0x26,
	0xd4, 0x2e, 0x8e, 0xef, 0xbc, 0x73, 0x73, 0xf0, 0xfe, 0x0d, 0xad, 0x02, 0x2f, 0x46, 0xef, 0xa2,
	0x23, 0x5b, 0x6d, 0x23, 0x78, 0xdb, 0x9a, 0xea, 0xd4, 0x3b, 0xd7, 0x1b, 0xa8, 0x13, 0xef, 0xa6,
	0x57, 0x1d, 0xf5, 0x00, 0x21, 0xb6, 0xc3, 0x38, 0xab, 0xe7, 0x0f, 0xc2, 0xc5, 0x2d, 0x6d, 0xc9,
	0x01, 0x67, 0x5a, 0x51, 0xc4, 0x10, 0x5f, 0xcb, 0x4c, 0x2b, 0x42, 0xf1, 0x66, 0x70, 0x6a, 0x32,
	0x10, 0x68, 0xc6, 0x72, 0x5e, 0xca, 0xe5, 0x24, 0x0d, 0x2e, 0x02, 0xd8, 0x78, 0x8d, 0x34, 0x67,
	0x88, 0xef, 0x9a, 0x4a, 0xcc, 0x19, 0xb1, 0x64, 0xc4, 0x7d, 0xc9, 0xc8, 0x9f, 0x49, 0x38, 0x3e,
	0xa6, 0xef, 0xd3, 0x99, 0x07, 0xf8, 0xa0, 0x9d, 0xa5, 0x2b, 0x86, 0x78, 0x29, 0xff, 0x71, 0x57,
	0x24, 0x70, 0xf9, 0x06, 0x00, 0x00, 0xff, 0xff, 0x66, 0x1e, 0xc7, 0xdd, 0xd4, 0x00, 0x00, 0x00,
}

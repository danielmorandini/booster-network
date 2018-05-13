/*
Copyright (C) 2018 Daniel Morandini

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payload.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PayloadCtrl struct {
	// operation is the operation that has to be performed.
	// See Ctrl operations in protocol.go
	Operation int32 `protobuf:"varint,1,opt,name=operation" json:"operation,omitempty"`
}

func (m *PayloadCtrl) Reset()                    { *m = PayloadCtrl{} }
func (m *PayloadCtrl) String() string            { return proto.CompactTextString(m) }
func (*PayloadCtrl) ProtoMessage()               {}
func (*PayloadCtrl) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *PayloadCtrl) GetOperation() int32 {
	if m != nil {
		return m.Operation
	}
	return 0
}

type PayloadBandwidth struct {
	// tot is the total number of bytes transmitted.
	Tot int64 `protobuf:"varint,1,opt,name=tot" json:"tot,omitempty"`
	// bandwidth is the current bandwidth.
	Bandwidth int64 `protobuf:"varint,2,opt,name=bandwidth" json:"bandwidth,omitempty"`
	// type is the transmission direction, i.e. dowload/upload
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *PayloadBandwidth) Reset()                    { *m = PayloadBandwidth{} }
func (m *PayloadBandwidth) String() string            { return proto.CompactTextString(m) }
func (*PayloadBandwidth) ProtoMessage()               {}
func (*PayloadBandwidth) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PayloadBandwidth) GetTot() int64 {
	if m != nil {
		return m.Tot
	}
	return 0
}

func (m *PayloadBandwidth) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *PayloadBandwidth) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type PayloadMonitor struct {
	// features contains the features that should be inspected.
	Features []int32 `protobuf:"varint,1,rep,packed,name=features" json:"features,omitempty"`
}

func (m *PayloadMonitor) Reset()                    { *m = PayloadMonitor{} }
func (m *PayloadMonitor) String() string            { return proto.CompactTextString(m) }
func (*PayloadMonitor) ProtoMessage()               {}
func (*PayloadMonitor) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *PayloadMonitor) GetFeatures() []int32 {
	if m != nil {
		return m.Features
	}
	return nil
}

type PayloadHello struct {
	// bport is the booster listening port.
	Bport string `protobuf:"bytes,1,opt,name=bport" json:"bport,omitempty"`
	// pport is the proxy listening port.
	Pport string `protobuf:"bytes,2,opt,name=pport" json:"pport,omitempty"`
}

func (m *PayloadHello) Reset()                    { *m = PayloadHello{} }
func (m *PayloadHello) String() string            { return proto.CompactTextString(m) }
func (*PayloadHello) ProtoMessage()               {}
func (*PayloadHello) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *PayloadHello) GetBport() string {
	if m != nil {
		return m.Bport
	}
	return ""
}

func (m *PayloadHello) GetPport() string {
	if m != nil {
		return m.Pport
	}
	return ""
}

type PayloadConnect struct {
	// target of the connect procedure.
	Target string `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
}

func (m *PayloadConnect) Reset()                    { *m = PayloadConnect{} }
func (m *PayloadConnect) String() string            { return proto.CompactTextString(m) }
func (*PayloadConnect) ProtoMessage()               {}
func (*PayloadConnect) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PayloadConnect) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type PayloadDisconnect struct {
	// id is the identifier of the node that should be disconnected
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *PayloadDisconnect) Reset()                    { *m = PayloadDisconnect{} }
func (m *PayloadDisconnect) String() string            { return proto.CompactTextString(m) }
func (*PayloadDisconnect) ProtoMessage()               {}
func (*PayloadDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *PayloadDisconnect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PayloadNode struct {
	// id is the identifier of the node. Usually a sha1 hash.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// baddr is the booster listening address.
	Baddr string `protobuf:"bytes,2,opt,name=baddr" json:"baddr,omitempty"`
	// paddr is the proxy listening address.
	Paddr string `protobuf:"bytes,3,opt,name=paddr" json:"paddr,omitempty"`
	// active tells the connection state of the node.
	Active bool `protobuf:"varint,4,opt,name=active" json:"active,omitempty"`
	// tunnels are the proxy tunnels managed by this node.
	Tunnels []*PayloadNode_Tunnel `protobuf:"bytes,5,rep,name=tunnels" json:"tunnels,omitempty"`
}

func (m *PayloadNode) Reset()                    { *m = PayloadNode{} }
func (m *PayloadNode) String() string            { return proto.CompactTextString(m) }
func (*PayloadNode) ProtoMessage()               {}
func (*PayloadNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *PayloadNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PayloadNode) GetBaddr() string {
	if m != nil {
		return m.Baddr
	}
	return ""
}

func (m *PayloadNode) GetPaddr() string {
	if m != nil {
		return m.Paddr
	}
	return ""
}

func (m *PayloadNode) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *PayloadNode) GetTunnels() []*PayloadNode_Tunnel {
	if m != nil {
		return m.Tunnels
	}
	return nil
}

type PayloadNode_Tunnel struct {
	// id is the tunnel identifier. Usally a sha1 hash.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// target is the remote endpoint address of the tunnel.
	Target string `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	// acks is the number of acknoledgments on this tunnel.
	Acks int32 `protobuf:"varint,3,opt,name=acks" json:"acks,omitempty"`
	// copies are the replications of this tunnel.
	Copies int32 `protobuf:"varint,4,opt,name=copies" json:"copies,omitempty"`
}

func (m *PayloadNode_Tunnel) Reset()                    { *m = PayloadNode_Tunnel{} }
func (m *PayloadNode_Tunnel) String() string            { return proto.CompactTextString(m) }
func (*PayloadNode_Tunnel) ProtoMessage()               {}
func (*PayloadNode_Tunnel) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

func (m *PayloadNode_Tunnel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PayloadNode_Tunnel) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *PayloadNode_Tunnel) GetAcks() int32 {
	if m != nil {
		return m.Acks
	}
	return 0
}

func (m *PayloadNode_Tunnel) GetCopies() int32 {
	if m != nil {
		return m.Copies
	}
	return 0
}

type PayloadHeartbeat struct {
	// id is the identifier of the heartbeat message. Should be unique.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// hops is the number of times that the heartbeat message has been reused.
	Hops int32 `protobuf:"varint,2,opt,name=hops" json:"hops,omitempty"`
	// ttl is the time to leave.
	Ttl *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *PayloadHeartbeat) Reset()                    { *m = PayloadHeartbeat{} }
func (m *PayloadHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*PayloadHeartbeat) ProtoMessage()               {}
func (*PayloadHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *PayloadHeartbeat) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PayloadHeartbeat) GetHops() int32 {
	if m != nil {
		return m.Hops
	}
	return 0
}

func (m *PayloadHeartbeat) GetTtl() *google_protobuf.Timestamp {
	if m != nil {
		return m.Ttl
	}
	return nil
}

type PayloadTunnelEvent struct {
	// target is the remote endpoint address of the tunnel.
	Target string `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	// event is the operation performed on the tunnel.
	Event int32 `protobuf:"varint,2,opt,name=event" json:"event,omitempty"`
}

func (m *PayloadTunnelEvent) Reset()                    { *m = PayloadTunnelEvent{} }
func (m *PayloadTunnelEvent) String() string            { return proto.CompactTextString(m) }
func (*PayloadTunnelEvent) ProtoMessage()               {}
func (*PayloadTunnelEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *PayloadTunnelEvent) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *PayloadTunnelEvent) GetEvent() int32 {
	if m != nil {
		return m.Event
	}
	return 0
}

func init() {
	proto.RegisterType((*PayloadCtrl)(nil), "internal.PayloadCtrl")
	proto.RegisterType((*PayloadBandwidth)(nil), "internal.PayloadBandwidth")
	proto.RegisterType((*PayloadMonitor)(nil), "internal.PayloadMonitor")
	proto.RegisterType((*PayloadHello)(nil), "internal.PayloadHello")
	proto.RegisterType((*PayloadConnect)(nil), "internal.PayloadConnect")
	proto.RegisterType((*PayloadDisconnect)(nil), "internal.PayloadDisconnect")
	proto.RegisterType((*PayloadNode)(nil), "internal.PayloadNode")
	proto.RegisterType((*PayloadNode_Tunnel)(nil), "internal.PayloadNode.Tunnel")
	proto.RegisterType((*PayloadHeartbeat)(nil), "internal.PayloadHeartbeat")
	proto.RegisterType((*PayloadTunnelEvent)(nil), "internal.PayloadTunnelEvent")
}

func init() { proto.RegisterFile("payload.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6a, 0xdc, 0x30,
	0x10, 0xc5, 0x76, 0xbc, 0xcd, 0xce, 0xb6, 0x21, 0x15, 0x25, 0x98, 0x25, 0x50, 0xe3, 0x5e, 0x0c,
	0x0d, 0x0e, 0xa4, 0xd0, 0x43, 0x8f, 0x49, 0x0b, 0xb9, 0xb4, 0x14, 0x11, 0x7a, 0xea, 0x45, 0xb6,
	0x27, 0x1b, 0x51, 0x47, 0x32, 0xf2, 0x6c, 0x4a, 0xbe, 0xbc, 0xd7, 0xa2, 0x91, 0xbc, 0x1b, 0x1a,
	0x7a, 0x9b, 0xf7, 0xf4, 0x34, 0xef, 0x49, 0x33, 0xf0, 0x6a, 0x54, 0x8f, 0x83, 0x55, 0x7d, 0x33,
	0x3a, 0x4b, 0x56, 0x1c, 0x6a, 0x43, 0xe8, 0x8c, 0x1a, 0xd6, 0x6f, 0x37, 0xd6, 0x6e, 0x06, 0x3c,
	0x67, 0xbe, 0xdd, 0xde, 0x9e, 0x93, 0xbe, 0xc7, 0x89, 0xd4, 0xfd, 0x18, 0xa4, 0xd5, 0x7b, 0x58,
	0x7d, 0x0f, 0x77, 0xaf, 0xc8, 0x0d, 0xe2, 0x14, 0x96, 0x76, 0x44, 0xa7, 0x48, 0x5b, 0x53, 0x24,
	0x65, 0x52, 0xe7, 0x72, 0x4f, 0x54, 0x3f, 0xe0, 0x38, 0x8a, 0x2f, 0x95, 0xe9, 0x7f, 0xeb, 0x9e,
	0xee, 0xc4, 0x31, 0x64, 0x64, 0x89, 0xb5, 0x99, 0xf4, 0xa5, 0xef, 0xd1, 0xce, 0xc7, 0x45, 0xca,
	0xfc, 0x9e, 0x10, 0x02, 0x0e, 0xe8, 0x71, 0xc4, 0x22, 0x2b, 0x93, 0x7a, 0x29, 0xb9, 0xae, 0xce,
	0xe0, 0x28, 0xf6, 0xfd, 0x6a, 0x8d, 0x26, 0xeb, 0xc4, 0x1a, 0x0e, 0x6f, 0x51, 0xd1, 0xd6, 0xe1,
	0x54, 0x24, 0x65, 0x56, 0xe7, 0x72, 0x87, 0xab, 0x4f, 0xf0, 0x32, 0xaa, 0xaf, 0x71, 0x18, 0xac,
	0x78, 0x03, 0x79, 0x3b, 0x5a, 0x17, 0x32, 0x2c, 0x65, 0x00, 0x9e, 0x1d, 0x99, 0x4d, 0x03, 0xcb,
	0xa0, 0xaa, 0x77, 0x4e, 0x57, 0xd6, 0x18, 0xec, 0x48, 0x9c, 0xc0, 0x82, 0x94, 0xdb, 0xe0, 0x7c,
	0x3d, 0xa2, 0xea, 0x1d, 0xbc, 0x8e, 0xca, 0xcf, 0x7a, 0xea, 0xa2, 0xf8, 0x08, 0x52, 0xdd, 0x47,
	0x61, 0xaa, 0xfb, 0xea, 0x4f, 0xb2, 0xfb, 0xbe, 0x6f, 0xb6, 0xc7, 0x7f, 0xcf, 0x39, 0x9a, 0xea,
	0x7b, 0x37, 0x87, 0x60, 0xc0, 0xd1, 0x98, 0xcd, 0x62, 0x34, 0x66, 0x4f, 0x60, 0xa1, 0x3a, 0xd2,
	0x0f, 0x58, 0x1c, 0x94, 0x49, 0x7d, 0x28, 0x23, 0x12, 0x1f, 0xe1, 0x05, 0x6d, 0x8d, 0xc1, 0x61,
	0x2a, 0xf2, 0x32, 0xab, 0x57, 0x17, 0xa7, 0xcd, 0x3c, 0xde, 0xe6, 0x89, 0x77, 0x73, 0xc3, 0x22,
	0x39, 0x8b, 0xd7, 0x3f, 0x61, 0x11, 0xa8, 0x67, 0xa9, 0xf6, 0x4f, 0x4e, 0x9f, 0x3e, 0xd9, 0x8f,
	0x46, 0x75, 0xbf, 0x26, 0x8e, 0x95, 0x4b, 0xae, 0xbd, 0xb6, 0xb3, 0xa3, 0xc6, 0x89, 0x53, 0xe5,
	0x32, 0xa2, 0xaa, 0xdf, 0xad, 0xc2, 0x35, 0x2a, 0x47, 0x2d, 0xaa, 0x67, 0xbf, 0xe3, 0xfb, 0xdd,
	0xd9, 0x71, 0x62, 0x97, 0x5c, 0x72, 0x2d, 0xce, 0x20, 0x23, 0x1a, 0xd8, 0x62, 0x75, 0xb1, 0x6e,
	0xc2, 0x7a, 0x36, 0xf3, 0x7a, 0x36, 0x37, 0xf3, 0x7a, 0x4a, 0x2f, 0xab, 0x2e, 0x41, 0x44, 0x97,
	0xf0, 0x94, 0x2f, 0x0f, 0x68, 0xfe, 0x3b, 0x32, 0xff, 0xaf, 0xe8, 0x05, 0xd1, 0x30, 0x80, 0x76,
	0xc1, 0xcd, 0x3f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x18, 0x29, 0xaa, 0x24, 0x03, 0x00,
	0x00,
}

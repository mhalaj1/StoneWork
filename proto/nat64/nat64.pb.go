// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: nat64/nat64.proto

package nat64

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Nat64Interface_Type int32

const (
	// Interface connecting inside/local network with IPv6 endpoints.
	Nat64Interface_IPV6_INSIDE Nat64Interface_Type = 0
	// Interface connecting outside/external network with IPv4 endpoints.
	Nat64Interface_IPV4_OUTSIDE Nat64Interface_Type = 1
)

// Enum value maps for Nat64Interface_Type.
var (
	Nat64Interface_Type_name = map[int32]string{
		0: "IPV6_INSIDE",
		1: "IPV4_OUTSIDE",
	}
	Nat64Interface_Type_value = map[string]int32{
		"IPV6_INSIDE":  0,
		"IPV4_OUTSIDE": 1,
	}
)

func (x Nat64Interface_Type) Enum() *Nat64Interface_Type {
	p := new(Nat64Interface_Type)
	*p = x
	return p
}

func (x Nat64Interface_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Nat64Interface_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_nat64_nat64_proto_enumTypes[0].Descriptor()
}

func (Nat64Interface_Type) Type() protoreflect.EnumType {
	return &file_nat64_nat64_proto_enumTypes[0]
}

func (x Nat64Interface_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Nat64Interface_Type.Descriptor instead.
func (Nat64Interface_Type) EnumDescriptor() ([]byte, []int) {
	return file_nat64_nat64_proto_rawDescGZIP(), []int{1, 0}
}

// Protocol to which the binding applies.
type Nat64StaticBIB_Protocol int32

const (
	Nat64StaticBIB_TCP  Nat64StaticBIB_Protocol = 0
	Nat64StaticBIB_UDP  Nat64StaticBIB_Protocol = 1
	Nat64StaticBIB_ICMP Nat64StaticBIB_Protocol = 2
)

// Enum value maps for Nat64StaticBIB_Protocol.
var (
	Nat64StaticBIB_Protocol_name = map[int32]string{
		0: "TCP",
		1: "UDP",
		2: "ICMP",
	}
	Nat64StaticBIB_Protocol_value = map[string]int32{
		"TCP":  0,
		"UDP":  1,
		"ICMP": 2,
	}
)

func (x Nat64StaticBIB_Protocol) Enum() *Nat64StaticBIB_Protocol {
	p := new(Nat64StaticBIB_Protocol)
	*p = x
	return p
}

func (x Nat64StaticBIB_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Nat64StaticBIB_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_nat64_nat64_proto_enumTypes[1].Descriptor()
}

func (Nat64StaticBIB_Protocol) Type() protoreflect.EnumType {
	return &file_nat64_nat64_proto_enumTypes[1]
}

func (x Nat64StaticBIB_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Nat64StaticBIB_Protocol.Descriptor instead.
func (Nat64StaticBIB_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_nat64_nat64_proto_rawDescGZIP(), []int{3, 0}
}

// IPv4-Embedded IPv6 Address Prefix used for NAT64.
// If no prefix is configured (at all or for a given VRF), then the well-known prefix (64:ff9b::/96) is used.
type Nat64IPv6Prefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// VRF id of tenant.
	// At most one IPv6 prefix can be configured for a given VRF (that's why VRF is part of the key but prefix is not).
	// Non-zero (and not all-ones) VRF has to be explicitly created (see proto/ligato/vpp/l3/vrf.proto).
	VrfId uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// NAT64 prefix in the <IPv6-Address>/<IPv6-Prefix> format.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *Nat64IPv6Prefix) Reset() {
	*x = Nat64IPv6Prefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nat64_nat64_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nat64IPv6Prefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nat64IPv6Prefix) ProtoMessage() {}

func (x *Nat64IPv6Prefix) ProtoReflect() protoreflect.Message {
	mi := &file_nat64_nat64_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nat64IPv6Prefix.ProtoReflect.Descriptor instead.
func (*Nat64IPv6Prefix) Descriptor() ([]byte, []int) {
	return file_nat64_nat64_proto_rawDescGZIP(), []int{0}
}

func (x *Nat64IPv6Prefix) GetVrfId() uint32 {
	if x != nil {
		return x.VrfId
	}
	return 0
}

func (x *Nat64IPv6Prefix) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

// Nat64Interface defines a local network interfaces enabled for NAT64.
type Nat64Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interface name (logical).
	Name string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type Nat64Interface_Type `protobuf:"varint,2,opt,name=type,proto3,enum=nat64.Nat64Interface_Type" json:"type,omitempty"`
}

func (x *Nat64Interface) Reset() {
	*x = Nat64Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nat64_nat64_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nat64Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nat64Interface) ProtoMessage() {}

func (x *Nat64Interface) ProtoReflect() protoreflect.Message {
	mi := &file_nat64_nat64_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nat64Interface.ProtoReflect.Descriptor instead.
func (*Nat64Interface) Descriptor() ([]byte, []int) {
	return file_nat64_nat64_proto_rawDescGZIP(), []int{1}
}

func (x *Nat64Interface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Nat64Interface) GetType() Nat64Interface_Type {
	if x != nil {
		return x.Type
	}
	return Nat64Interface_IPV6_INSIDE
}

// Nat44AddressPool defines an address pool used for NAT64.
type Nat64AddressPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// VRF id of tenant, 0xFFFFFFFF means independent of VRF.
	// Non-zero (and not all-ones) VRF has to be explicitly created (see proto/ligato/vpp/l3/vrf.proto).
	VrfId uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// First IP address of the pool.
	FirstIp string `protobuf:"bytes,2,opt,name=first_ip,json=firstIp,proto3" json:"first_ip,omitempty"`
	// Last IP address of the pool. Should be higher than first_ip or empty.
	LastIp string `protobuf:"bytes,3,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
}

func (x *Nat64AddressPool) Reset() {
	*x = Nat64AddressPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nat64_nat64_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nat64AddressPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nat64AddressPool) ProtoMessage() {}

func (x *Nat64AddressPool) ProtoReflect() protoreflect.Message {
	mi := &file_nat64_nat64_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nat64AddressPool.ProtoReflect.Descriptor instead.
func (*Nat64AddressPool) Descriptor() ([]byte, []int) {
	return file_nat64_nat64_proto_rawDescGZIP(), []int{2}
}

func (x *Nat64AddressPool) GetVrfId() uint32 {
	if x != nil {
		return x.VrfId
	}
	return 0
}

func (x *Nat64AddressPool) GetFirstIp() string {
	if x != nil {
		return x.FirstIp
	}
	return ""
}

func (x *Nat64AddressPool) GetLastIp() string {
	if x != nil {
		return x.LastIp
	}
	return ""
}

// Static NAT64 binding allowing IPv4 host from the outside to access IPv6 host from the inside.
type Nat64StaticBIB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// VRF (table) ID. Non-zero VRF has to be explicitly created (see proto/ligato/vpp/l3/vrf.proto).
	VrfId uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// IPv6 host from the inside/local network.
	InsideIpv6Address string `protobuf:"bytes,2,opt,name=inside_ipv6_address,json=insideIpv6Address,proto3" json:"inside_ipv6_address,omitempty"`
	// Inside port number (of the IPv6 host).
	InsidePort uint32 `protobuf:"varint,3,opt,name=inside_port,json=insidePort,proto3" json:"inside_port,omitempty"`
	// IPv4 host from the outside/external network.
	OutsideIpv4Address string `protobuf:"bytes,4,opt,name=outside_ipv4_address,json=outsideIpv4Address,proto3" json:"outside_ipv4_address,omitempty"`
	// Outside port number (of the IPv4 host).
	OutsidePort uint32                  `protobuf:"varint,5,opt,name=outside_port,json=outsidePort,proto3" json:"outside_port,omitempty"`
	Protocol    Nat64StaticBIB_Protocol `protobuf:"varint,6,opt,name=protocol,proto3,enum=nat64.Nat64StaticBIB_Protocol" json:"protocol,omitempty"`
}

func (x *Nat64StaticBIB) Reset() {
	*x = Nat64StaticBIB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nat64_nat64_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nat64StaticBIB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nat64StaticBIB) ProtoMessage() {}

func (x *Nat64StaticBIB) ProtoReflect() protoreflect.Message {
	mi := &file_nat64_nat64_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nat64StaticBIB.ProtoReflect.Descriptor instead.
func (*Nat64StaticBIB) Descriptor() ([]byte, []int) {
	return file_nat64_nat64_proto_rawDescGZIP(), []int{3}
}

func (x *Nat64StaticBIB) GetVrfId() uint32 {
	if x != nil {
		return x.VrfId
	}
	return 0
}

func (x *Nat64StaticBIB) GetInsideIpv6Address() string {
	if x != nil {
		return x.InsideIpv6Address
	}
	return ""
}

func (x *Nat64StaticBIB) GetInsidePort() uint32 {
	if x != nil {
		return x.InsidePort
	}
	return 0
}

func (x *Nat64StaticBIB) GetOutsideIpv4Address() string {
	if x != nil {
		return x.OutsideIpv4Address
	}
	return ""
}

func (x *Nat64StaticBIB) GetOutsidePort() uint32 {
	if x != nil {
		return x.OutsidePort
	}
	return 0
}

func (x *Nat64StaticBIB) GetProtocol() Nat64StaticBIB_Protocol {
	if x != nil {
		return x.Protocol
	}
	return Nat64StaticBIB_TCP
}

var File_nat64_nat64_proto protoreflect.FileDescriptor

var file_nat64_nat64_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x61, 0x74, 0x36, 0x34, 0x2f, 0x6e, 0x61, 0x74, 0x36, 0x34, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x61, 0x74, 0x36, 0x34, 0x22, 0x40, 0x0a, 0x0f, 0x4e, 0x61,
	0x74, 0x36, 0x34, 0x49, 0x50, 0x76, 0x36, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x15, 0x0a,
	0x06, 0x76, 0x72, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x72, 0x66, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x7f, 0x0a, 0x0e,
	0x4e, 0x61, 0x74, 0x36, 0x34, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x6e, 0x61, 0x74, 0x36, 0x34, 0x2e, 0x4e, 0x61, 0x74, 0x36, 0x34, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x50,
	0x56, 0x36, 0x5f, 0x49, 0x4e, 0x53, 0x49, 0x44, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49,
	0x50, 0x56, 0x34, 0x5f, 0x4f, 0x55, 0x54, 0x53, 0x49, 0x44, 0x45, 0x10, 0x01, 0x22, 0x5d, 0x0a,
	0x10, 0x4e, 0x61, 0x74, 0x36, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x72, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x72, 0x66, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x49, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x70, 0x22, 0xb1, 0x02, 0x0a,
	0x0e, 0x4e, 0x61, 0x74, 0x36, 0x34, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x42, 0x49, 0x42, 0x12,
	0x15, 0x0a, 0x06, 0x76, 0x72, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x76, 0x72, 0x66, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65,
	0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x49, 0x70, 0x76, 0x36, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x73,
	0x69, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6f, 0x75, 0x74, 0x73, 0x69,
	0x64, 0x65, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x70,
	0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x75, 0x74,
	0x73, 0x69, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x3a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x6e, 0x61, 0x74, 0x36, 0x34, 0x2e, 0x4e, 0x61, 0x74, 0x36, 0x34, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x42, 0x49, 0x42, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x26, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x43, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x55, 0x44, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x43, 0x4d, 0x50, 0x10, 0x02,
	0x42, 0x2b, 0x5a, 0x29, 0x70, 0x61, 0x6e, 0x74, 0x68, 0x65, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6e, 0x61, 0x74, 0x36, 0x34, 0x3b, 0x6e, 0x61, 0x74, 0x36, 0x34, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nat64_nat64_proto_rawDescOnce sync.Once
	file_nat64_nat64_proto_rawDescData = file_nat64_nat64_proto_rawDesc
)

func file_nat64_nat64_proto_rawDescGZIP() []byte {
	file_nat64_nat64_proto_rawDescOnce.Do(func() {
		file_nat64_nat64_proto_rawDescData = protoimpl.X.CompressGZIP(file_nat64_nat64_proto_rawDescData)
	})
	return file_nat64_nat64_proto_rawDescData
}

var file_nat64_nat64_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nat64_nat64_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nat64_nat64_proto_goTypes = []interface{}{
	(Nat64Interface_Type)(0),     // 0: nat64.Nat64Interface.Type
	(Nat64StaticBIB_Protocol)(0), // 1: nat64.Nat64StaticBIB.Protocol
	(*Nat64IPv6Prefix)(nil),      // 2: nat64.Nat64IPv6Prefix
	(*Nat64Interface)(nil),       // 3: nat64.Nat64Interface
	(*Nat64AddressPool)(nil),     // 4: nat64.Nat64AddressPool
	(*Nat64StaticBIB)(nil),       // 5: nat64.Nat64StaticBIB
}
var file_nat64_nat64_proto_depIdxs = []int32{
	0, // 0: nat64.Nat64Interface.type:type_name -> nat64.Nat64Interface.Type
	1, // 1: nat64.Nat64StaticBIB.protocol:type_name -> nat64.Nat64StaticBIB.Protocol
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nat64_nat64_proto_init() }
func file_nat64_nat64_proto_init() {
	if File_nat64_nat64_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nat64_nat64_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nat64IPv6Prefix); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nat64_nat64_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nat64Interface); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nat64_nat64_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nat64AddressPool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nat64_nat64_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nat64StaticBIB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nat64_nat64_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nat64_nat64_proto_goTypes,
		DependencyIndexes: file_nat64_nat64_proto_depIdxs,
		EnumInfos:         file_nat64_nat64_proto_enumTypes,
		MessageInfos:      file_nat64_nat64_proto_msgTypes,
	}.Build()
	File_nat64_nat64_proto = out.File
	file_nat64_nat64_proto_rawDesc = nil
	file_nat64_nat64_proto_goTypes = nil
	file_nat64_nat64_proto_depIdxs = nil
}

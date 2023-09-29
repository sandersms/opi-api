// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.
// {C} Copyright 2022 Pensando Systems Inc. All rights reserved
// Copyright (C) 2023 Intel Corporation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: opicommon.proto

package _go

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AES encryption types
type EncryptionType int32

const (
	// Encryption type is not specified
	EncryptionType_ENCRYPTION_TYPE_UNSPECIFIED EncryptionType = 0
	// AES CBC 128 encryption type
	EncryptionType_ENCRYPTION_TYPE_AES_CBC_128 EncryptionType = 1
	// AES CBC 192 encryption type
	EncryptionType_ENCRYPTION_TYPE_AES_CBC_192 EncryptionType = 2
	// AES CBC 256 encryption type
	EncryptionType_ENCRYPTION_TYPE_AES_CBC_256 EncryptionType = 3
	// AES XTS 128 encryption type
	EncryptionType_ENCRYPTION_TYPE_AES_XTS_128 EncryptionType = 4
	// AES XTS 192 encryption type
	EncryptionType_ENCRYPTION_TYPE_AES_XTS_192 EncryptionType = 5
	// AES XTS 256 encryption type
	EncryptionType_ENCRYPTION_TYPE_AES_XTS_256 EncryptionType = 6
)

// Enum value maps for EncryptionType.
var (
	EncryptionType_name = map[int32]string{
		0: "ENCRYPTION_TYPE_UNSPECIFIED",
		1: "ENCRYPTION_TYPE_AES_CBC_128",
		2: "ENCRYPTION_TYPE_AES_CBC_192",
		3: "ENCRYPTION_TYPE_AES_CBC_256",
		4: "ENCRYPTION_TYPE_AES_XTS_128",
		5: "ENCRYPTION_TYPE_AES_XTS_192",
		6: "ENCRYPTION_TYPE_AES_XTS_256",
	}
	EncryptionType_value = map[string]int32{
		"ENCRYPTION_TYPE_UNSPECIFIED": 0,
		"ENCRYPTION_TYPE_AES_CBC_128": 1,
		"ENCRYPTION_TYPE_AES_CBC_192": 2,
		"ENCRYPTION_TYPE_AES_CBC_256": 3,
		"ENCRYPTION_TYPE_AES_XTS_128": 4,
		"ENCRYPTION_TYPE_AES_XTS_192": 5,
		"ENCRYPTION_TYPE_AES_XTS_256": 6,
	}
)

func (x EncryptionType) Enum() *EncryptionType {
	p := new(EncryptionType)
	*p = x
	return p
}

func (x EncryptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_opicommon_proto_enumTypes[0].Descriptor()
}

func (EncryptionType) Type() protoreflect.EnumType {
	return &file_opicommon_proto_enumTypes[0]
}

func (x EncryptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptionType.Descriptor instead.
func (EncryptionType) EnumDescriptor() ([]byte, []int) {
	return file_opicommon_proto_rawDescGZIP(), []int{0}
}

// Transport type value options
type NvmeTransportType int32

const (
	// Transport type is not specified
	NvmeTransportType_NVME_TRANSPORT_TYPE_UNSPECIFIED NvmeTransportType = 0
	// Fibre channel transport type
	NvmeTransportType_NVME_TRANSPORT_FC NvmeTransportType = 1
	// Pcie transport type
	NvmeTransportType_NVME_TRANSPORT_PCIE NvmeTransportType = 2
	// RDMA transport type
	NvmeTransportType_NVME_TRANSPORT_RDMA NvmeTransportType = 3
	// TCP transport type
	NvmeTransportType_NVME_TRANSPORT_TCP NvmeTransportType = 4
	// Custom transport type
	NvmeTransportType_NVME_TRANSPORT_CUSTOM NvmeTransportType = 5
)

// Enum value maps for NvmeTransportType.
var (
	NvmeTransportType_name = map[int32]string{
		0: "NVME_TRANSPORT_TYPE_UNSPECIFIED",
		1: "NVME_TRANSPORT_FC",
		2: "NVME_TRANSPORT_PCIE",
		3: "NVME_TRANSPORT_RDMA",
		4: "NVME_TRANSPORT_TCP",
		5: "NVME_TRANSPORT_CUSTOM",
	}
	NvmeTransportType_value = map[string]int32{
		"NVME_TRANSPORT_TYPE_UNSPECIFIED": 0,
		"NVME_TRANSPORT_FC":               1,
		"NVME_TRANSPORT_PCIE":             2,
		"NVME_TRANSPORT_RDMA":             3,
		"NVME_TRANSPORT_TCP":              4,
		"NVME_TRANSPORT_CUSTOM":           5,
	}
)

func (x NvmeTransportType) Enum() *NvmeTransportType {
	p := new(NvmeTransportType)
	*p = x
	return p
}

func (x NvmeTransportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NvmeTransportType) Descriptor() protoreflect.EnumDescriptor {
	return file_opicommon_proto_enumTypes[1].Descriptor()
}

func (NvmeTransportType) Type() protoreflect.EnumType {
	return &file_opicommon_proto_enumTypes[1]
}

func (x NvmeTransportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NvmeTransportType.Descriptor instead.
func (NvmeTransportType) EnumDescriptor() ([]byte, []int) {
	return file_opicommon_proto_rawDescGZIP(), []int{1}
}

// Address family value options
type NvmeAddressFamily int32

const (
	// Address family is not specified
	NvmeAddressFamily_NVME_ADDRESS_FAMILY_UNSPECIFIED NvmeAddressFamily = 0
	// IPv4 address family
	NvmeAddressFamily_NVME_ADRFAM_IPV4 NvmeAddressFamily = 1
	// IPv6 address family
	NvmeAddressFamily_NVME_ADRFAM_IPV6 NvmeAddressFamily = 2
	// InfiniBand address family
	NvmeAddressFamily_NVME_ADRFAM_IB NvmeAddressFamily = 3
	// Fibre channel address family
	NvmeAddressFamily_NVME_ADRFAM_FC NvmeAddressFamily = 4
	// Intra host address family
	NvmeAddressFamily_NVME_ADRFAM_INTRA_HOST NvmeAddressFamily = 5
)

// Enum value maps for NvmeAddressFamily.
var (
	NvmeAddressFamily_name = map[int32]string{
		0: "NVME_ADDRESS_FAMILY_UNSPECIFIED",
		1: "NVME_ADRFAM_IPV4",
		2: "NVME_ADRFAM_IPV6",
		3: "NVME_ADRFAM_IB",
		4: "NVME_ADRFAM_FC",
		5: "NVME_ADRFAM_INTRA_HOST",
	}
	NvmeAddressFamily_value = map[string]int32{
		"NVME_ADDRESS_FAMILY_UNSPECIFIED": 0,
		"NVME_ADRFAM_IPV4":                1,
		"NVME_ADRFAM_IPV6":                2,
		"NVME_ADRFAM_IB":                  3,
		"NVME_ADRFAM_FC":                  4,
		"NVME_ADRFAM_INTRA_HOST":          5,
	}
)

func (x NvmeAddressFamily) Enum() *NvmeAddressFamily {
	p := new(NvmeAddressFamily)
	*p = x
	return p
}

func (x NvmeAddressFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NvmeAddressFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_opicommon_proto_enumTypes[2].Descriptor()
}

func (NvmeAddressFamily) Type() protoreflect.EnumType {
	return &file_opicommon_proto_enumTypes[2]
}

func (x NvmeAddressFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NvmeAddressFamily.Descriptor instead.
func (NvmeAddressFamily) EnumDescriptor() ([]byte, []int) {
	return file_opicommon_proto_rawDescGZIP(), []int{2}
}

// Many front-ends expose PCI devices to the host. This represents that endpoint.
// This device will ultimately be surfaced by the operating system at some
// Bus:Device:Function, but note that the values set here are not necessarily
// the same values that the operating system will choose. Instead, these are
// xPU-internal values.
// While the term "device" is often used for the entity attached to a PCI slot,
// we'll use the term "port" which also commonly used with PCI switches and avoids
// confusion with storage "devices".
type PciEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The "port" or "device". In other words, the connector/cable that's
	// plugged into a particular host. This number may end up matching
	// the host-assigned "device" value in the bus:device:function identifier,
	// but it does not strictly have to and that should not be relied upon.
	PortId *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// Physical function index. This may end up matching the host-assigned
	// "function" value in the bus:device:function identifier, but it does not
	// strictly have to and that should not be relied upon.
	PhysicalFunction *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=physical_function,json=physicalFunction,proto3" json:"physical_function,omitempty"`
	// Virtual function index. 1-based index.
	// The value 0 is reserved to represent the PCI physical "device".
	// This may end up matching the host-assigned "function" value in the
	// bus:device:function identifier, but it does not strictly have to and
	// that should not be relied upon.
	VirtualFunction *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=virtual_function,json=virtualFunction,proto3" json:"virtual_function,omitempty"`
}

func (x *PciEndpoint) Reset() {
	*x = PciEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opicommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PciEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PciEndpoint) ProtoMessage() {}

func (x *PciEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_opicommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PciEndpoint.ProtoReflect.Descriptor instead.
func (*PciEndpoint) Descriptor() ([]byte, []int) {
	return file_opicommon_proto_rawDescGZIP(), []int{0}
}

func (x *PciEndpoint) GetPortId() *wrapperspb.Int32Value {
	if x != nil {
		return x.PortId
	}
	return nil
}

func (x *PciEndpoint) GetPhysicalFunction() *wrapperspb.Int32Value {
	if x != nil {
		return x.PhysicalFunction
	}
	return nil
}

func (x *PciEndpoint) GetVirtualFunction() *wrapperspb.Int32Value {
	if x != nil {
		return x.VirtualFunction
	}
	return nil
}

// Represents Fabrics Endpoint
type FabricsEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ip address for TCP and RDMA
	Traddr string `protobuf:"bytes,1,opt,name=traddr,proto3" json:"traddr,omitempty"`
	// port for TCP and RDMA
	Trsvcid string `protobuf:"bytes,2,opt,name=trsvcid,proto3" json:"trsvcid,omitempty"`
	// address family
	Adrfam NvmeAddressFamily `protobuf:"varint,3,opt,name=adrfam,proto3,enum=opi_api.storage.v1.NvmeAddressFamily" json:"adrfam,omitempty"`
}

func (x *FabricsEndpoint) Reset() {
	*x = FabricsEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opicommon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FabricsEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FabricsEndpoint) ProtoMessage() {}

func (x *FabricsEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_opicommon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FabricsEndpoint.ProtoReflect.Descriptor instead.
func (*FabricsEndpoint) Descriptor() ([]byte, []int) {
	return file_opicommon_proto_rawDescGZIP(), []int{1}
}

func (x *FabricsEndpoint) GetTraddr() string {
	if x != nil {
		return x.Traddr
	}
	return ""
}

func (x *FabricsEndpoint) GetTrsvcid() string {
	if x != nil {
		return x.Trsvcid
	}
	return ""
}

func (x *FabricsEndpoint) GetAdrfam() NvmeAddressFamily {
	if x != nil {
		return x.Adrfam
	}
	return NvmeAddressFamily_NVME_ADDRESS_FAMILY_UNSPECIFIED
}

// Represents Volume statistics
type VolumeStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count of read bytes
	ReadBytesCount int32 `protobuf:"varint,1,opt,name=read_bytes_count,json=readBytesCount,proto3" json:"read_bytes_count,omitempty"`
	// Count of read operations
	ReadOpsCount int32 `protobuf:"varint,2,opt,name=read_ops_count,json=readOpsCount,proto3" json:"read_ops_count,omitempty"`
	// Count of written bytes
	WriteBytesCount int32 `protobuf:"varint,3,opt,name=write_bytes_count,json=writeBytesCount,proto3" json:"write_bytes_count,omitempty"`
	// Count of write opeations
	WriteOpsCount int32 `protobuf:"varint,4,opt,name=write_ops_count,json=writeOpsCount,proto3" json:"write_ops_count,omitempty"`
	// Count of unmapped bytes
	UnmapBytesCount int32 `protobuf:"varint,5,opt,name=unmap_bytes_count,json=unmapBytesCount,proto3" json:"unmap_bytes_count,omitempty"`
	// Count of unmap operations
	UnmapOpsCount int32 `protobuf:"varint,6,opt,name=unmap_ops_count,json=unmapOpsCount,proto3" json:"unmap_ops_count,omitempty"`
	// Read latency ticks
	ReadLatencyTicks int32 `protobuf:"varint,7,opt,name=read_latency_ticks,json=readLatencyTicks,proto3" json:"read_latency_ticks,omitempty"`
	// Write latency ticks
	WriteLatencyTicks int32 `protobuf:"varint,8,opt,name=write_latency_ticks,json=writeLatencyTicks,proto3" json:"write_latency_ticks,omitempty"`
	// Unmap latency ticks
	UnmapLatencyTicks int32 `protobuf:"varint,9,opt,name=unmap_latency_ticks,json=unmapLatencyTicks,proto3" json:"unmap_latency_ticks,omitempty"`
}

func (x *VolumeStats) Reset() {
	*x = VolumeStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opicommon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeStats) ProtoMessage() {}

func (x *VolumeStats) ProtoReflect() protoreflect.Message {
	mi := &file_opicommon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeStats.ProtoReflect.Descriptor instead.
func (*VolumeStats) Descriptor() ([]byte, []int) {
	return file_opicommon_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeStats) GetReadBytesCount() int32 {
	if x != nil {
		return x.ReadBytesCount
	}
	return 0
}

func (x *VolumeStats) GetReadOpsCount() int32 {
	if x != nil {
		return x.ReadOpsCount
	}
	return 0
}

func (x *VolumeStats) GetWriteBytesCount() int32 {
	if x != nil {
		return x.WriteBytesCount
	}
	return 0
}

func (x *VolumeStats) GetWriteOpsCount() int32 {
	if x != nil {
		return x.WriteOpsCount
	}
	return 0
}

func (x *VolumeStats) GetUnmapBytesCount() int32 {
	if x != nil {
		return x.UnmapBytesCount
	}
	return 0
}

func (x *VolumeStats) GetUnmapOpsCount() int32 {
	if x != nil {
		return x.UnmapOpsCount
	}
	return 0
}

func (x *VolumeStats) GetReadLatencyTicks() int32 {
	if x != nil {
		return x.ReadLatencyTicks
	}
	return 0
}

func (x *VolumeStats) GetWriteLatencyTicks() int32 {
	if x != nil {
		return x.WriteLatencyTicks
	}
	return 0
}

func (x *VolumeStats) GetUnmapLatencyTicks() int32 {
	if x != nil {
		return x.UnmapLatencyTicks
	}
	return 0
}

// QoS limits applied to volumes/devices
type QosLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Read kIOPS
	RdIopsKiops int64 `protobuf:"varint,1,opt,name=rd_iops_kiops,json=rdIopsKiops,proto3" json:"rd_iops_kiops,omitempty"`
	// Write kIOPS
	WrIopsKiops int64 `protobuf:"varint,2,opt,name=wr_iops_kiops,json=wrIopsKiops,proto3" json:"wr_iops_kiops,omitempty"`
	// Read/write kIOPS
	RwIopsKiops int64 `protobuf:"varint,3,opt,name=rw_iops_kiops,json=rwIopsKiops,proto3" json:"rw_iops_kiops,omitempty"`
	// Read bandwidth (MB/s)
	RdBandwidthMbs int64 `protobuf:"varint,4,opt,name=rd_bandwidth_mbs,json=rdBandwidthMbs,proto3" json:"rd_bandwidth_mbs,omitempty"`
	// Write bandwidth (MB/s)
	WrBandwidthMbs int64 `protobuf:"varint,5,opt,name=wr_bandwidth_mbs,json=wrBandwidthMbs,proto3" json:"wr_bandwidth_mbs,omitempty"`
	// Read/write bandwidth (MB/s)
	RwBandwidthMbs int64 `protobuf:"varint,6,opt,name=rw_bandwidth_mbs,json=rwBandwidthMbs,proto3" json:"rw_bandwidth_mbs,omitempty"`
}

func (x *QosLimit) Reset() {
	*x = QosLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opicommon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QosLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QosLimit) ProtoMessage() {}

func (x *QosLimit) ProtoReflect() protoreflect.Message {
	mi := &file_opicommon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QosLimit.ProtoReflect.Descriptor instead.
func (*QosLimit) Descriptor() ([]byte, []int) {
	return file_opicommon_proto_rawDescGZIP(), []int{3}
}

func (x *QosLimit) GetRdIopsKiops() int64 {
	if x != nil {
		return x.RdIopsKiops
	}
	return 0
}

func (x *QosLimit) GetWrIopsKiops() int64 {
	if x != nil {
		return x.WrIopsKiops
	}
	return 0
}

func (x *QosLimit) GetRwIopsKiops() int64 {
	if x != nil {
		return x.RwIopsKiops
	}
	return 0
}

func (x *QosLimit) GetRdBandwidthMbs() int64 {
	if x != nil {
		return x.RdBandwidthMbs
	}
	return 0
}

func (x *QosLimit) GetWrBandwidthMbs() int64 {
	if x != nil {
		return x.WrBandwidthMbs
	}
	return 0
}

func (x *QosLimit) GetRwBandwidthMbs() int64 {
	if x != nil {
		return x.RwBandwidthMbs
	}
	return 0
}

var File_opicommon_proto protoreflect.FileDescriptor

var file_opicommon_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x70, 0x69, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x6f, 0x70, 0x69, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0b, 0x50, 0x63, 0x69, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x4d, 0x0a, 0x11, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10,
	0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4b, 0x0a, 0x10, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01,
	0x0a, 0x0f, 0x46, 0x61, 0x62, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x74, 0x72, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1d,
	0x0a, 0x07, 0x74, 0x72, 0x73, 0x76, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x74, 0x72, 0x73, 0x76, 0x63, 0x69, 0x64, 0x12, 0x42, 0x0a,
	0x06, 0x61, 0x64, 0x72, 0x66, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e,
	0x6f, 0x70, 0x69, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x76, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x61, 0x64, 0x72, 0x66, 0x61,
	0x6d, 0x22, 0xc0, 0x03, 0x0a, 0x0b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x29, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x72,
	0x65, 0x61, 0x64, 0x4f, 0x70, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x11, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0f,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x4f, 0x70, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x11, 0x75, 0x6e, 0x6d,
	0x61, 0x70, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x75, 0x6e, 0x6d, 0x61, 0x70,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0f, 0x75, 0x6e,
	0x6d, 0x61, 0x70, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x75, 0x6e, 0x6d, 0x61, 0x70, 0x4f,
	0x70, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x12, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x72, 0x65, 0x61, 0x64, 0x4c, 0x61,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x33, 0x0a, 0x13, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x12,
	0x33, 0x0a, 0x13, 0x75, 0x6e, 0x6d, 0x61, 0x70, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x11, 0x75, 0x6e, 0x6d, 0x61, 0x70, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x54,
	0x69, 0x63, 0x6b, 0x73, 0x22, 0x92, 0x02, 0x0a, 0x08, 0x51, 0x6f, 0x73, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x27, 0x0a, 0x0d, 0x72, 0x64, 0x5f, 0x69, 0x6f, 0x70, 0x73, 0x5f, 0x6b, 0x69, 0x6f,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x72,
	0x64, 0x49, 0x6f, 0x70, 0x73, 0x4b, 0x69, 0x6f, 0x70, 0x73, 0x12, 0x27, 0x0a, 0x0d, 0x77, 0x72,
	0x5f, 0x69, 0x6f, 0x70, 0x73, 0x5f, 0x6b, 0x69, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x77, 0x72, 0x49, 0x6f, 0x70, 0x73, 0x4b, 0x69,
	0x6f, 0x70, 0x73, 0x12, 0x27, 0x0a, 0x0d, 0x72, 0x77, 0x5f, 0x69, 0x6f, 0x70, 0x73, 0x5f, 0x6b,
	0x69, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0b, 0x72, 0x77, 0x49, 0x6f, 0x70, 0x73, 0x4b, 0x69, 0x6f, 0x70, 0x73, 0x12, 0x2d, 0x0a, 0x10,
	0x72, 0x64, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6d, 0x62, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x72, 0x64, 0x42,
	0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4d, 0x62, 0x73, 0x12, 0x2d, 0x0a, 0x10, 0x77,
	0x72, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6d, 0x62, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x77, 0x72, 0x42, 0x61,
	0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4d, 0x62, 0x73, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x77,
	0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6d, 0x62, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x72, 0x77, 0x42, 0x61, 0x6e,
	0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4d, 0x62, 0x73, 0x2a, 0xf7, 0x01, 0x0a, 0x0e, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b,
	0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a,
	0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x45, 0x53, 0x5f, 0x43, 0x42, 0x43, 0x5f, 0x31, 0x32, 0x38, 0x10, 0x01, 0x12, 0x1f,
	0x0a, 0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x43, 0x42, 0x43, 0x5f, 0x31, 0x39, 0x32, 0x10, 0x02, 0x12,
	0x1f, 0x0a, 0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x43, 0x42, 0x43, 0x5f, 0x32, 0x35, 0x36, 0x10, 0x03,
	0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x58, 0x54, 0x53, 0x5f, 0x31, 0x32, 0x38, 0x10,
	0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x58, 0x54, 0x53, 0x5f, 0x31, 0x39, 0x32,
	0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x58, 0x54, 0x53, 0x5f, 0x32, 0x35,
	0x36, 0x10, 0x06, 0x2a, 0xb4, 0x01, 0x0a, 0x11, 0x4e, 0x76, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x4e, 0x56, 0x4d,
	0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x46, 0x43, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x50, 0x43, 0x49, 0x45, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x52, 0x44, 0x4d, 0x41, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x56, 0x4d, 0x45, 0x5f,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x43, 0x50, 0x10, 0x04, 0x12,
	0x19, 0x0a, 0x15, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x05, 0x2a, 0xa8, 0x01, 0x0a, 0x11, 0x4e,
	0x76, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x12, 0x23, 0x0a, 0x1f, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53,
	0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x41, 0x44,
	0x52, 0x46, 0x41, 0x4d, 0x5f, 0x49, 0x50, 0x56, 0x34, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4e,
	0x56, 0x4d, 0x45, 0x5f, 0x41, 0x44, 0x52, 0x46, 0x41, 0x4d, 0x5f, 0x49, 0x50, 0x56, 0x36, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x41, 0x44, 0x52, 0x46, 0x41, 0x4d,
	0x5f, 0x49, 0x42, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x56, 0x4d, 0x45, 0x5f, 0x41, 0x44,
	0x52, 0x46, 0x41, 0x4d, 0x5f, 0x46, 0x43, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x56, 0x4d,
	0x45, 0x5f, 0x41, 0x44, 0x52, 0x46, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x54, 0x52, 0x41, 0x5f, 0x48,
	0x4f, 0x53, 0x54, 0x10, 0x05, 0x42, 0x5d, 0x0a, 0x12, 0x6f, 0x70, 0x69, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4f, 0x70, 0x69,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x69, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6f, 0x70, 0x69, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opicommon_proto_rawDescOnce sync.Once
	file_opicommon_proto_rawDescData = file_opicommon_proto_rawDesc
)

func file_opicommon_proto_rawDescGZIP() []byte {
	file_opicommon_proto_rawDescOnce.Do(func() {
		file_opicommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_opicommon_proto_rawDescData)
	})
	return file_opicommon_proto_rawDescData
}

var file_opicommon_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_opicommon_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_opicommon_proto_goTypes = []interface{}{
	(EncryptionType)(0),           // 0: opi_api.storage.v1.EncryptionType
	(NvmeTransportType)(0),        // 1: opi_api.storage.v1.NvmeTransportType
	(NvmeAddressFamily)(0),        // 2: opi_api.storage.v1.NvmeAddressFamily
	(*PciEndpoint)(nil),           // 3: opi_api.storage.v1.PciEndpoint
	(*FabricsEndpoint)(nil),       // 4: opi_api.storage.v1.FabricsEndpoint
	(*VolumeStats)(nil),           // 5: opi_api.storage.v1.VolumeStats
	(*QosLimit)(nil),              // 6: opi_api.storage.v1.QosLimit
	(*wrapperspb.Int32Value)(nil), // 7: google.protobuf.Int32Value
}
var file_opicommon_proto_depIdxs = []int32{
	7, // 0: opi_api.storage.v1.PciEndpoint.port_id:type_name -> google.protobuf.Int32Value
	7, // 1: opi_api.storage.v1.PciEndpoint.physical_function:type_name -> google.protobuf.Int32Value
	7, // 2: opi_api.storage.v1.PciEndpoint.virtual_function:type_name -> google.protobuf.Int32Value
	2, // 3: opi_api.storage.v1.FabricsEndpoint.adrfam:type_name -> opi_api.storage.v1.NvmeAddressFamily
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_opicommon_proto_init() }
func file_opicommon_proto_init() {
	if File_opicommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opicommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PciEndpoint); i {
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
		file_opicommon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FabricsEndpoint); i {
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
		file_opicommon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeStats); i {
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
		file_opicommon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QosLimit); i {
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
			RawDescriptor: file_opicommon_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opicommon_proto_goTypes,
		DependencyIndexes: file_opicommon_proto_depIdxs,
		EnumInfos:         file_opicommon_proto_enumTypes,
		MessageInfos:      file_opicommon_proto_msgTypes,
	}.Build()
	File_opicommon_proto = out.File
	file_opicommon_proto_rawDesc = nil
	file_opicommon_proto_goTypes = nil
	file_opicommon_proto_depIdxs = nil
}

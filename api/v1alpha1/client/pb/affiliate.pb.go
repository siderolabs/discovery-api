// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: v1alpha1/client/pb/affiliate.proto

package pb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Affiliate protobuf spec, as seen internally in the client code.
//
// Discovery service recieves Affiliate marshaled to protobuf and encrypted.
type Affiliate struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	NodeId string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// IPs are encoded using binary marshaling.
	Addresses       [][]byte      `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Hostname        string        `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Nodename        string        `protobuf:"bytes,4,opt,name=nodename,proto3" json:"nodename,omitempty"`
	MachineType     string        `protobuf:"bytes,5,opt,name=machine_type,json=machineType,proto3" json:"machine_type,omitempty"`
	OperatingSystem string        `protobuf:"bytes,6,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
	Kubespan        *KubeSpan     `protobuf:"bytes,7,opt,name=kubespan,proto3,oneof" json:"kubespan,omitempty"`
	ControlPlane    *ControlPlane `protobuf:"bytes,8,opt,name=control_plane,json=controlPlane,proto3,oneof" json:"control_plane,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Affiliate) Reset() {
	*x = Affiliate{}
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Affiliate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Affiliate) ProtoMessage() {}

func (x *Affiliate) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Affiliate.ProtoReflect.Descriptor instead.
func (*Affiliate) Descriptor() ([]byte, []int) {
	return file_v1alpha1_client_pb_affiliate_proto_rawDescGZIP(), []int{0}
}

func (x *Affiliate) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Affiliate) GetAddresses() [][]byte {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *Affiliate) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Affiliate) GetNodename() string {
	if x != nil {
		return x.Nodename
	}
	return ""
}

func (x *Affiliate) GetMachineType() string {
	if x != nil {
		return x.MachineType
	}
	return ""
}

func (x *Affiliate) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *Affiliate) GetKubespan() *KubeSpan {
	if x != nil {
		return x.Kubespan
	}
	return nil
}

func (x *Affiliate) GetControlPlane() *ControlPlane {
	if x != nil {
		return x.ControlPlane
	}
	return nil
}

// KubeSpan optional configuration.
type KubeSpan struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	PublicKey           string                 `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Address             []byte                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	AdditionalAddresses []*IPPrefix            `protobuf:"bytes,3,rep,name=additional_addresses,json=additionalAddresses,proto3" json:"additional_addresses,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *KubeSpan) Reset() {
	*x = KubeSpan{}
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubeSpan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeSpan) ProtoMessage() {}

func (x *KubeSpan) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeSpan.ProtoReflect.Descriptor instead.
func (*KubeSpan) Descriptor() ([]byte, []int) {
	return file_v1alpha1_client_pb_affiliate_proto_rawDescGZIP(), []int{1}
}

func (x *KubeSpan) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *KubeSpan) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *KubeSpan) GetAdditionalAddresses() []*IPPrefix {
	if x != nil {
		return x.AdditionalAddresses
	}
	return nil
}

// IPPrefix contains CIDR.
type IPPrefix struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// IPs are encoded using binary marshaling.
	Ip            []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Bits          uint32 `protobuf:"varint,2,opt,name=bits,proto3" json:"bits,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPPrefix) Reset() {
	*x = IPPrefix{}
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPPrefix) ProtoMessage() {}

func (x *IPPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPPrefix.ProtoReflect.Descriptor instead.
func (*IPPrefix) Descriptor() ([]byte, []int) {
	return file_v1alpha1_client_pb_affiliate_proto_rawDescGZIP(), []int{2}
}

func (x *IPPrefix) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *IPPrefix) GetBits() uint32 {
	if x != nil {
		return x.Bits
	}
	return 0
}

// Endpoint for the Affiliate KubeSpan spec as seen internally in the client code.
//
// Discovery service receives Affiliate marshaled to protobuf and encrypted.
type Endpoint struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// IPs are encoded using binary marshaling.
	Ip            []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port          uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_v1alpha1_client_pb_affiliate_proto_rawDescGZIP(), []int{3}
}

func (x *Endpoint) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *Endpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// ControlPlane optional configuration.
type ControlPlane struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiServerPort uint32                 `protobuf:"varint,1,opt,name=api_server_port,json=apiServerPort,proto3" json:"api_server_port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ControlPlane) Reset() {
	*x = ControlPlane{}
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlPlane) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPlane) ProtoMessage() {}

func (x *ControlPlane) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_client_pb_affiliate_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPlane.ProtoReflect.Descriptor instead.
func (*ControlPlane) Descriptor() ([]byte, []int) {
	return file_v1alpha1_client_pb_affiliate_proto_rawDescGZIP(), []int{4}
}

func (x *ControlPlane) GetApiServerPort() uint32 {
	if x != nil {
		return x.ApiServerPort
	}
	return 0
}

var File_v1alpha1_client_pb_affiliate_proto protoreflect.FileDescriptor

var file_v1alpha1_client_pb_affiliate_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xfc, 0x02,
	0x0a, 0x09, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x42, 0x0a, 0x08, 0x6b, 0x75, 0x62, 0x65,
	0x73, 0x70, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x69, 0x64,
	0x65, 0x72, 0x6f, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x48, 0x00, 0x52,
	0x08, 0x6b, 0x75, 0x62, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x48, 0x01, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x22, 0x99, 0x01, 0x0a,
	0x08, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x54, 0x0a, 0x14, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x50, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x52, 0x13, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x49, 0x50, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x62, 0x69, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x36, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_v1alpha1_client_pb_affiliate_proto_rawDescOnce sync.Once
	file_v1alpha1_client_pb_affiliate_proto_rawDescData []byte
)

func file_v1alpha1_client_pb_affiliate_proto_rawDescGZIP() []byte {
	file_v1alpha1_client_pb_affiliate_proto_rawDescOnce.Do(func() {
		file_v1alpha1_client_pb_affiliate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1alpha1_client_pb_affiliate_proto_rawDesc), len(file_v1alpha1_client_pb_affiliate_proto_rawDesc)))
	})
	return file_v1alpha1_client_pb_affiliate_proto_rawDescData
}

var file_v1alpha1_client_pb_affiliate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1alpha1_client_pb_affiliate_proto_goTypes = []any{
	(*Affiliate)(nil),    // 0: sidero.discovery.client.Affiliate
	(*KubeSpan)(nil),     // 1: sidero.discovery.client.KubeSpan
	(*IPPrefix)(nil),     // 2: sidero.discovery.client.IPPrefix
	(*Endpoint)(nil),     // 3: sidero.discovery.client.Endpoint
	(*ControlPlane)(nil), // 4: sidero.discovery.client.ControlPlane
}
var file_v1alpha1_client_pb_affiliate_proto_depIdxs = []int32{
	1, // 0: sidero.discovery.client.Affiliate.kubespan:type_name -> sidero.discovery.client.KubeSpan
	4, // 1: sidero.discovery.client.Affiliate.control_plane:type_name -> sidero.discovery.client.ControlPlane
	2, // 2: sidero.discovery.client.KubeSpan.additional_addresses:type_name -> sidero.discovery.client.IPPrefix
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1alpha1_client_pb_affiliate_proto_init() }
func file_v1alpha1_client_pb_affiliate_proto_init() {
	if File_v1alpha1_client_pb_affiliate_proto != nil {
		return
	}
	file_v1alpha1_client_pb_affiliate_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1alpha1_client_pb_affiliate_proto_rawDesc), len(file_v1alpha1_client_pb_affiliate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1alpha1_client_pb_affiliate_proto_goTypes,
		DependencyIndexes: file_v1alpha1_client_pb_affiliate_proto_depIdxs,
		MessageInfos:      file_v1alpha1_client_pb_affiliate_proto_msgTypes,
	}.Build()
	File_v1alpha1_client_pb_affiliate_proto = out.File
	file_v1alpha1_client_pb_affiliate_proto_goTypes = nil
	file_v1alpha1_client_pb_affiliate_proto_depIdxs = nil
}

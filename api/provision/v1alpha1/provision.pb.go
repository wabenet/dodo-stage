// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: provision/v1alpha1/provision.proto

package v1alpha1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1alpha5 "github.com/wabenet/dodo-core/api/core/v1alpha5"
	v1alpha3 "github.com/wabenet/dodo-stage/api/stage/v1alpha3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	CaPath   string `protobuf:"bytes,2,opt,name=ca_path,json=caPath,proto3" json:"ca_path,omitempty"`
	CertPath string `protobuf:"bytes,3,opt,name=cert_path,json=certPath,proto3" json:"cert_path,omitempty"`
	KeyPath  string `protobuf:"bytes,4,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
}

func (x *ProxyConfig) Reset() {
	*x = ProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provision_v1alpha1_provision_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig) ProtoMessage() {}

func (x *ProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_provision_v1alpha1_provision_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig.ProtoReflect.Descriptor instead.
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return file_provision_v1alpha1_provision_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyConfig) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ProxyConfig) GetCaPath() string {
	if x != nil {
		return x.CaPath
	}
	return ""
}

func (x *ProxyConfig) GetCertPath() string {
	if x != nil {
		return x.CertPath
	}
	return ""
}

func (x *ProxyConfig) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

type ProvisionStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stage      *v1alpha3.StageInfo  `protobuf:"bytes,2,opt,name=stage,proto3" json:"stage,omitempty"`
	SshOptions *v1alpha3.SSHOptions `protobuf:"bytes,3,opt,name=ssh_options,json=sshOptions,proto3" json:"ssh_options,omitempty"`
}

func (x *ProvisionStageRequest) Reset() {
	*x = ProvisionStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provision_v1alpha1_provision_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionStageRequest) ProtoMessage() {}

func (x *ProvisionStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provision_v1alpha1_provision_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionStageRequest.ProtoReflect.Descriptor instead.
func (*ProvisionStageRequest) Descriptor() ([]byte, []int) {
	return file_provision_v1alpha1_provision_proto_rawDescGZIP(), []int{1}
}

func (x *ProvisionStageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProvisionStageRequest) GetStage() *v1alpha3.StageInfo {
	if x != nil {
		return x.Stage
	}
	return nil
}

func (x *ProvisionStageRequest) GetSshOptions() *v1alpha3.SSHOptions {
	if x != nil {
		return x.SshOptions
	}
	return nil
}

type CleanStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stage *v1alpha3.StageInfo `protobuf:"bytes,2,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *CleanStageRequest) Reset() {
	*x = CleanStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provision_v1alpha1_provision_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanStageRequest) ProtoMessage() {}

func (x *CleanStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provision_v1alpha1_provision_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanStageRequest.ProtoReflect.Descriptor instead.
func (*CleanStageRequest) Descriptor() ([]byte, []int) {
	return file_provision_v1alpha1_provision_proto_rawDescGZIP(), []int{2}
}

func (x *CleanStageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CleanStageRequest) GetStage() *v1alpha3.StageInfo {
	if x != nil {
		return x.Stage
	}
	return nil
}

type GetProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stage *v1alpha3.StageInfo `protobuf:"bytes,2,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *GetProxyRequest) Reset() {
	*x = GetProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provision_v1alpha1_provision_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProxyRequest) ProtoMessage() {}

func (x *GetProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provision_v1alpha1_provision_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProxyRequest.ProtoReflect.Descriptor instead.
func (*GetProxyRequest) Descriptor() ([]byte, []int) {
	return file_provision_v1alpha1_provision_proto_rawDescGZIP(), []int{3}
}

func (x *GetProxyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProxyRequest) GetStage() *v1alpha3.StageInfo {
	if x != nil {
		return x.Stage
	}
	return nil
}

type GetProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ProxyConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetProxyResponse) Reset() {
	*x = GetProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provision_v1alpha1_provision_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProxyResponse) ProtoMessage() {}

func (x *GetProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provision_v1alpha1_provision_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProxyResponse.ProtoReflect.Descriptor instead.
func (*GetProxyResponse) Descriptor() ([]byte, []int) {
	return file_provision_v1alpha1_provision_proto_rawDescGZIP(), []int{4}
}

func (x *GetProxyResponse) GetConfig() *ProxyConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_provision_v1alpha1_provision_proto protoreflect.FileDescriptor

var file_provision_v1alpha1_provision_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65, 0x6e, 0x65,
	0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x35, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68,
	0x22, 0xbb, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f,
	0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x4c, 0x0a, 0x0b, 0x73, 0x73, 0x68, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65,
	0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x53, 0x53, 0x48, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0a, 0x73, 0x73, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x69,
	0x0a, 0x11, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62,
	0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x67, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f,
	0x64, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x33, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x22, 0x5c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62,
	0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x32, 0xb3, 0x04, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x53, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65, 0x6e,
	0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x35, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x58, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62,
	0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x35, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x64, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x5c, 0x0a, 0x0a, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x36, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x77, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x77, 0x61, 0x62, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x62, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x64, 0x6f,
	0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x62, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x64, 0x6f, 0x64,
	0x6f, 0x2d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provision_v1alpha1_provision_proto_rawDescOnce sync.Once
	file_provision_v1alpha1_provision_proto_rawDescData = file_provision_v1alpha1_provision_proto_rawDesc
)

func file_provision_v1alpha1_provision_proto_rawDescGZIP() []byte {
	file_provision_v1alpha1_provision_proto_rawDescOnce.Do(func() {
		file_provision_v1alpha1_provision_proto_rawDescData = protoimpl.X.CompressGZIP(file_provision_v1alpha1_provision_proto_rawDescData)
	})
	return file_provision_v1alpha1_provision_proto_rawDescData
}

var file_provision_v1alpha1_provision_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_provision_v1alpha1_provision_proto_goTypes = []interface{}{
	(*ProxyConfig)(nil),                 // 0: com.wabenet.dodo.provision.v1alpha1.ProxyConfig
	(*ProvisionStageRequest)(nil),       // 1: com.wabenet.dodo.provision.v1alpha1.ProvisionStageRequest
	(*CleanStageRequest)(nil),           // 2: com.wabenet.dodo.provision.v1alpha1.CleanStageRequest
	(*GetProxyRequest)(nil),             // 3: com.wabenet.dodo.provision.v1alpha1.GetProxyRequest
	(*GetProxyResponse)(nil),            // 4: com.wabenet.dodo.provision.v1alpha1.GetProxyResponse
	(*v1alpha3.StageInfo)(nil),          // 5: com.wabenet.dodo.stage.v1alpha3.StageInfo
	(*v1alpha3.SSHOptions)(nil),         // 6: com.wabenet.dodo.stage.v1alpha3.SSHOptions
	(*empty.Empty)(nil),                 // 7: google.protobuf.Empty
	(*v1alpha5.PluginInfo)(nil),         // 8: com.wabenet.dodo.core.v1alpha5.PluginInfo
	(*v1alpha5.InitPluginResponse)(nil), // 9: com.wabenet.dodo.core.v1alpha5.InitPluginResponse
}
var file_provision_v1alpha1_provision_proto_depIdxs = []int32{
	5,  // 0: com.wabenet.dodo.provision.v1alpha1.ProvisionStageRequest.stage:type_name -> com.wabenet.dodo.stage.v1alpha3.StageInfo
	6,  // 1: com.wabenet.dodo.provision.v1alpha1.ProvisionStageRequest.ssh_options:type_name -> com.wabenet.dodo.stage.v1alpha3.SSHOptions
	5,  // 2: com.wabenet.dodo.provision.v1alpha1.CleanStageRequest.stage:type_name -> com.wabenet.dodo.stage.v1alpha3.StageInfo
	5,  // 3: com.wabenet.dodo.provision.v1alpha1.GetProxyRequest.stage:type_name -> com.wabenet.dodo.stage.v1alpha3.StageInfo
	0,  // 4: com.wabenet.dodo.provision.v1alpha1.GetProxyResponse.config:type_name -> com.wabenet.dodo.provision.v1alpha1.ProxyConfig
	7,  // 5: com.wabenet.dodo.provision.v1alpha1.Plugin.GetPluginInfo:input_type -> google.protobuf.Empty
	7,  // 6: com.wabenet.dodo.provision.v1alpha1.Plugin.InitPlugin:input_type -> google.protobuf.Empty
	7,  // 7: com.wabenet.dodo.provision.v1alpha1.Plugin.ResetPlugin:input_type -> google.protobuf.Empty
	1,  // 8: com.wabenet.dodo.provision.v1alpha1.Plugin.ProvisionStage:input_type -> com.wabenet.dodo.provision.v1alpha1.ProvisionStageRequest
	2,  // 9: com.wabenet.dodo.provision.v1alpha1.Plugin.CleanStage:input_type -> com.wabenet.dodo.provision.v1alpha1.CleanStageRequest
	3,  // 10: com.wabenet.dodo.provision.v1alpha1.Plugin.GetProxy:input_type -> com.wabenet.dodo.provision.v1alpha1.GetProxyRequest
	8,  // 11: com.wabenet.dodo.provision.v1alpha1.Plugin.GetPluginInfo:output_type -> com.wabenet.dodo.core.v1alpha5.PluginInfo
	9,  // 12: com.wabenet.dodo.provision.v1alpha1.Plugin.InitPlugin:output_type -> com.wabenet.dodo.core.v1alpha5.InitPluginResponse
	7,  // 13: com.wabenet.dodo.provision.v1alpha1.Plugin.ResetPlugin:output_type -> google.protobuf.Empty
	7,  // 14: com.wabenet.dodo.provision.v1alpha1.Plugin.ProvisionStage:output_type -> google.protobuf.Empty
	7,  // 15: com.wabenet.dodo.provision.v1alpha1.Plugin.CleanStage:output_type -> google.protobuf.Empty
	4,  // 16: com.wabenet.dodo.provision.v1alpha1.Plugin.GetProxy:output_type -> com.wabenet.dodo.provision.v1alpha1.GetProxyResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_provision_v1alpha1_provision_proto_init() }
func file_provision_v1alpha1_provision_proto_init() {
	if File_provision_v1alpha1_provision_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provision_v1alpha1_provision_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyConfig); i {
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
		file_provision_v1alpha1_provision_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionStageRequest); i {
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
		file_provision_v1alpha1_provision_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CleanStageRequest); i {
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
		file_provision_v1alpha1_provision_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProxyRequest); i {
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
		file_provision_v1alpha1_provision_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProxyResponse); i {
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
			RawDescriptor: file_provision_v1alpha1_provision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provision_v1alpha1_provision_proto_goTypes,
		DependencyIndexes: file_provision_v1alpha1_provision_proto_depIdxs,
		MessageInfos:      file_provision_v1alpha1_provision_proto_msgTypes,
	}.Build()
	File_provision_v1alpha1_provision_proto = out.File
	file_provision_v1alpha1_provision_proto_rawDesc = nil
	file_provision_v1alpha1_provision_proto_goTypes = nil
	file_provision_v1alpha1_provision_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*v1alpha5.PluginInfo, error)
	InitPlugin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*v1alpha5.InitPluginResponse, error)
	ResetPlugin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	ProvisionStage(ctx context.Context, in *ProvisionStageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CleanStage(ctx context.Context, in *CleanStageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetProxy(ctx context.Context, in *GetProxyRequest, opts ...grpc.CallOption) (*GetProxyResponse, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*v1alpha5.PluginInfo, error) {
	out := new(v1alpha5.PluginInfo)
	err := c.cc.Invoke(ctx, "/com.wabenet.dodo.provision.v1alpha1.Plugin/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) InitPlugin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*v1alpha5.InitPluginResponse, error) {
	out := new(v1alpha5.InitPluginResponse)
	err := c.cc.Invoke(ctx, "/com.wabenet.dodo.provision.v1alpha1.Plugin/InitPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ResetPlugin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.wabenet.dodo.provision.v1alpha1.Plugin/ResetPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ProvisionStage(ctx context.Context, in *ProvisionStageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.wabenet.dodo.provision.v1alpha1.Plugin/ProvisionStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) CleanStage(ctx context.Context, in *CleanStageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.wabenet.dodo.provision.v1alpha1.Plugin/CleanStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) GetProxy(ctx context.Context, in *GetProxyRequest, opts ...grpc.CallOption) (*GetProxyResponse, error) {
	out := new(GetProxyResponse)
	err := c.cc.Invoke(ctx, "/com.wabenet.dodo.provision.v1alpha1.Plugin/GetProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	GetPluginInfo(context.Context, *empty.Empty) (*v1alpha5.PluginInfo, error)
	InitPlugin(context.Context, *empty.Empty) (*v1alpha5.InitPluginResponse, error)
	ResetPlugin(context.Context, *empty.Empty) (*empty.Empty, error)
	ProvisionStage(context.Context, *ProvisionStageRequest) (*empty.Empty, error)
	CleanStage(context.Context, *CleanStageRequest) (*empty.Empty, error)
	GetProxy(context.Context, *GetProxyRequest) (*GetProxyResponse, error)
}

// UnimplementedPluginServer can be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (*UnimplementedPluginServer) GetPluginInfo(context.Context, *empty.Empty) (*v1alpha5.PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}
func (*UnimplementedPluginServer) InitPlugin(context.Context, *empty.Empty) (*v1alpha5.InitPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitPlugin not implemented")
}
func (*UnimplementedPluginServer) ResetPlugin(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPlugin not implemented")
}
func (*UnimplementedPluginServer) ProvisionStage(context.Context, *ProvisionStageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionStage not implemented")
}
func (*UnimplementedPluginServer) CleanStage(context.Context, *CleanStageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanStage not implemented")
}
func (*UnimplementedPluginServer) GetProxy(context.Context, *GetProxyRequest) (*GetProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProxy not implemented")
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.wabenet.dodo.provision.v1alpha1.Plugin/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetPluginInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_InitPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).InitPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.wabenet.dodo.provision.v1alpha1.Plugin/InitPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).InitPlugin(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ResetPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ResetPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.wabenet.dodo.provision.v1alpha1.Plugin/ResetPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ResetPlugin(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ProvisionStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ProvisionStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.wabenet.dodo.provision.v1alpha1.Plugin/ProvisionStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ProvisionStage(ctx, req.(*ProvisionStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_CleanStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).CleanStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.wabenet.dodo.provision.v1alpha1.Plugin/CleanStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).CleanStage(ctx, req.(*CleanStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_GetProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.wabenet.dodo.provision.v1alpha1.Plugin/GetProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetProxy(ctx, req.(*GetProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.wabenet.dodo.provision.v1alpha1.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPluginInfo",
			Handler:    _Plugin_GetPluginInfo_Handler,
		},
		{
			MethodName: "InitPlugin",
			Handler:    _Plugin_InitPlugin_Handler,
		},
		{
			MethodName: "ResetPlugin",
			Handler:    _Plugin_ResetPlugin_Handler,
		},
		{
			MethodName: "ProvisionStage",
			Handler:    _Plugin_ProvisionStage_Handler,
		},
		{
			MethodName: "CleanStage",
			Handler:    _Plugin_CleanStage_Handler,
		},
		{
			MethodName: "GetProxy",
			Handler:    _Plugin_GetProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provision/v1alpha1/provision.proto",
}

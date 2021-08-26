// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ova-obligation-api.proto

package ova_obligation_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateObligationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateObligationRequest) Reset() {
	*x = CreateObligationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_obligation_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateObligationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateObligationRequest) ProtoMessage() {}

func (x *CreateObligationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_obligation_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateObligationRequest.ProtoReflect.Descriptor instead.
func (*CreateObligationRequest) Descriptor() ([]byte, []int) {
	return file_api_ova_obligation_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateObligationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateObligationRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateObligationResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateObligationResponce) Reset() {
	*x = CreateObligationResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_obligation_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateObligationResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateObligationResponce) ProtoMessage() {}

func (x *CreateObligationResponce) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_obligation_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateObligationResponce.ProtoReflect.Descriptor instead.
func (*CreateObligationResponce) Descriptor() ([]byte, []int) {
	return file_api_ova_obligation_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateObligationResponce) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeObligationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeObligationRequest) Reset() {
	*x = DescribeObligationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_obligation_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeObligationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeObligationRequest) ProtoMessage() {}

func (x *DescribeObligationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_obligation_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeObligationRequest.ProtoReflect.Descriptor instead.
func (*DescribeObligationRequest) Descriptor() ([]byte, []int) {
	return file_api_ova_obligation_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeObligationRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeObligationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Obligation *Obligation `protobuf:"bytes,1,opt,name=obligation,proto3" json:"obligation,omitempty"`
}

func (x *DescribeObligationResponse) Reset() {
	*x = DescribeObligationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_obligation_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeObligationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeObligationResponse) ProtoMessage() {}

func (x *DescribeObligationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_obligation_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeObligationResponse.ProtoReflect.Descriptor instead.
func (*DescribeObligationResponse) Descriptor() ([]byte, []int) {
	return file_api_ova_obligation_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeObligationResponse) GetObligation() *Obligation {
	if x != nil {
		return x.Obligation
	}
	return nil
}

type RemoveObligationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveObligationRequest) Reset() {
	*x = RemoveObligationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_obligation_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveObligationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveObligationRequest) ProtoMessage() {}

func (x *RemoveObligationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_obligation_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveObligationRequest.ProtoReflect.Descriptor instead.
func (*RemoveObligationRequest) Descriptor() ([]byte, []int) {
	return file_api_ova_obligation_api_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveObligationRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListObligationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Obligation `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListObligationsResponse) Reset() {
	*x = ListObligationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_obligation_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObligationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObligationsResponse) ProtoMessage() {}

func (x *ListObligationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_obligation_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObligationsResponse.ProtoReflect.Descriptor instead.
func (*ListObligationsResponse) Descriptor() ([]byte, []int) {
	return file_api_ova_obligation_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListObligationsResponse) GetItems() []*Obligation {
	if x != nil {
		return x.Items
	}
	return nil
}

type Obligation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Obligation) Reset() {
	*x = Obligation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_obligation_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Obligation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Obligation) ProtoMessage() {}

func (x *Obligation) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_obligation_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Obligation.ProtoReflect.Descriptor instead.
func (*Obligation) Descriptor() ([]byte, []int) {
	return file_api_ova_obligation_api_proto_rawDescGZIP(), []int{6}
}

func (x *Obligation) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Obligation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Obligation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_api_ova_obligation_api_proto protoreflect.FileDescriptor

var file_api_ova_obligation_api_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x51, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b,
	0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x1a, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x6f, 0x62, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f,
	0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x17, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x54, 0x0a, 0x0a, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa4, 0x03, 0x0a, 0x0d,
	0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x70, 0x63, 0x12, 0x6d, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2b, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x12,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f,
	0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x56, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x6f,
	0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e,
	0x6f, 0x76, 0x61, 0x2e, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x6f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f,
	0x76, 0x61, 0x2d, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70,
	0x69, 0x3b, 0x6f, 0x76, 0x61, 0x5f, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ova_obligation_api_proto_rawDescOnce sync.Once
	file_api_ova_obligation_api_proto_rawDescData = file_api_ova_obligation_api_proto_rawDesc
)

func file_api_ova_obligation_api_proto_rawDescGZIP() []byte {
	file_api_ova_obligation_api_proto_rawDescOnce.Do(func() {
		file_api_ova_obligation_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ova_obligation_api_proto_rawDescData)
	})
	return file_api_ova_obligation_api_proto_rawDescData
}

var file_api_ova_obligation_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_ova_obligation_api_proto_goTypes = []interface{}{
	(*CreateObligationRequest)(nil),    // 0: ova.obligation.api.CreateObligationRequest
	(*CreateObligationResponce)(nil),   // 1: ova.obligation.api.CreateObligationResponce
	(*DescribeObligationRequest)(nil),  // 2: ova.obligation.api.DescribeObligationRequest
	(*DescribeObligationResponse)(nil), // 3: ova.obligation.api.DescribeObligationResponse
	(*RemoveObligationRequest)(nil),    // 4: ova.obligation.api.RemoveObligationRequest
	(*ListObligationsResponse)(nil),    // 5: ova.obligation.api.ListObligationsResponse
	(*Obligation)(nil),                 // 6: ova.obligation.api.Obligation
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_api_ova_obligation_api_proto_depIdxs = []int32{
	6, // 0: ova.obligation.api.DescribeObligationResponse.obligation:type_name -> ova.obligation.api.Obligation
	6, // 1: ova.obligation.api.ListObligationsResponse.items:type_name -> ova.obligation.api.Obligation
	0, // 2: ova.obligation.api.ObligationRpc.CreateObligation:input_type -> ova.obligation.api.CreateObligationRequest
	2, // 3: ova.obligation.api.ObligationRpc.DescribeObligation:input_type -> ova.obligation.api.DescribeObligationRequest
	7, // 4: ova.obligation.api.ObligationRpc.ListObligations:input_type -> google.protobuf.Empty
	4, // 5: ova.obligation.api.ObligationRpc.RemoveObligation:input_type -> ova.obligation.api.RemoveObligationRequest
	1, // 6: ova.obligation.api.ObligationRpc.CreateObligation:output_type -> ova.obligation.api.CreateObligationResponce
	3, // 7: ova.obligation.api.ObligationRpc.DescribeObligation:output_type -> ova.obligation.api.DescribeObligationResponse
	5, // 8: ova.obligation.api.ObligationRpc.ListObligations:output_type -> ova.obligation.api.ListObligationsResponse
	7, // 9: ova.obligation.api.ObligationRpc.RemoveObligation:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ova_obligation_api_proto_init() }
func file_api_ova_obligation_api_proto_init() {
	if File_api_ova_obligation_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ova_obligation_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateObligationRequest); i {
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
		file_api_ova_obligation_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateObligationResponce); i {
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
		file_api_ova_obligation_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeObligationRequest); i {
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
		file_api_ova_obligation_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeObligationResponse); i {
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
		file_api_ova_obligation_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveObligationRequest); i {
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
		file_api_ova_obligation_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObligationsResponse); i {
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
		file_api_ova_obligation_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Obligation); i {
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
			RawDescriptor: file_api_ova_obligation_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ova_obligation_api_proto_goTypes,
		DependencyIndexes: file_api_ova_obligation_api_proto_depIdxs,
		MessageInfos:      file_api_ova_obligation_api_proto_msgTypes,
	}.Build()
	File_api_ova_obligation_api_proto = out.File
	file_api_ova_obligation_api_proto_rawDesc = nil
	file_api_ova_obligation_api_proto_goTypes = nil
	file_api_ova_obligation_api_proto_depIdxs = nil
}
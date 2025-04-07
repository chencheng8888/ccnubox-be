// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.26.1
// source: infoSum/v1/infoSum.proto

package infoSumv1

import (
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

// web
type InfoSum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Link        string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Image       string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InfoSum) Reset() {
	*x = InfoSum{}
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoSum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoSum) ProtoMessage() {}

func (x *InfoSum) ProtoReflect() protoreflect.Message {
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoSum.ProtoReflect.Descriptor instead.
func (*InfoSum) Descriptor() ([]byte, []int) {
	return file_infoSum_v1_infoSum_proto_rawDescGZIP(), []int{0}
}

func (x *InfoSum) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InfoSum) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *InfoSum) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfoSum) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *InfoSum) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetInfoSumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoSumsRequest) Reset() {
	*x = GetInfoSumsRequest{}
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoSumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoSumsRequest) ProtoMessage() {}

func (x *GetInfoSumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoSumsRequest.ProtoReflect.Descriptor instead.
func (*GetInfoSumsRequest) Descriptor() ([]byte, []int) {
	return file_infoSum_v1_infoSum_proto_rawDescGZIP(), []int{1}
}

type GetInfoSumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoSums []*InfoSum `protobuf:"bytes,1,rep,name=InfoSums,proto3" json:"InfoSums,omitempty"`
}

func (x *GetInfoSumsResponse) Reset() {
	*x = GetInfoSumsResponse{}
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoSumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoSumsResponse) ProtoMessage() {}

func (x *GetInfoSumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoSumsResponse.ProtoReflect.Descriptor instead.
func (*GetInfoSumsResponse) Descriptor() ([]byte, []int) {
	return file_infoSum_v1_infoSum_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoSumsResponse) GetInfoSums() []*InfoSum {
	if x != nil {
		return x.InfoSums
	}
	return nil
}

type SaveInfoSumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoSum *InfoSum `protobuf:"bytes,1,opt,name=InfoSum,proto3" json:"InfoSum,omitempty"`
}

func (x *SaveInfoSumRequest) Reset() {
	*x = SaveInfoSumRequest{}
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveInfoSumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveInfoSumRequest) ProtoMessage() {}

func (x *SaveInfoSumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveInfoSumRequest.ProtoReflect.Descriptor instead.
func (*SaveInfoSumRequest) Descriptor() ([]byte, []int) {
	return file_infoSum_v1_infoSum_proto_rawDescGZIP(), []int{3}
}

func (x *SaveInfoSumRequest) GetInfoSum() *InfoSum {
	if x != nil {
		return x.InfoSum
	}
	return nil
}

type SaveInfoSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoSums []*InfoSum `protobuf:"bytes,1,rep,name=InfoSums,proto3" json:"InfoSums,omitempty"`
}

func (x *SaveInfoSumResponse) Reset() {
	*x = SaveInfoSumResponse{}
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveInfoSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveInfoSumResponse) ProtoMessage() {}

func (x *SaveInfoSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveInfoSumResponse.ProtoReflect.Descriptor instead.
func (*SaveInfoSumResponse) Descriptor() ([]byte, []int) {
	return file_infoSum_v1_infoSum_proto_rawDescGZIP(), []int{4}
}

func (x *SaveInfoSumResponse) GetInfoSums() []*InfoSum {
	if x != nil {
		return x.InfoSums
	}
	return nil
}

type DelInfoSumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelInfoSumRequest) Reset() {
	*x = DelInfoSumRequest{}
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelInfoSumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelInfoSumRequest) ProtoMessage() {}

func (x *DelInfoSumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelInfoSumRequest.ProtoReflect.Descriptor instead.
func (*DelInfoSumRequest) Descriptor() ([]byte, []int) {
	return file_infoSum_v1_infoSum_proto_rawDescGZIP(), []int{5}
}

func (x *DelInfoSumRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelInfoSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoSums []*InfoSum `protobuf:"bytes,1,rep,name=InfoSums,proto3" json:"InfoSums,omitempty"`
}

func (x *DelInfoSumResponse) Reset() {
	*x = DelInfoSumResponse{}
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelInfoSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelInfoSumResponse) ProtoMessage() {}

func (x *DelInfoSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infoSum_v1_infoSum_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelInfoSumResponse.ProtoReflect.Descriptor instead.
func (*DelInfoSumResponse) Descriptor() ([]byte, []int) {
	return file_infoSum_v1_infoSum_proto_rawDescGZIP(), []int{6}
}

func (x *DelInfoSumResponse) GetInfoSums() []*InfoSum {
	if x != nil {
		return x.InfoSums
	}
	return nil
}

var File_infoSum_v1_infoSum_proto protoreflect.FileDescriptor

var file_infoSum_v1_infoSum_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x53, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x66, 0x6f,
	0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0x79, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x75, 0x6d, 0x52, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73, 0x22,
	0x43, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x52, 0x07, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x75, 0x6d, 0x22, 0x46, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x75, 0x6d, 0x52, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73, 0x22, 0x23, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x45, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x52, 0x08,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73, 0x32, 0xfd, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x73, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x53,
	0x61, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x53, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x53,
	0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x63, 0x6e, 0x75, 0x2f,
	0x63, 0x63, 0x6e, 0x75, 0x62, 0x6f, 0x78, 0x2d, 0x62, 0x65, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x53, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x66, 0x6f, 0x53, 0x75, 0x6d, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infoSum_v1_infoSum_proto_rawDescOnce sync.Once
	file_infoSum_v1_infoSum_proto_rawDescData = file_infoSum_v1_infoSum_proto_rawDesc
)

func file_infoSum_v1_infoSum_proto_rawDescGZIP() []byte {
	file_infoSum_v1_infoSum_proto_rawDescOnce.Do(func() {
		file_infoSum_v1_infoSum_proto_rawDescData = protoimpl.X.CompressGZIP(file_infoSum_v1_infoSum_proto_rawDescData)
	})
	return file_infoSum_v1_infoSum_proto_rawDescData
}

var file_infoSum_v1_infoSum_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_infoSum_v1_infoSum_proto_goTypes = []any{
	(*InfoSum)(nil),             // 0: infoSum.v1.InfoSum
	(*GetInfoSumsRequest)(nil),  // 1: infoSum.v1.GetInfoSumsRequest
	(*GetInfoSumsResponse)(nil), // 2: infoSum.v1.GetInfoSumsResponse
	(*SaveInfoSumRequest)(nil),  // 3: infoSum.v1.SaveInfoSumRequest
	(*SaveInfoSumResponse)(nil), // 4: infoSum.v1.SaveInfoSumResponse
	(*DelInfoSumRequest)(nil),   // 5: infoSum.v1.DelInfoSumRequest
	(*DelInfoSumResponse)(nil),  // 6: infoSum.v1.DelInfoSumResponse
}
var file_infoSum_v1_infoSum_proto_depIdxs = []int32{
	0, // 0: infoSum.v1.GetInfoSumsResponse.InfoSums:type_name -> infoSum.v1.InfoSum
	0, // 1: infoSum.v1.SaveInfoSumRequest.InfoSum:type_name -> infoSum.v1.InfoSum
	0, // 2: infoSum.v1.SaveInfoSumResponse.InfoSums:type_name -> infoSum.v1.InfoSum
	0, // 3: infoSum.v1.DelInfoSumResponse.InfoSums:type_name -> infoSum.v1.InfoSum
	1, // 4: infoSum.v1.InfoSumService.GetInfoSums:input_type -> infoSum.v1.GetInfoSumsRequest
	3, // 5: infoSum.v1.InfoSumService.SaveInfoSum:input_type -> infoSum.v1.SaveInfoSumRequest
	5, // 6: infoSum.v1.InfoSumService.DelInfoSum:input_type -> infoSum.v1.DelInfoSumRequest
	2, // 7: infoSum.v1.InfoSumService.GetInfoSums:output_type -> infoSum.v1.GetInfoSumsResponse
	4, // 8: infoSum.v1.InfoSumService.SaveInfoSum:output_type -> infoSum.v1.SaveInfoSumResponse
	6, // 9: infoSum.v1.InfoSumService.DelInfoSum:output_type -> infoSum.v1.DelInfoSumResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_infoSum_v1_infoSum_proto_init() }
func file_infoSum_v1_infoSum_proto_init() {
	if File_infoSum_v1_infoSum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infoSum_v1_infoSum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infoSum_v1_infoSum_proto_goTypes,
		DependencyIndexes: file_infoSum_v1_infoSum_proto_depIdxs,
		MessageInfos:      file_infoSum_v1_infoSum_proto_msgTypes,
	}.Build()
	File_infoSum_v1_infoSum_proto = out.File
	file_infoSum_v1_infoSum_proto_rawDesc = nil
	file_infoSum_v1_infoSum_proto_goTypes = nil
	file_infoSum_v1_infoSum_proto_depIdxs = nil
}

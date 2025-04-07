// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.26.1
// source: static/v1/static.proto

package staticv1

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

type Static struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string            `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Labels  map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Static) Reset() {
	*x = Static{}
	mi := &file_static_v1_static_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Static) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Static) ProtoMessage() {}

func (x *Static) ProtoReflect() protoreflect.Message {
	mi := &file_static_v1_static_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Static.ProtoReflect.Descriptor instead.
func (*Static) Descriptor() ([]byte, []int) {
	return file_static_v1_static_proto_rawDescGZIP(), []int{0}
}

func (x *Static) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Static) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Static) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetStaticByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetStaticByNameRequest) Reset() {
	*x = GetStaticByNameRequest{}
	mi := &file_static_v1_static_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStaticByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticByNameRequest) ProtoMessage() {}

func (x *GetStaticByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_static_v1_static_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticByNameRequest.ProtoReflect.Descriptor instead.
func (*GetStaticByNameRequest) Descriptor() ([]byte, []int) {
	return file_static_v1_static_proto_rawDescGZIP(), []int{1}
}

func (x *GetStaticByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetStaticByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Static *Static `protobuf:"bytes,1,opt,name=static,proto3" json:"static,omitempty"`
}

func (x *GetStaticByNameResponse) Reset() {
	*x = GetStaticByNameResponse{}
	mi := &file_static_v1_static_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStaticByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticByNameResponse) ProtoMessage() {}

func (x *GetStaticByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_static_v1_static_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticByNameResponse.ProtoReflect.Descriptor instead.
func (*GetStaticByNameResponse) Descriptor() ([]byte, []int) {
	return file_static_v1_static_proto_rawDescGZIP(), []int{2}
}

func (x *GetStaticByNameResponse) GetStatic() *Static {
	if x != nil {
		return x.Static
	}
	return nil
}

type SaveStaticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Static *Static `protobuf:"bytes,1,opt,name=static,proto3" json:"static,omitempty"`
}

func (x *SaveStaticRequest) Reset() {
	*x = SaveStaticRequest{}
	mi := &file_static_v1_static_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveStaticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStaticRequest) ProtoMessage() {}

func (x *SaveStaticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_static_v1_static_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStaticRequest.ProtoReflect.Descriptor instead.
func (*SaveStaticRequest) Descriptor() ([]byte, []int) {
	return file_static_v1_static_proto_rawDescGZIP(), []int{3}
}

func (x *SaveStaticRequest) GetStatic() *Static {
	if x != nil {
		return x.Static
	}
	return nil
}

type SaveStaticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveStaticResponse) Reset() {
	*x = SaveStaticResponse{}
	mi := &file_static_v1_static_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveStaticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStaticResponse) ProtoMessage() {}

func (x *SaveStaticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_static_v1_static_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStaticResponse.ProtoReflect.Descriptor instead.
func (*SaveStaticResponse) Descriptor() ([]byte, []int) {
	return file_static_v1_static_proto_rawDescGZIP(), []int{4}
}

type GetStaticsByLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetStaticsByLabelsRequest) Reset() {
	*x = GetStaticsByLabelsRequest{}
	mi := &file_static_v1_static_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStaticsByLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticsByLabelsRequest) ProtoMessage() {}

func (x *GetStaticsByLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_static_v1_static_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticsByLabelsRequest.ProtoReflect.Descriptor instead.
func (*GetStaticsByLabelsRequest) Descriptor() ([]byte, []int) {
	return file_static_v1_static_proto_rawDescGZIP(), []int{5}
}

func (x *GetStaticsByLabelsRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetStaticsByLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statics []*Static `protobuf:"bytes,1,rep,name=statics,proto3" json:"statics,omitempty"`
}

func (x *GetStaticsByLabelsResponse) Reset() {
	*x = GetStaticsByLabelsResponse{}
	mi := &file_static_v1_static_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStaticsByLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticsByLabelsResponse) ProtoMessage() {}

func (x *GetStaticsByLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_static_v1_static_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticsByLabelsResponse.ProtoReflect.Descriptor instead.
func (*GetStaticsByLabelsResponse) Descriptor() ([]byte, []int) {
	return file_static_v1_static_proto_rawDescGZIP(), []int{6}
}

func (x *GetStaticsByLabelsResponse) GetStatics() []*Static {
	if x != nil {
		return x.Statics
	}
	return nil
}

var File_static_v1_static_proto protoreflect.FileDescriptor

var file_static_v1_static_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x22, 0xa8, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2c,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x22, 0x3e, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x42, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x42, 0x79, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x49, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x42, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x32, 0x97, 0x02, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x42, 0x79, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x42, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73,
	0x42, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x73, 0x79, 0x6e, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x63, 0x63, 0x6e, 0x75, 0x62, 0x6f, 0x78, 0x2d,
	0x62, 0x65, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_static_v1_static_proto_rawDescOnce sync.Once
	file_static_v1_static_proto_rawDescData = file_static_v1_static_proto_rawDesc
)

func file_static_v1_static_proto_rawDescGZIP() []byte {
	file_static_v1_static_proto_rawDescOnce.Do(func() {
		file_static_v1_static_proto_rawDescData = protoimpl.X.CompressGZIP(file_static_v1_static_proto_rawDescData)
	})
	return file_static_v1_static_proto_rawDescData
}

var file_static_v1_static_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_static_v1_static_proto_goTypes = []any{
	(*Static)(nil),                     // 0: static.v1.Static
	(*GetStaticByNameRequest)(nil),     // 1: static.v1.GetStaticByNameRequest
	(*GetStaticByNameResponse)(nil),    // 2: static.v1.GetStaticByNameResponse
	(*SaveStaticRequest)(nil),          // 3: static.v1.SaveStaticRequest
	(*SaveStaticResponse)(nil),         // 4: static.v1.SaveStaticResponse
	(*GetStaticsByLabelsRequest)(nil),  // 5: static.v1.GetStaticsByLabelsRequest
	(*GetStaticsByLabelsResponse)(nil), // 6: static.v1.GetStaticsByLabelsResponse
	nil,                                // 7: static.v1.Static.LabelsEntry
	nil,                                // 8: static.v1.GetStaticsByLabelsRequest.LabelsEntry
}
var file_static_v1_static_proto_depIdxs = []int32{
	7, // 0: static.v1.Static.labels:type_name -> static.v1.Static.LabelsEntry
	0, // 1: static.v1.GetStaticByNameResponse.static:type_name -> static.v1.Static
	0, // 2: static.v1.SaveStaticRequest.static:type_name -> static.v1.Static
	8, // 3: static.v1.GetStaticsByLabelsRequest.labels:type_name -> static.v1.GetStaticsByLabelsRequest.LabelsEntry
	0, // 4: static.v1.GetStaticsByLabelsResponse.statics:type_name -> static.v1.Static
	1, // 5: static.v1.StaticService.GetStaticByName:input_type -> static.v1.GetStaticByNameRequest
	3, // 6: static.v1.StaticService.SaveStatic:input_type -> static.v1.SaveStaticRequest
	5, // 7: static.v1.StaticService.GetStaticsByLabels:input_type -> static.v1.GetStaticsByLabelsRequest
	2, // 8: static.v1.StaticService.GetStaticByName:output_type -> static.v1.GetStaticByNameResponse
	4, // 9: static.v1.StaticService.SaveStatic:output_type -> static.v1.SaveStaticResponse
	6, // 10: static.v1.StaticService.GetStaticsByLabels:output_type -> static.v1.GetStaticsByLabelsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_static_v1_static_proto_init() }
func file_static_v1_static_proto_init() {
	if File_static_v1_static_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_static_v1_static_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_static_v1_static_proto_goTypes,
		DependencyIndexes: file_static_v1_static_proto_depIdxs,
		MessageInfos:      file_static_v1_static_proto_msgTypes,
	}.Build()
	File_static_v1_static_proto = out.File
	file_static_v1_static_proto_rawDesc = nil
	file_static_v1_static_proto_goTypes = nil
	file_static_v1_static_proto_depIdxs = nil
}

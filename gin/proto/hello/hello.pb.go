// 协议类型

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.1
// source: hello.proto

package hello

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HelloDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloDBRequest) Reset() {
	*x = HelloDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloDBRequest) ProtoMessage() {}

func (x *HelloDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloDBRequest.ProtoReflect.Descriptor instead.
func (*HelloDBRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloDBRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloDBResponse) Reset() {
	*x = HelloDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloDBResponse) ProtoMessage() {}

func (x *HelloDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloDBResponse.ProtoReflect.Descriptor instead.
func (*HelloDBResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloDBResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HelloDBResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GoodByeDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GoodByeDBRequest) Reset() {
	*x = GoodByeDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodByeDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodByeDBRequest) ProtoMessage() {}

func (x *GoodByeDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodByeDBRequest.ProtoReflect.Descriptor instead.
func (*GoodByeDBRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{2}
}

func (x *GoodByeDBRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GoodByeDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GoodByeDBResponse) Reset() {
	*x = GoodByeDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodByeDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodByeDBResponse) ProtoMessage() {}

func (x *GoodByeDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodByeDBResponse.ProtoReflect.Descriptor instead.
func (*GoodByeDBResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{3}
}

func (x *GoodByeDBResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GoodByeDBResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloHttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloHttpRequest) Reset() {
	*x = HelloHttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloHttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloHttpRequest) ProtoMessage() {}

func (x *HelloHttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloHttpRequest.ProtoReflect.Descriptor instead.
func (*HelloHttpRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{4}
}

func (x *HelloHttpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloHttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloHttpResponse) Reset() {
	*x = HelloHttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloHttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloHttpResponse) ProtoMessage() {}

func (x *HelloHttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloHttpResponse.ProtoReflect.Descriptor instead.
func (*HelloHttpResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{5}
}

func (x *HelloHttpResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HelloHttpResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GoodByeHttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GoodByeHttpRequest) Reset() {
	*x = GoodByeHttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodByeHttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodByeHttpRequest) ProtoMessage() {}

func (x *GoodByeHttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodByeHttpRequest.ProtoReflect.Descriptor instead.
func (*GoodByeHttpRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{6}
}

func (x *GoodByeHttpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GoodByeHttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GoodByeHttpResponse) Reset() {
	*x = GoodByeHttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodByeHttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodByeHttpResponse) ProtoMessage() {}

func (x *GoodByeHttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodByeHttpResponse.ProtoReflect.Descriptor instead.
func (*GoodByeHttpResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{7}
}

func (x *GoodByeHttpResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GoodByeHttpResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3f, 0x0a, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x44, 0x42, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x47, 0x6f,
	0x6f, 0x64, 0x42, 0x79, 0x65, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a,
	0x10, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x6f, 0x6f, 0x64,
	0x42, 0x79, 0x65, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x48, 0x74, 0x74,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x71, 0x0a, 0x07, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x44, 0x42, 0x12, 0x2f, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0f,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79,
	0x65, 0x12, 0x11, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x44, 0x42, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x44, 0x42,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xb1, 0x01, 0x0a, 0x09, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x74, 0x74, 0x70, 0x12, 0x4d, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x11, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x74, 0x74, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x73, 0x61, 0x79, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x47, 0x6f,
	0x6f, 0x64, 0x62, 0x79, 0x65, 0x12, 0x13, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65, 0x48,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x42, 0x79, 0x65, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2f, 0x73, 0x61, 0x79, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_hello_proto_goTypes = []interface{}{
	(*HelloDBRequest)(nil),      // 0: HelloDBRequest
	(*HelloDBResponse)(nil),     // 1: HelloDBResponse
	(*GoodByeDBRequest)(nil),    // 2: GoodByeDBRequest
	(*GoodByeDBResponse)(nil),   // 3: GoodByeDBResponse
	(*HelloHttpRequest)(nil),    // 4: HelloHttpRequest
	(*HelloHttpResponse)(nil),   // 5: HelloHttpResponse
	(*GoodByeHttpRequest)(nil),  // 6: GoodByeHttpRequest
	(*GoodByeHttpResponse)(nil), // 7: GoodByeHttpResponse
}
var file_hello_proto_depIdxs = []int32{
	0, // 0: HelloDB.SayHello:input_type -> HelloDBRequest
	2, // 1: HelloDB.SayGoodbye:input_type -> GoodByeDBRequest
	4, // 2: HelloHttp.SayHello:input_type -> HelloHttpRequest
	6, // 3: HelloHttp.SayGoodbye:input_type -> GoodByeHttpRequest
	1, // 4: HelloDB.SayHello:output_type -> HelloDBResponse
	3, // 5: HelloDB.SayGoodbye:output_type -> GoodByeDBResponse
	5, // 6: HelloHttp.SayHello:output_type -> HelloHttpResponse
	7, // 7: HelloHttp.SayGoodbye:output_type -> GoodByeHttpResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloDBRequest); i {
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
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloDBResponse); i {
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
		file_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodByeDBRequest); i {
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
		file_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodByeDBResponse); i {
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
		file_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloHttpRequest); i {
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
		file_hello_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloHttpResponse); i {
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
		file_hello_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodByeHttpRequest); i {
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
		file_hello_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodByeHttpResponse); i {
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
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}

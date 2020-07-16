// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: proto/store/store.proto

package go_micro_service_store_svc

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

type ErrorMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ErrorMsg) Reset() {
	*x = ErrorMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_store_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMsg) ProtoMessage() {}

func (x *ErrorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_store_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMsg.ProtoReflect.Descriptor instead.
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return file_proto_store_store_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetByObjectIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectIds []int64 `protobuf:"varint,1,rep,packed,name=objectIds,proto3" json:"objectIds,omitempty"`
}

func (x *GetByObjectIdsRequest) Reset() {
	*x = GetByObjectIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_store_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByObjectIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByObjectIdsRequest) ProtoMessage() {}

func (x *GetByObjectIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_store_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByObjectIdsRequest.ProtoReflect.Descriptor instead.
func (*GetByObjectIdsRequest) Descriptor() ([]byte, []int) {
	return file_proto_store_store_proto_rawDescGZIP(), []int{1}
}

func (x *GetByObjectIdsRequest) GetObjectIds() []int64 {
	if x != nil {
		return x.ObjectIds
	}
	return nil
}

type GetByObjectIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectInfos []*ObjectInfo `protobuf:"bytes,1,rep,name=objectInfos,proto3" json:"objectInfos,omitempty"`
}

func (x *GetByObjectIdsResponse) Reset() {
	*x = GetByObjectIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_store_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByObjectIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByObjectIdsResponse) ProtoMessage() {}

func (x *GetByObjectIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_store_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByObjectIdsResponse.ProtoReflect.Descriptor instead.
func (*GetByObjectIdsResponse) Descriptor() ([]byte, []int) {
	return file_proto_store_store_proto_rawDescGZIP(), []int{2}
}

func (x *GetByObjectIdsResponse) GetObjectInfos() []*ObjectInfo {
	if x != nil {
		return x.ObjectInfos
	}
	return nil
}

type ObjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bulk       string `protobuf:"bytes,2,opt,name=bulk,proto3" json:"bulk,omitempty"`
	ObjectName string `protobuf:"bytes,3,opt,name=objectName,proto3" json:"objectName,omitempty"`
}

func (x *ObjectInfo) Reset() {
	*x = ObjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_store_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectInfo) ProtoMessage() {}

func (x *ObjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_store_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectInfo.ProtoReflect.Descriptor instead.
func (*ObjectInfo) Descriptor() ([]byte, []int) {
	return file_proto_store_store_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjectInfo) GetBulk() string {
	if x != nil {
		return x.Bulk
	}
	return ""
}

func (x *ObjectInfo) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

type SaveStoreInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreInfos []*StoreInfo `protobuf:"bytes,1,rep,name=storeInfos,proto3" json:"storeInfos,omitempty"`
}

func (x *SaveStoreInfoRequest) Reset() {
	*x = SaveStoreInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_store_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveStoreInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStoreInfoRequest) ProtoMessage() {}

func (x *SaveStoreInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_store_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStoreInfoRequest.ProtoReflect.Descriptor instead.
func (*SaveStoreInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_store_store_proto_rawDescGZIP(), []int{4}
}

func (x *SaveStoreInfoRequest) GetStoreInfos() []*StoreInfo {
	if x != nil {
		return x.StoreInfos
	}
	return nil
}

type StoreInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	BulkName   string `protobuf:"bytes,2,opt,name=bulkName,proto3" json:"bulkName,omitempty"`
	ObjectName string `protobuf:"bytes,3,opt,name=objectName,proto3" json:"objectName,omitempty"`
	Size       int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Width      int32  `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height     int32  `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *StoreInfo) Reset() {
	*x = StoreInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_store_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfo) ProtoMessage() {}

func (x *StoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_store_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfo.ProtoReflect.Descriptor instead.
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return file_proto_store_store_proto_rawDescGZIP(), []int{5}
}

func (x *StoreInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *StoreInfo) GetBulkName() string {
	if x != nil {
		return x.BulkName
	}
	return ""
}

func (x *StoreInfo) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *StoreInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StoreInfo) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *StoreInfo) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type SaveStoreInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name2IdMap map[string]int64 `protobuf:"bytes,1,rep,name=name2IdMap,proto3" json:"name2IdMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ErrorMsg   *ErrorMsg        `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *SaveStoreInfoResponse) Reset() {
	*x = SaveStoreInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_store_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveStoreInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStoreInfoResponse) ProtoMessage() {}

func (x *SaveStoreInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_store_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStoreInfoResponse.ProtoReflect.Descriptor instead.
func (*SaveStoreInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_store_store_proto_rawDescGZIP(), []int{6}
}

func (x *SaveStoreInfoResponse) GetName2IdMap() map[string]int64 {
	if x != nil {
		return x.Name2IdMap
	}
	return nil
}

func (x *SaveStoreInfoResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

var File_proto_store_store_proto protoreflect.FileDescriptor

var file_proto_store_store_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x76, 0x63, 0x22, 0x1c, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x35, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x62, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x50,
	0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x75, 0x6c, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x75, 0x6c, 0x6b,
	0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x5d, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22,
	0xa5, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x6c,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x6c,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x61, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x49, 0x64, 0x4d, 0x61, 0x70, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73,
	0x76, 0x63, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x49, 0x64,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x49,
	0x64, 0x4d, 0x61, 0x70, 0x12, 0x41, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67,
	0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x73, 0x76, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x08, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x1a, 0x3d, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0x49, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xfa, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x76, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x76, 0x63,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_store_store_proto_rawDescOnce sync.Once
	file_proto_store_store_proto_rawDescData = file_proto_store_store_proto_rawDesc
)

func file_proto_store_store_proto_rawDescGZIP() []byte {
	file_proto_store_store_proto_rawDescOnce.Do(func() {
		file_proto_store_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_store_store_proto_rawDescData)
	})
	return file_proto_store_store_proto_rawDescData
}

var file_proto_store_store_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_store_store_proto_goTypes = []interface{}{
	(*ErrorMsg)(nil),               // 0: go.micro.service.store_svc.ErrorMsg
	(*GetByObjectIdsRequest)(nil),  // 1: go.micro.service.store_svc.GetByObjectIdsRequest
	(*GetByObjectIdsResponse)(nil), // 2: go.micro.service.store_svc.GetByObjectIdsResponse
	(*ObjectInfo)(nil),             // 3: go.micro.service.store_svc.ObjectInfo
	(*SaveStoreInfoRequest)(nil),   // 4: go.micro.service.store_svc.SaveStoreInfoRequest
	(*StoreInfo)(nil),              // 5: go.micro.service.store_svc.StoreInfo
	(*SaveStoreInfoResponse)(nil),  // 6: go.micro.service.store_svc.SaveStoreInfoResponse
	nil,                            // 7: go.micro.service.store_svc.SaveStoreInfoResponse.Name2IdMapEntry
}
var file_proto_store_store_proto_depIdxs = []int32{
	3, // 0: go.micro.service.store_svc.GetByObjectIdsResponse.objectInfos:type_name -> go.micro.service.store_svc.ObjectInfo
	5, // 1: go.micro.service.store_svc.SaveStoreInfoRequest.storeInfos:type_name -> go.micro.service.store_svc.StoreInfo
	7, // 2: go.micro.service.store_svc.SaveStoreInfoResponse.name2IdMap:type_name -> go.micro.service.store_svc.SaveStoreInfoResponse.Name2IdMapEntry
	0, // 3: go.micro.service.store_svc.SaveStoreInfoResponse.errorMsg:type_name -> go.micro.service.store_svc.ErrorMsg
	4, // 4: go.micro.service.store_svc.Store.SaveStoreInfo:input_type -> go.micro.service.store_svc.SaveStoreInfoRequest
	1, // 5: go.micro.service.store_svc.Store.GetByObjectIds:input_type -> go.micro.service.store_svc.GetByObjectIdsRequest
	6, // 6: go.micro.service.store_svc.Store.SaveStoreInfo:output_type -> go.micro.service.store_svc.SaveStoreInfoResponse
	2, // 7: go.micro.service.store_svc.Store.GetByObjectIds:output_type -> go.micro.service.store_svc.GetByObjectIdsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_store_store_proto_init() }
func file_proto_store_store_proto_init() {
	if File_proto_store_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_store_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMsg); i {
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
		file_proto_store_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByObjectIdsRequest); i {
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
		file_proto_store_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByObjectIdsResponse); i {
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
		file_proto_store_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectInfo); i {
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
		file_proto_store_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveStoreInfoRequest); i {
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
		file_proto_store_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreInfo); i {
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
		file_proto_store_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveStoreInfoResponse); i {
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
			RawDescriptor: file_proto_store_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_store_store_proto_goTypes,
		DependencyIndexes: file_proto_store_store_proto_depIdxs,
		MessageInfos:      file_proto_store_store_proto_msgTypes,
	}.Build()
	File_proto_store_store_proto = out.File
	file_proto_store_store_proto_rawDesc = nil
	file_proto_store_store_proto_goTypes = nil
	file_proto_store_store_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: favorite_rpc.proto

package favorite_rpc

import (
	context "context"
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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	ServiceTime   int64  `protobuf:"varint,3,opt,name=service_time,json=serviceTime,proto3" json:"service_time,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *BaseResp) GetServiceTime() int64 {
	if x != nil {
		return x.ServiceTime
	}
	return 0
}

type RPCFavoriteCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RPCFavoriteCreateReq) Reset() {
	*x = RPCFavoriteCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFavoriteCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFavoriteCreateReq) ProtoMessage() {}

func (x *RPCFavoriteCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFavoriteCreateReq.ProtoReflect.Descriptor instead.
func (*RPCFavoriteCreateReq) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RPCFavoriteCreateReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *RPCFavoriteCreateReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RPCFavoriteCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *RPCFavoriteCreateResp) Reset() {
	*x = RPCFavoriteCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFavoriteCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFavoriteCreateResp) ProtoMessage() {}

func (x *RPCFavoriteCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFavoriteCreateResp.ProtoReflect.Descriptor instead.
func (*RPCFavoriteCreateResp) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *RPCFavoriteCreateResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type RPCFavoriteDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RPCFavoriteDelReq) Reset() {
	*x = RPCFavoriteDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFavoriteDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFavoriteDelReq) ProtoMessage() {}

func (x *RPCFavoriteDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFavoriteDelReq.ProtoReflect.Descriptor instead.
func (*RPCFavoriteDelReq) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *RPCFavoriteDelReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *RPCFavoriteDelReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RPCFavoriteDelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *RPCFavoriteDelResp) Reset() {
	*x = RPCFavoriteDelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFavoriteDelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFavoriteDelResp) ProtoMessage() {}

func (x *RPCFavoriteDelResp) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFavoriteDelResp.ProtoReflect.Descriptor instead.
func (*RPCFavoriteDelResp) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *RPCFavoriteDelResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type RPCFavoriteListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RPCFavoriteListReq) Reset() {
	*x = RPCFavoriteListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFavoriteListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFavoriteListReq) ProtoMessage() {}

func (x *RPCFavoriteListReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFavoriteListReq.ProtoReflect.Descriptor instead.
func (*RPCFavoriteListReq) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *RPCFavoriteListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RPCFavoriteListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp    *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	VideoIdList []int64   `protobuf:"varint,2,rep,packed,name=video_id_list,json=videoIdList,proto3" json:"video_id_list,omitempty"`
}

func (x *RPCFavoriteListResp) Reset() {
	*x = RPCFavoriteListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFavoriteListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFavoriteListResp) ProtoMessage() {}

func (x *RPCFavoriteListResp) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFavoriteListResp.ProtoReflect.Descriptor instead.
func (*RPCFavoriteListResp) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *RPCFavoriteListResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *RPCFavoriteListResp) GetVideoIdList() []int64 {
	if x != nil {
		return x.VideoIdList
	}
	return nil
}

type RPCFilterFavoriteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoIdList []int64 `protobuf:"varint,2,rep,packed,name=video_id_list,json=videoIdList,proto3" json:"video_id_list,omitempty"`
}

func (x *RPCFilterFavoriteReq) Reset() {
	*x = RPCFilterFavoriteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFilterFavoriteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFilterFavoriteReq) ProtoMessage() {}

func (x *RPCFilterFavoriteReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFilterFavoriteReq.ProtoReflect.Descriptor instead.
func (*RPCFilterFavoriteReq) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *RPCFilterFavoriteReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RPCFilterFavoriteReq) GetVideoIdList() []int64 {
	if x != nil {
		return x.VideoIdList
	}
	return nil
}

type RPCFilterFavoriteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp    *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	VideoIdList []int64   `protobuf:"varint,2,rep,packed,name=video_id_list,json=videoIdList,proto3" json:"video_id_list,omitempty"`
}

func (x *RPCFilterFavoriteResp) Reset() {
	*x = RPCFilterFavoriteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCFilterFavoriteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCFilterFavoriteResp) ProtoMessage() {}

func (x *RPCFilterFavoriteResp) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCFilterFavoriteResp.ProtoReflect.Descriptor instead.
func (*RPCFilterFavoriteResp) Descriptor() ([]byte, []int) {
	return file_favorite_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *RPCFilterFavoriteResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *RPCFilterFavoriteResp) GetVideoIdList() []int64 {
	if x != nil {
		return x.VideoIdList
	}
	return nil
}

var File_favorite_rpc_proto protoreflect.FileDescriptor

var file_favorite_rpc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22, 0x75,
	0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x52, 0x50, 0x43, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x48, 0x0a, 0x15, 0x52, 0x50, 0x43, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x47, 0x0a, 0x11, 0x52,
	0x50, 0x43, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x12, 0x52, 0x50, 0x43, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x0a, 0x12, 0x52,
	0x50, 0x43, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x13, 0x52, 0x50,
	0x43, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x22, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x14, 0x52, 0x50, 0x43, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x15, 0x52,
	0x50, 0x43, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xce, 0x02, 0x0a, 0x0f, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x1e, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x1f, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x48, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x1b, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0c, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x72, 0x72, 0x69, 0x63, 0x6b, 0x62,
	0x61, 0x6c, 0x6c, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_favorite_rpc_proto_rawDescOnce sync.Once
	file_favorite_rpc_proto_rawDescData = file_favorite_rpc_proto_rawDesc
)

func file_favorite_rpc_proto_rawDescGZIP() []byte {
	file_favorite_rpc_proto_rawDescOnce.Do(func() {
		file_favorite_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorite_rpc_proto_rawDescData)
	})
	return file_favorite_rpc_proto_rawDescData
}

var file_favorite_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_favorite_rpc_proto_goTypes = []interface{}{
	(*BaseResp)(nil),              // 0: favorite.BaseResp
	(*RPCFavoriteCreateReq)(nil),  // 1: favorite.RPCFavoriteCreateReq
	(*RPCFavoriteCreateResp)(nil), // 2: favorite.RPCFavoriteCreateResp
	(*RPCFavoriteDelReq)(nil),     // 3: favorite.RPCFavoriteDelReq
	(*RPCFavoriteDelResp)(nil),    // 4: favorite.RPCFavoriteDelResp
	(*RPCFavoriteListReq)(nil),    // 5: favorite.RPCFavoriteListReq
	(*RPCFavoriteListResp)(nil),   // 6: favorite.RPCFavoriteListResp
	(*RPCFilterFavoriteReq)(nil),  // 7: favorite.RPCFilterFavoriteReq
	(*RPCFilterFavoriteResp)(nil), // 8: favorite.RPCFilterFavoriteResp
}
var file_favorite_rpc_proto_depIdxs = []int32{
	0, // 0: favorite.RPCFavoriteCreateResp.base_resp:type_name -> favorite.BaseResp
	0, // 1: favorite.RPCFavoriteDelResp.base_resp:type_name -> favorite.BaseResp
	0, // 2: favorite.RPCFavoriteListResp.base_resp:type_name -> favorite.BaseResp
	0, // 3: favorite.RPCFilterFavoriteResp.base_resp:type_name -> favorite.BaseResp
	1, // 4: favorite.FavoriteService.CreateFavorite:input_type -> favorite.RPCFavoriteCreateReq
	3, // 5: favorite.FavoriteService.DelFavorite:input_type -> favorite.RPCFavoriteDelReq
	5, // 6: favorite.FavoriteService.FavoriteList:input_type -> favorite.RPCFavoriteListReq
	7, // 7: favorite.FavoriteService.FilterFavorite:input_type -> favorite.RPCFilterFavoriteReq
	2, // 8: favorite.FavoriteService.CreateFavorite:output_type -> favorite.RPCFavoriteCreateResp
	4, // 9: favorite.FavoriteService.DelFavorite:output_type -> favorite.RPCFavoriteDelResp
	6, // 10: favorite.FavoriteService.FavoriteList:output_type -> favorite.RPCFavoriteListResp
	8, // 11: favorite.FavoriteService.FilterFavorite:output_type -> favorite.RPCFilterFavoriteResp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_favorite_rpc_proto_init() }
func file_favorite_rpc_proto_init() {
	if File_favorite_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorite_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_favorite_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFavoriteCreateReq); i {
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
		file_favorite_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFavoriteCreateResp); i {
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
		file_favorite_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFavoriteDelReq); i {
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
		file_favorite_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFavoriteDelResp); i {
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
		file_favorite_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFavoriteListReq); i {
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
		file_favorite_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFavoriteListResp); i {
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
		file_favorite_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFilterFavoriteReq); i {
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
		file_favorite_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCFilterFavoriteResp); i {
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
			RawDescriptor: file_favorite_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorite_rpc_proto_goTypes,
		DependencyIndexes: file_favorite_rpc_proto_depIdxs,
		MessageInfos:      file_favorite_rpc_proto_msgTypes,
	}.Build()
	File_favorite_rpc_proto = out.File
	file_favorite_rpc_proto_rawDesc = nil
	file_favorite_rpc_proto_goTypes = nil
	file_favorite_rpc_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type FavoriteService interface {
	CreateFavorite(ctx context.Context, req *RPCFavoriteCreateReq) (res *RPCFavoriteCreateResp, err error)
	DelFavorite(ctx context.Context, req *RPCFavoriteDelReq) (res *RPCFavoriteDelResp, err error)
	FavoriteList(ctx context.Context, req *RPCFavoriteListReq) (res *RPCFavoriteListResp, err error)
	FilterFavorite(ctx context.Context, req *RPCFilterFavoriteReq) (res *RPCFilterFavoriteResp, err error)
}

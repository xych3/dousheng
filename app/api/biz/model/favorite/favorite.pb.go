// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: favorite.proto

package favorite

import (
	_ "gitee.com/derrickball/douyin/app/api/biz/model/api"
	common "gitee.com/derrickball/douyin/app/api/biz/model/common"
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

type FavoriteActionType int32

const (
	FavoriteActionType__        FavoriteActionType = 0
	FavoriteActionType_FAVORITE FavoriteActionType = 1
	FavoriteActionType_CANCEL   FavoriteActionType = 2
)

// Enum value maps for FavoriteActionType.
var (
	FavoriteActionType_name = map[int32]string{
		0: "_",
		1: "FAVORITE",
		2: "CANCEL",
	}
	FavoriteActionType_value = map[string]int32{
		"_":        0,
		"FAVORITE": 1,
		"CANCEL":   2,
	}
)

func (x FavoriteActionType) Enum() *FavoriteActionType {
	p := new(FavoriteActionType)
	*p = x
	return p
}

func (x FavoriteActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FavoriteActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_favorite_proto_enumTypes[0].Descriptor()
}

func (FavoriteActionType) Type() protoreflect.EnumType {
	return &file_favorite_proto_enumTypes[0]
}

func (x FavoriteActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *FavoriteActionType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = FavoriteActionType(num)
	return nil
}

// Deprecated: Use FavoriteActionType.Descriptor instead.
func (FavoriteActionType) EnumDescriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{0}
}

type FavoriteActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *string             `protobuf:"bytes,1,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                                                                           // 用户鉴权token
	VideoId    *int64              `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,required" form:"video_id,required" query:"video_id,required"`                                                 // 视频id
	ActionType *FavoriteActionType `protobuf:"varint,3,req,name=action_type,json=actionType,enum=favorite.FavoriteActionType" json:"action_type,required" form:"action_type,required" query:"action_type,required"` // 1-点赞，2-取消点赞
}

func (x *FavoriteActionReq) Reset() {
	*x = FavoriteActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionReq) ProtoMessage() {}

func (x *FavoriteActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionReq.ProtoReflect.Descriptor instead.
func (*FavoriteActionReq) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *FavoriteActionReq) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *FavoriteActionReq) GetVideoId() int64 {
	if x != nil && x.VideoId != nil {
		return *x.VideoId
	}
	return 0
}

func (x *FavoriteActionReq) GetActionType() FavoriteActionType {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return FavoriteActionType__
}

type FavoriteActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
}

func (x *FavoriteActionResp) Reset() {
	*x = FavoriteActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionResp) ProtoMessage() {}

func (x *FavoriteActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionResp.ProtoReflect.Descriptor instead.
func (*FavoriteActionResp) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{1}
}

func (x *FavoriteActionResp) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *FavoriteActionResp) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type FavoriteListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Toke   *string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}

func (x *FavoriteListReq) Reset() {
	*x = FavoriteListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListReq) ProtoMessage() {}

func (x *FavoriteListReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListReq.ProtoReflect.Descriptor instead.
func (*FavoriteListReq) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{2}
}

func (x *FavoriteListReq) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *FavoriteListReq) GetToke() string {
	if x != nil && x.Toke != nil {
		return *x.Toke
	}
	return ""
}

type FavoriteListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32          `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string         `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	VideoList  []*common.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list" form:"video_list" query:"video_list"`                                  // 用户点赞视频列表
}

func (x *FavoriteListResp) Reset() {
	*x = FavoriteListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListResp) ProtoMessage() {}

func (x *FavoriteListResp) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListResp.ProtoReflect.Descriptor instead.
func (*FavoriteListResp) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{3}
}

func (x *FavoriteListResp) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *FavoriteListResp) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *FavoriteListResp) GetVideoList() []*common.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_favorite_proto protoreflect.FileDescriptor

var file_favorite_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a, 0x12, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22,
	0x3e, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x6f, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x6b, 0x65, 0x22,
	0x80, 0x01, 0x0a, 0x10, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x2a, 0x35, 0x0a, 0x12, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x41, 0x56, 0x4f, 0x52, 0x49, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x02, 0x32, 0xcf, 0x01, 0x0a, 0x0f, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x61, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x1c, 0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x12, 0x59, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x1a, 0xca, 0xc1, 0x18, 0x16, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x42, 0x23, 0x5a, 0x21, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69,
	0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
}

var (
	file_favorite_proto_rawDescOnce sync.Once
	file_favorite_proto_rawDescData = file_favorite_proto_rawDesc
)

func file_favorite_proto_rawDescGZIP() []byte {
	file_favorite_proto_rawDescOnce.Do(func() {
		file_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorite_proto_rawDescData)
	})
	return file_favorite_proto_rawDescData
}

var file_favorite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_favorite_proto_goTypes = []interface{}{
	(FavoriteActionType)(0),    // 0: favorite.FavoriteActionType
	(*FavoriteActionReq)(nil),  // 1: favorite.FavoriteActionReq
	(*FavoriteActionResp)(nil), // 2: favorite.FavoriteActionResp
	(*FavoriteListReq)(nil),    // 3: favorite.FavoriteListReq
	(*FavoriteListResp)(nil),   // 4: favorite.FavoriteListResp
	(*common.Video)(nil),       // 5: common.Video
}
var file_favorite_proto_depIdxs = []int32{
	0, // 0: favorite.FavoriteActionReq.action_type:type_name -> favorite.FavoriteActionType
	5, // 1: favorite.FavoriteListResp.video_list:type_name -> common.Video
	1, // 2: favorite.FavoriteHandler.Action:input_type -> favorite.FavoriteActionReq
	3, // 3: favorite.FavoriteHandler.List:input_type -> favorite.FavoriteListReq
	2, // 4: favorite.FavoriteHandler.Action:output_type -> favorite.FavoriteActionResp
	4, // 5: favorite.FavoriteHandler.List:output_type -> favorite.FavoriteListResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_favorite_proto_init() }
func file_favorite_proto_init() {
	if File_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionReq); i {
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
		file_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionResp); i {
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
		file_favorite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListReq); i {
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
		file_favorite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListResp); i {
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
			RawDescriptor: file_favorite_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorite_proto_goTypes,
		DependencyIndexes: file_favorite_proto_depIdxs,
		EnumInfos:         file_favorite_proto_enumTypes,
		MessageInfos:      file_favorite_proto_msgTypes,
	}.Build()
	File_favorite_proto = out.File
	file_favorite_proto_rawDesc = nil
	file_favorite_proto_goTypes = nil
	file_favorite_proto_depIdxs = nil
}

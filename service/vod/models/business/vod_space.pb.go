// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: vod/business/vod_space.proto

package business

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

type VodSpaceUploadConfigKey int32

const (
	VodSpaceUploadConfigKey_UndefinedVodSpaceUploadConfigKey            VodSpaceUploadConfigKey = 0 // 未定义的key
	VodSpaceUploadConfigKey_CustomUploadFilePathVodSpaceUploadConfigKey VodSpaceUploadConfigKey = 1 // 自定义文件路径key
)

// Enum value maps for VodSpaceUploadConfigKey.
var (
	VodSpaceUploadConfigKey_name = map[int32]string{
		0: "UndefinedVodSpaceUploadConfigKey",
		1: "CustomUploadFilePathVodSpaceUploadConfigKey",
	}
	VodSpaceUploadConfigKey_value = map[string]int32{
		"UndefinedVodSpaceUploadConfigKey":            0,
		"CustomUploadFilePathVodSpaceUploadConfigKey": 1,
	}
)

func (x VodSpaceUploadConfigKey) Enum() *VodSpaceUploadConfigKey {
	p := new(VodSpaceUploadConfigKey)
	*p = x
	return p
}

func (x VodSpaceUploadConfigKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VodSpaceUploadConfigKey) Descriptor() protoreflect.EnumDescriptor {
	return file_vod_business_vod_space_proto_enumTypes[0].Descriptor()
}

func (VodSpaceUploadConfigKey) Type() protoreflect.EnumType {
	return &file_vod_business_vod_space_proto_enumTypes[0]
}

func (x VodSpaceUploadConfigKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VodSpaceUploadConfigKey.Descriptor instead.
func (VodSpaceUploadConfigKey) EnumDescriptor() ([]byte, []int) {
	return file_vod_business_vod_space_proto_rawDescGZIP(), []int{0}
}

type VodSpaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName string `protobuf:"bytes,1,opt,name=SpaceName,proto3" json:"SpaceName,omitempty"` // 空间名
	//  uint32 AccountID = 2;        // 账号ID
	Region         string `protobuf:"bytes,3,opt,name=Region,proto3" json:"Region,omitempty"`                  // 所属区域
	ProjectName    string `protobuf:"bytes,4,opt,name=ProjectName,proto3" json:"ProjectName,omitempty"`        // IAM项目名称
	BucketName     string `protobuf:"bytes,5,opt,name=BucketName,proto3" json:"BucketName,omitempty"`          // 绑定的TOS Bucket名称
	BucketStatus   string `protobuf:"bytes,6,opt,name=BucketStatus,proto3" json:"BucketStatus,omitempty"`      // 绑定的TosBucket状态
	Description    string `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`        // 描述
	UserName       string `protobuf:"bytes,8,opt,name=UserName,proto3" json:"UserName,omitempty"`              // 创建人
	CreatedAt      string `protobuf:"bytes,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`            // 创建时间
	Type           string `protobuf:"bytes,10,opt,name=Type,proto3" json:"Type,omitempty"`                     // 类型
	MediaSyncLevel string `protobuf:"bytes,11,opt,name=MediaSyncLevel,proto3" json:"MediaSyncLevel,omitempty"` // 媒资同步类型
}

func (x *VodSpaceInfo) Reset() {
	*x = VodSpaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodSpaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodSpaceInfo) ProtoMessage() {}

func (x *VodSpaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodSpaceInfo.ProtoReflect.Descriptor instead.
func (*VodSpaceInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_space_proto_rawDescGZIP(), []int{0}
}

func (x *VodSpaceInfo) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *VodSpaceInfo) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *VodSpaceInfo) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *VodSpaceInfo) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *VodSpaceInfo) GetBucketStatus() string {
	if x != nil {
		return x.BucketStatus
	}
	return ""
}

func (x *VodSpaceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VodSpaceInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *VodSpaceInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *VodSpaceInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VodSpaceInfo) GetMediaSyncLevel() string {
	if x != nil {
		return x.MediaSyncLevel
	}
	return ""
}

type VodSpaceConfigInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID  uint32                 `protobuf:"varint,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`  // TopAccountID
	InstanceID string                 `protobuf:"bytes,2,opt,name=InstanceID,proto3" json:"InstanceID,omitempty"` // 计费实例ID
	SpaceName  string                 `protobuf:"bytes,3,opt,name=SpaceName,proto3" json:"SpaceName,omitempty"`   // 空间名
	Status     string                 `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`         // 空间状态
	Configs    []*VodCommonConfigInfo `protobuf:"bytes,5,rep,name=Configs,proto3" json:"Configs,omitempty"`       // 空间配置项
}

func (x *VodSpaceConfigInfo) Reset() {
	*x = VodSpaceConfigInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodSpaceConfigInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodSpaceConfigInfo) ProtoMessage() {}

func (x *VodSpaceConfigInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodSpaceConfigInfo.ProtoReflect.Descriptor instead.
func (*VodSpaceConfigInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_space_proto_rawDescGZIP(), []int{1}
}

func (x *VodSpaceConfigInfo) GetAccountID() uint32 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

func (x *VodSpaceConfigInfo) GetInstanceID() string {
	if x != nil {
		return x.InstanceID
	}
	return ""
}

func (x *VodSpaceConfigInfo) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *VodSpaceConfigInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VodSpaceConfigInfo) GetConfigs() []*VodCommonConfigInfo {
	if x != nil {
		return x.Configs
	}
	return nil
}

var File_vod_business_vod_space_proto protoreflect.FileDescriptor

var file_vod_business_vod_space_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x6f, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x1d,
	0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02,
	0x0a, 0x0c, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0xd7, 0x01, 0x0a, 0x12, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x56, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2a, 0x70, 0x0a, 0x17,
	0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x6e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x10, 0x00, 0x12, 0x2f, 0x0a,
	0x2b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x10, 0x01, 0x42, 0xcc,
	0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x6f, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x08, 0x56, 0x6f,
	0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01,
	0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x20, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x23, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vod_business_vod_space_proto_rawDescOnce sync.Once
	file_vod_business_vod_space_proto_rawDescData = file_vod_business_vod_space_proto_rawDesc
)

func file_vod_business_vod_space_proto_rawDescGZIP() []byte {
	file_vod_business_vod_space_proto_rawDescOnce.Do(func() {
		file_vod_business_vod_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_vod_business_vod_space_proto_rawDescData)
	})
	return file_vod_business_vod_space_proto_rawDescData
}

var file_vod_business_vod_space_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vod_business_vod_space_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vod_business_vod_space_proto_goTypes = []interface{}{
	(VodSpaceUploadConfigKey)(0), // 0: Volcengine.Vod.Models.Business.VodSpaceUploadConfigKey
	(*VodSpaceInfo)(nil),         // 1: Volcengine.Vod.Models.Business.VodSpaceInfo
	(*VodSpaceConfigInfo)(nil),   // 2: Volcengine.Vod.Models.Business.VodSpaceConfigInfo
	(*VodCommonConfigInfo)(nil),  // 3: Volcengine.Vod.Models.Business.VodCommonConfigInfo
}
var file_vod_business_vod_space_proto_depIdxs = []int32{
	3, // 0: Volcengine.Vod.Models.Business.VodSpaceConfigInfo.Configs:type_name -> Volcengine.Vod.Models.Business.VodCommonConfigInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vod_business_vod_space_proto_init() }
func file_vod_business_vod_space_proto_init() {
	if File_vod_business_vod_space_proto != nil {
		return
	}
	file_vod_business_vod_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vod_business_vod_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodSpaceInfo); i {
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
		file_vod_business_vod_space_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodSpaceConfigInfo); i {
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
			RawDescriptor: file_vod_business_vod_space_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vod_business_vod_space_proto_goTypes,
		DependencyIndexes: file_vod_business_vod_space_proto_depIdxs,
		EnumInfos:         file_vod_business_vod_space_proto_enumTypes,
		MessageInfos:      file_vod_business_vod_space_proto_msgTypes,
	}.Build()
	File_vod_business_vod_space_proto = out.File
	file_vod_business_vod_space_proto_rawDesc = nil
	file_vod_business_vod_space_proto_goTypes = nil
	file_vod_business_vod_space_proto_depIdxs = nil
}
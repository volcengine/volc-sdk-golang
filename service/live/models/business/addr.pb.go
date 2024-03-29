// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: live/business/addr.proto

package business

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GeneratePlayURLResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URLList []*PlayURL `protobuf:"bytes,1,rep,name=URLList,proto3" json:"URLList,omitempty"` //播放地址详情列表
}

func (x *GeneratePlayURLResult) Reset() {
	*x = GeneratePlayURLResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_addr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePlayURLResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePlayURLResult) ProtoMessage() {}

func (x *GeneratePlayURLResult) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_addr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePlayURLResult.ProtoReflect.Descriptor instead.
func (*GeneratePlayURLResult) Descriptor() ([]byte, []int) {
	return file_live_business_addr_proto_rawDescGZIP(), []int{0}
}

func (x *GeneratePlayURLResult) GetURLList() []*PlayURL {
	if x != nil {
		return x.URLList
	}
	return nil
}

type PlayURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL      string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`           //播放地址
	CDN      string `protobuf:"bytes,2,opt,name=CDN,proto3" json:"CDN,omitempty"`           // CDN类型：fcdn、hw、ws、ali
	Protocol string `protobuf:"bytes,3,opt,name=Protocol,proto3" json:"Protocol,omitempty"` //协议类型：hls、rtmp、flv、udp
	Type     string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`         //地址类型，pull、push、3rd_play(relay_sink)、3rd_play(relay_source)
}

func (x *PlayURL) Reset() {
	*x = PlayURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_addr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayURL) ProtoMessage() {}

func (x *PlayURL) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_addr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayURL.ProtoReflect.Descriptor instead.
func (*PlayURL) Descriptor() ([]byte, []int) {
	return file_live_business_addr_proto_rawDescGZIP(), []int{1}
}

func (x *PlayURL) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *PlayURL) GetCDN() string {
	if x != nil {
		return x.CDN
	}
	return ""
}

func (x *PlayURL) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *PlayURL) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GeneratePushURLResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushURLList        []string       `protobuf:"bytes,1,rep,name=PushURLList,proto3" json:"PushURLList,omitempty"`               // rtmp 推流地址列表
	PushURLListDetail  []*PushURLItem `protobuf:"bytes,2,rep,name=PushURLListDetail,proto3" json:"PushURLListDetail,omitempty"`   // rtmp 推流地址列表详情，会针对OBS推流方式对地址进行分离
	TsOverSrtURLList   []string       `protobuf:"bytes,3,rep,name=TsOverSrtURLList,proto3" json:"TsOverSrtURLList,omitempty"`     // TS over SRT地址
	RtmpOverSrtURLList []string       `protobuf:"bytes,4,rep,name=RtmpOverSrtURLList,proto3" json:"RtmpOverSrtURLList,omitempty"` // RTMP over SRT地址
	RtmURLList         []string       `protobuf:"bytes,5,rep,name=RtmURLList,proto3" json:"RtmURLList,omitempty"`                 // RTM推流地址
}

func (x *GeneratePushURLResult) Reset() {
	*x = GeneratePushURLResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_addr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePushURLResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePushURLResult) ProtoMessage() {}

func (x *GeneratePushURLResult) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_addr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePushURLResult.ProtoReflect.Descriptor instead.
func (*GeneratePushURLResult) Descriptor() ([]byte, []int) {
	return file_live_business_addr_proto_rawDescGZIP(), []int{2}
}

func (x *GeneratePushURLResult) GetPushURLList() []string {
	if x != nil {
		return x.PushURLList
	}
	return nil
}

func (x *GeneratePushURLResult) GetPushURLListDetail() []*PushURLItem {
	if x != nil {
		return x.PushURLListDetail
	}
	return nil
}

func (x *GeneratePushURLResult) GetTsOverSrtURLList() []string {
	if x != nil {
		return x.TsOverSrtURLList
	}
	return nil
}

func (x *GeneratePushURLResult) GetRtmpOverSrtURLList() []string {
	if x != nil {
		return x.RtmpOverSrtURLList
	}
	return nil
}

func (x *GeneratePushURLResult) GetRtmURLList() []string {
	if x != nil {
		return x.RtmURLList
	}
	return nil
}

type PushURLItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL        string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`               //推流地址
	DomainApp  string `protobuf:"bytes,2,opt,name=DomainApp,proto3" json:"DomainApp,omitempty"`   // OBS推流服务器地址
	StreamSign string `protobuf:"bytes,3,opt,name=StreamSign,proto3" json:"StreamSign,omitempty"` // OBS串流密钥
}

func (x *PushURLItem) Reset() {
	*x = PushURLItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_addr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushURLItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushURLItem) ProtoMessage() {}

func (x *PushURLItem) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_addr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushURLItem.ProtoReflect.Descriptor instead.
func (*PushURLItem) Descriptor() ([]byte, []int) {
	return file_live_business_addr_proto_rawDescGZIP(), []int{3}
}

func (x *PushURLItem) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *PushURLItem) GetDomainApp() string {
	if x != nil {
		return x.DomainApp
	}
	return ""
}

func (x *PushURLItem) GetStreamSign() string {
	if x != nil {
		return x.StreamSign
	}
	return ""
}

var File_live_business_addr_proto protoreflect.FileDescriptor

var file_live_business_addr_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x56, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x5b, 0x0a, 0x15, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x52,
	0x07, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79,
	0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x44, 0x4e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x43, 0x44, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x50, 0x75, 0x73, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x55, 0x52, 0x4c, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x55, 0x52, 0x4c, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x55, 0x52, 0x4c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x11, 0x50, 0x75,
	0x73, 0x68, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x2a, 0x0a, 0x10, 0x54, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x53, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x54, 0x73, 0x4f, 0x76, 0x65,
	0x72, 0x53, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x52,
	0x74, 0x6d, 0x70, 0x4f, 0x76, 0x65, 0x72, 0x53, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x52, 0x74, 0x6d, 0x70, 0x4f, 0x76, 0x65,
	0x72, 0x53, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52,
	0x74, 0x6d, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x52, 0x74, 0x6d, 0x55, 0x52, 0x4c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x0b, 0x50,
	0x75, 0x73, 0x68, 0x55, 0x52, 0x4c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x1c, 0x0a, 0x09,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x42, 0xcc, 0x01, 0x0a, 0x2a, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x04, 0x41, 0x64, 0x64, 0x72, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f,
	0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02,
	0x21, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x4c, 0x69,
	0x76, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0xe2, 0x02, 0x24, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x4c, 0x69, 0x76, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_live_business_addr_proto_rawDescOnce sync.Once
	file_live_business_addr_proto_rawDescData = file_live_business_addr_proto_rawDesc
)

func file_live_business_addr_proto_rawDescGZIP() []byte {
	file_live_business_addr_proto_rawDescOnce.Do(func() {
		file_live_business_addr_proto_rawDescData = protoimpl.X.CompressGZIP(file_live_business_addr_proto_rawDescData)
	})
	return file_live_business_addr_proto_rawDescData
}

var file_live_business_addr_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_live_business_addr_proto_goTypes = []interface{}{
	(*GeneratePlayURLResult)(nil), // 0: Volcengine.Live.Models.Business.GeneratePlayURLResult
	(*PlayURL)(nil),               // 1: Volcengine.Live.Models.Business.PlayURL
	(*GeneratePushURLResult)(nil), // 2: Volcengine.Live.Models.Business.GeneratePushURLResult
	(*PushURLItem)(nil),           // 3: Volcengine.Live.Models.Business.PushURLItem
}
var file_live_business_addr_proto_depIdxs = []int32{
	1, // 0: Volcengine.Live.Models.Business.GeneratePlayURLResult.URLList:type_name -> Volcengine.Live.Models.Business.PlayURL
	3, // 1: Volcengine.Live.Models.Business.GeneratePushURLResult.PushURLListDetail:type_name -> Volcengine.Live.Models.Business.PushURLItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_live_business_addr_proto_init() }
func file_live_business_addr_proto_init() {
	if File_live_business_addr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_live_business_addr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePlayURLResult); i {
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
		file_live_business_addr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayURL); i {
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
		file_live_business_addr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePushURLResult); i {
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
		file_live_business_addr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushURLItem); i {
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
			RawDescriptor: file_live_business_addr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_live_business_addr_proto_goTypes,
		DependencyIndexes: file_live_business_addr_proto_depIdxs,
		MessageInfos:      file_live_business_addr_proto_msgTypes,
	}.Build()
	File_live_business_addr_proto = out.File
	file_live_business_addr_proto_rawDesc = nil
	file_live_business_addr_proto_goTypes = nil
	file_live_business_addr_proto_depIdxs = nil
}

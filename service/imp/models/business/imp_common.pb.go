// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: imp/business/imp_common.proto

package business

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

type InputPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"` // 可选：VOD TOS
	TosBucket    string `protobuf:"bytes,2,opt,name=TosBucket,proto3" json:"TosBucket,omitempty"`
	VodSpaceName string `protobuf:"bytes,3,opt,name=VodSpaceName,proto3" json:"VodSpaceName,omitempty"`
	FileId       string `protobuf:"bytes,4,opt,name=FileId,proto3" json:"FileId,omitempty"`
}

func (x *InputPath) Reset() {
	*x = InputPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputPath) ProtoMessage() {}

func (x *InputPath) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputPath.ProtoReflect.Descriptor instead.
func (*InputPath) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{0}
}

func (x *InputPath) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InputPath) GetTosBucket() string {
	if x != nil {
		return x.TosBucket
	}
	return ""
}

func (x *InputPath) GetVodSpaceName() string {
	if x != nil {
		return x.VodSpaceName
	}
	return ""
}

func (x *InputPath) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type JobOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId    string `protobuf:"bytes,1,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"` // 任务模板id
	Properties    string `protobuf:"bytes,2,opt,name=Properties,proto3" json:"Properties,omitempty"` // 任务输出info
	Code          string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	FileMessageId string `protobuf:"bytes,4,opt,name=FileMessageId,proto3" json:"FileMessageId,omitempty"`
	TaskType      string `protobuf:"bytes,5,opt,name=TaskType,proto3" json:"TaskType,omitempty"` // 任务类型
	Status        string `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`     // 任务状态
	ActivityId    string `protobuf:"bytes,7,opt,name=ActivityId,proto3" json:"ActivityId,omitempty"`
	StartTime     string `protobuf:"bytes,8,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime       string `protobuf:"bytes,9,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	TemplateName  string `protobuf:"bytes,10,opt,name=TemplateName,proto3" json:"TemplateName,omitempty"`
}

func (x *JobOutput) Reset() {
	*x = JobOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOutput) ProtoMessage() {}

func (x *JobOutput) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOutput.ProtoReflect.Descriptor instead.
func (*JobOutput) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{1}
}

func (x *JobOutput) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *JobOutput) GetProperties() string {
	if x != nil {
		return x.Properties
	}
	return ""
}

func (x *JobOutput) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *JobOutput) GetFileMessageId() string {
	if x != nil {
		return x.FileMessageId
	}
	return ""
}

func (x *JobOutput) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *JobOutput) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobOutput) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *JobOutput) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *JobOutput) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *JobOutput) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

type JobExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId             string       `protobuf:"bytes,1,opt,name=JobId,proto3" json:"JobId,omitempty"`         // 工作流执行id
	InputPath         *InputPath   `protobuf:"bytes,2,opt,name=InputPath,proto3" json:"InputPath,omitempty"` // 输入文件info
	Output            []*JobOutput `protobuf:"bytes,3,rep,name=Output,proto3" json:"Output,omitempty"`       // 任务节点info
	Status            string       `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`       // 工作流状态
	CreateAt          string       `protobuf:"bytes,5,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	FinishedAt        string       `protobuf:"bytes,6,opt,name=FinishedAt,proto3" json:"FinishedAt,omitempty"`
	TemplateId        string       `protobuf:"bytes,7,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"`               // 工作流模板id
	EnableLowPriority string       `protobuf:"bytes,8,opt,name=EnableLowPriority,proto3" json:"EnableLowPriority,omitempty"` // 是否开启闲时任务
	JobSource         string       `protobuf:"bytes,9,opt,name=JobSource,proto3" json:"JobSource,omitempty"`                 // 任务来源
}

func (x *JobExecution) Reset() {
	*x = JobExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobExecution) ProtoMessage() {}

func (x *JobExecution) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobExecution.ProtoReflect.Descriptor instead.
func (*JobExecution) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{2}
}

func (x *JobExecution) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobExecution) GetInputPath() *InputPath {
	if x != nil {
		return x.InputPath
	}
	return nil
}

func (x *JobExecution) GetOutput() []*JobOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *JobExecution) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobExecution) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *JobExecution) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

func (x *JobExecution) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *JobExecution) GetEnableLowPriority() string {
	if x != nil {
		return x.EnableLowPriority
	}
	return ""
}

func (x *JobExecution) GetJobSource() string {
	if x != nil {
		return x.JobSource
	}
	return ""
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverrideParams *OverrideParams `protobuf:"bytes,1,opt,name=OverrideParams,proto3" json:"OverrideParams,omitempty"` // 动态参数
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{3}
}

func (x *Params) GetOverrideParams() *OverrideParams {
	if x != nil {
		return x.OverrideParams
	}
	return nil
}

type OverrideParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartErase []*SmartEraseOverrideParams `protobuf:"bytes,1,rep,name=SmartErase,proto3" json:"SmartErase,omitempty"`
	Output     []*OutputOverrideParams     `protobuf:"bytes,2,rep,name=Output,proto3" json:"Output,omitempty"`
}

func (x *OverrideParams) Reset() {
	*x = OverrideParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverrideParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideParams) ProtoMessage() {}

func (x *OverrideParams) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideParams.ProtoReflect.Descriptor instead.
func (*OverrideParams) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{4}
}

func (x *OverrideParams) GetSmartErase() []*SmartEraseOverrideParams {
	if x != nil {
		return x.SmartErase
	}
	return nil
}

func (x *OverrideParams) GetOutput() []*OutputOverrideParams {
	if x != nil {
		return x.Output
	}
	return nil
}

type SmartEraseOverrideParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId []string   `protobuf:"bytes,1,rep,name=ActivityId,proto3" json:"ActivityId,omitempty"`
	Watermark  *Watermark `protobuf:"bytes,2,opt,name=Watermark,proto3" json:"Watermark,omitempty"`
	OCR        *OCR       `protobuf:"bytes,3,opt,name=OCR,proto3" json:"OCR,omitempty"`
}

func (x *SmartEraseOverrideParams) Reset() {
	*x = SmartEraseOverrideParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartEraseOverrideParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartEraseOverrideParams) ProtoMessage() {}

func (x *SmartEraseOverrideParams) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartEraseOverrideParams.ProtoReflect.Descriptor instead.
func (*SmartEraseOverrideParams) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{5}
}

func (x *SmartEraseOverrideParams) GetActivityId() []string {
	if x != nil {
		return x.ActivityId
	}
	return nil
}

func (x *SmartEraseOverrideParams) GetWatermark() *Watermark {
	if x != nil {
		return x.Watermark
	}
	return nil
}

func (x *SmartEraseOverrideParams) GetOCR() *OCR {
	if x != nil {
		return x.OCR
	}
	return nil
}

type Watermark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetectRect []*DetectRect `protobuf:"bytes,1,rep,name=DetectRect,proto3" json:"DetectRect,omitempty"`
}

func (x *Watermark) Reset() {
	*x = Watermark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watermark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watermark) ProtoMessage() {}

func (x *Watermark) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watermark.ProtoReflect.Descriptor instead.
func (*Watermark) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{6}
}

func (x *Watermark) GetDetectRect() []*DetectRect {
	if x != nil {
		return x.DetectRect
	}
	return nil
}

type OCR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetectRect []*DetectRect `protobuf:"bytes,1,rep,name=DetectRect,proto3" json:"DetectRect,omitempty"`
}

func (x *OCR) Reset() {
	*x = OCR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCR) ProtoMessage() {}

func (x *OCR) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCR.ProtoReflect.Descriptor instead.
func (*OCR) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{7}
}

func (x *OCR) GetDetectRect() []*DetectRect {
	if x != nil {
		return x.DetectRect
	}
	return nil
}

type DetectRect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X1 float64 `protobuf:"fixed64,1,opt,name=X1,proto3" json:"X1,omitempty"`
	Y1 float64 `protobuf:"fixed64,2,opt,name=Y1,proto3" json:"Y1,omitempty"`
	X2 float64 `protobuf:"fixed64,3,opt,name=X2,proto3" json:"X2,omitempty"`
	Y2 float64 `protobuf:"fixed64,4,opt,name=Y2,proto3" json:"Y2,omitempty"`
}

func (x *DetectRect) Reset() {
	*x = DetectRect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectRect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectRect) ProtoMessage() {}

func (x *DetectRect) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectRect.ProtoReflect.Descriptor instead.
func (*DetectRect) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{8}
}

func (x *DetectRect) GetX1() float64 {
	if x != nil {
		return x.X1
	}
	return 0
}

func (x *DetectRect) GetY1() float64 {
	if x != nil {
		return x.Y1
	}
	return 0
}

func (x *DetectRect) GetX2() float64 {
	if x != nil {
		return x.X2
	}
	return 0
}

func (x *DetectRect) GetY2() float64 {
	if x != nil {
		return x.Y2
	}
	return 0
}

type OutputOverrideParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId []string    `protobuf:"bytes,1,rep,name=ActivityId,proto3" json:"ActivityId,omitempty"`
	OutputPath *OutputPath `protobuf:"bytes,2,opt,name=OutputPath,proto3" json:"OutputPath,omitempty"`
}

func (x *OutputOverrideParams) Reset() {
	*x = OutputOverrideParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputOverrideParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputOverrideParams) ProtoMessage() {}

func (x *OutputOverrideParams) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputOverrideParams.ProtoReflect.Descriptor instead.
func (*OutputOverrideParams) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{9}
}

func (x *OutputOverrideParams) GetActivityId() []string {
	if x != nil {
		return x.ActivityId
	}
	return nil
}

func (x *OutputOverrideParams) GetOutputPath() *OutputPath {
	if x != nil {
		return x.OutputPath
	}
	return nil
}

type OutputPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	TosBucket    string `protobuf:"bytes,2,opt,name=TosBucket,proto3" json:"TosBucket,omitempty"`
	VodSpaceName string `protobuf:"bytes,3,opt,name=VodSpaceName,proto3" json:"VodSpaceName,omitempty"`
	FileName     string `protobuf:"bytes,4,opt,name=FileName,proto3" json:"FileName,omitempty"`
}

func (x *OutputPath) Reset() {
	*x = OutputPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imp_business_imp_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputPath) ProtoMessage() {}

func (x *OutputPath) ProtoReflect() protoreflect.Message {
	mi := &file_imp_business_imp_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputPath.ProtoReflect.Descriptor instead.
func (*OutputPath) Descriptor() ([]byte, []int) {
	return file_imp_business_imp_common_proto_rawDescGZIP(), []int{10}
}

func (x *OutputPath) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OutputPath) GetTosBucket() string {
	if x != nil {
		return x.TosBucket
	}
	return ""
}

func (x *OutputPath) GetVodSpaceName() string {
	if x != nil {
		return x.VodSpaceName
	}
	return ""
}

func (x *OutputPath) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_imp_business_imp_common_proto protoreflect.FileDescriptor

var file_imp_business_imp_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6d, 0x70, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x69,
	0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22,
	0x79, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xb5, 0x02, 0x0a, 0x09, 0x4a,
	0x6f, 0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xf0, 0x02, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x09, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x56,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x41, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x77, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4a, 0x6f, 0x62, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x60, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x56, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x58, 0x0a, 0x0a, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x45, 0x72, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x72, 0x61, 0x73, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45,
	0x72, 0x61, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0xba, 0x01, 0x0a, 0x18, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x72, 0x61, 0x73,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x47, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x09, 0x57,
	0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x35, 0x0a, 0x03, 0x4f, 0x43, 0x52, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x43, 0x52, 0x52, 0x03, 0x4f, 0x43, 0x52, 0x22,
	0x57, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x4a, 0x0a, 0x0a,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6d,
	0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x74, 0x52, 0x0a, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x74, 0x22, 0x51, 0x0a, 0x03, 0x4f, 0x43, 0x52, 0x12,
	0x4a, 0x0a, 0x0a, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x74, 0x52,
	0x0a, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x74, 0x22, 0x4c, 0x0a, 0x0a, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x58, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x58, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x59, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x59, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x58, 0x32, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x58, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x59, 0x32, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x59, 0x32, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x4a, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x0a, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x7e,
	0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x56, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xcc,
	0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x6d, 0x70, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x49, 0x6d,
	0x70, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6d, 0x70, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01,
	0x01, 0xd8, 0x01, 0x01, 0xca, 0x02, 0x20, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x49, 0x6d, 0x70, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x23, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x49, 0x6d, 0x70, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imp_business_imp_common_proto_rawDescOnce sync.Once
	file_imp_business_imp_common_proto_rawDescData = file_imp_business_imp_common_proto_rawDesc
)

func file_imp_business_imp_common_proto_rawDescGZIP() []byte {
	file_imp_business_imp_common_proto_rawDescOnce.Do(func() {
		file_imp_business_imp_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_imp_business_imp_common_proto_rawDescData)
	})
	return file_imp_business_imp_common_proto_rawDescData
}

var file_imp_business_imp_common_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_imp_business_imp_common_proto_goTypes = []interface{}{
	(*InputPath)(nil),                // 0: Volcengine.Imp.Models.Business.InputPath
	(*JobOutput)(nil),                // 1: Volcengine.Imp.Models.Business.JobOutput
	(*JobExecution)(nil),             // 2: Volcengine.Imp.Models.Business.JobExecution
	(*Params)(nil),                   // 3: Volcengine.Imp.Models.Business.Params
	(*OverrideParams)(nil),           // 4: Volcengine.Imp.Models.Business.OverrideParams
	(*SmartEraseOverrideParams)(nil), // 5: Volcengine.Imp.Models.Business.SmartEraseOverrideParams
	(*Watermark)(nil),                // 6: Volcengine.Imp.Models.Business.Watermark
	(*OCR)(nil),                      // 7: Volcengine.Imp.Models.Business.OCR
	(*DetectRect)(nil),               // 8: Volcengine.Imp.Models.Business.DetectRect
	(*OutputOverrideParams)(nil),     // 9: Volcengine.Imp.Models.Business.OutputOverrideParams
	(*OutputPath)(nil),               // 10: Volcengine.Imp.Models.Business.OutputPath
}
var file_imp_business_imp_common_proto_depIdxs = []int32{
	0,  // 0: Volcengine.Imp.Models.Business.JobExecution.InputPath:type_name -> Volcengine.Imp.Models.Business.InputPath
	1,  // 1: Volcengine.Imp.Models.Business.JobExecution.Output:type_name -> Volcengine.Imp.Models.Business.JobOutput
	4,  // 2: Volcengine.Imp.Models.Business.Params.OverrideParams:type_name -> Volcengine.Imp.Models.Business.OverrideParams
	5,  // 3: Volcengine.Imp.Models.Business.OverrideParams.SmartErase:type_name -> Volcengine.Imp.Models.Business.SmartEraseOverrideParams
	9,  // 4: Volcengine.Imp.Models.Business.OverrideParams.Output:type_name -> Volcengine.Imp.Models.Business.OutputOverrideParams
	6,  // 5: Volcengine.Imp.Models.Business.SmartEraseOverrideParams.Watermark:type_name -> Volcengine.Imp.Models.Business.Watermark
	7,  // 6: Volcengine.Imp.Models.Business.SmartEraseOverrideParams.OCR:type_name -> Volcengine.Imp.Models.Business.OCR
	8,  // 7: Volcengine.Imp.Models.Business.Watermark.DetectRect:type_name -> Volcengine.Imp.Models.Business.DetectRect
	8,  // 8: Volcengine.Imp.Models.Business.OCR.DetectRect:type_name -> Volcengine.Imp.Models.Business.DetectRect
	10, // 9: Volcengine.Imp.Models.Business.OutputOverrideParams.OutputPath:type_name -> Volcengine.Imp.Models.Business.OutputPath
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_imp_business_imp_common_proto_init() }
func file_imp_business_imp_common_proto_init() {
	if File_imp_business_imp_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imp_business_imp_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputPath); i {
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
		file_imp_business_imp_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOutput); i {
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
		file_imp_business_imp_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobExecution); i {
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
		file_imp_business_imp_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_imp_business_imp_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverrideParams); i {
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
		file_imp_business_imp_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartEraseOverrideParams); i {
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
		file_imp_business_imp_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watermark); i {
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
		file_imp_business_imp_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCR); i {
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
		file_imp_business_imp_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectRect); i {
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
		file_imp_business_imp_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputOverrideParams); i {
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
		file_imp_business_imp_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputPath); i {
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
			RawDescriptor: file_imp_business_imp_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_imp_business_imp_common_proto_goTypes,
		DependencyIndexes: file_imp_business_imp_common_proto_depIdxs,
		MessageInfos:      file_imp_business_imp_common_proto_msgTypes,
	}.Build()
	File_imp_business_imp_common_proto = out.File
	file_imp_business_imp_common_proto_rawDesc = nil
	file_imp_business_imp_common_proto_goTypes = nil
	file_imp_business_imp_common_proto_depIdxs = nil
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log.proto

package pb

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LogContent struct {
	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *LogContent) Reset()         { *m = LogContent{} }
func (m *LogContent) String() string { return proto.CompactTextString(m) }
func (*LogContent) ProtoMessage()    {}
func (*LogContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}
func (m *LogContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogContent.Merge(m, src)
}
func (m *LogContent) XXX_Size() int {
	return m.Size()
}
func (m *LogContent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogContent.DiscardUnknown(m)
}

var xxx_messageInfo_LogContent proto.InternalMessageInfo

func (m *LogContent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LogContent) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Log struct {
	Time     int64         `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	Contents []*LogContent `protobuf:"bytes,2,rep,name=Contents,proto3" json:"Contents,omitempty"`
	// Types that are valid to be assigned to OptionalTimeNs:
	//
	//	*Log_TimeNs
	OptionalTimeNs isLog_OptionalTimeNs `protobuf_oneof:"OptionalTimeNs"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{1}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Log.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return m.Size()
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

type isLog_OptionalTimeNs interface {
	isLog_OptionalTimeNs()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Log_TimeNs struct {
	TimeNs uint32 `protobuf:"fixed32,3,opt,name=TimeNs,proto3,oneof" json:"TimeNs,omitempty"`
}

func (*Log_TimeNs) isLog_OptionalTimeNs() {}

func (m *Log) GetOptionalTimeNs() isLog_OptionalTimeNs {
	if m != nil {
		return m.OptionalTimeNs
	}
	return nil
}

func (m *Log) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetContents() []*LogContent {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Log) GetTimeNs() uint32 {
	if x, ok := m.GetOptionalTimeNs().(*Log_TimeNs); ok {
		return x.TimeNs
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Log) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Log_TimeNs)(nil),
	}
}

type LogTag struct {
	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *LogTag) Reset()         { *m = LogTag{} }
func (m *LogTag) String() string { return proto.CompactTextString(m) }
func (*LogTag) ProtoMessage()    {}
func (*LogTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{2}
}
func (m *LogTag) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogTag.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogTag.Merge(m, src)
}
func (m *LogTag) XXX_Size() int {
	return m.Size()
}
func (m *LogTag) XXX_DiscardUnknown() {
	xxx_messageInfo_LogTag.DiscardUnknown(m)
}

var xxx_messageInfo_LogTag proto.InternalMessageInfo

func (m *LogTag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LogTag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type LogGroup struct {
	Logs        []*Log    `protobuf:"bytes,1,rep,name=Logs,proto3" json:"Logs,omitempty"`
	Source      string    `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"`
	LogTags     []*LogTag `protobuf:"bytes,3,rep,name=LogTags,proto3" json:"LogTags,omitempty"`
	FileName    string    `protobuf:"bytes,4,opt,name=FileName,proto3" json:"FileName,omitempty"`
	ContextFlow string    `protobuf:"bytes,5,opt,name=ContextFlow,proto3" json:"ContextFlow,omitempty"`
}

func (m *LogGroup) Reset()         { *m = LogGroup{} }
func (m *LogGroup) String() string { return proto.CompactTextString(m) }
func (*LogGroup) ProtoMessage()    {}
func (*LogGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{3}
}
func (m *LogGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogGroup.Merge(m, src)
}
func (m *LogGroup) XXX_Size() int {
	return m.Size()
}
func (m *LogGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_LogGroup.DiscardUnknown(m)
}

var xxx_messageInfo_LogGroup proto.InternalMessageInfo

func (m *LogGroup) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *LogGroup) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *LogGroup) GetLogTags() []*LogTag {
	if m != nil {
		return m.LogTags
	}
	return nil
}

func (m *LogGroup) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *LogGroup) GetContextFlow() string {
	if m != nil {
		return m.ContextFlow
	}
	return ""
}

type LogGroupList struct {
	LogGroups []*LogGroup `protobuf:"bytes,1,rep,name=LogGroups,proto3" json:"LogGroups,omitempty"`
}

func (m *LogGroupList) Reset()         { *m = LogGroupList{} }
func (m *LogGroupList) String() string { return proto.CompactTextString(m) }
func (*LogGroupList) ProtoMessage()    {}
func (*LogGroupList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{4}
}
func (m *LogGroupList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogGroupList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogGroupList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogGroupList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogGroupList.Merge(m, src)
}
func (m *LogGroupList) XXX_Size() int {
	return m.Size()
}
func (m *LogGroupList) XXX_DiscardUnknown() {
	xxx_messageInfo_LogGroupList.DiscardUnknown(m)
}

var xxx_messageInfo_LogGroupList proto.InternalMessageInfo

func (m *LogGroupList) GetLogGroups() []*LogGroup {
	if m != nil {
		return m.LogGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*LogContent)(nil), "pb.LogContent")
	proto.RegisterType((*Log)(nil), "pb.Log")
	proto.RegisterType((*LogTag)(nil), "pb.LogTag")
	proto.RegisterType((*LogGroup)(nil), "pb.LogGroup")
	proto.RegisterType((*LogGroupList)(nil), "pb.LogGroupList")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_a153da538f858886) }

var fileDescriptor_a153da538f858886 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x4d, 0x9b, 0xb6, 0xd3, 0x52, 0xca, 0x20, 0xb2, 0x28, 0xc4, 0x10, 0x3c, 0x94,
	0x1e, 0xa2, 0xa8, 0x27, 0xbd, 0x55, 0xa8, 0x82, 0xa1, 0xc2, 0x5a, 0x3c, 0x78, 0x4b, 0x24, 0x2c,
	0x81, 0xb4, 0x13, 0x9b, 0x14, 0xf5, 0x2d, 0x7c, 0x07, 0x5f, 0xc6, 0x63, 0x8f, 0x1e, 0xa5, 0x7d,
	0x11, 0xc9, 0x66, 0xd3, 0x7a, 0xf4, 0x36, 0xff, 0xff, 0xe7, 0x9f, 0xf9, 0xc2, 0x42, 0x3b, 0x21,
	0xe9, 0xa5, 0x0b, 0xca, 0x09, 0x6b, 0x69, 0xe8, 0x5e, 0x00, 0xf8, 0x24, 0xaf, 0x69, 0x9e, 0x47,
	0xf3, 0x1c, 0xfb, 0x60, 0xde, 0x45, 0xef, 0x9c, 0x39, 0x6c, 0xd0, 0x16, 0xc5, 0x88, 0x7b, 0xd0,
	0x78, 0x0c, 0x92, 0x65, 0xc4, 0x6b, 0xca, 0x2b, 0x85, 0xfb, 0x02, 0xa6, 0x4f, 0x12, 0x11, 0xea,
	0xd3, 0x78, 0x16, 0xa9, 0xef, 0x4d, 0xa1, 0x66, 0x1c, 0x42, 0x4b, 0x6f, 0xcb, 0x78, 0xcd, 0x31,
	0x07, 0x9d, 0xb3, 0x9e, 0x97, 0x86, 0xde, 0xee, 0x88, 0xd8, 0xe6, 0xc8, 0xc1, 0x2a, 0x3a, 0x93,
	0x8c, 0x9b, 0x0e, 0x1b, 0x34, 0x6f, 0x0d, 0xa1, 0xf5, 0xa8, 0x0f, 0xbd, 0xfb, 0x34, 0x8f, 0x69,
	0x1e, 0x24, 0xa5, 0xe3, 0x9e, 0x82, 0xe5, 0x93, 0x9c, 0x06, 0xf2, 0xdf, 0x90, 0x9f, 0x0c, 0x5a,
	0x3e, 0xc9, 0x9b, 0x05, 0x2d, 0x53, 0x3c, 0x84, 0xba, 0x4f, 0x32, 0xe3, 0x4c, 0x21, 0x35, 0x35,
	0x92, 0x50, 0x26, 0xee, 0x83, 0xf5, 0x40, 0xcb, 0xc5, 0x73, 0xb5, 0x40, 0x2b, 0x3c, 0x86, 0x66,
	0x79, 0xb3, 0x00, 0x2c, 0x7a, 0xa0, 0x7b, 0xd3, 0x40, 0x8a, 0x2a, 0xc2, 0x03, 0x68, 0x8d, 0xe3,
	0x24, 0x9a, 0x04, 0xb3, 0x88, 0xd7, 0x55, 0x7f, 0xab, 0xd1, 0x81, 0x8e, 0xfa, 0xdb, 0xb7, 0x7c,
	0x9c, 0xd0, 0x2b, 0x6f, 0xa8, 0xf8, 0xaf, 0xe5, 0x5e, 0x42, 0xb7, 0x82, 0xf4, 0xe3, 0x2c, 0xc7,
	0x21, 0xb4, 0x2b, 0x5d, 0xd1, 0x76, 0xf5, 0x55, 0x65, 0x8a, 0x5d, 0x3c, 0x3a, 0xfa, 0x5a, 0xdb,
	0x6c, 0xb5, 0xb6, 0xd9, 0xcf, 0xda, 0x66, 0x1f, 0x1b, 0xdb, 0x58, 0x6d, 0x6c, 0xe3, 0x7b, 0x63,
	0x1b, 0x4f, 0x0d, 0xef, 0xe4, 0x2a, 0x0d, 0x43, 0x4b, 0x3d, 0xf4, 0xf9, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x96, 0xe6, 0xb7, 0xdc, 0xf5, 0x01, 0x00, 0x00,
}

func (m *LogContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OptionalTimeNs != nil {
		{
			size := m.OptionalTimeNs.Size()
			i -= size
			if _, err := m.OptionalTimeNs.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Contents) > 0 {
		for iNdEx := len(m.Contents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLog(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Time != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Log_TimeNs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log_TimeNs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.TimeNs))
	i--
	dAtA[i] = 0x1d
	return len(dAtA) - i, nil
}
func (m *LogTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogTag) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogTag) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LogGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContextFlow) > 0 {
		i -= len(m.ContextFlow)
		copy(dAtA[i:], m.ContextFlow)
		i = encodeVarintLog(dAtA, i, uint64(len(m.ContextFlow)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FileName) > 0 {
		i -= len(m.FileName)
		copy(dAtA[i:], m.FileName)
		i = encodeVarintLog(dAtA, i, uint64(len(m.FileName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LogTags) > 0 {
		for iNdEx := len(m.LogTags) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LogTags[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLog(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLog(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LogGroupList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogGroupList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogGroupList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LogGroups) > 0 {
		for iNdEx := len(m.LogGroups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LogGroups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLog(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	offset -= sovLog(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LogContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	return n
}

func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovLog(uint64(m.Time))
	}
	if len(m.Contents) > 0 {
		for _, e := range m.Contents {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.OptionalTimeNs != nil {
		n += m.OptionalTimeNs.Size()
	}
	return n
}

func (m *Log_TimeNs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 5
	return n
}
func (m *LogTag) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	return n
}

func (m *LogGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if len(m.LogTags) > 0 {
		for _, e := range m.LogTags {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	l = len(m.FileName)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.ContextFlow)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	return n
}

func (m *LogGroupList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LogGroups) > 0 {
		for _, e := range m.LogGroups {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	return n
}

func sovLog(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LogContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents, &LogContent{})
			if err := m.Contents[len(m.Contents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeNs", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalTimeNs = &Log_TimeNs{v}
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogTag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogTags = append(m.LogTags, &LogTag{})
			if err := m.LogTags[len(m.LogTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContextFlow", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContextFlow = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogGroupList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogGroupList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogGroupList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogGroups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogGroups = append(m.LogGroups, &LogGroup{})
			if err := m.LogGroups[len(m.LogGroups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLog
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLog
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLog
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLog        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLog = fmt.Errorf("proto: unexpected end of group")
)

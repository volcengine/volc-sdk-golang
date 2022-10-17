package notify

import (
	"fmt"
	"strings"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

type JsonTime time.Time

//实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format(TIME_LAYOUT))
	return []byte(stamp), nil
}

func (this JsonTime) UnmarshalJSON(bytes []byte) error {
	location, _ := time.LoadLocation("Asia/Shanghai")

	str := string(bytes)
	str = strings.Trim(str, "\"")
	t, err := time.ParseInLocation(TIME_LAYOUT, str, location)
	if err == nil {
		this = JsonTime(t)
	}
	return err
}

type PhoneParam struct {
	Phone      string
	PhoneParam map[string]interface{}
	TtsContent string
	Ext        string
	SmsParam   map[string]interface{}
}

type ForbidTimeItem struct {
	//星期几，不填代表每天都生效,1表示周日，2表示周一，以此类推
	Weekdays []int32 `json:",omitempty"`

	//时间段，至少应该有一个时间段,example:["09:00-10:00","15:00-16:00"]
	Times []string
}

type CreateTaskRequest struct {
	Name            string
	PhoneList       []*PhoneParam
	Resource        string
	StartTime       JsonTime
	EndTime         JsonTime
	Concurrency     int32
	Start           bool  `json:",omitempty"`
	MaxRingDuration int32 `json:",omitempty"`
	//max 3
	RingAgainTimes int32 `json:",omitempty"`
	//min 5
	RingAgainInterval int32 `json:",omitempty"`
	Unique            bool
	ForbidTimeList    []*ForbidTimeItem `json:",omitempty"`
	NumberPoolNo      string
	NumberList        []string `json:",omitempty"`
	SelectNumberType  int32
	SelectNumberRule  int32
	Type              int32
	FinishWhenListEnd bool `json:",omitempty"`
}

type FailItem struct {
	Index int32
	Phone string
	Type  string
}

type TaskAppendResult struct {
	TaskOpenId string
	FailList   []*FailItem
}

type TaskAppendResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *TaskAppendResult
}

type CommonResponse struct {
	ResponseMetadata *base.ResponseMetadata
}

type BatchAppendRequest struct {
	TaskOpenId string
	PhoneList  []*PhoneParam
}

type CreateTtsResourceRequest struct {
	Name               string
	TtsTemplateContent string
	Remark             string
}

type EditTaskRequest struct {
	TaskOpenId  string
	StartTime   JsonTime
	EndTime     JsonTime
	Recall      bool
	Concurrency int32
	//max 3
	RingAgainTimes int32
	//min 5
	RingAgainInterval int32
	ForbidTimeList    []*ForbidTimeItem `json:",omitempty"`
}

type FetchVoiceResourceRequest struct {
	Url  string
	Name string
}

type SingleParam struct {
	SingleOpenId string
	Phone        string
	Resource     string
	TriggerTime  JsonTime `json:",omitempty"`
	//max 3
	RingAgainTimes int32 `json:",omitempty"`
	//min 5
	RingAgainInterval int32 `json:",omitempty"`
	PhoneParam        map[string]interface{}
	TtsContent        string
	Ext               string
	NumberPoolNo      string
	NumberList        []string `json:",omitempty"`
	NumberType        int32
	Type              int32
	SelectNumberRule  int32             `json:",omitempty"`
	ForbidTimeList    []*ForbidTimeItem `json:",omitempty"`
}

type SingleAppendRequest struct {
	List []*SingleParam
}

type UploadVoiceResourceRequest struct {
	FileName string
}

type BasicResourceResult struct {
	ResourceKey string
	Suffix      string
	Name        string
	UploadUrl   string
}

type SingleResultItem struct {
	Phone        string
	SingleOpenId string
	Type         string
}

type SingleInfo struct {
	SingleOpenId string
	Resource     string
	Phone        string
	Ext          string
	CreateTime   JsonTime
	State        string
	ReleaseTime  JsonTime
	Duration     int32
}

type SingleInfoResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *SingleInfo
}

type SingleAppendResult struct {
	SuccessList []*SingleResultItem
	FailList    []*SingleResultItem
}

type SingleAppendResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *SingleAppendResult
}

type BasicResourceResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BasicResourceResult
}

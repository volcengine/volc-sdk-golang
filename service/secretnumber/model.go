package secretnumber

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type BindAXBRequest struct {
	PhoneNoA        string
	PhoneNoB        string
	PhoneNoX        string
	NumberPoolNo    string
	ExpireTime      int64
	AudioRecordFlag int32
	UserData        string
}

type SelectNumberAndBindAXBRequest struct {
	PhoneNoA          string
	PhoneNoB          string
	NumberPoolNo      string
	ExpireTime        int64
	AudioRecordFlag   int32
	CityCode          string
	CityCodeByPhoneNo string
	DegradeCityList   string
	UserData          string
}

type SecretBindResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SecretBindResult
}

type SecretBindResult struct {
	PhoneNoX string
	SubId    string
	Status   int32
}

type OperationResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           bool
}

type SpecificSubIdRequest struct {
	NumberPoolNo string
	SubId        string
}

type QuerySubscriptionResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           Subscription
}

type QuerySubscriptionForListRequest struct {
	NumberPoolNo string
	PhoneNoX     string
	PhoneNoA     string
	PhoneNoB     string
	Status       int32
	SubId        string
	Offset       int32
	Limit        int32
}

type QuerySubscriptionForListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SubscriptionList
}

type SubscriptionList struct {
	Records []Subscription
	Offset  int32
	Limit   int32
	Total   int32
}

type Subscription struct {
	SubId      string
	PhoneNoA   string
	PhoneNoB   string
	PhoneNoX   string
	Status     int32
	RecordFlag int32
	EnableTime int64
	ExpireTime int64
}

type UpgradeAXToAXBRequest struct {
	NumberPoolNo string
	SubId        string
	PhoneNoB     string
	UserData     string
}

type UpdateAXBRequest struct {
	UpdateType   string
	NumberPoolNo string
	SubId        string
	ExpireTime   int64
	PhoneNoB     string
}

type BindAXNRequest struct {
	PhoneNoA        string
	PhoneNoB        string
	PhoneNoX        string
	NumberPoolNo    string
	ExpireTime      int64
	AudioRecordFlag int32
	UserData        string
}

type UpdateAXNRequest struct {
	UpdateType   string
	NumberPoolNo string
	SubId        string
	ExpireTime   int64
	PhoneNoB     string
}

type Click2CallRequest struct {
	Caller                      string
	CallerNumber                string
	CallerNumberPoolNo          string
	CallerNumberCityCode        string
	CallerNumberDegradeCityList string
	Callee                      string
	CalleeNumber                string
	CalleeNumberPoolNo          string
	CalleeNumberCityCode        string
	CalleeNumberDegradeCityList string
	MaxTime                     int32
	PreVoice                    string
	PreVoicePlay                bool
	LastMinutes                 int32
	LastVoice                   string
	LastVoiceTo                 string
	UserData                    string
	CityCodeByPhoneNo           string
}

type Click2CallResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           Click2CallResult
}

type Click2CallResult struct {
	CallId string
}

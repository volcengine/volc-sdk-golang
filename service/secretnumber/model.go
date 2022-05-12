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
	UserData     string
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

type SelectNumberAndBindAXNRequest struct {
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

type UpdateAXNRequest struct {
	UpdateType   string
	NumberPoolNo string
	SubId        string
	ExpireTime   int64
	PhoneNoB     string
	PhoneNoA     string
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

type QueryAudioRecordFileUrlRequest struct {
	CallId string
}

type QueryAudioRecordFileUrlData struct {
	AudioRecordFileUrl      string
	AudioRecordLeftFileUrl  string
	AudioRecordRightFileUrl string
}

type QueryAudioRecordFileUrlResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           QueryAudioRecordFileUrlData
}

type QueryAudioRecordToTextFileRequest struct {
	CallIdList string
}

type QueryAudioRecordToTextFile struct {
	CallId                   string
	AudioRecordToTextFileUrl string
}

type QueryAudioRecordToTextFileResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []QueryAudioRecordToTextFile
}

type QueryCallRecordMsgRequest struct {
	CallIdList string
}

type QueryCallRecordMsg struct {
	AccountId              string
	CallId                 string
	ServiceType            int32
	SubServiceType         int32
	Caller                 string
	CallerCountryIsoCode   string
	CallerProvinceCode     string
	CallerCityCode         string
	Callee                 string
	CalleeCountryIsoCode   string
	CalleeProvinceCode     string
	CalleeCityCode         string
	BeginCallTime          string
	EndTime                string
	ReleaseType            int32
	CallDuration           int32
	CallResult             int32
	AudioRecordFlag        int32
	CdrCreateTime          string
	UserData               string
	CallType               int32
	CallerShowNumber       string
	CallerShowNumberPoolNo string
	CalleeShowNumber       string
	CalleeShowNumberPoolNo string
	CallerCallingTime      string
	CallerRingingTime      string
	CallerDuration         int32
}

type QueryCallRecordMsgResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []QueryCallRecordMsg
}

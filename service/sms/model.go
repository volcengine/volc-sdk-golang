package sms

import "github.com/volcengine/volc-sdk-golang/base"

type SmsRequest struct {
	SmsAccount    string
	Sign          string
	TemplateID    string
	TemplateParam string
	PhoneNumbers  string
	Tag           string
	UserExtCode   string
}

type SmsBatchRequest struct {
	SmsAccount string
	Sign       string
	TemplateID string

	Tag      string
	From     string
	Messages []*SmsBatchMessages
}

type SmsBatchMessages struct {
	TemplateParam string
	PhoneNumber   string
}

// AssumeRole
type SmsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SmsResult `json:"Result,omitempty"`
}

type SmsResult struct {
	MessageID []string `json:"MessageID"`
}

type SmsVerifyCodeRequest struct {
	SmsAccount  string
	Sign        string
	TemplateID  string
	PhoneNumber string
	Tag         string
	UserExtCode string
	Scene       string
	CodeType    int32
	ExpireTime  int32
	TryCount    int32
}

type CheckSmsVerifyCodeRequest struct {
	SmsAccount  string
	PhoneNumber string
	Scene       string
	Code        string
}

type CheckSmsVerifyCodeResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}
type ConversionRequest struct {
	MessageIDs []string
	Type       int
}
type ConversionResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type SmsChannelType string

const (
	SmsChannelTypeCnOTP   SmsChannelType = "CN_OTP"
	SmsChannelTypeCnNTC   SmsChannelType = "CN_NTC"
	SmsChannelTypeCnMKT   SmsChannelType = "CN_MKT"
	SmsChannelTypeCnVms   SmsChannelType = "CN_VMS"
	SmsChannelTypeI18nOTP SmsChannelType = "I18N_OTP"
	SmsChannelTypeI18nMKT SmsChannelType = "I18N_MKT"
)

type Area string

const (
	AreaCN       Area = "cn"
	AreaOverseas Area = "overseas"
	AreaAll      Area = "all"
)

type SmsOrderStatus int64

const (
	SmsOrder_REVIEWING SmsOrderStatus = 1
	SmsOrder_REJECTED  SmsOrderStatus = 2
	SmsOrder_PASSED    SmsOrderStatus = 3
	SmsOrder_CLOSE     SmsOrderStatus = 4
	SmsOrder_EXEMPTED  SmsOrderStatus = 5
)

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

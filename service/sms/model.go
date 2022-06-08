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

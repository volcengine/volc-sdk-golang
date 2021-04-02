package sms

import "github.com/volcengine/volc-sdk-golang/base"

type SmsRequest struct {
	SmsAccount    string
	Sign          string
	TemplateID    string
	TemplateParam string
	PhoneNumbers  string
	Tag           string
}

// AssumeRole
type SmsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SmsResult `json:"Result,omitempty"`
}

type SmsResult struct {
	MessageID []string `json:"MessageID"`
}

package sms

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	SourceTypeText     = "text/string"
	SourceTypeImageJPG = "image/jpg"
	SourceTypeImagePNG = "image/png"
	SourceTypeImageGIF = "image/gif"
	SourceTypeVideo    = "video/mp4"
	SourceTypeAudio    = "audio/mp3"
)

type ApplyVmsTemplateRequest struct {
	SubAccount  string         `json:"subAccount"`
	ChannelType SmsChannelType `json:"channelType"`
	Name        string         `json:"name"`
	Theme       string         `json:"theme"`
	Signature   string         `json:"signature"`
	Contents    []VmsElement   `json:"contents"`
}

type VmsElement struct {
	SourceType string `json:"sourceType"`
	Content    string `json:"content"`
}

type ApplyVmsTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *VmsTemplateInfo
}

type VmsTemplateInfo struct {
	TemplateId string `json:"templateId"`
}

type GetVmsTemplateStatusRequest struct {
	SubAccount string `json:"SubAccount"`
	TemplateId string `json:"TemplateId"`
}

type GetVmsTemplateStatusResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           struct {
		ApplyResult   []CarrierReviewInfo `json:"ApplyResult"`
		ApproveTime   int64               `json:"approveTime"`   // 视频模版；模版通过的第一个运营商审批时间
		ExpireTime    int64               `json:"expireTime"`    // 视频模版；模版的到期时间
		RestValidDays int64               `json:"restValidDays"` // 视频模版；模版剩余有效天数
	}
}

type CarrierReviewInfo struct {
	Carrier string         `json:"carrier"`
	Status  SmsOrderStatus `json:"status"`
	Reason  string         `json:"reason"`
}

type SendVmsRequest struct {
	SmsAccount    string
	TemplateID    string
	TemplateParam string
	PhoneNumbers  string
	Tag           string
}

type SendVmsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *VmsResult `json:"Result,omitempty"`
}

type VmsResult struct {
	MessageID []string `json:"MessageID"`
}

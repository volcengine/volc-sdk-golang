package sms

import "github.com/volcengine/volc-sdk-golang/base"

type ShortUrlEnableStatus string

const (
	EnableStatusNotEnabled ShortUrlEnableStatus = "0"
	EnableStatusEnabled    ShortUrlEnableStatus = "1"
)

type GetSmsTemplateAndOrderListRequest struct {
	SubAccount  string `url:"subAccount,omitempty"`
	TemplateId  string `url:"templateId"`
	Name        string `url:"name"`
	Area        Area   `url:"area"`
	ChannelType string `url:"channelType"`
	Content     string `url:"content"`
	PageIndex   int    `url:"pageIndex"`
	PageSize    int    `url:"pageSize"`
}

type GetSmsTemplateAndOrderListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *struct {
		List  []*SmsTemplateInfo `json:"list"`
		Total int                `json:"total"`
	}
}

type SmsTemplateInfo struct {
	Id              string          `json:"id"`
	ApplyId         string          `json:"applyId,omitempty"`
	TemplateId      string          `json:"templateId"`
	ChannelType     SmsChannelType  `json:"channelType"`
	ChannelTypeName string          `json:"channelTypeName"`
	Name            string          `json:"name"`
	Content         string          `json:"content"`
	Status          SmsOrderStatus  `json:"status"`
	IsOrder         bool            `json:"isOrder"`
	Reason          string          `json:"reason"`
	CreatedTime     int64           `json:"createdTime"`
	ShortUrlConfig  *ShortUrlConfig `json:"shortUrlConfig"`
}

type ShortUrlConfig struct {
	IsEnabled          ShortUrlEnableStatus `json:"isEnabled"`
	IsNeedClickDetails ShortUrlEnableStatus `json:"isNeedClickDetails"`
	RawUrl             string               `json:"rawUrl"`
}

type ApplySmsTemplateRequest struct {
	SubAccount     string          `json:"subAccount"`
	Area           Area            `json:"area"`
	ChannelType    SmsChannelType  `json:"channelType"`
	Name           string          `json:"name"`
	Content        string          `json:"content"`
	Desc           string          `json:"desc"`
	ShortUrlConfig *ShortUrlConfig `json:"shortUrlConfig"`
}

type ApplySmsTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SmsTemplateInfo
}

type DeleteSmsTemplateRequest struct {
	SubAccount string `json:"subAccount"`
	Id         string `json:"id"`
	IsOrder    bool   `json:"isOrder"`
}

type DeleteSmsTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

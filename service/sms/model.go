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

	VmsOrder_VENDOR_REVIEWING SmsOrderStatus = 6 // 视频短信供应商审核中
	VmsOrder_VENDOR_REJECTED  SmsOrderStatus = 7 // 视频短信供应商审核失败
	VmsOrder_VENDOR_PASSED    SmsOrderStatus = 8 // 视频短信供应商审核通过
	VmsOrder_VENDOR_EXPIRED   SmsOrderStatus = 9 // 视频短信模版过期
)

type SendLogStatus int32

const (
	DefaultSendLog = SendLogStatus(0)
	SendNoReceipt  = SendLogStatus(1)
	SendFail       = SendLogStatus(2)
	SendAndReceipt = SendLogStatus(3)
	Verify         = SendLogStatus(4)
	Click          = SendLogStatus(5)
)

type GetSmsSendDetailsRequest struct {
	SmsAccount  string `json:"subAccount"`
	PhoneNumber string `json:"phoneNumber"`
	MessageId   string `json:"messageId"`
	SendDate    string `json:"sendDate"`
	PageSize    int64  `json:"pageSize"`
	PageIndex   int64  `json:"pageIndex"`
}

type GetSmsSendDetailsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *struct {
		List       []*SendDetailsResult `json:"sendDetailsResults"`
		Total      int                  `json:"total"`
		Account    string               `json:"account"`
		SubAccount string               `json:"subAccount"`
	}
}

type SendDetailsResult struct {
	Status       SendLogStatus `json:"status"`
	ErrorCode    string        `json:"errorCode"`
	ErrorMessage string        `json:"errorMessage"`
	PhoneNumber  string        `json:"phoneNumber"`
	Signature    string        `json:"signature"`
	TemplateId   string        `json:"templateID"`
	Content      string        `json:"content"`
	ChannelType  string        `json:"channelType"`
	MessageId    string        `json:"messageId"`
	MsgCount     int32         `json:"msgCount"`
	SendTime     int64         `json:"sendTime"`
	ReceiptTime  int64         `json:"receiptTime"`
}

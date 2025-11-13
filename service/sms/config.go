package sms

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20200101 = "2020-01-01"
	ServiceVersion20210101 = "2021-01-01"
	ServiceVersion20210111 = "2021-01-11"
	ServiceVersion20250101 = "2025-01-01"
	ServiceName            = "volcSMS"
)

var (
	ServiceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Timeout: 5 * time.Second,
			Host:    "sms.volcengineapi.com",
			Scheme:  "https",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"SendSms": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendSms"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"SendBatchSms": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendBatchSms"},
				"Version": []string{ServiceVersion20210101},
			},
		},
		"SendSmsVerifyCode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendSmsVerifyCode"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"CheckSmsVerifyCode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CheckSmsVerifyCode"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"Conversion": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Conversion"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"GetSmsTemplateAndOrderList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSmsTemplateAndOrderList"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ApplySmsTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplySmsTemplate"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"DeleteSmsTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSmsTemplate"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetSubAccountList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSubAccountList"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetSubAccountDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSubAccountDetail"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"InsertSubAccount": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"InsertSubAccount"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetSignatureAndOrderList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSignatureAndOrderList"},
				"Version": []string{ServiceVersion20250101},
			},
		},
		"ApplySmsSignature": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplySmsSignature"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"DeleteSignature": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSignature"},
				"Version": []string{ServiceVersion20250101},
			},
		},
		"ApplyVmsTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyVmsTemplate"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetVmsTemplateStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetVmsTemplateStatus"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetSmsSendDetails": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSmsSendDetails"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ApplySignatureIdent": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplySignatureIdent"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetSignatureIdentList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSignatureIdentList"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"BatchBindSignatureIdent": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchBindSignatureIdent"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ApplySmsSignatureV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplySmsSignatureV2"},
				"Version": []string{ServiceVersion20250101},
			},
		},
		"UpdateSmsSignature": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSmsSignature"},
				"Version": []string{ServiceVersion20250101},
			},
		},
		"UpdateSignatureIdentification": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSignatureIdent"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"BulkCreateTobTrafficDrivingLink": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BulkCreateTobTrafficDrivingLink"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"BulkCreateTobTrafficDrivingPhone": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BulkCreateTobTrafficDrivingPhone"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"DeleteTobTrafficDrivingLink": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTobTrafficDrivingLinkSDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"DeleteTobTrafficDrivingPhone": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTobTrafficDrivingPhoneSDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetRelationTemplateList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRelationTemplateSDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetTobTrafficDrivingLinkList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTobTrafficDrivingLinkList"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetTobTrafficDrivingPhoneList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTobTrafficDrivingPhoneList"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"UpdateTobTrafficDrivingPhone": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTobTrafficDrivingPhoneSDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"BindTrafficDrivingParams": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindTrafficDrivingParamsSDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ApplySmsTemplateV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplySmsTemplateV2SDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ApplySmsSubContentTemplateV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplySmsSubContentTemplateV2SDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ListSmsTemplateV2": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSmsTemplateV2SDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"BindSignatures": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindSignaturesSDK"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ListSecondTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSecondTemplate"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ListSubContent": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSubContent"},
				"Version": []string{ServiceVersion20210111},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type SMS struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *SMS {
	instance := &SMS{}
	instance.Client = base.NewClient(ServiceInfo[DefaultRegion], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *SMS) GetServiceInfo(region string) *base.ServiceInfo {
	if serviceInfo, ok := ServiceInfo[region]; ok {
		return serviceInfo
	}
	return nil
}

// GetAPIInfo interface
func (p *SMS) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (s *SMS) SetRegion(region string) {
	if serviceInfo := s.GetServiceInfo(region); serviceInfo != nil {
		si := serviceInfo.Clone()
		si.Credentials = s.Client.ServiceInfo.Credentials.Clone()
		si.Credentials.Region = region
		s.Client.ServiceInfo = si
	}
}

// SetHost .
func (s *SMS) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *SMS) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}

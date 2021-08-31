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
	ServiceName            = "volcSMS"
)

var (
	ServiceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Timeout: 5 * time.Second,
			Host:    "sms.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "sms.byteplusapi.com",
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

func NewInstanceI18n(region string) *SMS {
	instance := &SMS{}
	instance.Client = base.NewClient(ServiceInfo[base.RegionApSingapore], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = region
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
		serviceInfo.Credentials.Region = region
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

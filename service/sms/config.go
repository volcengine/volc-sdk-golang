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
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "sms.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
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
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *SMS) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
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
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (s *SMS) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *SMS) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}

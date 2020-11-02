package iam

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20180101 = "2018-01-01"
	ServiceName            = "iam"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"ListUsers": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUsers"},
				"Version": []string{ServiceVersion20180101},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type IAM struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *IAM {
	instance := &IAM{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	ServiceInfo.Credentials.Service = ServiceName
	ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *IAM) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *IAM) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *IAM) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *IAM) SetHost(host string) {
	ServiceInfo.Host = host
}

// SetSchema .
func (p *IAM) SetSchema(schema string) {
	ServiceInfo.Scheme = schema
}

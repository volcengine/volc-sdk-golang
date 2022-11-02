package sami

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20210727 = "2021-07-27"
	ServiceName            = "sami"
	SamiDomain             = "https://sami.bytedance.com"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 10 * time.Second,
		Host:    "open.volcengineapi.com",
		Header:  http.Header{},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"GetToken": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetToken"},
				"Version": []string{ServiceVersion20210727},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type Sami struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *Sami {
	instance := &Sami{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *Sami) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

// GetAPIInfo interface
func (p *Sami) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion
func (p *Sami) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Sami) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *Sami) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}

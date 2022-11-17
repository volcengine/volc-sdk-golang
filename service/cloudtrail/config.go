package cloudtrail

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20210901 = "2021-09-01"
	ServiceName            = "cloud_trail"
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
		// events
		"LookupEvents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"LookupEvents"},
				"Version": []string{ServiceVersion20210901},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// CloudTrail .
type CloudTrail struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *CloudTrail {
	instance := &CloudTrail{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *CloudTrail) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

// GetAPIInfo interface
func (p *CloudTrail) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion .
func (p *CloudTrail) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *CloudTrail) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *CloudTrail) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}

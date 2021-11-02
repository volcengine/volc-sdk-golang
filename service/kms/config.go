package kms

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20210218 = "2021-02-18"
	ServiceName            = "kms"
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
		"CreateKeyring": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateKeyring"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"DescribeKeyrings": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeKeyrings"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"UpdateKeyring": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateKeyring"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"QueryKeyring": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryKeyring"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"CreateKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateKey"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"DescribeKeys": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeKeys"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"UpdateKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateKey"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"GenerateDataKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GenerateDataKey"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"Encrypt": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Encrypt"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"Decrypt": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Decrypt"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"EnableKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableKey"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"DisableKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableKey"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"ScheduleKeyDeletion": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ScheduleKeyDeletion"},
				"Version": []string{ServiceVersion20210218},
			},
		},
		"CancelKeyDeletion": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CancelKeyDeletion"},
				"Version": []string{ServiceVersion20210218},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// KMS .
type KMS struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *KMS {
	instance := &KMS{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *KMS) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *KMS) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion .
func (p *KMS) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *KMS) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *KMS) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}

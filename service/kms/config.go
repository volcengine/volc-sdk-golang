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
	defaultRetrySetting = base.RetrySettings{
		AutoRetry: true,
	}
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
			Retry: defaultRetrySetting,
		},
		"DescribeKeyrings": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeKeyrings"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"UpdateKeyring": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateKeyring"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"QueryKeyring": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryKeyring"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"CreateKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"DescribeKeys": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeKeys"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"DescribeKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"UpdateKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"GenerateDataKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GenerateDataKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"Encrypt": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Encrypt"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"Decrypt": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Decrypt"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"EnableKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"DisableKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"ScheduleKeyDeletion": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ScheduleKeyDeletion"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"CancelKeyDeletion": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CancelKeyDeletion"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"ArchiveKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ArchiveKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"CancelArchiveKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CancelArchiveKey"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"EnableKeyRotation": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableKeyRotation"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"DisableKeyRotation": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableKeyRotation"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
		},
		"ReEncrypt": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ReEncrypt"},
				"Version": []string{ServiceVersion20210218},
			},
			Retry: defaultRetrySetting,
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

package notify

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion  = "cn-north-1"
	ServiceVersion = "2021-01-01"
	ServiceName    = "volc_voice_notify"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "cloud-vms.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{

		// accessKey
		"CreateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"BatchAppend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchAppend"},
				"Version": []string{ServiceVersion},
			},
		},
		"PauseTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PauseTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"ResumeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeTask"},
				"Version": []string{ServiceVersion},
			},
		},

		// policy
		"StopTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"UpdateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"SingleBatchAppend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SingleBatchAppend"},
				"Version": []string{ServiceVersion},
			},
		},
		"SingleInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SingleInfo"},
				"Version": []string{ServiceVersion},
			},
		},
		"SingleCancel": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SingleCancel"},
				"Version": []string{ServiceVersion},
			},
		},
		"FetchResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FetchResource"},
				"Version": []string{ServiceVersion},
			},
		},
		"OpenCreateTts": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OpenCreateTts"},
				"Version": []string{ServiceVersion},
			},
		},

		// role
		"OpenDeleteResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OpenDeleteResource"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetResourceUploadUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResourceUploadUrl"},
				"Version": []string{ServiceVersion},
			},
		},
		"CommitResourceUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitResourceUpload"},
				"Version": []string{ServiceVersion},
			},
		},
		"OpenUpdateResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OpenUpdateResource"},
				"Version": []string{ServiceVersion},
			},
		},
		"QueryUsableResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryUsableResource"},
				"Version": []string{ServiceVersion},
			},
		},
		"QueryOpenGetResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryOpenGetResource"},
				"Version": []string{ServiceVersion},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type Notify struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *Notify {
	instance := &Notify{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *Notify) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

// GetAPIInfo interface
func (p *Notify) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *Notify) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Notify) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *Notify) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}

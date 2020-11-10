package visual

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20200826 = "2020-08-26"
	ServiceName            = "cv"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 10 * time.Second,
		Host:    "open.volcengineapi.com",
		Header:  http.Header{},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"BankCard": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BankCard"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"IDCard": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"IDCard"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OCRNormal": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OCRNormal"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"FaceSwap": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FaceSwap"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"JPCartoon": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"JPCartoon"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"JPCartoonCut": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"JPCartoonCut"},
				"Version": []string{ServiceVersion20200826},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type Visual struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *Visual {
	instance := &Visual{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	ServiceInfo.Credentials.Service = ServiceName
	ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *Visual) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *Visual) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion
func (p *Visual) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Visual) SetHost(host string) {
	ServiceInfo.Host = host
}

// SetSchema .
func (p *Visual) SetSchema(schema string) {
	ServiceInfo.Scheme = schema
}

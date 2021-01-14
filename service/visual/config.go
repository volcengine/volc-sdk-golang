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
		Host:    "visual.volcengineapi.com",
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
		"VideoSceneDetect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoSceneDetect"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OverResolution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OverResolution"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"GoodsSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GoodsSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageOutpaint": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageOutpaint"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageInpaint": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageInpaint"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageCut": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageCut"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EntityDetect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EntityDetect"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"GoodsDetect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GoodsDetect"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ConvertPhoto": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ConvertPhoto"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EnhancePhoto": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnhancePhoto"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"GeneralSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneralSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"HumanSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"HumanSegment"},
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
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
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

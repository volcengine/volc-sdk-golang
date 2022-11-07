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
		"DistortionFree": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DistortionFree"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"StretchRecovery": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StretchRecovery"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EmoticonEdit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EmoticonEdit"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EyeClose2Open": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EyeClose2Open"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageScore": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageScore"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageFlow": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageFlow"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"PoemMaterial": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PoemMaterial"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"CarPlateDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CarPlateDetection"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"CarSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CarSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"CarDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CarDetection"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"SkySegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SkySegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
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
		"VideoInpaintSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoInpaintSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoInpaintQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoInpaintQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoRetargetingSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoRetargetingSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoRetargetingQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoRetargetingQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoSummarizationSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoSummarizationSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoSummarizationQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoSummarizationQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoOverResolutionSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoOverResolutionSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoOverResolutionQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoOverResolutionQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OverResolutionV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OverResolutionV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
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
	return p.Client.ServiceInfo
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
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Visual) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *Visual) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}

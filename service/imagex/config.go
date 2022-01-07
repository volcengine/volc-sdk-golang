package imagex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "ImageX"
	ApiVersion     = "2018-08-01"
	DefaultTimeout = 10 * time.Second

	ResourceServiceIdTRN = "trn:ImageX:*:*:ServiceId/%s"
	ResourceStoreKeyTRN  = "trn:ImageX:*:*:StoreKeys/%s"
)

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: ServiceName,
			},
		},
		base.RegionUsEast1: {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex-us-east-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionUsEast1,
				Service: ServiceName,
			},
		},
		base.RegionApSingapore: {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionApSingapore,
				Service: ServiceName,
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		// 服务管理
		"GetImageService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageService"},
				"Version": []string{ApiVersion},
			},
		},
		"GetAllImageServices": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllImageServices"},
				"Version": []string{ApiVersion},
			},
		},
		"UpdateImageAuthKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageAuthKey"},
				"Version": []string{ApiVersion},
			},
		},
		"UpdateImageObjectAccess": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageObjectAccess"},
				"Version": []string{ApiVersion},
			},
		},
		"UpdateImageMirrorConf": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageMirrorConf"},
				"Version": []string{ApiVersion},
			},
		},
		// 域名管理
		"GetServiceDomains": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetServiceDomains"},
				"Version": []string{ApiVersion},
			},
		},
		"GetDomainConfig": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDomainConfig"},
				"Version": []string{ApiVersion},
			},
		},
		"UpdateResponseHeader": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateResponseHeader"},
				"Version": []string{ApiVersion},
			},
		},
		"UpdateRefer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRefer"},
				"Version": []string{ApiVersion},
			},
		},
		"UpdateHttps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHttps"},
				"Version": []string{ApiVersion},
			},
		},
		"GetResponseHeaderValidateKeys": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResponseHeaderValidateKeys"},
				"Version": []string{ApiVersion},
			},
		},
		// 模板管理
		"CreateImageTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageTemplate"},
				"Version": []string{ApiVersion},
			},
		},
		"DeleteImageTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageTemplate"},
				"Version": []string{ApiVersion},
			},
		},
		"PreviewImageTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PreviewImageTemplate"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageTemplate"},
				"Version": []string{ApiVersion},
			},
		},
		"GetAllImageTemplates": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllImageTemplates"},
				"Version": []string{ApiVersion},
			},
		},
		// 资源管理相关
		"DeleteImageUploadFiles": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageUploadFiles"},
				"Version": []string{ApiVersion},
			},
		},
		"ApplyImageUpload": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyImageUpload"},
				"Version": []string{ApiVersion},
			},
		},
		"CommitImageUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitImageUpload"},
				"Version": []string{ApiVersion},
			},
		},
		"UpdateImageUploadFiles": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageUploadFiles"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageUploadFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUploadFile"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageUploadFiles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUploadFiles"},
				"Version": []string{ApiVersion},
			},
		},
		"PreviewImageUploadFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PreviewImageUploadFile"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageUpdateFiles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUpdateFiles"},
				"Version": []string{ApiVersion},
			},
		},
		"FetchImageUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FetchImageUrl"},
				"Version": []string{ApiVersion},
			},
		},

		"GetImageStyleResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageStyleResult"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageOCR": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageOCR"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageBgFillResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageBgFillResult"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageEnhanceResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEnhanceResult"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageEraseModels": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEraseModels"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageEraseResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEraseResult"},
				"Version": []string{ApiVersion},
			},
		},
		"GetImageQuality": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageQuality"},
				"Version": []string{ApiVersion},
			},
		},
	}
)

// DefaultInstance 默认实例，Region: cn-north-1
var DefaultInstance = NewInstance()

type ImageX struct {
	*base.Client
}

func NewInstance() *ImageX {
	instance := &ImageX{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *ImageX {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("ImageX not support region %s", region))
	}
	instance := &ImageX{
		Client: base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

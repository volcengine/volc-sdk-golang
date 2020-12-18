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

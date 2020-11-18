package vod

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	UPDATE_INTERVAL = 10
)

type Vod struct {
	*base.Client
	DomainCache map[string]map[string]int
	Lock        sync.RWMutex
}

const (
	// ResourceSpaceFormat space的资源结构
	ResourceSpaceFormat = "trn:vod:%s:*:space/%s"
	// ResourceVideoFormat vid的视频结构
	ResourceVideoFormat = "trn:vod::*:video_id/%s"
	// ResourceStreamTypeFormat  stream type理论和账号无关
	ResourceStreamTypeFormat = "trn:vod:::stream_type/%s"
	// ResourceWatermarkFormat 水印信息与账号有关,与region无关
	ResourceWatermarkFormat = "trn:vod::*:watermark/%s"

	ActionGetPalyInfo = "vod:GetPlayInfo"

	Star = "*"
)

func NewInstance() *Vod {
	instance := &Vod{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *Vod {
	var serviceInfo *base.ServiceInfo
	var ok bool
	if serviceInfo, ok = ServiceInfoMap[region]; !ok {
		panic("Cant find the region, please check it carefully")
	}

	instance := &Vod{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			//Host:    "vod.volcengineapi.com",
			Host: "volcengineapi-boe.byted.org",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "vod"},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "vod.ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "vod"},
		},
		base.RegionUsEast1: {
			Timeout: 5 * time.Second,
			Host:    "vod.us-east-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionUsEast1, Service: "vod"},
		},
	}

	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "vod.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "vod"},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"GetPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfo"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetOriginalPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetOriginalPlayInfo"},
				"Version": []string{"2020-08-01"},
			},
		},
		"StartWorkflow": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWorkflow"},
				"Version": []string{"2020-08-01"},
			},
		},
		"UploadMediaByUrl": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadMediaByUrl"},
				"Version": []string{"2018-01-01"},
			},
		},
		"ApplyUpload": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyUpload"},
				"Version": []string{"2018-01-01"},
			},
		},
		"CommitUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUpload"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetCdnDomainWeights": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCdnDomainWeights"},
				"Version": []string{"2019-07-01"},
			},
		},
		"UpdateMediaInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateMediaInfo"},
				"Version": []string{"2020-08-01"},
			},
		},
		"UpdateMediaPublishStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateMediaPublishStatus"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetMediaInfos": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMediaInfos"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetRecommendedPoster": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecommendedPoster"},
				"Version": []string{"2020-08-01"},
			},
		},
	}
)

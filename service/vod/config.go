package vod

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

type Vod struct {
	*base.Client
	DomainCache map[string]map[string]int
	Lock        sync.RWMutex
}

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
			Host:    "vod.volcengineapi.com",
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

	ApiInfoList = map[string]*base.ApiInfo{
		"GetPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfo"},
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
		"QueryUploadTaskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryUploadTaskInfo"},
				"Version": []string{"2020-08-01"},
			},
		},
		"UploadMediaByUrl": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadMediaByUrl"},
				"Version": []string{"2020-08-01"},
			},
		},
		"ApplyUploadInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyUploadInfo"},
				"Version": []string{"2020-08-01"},
			},
		},
		"CommitUploadInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUploadInfo"},
				"Version": []string{"2020-08-01"},
			},
			Timeout: 8 * time.Second,
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
		"DeleteMedia": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteMedia"},
				"Version": []string{"2020-08-01"},
			},
		},
		"DeleteTranscodes": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTranscodes"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetMediaList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMediaList"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetHlsDecryptionKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetHlsDecryptionKey"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetPrivateDrmPlayAuth": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPrivateDrmPlayAuth"},
				"Version": []string{"2020-08-01"},
			},
		},
	}
)

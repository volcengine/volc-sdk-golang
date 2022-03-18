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
	}

	ApiInfoList = map[string]*base.ApiInfo{
		// **********************************************************************
		// 播放
		// **********************************************************************
		"GetPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfo"},
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
		"GetHlsDecryptionKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetHlsDecryptionKey"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetPlayInfoWithLiveTimeShiftScene": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfoWithLiveTimeShiftScene"},
				"Version": []string{"2021-11-01"},
			},
		},

		// **********************************************************************
		// 上传
		// **********************************************************************
		"UploadMediaByUrl": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadMediaByUrl"},
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

		// **********************************************************************
		// 媒资
		// **********************************************************************
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
		"GetSubtitleInfoList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSubtitleInfoList"},
				"Version": []string{"2020-08-01"},
			},
		},
		"UpdateSubtitleStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSubtitleStatus"},
				"Version": []string{"2020-08-01"},
			},
		},
		"UpdateSubtitleInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSubtitleInfo"},
				"Version": []string{"2020-08-01"},
			},
		},
		//"GetMLFramesForAudit": {
		//	Method: http.MethodGet,
		//	Path:   "/",
		//	Query: url.Values{
		//		"Action":  []string{"GetMLFramesForAudit"},
		//		"Version": []string{"2021-11-01"},
		//	},
		//},
		//"GetBetterFramesForAudit": {
		//	Method: http.MethodGet,
		//	Path:   "/",
		//	Query: url.Values{
		//		"Action":  []string{"GetBetterFramesForAudit"},
		//		"Version": []string{"2021-11-01"},
		//	},
		//},
		//"GetAudioInfoForAudit": {
		//	Method: http.MethodGet,
		//	Path:   "/",
		//	Query: url.Values{
		//		"Action":  []string{"GetAudioInfoForAudit"},
		//		"Version": []string{"2021-11-01"},
		//	},
		//},
		//"GetAutomaticSpeechRecognitionForAudit": {
		//	Method: http.MethodGet,
		//	Path:   "/",
		//	Query: url.Values{
		//		"Action":  []string{"GetAutomaticSpeechRecognitionForAudit"},
		//		"Version": []string{"2021-11-01"},
		//	},
		//},
		//"GetAudioEventDetectionForAudit": {
		//	Method: http.MethodGet,
		//	Path:   "/",
		//	Query: url.Values{
		//		"Action":  []string{"GetAudioEventDetectionForAudit"},
		//		"Version": []string{"2021-11-01"},
		//	},
		//},
		"CreateVideoClassification": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateVideoClassification"},
				"Version": []string{"2021-01-01"},
			},
		},
		"UpdateVideoClassification": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateVideoClassification"},
				"Version": []string{"2021-01-01"},
			},
		},
		"DeleteVideoClassification": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteVideoClassification"},
				"Version": []string{"2021-01-01"},
			},
		},
		"ListVideoClassifications": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVideoClassifications"},
				"Version": []string{"2021-01-01"},
			},
		},
		"ListSnapshots": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSnapshots"},
				"Version": []string{"2021-01-01"},
			},
		},

		// **********************************************************************
		// 转码
		// **********************************************************************
		"StartWorkflow": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWorkflow"},
				"Version": []string{"2020-08-01"},
			},
		},

		// **********************************************************************
		// 空间管理
		// **********************************************************************
		"CreateSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSpace"},
				"Version": []string{"2021-01-01"},
			},
		},
		"ListSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSpace"},
				"Version": []string{"2021-01-01"},
			},
		},
		"GetSpaceDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSpaceDetail"},
				"Version": []string{"2022-01-01"},
			},
		},
		//"GetSpaceConfig": {
		//	Method: http.MethodGet,
		//	Path:   "/",
		//	Query: url.Values{
		//		"Action":  []string{"GetSpaceConfig"},
		//		"Version": []string{"2022-01-01"},
		//	},
		//},
		"UpdateSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSpace"},
				"Version": []string{"2021-01-01"},
			},
		},
		"UpdateSpaceUploadConfig": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSpaceUploadConfig"},
				"Version": []string{"2022-01-01"},
			},
		},

		// **********************************************************************
		// 分发加速管理
		// **********************************************************************
		"ListDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDomain"},
				"Version": []string{"2021-01-01"},
			},
		},
		"CreateCdnRefreshTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCdnRefreshTask"},
				"Version": []string{"2021-01-01"},
			},
		},
		"CreateCdnPreloadTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCdnPreloadTask"},
				"Version": []string{"2021-01-01"},
			},
		},

		// **********************************************************************
		// 回调管理
		// **********************************************************************
		"AddCallbackSubscription": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddCallbackSubscription"},
				"Version": []string{"2021-12-01"},
			},
		},
		"SetCallbackEvent": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetCallbackEvent"},
				"Version": []string{"2021-01-01"},
			},
		},
	}
)

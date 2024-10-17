package vod

import (
	"fmt"
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
		serviceInfo = &base.ServiceInfo{
			Timeout: 60 * time.Second,
			Scheme:  "https",
			Host:    fmt.Sprintf("vod.%s.volcengineapi.com", region),
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: region, Service: "vod"},
		}
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
			Timeout: 60 * time.Second,
			Scheme:  "https",
			Host:    "vod.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "vod"},
		},
		base.RegionApSouthEast1: {
			Timeout: 60 * time.Second,
			Scheme:  "https",
			Host:    "vod.ap-southeast-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSouthEast1, Service: "vod"},
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
		"DescribeDrmDataKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDrmDataKey"},
				"Version": []string{"2023-07-01"},
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
		"GetAllPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllPlayInfo"},
				"Version": []string{"2022-01-01"},
			},
		},
		// **********************************************************************
		// 通用文件操作
		// **********************************************************************

		"SubmitBlockObjectTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitBlockObjectTasks"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListBlockObjectTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListBlockObjectTasks"},
				"Version": []string{"2023-07-01"},
			},
		},

		// **********************************************************************
		// 文件上传
		// **********************************************************************
		"SubmitMoveObjectTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitMoveObjectTask"},
				"Version": []string{"2023-07-01"},
			},
		},
		"QueryMoveObjectTaskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryMoveObjectTaskInfo"},
				"Version": []string{"2023-07-01"},
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
				"Version": []string{"2022-01-01"},
			},
		},
		"CommitUploadInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUploadInfo"},
				"Version": []string{"2022-01-01"},
			},
			Timeout: 30 * time.Second,
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
		"UpdateMediaStorageClass": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateMediaStorageClass"},
				"Version": []string{"2022-12-01"},
			},
		},
		"GetMediaInfos": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMediaInfos"},
				"Version": []string{"2022-12-01"},
			},
		},
		"GetMediaInfos20230701": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMediaInfos"},
				"Version": []string{"2023-07-01"},
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
		"DeleteMediaTosFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteMediaTosFile"},
				"Version": []string{"2022-12-01"},
			},
		},
		"GetFileInfos": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetFileInfos"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateFileStorageClass": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateFileStorageClass"},
				"Version": []string{"2023-07-01"},
			},
		},
		"GetMediaList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMediaList"},
				"Version": []string{"2022-12-01"},
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
		"GetAuditFramesForAudit": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAuditFramesForAudit"},
				"Version": []string{"2021-11-01"},
			},
		},
		"GetMLFramesForAudit": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMLFramesForAudit"},
				"Version": []string{"2021-11-01"},
			},
		},
		"GetBetterFramesForAudit": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetBetterFramesForAudit"},
				"Version": []string{"2021-11-01"},
			},
		},
		"GetAudioInfoForAudit": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAudioInfoForAudit"},
				"Version": []string{"2021-11-01"},
			},
		},
		"GetAutomaticSpeechRecognitionForAudit": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAutomaticSpeechRecognitionForAudit"},
				"Version": []string{"2021-11-01"},
			},
		},
		"GetAudioEventDetectionForAudit": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAudioEventDetectionForAudit"},
				"Version": []string{"2021-11-01"},
			},
		},
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
		"ExtractMediaMetaTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ExtractMediaMetaTask"},
				"Version": []string{"2022-01-01"},
			},
		},
		"SubmitBlockMediaTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitBlockMediaTask"},
				"Version": []string{"2022-12-01"},
			},
		},
		"ListFileMetaInfosByFileNames": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListFileMetaInfosByFileNames"},
				"Version": []string{"2023-07-01"},
			},
		},
		"SubmitUnblockMediaTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitUnblockMediaTask"},
				"Version": []string{"2022-12-01"},
			},
		},
		"QueryMediaBlockStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryMediaBlockStatus"},
				"Version": []string{"2022-12-01"},
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
		"RetrieveTranscodeResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RetrieveTranscodeResult"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetWorkflowExecution": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWorkflowExecution"},
				"Version": []string{"2020-08-01"},
			},
		},
		"GetWorkflowExecutionResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWorkflowExecutionResult"},
				"Version": []string{"2022-12-01"},
			},
		},
		"GetTaskTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTaskTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"CreateTaskTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTaskTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateTaskTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTaskTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListTaskTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListTaskTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DeleteTaskTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTaskTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"GetWorkflowTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWorkflowTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"CreateWorkflowTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateWorkflowTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateWorkflowTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateWorkflowTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListWorkflowTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListWorkflowTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DeleteWorkflowTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteWorkflowTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"GetWatermarkTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWatermarkTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"CreateWatermarkTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateWatermarkTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateWatermarkTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateWatermarkTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListWatermarkTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListWatermarkTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DeleteWatermarkTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteWatermarkTemplate"},
				"Version": []string{"2023-07-01"},
			},
		},

		// **********************************************************************
		// 视频编辑
		// **********************************************************************
		"SubmitDirectEditTaskAsync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitDirectEditTaskAsync"},
				"Version": []string{"2018-01-01"},
			},
		},
		"SubmitDirectEditTaskSync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitDirectEditTaskSync"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetDirectEditResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDirectEditResult"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetDirectEditProgress": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDirectEditProgress"},
				"Version": []string{"2018-01-01"},
			},
		},
		"CancelDirectEditTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CancelDirectEditTask"},
				"Version": []string{"2018-01-01"},
			},
		},

		// **********************************************************************
		// 空间管理
		// **********************************************************************
		"DeleteSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSpace"},
				"Version": []string{"2023-07-01"},
			},
		},
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
		"DescribeVodSpaceStorageData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceStorageData"},
				"Version": []string{"2023-07-01"},
			},
		},

		// **********************************************************************
		// 分发加速管理
		// **********************************************************************
		"DeleteDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDomain"},
				"Version": []string{"2023-07-01"},
			},
		},
		"StopDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopDomain"},
				"Version": []string{"2023-07-01"},
			},
		},
		"StartDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartDomain"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateDomainPlayRule": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainPlayRule"},
				"Version": []string{"2023-07-01"},
			},
		},
		"RemoveDomainFromScheduler": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemoveDomainFromScheduler"},
				"Version": []string{"2023-07-01"},
			},
		},
		"AddDomainToScheduler": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddDomainToScheduler"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDomain"},
				"Version": []string{"2023-01-01"},
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
		"ListCdnTasks": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnTasks"},
				"Version": []string{"2022-01-01"},
			},
		},
		"ListCdnAccessLog": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnAccessLog"},
				"Version": []string{"2022-01-01"},
			},
		},
		"ListCdnTopAccessUrl": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnTopAccessUrl"},
				"Version": []string{"2022-01-01"},
			},
		},
		"ListCdnTopAccess": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnTopAccess"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodDomainBandwidthData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodDomainBandwidthData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListCdnUsageData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnUsageData"},
				"Version": []string{"2022-12-01"},
			},
		},
		"ListCdnStatusData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnStatusData"},
				"Version": []string{"2022-01-01"},
			},
		},
		"DescribeIpInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeIpInfo"},
				"Version": []string{"2022-01-01"},
			},
		},
		"ListCdnPvData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnPvData"},
				"Version": []string{"2022-01-01"},
			},
		},
		"DescribeVodDomainTrafficData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodDomainTrafficData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"SubmitBlockTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitBlockTasks"},
				"Version": []string{"2022-01-01"},
			},
		},
		"GetContentBlockTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetContentBlockTasks"},
				"Version": []string{"2022-01-01"},
			},
		},
		"CreateDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDomain"},
				"Version": []string{"2023-02-01"},
			},
		},
		"UpdateDomainExpire": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainExpire"},
				"Version": []string{"2023-02-01"},
			},
		},
		"UpdateDomainAuthConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainAuthConfig"},
				"Version": []string{"2023-02-01"},
			},
		},
		"ListCdnHitrateData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnHitrateData"},
				"Version": []string{"2022-01-01"},
			},
		},
		"AddOrUpdateCertificate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddOrUpdateCertificate"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateDomainUrlAuthConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainUrlAuthConfig"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeDomainConfig": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDomainConfig"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateDomainConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainConfig"},
				"Version": []string{"2023-07-01"},
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
				"Version": []string{"2022-01-01"},
			},
		},
		// **********************************************************************
		// 计量计费
		// **********************************************************************
		"DescribeVodSpaceTranscodeData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceTranscodeData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodSpaceAIStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceAIStatisData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodSpaceSubtitleStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceSubtitleStatisData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodSpaceDetectStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceDetectStatisData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodSnapshotData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSnapshotData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodSpaceWorkflowDetailData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceWorkflowDetailData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodSpaceEditDetailData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceEditDetailData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodPlayFileLogByDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodPlayFileLogByDomain"},
				"Version": []string{"2022-12-01"},
			},
		},
		"DescribeVodEnhanceImageData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodEnhanceImageData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodSpaceEditStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceEditStatisData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodPlayedStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodPlayedStatisData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodMostPlayedStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodMostPlayedStatisData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodRealtimeMediaData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodRealtimeMediaData"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DescribeVodRealtimeMediaDetailData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodRealtimeMediaDetailData"},
				"Version": []string{"2023-07-01"},
			},
		},
		// **********************************************************************
		// 迁移平台
		// **********************************************************************
		"SetCloudMigrateJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetCloudMigrateJob"},
				"Version": []string{"2023-07-01"},
			},
		},
		"GetCloudMigrateJob": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCloudMigrateJob"},
				"Version": []string{"2023-07-01"},
			},
		},
		"SubmitCloudMigrateJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitCloudMigrateJob"},
				"Version": []string{"2023-07-01"},
			},
		},
	}
)

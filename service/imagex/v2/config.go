package imagex

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "ImageX"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-north-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-north-1",
				Service: ServiceName,
			},
		},
		"ap-singapore-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "ap-singapore-1",
				Service: ServiceName,
			},
		},
		"us-east-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex-us-east-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "us-east-1",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"UpdateImageDomainVolcOrigin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageDomainVolcOrigin"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DelDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DelDomain"},
				"Version": []string{"2023-05-01"},
			},
		},
		"AddDomainV1": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddDomainV1"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageDomainIPAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageDomainIPAuth"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateRefer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRefer"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageDomainUaAccess": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageDomainUaAccess"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateHttps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHttps"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageDomainDownloadSpeedLimit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageDomainDownloadSpeedLimit"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateResponseHeader": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateResponseHeader"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageDomainAreaAccess": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageDomainAreaAccess"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateDomainAdaptiveFmt": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainAdaptiveFmt"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageDomainConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageDomainConfig"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateAdvance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAdvance"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageDomainBandwidthLimit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageDomainBandwidthLimit"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateSlimConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSlimConfig"},
				"Version": []string{"2023-05-01"},
			},
		},
		"SetDefaultDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetDefaultDomain"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageVolcCdnAccessLog": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageVolcCdnAccessLog"},
				"Version": []string{"2018-08-01"},
			},
		},
		"VerifyDomainOwner": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VerifyDomainOwner"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetResponseHeaderValidateKeys": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResponseHeaderValidateKeys"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetDomainConfig": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDomainConfig"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetDomainOwnerVerifyContent": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDomainOwnerVerifyContent"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetServiceDomains": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetServiceDomains"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteImageMonitorRules": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageMonitorRules"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteImageMonitorRecords": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageMonitorRecords"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageMonitorRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageMonitorRule"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageMonitorRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageMonitorRule"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageMonitorRuleStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageMonitorRuleStatus"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageAlertRecords": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAlertRecords"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageMonitorRules": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageMonitorRules"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageSettingRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageSettingRule"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageSettingRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageSettingRule"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageSettingRulePriority": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageSettingRulePriority"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageSettingRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageSettingRule"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageSettings": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageSettings"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageSettingRuleHistory": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageSettingRuleHistory"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageSettingRules": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageSettingRules"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageMigrateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageMigrateTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageMigrateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageMigrateTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"ExportFailedMigrateTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ExportFailedMigrateTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageTaskStrategy": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageTaskStrategy"},
				"Version": []string{"2023-05-01"},
			},
		},
		"TerminateImageMigrateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TerminateImageMigrateTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetVendorBuckets": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetVendorBuckets"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageMigrateTasks": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageMigrateTasks"},
				"Version": []string{"2023-05-01"},
			},
		},
		"RerunImageMigrateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RerunImageMigrateTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAddOnTag": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAddOnTag"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCubeUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCubeUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSourceRequestBandwidth": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSourceRequestBandwidth"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSourceRequestTraffic": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSourceRequestTraffic"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSourceRequest": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSourceRequest"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXStorageUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXStorageUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXBucketRetrievalUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXBucketRetrievalUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXAIRequestCntUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXAIRequestCntUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSummary": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSummary"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXDomainTrafficData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXDomainTrafficData"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXDomainBandwidthData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXDomainBandwidthData"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXDomainBandwidthNinetyFiveData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXDomainBandwidthNinetyFiveData"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXBucketUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXBucketUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXBillingRequestCntUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXBillingRequestCntUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXRequestCntUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXRequestCntUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXBaseOpUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXBaseOpUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCompressUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCompressUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXScreenshotUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXScreenshotUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXVideoClipDurationUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXVideoClipDurationUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXMultiCompressUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXMultiCompressUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXEdgeRequest": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXEdgeRequest"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXEdgeRequestBandwidth": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXEdgeRequestBandwidth"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXEdgeRequestTraffic": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXEdgeRequestTraffic"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXEdgeRequestRegions": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXEdgeRequestRegions"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXMirrorRequestHttpCodeByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXMirrorRequestHttpCodeByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXMirrorRequestHttpCodeOverview": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXMirrorRequestHttpCodeOverview"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXMirrorRequestTraffic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXMirrorRequestTraffic"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXMirrorRequestBandwidth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXMirrorRequestBandwidth"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXServerQPSUsage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXServerQPSUsage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHitRateTrafficData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHitRateTrafficData"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHitRateRequestData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHitRateRequestData"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCDNTopRequestData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCDNTopRequestData"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHeifEncodeFileInSizeByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHeifEncodeFileInSizeByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHeifEncodeFileOutSizeByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHeifEncodeFileOutSizeByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHeifEncodeSuccessCountByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHeifEncodeSuccessCountByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHeifEncodeSuccessRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHeifEncodeSuccessRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHeifEncodeDurationByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHeifEncodeDurationByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXHeifEncodeErrorCodeByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXHeifEncodeErrorCodeByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXExceedResolutionRatioAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXExceedResolutionRatioAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXExceedFileSize": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXExceedFileSize"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXExceedCountByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXExceedCountByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXServiceQuality": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXServiceQuality"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageXQueryApps": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageXQueryApps"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageXQueryRegions": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageXQueryRegions"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageXQueryDims": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageXQueryDims"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageXQueryVals": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageXQueryVals"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadCountByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadCountByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadDuration": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadDuration"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadSuccessRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadSuccessRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadFileSize": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadFileSize"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadErrorCodeByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadErrorCodeByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadErrorCodeAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadErrorCodeAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadSpeed": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadSpeed"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXUploadSegmentSpeedByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXUploadSegmentSpeedByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnSuccessRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnSuccessRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnSuccessRateAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnSuccessRateAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnErrorCodeByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnErrorCodeByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnErrorCodeAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnErrorCodeAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnDurationDetailByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnDurationDetailByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnDurationAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnDurationAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnReuseRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnReuseRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnReuseRateAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnReuseRateAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXCdnProtocolRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXCdnProtocolRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientFailureRate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientFailureRate"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientDecodeSuccessRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientDecodeSuccessRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientDecodeDurationByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientDecodeDurationByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientQueueDurationByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientQueueDurationByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientErrorCodeByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientErrorCodeByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientErrorCodeAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientErrorCodeAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientLoadDuration": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientLoadDuration"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientLoadDurationAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientLoadDurationAll"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientSdkVerByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientSdkVerByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientFileSize": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientFileSize"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientTopFileSize": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientTopFileSize"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientCountByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientCountByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientQualityRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientQualityRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientTopQualityURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientTopQualityURL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientDemotionRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientDemotionRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientTopDemotionURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientTopDemotionURL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXClientScoreByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXClientScoreByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSensibleCountByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSensibleCountByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSensibleCacheHitRateByTime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSensibleCacheHitRateByTime"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSensibleTopSizeURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSensibleTopSizeURL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSensibleTopResolutionURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSensibleTopResolutionURL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSensibleTopRamURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSensibleTopRamURL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DescribeImageXSensibleTopUnknownURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageXSensibleTopUnknownURL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateBatchProcessTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateBatchProcessTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetBatchProcessResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetBatchProcessResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetBatchTaskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetBatchTaskInfo"},
				"Version": []string{"2023-05-01"},
			},
		},
		"AIProcess": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AIProcess"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageAITask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageAITask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAITasks": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAITasks"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAIDetails": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAIDetails"},
				"Version": []string{"2023-05-01"},
			},
		},
		"ReportEvent": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ReportEvent"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageResourceStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageResourceStatus"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateFileStorageClass": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateFileStorageClass"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageStorageFiles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageStorageFiles"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteImageUploadFiles": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageUploadFiles"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateFileRestore": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateFileRestore"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageUploadFiles": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageUploadFiles"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CommitImageUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitImageUpload"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageFileCT": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageFileCT"},
				"Version": []string{"2018-08-01"},
			},
		},
		"ApplyVpcUploadInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyVpcUploadInfo"},
				"Version": []string{"2023-05-01"},
			},
		},
		"ApplyImageUpload": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyImageUpload"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageUploadFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUploadFile"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageUploadFiles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUploadFiles"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageUpdateFiles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUpdateFiles"},
				"Version": []string{"2023-05-01"},
			},
		},
		"PreviewImageUploadFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PreviewImageUploadFile"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageEraseResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEraseResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageService"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetAllImageServices": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllImageServices"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageCompressTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageCompressTask"},
				"Version": []string{"2018-08-01"},
			},
		},
		"FetchImageUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FetchImageUrl"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageStorageTTL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageStorageTTL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetCompressTaskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCompressTaskInfo"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetUrlFetchTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetUrlFetchTask"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetResourceURL": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResourceURL"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageFromUri": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageFromUri"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageFileKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageFileKey"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageContentTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageContentTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageContentTaskDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageContentTaskDetail"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageContentBlockList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageContentBlockList"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageTranscodeQueue": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageTranscodeQueue"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageTranscodeQueue": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageTranscodeQueue"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageTranscodeQueue": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageTranscodeQueue"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageTranscodeQueueStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageTranscodeQueueStatus"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageTranscodeQueues": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageTranscodeQueues"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageTranscodeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageTranscodeTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageTranscodeDetails": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageTranscodeDetails"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageTranscodeCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageTranscodeCallback"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageTranscodeDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageTranscodeDetail"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImagePSDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImagePSDetection"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageSuperResolutionResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageSuperResolutionResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetDenoisingImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDenoisingImage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageDuplicateDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageDuplicateDetection"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageOCRV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageOCRV2"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageBgFillResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageBgFillResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetSegmentImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSegmentImage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageSmartCropResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageSmartCropResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageComicResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageComicResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageEnhanceResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEnhanceResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageQuality": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageQuality"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetLicensePlateDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetLicensePlateDetection"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetPrivateImageType": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPrivateImageType"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateCVImageGenerateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCVImageGenerateTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateHiddenWatermarkImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateHiddenWatermarkImage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateHmExtractTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateHmExtractTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageExifData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageExifData"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageDetectResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageDetectResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetCVImageGenerateResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCVImageGenerateResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageHmExtract": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageHmExtract"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetCVTextGenerateImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCVTextGenerateImage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetCVImageGenerateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCVImageGenerateTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageHmEmbed": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageHmEmbed"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetCVAnimeGenerateImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCVAnimeGenerateImage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetComprehensiveEnhanceImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetComprehensiveEnhanceImage"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAiGenerateTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAiGenerateTask"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetProductAIGCResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetProductAIGCResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageEraseModels": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEraseModels"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetDedupTaskStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDedupTaskStatus"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageHmExtractTaskInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageHmExtractTaskInfo"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageService": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageService"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageService": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageService"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageAuthKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageAuthKey"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateResEventRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateResEventRule"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateServiceName": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateServiceName"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateStorageRules": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateStorageRules"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateStorageRulesV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateStorageRulesV2"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageObjectAccess": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageObjectAccess"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageUploadOverwrite": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageUploadOverwrite"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageMirrorConf": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageMirrorConf"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageServiceSubscription": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageServiceSubscription"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAuthKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAuthKey"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageAnalyzeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageAnalyzeTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageAnalyzeTaskRun": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageAnalyzeTaskRun"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageAnalyzeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageAnalyzeTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageAnalyzeTaskStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageAnalyzeTaskStatus"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageAnalyzeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageAnalyzeTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAnalyzeTasks": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAnalyzeTasks"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAnalyzeResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAnalyzeResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageElements": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageElements"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageBackgroundColors": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageBackgroundColors"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageStyle": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageStyle"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageStyle": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageStyle"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageStyleMeta": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageStyleMeta"},
				"Version": []string{"2023-05-01"},
			},
		},
		"AddImageElements": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddImageElements"},
				"Version": []string{"2023-05-01"},
			},
		},
		"AddImageBackgroundColors": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddImageBackgroundColors"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageStyle": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageStyle"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageFonts": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageFonts"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageElements": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageElements"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageBackgroundColors": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageBackgroundColors"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageStyles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageStyles"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageStyleDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageStyleDetail"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageStyleResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageStyleResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DownloadCert": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DownloadCert"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAllDomainCert": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAllDomainCert"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetCertInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCertInfo"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetAllCerts": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllCerts"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageTemplate"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteTemplatesFromBin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTemplatesFromBin"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageTemplate"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageTemplatesByImport": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageTemplatesByImport"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateTemplatesFromBin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTemplatesFromBin"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageTemplate"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetTemplatesFromBin": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTemplatesFromBin"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetAllImageTemplates": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllImageTemplates"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageAuditTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageAuditTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"DeleteImageAuditResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageAuditResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetSyncAuditResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSyncAuditResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageAuditTaskStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageAuditTaskStatus"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateImageAuditTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageAuditTask"},
				"Version": []string{"2023-05-01"},
			},
		},
		"UpdateAuditImageStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAuditImageStatus"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAuditTasks": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAuditTasks"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetImageAuditResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAuditResult"},
				"Version": []string{"2023-05-01"},
			},
		},
		"GetAuditEntrysCount": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAuditEntrysCount"},
				"Version": []string{"2023-05-01"},
			},
		},
		"CreateImageRetryAuditTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageRetryAuditTask"},
				"Version": []string{"2023-05-01"},
			},
		},
	}
)

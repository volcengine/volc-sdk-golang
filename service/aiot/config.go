package aiot

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20210101 = "2021-01-01"
	ServiceName            = "aiotvideo"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 30 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}
	ApiInfoList = map[string]*base.ApiInfo{
		"AddGroupTreeNode": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"AddGroupTreeNode"}, "Version": []string{ServiceVersion20210101}},
		},
		"AddSpaceDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"AddSpaceDomain"}, "Version": []string{ServiceVersion20210101}},
		},
		"AddVehicleImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"AddVehicleImage"}, "Version": []string{ServiceVersion20210101}},
		},
		"BatchDeleteCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"BatchDeleteCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"BatchDeleteCascadeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"BatchDeleteCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"BatchDeleteSlowLive": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"BatchDeleteSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"BindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"BindCert"}, "Version": []string{ServiceVersion20210101}},
		},
		"BindDeviceToGroupTreeNode": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"BindDeviceToGroupTreeNode"}, "Version": []string{ServiceVersion20210101}},
		},
		"CalculatePkg": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CalculatePkg"}, "Version": []string{ServiceVersion20210101}},
		},
		"CancelBindTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CancelBindTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"CertDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CertDetail"}, "Version": []string{ServiceVersion20210101}},
		},
		"CheckBindTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CheckBindTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"CloudControl": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CloudControl"}, "Version": []string{ServiceVersion20210101}},
		},
		"CloudRecordPlay": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CloudRecordPlay"}, "Version": []string{ServiceVersion20210101}},
		},
		"ConnectOBSServer": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ConnectOBSServer"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateAIApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateAIApp"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateAITemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateAITemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateAccount": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateAccount"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateBucket": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateBucket"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateCascadeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateDevice": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateDevice"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateForward": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateForward"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateLib": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateLib"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateLibInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateLibInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateRecordTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateRecordTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateScreenshotTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateScreenshotTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateSlowLive": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateStructuredView": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateStructuredView"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateStructuredViewCascadeJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateStructuredViewCascadeJob"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateStructuredViewCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateStructuredViewCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateStructuredViewCode": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateStructuredViewCode"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateStructuredViewSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateStructuredViewSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"CreateTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CreateTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"CruiseControl": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"CruiseControl"}, "Version": []string{ServiceVersion20210101}},
		},
		"DelVehicleImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DelVehicleImage"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteAIApp": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteAIApp"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteAITemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteAITemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteAlarmNotify": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteAlarmNotify"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteAlarmNotifyAll": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteAlarmNotifyAll"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteBucket": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteBucket"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteCascadeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteCert"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteDevice": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteDevice"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteForward": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteForward"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteGroupTreeNodes": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteGroupTreeNodes"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteLib": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteLib"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteLibInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteLibInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteLocalDownload": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteLocalDownload"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteRecordTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteRecordTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteScreenshotTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteScreenshotTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteSlowLive": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteSpaceDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteSpaceDomain"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteStreamLog": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteStreamLog"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteStructuredView": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteStructuredView"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteStructuredViewCascadeJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteStructuredViewCascadeJob"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteStructuredViewCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteStructuredViewCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteStructuredViewSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteStructuredViewSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"DeleteTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"DescribePushStreamMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DescribePushStreamMetrics"}, "Version": []string{ServiceVersion20210101}},
		},
		"DisableAuthInSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DisableAuthInSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"DisableStructuredView": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DisableStructuredView"}, "Version": []string{ServiceVersion20210101}},
		},
		"DisableStructuredViewCascadeJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DisableStructuredViewCascadeJob"}, "Version": []string{ServiceVersion20210101}},
		},
		"DisableStructuredViewCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DisableStructuredViewCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"DisableStructuredViewSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DisableStructuredViewSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"EnableStructuredView": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"EnableStructuredView"}, "Version": []string{ServiceVersion20210101}},
		},
		"EnableStructuredViewCascadeJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"EnableStructuredViewCascadeJob"}, "Version": []string{ServiceVersion20210101}},
		},
		"EnableStructuredViewCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"EnableStructuredViewCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"EnableStructuredViewSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"EnableStructuredViewSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"ForbidStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ForbidStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"FreshDevice": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"FreshDevice"}, "Version": []string{ServiceVersion20210101}},
		},
		"GenSipID": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GenSipID"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetAIAnalysisResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetAIAnalysisResult"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetAIApp": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetAIApp"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetAITemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetAITemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetAccount": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetAccount"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetCascadePlatform": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetCascadeTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetDataProjectWithBindWidthAndFlow": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetDataProjectWithBindWidthAndFlow"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetDevice": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetDevice"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetDeviceChannels": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetDeviceChannels"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetDeviceMaxDownloadSpeed": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetDeviceMaxDownloadSpeed"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetDevicesByGroupTreeNode": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetDevicesByGroupTreeNode"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetForward": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetForward"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetGBMediaFinish": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetGBMediaFinish"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetGroupNodesByCascadeTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetGroupNodesByCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetGroupTreeNode": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetGroupTreeNode"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetLib": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetLib"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetLibInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetLibInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetLocalDownload": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetLocalDownload"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetPkgInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetPkgInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetPushStreamCnt": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetPushStreamCnt"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetRecordList": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetRecordList"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetRecordTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetRecordTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetScreenshotTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetScreenshotTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetSlowLive": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetSpaceDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetSpaceDomain"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetSpaceTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetSpaceTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetStream": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetStreamData": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetStreamData"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetStructuredView": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetStructuredView"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetStructuredViewCascadeJob": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetStructuredViewCascadeJob"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetStructuredViewCascadePlatform": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetStructuredViewCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetStructuredViewDataInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetStructuredViewDataInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetStructuredViewSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetStructuredViewSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetTotalData": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetTotalData"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetTotalDeviceAndStream": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetTotalDeviceAndStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetVehicleAIAnalysisResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetVehicleAIAnalysisResult"}, "Version": []string{ServiceVersion20210101}},
		},
		"GetVehicleImage": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetVehicleImage"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListAIApp": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListAIApp"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListAITemplates": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListAITemplates"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListAlarmNotify": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListAlarmNotify"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListAppResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListAppResult"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListBucket": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListBucket"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListCascadePlatform": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListCascadeTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListCertificates": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListCertificates"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListDeviceRecords": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListDeviceRecords"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListDeviceScreenshots": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListDeviceScreenshots"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListDevices": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListDevices"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListForward": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListForward"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListGroupTreeNodes": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListGroupTreeNodes"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListLib": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListLib"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListLibInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListLibInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListLibItem": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListLibItem"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListRecordTemplates": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListRecordTemplates"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListScreenshotTemplates": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListScreenshotTemplates"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListSlowLive": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListSlowLiveStreams": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListSlowLiveStreams"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListSpaceDomains": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListSpaceDomains"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListSpaces": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListSpaces"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListStreams": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListStreams"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListStructuredViewCascadeJob": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListStructuredViewCascadeJob"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListStructuredViewCascadePlatform": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListStructuredViewCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListStructuredViewData": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListStructuredViewData"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListStructuredViewSpaces": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListStructuredViewSpaces"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListStructuredViews": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListStructuredViews"}, "Version": []string{ServiceVersion20210101}},
		},
		"ListTemplates": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ListTemplates"}, "Version": []string{ServiceVersion20210101}},
		},
		"LocalMediaDownload": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"LocalMediaDownload"}, "Version": []string{ServiceVersion20210101}},
		},
		"NewPackage": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"NewPackage"}, "Version": []string{ServiceVersion20210101}},
		},
		"PlayBackControl": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"PlayBackControl"}, "Version": []string{ServiceVersion20210101}},
		},
		"PlayBackStart": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"PlayBackStart"}, "Version": []string{ServiceVersion20210101}},
		},
		"PlayBackStop": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"PlayBackStop"}, "Version": []string{ServiceVersion20210101}},
		},
		"PlayStreamRes": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"PlayStreamRes"}, "Version": []string{ServiceVersion20210101}},
		},
		"PlayTs": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"PlayTs"}, "Version": []string{ServiceVersion20210101}},
		},
		"QueryPresetInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"QueryPresetInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"RefreshPlayBack": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"RefreshPlayBack"}, "Version": []string{ServiceVersion20210101}},
		},
		"ResetAlarm": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ResetAlarm"}, "Version": []string{ServiceVersion20210101}},
		},
		"SetAlarmGuard": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"SetAlarmGuard"}, "Version": []string{ServiceVersion20210101}},
		},
		"SetSpaceTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"SetSpaceTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"ShareGroupNodesToCascadeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"ShareGroupNodesToCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"SignSlowliveWsToken": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"SignSlowliveWsToken"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartAIApp": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartAIApp"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartCascadeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartDevice": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartDevice"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartForward": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartForward"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartSlowLive": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"StartVoiceTalk": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StartVoiceTalk"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopAIApp": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopAIApp"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopCascadeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopDevice": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopDevice"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopForward": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopForward"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopSlowLive": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"StopVoiceTalk": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StopVoiceTalk"}, "Version": []string{ServiceVersion20210101}},
		},
		"StreamLogs": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"StreamLogs"}, "Version": []string{ServiceVersion20210101}},
		},
		"UnbindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UnbindCert"}, "Version": []string{ServiceVersion20210101}},
		},
		"UnforbidStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UnforbidStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateAIApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateAIApp"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateAITemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateAITemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateAccount": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateAccount"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateAuthInSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateAuthInSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateBucket": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateBucket"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateCascadeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateCascadeTask"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateCert"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateDevice": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateDevice"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateDomainHTTPS": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateDomainHTTPS"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateGroupTreeNode": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateGroupTreeNode"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateLib": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateLib"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateLibInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateLibInfo"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateRecordTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateRecordTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateScreenshotTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateScreenshotTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateSlowLive": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateSlowLive"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateSlowLiveStreamConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateSlowLiveStreamConfig"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateSpaceDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateSpaceDomain"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateStream"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateStructuredView": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateStructuredView"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateStructuredViewCascadeJob": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateStructuredViewCascadeJob"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateStructuredViewCascadePlatform": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateStructuredViewCascadePlatform"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateStructuredViewSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateStructuredViewSpace"}, "Version": []string{ServiceVersion20210101}},
		},
		"UpdateTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UpdateTemplate"}, "Version": []string{ServiceVersion20210101}},
		},
		"UploadCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UploadCert"}, "Version": []string{ServiceVersion20210101}},
		},
		"UploadToken": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UploadToken"}, "Version": []string{ServiceVersion20210101}},
		},
		"UploadVehicleImage": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"UploadVehicleImage"}, "Version": []string{ServiceVersion20210101}},
		},
		"EdgeInvite": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"EdgeInvite"}, "Version": []string{ServiceVersion20210101}},
		},
		"EdgeBye": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"EdgeBye"}, "Version": []string{ServiceVersion20210101}},
		},
		"EdgeStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"EdgeStatus"}, "Version": []string{ServiceVersion20210101}},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

type QueryInterceptorFunc func(api string, v url.Values) url.Values

// IAM .
type AIoT struct {
	Client *base.Client

	CostCallback func(api string, cost int64, query url.Values, payload string) //todo 用于api调用耗时测试，正式使用需要删除

	queryInterceptors []QueryInterceptorFunc
}

// NewInstance 创建一个实例
func NewInstance() *AIoT {
	instance := &AIoT{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *AIoT) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *AIoT) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *AIoT) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *AIoT) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *AIoT) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}

func (p *AIoT) AddRequestInterceptor(interceptor QueryInterceptorFunc) {
	p.queryInterceptors = append(p.queryInterceptors, interceptor)
}

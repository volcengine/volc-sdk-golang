package acep

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "ACEP"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-north-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "open.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-north-1",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"ListContainerImages": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListContainerImages"},
				"Version": []string{"2022-08-01"},
			},
		},
		"DeleteContainerImages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteContainerImages"},
				"Version": []string{"2022-08-01"},
			},
		},
		"ImportContainerImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImportContainerImage"},
				"Version": []string{"2022-08-01"},
			},
		},
		"PullFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PullFile"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UpdatePodProperty": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePodProperty"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CloseApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CloseApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CreatePod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CreatePodOneStep": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePodOneStep"},
				"Version": []string{"2023-10-30"},
			},
		},
		"AddPropertyRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddPropertyRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DeletePod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"RunSyncCommand": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RunSyncCommand"},
				"Version": []string{"2023-10-30"},
			},
		},
		"LaunchApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"LaunchApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"LaunchApps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"LaunchApps"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PowerOffPod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PowerOffPod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PodStop": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PodStop"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PowerOnPod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PowerOnPod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CopyPod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CopyPod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PodDataTransfer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PodDataTransfer"},
				"Version": []string{"2023-10-30"},
			},
		},
		"RebootPod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RebootPod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ResetPod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResetPod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"BanUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BanUser"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PushFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PushFile"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DistributeFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DistributeFile"},
				"Version": []string{"2023-10-30"},
			},
		},
		"StartRecording": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRecording"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ScreenShot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ScreenShot"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PodAdb": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PodAdb"},
				"Version": []string{"2023-10-30"},
			},
		},
		"RunCommand": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RunCommand"},
				"Version": []string{"2023-10-30"},
			},
		},
		"BatchScreenShot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchScreenShot"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PodMute": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PodMute"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UpdatePod": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListDc": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDc"},
				"Version": []string{"2023-10-30"},
			},
		},
		"GetPodMetric": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPodMetric"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DetailPod": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailPod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListPod": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPod"},
				"Version": []string{"2023-10-30"},
			},
		},
		"GetPodProperty": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPodProperty"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListPropertyRule": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPropertyRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"PodDataDelete": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PodDataDelete"},
				"Version": []string{"2023-10-30"},
			},
		},
		"RemovePropertyRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemovePropertyRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"StopRecording": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRecording"},
				"Version": []string{"2023-10-30"},
			},
		},
		"GetPodAppList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPodAppList"},
				"Version": []string{"2023-10-30"},
			},
		},
		"SetProxy": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetProxy"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListTask"},
				"Version": []string{"2023-10-30"},
			},
		},
		"GetTaskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTaskInfo"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CreatePortMappingRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePortMappingRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListPortMappingRule": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPortMappingRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DetailPortMappingRule": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailPortMappingRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"BindPortMappingRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindPortMappingRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UnbindPortMappingRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindPortMappingRule"},
				"Version": []string{"2023-10-30"},
			},
		},
		"AttachTag": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AttachTag"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListTag": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListTag"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CreateTag": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTag"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DeleteTag": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTag"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UpdateTag": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTag"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UploadApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DetailApp": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UpdateApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListApp": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DeleteApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UninstallApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UninstallApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"InstallApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"InstallApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"InstallApps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"InstallApps"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListAppVersionDeploy": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAppVersionDeploy"},
				"Version": []string{"2023-10-30"},
			},
		},
		"AutoInstallApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AutoInstallApp"},
				"Version": []string{"2023-10-30"},
			},
		},
		"GetAppCrashLog": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAppCrashLog"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CreateDisplayLayoutMini": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDisplayLayoutMini"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DeleteDisplayLayout": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDisplayLayout"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListDisplayLayoutMini": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDisplayLayoutMini"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DetailDisplayLayoutMini": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailDisplayLayoutMini"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CreateAppImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateAppImage"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DetailAppVersionImage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailAppVersionImage"},
				"Version": []string{"2023-10-30"},
			},
		},
		"CreateImageOneStep": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageOneStep"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListImageResource": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListImageResource"},
				"Version": []string{"2023-10-30"},
			},
		},
		"GetImagePreheating": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImagePreheating"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListConfiguration": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListConfiguration"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListPodResourceSet": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPodResourceSet"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListPodResource": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPodResource"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListResourceQuota": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListResourceQuota"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UpdatePodResourceApplyNum": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePodResourceApplyNum"},
				"Version": []string{"2023-10-30"},
			},
		},
		"SubscribeResourceAuto": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubscribeResourceAuto"},
				"Version": []string{"2023-10-30"},
			},
		},
		"RenewResourceAuto": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RenewResourceAuto"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UpdateProductResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateProductResource"},
				"Version": []string{"2023-10-30"},
			},
		},
		"GetProductResource": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetProductResource"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UnsubscribeHostResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnsubscribeHostResource"},
				"Version": []string{"2023-10-30"},
			},
		},
		"UpdateHost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHost"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ListHost": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListHost"},
				"Version": []string{"2023-10-30"},
			},
		},
		"DetailHost": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailHost"},
				"Version": []string{"2023-10-30"},
			},
		},
		"RebootHost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RebootHost"},
				"Version": []string{"2023-10-30"},
			},
		},
		"ResetHost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResetHost"},
				"Version": []string{"2023-10-30"},
			},
		},
	}
)

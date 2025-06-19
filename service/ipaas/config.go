package ipaas

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "iPaaS"
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

		"UpdateHostUniversalInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHostUniversalInfo"},
				"Version": []string{"2023-08-01"},
			},
		},
		"UpdateInstanceUniversalInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateInstanceUniversalInfo"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ListAsyncAction": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAsyncAction"},
				"Version": []string{"2020-10-25"},
			},
		},
		"DetailProduct": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailProduct"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListInstancePackageBrief": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListInstancePackageBrief"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListProductIDCData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListProductIDCData"},
				"Version": []string{"2022-06-30"},
			},
		},
		"GetSecurityGroupVisibleConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSecurityGroupVisibleConfig"},
				"Version": []string{"2020-10-25"},
			},
		},
		"CreateSecurityGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSecurityGroup"},
				"Version": []string{"2020-10-25"},
			},
		},
		"CreateSecurityRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSecurityRule"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UpdateSecurityRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSecurityRule"},
				"Version": []string{"2020-10-25"},
			},
		},
		"BindInstanceSecurityGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindInstanceSecurityGroup"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UnbindInstanceSecurityGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindInstanceSecurityGroup"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ModifyInstanceWindowDisplaySpec": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyInstanceWindowDisplaySpec"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ModifyInstanceFps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyInstanceFps"},
				"Version": []string{"2023-08-01"},
			},
		},
		"PowerDownInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PowerDownInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ColdRebootInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ColdRebootInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UpgradeInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpgradeInstances"},
				"Version": []string{"2023-08-01"},
			},
		},
		"PowerUpInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PowerUpInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"AdbCommand": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AdbCommand"},
				"Version": []string{"2020-10-25"},
			},
		},
		"SetInstanceBandwidth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetInstanceBandwidth"},
				"Version": []string{"2023-08-01"},
			},
		},
		"InstallApplication": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"InstallApplication"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ControlApplication": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ControlApplication"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GetInstanceProperty": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetInstanceProperty"},
				"Version": []string{"2020-10-25"},
			},
		},
		"PullFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PullFile"},
				"Version": []string{"2020-10-25"},
			},
		},
		"PushFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PushFile"},
				"Version": []string{"2020-10-25"},
			},
		},
		"RecordScreen": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RecordScreen"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ExecCmdSync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ExecCmdSync"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ResetInstanceToFactory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResetInstanceToFactory"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UpdateInstanceProperty": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateInstanceProperty"},
				"Version": []string{"2020-10-25"},
			},
		},
		"WarmRebootInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WarmRebootInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ResetInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResetInstances"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GetInstanceProperties": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetInstanceProperties"},
				"Version": []string{"2020-10-25"},
			},
		},
		"SetInstanceProperties": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetInstanceProperties"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GetJobDetails": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetJobDetails"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListTaskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListTaskInfo"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UpgradeImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpgradeImage"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListAOSPRepoAclEntries": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAOSPRepoAclEntries"},
				"Version": []string{"2020-10-25"},
			},
		},
		"AddAOSPRepoAclEntries": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddAOSPRepoAclEntries"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GrantGitRepoPermission": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GrantGitRepoPermission"},
				"Version": []string{"2020-10-25"},
			},
		},
		"RemoveAOSPRepoAclEntries": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemoveAOSPRepoAclEntries"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GetGitRepoUserInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetGitRepoUserInfo"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListHost": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListHost"},
				"Version": []string{"2020-10-25"},
			},
		},
		"InitializeHost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"InitializeHost"},
				"Version": []string{"2023-08-01"},
			},
		},
		"GetInfoAfterOrder": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetInfoAfterOrder"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ListHostMetricData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListHostMetricData"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ReconfigureDevicesPackage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ReconfigureDevicesPackage"},
				"Version": []string{"2023-08-01"},
			},
		},
		"RebootHost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RebootHost"},
				"Version": []string{"2020-10-25"},
			},
		},
		"DistributeFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DistributeFile"},
				"Version": []string{"2020-10-25"},
			},
		},
		"DistributeFileToInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DistributeFileToInstances"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GetFileDistributionJobDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetFileDistributionJobDetail"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GetFileDistributionResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetFileDistributionResult"},
				"Version": []string{"2020-10-25"},
			},
		},
		"FixInstancesSGBound": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FixInstancesSGBound"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListSecurityGroup": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSecurityGroup"},
				"Version": []string{"2020-10-25"},
			},
		},
		"DetailSecurityGroup": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailSecurityGroup"},
				"Version": []string{"2020-10-25"},
			},
		},
		"BindInstancesSecurityGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindInstancesSecurityGroup"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UnbindInstancesSecurityGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindInstancesSecurityGroup"},
				"Version": []string{"2020-10-25"},
			},
		},
		"CreateDevices": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDevices"},
				"Version": []string{"2023-08-01"},
			},
		},
		"DeleteDevices": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDevices"},
				"Version": []string{"2023-08-01"},
			},
		},
		"TosPreSignUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TosPreSignUrl"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GetResourceNetworkCurveConsole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResourceNetworkCurveConsole"},
				"Version": []string{"2023-08-01"},
			},
		},
		"GetResourcePodCurrentConsole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResourcePodCurrentConsole"},
				"Version": []string{"2023-08-01"},
			},
		},
		"GetResourcePodCurveConsole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResourcePodCurveConsole"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ListInstance": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ExportInstance": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ExportInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListInstanceMetricData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListInstanceMetricData"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ListPortMapping": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPortMapping"},
				"Version": []string{"2020-10-25"},
			},
		},
		"DetailInstance": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetailInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"LatestMetricInstance": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"LatestMetricInstance"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListContainerImagesInner": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListContainerImagesInner"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ImportContainerImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImportContainerImage"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListContainerImages": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListContainerImages"},
				"Version": []string{"2020-10-25"},
			},
		},
		"DeleteContainerImages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteContainerImages"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UpdateContainerImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateContainerImage"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListAdbKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAdbKey"},
				"Version": []string{"2020-10-25"},
			},
		},
		"BindInstancesAdbKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindInstancesAdbKey"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UnbindInstancesAdbKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindInstancesAdbKey"},
				"Version": []string{"2020-10-25"},
			},
		},
		"BindInstanceAdbKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindInstanceAdbKey"},
				"Version": []string{"2020-10-25"},
			},
		},
		"UnbindInstanceAdbKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindInstanceAdbKey"},
				"Version": []string{"2020-10-25"},
			},
		},
		"GenerateAdbKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GenerateAdbKey"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListGitRepoWhiteIP": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListGitRepoWhiteIP"},
				"Version": []string{"2023-08-01"},
			},
		},
		"DeleteGitRepoWhiteIP": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteGitRepoWhiteIP"},
				"Version": []string{"2023-08-01"},
			},
		},
		"OperateGitRepoWhiteIP": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OperateGitRepoWhiteIP"},
				"Version": []string{"2023-08-01"},
			},
		},
		"AddGitRepoWhiteIP": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddGitRepoWhiteIP"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ListGitRepoSSHKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListGitRepoSSHKey"},
				"Version": []string{"2023-08-01"},
			},
		},
		"DeleteGitRepoSSHKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteGitRepoSSHKey"},
				"Version": []string{"2023-08-01"},
			},
		},
		"OperateGitRepoSSHKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OperateGitRepoSSHKey"},
				"Version": []string{"2023-08-01"},
			},
		},
		"AddGitRepoSSHKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddGitRepoSSHKey"},
				"Version": []string{"2023-08-01"},
			},
		},
		"GenerateGitRepoSSHKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GenerateGitRepoSSHKey"},
				"Version": []string{"2023-08-01"},
			},
		},
		"AcquireIdempotentToken": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AcquireIdempotentToken"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ListProduct": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListProduct"},
				"Version": []string{"2020-10-25"},
			},
		},
		"AddEventCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddEventCallback"},
				"Version": []string{"2020-10-25"},
			},
		},
		"DelEventCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DelEventCallback"},
				"Version": []string{"2020-10-25"},
			},
		},
		"SetEventCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetEventCallback"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListEventCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListEventCallback"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ValidateEventCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ValidateEventCallback"},
				"Version": []string{"2020-10-25"},
			},
		},
		"ListPackage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPackage"},
				"Version": []string{"2023-08-01"},
			},
		},
		"ListDcCapacity": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDcCapacity"},
				"Version": []string{"2023-08-01"},
			},
		},
	}
)

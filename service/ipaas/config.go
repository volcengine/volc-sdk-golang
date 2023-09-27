package ipaas

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "IPaaS"
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
		"ModifyInstanceWindowDisplaySpec": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyInstanceWindowDisplaySpec"},
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
		"GetJobDetails": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetJobDetails"},
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
		"ListHost": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListHost"},
				"Version": []string{"2020-10-25"},
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
		"RebootHost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RebootHost"},
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
		"ListPackage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPackage"},
				"Version": []string{"2023-08-01"},
			},
		},
	}
)

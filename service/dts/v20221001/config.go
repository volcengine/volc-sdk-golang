package dts_v20221001

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "dts"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-beijing": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "dts.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-beijing",
				Service: ServiceName,
			},
		},
		"cn-guangzhou": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "dts.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-guangzhou",
				Service: ServiceName,
			},
		},
		"cn-shanghai": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "dts.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-shanghai",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"CreateDataValidationTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDataValidationTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"CreateSubscriptionGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSubscriptionGroup"},
				"Version": []string{"2022-10-01"},
			},
		},
		"CreateTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DeleteDataValidationTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDataValidationTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DeleteSubscriptionGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSubscriptionGroup"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DeleteTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DeleteTransmissionTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTransmissionTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeDataValidationResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDataValidationResult"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeDataValidationTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDataValidationTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeSubscriptionGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeSubscriptionGroup"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeSubscriptionGroupProgress": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeSubscriptionGroupProgress"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeSubscriptionGroups": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeSubscriptionGroups"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeTransmissionTaskInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeTransmissionTaskInfo"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeTransmissionTaskProgress": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeTransmissionTaskProgress"},
				"Version": []string{"2022-10-01"},
			},
		},
		"DescribeTransmissionTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeTransmissionTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"ModifyTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"ResumeDataValidationTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeDataValidationTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"ResumeTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"ResumeTransmissionTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeTransmissionTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"RetryTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RetryTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"RetryTransmissionTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RetryTransmissionTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"SetBiSyncDDLSource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetBiSyncDDLSource"},
				"Version": []string{"2022-10-01"},
			},
		},
		"StartTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"StartTransmissionTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartTransmissionTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"StopDataValidationTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopDataValidationTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"StopTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"StopTransmissionTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopTransmissionTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"SuspendDataValidationTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SuspendDataValidationTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"SuspendTransmissionTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SuspendTransmissionTask"},
				"Version": []string{"2022-10-01"},
			},
		},
		"SuspendTransmissionTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SuspendTransmissionTasks"},
				"Version": []string{"2022-10-01"},
			},
		},
		"UpdateSubscriptionGroup": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSubscriptionGroup"},
				"Version": []string{"2022-10-01"},
			},
		},
	}
)

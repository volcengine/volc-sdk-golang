package rtc

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

// RTC ... use base clie

const (
	DefaultRegion          = "cn-north-1"
	DefaultTimeout         = 10 * time.Second
	ServiceVersion20201201 = "2020-12-01"
	ServiceVersion20220601 = "2022-06-01"
	ServiceName            = "rtc"
	ServiceHost            = "rtc.volcengineapi.com"

	// action name
	ActionStartRecord      = "StartRecord"
	ActionGetRecordTask    = "GetRecordTask"
	ActionStartWebRecord   = "StartWebRecord"
	ActionStopWebRecord    = "StopWebRecord"
	ActionGetWebRecordTask = "GetWebRecordTask"
	ActionGetWebRecordList = "GetWebRecordList"
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-north-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "rtc.volcengineapi.com",
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

		"StartWBRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWBRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"WbTranscodeCreate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeCreate"},
				"Version": []string{"2020-12-01"},
			},
		},
		"WbTranscodeQuery": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeQuery"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopWBRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopWBRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"WbTranscodeGet": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeGet"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRecord"},
				"Version": []string{"2022-06-01"},
			},
		},
		"GetRecordTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecordTask"},
				"Version": []string{"2022-06-01"},
			},
		},
		"StartWebRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecordTask"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopWebRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopWebRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetWebRecordTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWebRecordTask"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetWebRecordList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWebRecordList"},
				"Version": []string{"2020-12-01"},
			},
		},
		//ActionExample, add new action
		/*
			"ActionExample": {
				Method: http.MethodGet,
				Path:   "/",
				Query: url.Values{
					"Action":  []string{"ActionExample"},
					"Version": []string{ServiceVersion20201201},
				},
			},
		*/
	}
)

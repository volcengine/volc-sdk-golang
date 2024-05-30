package rtc_v20220601

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "rtc"
	DefaultTimeout = 10 * time.Second
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

		"StopSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopSnapshot"},
				"Version": []string{"2022-06-01"},
			},
		},
		"StartSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartSnapshot"},
				"Version": []string{"2022-06-01"},
			},
		},
		"StartSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartSegment"},
				"Version": []string{"2022-06-01"},
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
		"UpdateRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRecord"},
				"Version": []string{"2022-06-01"},
			},
		},
		"UpdateSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshot"},
				"Version": []string{"2022-06-01"},
			},
		},
		"UpdateSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSegment"},
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
		"GetSnapshotTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSnapshotTask"},
				"Version": []string{"2022-06-01"},
			},
		},
		"GetSegmentTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSegmentTask"},
				"Version": []string{"2022-06-01"},
			},
		},
		"StopRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRecord"},
				"Version": []string{"2022-06-01"},
			},
		},
		"StopSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopSegment"},
				"Version": []string{"2022-06-01"},
			},
		},
		"GetRealtimePostProcessing": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRealtimePostProcessing"},
				"Version": []string{"2022-06-01"},
			},
		},
	}
)

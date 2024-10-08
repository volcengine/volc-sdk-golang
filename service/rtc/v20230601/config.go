package rtc_v20230601

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
		"ap-singapore-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "open-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "ap-singapore-1",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"StartRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRecord"},
				"Version": []string{"2023-06-01"},
			},
		},
		"StartPushSingleStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushSingleStreamToCDN"},
				"Version": []string{"2023-06-01"},
			},
		},
		"StartPushMixedStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushMixedStreamToCDN"},
				"Version": []string{"2023-06-01"},
			},
		},
		"UpdateRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRecord"},
				"Version": []string{"2023-06-01"},
			},
		},
		"UpdatePushMixedStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePushMixedStreamToCDN"},
				"Version": []string{"2023-06-01"},
			},
		},
		"GetPushSingleStreamToCDNTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPushSingleStreamToCDNTask"},
				"Version": []string{"2023-06-01"},
			},
		},
		"GetPushMixedStreamToCDNTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPushMixedStreamToCDNTask"},
				"Version": []string{"2023-06-01"},
			},
		},
		"GetRecordTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecordTask"},
				"Version": []string{"2023-06-01"},
			},
		},
		"StopRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRecord"},
				"Version": []string{"2023-06-01"},
			},
		},
		"StopPushStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPushStreamToCDN"},
				"Version": []string{"2023-06-01"},
			},
		},
	}
)

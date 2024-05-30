package rtc_v20230720

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

		"SendBroadcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendBroadcast"},
				"Version": []string{"2023-07-20"},
			},
		},
		"SendRoomUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendRoomUnicast"},
				"Version": []string{"2023-07-20"},
			},
		},
		"SendUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendUnicast"},
				"Version": []string{"2023-07-20"},
			},
		},
		"BatchSendRoomUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchSendRoomUnicast"},
				"Version": []string{"2023-07-20"},
			},
		},
	}
)

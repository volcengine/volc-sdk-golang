package dts_v20180101

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

		"GetAsyncPreCheckResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAsyncPreCheckResult"},
				"Version": []string{"2018-01-01"},
			},
		},
		"PreCheckAsync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PreCheckAsync"},
				"Version": []string{"2018-01-01"},
			},
		},
	}
)

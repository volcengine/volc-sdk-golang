package vms

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultTimeout = 10 * time.Second
)

type Risk struct {
	*base.Client
}

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

func NewInstance() *Risk {
	instance := &Risk{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: DefaultTimeout,
			Host:    "cloud-vms.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: "volc_risk_http",
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"QueryCanCall": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryCanCall"},
				"Version": []string{"2021-01-01"},
			},
		},
	}
)

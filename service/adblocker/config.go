package adblocker

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type AdBlocker struct {
	*base.Client
	retry bool
}

func (p *AdBlocker) Retry() bool {
	return p.retry
}

func (p *AdBlocker) CloseRetry() {
	p.retry = false
}

var DefaultInstance = NewInstance()

func NewInstance() *AdBlocker {
	instance := &AdBlocker{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
		retry:  false,
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "open.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "AdBlocker"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"AdBlock": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AdBlock"},
				"Version": []string{"2021-01-06"},
			},
		},
	}
)

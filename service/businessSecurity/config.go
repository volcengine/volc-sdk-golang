package businessSecurity

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type BusinessSecurity struct {
	*base.Client
	retry bool
}

var DefaultInstance = NewInstance()

func NewInstance() *BusinessSecurity {
	instance := &BusinessSecurity{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
		retry:  true,
	}
	return instance
}

func (p *BusinessSecurity) Retry() bool {
	return p.retry
}

func (p *BusinessSecurity) CloseRetry() {
	p.retry = false
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "open.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "BusinessSecurity"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"RiskDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RiskDetection"},
				"Version": []string{"2021-02-02"},
			},
		},
	}
)

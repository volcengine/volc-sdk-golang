package businessSecurity

import (
	"fmt"
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

func (p *BusinessSecurity) SetRegion(region string) error {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		return fmt.Errorf("region does not spport or unknown region")
	}
	p.ServiceInfo = serviceInfo
	p.SetScheme("http")
	return nil
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "riskcontrol.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "BusinessSecurity"},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "open-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "BusinessSecurity"},
		},
		base.RegionUsEast1: {
			Timeout: 5 * time.Second,
			Host:    "open-us-east-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionUsEast1, Service: "BusinessSecurity"},
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
		"AsyncRiskDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncRiskDetection"},
				"Version": []string{"2021-02-25"},
			},
		},
		"RiskResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RiskResult"},
				"Version": []string{"2021-03-10"},
			},
		},
		"DataReport": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DataReport"},
				"Version": []string{"2021-08-31"},
			},
		},
		"ElementVerify": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ElementVerify"},
				"Version": []string{"2021-11-23"},
			},
		},
		"MobileStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MobileStatus"},
				"Version": []string{"2020-12-25"},
			},
		},
	}
)

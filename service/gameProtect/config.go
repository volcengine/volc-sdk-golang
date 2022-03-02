package gameProtect

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type GameProtector struct {
	*base.Client
	retry bool
}

func (p *GameProtector) Retry() bool {
	return p.retry
}

func (p *GameProtector) CloseRetry() {
	p.retry = false
}

var DefaultInstance = NewInstance()

func NewInstance() *GameProtector {
	instance := &GameProtector{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
		retry:  false,
	}
	return instance
}

func (p *GameProtector) SetRegion(region string) error {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		return fmt.Errorf("region does not spport or unknown region")
	}
	p.ServiceInfo = serviceInfo
	return nil
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "open.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "game_protect"},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "open-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "game_protect"},
		},
		base.RegionUsEast1: {
			Timeout: 5 * time.Second,
			Host:    "open-us-east-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionUsEast1, Service: "game_protect"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"RiskResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RiskResult"},
				"Version": []string{"2021-04-25"},
			},
		},
	}
)

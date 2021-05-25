package gameProtect

import (
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
	instance := &GameProtector {
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
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "game_protect"},
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

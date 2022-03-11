package imp

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

type Imp struct {
	*base.Client
	DomainCache map[string]map[string]int
	Lock        sync.RWMutex
}

func NewInstance() *Imp {
	instance := &Imp{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *Imp {
	var serviceInfo *base.ServiceInfo
	var ok bool
	if serviceInfo, ok = ServiceInfoMap[region]; !ok {
		panic("Cant find the region, please check it carefully")
	}

	instance := &Imp{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

const (
	Host = "open.volcengineapi.com"
)

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    Host,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "imp"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"SubmitJob": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitJob"},
				"Version": []string{"2021-06-11"},
			},
		},
		"KillJob": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"KillJob"},
				"Version": []string{"2021-06-11"},
			},
		},
		"RetrieveJob": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RetrieveJob"},
				"Version": []string{"2021-06-11"},
			},
		},
	}
)

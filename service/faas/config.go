package faas

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	defaultRegion          = "cn-beijing"
	ServiceVersion20220101 = "2021-03-03"
	serviceName            = "vefaas"
)

var (
	serviceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "volcengineapi.byted.org",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	apiInfoList = map[string]*base.ApiInfo{
		"ListTriggers": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListTriggers"},
				"Version": []string{ServiceVersion20220101},
			},
		},
	}
)

type FaaS struct {
	*base.Client
}

// NewInstance creates new instance
func NewInstance() *FaaS {
	instance := &FaaS{}
	instance.Client = base.NewClient(serviceInfo, apiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = serviceName
	instance.Client.ServiceInfo.Credentials.Region = defaultRegion
	return instance
}

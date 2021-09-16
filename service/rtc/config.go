package rtc

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

// RTC ... use base client
type RTC struct {
	Client *base.Client
}

// NewInstance ...
func NewInstance() *RTC {
	instance := &RTC{}
	instance.Client = base.NewClient(ServiceInfo, DefaultApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20201201 = "2020-12-01"
	ServiceName            = "rtc"
	ServiceHost            = "open.volcengineapi.com"

	// action name
	ActionListRooms      = "ListRooms"
	ActionListIndicators = "ListIndicators"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    ServiceHost,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	DefaultApiInfoList = map[string]*base.ApiInfo{
		ActionListRooms: {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionListRooms},
				"Version": []string{ServiceVersion20201201},
			},
		},
		ActionListIndicators: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionListIndicators},
				"Version": []string{ServiceVersion20201201},
			},
		},
		//ActionExample, add new action
		/*
		"ActionExample": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ActionExample"},
				"Version": []string{ServiceVersion20201201},
			},
		},
		*/
	}
)

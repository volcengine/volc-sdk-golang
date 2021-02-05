package vedit

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultTimeout = 10 * time.Second
)

type VEdit struct {
	*base.Client
}

func NewInstance() *VEdit {
	instance := &VEdit{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: DefaultTimeout,
			Host:    "vedit.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: "edit",
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"SubmitDirectEditTaskAsync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitDirectEditTaskAsync"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetDirectEditResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDirectEditResult"},
				"Version": []string{"2018-01-01"},
			},
		},
		"SubmitTemplateTaskAsync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitTemplateTaskAsync"},
				"Version": []string{"2018-01-01"},
			},
		},
	}
)

package secretnumber

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultTimeout = 10 * time.Second
)

type SecretNumber struct {
	*base.Client
}

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

func NewInstance() *SecretNumber {
	instance := &SecretNumber{
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
				Service: "volc_secret_number",
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"BindAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindAXB"},
				"Version": []string{"2020-09-01"},
			},
		},
		"SelectNumberAndBindAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SelectNumberAndBindAXB"},
				"Version": []string{"2020-09-01"},
			},
		},
		"UnbindAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindAXB"},
				"Version": []string{"2020-09-01"},
			},
		},
		"QuerySubscription": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QuerySubscription"},
				"Version": []string{"2020-09-01"},
			},
		},
		"QuerySubscriptionForList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QuerySubscriptionForList"},
				"Version": []string{"2020-09-01"},
			},
		},
		"UpgradeAXToAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpgradeAXToAXB"},
				"Version": []string{"2020-09-01"},
			},
		},
		"UpdateAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAXB"},
				"Version": []string{"2020-09-01"},
			},
		},
		"BindAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindAXN"},
				"Version": []string{"2020-09-01"},
			},
		},
		"UpdateAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAXN"},
				"Version": []string{"2020-09-01"},
			},
		},
		"UnbindAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindAXN"},
				"Version": []string{"2020-09-01"},
			},
		},
	}
)

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

type DataCenter struct {
	*base.Client
}

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// DefaultDataCenterInstance 默认的实例
var DefaultDataCenterInstance = NewDataCenterInstance()

func NewInstance() *SecretNumber {
	instance := &SecretNumber{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewDataCenterInstance() *DataCenter {
	instance := &DataCenter{
		Client: base.NewClient(DataCenterServiceInfoMap[base.RegionCnNorth1], DataCenterApiInfoList),
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

	DataCenterServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: DefaultTimeout,
			Host:    "cloud-vms.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: "volc_datacenter_http",
			},
		},
	}

	DataCenterApiInfoList = map[string]*base.ApiInfo{
		"QueryAudioRecordFileUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryAudioRecordFileUrl"},
				"Version": []string{"2020-09-01"},
			},
		},
		"QueryAudioRecordToTextFileUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryAudioRecordToTextFileUrl"},
				"Version": []string{"2021-01-01"},
			},
		},
		"QueryCallRecordMsg": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryCallRecordMsg"},
				"Version": []string{"2021-01-01"},
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
		"SelectNumberAndBindAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SelectNumberAndBindAXN"},
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
		"Click2Call": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Click2Call"},
				"Version": []string{"2021-09-01"},
			},
		},
	}
)

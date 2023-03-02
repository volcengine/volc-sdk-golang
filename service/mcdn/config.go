package mcdn

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion  = base.RegionCnNorth1
	ServiceName    = "mcdn"
	ServiceVersion = "2022-03-01"
)

var (
	serviceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Host:    "open.volcengineapi.com",
			Timeout: 10 * time.Minute,
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
				"X-Upstream":   []string{"volcano"},
			},
		},
	}

	apiInfo = map[string]*base.ApiInfo{
		"ListCloudAccounts": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCloudAccounts"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListCdnDomains": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnDomains"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListDnsSchedules": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDnsSchedules"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeDnsSchedule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDnsSchedule"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnData"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnOriginData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnOriginData"},
				"Version": []string{ServiceVersion},
			},
		},
		"SubmitRefreshTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitRefreshTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"SubmitPreloadTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitPreloadTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListContentTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListContentTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeContentTaskByTaskId": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeContentTaskByTaskId"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListVendorContentTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVendorContentTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeContentQuota": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeContentQuota"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListViews": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListViews"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnRegionAndIsp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnRegionAndIsp"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnDataOffline": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnDataOffline"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnOriginDataOffline": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnOriginDataOffline"},
				"Version": []string{ServiceVersion},
			},
		},
	}
)

// DefaultInstance Package level default instance
var DefaultInstance = NewInstance()

type MCDN struct {
	Client *base.Client
}

func NewInstance() *MCDN {
	instance := &MCDN{
		Client: base.NewClient(serviceInfo[DefaultRegion], apiInfo),
	}
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

func (s *MCDN) GetServiceInfo(region string) *base.ServiceInfo {
	if serviceInfo, ok := serviceInfo[region]; ok {
		return serviceInfo
	}
	return nil
}

func (s *MCDN) GetAPIInfo(api string) *base.ApiInfo {
	if inf, ok := apiInfo[api]; ok {
		return inf
	}
	return nil
}

// SetHost .
func (s *MCDN) SetRegion(region string) {
	s.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (s *MCDN) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *MCDN) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}

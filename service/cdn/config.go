package cdn

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultRegion  = "cn-north-1"
	ServiceVersion = "2021-03-01"
	ServiceName    = "CDN"
)

var (
	ServiceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Timeout: 30 * time.Second,
			Host:    "cdn.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
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
		"DescribeContentTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeContentTasks"},
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
		"DescribeCdnData": {
			Timeout: 60 * time.Second,
			Method:  http.MethodPost,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnData"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnOriginData": {
			Timeout: 60 * time.Second,
			Method:  http.MethodPost,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnOriginData"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnRegionAndIsp": {
			Timeout: 60 * time.Second,
			Method:  http.MethodPost,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnRegionAndIsp"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnDomainTopData": {
			Timeout: 60 * time.Second,
			Method:  http.MethodPost,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnDomainTopData"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnAccountingData": {
			Timeout: 60 * time.Second,
			Method:  http.MethodPost,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnAccountingData"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnAccessLog": {
			Timeout: 60 * time.Second,
			Method:  http.MethodPost,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnAccessLog"},
				"Version": []string{ServiceVersion},
			},
		},
		"StartCdnDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartCdnDomain"},
				"Version": []string{ServiceVersion},
			},
		},
		"StopCdnDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopCdnDomain"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteCdnDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCdnDomain"},
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
		"DescribeCdnUpperIp": {
			Timeout: 60 * time.Second,
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnUpperIp"},
				"Version": []string{ServiceVersion},
			},
		},
	}
)

// DefaultInstance Package level default instance
var DefaultInstance = NewInstance()

type CDN struct {
	Client *base.Client
}

func NewInstance() *CDN {
	instance := new(CDN)
	instance.Client = base.NewClient(ServiceInfo[DefaultRegion], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

func (s *CDN) GetServiceInfo(region string) *base.ServiceInfo {
	if serviceInfo, ok := ServiceInfo[region]; ok {
		return serviceInfo
	}
	return nil
}

func (s *CDN) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

func (s *CDN) SetRegion(region string) {
	if serviceInfo := s.GetServiceInfo(region); serviceInfo != nil {
		serviceInfo.Credentials.Region = region
	}
}

func (s *CDN) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *CDN) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}

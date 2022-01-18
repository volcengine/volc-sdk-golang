package cdn

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
)

const (
	DefaultRegion  = "cn-north-1"
	ServiceVersion = "2021-03-01"
	ServiceName    = "CDN"
)

var (
	ServiceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Host: "cdn.volcengineapi.com",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"AddCdnDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddCdnDomain"},
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
		"DescribeCdnConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnConfig"},
				"Version": []string{ServiceVersion},
			},
		},
		"UpdateCdnConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCdnConfig"},
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
		"DescribeEdgeNrtDataSummary": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeEdgeNrtDataSummary"},
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
		"DescribeOriginNrtDataSummary": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeOriginNrtDataSummary"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnDataDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnDataDetail"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeEdgeStatisticalData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeEdgeStatisticalData"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeEdgeTopNrtData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeEdgeTopNrtData"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeEdgeTopStatisticalData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeEdgeTopStatisticalData"},
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
		"DescribeCdnService": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnService"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeAccountingData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeAccountingData"},
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
		"SubmitBlockTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitBlockTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"SubmitUnblockTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitUnblockTask"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeContentBlockTasks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeContentBlockTasks"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnAccessLog": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnAccessLog"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeIPInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeIPInfo"},
				"Version": []string{ServiceVersion},
			},
		},
		"DescribeCdnUpperIp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCdnUpperIp"},
				"Version": []string{ServiceVersion},
			},
		},
		"AddResourceTags": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddResourceTags"},
				"Version": []string{ServiceVersion},
			},
		},
		"UpdateResourceTags": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateResourceTags"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListResourceTags": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListResourceTags"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteResourceTags": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteResourceTags"},
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

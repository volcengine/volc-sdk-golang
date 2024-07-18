package cdn

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
)

var ApiInfoList = map[string]*base.ApiInfo{
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
	"DescribeDistrictIspData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeDistrictIspData"},
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
	"DescribeOriginTopNrtData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeOriginTopNrtData"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeEdgeTopStatusCode": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeEdgeTopStatusCode"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeOriginTopStatusCode": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeOriginTopStatusCode"},
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
	"DescribeIPListInfo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeIPListInfo"},
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
	"AddCdnCertificate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"AddCdnCertificate"},
			"Version": []string{ServiceVersion},
		},
	},
	"ListCertInfo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"ListCertInfo"},
			"Version": []string{ServiceVersion},
		},
	},
	"ListCdnCertInfo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"ListCdnCertInfo"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeCertConfig": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeCertConfig"},
			"Version": []string{ServiceVersion},
		},
	},
	"BatchDeployCert": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"BatchDeployCert"},
			"Version": []string{ServiceVersion},
		},
	},
	"DeleteCdnCertificate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DeleteCdnCertificate"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeAccountingSummary": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeAccountingSummary"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeDistrictData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeDistrictData"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeEdgeData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeEdgeData"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeDistrictSummary": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeDistrictSummary"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeEdgeSummary": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeEdgeSummary"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeOriginData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeOriginData"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeOriginSummary": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeOriginSummary"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeUserData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeUserData"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeDistrictRanking": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeDistrictRanking"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeEdgeRanking": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeEdgeRanking"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeOriginRanking": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeOriginRanking"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeEdgeStatusCodeRanking": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeEdgeStatusCodeRanking"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeOriginStatusCodeRanking": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeOriginStatusCodeRanking"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeStatisticalRanking": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeStatisticalRanking"},
			"Version": []string{ServiceVersion},
		},
	},
}

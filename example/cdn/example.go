package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
	"time"
)

// Warning: these tests may fail when the test interval is less than 2 minutes due to the configuring.

var (
	DefaultInstance = cdn.DefaultInstance
	ak              = "ak"
	sk              = "sk"
	typeFile        = "file"
	testDomain1     = "test1.com"
	testDomain2     = "test2.com"
	testDomain3     = "test3.com"
	now             = time.Now()
	testStartTime   = now.Unix()
	testEndTime     = now.Add(time.Minute * 10).Unix()
	exampleDomain   = "example.com"
)

func init() {
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
}

// 域名操作

func AddCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.AddCdnDomain(&cdn.AddCdnDomainRequest{
		Domain:      testDomain2,
		ServiceType: "web",
		Origin: []cdn.OriginRule{
			{OriginAction: cdn.OriginAction{
				OriginLines: []cdn.OriginLine{
					{
						OriginType:   "primary",
						InstanceType: "ip",
						Address:      "1.1.1.1",
						HttpPort:     "80",
						HttpsPort:    "80",
						Weight:       "100",
					},
				},
			}},
		},
		OriginProtocol: "http",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainRequest{Domain: testDomain1})
	if err != nil {
		resp, err = DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainRequest{Domain: testDomain1})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainRequest{Domain: testDomain2})
	if err != nil {
		resp, err = DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainRequest{Domain: testDomain2})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func DeleteDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&cdn.DeleteCdnDomainRequest{Domain: testDomain2})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&cdn.ListCdnDomainsRequest{Domain: &testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}

// 域名配置

func DescribeCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnConfig(&cdn.DescribeCdnConfigRequest{
		Domain: testDomain2,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.Result.DomainConfig)
	domain := resp.Result.DomainConfig
	fmt.Printf("%+v\n", domain)
}

func UpdateCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.UpdateCdnConfig(&cdn.UpdateCdnConfigRequest{
		Domain: testDomain2,
		Origin: []cdn.OriginRule{
			{OriginAction: cdn.OriginAction{
				OriginLines: []cdn.OriginLine{
					{
						OriginType:   "primary",
						InstanceType: "ip",
						Address:      "1.1.1.1",
						HttpPort:     "80",
						HttpsPort:    "80",
						Weight:       "100",
					},
				},
			}},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

// 数据统计

func DescribeCdnData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnData(&cdn.DescribeCdnDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func DescribeEdgeNrtDataSummary(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeNrtDataSummary(&cdn.DescribeEdgeNrtDataSummaryRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func DescribeCdnOriginData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnOriginData(&cdn.DescribeCdnOriginDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func DescribeOriginNrtDataSummary(t *testing.T) {
	resp, err := DefaultInstance.DescribeOriginNrtDataSummary(&cdn.DescribeOriginNrtDataSummaryRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func DescribeCdnDataDetail(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnDataDetail(&cdn.DescribeCdnDataDetailRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Domain)
	assert.NotEmpty(t, resp.Result.DataDetails)
}

func DescribeEdgeStatisticalData(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeStatisticalData(&cdn.DescribeEdgeStatisticalDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func DescribeEdgeTopNrtData(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeTopNrtData(&cdn.DescribeEdgeTopNrtDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

func DescribeEdgeTopStatisticalData(t *testing.T) {
	metric := "flux"
	resp, err := DefaultInstance.DescribeEdgeTopStatisticalData(&cdn.DescribeEdgeTopStatisticalDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    &metric,
		Domain:    "yourexample.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

func DescribeCdnRegionAndIsp(t *testing.T) {
	area := "China"
	resp, err := DefaultInstance.DescribeCdnRegionAndIsp(&cdn.DescribeCdnRegionAndIspRequest{Area: &area})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Isps)
	assert.NotEmpty(t, resp.Result.Regions)
}

// 计费查询

func DescribeCdnService(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnService()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.ServiceInfos)
}

func DescribeCdnAccountingData(t *testing.T) {
	domain := "example.com"
	resp, err := DefaultInstance.DescribeAccountingData(&cdn.DescribeAccountingDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    &domain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
}

// 内容管理

func SubmitRefreshTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&cdn.SubmitRefreshTaskRequest{
		Type: &typeFile,
		Urls: "http://example.com/1.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Result.TaskID)
	assert.NotEmpty(t, resp.Result.TaskID)

	// should cause an error when domain is not belong to this account
	_, err = DefaultInstance.SubmitRefreshTask(&cdn.SubmitRefreshTaskRequest{
		Type: &typeFile,
		Urls: "http://foo.com/1.txt",
	})
	assert.Error(t, err)
}

func SubmitRefreshTaskWithCustomExpiresTime(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&cdn.SubmitRefreshTaskRequest{
		Type: &typeFile,
		Urls: "http://example.com/1.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func SubmitPreloadTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitPreloadTask(&cdn.SubmitPreloadTaskRequest{
		Urls: "http://example.com/1.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Result.TaskID)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func DescribeContentTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentTasks(&cdn.DescribeContentTasksRequest{
		TaskType:  typeFile,
		StartTime: testStartTime,
		EndTime:   testEndTime,
	})
	assert.NoError(t, err)
	assert.Greater(t, int(resp.Result.Total), 0)
}

func DescribeContentQuota(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentQuota()
	assert.NoError(t, err)
	assert.Greater(t, int(resp.Result.RefreshQuota), 10)
}

func SubmitBlockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitBlockTask(&cdn.SubmitBlockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func SubmitUnblockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitUnblockTask(&cdn.SubmitUnblockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func DescribeContentBlockTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentBlockTasks(&cdn.DescribeContentBlockTasksRequest{
		TaskType:  typeFile,
		StartTime: testStartTime,
		EndTime:   testEndTime,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Data)
}

// 日志查询

func DescribeCdnAccessLog(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnAccessLog(&cdn.DescribeCdnAccessLogRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if resp.Result.TotalCount > 0 {
		assert.NotEmpty(t, resp.Result.DomainLogDetails)
	}
	assert.Greater(t, int(resp.Result.PageNum), 0)
}

// 服务查询

func DescribeIPInfo(t *testing.T) {
	resp, err := DefaultInstance.DescribeIPInfo(&cdn.DescribeIPInfoRequest{
		IP: testDomain3,
	})
	assert.NoError(t, err)
	assert.Equal(t, testDomain3, resp.Result.IP)
}

func DescribeCdnUpperIp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnUpperIp(&cdn.DescribeCdnUpperIpRequest{Domain: testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

// 标签操作
func AddResourceTags(t *testing.T) {
	resp, err := DefaultInstance.AddResourceTags(&cdn.AddResourceTagsRequest{
		Resources: []string{"www.example1.com", "www.example2.com"},
		ResourceTags: []cdn.ResourceTagEntry{
			{Key: "userKey", Value: "userValue"},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}

func UpdateResourceTags(t *testing.T) {
	resp, err := DefaultInstance.UpdateResourceTags(&cdn.UpdateResourceTagsRequest{
		Resources: []string{"www.example1.com", "www.example2.com"},
		ResourceTags: []cdn.ResourceTagEntry{
			{Key: "userKey", Value: "userValue"},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}

func ListResourceTags(t *testing.T) {
	resp, err := DefaultInstance.ListResourceTags()
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
	assert.NotEmpty(t, resp.Result.Resources)
}

func DeleteResourceTags(t *testing.T) {
	resp, err := DefaultInstance.DeleteResourceTags(&cdn.DeleteResourceTagsRequest{
		Resources: []string{"www.example1.com", "www.example2.com"},
		ResourceTags: []cdn.ResourceTagEntry{
			{Key: "userKey", Value: "userValue"},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}

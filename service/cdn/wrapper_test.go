package cdn

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Warning: these tests may fail when the test interval is less than 2 minutes due to the configuring.

var (
	testAk        = "ak"
	testSk        = "sk"
	typeFile      = "file"
	testDomain1   = "test1.com"
	testDomain2   = "test2.com"
	testDomain3   = "test3.com"
	now           = time.Now()
	testStartTime = now.Unix()
	testEndTime   = now.Add(time.Minute * 10).Unix()
	exampleDomain = "example.com"
)

func TestMain(m *testing.M) {
	ak := os.Getenv("testAk")
	sk := os.Getenv("testSk")
	host := os.Getenv("host")
	if ak != "" && sk != "" {
		testAk = ak
		testSk = sk
	}
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	if host != "" {
		DefaultInstance.Client.SetHost(host)
	}
	os.Exit(m.Run())
}

// 域名操作

func TestCDN_AddCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.AddCdnDomain(&AddCdnDomainRequest{
		Domain:      testDomain2,
		ServiceType: GetStrPtr("web"),
		Origin: []OriginRule{
			{OriginAction: &OriginAction{
				OriginLines: []OriginLine{
					{
						OriginType:   GetStrPtr("primary"),
						InstanceType: GetStrPtr("ip"),
						Address:      GetStrPtr("1.1.1.1"),
						HttpPort:     GetStrPtr("80"),
						HttpsPort:    GetStrPtr("80"),
						Weight:       GetStrPtr("100"),
					},
				},
			}},
		},
		OriginProtocol: GetStrPtr("http"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&StartCdnDomainRequest{Domain: testDomain1})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&StopCdnDomainRequest{Domain: testDomain2})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_DeleteDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&DeleteCdnDomainRequest{Domain: testDomain2})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&ListCdnDomainsRequest{Domain: &testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}

// 域名配置

func TestCDN_DescribeCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnConfig(&DescribeCdnConfigRequest{
		Domain: testDomain2,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.Result.DomainConfig)
	domain := resp.Result.DomainConfig
	fmt.Printf("%+v\n", domain)
}

func TestCDN_UpdateCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.UpdateCdnConfig(&UpdateCdnConfigRequest{
		Domain: &testDomain2,
		Origin: []OriginRule{
			{OriginAction: &OriginAction{
				OriginLines: []OriginLine{
					{
						OriginType:   GetStrPtr("primary"),
						InstanceType: GetStrPtr("ip"),
						Address:      GetStrPtr("1.1.1.1"),
						HttpPort:     GetStrPtr("80"),
						HttpsPort:    GetStrPtr("80"),
						Weight:       GetStrPtr("100"),
					},
				},
			}},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

// 数据统计

func TestCDN_DescribeCdnData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnData(&DescribeCdnDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeEdgeNrtDataSummary(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeNrtDataSummary(&DescribeEdgeNrtDataSummaryRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeCdnOriginData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnOriginData(&DescribeCdnOriginDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeOriginNrtDataSummary(t *testing.T) {
	resp, err := DefaultInstance.DescribeOriginNrtDataSummary(&DescribeOriginNrtDataSummaryRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeCdnDataDetail(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnDataDetail(&DescribeCdnDataDetailRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Name)
	assert.NotEmpty(t, resp.Result.DataDetails)
}

func TestCDN_DescribeEdgeStatisticalData(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeStatisticalData(&DescribeEdgeStatisticalDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "clientIp",
		Domain:    &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeEdgeTopNrtData(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeTopNrtData(&DescribeEdgeTopNrtDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Item:      "isp",
		Domain:    &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

func TestCDN_DescribeOriginTopNrtData(t *testing.T) {
	resp, err := DefaultInstance.DescribeOriginTopNrtData(&DescribeOriginTopNrtDataRequest{
		Metric: "flux",
		Item:   "domain",
		Domain: &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

func TestCDN_DescribeEdgeTopStatusCode(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeTopStatusCode(&DescribeEdgeTopStatusCodeRequest{
		Metric: "status_5xx",
		Item:   "domain",
		Domain: &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

func TestCDN_DescribeOriginTopStatusCode(t *testing.T) {
	resp, err := DefaultInstance.DescribeOriginTopStatusCode(&DescribeOriginTopStatusCodeRequest{
		Metric: "status_5xx",
		Item:   "domain",
		Domain: &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

func TestCDN_DescribeEdgeTopStatisticalData(t *testing.T) {
	metric := "flux"
	resp, err := DefaultInstance.DescribeEdgeTopStatisticalData(&DescribeEdgeTopStatisticalDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    metric,
		Item:      "url",
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

func TestCDN_DescribeCdnRegionAndIsp(t *testing.T) {
	area := "China"
	resp, err := DefaultInstance.DescribeCdnRegionAndIsp(&DescribeCdnRegionAndIspRequest{Area: &area})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Isps)
	assert.NotEmpty(t, resp.Result.Regions)
}

// 计费查询

func TestCDN_DescribeCdnService(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnService()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.ServiceInfos)
}

func TestCDN_DescribeCdnAccountingData(t *testing.T) {
	domain := "example.com"
	resp, err := DefaultInstance.DescribeAccountingData(&DescribeAccountingDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    &domain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
}

// 内容管理

func TestCDN_SubmitRefreshTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&SubmitRefreshTaskRequest{
		Type: &typeFile,
		Urls: "http://example.com/1.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Result.TaskID)
	assert.NotEmpty(t, resp.Result.TaskID)

	// should cause an error when domain is not belong to this account
	_, err = DefaultInstance.SubmitRefreshTask(&SubmitRefreshTaskRequest{
		Type: &typeFile,
		Urls: "http://foo.com/1.txt",
	})
	assert.Error(t, err)
}

func TestCDN_SubmitRefreshTaskWithCustomExpiresTime(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&SubmitRefreshTaskRequest{
		Type: &typeFile,
		Urls: "http://example.com/1.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func TestCDN_SubmitPreloadTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitPreloadTask(&SubmitPreloadTaskRequest{
		Urls: "http://example.com/1.txt",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Result.TaskID)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func TestCDN_DescribeContentTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentTasks(&DescribeContentTasksRequest{
		TaskType:  typeFile,
		StartTime: &testStartTime,
		EndTime:   &testEndTime,
	})
	assert.NoError(t, err)
	assert.Greater(t, int(resp.Result.Total), 0)
}

func TestCDN_DescribeContentQuota(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentQuota()
	assert.NoError(t, err)
	assert.Greater(t, int(resp.Result.RefreshQuota), 10)
}

func TestCDN_SubmitBlockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitBlockTask(&SubmitBlockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func TestCDN_SubmitUnblockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitUnblockTask(&SubmitUnblockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func TestCDN_DescribeContentBlockTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentBlockTasks(&DescribeContentBlockTasksRequest{
		TaskType:  typeFile,
		StartTime: &testStartTime,
		EndTime:   &testEndTime,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Data)
}

// 日志查询

func TestCDN_DescribeCdnAccessLog(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnAccessLog(&DescribeCdnAccessLogRequest{
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

func TestCDN_DescribeIPInfo(t *testing.T) {
	resp, err := DefaultInstance.DescribeIPInfo(&DescribeIPInfoRequest{
		IP: "1.1.1.1",
	})
	assert.NoError(t, err)
	assert.Equal(t, "1.1.1.1", resp.Result.IP)
}

func TestCDN_DescribeCdnUpperIp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnUpperIp(&DescribeCdnUpperIpRequest{Domain: testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

// 标签操作
func TestCDN_AddResourceTags(t *testing.T) {
	resp, err := DefaultInstance.AddResourceTags(&AddResourceTagsRequest{
		Resources: []string{"www.example1.com", "www.example2.com"},
		ResourceTags: []ResourceTag{
			{Key: GetStrPtr("userKey"), Value: GetStrPtr("userValue")},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}

func TestCDN_UpdateResourceTags(t *testing.T) {
	resp, err := DefaultInstance.UpdateResourceTags(&UpdateResourceTagsRequest{
		Resources: []string{"www.example1.com", "www.example2.com"},
		ResourceTags: []ResourceTag{
			{Key: GetStrPtr("userKey"), Value: GetStrPtr("userValue")},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}

func TestCDN_ListResourceTags(t *testing.T) {
	resp, err := DefaultInstance.ListResourceTags()
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
	assert.NotEmpty(t, resp.Result.ResourceTags)
}

func TestCDN_DeleteResourceTags(t *testing.T) {
	resp, err := DefaultInstance.DeleteResourceTags(&DeleteResourceTagsRequest{
		Resources: []string{"www.example1.com", "www.example2.com"},
		ResourceTags: []ResourceTag{
			{Key: GetStrPtr("userKey"), Value: GetStrPtr("userValue")},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
}

package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
	"time"
)

var now = time.Now()
var testStartTime = now.Unix()
var testEndTime = now.Add(time.Minute * 10).Unix()

func init() {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestCDN_DescribeCdnAccountingData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnAccountingData(&cdn.DescribeCdnAccountingDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    "yourdomain.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
}

func TestCDN_DescribeCdnRegionAndIsp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnRegionAndIsp()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Isps)
	assert.NotEmpty(t, resp.Result.Regions)
}

func TestCDN_DescribeCdnData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnData(&cdn.DescribeCdnDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeCdnOriginData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnOriginData(&cdn.DescribeCdnOriginDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeCdnDomainTopData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnDomainTopData(&cdn.DescribeCdnDomainTopDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    "yourdomain.com",
		Item:      "isp",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Domain)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

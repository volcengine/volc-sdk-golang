package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var now = time.Now()
var testStartTime = now.Unix()
var testEndTime = now.Add(time.Minute * 10).Unix()

func TestCDN_DescribeCdnAccountingData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnAccountingData(&DescribeCdnAccountingDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    "qs0902001-auto-test.byteimg.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
}

func TestCDN_DescribeCdnRegionAndIsp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnRegionAndIsp(&DescribeCdnRegionAndIspParam{Area: "China"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Isps)
	assert.NotEmpty(t, resp.Result.Regions)
}

func TestCDN_DescribeCdnData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnData(&DescribeCdnDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeCdnOriginData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnOriginData(&DescribeCdnOriginDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

func TestCDN_DescribeCdnDomainTopData(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnDomainTopData(&DescribeCdnDomainTopDataParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    "qs0902001-auto-test.byteimg.com",
		Item:      "isp",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Domain)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

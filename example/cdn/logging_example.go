package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func init() {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestCDN_DescribeCdnAccessLog(t *testing.T) {
	resp, err := cdn.DefaultInstance.DescribeCdnAccessLog(&cdn.DescribeCdnAccessLogParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    "yourdomain.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if resp.Result.TotalCount > 0 {
		assert.NotEmpty(t, resp.Result.DomainLogDetails)
	}
	assert.Greater(t, resp.Result.PageNum, 0)
}

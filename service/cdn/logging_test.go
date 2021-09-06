package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCDN_DescribeCdnAccessLog(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnAccessLog(&DescribeCdnAccessLogParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    "qs0902001-auto-test.byteimg.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if resp.Result.TotalCount > 0 {
		assert.NotEmpty(t, resp.Result.DomainLogDetails)
	}
	assert.Greater(t, int(resp.Result.PageNum), 0)
}

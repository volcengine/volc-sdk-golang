package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

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

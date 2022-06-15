package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

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

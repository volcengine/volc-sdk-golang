package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeEdgeStatisticalData(t *testing.T) {
	resp, err := DefaultInstance.DescribeEdgeStatisticalData(&cdn.DescribeEdgeStatisticalDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "clientIp",
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Resources)
}

package cdn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DescribeEdgeTopStatisticalData(t *testing.T) {
	metric := "flux"
	resp, err := DefaultInstance.DescribeEdgeTopStatisticalData(&cdn.DescribeEdgeTopStatisticalDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    metric,
		Domain:    exampleDomain,
		Item:      "url",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

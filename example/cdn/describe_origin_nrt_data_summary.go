package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

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

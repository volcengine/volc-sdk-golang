package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeContentTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentTasks(&cdn.DescribeContentTasksRequest{
		TaskType:  "refresh_file",
		StartTime: testStartTime,
		EndTime:   testEndTime,
	})
	assert.NoError(t, err)
	assert.Greater(t, int(resp.Result.Total), 0)
}

package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeAccountingData(t *testing.T) {
	domain := exampleDomain
	resp, err := DefaultInstance.DescribeAccountingData(&cdn.DescribeAccountingDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    &domain,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
}

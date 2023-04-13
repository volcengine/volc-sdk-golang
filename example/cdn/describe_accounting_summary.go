package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
)

func DescribeAccountingSummary(t *testing.T) {
	resp, err := DefaultInstance.DescribeAccountingSummary(&cdn.DescribeAccountingSummaryRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}

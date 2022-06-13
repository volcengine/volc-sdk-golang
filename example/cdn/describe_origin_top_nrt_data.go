package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeOriginTopNrtData(t *testing.T) {
	resp, err := DefaultInstance.DescribeOriginTopNrtData(&cdn.DescribeOriginTopNrtDataRequest{
		Metric: "flux",
		Item:   "domain",
		Domain: &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

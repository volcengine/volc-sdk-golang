package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeOriginTopStatusCode(t *testing.T) {
	resp, err := DefaultInstance.DescribeOriginTopStatusCode(&cdn.DescribeOriginTopStatusCodeRequest{
		Metric: "status_5xx",
		Item:   "domain",
		Domain: &exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}

package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeIPInfo(t *testing.T) {
	resp, err := DefaultInstance.DescribeIPInfo(&cdn.DescribeIPInfoRequest{
		IP: exampleDomain,
	})
	assert.NoError(t, err)
	assert.Equal(t, exampleDomain, resp.Result.IP)
}

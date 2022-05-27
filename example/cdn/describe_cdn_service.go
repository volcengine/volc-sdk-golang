package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func DescribeCdnService(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnService()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.ServiceInfos)
}

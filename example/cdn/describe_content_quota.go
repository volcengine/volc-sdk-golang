package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func DescribeContentQuota(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentQuota()
	assert.NoError(t, err)
	assert.Greater(t, int(resp.Result.RefreshQuota), 10)
}

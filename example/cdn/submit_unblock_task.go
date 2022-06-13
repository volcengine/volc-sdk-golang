package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func SubmitUnblockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitUnblockTask(&cdn.SubmitUnblockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}

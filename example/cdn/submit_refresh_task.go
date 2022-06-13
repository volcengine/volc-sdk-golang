package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func SubmitRefreshTask(t *testing.T) {
	fileType := "file"
	resp, err := DefaultInstance.SubmitRefreshTask(&cdn.SubmitRefreshTaskRequest{
		Type: &fileType,
		Urls: fmt.Sprintf("http://%s/1.txt", exampleDomain),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TaskID)

	// should cause an error when domain is not belong to this account
	_, err = DefaultInstance.SubmitRefreshTask(&cdn.SubmitRefreshTaskRequest{
		Type: &fileType,
		Urls: "http://foo.com/1.txt",
	})
	assert.Error(t, err)
}

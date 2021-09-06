package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

var (
	DefaultInstance = cdn.DefaultInstance
	testAk          = "testAK"
	testSk          = "testSk"
)

func init() {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestCDN_SubmitRefreshTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&cdn.SubmitRefreshTaskParam{
		Type: "file",
		Urls: []string{"http://yourdomain.com/1.txt"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TaskID)
}

func TestCDN_DescribeContentQuota(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentQuota()
	assert.NoError(t, err)
	assert.Greater(t, int(resp.Result.RefreshQuota), 10)
}

func TestCDN_SubmitRefreshTaskWithCustomExpiresTime(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&cdn.SubmitRefreshTaskParam{
		Type: "file",
		Urls: []string{"http://yourdomain.com/1.txt"},
	}, cdn.QueryOption{
		Key:   "X-Expires",
		Value: "300",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TaskID)
}

package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	testAk = "testAK"
	testSk = "testSk"
)

func TestMain(m *testing.M) {
	ak := os.Getenv("testAk")
	sk := os.Getenv("testSk")
	host := os.Getenv("host")
	if ak != "" && sk != "" {
		testAk = ak
		testSk = sk
	}
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	if host != "" {
		DefaultInstance.Client.SetHost(host)
	}
	os.Exit(m.Run())
}

func TestCDN_SubmitRefreshTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&SubmitRefreshTaskParam{
		Type: "file",
		Urls: []string{"http://qs0902002-auto-test.byteimg.com/1.txt"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Result.TaskID)
	assert.NotEmpty(t, resp.Result.TaskID)

	// should cause an error when domain is not belong to this account
	_, err = DefaultInstance.SubmitRefreshTask(&SubmitRefreshTaskParam{
		Type: "file",
		Urls: []string{"http://foo.bar/1.txt"},
	})
	assert.Error(t, err)
}

func TestCDN_SubmitRefreshTaskWithCustomExpiresTime(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(&SubmitRefreshTaskParam{
		Type: "file",
		Urls: []string{"http://dlctest0820003.byteimg.com/1.txt"},
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

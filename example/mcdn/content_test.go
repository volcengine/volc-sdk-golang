package mcdn

import (
	"github.com/stretchr/testify/require"
	"github.com/volcengine/volc-sdk-golang/service/mcdn"
	"testing"
)

func TestSubmitRefreshTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitRefreshTask(mcdn.SubmitRefreshTaskRequest{
		Urls:   "http://example1.com/1.txt\nhttp://example2.com/1.txt",
		Type:   "file",
		Vendor: "ksyun,wangsu",
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.TaskId)
}

func TestSubmitPreloadTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitPreloadTask(mcdn.SubmitPreloadTaskRequest{
		Urls:   "http://example1.com/1.txt\nhttp://example2.com/1.txt",
		Vendor: "ksyun,wangsu",
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.TaskId)
}

func TestListContentTask(t *testing.T) {
	resp, err := DefaultInstance.ListContentTask(mcdn.ListContentTaskRequest{
		TaskId:    "62dc99fd56139eaa6c426cbd",
		TaskType:  "preload",
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    "www.example.com",
		Url:       "https://www.example.com/img",
		Pagination: &mcdn.PagingOption{
			PageSize: 10,
			PageNum:  1,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Tasks)
}

func TestDescribeContentTaskByTaskId(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentTaskByTaskId(mcdn.DescribeContentTaskByTaskIdRequest{
		TaskId: "62dc99fd56139eaa6c426cbd",
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.SubTasks)
}

func TestListVendorContentTask(t *testing.T) {
	resp, err := DefaultInstance.ListVendorContentTask(mcdn.ListVendorContentTaskRequest{
		VendorTaskId:   "task_id_1",
		TaskType:       "refresh_dir",
		CloudAccountId: "cloud_account_1",
		StartTime:      testStartTime,
		EndTime:        testEndTime,
		ProductType:    "cdn",
		Pagination: &mcdn.PagingOption{
			PageSize: 10,
			PageNum:  1,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Tasks)
}

func TestDescribeContentQuota(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentQuota(mcdn.DescribeContentQuotaRequest{
		CloudAccountId: "cloud_account_1,cloud_account_2",
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Quotas)
}

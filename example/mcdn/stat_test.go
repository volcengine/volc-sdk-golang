package mcdn

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/volcengine/volc-sdk-golang/service/mcdn"
)

func TestDescribeCdnDataOffline(t *testing.T) {
	cloudAccountId := "cloud_account_1"
	resp, err := DefaultInstance.DescribeCdnDataOffline(mcdn.DescribeCdnDataOfflineRequest{
		StartTime:       testStartTime,
		EndTime:         testEndTime,
		Metric:          "flux",
		Interval:        "1min",
		CloudAccountIds: []string{cloudAccountId},
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Resources)
}

func TestDescribeCdnOriginDataOffline(t *testing.T) {
	cloudAccountId := "cloud_account_1"
	resp, err := DefaultInstance.DescribeCdnOriginDataOffline(mcdn.DescribeCdnDataOfflineRequest{
		StartTime:       testStartTime,
		EndTime:         testEndTime,
		Metric:          "flux",
		Interval:        "1min",
		CloudAccountIds: []string{cloudAccountId},
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Resources)
}

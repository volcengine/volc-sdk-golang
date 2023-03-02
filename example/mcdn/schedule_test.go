package mcdn

import (
	"github.com/stretchr/testify/require"
	"github.com/volcengine/volc-sdk-golang/service/mcdn"
	"testing"
)

func TestListDnsSchedules(t *testing.T) {
	resp, err := DefaultInstance.ListDnsSchedules(mcdn.ListDnsSchedulesRequest{
		Pagination: &mcdn.PagingOption{
			PageSize: 10,
			PageNum:  1,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.DnsSchedules)
}

func TestDescribeDnsSchedule(t *testing.T) {
	resp, err := DefaultInstance.DescribeDnsSchedule(mcdn.DescribeDnsScheduleRequest{
		Id: "121194655732985856",
	})
	require.NoError(t, err)
	require.NotNil(t, resp.Result.DnsScheduleInfo)
	require.NotEmpty(t, resp.Result.DnsScheduleInfo.Id)
}

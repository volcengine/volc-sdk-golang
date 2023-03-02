package mcdn

import (
	"github.com/stretchr/testify/require"
	"github.com/volcengine/volc-sdk-golang/service/mcdn"
	"testing"
)

func TestListCloudAccounts(t *testing.T) {
	resp, err := DefaultInstance.ListCloudAccounts(mcdn.ListCloudAccountsRequest{
		Pagination: &mcdn.PagingOption{
			PageSize: 10,
			PageNum:  1,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.CloudAccounts)
}

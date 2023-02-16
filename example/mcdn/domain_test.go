package mcdn

import (
	"github.com/stretchr/testify/require"
	"github.com/volcengine/volc-sdk-golang/service/mcdn"
	"testing"
)

func TestListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(mcdn.ListDomainsRequest{
		Vendor: []string{"aliyun"},
		Pagination: &mcdn.PagingOption{
			PageSize: 10,
			PageNum:  1,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Result.Domains)
}

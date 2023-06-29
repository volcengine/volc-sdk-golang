package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func ListCdnDomains(t *testing.T) {
	var a, b = int64(10), int64(1)
	resp, err := DefaultInstance.ListCdnDomains(&cdn.ListCdnDomainsRequest{PageSize: &a, PageNum: &b})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)

	// use get method
	resp, err = DefaultInstance.ListCdnDomains(&cdn.ListCdnDomainsRequest{PageSize: &a, PageNum: &b}, cdn.UseGet())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}

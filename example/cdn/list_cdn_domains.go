package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&cdn.ListCdnDomainsRequest{Domain: &exampleDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}

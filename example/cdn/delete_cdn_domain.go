package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DeleteCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&cdn.DeleteCdnDomainRequest{Domain: operateDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

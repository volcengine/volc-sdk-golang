package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeCdnUpperIp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnUpperIp(&cdn.DescribeCdnUpperIpRequest{Domain: exampleDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

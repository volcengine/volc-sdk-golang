package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCDN_DescribeCdnUpperIp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnUpperIp(&DescribeCdnUpperIpParam{Domain: testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

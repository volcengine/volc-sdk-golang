package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func TestCDN_DescribeCdnUpperIp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnUpperIp(&cdn.DescribeCdnUpperIpParam{Domain: testDomain, IpVersion: "1.1.1.1,114.114.114.114"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

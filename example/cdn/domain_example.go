package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

var testDomain = "www.yourdomain.com"

func TestCDN_StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainParam{DomainName: testDomain})
	if err != nil {
		resp, err = DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainParam{DomainName: testDomain})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainParam{DomainName: testDomain})
	if err != nil {
		resp, err = DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainParam{DomainName: testDomain})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_DeleteDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&cdn.DeleteCdnDomainParam{DomainName: testDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

var testDomain = "www.yourdomain.com"

func TestCDN_StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&cdn.StartDomainParam{DomainName: testDomain})
	if err != nil {
		resp, err = DefaultInstance.StopCdnDomain(&cdn.StopDomainParam{DomainName: testDomain})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&cdn.StopDomainParam{DomainName: testDomain})
	if err != nil {
		resp, err = DefaultInstance.StartCdnDomain(&cdn.StartDomainParam{DomainName: testDomain})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_DeleteDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&cdn.DeleteDomainParam{DomainName: testDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

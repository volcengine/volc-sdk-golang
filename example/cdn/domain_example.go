package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

var testDomain = "www.yourdomain.com"

func TestCDN_StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainParam{Domain: testDomain})
	if err != nil {
		resp, err = DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainParam{Domain: testDomain})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainParam{Domain: testDomain})
	if err != nil {
		resp, err = DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainParam{Domain: testDomain})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_DeleteDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&cdn.DeleteCdnDomainParam{Domain: testDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&cdn.ListCdnDomainsParam{Domain: testDomain, ServiceType: "web", ResourceTag: "tagk:tagv", PageSize: 3, PageNum: 50})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

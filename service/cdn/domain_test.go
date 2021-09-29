package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Warning: these tests may fail when the test interval is less than 2 minutes due to the configuring.

var testDomain1 = "test1.com"
var testDomain2 = "test2.com"
var testDomain3 = "test3.com"

func TestCDN_StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&StartCdnDomainParam{Domain: testDomain1})
	if err != nil {
		resp, err = DefaultInstance.StopCdnDomain(&StopCdnDomainParam{Domain: testDomain1})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&StopCdnDomainParam{Domain: testDomain2})
	if err != nil {
		resp, err = DefaultInstance.StartCdnDomain(&StartCdnDomainParam{Domain: testDomain2})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

//func TestCDN_DeleteDomain(t *testing.T) {
//	resp, err := DefaultInstance.DeleteCdnDomain(&DeleteCdnDomainParam{Domain: "test3.com"})
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//}

func TestCDN_ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&ListCdnDomainsParam{Domain: testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}

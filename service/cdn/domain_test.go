package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Warning: these tests may fail when the test interval is less than 2 minutes due to the configuring.

var testDomain1 = "sdk.cdn-test.bytedance.com"
var testDomain2 = "sdk-online.cdn-test.bytedance.com"
var testDomain3 = "qwer001.byteimg.com"

func TestCDN_StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&StartCdnDomainParam{DomainName: testDomain1})
	if err != nil {
		resp, err = DefaultInstance.StopCdnDomain(&StopCdnDomainParam{DomainName: testDomain1})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&StopCdnDomainParam{DomainName: testDomain2})
	if err != nil {
		resp, err = DefaultInstance.StartCdnDomain(&StartCdnDomainParam{DomainName: testDomain2})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

//func TestCDN_DeleteDomain(t *testing.T) {
//	resp, err := DefaultInstance.DeleteCdnDomain(&DeleteCdnDomainParam{DomainName: "sdk-deleted.cdn-test.bytedance.com"})
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//}

func TestCDN_ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&ListCdnDomainsParam{Domain: testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}

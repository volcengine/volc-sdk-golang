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
//	resp, err := DefaultInstance.DeleteCdnDomain(&DeleteCdnDomainParam{Domain: "sdk-deleted.cdn-test.bytedance.com"})
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//}

func TestCDN_ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&ListCdnDomainsParam{Domain: testDomain3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}

package cdn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Warning: these tests may fail when the test interval is less than 2 minutes due to the configuring.

var testDomain1 = "sdk.cdn-test.bytedance.com"
var testDomain2 = "sdk-online.cdn-test.bytedance.com"

func TestCDN_StartDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&StartDomainParam{DomainName: testDomain1})
	if err != nil {
		resp, err = DefaultInstance.StopCdnDomain(&StopDomainParam{DomainName: testDomain1})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCDN_StopDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&StopDomainParam{DomainName: testDomain2})
	if err != nil {
		resp, err = DefaultInstance.StartCdnDomain(&StartDomainParam{DomainName: testDomain2})
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

//func TestCDN_DeleteDomain(t *testing.T) {
//	resp, err := DefaultInstance.DeleteCdnDomain(&DeleteDomainParam{DomainName: "sdk-deleted.cdn-test.bytedance.com"})
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//}

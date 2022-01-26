package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"
)

type commonResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           interface{} `json:",omitempty"`
}

func TestCommonRequest(t *testing.T) {
	resp := commonResponse{}
	DefaultInstance.Client.SetAccessKey("ak")
	DefaultInstance.Client.SetSecretKey("sk")
	err := DefaultInstance.SendCommonRequest("DescribeCdnAccessLog", &DescribeCdnAccessLogParam{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    "qs0902001-auto-test.byteimg.com",
	}, &resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
	assert.Nil(t, resp.ResponseMetadata.Error)
	fmt.Println(resp.Result)
}

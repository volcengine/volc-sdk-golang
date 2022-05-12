package cdn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type commonResponse struct {
	ResponseMetadata *ResponseMetadata
	Result           interface{} `json:",omitempty"`
}

func TestCommonRequest(t *testing.T) {
	resp := commonResponse{}
	DefaultInstance.Client.SetAccessKey("ak")
	DefaultInstance.Client.SetSecretKey("sk")
	err := DefaultInstance.SendCommonRequest("DescribeCdnAccessLog", &DescribeCdnAccessLogRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    "qs0902001-auto-test.byteimg.com",
	}, &resp)
	assert.NoError(t, err)
	err = DefaultInstance.ValidateResponse(resp.ResponseMetadata)
	assert.NoError(t, err)
	fmt.Println(resp.Result)
}

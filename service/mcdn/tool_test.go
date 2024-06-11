package mcdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/volcengine/volc-sdk-golang/base"
)

type commonResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           interface{} `json:",omitempty"`
}

func TestCommonRequest(t *testing.T) {
	resp := commonResponse{}
	DefaultInstance.Client.SetAccessKey("ak")
	DefaultInstance.Client.SetSecretKey("sk")
	err := DefaultInstance.SendCommonRequest("ListCloudAccounts", &ListCloudAccountsRequest{}, &resp)
	require.NoError(t, err)
	require.NotNil(t, resp.ResponseMetadata)
	err = DefaultInstance.ValidateResponse(*resp.ResponseMetadata)
	require.NoError(t, err)
	fmt.Println(resp.Result)
}

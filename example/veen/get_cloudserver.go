package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func GetCloudServer(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).GetCloudServer(&veen.GetCloudServerReq{
		CloudServerID: "cloudserver-c4zxqrb7jz2jchd",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

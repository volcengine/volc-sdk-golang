package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func RebootCloudServer(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).RebootCloudServer(&veen.RebootCloudServerReq{
		CloudServerID: "cloudserver-2vgvjzqt9m7567f",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

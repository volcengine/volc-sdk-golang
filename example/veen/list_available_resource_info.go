package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ListAvailableResourceInfo(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).ListAvailableResourceInfo(&veen.ListAvailableResourceInfoReq{
		InstanceType:  "veEN.C1.large",
		CloudDiskType: "CloudSSD",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

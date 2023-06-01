package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func SetInstanceName(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).SetInstanceName(&veen.SetInstanceNameReq{
		InstanceIdentity: "veen20020201213020320353",
		InstanceName:     "golang_sdk测试",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

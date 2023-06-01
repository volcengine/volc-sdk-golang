package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func GetInstance(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).GetInstance(&veen.GetInstanceReq{
		InstanceIdentity: "veen20020201213020320353",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

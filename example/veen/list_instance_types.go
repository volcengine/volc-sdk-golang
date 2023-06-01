package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ListInstanceTypes(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).ListInstanceTypes(&veen.ListInstanceTypesReq{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

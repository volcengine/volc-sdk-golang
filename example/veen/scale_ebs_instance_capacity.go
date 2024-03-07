package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ScaleEbsInstanceCapacity(t *testing.T) {
	withReboot := true
	resp, err := veen.NewInstance(ak, sk).ScaleEbsInstanceCapacity(&veen.ScaleEbsInstanceCapacityReq{
		EbsID:      "disk-t9p44586fn6cbs9",
		Capacity:   "200",
		WithReboot: &withReboot,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

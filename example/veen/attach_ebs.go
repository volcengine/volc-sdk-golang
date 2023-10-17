package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func AttachEbs(t *testing.T) {
	deleteWithRes := true
	resp, err := veen.NewInstance(ak, sk).AttachEbs(&veen.AttachEbsReq{
		EbsID:         "disk-t9p44586fn6cbs9",
		EbsIds:        []string{"disk-t9p44586fn6cbs9", "disk-fz26fhklggjnbhj"},
		ResType:       "veen",
		ResID:         "testing2-veen92023710333272539522",
		DeleteWithRes: &deleteWithRes,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

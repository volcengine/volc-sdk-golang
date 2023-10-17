package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func GetEbsInstance(t *testing.T) {
	withAttachmentInfo := true
	resp, err := veen.NewInstance(ak, sk).GetEbsInstance(&veen.GetEbsInstanceReq{
		EbsID:              "disk-t9p44586fn6cbs9",
		WithAttachmentInfo: &withAttachmentInfo,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

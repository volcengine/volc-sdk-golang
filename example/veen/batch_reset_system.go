package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func BatchResetSystem(t *testing.T) {
	clearDataDisk := false
	imageIdentity := "image7ajpdaodf7"
	resp, err := veen.NewInstance(ak, sk).BatchResetSystem(&veen.BatchResetSystemReq{
		InstanceIdentities: []string{"testing-veen22052022220212262231", "testing-veen23323035425312030023"},
		ImageIdentity:      &imageIdentity,
		ClearDataDisk:      &clearDataDisk,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

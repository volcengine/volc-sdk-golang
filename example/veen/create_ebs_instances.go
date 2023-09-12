package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func CreateEbsInstances(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).CreateEbsInstances(&veen.CreateEbsInstancesReq{
		ClusterName: "bdcdn-chdcu",
		ChargeType:  "HourUsed",
		EbsType:     "data",
		StorageType: "CloudBlockHDD",
		Capacity:    "20",
		Number:      1,
		Name:        "test",
		Desc:        "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

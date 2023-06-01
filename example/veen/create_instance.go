package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func CreateInstance(t *testing.T) {
	clusterName := "bdcdn-ycct"
	resp, err := veen.NewInstance(ak, sk).CreateInstance(&veen.CreateInstanceReq{
		CloudServerIdentity: "cloudserver-2vgvjzqt9m7567f",
		InstanceAreaNums: []*veen.InstanceAreaNum{
			{
				ClusterName: &clusterName,
				Num:         1,
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

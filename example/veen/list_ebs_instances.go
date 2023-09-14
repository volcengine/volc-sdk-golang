package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ListEbsInstances(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).ListEbsInstances(&veen.ListEbsInstancesReq{
		WithAttachmentInfo:  true,
		ResIds:              "testing2-veen92023710333272539522,testing2-veen23303021220530991200",
		EbsIds:              "disk-t9p44586fn6cbs9,disk-fz26fhklggjnbhj",
		EbsNames:            "本地_虚机测试用-3,本地_虚机测试用-2",
		Regions:             "CentralChina,NorthChina",
		ClusterNames:        "bdcdn-chdcu,bdcdn-sjzmp03",
		Status:              "attached,detached",
		EbsType:             "data",
		ChargeType:          "HourUsed",
		FuzzyVeenExternalIP: "172.19.23.230",
		PageOption: veen.EbsPagination{
			PageNo:   1,
			PageSize: 10,
		},
		OrderOption: veen.EbsOrderOption{
			OrderBy: "ebs_id",
			Asc:     false,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

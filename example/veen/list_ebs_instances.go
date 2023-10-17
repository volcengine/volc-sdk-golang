package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ListEbsInstances(t *testing.T) {
	withAttachmentInfo := true
	fuzzyIp := "172.19.23.230"
	resp, err := veen.NewInstance(ak, sk).ListEbsInstances(&veen.ListEbsInstancesReq{
		WithAttachmentInfo:  &withAttachmentInfo,
		ResIds:              []string{"testing2-veen92023710333272539522", "testing2-veen23303021220530991200"},
		EbsIds:              []string{"disk-t9p44586fn6cbs9", "disk-fz26fhklggjnbhj"},
		EbsNames:            []string{"本地_虚机测试用-3", "本地_虚机测试用-2"},
		Regions:             []string{"CentralChina", "NorthChina"},
		ClusterNames:        []string{"bdcdn-chdcu", "bdcdn-sjzmp03"},
		Status:              []string{"attached", "detached"},
		EbsType:             []string{"data"},
		ChargeType:          []string{"HourUsed"},
		FuzzyVeenExternalIP: &fuzzyIp,
		PageOption: &veen.EbsPagination{
			PageNo:   1,
			PageSize: 10,
		},
		OrderOption: &veen.EbsOrderOption{
			OrderBy: "id",
			Asc:     false,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

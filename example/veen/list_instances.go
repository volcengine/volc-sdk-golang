package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ListInstances(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).ListInstances(&veen.ListInstancesReq{
		Regions: "SouthChina,Northwest",
		Cities:  "440300,640100",
		Isps:    "CTCC,CMCC_CTCC_CUCC",
		Status:  "opening,running",
		Pagination: veen.Pagination{
			Page:    2,
			Limit:   2,
			OrderBy: 2,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

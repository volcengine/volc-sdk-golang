package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func ScaleInstanceCloudDiskCapacity(t *testing.T) {
	withReboot := true
	resp, err := veen.NewInstance(ak, sk).ScaleInstanceCloudDiskCapacity(&veen.ScaleInstanceCloudDiskCapacityReq{
		InstanceIdentity: "veen20020201213020320353",
		ScaleSystemCloudDiskInfo: &veen.ScaleCloudDiskInfo{
			DeviceName: "vda",
			Capacity:   "1000",
		},
		ScaleDataCloudDiskInfoList: []*veen.ScaleCloudDiskInfo{
			{
				DeviceName: "vdb",
				Capacity:   "1000",
			},
		},
		WithReboot: &withReboot,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}

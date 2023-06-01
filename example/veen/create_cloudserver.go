package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func CreateCloudServer(t *testing.T) {
	bandwidthPeak := "200"
	enableIpv6 := false
	secretData := "Abcd1234*"
	clusterName := "bdcdn-ycct"
	resp, err := veen.NewInstance(ak, sk).CreateCloudServer(&veen.CreateCloudServerReq{
		CloudServerName: "test-golang-sdk",
		ImageID:         "imagekhvhgzhna6",
		SpecName:        "veEN.C1.large",
		StorageConfig: &veen.StorageConfig{
			SystemDisk: &veen.DiskSpec{
				StorageType: "CloudBlockSSD",
				Capacity:    "40",
			},
			DataDiskList: []*veen.DiskSpec{
				{
					StorageType: "CloudBlockSSD",
					Capacity:    "20",
				},
				{
					StorageType: "CloudBlockSSD",
					Capacity:    "40",
				},
			},
		},
		NetworkConfig: &veen.CloudServerNetworkConfig{
			BandwidthPeak: &bandwidthPeak,
			EnableIpv6:    &enableIpv6,
		},
		SecretConfig: &veen.SecretConfig{
			SecretType: veen.SecretType_Custom,
			SecretData: &secretData,
		},
		BillingConfig: &veen.CloudServerBillingConfigs{
			ComputingBillingMethod: "MonthlyPeak",
			BandwidthBillingMethod: "MonthlyP95",
		},
		InstanceAreaNums: []*veen.InstanceAreaNum{
			{
				ClusterName: &clusterName,
				Num:         2,
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.ResponseMetadata)
	fmt.Printf("%+v", veen.ToJson(resp))
}

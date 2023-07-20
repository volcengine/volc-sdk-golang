// Code generated by protoc-gen-volcengine-sdk
// source: VodCdnService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_ListDomain(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListDomainRequest{
		SpaceName:         "your SpaceName",
		DomainType:        "your DomainType",
		SourceStationType: 0,
		Offset:            0,
	}

	resp, status, err := instance.ListDomain(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreateCdnRefreshTask(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCreateCdnRefreshTaskRequest{
		SpaceName: "your SpaceName",
		Urls:      "your Urls",
		Type:      "your Type",
	}

	resp, status, err := instance.CreateCdnRefreshTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreateCdnPreloadTask(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCreateCdnPreloadTaskRequest{
		SpaceName: "your SpaceName",
		Urls:      "your Urls",
	}

	resp, status, err := instance.CreateCdnPreloadTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnTasks(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListCdnTasksRequest{
		SpaceName:      "your SpaceName",
		TaskId:         "your TaskId",
		DomainName:     "your DomainName",
		TaskType:       "your TaskType",
		Status:         "your Status",
		StartTimestamp: 0,
		EndTimestamp:   0,
		PageNum:        0,
		PageSize:       0,
	}

	resp, status, err := instance.ListCdnTasks(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnAccessLog(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListCdnAccessLogRequest{
		Domains:        "your Domains",
		StartTimestamp: 0,
		EndTimestamp:   0,
		SpaceName:      "your SpaceName",
	}

	resp, status, err := instance.ListCdnAccessLog(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnTopAccessUrl(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListCdnTopAccessUrlRequest{
		Domains:        "your Domains",
		StartTimestamp: 0,
		EndTimestamp:   0,
		SortType:       "your SortType",
	}

	resp, status, err := instance.ListCdnTopAccessUrl(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeVodDomainBandwidthData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDescribeVodDomainBandwidthDataRequest{
		DomainList:    "your DomainList",
		StartTime:     "your StartTime",
		EndTime:       "your EndTime",
		Aggregation:   0,
		BandwidthType: "your BandwidthType",
		Area:          "your Area",
	}

	resp, status, err := instance.DescribeVodDomainBandwidthData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnUsageData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListCdnUsageDataRequest{
		Domains:        "your Domains",
		Interval:       "your Interval",
		StartTimestamp: 0,
		EndTimestamp:   0,
		DataType:       "your DataType",
		Metric:         "your Metric",
		NeedDetail:     false,
		Area:           "your Area",
		Region:         "your Region",
		Isp:            "your Isp",
		Protocol:       "your Protocol",
		IpVersion:      "your IpVersion",
		BillingRegion:  "your BillingRegion",
	}

	resp, status, err := instance.ListCdnUsageData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnStatusData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListCdnStatusDataRequest{
		Domains:        "your Domains",
		Interval:       "your Interval",
		StartTimestamp: 0,
		EndTimestamp:   0,
		DataType:       "your DataType",
		Metric:         "your Metric",
		NeedDetail:     false,
	}

	resp, status, err := instance.ListCdnStatusData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnHitrateData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListCdnHitrateDataRequest{
		Domains:        "your domain",
		Interval:       "auto",
		StartTimestamp: 1689414058,
		EndTimestamp:   1689759658,
		Metric:         "hitrate",
		NeedDetail:     false,
	}

	resp, status, err := instance.ListCdnHitrateData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeIpInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDescribeIPInfoRequest{
		Ips: "your Ips",
	}

	resp, status, err := instance.DescribeIpInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeVodDomainTrafficData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDescribeVodDomainTrafficDataRequest{
		DomainList:  "your DomainList",
		StartTime:   "your StartTime",
		EndTime:     "your EndTime",
		Aggregation: 0,
		TrafficType: "your TrafficType",
	}

	resp, status, err := instance.DescribeVodDomainTrafficData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnPvData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListCdnPvDataRequest{
		Domains:        "your Domains",
		Interval:       "your Interval",
		StartTimestamp: 0,
		EndTimestamp:   0,
		DataType:       "your DataType",
		NeedDetail:     false,
	}

	resp, status, err := instance.ListCdnPvData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_SubmitBlockTasks(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodSubmitBlockTasksRequest{
		FileUrls:  "your FileUrls",
		Operation: "your Operation",
	}

	resp, status, err := instance.SubmitBlockTasks(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetContentBlockTasks(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetContentBlockTasksRequest{
		Url:       "your Url",
		Domain:    "your Domain",
		TaskID:    "your TaskID",
		TaskType:  "your TaskType",
		Status:    "your Status",
		StartTime: 0,
		EndTime:   0,
		PageNum:   0,
		PageSize:  0,
	}

	resp, status, err := instance.GetContentBlockTasks(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreateDomain(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCreateDomainV2Request{
		SpaceName:                "your SpaceName",
		DomainType:               "your DomainType",
		Domain:                   "your Domain",
		SourceStationType:        0,
		SourceStationAddressType: 0,
		Origins:                  "your Origins",
		Area:                     "your Area",
		BucketName:               "your BucketName",
	}

	resp, status, err := instance.CreateDomain(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateDomainExpire(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateDomainExpireV2Request{
		SpaceName:  "your SpaceName",
		DomainType: "your DomainType",
		Domain:     "your Domain",
		Expire:     0,
	}

	resp, status, err := instance.UpdateDomainExpire(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateDomainAuthConfig(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateDomainAuthConfigV2Request{
		SpaceName:  "your SpaceName",
		DomainType: "your DomainType",
		Domain:     "your Domain",
		MainKey:    "your MainKey",
		BackupKey:  "your BackupKey",
		Status:     "your Status",
	}

	resp, status, err := instance.UpdateDomainAuthConfig(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

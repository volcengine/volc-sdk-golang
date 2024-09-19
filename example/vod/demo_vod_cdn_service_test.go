// Code generated by protoc-gen-volcengine-sdk
// source: VodCdnService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_AddDomainToScheduler(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodAddDomainToSchedulerRequest{
		SpaceName:         "your SpaceName",
		DomainType:        "your DomainType",
		Domain:            "your Domain",
		SourceStationType: 0,
	}

	resp, status, err := instance.AddDomainToScheduler(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_RemoveDomainFromScheduler(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodRemoveDomainFromSchedulerRequest{
		SpaceName:         "your SpaceName",
		DomainType:        "your DomainType",
		Domain:            "your Domain",
		SourceStationType: 0,
	}

	resp, status, err := instance.RemoveDomainFromScheduler(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateDomainPlayRule(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodUpdateDomainPlayRuleRequest{
		SpaceName:     "your SpaceName",
		DefaultDomain: "your DefaultDomain",
		PlayRule:      0,
	}

	resp, status, err := instance.UpdateDomainPlayRule(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_StartDomain(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodStartDomainRequest{
		SpaceName:         "your SpaceName",
		DomainType:        "your DomainType",
		Domain:            "your Domain",
		SourceStationType: 0,
	}

	resp, status, err := instance.StartDomain(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_StopDomain(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodStopDomainRequest{
		SpaceName:         "your SpaceName",
		DomainType:        "your DomainType",
		Domain:            "your Domain",
		SourceStationType: 0,
	}

	resp, status, err := instance.StopDomain(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteDomain(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodDeleteDomainRequest{
		SpaceName:  "your SpaceName",
		DomainType: "your DomainType",
		Domain:     "your Domain",
	}

	resp, status, err := instance.DeleteDomain(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListDomain(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodListDomainRequest{
		SpaceName:         "your SpaceName",
		DomainType:        "your DomainType",
		SourceStationType: 0,
		Offset:            0,
		Limit:             0,
	}

	resp, status, err := instance.ListDomain(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreateCdnRefreshTask(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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

func Test_ListCdnTopAccess(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodListCdnTopAccessRequest{
		Domains:        "your Domains",
		StartTimestamp: 0,
		EndTimestamp:   0,
		SortType:       "your SortType",
		Item:           "your Item",
	}

	resp, status, err := instance.ListCdnTopAccess(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnUsageData(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodListCdnHitrateDataRequest{
		Domains:        "your Domains",
		Interval:       "your Interval",
		StartTimestamp: 0,
		EndTimestamp:   0,
		Metric:         "your Metric",
		NeedDetail:     false,
	}

	resp, status, err := instance.ListCdnHitrateData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeIpInfo(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodDescribeIPInfoRequest{
		Ips: "your Ips",
	}

	resp, status, err := instance.DescribeIpInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListCdnPvData(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

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

func Test_AddOrUpdateCertificate(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.AddOrUpdateCertificateV2Request{
		SpaceName:     "your SpaceName",
		Domain:        "your Domain",
		DomainType:    "your DomainType",
		CertificateId: "your CertificateId",
		HttpsStatus:   "your HttpsStatus",
	}

	resp, status, err := instance.AddOrUpdateCertificate(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateDomainUrlAuthConfig(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodUpdateDomainUrlAuthConfigV2Request{
		SpaceName:  "your SpaceName",
		DomainType: "your DomainType",
		Domain:     "your Domain",
		MainKey:    "your MainKey",
		BackupKey:  "your BackupKey",
		Status:     "your Status",
	}

	resp, status, err := instance.UpdateDomainUrlAuthConfig(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateDomainConfig(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodUpdateDomainConfigRequest{
		SpaceName:  "your SpaceName",
		DomainType: "your DomainType",
		Domain:     "your Domain",
		Config:     nil,
	}

	resp, status, err := instance.UpdateDomainConfig(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeDomainConfig(t *testing.T) {
	// Create a VOD instance in the specified region.
	// instance := vod.NewInstanceWithRegion("cn-north-1")
	instance := vod.NewInstance()

	// Configure your Access Key ID (AK) and Secret Access Key (SK) in the environment variables or in the local ~/.volc/config file. For detailed instructions, see  https://www.volcengine.com/docs/4/65655.
	// The SDK will automatically fetch the AK and SK from the environment variables or the ~/.volc/config file as needed.
	// During testing, you may use the following code snippet. However, do not store the AK and SK directly in your project code to prevent potential leakage and safeguard the security of all resources associated with your account.
	// instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	query := &request.VodDescribeDomainConfigRequest{
		SpaceName:  "your SpaceName",
		DomainType: "your DomainType",
		Domain:     "your Domain",
	}

	resp, status, err := instance.DescribeDomainConfig(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

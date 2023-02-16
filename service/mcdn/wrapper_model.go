package mcdn

import "github.com/volcengine/volc-sdk-golang/base"

type ListCloudAccountsResponseWrapper struct {
	Result           ListCloudAccountsResponse
	ResponseMetadata base.ResponseMetadata
}

type ListDomainsResponseWrapper struct {
	Result           ListDomainsResponse
	ResponseMetadata base.ResponseMetadata
}

type ListDnsSchedulesResponseWrapper struct {
	Result           ListDnsSchedulesResponse
	ResponseMetadata base.ResponseMetadata
}

type DescribeDnsScheduleResponseWrapper struct {
	Result           DescribeDnsScheduleResponse
	ResponseMetadata base.ResponseMetadata
}

type DescribeCdnDataResponseWrapper struct {
	Result           DescribeCdnDataResponse
	ResponseMetadata base.ResponseMetadata
}

type DescribeCdnDataOfflineResponseWrapper struct {
	Result           DescribeCdnDataOfflineResponse
	ResponseMetadata base.ResponseMetadata
}

type SubmitTaskResponseWrapper struct {
	Result           SubmitTaskResponse
	ResponseMetadata base.ResponseMetadata
}

type ListContentTaskResponseWrapper struct {
	Result           ListContentTaskResponse
	ResponseMetadata base.ResponseMetadata
}

type DescribeContentTaskByTaskIdResponseWrapper struct {
	Result           DescribeContentTaskByTaskIdResponse
	ResponseMetadata base.ResponseMetadata
}

type ListVendorContentTaskResponseWrapper struct {
	Result           ListVendorContentTaskResponse
	ResponseMetadata base.ResponseMetadata
}

type DescribeContentQuotaResponseWrapper struct {
	Result           DescribeContentQuotaResponse
	ResponseMetadata base.ResponseMetadata
}

type ListViewsResponseWrapper struct {
	Result           ListViewsResponse
	ResponseMetadata base.ResponseMetadata
}

type DescribeCdnRegionAndIspResponseWrapper struct {
	Result           DescribeCdnRegionAndIspResponse
	ResponseMetadata base.ResponseMetadata
}

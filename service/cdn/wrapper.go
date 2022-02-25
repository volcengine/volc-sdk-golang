package cdn

func (s *CDN) AddCdnDomain(dto *AddCdnDomainRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("AddCdnDomain", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StartCdnDomain(dto *StartCdnDomainRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("StartCdnDomain", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StopCdnDomain(dto *StopCdnDomainRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("StopCdnDomain", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnDomain(dto *DeleteCdnDomainRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("DeleteCdnDomain", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnDomains(dto *ListCdnDomainsRequest, opts ...RequestOption) (responseBody *ListCdnDomainsResponse, err error) {
	responseBody = new(ListCdnDomainsResponse)
	if err = s.doRequest("ListCdnDomains", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnConfig(dto *DescribeCdnConfigRequest, opts ...RequestOption) (responseBody *DescribeCdnConfigResponse, err error) {
	responseBody = new(DescribeCdnConfigResponse)
	if err = s.doRequest("DescribeCdnConfig", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateCdnConfig(dto *UpdateCdnConfigRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("UpdateCdnConfig", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnData(dto *DescribeCdnDataRequest, opts ...RequestOption) (responseBody *DescribeCdnDataResponse, err error) {
	responseBody = new(DescribeCdnDataResponse)
	if err = s.doRequest("DescribeCdnData", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeNrtDataSummary(dto *DescribeEdgeNrtDataSummaryRequest, opts ...RequestOption) (responseBody *DescribeEdgeNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeEdgeNrtDataSummaryResponse)
	if err = s.doRequest("DescribeEdgeNrtDataSummary", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnOriginData(dto *DescribeCdnOriginDataRequest, opts ...RequestOption) (responseBody *DescribeCdnOriginDataResponse, err error) {
	responseBody = new(DescribeCdnOriginDataResponse)
	if err = s.doRequest("DescribeCdnOriginData", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginNrtDataSummary(dto *DescribeOriginNrtDataSummaryRequest, opts ...RequestOption) (responseBody *DescribeOriginNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeOriginNrtDataSummaryResponse)
	if err = s.doRequest("DescribeOriginNrtDataSummary", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnDataDetail(dto *DescribeCdnDataDetailRequest, opts ...RequestOption) (responseBody *DescribeCdnDataDetailResponse, err error) {
	responseBody = new(DescribeCdnDataDetailResponse)
	if err = s.doRequest("DescribeCdnDataDetail", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeStatisticalData(dto *DescribeEdgeStatisticalDataRequest, opts ...RequestOption) (responseBody *DescribeEdgeStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeStatisticalDataResponse)
	if err = s.doRequest("DescribeEdgeStatisticalData", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopNrtData(dto *DescribeEdgeTopNrtDataRequest, opts ...RequestOption) (responseBody *DescribeEdgeTopNrtDataResponse, err error) {
	responseBody = new(DescribeEdgeTopNrtDataResponse)
	if err = s.doRequest("DescribeEdgeTopNrtData", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopStatisticalData(dto *DescribeEdgeTopStatisticalDataRequest, opts ...RequestOption) (responseBody *DescribeEdgeTopStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeTopStatisticalDataResponse)
	if err = s.doRequest("DescribeEdgeTopStatisticalData", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnRegionAndIsp(dto *DescribeCdnRegionAndIspRequest, opts ...RequestOption) (responseBody *DescribeCdnRegionAndIspResponse, err error) {
	responseBody = new(DescribeCdnRegionAndIspResponse)
	if err = s.doRequest("DescribeCdnRegionAndIsp", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnDomainTopData(dto *DescribeCdnDomainTopDataRequest, opts ...RequestOption) (responseBody *DescribeCdnDomainTopDataResponse, err error) {
	responseBody = new(DescribeCdnDomainTopDataResponse)
	if err = s.doRequest("DescribeCdnDomainTopData", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnService(opts ...RequestOption) (responseBody *DescribeCdnServiceResponse, err error) {
	responseBody = new(DescribeCdnServiceResponse)
	if err = s.doRequest("DescribeCdnService", nil, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeAccountingData(dto *DescribeAccountingDataRequest, opts ...RequestOption) (responseBody *DescribeAccountingDataResponse, err error) {
	responseBody = new(DescribeAccountingDataResponse)
	if err = s.doRequest("DescribeAccountingData", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitRefreshTask(dto *SubmitRefreshTaskRequest, opts ...RequestOption) (responseBody *SubmitRefreshTaskResponse, err error) {
	responseBody = new(SubmitRefreshTaskResponse)
	if err = s.doRequest("SubmitRefreshTask", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitPreloadTask(dto *SubmitPreloadTaskRequest, opts ...RequestOption) (responseBody *SubmitPreloadTaskResponse, err error) {
	responseBody = new(SubmitPreloadTaskResponse)
	if err = s.doRequest("SubmitPreloadTask", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentTasks(dto *DescribeContentTasksRequest, opts ...RequestOption) (responseBody *DescribeContentTasksResponse, err error) {
	responseBody = new(DescribeContentTasksResponse)
	if err = s.doRequest("DescribeContentTasks", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentQuota(opts ...RequestOption) (responseBody *DescribeContentQuotaResponse, err error) {
	responseBody = new(DescribeContentQuotaResponse)
	if err = s.doRequest("DescribeContentQuota", nil, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitBlockTask(dto *SubmitBlockTaskRequest, opts ...RequestOption) (responseBody *SubmitBlockTaskResponse, err error) {
	responseBody = new(SubmitBlockTaskResponse)
	if err = s.doRequest("SubmitBlockTask", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitUnblockTask(dto *SubmitUnblockTaskRequest, opts ...RequestOption) (responseBody *SubmitUnblockTaskResponse, err error) {
	responseBody = new(SubmitUnblockTaskResponse)
	if err = s.doRequest("SubmitUnblockTask", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentBlockTasks(dto *DescribeContentBlockTasksRequest, opts ...RequestOption) (responseBody *DescribeContentBlockTasksResponse, err error) {
	responseBody = new(DescribeContentBlockTasksResponse)
	if err = s.doRequest("DescribeContentBlockTasks", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnAccessLog(dto *DescribeCdnAccessLogRequest, opts ...RequestOption) (responseBody *DescribeCdnAccessLogResponse, err error) {
	responseBody = new(DescribeCdnAccessLogResponse)
	if err = s.doRequest("DescribeCdnAccessLog", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeIPInfo(dto *DescribeIPInfoRequest, opts ...RequestOption) (responseBody *DescribeIPInfoResponse, err error) {
	responseBody = new(DescribeIPInfoResponse)
	if err = s.doRequest("DescribeIPInfo", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnUpperIp(dto *DescribeCdnUpperIpRequest, opts ...RequestOption) (responseBody *DescribeCdnUpperIpResponse, err error) {
	responseBody = new(DescribeCdnUpperIpResponse)
	if err = s.doRequest("DescribeCdnUpperIp", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) AddResourceTags(dto *AddResourceTagsRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("AddResourceTags", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateResourceTags(dto *UpdateResourceTagsRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("UpdateResourceTags", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListResourceTags(opts ...RequestOption) (responseBody *ListResourceTagsResponse, err error) {
	responseBody = new(ListResourceTagsResponse)
	if err = s.doRequest("ListResourceTags", nil, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteResourceTags(dto *DeleteResourceTagsRequest, opts ...RequestOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("DeleteResourceTags", &dto, responseBody, opts...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

package cdn

func (s *CDN) AddCdnDomain(dto *AddCdnDomainRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("AddCdnDomain", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StartCdnDomain(dto *StartCdnDomainRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("StartCdnDomain", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StopCdnDomain(dto *StopCdnDomainRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("StopCdnDomain", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnDomain(dto *DeleteCdnDomainRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("DeleteCdnDomain", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnDomains(dto *ListCdnDomainsRequest) (responseBody *ListCdnDomainsResponse, err error) {
	responseBody = new(ListCdnDomainsResponse)
	if err = s.doRequest("ListCdnDomains", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnConfig(dto *DescribeCdnConfigRequest) (responseBody *DescribeCdnConfigResponse, err error) {
	responseBody = new(DescribeCdnConfigResponse)
	if err = s.doRequest("DescribeCdnConfig", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateCdnConfig(dto *UpdateCdnConfigRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("UpdateCdnConfig", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnData(dto *DescribeCdnDataRequest) (responseBody *DescribeCdnDataResponse, err error) {
	responseBody = new(DescribeCdnDataResponse)
	if err = s.doRequest("DescribeCdnData", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeNrtDataSummary(dto *DescribeEdgeNrtDataSummaryRequest) (responseBody *DescribeEdgeNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeEdgeNrtDataSummaryResponse)
	if err = s.doRequest("DescribeEdgeNrtDataSummary", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnOriginData(dto *DescribeCdnOriginDataRequest) (responseBody *DescribeCdnOriginDataResponse, err error) {
	responseBody = new(DescribeCdnOriginDataResponse)
	if err = s.doRequest("DescribeCdnOriginData", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginNrtDataSummary(dto *DescribeOriginNrtDataSummaryRequest) (responseBody *DescribeOriginNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeOriginNrtDataSummaryResponse)
	if err = s.doRequest("DescribeOriginNrtDataSummary", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnDataDetail(dto *DescribeCdnDataDetailRequest) (responseBody *DescribeCdnDataDetailResponse, err error) {
	responseBody = new(DescribeCdnDataDetailResponse)
	if err = s.doRequest("DescribeCdnDataDetail", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeStatisticalData(dto *DescribeEdgeStatisticalDataRequest) (responseBody *DescribeEdgeStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeStatisticalDataResponse)
	if err = s.doRequest("DescribeEdgeStatisticalData", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopNrtData(dto *DescribeEdgeTopNrtDataRequest) (responseBody *DescribeEdgeTopNrtDataResponse, err error) {
	responseBody = new(DescribeEdgeTopNrtDataResponse)
	if err = s.doRequest("DescribeEdgeTopNrtData", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopStatisticalData(dto *DescribeEdgeTopStatisticalDataRequest) (responseBody *DescribeEdgeTopStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeTopStatisticalDataResponse)
	if err = s.doRequest("DescribeEdgeTopStatisticalData", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnRegionAndIsp(dto *DescribeCdnRegionAndIspRequest) (responseBody *DescribeCdnRegionAndIspResponse, err error) {
	responseBody = new(DescribeCdnRegionAndIspResponse)
	if err = s.doRequest("DescribeCdnRegionAndIsp", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnService() (responseBody *DescribeCdnServiceResponse, err error) {
	responseBody = new(DescribeCdnServiceResponse)
	if err = s.doRequest("DescribeCdnService", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeAccountingData(dto *DescribeAccountingDataRequest) (responseBody *DescribeAccountingDataResponse, err error) {
	responseBody = new(DescribeAccountingDataResponse)
	if err = s.doRequest("DescribeAccountingData", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitRefreshTask(dto *SubmitRefreshTaskRequest) (responseBody *SubmitRefreshTaskResponse, err error) {
	responseBody = new(SubmitRefreshTaskResponse)
	if err = s.doRequest("SubmitRefreshTask", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitPreloadTask(dto *SubmitPreloadTaskRequest) (responseBody *SubmitPreloadTaskResponse, err error) {
	responseBody = new(SubmitPreloadTaskResponse)
	if err = s.doRequest("SubmitPreloadTask", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentTasks(dto *DescribeContentTasksRequest) (responseBody *DescribeContentTasksResponse, err error) {
	responseBody = new(DescribeContentTasksResponse)
	if err = s.doRequest("DescribeContentTasks", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentQuota() (responseBody *DescribeContentQuotaResponse, err error) {
	responseBody = new(DescribeContentQuotaResponse)
	if err = s.doRequest("DescribeContentQuota", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitBlockTask(dto *SubmitBlockTaskRequest) (responseBody *SubmitBlockTaskResponse, err error) {
	responseBody = new(SubmitBlockTaskResponse)
	if err = s.doRequest("SubmitBlockTask", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitUnblockTask(dto *SubmitUnblockTaskRequest) (responseBody *SubmitUnblockTaskResponse, err error) {
	responseBody = new(SubmitUnblockTaskResponse)
	if err = s.doRequest("SubmitUnblockTask", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentBlockTasks(dto *DescribeContentBlockTasksRequest) (responseBody *DescribeContentBlockTasksResponse, err error) {
	responseBody = new(DescribeContentBlockTasksResponse)
	if err = s.doRequest("DescribeContentBlockTasks", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnAccessLog(dto *DescribeCdnAccessLogRequest) (responseBody *DescribeCdnAccessLogResponse, err error) {
	responseBody = new(DescribeCdnAccessLogResponse)
	if err = s.doRequest("DescribeCdnAccessLog", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeIPInfo(dto *DescribeIPInfoRequest) (responseBody *DescribeIPInfoResponse, err error) {
	responseBody = new(DescribeIPInfoResponse)
	if err = s.doRequest("DescribeIPInfo", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnUpperIp(dto *DescribeCdnUpperIpRequest) (responseBody *DescribeCdnUpperIpResponse, err error) {
	responseBody = new(DescribeCdnUpperIpResponse)
	if err = s.doRequest("DescribeCdnUpperIp", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) AddResourceTags(dto *AddResourceTagsRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("AddResourceTags", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateResourceTags(dto *UpdateResourceTagsRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("UpdateResourceTags", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListResourceTags() (responseBody *ListResourceTagsResponse, err error) {
	responseBody = new(ListResourceTagsResponse)
	if err = s.doRequest("ListResourceTags", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteResourceTags(dto *DeleteResourceTagsRequest) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if err = s.doRequest("DeleteResourceTags", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

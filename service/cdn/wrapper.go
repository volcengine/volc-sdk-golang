package cdn

func (s *CDN) AddCdnDomain(dto *AddCdnDomainRequest) (responseBody *AddCdnDomainResponse, err error) {
	responseBody = new(AddCdnDomainResponse)
	if err = s.post("AddCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StartCdnDomain(dto *StartCdnDomainRequest) (responseBody *StartCdnDomainResponse, err error) {
	responseBody = new(StartCdnDomainResponse)
	if err = s.post("StartCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StopCdnDomain(dto *StopCdnDomainRequest) (responseBody *StopCdnDomainResponse, err error) {
	responseBody = new(StopCdnDomainResponse)
	if err = s.post("StopCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnDomain(dto *DeleteCdnDomainRequest) (responseBody *DeleteCdnDomainResponse, err error) {
	responseBody = new(DeleteCdnDomainResponse)
	if err = s.post("DeleteCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnDomains(dto *ListCdnDomainsRequest, options ...OptionArg) (responseBody *ListCdnDomainsResponse, err error) {
	responseBody = new(ListCdnDomainsResponse)
	if err = s.post("ListCdnDomains", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnConfig(dto *DescribeCdnConfigRequest) (responseBody *DescribeCdnConfigResponse, err error) {
	responseBody = new(DescribeCdnConfigResponse)
	if err = s.post("DescribeCdnConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateCdnConfig(dto *UpdateCdnConfigRequest) (responseBody *UpdateCdnConfigResponse, err error) {
	responseBody = new(UpdateCdnConfigResponse)
	if err = s.post("UpdateCdnConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnData(dto *DescribeCdnDataRequest, options ...OptionArg) (responseBody *DescribeCdnDataResponse, err error) {
	responseBody = new(DescribeCdnDataResponse)
	if err = s.post("DescribeCdnData", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeNrtDataSummary(dto *DescribeEdgeNrtDataSummaryRequest, options ...OptionArg) (responseBody *DescribeEdgeNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeEdgeNrtDataSummaryResponse)
	if err = s.post("DescribeEdgeNrtDataSummary", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnOriginData(dto *DescribeCdnOriginDataRequest, options ...OptionArg) (responseBody *DescribeCdnOriginDataResponse, err error) {
	responseBody = new(DescribeCdnOriginDataResponse)
	if err = s.post("DescribeCdnOriginData", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginNrtDataSummary(dto *DescribeOriginNrtDataSummaryRequest, options ...OptionArg) (responseBody *DescribeOriginNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeOriginNrtDataSummaryResponse)
	if err = s.post("DescribeOriginNrtDataSummary", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnDataDetail(dto *DescribeCdnDataDetailRequest, options ...OptionArg) (responseBody *DescribeCdnDataDetailResponse, err error) {
	responseBody = new(DescribeCdnDataDetailResponse)
	if err = s.post("DescribeCdnDataDetail", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeDistrictIspData(dto *DescribeDistrictIspDataRequest, options ...OptionArg) (responseBody *DescribeDistrictIspDataResponse, err error) {
	responseBody = new(DescribeDistrictIspDataResponse)
	if err = s.post("DescribeDistrictIspData", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeStatisticalData(dto *DescribeEdgeStatisticalDataRequest, options ...OptionArg) (responseBody *DescribeEdgeStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeStatisticalDataResponse)
	if err = s.post("DescribeEdgeStatisticalData", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopNrtData(dto *DescribeEdgeTopNrtDataRequest, options ...OptionArg) (responseBody *DescribeEdgeTopNrtDataResponse, err error) {
	responseBody = new(DescribeEdgeTopNrtDataResponse)
	if err = s.post("DescribeEdgeTopNrtData", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginTopNrtData(dto *DescribeOriginTopNrtDataRequest, options ...OptionArg) (responseBody *DescribeOriginTopNrtDataResponse, err error) {
	responseBody = new(DescribeOriginTopNrtDataResponse)
	if err = s.post("DescribeOriginTopNrtData", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopStatusCode(dto *DescribeEdgeTopStatusCodeRequest, options ...OptionArg) (responseBody *DescribeEdgeTopStatusCodeResponse, err error) {
	responseBody = new(DescribeEdgeTopStatusCodeResponse)
	if err = s.post("DescribeEdgeTopStatusCode", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginTopStatusCode(dto *DescribeOriginTopStatusCodeRequest, options ...OptionArg) (responseBody *DescribeOriginTopStatusCodeResponse, err error) {
	responseBody = new(DescribeOriginTopStatusCodeResponse)
	if err = s.post("DescribeOriginTopStatusCode", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopStatisticalData(dto *DescribeEdgeTopStatisticalDataRequest, options ...OptionArg) (responseBody *DescribeEdgeTopStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeTopStatisticalDataResponse)
	if err = s.post("DescribeEdgeTopStatisticalData", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnRegionAndIsp(dto *DescribeCdnRegionAndIspRequest, options ...OptionArg) (responseBody *DescribeCdnRegionAndIspResponse, err error) {
	responseBody = new(DescribeCdnRegionAndIspResponse)
	if err = s.post("DescribeCdnRegionAndIsp", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnService() (responseBody *DescribeCdnServiceResponse, err error) {
	responseBody = new(DescribeCdnServiceResponse)
	if err = s.post("DescribeCdnService", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeAccountingData(dto *DescribeAccountingDataRequest) (responseBody *DescribeAccountingDataResponse, err error) {
	responseBody = new(DescribeAccountingDataResponse)
	if err = s.post("DescribeAccountingData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitRefreshTask(dto *SubmitRefreshTaskRequest) (responseBody *SubmitRefreshTaskResponse, err error) {
	responseBody = new(SubmitRefreshTaskResponse)
	if err = s.post("SubmitRefreshTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitPreloadTask(dto *SubmitPreloadTaskRequest) (responseBody *SubmitPreloadTaskResponse, err error) {
	responseBody = new(SubmitPreloadTaskResponse)
	if err = s.post("SubmitPreloadTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentTasks(dto *DescribeContentTasksRequest) (responseBody *DescribeContentTasksResponse, err error) {
	responseBody = new(DescribeContentTasksResponse)
	if err = s.post("DescribeContentTasks", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentQuota(options ...OptionArg) (responseBody *DescribeContentQuotaResponse, err error) {
	responseBody = new(DescribeContentQuotaResponse)
	if err = s.post("DescribeContentQuota", nil, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitBlockTask(dto *SubmitBlockTaskRequest) (responseBody *SubmitBlockTaskResponse, err error) {
	responseBody = new(SubmitBlockTaskResponse)
	if err = s.post("SubmitBlockTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitUnblockTask(dto *SubmitUnblockTaskRequest) (responseBody *SubmitUnblockTaskResponse, err error) {
	responseBody = new(SubmitUnblockTaskResponse)
	if err = s.post("SubmitUnblockTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentBlockTasks(dto *DescribeContentBlockTasksRequest) (responseBody *DescribeContentBlockTasksResponse, err error) {
	responseBody = new(DescribeContentBlockTasksResponse)
	if err = s.post("DescribeContentBlockTasks", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnAccessLog(dto *DescribeCdnAccessLogRequest, options ...OptionArg) (responseBody *DescribeCdnAccessLogResponse, err error) {
	responseBody = new(DescribeCdnAccessLogResponse)
	if err = s.post("DescribeCdnAccessLog", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeIPInfo(dto *DescribeIPInfoRequest) (responseBody *DescribeIPInfoResponse, err error) {
	responseBody = new(DescribeIPInfoResponse)
	if err = s.post("DescribeIPInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeIPListInfo(dto *DescribeIPListInfoRequest) (responseBody *DescribeIPListInfoResponse, err error) {
	responseBody = new(DescribeIPListInfoResponse)
	if err = s.post("DescribeIPListInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnUpperIp(dto *DescribeCdnUpperIpRequest) (responseBody *DescribeCdnUpperIpResponse, err error) {
	responseBody = new(DescribeCdnUpperIpResponse)
	if err = s.post("DescribeCdnUpperIp", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) AddResourceTags(dto *AddResourceTagsRequest) (responseBody *AddResourceTagsResponse, err error) {
	responseBody = new(AddResourceTagsResponse)
	if err = s.post("AddResourceTags", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateResourceTags(dto *UpdateResourceTagsRequest) (responseBody *UpdateResourceTagsResponse, err error) {
	responseBody = new(UpdateResourceTagsResponse)
	if err = s.post("UpdateResourceTags", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListResourceTags() (responseBody *ListResourceTagsResponse, err error) {
	responseBody = new(ListResourceTagsResponse)
	if err = s.post("ListResourceTags", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteResourceTags(dto *DeleteResourceTagsRequest) (responseBody *DeleteResourceTagsResponse, err error) {
	responseBody = new(DeleteResourceTagsResponse)
	if err = s.post("DeleteResourceTags", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) AddCdnCertificate(dto *AddCdnCertificateRequest) (responseBody *AddCdnCertificateResponse, err error) {
	responseBody = new(AddCdnCertificateResponse)
	if err = s.post("AddCdnCertificate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCertInfo(dto *ListCertInfoRequest) (responseBody *ListCertInfoResponse, err error) {
	responseBody = new(ListCertInfoResponse)
	if err = s.post("ListCertInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnCertInfo(dto *ListCdnCertInfoRequest) (responseBody *ListCdnCertInfoResponse, err error) {
	responseBody = new(ListCdnCertInfoResponse)
	if err = s.post("ListCdnCertInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCertConfig(dto *DescribeCertConfigRequest) (responseBody *DescribeCertConfigResponse, err error) {
	responseBody = new(DescribeCertConfigResponse)
	if err = s.post("DescribeCertConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) BatchDeployCert(dto *BatchDeployCertRequest) (responseBody *BatchDeployCertResponse, err error) {
	responseBody = new(BatchDeployCertResponse)
	if err = s.post("BatchDeployCert", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnCertificate(dto *DeleteCdnCertificateRequest) (responseBody *DeleteCdnCertificateResponse, err error) {
	responseBody = new(DeleteCdnCertificateResponse)
	if err = s.post("DeleteCdnCertificate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeAccountingSummary(dto *DescribeAccountingSummaryRequest) (responseBody *DescribeAccountingSummaryResponse, err error) {
	responseBody = new(DescribeAccountingSummaryResponse)
	if err = s.post("DescribeAccountingSummary", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

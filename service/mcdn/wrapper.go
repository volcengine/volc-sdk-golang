package mcdn

func (s *MCDN) ListCloudAccounts(req ListCloudAccountsRequest, opts ...withRequestOption) (resp *ListCloudAccountsResponseWrapper, err error) {
	resp = new(ListCloudAccountsResponseWrapper)
	if err = s.doRequest("ListCloudAccounts", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) ListCdnDomains(req ListDomainsRequest, opts ...withRequestOption) (resp *ListDomainsResponseWrapper, err error) {
	resp = new(ListDomainsResponseWrapper)
	if err = s.doRequest("ListCdnDomains", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) ListDnsSchedules(req ListDnsSchedulesRequest, opts ...withRequestOption) (resp *ListDnsSchedulesResponseWrapper, err error) {
	resp = new(ListDnsSchedulesResponseWrapper)
	if err = s.doRequest("ListDnsSchedules", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) DescribeDnsSchedule(req DescribeDnsScheduleRequest, opts ...withRequestOption) (resp *DescribeDnsScheduleResponseWrapper, err error) {
	resp = new(DescribeDnsScheduleResponseWrapper)
	if err = s.doRequest("DescribeDnsSchedule", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) DescribeCdnDataOffline(req DescribeCdnDataOfflineRequest, opts ...withRequestOption) (resp *DescribeCdnDataOfflineResponseWrapper, err error) {
	resp = new(DescribeCdnDataOfflineResponseWrapper)
	if err = s.doRequest("DescribeCdnDataOffline", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) DescribeCdnOriginDataOffline(req DescribeCdnDataOfflineRequest, opts ...withRequestOption) (resp *DescribeCdnDataOfflineResponseWrapper, err error) {
	resp = new(DescribeCdnDataOfflineResponseWrapper)
	if err = s.doRequest("DescribeCdnOriginDataOffline", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) SubmitRefreshTask(req SubmitRefreshTaskRequest, opts ...withRequestOption) (resp *SubmitTaskResponseWrapper, err error) {
	resp = new(SubmitTaskResponseWrapper)
	if err = s.doRequest("SubmitRefreshTask", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) SubmitPreloadTask(req SubmitPreloadTaskRequest, opts ...withRequestOption) (resp *SubmitTaskResponseWrapper, err error) {
	resp = new(SubmitTaskResponseWrapper)
	if err = s.doRequest("SubmitPreloadTask", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) ListContentTask(req ListContentTaskRequest, opts ...withRequestOption) (resp *ListContentTaskResponseWrapper, err error) {
	resp = new(ListContentTaskResponseWrapper)
	if err = s.doRequest("ListContentTask", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) DescribeContentTaskByTaskId(req DescribeContentTaskByTaskIdRequest, opts ...withRequestOption) (resp *DescribeContentTaskByTaskIdResponseWrapper, err error) {
	resp = new(DescribeContentTaskByTaskIdResponseWrapper)
	if err = s.doRequest("DescribeContentTaskByTaskId", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) ListVendorContentTask(req ListVendorContentTaskRequest, opts ...withRequestOption) (resp *ListVendorContentTaskResponseWrapper, err error) {
	resp = new(ListVendorContentTaskResponseWrapper)
	if err = s.doRequest("ListVendorContentTask", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) DescribeContentQuota(req DescribeContentQuotaRequest, opts ...withRequestOption) (resp *DescribeContentQuotaResponseWrapper, err error) {
	resp = new(DescribeContentQuotaResponseWrapper)
	if err = s.doRequest("DescribeContentQuota", req, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) ListViews(opts ...withRequestOption) (resp *ListViewsResponseWrapper, err error) {
	resp = new(ListViewsResponseWrapper)
	if err = s.doRequest("ListViews", nil, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *MCDN) DescribeCdnRegionAndIsp(opts ...withRequestOption) (resp *DescribeCdnRegionAndIspResponseWrapper, err error) {
	resp = new(DescribeCdnRegionAndIspResponseWrapper)
	if err = s.doRequest("DescribeCdnRegionAndIsp", nil, &resp, opts); err != nil {
		return
	}
	if err = validateResponse(resp.ResponseMetadata); err != nil {
		return
	}
	return
}

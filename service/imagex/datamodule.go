package imagex

import (
	"net/url"
)

func (c *ImageX) DescribeImageXDomainTrafficData(req *DescribeImageXDomainTrafficDataReq) (*DescribeImageXDomainTrafficDataResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXDomainTrafficDataResp{}
	err = c.ImageXGet("DescribeImageXDomainTrafficData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXDomainBandwidthData(req *DescribeImageXDomainBandwidthDataReq) (*DescribeImageXDomainBandwidthDataResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXDomainBandwidthDataResp{}
	err = c.ImageXGet("DescribeImageXDomainBandwidthData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXBucketUsage(req *DescribeImageXBucketUsageReq) (*DescribeImageXBucketUsageResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXBucketUsageResp{}
	err = c.ImageXGet("DescribeImageXBucketUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXRequestCntUsage(req *DescribeImageXRequestCntUsageReq) (*DescribeImageXRequestCntUsageResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXRequestCntUsageResp{}
	err = c.ImageXGet("DescribeImageXRequestCntUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXBaseOpUsage(req *DescribeImageXBaseOpUsageReq) (*DescribeImageXBaseOpUsageResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXBaseOpUsageResp{}
	err = c.ImageXGet("DescribeImageXBaseOpUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCompressUsage(req *DescribeImageXCompressUsageReq) (*DescribeImageXCompressUsageResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXCompressUsageResp{}
	err = c.ImageXGet("DescribeImageXCompressUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXEdgeRequest(req *DescribeImageXEdgeRequestReq) (*DescribeImageXEdgeRequestResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXEdgeRequestResp{}
	err = c.ImageXGet("DescribeImageXEdgeRequest", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXMirrorRequestTraffic(req *DescribeImageXMirrorRequestTrafficReq) (*DescribeImageXMirrorRequestTrafficResp, error) {
	resp := &DescribeImageXMirrorRequestTrafficResp{}
	err := c.ImageXPost("DescribeImageXMirrorRequestTraffic", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXMirrorRequestBandwidth(req *DescribeImageXMirrorRequestBandwidthReq) (*DescribeImageXMirrorRequestBandwidthResp, error) {
	resp := &DescribeImageXMirrorRequestBandwidthResp{}
	err := c.ImageXPost("DescribeImageXMirrorRequestBandwidth", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXMirrorRequestHttpCodeByTime(req *DescribeImageXMirrorRequestHttpCodeByTimeReq) (*DescribeImageXMirrorRequestHttpCodeByTimeResp, error) {
	resp := &DescribeImageXMirrorRequestHttpCodeByTimeResp{}
	err := c.ImageXPost("DescribeImageXMirrorRequestHttpCodeByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXHitRateTrafficData(req *DescribeImageXHitRateTrafficDataReq) (*DescribeImageXHitRateTrafficDataResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXHitRateTrafficDataResp{}
	err = c.ImageXGet("DescribeImageXHitRateTrafficData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXHitRateRequestData(req *DescribeImageXHitRateRequestDataReq) (*DescribeImageXHitRateRequestDataResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXHitRateRequestDataResp{}
	err = c.ImageXGet("DescribeImageXHitRateRequestData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCDNTopRequestData(req *DescribeImageXCDNTopRequestDataReq) (*DescribeImageXCDNTopRequestDataResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXCDNTopRequestDataResp{}
	err = c.ImageXGet("DescribeImageXCDNTopRequestData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXSummary(req *DescribeImageXSummaryReq) (*DescribeImageXSummaryResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXSummaryResp{}
	err = c.ImageXGet("DescribeImageXSummary", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXMirrorRequestHttpCodeOverview(req *DescribeImageXMirrorRequestHttpCodeOverviewReq) (*DescribeImageXMirrorRequestHttpCodeOverviewResp, error) {
	resp := &DescribeImageXMirrorRequestHttpCodeOverviewResp{}
	err := c.ImageXPost("DescribeImageXMirrorRequestHttpCodeOverview", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXEdgeRequestBandwidth(req *DescribeImageXEdgeRequestBandwidthReq) (*DescribeImageXEdgeRequestBandwidthResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXEdgeRequestBandwidthResp{}
	err = c.ImageXGet("DescribeImageXEdgeRequestBandwidth", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXEdgeRequestTraffic(req *DescribeImageXEdgeRequestTrafficReq) (*DescribeImageXEdgeRequestTrafficResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXEdgeRequestTrafficResp{}
	err = c.ImageXGet("DescribeImageXEdgeRequestTraffic", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXEdgeRequestRegions(req *DescribeImageXEdgeRequestRegionsReq) (*DescribeImageXEdgeRequestRegionsResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXEdgeRequestRegionsResp{}
	err = c.ImageXGet("DescribeImageXEdgeRequestRegions", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXServiceQuality(req *DescribeImageXServiceQualityReq) (*DescribeImageXServiceQualityResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXServiceQualityResp{}
	err = c.ImageXGet("DescribeImageXServiceQuality", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageXQueryApps(req *GetImageXQueryAppsReq) (*GetImageXQueryAppsResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageXQueryAppsResp{}
	err = c.ImageXGet("GetImageXQueryApps", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageXQueryRegions(req *GetImageXQueryRegionsReq) (*GetImageXQueryRegionsResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageXQueryRegionsResp{}
	err = c.ImageXGet("GetImageXQueryRegions", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageXQueryDims(req *GetImageXQueryDimsReq) (*GetImageXQueryDimsResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageXQueryDimsResp{}
	err = c.ImageXGet("GetImageXQueryDims", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageXQueryVals(req *GetImageXQueryValsReq) (*GetImageXQueryValsResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageXQueryValsResp{}
	err = c.ImageXGet("GetImageXQueryVals", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadSuccessRateByTime(req *DescribeImageXUploadSuccessRateByTimeReq) (*DescribeImageXUploadSuccessRateByTimeResp, error) {
	resp := &DescribeImageXUploadSuccessRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXUploadSuccessRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadErrorCodeAll(req *DescribeImageXUploadErrorCodeAllReq) (*DescribeImageXUploadErrorCodeAllResp, error) {
	resp := &DescribeImageXUploadErrorCodeAllResp{}
	err := c.ImageXPost("DescribeImageXUploadErrorCodeAll", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadErrorCodeByTime(req *DescribeImageXUploadErrorCodeByTimeReq) (*DescribeImageXUploadErrorCodeByTimeResp, error) {
	resp := &DescribeImageXUploadErrorCodeByTimeResp{}
	err := c.ImageXPost("DescribeImageXUploadErrorCodeByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadCountByTime(req *DescribeImageXUploadCountByTimeReq) (*DescribeImageXUploadCountByTimeResp, error) {
	resp := &DescribeImageXUploadCountByTimeResp{}
	err := c.ImageXPost("DescribeImageXUploadCountByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadFileSize(req *DescribeImageXUploadFileSizeReq) (*DescribeImageXUploadFileSizeResp, error) {
	resp := &DescribeImageXUploadFileSizeResp{}
	err := c.ImageXPost("DescribeImageXUploadFileSize", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadSpeed(req *DescribeImageXUploadSpeedReq) (*DescribeImageXUploadSpeedResp, error) {
	resp := &DescribeImageXUploadSpeedResp{}
	err := c.ImageXPost("DescribeImageXUploadSpeed", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadDuration(req *DescribeImageXUploadDurationReq) (*DescribeImageXUploadDurationResp, error) {
	resp := &DescribeImageXUploadDurationResp{}
	err := c.ImageXPost("DescribeImageXUploadDuration", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXUploadSegmentSpeedByTime(req *DescribeImageXUploadSegmentSpeedByTimeReq) (*DescribeImageXUploadSegmentSpeedByTimeResp, error) {
	resp := &DescribeImageXUploadSegmentSpeedByTimeResp{}
	err := c.ImageXPost("DescribeImageXUploadSegmentSpeedByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnSuccessRateByTime(req *DescribeImageXCdnSuccessRateByTimeReq) (*DescribeImageXCdnSuccessRateByTimeResp, error) {
	resp := &DescribeImageXCdnSuccessRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXCdnSuccessRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnSuccessRateAll(req *DescribeImageXCdnSuccessRateAllReq) (*DescribeImageXCdnSuccessRateAllResp, error) {
	resp := &DescribeImageXCdnSuccessRateAllResp{}
	err := c.ImageXPost("DescribeImageXCdnSuccessRateAll", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnErrorCodeByTime(req *DescribeImageXCdnErrorCodeByTimeReq) (*DescribeImageXCdnErrorCodeByTimeResp, error) {
	resp := &DescribeImageXCdnErrorCodeByTimeResp{}
	err := c.ImageXPost("DescribeImageXCdnErrorCodeByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnErrorCodeAll(req *DescribeImageXCdnErrorCodeAllReq) (*DescribeImageXCdnErrorCodeAllResp, error) {
	resp := &DescribeImageXCdnErrorCodeAllResp{}
	err := c.ImageXPost("DescribeImageXCdnErrorCodeAll", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnDurationDetailByTime(req *DescribeImageXCdnDurationDetailByTimeReq) (*DescribeImageXCdnDurationDetailByTimeResp, error) {
	resp := &DescribeImageXCdnDurationDetailByTimeResp{}
	err := c.ImageXPost("DescribeImageXCdnDurationDetailByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnDurationAll(req *DescribeImageXCdnDurationAllReq) (*DescribeImageXCdnDurationAllResp, error) {
	resp := &DescribeImageXCdnDurationAllResp{}
	err := c.ImageXPost("DescribeImageXCdnDurationAll", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnReuseRateByTime(req *DescribeImageXCdnReuseRateByTimeReq) (*DescribeImageXCdnReuseRateByTimeResp, error) {
	resp := &DescribeImageXCdnReuseRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXCdnReuseRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnReuseRateAll(req *DescribeImageXCdnReuseRateAllReq) (*DescribeImageXCdnReuseRateAllResp, error) {
	resp := &DescribeImageXCdnReuseRateAllResp{}
	err := c.ImageXPost("DescribeImageXCdnReuseRateAll", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXCdnProtocolRateByTime(req *DescribeImageXCdnProtocolRateByTimeReq) (*DescribeImageXCdnProtocolRateByTimeResp, error) {
	resp := &DescribeImageXCdnProtocolRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXCdnProtocolRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientErrorCodeAll(req *DescribeImageXClientErrorCodeAllReq) (*DescribeImageXClientErrorCodeAllResp, error) {
	resp := &DescribeImageXClientErrorCodeAllResp{}
	err := c.ImageXPost("DescribeImageXClientErrorCodeAll", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientErrorCodeByTime(req *DescribeImageXClientErrorCodeByTimeReq) (*DescribeImageXClientErrorCodeByTimeResp, error) {
	resp := &DescribeImageXClientErrorCodeByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientErrorCodeByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientDecodeSuccessRateByTime(req *DescribeImageXClientDecodeSuccessRateByTimeReq) (*DescribeImageXClientDecodeSuccessRateByTimeResp, error) {
	resp := &DescribeImageXClientDecodeSuccessRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientDecodeSuccessRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientDecodeDurationByTime(req *DescribeImageXClientDecodeDurationByTimeReq) (*DescribeImageXClientDecodeDurationByTimeResp, error) {
	resp := &DescribeImageXClientDecodeDurationByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientDecodeDurationByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientQueueDurationByTime(req *DescribeImageXClientQueueDurationByTimeReq) (*DescribeImageXClientQueueDurationByTimeResp, error) {
	resp := &DescribeImageXClientQueueDurationByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientQueueDurationByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientLoadDurationAll(req *DescribeImageXClientLoadDurationAllReq) (*DescribeImageXClientLoadDurationAllResp, error) {
	resp := &DescribeImageXClientLoadDurationAllResp{}
	err := c.ImageXPost("DescribeImageXClientLoadDurationAll", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientLoadDuration(req *DescribeImageXClientLoadDurationReq) (*DescribeImageXClientLoadDurationResp, error) {
	resp := &DescribeImageXClientLoadDurationResp{}
	err := c.ImageXPost("DescribeImageXClientLoadDuration", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientFailureRate(req *DescribeImageXClientFailureRateReq) (*DescribeImageXClientFailureRateResp, error) {
	resp := &DescribeImageXClientFailureRateResp{}
	err := c.ImageXPost("DescribeImageXClientFailureRate", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientSdkVerByTime(req *DescribeImageXClientSdkVerByTimeReq) (*DescribeImageXClientSdkVerByTimeResp, error) {
	resp := &DescribeImageXClientSdkVerByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientSdkVerByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientFileSize(req *DescribeImageXClientFileSizeReq) (*DescribeImageXClientFileSizeResp, error) {
	resp := &DescribeImageXClientFileSizeResp{}
	err := c.ImageXPost("DescribeImageXClientFileSize", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientTopFileSize(req *DescribeImageXClientTopFileSizeReq) (*DescribeImageXClientTopFileSizeResp, error) {
	resp := &DescribeImageXClientTopFileSizeResp{}
	err := c.ImageXPost("DescribeImageXClientTopFileSize", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientCountByTime(req *DescribeImageXClientCountByTimeReq) (*DescribeImageXClientCountByTimeResp, error) {
	resp := &DescribeImageXClientCountByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientCountByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientScoreByTime(req *DescribeImageXClientScoreByTimeReq) (*DescribeImageXClientScoreByTimeResp, error) {
	resp := &DescribeImageXClientScoreByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientScoreByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientDemotionRateByTime(req *DescribeImageXClientDemotionRateByTimeReq) (*DescribeImageXClientDemotionRateByTimeResp, error) {
	resp := &DescribeImageXClientDemotionRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientDemotionRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientTopDemotionURL(req *DescribeImageXClientTopDemotionURLReq) (*DescribeImageXClientTopDemotionURLResp, error) {
	resp := &DescribeImageXClientTopDemotionURLResp{}
	err := c.ImageXPost("DescribeImageXClientTopDemotionURL", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientQualityRateByTime(req *DescribeImageXClientQualityRateByTimeReq) (*DescribeImageXClientQualityRateByTimeResp, error) {
	resp := &DescribeImageXClientQualityRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXClientQualityRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXClientTopQualityURL(req *DescribeImageXClientTopQualityURLReq) (*DescribeImageXClientTopQualityURLResp, error) {
	resp := &DescribeImageXClientTopQualityURLResp{}
	err := c.ImageXPost("DescribeImageXClientTopQualityURL", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXSensibleCountByTime(req *DescribeImageXSensibleCountByTimeReq) (*DescribeImageXSensibleCountByTimeResp, error) {
	resp := &DescribeImageXSensibleCountByTimeResp{}
	err := c.ImageXPost("DescribeImageXSensibleCountByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXSensibleCacheHitRateByTime(req *DescribeImageXSensibleCacheHitRateByTimeReq) (*DescribeImageXSensibleCacheHitRateByTimeResp, error) {
	resp := &DescribeImageXSensibleCacheHitRateByTimeResp{}
	err := c.ImageXPost("DescribeImageXSensibleCacheHitRateByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXSensibleTopSizeURL(req *DescribeImageXSensibleTopSizeURLReq) (*DescribeImageXSensibleTopSizeURLResp, error) {
	resp := &DescribeImageXSensibleTopSizeURLResp{}
	err := c.ImageXPost("DescribeImageXSensibleTopSizeURL", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXSensibleTopRamURL(req *DescribeImageXSensibleTopRamURLReq) (*DescribeImageXSensibleTopRamURLResp, error) {
	resp := &DescribeImageXSensibleTopRamURLResp{}
	err := c.ImageXPost("DescribeImageXSensibleTopRamURL", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXSensibleTopResolutionURL(req *DescribeImageXSensibleTopResolutionURLReq) (*DescribeImageXSensibleTopResolutionURLResp, error) {
	resp := &DescribeImageXSensibleTopResolutionURLResp{}
	err := c.ImageXPost("DescribeImageXSensibleTopResolutionURL", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) DescribeImageXSensibleTopUnknownURL(req *DescribeImageXSensibleTopUnknownURLReq) (*DescribeImageXSensibleTopUnknownURLResp, error) {
	resp := &DescribeImageXSensibleTopUnknownURLResp{}
	err := c.ImageXPost("DescribeImageXSensibleTopUnknownURL", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

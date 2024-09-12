package imagex

import (
	"context"
	"net/url"
)

func (c *Imagex) DelDomain(ctx context.Context, arg *DelDomainReq) (*DelDomainRes, error) {
	query, err := marshalToQuery(arg.DelDomainQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DelDomainBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DelDomain", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DelDomainRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateRefer(ctx context.Context, arg *UpdateReferReq) (*UpdateReferRes, error) {
	query, err := marshalToQuery(arg.UpdateReferQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateReferBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRefer", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateReferRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateHTTPS(ctx context.Context, arg *UpdateHTTPSReq) (*UpdateHTTPSRes, error) {
	query, err := marshalToQuery(arg.UpdateHTTPSQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateHTTPSBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHttps", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateHTTPSRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateResponseHeader(ctx context.Context, arg *UpdateResponseHeaderReq) (*UpdateResponseHeaderRes, error) {
	query, err := marshalToQuery(arg.UpdateResponseHeaderQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateResponseHeaderBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateResponseHeader", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateResponseHeaderRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) SetDefaultDomain(ctx context.Context, arg *SetDefaultDomainBody) (*SetDefaultDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SetDefaultDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SetDefaultDomainRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageVolcCdnAccessLog(ctx context.Context, arg *DescribeImageVolcCdnAccessLogReq) (*DescribeImageVolcCdnAccessLogRes, error) {
	query, err := marshalToQuery(arg.DescribeImageVolcCdnAccessLogQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageVolcCdnAccessLogBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageVolcCdnAccessLog", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageVolcCdnAccessLogRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetResponseHeaderValidateKeys(ctx context.Context) (*GetResponseHeaderValidateKeysRes, error) {

	data, _, err := c.CtxQuery(ctx, "GetResponseHeaderValidateKeys", url.Values{})
	if err != nil {
		return nil, err
	}

	result := new(GetResponseHeaderValidateKeysRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetDomainConfig(ctx context.Context, arg *GetDomainConfigQuery) (*GetDomainConfigRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetDomainConfig", query)
	if err != nil {
		return nil, err
	}

	result := new(GetDomainConfigRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetServiceDomains(ctx context.Context, arg *GetServiceDomainsQuery) (*GetServiceDomainsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetServiceDomains", query)
	if err != nil {
		return nil, err
	}

	result := new(GetServiceDomainsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageMigrateTask(ctx context.Context, arg *CreateImageMigrateTaskBody) (*CreateImageMigrateTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageMigrateTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageMigrateTask(ctx context.Context, arg *DeleteImageMigrateTaskQuery) (*DeleteImageMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageMigrateTask", query, "")
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) ExportFailedMigrateTask(ctx context.Context, arg *ExportFailedMigrateTaskQuery) (*ExportFailedMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ExportFailedMigrateTask", query)
	if err != nil {
		return nil, err
	}

	result := new(ExportFailedMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageTaskStrategy(ctx context.Context, arg *UpdateImageTaskStrategyBody) (*UpdateImageTaskStrategyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageTaskStrategy", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageTaskStrategyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) TerminateImageMigrateTask(ctx context.Context, arg *TerminateImageMigrateTaskQuery) (*TerminateImageMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "TerminateImageMigrateTask", query, "")
	if err != nil {
		return nil, err
	}

	result := new(TerminateImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetVendorBuckets(ctx context.Context, arg *GetVendorBucketsBody) (*GetVendorBucketsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetVendorBuckets", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetVendorBucketsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageMigrateTasks(ctx context.Context, arg *GetImageMigrateTasksQuery) (*GetImageMigrateTasksRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageMigrateTasks", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageMigrateTasksRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) RerunImageMigrateTask(ctx context.Context, arg *RerunImageMigrateTaskQuery) (*RerunImageMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RerunImageMigrateTask", query, "")
	if err != nil {
		return nil, err
	}

	result := new(RerunImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSourceRequestBandwidth(ctx context.Context, arg *DescribeImageXSourceRequestBandwidthQuery) (*DescribeImageXSourceRequestBandwidthRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSourceRequestBandwidth", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSourceRequestBandwidthRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSourceRequestTraffic(ctx context.Context, arg *DescribeImageXSourceRequestTrafficQuery) (*DescribeImageXSourceRequestTrafficRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSourceRequestTraffic", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSourceRequestTrafficRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSourceRequest(ctx context.Context, arg *DescribeImageXSourceRequestQuery) (*DescribeImageXSourceRequestRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSourceRequest", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSourceRequestRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBucketRetrievalUsage(ctx context.Context, arg *DescribeImageXBucketRetrievalUsageQuery) (*DescribeImageXBucketRetrievalUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBucketRetrievalUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBucketRetrievalUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSummary(ctx context.Context, arg *DescribeImageXSummaryQuery) (*DescribeImageXSummaryRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSummary", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSummaryRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXDomainTrafficData(ctx context.Context, arg *DescribeImageXDomainTrafficDataQuery) (*DescribeImageXDomainTrafficDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXDomainTrafficData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXDomainTrafficDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXDomainBandwidthData(ctx context.Context, arg *DescribeImageXDomainBandwidthDataQuery) (*DescribeImageXDomainBandwidthDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXDomainBandwidthData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXDomainBandwidthDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXDomainBandwidthNinetyFiveData(ctx context.Context, arg *DescribeImageXDomainBandwidthNinetyFiveDataQuery) (*DescribeImageXDomainBandwidthNinetyFiveDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXDomainBandwidthNinetyFiveData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXDomainBandwidthNinetyFiveDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBucketUsage(ctx context.Context, arg *DescribeImageXBucketUsageQuery) (*DescribeImageXBucketUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBucketUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBucketUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBillingRequestCntUsage(ctx context.Context, arg *DescribeImageXBillingRequestCntUsageQuery) (*DescribeImageXBillingRequestCntUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBillingRequestCntUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBillingRequestCntUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXRequestCntUsage(ctx context.Context, arg *DescribeImageXRequestCntUsageQuery) (*DescribeImageXRequestCntUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXRequestCntUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXRequestCntUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBaseOpUsage(ctx context.Context, arg *DescribeImageXBaseOpUsageQuery) (*DescribeImageXBaseOpUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBaseOpUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBaseOpUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCompressUsage(ctx context.Context, arg *DescribeImageXCompressUsageQuery) (*DescribeImageXCompressUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXCompressUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCompressUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXScreenshotUsage(ctx context.Context, arg *DescribeImageXScreenshotUsageQuery) (*DescribeImageXScreenshotUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXScreenshotUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXScreenshotUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXVideoClipDurationUsage(ctx context.Context, arg *DescribeImageXVideoClipDurationUsageQuery) (*DescribeImageXVideoClipDurationUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXVideoClipDurationUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXVideoClipDurationUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMultiCompressUsage(ctx context.Context, arg *DescribeImageXMultiCompressUsageQuery) (*DescribeImageXMultiCompressUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXMultiCompressUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMultiCompressUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequest(ctx context.Context, arg *DescribeImageXEdgeRequestQuery) (*DescribeImageXEdgeRequestRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequest", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequestBandwidth(ctx context.Context, arg *DescribeImageXEdgeRequestBandwidthQuery) (*DescribeImageXEdgeRequestBandwidthRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequestBandwidth", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestBandwidthRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequestTraffic(ctx context.Context, arg *DescribeImageXEdgeRequestTrafficQuery) (*DescribeImageXEdgeRequestTrafficRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequestTraffic", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestTrafficRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequestRegions(ctx context.Context, arg *DescribeImageXEdgeRequestRegionsQuery) (*DescribeImageXEdgeRequestRegionsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequestRegions", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestRegionsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestHTTPCodeByTime(ctx context.Context, arg *DescribeImageXMirrorRequestHTTPCodeByTimeBody) (*DescribeImageXMirrorRequestHTTPCodeByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestHttpCodeByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestHTTPCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestHTTPCodeOverview(ctx context.Context, arg *DescribeImageXMirrorRequestHTTPCodeOverviewBody) (*DescribeImageXMirrorRequestHTTPCodeOverviewRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestHttpCodeOverview", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestHTTPCodeOverviewRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestTraffic(ctx context.Context, arg *DescribeImageXMirrorRequestTrafficBody) (*DescribeImageXMirrorRequestTrafficRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestTraffic", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestTrafficRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestBandwidth(ctx context.Context, arg *DescribeImageXMirrorRequestBandwidthBody) (*DescribeImageXMirrorRequestBandwidthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestBandwidth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestBandwidthRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXServerQPSUsage(ctx context.Context, arg *DescribeImageXServerQPSUsageQuery) (*DescribeImageXServerQPSUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXServerQPSUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXServerQPSUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHitRateTrafficData(ctx context.Context, arg *DescribeImageXHitRateTrafficDataQuery) (*DescribeImageXHitRateTrafficDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXHitRateTrafficData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHitRateTrafficDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHitRateRequestData(ctx context.Context, arg *DescribeImageXHitRateRequestDataQuery) (*DescribeImageXHitRateRequestDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXHitRateRequestData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHitRateRequestDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCDNTopRequestData(ctx context.Context, arg *DescribeImageXCDNTopRequestDataQuery) (*DescribeImageXCDNTopRequestDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXCDNTopRequestData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCDNTopRequestDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXExceedResolutionRatioAll(ctx context.Context, arg *DescribeImageXExceedResolutionRatioAllBody) (*DescribeImageXExceedResolutionRatioAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXExceedResolutionRatioAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXExceedResolutionRatioAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXExceedFileSize(ctx context.Context, arg *DescribeImageXExceedFileSizeBody) (*DescribeImageXExceedFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXExceedFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXExceedFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXExceedCountByTime(ctx context.Context, arg *DescribeImageXExceedCountByTimeBody) (*DescribeImageXExceedCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXExceedCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXExceedCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXServiceQuality(ctx context.Context, arg *DescribeImageXServiceQualityQuery) (*DescribeImageXServiceQualityRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXServiceQuality", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXServiceQualityRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryApps(ctx context.Context, arg *GetImageXQueryAppsQuery) (*GetImageXQueryAppsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryApps", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryAppsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryRegions(ctx context.Context, arg *GetImageXQueryRegionsQuery) (*GetImageXQueryRegionsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryRegions", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryRegionsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryDims(ctx context.Context, arg *GetImageXQueryDimsQuery) (*GetImageXQueryDimsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryDims", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryDimsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryVals(ctx context.Context, arg *GetImageXQueryValsQuery) (*GetImageXQueryValsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryVals", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryValsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadCountByTime(ctx context.Context, arg *DescribeImageXUploadCountByTimeBody) (*DescribeImageXUploadCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadDuration(ctx context.Context, arg *DescribeImageXUploadDurationBody) (*DescribeImageXUploadDurationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadDuration", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadDurationRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadSuccessRateByTime(ctx context.Context, arg *DescribeImageXUploadSuccessRateByTimeBody) (*DescribeImageXUploadSuccessRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadSuccessRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadSuccessRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadFileSize(ctx context.Context, arg *DescribeImageXUploadFileSizeBody) (*DescribeImageXUploadFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadErrorCodeByTime(ctx context.Context, arg *DescribeImageXUploadErrorCodeByTimeBody) (*DescribeImageXUploadErrorCodeByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadErrorCodeByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadErrorCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadErrorCodeAll(ctx context.Context, arg *DescribeImageXUploadErrorCodeAllBody) (*DescribeImageXUploadErrorCodeAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadErrorCodeAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadErrorCodeAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadSpeed(ctx context.Context, arg *DescribeImageXUploadSpeedBody) (*DescribeImageXUploadSpeedRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadSpeed", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadSpeedRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadSegmentSpeedByTime(ctx context.Context, arg *DescribeImageXUploadSegmentSpeedByTimeBody) (*DescribeImageXUploadSegmentSpeedByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadSegmentSpeedByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadSegmentSpeedByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnSuccessRateByTime(ctx context.Context, arg *DescribeImageXCdnSuccessRateByTimeBody) (*DescribeImageXCdnSuccessRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnSuccessRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnSuccessRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnSuccessRateAll(ctx context.Context, arg *DescribeImageXCdnSuccessRateAllBody) (*DescribeImageXCdnSuccessRateAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnSuccessRateAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnSuccessRateAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnErrorCodeByTime(ctx context.Context, arg *DescribeImageXCdnErrorCodeByTimeBody) (*DescribeImageXCdnErrorCodeByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnErrorCodeByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnErrorCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnErrorCodeAll(ctx context.Context, arg *DescribeImageXCdnErrorCodeAllBody) (*DescribeImageXCdnErrorCodeAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnErrorCodeAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnErrorCodeAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnDurationDetailByTime(ctx context.Context, arg *DescribeImageXCdnDurationDetailByTimeBody) (*DescribeImageXCdnDurationDetailByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnDurationDetailByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnDurationDetailByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnDurationAll(ctx context.Context, arg *DescribeImageXCdnDurationAllBody) (*DescribeImageXCdnDurationAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnDurationAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnDurationAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnReuseRateByTime(ctx context.Context, arg *DescribeImageXCdnReuseRateByTimeBody) (*DescribeImageXCdnReuseRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnReuseRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnReuseRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnReuseRateAll(ctx context.Context, arg *DescribeImageXCdnReuseRateAllBody) (*DescribeImageXCdnReuseRateAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnReuseRateAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnReuseRateAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnProtocolRateByTime(ctx context.Context, arg *DescribeImageXCdnProtocolRateByTimeBody) (*DescribeImageXCdnProtocolRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnProtocolRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnProtocolRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientFailureRate(ctx context.Context, arg *DescribeImageXClientFailureRateBody) (*DescribeImageXClientFailureRateRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientFailureRate", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientFailureRateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientDecodeSuccessRateByTime(ctx context.Context, arg *DescribeImageXClientDecodeSuccessRateByTimeBody) (*DescribeImageXClientDecodeSuccessRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientDecodeSuccessRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientDecodeSuccessRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientDecodeDurationByTime(ctx context.Context, arg *DescribeImageXClientDecodeDurationByTimeBody) (*DescribeImageXClientDecodeDurationByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientDecodeDurationByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientDecodeDurationByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientQueueDurationByTime(ctx context.Context, arg *DescribeImageXClientQueueDurationByTimeBody) (*DescribeImageXClientQueueDurationByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientQueueDurationByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientQueueDurationByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientErrorCodeByTime(ctx context.Context, arg *DescribeImageXClientErrorCodeByTimeBody) (*DescribeImageXClientErrorCodeByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientErrorCodeByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientErrorCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientErrorCodeAll(ctx context.Context, arg *DescribeImageXClientErrorCodeAllBody) (*DescribeImageXClientErrorCodeAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientErrorCodeAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientErrorCodeAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientLoadDuration(ctx context.Context, arg *DescribeImageXClientLoadDurationBody) (*DescribeImageXClientLoadDurationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientLoadDuration", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientLoadDurationRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientLoadDurationAll(ctx context.Context, arg *DescribeImageXClientLoadDurationAllBody) (*DescribeImageXClientLoadDurationAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientLoadDurationAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientLoadDurationAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientSdkVerByTime(ctx context.Context, arg *DescribeImageXClientSdkVerByTimeBody) (*DescribeImageXClientSdkVerByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientSdkVerByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientSdkVerByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientFileSize(ctx context.Context, arg *DescribeImageXClientFileSizeBody) (*DescribeImageXClientFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientTopFileSize(ctx context.Context, arg *DescribeImageXClientTopFileSizeBody) (*DescribeImageXClientTopFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientTopFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientTopFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientCountByTime(ctx context.Context, arg *DescribeImageXClientCountByTimeBody) (*DescribeImageXClientCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientQualityRateByTime(ctx context.Context, arg *DescribeImageXClientQualityRateByTimeBody) (*DescribeImageXClientQualityRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientQualityRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientQualityRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientTopQualityURL(ctx context.Context, arg *DescribeImageXClientTopQualityURLBody) (*DescribeImageXClientTopQualityURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientTopQualityURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientTopQualityURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientDemotionRateByTime(ctx context.Context, arg *DescribeImageXClientDemotionRateByTimeBody) (*DescribeImageXClientDemotionRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientDemotionRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientDemotionRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientTopDemotionURL(ctx context.Context, arg *DescribeImageXClientTopDemotionURLBody) (*DescribeImageXClientTopDemotionURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientTopDemotionURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientTopDemotionURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientScoreByTime(ctx context.Context, arg *DescribeImageXClientScoreByTimeBody) (*DescribeImageXClientScoreByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientScoreByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientScoreByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleCountByTime(ctx context.Context, arg *DescribeImageXSensibleCountByTimeBody) (*DescribeImageXSensibleCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleCacheHitRateByTime(ctx context.Context, arg *DescribeImageXSensibleCacheHitRateByTimeBody) (*DescribeImageXSensibleCacheHitRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleCacheHitRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleCacheHitRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopSizeURL(ctx context.Context, arg *DescribeImageXSensibleTopSizeURLBody) (*DescribeImageXSensibleTopSizeURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopSizeURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopSizeURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopResolutionURL(ctx context.Context, arg *DescribeImageXSensibleTopResolutionURLBody) (*DescribeImageXSensibleTopResolutionURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopResolutionURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopResolutionURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopRAMURL(ctx context.Context, arg *DescribeImageXSensibleTopRAMURLBody) (*DescribeImageXSensibleTopRAMURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopRamURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopRAMURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopUnknownURL(ctx context.Context, arg *DescribeImageXSensibleTopUnknownURLBody) (*DescribeImageXSensibleTopUnknownURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopUnknownURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopUnknownURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageResourceStatus(ctx context.Context, arg *UpdateImageResourceStatusReq) (*UpdateImageResourceStatusRes, error) {
	query, err := marshalToQuery(arg.UpdateImageResourceStatusQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageResourceStatusBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageResourceStatus", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageResourceStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageStorageFiles(ctx context.Context, arg *GetImageStorageFilesQuery) (*GetImageStorageFilesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageStorageFiles", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageStorageFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageUploadFiles(ctx context.Context, arg *DeleteImageUploadFilesReq) (*DeleteImageUploadFilesRes, error) {
	query, err := marshalToQuery(arg.DeleteImageUploadFilesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageUploadFilesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageUploadFiles", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageUploadFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageUploadFiles(ctx context.Context, arg *UpdateImageUploadFilesReq) (*UpdateImageUploadFilesRes, error) {
	query, err := marshalToQuery(arg.UpdateImageUploadFilesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageUploadFilesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageUploadFiles", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageUploadFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CommitImageUpload(ctx context.Context, arg *CommitImageUploadReq) (*CommitImageUploadRes, error) {
	query, err := marshalToQuery(arg.CommitImageUploadQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CommitImageUploadBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CommitImageUpload", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CommitImageUploadRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) ApplyImageUpload(ctx context.Context, arg *ApplyImageUploadQuery) (*ApplyImageUploadRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ApplyImageUpload", query)
	if err != nil {
		return nil, err
	}

	result := new(ApplyImageUploadRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageUploadFile(ctx context.Context, arg *GetImageUploadFileQuery) (*GetImageUploadFileRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageUploadFile", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageUploadFileRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageUploadFiles(ctx context.Context, arg *GetImageUploadFilesQuery) (*GetImageUploadFilesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageUploadFiles", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageUploadFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageUpdateFiles(ctx context.Context, arg *GetImageUpdateFilesQuery) (*GetImageUpdateFilesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageUpdateFiles", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageUpdateFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) PreviewImageUploadFile(ctx context.Context, arg *PreviewImageUploadFileQuery) (*PreviewImageUploadFileRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "PreviewImageUploadFile", query)
	if err != nil {
		return nil, err
	}

	result := new(PreviewImageUploadFileRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageService(ctx context.Context, arg *GetImageServiceQuery) (*GetImageServiceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageService", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageServiceRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAllImageServices(ctx context.Context, arg *GetAllImageServicesQuery) (*GetAllImageServicesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAllImageServices", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAllImageServicesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageCompressTask(ctx context.Context, arg *CreateImageCompressTaskReq) (*CreateImageCompressTaskRes, error) {
	query, err := marshalToQuery(arg.CreateImageCompressTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageCompressTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageCompressTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageCompressTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) FetchImageURL(ctx context.Context, arg *FetchImageURLBody) (*FetchImageURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "FetchImageUrl", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(FetchImageURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageStorageTTL(ctx context.Context, arg *UpdateImageStorageTTLBody) (*UpdateImageStorageTTLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageStorageTTL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageStorageTTLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetCompressTaskInfo(ctx context.Context, arg *GetCompressTaskInfoQuery) (*GetCompressTaskInfoRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetCompressTaskInfo", query)
	if err != nil {
		return nil, err
	}

	result := new(GetCompressTaskInfoRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetURLFetchTask(ctx context.Context, arg *GetURLFetchTaskQuery) (*GetURLFetchTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetUrlFetchTask", query)
	if err != nil {
		return nil, err
	}

	result := new(GetURLFetchTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetResourceURL(ctx context.Context, arg *GetResourceURLQuery) (*GetResourceURLRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetResourceURL", query)
	if err != nil {
		return nil, err
	}

	result := new(GetResourceURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageFromURI(ctx context.Context, arg *CreateImageFromURIReq) (*CreateImageFromURIRes, error) {
	query, err := marshalToQuery(arg.CreateImageFromURIQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageFromURIBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageFromUri", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageFromURIRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageFileKey(ctx context.Context, arg *UpdateImageFileKeyReq) (*UpdateImageFileKeyRes, error) {
	query, err := marshalToQuery(arg.UpdateImageFileKeyQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageFileKeyBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageFileKey", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageFileKeyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageContentTask(ctx context.Context, arg *CreateImageContentTaskBody) (*CreateImageContentTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageContentTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageContentTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageContentTaskDetail(ctx context.Context, arg *GetImageContentTaskDetailBody) (*GetImageContentTaskDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageContentTaskDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageContentTaskDetailRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageContentBlockList(ctx context.Context, arg *GetImageContentBlockListBody) (*GetImageContentBlockListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageContentBlockList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageContentBlockListRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTranscodeQueue(ctx context.Context, arg *CreateImageTranscodeQueueBody) (*CreateImageTranscodeQueueRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTranscodeQueue", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTranscodeQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageTranscodeQueue(ctx context.Context, arg *DeleteImageTranscodeQueueBody) (*DeleteImageTranscodeQueueRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageTranscodeQueue", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageTranscodeQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageTranscodeQueue(ctx context.Context, arg *UpdateImageTranscodeQueueBody) (*UpdateImageTranscodeQueueRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageTranscodeQueue", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageTranscodeQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageTranscodeQueueStatus(ctx context.Context, arg *UpdateImageTranscodeQueueStatusBody) (*UpdateImageTranscodeQueueStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageTranscodeQueueStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageTranscodeQueueStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageTranscodeQueues(ctx context.Context, arg *GetImageTranscodeQueuesQuery) (*GetImageTranscodeQueuesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageTranscodeQueues", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageTranscodeQueuesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTranscodeTask(ctx context.Context, arg *CreateImageTranscodeTaskBody) (*CreateImageTranscodeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTranscodeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTranscodeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageTranscodeDetails(ctx context.Context, arg *GetImageTranscodeDetailsQuery) (*GetImageTranscodeDetailsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageTranscodeDetails", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageTranscodeDetailsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTranscodeCallback(ctx context.Context, arg *CreateImageTranscodeCallbackBody) (*CreateImageTranscodeCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTranscodeCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTranscodeCallbackRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageTranscodeDetail(ctx context.Context, arg *DeleteImageTranscodeDetailBody) (*DeleteImageTranscodeDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageTranscodeDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageTranscodeDetailRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImagePSDetection(ctx context.Context, arg *GetImagePSDetectionReq) (*GetImagePSDetectionRes, error) {
	query, err := marshalToQuery(arg.GetImagePSDetectionQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImagePSDetectionBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImagePSDetection", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImagePSDetectionRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageEraseResult(ctx context.Context, arg *GetImageEraseResultBody) (*GetImageEraseResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageEraseResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageEraseResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageSuperResolutionResult(ctx context.Context, arg *GetImageSuperResolutionResultBody) (*GetImageSuperResolutionResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageSuperResolutionResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageSuperResolutionResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageDuplicateDetection(ctx context.Context, arg *GetImageDuplicateDetectionReq) (*GetImageDuplicateDetectionRes, error) {
	query, err := marshalToQuery(arg.GetImageDuplicateDetectionQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageDuplicateDetectionBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageDuplicateDetection", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageDuplicateDetectionRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageOCRV2(ctx context.Context, arg *GetImageOCRV2Req) (*GetImageOCRV2Res, error) {
	query, err := marshalToQuery(arg.GetImageOCRV2Query)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageOCRV2Body)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageOCRV2", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageOCRV2Res)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageBgFillResult(ctx context.Context, arg *GetImageBgFillResultBody) (*GetImageBgFillResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageBgFillResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageBgFillResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetSegmentImage(ctx context.Context, arg *GetSegmentImageReq) (*GetSegmentImageRes, error) {
	query, err := marshalToQuery(arg.GetSegmentImageQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetSegmentImageBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetSegmentImage", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetSegmentImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageSmartCropResult(ctx context.Context, arg *GetImageSmartCropResultBody) (*GetImageSmartCropResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageSmartCropResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageSmartCropResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageComicResult(ctx context.Context, arg *GetImageComicResultBody) (*GetImageComicResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageComicResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageComicResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageEnhanceResult(ctx context.Context, arg *GetImageEnhanceResultBody) (*GetImageEnhanceResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageEnhanceResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageEnhanceResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageQuality(ctx context.Context, arg *GetImageQualityReq) (*GetImageQualityRes, error) {
	query, err := marshalToQuery(arg.GetImageQualityQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageQualityBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageQuality", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageQualityRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetPrivateImageType(ctx context.Context, arg *GetPrivateImageTypeReq) (*GetPrivateImageTypeRes, error) {
	query, err := marshalToQuery(arg.GetPrivateImageTypeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetPrivateImageTypeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetPrivateImageType", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetPrivateImageTypeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateHiddenWatermarkImage(ctx context.Context, arg *CreateHiddenWatermarkImageReq) (*CreateHiddenWatermarkImageRes, error) {
	query, err := marshalToQuery(arg.CreateHiddenWatermarkImageQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateHiddenWatermarkImageBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateHiddenWatermarkImage", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateHiddenWatermarkImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageExifData(ctx context.Context, arg *UpdateImageExifDataReq) (*UpdateImageExifDataRes, error) {
	query, err := marshalToQuery(arg.UpdateImageExifDataQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageExifDataBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageExifData", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageExifDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageDetectResult(ctx context.Context, arg *GetImageDetectResultReq) (*GetImageDetectResultRes, error) {
	query, err := marshalToQuery(arg.GetImageDetectResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageDetectResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageDetectResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageDetectResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageHmExtract(ctx context.Context, arg *CreateImageHmExtractQuery) (*CreateImageHmExtractRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageHmExtract", query, "")
	if err != nil {
		return nil, err
	}

	result := new(CreateImageHmExtractRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageHmEmbed(ctx context.Context, arg *CreateImageHmEmbedBody) (*CreateImageHmEmbedRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageHmEmbed", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageHmEmbedRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetComprehensiveEnhanceImage(ctx context.Context, arg *GetComprehensiveEnhanceImageBody) (*GetComprehensiveEnhanceImageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetComprehensiveEnhanceImage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetComprehensiveEnhanceImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageEraseModels(ctx context.Context, arg *GetImageEraseModelsQuery) (*GetImageEraseModelsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageEraseModels", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageEraseModelsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetDedupTaskStatus(ctx context.Context, arg *GetDedupTaskStatusQuery) (*GetDedupTaskStatusRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetDedupTaskStatus", query)
	if err != nil {
		return nil, err
	}

	result := new(GetDedupTaskStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageService(ctx context.Context, arg *CreateImageServiceBody) (*CreateImageServiceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageService", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageServiceRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageService(ctx context.Context, arg *DeleteImageServiceQuery) (*DeleteImageServiceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageService", query, "")
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageServiceRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAuthKey(ctx context.Context, arg *UpdateImageAuthKeyReq) (*UpdateImageAuthKeyRes, error) {
	query, err := marshalToQuery(arg.UpdateImageAuthKeyQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageAuthKeyBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAuthKey", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAuthKeyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateServiceName(ctx context.Context, arg *UpdateServiceNameReq) (*UpdateServiceNameRes, error) {
	query, err := marshalToQuery(arg.UpdateServiceNameQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateServiceNameBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateServiceName", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateServiceNameRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageObjectAccess(ctx context.Context, arg *UpdateImageObjectAccessReq) (*UpdateImageObjectAccessRes, error) {
	query, err := marshalToQuery(arg.UpdateImageObjectAccessQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageObjectAccessBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageObjectAccess", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageObjectAccessRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageMirrorConf(ctx context.Context, arg *UpdateImageMirrorConfReq) (*UpdateImageMirrorConfRes, error) {
	query, err := marshalToQuery(arg.UpdateImageMirrorConfQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageMirrorConfBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageMirrorConf", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageMirrorConfRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageServiceSubscription(ctx context.Context, arg *GetImageServiceSubscriptionQuery) (*GetImageServiceSubscriptionRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageServiceSubscription", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageServiceSubscriptionRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAuthKey(ctx context.Context, arg *GetImageAuthKeyQuery) (*GetImageAuthKeyRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAuthKey", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAuthKeyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageAnalyzeTask(ctx context.Context, arg *CreateImageAnalyzeTaskBody) (*CreateImageAnalyzeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageAnalyzeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageAnalyzeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAnalyzeTaskRun(ctx context.Context, arg *DeleteImageAnalyzeTaskRunBody) (*DeleteImageAnalyzeTaskRunRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAnalyzeTaskRun", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAnalyzeTaskRunRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAnalyzeTask(ctx context.Context, arg *DeleteImageAnalyzeTaskBody) (*DeleteImageAnalyzeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAnalyzeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAnalyzeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAnalyzeTaskStatus(ctx context.Context, arg *UpdateImageAnalyzeTaskStatusBody) (*UpdateImageAnalyzeTaskStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAnalyzeTaskStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAnalyzeTaskStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAnalyzeTask(ctx context.Context, arg *UpdateImageAnalyzeTaskBody) (*UpdateImageAnalyzeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAnalyzeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAnalyzeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAnalyzeTasks(ctx context.Context, arg *GetImageAnalyzeTasksQuery) (*GetImageAnalyzeTasksRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAnalyzeTasks", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAnalyzeTasksRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAnalyzeResult(ctx context.Context, arg *GetImageAnalyzeResultQuery) (*GetImageAnalyzeResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAnalyzeResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAnalyzeResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageStyleResult(ctx context.Context, arg *GetImageStyleResultReq) (*GetImageStyleResultRes, error) {
	query, err := marshalToQuery(arg.GetImageStyleResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageStyleResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageStyleResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageStyleResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTemplate(ctx context.Context, arg *CreateImageTemplateReq) (*CreateImageTemplateRes, error) {
	query, err := marshalToQuery(arg.CreateImageTemplateQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageTemplateBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTemplate", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTemplateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteTemplatesFromBin(ctx context.Context, arg *DeleteTemplatesFromBinReq) (*DeleteTemplatesFromBinRes, error) {
	query, err := marshalToQuery(arg.DeleteTemplatesFromBinQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteTemplatesFromBinBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTemplatesFromBin", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTemplatesFromBinRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageTemplate(ctx context.Context, arg *DeleteImageTemplateReq) (*DeleteImageTemplateRes, error) {
	query, err := marshalToQuery(arg.DeleteImageTemplateQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageTemplateBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageTemplate", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageTemplateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateTemplatesFromBin(ctx context.Context, arg *CreateTemplatesFromBinReq) (*CreateTemplatesFromBinRes, error) {
	query, err := marshalToQuery(arg.CreateTemplatesFromBinQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateTemplatesFromBinBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTemplatesFromBin", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTemplatesFromBinRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageTemplate(ctx context.Context, arg *GetImageTemplateQuery) (*GetImageTemplateRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageTemplate", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageTemplateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetTemplatesFromBin(ctx context.Context, arg *GetTemplatesFromBinQuery) (*GetTemplatesFromBinRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetTemplatesFromBin", query)
	if err != nil {
		return nil, err
	}

	result := new(GetTemplatesFromBinRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAllImageTemplates(ctx context.Context, arg *GetAllImageTemplatesQuery) (*GetAllImageTemplatesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAllImageTemplates", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAllImageTemplatesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageAuditTask(ctx context.Context, arg *CreateImageAuditTaskBody) (*CreateImageAuditTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageAuditTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAuditResult(ctx context.Context, arg *DeleteImageAuditResultBody) (*DeleteImageAuditResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAuditResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAuditResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAuditTaskStatus(ctx context.Context, arg *UpdateImageAuditTaskStatusBody) (*UpdateImageAuditTaskStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAuditTaskStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAuditTaskStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAuditTask(ctx context.Context, arg *UpdateImageAuditTaskBody) (*UpdateImageAuditTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAuditTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateAuditImageStatus(ctx context.Context, arg *UpdateAuditImageStatusBody) (*UpdateAuditImageStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateAuditImageStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAuditImageStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAuditTasks(ctx context.Context, arg *GetImageAuditTasksQuery) (*GetImageAuditTasksRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAuditTasks", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAuditTasksRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAuditResult(ctx context.Context, arg *GetImageAuditResultQuery) (*GetImageAuditResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAuditResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAuditResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAuditEntrysCount(ctx context.Context, arg *GetAuditEntrysCountQuery) (*GetAuditEntrysCountRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAuditEntrysCount", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAuditEntrysCountRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageRetryAuditTask(ctx context.Context, arg *CreateImageRetryAuditTaskBody) (*CreateImageRetryAuditTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageRetryAuditTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageRetryAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

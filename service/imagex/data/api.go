package data

import (
	"net/http"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imagex"
)

func AddDataModule(c *imagex.ImageX) error {
	g := []string{"DescribeImageXDomainTrafficData",
		"DescribeImageXDomainBandwidthData",
		"DescribeImageXBucketUsage",
		"DescribeImageXRequestCntUsage",
		"DescribeImageXBaseOpUsage",
		"DescribeImageXCompressUsage",
		"DescribeImageXEdgeRequest",
		"DescribeImageXHitRateTrafficData",
		"DescribeImageXHitRateRequestData",
		"DescribeImageXCDNTopRequestData",
		"DescribeImageXSummary",
	}
	p := []string{"DescribeImageXMirrorRequestTraffic",
		"DescribeImageXMirrorRequestBandwidth",
		"DescribeImageXMirrorRequestHttpCodeByTime",
		"DescribeImageXMirrorRequestHttpCodeOverview",
	}
	for _, i := range g {
		c.ApiInfoList[i] = &base.ApiInfo{
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{i},
				"Version": []string{"2018-08-01"},
			},
		}
	}
	for _, i := range p {
		c.ApiInfoList[i] = &base.ApiInfo{
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{i},
				"Version": []string{"2018-08-01"},
			},
		}
	}
	return nil
}

func DescribeImageXDomainTrafficData(instance *imagex.ImageX, req *DescribeImageXDomainTrafficDataReq) (*DescribeImageXDomainTrafficDataResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXDomainTrafficDataResp{}
	err = instance.ImageXGet("DescribeImageXDomainTrafficData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXDomainBandwidthData(instance *imagex.ImageX, req *DescribeImageXDomainBandwidthDataReq) (*DescribeImageXDomainBandwidthDataResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXDomainBandwidthDataResp{}
	err = instance.ImageXGet("DescribeImageXDomainBandwidthData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXBucketUsage(instance *imagex.ImageX, req *DescribeImageXBucketUsageReq) (*DescribeImageXBucketUsageResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXBucketUsageResp{}
	err = instance.ImageXGet("DescribeImageXBucketUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXRequestCntUsage(instance *imagex.ImageX, req *DescribeImageXRequestCntUsageReq) (*DescribeImageXRequestCntUsageResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXRequestCntUsageResp{}
	err = instance.ImageXGet("DescribeImageXRequestCntUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXBaseOpUsage(instance *imagex.ImageX, req *DescribeImageXBaseOpUsageReq) (*DescribeImageXBaseOpUsageResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXBaseOpUsageResp{}
	err = instance.ImageXGet("DescribeImageXBaseOpUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXCompressUsage(instance *imagex.ImageX, req *DescribeImageXCompressUsageReq) (*DescribeImageXCompressUsageResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXCompressUsageResp{}
	err = instance.ImageXGet("DescribeImageXCompressUsage", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXEdgeRequest(instance *imagex.ImageX, req *DescribeImageXEdgeRequestReq) (*DescribeImageXEdgeRequestResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXEdgeRequestResp{}
	err = instance.ImageXGet("DescribeImageXEdgeRequest", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXMirrorRequestTraffic(instance *imagex.ImageX, req *DescribeImageXMirrorRequestTrafficReq) (*DescribeImageXMirrorRequestTrafficResp, error) {
	resp := &DescribeImageXMirrorRequestTrafficResp{}
	err := instance.ImageXPost("DescribeImageXMirrorRequestTraffic", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXMirrorRequestBandwidth(instance *imagex.ImageX, req *DescribeImageXMirrorRequestBandwidthReq) (*DescribeImageXMirrorRequestBandwidthResp, error) {
	resp := &DescribeImageXMirrorRequestBandwidthResp{}
	err := instance.ImageXPost("DescribeImageXMirrorRequestBandwidth", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXMirrorRequestHttpCodeByTime(instance *imagex.ImageX, req *DescribeImageXMirrorRequestHttpCodeByTimeReq) (*DescribeImageXMirrorRequestHttpCodeByTimeResp, error) {
	resp := &DescribeImageXMirrorRequestHttpCodeByTimeResp{}
	err := instance.ImageXPost("DescribeImageXMirrorRequestHttpCodeByTime", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXHitRateTrafficData(instance *imagex.ImageX, req *DescribeImageXHitRateTrafficDataReq) (*DescribeImageXHitRateTrafficDataResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXHitRateTrafficDataResp{}
	err = instance.ImageXGet("DescribeImageXHitRateTrafficData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXHitRateRequestData(instance *imagex.ImageX, req *DescribeImageXHitRateRequestDataReq) (*DescribeImageXHitRateRequestDataResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXHitRateRequestDataResp{}
	err = instance.ImageXGet("DescribeImageXHitRateRequestData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXCDNTopRequestData(instance *imagex.ImageX, req *DescribeImageXCDNTopRequestDataReq) (*DescribeImageXCDNTopRequestDataResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXCDNTopRequestDataResp{}
	err = instance.ImageXGet("DescribeImageXCDNTopRequestData", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXSummary(instance *imagex.ImageX, req *DescribeImageXSummaryReq) (*DescribeImageXSummaryResp, error) {
	query, err := marshalToQuery(req)
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageXSummaryResp{}
	err = instance.ImageXGet("DescribeImageXSummary", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DescribeImageXMirrorRequestHttpCodeOverview(instance *imagex.ImageX, req *DescribeImageXMirrorRequestHttpCodeOverviewReq) (*DescribeImageXMirrorRequestHttpCodeOverviewResp, error) {
	resp := &DescribeImageXMirrorRequestHttpCodeOverviewResp{}
	err := instance.ImageXPost("DescribeImageXMirrorRequestHttpCodeOverview", url.Values{}, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

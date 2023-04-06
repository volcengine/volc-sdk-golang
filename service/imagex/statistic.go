package imagex

func (c *ImageX) DescribeImageVolcCdnAccessLog(req *DescribeImageVolcCdnAccessLogReq) (*DescribeImageVolcCdnAccessLogResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &DescribeImageVolcCdnAccessLogResp{}
	err = c.ImageXPost("DescribeImageVolcCdnAccessLog", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

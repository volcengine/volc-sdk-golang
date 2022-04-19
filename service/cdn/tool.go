package cdn

func (s *CDN) SendCommonRequest(action string, in interface{}, out interface{}, reqOptions ...RequestOption) error {
	return s.doRequest(action, in, out, reqOptions...)
}

func (s *CDN) ValidateResponse(meta *ResponseMetadata) error {
	return validateResponse(meta)
}

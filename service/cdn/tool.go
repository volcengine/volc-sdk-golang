package cdn

func (s *CDN) SendCommonRequest(action string, in interface{}, out interface{}) error {
	return s.post(action, in, out)
}

func (s *CDN) ValidateResponse(meta *ResponseMetadata) error {
	return validateResponse(meta)
}

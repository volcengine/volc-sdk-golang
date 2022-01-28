package cdn

import "github.com/volcengine/volc-sdk-golang/base"

func (s *CDN) SendCommonRequest(action string, in interface{}, out interface{}) error {
	return s.post(action, in, out)
}

func (s *CDN) ValidateResponse(meta *base.ResponseMetadata) error {
	return validateResponse(meta)
}

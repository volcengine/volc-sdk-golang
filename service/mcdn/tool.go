package mcdn

import "github.com/volcengine/volc-sdk-golang/base"

func (s *MCDN) SendCommonRequest(action string, in interface{}, out interface{}, opts ...withRequestOption) error {
	return s.doRequest(action, in, out, opts)
}

func (s *MCDN) ValidateResponse(meta base.ResponseMetadata) error {
	return validateResponse(meta)
}

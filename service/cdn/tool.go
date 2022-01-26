package cdn

func (s *CDN) SendCommonRequest(action string, in interface{}, out interface{}) error {
	return s.post(action, in, out)
}

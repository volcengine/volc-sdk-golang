package cdn

import "errors"

func (s *CDN) DescribeCdnUpperIp(dto *DescribeCdnUpperIpParam) (responseBody *DescribeCdnUpperIpResponse, err error) {
	responseBody = new(DescribeCdnUpperIpResponse)
	if dto.Domain == "" {
		err = errors.New("domain name cannot be empty")
		return
	}
	if err = s.post("DescribeCdnUpperIp", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

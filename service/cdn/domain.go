package cdn

import (
	"errors"
)

func (s *CDN) StartCdnDomain(dto *StartCdnDomainParam) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if dto.Domain == "" {
		err = errors.New("domain name cannot be empty")
		return
	}
	if err = s.post("StartCdnDomain", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StopCdnDomain(dto *StopCdnDomainParam) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if dto.Domain == "" {
		err = errors.New("domain name cannot be empty")
		return
	}
	if err = s.post("StopCdnDomain", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnDomain(dto *DeleteCdnDomainParam) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if dto.Domain == "" {
		err = errors.New("domain name cannot be empty")
		return
	}
	if err = s.post("DeleteCdnDomain", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnDomains(dto *ListCdnDomainsParam) (responseBody *ListCdnDomainsResponse, err error) {
	responseBody = new(ListCdnDomainsResponse)
	if err = s.post("ListCdnDomains", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

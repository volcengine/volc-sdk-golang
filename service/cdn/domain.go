package cdn

import (
	"errors"
)

func (s *CDN) StartCdnDomain(dto *StartCdnDomainParam, queryOptions ...QueryOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if dto.DomainName == "" {
		err = errors.New("domain name cannot be empty")
		return
	}
	if err = s.post("StartCdnDomain", &dto, responseBody, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StopCdnDomain(dto *StopCdnDomainParam, queryOptions ...QueryOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if dto.DomainName == "" {
		err = errors.New("domain name cannot be empty")
		return
	}
	if err = s.post("StopCdnDomain", &dto, responseBody, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnDomain(dto *DeleteCdnDomainParam, queryOptions ...QueryOption) (responseBody *EmptyResponse, err error) {
	responseBody = new(EmptyResponse)
	if dto.DomainName == "" {
		err = errors.New("domain name cannot be empty")
		return
	}
	if err = s.post("DeleteCdnDomain", &dto, responseBody, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnDomains(dto *ListCdnDomainsParam, queryOptions ...QueryOption) (responseBody *ListCdnDomainsResponse, err error) {
	responseBody = new(ListCdnDomainsResponse)
	if err = s.post("ListCdnDomains", &dto, responseBody, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

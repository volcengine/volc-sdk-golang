package cdn

import (
	"errors"
)

func (s *CDN) StartCdnDomain(dto *StartDomainParam, queryOptions ...QueryOption) (responseBody *EmptyResponse, err error) {
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

func (s *CDN) StopCdnDomain(dto *StopDomainParam, queryOptions ...QueryOption) (responseBody *EmptyResponse, err error) {
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

func (s *CDN) DeleteCdnDomain(dto *DeleteDomainParam, queryOptions ...QueryOption) (responseBody *EmptyResponse, err error) {
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

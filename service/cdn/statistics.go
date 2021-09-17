package cdn

import (
	"errors"
)

func (s *CDN) DescribeCdnData(dto *DescribeCdnDataParam, queryOptions ...QueryOption) (response *DescribeCdnDataResponse, err error) {
	response = new(DescribeCdnDataResponse)
	if dto.EndTime == 0 {
		err = errors.New("end time is required")
		return
	}
	if dto.StartTime == 0 {
		err = errors.New("start time is required")
		return
	}
	if dto.Metric == "" {
		err = errors.New("metric name is required")
		return
	}
	if err = s.post("DescribeCdnData", dto, response, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(response.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnOriginData(dto *DescribeCdnOriginDataParam, queryOptions ...QueryOption) (response *DescribeCdnOriginDataResponse, err error) {
	response = new(DescribeCdnOriginDataResponse)
	if dto.EndTime == 0 {
		err = errors.New("end time is required")
		return
	}
	if dto.StartTime == 0 {
		err = errors.New("start time is required")
		return
	}
	if dto.Metric == "" {
		err = errors.New("metric name is required")
		return
	}
	if err = s.post("DescribeCdnOriginData", dto, response, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(response.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnRegionAndIsp(queryOptions ...QueryOption) (response *DescribeCdnRegionAndIspResponse, err error) {
	response = new(DescribeCdnRegionAndIspResponse)
	if err = s.post("DescribeCdnRegionAndIsp", nil, response, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(response.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnDomainTopData(dto *DescribeCdnDomainTopDataParam, queryOptions ...QueryOption) (response *DescribeCdnDomainTopDataResponse, err error) {
	response = new(DescribeCdnDomainTopDataResponse)
	if dto.EndTime == 0 {
		err = errors.New("end time is required")
		return
	}
	if dto.StartTime == 0 {
		err = errors.New("start time is required")
		return
	}
	if dto.Metric == "" {
		err = errors.New("metric name is required")
		return
	}
	if dto.Domain == "" {
		err = errors.New("domain name is required")
		return
	}
	if dto.Item == "" {
		err = errors.New("item name is required")
		return
	}
	if err = s.post("DescribeCdnDomainTopData", dto, response, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(response.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnAccountingData(dto *DescribeCdnAccountingDataParam, queryOptions ...QueryOption) (response *DescribeCdnAccountingDataResponse, err error) {
	response = new(DescribeCdnAccountingDataResponse)
	if dto.EndTime == 0 {
		err = errors.New("end time is required")
		return
	}
	if dto.StartTime == 0 {
		err = errors.New("start time is required")
		return
	}
	if dto.Domain == "" {
		err = errors.New("domain name is required")
		return
	}
	if err = s.post("DescribeCdnAccountingData", dto, response, queryOptions...); err != nil {
		return
	}
	if err = validateResponse(response.ResponseMetadata); err != nil {
		return
	}
	return
}

package cdn

import "errors"

func (s *CDN) DescribeCdnAccessLog(dto *DescribeCdnAccessLogParam) (response *DescribeCdnAccessLogResponse, err error) {
	response = new(DescribeCdnAccessLogResponse)
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
	if err = s.post("DescribeCdnAccessLog", dto, response); err != nil {
		return
	}
	if err = validateResponse(response.ResponseMetadata); err != nil {
		return
	}
	return
}

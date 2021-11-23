package cdn

import (
	"github.com/pkg/errors"
	"strings"
)

func (s *CDN) SubmitRefreshTask(dto *SubmitRefreshTaskParam) (responseBody *SubmitRefreshTaskResponse, err error) {
	responseBody = new(SubmitRefreshTaskResponse)
	if dto.Type != "file" && dto.Type != "dir" {
		err = errors.New("type can only be filled by either 'file' or 'dir'")
		return
	}
	if len(dto.Urls) == 0 {
		err = errors.New("urls cannot be empty")
		return
	}
	requestBody := SubmitRefreshTaskRequest{
		Type: dto.Type,
		Urls: strings.Join(dto.Urls, "\n"),
	}
	if err = s.post("SubmitRefreshTask", &requestBody, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitPreloadTask(dto *SubmitPreloadTaskParam) (responseBody *SubmitPreloadTaskResponse, err error) {
	responseBody = new(SubmitPreloadTaskResponse)
	if len(dto.Urls) == 0 {
		err = errors.New("urls cannot be empty")
		return
	}
	requestBody := SubmitPreloadTaskRequest{
		Urls: strings.Join(dto.Urls, "\n"),
	}
	if err = s.post("SubmitPreloadTask", &requestBody, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentTasks(dto *DescribeContentTasksParam) (responseBody *DescribeContentTasksResponse, err error) {
	responseBody = new(DescribeContentTasksResponse)
	if err = s.post("DescribeContentTasks", &dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentQuota() (responseBody *DescribeContentQuotaResponse, err error) {
	responseBody = new(DescribeContentQuotaResponse)
	if err = s.post("DescribeContentQuota", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

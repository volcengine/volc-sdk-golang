package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type VideoEditSubmitTaskData struct {
	TaskId string `json:"task_id"`
}

type VideoEditQueryTaskData struct {
	Status   string `json:"status"`
	VideoUrl string `json:"video_url"`
}

type VideoEditSubmitTaskResult struct {
	ResponseMetadata *base.ResponseMetadata   `json:",omitempty"`
	RequestId        string                   `json:"request_id"`
	TimeElapsed      string                   `json:"time_elapsed"`
	Code             int                      `json:"code"`
	Message          string                   `json:"message"`
	Data             *VideoEditSubmitTaskData `json:"data"`
}

type VideoEditQueryTaskResult struct {
	ResponseMetadata *base.ResponseMetadata  `json:",omitempty"`
	RequestId        string                  `json:"request_id"`
	TimeElapsed      string                  `json:"time_elapsed"`
	Code             int                     `json:"code"`
	Message          string                  `json:"message"`
	Data             *VideoEditQueryTaskData `json:"data"`
}

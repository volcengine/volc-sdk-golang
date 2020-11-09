package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type JPCartoonData struct {
	Image string `json:"image"`
	Clip  string `json:"clip"`
}

type JPCartoonResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *JPCartoonData         `json:"data"`
}

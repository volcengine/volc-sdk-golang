package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type EyeClose2OpenData struct {
	Image string `json:"image"`
}

type EyeClose2OpenResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *EyeClose2OpenData     `json:"data"`
}

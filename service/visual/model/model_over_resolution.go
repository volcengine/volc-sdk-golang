package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type OverResolutionData struct {
	ImgData string `json:"img_data"`
}

type OverResolutionResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *OverResolutionData    `json:"data"`
}

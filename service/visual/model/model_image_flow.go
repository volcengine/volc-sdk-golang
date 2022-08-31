package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type ImageFlowData struct {
	Video string `json:"video"`
}

type ImageFlowResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *ImageFlowData         `json:"data"`
}

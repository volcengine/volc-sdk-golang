package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type FaceSwapAIResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`

	Code        int                   `json:"code"`
	Data        *VisualBaseRespDataV2 `json:"data,omitempty"`
	Message     string                `json:"message"`
	RequestId   string                `json:"request_id"`
	Status      int                   `json:"status"`
	TimeElapsed string                `json:"time_elapsed"`
}

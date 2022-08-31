package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type CarSegmentData struct {
	Mask string `json:"mask"`
}

type CarSegmentResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *CarSegmentData        `json:"data"`
}

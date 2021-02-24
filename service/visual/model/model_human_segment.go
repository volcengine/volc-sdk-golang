package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type HumanSegmentData struct {
	Mask string `json:"mask"`
}

type HumanSegmentResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *HumanSegmentData      `json:"data"`
}

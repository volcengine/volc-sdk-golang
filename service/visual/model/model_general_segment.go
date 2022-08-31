package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type GeneralSegmentData struct {
	Mask string `json:"mask"`
}

type GeneralSegmentResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *GeneralSegmentData    `json:"data"`
}

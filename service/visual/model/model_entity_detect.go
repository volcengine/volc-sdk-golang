package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type EntityDetectData struct {
	Entities []string `json:"entities"`
}

type EntityDetectResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *EntityDetectData      `json:"data"`
}

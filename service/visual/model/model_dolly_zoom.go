package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type DollyZoomData struct {
	Video string `json:"video"`
}

type DollyZoomResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *DollyZoomData         `json:"data"`
}

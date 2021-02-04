package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type EnhancePhotoData struct {
	Image string `json:"image"`
}

type EnhancePhotoResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *EnhancePhotoData      `json:"data"`
}

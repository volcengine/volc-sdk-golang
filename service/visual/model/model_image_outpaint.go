package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type ImageOutpaintData struct {
	ImgData string `json:"img_data"`
}

type ImageOutpaintResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *ImageOutpaintData     `json:"data"`
}

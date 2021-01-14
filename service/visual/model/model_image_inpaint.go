package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type ImageInpaintData struct {
	ImgData string `json:"img_data"`
}

type ImageInpaintResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *ImageInpaintData      `json:"data"`
}

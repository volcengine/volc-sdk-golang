package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type VisionInfo struct {
	Url  string `json:"url"`
	Data string `json:"data"`
}

type ImageCutBoundingBox struct {
	Xmin float64 `json:"x_min"`
	Xmax float64 `json:"x_max"`
	Ymin float64 `json:"y_min"`
	Ymax float64 `json:"y_max"`
}

type ImageCutData struct {
	Image       *VisionInfo          `json:"image"`
	BoundingBox *ImageCutBoundingBox `json:"boundingbox"`
}

type ImageCutResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *ImageCutData          `json:"data"`
}

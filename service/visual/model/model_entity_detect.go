package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type EntityObject struct {
	Xmin      float64 `json:"x_min"`
	Xmax      float64 `json:"x_max"`
	Ymin      float64 `json:"y_min"`
	Ymax      float64 `json:"y_max"`
	Prob      float64 `json:"prob"`
	ClassID   int     `json:"class_id"`
	ClassName string  `json:"class_name"`
}

type EntityFrame struct {
	Objects []EntityObject `json:"objects"`
	Height  int            `json:"height"`
	Width   int            `json:"width"`
}

type EntityDetectData struct {
	Entities *EntityFrame `json:"entities"`
}

type EntityDetectResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *EntityDetectData      `json:"data"`
}

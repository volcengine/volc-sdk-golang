package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type BoundingBox struct {
	Xmin int     `json:"x_min"`
	Xmax int     `json:"x_max"`
	Ymin int     `json:"y_min"`
	Ymax int     `json:"y_max"`
	Prob float64 `json:"prob"`
}

type Object struct {
	BBox     BoundingBox `json:"boundingbox"`
	Category int         `json:"category"`
}

type Frame struct {
	Objects []Object `json:"objects"`
	Width   int      `json:"width"`
	Height  int      `json:"height"`
}

type GoodsDetectData struct {
	Frame *Frame `json:"Frame"`
}

type GoodsDetectResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *GoodsDetectData       `json:"data"`
}

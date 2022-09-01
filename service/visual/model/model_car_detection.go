package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type CarBox struct {
	MinX  int     `json:"min_x"`
	MinY  int     `json:"min_y"`
	MaxX  int     `json:"max_x"`
	MaxY  int     `json:"max_y"`
	Score float64 `json:"score"`
}

type CarDetectionData struct {
	CarBox []CarBox `json:"car_box"`
}

type CarDetectionResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *CarDetectionData      `json:"data"`
}

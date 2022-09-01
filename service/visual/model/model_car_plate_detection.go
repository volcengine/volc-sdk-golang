package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type CarPlateBox struct {
	MinX  int     `json:"min_x"`
	MinY  int     `json:"min_y"`
	MaxX  int     `json:"max_x"`
	MaxY  int     `json:"max_y"`
	Score float64 `json:"score"`
}

type CarPlateDetectionData struct {
	CarPlateBox []CarBox `json:"car_plate_box"`
}

type CarPlateDetectionResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *CarPlateDetectionData `json:"data"`
}

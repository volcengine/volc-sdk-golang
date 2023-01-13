package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type PointInfo struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type CharInfo struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Score  float64 `json:"score"`
	Char   string  `json:"char"`
}

type RectInfo struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type OCRNormalData struct {
	LineTexts []string      `json:"line_texts"`
	LineRects []*RectInfo   `json:"line_rects"`
	Chars     [][]*CharInfo `json:"chars"`
	Polygons  [][][]float64 `json:"polygons"`
}

type OCRNormalResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *OCRNormalData         `json:"data"`
}

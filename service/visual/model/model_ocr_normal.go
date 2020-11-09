package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type PointInfo struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type CharInfo struct {
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Score  float64 `json:"score"`
	Char   string  `json:"char"`
}

type RectInfo struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type OCRNormalData struct {
	LineTexts []string      `json:"line_texts"`
	LineRects []*RectInfo   `json:"line_rects"`
	Chars     [][]*CharInfo `json:"chars"`
	Polygons  [][][]int     `json:"polygons"`
}

type OCRNormalResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *OCRNormalData         `json:"data"`
}

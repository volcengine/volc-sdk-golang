package model

import "github.com/volcengine/volc-sdk-golang/base"

type ImageScoreV2Data struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64  []string `json:"binary_data_base64"`
	Brightness        float64  `json:"brightness"`
	Contrast          float64  `json:"contrast"`
	HasBlackEdge      int      `json:"has_black_edge"`
	IsNight           int      `json:"is_night"`
	IsPureBackground  int      `json:"is_pure_background"`
	IsSandwitchVideo  int      `json:"is_sandwitch_video"`
	OverExposure      float64  `json:"over_exposure"`
	SandwitchBottom   int      `json:"sandwitch_bottom"`
	SandwitchLeft     int      `json:"sandwitch_left"`
	SandwitchRight    int      `json:"sandwitch_right"`
	SandwitchTop      int      `json:"sandwitch_top"`
	ScoreAestheticsV2 float64  `json:"score_aesthetics_v2"`
	ScoreDetail       float64  `json:"score_detail"`
	ScoreFace         float64  `json:"score_face"`
	ScoreLuma         float64  `json:"score_luma"`
	ScoreTotal        float64  `json:"score_total"`
	UnderExposure     float64  `json:"under_exposure"`
}

type ImageScoreV2Result struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *ImageScoreV2Data      `json:"data"`
}

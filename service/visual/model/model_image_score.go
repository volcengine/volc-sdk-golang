package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type ImageScoreData struct {
	Score            float64 `json:"score"`
	SharpnessScore   float64 `json:"sharpness_score"`
	QualityScore     float64 `json:"quality_score"`
	MeaninglessScore float64 `json:"meaningless_score"`
	FaceScore        float64 `json:"face_score"`
}

type ImageScoreResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *ImageScoreData        `json:"data"`
}

package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type FrontInfo struct {
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Ethnicity   string `json:"ethnicity"`
	DateOfBirth string `json:"data_of_birth"`
	Domicile    string `json:"domicile"`
	IdNumber    string `json:"id_number"`
	FaceCorners []int  `json:"face_corners"`
}

type BackInfo struct {
	IssueAuthority string `json:"issue_authority"`
	ValidPeriod    string `json:"valid_period"`
}

type IDCardData struct {
	CardFront   FrontInfo `json:"card_front"`
	CardBack    BackInfo  `json:"card_back"`
	CardCorners []float64 `json:"card_corners"`
}

type IDCardResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *IDCardData            `json:"data"`
}

type IDCardFrontInfoV2 struct {
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	Ethnicity   string    `json:"ethnicity"`
	DataOfBirth string    `json:"data_of_birth"`
	Domicile    string    `json:"domicile"`
	IDNumber    string    `json:"id_number"`
	FaceCorners []float64 `json:"face_corners"`
	CardCorners []float64 `json:"card_corners"`
	CardImg     string    `json:"card_img"`
	FaceImg     string    `json:"face_img"`
	CardType    string    `json:"card_type"`
}

type IDCardBackInfoV2 struct {
	IssueAuthority string    `json:"issue_authority"`
	ValidPeriod    string    `json:"valid_period"`
	CardCorners    []float64 `json:"card_corners"`
	CardImg        string    `json:"card_img"`
	CardType       string    `json:"card_type"`
}

type IDCardDataV2 struct {
	CardFront IDCardFrontInfoV2 `json:"card_front"`
	CardBack  IDCardBackInfoV2  `json:"card_back"`
}

type IDCardResultV2 struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *IDCardDataV2          `json:"data"`
}

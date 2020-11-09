package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type FrontInfo struct {
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Ethnicity   string `json:"ethnicity"`
	DateOfBirth string `json:"date_of_birth"`
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
	CardCorners []int     `json:"card_corners"`
}

type IDCardResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *IDCardData            `json:"data"`
}

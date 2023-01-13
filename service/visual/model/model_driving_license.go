package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type DrivingLicenseMainInfo struct {
	IdNumber         string `json:"id_number"`
	Name             string `json:"name"`
	Sex              string `json:"sex"`
	Nationality      string `json:"nationality"`
	Address          string `json:"address"`
	DateOfBirth      string `json:"date_of_birth"`
	DateOfFirstIssue string `json:"date_of_first_issue"`
	Class            string `json:"class"`
	ValidBegin       string `json:"valid_begin"`
	ValidEnd         string `json:"valid_end"`
}

type DrivingLicenseData struct {
	MainInfo *DrivingLicenseMainInfo `json:"license_main"`
}

type DrivingLicenseResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *DrivingLicenseData    `json:"data"`
}

package vms

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type RiskControlReq struct {
	CustomerNumberList string
	BusinessLineId     string
	CallType           int32
}

type RiskControlResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []NumberStatus
}

type NumberStatus struct {
	CustomerNumber string
	BusinessLineId int32
	Status         int32
	StartTime      string
	ExpiryTime     string
	Reason         string
	HighRisk       int32
	TreatmentId    int32
	TreatmentName  string
}

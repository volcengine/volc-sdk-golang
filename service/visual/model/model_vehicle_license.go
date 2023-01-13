package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type VehicleLicenseMainInfo struct {
	PlateNumber  string `json:"plate_number"`
	VehicleType  string `json:"vehicle_type"`
	Owner        string `json:"owner"`
	Address      string `json:"address"`
	UseCharacter string `json:"use_character"`
	Model        string `json:"model"`
	VIN          string `json:"vin"`
	EngineNumber string `json:"engine_number"`
	RegisterDate string `json:"register_date"`
	IssueDate    string `json:"issue_date"`
}

type VehicleLicenseData struct {
	MainInfo *VehicleLicenseMainInfo `json:"license_main"`
}

type VehicleLicenseResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *VehicleLicenseData    `json:"data"`
}

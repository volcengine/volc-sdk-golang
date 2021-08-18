package rtc

import (
	"encoding/json"

	"github.com/volcengine/volc-sdk-golang/base"
)

// CommonResponse ... need to decode result by type
type CommonResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           json.RawMessage `json:"Result,omitempty"`
}

// ListRoomsResponse ...
type ListRoomsResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ListRoomsResult `json:"Result,omitempty"`
}

// ListRoomsResult ...
type ListRoomsResult struct {
	Total       int    `json:"Total"`
	ActiveNum   int    `json:"ActiveNum"`
	InactiveNum int    `json:"InactiveNum"`
	Offset      int    `json:"Offset"`
	Limit       int    `json:"Limit"`
	Rooms       []Room `json:"Rooms"`
}

// Room ...
type Room struct {
	RoomId    string `json:"RoomId"`
	AppId     string `json:"AppId"`
	UserNum   int    `json:"UserNum"`
	StreamNum int    `json:"StreamNum"`
	State     int    `json:"State"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

// ListIndicatorsRequest...
type ListIndicatorsRequest struct {
	AppId     string `json:"AppId"`
	StartTime string `json:"StartTime"`
	EndTime   string `json:"EndTime"`
	Indicator string `json:"Indicator"`
	OS        string `json:"OS,omitempty"`
	Network   string `json:"Network,omitempty"`
}

// ListIndicatorsResponse ...
type ListIndicatorsResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ListIndicatorsResult `json:"Result,omitempty"`
}

type ListIndicatorsResult struct {
	Indicators []Indicator `json:"Indicators"`
}

// Indicator ...
type Indicator struct {
	Name  string `json:"Name"`
	Unit  string `json:"Unit"`
	Datas []Data `json:"Data"`
}

//Data ...
type Data struct {
	TimeStamp string  `json:"TimeStamp"`
	Value     float64 `json:"Value"`
}

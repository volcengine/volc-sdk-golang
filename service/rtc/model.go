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

type (
	// ListRoomInformationResponse ...
	ListRoomInformationResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *ListRoomInformationResult `json:"Result,omitempty"`
	}

	// ListRoomInformationResult ...
	ListRoomInformationResult struct {
		Total    int    `json:"Total"`
		PageNum  int    `json:"PageNum"`
		PageSize int    `json:"PageSize"`
		HasMore  bool   `json:"HasMore"`
		RoomList []Room `json:"RoomList"`
	}

	// Room ...
	Room struct {
		RoomId      string `json:"RoomId"`
		CreatedTime string `json:"CreatedTime"`
		DestroyTime string `json:"DestroyTime"`
		IsFinished  bool   `json:"IsFinished"`
	}
)

type (
	// ListIndicatorsRequest...
	ListIndicatorsRequest struct {
		AppId     string `json:"AppId"`
		StartTime string `json:"StartTime"`
		EndTime   string `json:"EndTime"`
		Indicator string `json:"Indicator"`
		OS        string `json:"OS,omitempty"`
		Network   string `json:"Network,omitempty"`
	}

	// ListIndicatorsResponse ...
	ListIndicatorsResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *ListIndicatorsResult `json:"Result,omitempty"`
	}

	ListIndicatorsResult struct {
		Indicators []Indicator `json:"Indicators"`
	}

	// Indicator ...
	Indicator struct {
		Name  string `json:"Name"`
		Unit  string `json:"Unit"`
		Datas []Data `json:"Data"`
	}

	// Data ...
	Data struct {
		TimeStamp string  `json:"TimeStamp"`
		Value     float64 `json:"Value"`
	}
)

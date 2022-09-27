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
	// StartRecordRequest ...
	StartRecordRequest struct {
		AppId            string            `json:"AppId"`
		BusinessId       string            `json:"BusinessId"`
		RoomId           string            `json:"RoomId"`
		TaskId           string            `json:"TaskId"`
		RecordMode       uint32            `json:"RecordMode"`
		Encode           *Encode           `json:"Encode,omitempty"`
		FileFormatConfig *FileFormatConfig `json:"FileFormatConfig,omitempty"`
		StorageConfig    StorageConfig     `json:"StorageConfig"`
	}

	// StartRecordResponse ...
	StartRecordResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           string `json:"Result,omitempty"`
	}

	Encode struct {
		VideoWidth   uint32 `json:"VideoWidth"`
		VideoHeight  uint32 `json:"VideoHeight"`
		VideoFps     uint32 `json:"VideoFps"`
		VideoBitrate uint32 `json:"VideoBitrate"`
	}

	FileFormatConfig struct {
		FileFormat []string
	}

	StorageConfig struct {
		Type         uint32        `json:"Type"`
		TosConfig    *TosConfig    `json:"TosConfig,omitempty"`
		CustomConfig *CustomConfig `json:"CustomConfig,omitempty"`
	}

	TosConfig struct {
		AccountId string `json:"AccountId"`
		Region    uint32 `json:"Region"`
		Bucket    string `json:"Bucket"`
	}

	CustomConfig struct {
		Vendor    uint32 `json:"Vendor"`
		Region    string `json:"Region"`
		Bucket    string `json:"Bucket"`
		AccessKey string `json:"AccessKey"`
		SecretKey string `json:"SecretKey"`
	}
)

type (
	// GetRecordTaskResponse ...
	GetRecordTaskResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *GetRecordTaskResult `json:"Result,omitempty"`
	}

	GetRecordTaskResult struct {
		RecordTask RecordTask `json:"RecordTask"`
	}

	RecordTask struct {
		StartTime      uint64       `json:"StartTime"`
		EndTime        uint64       `json:"EndTime"`
		Status         uint64       `json:"Status"`
		StopReason     string       `json:"StopReason"`
		RecordFileList []RecordFile `json:"RecordFileList"`
	}

	RecordFile struct {
		Vid         string   `json:"Vid"`
		ObjectKey   string   `json:"ObjectKey"`
		Duration    uint64   `json:"Duration"`
		Size        uint64   `json:"Size"`
		StartTime   uint64   `json:"StartTime"`
		StreamList  []Stream `json:"StreamList"`
		VideoCodec  string   `json:"VideoCodec"`
		AudioCodec  string   `json:"AudioCodec"`
		VideoWidth  int      `json:"VideoWidth"`
		VideoHeight int      `json:"VideoHeight"`
	}

	Stream struct {
		Index      uint32 `json:"Index"`
		StreamType uint32 `json:"StreamType"`
		UserId     string `json:"UserId"`
	}
)

type (
	StartWebRecordRequest struct {
		AppId           string
		TaskId          string
		InputURL        string
		MaxRecordSecond int
		JsCommand       []string
		Bucket          string
		VideoSpace      string
		VideoInfo       WebVideoInfo
		PageInfo        WebPageInfo
		Duration        int
	}

	WebVideoInfo struct {
		ResolutionWidth  int
		ResolutionHeight int
	}

	WebPageInfo struct {
		PageWidth  int
		PageHeight int
	}

	StartWebRecordResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *StartWebRecordResult `json:"Result,omitempty"`
	}

	StartWebRecordResult struct {
		Message string `json:"Message"`
	}

	StopWebRecordRequest struct {
		AppId  string
		TaskId string
	}

	StopWebRecordResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *StopWebRecordResult `json:"Result,omitempty"`
	}

	StopWebRecordResult struct {
		Message string `json:"Message"`
	}

	GetWebRecordTaskResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *GetWebRecordTaskResult `json:"Result,omitempty"`
	}

	GetWebRecordTaskResult struct {
		Message   string        `json:"Message"`
		EventData TaskEventData `json:"EventData"`
	}

	TaskEventData struct {
		TaskId     string `json:"TaskId"`
		Status     int    `json:"Status"`
		CreateTime int    `json:"CreateTime"`
		FinishTime int    `json:"FinishTime"`
		Vid        string `json:"Vid"`
		FinalFile  File   `json:"FinalFile"`
		Files      []File `json:"Files"`
	}

	File struct {
		Index     int
		Bucket    string
		ObjectKey string
	}

	GetWebRecordListResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *GetWebRecordListResult `json:"Result,omitempty"`
	}

	GetWebRecordListResult struct {
		Message       string        `json:"Message"`
		WebRecordList WebRecordList `json:"WebRecordList"`
	}

	WebRecordList struct {
		AppId      string  `json:"AppId"`
		Tasks      []Tasks `json:"Tasks"`
		PageNumber int     `json:"PageNumber"`
		PageSize   int     `json:"PageSize"`
		TotalCount int     `json:"TotalCount"`
	}

	Tasks struct {
		TaskId          string `json:"TaskId"`
		CreateTime      int    `json:"CreateTime"`
		FinishTime      int    `json:"FinishTime"`
		Status          int    `json:"Status"`
		InputURL        string `json:"InputURL"`
		VideoSpace      string `json:"VideoSpace"`
		Vid             string `json:"Vid"`
		MaxRecordSecond int    `json:"MaxRecordSecond"`
		Duration        int    `json:"Duration"`
		Bucket          string `json:"Bucket"`
		ObjectKey       string `json:"ObjectKey"`
	}
)

package tls

import "errors"

type LogBackFlowTaskTopicSource struct {
	ProjectID string `json:"ProjectId"`
	TopicID   string `json:"TopicID"`
}

type LogBackFlowTaskSource struct {
	SourceType                 string                      `json:"SourceType"`
	LogBackFlowTaskTopicSource *LogBackFlowTaskTopicSource `json:"LogBackFlowTaskTopicSource,omitempty"`
}

type LogBackFlowScheduleSqlTaskInfo struct {
	DestTopicID       string        `json:"DestTopicID"`
	DestRegion        string        `json:"DestRegion,omitempty"`
	RequestCycle      *RequestCycle `json:"RequestCycle,omitempty"`
	ProcessTimeWindow string        `json:"ProcessTimeWindow,omitempty"`
	ProcessSqlDelay   int64         `json:"ProcessSqlDelay,omitempty"`
	MaxRetryTimes     int           `json:"MaxRetryTimes,omitempty"`
	MaxTimeout        int64         `json:"MaxTimeout,omitempty"`
}

type LogBackFlowShipperToTosInfo struct {
	TosShipperInfo *TosShipperInfo `json:"TosShipperInfo,omitempty"`
	ContentInfo    *ContentInfo    `json:"ContentInfo,omitempty"`
}

type LogBackFlowQueryField struct {
	Alias  string `json:"Alias,omitempty"`
	Column string `json:"Column,omitempty"`
}

type LogBackFlowQueryFilter struct {
	Field    string      `json:"Field,omitempty"`
	Value    interface{} `json:"Value,omitempty"`
	Values   []string    `json:"Values,omitempty"`
	Operator string      `json:"Operator,omitempty"`
}

type LogBackFlowQueryParams struct {
	Asc     *bool                    `json:"Asc,omitempty"`
	Limit   *int                     `json:"Limit,omitempty"`
	Order   string                   `json:"Order,omitempty"`
	Fields  []LogBackFlowQueryField  `json:"Fields,omitempty"`
	Filters []LogBackFlowQueryFilter `json:"Filters,omitempty"`
}

type CreateLogBackFlowTaskRequest struct {
	CommonRequest
	BackFlowEndTime       *int64                          `json:"BackFlowEndTime,omitempty"`
	BackFlowStartTime     int64                           `json:"BackFlowStartTime,omitempty"`
	Description           *string                         `json:"Description,omitempty"`
	IamProjectName        *string                         `json:"IamProjectName,omitempty"`
	LogBackFlowTaskSource *LogBackFlowTaskSource          `json:"LogBackFlowTaskSource"`
	QueryParams           *LogBackFlowQueryParams         `json:"QueryParams"`
	ScheduleSqlTaskInfo   *LogBackFlowScheduleSqlTaskInfo `json:"ScheduleSqlTaskInfo,omitempty"`
	ShipperToTosInfo      *LogBackFlowShipperToTosInfo    `json:"ShipperToTosInfo,omitempty"`
	TaskName              string                          `json:"TaskName"`
}

func (v *CreateLogBackFlowTaskRequest) CheckValidation() error {
	if v.TaskName == "" {
		return errors.New("Invalid argument, empty TaskName")
	}
	if v.LogBackFlowTaskSource == nil {
		return errors.New("Invalid argument, empty LogBackFlowTaskSource")
	}
	if v.QueryParams == nil {
		return errors.New("Invalid argument, empty QueryParams")
	}
	return nil
}

type CreateLogBackFlowTaskResponse struct {
	CommonResponse
	TaskID string `json:"TaskId"`
}

type DeleteLogBackFlowTaskRequest struct {
	CommonRequest
	TaskID string `json:"TaskId"`
}

func (v *DeleteLogBackFlowTaskRequest) CheckValidation() error {
	if v.TaskID == "" {
		return errors.New("Invalid argument, empty TaskID")
	}
	return nil
}

type DeleteLogBackFlowTaskResponse struct {
	CommonResponse
}

type DescribeLogBackFlowTasksRequest struct {
	CommonRequest
	PageNumber        *int     `json:"PageNumber,omitempty"`
	PageSize          *int     `json:"PageSize,omitempty"`
	TopicIDList       []string `json:"TopicIDList,omitempty"`
	TaskID            *string  `json:"TaskId,omitempty"`
	TaskName          *string  `json:"TaskName,omitempty"`
	Status            *int     `json:"Status,omitempty"`
	ScheduleSQLTaskID *string  `json:"ScheduleSQLTaskId,omitempty"`
	ShipperID         *string  `json:"ShipperId,omitempty"`
}

func (v *DescribeLogBackFlowTasksRequest) CheckValidation() error {
	return nil
}

type LogBackFlowRelaTasksInfo struct {
	ScheduleSQLTaskID   string `json:"ScheduleSQLTaskId,omitempty"`
	ScheduleSQLTaskName string `json:"ScheduleSQLTaskName,omitempty"`
	DestRegion          string `json:"DestRegion,omitempty"`
	ShipperID           string `json:"ShipperID,omitempty"`
	ShipperName         string `json:"ShipperName,omitempty"`
}

type LogBackFlowTaskInfo struct {
	TaskID                string                          `json:"TaskId"`
	TaskName              string                          `json:"TaskName"`
	Status                int                             `json:"Status"`
	LogBackFlowTaskSource *LogBackFlowTaskSource          `json:"LogBackFlowTaskSource,omitempty"`
	ScheduleSqlTaskInfo   *LogBackFlowScheduleSqlTaskInfo `json:"ScheduleSqlTaskInfo,omitempty"`
	QueryParams           *LogBackFlowQueryParams         `json:"QueryParams,omitempty"`
	ShipperToTosInfo      *LogBackFlowShipperToTosInfo    `json:"ShipperToTosInfo,omitempty"`
	Description           *string                         `json:"Description,omitempty"`
	RelaTasksInfo         *LogBackFlowRelaTasksInfo       `json:"RelaTasksInfo,omitempty"`
	BackFlowStartTime     int64                           `json:"BackFlowStartTime,omitempty"`
	BackFlowEndTime       *int64                          `json:"BackFlowEndTime,omitempty"`
	CreateTime            int64                           `json:"CreateTime,omitempty"`
	ModifyTime            int64                           `json:"ModifyTime,omitempty"`
	IamProjectName        string                          `json:"IamProjectName,omitempty"`
}

type DescribeLogBackFlowTasksResponse struct {
	CommonResponse
	LogBackFlowTasks []*LogBackFlowTaskInfo `json:"LogBackFlowTasks"`
	Total            int64                  `json:"Total"`
}

type ModifyLogBackFlowTaskRequest struct {
	CommonRequest
	QueryParams         *LogBackFlowQueryParams         `json:"QueryParams,omitempty"`
	ScheduleSqlTaskInfo *LogBackFlowScheduleSqlTaskInfo `json:"ScheduleSqlTaskInfo,omitempty"`
	ShipperToTosInfo    *LogBackFlowShipperToTosInfo    `json:"ShipperToTosInfo,omitempty"`
	TaskID              string                          `json:"TaskId"`
}

func (v *ModifyLogBackFlowTaskRequest) CheckValidation() error {
	if v.TaskID == "" {
		return errors.New("Invalid argument, empty TaskID")
	}
	return nil
}

type ModifyLogBackFlowTaskResponse struct {
	CommonResponse
}

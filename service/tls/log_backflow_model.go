package tls

import (
	"encoding/json"
	"errors"
)

const (
	LogBackFlowTaskStatusDone         = "DONE"
	LogBackFlowTaskStatusCreating     = "CREATING"
	LogBackFlowTaskStatusFinished     = "FINISHED"
	LogBackFlowTaskStatusDeleting     = "DELETING"
	LogBackFlowTaskStatusCreateFailed = "CREATEFAILED"
	LogBackFlowTaskStatusRunFailed    = "RUNFAILED"
)

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

type EvaluationSetFieldSchema struct {
	Name                 string `json:"Name"`
	Description          string `json:"Description"`
	ContentType          string `json:"ContentType"`
	DefaultDisplayFormat int32  `json:"DefaultDisplayFormat"`
	IsRequired           bool   `json:"IsRequired"`
	TextSchema           string `json:"TextSchema"`
	Key                  string `json:"Key"`
}

type EvaluationSetSchema struct {
	FieldSchemas []*EvaluationSetFieldSchema `json:"FieldSchemas"`
}

type EvaluationSetFieldMapping struct {
	Source string `json:"Source"`
	Target string `json:"Target"`
}

type EvaluationSetShipperInfo struct {
	WorkspaceID              string                       `json:"WorkspaceId"`
	ProjectName              string                       `json:"ProjectName,omitempty"`
	EvaluationSetID          string                       `json:"EvaluationSetId,omitempty"`
	EvaluationSetName        string                       `json:"EvaluationSetName,omitempty"`
	EvaluationSetDescription string                       `json:"EvaluationSetDescription,omitempty"`
	BizCategory              string                       `json:"BizCategory,omitempty"`
	EvaluationSetSchema      *EvaluationSetSchema         `json:"EvaluationSetSchema,omitempty"`
	FieldMappings            []*EvaluationSetFieldMapping `json:"FieldMappings"`
	ItemKeyField             string                       `json:"ItemKeyField,omitempty"`
	BatchSize                int32                        `json:"BatchSize,omitempty"`
	SkipInvalidItems         bool                         `json:"SkipInvalidItems,omitempty"`
	AllowPartialAdd          bool                         `json:"AllowPartialAdd,omitempty"`
}

type LogBackFlowShipperToAgentLoopInfo struct {
	EvaluationSetShipperInfo *EvaluationSetShipperInfo `json:"EvaluationSetShipperInfo"`
	ContentInfo              *ContentInfo              `json:"ContentInfo"`
}

type LogBackFlowETLTaskInfo struct {
	Script          string           `json:"Script"`
	TargetResources []TargetResource `json:"TargetResources"`
}

type LogBackFlowQueryField struct {
	Alias  string `json:"Alias,omitempty"`
	Column string `json:"Column,omitempty"`
}

type LogBackFlowQueryFilter struct {
	Field     string        `json:"Field,omitempty"`
	Value     interface{}   `json:"Value,omitempty"`
	Values    []string      `json:"-"`
	ValuesAny []interface{} `json:"-"`
	Operator  string        `json:"Operator,omitempty"`
}

func (v LogBackFlowQueryFilter) MarshalJSON() ([]byte, error) {
	wire := map[string]interface{}{}
	if v.Field != "" {
		wire["Field"] = v.Field
	}
	if v.Value != nil {
		wire["Value"] = v.Value
	}
	if len(v.ValuesAny) > 0 {
		wire["Values"] = v.ValuesAny
	} else if len(v.Values) > 0 {
		wire["Values"] = v.Values
	}
	if v.Operator != "" {
		wire["Operator"] = v.Operator
	}
	return json.Marshal(wire)
}

func (v *LogBackFlowQueryFilter) UnmarshalJSON(data []byte) error {
	var wire struct {
		Field    string        `json:"Field"`
		Value    interface{}   `json:"Value"`
		Values   []interface{} `json:"Values"`
		Operator string        `json:"Operator"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	v.Field = wire.Field
	v.Value = wire.Value
	v.ValuesAny = wire.Values
	v.Operator = wire.Operator
	v.Values = nil
	for _, value := range wire.Values {
		stringValue, ok := value.(string)
		if !ok {
			v.Values = nil
			return nil
		}
		v.Values = append(v.Values, stringValue)
	}
	return nil
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
	BackFlowEndTime        *int64                             `json:"BackFlowEndTime,omitempty"`
	BackFlowStartTime      int64                              `json:"BackFlowStartTime,omitempty"`
	Description            *string                            `json:"Description,omitempty"`
	IamProjectName         *string                            `json:"IamProjectName,omitempty"`
	LogBackFlowTaskSource  *LogBackFlowTaskSource             `json:"LogBackFlowTaskSource"`
	ETLTaskInfo            *LogBackFlowETLTaskInfo            `json:"ETLTaskInfo"`
	QueryParams            *LogBackFlowQueryParams            `json:"QueryParams,omitempty"`
	ShipperToTosInfo       *LogBackFlowShipperToTosInfo       `json:"ShipperToTosInfo,omitempty"`
	ShipperToAgentLoopInfo *LogBackFlowShipperToAgentLoopInfo `json:"ShipperToAgentLoopInfo,omitempty"`
	TaskName               string                             `json:"TaskName"`
	// Deprecated: current log-service no longer accepts ScheduleSqlTaskInfo in create requests.
	ScheduleSqlTaskInfo *LogBackFlowScheduleSqlTaskInfo `json:"-"`
}

func (v *CreateLogBackFlowTaskRequest) CheckValidation() error {
	if v.TaskName == "" {
		return errors.New("Invalid argument, empty TaskName")
	}
	if v.LogBackFlowTaskSource == nil {
		return errors.New("Invalid argument, empty LogBackFlowTaskSource")
	}
	if v.LogBackFlowTaskSource.SourceType != "Topic" || v.LogBackFlowTaskSource.LogBackFlowTaskTopicSource == nil ||
		v.LogBackFlowTaskSource.LogBackFlowTaskTopicSource.ProjectID == "" || v.LogBackFlowTaskSource.LogBackFlowTaskTopicSource.TopicID == "" {
		return errors.New("Invalid argument, invalid LogBackFlowTaskSource")
	}
	if v.ETLTaskInfo == nil {
		return errors.New("Invalid argument, empty ETLTaskInfo")
	}
	if !isValidLogBackFlowETLTaskInfo(v.ETLTaskInfo) {
		return errors.New("Invalid argument, invalid ETLTaskInfo")
	}
	if v.QueryParams != nil && len(v.QueryParams.Fields) == 0 {
		return errors.New("Invalid argument, empty QueryParams.Fields")
	}
	if v.BackFlowStartTime <= 0 {
		return errors.New("Invalid argument, invalid BackFlowStartTime")
	}
	if v.ShipperToTosInfo != nil && v.ShipperToAgentLoopInfo != nil {
		return errors.New("Invalid argument, ShipperToTosInfo and ShipperToAgentLoopInfo are mutually exclusive")
	}
	if v.ScheduleSqlTaskInfo != nil {
		return errors.New("Invalid argument, ScheduleSqlTaskInfo is no longer supported")
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
	PageNumber  *int     `json:"PageNumber,omitempty"`
	PageSize    *int     `json:"PageSize,omitempty"`
	TopicIDList []string `json:"TopicIDList,omitempty"`
	TaskID      *string  `json:"TaskId,omitempty"`
	TaskName    *string  `json:"TaskName,omitempty"`
	Status      *string  `json:"Status,omitempty"`
	ETLTaskID   *string  `json:"ETLTaskId,omitempty"`
	ShipperID   *string  `json:"ShipperId,omitempty"`
	// Deprecated: current log-service no longer supports ScheduleSQLTaskId as a query filter.
	ScheduleSQLTaskID *string `json:"-"`
}

func (v *DescribeLogBackFlowTasksRequest) CheckValidation() error {
	if v.ScheduleSQLTaskID != nil {
		return errors.New("Invalid argument, ScheduleSQLTaskID is no longer supported; use ETLTaskID")
	}
	if v.Status != nil && !isValidLogBackFlowTaskStatus(*v.Status) {
		return errors.New("Invalid argument, invalid Status")
	}
	return nil
}

func isValidLogBackFlowTaskStatus(status string) bool {
	switch status {
	case LogBackFlowTaskStatusDone, LogBackFlowTaskStatusCreating, LogBackFlowTaskStatusFinished,
		LogBackFlowTaskStatusDeleting, LogBackFlowTaskStatusCreateFailed, LogBackFlowTaskStatusRunFailed:
		return true
	default:
		return false
	}
}

func isValidLogBackFlowETLTaskInfo(info *LogBackFlowETLTaskInfo) bool {
	return info != nil && info.Script != "" && len(info.TargetResources) == 1 &&
		info.TargetResources[0].TopicID != "" && info.TargetResources[0].Alias != ""
}

type LogBackFlowRelaTasksInfo struct {
	ETLTaskID           string `json:"ETLTaskId,omitempty"`
	ETLTaskName         string `json:"ETLTaskName,omitempty"`
	ScheduleSQLTaskID   string `json:"ScheduleSQLTaskId,omitempty"`
	ScheduleSQLTaskName string `json:"ScheduleSQLTaskName,omitempty"`
	DestRegion          string `json:"DestRegion,omitempty"`
	ShipperID           string `json:"ShipperID,omitempty"`
	ShipperName         string `json:"ShipperName,omitempty"`
}

type LogBackFlowTaskInfo struct {
	TaskID                 string                             `json:"TaskId"`
	TaskName               string                             `json:"TaskName"`
	Status                 int                                `json:"Status"`
	LogBackFlowTaskSource  *LogBackFlowTaskSource             `json:"LogBackFlowTaskSource,omitempty"`
	ETLTaskInfo            *LogBackFlowETLTaskInfo            `json:"ETLTaskInfo,omitempty"`
	ScheduleSqlTaskInfo    *LogBackFlowScheduleSqlTaskInfo    `json:"ScheduleSqlTaskInfo,omitempty"`
	QueryParams            *LogBackFlowQueryParams            `json:"QueryParams,omitempty"`
	ShipperToTosInfo       *LogBackFlowShipperToTosInfo       `json:"ShipperToTosInfo,omitempty"`
	ShipperToAgentLoopInfo *LogBackFlowShipperToAgentLoopInfo `json:"ShipperToAgentLoopInfo,omitempty"`
	Description            *string                            `json:"Description,omitempty"`
	RelaTasksInfo          *LogBackFlowRelaTasksInfo          `json:"RelaTasksInfo,omitempty"`
	BackFlowStartTime      int64                              `json:"BackFlowStartTime,omitempty"`
	BackFlowEndTime        *int64                             `json:"BackFlowEndTime,omitempty"`
	CreateTime             int64                              `json:"CreateTime,omitempty"`
	ModifyTime             int64                              `json:"ModifyTime,omitempty"`
	IamProjectName         string                             `json:"IamProjectName,omitempty"`
}

type DescribeLogBackFlowTasksResponse struct {
	CommonResponse
	LogBackFlowTasks []*LogBackFlowTaskInfo `json:"LogBackFlowTasks"`
	Total            int64                  `json:"Total"`
}

type ModifyLogBackFlowTaskRequest struct {
	CommonRequest
	ETLTaskInfo            *LogBackFlowETLTaskInfo            `json:"ETLTaskInfo,omitempty"`
	QueryParams            *LogBackFlowQueryParams            `json:"QueryParams,omitempty"`
	ShipperToTosInfo       *LogBackFlowShipperToTosInfo       `json:"ShipperToTosInfo,omitempty"`
	ShipperToAgentLoopInfo *LogBackFlowShipperToAgentLoopInfo `json:"ShipperToAgentLoopInfo,omitempty"`
	TaskID                 string                             `json:"TaskId"`
	// Deprecated: current log-service no longer accepts ScheduleSqlTaskInfo in modify requests.
	ScheduleSqlTaskInfo *LogBackFlowScheduleSqlTaskInfo `json:"-"`
}

func (v *ModifyLogBackFlowTaskRequest) CheckValidation() error {
	if v.TaskID == "" {
		return errors.New("Invalid argument, empty TaskID")
	}
	if v.QueryParams != nil && v.ETLTaskInfo == nil {
		return errors.New("Invalid argument, QueryParams requires ETLTaskInfo")
	}
	if v.ETLTaskInfo != nil && !isValidLogBackFlowETLTaskInfo(v.ETLTaskInfo) {
		return errors.New("Invalid argument, invalid ETLTaskInfo")
	}
	if v.QueryParams != nil && len(v.QueryParams.Fields) == 0 {
		return errors.New("Invalid argument, empty QueryParams.Fields")
	}
	if v.ShipperToTosInfo != nil && v.ShipperToAgentLoopInfo != nil {
		return errors.New("Invalid argument, ShipperToTosInfo and ShipperToAgentLoopInfo are mutually exclusive")
	}
	if v.ScheduleSqlTaskInfo != nil {
		return errors.New("Invalid argument, ScheduleSqlTaskInfo is no longer supported")
	}
	return nil
}

type ModifyLogBackFlowTaskResponse struct {
	CommonResponse
}

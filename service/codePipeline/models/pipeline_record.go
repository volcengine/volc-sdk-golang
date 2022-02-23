package models

type ListPipelineRecordsRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	PipelineId  string `json:"PipelineId"`
	Page
	Desc   bool                  `json:"Desc,omitempty"`
	Filter *PipelineRecordFilter `json:"Filter,omitempty"`
}

type ListPipelineRecordsResponse struct {
	Total int64            `json:"Total"`
	Items []PipelineRecord `json:"Items"`
}

type GetPipelineRecordRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	PipelineId  string `json:"PipelineId"`
	Id          string `json:"Id"`
}

type GetPipelineRecordResponse struct {
	PipelineRecord
}

type StopPipelineRecordRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	PipelineId  string `json:"PipelineId"`
	Id          string `json:"Id"`
}

type StopPipelineRecordResponse struct {
}

type RetryPipelineRecordRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	PipelineId  string `json:"PipelineId"`
	Id          string `json:"Id"`
}

type RetryPipelineRecordResponse struct {
}

type DeletePipelineRecordRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	PipelineId  string `json:"PipelineId"`
	Id          string `json:"Id"`
}

type DeletePipelineRecordResponse struct {
}

type PipelineRecordFilter struct {
	Statuses string `json:"Name,omitempty"`
}

type PipelineRecord struct {
	Id          string                `json:"Id"`
	Status      string                `json:"Status"`
	Creator     string                `json:"Creator"`
	StartTime   string                `json:"StartTime"`
	EndTime     string                `json:"EndTime,omitempty"`
	TriggerMode string                `json:"TriggerMode"`
	DynamicEnvs []*KVPair             `json:"DynamicEnvs,omitempty"`
	Description string                `json:"Description,omitempty"`
	WebhookUrl  string                `json:"WebhookUrl,omitempty"`
	Stages      []PipelineRecordStage `json:"Stages"`
	LogStatus   bool                  `json:"LogStatus"`
	ClusterPool string                `json:"ClusterPool"`
	ClusterId   string                `json:"ClusterId"`
}

type PipelineRecordStage struct {
	Name   string               `json:"Name"`
	Id     string               `json:"Id"`
	Status string               `json:"Status"`
	Infos  []KVPair             `json:"Infos,omitempty"`
	Tasks  []PipelineRecordTask `json:"Tasks"`
}

type PipelineRecordTask struct {
	Name             string               `json:"Name"`
	Id               string               `json:"Id"`
	Status           string               `json:"Status"`
	StartTime        string               `json:"StartTime,omitempty"`
	EndTime          string               `json:"EndTime,omitempty"`
	Type             string               `json:"Type"`
	Steps            []PipelineRecordStep `json:"Steps"`
	TotalRetryNumber int8                 `json:"TotalRetryNumber"`
}

type PipelineRecordStep struct {
	Name                string               `json:"Name"`
	Id                  string               `json:"Id"`
	Status              string               `json:"Status"`
	StartTime           string               `json:"StartTime,omitempty"`
	EndTime             string               `json:"EndTime,omitempty"`
	Type                string               `json:"Type"`
	ApprovalStatus      *ApprovalStatus      `json:"ApprovalStatus,omitempty"`
	RollingUpdateStatus *RollingUpdateStatus `json:"RollingUpdateStatus,omitempty"`
	RollbackTrigger     *RollbackTrigger     `json:"RollbackTrigger,omitempty"`
	UpdateImageStatus   string               `json:"UpdateImageStatus,omitempty"`
	Result              []KVPair             `json:"Result"`
	Params              []KVPair             `json:"Params"`
	CustomParams        []*KVPair            `json:"CustomParams,omitempty"`
	TotalRetryNumber    int8                 `json:"TotalRetryNumber"`
}

type ApprovalStatus struct {
	Type      string      `json:"Type"`
	Approvers []*Approver `json:"Approvers,omitempty"`
	Status    string      `json:"Status"`
}

type RollingUpdateStatus struct {
	RollingUpdateConfig RollingUpdateConfig `json:"RollingUpdateConfig"`
	BatchStatus         []string            `json:"BatchStatus"`
	Status              string              `json:"Status"`
}

type RollingUpdateConfig struct {
	Cluster       string `json:"Cluster"`
	Namespace     string `json:"Namespace"`
	Kind          string `json:"Kind"`
	Name          string `json:"Name"`
	ContainerName string `json:"ContainerName"`
	UpdateImage   string `json:"UpdateImage"`
	Batch         int8   `json:"Batch"`
	BatchPercent  []int8 `json:"BatchPercent"`
	Timeout       string `json:"Timeout"`
}

type RollbackTrigger struct {
	AccountId int64  `json:"AccountId,omitempty"`
	UserId    int64  `json:"UserId,omitempty"`
	Name      string `json:"Name"`
}

type Approver struct {
	AccountId   int64  `json:"AccountId,omitempty"`
	UserId      int64  `json:"UserId,omitempty"`
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
	Status      string `json:"Status"`
}

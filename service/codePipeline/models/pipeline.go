package models

type CreatePipelineRequest struct {
	WorkspaceId string   `json:"WorkspaceId"`
	Item        Pipeline `json:"Item"`
	ClientToken string   `json:"ClientToken,omitempty"`
}

type CreatePipelineResponse struct {
	Id string `json:"Id"`
}

type ListPipelinesRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	Page
	Filter  *PipelineFilter `json:"Filter,omitempty"`
	Desc    bool            `json:"Desc,omitempty"`
	OrderBy string          `json:"OrderBy,omitempty"`
}

type ListPipelinesResponse struct {
	Total int64      `json:"Total"`
	Items []Pipeline `json:"Items"`
}

type GetPipelineRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	Id          string `json:"Id"`
}

type GetPipelineResponse struct {
	Pipeline
}

type UpdatePipelineRequest struct {
	WorkspaceId string   `json:"WorkspaceId"`
	Id          string   `json:"Id"`
	Item        Pipeline `json:"Item"`
}

type UpdatePipelineResponse struct {
}

type DeletePipelineRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	Id          string `json:"Id"`
}

type DeletePipelineResponse struct {
}

type UpdatePipelinePropertiesRequest struct {
	WorkspaceId string `json:"WorkspaceId"`
	Id          string `json:"Id"`
	Name        string `json:"Name"`
}

type UpdatePipelinePropertiesResponse struct {
}

type RunPipelineRequest struct {
	WorkspaceId string    `json:"WorkspaceId"`
	Id          string    `json:"Id"`
	Branch      string    `json:"Branch,omitempty"`
	DynamicEnvs []*KVPair `json:"DynamicEnvs,omitempty"`
	Description string    `json:"Description,omitempty"`
}

type RunPipelineResponse struct {
	RecordId string `json:"RecordId"`
}

type Pipeline struct {
	Id               string             `json:"Id,omitempty"`
	Name             string             `json:"Name"`
	ClusterPool      string             `json:"ClusterPool"`
	Scm              SCM                `json:"Scm"`
	Trigger          *Trigger           `json:"Trigger,omitempty"`
	Env              *Env               `json:"Env,omitempty"`
	Stages           []Stage            `json:"Stages"`
	CreateTime       string             `json:"CreateTime,omitempty"`
	UpdateTime       string             `json:"UpdateTime,omitempty"`
	LastStatus       string             `json:"LastStatus,omitempty"`
	LastStagesStatus []*LastStageStatus `json:"LastStagesStatus,omitempty"`
	Triggerer        string             `json:"Triggerer,omitempty"`
	LastTriggerTime  string             `json:"LastTriggerTime,omitempty"`
	Language         string             `json:"Language"`
	TemplateId       string             `json:"TemplateId,omitempty"`
	CustomTemplate   bool               `json:"CustomTemplate,omitempty"`
	Cache            *Cache             `json:"Cache,omitempty"`
	Notification     *Notification      `json:"Notification,omitempty"`
	CleaningCache    bool               `json:"CleaningCache,omitempty"`
	ClusterId        string             `json:"ClusterId,omitempty"`
}

type SCM struct {
	Id            string   `json:"Id,omitempty"`
	Type          string   `json:"Type,omitempty"`
	Server        string   `json:"Server,omitempty"`
	DefaultBranch string   `json:"DefaultBranch,omitempty"`
	Disabled      bool     `json:"Disabled,omitempty"`
	Webhook       *Webhook `json:"Webhook,omitempty"`
}

type Webhook struct {
	Events  []string  `json:"Events,omitempty"`
	Filters []*KVPair `json:"Filters,omitempty"`
	Url     string    `json:"Url,omitempty"`
}

type PipelineFilter struct {
	Name string `json:"Name,omitempty"`
}

type Trigger struct {
	Scheduled *Scheduled `json:"Scheduled,omitempty"`
}

type Scheduled struct {
	Type string `json:"Type"`
	At   string `json:"At"`
}

type Env struct {
	Vars []*Var `json:"Vars,omitempty"`
}

type Var struct {
	Key     string `json:"Key"`
	Dynamic bool   `json:"Dynamic"`
	Secret  bool   `json:"Secret"`
	Value   string `json:"Value"`
}

type Stage struct {
	Id    string  `json:"Id,omitempty"`
	Name  string  `json:"Name"`
	Tasks []*Task `json:"Tasks"`
}

type Task struct {
	Id      string       `json:"Id,omitempty"`
	Name    string       `json:"Name"`
	Type    string       `json:"Type"`
	Steps   []Step       `json:"Steps"`
	Webhook *TaskWebhook `json:"Webhook,omitempty"`
}

type TaskWebhook struct {
	Events []string `json:"Events"`
	WebhookNotification
}

type Step struct {
	Id             string          `json:"Id,omitempty"`
	Name           string          `json:"Name"`
	Language       string          `json:"Language,omitempty"`
	Type           string          `json:"Type"`
	ApprovalConfig *ApprovalConfig `json:"ApprovalConfig,omitempty"`
	Params         []KVPair        `json:"Params"`
	CustomParams   []*KVPair       `json:"CustomParams,omitempty"`
}

type ApprovalConfig struct {
	Type         string                `json:"Type"`
	Approvers    []*User               `json:"Approvers,omitempty"`
	Notification *ApprovalNotification `json:"Notification,omitempty"`
}

type ApprovalNotification struct {
	Enable        bool     `json:"Enable"`
	Rules         []string `json:"Rules"`
	NotifyMethods []string `json:"NotifyMethods"`
}

type LastStageStatus struct {
	Name      string    `json:" Name"`
	Status    string    `json:" Status"`
	TaskNames []string  `json:" TaskNames,omitempty"`
	Infos     []*KVPair `json:" Infos,omitempty"`
}

type Cache struct {
	Dirs []Dir `json:"Dirs,omitempty"`
}

type Dir struct {
	Path        string `json:"Path"`
	Description string `json:"Description"`
}

type Notification struct {
	Events              []string               `json:"Events"`
	EnableStationLetter string                 `json:"EnableStationLetter,omitempty"`
	Sms                 []*SmsNotification     `json:"Sms,omitempty"`
	Email               []*EmailNotification   `json:"Email,omitempty"`
	Webhook             []*WebhookNotification `json:"Webhook,omitempty"`
	CustomContent       string                 `json:"CustomContent,omitempty"`
}

type SmsNotification struct {
	Enable     bool     `json:" Enable"`
	Targets    []string `json:" Targets"`
	OtherUsers []User   `json:" OtherUsers"`
}

type EmailNotification struct {
	Enable     bool     `json:" Enable"`
	Targets    []string `json:" Targets"`
	OtherUsers []User   `json:" OtherUsers"`
}

type WebhookNotification struct {
	Enable bool   `json:" Enable"`
	Url    string `json:" Url"`
	Token  string `json:" Token,omitempty"`
	Type   string `json:" Type"`
}

type User struct {
	AccountId int64  `json:"AccountId,omitempty"`
	UserId    int64  `json:"UserId,omitempty"`
	Name      string `json:"Name"`
}

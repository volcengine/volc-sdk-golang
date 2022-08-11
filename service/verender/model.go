package verender

import (
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/volcengine/volc-sdk-golang/base"
)

type GetStorageAccessResponse struct {
	ResponseMetaData `json:"ResponseMetaData"`
	WorkspaceResult  `json:"Result"`
	Time             time.Time `json:"Time"`
}

type Error struct {
	CodeN   int    `json:"CodeN"`
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

type ResponseMetaData struct {
	RequestId string `json:"RequestId"`
	Action    string `json:"Action"`
	Version   string `json:"Version"`
	Service   string `json:"Service"`
	Region    string `json:"Region"`
	Error     Error  `json:"Error"`
}

type WorkspaceResult struct {
	ClusterEndpoint string `json:"cluster_endpoint"`
	BucketName      string `json:"bucket_name"`
	BucketToken     string `json:"bucket_token"`
	VolumeName      string `json:"volume_name"`
	ClusterName     string `json:"cluster_name"`
	AccessKey       string `json:"access_key"`
	SecretKey       string `json:"secret_key"`
	WorkspaceId     int    `json:"workspace_id"`
}

type Resolutions struct {
	Width  int `json:"Width"`
	Height int `json:"Height"`
}

type FrameSettings struct {
	Start int `json:"Start"`
	End   int `json:"End"`
	Step  int `json:"Step"`
}

type FrameProcess struct {
	Total      int `json:"total"`
	Waiting    int `json:"waiting"`
	Processing int `json:"processing"`
	Success    int `json:"success"`
	Failed     int `json:"failed"`
	Stopped    int `json:"stopped"`
	Paused     int `json:"paused"`
	ToStart    int `json:"to_start"`
}

type LayerInformation struct {
	LayerIndex   int          `json:"layer_index"`
	FrameProcess FrameProcess `json:"frame_process"`
	Cost         float64      `json:"cost"`
}

type Statistics struct {
	CpuTime int     `json:"CpuTime"`
	GpuTime int     `json:"GpuTime"`
	Cost    float64 `json:"Cost"`
}

type Progress struct {
	Finished int `json:"Finished"`
	Total    int `json:"Total"`
}

type OutputImageTemple struct {
	Padding         int    `json:"Padding"`
	ImageNameTemple string `json:"ImageNameTemple"`
	SceneName       string `json:"SceneName"`
	Extension       string `json:"Extension"`
}

type CellSpec struct {
	ComputerResourceType  string `json:"ComputerResourceType"`
	ComputerResourceCount int    `json:"ComputerResourceCount"`
	ID                    int64  `json:"id"`
}

type JobInfo struct {
	// request
	UserId                int64              `json:"UserId,omitempty"`
	AccountId             int64              `json:"AccountId,omitempty"`
	UserName              string             `json:"UserName"`
	WorkspaceId           int64              `json:"WorkspaceId,omitempty"`
	Title                 string             `json:"Title"`
	Description           string             `json:"Description,omitempty"`
	DccTool               string             `json:"DccTool"`
	DccToolVersion        string             `json:"DccToolVersion,omitempty"`
	Renderer              string             `json:"Renderer"`
	RendererVersion       string             `json:"RendererVersion,omitempty"`
	Tryout                bool               `json:"Tryout"`
	TryoutFrameNumbers    []string           `json:"TryoutFrameNumbers,omitempty"`
	SceneFile             string             `json:"SceneFile"`
	OutputFormat          string             `json:"OutputFormat,omitempty"`
	Resolutions           []Resolutions      `json:"Resolutions,omitempty"`
	PathMapping           map[string]string  `json:"PathMapping,omitempty"`
	FrameSettings         FrameSettings      `json:"FrameSettings,omitempty"`
	Cameras               []string           `json:"Cameras,omitempty"`
	Layers                []string           `json:"Layers,omitempty"`
	LayerInformationList  []LayerInformation `json:"LayerInformationList,omitempty"`
	TimeoutReminder       int64              `json:"TimeoutReminder,omitempty"`
	TimeoutStopper        int64              `json:"TimeoutStopper,omitempty"`
	FrameProcess          FrameProcess       `json:"FrameProcess,omitempty"`
	OutputImageTemple     OutputImageTemple  `json:"OutputImageTemple,omitempty"`
	WantedCellSpecs       []CellSpec         `json:"WantedCellSpecs,omitempty"`
	UseLegacyRenderLayers *bool              `json:"UseLegacyRenderLayers,omitempty"`
	RenderImageTag        *string            `json:"RenderImageTag,omitempty"`
	Environment           map[string]string  `json:"Environment,omitempty"`
	CallbackTarget        *string            `json:"CallbackTarget,omitempty"`
	UserData              *string            `json:"UserData,omitempty"`

	// response
	JobId               string      `json:"JobId,omitempty"`
	Priority            int         `json:"Priority,omitempty"`
	CreatedAt           string      `json:"CreatedAt,omitempty"`
	StoppedAt           string      `json:"StoppedAt,omitempty"`
	Stage               string      `json:"Stage,omitempty"`
	Statistics          Statistics  `json:"Statistics,omitempty"`
	Error               string      `json:"Error,omitempty"`
	ResultBucket        string      `json:"ResultBucket,omitempty"`
	ResultDirectoryPath string      `json:"ResultDirectoryPath,omitempty"`
	FrameList           []FrameInfo `json:"FrameList,omitempty"`
	Paused              bool        `json:"Paused,omitempty"`
	CompletionAt        string      `json:"CompletionAt,omitempty"`
	OutputBytes         int64       `json:"OutputBytes,omitempty"`
	Progress            *Progress   `json:"Progress,omitempty"`
}

type JobResult struct {
	Job JobInfo `json:"Job"`
}

type GetJobResponse struct {
	ResponseMetaData `json:"ResponseMetaData"`
	JobResult        `json:"Result"`
}

type ListWorkspaceRequest struct {
	WorkspaceId        int64
	WorkspaceName      string
	PageNum            int64
	PageSize           int64
	StartPurchaseTime  string
	EndPurchaseTime    string
	OrderBy            string
	OrderType          string
	FuzzyWorkspaceName string
	FuzzySearchContent string
}

type WorkspaceList struct {
	Total      int         `json:"Total"`
	Workspaces []Workspace `json:"Workspaces"`
}

type ListWorkspaceResponse struct {
	MetaData   ResponseMetaData `json:"ResponseMetadata"`
	Workspaces WorkspaceList    `json:"Result,omitempty"`
}

type ListJobRequest struct {
	OrderBy  string
	PageNum  int
	PageSize int
	Tryout   int
	Keyword  string
	Status   string
	JobType  string
}

type JobList struct {
	Jobs      []JobInfo `json:"Jobs"`
	TotalJobs int       `json:"TotalJobs"`
}

type ListJobResponse struct {
	MetaData ResponseMetaData `json:"ResponseMetadata"`
	JobsData JobList          `json:"Result,omitempty"`
}

type UserInfo struct {
	AccountId int    `json:"AccountId,omitempty"`
	UserId    int    `json:"UserId"`
	UserName  string `json:"UserName"`
}

type GetUserInfoResult struct {
	MetaData ResponseMetaData `json:"ResponseMetadata"`
	User     UserInfo         `json:"Result,omitempty"`
}

type RenderJobInfo struct {
	Job JobInfo `json:"RenderJob"`
}

type CreateJobResponse struct {
	Metadata  ResponseMetaData `json:"ResponseMetadata"`
	JobDetail RenderJobInfo    `json:"Result,omitempty"`
}

type BatchJobRequest struct {
	JobIds []string `json:"JobIds"`
}

type GetFramesRequest struct {
	RenderJobId string
	PageNum     int
	PageSize    int
}

type FrameInfo struct {
	Start        int    `json:"Start"`
	End          int    `json:"End"`
	Status       string `json:"Status"`
	CpuTime      int64  `json:"CpuTime"`
	GpuTime      int64  `json:"GpuTime"`
	StartingTime string `json:"StartingTime"`
	FinishedTime string `json:"FinishedTime"`
	ExitCode     int32  `json:"ExitCode,omitempty"`
	Error        string `json:"Error,omitempty"`
}

type FrameList struct {
	Frames      []FrameInfo `json:"FrameList"`
	TotalFrames int         `json:"TotalFrames"`
}

type GetFramesResponse struct {
	MetaData ResponseMetaData `json:"ResponseMetadata"`
	Frames   FrameList        `json:"Result"`
}

type LayerRequest struct {
	LayerIndex int   `json:"LayerIndex"`
	PageNum    int   `json:"PageNum,omitempty"`
	PageSize   int   `json:"PageSize,omitempty"`
	Statuses   []int `json:"Statuses,omitempty"`
}

type GetLayerFramesRequest struct {
	RenderJobId   string         `json:"RenderJobId"`
	LayerRequests []LayerRequest `json:"LayerRequests"`
}

type LayerFrameInfo struct {
	FrameNumber  int     `json:"frame_number"`
	Status       int     `json:"status"`
	StartingTime string  `json:"starting_time,omitempty"`
	FinishedTime string  `json:"finished_time,omitempty"`
	CpuTime      int64   `json:"cpu_time,omitempty"`
	GpuTime      int64   `json:"gpu_time,omitempty"`
	Cost         float64 `json:"cost,omitempty"`
}

type LayerFrameList struct {
	LayerFrameList []LayerFrameInfo `json:"FrameInformationList"`
	Summary        FrameProcess     `json:"FrameProcess"`
	CpuTime        int64            `json:"CpuTime"`
	GpuTime        int64            `json:"GpuTime"`
	Cost           float64          `json:"Cost"`
	Total          int              `json:"Total"`
}

type GetLayerFramesResponse struct {
	MetaData    ResponseMetaData `json:"ResponseMetadata"`
	LayerFrames []LayerFrameList `json:"Result"`
}

type Workspace struct {
	ClusterEndpoint       string        `json:"-"`
	BucketName            string        `json:"-"`
	BucketToken           string        `json:"-"`
	VolumeName            string        `json:"-"`
	ClusterName           string        `json:"-"`
	AccessKey             string        `json:"-"`
	SecretKey             string        `json:"-"`
	Time                  time.Time     `json:"-"`
	MinioClient           *minio.Client `json:"-"`
	Client                *base.Client  `json:"-"`
	WorkspaceId           int           `json:"Id,omitempty"`
	WorkspaceName         string        `json:"Name,omitempty"`
	Description           string        `json:"Description,omitempty"`
	ResourcePoolId        int           `json:"ResourcePoolId,omitempty"`
	CpsId                 int           `json:"CpsId,omitempty"`
	StorageTotal          int64         `json:"StorageTotal,omitempty"`
	AccountId             int           `json:"AccountId,omitempty"`
	UserId                int           `json:"UserId,omitempty"`
	AccountUserName       string        `json:"AccountUserName,omitempty"`
	CreatedAt             string        `json:"CreatedAt,omitempty"`
	UpdateAt              string        `json:"UpdateAt,omitempty"`
	Specification         *StorageSpec  `json:"Specification,omitempty"`
	Status                string        `json:"Status,omitempty"`
	HardwareSpecification *HardwareSpec `json:"HardwareSpecification,omitempty"`
}

type ResourcePool struct {
	ResourcePoolId     int    `json:"ResourcePoolId"`
	ResourcePoolName   string `json:"ResourcePoolName"`
	ResourcePoolRegion string `json:"ResourcePoolRegion"`
	CpsId              int    `json:"CpsId"`
	CpsName            string `json:"CpsName"`
	SupportCpu         string `json:"SupportCpu"`
	SupportGpu         string `json:"SupportGpu"`
}

type ListResourcePoolResponse struct {
	MetaData      ResponseMetaData `json:"ResponseMetadata"`
	ResourcePools []ResourcePool   `json:"Result"`
}

type GetWorkspaceLimitResponse struct {
	MetaData       ResponseMetaData `json:"ResponseMetadata"`
	WorkspaceLimit WorkspaceLimit   `json:"Result"`
}

type WorkspaceLimit struct {
	WorkspaceSizeUpperLimit   int `json:"WorkspaceSizeUpperLimit"`
	WorkspaceAmountUpperLimit int `json:"WorkspaceAmountUpperLimit"`
}

type HardwareSpec struct {
	ResourcePoolId     int    `json:"ResourcePoolId,omitempty"`
	ResourcePoolName   string `json:"ResourcePoolName,omitempty"`
	ResourcePoolRegion string `json:"ResourcePoolRegion,omitempty"`
	CpsId              int    `json:"CpsId,omitempty"`
	CpsName            string `json:"CpsName,omitempty"`
	SupportCpu         string `json:"SupportCpu,omitempty"`
	SupportGpu         string `json:"SupportGpu,omitempty"`
}

type StorageSpec struct {
	Id               int    `json:"Id,omitempty"`
	StorageTotal     int64  `json:"StorageTotal,omitempty"`
	StorageAvailable int64  `json:"StorageAvailable,omitempty"`
	ValidStartTime   string `json:"ValidStartTime,omitempty"`
}

type PurchaseWorkspaceResponse struct {
	MetaData      ResponseMetaData `json:"ResponseMetadata"`
	WorkspaceInfo Workspace        `json:"Result"`
}

type Job struct {
	JobInfo
	Client *base.Client
}

type GetWorkspaceHardwareSpecificationResponse struct {
	MetaData      ResponseMetaData                 `json:"ResponseMetaData"`
	HardwareSpecs []WorkspaceHardwareSpecification `json:"Result"`
}

type WorkspaceHardwareSpecification struct {
	Type   string `json:"Type"`
	Count  int    `json:"Count"`
	SpecID int    `json:"SpecId"`
}

type StatisticsFilter struct {
	WorkspaceIds []int    `json:"WorkspaceIds,omitempty"`
	UserIds      []int    `json:"UserIds,omitempty"`
	JobTypes     []string `json:"JobTypes,omitempty"`
	Keyword      string   `json:"Keyword,omitempty"`
}

type StatisticsRequest struct {
	StartTime  string           `json:"StartTime,omitempty"`
	EndTime    string           `json:"EndTime,omitempty"`
	TimeType   string           `json:"TimeType,omitempty"`
	PageNum    int64            `json:"PageNum,omitempty"`
	PageSize   int64            `json:"PageSize,omitempty"`
	OrderField string           `json:"OrderField,omitempty"`
	OrderBy    string           `json:"OrderBy,omitempty"`
	FileName   string           `json:"FileName,omitempty"`
	Filter     StatisticsFilter `json:"-"`
}

type StatisticsSummary struct {
	TotalCost    float64 `json:"TotalCost"`
	TotalJobs    int64   `json:"TotalJobs"`
	CpuTime      int64   `json:"CpuTime"`
	GpuTime      int64   `json:"GpuTime"`
	TotalFrames  int64   `json:"TotalFrames"`
	TotalBytes   int64   `json:"TotalBytes"`
	TotalJobTime int64   `json:"TotalJobTime"`
}

type StatisticsRecord struct {
	StartTime     string  `json:"StartTime"`
	EndTime       string  `json:"EndTime"`
	JobType       string  `json:"JobType"`
	CpuTime       int64   `json:"CpuTime"`
	GpuTime       int64   `json:"GpuTime"`
	JobTime       int64   `json:"JobTime"`
	TotalFrames   int64   `json:"TotalFrames"`
	TotalJobs     int64   `json:"TotalJobs"`
	TotalBytes    int64   `json:"TotalBytes,omitempty"`
	TotalCost     float64 `json:"TotalCost,omitempty"`
	JobId         string  `json:"JobId,omitempty"`
	JobName       string  `json:"JobName,omitempty"`
	JobDesc       string  `json:"JobDescription,omitempty"`
	WorkspaceId   int     `json:"WorkspaceId,omitempty"`
	WorkspaceName string  `json:"WorkspaceName,omitempty"`
	UserId        int     `json:"UserId,omitempty"`
	UserName      string  `json:"UserName,omitempty"`
}

type AccountStatistics struct {
	Summary    StatisticsSummary  `json:"Summary"`
	Statistics []StatisticsRecord `json:"Statistics"`
	Workspaces []WorkspaceInfo    `json:"Workspaces"`
	Users      []UserInfo         `json:"Users"`
}

type GetAccountStatisticsResponse struct {
	MetaData    ResponseMetaData  `json:"ResponseMetadata"`
	AccountStat AccountStatistics `json:"Result"`
}

type WorkspaceInfo struct {
	WorkspaceId   int    `json:"WorkspaceId"`
	WorkspaceName string `json:"WorkspaceName"`
}

type AccountStatisticDetails struct {
	Total      int64              `json:"Total"`
	Items      []StatisticsRecord `json:"Items"`
	Workspaces []WorkspaceInfo    `json:"Workspaces"`
	Users      []UserInfo         `json:"Users"`
	JobTypes   []string           `json:"JobTypes"`
}

type GetAccountStatisticDetailsResponse struct {
	MetaData ResponseMetaData        `json:"ResponseMetadata"`
	Details  AccountStatisticDetails `json:"Result"`
}

type JobOverallStatistics struct {
	Failed  int64 `json:"Failed"`
	Running int64 `json:"Running"`
	Success int64 `json:"Success"`
	Waiting int64 `json:"Waiting"`
}

type GetJobOverallStatisticsResponse struct {
	MetaData ResponseMetaData     `json:"ResponseMetadata"`
	Stat     JobOverallStatistics `json:"Result"`
}

type MessageInfo struct {
	Id             int    `json:"Id"`
	Publisher      string `json:"Publisher"`
	MessageType    string `json:"MessageType"`
	MessageSubType string `json:"MessageSubType"`
	Title          string `json:"Title"`
	Content        string `json:"Content"`
	MarkAsRead     bool   `json:"MarkAsRead"`
	CreatedAt      string `json:"CreatedAt"`
}

type MessageList struct {
	Messages    []MessageInfo `json:"Messages"`
	Total       int64         `json:"Total"`
	TotalUnread int64         `json:"TotalUnread"`
}

type ListMessageResponse struct {
	MetaData ResponseMetaData `json:"ResponseMetadata"`
	Messages MessageList      `json:"Result"`
}

type ListMessageRequest struct {
	PageNum  int64 `json:"PageNum,omitempty"`
	PageSize int64 `json:"PageSize,omitempty"`
}

type MessageIdList struct {
	MessageIds []int `json:"MessageIds"`
}

type ObjectInfo struct {
	Etag          string
	Key           string
	ContentLength int64
	ContentType   string
	LastModified  time.Time
}

type DeleteWorkspaceResp struct {
	MetaData ResponseMetaData `json:"ResponseMetadata"`
}

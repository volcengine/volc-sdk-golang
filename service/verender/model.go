package verender

import (
    "time"

    "github.com/volcengine/volc-sdk-golang/base"
)

type Verender struct {
    ftransClientAddr     string `json:"ftrans_client_addr"`
    ftransClientAclToken string `json:"ftrans_client_acl_token"`
    ftransProxyAddr      string `json:"ftrans_proxy_addr"`
    ftransServerIsp      string `json:"ftrans_server_isp"`
    Client               *base.Client
}

type Error struct {
    CodeN   int    `json:"CodeN"`
    Code    string `json:"Code"`
    Message string `json:"Message"`
}

type ResponseMetadata struct {
    RequestID string `json:"RequestId"`
    Action    string `json:"Action"`
    Version   string `json:"Version"`
    Service   string `json:"Service"`
    Region    string `json:"Region"`
    Error     Error  `json:"Error"`
}

type Specification struct {
    ID               int64     `json:"Id"`
    StorageAvailable int64     `json:"StorageAvailable"`
    StorageTotal     int64     `json:"StorageTotal"`
    ValidStartTime   time.Time `json:"ValidStartTime"`
}

type CellSpec struct {
    ID             int64   `json:"id"`
    CPUCount       int64   `json:"cpu_count"`
    CPUType        string  `json:"cpu_type"`
    GPUCount       int64   `json:"gpu_count"`
    GPUType        string  `json:"gpu_type"`
    SystemInfo     string  `json:"system_info"`
    CellType       string  `json:"cell_type"`
    ResourcePoolID int64   `json:"resource_pool_id"`
    UintPrice      float64 `json:"unit_price"`
}

type Dcc struct {
    Dcc        string `json:"Dcc"`
    DccVersion string `json:"DccVersion"`
}

type ListDccsResult struct {
    Dccs []Dcc `json:"Dccs"`
}

type ListDccsResponse struct {
    ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
    Result           *ListDccsResult  `json:"Result"`
}

type ListAccountDccPluginsReq struct {
    SpecTemplateId int64  `json:"SpecTemplateId"`
    Dcc            string `json:"Dcc,omitempty"`
    DccVersion     string `json:"DccVersion"`
    RegionId       int64  `json:"RegionId"`
}

type DccPlugin struct {
    Id             int64  `json:"Id"`
    Dcc            string `json:"Dcc"`
    DccVersion     string `json:"DccVersion"`
    Plugin         string `json:"Plugin"`
    PluginVersion  string `json:"PluginVersion"`
    NeedLicense    bool   `json:"NeedLicense"`
    RendererPlugin bool   `json:"RendererPlugin"`
    DccLicense     bool   `json:"DccLicense"`
}

type ListAccountDccPluginsResult struct {
    Plugins []DccPlugin `json:"Plugins"`
}

type ListAccountDccPluginsResp struct {
    ResponseMetadata ResponseMetadata            `json:"ResponseMetadata"`
    Result           ListAccountDccPluginsResult `json:"Result"`
}

type StorageAccess struct {
    BucketName          string        `json:"bucket_name"`
    FtransAclToken      string        `json:"ftrans_acl_token"`
    FtransSecurityToken string        `json:"ftrans_security_token"`
    FtransQuicServer    string        `json:"ftrans_quic_server"`
    FtransCertName      string        `json:"ftrans_cert_name"`
    FtransS10Server     string        `json:"ftrans_s10_server"`
    FtransCert          string        `json:"cert_pem"`
    FtransKey           string        `json:"private_key_pem"`
    FtransClient        *FtransClient `json:"-"`
    ExpiredNSec         int64         `json:"-"`
}

type HardwareSpecification struct {
    SupportCpu string `json:"SupportCpu"`
    SupportGpu string `json:"SupportGpu"`
    RegionName string `json:"RegionName"`
    RegionId   int64  `json:"RegionId"`
}

type Workspace struct {
    Id                    int64                 `json:"Id"`
    CreatedAt             time.Time             `json:"CreatedAt"`
    Name                  string                `json:"Name"`
    Description           string                `json:"Description"`
    AccountID             int64                 `json:"AccountId"`
    UserID                int64                 `json:"UserId"`
    AccountUserName       string                `json:"AccountUserName"`
    RegionId              int64                 `json:"RegionId"`
    SpecIdList            []int64               `json:"SpecIdList"`
    Status                string                `json:"Status"`
    Specification         Specification         `json:"Specification,omitempty"`
    HardwareSpecification HardwareSpecification `json:"HardwareSpecification"`
    client                *base.Client          `json:"-"`
    storageAccess         *StorageAccess        `json:"-"`
    ConvertToLowerCase    bool                  `json:"ConvertToLowerCase"`
}

type WorkspaceLimit struct {
    WorkspaceSizeUpperLimit   int
    WorkspaceAmountUpperLimit int
}

type FileInfo struct {
    Key           string    `json:"Key"` // needed in CreateRenderJob
    ContentLength int64     `json:"ContentLength"`
    LastModified  time.Time `json:"LastModified"`
}

type Plugin struct {
    PluginName     string `json:"Plugin,omitempty"`
    PluginVersion  string `json:"PluginVersion"`
    Name           string `json:"Name,omitempty"`
    Version        string `json:"Version,omitempty"`
    NeedLicense    bool   `json:"NeedLicense"`
    RendererPlugin bool   `json:"RendererPlugin"`
}

type RenderSetting struct {
    AccountID          int64     `json:"AccountId,omitempty"`
    UserID             int64     `json:"UserId,omitempty"`
    WorkspaceID        int64     `json:"WorkspaceId,omitempty"`
    Name               string    `json:"Name,omitempty"`
    Dcc                string    `json:"Dcc,omitempty"`
    DccVersion         string    `json:"DccVersion,omitempty"`
    Plugins            []*Plugin `json:"Plugins,omitempty"`
    RenderLayerMode    string    `json:"RenderLayerMode,omitempty"`
    ProjectPath        string    `json:"ProjectPath,omitempty"`
    FramesCountOneCell int8      `json:"FramesCountOneCell,omitempty"`
    CellSpecID         int64     `json:"CellSpecId,omitempty"`
    Id                 int64     `json:"Id,omitempty"`
    IsDeleted          bool      `json:"IsDeleted,omitempty"`
}

type User struct {
    AccountID int64  `json:"AccountId"`
    UserID    int64  `json:"UserId"`
    UserName  string `json:"UserName"`
}

type ListWorkspaceRequest struct {
    WorkspaceID        int64  `json:"WorkspaceId"`
    WorkspaceName      string `json:"WorkspaceName"`
    PageNum            int64  `json:"PageNum"`
    PageSize           int64  `json:"PageSize"`
    StartPurchaseTime  string `json:"StartPurchaseTime"`
    EndPurchaseTime    string `json:"EndPurchaseTime"`
    OrderType          string `json:"OrderType"`
    OrderField         string `json:"OrderField"`
    FuzzyWorkspaceName string `json:"FuzzyWorkspaceName"`
    FuzzySearchContent string `json:"FuzzySearchContent"`
    ListScope          string `json:"ListScope"`
}

type WorkspaceList struct {
    Total      int          `json:"Total"`
    Workspaces []*Workspace `json:"Workspaces"`
}

type ListWorkspaceResponse struct {
    ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
    Result           WorkspaceList    `json:"Result"`
}

type GetStorageAccessResponse struct {
    ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
    Result           StorageAccess    `json:"Result"`
}

type GetCurrentUserResponse struct {
    ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
    Result           *User            `json:"Result"`
}

type CellSpecList struct {
    CellSpecs []*CellSpec `json:"cell_specs"`
}

type ListCellSpecResponse struct {
    ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
    Result           *CellSpecList    `json:"Result"`
}

type AddRenderSettingResult struct {
    RenderSettingName string `json:"RenderSettingName"`
    RenderSettingID   int64  `json:"RenderSettingId"`
}

type AddRenderSettingResponse struct {
    ResponseMetadata ResponseMetadata        `json:"ResponseMetadata"`
    Result           *AddRenderSettingResult `json:"Result"`
}

type VoidResponse struct {
    ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
}

type ListRenderSettingResponse struct {
    ResponseMetadata ResponseMetadata            `json:"ResponseMetadata"`
    Result           map[string][]*RenderSetting `json:"Result"`
}

type RenderSettingList struct {
    Settings []*RenderSetting `json:"Settings"`
}

type GetRenderSettingResponse struct {
    ResponseMetadata ResponseMetadata   `json:"ResponseMetadata"`
    Result           *RenderSettingList `json:"Result"`
}

type OutputImageTemplate struct {
    Padding         int8   `json:"Padding"`
    ImageNameTemple string `json:"ImageNameTemple"`
    SceneName       string `json:"SceneName"`
    Extension       string `json:"Extension"`
}

type Resolution struct {
    Height int64 `json:"Height"`
    Width  int64 `json:"Width"`
}

type Frame struct {
    Start         float64  `json:"Start"`
    End           float64  `json:"End"`
    Step          float64  `json:"Step"`
    RenumberStart *float64 `json:"RenumberStart,omitempty"`
}

type LayerConfig struct {
    LayerIndex  int8       `json:"LayerIndex"`
    LayerName   string     `json:"LayerName"`
    Frame       Frame      `json:"Frame"`
    Resolutions Resolution `json:"Resolutions"`
    Cameras     []string   `json:"Cameras"`
}

type RenderJob struct {
    // request
    AccountID                int64             `json:"AccountId"`
    UserID                   int64             `json:"UserId"`
    UserName                 string            `json:"UserName"`
    WorkspaceID              int64             `json:"WorkspaceId"`
    Title                    string            `json:"Title,omitempty"`
    Description              string            `json:"Description,omitempty"`
    Tryout                   bool              `json:"Tryout"`
    TryoutFrameNumbers       []string          `json:"TryoutFrameNumbers"`
    SceneFile                string            `json:"SceneFile"`
    PathMapping              map[string]string `json:"PathMapping"`
    TimeoutReminderEachFrame int64             `json:"TimeoutReminderEachFrame"`
    TimeoutStopperEachFrame  int64             `json:"TimeoutStopperEachFrame"`
    LayerConfig              []*LayerConfig    `json:"LayerConfig"`
    RenderSetting            *RenderSetting    `json:"RenderSetting"`
    FramesCountEachCell      int64             `json:"FramesCountEachCell"`
    PluginData               string            `json:"PluginData"`
    Renderer                 string            `json:"Renderer"`
    DeviceName               string            `json:"DeviceName"`
    AutoRetryNumber          int64             `json:"AutoRetryNumber"`

    // response
    JobID             string    `json:"JobId,omitempty"`
    Priority          int64     `json:"Priority,omitempty"`
    CreatedAt         string    `json:"CreatedAt,omitempty"`
    StoppedAt         string    `json:"StoppedAt,omitempty"`
    CompletionAt      string    `json:"CompletionAt"`
    Stage             string    `json:"Stage,omitempty"`
    Error             string    `json:"Error,omitempty"`
    RenderSettingID   int64     `json:"RenderSettingId"`
    RenderSettingName string    `json:"RenderSettingName"`
    Plugins           []*Plugin `json:"Plugins"`
}

type CreateRenderJobResult struct {
    RenderJob *RenderJob `json:"RenderJob"`
}

type CreateRenderJobResponse struct {
    ResponseMetadata ResponseMetadata       `json:"ResponseMetadata"`
    Result           *CreateRenderJobResult `json:"Result"`
}

type ListRenderJobRequest struct {
    PageNum        int64    `json:"PageNum"`
    PageSize       int64    `json:"PageSize"`
    OrderType      string   `json:"OrderType"`
    OrderField     string   `json:"OrderField"`
    Keyword        string   `json:"Keyword"`
    Status         []string `json:"Status"`
    RenderSettings []int64  `json:"RenderSetting"`
    DeviceName     string   `json:"DeviceName"`
    UserId         int64    `json:"UserId"`
}

type ListRenderJobResult struct {
    TotalJobs      int64                     `json:"TotalJobs"`
    Jobs           []*RenderJob              `json:"Jobs"`
    RenderSettings []*AddRenderSettingResult `json:"RenderSettings"`
}

type ListRenderJobResponse struct {
    ResponseMetadata ResponseMetadata     `json:"ResponseMetadata"`
    Result           *ListRenderJobResult `json:"Result"`
}

type GetRenderJobResult struct {
    Job *RenderJob `json:"Job"`
}

type GetRenderJobResponse struct {
    ResponseMetadata ResponseMetadata    `json:"ResponseMetadata"`
    Result           *GetRenderJobResult `json:"Result"`
}

type RetryLayerFrame struct {
    LayerIndex   int64  `json:"LayerIndex"`
    FrameIndexes string `json:"FrameIndexes"`
}

type RetryJobRequest struct {
    JobID           string             `json:"JobId"`
    AllFailedFrames bool               `json:"AllFailedFrames"`
    CustomFrames    []*RetryLayerFrame `json:"CustomFrames"`
}

type BatchJobIDList struct {
    JobIDList []string `json:"JobIds"`
}

type BatchJobPriority struct {
    JobIDList []string `json:"JobIds"`
    Priority  int8     `json:"Priority"`
}

type ListJobOutputRequest struct {
    StartTime  string   `json:"start_time"`
    EndTime    string   `json:"end_time"`
    PageNum    int      `json:"page_num"`
    PageSize   int      `json:"page_size"`
    Type       string   `json:"type"`
    Status     string   `json:"status"`
    OrderType  string   `json:"order_type"`
    OrderField string   `json:"order_field"`
    JobIDList  []string `json:"job_id_list"`
}

type ListJobOutputResult struct {
    Total int64    `json:"Total"`
    Files []string `json:"Files"`
}

type ListJobOutputResponse struct {
    ResponseMetadata ResponseMetadata     `json:"ResponseMetadata"`
    Result           *ListJobOutputResult `json:"Result"`
}

type FrameOutputFilter struct {
    Frames       []float64 `json:"frames"`
    IncludeThumb bool      `json:"include_thumb"`
    IncludeImage bool      `json:"include_image"`
}

type GetJobOutputFilter struct {
    IncludeLog bool                          `json:"include_log"`
    Layers     map[string]*FrameOutputFilter `json:"layers"`
}

type GetJobOutputRequest struct {
    JobID  string `json:"job_id"`
    Filter *GetJobOutputFilter
}

type FrameOutput struct {
    Frame  float64  `json:"frame"`
    Images []string `json:"images"`
    Thumb  string   `json:"thumb"`
}

type GetJobOutputResult struct {
    ImageList map[string][]*FrameOutput `json:"image_list"`
    LogList   []string                  `json:"log_list"`
}

type GetJobOutputResponse struct {
    ResponseMetadata ResponseMetadata    `json:"ResponseMetadata"`
    Result           *GetJobOutputResult `json:"Result"`
}

type UpdateJobOutputRequest struct {
    AccountID   int64    `json:"account_id"`
    UserID      int64    `json:"user_id"`
    WorkspaceID int64    `json:"workspace_id"`
    JobID       string   `json:"job_id"`
    Files       []string `json:"files"`
}

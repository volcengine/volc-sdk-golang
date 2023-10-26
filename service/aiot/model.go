package aiot

import (
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

// CascadeStatus 级联任务状态
type CascadeStatus string

const (
	CascadeStatusCreated        = "created"         // 初始状态
	CascadeStatusStopped        = "stopped"         // 已停用
	CascadeStatusTryingRegister = "trying_register" // 尝试注册中
	CascadeStatusRegistered     = "registered"      // 注册成功
	CascadeStatusRegisterFail   = "register_fail"   // 注册失败
)

type DeviceType string

const (
	DeviceTypeCamera   DeviceType = "IPC"
	DeviceTypeNVR      DeviceType = "NVR"
	DeviceTypePlatform DeviceType = "Platform" //国标平台
)

const (
	DeviceStatusOnline          = "online"
	DeviceStatusOffline         = "offline"
	DeviceStatusUnregistered    = "unregistered"
	DeviceStatusInvalidPassword = "invalid_password"
	DeviceStatusChannelUpdate   = "channel_update"
)

// CloudControl - Action type
const (
	PTZControl    = "PTZControl"
	FiControl     = "FiControl"
	PresetControl = "PresetControl"
)

// CloudControl - PTZ控制 cmd type
const (
	PtzCmdStop      = "stop"
	PtzCmdRight     = "right"
	PtzCmdLeft      = "left"
	PtzCmdDown      = "down"
	PtzCmdUp        = "up"
	PtzCmdRightUp   = "rightup"
	PtzCmdRightDown = "rightdown"
	PtzCmdLeftUp    = "leftup"
	PtzCmdLeftDown  = "leftdown"
	PtzCmdZoomIn    = "zoomin"
	PtzCmdZoomOut   = "zoomout"
)

// CloudControl - 光圈控制 cmd type
const (
	FiCmdStop      = "stop"
	FiCmdFocusFar  = "focusfar"
	FiCmdFocusNear = "focusnear"
	FiCmdIrisIn    = "irisin"
	FiCmdIrisOut   = "irisout"
)

// CloudControl - 预置位控制 cmd type
const (
	PresetCmdSet    = "set"
	PresetCmdGoto   = "goto"
	PresetCmdRemove = "remove"
)

// CloudControl - 巡航控制 cmd type
const (
	CruiseCmdAdd         = "add"
	CruiseCmdRemove      = "remove"
	CruiseCmdSetSpeed    = "setspeed"
	CruiseCmdSetStayTime = "setstay"
	CruiseCmdStart       = "start"
	CruiseCmdStop        = "stop"
)

func (d DeviceType) String() string {
	return string(d)
}

type PageRequest struct {
	PageNumber int
	PageSize   int
}

type PageResult struct {
	PageNumber int64
	PageSize   int64
	TotalCount int64
}

type IDResponse struct {
	ID string `json:"ID"`
}

type RawResponse struct {
	ResponseMetadata base.ResponseMetadata
}

type SpaceResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type SpaceResponseV3 struct {
	ResponseMetadata base.ResponseMetadata
}

type GetSpaceResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SpaceView `json:"Result,omitempty"`
}

type SpaceGB struct {
	Domain       string `json:"-"`
	PullOnDemand bool   `json:"PullOnDemand"`
}

type SpaceDomain struct {
	Domain        string           `json:"Domain"`
	Type          string           `json:"Type"`
	Status        string           `json:"Status"`
	CNAME         string           `json:"CNAME"`
	CNAMEStatus   string           `json:"CNAMEStatus"`
	Default       bool             `json:"Default"`
	Https         SpaceDomainHttps `json:"https"`
	CreatedAt     string           `json:"CreatedAt"`
	UpdatedAt     string           `json:"UpdatedAt"`
	TimeFlag      bool             `json:"TimeFlag"`
	MainKey       string           `json:"MainKey"`
	SpareKey      string           `json:"SpareKey"`
	ValidDuration int64            `json:"ValidDuration"`
}

type SpaceDomainHttps struct {
	Enable        bool   `json:"Enable"`
	CertificateID string `json:"CertificateID"`
}

type SipServer struct {
	SipId     string         `json:"SipId"`
	SipDomain string         `json:"SipDomain"`
	SipIp     string         `json:"SipIp"`
	SipPort   map[string]int `json:"SipPort"`
}

type SimpleTemplate struct {
	TemplateID       string `json:"TemplateID"`
	TemplateName     string `json:"TemplateName"`
	ApplyToALLStream bool   `json:"-"`
	BindTime         string `json:"BindTime"`
	FCDNTemplateName string `json:"FCDNTemplateName,omitempty"`
}

type SpaceView struct {
	SpaceName     string                 `json:"SpaceName"`
	Description   string                 `json:"Description"`
	SpaceID       string                 `json:"SpaceID"`
	Region        string                 `json:"Region"`
	AccessType    string                 `json:"AccessType"` //GB/RTMP
	Status        string                 `json:"Status"`
	GB            SpaceGB                `json:"GB"`
	HLSLowLatency bool                   `json:"HLSLowLatency"`
	CreatedAt     string                 `json:"CreatedAt"`
	UpdatedAt     string                 `json:"UpdatedAt"`
	Domains       map[string]SpaceDomain `json:"Domains"`
	SipServer     *SipServer             `json:"SipServer"`
	CallbackURL   *string                `json:"CallbackURL"`
	Template      struct {
		Screenshot SimpleTemplate   `json:"Screenshot"`
		Record     SimpleTemplate   `json:"Record"`
		AI         SimpleTemplate   `json:"AI"`
		Trans      []SimpleTemplate `bson:"Trans"`
	} `bson:"Template"`
}

type ListSpacesResult struct {
	PageResult
	Spaces []*SpaceView
}

type ListSpacesResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListSpacesResult `json:"Result,omitempty"`
}

type CreateDeviceResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *IDDevice `json:"Result,omitempty"`
}

type IDDevice struct {
	ID string `json:"ID"`
	*DeviceView
}

type AlertNotification struct {
	Enabled     bool `bson:"Enabled"`
	Phone       bool `bson:"Phone"`
	Device      bool `bson:"Device"`
	SMS         bool `bson:"SMS"`
	GPS         bool `bson:"GPS"`
	Video       bool `bson:"Video"`
	DeviceFault bool `bson:"DeviceFault"`
	Other       bool `bson:"Other"`
}

type DeviceStreams struct {
	StreamID string   `bson:"stream_id"`
	PullUrls []string `bson:"pull_urls"`
	PushUrl  string   `bson:"push_url"`
	Status   string   `bson:"status"`
}

type Coordinates struct {
	Longitude *float64 `json:"Longitude"` //经度
	Latitude  *float64 `json:"Latitude"`  //纬度
}

type CommonCount struct {
	OnlineCount int64 `json:"OnlineCount"`
	AllCount    int64 `json:"AllCount"`
}

type DeviceView struct {
	SpaceID              string                      `json:"SpaceID"`
	SpaceName            string                      `json:"SpaceName"`
	Type                 string                      `json:"Type"`
	DeviceName           string                      `json:"DeviceName"`
	DeviceID             string                      `json:"DeviceID"`
	DeviceNSID           string                      `json:"DeviceNSID"`
	Username             string                      `json:"Username"`
	Password             string                      `json:"Password"`
	Description          string                      `json:"Description"`
	AutoPullAfterRegiter bool                        `json:"AutoPullAfterRegiter"`
	CreatedAt            string                      `json:"CreatedAt"`
	UpdatedAt            string                      `json:"UpdatedAt"`
	Manufactory          string                      `json:"Manufactory"`
	AlertNotification    *AlertNotification          `json:"AlertNotification"`
	DeviceStreams        map[string]*DeviceStreams   `json:"DeviceStreams"`
	DeviceSubStreams     map[string][]*DeviceStreams `json:"DeviceSubStreams"`
	Status               string                      `json:"Status"`
	ChannelNum           int                         `json:"ChannelNum"`
	RtpTransportTcp      bool                        `json:"RtpTransportTcp"` //流传输协议tcp为true，否则默认udp
	ContactCount         int                         `json:"ContactCount"`    //Device的contact数量, 用于前端提示多device注册同一个账号

	Location     string       `json:"Location"`
	Coordinates  *Coordinates `json:"Coordinates"`
	DeviceIP     string       `json:"DeviceIP"`
	CommonCount  *CommonCount `json:"CommonCount,omitempty" bson:"CommonCount,omitempty"`
	UseSubStream bool         `json:"UseSubStream"` // 是否启用子码流
}

type DeviceResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type CloudRecordPlayResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           map[string]interface{} `json:"Result,omitempty"`
}

type LocalMediaDownloadResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           LocalMediaDownloadResp `json:"Result,omitempty"`
}

type LocalMediaDownloadResp struct {
	Result bool   `json:"Result"`
	ID     string `json:"ID"`
}

type LocalMediaDownloadResponseV3 struct {
	ResponseMetadata base.ResponseMetadata
	Result           LocalMediaDownloadRespV3 `json:"Result,omitempty"`
}

type LocalMediaDownloadRespV3 struct {
	DownloadID string `json:"DownloadID"`
}

type ListGBMediaResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListGBMediaResp `json:"Result"`
}

type ListGBMediaResp struct {
	PageResult
	GBMedias []*GBMediaDownload
}

type GBMediaDownload struct {
	ID               string        `json:"ID"` //=streamID
	StartTime        int64         `json:"StartTime"`
	EndTime          int64         `json:"EndTime"`
	CallID           string        `json:"CallID"`
	Status           string        `json:"Status"`
	Url              string        `json:"Url"`
	FileSize         int64         `json:"FileSize"`
	Msg              string        `json:"Msg"`
	FileStreamLength time.Duration `json:"FileStreamLength"`
	FileName         string        `json:"FileName"`
	DeviceMsg        string        `json:"DeviceMsg"` //设备上传时异常bye存储信息
	CreateAt         int64         `json:"CreateAt"`
	SubtitleUrl      string        `json:"SubtitleUrl"`
	FirstTs          int64         `json:"FirstTs"`
	FirstTranscode   int64         `json:"FirstTranscode"`
	DeviceInfoStatus int8          `json:"DeviceInfoStatus"` //0=waiting,-1=failed,1=success
}

type GetDeviceResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DeviceView `json:"Result,omitempty"`
}

type ListDevicesResult struct {
	PageResult
	Devices []*DeviceView
}

type ListDevicesResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListDevicesResult `json:"Result,omitempty"`
}

type GetDeviceChannelsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           DeviceChannelResp `json:"Result,omitempty"`
}

type DeviceChannelResp struct {
	DeviceName string                  `json:"DeviceName"`
	DeviceID   string                  `json:"DeviceID"`
	DeviceNSID string                  `json:"DeviceNSID"`
	Channels   []*DeviceItemWithStream `json:"Channels"`
}
type DeviceItemWithStream struct {
	DeviceItem DeviceItem `json:"DeviceItem"`
	StreamID   string     `json:"StreamID"`
	CreateAt   string     `json:"CreateAt"`
	SubStreams []string   `json:"SubStreams"`
}

type CatalogItemInfo struct {
	PTZType       string `json:"PTZType,omitempty"` // 1-球机;2-半球机;3-固定枪机;4-遥控枪机
	Resolution    string `json:"Resolution,omitempty"`
	DownloadSpeed string `json:"DownloadSpeed,omitempty"`
}

type DeviceItem struct {
	DeviceID     string          `json:"DeviceID"`
	Event        string          `json:"Event,omitempty"`
	Name         string          `json:"Name"`
	Manufacturer string          `json:"Manufacturer,omitempty"`
	Model        string          `json:"Model,omitempty"`
	Owner        string          `json:"Owner,omitempty"`
	CivilCode    string          `json:"CivilCode,omitempty"`
	Address      string          `json:"Address,omitempty"`
	Parental     string          `json:"Parental,omitempty"`
	ParentID     string          `json:"ParentID,omitempty"`
	RegisterWay  string          `json:"RegisterWay,omitempty"`
	Secrecy      string          `json:"Secrecy,omitempty"`
	StreamNum    string          `json:"StreamNum,omitempty"`
	IPAddress    string          `json:"IPAddress,omitempty"`
	Port         string          `json:"Port,omitempty"`
	Password     string          `json:"Password,omitempty"`
	Status       string          `json:"Status"`
	Info         CatalogItemInfo `json:"Info,omitempty"`
}

type PresetList struct {
	Num   int          `json:"Num"`
	Items []PresetItem `json:"Items"`
}

type PresetItem struct {
	PresetID   int    `json:"PresetID"`
	PresetName string `json:"PresetName"`
}

type QueryPresetInfoResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           PresetList `json:"Result,omitempty"`
}

type RecordList struct {
	Num   int          `json:"Num"`
	Items []RecordItem `json:"Items"`
}

type RecordItem struct {
	DeviceID       string `json:"DeviceID"`
	Name           string `json:"Name"`
	StartTime      string `json:"StartTime"`
	EndTime        string `json:"EndTime"`
	StartTimestamp int64  `json:"StartTimeStamp"`
	EndTimestamp   int64  `json:"EndTimeStamp"`
	Secrecy        string `json:"Secrecy"`
	Type           string `json:"Type"`
	FileSize       string `json:"FileSize"`
}

type GetRecordListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           RecordList `json:"Result,omitempty"`
}

type PlayBackStartResp struct {
	StreamID string   `json:"StreamID"`
	PullUrls []string `json:"PullUrls"`
	PushUrl  string   `json:"PushUrl"`
}

type PlayBackStartResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           PlayBackStartResp `json:"Result,omitempty"`
}

type StreamResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type GetStreamResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           StreamResp `json:"Result,omitempty"`
}

type StreamResp struct {
	*StreamView
	Logs LogViews
}

type LogViews []*LogView
type LogView struct {
	Message string `json:"Message"`
}

type StreamView struct {
	AccountID       string              `json:"AccountID,omitempty"`
	SpaceID         string              `json:"SpaceID"`
	SIPID           string              `json:"SIPID"`
	StreamName      string              `json:"StreamName"`
	StreamID        string              `json:"StreamID"`
	SpaceAccessType string              `json:"SpaceAccessType"`
	DeviceID        string              `json:"DeviceID"`
	DeviceNSID      string              `json:"DeviceNSID"`
	ChannelID       string              `json:"ChannelID"`
	Status          string              `json:"Status"`
	Description     string              `json:"Description"`
	Screenshot      StreamTemplate      `json:"Screenshot"`
	Record          StreamTemplate      `json:"Record"`
	AI              StreamTemplate      `json:"AI"`
	CreatedAt       string              `json:"CreatedAt"`
	UpdatedAt       string              `json:"UpdatedAt"`
	PushUrlDDL      int64               `json:"PushUrlDDL,omitempty"`
	PushUrl         string              `json:"PushUrl,omitempty"`
	PullUrls        []string            `json:"PullUrls"`
	TransPullUrls   map[string][]string `json:"TransPullUrls"`
	FailedTimes     int                 `json:"FailedTimes"`
	RecentPushTs    string              `json:"RecentPushTs"`    //最近推流时间
	StreamLogs      LogViews            `json:"StreamLogs"`      //stream日志
	RtpTransportTcp bool                `json:"RtpTransportTcp"` //rtp传输协议是否使用tcp
	DemandLive      bool                `json:"DemandLive"`
	Resolution      string              `json:"Resolution"` // 通道支持的分辨率列表, 分辨率列表由‘/’隔开，国标协议样例: 6/3
}

type StreamTemplate struct {
	TemplateName string `json:"-"`
	TemplateID   string `json:"TemplateID"`
}

type ListStreamsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListStreamsResult `json:"Result,omitempty"`
}

type GetStreamDataResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GetStreamDataRes `json:"Result,omitempty"`
}

type GetStreamDataRes struct {
	BVideo []DataProject
	BAudio []DataProject
	FPS    []DataProject
	Height int
	Width  int
}

type StreamLogsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListStreamRecordsResult `json:"Result,omitempty"`
}

type ListStreamRecordsResult struct {
	PageResult
	StreamRecords []*StreamRecord
}

type StreamRecord struct {
	RecordID  string `bson:"RecordID"`
	StreamID  string `bson:"StreamID"`
	StartTime int    `bson:"StartTime"`
	EndTime   int    `bson:"EndTime"`
	Duration  int    `bson:"Duration"`
}

type ListStreamsResult struct {
	PageResult
	Streams []*StreamResp
}

type GetSpaceTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SpaceTemplate `json:"Result,omitempty"`
}

type SpaceTemplate struct {
	Screenshot SimpleTemplate   `json:"Screenshot"`
	Record     SimpleTemplate   `json:"Record"`
	AI         SimpleTemplate   `json:"AI"`
	Trans      []SimpleTemplate `bson:"Trans"`
}

type TemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type GetTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           Template `json:"Result,omitempty"`
}

type ListTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListTemplatesResult `json:"Result,omitempty"`
}

type ListTemplatesResult struct {
	PageResult
	Templates []*Template
}

type Template struct {
	AccountID    string              `json:"-"`
	TemplateID   string              `json:"TemplateID"`
	TemplateName string              `json:"TemplateName"`
	TemplateType string              `json:"TemplateType"`
	Screenshot   *ScreenshotTemplate `json:"Screenshot,omitempty"`
	Record       *RecordTemplate     `json:"Record,omitempty"`
	AI           *AITemplate         `json:"AI,omitempty"`
	Trans        *TransTemplate      `json:"Trans,omitempty"`
	CreatedAt    string              `json:"CreatedAt"`
	UpdatedAt    string              `json:"UpdatedAt"`
	Status       string              `json:"-"`
}

type ScreenshotTemplate struct {
	ScreenshotPeriod int                  `json:"ScreenshotPeriod"` // ms
	Type             []ScreenshotType     `json:"Type"`             //多种截图方式
	Bucket           TemplateBucketConfig `json:"-"`
	TTL              TemplateTTLConfig    `json:"TTL"`
}
type TemplateBucketConfig struct {
	Region    string `json:"Region"`
	Name      string `json:"Name"`
	Path      string `json:"Path"`
	AccessKey string `json:"AccessKey"`
}
type TemplateTTLConfig struct {
	Days int `json:"Days"`
}

type ScreenshotType string

type ScreenshotTypeInt int32

type RecordTemplate struct {
	RecordDuration int                  `bson:"RecordDuration"` // ms
	SliceDuration  int                  `bson:"SliceDuration"`
	TTL            TemplateTTLConfig    `bson:"TTL"`
	Type           string               `bson:"Type"`
	Format         string               `bson:"Format"`
	Path           string               `bson:"Path" json:"-"`
	Bucket         TemplateBucketConfig `bson:"Bucket"`
	EnableTimes    *EnableTimes         `bson:"EnableTimes"`
}

type AITemplate struct {
	CapabilitySet []string                  `json:"CapabilitySet"`
	TemplateItems map[string]AITemplateItem `json:"TemplateItems"`
	ImageTTL      int                       `json:"ImageTTL"`
	Path          string                    `json:"-"`
	Bucket        TemplateBucketConfig      `json:"-"`
}
type AITemplateItem struct {
	CapabilityType      string                 `json:"CapabilityType"`
	ConfidenceThreshold float64                `json:"ConfidenceThreshold"`
	FrameInterval       int                    `json:"FrameInterval"`
	Extra               map[string]interface{} `json:"Extra"`
	Status              bool                   `json:"Status"`
	GrabStrategy        string                 `json:"GrabStrategy"`
	EnableTimes         EnableTimes            `json:"EnableTimes"`
}
type EnableTime struct {
	EnableFrom string
	EnableTo   string
}
type EnableTimes []EnableTime

type TransTemplate struct {
	SuffixName   *string `json:"SuffixName"`
	VideoBitrate *int    `json:"VideoBitrate,omitempty"`
	Vcodec       *string `json:"Vcodec"`
	AudioBitrate *int    `json:"AudioBitrate,omitempty"`
	Acodec       *string `json:"Acodec"`
	FPS          *int    `json:"FPS,omitempty"`
	GOP          *int    `json:"GOP"`
	Width        *int    `json:"Width"`
	Height       *int    `json:"Height"`
	As           *string `json:"As,omitempty"`
	ShortSide    *int    `json:"ShortSide,omitempty"`
	LongSide     *int    `json:"LongSide,omitempty"`
	BFrames      *int    `json:"BFrames,omitempty"`
	Roi          *bool   `json:"Roi"`
}

type ListHistoryResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []*HistoryResult `json:"Result,omitempty"`
}

type HistoryResult struct {
	Type         string        `json:"Type"`
	Path         string        `json:"Path"`
	ScreenResult *ScreenResult `json:"Screenshot,omitempty"`
	RecordResult *RecordResult `json:"Record,omitempty"`
	RecordMeta   *RecordMeta   `json:"RecordMeta,omitempty"`
}

type RecordMeta struct {
	StartTs     string  `json:"StartTs"`
	EndTs       string  `json:"EndTs"`
	Duration    float64 `json:"Duration"`
	Format      string  `json:"Format"`
	Bucket      string  `json:"Bucket"`
	StorageType string  `json:"StorageType"`
}

type ScreenResult struct {
	BornTs string `json:"BornTs"`
}

type ListHistoryResponseV3 struct {
	ResponseMetadata base.ResponseMetadata
	Result           []*HistoryResultV3 `json:"Result,omitempty"`
}

type HistoryResultV3 struct {
	Type         string          `json:"Type"`
	Path         string          `json:"Path"`
	ScreenResult *ScreenResultV3 `json:"Screenshot,omitempty"`
	RecordResult *RecordResultV3 `json:"Record,omitempty"`
	RecordMeta   *RecordMetaV3   `json:"RecordMeta,omitempty"`
}

type RecordResultV3 struct {
	StartTime int64   `json:"StartTime"`
	EndTime   int64   `json:"EndTime"`
	Duration  float64 `json:"Duration"`
	Cover     string  `json:"Cover"`
	Format    string  `json:"Format"`
}

type RecordMetaV3 struct {
	StartTime   int64   `json:"StartTime"`
	EndTime     int64   `json:"EndTime"`
	Duration    float64 `json:"Duration"`
	Format      string  `json:"Format"`
	Bucket      string  `json:"Bucket"`
	StorageType string  `json:"StorageType"`
}

type ScreenResultV3 struct {
	BornTs int64 `json:"BornTs"`
}

type RecordResult struct {
	StartTs  string  `json:"StartTs"`
	EndTs    string  `json:"EndTs"`
	Duration float64 `json:"Duration"`
	Cover    string  `json:"Cover"`
	Format   string  `json:"Format"`
}

type CascadeSipConfig struct {
	SipServerID     string `json:"SipServerID"`     // SIP服务器ID
	Realm           string `json:"Realm"`           // SIP服务器域
	SipServerHost   string `json:"SipServerHost"`   // SIP服务器地址
	SipServerPort   int    `json:"SipServerPort"`   // SIP服务器端口
	SipUserID       string `json:"SipUserID"`       // 平台SIP用户名
	Password        string `json:"Password"`        // 平台SIP密码
	SipTransport    string `json:"SipTransport"`    // SIP信令传输协议, udp/tcp; default:udp
	RegisterExpires uint32 `json:"RegisterExpires"` // 注册有效期(秒)
	KeepalivePeriod uint32 `json:"KeepalivePeriod"` // 心跳周期(秒)
}

type CreateCascadePlatformResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type UpdateCascadePlatformResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type CascadePlatform struct {
	AccountID      string           `json:"AccountID"`
	PlatformID     string           `json:"PlatformID"`
	SipConfig      CascadeSipConfig `json:"SipConfig"`
	PlatformName   string           `json:"PlatformName"`   // 上级平台名称
	Description    string           `json:"Description"`    // 上级平台描述
	Enable         bool             `json:"Enable"`         // 是否使能级联平台
	EnableAccess   bool             `json:"EnableAccess"`   // 平台授权
	EnablePTZ      bool             `json:"EnablePTZ"`      // PTZ授权
	Status         CascadeStatus    `json:"Status"`         // 注册状态
	SharedNodesNum int              `json:"SharedNodesNum"` // 共享节点数量
	TaskIDs        []string         `json:"TaskIDs"`        // 关联的级联任务列表
	CreateAt       string           `json:"CreateAt"`
	UpdateAt       string           `json:"UpdateAt"`
}

type GetCascadePlatformResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           CascadePlatform `json:"Result,omitempty"`
}

type ListCascadePlatformResult struct {
	PageResult
	CascadePlatforms []*CascadePlatform `json:"CascadePlatforms"`
}

type ListCascadePlatformResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListCascadePlatformResult `json:"Result,omitempty"`
}

type CreateCascadeTaskResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type ListCascadeTaskResult struct {
	PageResult
	CascadeTasks []*CascadeTask `json:"CascadeTasks"`
}

type ListCascadeTaskResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListCascadeTaskResult `json:"Result,omitempty"`
}

type CascadeTask struct {
	AccountID     string        `json:"AccountID"`
	TaskID        string        `json:"TaskID"`        // 级联任务ID, 2位数字
	TaskName      string        `json:"TaskName"`      // 级联任务名称
	Description   string        `json:"Description"`   // 级联任务描述
	GroupTreeID   string        `json:"GroupTreeID"`   // 关联的虚拟组织树ID
	GroupTreeName string        `json:"GroupTreeName"` // 关联的虚拟组织树名称
	PlatformID    string        `json:"PlatformID"`    // 关联的上级平台ID
	PlatformName  string        `json:"PlatformName"`  // 关联的上级平台名称
	Status        CascadeStatus `json:"Status"`        // 级联任务状态
	Enable        bool          `json:"Enable"`        // 是否启用
	SharedNodes   []string      `json:"SharedNodes"`   // 共享节点ID列表
	CreateAt      string        `json:"CreateAt"`
	UpdateAt      string        `json:"UpdateAt"`
}

type GetCascadeTaskResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           CascadeTask `json:"Result,omitempty"`
}

type GetGroupTreeNodeResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GroupTreeNode `json:"Result,omitempty"`
}

type SlowLiveResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type GetSlowLiveResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GetSlowLiveResult `json:"Result,omitempty"`
}

type GetSlowLiveResult struct {
	SlowLive *SlowLiveView `json:"SlowLive"`
}

type SlowLiveView struct {
	ID                      string          `json:"ID"`
	RelatedSpaces           []string        `json:"RelatedSpaces"`
	Name                    string          `json:"Name"`
	MergingMode             string          `json:"MergingMode"`
	OutputResolution        string          `json:"OutputResolution"`
	OutputFrameRate         int             `json:"OutputFrameRate"`
	OutputEncoding          string          `json:"OutputEncoding"`
	CreateTimestamp         int64           `json:"CreateTimestamp"`
	PreviousClosedTimestamp int64           `json:"PreviousClosedTimestamp"`
	FinishTimestamp         int64           `json:"FinishTimestamp"`
	Status                  SlowLiveStatus  `json:"Status"`
	EnablePushingStream     *bool           `json:"EnablePushingStream"`
	Config                  *SlowLiveConfig `json:"Config"`
	AccountID               string          `json:"AccountID"`
}

type ListSlowLiveResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListSlowLiveResult `json:"Result,omitempty"`
}

type ListSlowLiveResult struct {
	PageResult
	SlowLiveStreamings []*SlowLiveView `json:"SlowLiveStreamings"`
}

type SignSlowliveWsTokenResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SignSlowliveWsTokenResult `json:"Result,omitempty"`
}

type SignSlowliveWsTokenResult struct {
	Token      string `json:"Token"`
	ExpireAt   int64  `json:"ExpireAt"`
	SlowliveID string `json:"SlowliveID"`
}

type ListGroupTreeNodesResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListGroupTreeNodesRes `json:"Result,omitempty"`
}

type GroupTreeNode struct {
	ID          string           `json:"ID"`
	Name        string           `json:"Name"`
	NSID        string           `json:"NSID"`
	Description string           `json:"Description"`
	Shared      bool             `json:"Shared"` // only used to mark node share flag for cascade task
	SpaceID     string           `json:"SpaceID"`
	SpaceName   string           `json:"SpaceName"`
	Region      string           `json:"Region"`
	ParentID    string           `json:"ParentID"`
	Nodes       []*GroupTreeNode `json:"Nodes"`
	Device      *GroupDevice     `json:"Device"`
}

type StartVoiceTalkResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           StartVoiceTalkResp
}

type StartVoiceTalkResp struct {
	AudioSendUrl string `json:"AudioSendUrl"`
}

type StopVoiceTalkResponse struct {
	ResponseMetadata base.ResponseMetadata
}

type GroupDevice struct {
	ID          string         `json:"ID"`
	Type        string         `json:"Type"`
	DeviceName  string         `json:"DeviceName"`
	Description string         `json:"Description"`
	Status      string         `json:"Status"`
	DeviceNSID  string         `json:"DeviceNSID"`
	Protocol    string         `json:"Protocol"`
	Vendor      string         `json:"Vendor"`
	Channels    []*GroupStream `json:"Channels"`
	Available   bool           `json:"Available"`
}

type GroupStream struct {
	ID          string `json:"ID"`
	SipID       string `json:"SipID"`
	StreamName  string `json:"StreamName"`
	DeviceID    string `json:"DeviceID"`
	Status      string `json:"Status"`
	Description string `json:"Description"`
	Available   bool   `json:"Available"`
}

type ListGroupTreeNodesRes struct {
	PageNumber int64            `json:"PageNumber"`
	PageSize   int64            `json:"PageSize"`
	TotalCount int64            `json:"TotalCount"`
	Nodes      []*GroupTreeNode `json:"Nodes"`
}

type AddGroupTreeNodeResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string `json:"Result,omitempty"`
}

type GetDevicesByGroupTreeNodeResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GetDevicesByGroupTreeNodeRes `json:"Result,omitempty"`
}

type GetDevicesByGroupTreeNodeRes struct {
	PageNumber int64          `json:"PageNumber"`
	PageSize   int64          `json:"PageSize"`
	TotalCount int64          `json:"TotalCount"`
	Data       []*GroupDevice `json:"Data"`
}

type StructuredViewResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type StructuredViewSpace struct {
	ID              string                    `json:"ID"`
	AccountID       string                    `json:"AccountID"`
	SpaceName       string                    `json:"SpaceName"`
	ViewNumber      int64                     `json:"ViewNumber"`
	StoragePeriod   int                       `json:"StoragePeriod"`
	StorageType     string                    `json:"StorageType"`
	Region          string                    `json:"Region"`
	Description     string                    `json:"Description"`
	SpaceCode       string                    `json:"SpaceCode"`
	SpaceIP         string                    `json:"SpaceIP"`
	SpacePort       int                       `json:"SpacePort"`
	CreateTimestamp int64                     `json:"CreateTimestamp"`
	UpdateTimestamp int64                     `json:"UpdateTimestamp"`
	Status          StructuredViewSpaceStatus `json:"Status"`
}

type GetStructuredViewSpaceResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           StructuredViewSpace `json:"Result,omitempty"`
}

type GetStructuredViewResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           StructuredView `json:"Result,omitempty"`
}

type StructuredView struct {
	ID               string               `json:"ID"`
	AccountID        string               `json:"AccountID"`
	ViewName         string               `json:"ViewName"`
	ViewType         string               `json:"ViewType"`
	ViewCode         string               `json:"ViewCode"`
	ViewSpaceID      string               `json:"ViewSpaceID"`
	ViewSpaceName    string               `json:"ViewSpaceName"`
	ViewIP           string               `json:"ViewIP"`
	ViewPort         int                  `json:"ViewPort"`
	Protocol         string               `json:"Protocol"`
	Username         string               `json:"Username"`
	Password         string               `json:"Password"`
	Description      string               `json:"Description"`
	Location         string               `json:"Location"`
	AdministrativeID string               `json:"AdministrativeID"`
	CreateTimestamp  int64                `json:"CreateTimestamp"`
	UpdateTimestamp  int64                `json:"UpdateTimestamp"`
	Status           StructuredViewStatus `bson:"Status" json:"Status"`
}

type StructuredViewStatus string

type ListStructuredViewSpacesResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListStructuredViewSpaceResult `json:"Result,omitempty"`
}

type ListStructuredViewsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListStructuredViewResult `json:"Result,omitempty"`
}

type ListStructuredViewResult struct {
	PageResult
	StructuredView []*StructuredView `json:"ViewList"`
}

type ListStructuredViewDataResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListStructuredViewDataResult `json:"Result,omitempty"`
}

type ListStructuredViewDataResult struct {
	PageResult
	StructuredViewData []*StructuredViewData `json:"DataList"`
}

type StructuredViewData struct {
	ID           string `json:"ID"`
	ViewDataType string `json:"ViewDataType"`
	Timestamp    int64  `json:"Timestamp"`
	Url          string `json:"Url"`
}

type PhotoObjectDetail struct {
	PhotoID             string `json:"PhotoID"`
	InfoKind            string `json:"InfoKind"` //人工采集还是自动采集
	ImageSource         string `json:"ImageSource"`
	StoragePath         string `json:"StoragePath"`
	Type                string `json:"Type"` //"11"人脸小图， "14"场景大图
	Title               string `json:"Title"`
	FileFormat          string `json:"FileFormat"`
	ShotTime            string `json:"ShotTime"`
	ContentDescription  string `json:"ContentDescription"`  //对图像内容的简要描述
	ShotPlaceFullAdress string `json:"ShotPlaceFullAdress"` //拍摄地点区划内详细地址
	SecurityLevel       string `json:"SecurityLevel"`
	Width               int    `json:"Width"`
	Height              int    `json:"Height"`
	CreateTime          string `json:"CreateTime"`
	DownloadUrl         string `json:"DownloadUrl"`
	DeviceID            string `json:"device_id"`
}

type GetStructuredViewDataInfoResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           PhotoObjectDetail `json:"Result,omitempty"`
}

type StructuredViewCascadePlatform struct {
	ID              string                              `json:"ID"`
	AccountID       string                              `json:"AccountID"`
	PlatformName    string                              `json:"PlatformName"`
	PlatformCode    string                              `json:"PlatformCode"`
	PlatformIP      string                              `json:"PlatformIP"`
	PlatformPort    int                                 `json:"PlatformPort"`
	Description     string                              `json:"Description"`
	Username        string                              `json:"Username"`
	Password        string                              `json:"Password"`
	CreateTimestamp int64                               `json:"CreateTimestamp"`
	UpdateTimestamp int64                               `json:"UpdateTimestamp"`
	Status          StructuredViewCascadePlatformStatus `json:"Status"`
}

type GetStructuredViewCascadePlatformResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           StructuredViewCascadePlatform `json:"Result,omitempty"`
}

type ListStructuredViewCascadePlatformResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListStructuredViewCascadePlatformResult `json:"Result,omitempty"`
}

type ListStructuredViewCascadePlatformResult struct {
	PageResult
	PlatformList []*StructuredViewCascadePlatform `json:"PlatformList"`
}

type StructuredViewCascadeJob struct {
	ID              string                         `json:"ID"`
	AccountID       string                         `json:"AccountID"`
	JobName         string                         `json:"JobName"`
	ViewSpaceID     string                         `json:"ViewSpaceID"`
	ViewSpaceName   string                         `json:"ViewSpaceName"`
	ViewSpaceCode   string                         `json:"ViewSpaceCode"`
	PlatformID      string                         `json:"PlatformID"`
	PlatformCode    string                         `json:"PlatformCode"`
	PlatformName    string                         `json:"PlatformName"`
	CreateTimestamp int64                          `json:"CreateTimestamp"`
	UpdateTimestamp int64                          `json:"UpdateTimestamp"`
	Status          StructuredViewCascadeJobStatus `json:"Status"`
	Description     string                         `json:"Description"`
}

type ListStructuredViewCascadeJobResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListStructuredViewCascadeJobResult `json:"Result,omitempty"`
}

type ListStructuredViewCascadeJobResult struct {
	PageResult
	JobList []*StructuredViewCascadeJob `json:"JobList"`
}

type GetStructuredViewCascadeJobResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           StructuredViewCascadeJob `json:"Result,omitempty"`
}

type ListStructuredViewSpaceResult struct {
	PageResult
	StructuredViewSpace []*StructuredViewSpace `json:"ViewSpaceList"`
}

type Permissions struct {
	Charge *Charge `json:"Charge"`
}

type Charge struct {
	Bandwidth bool `json:"bandwidth"`
	Flow      bool `json:"flow"`
	Channel   bool `json:"channel"`
}

type ChargeMode string

type AccountStatus struct {
	Status      int          `json:"Status"`
	CreateTime  int64        `json:"CreateTime"`
	ChargeMode  ChargeMode   `json:"ChargeMode"`
	Permissions *Permissions `json:"Permissions"`
}

type GetAccountResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *AccountStatus `json:"Result,omitempty"`
}

type CreateAccountResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CreateAccountRes `json:"Result,omitempty"`
}

type CreateAccountRes struct {
	PreorderNumber string `json:"PreorderNumber"`
}

type NewPackageResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CreateAccountRes `json:"Result,omitempty"`
}

type CalculatePkgResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CalculatePkgRes `json:"Result,omitempty"`
}

type CalculatePkgRes struct {
	Price float64 `json:"Price"`
}

type ChargeItem struct {
	Name          string           `json:"Name"`
	ValidDur      string           `json:"ValidDur"`
	Region        string           `json:"Region"`
	Types         []*TypeItem      `json:"Types"`
	Specification []*Specification `json:"Specification"`
}

type TypeItem struct {
	Name  string
	Extra string
}

type Specification struct {
	Name  string  `json:"Name"`
	Type  string  `json:"Type"`
	Code  string  `json:"Code"`
	Price float64 `json:"Price"`
}

type GetPkgInfoResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []*ChargeItem `json:"Result,omitempty"`
}

type ListSpaceDomainsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListSpaceDomainsRes `json:"Result,omitempty"`
}

type ListSpaceDomainsRes struct {
	Domains interface{} `json:"Domains"`
}

type DeleteSpaceDomainResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           DeleteSpaceDomainRes `json:"Result,omitempty"`
}

type DeleteSpaceDomainRes struct {
	Domain interface{} `json:"Domain"`
}

type GetSpaceDomainResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SpaceDomain `json:"Result,omitempty"`
}

type UpdateSpaceDomainResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           map[string]SpaceDomain `json:"Result,omitempty"`
}

type GetDataProject struct {
	TotalUp   float64
	TotalDown float64
	DataUp    []DataProject
	DataDown  []DataProject
}

type DataProject struct {
	TimeStamp string  `json:"TimeStamp"`
	Value     float64 `json:"Value"`
}

type GetDataProjectWithBindWidthAndFlowResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GetDataProject `json:"Result,omitempty"`
}

type GetTotalDataRes struct {
	MaxUp           float64
	MaxDown         float64
	UpChainGrowth   string
	DownChainGrowth string
}

type GetTotalDataResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GetTotalDataRes `json:"Result,omitempty"`
}

type GetPushStreamCntResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GetPushStreamCntRes `json:"Result,omitempty"`
}

type GetPushStreamCntRes struct {
	Cnt []DataProject
}

type ListAlarmNotifyResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListAlarmNotifyResult `json:"Result,omitempty"`
}

type AlarmNotify struct {
	AccountID        string `json:"AccountID"`
	SpaceID          string `json:"SpaceID"`
	DeviceNSID       string `json:"DeviceNSID"`
	CmdType          string `json:"CmdType"`
	SN               uint32 `json:"SN"`
	ChannelID        string `json:"ChannelID"`
	AlarmPriority    uint32 `json:"AlarmPriority"`
	AlarmMethod      uint32 `json:"AlarmMethod"`
	AlarmTime        int64  `json:"AlarmTime"`
	AlarmDescription string `json:"AlarmDescription"`
	Longitude        string `json:"Longitude"`
	Latitude         string `json:"Latitude"`
	AlarmInfo        string `json:"AlarmInfo"`
	AlarmNotifyID    string `json:"AlarmNotifyID"`
	Info             struct {
		AlarmType      uint32 `json:"AlarmType"`
		AlarmTypeParam struct {
			EventType uint32 `json:"EventType"`
		} `bson:"AlarmTypeParam"`
	} `bson:"Info"`
}

type ListAlarmNotifyResult struct {
	PageResult
	AlarmNotifies []*AlarmNotify
}

type UploadCertResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           map[string]string `json:"Result,omitempty"`
}

type ListCertificatesResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListCertsResult `json:"Result,omitempty"`
}

type SimpleCertInfo struct {
	ChainID   string `json:"ChainID"`
	CertName  string `json:"CertName"`
	Status    string `json:"Status"`
	Domain    string `json:"Domain"`
	CreateAt  int64  `json:"CreateAt"`
	UpdateAt  int64  `json:"UpdateAt"`
	NotBefore int64  `json:"NotBefore"`
	NotAfter  int64  `json:"NotAfter"`
}

type ListCertsResult struct {
	PageResult
	Certs []*SimpleCertInfo
}

type CertDetail struct {
	ChainID   string `bson:"ChainID"`
	Status    string `bson:"Status"`
	Domain    string `bson:"Domain"`
	CreateAt  int64  `bson:"CreateAt"`
	UpdateAt  int64  `bson:"UpdateAt"`
	NotBefore int64  `bson:"NotBefore"`
	NotAfter  int64  `bson:"NotAfter"`
	CertName  string `bson:"CertName"`
	CertDesc  string `bson:"CertDesc"`
	CertPem   string `bson:"CertPem"`
}

type CertDetailResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           CertDetail `json:"Result,omitempty"`
}

type EdgeResponseDetail struct {
	Status bool
}

type EdgeResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           EdgeResponseDetail `json:"Result,omitempty"`
}

// ai releated model
type ParseRecord struct {
	AccountID   string        `json:"AccountID" bson:"AccountID"` // 账号ID
	LibID       string        `json:"LibID" bson:"LibID"`         // 车辆库ID
	ItemID      string        `json:"ItemID" bson:"ItemID"`       // 车辆库车辆 ID
	AITaskID    string        `json:"AITaskID" bson:"AITaskID"`   // AI 任务ID
	Plate       string        `json:"Plate" bson:"Plate"`         // 车牌
	Ts          time.Duration `json:"Ts" bson:"Ts"`               // 时间戳
	TosKey      string        `json:"TosKey" bson:"TosKey"`       // 车辆截图
	Score       float32       `json:"Score" bson:"Score"`         // AI识别打分
	DownloadUrl string        `json:"DownloadUrl,omitempty"`      // 存储网关下载地址
}

type RectParams struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type VehicleRecord struct {
	ObjectID       int          `json:"ObjectID,omitempty"`   // 识别ObjectID
	Plate          string       `json:"Plate,omitempty"`      // 车牌
	Brand          string       `json:"Brand,omitempty"`      // 品牌
	Type           string       `json:"Type,omitempty"`       // 车辆类型
	Color          string       `json:"Color,omitempty"`      // 车颜色
	Status         string       `json:"Status,omitempty"`     // 车损状态
	RectParams     *RectParams  `json:"RectParams,omitempty"` // 车辆坐标
	Score          float32      `json:"Score,omitempty"`      // 打分
	CarParseRecord *ParseRecord `json:"ParseRecord,omitempty"`
}

type FaceRecord struct {
}

type QualityDetection struct {
	AICapabilityTypeBlackout     *int `json:"blackout"`
	AICapabilityTypeHinder       *int `json:"hinder"`
	AICapabilityTypeOverexposure *int `json:"overexposure"`
	AICapabilityTypeBlur         *int `json:"blur"`
	AICapabilityTypeNoise        *int `json:"noise"`
}

type SegInfo struct {
	SegUrl     string           `json:"SegUrl"`
	RectParams RectParams       `json:"RectParams"`
	Vehicle    VehicleRecord    `json:"Vehicle"`
	Face       FaceRecord       `json:"Face"`
	VQ         QualityDetection `json:"VQ"`
}

type AIAnalysisResponse struct {
	SegInformation []*SegInfo `json:"SegInformation"`
	OriUrl         string     `json:"Url"`
	Time           int64      `json:"Time"`
}

type GetAIAnalysisResult struct {
	PageResult
	AiResult []*AIAnalysisResponse
}

type GetAIAnalysisResultResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           GetAIAnalysisResult `json:"Result,omitempty"`
}

type RecordPlanView struct {
	PlanID       string `bson:"PlanID"`
	PlanName     string `bson:"PlanName"`
	Description  string `bson:"Description"`
	Status       string `bson:"Status"`
	CreatedAt    int64  `bson:"CreatedAt"`
	UpdatedAt    int64  `bson:"UpdatedAt"`
	BindTemplate string `bson:"BindTemplate"`
}

type RecordPlanResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *RecordPlanView
}

type BindChannel struct {
	ChannelID string `json:"ChannelID"`
}

type CreateRecordPlanRequest struct {
	PlanName       string                    `json:"PlanName"`
	Description    string                    `json:"Description"`
	Status         string                    `json:"Status"`
	BindTemplate   string                    `json:"BindTemplate"`
	StreamingIndex int                       `json:"StreamingIndex"`
	Resolution     string                    `json:"Resolution"`
	BindChannels   map[string][]*BindChannel `json:"BindChannels"`
	BindStreams    []string                  `json:"BindStreams"`
}

type UpdateRecordPlanList struct {
	StreamingIndex int                       `json:"StreamingIndex"`
	Resolution     string                    `json:"Resolution"`
	Devices        map[string][]*BindChannel `json:"Devices"`
	Streams        []string                  `json:"Streams"`
}

type UpdateRecordPlanRequest struct {
	PlanID       string                `json:"PlanID"`
	PlanName     string                `json:"PlanName"`
	Description  string                `json:"Description"`
	BindTemplate string                `json:"BindTemplate"`
	Status       string                `json:"Status"`
	AddList      *UpdateRecordPlanList `json:"AddList"`
	DelList      *UpdateRecordPlanList `json:"DelList"`
}

type ListRecordPlanResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListRecordPlanResp
}

type ListRecordPlanResp struct {
	PageResult PageResult        `json:"PageResult"`
	List       []*RecordPlanView `json:"List"`
}

type StreamingType = string //码流类型

type RecordPlanStreamView struct {
	DeviceNSID    string        `json:"DeviceNSID"`
	ChannelID     string        `json:"ChannelID"`
	StreamingType StreamingType `json:"StreamingType"`
	StreamID      string        `json:"StreamID"`
	StreamName    string        `json:"StreamName"`
	PullUrls      []string      `json:"PullUrls"`
}

type ListRecordPlanChannelResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListRecordPlanChannelResp
}

type ListRecordPlanChannelResp struct {
	PageResult
	List map[string][]*RecordPlanStreamView `json:"List"`
}

type RecordPlanResultResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           IDResponse `json:"Result,omitempty"`
}

type CruisePoint struct {
	PresetID   uint8  `json:"PresetID" bson:"PresetID"` // 预置点编号(1~255)
	PresetName string `json:"PresetName" bson:"PresetName"`
}

type CruiseTrack struct {
	DeviceNSID  string        `json:"DeviceNSID"`
	ChannelID   string        `json:"ChannelID"`
	TrackID     uint8         `json:"TrackID"`     // 巡航轨迹组编号(1~255)
	TrackList   []CruisePoint `json:"TrackList"`   // 巡航轨迹列表, 最多可添加255个预置点
	StaySeconds uint32        `json:"StaySeconds"` // 预置点停留时间, 单位:秒, 取值范围:1~4095
	Speed       uint32        `json:"Speed"`       // 巡航速度, 取值范围: 1~4095
}

type SetCruiseTrackArgs struct {
	DeviceNSID  string        `json:"DeviceNSID"`
	ChannelID   string        `json:"ChannelID"`
	TrackID     uint8         `json:"TrackID"`     // 巡航轨迹组编号(1~255)
	TrackList   []CruisePoint `json:"TrackList"`   // 巡航轨迹列表, 最多可添加255个预置点
	StaySeconds uint32        `json:"StaySeconds"` // 预置点停留时间, 单位:秒, 取值范围:1~4095
	Speed       uint32        `json:"Speed"`       // 巡航速度, 取值范围: 1~4095
}

type SetCruiseTrackArgsV3 struct {
	StreamID    string        `json:"StreamID"`
	TrackID     uint8         `json:"TrackID"`     // 巡航轨迹组编号(1~255)
	TrackList   []CruisePoint `json:"TrackList"`   // 巡航轨迹列表, 最多可添加255个预置点
	StaySeconds uint32        `json:"StaySeconds"` // 预置点停留时间, 单位:秒, 取值范围:1~4095
	Speed       uint32        `json:"Speed"`       // 巡航速度, 取值范围: 1~4095
}

type DeleteCruiseTrackArgs struct {
	DeviceNSID string `json:"DeviceNSID"`
	ChannelID  string `json:"ChannelID"`
	TrackID    uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type DeleteCruiseTrackArgsV3 struct {
	StreamID string `json:"StreamID"`
	TrackID  uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type StartCruiseTrackArgs struct {
	DeviceNSID string `json:"DeviceNSID"`
	ChannelID  string `json:"ChannelID"`
	TrackID    uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type StartCruiseTrackArgsV3 struct {
	StreamID string `json:"StreamID"`
	TrackID  uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type StopCruiseTrackArgs struct {
	DeviceNSID string `json:"DeviceNSID"`
	ChannelID  string `json:"ChannelID"`
}

type StopCruiseTrackArgsV3 struct {
	StreamID string `json:"StreamID"`
}

type GetCruiseTrackResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           CruiseTrack `json:"Result,omitempty"`
}

type ListCruiseTracksResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []CruiseTrack `json:"Result,omitempty"`
}

type Resource struct {
	Mem  int64 `bson:"Mem"`
	Cpu  int64 `bson:"Cpu"`
	Disk int64 `bson:"Disk"`
}

type NSSArg struct {
	Addr             string            `bson:"Addr"`
	IPv4             string            `bson:"IPv4"`
	IPv6             string            `bson:"IPv6"`
	EnableRecord     bool              `bson:"EnableRecord"`     //true表示录制集群
	EnableRelay      bool              `bson:"EnableRelay"`      //true表示支持relay
	RecordConfigHost string            `bson:"RecordConfigHost"` //aiot_nss heartbeat请求的路径
	ScanDurationSec  int               `bson:"ScanDurationSec"`  //nss扫描的时间间隔
	QuotaKey         string            `bson:"QuotaKey"`         //Quota key
	WantStatus       string            `bson:"WantStatus"`
	Version          string            `bson:"Version"`          //配置的版本
	StatusUpdateTime int64             `bson:"StatusUpdateTime"` //状态更新时间
	OnTime           int64             `bson:"OnTime"`
	Resource         *Resource         `bson:"Resource"`
	Env              map[string]string `bson:"Env"`
}

type NssInfoListResp struct {
	PageResult
	NssInfo []*NSSArg `json:"NssInfo"`
}

type NssInfoListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           NssInfoListResp `json:"Result,omitempty"`
}

type GetDeviceChannelsMode string

const (
	GetDeviceChannelsModeBrief  GetDeviceChannelsMode = "brief"
	GetDeviceChannelsModeDetail GetDeviceChannelsMode = "detail"
)

type getDeviceChannelsOptions struct {
	Mode GetDeviceChannelsMode
}

type GetDeviceChannelsOption func(*getDeviceChannelsOptions)

func WithGetDeviceChannelsMode(mode GetDeviceChannelsMode) GetDeviceChannelsOption {
	return func(c *getDeviceChannelsOptions) {
		c.Mode = mode
	}
}

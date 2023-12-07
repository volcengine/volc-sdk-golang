package aiot

type SDKProductValidDurationEnum int
type SDKGeneralDeviceAVAbilityBitmask int
type SDKNVRAVAbilityEnum int

type CreateSpaceRequest struct {
	SpaceName   string `json:"SpaceName"`
	Description string `json:"Description"`
	Region      string `json:"Region"`
	AccessType  string `json:"AccessType"`
	GB          struct {
		PullOnDemand bool `json:"PullOnDemand"`
	} `json:"GB"`
	SDK struct {
		DeviceType             string                           `json:"DeviceType"`
		ProductValidDuration   SDKProductValidDurationEnum      `json:"ValidDuration"`
		GeneralDeviceAVAbility SDKGeneralDeviceAVAbilityBitmask `json:"GeneralDeviceAVAbility"`
		NVRAVAbility           SDKNVRAVAbilityEnum              `json:"NVRAVAbility"`
		OS                     string                           `json:"OS"`
		ChipManufacture        string                           `json:"ChipManufacture"`
		ChipModel              string                           `json:"ChipModel"`
	} `json:"SDK"`
	HLSLowLatency bool   `json:"HLSLowLatency"`
	CallbackURL   string `json:"CallbackURL"`
}

type SpaceRequest struct {
	SpaceID string
}

type UpdateSpaceRequest struct {
	SpaceID       string
	SpaceName     *string `json:"SpaceName"`
	Description   *string `json:"Description"`
	HLSLowLatency *bool   `json:"HLSLowLatency"`
	GB            struct {
		PullOnDemand *bool `json:"PullOnDemand"`
	} `json:"GB"`
	CallbackURL *string `json:"CallbackURL"`
}

type ListSpacesRequest struct {
	PageNumber int
	PageSize   int
	Order      string
}

type CheckBindTemplateRequest struct {
	TemplateID   string
	TemplateType string
	PageNumber   int
	PageSize     int
}

type CancelBindTemplateRequest struct {
	SpaceID      string
	TemplateType string
}

type SetSpaceTemplateRequest struct {
	SpaceID      string
	TemplateType string

	TemplateID      string   `json:"TemplateID"`
	TemplateTransID []string `json:"TemplateTransID"`
}

type CreateDeviceRequest struct {
	SpaceID               string
	Type                  string             `json:"Type"`
	DeviceName            string             `json:"DeviceName"`
	DeviceNSID            *string            `json:"DeviceNSID"`
	Username              *string            `json:"Username"`
	Password              *string            `json:"Password"`
	Description           string             `json:"Description"`
	AutoPullAfterRegister *bool              `json:"AutoPullAfterRegiter"`
	PasswordLevel         string             `json:"PasswordLevel"`
	RtpTransportTcp       *bool              `json:"RtpTransportTcp"`
	DownloadSpeed         *int               `json:"DownloadSpeed"`
	Location              *string            `json:"Location"`
	Coordinates           *Coordinates       `json:"Coordinates"`
	AlertNotification     *AlertNotification `json:"AlertNotification"`
	UseSubStream          *bool              `json:"UseSubStream"`
}

type DeviceRequest struct {
	DeviceID string
}

type FreshDeviceRequest struct {
	SpaceID  string
	DeviceID string
}

type LocalMediaDownloadRequest struct {
	SpaceID             string
	DeviceID            string `json:"DeviceID"`
	ChannelID           string `json:"ChannelID"`
	StartTime           int64  `json:"StartTime"`
	EndTime             int64  `json:"EndTime"`
	Version             string `json:"Version"`
	MediaProcess        *Mps   `json:"MediaProcess"`
	DownloadSpeedSingle *int   `json:"DownloadSpeedSingle"` //在本地录像下载的时候时候，不填就以设备维度所设置的速度进行下载
	SchedQuotaKey       string `json:"SchedQuotaKey"`       //调度到哪个集群
}

type LocalMediaDownloadRequestV3 struct {
	StreamID      string `json:"StreamID"`
	StartTime     int64  `json:"StartTime"`
	EndTime       int64  `json:"EndTime"`
	MediaProcess  *Mps   `json:"MediaProcess"`
	DownloadSpeed *int   `json:"DownloadSpeedSingle"` //在本地录像下载的时候时候，不填就以设备维度所设置的速度进行下载
	Expire        int    `json:"Expire"`              //自定义文件过期时间
}

type GetLocalMediaDownloadRequest struct {
	ID      string
	SpaceID string
}
type Mps struct {
	M3U8Option *M3U8Option `json:"M3U8Option"`
	Subtitle   *Subtitle   `json:"Subtitle"`
}

type M3U8Option struct {
	ToMp4 *bool
}

type Subtitle struct {
	SubtitleList SubtitleList `json:"SubtitleList"`
	SubtitleSrc  string       `json:"SubtitleSrc"`
	Font         Font         `json:"Font"`
}

type Font struct {
	FontSize     *int    `json:"FontSize"`
	PrimaryColor *string `json:"PrimaryColor"`
	Alignment    *string `json:"Alignment"`
}

type SubtitleList []*SubtitleItem

type SubtitleItem struct {
	Start   int64  `json:"Start"` //millsecond
	End     int64  `json:"End"`   //millsecond
	Content string `json:"Content"`
	Font    *Font  `json:"Font"`
}

type GenSipIDRequest struct {
	DeviceType  string
	SipServerID string
}

type GetDeviceRequest struct {
	SpaceID     string
	DeviceID    string
	SipServerID string
}

type ListDevicesRequest struct {
	SpaceID    string
	DeviceID   string
	DeviceNSID string
	DeviceName string
	PageNumber int
	PageSize   int
	Order      string
}

type UpdateDeviceRequest struct {
	DeviceID             string
	DeviceName           *string            `json:"DeviceName"`
	DeviceNSID           *string            `json:"DeviceNSID"`
	Username             *string            `json:"Username"`
	Password             *string            `json:"Password"`
	Description          *string            `json:"Description"`
	AutoPullAfterRegiter *bool              `json:"AutoPullAfterRegiter"`
	AlertNotification    *AlertNotification `json:"AlertNotification"`
	Type                 string             `json:"Type"`
	RtpTransportTcp      *bool              `json:"RtpTransportTcp"` //流传输协议tcp为true，否则默认udp
	Location             *string            `json:"Location"`        //可以修改设备的地址位置和坐标
	Coordinates          *Coordinates       `json:"Coordinates"`
}

type CloudRecordPlayRequest struct {
	StreamID   string `json:"StreamID"`
	StartTs    int64  `json:"StartTs"`
	EndTs      int64  `json:"EndTs"`
	TokenValid *int   `json:"TokenValid"`
}

type CloudRecordPlayRequestV3 struct {
	StreamID       string `json:"StreamID"`
	StreamingIndex int    `json:"StreamingIndex,omitempty"` // 码流类型(可选参数), 0-主码流;1-子码流1;2-子码流2, 以此类推
	Resolution     string `json:"Resolution,omitempty"`     // 国标分辨率参数(可选参数), 1-QCIF(176x144); 2-CIF(320x288); 3-4CIF(704x576); 4-D1(720x576); 5-720p; 6-1080p; 其他：WxH表示
	StartTime      int64  `json:"StartTime"`
	EndTime        int64  `json:"EndTime"`
	TokenValid     *int   `json:"TokenValid"`
}

type ListHistoryRequest struct {
	SpaceID    string `json:"-"`
	StreamID   string `json:"-"`
	PageNumber int    `json:"-"`
	PageSize   int    `json:"-"`
	StartTs    string `json:"StartTs"`
	EndTs      string `json:"EndTs"`
	ReqType    string `json:"-"`
}

type ListHistoryRequestV3 struct {
	StreamID       string
	StreamingIndex int
	Resolution     string
	PageNumber     int   `json:"-"`
	PageSize       int   `json:"-"`
	StartTime      int64 `json:"StartTime"`
	EndTime        int64 `json:"EndTime"`
	ReqType        string
}

type DeleteStreamRecordRequest struct {
	StreamID string
	ObjKey   []string
	Bucket   string
}

type AiotPlayBackControlRequest struct {
	StreamName string `json:"StreamName"`
	Ntp        string `json:"Ntp"`
	Scale      int    `json:"Scale"`
}

type CreateDeviceArg struct {
	Type                 string             `json:"Type"`
	DeviceName           string             `json:"DeviceName"`
	DeviceNSID           *string            `json:"DeviceNSID"`
	Username             *string            `json:"Username"`
	Password             *string            `json:"Password"`
	Description          string             `json:"Description"`
	AutoPullAfterRegiter *bool              `json:"AutoPullAfterRegiter"`
	PasswordLevel        string             `json:"PasswordLevel"`
	RtpTransportTcp      *bool              `json:"RtpTransportTcp"`
	DownloadSpeed        *int               `json:"DownloadSpeed"`
	Location             *string            `json:"Location"`
	Coordinates          *Coordinates       `json:"Coordinates"`
	AlertNotification    *AlertNotification `json:"AlertNotification"`
	UseSubStream         *bool              `json:"UseSubStream"`
}

type UpdateDeviceArg struct {
	DeviceName           *string            `json:"DeviceName"`
	DeviceNSID           *string            `json:"DeviceNSID"`
	Username             *string            `json:"Username"`
	Password             *string            `json:"Password"`
	Description          *string            `json:"Description"`
	AutoPullAfterRegiter *bool              `json:"AutoPullAfterRegiter"`
	AlertNotification    *AlertNotification `json:"AlertNotification"`
	ChannelNum           *int               `json:"-"`
	Status               string             `json:"-"`
	Type                 string             `json:"Type"`
	RtpTransportTcp      *bool              `json:"RtpTransportTcp"` //流传输协议tcp为true，否则默认udp
	Location             *string            `json:"Location"`        //可以修改设备的地址位置和坐标
	Coordinates          *Coordinates       `json:"Coordinates"`
	OnChannelNum         *int               `json:"-"`
	UseSubStream         *bool              `json:"UseSubStream"`
}

type CloudControlRequest struct {
	SipID     string
	DeviceID  string `json:"DeviceID"`
	ChannelID string `json:"ChannelID"`
	Cmd       string `json:"Cmd"`
	Para      uint8  `json:"Para"`
	Action    string `json:"action"`
}

type CloudControlRequestV3 struct {
	StreamID string `json:"StreamID"`
	Cmd      string `json:"Cmd"`
	Para     uint8  `json:"Para"`
	Action   string `json:"action"`
}

type QueryPresetInfoRequest struct {
	SipID     string
	DeviceID  string `json:"DeviceID"`
	ChannelID string `json:"ChannelID"`
	TimeOut   int    `json:"TimeOut"`
}

type QueryPresetInfoRequestV3 struct {
	StreamID string `json:"StreamID"`
	TimeOut  int    `json:"TimeOut"`
}

type GetRecordListRequest struct {
	DeviceID     string `json:"DeviceID"`
	ChannelID    string `json:"ChannelID"`
	StartTime    int64  `json:"StartTime"`
	EndTime      int64  `json:"EndTime"`
	RecordType   string `json:"RecordType"`
	TimeoutInSec int    `json:"TimeoutInSec"`
	Order        bool   `json:"Order"`
}

type GetRecordListRequestV3 struct {
	StreamID     string `json:"StreamID"`
	StartTime    int64  `json:"StartTime"`
	EndTime      int64  `json:"EndTime"`
	RecordType   string `json:"RecordType"`
	TimeoutInSec int    `json:"TimeoutInSec"`
	Order        bool   `json:"Order"`
}

type PlaybackRequest struct {
	PlaybackID string `json:"PlaybackID"`
}

type PlayBackStartRequest struct {
	DeviceID            string `json:"DeviceID"`
	ChannelID           string `json:"ChannelID"`
	StartTime           int64  `json:"StartTime"`
	EndTime             int64  `json:"EndTime"`
	Version             string `json:"Version"`
	MediaProcess        *Mps   `json:"MediaProcess"`
	DownloadSpeedSingle *int   `json:"DownloadSpeedSingle"` //在本地录像下载的时候时候，不填就以设备维度所设置的速度进行下载
	SchedQuotaKey       string `json:"SchedQuotaKey"`       //调度到哪个集群
}

type StartPlaybackRequestV3 struct {
	StreamID            string `json:"StreamID"`
	StartTime           int64  `json:"StartTime"`
	EndTime             int64  `json:"EndTime"`
	Version             string `json:"Version"`
	MediaProcess        *Mps   `json:"MediaProcess"`
	DownloadSpeedSingle *int   `json:"DownloadSpeedSingle"` //在本地录像下载的时候时候，不填就以设备维度所设置的速度进行下载
	SchedQuotaKey       string `json:"SchedQuotaKey"`       //调度到哪个集群
}

type PlayBackControlRequest struct {
	StreamID string  `json:"StreamID"`
	Cmd      int     `json:"Cmd"`
	Ntp      string  `json:"Ntp"`
	Scale    float32 `json:"Scale"`
}

type ControlPlaybackRequestV3 struct {
	PlaybackID string  `json:"PlaybackID"`
	Cmd        int     `json:"Cmd"`
	Ntp        string  `json:"Ntp"`
	Scale      float32 `json:"Scale"`
}

type CruiseControlRequest struct {
	DeviceNSID  string `json:"DeviceNSID"`
	ChannelID   string `json:"ChannelID"`
	Action      string `json:"Action"`
	TrackID     uint8  `json:"TrackID"`     // 巡航轨迹组ID
	PresetID    uint8  `json:"PresetID"`    // 预置位ID
	Speed       uint32 `json:"Speed"`       // 巡航速度
	StaySeconds uint32 `json:"StaySeconds"` // 巡航停留时间
}

type CruiseControlRequestV3 struct {
	StreamID    string `json:"StreamID"`
	Action      string `json:"Action"`
	TrackID     uint8  `json:"TrackID"`     // 巡航轨迹组ID
	PresetID    uint8  `json:"PresetID"`    // 预置位ID
	Speed       uint32 `json:"Speed"`       // 巡航速度
	StaySeconds uint32 `json:"StaySeconds"` // 巡航停留时间
}

type SetCruiseTrackRequest struct {
	DeviceNSID  string        `json:"DeviceNSID"`
	ChannelID   string        `json:"ChannelID"`
	TrackID     uint8         `json:"TrackID"`     // 巡航轨迹组编号(1~255)
	TrackList   []CruisePoint `json:"TrackList"`   // 巡航轨迹列表, 最多可添加255个预置点
	StaySeconds uint32        `json:"StaySeconds"` // 预置点停留时间, 单位:秒, 取值范围:1~4095
	Speed       uint32        `json:"Speed"`       // 巡航速度, 取值范围: 1~4095
}

type SetCruiseTrackRequestV3 struct {
	StreamID    string        `json:"StreamID"`
	TrackID     uint8         `json:"TrackID"`     // 巡航轨迹组编号(1~255)
	TrackList   []CruisePoint `json:"TrackList"`   // 巡航轨迹列表, 最多可添加255个预置点
	StaySeconds uint32        `json:"StaySeconds"` // 预置点停留时间, 单位:秒, 取值范围:1~4095
	Speed       uint32        `json:"Speed"`       // 巡航速度, 取值范围: 1~4095
}

type DeleteCruiseTrackRequest struct {
	DeviceNSID string `json:"DeviceNSID"`
	ChannelID  string `json:"ChannelID"`
	TrackID    uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type DeleteCruiseTrackRequestV3 struct {
	StreamID string `json:"StreamID"`
	TrackID  uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type StartCruiseTrackRequest struct {
	DeviceNSID string `json:"DeviceNSID"`
	ChannelID  string `json:"ChannelID"`
	TrackID    uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type StartCruiseTrackRequestV3 struct {
	StreamID string `json:"StreamID"`
	TrackID  uint8  `json:"TrackID"` // 巡航轨迹组编号(1~255)
}

type StopCruiseTrackRequest struct {
	DeviceNSID string `json:"DeviceNSID"`
	ChannelID  string `json:"ChannelID"`
}

type StopCruiseTrackRequestV3 struct {
	StreamID string `json:"StreamID"`
}

type CreateStreamRequest struct {
	SpaceID     string
	StreamName  string `json:"StreamName"` //流的名称，用户输入
	Description string `json:"Description"`
}

type UpdateStreamRequest struct {
	SpaceID        string
	StreamID       string
	StreamName     *string `json:"StreamName"`
	Description    *string `json:"Description"`
	TemplateEnable bool    `json:"template_enable"`
	Screenshot     *struct {
		TemplateName string `json:"TemplateName"`
		TemplateID   string `json:"TemplateID"`
	} `json:"Screenshot"`
	Record *struct {
		TemplateName string `json:"TemplateName"`
		TemplateID   string `json:"TemplateID"`
	} `json:"Record"`
	PushUrl     string `json:"push_url"`
	Status      string `json:"Status"`
	PushUrlDDL  *int64 `json:"PushUrlDDL"`
	IgnoreAudio bool   `json:"IgnoreAudio"`
}

type StreamRequest struct {
	StreamID           string `json:"StreamID,omitempty"`
	StreamingIndex     int    `json:"StreamingIndex,omitempty"` // 码流类型(可选参数), 0-主码流;1-子码流1;2-子码流2, 以此类推
	Resolution         string `json:"Resolution,omitempty"`     // 国标分辨率参数(可选参数), 1-QCIF(176x144); 2-CIF(320x288); 3-4CIF(704x576); 4-D1(720x576); 5-720p; 6-1080p; 其他：WxH表示
	EnableAudioTrans   bool   `json:"-"`
}

type ListStreamsRequest struct {
	SpaceID    string
	StreamName string
	PageNumber int
	PageSize   int
	Order      string
}

type GetStreamDataRequest struct {
	StreamID       string
	StartTime      string
	EndTime        string
	Resolution     string
	StreamingIndex int64
}

type StreamLogsRequest struct {
	StartTs    string
	EndTs      string
	StreamID   string
	Order      string
	PageNumber int
	PageSize   int
}

type CreateTemplateRequest struct {
	TemplateName string              `json:"TemplateName"`
	Screenshot   *ScreenshotTemplate `json:"Screenshot"`
	Record       *RecordTemplate     `json:"Record"`
	AI           *AITemplate         `json:"AI"`
	Trans        *TransTemplate      `json:"Trans"`
}

type DeleteTemplateRequest struct {
	TemplateID   string `json:"TemplateID"`
	TemplateType string `json:"TemplateType"`
}

type TemplateRequest struct {
	TemplateID string
}

type ListTemplateRequest struct {
	PageNumber int
	PageSize   int
	Order      string
}

type UpdateTemplateRequest struct {
	TemplateID   string
	TemplateName *string             `json:"TemplateName"`
	Screenshot   *ScreenshotTemplate `json:"Screenshot"`
	Record       *RecordTemplate     `json:"Record"`
	AI           *AITemplate         `json:"AI"`
	Trans        *TransTemplate      `json:"Trans"`
	Enable       *bool               `json:"enable"`
}

type CreateSlowLiveRequest struct {
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

type SlowLiveStreamConfig struct {
	StreamID    string `json:"StreamID"`
	EnableAudio bool   `json:"EnableAudio"`
}

type SlowLiveStatus string

type SlowLiveConfig struct {
	EnableRolling   bool                    `json:"EnableRolling"`
	RollingInterval int                     `json:"RollingInterval"`
	StreamList      []*SlowLiveStreamConfig `json:"StreamList"`
}

type UpdateSlowLiveRequest struct {
	LiveStreamID            string
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

type SlowLiveRequest struct {
	PageNumber int
	PageSize   int
}

type BatchDeleteSlowLiveRequest struct {
	LiveStreamIDs []string `json:"LiveStreamIDs"`
}

type SignSlowliveWsTokenRequest struct {
	ValidDuration int
	SlowliveID    string
}

type UpdateGroupTreeNodeRequest struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	DeviceID    string `json:"DeviceID"`
}

type DeleteGroupTreeNodesRequest struct {
	IDs []string `json:"IDs"`
}

type GetDevicesByGroupTreeNodeRequest struct {
	NodeID     string `json:"NodeID"`
	SearchName string `json:"SearchName"`
}

type StartVoiceTalkRequest struct {
	SpaceID    string
	DeviceNSID string
	UseTcp     bool
}

type StopVoiceTalkRequest struct {
	SpaceID    string
	DeviceNSID string
}

type CreateCascadePlatformRequest struct {
	PlatformName string           `json:"PlatformName"`
	Description  string           `json:"Description"`
	SipConfig    CascadeSipConfig `json:"SipConfig"`
	Enable       *bool            `json:"Enable"`       // 是否使能级联平台,不填默认使能
	EnableAccess bool             `json:"EnableAccess"` // 平台授权
	EnablePTZ    bool             `json:"EnablePTZ"`    // PTZ授权
}

type UpdateCascadePlatformRequest struct {
	PlatformID   string           `json:"PlatformID"`
	PlatformName string           `json:"PlatformName"` // 上级平台名称
	Description  string           `json:"Description"`  // 上级平台描述
	SipConfig    CascadeSipConfig `json:"SipConfig"`
	EnableAccess bool             `json:"EnableAccess"` // 平台授权
	EnablePTZ    bool             `json:"EnablePTZ"`    // PTZ授权
}

type ListCascadePlatformRequest struct {
	PlatformName string
	PageNumber   int64
	PageSize     int64
	Order        string
}

type BatchDeleteCascadePlatformRequest struct {
	PlatformIDs []string `json:"IDs"`
}

type CreateCascadeTaskRequest struct {
	TaskName    string `json:"TaskName"`    // 级联任务名称
	GroupTreeID string `json:"GroupTreeID"` // 关联的虚拟组织树ID
	PlatformID  string `json:"PlatformID"`  // 关联的上级平台ID
	Description string `json:"Description"`
}

type UpdateCascadeTaskArg struct {
	TaskName string `json:"TaskName"`
}

type BatchDeleteCascadeTaskRequest struct {
	TaskIDs []string `json:"IDs"`
}

type ListCascadeTaskRequest struct {
	PlatformName  string `json:"PlatformName"`
	PlatformID    string `json:"PlatformID"`
	GroupTreeID   string `json:"GroupTreeID"`
	GroupTreeName string `json:"GroupTreeName"`
	PageNumber    int    `json:"PageNumber"`
	PageSize      int    `json:"PageSize"`
	Order         string `json:"Order"`
}

type ShareGroupNodesRequest struct {
	TaskID       string   `json:"TaskID"`
	GroupNodeIDs []string `json:"GroupNodeIDs"`
}

type AddGroupTreeNodeRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	SpaceID     string `json:"SpaceID"`
	ParentID    string `json:"ParentID"`
}

type CreateStructuredViewSpaceRequest struct {
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

type StructuredViewSpaceStatus string

type UpdateStructuredViewSpaceRequest struct {
	SpaceID         string
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

type ListStructuredViewSpacesRequest struct {
	SpaceName  string
	Order      string
	PageNumber int
	PageSize   int
}

type CreateStructuredViewRequest struct {
	ID               string               ` json:"ID"`
	AccountID        string               ` json:"AccountID"`
	ViewName         string               ` json:"ViewName"`
	ViewType         string               `json:"ViewType"`
	ViewCode         string               ` json:"ViewCode"`
	ViewSpaceID      string               ` json:"ViewSpaceID"`
	ViewSpaceName    string               ` json:"ViewSpaceName"`
	ViewIP           string               `json:"ViewIP"`
	ViewPort         int                  ` json:"ViewPort"`
	Protocol         string               ` json:"Protocol"`
	Username         string               ` json:"Username"`
	Password         string               ` json:"Password"`
	Description      string               ` json:"Description"`
	Location         string               ` json:"Location"`
	AdministrativeID string               ` json:"AdministrativeID"`
	CreateTimestamp  int64                ` json:"CreateTimestamp"`
	UpdateTimestamp  int64                ` json:"UpdateTimestamp"`
	Status           StructuredViewStatus ` json:"Status"`
}

type UpdateStructuredViewRequest struct {
	ViewID           string
	ID               string               ` json:"ID"`
	AccountID        string               ` json:"AccountID"`
	ViewName         string               ` json:"ViewName"`
	ViewType         string               `json:"ViewType"`
	ViewCode         string               ` json:"ViewCode"`
	ViewSpaceID      string               ` json:"ViewSpaceID"`
	ViewSpaceName    string               ` json:"ViewSpaceName"`
	ViewIP           string               `json:"ViewIP"`
	ViewPort         int                  ` json:"ViewPort"`
	Protocol         string               ` json:"Protocol"`
	Username         string               ` json:"Username"`
	Password         string               ` json:"Password"`
	Description      string               ` json:"Description"`
	Location         string               ` json:"Location"`
	AdministrativeID string               ` json:"AdministrativeID"`
	CreateTimestamp  int64                ` json:"CreateTimestamp"`
	UpdateTimestamp  int64                ` json:"UpdateTimestamp"`
	Status           StructuredViewStatus ` json:"Status"`
}

type ListStructuredViewsRequest struct {
	ViewSpaceName string
	ViewName      string
	PageNumber    int
	PageSize      int
	Order         string
}

type ListStructuredViewDataRequest struct {
	ViewDataType string
	StartTs      string
	EndTs        string
	ViewID       string
	PageNumber   int
	PageSize     int
	Order        string
}

type StructuredViewCascadePlatformStatus string

type CreateStructuredViewCascadePlatformRequest struct {
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

type ListStructuredViewCascadePlatformRequest struct {
	PlatformName string
	PageNumber   int
	PageSize     int
	Order        string
}

type CreateStructuredViewCascadeJobRequest struct {
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

type UpdateStructuredCascadeJobRequest struct {
	JobID           string
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

type StructuredViewCascadeJobStatus string

type UpdateStructuredViewCascadePlatformRequest struct {
	PlatformID      string
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

type ListStructuredViewCascadeJobRequest struct {
	JobName    string
	PageNumber int
	PageSize   int
	Order      string
}

type CreateAccountRequest struct {
	ChargeMode ChargeMode  `json:"ChargeMode"` //bind=带宽计费 flow=流量计费
	ChannelPkg *ChannelPkg `json:"ChannelPkg"`
}

type ChannelPkg struct {
	Level     string `json:"Level"`
	Cnt       int    `json:"Cnt"`       // 购买的路数
	RenewType *int32 `json:"RenewType"` //自动续费方式
}

type UpdateAccountRequest struct {
	ChargeMode *ChargeMode `json:"ChargeMode"`
}

type AddSpaceDomainRequest struct {
	SpaceID string
	Domain  string `json:"Domain"`
	Type    string `json:"Type"`
}

type UpdateDomainHTTPSRequest struct {
	SpaceID       string
	Domain        string
	Enable        bool   `json:"Enable"`
	CertificateID string `json:"CertificateID"`
}

type UpdateSpaceDomainRequest struct {
	SpaceID     string
	SpaceDomain string `json:"SpaceDomain"`
	Default     bool   `json:"Default"`
}

type UpdateAuthInSpaceRequest struct {
	SpaceID       string `json:"SpaceID"`
	Domain        string `json:"Domain"`
	MainKey       string `json:"MainKey"`
	SpareKey      string `json:"SpareKey"`
	ValidDuration int64  `json:"ValidDuration"`
}

type GetDataProjectWithBindWidthAndFlowRequest struct {
	SpaceID    string
	StreamName string
	StartTime  string
	EndTime    string
	Data       string
}

type GetPushStreamCntRequest struct {
	SpaceID   string
	StartTime string
	EndTime   string
}

type SetAlarmGuardRequest struct {
	SipID      string
	DeviceNSID string
	ChannelID  string
	Enable     string
}

type SetDeviceAlarmRequestV3 struct {
	DeviceID string `json:"DeviceID"`
	StreamID string `json:"StreamID"`
	Enable   string `json:"Enable"`
}

type ResetAlarmRequest struct {
	SipID      string
	DeviceNSID string
	ChannelID  string
}

type ResetDeviceAlarmRequestV3 struct {
	DeviceID string `json:"DeviceID"`
	StreamID string `json:"StreamID"`
}

type ListAlarmNotifyRequest struct {
	PageNumber    int      `json:"PageNumber"`
	PageSize      int      `json:"PageSize"`
	DeviceNSID    string   `json:"DeviceNSID"`
	ChannelID     string   `json:"ChannelID"`
	StartTime     int64    `json:"StartTime"`
	EndTime       int64    `json:"EndTime"`
	AlarmMethod   []uint32 `json:"AlarmMethod"`
	AlarmPriority []uint32 `json:"AlarmPriority"`
	AlarmType2    []uint32 `json:"AlarmType2"`
	AlarmType5    []uint32 `json:"AlarmType5"`
	AlarmType6    []uint32 `json:"AlarmType6"`
	Order         uint32   `json:"Order"`
}

type ListDeviceAlarmsRequestV3 struct {
	PageNumber int
	PageSize   int
	DeviceID   string `json:"DeviceID"`
	StreamID   string `json:"StreamID"`
	StartTime  int64  `bson:"StartTime"`
	EndTime    int64  `bson:"EndTime"`
	// 告警方式筛选条件，可选枚举值为（如传入空数组，即返回所有告警方式的报警记录）
	//1（电话报警）
	//2（设备报警）
	//3（短信报警
	//4（GPS报警）
	//5（视频报警）
	//6（设备故障报警）
	//7（其他报警）
	AlarmMethod []uint32 `bson:"AlarmMethod"`
	// 报警级别筛选条件，可选枚举值为（如传入空数组，即返回所有告警级别的报警记录）：
	//1（一级警情）
	//2（二级警情）
	//3（三级警情）
	//4（四级警情）
	AlarmPriority []uint32 `bson:"AlarmPriority"`
	Order         uint32   `bson:"Order"` // 排序方式：0（正序）,其他（倒序）
}

type UploadCertRequest struct {
	CertName   string `json:"CertName"`
	CertDesc   string `json:"CertDesc"`
	SSLPublic  string `json:"SSLPublic"`
	SSLPrivate string `json:"SSLPrivate"`
}

type UpdateCertRequest struct {
	SSLPublic  string `json:"SSLPublic"`
	SSLPrivate string `json:"SSLPrivate"`
	ChainID    string `json:"ChainID"`
	CertName   string `json:"CertName"`
	CertDesc   string `json:"CertDesc"`
}

type BindCertRequest struct {
	SpaceID    string
	Domain     string `json:"Domain"`
	ChainID    string `json:"ChainID"`
	DomainType string `json:"DomainType"`
}

type ListCertificatesRequest struct {
	PageNumber int
	PageSize   int
}

type UnbindCertRequest struct {
	SpaceID string
	Domain  string `json:"Domain"`
}

type DeleteCertRequest struct {
	ChainID string `json:"ChainID"`
}

type EdgeInviteRequest struct {
	StreamID      string
	LocalFlvHttps string
	Status        string
	SSRC          int
	RTPPort       int
	IP            string
}

type EdgeStatusRequest struct {
	StreamID      string
	LocalFlvHttps string
	Status        string
	SSRC          int
	RTPPort       int
	IP            string
}

type GetAIAnalysisResultRequest struct {
	StreamID string              `json:"StreamID"`
	StartT   int64               `json:"StartTime"`
	EndT     int64               `json:"EndTime"`
	AI       map[string][]string `json:"AI"`
}

package live

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type (
	EResourceUsed  int
	EResourceType  int
	EResourceCate  int
	EResourceForm  string
	ERTCStatus     int
	ERTCUserID     string
	EScene         int
	ECasterStatus  int
	CasterTaskCode int
)

const (
	CasterInit    ECasterStatus = 1
	CasterRunning ECasterStatus = 2
	CasterFinish  ECasterStatus = 3

	SceneLayout EScene = 0
	ScenePVW    EScene = 1
	ScenePGM    EScene = 2

	ResourceNoneUsed   EResourceUsed = 0
	ResourceLayoutUsed EResourceUsed = 1
	ResourcePVWUsed    EResourceUsed = 2
	ResourcePGMUsed    EResourceUsed = 3

	ResourceTypePull            EResourceType = 1
	ResourceTypePush            EResourceType = 2
	ResourceTypeVideo           EResourceType = 3
	ResourceTypeVideoURL        EResourceType = 4
	ResourceTypePicture         EResourceType = 5
	ResourceTypePushMate        EResourceType = 6
	ResourceTypeImageCollection EResourceType = 7
	ResourceTypeOPED            EResourceType = 8

	ResourceCateVideo           EResourceCate = 1
	ResourceCatePicture         EResourceCate = 2
	ResourceCateText            EResourceCate = 3
	ResourceCateImageCollection EResourceCate = 4
	ResourceCateOPED            EResourceCate = 5

	RTCStatusNone ERTCStatus = 0
	RTCStatusOk   ERTCStatus = 1
	RTCStatusErr  ERTCStatus = 2

	ResourceFormNone EResourceForm = ""
	ResourceFormRts  EResourceForm = "rts"
	ResourceFormFlv  EResourceForm = "flv"

	CasterTaskCodeSuccess    CasterTaskCode = 0
	CasterTaskCodeInit       CasterTaskCode = 1
	CasterTaskCodeStop       CasterTaskCode = 2
	CasterTaskCodeNotCreated CasterTaskCode = 3
	CasterTaskCodeStartField CasterTaskCode = 500

	CloudCastDownloadStatusProcessing = 1
	CloudCastDownloadStatusSuc        = 2
	CloudCastDownloadStatusFail       = 3
)

type CreateCasterResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CreateCasterResponse `json:"Result,omitempty"`
}

type CreateCasterResponse struct {
	CasterID       *int64
	BackupCasterID *int64
}

type ListCastersResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListCastersResponse `json:"Result,omitempty"`
}

type ListCastersResponse struct {
	Casters []*CasterBaseCondition
	Count   int64
}

type GetCasterResourceResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetCasterResourceResponse `json:"Result,omitempty"`
}

type GetCasterResourceResponse struct {
	Resources *CasterInfoResources
}

type GetCasterLayoutResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetCasterLayoutResponse `json:"Result,omitempty"`
}

type GetCasterLayoutResponse struct {
	Layouts    []*CasterLayoutCondition
	LayoutTpls []*CasterLayoutTplCondition
}

type GetCasterConfigResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetCasterConfigResponse `json:"Result,omitempty"`
}

type GetCasterConfigResponse struct {
	Caster     *CasterCondition
	PushParams *PushParams
}

type StartCasterResp struct {
	ResponseMetadata base.ResponseMetadata
}

type StopCasterResp struct {
	ResponseMetadata base.ResponseMetadata
}

type CreateCasterResourceResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CasterResourceCondition `json:"Result,omitempty"`
}

type UpdateCasterLayoutResp struct {
	ResponseMetadata base.ResponseMetadata
}

type DeleteCasterLayoutResp struct {
	ResponseMetadata base.ResponseMetadata
}

type CreateCasterResourceImageResp struct {
	ResponseMetadata base.ResponseMetadata
}

type DeleteCasterResourceImageResp struct {
	ResponseMetadata base.ResponseMetadata
}

type CreateCasterResourceImagesResp struct {
	ResponseMetadata base.ResponseMetadata
}

type RemoveCasterResourceResp struct {
	ResponseMetadata base.ResponseMetadata
}

type UpdateCasterConfigResp struct {
	ResponseMetadata base.ResponseMetadata
}

type CreateCasterResourceOPEDResp struct {
	ResponseMetadata base.ResponseMetadata
}

type DeleteCasterResourceOPEDResp struct {
	ResponseMetadata base.ResponseMetadata
}

type SwitchCasterResourceResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SwitchCasterResourceResponse `json:"Result,omitempty"`
}

type SwitchCasterResourceResponse struct {
	NotStart bool
}

type SwitchCasterResourceImageResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SwitchCasterResourceImageResponse `json:"Result,omitempty"`
}

type SwitchCasterResourceImageResponse struct {
	NotStart bool
}

type SwitchCasterResourceOPEDResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SwitchCasterResourceOPEDResponse `json:"Result,omitempty"`
}

type SwitchCasterResourceOPEDResponse struct {
	NotStart bool
}

type StartCasterOPEDArrangeResp struct {
	ResponseMetadata base.ResponseMetadata
}

type SwitchCasterLayoutResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SwitchCasterLayoutResponse `json:"Result,omitempty"`
}

type SwitchCasterLayoutResponse struct {
	NotStart bool
}

type CopyCasterPVWToPGMResp struct {
	ResponseMetadata base.ResponseMetadata
}

type GetCasterResourceVodInfoResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetCasterResourceVodInfoResponse `json:"Result,omitempty"`
}

type GetCasterResourceVodInfoResponse struct {
	Resource []*ResourceVodPlayInfo
}

type GetCasterArrangeResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetCasterArrangeResponse `json:"Result,omitempty"`
}

type GetCasterArrangeResponse struct {
	Arranges []*CasterArrange
}

type CreateCasterArrangeResp struct {
	ResponseMetadata base.ResponseMetadata
}

type UpdateCasterArrangeResp struct {
	ResponseMetadata base.ResponseMetadata
}

type DeleteCasterArrangeResp struct {
	ResponseMetadata base.ResponseMetadata
}

type CasterBaseCondition struct {
	ID            int64
	CloudCasterID string
	IsBackup      bool
	Name          string
	Status        ECasterStatus
	CreateTime    int64
	UpdateTime    int64
	Creator       string
}

type CasterCondition struct {
	Name          string
	ID            int64
	UID           string
	CloudCasterID string
	Status        ECasterStatus
	HasBackup     bool
	CreateTime    int64
	UpdateTime    int64
}

type PushParams struct {
	Direction int
	Size      int
	Fps       int
	VBitRate  int
	Delay     int
	SEI       string
	SEIs      []string
	Logo      Logo
	Gop       int
	ABitRate  int
	AChannel  string
	Codec     string
	Profile   string
	ASample   int
	Tune      string
	Options   string
	Streams   []*Stream
}

type Logo struct {
	URL string
	X   int
	Y   int
	W   int
	H   int
}

type Stream struct {
	ID  int64
	URL string
}

type CasterResourceCondition struct {
	ID                  int64
	Name                string
	No                  int
	Type                int
	URL                 string
	Volume              int
	Duration            int
	BreakAt             int
	Mute                bool
	VodHeadSwitch       bool
	VideoUsed           int
	AudioUsed           int
	RtcPushUserID       string
	Status              int
	PGMUsed             bool
	PVWUsed             bool
	AdjustDts           int64
	AdjustAudio         int
	AdjustRtcPushUserID string
	VodUseDownload      bool
}

type CasterLayoutCondition struct {
	ID       int64
	Name     string
	TplId    int64
	Elements []*Element
	PVWUsed  *bool
	PGMUsed  *bool
	BackupID int64
}

type CasterLayoutTplCondition struct {
	ID         int64
	Name       string
	Elements   []*Element
	W          int
	H          int
	ScreenType int
	CustomType int
}

type Element struct {
	Name         string
	W            float32
	H            float32
	X            float32
	Y            float32
	ResourceID   int
	ResourceNO   int
	Opacity      int
	ZIndex       int
	Cate         int
	IsHide       bool
	ImageIndexID int64
}

type CasterInfoResources struct {
	Video            []*Video
	Image            []*Image
	ImageCollections []*ImageCollections
	OPED             []*OPED
}

type Video struct {
	ID                  int64
	Name                string
	No                  int64
	Type                EResourceType
	URL                 string
	Volume              int
	Duration            int
	BreakAt             int64
	Mute                bool
	VodHeadSwitch       bool
	VideoUsed           EResourceUsed
	AudioUsed           EResourceUsed
	RtcPushUserID       ERTCUserID
	Status              ERTCStatus
	PGMUsed             bool
	PVWUsed             bool
	AdjustDts           int64
	AdjustAudio         int
	AdjustRtcPushUserID ERTCUserID
	VodUseDownload      bool
}

type Image struct {
	ID      int64
	ImageID string
	Name    string
	URL     string
	PGMUsed bool
	PVWUsed bool
}

type ImageCollections struct {
	ID              int64
	Name            string
	ImageCollection []ImageCollection
	DownloadStatus  int
	DownloadPercent float32
}

type ImageCollection struct {
	ImageIndexID   int64
	ImageID        string
	URL            string
	DownloadStatus int
}

type OPED struct {
	ID          int64
	Name        string
	URL         string
	PVWUsed     bool
	PGMUsed     bool
	Duration    int
	ArrangeInfo ArrangeInfo
}

type ArrangeInfo struct {
	OpType     int
	ResourceNo int64
	ResourceID int64
	LayoutID   int64
}

type CasterArrange struct {
	ID            int64
	OpType        int
	ExecType      int
	ExecTime      int64
	CountdownTime int64
	ResourceNo    int64
	ResourceID    int64
	LayoutID      int64
	MainBackup    bool
}

type ResourceVodPlayInfo struct {
	ID          int64
	Duration    int
	Type        int
	VodPlayTime int
}

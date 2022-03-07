package live

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type ListCommonTransPresetDetailReq struct {
	PresetList []string `json:"PresetList"`
}

type ListCommonTransPresetDetailResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListCommonTransPresetDetailResult
}
type CallbackDetail struct {
	CallbackType string `json:"CallbackType" binding:"oneof=http nsq kafka rpc"` // nsq kafka rpc not supported temporally
	URL          string `json:"URL" binding:"required"`
}

type CallbackInfo struct {
	Vhost               string            `json:"Vhost"`
	App                 string            `json:"App" `
	MessageType         string            `json:"MessageType"`
	TranscodeCallback   int64             `json:"TranscodeCallback"`
	CallbackDetailList  []*CallbackDetail `json:"CallbackDetailList"`
	AuthEnable          bool              `json:"AuthEnable"`
	AuthKeyPrimary      string            `json:"AuthKeyPrimary"`
	AuthKeySecond       string            `json:"AuthKeySecond"`
	ValidDuration       int64             `json:"ValidDuration"`
	EncryptionAlgorithm string            `json:"EncryptionAlgorithm"`
	CallbackField       []string          `json:"CallbackField"`
	AuthField           map[string]string `json:"AuthField"`
	EncryptField        []string          `json:"EncryptField"`
	AppendField         map[string]string `json:"AppendField"`
	TimeoutSecond       int64             `json:"TimeoutSecond"`
	RetryTimes          int64             `json:"RetryTimes"`
	RetryInternalSecond int64             `json:"RetryInternalSecond"`
}
type DescribeCallbackOutput struct {
	CallbackList []*CallbackInfo `json:"CallbackList"`
}
type DescribeCallbackResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeCallbackOutput `json:"Result,omitempty"`
}
type UpdateCallbackResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DeleteCallbackResp struct {
	ResponseMetadata base.ResponseMetadata
}
type CreateDomainResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DeleteDomainResp struct {
	ResponseMetadata base.ResponseMetadata
}

type DomainInfo struct {
	Vhost       string `json:"Vhost"`
	Domain      string `json:"Domain"`
	Type        string `json:"Type"`
	Region      string `json:"Region"`
	Status      int64  `json:"Status"`
	CName       string `json:"CName"`
	ChainID     string `json:"ChainID"`
	CertDomain  string `json:"CertDomain"`
	CnameCheck  int64  `json:"CnameCheck"`
	DomainCheck int64  `json:"DomainCheck"`
	ICPCheck    int64  `json:"ICPCheck"`
	CreateTime  string `json:"CreateTime"`
	CertName    string `json:"CertName"`
	PushDomain  string `json:"PushDomain"`
}

type DescribeDomainOutput struct {
	DomainList []*DomainInfo `json:"DomainList"`
}
type ListDomainDetailOutput struct {
	DomainList []*DomainInfo `json:"DomainList"`
	Total      int64         `json:"Total"`
}

type ListDomainDetailResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListDomainDetailOutput `json:"Result,omitempty"`
}
type DescribeDomainResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeDomainOutput `json:"Result,omitempty"`
}
type EnableDomainResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DisableDomainResp struct {
	ResponseMetadata base.ResponseMetadata
}
type ManagerPullPushDomainBindResp struct {
	ResponseMetadata base.ResponseMetadata
}
type UpdateAuthKeyResp struct {
	ResponseMetadata base.ResponseMetadata
}
type EnableAuthResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DisableAuthResp struct {
	ResponseMetadata base.ResponseMetadata
}
type AuthDetail struct {
}

type AuthInfoItem struct {
	Vhost          string      `json:"Vhost"`
	App            string      `json:"App"`
	SceneType      string      `json:"SceneType"`
	AuthStatus     bool        `json:"AuthStatus"`
	ValidDuration  int64       `json:"ValidDuration"`
	AuthDetailList interface{} `json:"AuthDetailList"`
}

type DescribeAuthOutput struct {
	AuthList []*AuthInfoItem `json:"AuthList"`
}

type DescribeAuthResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeAuthOutput `json:"Result,omitempty"`
}

type ForbidStreamResp struct {
	ResponseMetadata base.ResponseMetadata
}

type ResumeStreamResp struct {
	ResponseMetadata base.ResponseMetadata
}
type ListCertConsoleResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           ListCertResp `json:"Result,omitempty"`
}

type ListCertResp struct {
	CertList []SimpleCertInfo `json:"CertList"`
}

type SimpleCertInfo struct {
	CertDomain string `json:"CertDomain"`
	ChainID    string `json:"ChainID"`
	NotBefore  string `json:"NotBefore"`
	NotAfter   string `json:"NotAfter"`
	Status     string `json:"Status"`
	StatusCode int    `json:"statuscode"`
}
type CreateCertConsoleResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           OpenAPICreatChainResponse `json:"Result,omitempty"`
}

type OpenAPICreatChainResponse struct {
	ChainID   string `json:"ChainID"`
	Domain    string `json:"Domain"`
	UseWay    string `json:"UseWay"`
	AccountID string `json:"AccountID"`
}
type DescribeCertDetailResp struct {
	ChainID    string           `json:"ChainID"`
	CertDomain string           `json:"Domain"`
	UseWay     string           `json:"UseWay,omitempty"`
	Status     string           `json:"Status,omitempty"`
	Rsa        *OpenAPICertData `json:"Rsa,omitempty"`
}

type OpenAPICertData struct {
	PubKey       string `json:"PubKey"`
	PriKey       string `json:"PriKey"`
	CertType     string `json:"CertType"`
	PubName      string `json:"PubName"`
	PriName      string `json:"PriName"`
	NotBefore    string `json:"NotBefore"`
	NotAfter     string `json:"NotAfter"`
	FingerPrint  string `json:"FingerPrint"`
	SerialNumber string `json:"SerialNumber"`
}
type DescribeCertDetailConsoleResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           DescribeCertDetailResp `json:"Result,omitempty"`
}
type UpdateCertConsoleResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           OpenAPICreatChainResponse `json:"Result,omitempty"`
}
type BindCertConsoleResp struct {
	ResponseMetadata base.ResponseMetadata
}
type UnbindCertConsoleResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DeleteCertConsoleResp struct {
	ResponseMetadata base.ResponseMetadata
}
type UpdateRefererResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DeleteRefererResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DescribeRefererResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeRefererResult `json:"Result,omitempty"`
}
type DescribeRefererResult struct {
	RefererList []*RefererItem `json:"RefererList,omitempty"`
}
type RefererItem struct {
	Vhost           string         `json:"Vhost"`
	App             string         `json:"app"`
	RefererInfoList []*RefererInfo `json:"RefererInfoList"`
}
type RefererInfo struct {
	Type     string `json:"Type"`
	Key      string `json:"Key"`
	Value    string `json:"Value"`
	Priority int64  `json:"Priority"`
}
type CreateRecordPresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type UpdateRecordPresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DeleteRecordPresetResp struct {
	ResponseMetadata base.ResponseMetadata
}

type SlicePreset struct {
	Preset         *string         `json:"Preset,omitempty"`
	Description    *string         `json:"Description,omitempty"`
	AccountID      *string         `json:"AccountID,omitempty"`
	Bucket         *string         `json:"Bucket,omitempty"`
	Status         *int64          `json:"Status,omitempty"`
	AccessKey      *string         `json:"AccessKey,omitempty"`
	Interval       *int64          `json:"Interval,omitempty"`
	Format         []string        `json:"Format,omitempty"`
	Duration       *int64          `json:"Duration,omitempty"`
	SliceDuration  *int64          `json:"SliceDuration,omitempty"`
	ReserveDays    *int64          `json:"ReserveDays,omitempty"`
	SnapshotFormat *string         `json:"SnapshotFormat,omitempty"`
	CallbackDetail *CallbackDetail `json:"CallbackDetail"`
}
type SlicePresetsVhostAPP struct {
	Vhost       string       `json:"Vhost"`
	App         string       `json:"App"`
	SlicePreset *SlicePreset `json:"SlicePreset"`
}
type ListVhostRecordPresetRespOutput struct {
	PresetList []*SlicePresetsVhostAPP `json:"PresetList"`
}
type ListVhostRecordPresetResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListVhostRecordPresetRespOutput `json:"Result,omitempty"`
}
type CreateTranscodePresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type UpdateTranscodePresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DeleteTranscodePresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type ListVhostTransCodePresetResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListVhostTransCodePresetRespOutput `json:"Result,omitempty"`
}
type ListVhostTransCodePresetRespOutput struct {
	CommonPresetList    []TranscodePresetVhostAPP `json:"CommonPresetList"`
	CustomizePresetList []TranscodePresetVhostAPP `json:"CustomizePresetList"`
}
type TranscodePresetVhostAPP struct {
	Vhost           string           `json:"Vhost"`
	App             string           `json:"App"`
	TranscodePreset *TranscodePreset `json:"TranscodePreset"`
}
type TranscodePreset struct {
	Preset       *string `json:"Preset,omitempty"`
	Status       *int64  `json:"Status,omitempty"`
	SuffixName   *string `json:"SuffixName,omitempty"`
	StopInterval *int64  `json:"StopInterval,omitempty"`
	Describe     *string `json:"Describe,omitempty"`
	PresetKind   *int64  `json:"PresetKind,omitempty"`
	PresetType   *int    `json:"PresetType,omitempty"`
	Roi          *bool   `json:"Roi,omitempty"`
	Vclass       *bool   `json:"Vclass,omitempty"`
	Ocr          *bool   `json:"Ocr,omitempty"`
	Modifier     *string `json:"Modifier,omitempty"`
	Revision     *string `json:"Revision,omitempty"`
	//*****video param ******
	Vn           *int64  `json:"Vn,omitempty"`
	FPS          *int64  `json:"FPS,omitempty"`
	VideoBitrate *int64  `json:"VideoBitrate,omitempty"`
	VbThreshold  *string `json:"VbThreshold,omitempty"`
	Vcodec       *string `json:"Vcodec,omitempty"`
	VProfile     *string `json:"VProfile,omitempty"`
	VLevel       *string `json:"VLevel,omitempty"`
	VRateCtrl    *string `json:"VRateCtrl,omitempty"`
	GopMin       *int64  `json:"GopMin,omitempty"`
	GOP          *int64  `json:"GOP,omitempty"`
	BFrames      *int64  `json:"BFrames,omitempty"`
	LookAhead    *int64  `json:"LookAhead,omitempty"`
	VPreset      *string `json:"VPreset,omitempty"`
	Threads      *int64  `json:"Threads,omitempty"`
	Width        *int64  `json:"Width,omitempty"`
	Height       *int64  `json:"Height,omitempty"`
	As           *string `json:"As,omitempty"`
	AutoTrans    *int64  `json:"AutoTrans,omitempty"`
	LongSide     *int64  `json:"LongSide,omitempty"`
	ShortSide    *int64  `json:"ShortSide,omitempty"`
	Abr          *bool   `json:"Abr,omitempty"`
	VBVBufSize   *int64  `json:"VBVBufSize,omitempty"`
	VBVMaxRate   *int64  `json:"VBVMaxRate,omitempty"`
	Qp           *int64  `json:"Qp,omitempty"`
	HVSPre       *bool   `json:"HVSPre,omitempty"`
	BCM          *int64  `json:"BCM,omitempty"`
	VBRatio      *int64  `json:"VBRatio,omitempty"`
	SITI         *bool   `json:"SITI,omitempty"`

	// Nvidia hardware encoding related parameters, Vcodec, Width, Height, Vr, Vb, gop reuse the above general parameters, other software encoding parameters are ignored
	// When NvPriority = 0, it means that nvidia hardware transcoding is not enabled, and all nv parameters are not used
	NvPriority  *int64  `json:"NvPriority,omitempty"` // nvidia transcoding priority, 0 means off, >0 means on, the larger the priority, the higher the priority
	NvCodec     *string `json:"NvCodec,omitempty"`
	NvPreset    *string `json:"NvPreset,omitempty"`
	NvProfile   *string `json:"NvProfile,omitempty"`
	NvGop       *int64  `json:"NvGop,omitempty"`
	NvBf        *int64  `json:"NvBf,omitempty"`
	NvRefs      *int64  `json:"NvRefs,omitempty"`
	NvLookahead *int64  `json:"NvLookahead,omitempty"`
	NvTempAQ    *int64  `json:"NvTempAQ,omitempty"`
	NvHVSPre    *bool   `json:"NvHVSPre,omitempty"`
	NvPercent   *int64  `json:"NvPercent,omitempty"`
	//***** audio param *****
	An            *int64  `json:"An,omitempty"`
	AR            *int64  `json:"AR,omitempty"`
	AudioBitrate  *int64  `json:"AudioBitrate,omitempty"`
	Acodec        *string `json:"Acodec,omitempty"`
	AProfile      *string `json:"AProfile,omitempty"`
	RegionConfig  *string `json:"RegionConfig,omitempty"`
	AdvancedParam *string `json:"AdvancedParam,omitempty"`
}
type CreateSnapshotPresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type UpdateSnapshotPresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type DeleteSnapshotPresetResp struct {
	ResponseMetadata base.ResponseMetadata
}
type ListVhostSnapshotPresetRespOutput struct {
	PresetList []*SlicePresetsVhostAPP `json:"PresetList"`
}
type ListVhostSnapshotPresetResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListVhostSnapshotPresetRespOutput `json:"Result,omitempty"`
}
type ListCommonTransPresetDetailResult struct {
	StandardPresetDetail     []*TranscodePresetResult `json:"StandardPresetDetail"`
	NarrowBandHDPresetDetail []*TranscodePresetResult `json:"NarrowBandHDPresetDetail"`
}

type TranscodePresetResult struct {
	Preset       *string `json:"Preset,omitempty"`
	Status       *int64  `json:"Status,omitempty"`
	SuffixName   *string `json:"SuffixName,omitempty"`
	VideoBitrate *int64  `json:"VideoBitrate,omitempty"`
	Vcodec       *string `json:"Vcodec,omitempty"`
	AudioBitrate *int64  `json:"AudioBitrate,omitempty"`
	Acodec       *string `json:"Acodec,omitempty"`
	FPS          *int64  `json:"FPS,omitempty"`
	GOP          *int64  `json:"GOP,omitempty"`
	Width        *int64  `json:"Width,omitempty"`
	Height       *int64  `json:"Height,omitempty"`
	AutoTrans    *int64  `json:"AutoTrans,omitempty"`
	As           *string `json:"As,omitempty"`
	ShortSide    *int64  `json:"ShortSide,omitempty"`
	LongSide     *int64  `json:"LongSide,omitempty"`
	Roi          *bool   `json:"Roi,omitempty"`
}

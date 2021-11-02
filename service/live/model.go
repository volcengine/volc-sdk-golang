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
type UpdateAllAuthUnderVhostResp struct {
	ResponseMetadata base.ResponseMetadata
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
type DescribeRecordPresetResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeRecordPresetOutput `json:"Result,omitempty"`
}
type DescribeRecordPresetOutput struct {
	PresetList []string `json:"PresetList"`
}
type DescribeRecordPresetDetailResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeRecordPresetDetailOutput `json:"Result,omitempty"`
}
type DescribeRecordPresetDetailOutput struct {
	PresetDetailList []*RecordPreset `json:"PresetDetailList"`
}
type RecordPreset struct {
	Preset          string          `json:"Preset"`
	Status          int             `json:"Status"`
	Format          []string        `json:"Format"`
	Duration        int             `json:"Duration"`
	SliceDuration   int             `json:"SliceDuration"`
	Bucket          string          `json:"Bucket"`
	ReserveDuration int             `json:"ReserveDuration"`
	AccessKey       string          `json:"AccessKey"`
	ReserveDays     int64           `json:"ReserveDays"`
	CallbackDetail  *CallbackDetail `json:"CallbackDetail"`
	PullDomain      string          `json:"PullDomain"`
	VodNamespace    string          `json:"VodNamespace"`
	WorkflowID      string          `json:"WorkflowID"`
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
type DescribeTranscodePresetOutput struct {
	PresetList []string `json:"PresetList"`
}
type DescribeTranscodePresetResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeTranscodePresetOutput `json:"Result,omitempty"`
}

type TranscodePreset struct {
	Preset       string `json:"Preset"`
	Status       int    `json:"Status"`
	SuffixName   string `json:"SuffixName"`
	VideoBitrate int    `json:"VideoBitrate"`
	Vcodec       string `json:"Vcodec"`
	AudioBitrate int    `json:"AudioBitrate"`
	Acodec       string `json:"Acodec"`
	FPS          int    `json:"FPS"`
	GOP          int    `json:"GOP"`
	Width        int    `json:"Width"`
	Height       int    `json:"Height"`
	AutoTrans    int    `json:"-"`
	As           string `json:"-"`
	ShortSide    int    `json:"-"`
	LongSide     int    `json:"-"`
	Roi          bool   `json:"-"`
}
type DescribeTranscodePresetDetailOutput struct {
	PresetDetailList []*TranscodePreset `json:"PresetDetailList"`
}
type DescribeTranscodePresetDetailResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeTranscodePresetDetailOutput `json:"Result,omitempty"`
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
type DescribeSnapshotPresetOutput struct {
	PresetList []string `json:"PresetList"`
}
type DescribeSnapshotPresetResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeSnapshotPresetOutput `json:"Result,omitempty"`
}
type SnapshotPreset struct {
	Preset         string          `json:"Preset"`
	Status         int             `json:"Status"`
	Interval       int             `json:"Interval"`
	Bucket         string          `json:"Bucket"`
	AccessKey      string          `json:"AccessKey"`
	CallbackDetail *CallbackDetail `json:"CallbackDetail"`
}
type DescribeSnapshotPresetDetailOutput struct {
	PresetDetailList []*SnapshotPreset `json:"PresetDetailList"`
}
type DescribeSnapshotPresetDetailResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeSnapshotPresetDetailOutput `json:"Result,omitempty"`
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

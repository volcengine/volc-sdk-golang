package imagex

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ActionRefresh    = 0
	ActionDisable    = 1
	ActionEnable     = 2
	ActionPreload    = 4
	ActionRefreshDir = 5

	SEGMENT_CLASS_GENERAL    = "general"
	SEGMENT_CLASS_HUMAN      = "human"
	SEGMENT_CLASS_PRODUCT    = "product"
	SEGMENT_CLASS_HUMAN_V2   = "humanv2"
	SEGMENT_CLASS_PRODUCT_V2 = "productv2"

	SMARTCROP_POLICY_DEMOTION_CENTER = "center"
	SMARTCROP_POLICY_DEMOTION_FGLASS = "fglass"
	SMARTCROP_POLICY_DEMOTION_TOP    = "top"

	SMARTCROP_SCENE_NORMAL  = "normal"
	SMARTCROP_SCENE_CARTOON = "cartoon"

	FunctionEncryption = "Encryption"
)

// GetAllImageServices
type GetServicesResult struct {
	Services []Service `json:"Services"`
}

// GetImageService
type Service struct {
	ServiceName    string   `json:"ServiceName"`
	ServiceId      string   `json:"ServiceId"`
	ServiceRegion  string   `json:"ServiceRegion"`
	CreateAt       string   `json:"CreateAt"`
	ServiceStatus  string   `json:"ServiceStatus"`
	HasSigkey      bool     `json:"HasSigkey"`
	TemplatePrefix string   `json:"TemplatePrefix"`
	DomainInfos    []Domain `json:"DomainInfos"`
	PrimaryKey     string   `json:"PrimaryKey"`
	SecondaryKey   string   `json:"SecondaryKey"`
	ObjectAccess   bool     `json:"ObjectAccess"`
	CompactURL     bool     `json:"CompactURL"`
	Mirror         Mirror   `json:"Mirror"`
	Storage        Storage  `json:"Storage"`
	AllowBkts      []string `json:"AllowBkts"`
}

type Mirror struct {
	Schema  string            `json:"Schema"`
	Host    string            `json:"Host"`
	Source  string            `json:"Source"`
	Headers map[string]string `json:"Headers"`
}

type Domain struct {
	DomainName string `json:"DomainName"`
	CNAME      string `json:"CNAME"`
	Status     string `json:"Status"`
	IsDefault  bool   `json:"IsDefault"`
}

type Storage struct {
	BktName  string `json:"BktName"`
	TTL      int    `json:"TTL"`
	AllTypes bool   `json:"AllTypes"`
}

// GetServiceDomains
type DomainResult struct {
	Domain      string      `json:"domain"`
	CNAME       string      `json:"cname"`
	Status      string      `json:"status"`
	HttpsConfig HttpsConfig `json:"https_config"`
	IsDefault   bool        `json:"is_default"`
	Resolved    bool        `json:"resolved"`
}

type HttpsConfig struct {
	EnableHttps bool   `json:"enable_https"`
	ForceHttps  bool   `json:"force_https"`
	CertId      string `json:"cert_id"`
}

type AccessControlConfig struct {
	ReferLink ReferLink `json:"refer_link"`
}

type ReferLink struct {
	Enabled         bool     `json:"enabled"`
	IsWhiteMode     bool     `json:"is_white_mode"`
	Values          []string `json:"values"`
	AllowEmptyRefer bool     `json:"allow_empty_refer"`
}

type RespHdr struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

// DeleteImageUploadFiles
type DeleteImageParam struct {
	StoreUris []string `json:"StoreUris"`
}

type DeleteImageResult struct {
	ServiceId    string   `json:"ServiceId"`
	DeletedFiles []string `json:"DeletedFiles"`
}

// ApplyImageUpload
type ApplyUploadImageParam struct {
	ServiceId   string
	SessionKey  string
	UploadNum   int
	StoreKeys   []string
	CommitParam *CommitUploadImageParam
}

type ApplyUploadImageResult struct {
	UploadAddress UploadAddress `json:"UploadAddress"`
	RequestId     string        `json:"RequestId"`
}

type UploadAddress struct {
	SessionKey  string      `json:"SessionKey"`
	UploadHosts []string    `json:"UploadHosts"`
	StoreInfos  []StoreInfo `json:"StoreInfos"`
}

type StoreInfo struct {
	StoreUri string `json:"StoreUri"`
	Auth     string `json:"Auth"`
}

// CommitImageUpload
type CommitUploadImageParam struct {
	ServiceId   string     `json:"-"`
	SessionKey  string     `json:"SessionKey"`
	SuccessOids []string   `json:"SuccessOids"`
	Functions   []Function `json:"Functions"`
}

type Function struct {
	Name  string      `json:"Name"`
	Input interface{} `json:"Input"`
}

type EncryptionInput struct {
	Config       map[string]string `json:"Config"`
	PolicyParams map[string]string `json:"PolicyParams"`
}

type CommitUploadImageResult struct {
	Results    []Result    `json:"Results"`
	RequestId  string      `json:"RequestId"`
	ImageInfos []ImageInfo `json:"PluginResult"`
}

type Result struct {
	Uri        string     `json:"Uri"`
	UriStatus  int        `json:"UriStatus"` // 图片上传结果（2000:成功，2001:失败）需要传 SuccessOids 才会返回
	Encryption Encryption `json:"Encryption"`
}

type Encryption struct {
	Uri       string            `json:"Uri"`
	SecretKey string            `json:"SecretKey"`
	Algorithm string            `json:"Algorithm"`
	Version   string            `json:"Version"`
	SourceMd5 string            `json:"SourceMd5"`
	Extra     map[string]string `json:"Extra"`
}

type ImageInfo struct {
	FileName    string `json:"FileName"`
	ImageUri    string `json:"ImageUri"`
	ImageWidth  int    `json:"ImageWidth"`
	ImageHeight int    `json:"ImageHeight"`
	ImageMd5    string `json:"ImageMd5"`
	ImageFormat string `json:"ImageFormat"`
	ImageSize   int    `json:"ImageSize"`
	FrameCnt    int    `json:"FrameCnt"`
	Duration    int    `json:"Duration"`
}

// UpdateImageUploadFiles
type UpdateImageUrlPayload struct {
	Action    int      `json:"Action"`
	ImageUrls []string `json:"ImageUrls"`
}

// FetchImageUrl
type FetchUrlReq struct {
	Url           string              `json:"Url"`
	ServiceId     string              `json:"ServiceId"`
	StoreKey      string              `json:"StoreKey"`
	RequestHeader map[string][]string `json:"RequestHeader"`
	TimeOut       int                 `json:"TimeOut"`
}

type FetchUrlResp struct {
	Url      string `json:"Url"`
	StoreUri string `json:"StoreUri"`
	FSize    int    `json:"FSize"`
}

// GetImageStyleResult
type GetImageStyleResultReq struct {
	ServiceId     string
	StyleId       string            `json:"StyleId"`
	Params        map[string]string `json:"Params"`
	OutputFormat  string            `json:"OutputFormat"`
	OutputQuality int               `json:"OutputQuality"`
}

type GetImageStyleResultResp struct {
	ResUri       string `json:"ResUri"`
	RenderDetail []struct {
		Element string `json:"Element"`
		ErrMsg  string `json:"ErrMsg"`
	} `json:"RenderDetail"`
}

// GetImageOCR
type GetImageOCRLicenseResult struct {
	Scene     string                                 `json:"Scene"`
	OCRResult map[string]*GetImageOCRLicenseTextInfo `json:"OCR_result"`
}

type GetImageOCRLicenseTextInfo struct {
	Content  string `json:"Content"`
	Location []int  `json:"Location"`
}

type GetImageOCRGeneralResult struct {
	Scene     string                        `json:"Scene"`
	OCRResult []*GetImageOCRGeneralTextInfo `json:"OCR_result"`
}

type GetImageOCRGeneralTextInfo struct {
	Content    string  `json:"Content"`
	Confidence string  `json:"Confidence"`
	Location   [][]int `json:"Location"`
}

type GetImageOCRParam struct {
	ServiceId string
	StoreUri  string
	ImageUrl  string
}

// GetImageBgFillResult
type GetImageBgFillParam struct {
	ServiceId string  `json:"ServiceId"`
	StoreUri  string  `json:"StoreUri"`
	Model     int     `json:"Model"`
	Top       float64 `json:"Top"`
	Bottom    float64 `json:"Bottom"`
	Left      float64 `json:"Left"`
	Right     float64 `json:"Right"`
}

type GetImageBgFillResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageEnhanceResult
type GetImageEnhanceParam struct {
	ServiceId    string   `json:"ServiceId"`
	StoreUri     string   `json:"StoreUri"`
	Model        int      `json:"Model"`
	DisableAr    bool     `json:"DisableAr"`
	DisableSharp bool     `json:"DisableSharp"`
	Resolution   string   `json:"Resolution"`
	Actions      []string `json:"Actions"`
}

type GetImageEnhanceResult struct {
	ResUri string `json:"ResUri"`
	Method string `json:"Method"`
}

// GetImageEraseResult
type GetImageEraseParam struct {
	ServiceId string     `json:"ServiceId"`
	StoreUri  string     `json:"StoreUri"`
	Model     string     `json:"Model"`
	BBox      []EraseBox `json:"BBox"`
}

type EraseBox struct {
	X1 float64 `json:"X1"`
	Y1 float64 `json:"Y1"`
	X2 float64 `json:"X2"`
	Y2 float64 `json:"Y2"`
}

type GetImageEraseResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageQuality
type GetImageQualityParam struct {
	ServiceId string
	ImageUrl  string `json:"ImageUrl"`
	VqType    string `json:"VqType"`
}

type GetImageQualityResult struct {
	VqType   string                 `json:"VqType"`
	NrScores map[string]interface{} `json:"NrScores"`
}

// GetImageSegment
type GetImageSegmentParam struct {
	ServiceId string
	StoreUri  string   `json:"StoreUri"`
	Class     string   `json:"Class"`
	Refine    bool     `json:"Refine"`
	OutFormat string   `json:"OutFormat"`
	TransBg   bool     `json:"TransBg"`
	Contour   *Contour `json:"Contour"`
}

type Contour struct {
	Color string `json:"Color"`
	Size  int    `json:"Size"`
}

type GetImageSegmentResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageSuperResolutionResult
type GetImageSuperResolutionParam struct {
	ServiceId string  `json:"ServiceId"`
	StoreUri  string  `json:"StoreUri"`
	Multiple  float64 `json:"Multiple"`
}

type GetImageSuperResolutionResp struct {
	ResUri string `json:"ResUri"`
}

// GetImageUploadFile
type GetImageUploadFileParam struct {
	ServiceId string `json:"ServiceId"`
	StoreUri  string `json:"StoreUri"`
}

type GetImageUploadFileResult struct {
	ServiceId    string `json:"ServiceId"`
	StoreUri     string `json:"StoreUri"`
	LastModified string `json:"LastModified"`
	FileSize     int    `json:"FileSize"`
	Marker       int64  `json:"Marker"`
}

// GetImageUploadFiles
type GetImageUploadFilesParam struct {
	ServiceId string `json:"ServiceId"`
	Limit     int    `json:"Limit"`
	Marker    int64  `json:"Marker"`
}

type GetImageUploadFilesResult struct {
	ServiceId   string       `json:"ServiceId"`
	FileObjects []FileObject `json:"FileObjects"`
	Status      string       `json:"Status"`
	HasMore     bool         `json:"hasMore"`
}

type FileObject struct {
	StoreUri     string `json:"StoreUri"`
	Key          string `json:"Key"`
	LastModified string `json:"LastModified"`
	FileSize     int    `json:"FileSize"`
	Marker       int64  `json:"Marker"`
}

// GetImageDuplicateDetection
type GetImageDuplicateDetectionParam struct {
	Urls      []string `json:"Urls"`
	ServiceId string   `json:"-"`
}

type GetImageDuplicateDetectionAsyncParam struct {
	GetImageDuplicateDetectionParam
	Callback string `json:"Callback"`
}

type getImageDuplicateDetectionParam struct {
	Urls     []string `json:"Urls"`
	Async    bool     `json:"Async"`
	Callback string   `json:"Callback"`
}

type GetImageDuplicateDetectionResult struct {
	Score  string              `json:"Score"`
	Groups map[string][]string `json:"Groups"`
}

type GetImageDuplicateDetectionAsyncResult struct {
	TaskId   string `json:"TaskId"`
	Callback string `json:"Callback"`
}

type GetImageDuplicateDetectionCallbackBody struct {
	Status int                              `json:"Status"`
	TaskId string                           `json:"TaskId"`
	Result GetImageDuplicateDetectionResult `json:"Result"`
}

type GetDedupTaskStatusResult struct {
	Status int                              `json:"Status"`
	TaskId string                           `json:"TaskId"`
	Result GetImageDuplicateDetectionResult `json:"Result"`
}

// GetDenoisingImage
type GetDenoisingImageParam struct {
	ServiceId   string  `json:"-"`
	StoreUri    string  `json:"StoreUri"`
	ImageUrl    string  `json:"ImageUrl"`
	OutFormat   string  `json:"OutFormat"`
	Intensity   float64 `json:"Intensity"`
	CanDemotion bool    `json:"CanDemotion"`
}

type GetDenoisingImageResult struct {
	ResUri   string `json:"ResUri"`
	Demotion bool   `json:"Demotion"`
}

// GetImageComicResult
type GetImageComicParam struct {
	ServiceId string `json:"ServiceId"`
	StoreUri  string `json:"StoreUri"`
}

type GetImageComicResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageSmartCropResult
type GetImageSmartCropParam struct {
	ServiceId string  `json:"ServiceId"`
	StoreUri  string  `json:"StoreUri"`
	Policy    string  `json:"Policy"`
	Scene     string  `json:"Scene"`
	Sigma     float64 `json:"Sigma"`
	Width     int     `json:"Width"`
	Height    int     `json:"Height"`
}

type GetImageSmartCropResult struct {
	ResUri   string `json:"ResUri"`
	Demotion bool   `json:"Demotion"`
}

// GetLicensePlateDetection
type GetLicensePlateDetectionParam struct {
	ServiceId string `json:"-"`
	ImageUri  string `json:"ImageUri"`
}

type LocationsLower struct {
	X1 int `json:"x1"`
	X2 int `json:"x2"`
	Y1 int `json:"y1"`
	Y2 int `json:"y2"`
}

type GetLicensePlateDetectionResult struct {
	Locations []LocationsLower `json:"Locations"`
}

// GetImagePSDetection
type GetImagePSDetectionParam struct {
	ServiceId string `json:"-"`
	ImageUri  string `json:"ImageUri"`
	Method    string `json:"Method"`
}

type GetImagePSDetectionResult struct {
	ElaImage []byte  `json:"ElaImage"`
	NaImage  []byte  `json:"NaImage"`
	HeImage  []byte  `json:"HeImage"`
	Score    float64 `json:"Score"`
	Label    int     `json:"Label"`
}

// GetPrivateImageType
type GetPrivateImageTypeParam struct {
	ServiceId  string  `json:"-"`
	ImageUri   string  `json:"ImageUri"`
	Method     string  `json:"Method"`
	ThresFace  float64 `json:"ThresFace"`
	ThresCloth float64 `json:"ThresCloth"`
}

type GetPrivateImageTypeResult struct {
	Face  int `json:"Face"`
	Cloth int `json:"Cloth"`
}

// CreateImageHmEmbed
type CreateImageHmEmbedParam struct {
	ServiceId     string `json:"ServiceId"`
	StoreUri      string `json:"StoreUri"`
	Algorithm     string `json:"Algorithm"`
	Info          string `json:"Info"`
	OutFormat     string `json:"OutFormat"`
	OutQuality    int    `json:"OutQuality"`
	StrengthLevel string `json:"StrengthLevel"`
}

type CreateImageHmEmbedResult struct {
	StoreUri string `json:"StoreUri"`
}

// CreateImageHmExtract
type CreateImageHmExtractParam struct {
	ServiceId string `json:"ServiceId"`
	StoreUri  string `json:"StoreUri"`
	Algorithm string `json:"Algorithm"`
}

type CreateImageHmExtractResult struct {
	Info       string `json:"Info"`
	StatusCode int    `json:"StatusCode"`
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return fmt.Errorf("fail to marshal result, %v", err)
	}
	if err = json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

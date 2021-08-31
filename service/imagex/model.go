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
	ServiceId   string       `json:"-"`
	SessionKey  string       `json:"SessionKey"`
	SuccessOids []string     `json:"SuccessOids"`
	OptionInfos []OptionInfo `json:"OptionInfos"`
	Functions   []Function   `json:"Functions"`
}

type OptionInfo struct {
	StoreUri string `json:"StoreUri"`
	FileName string `json:"FileName"`
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
	Url       string `json:"Url"`
	ServiceId string `json:"ServiceId"`
	StoreKey  string `json:"StoreKey"`
}

type FetchUrlResp struct {
	Url      string `json:"Url"`
	StoreUri string `json:"StoreUri"`
	FSize    int    `json:"FSize"`
}

//GetImageOCR
type GetImageOCRResult struct {
	Scene     string                          `json:"Scene"`
	OCRResult map[string]*GetImageOCRTextInfo `json:"OCR_result"`
}

type GetImageOCRTextInfo struct {
	Content  string `json:"Content"`
	Location []int  `json:"Location"`
}

type GetImageOCRParam struct {
	ServiceId string
	Scene     string
	StoreUri  string
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

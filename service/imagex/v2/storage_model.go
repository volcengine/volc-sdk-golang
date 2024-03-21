package imagex

// ApplyImageUpload
type ApplyUploadImageParam struct {
	ServiceId   string
	SessionKey  string
	UploadNum   int
	StoreKeys   []string
	CommitParam *CommitUploadImageParam
	SkipMeta    bool
	SkipCommit  bool
	Overwrite   bool
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
	SkipMeta    bool       `json:"-"`
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
	PutError   *PutError  `json:"-"` // 上传阶段错误
}

type PutError struct {
	ErrorCode int
	Error     string
	Message   string
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

type UploadPolicy struct {
	ContentTypeBlackList []string `json:"ContentTypeBlackList,omitempty"`
	ContentTypeWhiteList []string `json:"ContentTypeWhiteList,omitempty"`
	FileSizeUpLimit      string   `json:"FileSizeUpLimit,omitempty"`
	FileSizeBottomLimit  string   `json:"FileSizeBottomLimit,omitempty"`
}

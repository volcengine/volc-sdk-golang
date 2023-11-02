package rtc

import (
	"encoding/json"

	"github.com/volcengine/volc-sdk-golang/base"
)

// CommonResponse ... need to decode result by type
type CommonResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           json.RawMessage `json:"Result,omitempty"`
}

type (
	// StartRecordRequest ...
	StartRecordRequest struct {
		AppId            string            `json:"AppId"`
		BusinessId       string            `json:"BusinessId"`
		RoomId           string            `json:"RoomId"`
		TaskId           string            `json:"TaskId"`
		RecordMode       uint32            `json:"RecordMode"`
		Encode           *Encode           `json:"Encode,omitempty"`
		FileFormatConfig *FileFormatConfig `json:"FileFormatConfig,omitempty"`
		StorageConfig    StorageConfig     `json:"StorageConfig"`
	}

	// StartRecordResponse ...
	StartRecordResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           string `json:"Result,omitempty"`
	}

	Encode struct {
		VideoWidth   uint32 `json:"VideoWidth"`
		VideoHeight  uint32 `json:"VideoHeight"`
		VideoFps     uint32 `json:"VideoFps"`
		VideoBitrate uint32 `json:"VideoBitrate"`
	}

	FileFormatConfig struct {
		FileFormat []string
	}

	StorageConfig struct {
		Type         uint32        `json:"Type"`
		TosConfig    *TosConfig    `json:"TosConfig,omitempty"`
		CustomConfig *CustomConfig `json:"CustomConfig,omitempty"`
	}

	TosConfig struct {
		AccountId string `json:"AccountId"`
		Region    uint32 `json:"Region"`
		Bucket    string `json:"Bucket"`
	}

	CustomConfig struct {
		Vendor    uint32 `json:"Vendor"`
		Region    string `json:"Region"`
		Bucket    string `json:"Bucket"`
		AccessKey string `json:"AccessKey"`
		SecretKey string `json:"SecretKey"`
	}
)

type (
	// GetRecordTaskResponse ...
	GetRecordTaskResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *GetRecordTaskResult `json:"Result,omitempty"`
	}

	GetRecordTaskResult struct {
		RecordTask RecordTask `json:"RecordTask"`
	}

	RecordTask struct {
		StartTime      uint64       `json:"StartTime"`
		EndTime        uint64       `json:"EndTime"`
		Status         uint64       `json:"Status"`
		StopReason     string       `json:"StopReason"`
		RecordFileList []RecordFile `json:"RecordFileList"`
	}

	RecordFile struct {
		Vid         string   `json:"Vid"`
		ObjectKey   string   `json:"ObjectKey"`
		Duration    uint64   `json:"Duration"`
		Size        uint64   `json:"Size"`
		StartTime   uint64   `json:"StartTime"`
		StreamList  []Stream `json:"StreamList"`
		VideoCodec  string   `json:"VideoCodec"`
		AudioCodec  string   `json:"AudioCodec"`
		VideoWidth  int      `json:"VideoWidth"`
		VideoHeight int      `json:"VideoHeight"`
	}

	Stream struct {
		Index      uint32 `json:"Index"`
		StreamType uint32 `json:"StreamType"`
		UserId     string `json:"UserId"`
	}
)

type (
	StartWebRecordRequest struct {
		AppId           string
		TaskId          string
		InputURL        string
		MaxRecordSecond int
		JsCommand       []string
		Bucket          string
		VideoSpace      string
		VideoInfo       WebVideoInfo
		PageInfo        WebPageInfo
		Duration        int
	}

	WebVideoInfo struct {
		ResolutionWidth  int
		ResolutionHeight int
	}

	WebPageInfo struct {
		PageWidth  int
		PageHeight int
	}

	StartWebRecordResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *StartWebRecordResult `json:"Result,omitempty"`
	}

	StartWebRecordResult struct {
		Message string `json:"Message"`
	}

	StopWebRecordRequest struct {
		AppId  string
		TaskId string
	}

	StopWebRecordResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *StopWebRecordResult `json:"Result,omitempty"`
	}

	StopWebRecordResult struct {
		Message string `json:"Message"`
	}

	GetWebRecordTaskResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *GetWebRecordTaskResult `json:"Result,omitempty"`
	}

	GetWebRecordTaskResult struct {
		Message   string        `json:"Message"`
		EventData TaskEventData `json:"EventData"`
	}

	TaskEventData struct {
		TaskId     string `json:"TaskId"`
		Status     int    `json:"Status"`
		CreateTime int    `json:"CreateTime"`
		FinishTime int    `json:"FinishTime"`
		Vid        string `json:"Vid"`
		FinalFile  File   `json:"FinalFile"`
		Files      []File `json:"Files"`
	}

	File struct {
		Index     int
		Bucket    string
		ObjectKey string
	}

	GetWebRecordListResponse struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *GetWebRecordListResult `json:"Result,omitempty"`
	}

	GetWebRecordListResult struct {
		Message       string        `json:"Message"`
		WebRecordList WebRecordList `json:"WebRecordList"`
	}

	WebRecordList struct {
		AppId      string  `json:"AppId"`
		Tasks      []Tasks `json:"Tasks"`
		PageNumber int     `json:"PageNumber"`
		PageSize   int     `json:"PageSize"`
		TotalCount int     `json:"TotalCount"`
	}

	Tasks struct {
		TaskId          string `json:"TaskId"`
		CreateTime      int    `json:"CreateTime"`
		FinishTime      int    `json:"FinishTime"`
		Status          int    `json:"Status"`
		InputURL        string `json:"InputURL"`
		VideoSpace      string `json:"VideoSpace"`
		Vid             string `json:"Vid"`
		MaxRecordSecond int    `json:"MaxRecordSecond"`
		Duration        int    `json:"Duration"`
		Bucket          string `json:"Bucket"`
		ObjectKey       string `json:"ObjectKey"`
	}
)

type StartWBRecordBody struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"AppId"`

	// REQUIRED; 需要录制的白板房间 ID，同一个 appId 中，为每个房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制任务 ID。你可以自行设定 TaskId 以区分不同的白板录制任务。 关于 TaskId 及以上 Id 字段的命名规则，参看ID [https://www.volcengine.com/docs/6348/69835#idname]。
	TaskID string `json:"TaskId"`

	// REQUIRED; 任务发起方的用户 ID。不能与房间中其他用户的 ID 重复，否则先进房的用户会被移出房间。
	UserID string `json:"UserId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 自定义 UI 的 Web 页面地址。 不填表示使用默认白板页面。
	SourceURL *string `json:"SourceURL,omitempty"`
}

type StartWBRecordRes struct {

	// REQUIRED
	ResponseMetadata StartWBRecordResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StartWBRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type StopWBRecordBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，同一个 appId 中，每个房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制任务 ID。调用StartWBRecord时使用的任务 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 调用接口的用户 ID
	UserID string `json:"UserId"`
}

type StopWBRecordRes struct {

	// REQUIRED
	ResponseMetadata base.ResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopWBRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type WbTranscodeCreateBody struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"app_id"`

	// REQUIRED; 用户 ID
	Operator string `json:"operator"`

	// REQUIRED; 需要转换为图片的文档链接地址。每次调用接口只能请求处理一份文档。
	Resource string `json:"resource"`

	// 转码成功后返回的 URL 预签名有效期。有效范围：0~604800。
	// * 【默认值】0：无限期。你需要在 TOS 服务的 bucket 为公共读。
	// * 1～604800：URL 预签名有效期，单位是秒。
	PreSignDuration *int32 `json:"pre_sign_duration,omitempty"`

	// 静态转码的转码优先级
	// * 【默认值】0: 非实时转码
	// * 1: 实时转码
	Priority *int32 `json:"priority,omitempty"`

	// 动态转码文件设置
	ResourceAttr *WbTranscodeCreateBodyResourceAttr `json:"resource_attr,omitempty"`

	// 对象存储属性
	// * 使用火山引擎的对象存储服务，且本次传入的参数与控制台设置的属性有差异，则以传入参数为准。
	StorageConfig *WbTranscodeCreateBodyStorageConfig `json:"storage_config,omitempty"`

	// 转码参数设置
	TranscodeConfig *WbTranscodeCreateBodyTranscodeConfig `json:"transcode_config,omitempty"`

	// 转码类型
	// * 0: 静态转码(默认)
	// * 1: 动态转码
	TranscodeMode *int32 `json:"transcode_mode,omitempty"`
}

// WbTranscodeCreateBodyResourceAttr - 动态转码文件设置
type WbTranscodeCreateBodyResourceAttr struct {

	// REQUIRED; 文件名
	FileName string `json:"file_name"`

	// REQUIRED; 文件大小，单位：byte
	Size int32 `json:"size"`
}

// WbTranscodeCreateBodyStorageConfig - 对象存储属性
// * 使用火山引擎的对象存储服务，且本次传入的参数与控制台设置的属性有差异，则以传入参数为准。
type WbTranscodeCreateBodyStorageConfig struct {

	// REQUIRED; 存储类型 【默认值】0: Tos1: 第三方对象存储接口，支持阿里云和亚马逊
	Type int32 `json:"type"`

	// 第三方对象存储服务参数设置
	CustomConfig *WbTranscodeCreateBodyStorageConfigCustomConfig `json:"custom_config,omitempty"`

	// 火山引擎的对象存储服务参数设置
	TosConfig *WbTranscodeCreateBodyStorageConfigTosConfig `json:"tos_config,omitempty"`
}

// WbTranscodeCreateBodyStorageConfigCustomConfig - 第三方对象存储服务参数设置
type WbTranscodeCreateBodyStorageConfigCustomConfig struct {

	// REQUIRED; Access Key
	AccessKey string `json:"access_key"`

	// REQUIRED; 桶名称
	Bucket string `json:"bucket"`

	// REQUIRED; 区域
	// {
	// "Custom": {
	// "AmazonS3": {
	// "EndPointFormat": "",
	// "RegionIDs": {
	// "0": "us-east-1",
	// "1": "us-east-2",
	// "2": "us-west-1",
	// "3": "us-west-2",
	// "4": "eu-west-1",
	// "5": "eu-west-2",
	// "6": "eu-west-3",
	// "8": "ap-southeast-1",
	// "10": "ap-northeast-1",
	// "11": "ap-northeast-2",
	// "12": "sa-east-1",
	// "13": "ca-central-1",
	// "14": "ap-south-1",
	// "18": "ap-east-1",
	// "19": "eu-south-1",
	// "20": "ap-northeast-3",
	// "21": "eu-north-1",
	// "22": "me-south-1",
	// "24": "af-south-1"
	// }
	// },
	// "AlicloudOSS": {
	// "EndPointFormat": "%v.aliyuncs.com",
	// "RegionIDs": {
	// "0": "oss-cn-hangzhou",
	// "1": "oss-cn-shanghai",
	// "2": "oss-cn-qingdao",
	// "3": "oss-cn-beijing",
	// "4": "oss-cn-zhangjiakou",
	// "5": "oss-cn-huhehaote",
	// "6": "oss-cn-wulanchabu",
	// "7": "oss-cn-shenzhen",
	// "8": "oss-cn-heyuan",
	// "9": "oss-cn-guangzhou",
	// "10": "oss-cn-chengdu",
	// "11": "oss-cn-hongkong",
	// "12": "oss-us-west-1",
	// "13": "oss-us-east-1",
	// "14": "oss-ap-southeast-1",
	// "15": "oss-ap-southeast-2",
	// "17": "oss-ap-southeast-5",
	// "18": "oss-ap-northeast-1",
	// "19": "oss-ap-south-1",
	// "20": "oss-eu-central-1",
	// "21": "oss-eu-west-1",
	// "22": "oss-me-east-1",
	// "23": "oss-ap-southeast-6"
	// }
	// }
	// }
	Region int32 `json:"region"`

	// REQUIRED; Secret Key AK/SK建议只开通写权限，关闭读权限。
	SecretKey string `json:"secret_key"`

	// REQUIRED; 第三方存储供应商0: Amazon 亚马逊1: Alicloud 阿里云
	Vendor int32 `json:"vendor"`
}

// WbTranscodeCreateBodyStorageConfigTosConfig - 火山引擎的对象存储服务参数设置
type WbTranscodeCreateBodyStorageConfigTosConfig struct {

	// REQUIRED; Bucket 所属的火山引擎账号 ID。在登录火山引擎后，可在头像的悬浮菜单中找到账号 ID。[https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_0819c44c6aadff358a7dfc52c5daab57.png]
	AccountID string `json:"account_id"`

	// REQUIRED; 桶名称。登录TOS 控制台 [https://console.volcengine.com/tos/bucket]开通和获取。
	Bucket string `json:"bucket"`

	// REQUIRED; 0: 北京
	Region int32 `json:"region"`
}

// WbTranscodeCreateBodyTranscodeConfig - 转码参数设置
type WbTranscodeCreateBodyTranscodeConfig struct {

	// REQUIRED; 输入文件类型1: ppt2: pptx3: doc4: docx5: pdf
	InputFormat int32 `json:"input_format"`

	// REQUIRED; 输出文件类型1: png2: jpg/jpeg
	OutputFormat int32 `json:"output_format"`

	// REQUIRED; 转码后的页面高度，单位为像素
	OutputHeight int32 `json:"output_height"`

	// REQUIRED; 转码后的页面宽度，单位为像素
	OutputWidth int32 `json:"output_width"`

	// 是否按照指定分辨率拉伸页面 默认为false，按照文件的原始宽高比适配指定分辨率。
	ForceUseResolution *bool `json:"force_use_resolution,omitempty"`

	// 是否生成缩略图，默认为false
	Thumbnail *bool `json:"thumbnail,omitempty"`

	// 缩略图分辨率高，默认为180
	ThumbnailHeight *int32 `json:"thumbnail_height,omitempty"`

	// 缩略图分辨率宽，默认为320
	ThumbnailWidth *int32 `json:"thumbnail_width,omitempty"`
}

type WbTranscodeCreateRes struct {

	// REQUIRED
	ResponseMetadata base.ResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *WbTranscodeCreateResResult `json:"Result,omitempty"`
}

type WbTranscodeCreateResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// WbTranscodeCreateResResult - 视请求的接口而定
type WbTranscodeCreateResResult struct {

	// REQUIRED; task_id 为查询任务进度和获取转码后链接的必填参数
	TaskID string `json:"task_id"`
}

type WbTranscodeGetQuery struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"app_id" query:"app_id"`

	// REQUIRED; 成功调用WbTranscodeQuery后返回的任务ID
	TaskID string `json:"task_id" query:"task_id"`
}

type WbTranscodeGetRes struct {

	// REQUIRED
	ResponseMetadata base.ResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result WbTranscodeGetResResult `json:"Result"`
}

type WbTranscodeGetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type WbTranscodeGetResResult struct {

	// REQUIRED
	FileName string `json:"file_name"`

	// REQUIRED
	H5URL string `json:"h5_url"`

	// REQUIRED
	Height int32 `json:"height"`

	// REQUIRED
	Images []WbTranscodeGetResResultImagesItem `json:"images"`

	// REQUIRED
	TranscodeMode int32 `json:"transcode_mode"`

	// REQUIRED
	Width int32 `json:"width"`
}

type WbTranscodeGetResResultImagesItem struct {

	// REQUIRED
	Img string `json:"img"`

	// REQUIRED
	PageID int32 `json:"page_id"`

	// REQUIRED
	ThumbnailURL string `json:"thumbnail_url"`
}

type WbTranscodeQueryQuery struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"app_id" query:"app_id"`

	// REQUIRED; 成功调用WbTranscodeQuery后返回的任务ID
	TaskID string `json:"task_id" query:"task_id"`
}

type WbTranscodeQueryRes struct {

	// REQUIRED
	ResponseMetadata base.ResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result WbTranscodeQueryResResult `json:"Result"`
}

type WbTranscodeQueryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type WbTranscodeQueryResResult struct {

	// REQUIRED; 转码任务状态
	// * 0：未开始
	// * 1：进行中
	// * 2：已完成
	// * 3：失败
	Status int32 `json:"status"`

	// 错误码
	ErrCode *int32 `json:"err_code,omitempty"`

	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty"`

	// 静态转码进度。100表示已完成。 动态转码恒定为0。
	Progress *int32 `json:"progress,omitempty"`
}
type StartWBRecord struct{}
type StartWBRecordQuery struct{}
type WbTranscodeGet struct{}
type WbTranscodeGetBody struct{}
type WbTranscodeCreateQuery struct{}
type WbTranscodeCreate struct{}
type WbTranscodeQueryBody struct{}
type StopWBRecord struct{}
type StopWBRecordQuery struct{}
type StopWBRecordReq struct {
	*StopWBRecordQuery
	*StopWBRecordBody
}
type WbTranscodeGetReq struct {
	*WbTranscodeGetQuery
	*WbTranscodeGetBody
}
type StartWBRecordReq struct {
	*StartWBRecordQuery
	*StartWBRecordBody
}
type WbTranscodeCreateReq struct {
	*WbTranscodeCreateQuery
	*WbTranscodeCreateBody
}
type WbTranscodeQueryReq struct {
	*WbTranscodeQueryQuery
	*WbTranscodeQueryBody
}

package rtc_v20220601

type Event20231101 struct {

	// 原因，只有error事件有
	Cause *string `json:"Cause,omitempty"`

	// 事件描述
	Description *string `json:"Description,omitempty"`

	// 原始日志信息，这一期不做在前端展示
	Detail *string `json:"Detail,omitempty"`

	// 展示名
	DisplayName *string `json:"DisplayName,omitempty"`

	// 事件级别 Info、Warning、Error
	Level *int64 `json:"Level,omitempty"`

	// 事件名，英文
	Name *string `json:"Name,omitempty"`

	// 子事件，构成该聚合事件的事件，这一期不做在前端展示
	SubEvents []*Event20231101SubEventsItem `json:"SubEvents,omitempty"`

	// 标签
	Tags []*string `json:"Tags,omitempty"`

	// 事件时间，若为聚合类事件，取最大时间
	Time *int64 `json:"Time,omitempty"`
}

type Event20231101SubEventsItem struct {

	// 原因，只有error事件有
	Cause *string `json:"Cause,omitempty"`

	// 事件描述
	Description *string `json:"Description,omitempty"`

	// 原始日志信息，这一期不做在前端展示
	Detail *string `json:"Detail,omitempty"`

	// 展示名
	DisplayName *string `json:"DisplayName,omitempty"`

	// 事件级别 Info、Warning、Error
	Level *int64 `json:"Level,omitempty"`

	// 事件名，英文
	Name *string `json:"Name,omitempty"`

	// 子事件，构成该聚合事件的事件，这一期不做在前端展示
	SubEvents []*Event20231101 `json:"SubEvents,omitempty"`

	// 标签
	Tags []*string `json:"Tags,omitempty"`

	// 事件时间，若为聚合类事件，取最大时间
	Time *int64 `json:"Time,omitempty"`
}

type GetRealtimePostProcessingQuery struct {

	// REQUIRED
	AccountID string `json:"AccountID" query:"AccountID"`

	// REQUIRED
	AppID string `json:"AppID" query:"AppID"`

	// REQUIRED
	BusinessID string `json:"BusinessID" query:"BusinessID"`

	// REQUIRED
	End int32 `json:"end" query:"end"`

	// REQUIRED
	IndicatorName string `json:"IndicatorName" query:"IndicatorName"`

	// REQUIRED
	IsDemo bool `json:"IsDemo" query:"IsDemo"`

	// REQUIRED
	Page int32 `json:"page" query:"page"`

	// REQUIRED
	Size int32 `json:"size" query:"size"`

	// REQUIRED
	Start int32 `json:"start" query:"start"`

	// REQUIRED
	Timestamp int32 `json:"timestamp" query:"timestamp"`
}

type GetRealtimePostProcessingRes struct {

	// REQUIRED
	ResponseMetadata GetRealtimePostProcessingResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetRealtimePostProcessingResResult `json:"Result,omitempty"`
}

type GetRealtimePostProcessingResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *GetRealtimePostProcessingResResponseMetadataError `json:"Error,omitempty"`
}

// GetRealtimePostProcessingResResponseMetadataError - 仅在请求失败时返回。
type GetRealtimePostProcessingResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// GetRealtimePostProcessingResResult - 视请求的接口而定
type GetRealtimePostProcessingResResult struct {

	// REQUIRED
	Common GetRealtimePostProcessingResResultCommon `json:"common"`
}

type GetRealtimePostProcessingResResultCommon struct {

	// REQUIRED
	Abnormal int32 `json:"abnormal"`

	// REQUIRED
	Timestamp int32 `json:"timestamp"`

	// REQUIRED
	Value float32 `json:"value"`
}

type GetRecordTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的云端录制任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata GetRecordTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetRecordTaskResResult          `json:"Result,omitempty"`
}

type GetRecordTaskResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *GetRecordTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetRecordTaskResResponseMetadataError - 仅在请求失败时返回。
type GetRecordTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetRecordTaskResResult struct {

	// REQUIRED; 录制任务信息
	RecordTask GetRecordTaskResResultRecordTask `json:"RecordTask"`
}

// GetRecordTaskResResultRecordTask - 录制任务信息
type GetRecordTaskResResultRecordTask struct {

	// 任务结束的时间，为 Unix 时间戳，单位毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 录制生成的文件列表。
	RecordFileList []*GetRecordTaskResResultRecordTaskRecordFileListItem `json:"RecordFileList,omitempty"`

	// 任务开始的时间，为 Unix 时间戳，单位毫秒。
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态:
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int32 `json:"Status,omitempty"`

	// 任务停止的原因：空：表示任务未结束UnknownStopReason：未知停止原因StopByAPI：用户主动通过 API 停止StartTaskFailed：任务启动失败IdleTimeOut：超过了最大空闲时间
	StopReason *string `json:"StopReason,omitempty"`
}

type GetRecordTaskResResultRecordTaskRecordFileListItem struct {

	// 音频录制编码器
	AudioCodec *string `json:"AudioCodec,omitempty"`

	// 文件时长，单位毫秒。
	Duration *int64 `json:"Duration,omitempty"`

	// 文件在对象存储平台中的完整路径，如abc/efg/123.mp4。仅在你选择配置存储到对象存储平台时，此参数有效。
	ObjectKey *string `json:"ObjectKey,omitempty"`

	// 文件大小，单位字节。
	Size *int64 `json:"Size,omitempty"`

	// 当前录制文件创建的时间，为 Unix 时间戳，单位毫秒。
	StartTime *int64 `json:"StartTime,omitempty"`

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetRecordTaskResResultRecordTaskRecordFileListPropertiesItemsItem `json:"StreamList,omitempty"`

	// 文件在点播平台的唯一标识。你可以根据 vid 可以在点播平台上找到对应的文件。仅在你选择配置存储到 Vod 平台时，此参数有效。
	Vid *string `json:"Vid,omitempty"`

	// 视频录制编码协议
	VideoCodec *string `json:"VideoCodec,omitempty"`

	// 录制视频高度，单位像素。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 录制视频宽度，单位像素。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type GetRecordTaskResResultRecordTaskRecordFileListPropertiesItemsItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetSegmentTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的音频切片任务 ID。自动切片任务下，该字段可传入用户 UserId。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetSegmentTaskRes struct {

	// REQUIRED
	ResponseMetadata GetSegmentTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetSegmentTaskResResult          `json:"Result,omitempty"`
}

type GetSegmentTaskResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *GetSegmentTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetSegmentTaskResResponseMetadataError - 仅在请求失败时返回。
type GetSegmentTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetSegmentTaskResResult struct {

	// 音频切片任务信息
	SegmentTask *GetSegmentTaskResResultSegmentTask `json:"SegmentTask,omitempty"`
}

// GetSegmentTaskResResultSegmentTask - 音频切片任务信息
type GetSegmentTaskResResultSegmentTask struct {
	Control *GetSegmentTaskResResultSegmentTaskControl `json:"Control,omitempty"`

	// 每个音频切片的时长。
	Duration *int32 `json:"Duration,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 是否在开启音视频切片时，立刻开始切片。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。
	Identifier *string `json:"Identifier,omitempty"`

	// 任务最大的空闲超时时间。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因- 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：自动切片任务中，切片任务对应的客户端用户主动退房。
	StopReason    *string                                          `json:"StopReason,omitempty"`
	StorageConfig *GetSegmentTaskResResultSegmentTaskStorageConfig `json:"StorageConfig,omitempty"`

	// 流的类型
	TargetStreams *GetSegmentTaskResResultSegmentTaskTargetStreams `json:"TargetStreams,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskControl struct {

	// 是否开启切片对齐功能。默认为 False。你可以使用音频切片对齐功能，对齐各个用户音频切片的开始和结束时刻。
	// * 当 Align=False 时，关闭音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。如果用户在切片的周期中，有部分时间未发布音频，返回的音频切片时长会小于切片周期。各个用户音频切片开始时间不一定一致。
	// * 当 Align=True 时，开启音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。切片长度和切片周期相等，且各个用户音频切片开始的时间一致。如果用户在切片的周期中，有部分时间未发布音频，切片长度不变，这段时间呈现静音帧。如果用户在某个切片周期中始终没有发布音频，则不生成音频切片。
	Align *bool `json:"Align,omitempty"`

	// 是否忽略静音切片。默认值为 false
	IgnoreSilence *bool `json:"IgnoreSilence,omitempty"`

	// 是否开启合流切片功能。默认为 False。
	// * 当 Mixed=False 时，只会对 TargetStreams 中指定的音频流分别切片。
	// * 当 Mixed=True 时，除了会对 TargetStreams 中指定的音频流分别切片，还会对指定的音频流进行混音，生成合流切片，合流切片对应的用户名为 mixed。此时，任务创建后，不管是否有人上麦，会持续回调混音切片。
	// 不同平台的回调参看：
	// 操作 ANDROID API IOS API WINDOWS API
	// 本地麦克风录制和远端所有用户混音后的音频数据回调 onMixedAudioFrame [70081#onmixedaudioframe] onMixedAudioFrame: [70087#onmixedaudioframe] onMixedAudioFrame
	// [70096#onmixedaudioframe]
	Mixed *bool `json:"Mixed,omitempty"`

	// 冗余切片时长，单位为毫秒。
	// 当前 RTC 按照传入的Duration值进行固定时长切片,可能存在敏感词被截断，未被识别的情况。此时你可以添加冗余切片，即上一段切片的末尾指定时长，来避免此情况，此时切片的时长变为RedundantDuration + Duration。
	// 例如：当 Duration = 20，redundantDuration = 3 时，最终输出的前三个文件时长为：[0-20], [17-40],
	// [37-60]。
	RedundantDuration *int32 `json:"RedundantDuration,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfig struct {
	CustomConfig *GetSegmentTaskResResultSegmentTaskStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *GetSegmentTaskResResultSegmentTaskStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 第三方存储平台(需支持 S3 协议)
	// * 3: VeImageX (当前仅支持抽帧截图功能) 默认值为 0。
	Type           *int32                                                         `json:"Type,omitempty"`
	VeImageXConfig *GetSegmentTaskResResultSegmentTaskStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *GetSegmentTaskResResultSegmentTaskStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。
	// * 0： Amazon S3
	// * 1： 阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。- 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。
	// 你也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

// GetSegmentTaskResResultSegmentTaskTargetStreams - 流的类型
type GetSegmentTaskResResultSegmentTaskTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个StreamList中，Stream.Index不能重复。
	StreamList []*GetSegmentTaskResResultSegmentTaskTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetSnapshotTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的抽帧截图任务 ID。自动抽帧任务下，该字段可传入用户 UserId。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata GetSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetSnapshotTaskResResult          `json:"Result,omitempty"`
}

type GetSnapshotTaskResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *GetSnapshotTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetSnapshotTaskResResponseMetadataError - 仅在请求失败时返回。
type GetSnapshotTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetSnapshotTaskResResult struct {

	// 抽帧截图任务信息
	SnapshotTask *GetSnapshotTaskResResultSnapshotTask `json:"SnapshotTask,omitempty"`
}

// GetSnapshotTaskResResultSnapshotTask - 抽帧截图任务信息
type GetSnapshotTaskResResultSnapshotTask struct {

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime     *int64                                           `json:"EndTime,omitempty"`
	ImageConfig *GetSnapshotTaskResResultSnapshotTaskImageConfig `json:"ImageConfig,omitempty"`

	// 任务最大的空闲超时时间。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：自动截图任务中，截图任务对应的客户端用户主动退房。
	StopReason    *string                                            `json:"StopReason,omitempty"`
	StorageConfig *GetSnapshotTaskResResultSnapshotTaskStorageConfig `json:"StorageConfig,omitempty"`

	// 流的类型
	TargetStreams *GetSnapshotTaskResResultSnapshotTaskTargetStreams `json:"TargetStreams,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskImageConfig struct {

	// 图片的格式。值可取0或1，默认为0。选择0时，图片格式为 JEPG；选择1时，图片格式为 PNG。默认值为0。值不合法时，自动调整为默认值。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，单位为像素，取值范围为[0, 1920]，默认值为0，此时，和视频流的实际高度相同。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，单位为秒，取值范围为[1, 600]，默认值为2。值不合法时，自动调整为默认值。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，单位为像素，取值范围为[0, 1920]。默认值为0，此时，和视频流的实际宽度相同。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfig struct {
	CustomConfig *GetSnapshotTaskResResultSnapshotTaskStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *GetSnapshotTaskResResultSnapshotTaskStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 第三方存储平台(需支持 S3 协议)
	// * 3: VeImageX (当前仅支持抽帧截图功能) 默认值为 0。
	Type           *int32                                                           `json:"Type,omitempty"`
	VeImageXConfig *GetSnapshotTaskResResultSnapshotTaskStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *GetSnapshotTaskResResultSnapshotTaskStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。
	// * 0： Amazon S3
	// * 1： 阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。- 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。
	// 你也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

// GetSnapshotTaskResResultSnapshotTaskTargetStreams - 流的类型
type GetSnapshotTaskResResultSnapshotTaskTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个StreamList中，Stream.Index不能重复。
	StreamList []*GetSnapshotTaskResResultSnapshotTaskTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type IndicatorTag struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                      `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTagSubTagsItem `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTag202201 struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                            `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag202201SubTagsItem `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTag202201SubTagsItem struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                 `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag202201 `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTagSubTagsItem struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool           `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type StartRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED
	StorageConfig StartRecordBodyStorageConfig `json:"StorageConfig"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 配置选项
	Control *StartRecordBodyControl `json:"Control,omitempty"`
	Encode  *StartRecordBodyEncode  `json:"Encode,omitempty"`

	// 流的类型
	ExcludeStreams   *StartRecordBodyExcludeStreams   `json:"ExcludeStreams,omitempty"`
	FileFormatConfig *StartRecordBodyFileFormatConfig `json:"FileFormatConfig,omitempty"`
	FileNameConfig   *StartRecordBodyFileNameConfig   `json:"FileNameConfig,omitempty"`
	Layout           *StartRecordBodyLayout           `json:"Layout,omitempty"`

	// 使用此参数设定录制模式：单流/合流录制。0 表示合流录制，1 表示单流录制。默认值为 0。
	RecordMode *int32 `json:"RecordMode,omitempty"`

	// 流的类型
	TargetStreams *StartRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

// StartRecordBodyControl - 配置选项
type StartRecordBodyControl struct {

	// 选择补帧模式。默认值为0，可以取0和1。0为补最后一帧，1为补黑帧。值不合法时会自动调整为默认值。
	// 自动布局模式下，没有补帧的逻辑。
	// 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系:
	// 你可以在 Region 中传入 Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// （仅对录制有效）最大录制时长，单位为秒。默认值为 0。0 表示不限制录制时长。
	MaxRecordTime *int32 `json:"MaxRecordTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type StartRecordBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
	// 当AudioProfile=1时：
	// * 若输入参数取值范围为 [32,128]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [128,192]，编码码率固定为128。
	// 当AudioProfile=2时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。
	// * 1：单声道
	// * 2：双声道。 默认值为 2。值不合法时，自动调整为默认值。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。默认值为 0，此时使用 aac 编码协议。目前只支持 aac。值不合法时，自动调整为默认值。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型，在使用 aac 编码时生效。取值范围为 {0, 1, 2}。
	// * 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。取 2 时，只支持输出流音频声道数为双声道。 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。默认值 48000，取值为 [32000,44100,48000]，单位是 Hz。值不合法时，自动调整为默认值。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 输出视频码率。取值范围 [1,10000]，单位为 Kbps，默认值为自适应。值不合法时，自动调整为默认值。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。默认值为 0，可以取 0 或 1。取 0 时使用 H.264,取 1 时使用 ByteVC1 编码器。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出视频帧率。默认为 15，取值范围为 [1,60]。值不合法时，自动调整为默认值。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。默认为 4，取值范围为 [1,5]，单位为秒。值不合法时，自动调整为默认值。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 输出画面的高度，范围为[2, 1920]，必须是偶数，默认值为480。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面高度以 canvas 中的 Height 为主。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 输出画面的宽度。默认值为640，范围为 [2, 1920]，必须是偶数。值不合法时，自动调整为默认值。自定义布局下此参数不生效，整体画面宽度以 canvas 中的 Width 为主。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

// StartRecordBodyExcludeStreams - 流的类型
type StartRecordBodyExcludeStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个StreamList中，Stream.Index不能重复。
	StreamList []*StartRecordBodyExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordBodyFileFormatConfig struct {

	// 存储文件格式。可取值为：MP4、HLS、FLV、MP3、 AAC、M4A。可同时选择多个存储文件格式，最终生成所选存储格式对应的多个文件。
	// MP3、AAC和M4A仅在编码纯音频时有效。
	FileFormat []*string `json:"FileFormat,omitempty"`
}

type StartRecordBodyFileNameConfig struct {

	// 自定义录制文件名模式，具体参看自定义录制文件名 [106873]。
	Pattern *string `json:"Pattern,omitempty"`

	// 制定录制文件名的前缀，对TosConfig和CustomConfig生效，具体参看注释1
	Prefix []*string `json:"Prefix,omitempty"`
}

type StartRecordBodyLayout struct {
	CustomLayout *StartRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                `json:"LayoutMode,omitempty"`
	MainVideoStream *StartRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type StartRecordBodyLayoutCustomLayout struct {
	Canvas *StartRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即
	// Regions.Region.StreamIndex 要与
	// TargetStreams.StreamList.Stream.Index 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*StartRecordBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type StartRecordBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type StartRecordBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：- 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 值不合法时，自动调整为默认值。
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartRecordBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordBodyStorageConfig struct {
	CustomConfig *StartRecordBodyStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *StartRecordBodyStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 第三方存储平台(需支持 S3 协议)
	// * 3: VeImageX (当前仅支持抽帧截图功能) 默认值为 0。
	Type           *int32                                      `json:"Type,omitempty"`
	VeImageXConfig *StartRecordBodyStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *StartRecordBodyStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type StartRecordBodyStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。
	// * 0： Amazon S3
	// * 1： 阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type StartRecordBodyStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartRecordBodyStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。- 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。
	// 你也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartRecordBodyStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

// StartRecordBodyTargetStreams - 流的类型
type StartRecordBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个StreamList中，Stream.Index不能重复。
	StreamList []*StartRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordRes struct {

	// REQUIRED
	ResponseMetadata StartRecordResResponseMetadata `json:"ResponseMetadata"`

	// 仅在请求成功时返回"ok",失败时为空
	Result *string `json:"Result,omitempty"`
}

type StartRecordResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *StartRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StartRecordResResponseMetadataError - 仅在请求失败时返回。
type StartRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartSegmentBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED
	StorageConfig StartSegmentBodyStorageConfig `json:"StorageConfig"`

	// REQUIRED; 截图任务 ID。你可以自行设定TaskId以区分不同的切片任务，且在后续更新和结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识。添加 BusinessId 后，你可以在控制台上查看根据 BusinessId 划分的质量和用量数据。
	BusinessID *string                  `json:"BusinessId,omitempty"`
	Control    *StartSegmentBodyControl `json:"Control,omitempty"`

	// 每个音频切片的时长。单位为秒。范围为[1, 600]，默认值为20秒。值不合法时，自动调整为默认值。
	Duration *int32 `json:"Duration,omitempty"`

	// 是否在开启音视频切片时，立刻开始切片。 true 为在开启音视频切片时，立刻开启切片；false 时，在开启音视频切片时，不开始切片，热启动任务。 默认值为true。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。 切片文件名由 Identifier + UserId + 时间戳 + 序号组成。默认情况下 Identifier 为 随机生成的 UUID。在自定义文件名时，Identifier 的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。
	// 在自定义文件名时，你需确保 Identifier 命名全局唯一，否则在 TOS 平台会因文件名重复被覆盖。
	Identifier *string `json:"Identifier,omitempty"`

	// 任务最大的空闲超时时间。 如果切片任务订阅的所有流都已停止发布，那么任务会在空闲时间超过设定值后自动停止。值的范围为[1, 86400] ，单位为秒。默认值为 180秒。值不合法时，自动调整为默认值。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型
	TargetStreams *StartSegmentBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type StartSegmentBodyControl struct {

	// 是否开启切片对齐功能。默认为 False。你可以使用音频切片对齐功能，对齐各个用户音频切片的开始和结束时刻。
	// * 当 Align=False 时，关闭音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。如果用户在切片的周期中，有部分时间未发布音频，返回的音频切片时长会小于切片周期。各个用户音频切片开始时间不一定一致。
	// * 当 Align=True 时，开启音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。切片长度和切片周期相等，且各个用户音频切片开始的时间一致。如果用户在切片的周期中，有部分时间未发布音频，切片长度不变，这段时间呈现静音帧。如果用户在某个切片周期中始终没有发布音频，则不生成音频切片。
	Align *bool `json:"Align,omitempty"`

	// 是否忽略静音切片。默认值为 false
	IgnoreSilence *bool `json:"IgnoreSilence,omitempty"`

	// 是否开启合流切片功能。默认为 False。
	// * 当 Mixed=False 时，只会对 TargetStreams 中指定的音频流分别切片。
	// * 当 Mixed=True 时，除了会对 TargetStreams 中指定的音频流分别切片，还会对指定的音频流进行混音，生成合流切片，合流切片对应的用户名为 mixed。此时，任务创建后，不管是否有人上麦，会持续回调混音切片。
	// 不同平台的回调参看：
	// 操作 ANDROID API IOS API WINDOWS API
	// 本地麦克风录制和远端所有用户混音后的音频数据回调 onMixedAudioFrame [70081#onmixedaudioframe] onMixedAudioFrame: [70087#onmixedaudioframe] onMixedAudioFrame
	// [70096#onmixedaudioframe]
	Mixed *bool `json:"Mixed,omitempty"`

	// 冗余切片时长，单位为毫秒。
	// 当前 RTC 按照传入的Duration值进行固定时长切片,可能存在敏感词被截断，未被识别的情况。此时你可以添加冗余切片，即上一段切片的末尾指定时长，来避免此情况，此时切片的时长变为RedundantDuration + Duration。
	// 例如：当 Duration = 20，redundantDuration = 3 时，最终输出的前三个文件时长为：[0-20], [17-40],
	// [37-60]。
	RedundantDuration *int32 `json:"RedundantDuration,omitempty"`
}

type StartSegmentBodyStorageConfig struct {
	CustomConfig *StartSegmentBodyStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *StartSegmentBodyStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 第三方存储平台(需支持 S3 协议)
	// * 3: VeImageX (当前仅支持抽帧截图功能) 默认值为 0。
	Type           *int32                                       `json:"Type,omitempty"`
	VeImageXConfig *StartSegmentBodyStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *StartSegmentBodyStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type StartSegmentBodyStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。
	// * 0： Amazon S3
	// * 1： 阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type StartSegmentBodyStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSegmentBodyStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。- 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。
	// 你也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSegmentBodyStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

// StartSegmentBodyTargetStreams - 流的类型
type StartSegmentBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个StreamList中，Stream.Index不能重复。
	StreamList []*StartSegmentBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartSegmentBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartSegmentRes struct {

	// REQUIRED
	ResponseMetadata StartSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type StartSegmentResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *StartSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// StartSegmentResResponseMetadataError - 仅在请求失败时返回。
type StartSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartSnapshotBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED
	StorageConfig StartSnapshotBodyStorageConfig `json:"StorageConfig"`

	// REQUIRED; 截图任务 ID。你可以自行设定TaskId以区分不同的切片任务，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID  *string                       `json:"BusinessId,omitempty"`
	ImageConfig *StartSnapshotBodyImageConfig `json:"ImageConfig,omitempty"`

	// 任务最大的空闲超时时间。如果抽帧截图任务订阅的所有流都已停止发布，那么任务会在空闲时间超过设定值后自动停止。值的范围为[1, 86400] ，单位为秒。默认值为 180秒。值不合法时，自动调整为默认值。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型
	TargetStreams *StartSnapshotBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type StartSnapshotBodyImageConfig struct {

	// 图片的格式。值可取0或1，默认为0。选择0时，图片格式为 JEPG；选择1时，图片格式为 PNG。默认值为0。值不合法时，自动调整为默认值。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，单位为像素，取值范围为[0, 1920]，默认值为0，此时，和视频流的实际高度相同。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，单位为秒，取值范围为[1, 600]，默认值为2。值不合法时，自动调整为默认值。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，单位为像素，取值范围为[0, 1920]。默认值为0，此时，和视频流的实际宽度相同。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type StartSnapshotBodyStorageConfig struct {
	CustomConfig *StartSnapshotBodyStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *StartSnapshotBodyStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 第三方存储平台(需支持 S3 协议)
	// * 3: VeImageX (当前仅支持抽帧截图功能) 默认值为 0。
	Type           *int32                                        `json:"Type,omitempty"`
	VeImageXConfig *StartSnapshotBodyStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *StartSnapshotBodyStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type StartSnapshotBodyStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。
	// * 0： Amazon S3
	// * 1： 阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type StartSnapshotBodyStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSnapshotBodyStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。- 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。
	// 你也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSnapshotBodyStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

// StartSnapshotBodyTargetStreams - 流的类型
type StartSnapshotBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个StreamList中，Stream.Index不能重复。
	StreamList []*StartSnapshotBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartSnapshotBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartSnapshotRes struct {

	// REQUIRED
	ResponseMetadata StartSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                          `json:"Result,omitempty"`
}

type StartSnapshotResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *StartSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// StartSnapshotResResponseMetadataError - 仅在请求失败时返回。
type StartSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopRecordRes struct {

	// REQUIRED
	ResponseMetadata StopRecordResResponseMetadata `json:"ResponseMetadata"`

	// 仅在请求成功时返回"ok",失败时为空
	Result *string `json:"Result,omitempty"`
}

type StopRecordResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *StopRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StopRecordResResponseMetadataError - 仅在请求失败时返回。
type StopRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopSegmentBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 你需要关闭的音频切片任务的 ID。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopSegmentRes struct {

	// REQUIRED
	ResponseMetadata StopSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
}

type StopSegmentResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *StopSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// StopSegmentResResponseMetadataError - 仅在请求失败时返回。
type StopSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopSnapshotBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 你需要关闭的抽帧截图任务的 ID。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopSnapshotRes struct {

	// REQUIRED
	ResponseMetadata StopSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type StopSnapshotResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *StopSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// StopSnapshotResResponseMetadataError - 仅在请求失败时返回。
type StopSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string                 `json:"BusinessId,omitempty"`
	Layout     *UpdateRecordBodyLayout `json:"Layout,omitempty"`

	// 流的类型
	TargetStreams *UpdateRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type UpdateRecordBodyLayout struct {
	CustomLayout *UpdateRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为0，值的范围为{0, 1, 2, 3}。
	// * 0为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2为自定义布局模式。
	// * 3为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                 `json:"LayoutMode,omitempty"`
	MainVideoStream *UpdateRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayout struct {
	Canvas *UpdateRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即
	// Regions.Region.StreamIndex 要与
	// TargetStreams.StreamList.Stream.Index 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*UpdateRecordBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：- 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// 值不合法时，自动调整为默认值。
	// 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdateRecordBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

// UpdateRecordBodyTargetStreams - 流的类型
type UpdateRecordBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个StreamList中，Stream.Index不能重复。
	StreamList []*UpdateRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type UpdateRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用Index对流进行标志。后续在Layout.regions.StreamIndex中，你需要使用Index指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdateRecordRes struct {

	// REQUIRED
	ResponseMetadata UpdateRecordResResponseMetadata `json:"ResponseMetadata"`

	// 仅在请求成功时返回"ok",失败时为空
	Result *string `json:"Result,omitempty"`
}

type UpdateRecordResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *UpdateRecordResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateRecordResResponseMetadataError - 仅在请求失败时返回。
type UpdateRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateSegmentBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 切片任务 ID。你必须对每个切片任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 音频切片的时长。单位为秒。范围为 [1, 600]，默认值为 20秒。值不合法时，自动调整为默认值。 更新该字段后，计时器会重新启动，当前切片立即停止，生成一个新切片。
	Duration *int32 `json:"Duration,omitempty"`

	// 是否在开启音视切片时，立刻开始切片。选择 True 时，开启切片；选择 False 时，不切片。默认值 true。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。 切片文件名由 Identifier + UserId + 时间戳 + 序号组成。默认情况下 Identifier 为 随机生成的 UUID。在自定义文件名时，Identifier 的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。
	// 在自定义文件名时，你需确保 identifier 命名全局唯一，否则在 TOS 平台会因文件名重复被覆盖。
	Identifier *string `json:"Identifier,omitempty"`
}

type UpdateSegmentRes struct {

	// REQUIRED
	ResponseMetadata UpdateSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                          `json:"Result,omitempty"`
}

type UpdateSegmentResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *UpdateSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateSegmentResResponseMetadataError - 仅在请求失败时返回。
type UpdateSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateSnapshotBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 截图任务 ID。你必须对每个截图任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID  *string                        `json:"BusinessId,omitempty"`
	ImageConfig *UpdateSnapshotBodyImageConfig `json:"ImageConfig,omitempty"`
}

type UpdateSnapshotBodyImageConfig struct {

	// 图片的格式。值可取0或1，默认为0。选择0时，图片格式为 JEPG；选择1时，图片格式为 PNG。默认值为0。值不合法时，自动调整为默认值。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，单位为像素，取值范围为[0, 1920]，默认值为0，此时，和视频流的实际高度相同。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，单位为秒，取值范围为[1, 600]，默认值为2。值不合法时，自动调整为默认值。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，单位为像素，取值范围为[0, 1920]。默认值为0，此时，和视频流的实际宽度相同。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type UpdateSnapshotRes struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                           `json:"Result,omitempty"`
}

type UpdateSnapshotResResponseMetadata struct {

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

	// 仅在请求失败时返回。
	Error *UpdateSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateSnapshotResResponseMetadataError - 仅在请求失败时返回。
type UpdateSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（仅后处理模块返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}
type GetRealtimePostProcessing struct{}
type GetRealtimePostProcessingBody struct{}
type GetRecordTask struct{}
type GetRecordTaskBody struct{}
type GetSegmentTask struct{}
type GetSegmentTaskBody struct{}
type GetSnapshotTask struct{}
type GetSnapshotTaskBody struct{}
type StartRecord struct{}
type StartRecordQuery struct{}
type StartSegment struct{}
type StartSegmentQuery struct{}
type StartSnapshot struct{}
type StartSnapshotQuery struct{}
type StopRecord struct{}
type StopRecordQuery struct{}
type StopSegment struct{}
type StopSegmentQuery struct{}
type StopSnapshot struct{}
type StopSnapshotQuery struct{}
type UpdateRecord struct{}
type UpdateRecordQuery struct{}
type UpdateSegment struct{}
type UpdateSegmentQuery struct{}
type UpdateSnapshot struct{}
type UpdateSnapshotQuery struct{}
type GetRealtimePostProcessingReq struct {
	*GetRealtimePostProcessingQuery
	*GetRealtimePostProcessingBody
}
type GetRecordTaskReq struct {
	*GetRecordTaskQuery
	*GetRecordTaskBody
}
type GetSegmentTaskReq struct {
	*GetSegmentTaskQuery
	*GetSegmentTaskBody
}
type GetSnapshotTaskReq struct {
	*GetSnapshotTaskQuery
	*GetSnapshotTaskBody
}
type StartRecordReq struct {
	*StartRecordQuery
	*StartRecordBody
}
type StartSegmentReq struct {
	*StartSegmentQuery
	*StartSegmentBody
}
type StartSnapshotReq struct {
	*StartSnapshotQuery
	*StartSnapshotBody
}
type StopRecordReq struct {
	*StopRecordQuery
	*StopRecordBody
}
type StopSegmentReq struct {
	*StopSegmentQuery
	*StopSegmentBody
}
type StopSnapshotReq struct {
	*StopSnapshotQuery
	*StopSnapshotBody
}
type UpdateRecordReq struct {
	*UpdateRecordQuery
	*UpdateRecordBody
}
type UpdateSegmentReq struct {
	*UpdateSegmentQuery
	*UpdateSegmentBody
}
type UpdateSnapshotReq struct {
	*UpdateSnapshotQuery
	*UpdateSnapshotBody
}

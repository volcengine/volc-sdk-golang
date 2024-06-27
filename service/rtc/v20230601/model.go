package rtc_v20230601

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

type GetPushMixedStreamToCDNTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 要查询的转推直播任务 ID。通过服务端发起时，该值为调用 OpenAPI 时传入的 TaskId。通过客户端 SDK 发起时，TaskId 是按照 userId@@taskId 格式拼接而成的字符串；当传入的 taskId 为空时，这里的
	// TaskId 为 userId。
	// TaskId和UserId均为非必填参数，但是你需要至少填一个参数以保证可以正常发起请求。
	TaskID *string `json:"TaskId,omitempty" query:"TaskId"`

	// 客户端发起转推任务的用户 ID。
	// 你在客户端发起多个任务，当使用该接口进行查询时：
	// * 查询第一个任务时，UserId 可以传入发起转推任务的用户 UserId，TaskId 传入空字符串；或直接将该用户的 UserId 传入 TaskId。
	// * 查询第二个以上任务时，UserId 和 TaskId 为发起转推任务的用户 UserId 和 TaskId。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type GetPushMixedStreamToCDNTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPushMixedStreamToCDNTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetPushMixedStreamToCDNTaskResResult          `json:"Result,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResponseMetadata struct {

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
	Error *GetPushMixedStreamToCDNTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResponseMetadataError - 仅在请求失败时返回。
type GetPushMixedStreamToCDNTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResult struct {

	// 合流转推任务信息
	PushMixedStreamToCDNTask *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask `json:"PushMixedStreamToCDNTask,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask - 合流转推任务信息
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask struct {
	Control *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControl `json:"Control,omitempty"`
	Encode  *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskEncode  `json:"Encode,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 转推任务包含的音视频流
	ExcludeStreams *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreams `json:"ExcludeStreams,omitempty"`
	Layout         *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayout         `json:"Layout,omitempty"`

	// 推流状态
	// * 0：运行中，未获取到任务状态，建议稍后重新查询
	// * 1：未开始推流
	// * 2：首次连接 CDN 服务
	// * 3：正在重连 CDN 服务
	// * 4：连接 CDN 服务成功，正在尝试推流。
	// * 5：连接 CDN 服务成功，推流成功
	// * 6：已停止推流。
	// 仅当Status=2 时，PushStreamState 有实际意义；当Status=3 时，PushStreamState=6; status 为其他值时，PushStreamState 均为0。
	PushStreamState *int32 `json:"PushStreamState,omitempty"`

	// 推流 CDN 地址。仅支持 RTMP 协议。
	PushURL *string `json:"PushURL,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0：未知异常状态
	// * 1：未开始
	// * 2： 运行中
	// * 3： 已结束
	// * 4： 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：客户端用户主动退房/调用停止转推接口
	StopReason *string `json:"StopReason,omitempty"`

	// 转推任务包含的音视频流
	TargetStreams *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreams `json:"TargetStreams,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControl struct {

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

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                                                            `json:"PushStreamMode,omitempty"`
	SEIConfig      *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSEIConfig     `json:"SEIConfig,omitempty"`
	SpatialConfig  *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfig `json:"SpatialConfig,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSEIConfig struct {

	// SEI 中是否包含用户说话音量值。
	// * false：否。
	// * true：是。
	// 默认值为 false。值不合法时自动调整为默认值。
	AddVolumeValue *bool `json:"AddVolumeValue,omitempty"`

	// 是否透传 RTC 流里的 SEI 信息。
	// * true：是。
	// * false：否。
	// 默认值为 true。
	PassthroughRTCSEIMessage *bool `json:"PassthroughRTCSEIMessage,omitempty"`

	// 开启音量指示模式后，非关键帧携带 SEI 包含的信息类型。
	// * 0：全量信息。
	// * 1：只有音量信息和时间戳。
	// 关于 SEI 结构，参看服务端合流转推 SEI 结构 [1163740#sei]
	SEIContentMode *int32 `json:"SEIContentMode,omitempty"`

	// 设置 SEI 的 Payload Type，对 服务端合流转推 SEI 和 RTC 透传SEI 均生效。取值为 5 或 100，默认为 100。
	SEIPayLoadType *int32 `json:"SEIPayLoadType,omitempty"`

	// 该字段为长度为 32 位的 16 进制字符，每个字符的范围为 a-f，A-F，0-9，不可包含 -。如果校验失败，会返回错误码 InvalidParameter。
	// 仅当 SEIPayLoadType=5时，该字段需要填写，SEIPayLoadType=100时，该字段内容会被忽略。
	// 此参数仅对 RTC 透传SEI 生效。当用户设置 SEIPayloadType = 5 时，服务端合流转推SEI 的 SEIPayloadUUID 为固定值：566f6c63656e67696e65525443534549，对应16位字符串
	// VolcengineRTCSEI。
	SEIPayloadUUID *string `json:"SEIPayloadUUID,omitempty"`

	// 设置SEI数据结构 [1163740#sei]中 app_data 参数的值，最大长度为 4096。此参数支持动态更新。
	UserConfigExtraInfo *string `json:"UserConfigExtraInfo,omitempty"`

	// 用户说话音量的回调间隔，单位为秒，取值范围为[0.5,∞]，默认值为 2。仅当用户说话音量发生变化时此回调才会被触发。
	VolumeIndicationInterval *float32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。
	// * true：是。此时非关键帧中也可能携带 SEI 信息。
	// * false：否。此时只会在关键帧中携带 SEI 信息。
	// 默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。 关于音量指示模式的用法，参看 音量指示模式 [1163740#volume]。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfig struct {
	AudienceSpatialOrientation *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。 false：关闭。true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfigAudienceSpatialOrientation struct {

	// 前方朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 头顶朝向
	Up []*float32 `json:"Up,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。
	// 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
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

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreams - 转推任务包含的音视频流
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayout struct {
	CustomLayout *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为 0，值的范围为 {0, 1, 2, 3}。
	// * 0 为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1 为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2 为自定义布局模式。
	// * 3 为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                                                             `json:"LayoutMode,omitempty"`
	MainVideoStream *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayout struct {
	Canvas *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。- 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage
	// 不造成任何影响。- 可以传入的图片的格式包括：JPG, JPEG, PNG。- 当图片和画布尺寸不一致时，图片根据
	// RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形 值不合法时，自动调整为默认值。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreams - 转推任务包含的音视频流
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushSingleStreamToCDNTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 要查询的转推直播任务 ID。通过服务端发起时，该值为调用 OpenAPI 时传入的 TaskId。通过客户端 SDK 发起时，TaskId 是按照 userId@@taskId 格式拼接而成的字符串；当传入的 taskId 为空时，这里的
	// TaskId 为 userId。
	// TaskId和UserId均为非必填参数，但是你需要至少填一个参数以保证可以正常发起请求。
	TaskID *string `json:"TaskId,omitempty" query:"TaskId"`

	// 客户端发起转推任务的用户 ID。
	// 你在客户端发起多个任务，当使用该接口进行查询时：
	// * 查询第一个任务时，UserId 可以传入发起转推任务的用户 UserId，TaskId 传入空字符串；或直接将该用户的 UserId 传入 TaskId。
	// * 查询第二个以上任务时，UserId 和 TaskId 为发起转推任务的用户 UserId 和 TaskId。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type GetPushSingleStreamToCDNTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPushSingleStreamToCDNTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetPushSingleStreamToCDNTaskResResult          `json:"Result,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResponseMetadata struct {

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
	Error *GetPushSingleStreamToCDNTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetPushSingleStreamToCDNTaskResResponseMetadataError - 仅在请求失败时返回。
type GetPushSingleStreamToCDNTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResult struct {

	// 单流转推任务信息
	PushSingleStreamToCDNTask *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask `json:"PushSingleStreamToCDNTask,omitempty"`
}

// GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask - 单流转推任务信息
type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask struct {
	Control *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskControl `json:"Control,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 推流状态
	// * 0：未获取到任务状态。建议稍后重新查询
	// * 1：未开始推流
	// * 2：首次连接 CDN 服务
	// * 3：正在重连 CDN 服务
	// * 4：连接 CDN 服务成功，正在尝试推流。
	// * 5：连接 CDN 服务成功，推流成功
	// * 6：已停止推流。
	// 仅当Status=2 时，PushStreamState 有实际意义；当Status=3 时，PushStreamState=6; status 为其他值时，PushStreamState 均为0。
	PushStreamState *int32 `json:"PushStreamState,omitempty"`

	// 推流地址。目前仅支持 rtmp 协议
	PushURL *string `json:"PushURL,omitempty"`

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
	// * UserDisconnect：客户端用户主动退房/调用停止转推接口
	StopReason *string                                                               `json:"StopReason,omitempty"`
	Stream     *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskStream `json:"Stream,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskControl struct {

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为 0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetRecordTaskResResult struct {

	// -
	RecordTask *GetRecordTaskResResultRecordTask `json:"RecordTask,omitempty"`
}

// GetRecordTaskResResultRecordTask - -
type GetRecordTaskResResultRecordTask struct {

	// 任务结束的时间，为 Unix 时间戳，单位毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 录制生成的文件列表。
	RecordFileList []*GetRecordTaskResResultRecordTaskRecordFileListItem `json:"RecordFileList,omitempty"`

	// 任务开始的时间，为 Unix 时间戳，单位毫秒。
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态:0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因：
	// * 空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动通过 API 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：超过了最大空闲时间
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

	// 录制文件中包含流的列表。
	StreamList []*GetRecordTaskResResultRecordTaskRecordFileListPropertiesItemsItem `json:"StreamList,omitempty"`

	// 文件在火山引擎视频点播 VOD [https://www.volcengine.com/product/vod] 平台的唯一标识。你可以根据 vid 在点播平台上找到对应的文件。仅在你选择配置存储到 Vod 平台时，此参数有效。
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

type StartPushMixedStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 推流 CDN 地址。仅支持 RTMP 协议。
	PushURL string `json:"PushURL"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 合流转推任务 ID。你必须对每个合流转推任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string                               `json:"BusinessId,omitempty"`
	Control    *StartPushMixedStreamToCDNBodyControl `json:"Control,omitempty"`
	Encode     *StartPushMixedStreamToCDNBodyEncode  `json:"Encode,omitempty"`

	// 转推任务包含的音视频流
	ExcludeStreams *StartPushMixedStreamToCDNBodyExcludeStreams `json:"ExcludeStreams,omitempty"`
	Layout         *StartPushMixedStreamToCDNBodyLayout         `json:"Layout,omitempty"`

	// 转推任务包含的音视频流
	TargetStreams *StartPushMixedStreamToCDNBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type StartPushMixedStreamToCDNBodyControl struct {

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

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                             `json:"PushStreamMode,omitempty"`
	SEIConfig      *StartPushMixedStreamToCDNBodyControlSEIConfig     `json:"SEIConfig,omitempty"`
	SpatialConfig  *StartPushMixedStreamToCDNBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`
}

type StartPushMixedStreamToCDNBodyControlSEIConfig struct {

	// SEI 中是否包含用户说话音量值。
	// * false：否。
	// * true：是。
	// 默认值为 false。值不合法时自动调整为默认值。
	AddVolumeValue *bool `json:"AddVolumeValue,omitempty"`

	// 是否透传 RTC 流里的 SEI 信息。
	// * true：是。
	// * false：否。
	// 默认值为 true。
	PassthroughRTCSEIMessage *bool `json:"PassthroughRTCSEIMessage,omitempty"`

	// 开启音量指示模式后，非关键帧携带 SEI 包含的信息类型。
	// * 0：全量信息。
	// * 1：只有音量信息和时间戳。
	// 关于 SEI 结构，参看服务端合流转推 SEI 结构 [1163740#sei]
	SEIContentMode *int32 `json:"SEIContentMode,omitempty"`

	// 设置 SEI 的 Payload Type，对 服务端合流转推 SEI 和 RTC 透传SEI 均生效。取值为 5 或 100，默认为 100。
	SEIPayLoadType *int32 `json:"SEIPayLoadType,omitempty"`

	// 该字段为长度为 32 位的 16 进制字符，每个字符的范围为 a-f，A-F，0-9，不可包含 -。如果校验失败，会返回错误码 InvalidParameter。
	// 仅当 SEIPayLoadType=5时，该字段需要填写，SEIPayLoadType=100时，该字段内容会被忽略。
	// 此参数仅对 RTC 透传SEI 生效。当用户设置 SEIPayloadType = 5 时，服务端合流转推SEI 的 SEIPayloadUUID 为固定值：566f6c63656e67696e65525443534549，对应16位字符串
	// VolcengineRTCSEI。
	SEIPayloadUUID *string `json:"SEIPayloadUUID,omitempty"`

	// 设置SEI数据结构 [1163740#sei]中 app_data 参数的值，最大长度为 4096。此参数支持动态更新。
	UserConfigExtraInfo *string `json:"UserConfigExtraInfo,omitempty"`

	// 用户说话音量的回调间隔，单位为秒，取值范围为[0.5,∞]，默认值为 2。仅当用户说话音量发生变化时此回调才会被触发。
	VolumeIndicationInterval *float32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。
	// * true：是。此时非关键帧中也可能携带 SEI 信息。
	// * false：否。此时只会在关键帧中携带 SEI 信息。
	// 默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。 关于音量指示模式的用法，参看 音量指示模式 [1163740#volume]。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type StartPushMixedStreamToCDNBodyControlSpatialConfig struct {
	AudienceSpatialOrientation *StartPushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。 false：关闭。true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

type StartPushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 前方朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 头顶朝向
	Up []*float32 `json:"Up,omitempty"`
}

type StartPushMixedStreamToCDNBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。
	// 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
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

// StartPushMixedStreamToCDNBodyExcludeStreams - 转推任务包含的音视频流
type StartPushMixedStreamToCDNBodyExcludeStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartPushMixedStreamToCDNBodyExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartPushMixedStreamToCDNBodyExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayout struct {
	CustomLayout *StartPushMixedStreamToCDNBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为 0，值的范围为 {0, 1, 2, 3}。
	// * 0 为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1 为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2 为自定义布局模式。
	// * 3 为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                              `json:"LayoutMode,omitempty"`
	MainVideoStream *StartPushMixedStreamToCDNBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayout struct {
	Canvas *StartPushMixedStreamToCDNBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。- 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage
	// 不造成任何影响。- 可以传入的图片的格式包括：JPG, JPEG, PNG。- 当图片和画布尺寸不一致时，图片根据
	// RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形 值不合法时，自动调整为默认值。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

// StartPushMixedStreamToCDNBodyTargetStreams - 转推任务包含的音视频流
type StartPushMixedStreamToCDNBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartPushMixedStreamToCDNBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartPushMixedStreamToCDNBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StartPushMixedStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                      `json:"Result,omitempty"`
}

type StartPushMixedStreamToCDNResResponseMetadata struct {

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
	Error *StartPushMixedStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushMixedStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StartPushMixedStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartPushSingleStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED
	Control StartPushSingleStreamToCDNBodyControl `json:"Control"`

	// REQUIRED; 推流地址。目前仅支持 rtmp 协议
	PushURL string `json:"PushURL"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED
	Stream StartPushSingleStreamToCDNBodyStream `json:"Stream"`

	// REQUIRED; 转推任务 ID。你必须对每个转推任务设定 TaskId，且在后续进行任务结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StartPushSingleStreamToCDNBodyControl struct {

	// 任务的空闲超时时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10, 86400], 默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为 0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。值不合法时，自动调整为默认值。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type StartPushSingleStreamToCDNBodyStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushSingleStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StartPushSingleStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                       `json:"Result,omitempty"`
}

type StartPushSingleStreamToCDNResResponseMetadata struct {

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
	Error *StartPushSingleStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushSingleStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StartPushSingleStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
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
	BusinessID *string                 `json:"BusinessId,omitempty"`
	Control    *StartRecordBodyControl `json:"Control,omitempty"`
	Encode     *StartRecordBodyEncode  `json:"Encode,omitempty"`

	// 转推任务包含的音视频流
	ExcludeStreams   *StartRecordBodyExcludeStreams   `json:"ExcludeStreams,omitempty"`
	FileFormatConfig *StartRecordBodyFileFormatConfig `json:"FileFormatConfig,omitempty"`

	// 录制文件的命名设置。
	// 录制文件的名称由 [StorageConfig.Type] 和 [FileNameConfig] 共同决定。
	// * 当StorageConfig.Type配置为 0 ，即存储在 TOS 平台时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern +随机数组成
	// * 当StorageConfig.Type配置为 1 ，即存储在 火山引擎视频点播平台 [https://www.volcengine.com/product/vod] 平台时，录制文件名称由FilenameConfig.Pattern
	// + 随机数组成
	// * 当StorageConfig.Type配置为 2 ，即存储在指定第三方S3 对象存储平台 [155125#storage]时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern
	// +随机数组成。 例如：当 StorageConfig.Type 配置为 0 ，FilenameConfig.Prefix配置为
	// directory1/directory2/，FilenameConfig.Pattern 配置为 {TaskId}_{RoomId}_{StartTime}_{Duration}， 若此时该文件的 TaskId = mytask123456789,
	// RoomId = myroom99991234，StartTime =1645769481126，Duration = 30s
	// 且生成的随机八位字符为 c4694fa1，则生成录制文件的文件名为：directory1/directory2/mytask123456789_myroom99991234_1645769481126_000030_c4694fa1.mp4。
	// > 文件名在 视频点播平台 [https://www.volcengine.com/product/vod] 上可以重复，但在 TOS [https://www.volcengine.com/product/tos] 或第三方存储平台 [155125#storage]上默认不可以重复。
	// 如果你设置了 Pattern，需自行保证最终文件名的唯一性，否则在 TOS 或第三方存储平台上同名文件将被覆盖;
	// 你也可以给对应 bucket 开启版本控制，允许文件名重复，防止被覆盖的情况发生。
	FileNameConfig *StartRecordBodyFileNameConfig `json:"FileNameConfig,omitempty"`
	Layout         *StartRecordBodyLayout         `json:"Layout,omitempty"`

	// 使用此参数设定录制模式：单流/合流录制。0 表示合流录制，1 表示单流录制。默认值为 0。
	RecordMode *int32 `json:"RecordMode,omitempty"`

	// 转推任务包含的音视频流
	TargetStreams *StartRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

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

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                               `json:"PushStreamMode,omitempty"`
	SEIConfig      *StartRecordBodyControlSEIConfig     `json:"SEIConfig,omitempty"`
	SpatialConfig  *StartRecordBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`
}

type StartRecordBodyControlSEIConfig struct {

	// SEI 中是否包含用户说话音量值。
	// * false：否。
	// * true：是。
	// 默认值为 false。值不合法时自动调整为默认值。
	AddVolumeValue *bool `json:"AddVolumeValue,omitempty"`

	// 是否透传 RTC 流里的 SEI 信息。
	// * true：是。
	// * false：否。
	// 默认值为 true。
	PassthroughRTCSEIMessage *bool `json:"PassthroughRTCSEIMessage,omitempty"`

	// 开启音量指示模式后，非关键帧携带 SEI 包含的信息类型。
	// * 0：全量信息。
	// * 1：只有音量信息和时间戳。
	// 关于 SEI 结构，参看服务端合流转推 SEI 结构 [1163740#sei]
	SEIContentMode *int32 `json:"SEIContentMode,omitempty"`

	// 设置 SEI 的 Payload Type，对 服务端合流转推 SEI 和 RTC 透传SEI 均生效。取值为 5 或 100，默认为 100。
	SEIPayLoadType *int32 `json:"SEIPayLoadType,omitempty"`

	// 该字段为长度为 32 位的 16 进制字符，每个字符的范围为 a-f，A-F，0-9，不可包含 -。如果校验失败，会返回错误码 InvalidParameter。
	// 仅当 SEIPayLoadType=5时，该字段需要填写，SEIPayLoadType=100时，该字段内容会被忽略。
	// 此参数仅对 RTC 透传SEI 生效。当用户设置 SEIPayloadType = 5 时，服务端合流转推SEI 的 SEIPayloadUUID 为固定值：566f6c63656e67696e65525443534549，对应16位字符串
	// VolcengineRTCSEI。
	SEIPayloadUUID *string `json:"SEIPayloadUUID,omitempty"`

	// 设置SEI数据结构 [1163740#sei]中 app_data 参数的值，最大长度为 4096。此参数支持动态更新。
	UserConfigExtraInfo *string `json:"UserConfigExtraInfo,omitempty"`

	// 用户说话音量的回调间隔，单位为秒，取值范围为[0.5,∞]，默认值为 2。仅当用户说话音量发生变化时此回调才会被触发。
	VolumeIndicationInterval *float32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。
	// * true：是。此时非关键帧中也可能携带 SEI 信息。
	// * false：否。此时只会在关键帧中携带 SEI 信息。
	// 默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。 关于音量指示模式的用法，参看 音量指示模式 [1163740#volume]。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type StartRecordBodyControlSpatialConfig struct {
	AudienceSpatialOrientation *StartRecordBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。 false：关闭。true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

type StartRecordBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 前方朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 头顶朝向
	Up []*float32 `json:"Up,omitempty"`
}

type StartRecordBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。
	// 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
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

// StartRecordBodyExcludeStreams - 转推任务包含的音视频流
type StartRecordBodyExcludeStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartRecordBodyExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordBodyFileFormatConfig struct {

	// 存储文件格式。可取值为：MP4、HLS、FLV、MP3、 AAC、M4A。可同时选择多个存储文件格式，最终生成所选存储格式对应的多个文件。
	// MP3、AAC和M4A仅在编码纯音频时有效。
	FileFormat []*string `json:"FileFormat,omitempty"`
}

// StartRecordBodyFileNameConfig - 录制文件的命名设置。
// 录制文件的名称由 [StorageConfig.Type] 和 [FileNameConfig] 共同决定。
// * 当StorageConfig.Type配置为 0 ，即存储在 TOS 平台时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern +随机数组成
// * 当StorageConfig.Type配置为 1 ，即存储在 火山引擎视频点播平台 [https://www.volcengine.com/product/vod] 平台时，录制文件名称由FilenameConfig.Pattern
// + 随机数组成
// * 当StorageConfig.Type配置为 2 ，即存储在指定第三方S3 对象存储平台 [155125#storage]时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern
// +随机数组成。 例如：当 StorageConfig.Type 配置为 0 ，FilenameConfig.Prefix配置为
// directory1/directory2/，FilenameConfig.Pattern 配置为 {TaskId}_{RoomId}_{StartTime}_{Duration}， 若此时该文件的 TaskId = mytask123456789,
// RoomId = myroom99991234，StartTime =1645769481126，Duration = 30s
// 且生成的随机八位字符为 c4694fa1，则生成录制文件的文件名为：directory1/directory2/mytask123456789_myroom99991234_1645769481126_000030_c4694fa1.mp4。
// > 文件名在 视频点播平台 [https://www.volcengine.com/product/vod] 上可以重复，但在 TOS [https://www.volcengine.com/product/tos] 或第三方存储平台 [155125#storage]上默认不可以重复。
// 如果你设置了 Pattern，需自行保证最终文件名的唯一性，否则在 TOS 或第三方存储平台上同名文件将被覆盖;
// 你也可以给对应 bucket 开启版本控制，允许文件名重复，防止被覆盖的情况发生。
type StartRecordBodyFileNameConfig struct {

	// 自定义录制文件名模式，具体参看自定义录制文件名 [106873]。
	Pattern *string `json:"Pattern,omitempty"`

	// 制定录制文件名的前缀，对TosConfig和CustomConfig生效。
	// Prefix 为指定录制文件名的前缀，是一个由多个字符串组成的数组，在 TOS 以及支持 S3 的第三方存储平台上，可以实现指定文件夹的功能。如：Prefix=["directory1","directory2"]，将在录制文件名前加上前缀
	// "directory1/directory2/"。 前缀长度（包括斜杠）不得超过 128 个字符。Prefix中不得出现斜杠、下划线、括号等符号字符。
	// 仅支持以下字符：
	// * 26 个小写英文字母 a-z
	// * 26 个大写英文字母 A-Z
	// * 10 个数字 0-9
	Prefix []*string `json:"Prefix,omitempty"`
}

type StartRecordBodyLayout struct {
	CustomLayout *StartRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为 0，值的范围为 {0, 1, 2, 3}。
	// * 0 为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1 为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2 为自定义布局模式。
	// * 3 为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                `json:"LayoutMode,omitempty"`
	MainVideoStream *StartRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type StartRecordBodyLayoutCustomLayout struct {
	Canvas *StartRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
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

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。- 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage
	// 不造成任何影响。- 可以传入的图片的格式包括：JPG, JPEG, PNG。- 当图片和画布尺寸不一致时，图片根据
	// RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形 值不合法时，自动调整为默认值。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
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

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
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

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。* 火山引擎平台账号 ID 查看路径：控制台->账号管理->账号信息 ![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_2645e01f7a4e27c77038348297d2849d.png
	// =60%x)
	// * 此账号 ID 为火山引擎主账号 ID。
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。
	// * 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。 你也可以通过 OpenAPI 的方式获取服务
	// ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]默认值为0。
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

// StartRecordBodyTargetStreams - 转推任务包含的音视频流
type StartRecordBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordRes struct {

	// REQUIRED
	ResponseMetadata StartRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopPushStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 转推任务 ID。你必须对每个转推任务设定 TaskId，且在进行任务结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type StopPushStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StopPushStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                `json:"Result,omitempty"`
}

type StopPushStreamToCDNResResponseMetadata struct {

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
	Error *StopPushStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StopPushStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StopPushStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
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
	Result           *string                       `json:"Result,omitempty"`
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdatePushMixedStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 合流转推任务 ID。你必须对每个合流转推任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string                                `json:"BusinessId,omitempty"`
	Control    *UpdatePushMixedStreamToCDNBodyControl `json:"Control,omitempty"`
	Encode     *UpdatePushMixedStreamToCDNBodyEncode  `json:"Encode,omitempty"`

	// 是否更新部分参数。
	// * False：否。
	// * True：是。
	// 默认值为 False。 开启部分更新后，必须按照参数层级传入，且数组类参数需要传入该数组中所有参数。
	IsUpdatePartialParam *bool                                 `json:"IsUpdatePartialParam,omitempty"`
	Layout               *UpdatePushMixedStreamToCDNBodyLayout `json:"Layout,omitempty"`

	// 更新请求序列号。填写该参数后，服务端会对请求进行校验，请确保最后一次更新请求的序列号大于前一次请求的序列号。
	// 建议更新部分参数场景下传入此参数，以确保服务端按照最新请求进行更新。
	SequenceNumber *int32 `json:"SequenceNumber,omitempty"`

	// 转推任务包含的音视频流
	TargetStreams *UpdatePushMixedStreamToCDNBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControl struct {

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

	// 转推直播推流模式，用于控制触发推流的时机。取值为0或1。默认值为0。
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。 值不合法时，自动调整为默认值。
	// 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                              `json:"PushStreamMode,omitempty"`
	SEIConfig      *UpdatePushMixedStreamToCDNBodyControlSEIConfig     `json:"SEIConfig,omitempty"`
	SpatialConfig  *UpdatePushMixedStreamToCDNBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControlSEIConfig struct {

	// SEI 中是否包含用户说话音量值。
	// * false：否。
	// * true：是。
	// 默认值为 false。值不合法时自动调整为默认值。
	AddVolumeValue *bool `json:"AddVolumeValue,omitempty"`

	// 是否透传 RTC 流里的 SEI 信息。
	// * true：是。
	// * false：否。
	// 默认值为 true。
	PassthroughRTCSEIMessage *bool `json:"PassthroughRTCSEIMessage,omitempty"`

	// 开启音量指示模式后，非关键帧携带 SEI 包含的信息类型。
	// * 0：全量信息。
	// * 1：只有音量信息和时间戳。
	// 关于 SEI 结构，参看服务端合流转推 SEI 结构 [1163740#sei]
	SEIContentMode *int32 `json:"SEIContentMode,omitempty"`

	// 设置 SEI 的 Payload Type，对 服务端合流转推 SEI 和 RTC 透传SEI 均生效。取值为 5 或 100，默认为 100。
	SEIPayLoadType *int32 `json:"SEIPayLoadType,omitempty"`

	// 该字段为长度为 32 位的 16 进制字符，每个字符的范围为 a-f，A-F，0-9，不可包含 -。如果校验失败，会返回错误码 InvalidParameter。
	// 仅当 SEIPayLoadType=5时，该字段需要填写，SEIPayLoadType=100时，该字段内容会被忽略。
	// 此参数仅对 RTC 透传SEI 生效。当用户设置 SEIPayloadType = 5 时，服务端合流转推SEI 的 SEIPayloadUUID 为固定值：566f6c63656e67696e65525443534549，对应16位字符串
	// VolcengineRTCSEI。
	SEIPayloadUUID *string `json:"SEIPayloadUUID,omitempty"`

	// 设置SEI数据结构 [1163740#sei]中 app_data 参数的值，最大长度为 4096。此参数支持动态更新。
	UserConfigExtraInfo *string `json:"UserConfigExtraInfo,omitempty"`

	// 用户说话音量的回调间隔，单位为秒，取值范围为[0.5,∞]，默认值为 2。仅当用户说话音量发生变化时此回调才会被触发。
	VolumeIndicationInterval *float32 `json:"VolumeIndicationInterval,omitempty"`

	// (仅对转推直播有效）是否开启音量指示模式。
	// * true：是。此时非关键帧中也可能携带 SEI 信息。
	// * false：否。此时只会在关键帧中携带 SEI 信息。
	// 默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。 关于音量指示模式的用法，参看 音量指示模式 [1163740#volume]。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControlSpatialConfig struct {
	AudienceSpatialOrientation *UpdatePushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。 false：关闭。true：开启
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 前方朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 头顶朝向
	Up []*float32 `json:"Up,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64，值不合法时，自动调整为默认值。
	// 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
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

type UpdatePushMixedStreamToCDNBodyLayout struct {
	CustomLayout *UpdatePushMixedStreamToCDNBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为 0，值的范围为 {0, 1, 2, 3}。
	// * 0 为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1 为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2 为自定义布局模式。
	// * 3 为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                               `json:"LayoutMode,omitempty"`
	MainVideoStream *UpdatePushMixedStreamToCDNBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayout struct {
	Canvas *UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。 当传入的选填参数值不合法时，自动调整为默认值。
	Regions []*UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。值不合法时，自动调整为默认值。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 480。值不合法时，自动调整为默认值。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，单位为像素，范围为 [2, 1920]，必须是偶数。默认值为 640。值不合法时，自动调整为默认值。
	Width *int32 `json:"Width,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。- 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage
	// 不造成任何影响。- 可以传入的图片的格式包括：JPG, JPEG, PNG。- 当图片和画布尺寸不一致时，图片根据
	// RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形 值不合法时，自动调整为默认值。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	// 不同渲染模式下，效果如下：![alt](https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_5e4ddbcdbefe2a108f6f9810bfa0b030.png
	// =100%x)
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为[-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]：0 表示该区域图像位于最下层，100 表示该区域图像位于最上层, 默认值为 0。值不合法时，自动调整为默认值。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。值不合法时，自动调整为默认值。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。值不合法时，自动调整为默认值。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

// UpdatePushMixedStreamToCDNBodyTargetStreams - 转推任务包含的音视频流
type UpdatePushMixedStreamToCDNBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*UpdatePushMixedStreamToCDNBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdatePushMixedStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata UpdatePushMixedStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                       `json:"Result,omitempty"`
}

type UpdatePushMixedStreamToCDNResResponseMetadata struct {

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
	Error *UpdatePushMixedStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// UpdatePushMixedStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type UpdatePushMixedStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
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

	// 转推任务包含的音视频流
	TargetStreams *UpdateRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type UpdateRecordBodyLayout struct {
	CustomLayout *UpdateRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。默认值为 0，值的范围为 {0, 1, 2, 3}。
	// * 0 为自适应布局模式。参看自适应布局 [1167930#adapt]。
	// * 1 为垂直布局模式。参看垂直布局 [1167930#vertical]。
	// * 2 为自定义布局模式。
	// * 3 为并排布局模式。参看并排布局 [1167930#horizontal]
	LayoutMode      *int32                                 `json:"LayoutMode,omitempty"`
	MainVideoStream *UpdateRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayout struct {
	Canvas *UpdateRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
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

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]。0.0 表示完全透明，1.0 表示完全不透明，默认值为 1.0。值不合法时，自动调整为默认值。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。- 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage
	// 不造成任何影响。- 可以传入的图片的格式包括：JPG, JPEG, PNG。- 当图片和画布尺寸不一致时，图片根据
	// RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。值不合法时，自动调整为默认值。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 (0,1]。圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。值不合法时，自动调整为默认值。
	// 假如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式，值的范围为 {0, 1, 2，3}, 默认值为 0：
	// * 0 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 表示按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 表示按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 表示按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形 值不合法时，自动调整为默认值。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
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

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

// UpdateRecordBodyTargetStreams - 转推任务包含的音视频流
type UpdateRecordBodyTargetStreams struct {

	// 由Stream组成的列表，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*UpdateRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type UpdateRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型，值可以取0或1，默认值为0。0表示普通音视频流，1表示屏幕流。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdateRecordRes struct {

	// REQUIRED
	ResponseMetadata UpdateRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
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

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}
type GetPushMixedStreamToCDNTask struct{}
type GetPushMixedStreamToCDNTaskBody struct{}
type GetPushSingleStreamToCDNTask struct{}
type GetPushSingleStreamToCDNTaskBody struct{}
type GetRecordTask struct{}
type GetRecordTaskBody struct{}
type StartPushMixedStreamToCDN struct{}
type StartPushMixedStreamToCDNQuery struct{}
type StartPushSingleStreamToCDN struct{}
type StartPushSingleStreamToCDNQuery struct{}
type StartRecord struct{}
type StartRecordQuery struct{}
type StopPushStreamToCDN struct{}
type StopPushStreamToCDNQuery struct{}
type StopRecord struct{}
type StopRecordQuery struct{}
type UpdatePushMixedStreamToCDN struct{}
type UpdatePushMixedStreamToCDNQuery struct{}
type UpdateRecord struct{}
type UpdateRecordQuery struct{}
type GetPushMixedStreamToCDNTaskReq struct {
	*GetPushMixedStreamToCDNTaskQuery
	*GetPushMixedStreamToCDNTaskBody
}
type GetPushSingleStreamToCDNTaskReq struct {
	*GetPushSingleStreamToCDNTaskQuery
	*GetPushSingleStreamToCDNTaskBody
}
type GetRecordTaskReq struct {
	*GetRecordTaskQuery
	*GetRecordTaskBody
}
type StartPushMixedStreamToCDNReq struct {
	*StartPushMixedStreamToCDNQuery
	*StartPushMixedStreamToCDNBody
}
type StartPushSingleStreamToCDNReq struct {
	*StartPushSingleStreamToCDNQuery
	*StartPushSingleStreamToCDNBody
}
type StartRecordReq struct {
	*StartRecordQuery
	*StartRecordBody
}
type StopPushStreamToCDNReq struct {
	*StopPushStreamToCDNQuery
	*StopPushStreamToCDNBody
}
type StopRecordReq struct {
	*StopRecordQuery
	*StopRecordBody
}
type UpdatePushMixedStreamToCDNReq struct {
	*UpdatePushMixedStreamToCDNQuery
	*UpdatePushMixedStreamToCDNBody
}
type UpdateRecordReq struct {
	*UpdateRecordQuery
	*UpdateRecordBody
}

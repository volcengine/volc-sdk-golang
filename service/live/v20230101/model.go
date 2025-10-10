package live_v20230101

type BindCertBody struct {

	// REQUIRED; 需要绑定的 HTTPS 证书的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/1126822]接口获取。
	ChainID string `json:"ChainID"`

	// REQUIRED; 填写需要配置 HTTPS 证书的域名。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要绑定证书的域名。
	Domain string `json:"Domain"`

	// 是否开启 HTTP/2 协议，默认为 false。取值如下：
	// * false: 关闭 HTTP/2 协议。
	// * true: 开启 HTTP/2 协议。
	HTTP2 *bool `json:"HTTP2,omitempty"`

	// 是否启用 HTTPS 协议，默认值为 false，取值及含义如下所示。
	// * false：关闭；
	// * true：启用。
	HTTPS *bool `json:"HTTPS,omitempty"`

	// 最大支持的TLS版本，不填默认不校验，可选值为：TLSv1.0、TLSv1.1、TLSv1.2、TLSv1.3
	MaxTLSVersion *string `json:"MaxTLSVersion,omitempty"`

	// 最小支持的TLS版本，不填默认为TLSv1.2，可选值为：TLSv1.0、TLSv1.1、TLSv1.2、TLSv1.3
	MinTLSVersion *string `json:"MinTLSVersion,omitempty"`
}

type BindCertRes struct {

	// REQUIRED
	ResponseMetadata BindCertResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type BindCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                            `json:"Version"`
	Error   *BindCertResResponseMetadataError `json:"Error,omitempty"`
}

type BindCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type BindEncryptDRMBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 加密类型，支持的取值及含义如下所示。
	// * FairPlay：使用 FairPlay 技术的商业 DRM 加密；
	// * Widevine：使用 Widevine 技术的商业 DRM 加密；
	// * PlayReady：使用 PlayReady 技术的商业 DRM 加密；
	// * ClearKey：HLS 标准加密。
	// :::tip DRM 加密与 HLS 标准加密不可同时配置。 :::
	DRMSystems []string `json:"DRMSystems"`

	// REQUIRED; 是否开启源流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。 :::tip 源流和转码流至少有一个需要开启录制。 :::
	EncryptOriginStream bool `json:"EncryptOriginStream"`

	// REQUIRED; 是否开启转码流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。 :::tip 源流和转码流至少有一个需要开启录制。 :::
	EncryptTranscodeStream bool `json:"EncryptTranscodeStream"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 开启转码流加密时待加密的转码流对应的转码流后缀配置。您可以调用查询转码配置列表 [https://www.volcengine.com/docs/6469/1126853]接口或在视频直播控制台的转码配置 [https://console.volcengine.com/live/main/application/transcode]页面，查看转码配置的转码流后缀。
	EncryptTranscodeSuffix []*string `json:"EncryptTranscodeSuffix,omitempty"`
}

type BindEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata BindEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type BindEncryptDRMResResponseMetadata struct {

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

// Components130L268SchemasCreatehighlighttaskbodyPropertiesSellpointparamPropertiesEcommerceinfoAdditionalproperties - 对应输入视频的商家卖点信息
type Components130L268SchemasCreatehighlighttaskbodyPropertiesSellpointparamPropertiesEcommerceinfoAdditionalproperties struct {
	ProductInfo []*CreateHighLightTaskBodySellPointParamECommerceInfoPropertiesItemsItem `json:"ProductInfo,omitempty"`
}

// Components1404CjzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesTosparam
// - TOS 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type Components1404CjzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesTosparam struct {
	Bucket string `json:"Bucket"`

	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	StorageDir string `json:"StorageDir"`
}

// Components1523StvSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguage
// - 原文字幕展示参数配置。
type Components1523StvSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguage struct {
	Border Components1O8E0AlSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguagePropertiesBorder `json:"Border"`

	Display bool `json:"Display"`

	Font string `json:"Font"`

	FontColor string `json:"FontColor"`

	Language string `json:"Language"`
}

type Components1B7Y2U1SchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutputPropertiesHlsourceclipsItems struct {
	HLClips []ComponentsHqpv97SchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutputPropertiesHlsourceclipsItemsPropertiesHlclipsItems `json:"HLClips"`

	SourceURL string `json:"SourceUrl"`
}

type Components1C398ShSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesTargetlanguageItems struct {
	Border ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePresetTargetLanguageItemBorder `json:"Border"`

	Font string `json:"Font"`

	FontColor string `json:"FontColor"`

	Language string `json:"Language"`
}

type Components1Nf1A8CSchemasListpulltopushtaskv2ResPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems struct {
	SrcAddr string `json:"SrcAddr"`

	EndOffset *float32 `json:"EndOffset,omitempty"`

	StartOffset *float32 `json:"StartOffset,omitempty"`
}

// Components1O8E0AlSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguagePropertiesBorder
// - 原文字幕的字体描边配置。
type Components1O8E0AlSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguagePropertiesBorder struct {
	Color string `json:"Color"`

	Width int32 `json:"Width"`
}

type Components1Pvao98SchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutput struct {
	HLSourceClips []Components1B7Y2U1SchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutputPropertiesHlsourceclipsItems `json:"HLSourceClips"`
}

type Components1Tzc8QlSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesTosparam struct {
	Bucket          *string `json:"Bucket,omitempty"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	StorageDir      *string `json:"StorageDir,omitempty"`
}

type Components1UtnxifSchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutputPropertiesHlsourceclipsItemsPropertiesHlclipsItemsPropertiesSellpointsresItems struct {
	EndTime int64 `json:"EndTime"`

	Message string `json:"Message"`

	SellPoints string `json:"SellPoints"`

	StartTime int64 `json:"StartTime"`

	Status string `json:"Status"`

	EffectType *string `json:"EffectType,omitempty"`
}

type Components1UxazjaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesImagexparam struct {
	ServiceID       string  `json:"ServiceID"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
}

// Components1Via6UrSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesTosparam
// - TOS 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type Components1Via6UrSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesTosparam struct {
	Bucket string `json:"Bucket"`

	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	StorageDir string `json:"StorageDir"`
}

// Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam
// - 录制为 FLV 格式时的录制参数。
type Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam `json:"VODParam,omitempty"`
}

// ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam
// - 录制为 HLS 格式时的录制参数。
type ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *Components1404CjzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsS0Ofr3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesVodparam `json:"VODParam,omitempty"`
}

// ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam
// - TOS 存储相关配置。
type ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam struct {
	Bucket string `json:"Bucket"`

	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	StorageDir string `json:"StorageDir"`
}

// ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig
// - 录制模板详细配置。
type ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig struct {
	FlvParam *Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam `json:"FlvParam,omitempty"`

	HlsParam *ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam `json:"HlsParam,omitempty"`

	Mp4Param *ComponentsKqy98ZSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4Param `json:"Mp4Param,omitempty"`

	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

type ComponentsGg7M1TSchemasListpulltopushtaskresPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems struct {
	SrcAddr string `json:"SrcAddr"`

	EndOffset *float32 `json:"EndOffset,omitempty"`

	StartOffset *float32 `json:"StartOffset,omitempty"`
}

type ComponentsHqpv97SchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutputPropertiesHlsourceclipsItemsPropertiesHlclipsItems struct {
	Confidence float32 `json:"Confidence"`

	Description string `json:"Description"`

	Duration int64 `json:"Duration"`

	HLEnd int64 `json:"HLEnd"`

	HLStart int64 `json:"HLStart"`

	Index int32 `json:"Index"`

	Label int32 `json:"Label"`

	SEnd int64 `json:"SEnd"`

	SStart int64 `json:"SStart"`

	Score float32 `json:"Score"`

	SellPointsRes []Components1UtnxifSchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutputPropertiesHlsourceclipsItemsPropertiesHlclipsItemsPropertiesSellpointsresItems `json:"SellPointsRes"`

	SourceURL string `json:"SourceUrl"`

	ClipURL *string `json:"ClipUrl,omitempty"`
}

// ComponentsJ1MbxoSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesPosition
// - 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
// :::tip
// * 使用预设配置时，字幕位置设置不生效。
// * 不使用预设配置时，字幕位置设置必填。 :::
type ComponentsJ1MbxoSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesPosition struct {
	MarginHorizontal float32 `json:"MarginHorizontal"`

	MarginVertical float32 `json:"MarginVertical"`
}

// ComponentsK46Cw0SchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpegparamPropertiesImagexparam
// - 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
type ComponentsK46Cw0SchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpegparamPropertiesImagexparam struct {
	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	OverwriteObject string `json:"OverwriteObject"`

	ServiceID string `json:"ServiceID"`
}

// ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam
// - VOD 存储相关配置。
type ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam struct {
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	ExactObject *string `json:"ExactObject,omitempty"`

	StorageClass *int32 `json:"StorageClass,omitempty"`

	VodNamespace *string `json:"VodNamespace,omitempty"`

	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// ComponentsKqy98ZSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4Param
// - 录制为 HLS 格式时的录制参数。
type ComponentsKqy98ZSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4Param struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *Components1Via6UrSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsQms0JiSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesVodparam `json:"VODParam,omitempty"`
}

// ComponentsQms0JiSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesVodparam
// - VOD 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type ComponentsQms0JiSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesVodparam struct {
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	ExactObject *string `json:"ExactObject,omitempty"`

	StorageClass *int32 `json:"StorageClass,omitempty"`

	VodNamespace *string `json:"VodNamespace,omitempty"`

	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// ComponentsS0Ofr3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesVodparam
// - VOD 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type ComponentsS0Ofr3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesVodparam struct {
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	ExactObject *string `json:"ExactObject,omitempty"`

	StorageClass *int32 `json:"StorageClass,omitempty"`

	VodNamespace *string `json:"VodNamespace,omitempty"`

	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// ComponentsSlabtaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparam
// - 截图格式为 JPG 时的截图参数。
type ComponentsSlabtaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparam struct {
	Enable      *bool                                                                                                                                                                                      `json:"Enable,omitempty"`
	ImageXParam *Components1UxazjaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesImagexparam `json:"ImageXParam,omitempty"`
	TOSParam    *Components1Tzc8QlSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesTosparam    `json:"TOSParam,omitempty"`
}

type ContinuePullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取状态为停用的任务 ID。
	TaskID string `json:"TaskId"`

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
}

type ContinuePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata ContinuePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type ContinuePullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *ContinuePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type ContinuePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateCarouselTaskBody struct {

	// REQUIRED; 轮播任务名称，运行中的任务，名称不能重复
	Name string `json:"Name"`

	// REQUIRED; 轮播规则，用于指定轮播播放的素材和行为等。
	Rule CreateCarouselTaskBodyRule `json:"Rule"`
}

// CreateCarouselTaskBodyRule - 轮播规则，用于指定轮播播放的素材和行为等。
type CreateCarouselTaskBodyRule struct {

	// REQUIRED; -1代表无限循环，最终停止由StopTime字段控制，或者系统默认的停止时间（3天）
	Loop int32 `json:"Loop"`

	// REQUIRED; 0为普通模式，此模式下系统会根据前后两个点播素材的头信息来判断是否能不断流拼接，如果不满足拼接条件，在进行素材切换时会断流。 2为转码模式，此模式下系统会将所有素材格式化为固定参数，用户可以配置这个音视频参数，如果不配置默认参数跟随第一个素材，在进行素材切换时不会断流
	Mode int32 `json:"Mode"`

	// REQUIRED; 轮播任务的推流参数，包括视频、音频、推流地址及回调信息。
	Output CreateCarouselTaskBodyRuleOutput `json:"Output"`

	// REQUIRED; 轮播素材列表，用于指定在轮播过程中播放的素材资源。
	Source []CreateCarouselTaskBodyRuleSourceItem `json:"Source"`

	// 播放时间，选填，默认会等待第一个视频缓存完毕，如果系统时间大于此值，则开始播放
	PlayTime *int32 `json:"PlayTime,omitempty"`

	// 停止时间，选填，当此字段被设置时，系统会遵循此时间设置关闭任务
	StopTime *int32 `json:"StopTime,omitempty"`
}

// CreateCarouselTaskBodyRuleOutput - 轮播任务的推流参数，包括视频、音频、推流地址及回调信息。
type CreateCarouselTaskBodyRuleOutput struct {

	// REQUIRED; 推流rtmp地址或者rtmps地址，支持多推，最多填写8条地址，最少1条地址
	URL []string `json:"Url"`

	// 转码模式下有效，可选配置推流的音频参数
	Audio *CreateCarouselTaskBodyRuleOutputAudio `json:"Audio,omitempty"`

	// 回调函数。
	Callback *CreateCarouselTaskBodyRuleOutputCallback `json:"Callback,omitempty"`

	// 转码模式下有效，可选配置推流的视频参数
	Video *CreateCarouselTaskBodyRuleOutputVideo `json:"Video,omitempty"`
}

// CreateCarouselTaskBodyRuleOutputAudio - 转码模式下有效，可选配置推流的音频参数
type CreateCarouselTaskBodyRuleOutputAudio struct {

	// 音频码率设置
	BitRate *int32 `json:"BitRate,omitempty"`

	// mono：单声道；stereo：双声道
	ChannelLayout *string `json:"ChannelLayout,omitempty"`

	// 采样率，可选：22000、32000、44100、48000
	SampleRate *int32 `json:"SampleRate,omitempty"`
}

// CreateCarouselTaskBodyRuleOutputCallback - 回调函数。
type CreateCarouselTaskBodyRuleOutputCallback struct {

	// REQUIRED; 回调地址，系统会将部分信息回调出去
	URL string `json:"Url"`
}

// CreateCarouselTaskBodyRuleOutputVideo - 转码模式下有效，可选配置推流的视频参数
type CreateCarouselTaskBodyRuleOutputVideo struct {

	// 视频码率，单位是bit
	BitRate *int32 `json:"BitRate,omitempty"`

	// 视频帧率取值[10,60]
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 取值[1,10]
	GOP *int32 `json:"GOP,omitempty"`

	// 取值范围[10,2160]
	Height *int32 `json:"Height,omitempty"`

	// 取值范围[10,2160]
	Width *int32 `json:"Width,omitempty"`
}

type CreateCarouselTaskBodyRuleSourceItem struct {

	// REQUIRED; 播放列表内不允许重复
	ID string `json:"ID"`

	// REQUIRED; vod：点播文件；m3u8：m3u8文件
	Type string `json:"Type"`

	// REQUIRED; 轮播素材的公网可访问地址。确保提供的地址能够被公网正常访问，以便正确加载轮播素材内容。
	URL string `json:"Url"`

	// 此素材连续播放几次，字段必须大于等于0
	Loop *int32 `json:"Loop,omitempty"`

	// 可以控制当前素材跳过开头进行播放，单位是秒，注意此字段如果小于等于0或者大于视频长度不生效，只在素材type为vod时生效
	Seek *int32 `json:"Seek,omitempty"`
}

type CreateCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateCarouselTaskResResult `json:"Result"`
}

type CreateCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type CreateCarouselTaskResResult struct {

	// REQUIRED; 轮播任务数据对象。
	Data CreateCarouselTaskResResultData `json:"Data"`
}

// CreateCarouselTaskResResultData - 轮播任务数据对象。
type CreateCarouselTaskResResultData struct {

	// REQUIRED; 任务唯一标识
	TaskID string `json:"TaskID"`
}

type CreateCertBody struct {

	// REQUIRED; 证书信息。
	Rsa CreateCertBodyRsa `json:"Rsa"`

	// REQUIRED; 证书用途，当前仅支持设置为 https，表示用于 HTTPS 加密；
	UseWay string `json:"UseWay"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty"`

	// 证书链 ID，用于标识整个证书链，包括叶子证书（服务器证书）、中间证书（中间 CA 证书）以及根证书（根 CA 证书）。 :::tip 使用当前接口更新证书时， ChainID 为必选参数。 :::
	ChainID *string `json:"ChainID,omitempty"`

	// 项目名称，默认值为 default，您可以登录访问控制 [https://console.volcengine.com/iam/resourcemanage/project]获取项目名称。
	ProjectName *string `json:"ProjectName,omitempty"`
}

// CreateCertBodyRsa - 证书信息。
type CreateCertBodyRsa struct {

	// REQUIRED; 证书私钥的内容，你需要在计算机上使用文本编辑器打开证书私钥，并将所有内容复制粘贴作为参数。 :::tip 请确保证书私钥没有密码保护。 :::
	Prikey string `json:"Prikey"`

	// REQUIRED; 证书内容，你需要在计算机上使用文本编辑器打开证书，并将所有内容复制粘贴作为参数。 :::tip
	// * 视频直播支持证书链校验。你只需要上传为你的域名颁发的证书，系统将自动检索完整的证书链。
	// * 如果你选择上传证书链，请务必包含服务器证书、中间证书和根证书，并按正确的顺序排列：首先是服务器证书，其次是中间证书，然后是根证书。错误的顺序将使证书链无效。 :::
	Pubkey string `json:"Pubkey"`
}

type CreateCertRes struct {

	// REQUIRED
	ResponseMetadata CreateCertResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateCertResResult          `json:"Result,omitempty"`
}

type CreateCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *CreateCertResResponseMetadataError `json:"Error,omitempty"`
}

type CreateCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateCertResResult struct {

	// 证书链 ID。
	ChainID *string `json:"ChainID,omitempty"`

	// 使用该证书的域名。
	Domain *string `json:"Domain,omitempty"`

	// 证书用途，取值及含义如下所示。
	// * https：用于 HTTPS 加密；
	// * sign：用于签名加密。
	UseWay *string `json:"UseWay,omitempty"`
}

type CreateCloudMixTaskBody struct {

	// REQUIRED; 混流任务详细配置。
	MixedRules CreateCloudMixTaskBodyMixedRules `json:"MixedRules"`

	// REQUIRED; 混流任务名称，与正在进行中的任务名称不能重复。
	Name string `json:"Name"`
}

// CreateCloudMixTaskBodyMixedRules - 混流任务详细配置。
type CreateCloudMixTaskBodyMixedRules struct {

	// REQUIRED; 混流输出布局配置。
	InputLayout CreateCloudMixTaskBodyMixedRulesInputLayout `json:"InputLayout"`

	// REQUIRED; 混流素材列表，最多支持配置 8 路素材。
	InputSource []CreateCloudMixTaskBodyMixedRulesInputSourceItem `json:"InputSource"`

	// REQUIRED; 混流输出视频质量参数配置。
	Output CreateCloudMixTaskBodyMixedRulesOutput `json:"Output"`
}

// CreateCloudMixTaskBodyMixedRulesInputLayout - 混流输出布局配置。
type CreateCloudMixTaskBodyMixedRulesInputLayout struct {

	// REQUIRED; 混流输出画布配置及素材布局配置。
	Scene CreateCloudMixTaskBodyMixedRulesInputLayoutScene `json:"Scene"`

	// 混流输出视频中 Logo 布局配置。
	// :::tip 支持最多配置 4 个 Logo，展示层级以添加顺序为准。 :::
	Logo []*CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItem `json:"Logo,omitempty"`
}

type CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItem struct {

	// REQUIRED; Logo 图片在混流输出整体画面中的布局配置。
	Layout CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout `json:"Layout"`

	// REQUIRED; Logo 图片访问地址。
	URL string `json:"Url"`
}

// CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout - Logo 图片在混流输出整体画面中的布局配置。
type CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout struct {

	// REQUIRED
	H int32 `json:"H"`

	// REQUIRED
	W int32 `json:"W"`

	// REQUIRED
	X int32 `json:"X"`

	// REQUIRED
	Y int32 `json:"Y"`
}

// CreateCloudMixTaskBodyMixedRulesInputLayoutScene - 混流输出画布配置及素材布局配置。
type CreateCloudMixTaskBodyMixedRulesInputLayoutScene struct {

	// REQUIRED; 混流输出整体画布高度，单位为 px，取值范围为 [10,2160]。
	Height int32 `json:"Height"`

	// REQUIRED; 混流素材在混流输出整体画面中的布局配置。 :::tip 混流素材布局中需包含所有素材的配置，且需与通过 Layer 参数与混流素材一一匹配。 :::
	Layout []CreateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem `json:"Layout"`

	// REQUIRED; 混流输出画布整体宽度，单位为 px，取值范围为 [10,2160]。
	Width int32 `json:"Width"`
}

type CreateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem struct {

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制高度，单位为 px，取值范围为 [10,2160]。
	// :::tip 限制宽度和限制高度指定了素材展示的限制范围，当素材尺寸和限制值不一致时，素材将在限制范围内根据长边进行等比缩放，其余区域透明展示。 :::
	H int32 `json:"H"`

	// REQUIRED; 当配置素材布局时需要通过 Layer 参数与素材进行一一对应。 :::tip 配置 Logo 图片的布局时此参数不生效。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制宽度，单位为 px，取值范围为 [10,2160]。
	W int32 `json:"W"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 X 偏移位置，单位为 px，取值范围为 0 到设置的画面宽度。
	X int32 `json:"X"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 Y 偏移位置，单位为 px，取值范围为 0 到设置的画面高度。
	Y int32 `json:"Y"`
}

type CreateCloudMixTaskBodyMixedRulesInputSourceItem struct {

	// REQUIRED; 混流素材 ID，一个任务中素材 ID 不能重复，此 ID 用于任务状态回调消息中标识素材。
	ID string `json:"ID"`

	// REQUIRED; 混流素材的叠放顺序，1 为最底层，2 层在 1 层之上，以此类推，取值范围为[1,9999]。 :::tip 当前准备使用某个素材作为布局背景时，其叠放顺序应设置为所有素材中的最小值。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 混流素材类型，支持的取值及含义如下所示。
	// * vod：视频点播中的素材，支持 MP4、FLV 格式素材；
	// * live：直播源素材，支持 RTMP、FLV 协议拉流地址；
	// * pic：图片素材，支持 png、jpg 格式图片。
	Type string `json:"Type"`

	// REQUIRED; 混流素材的访问地址。 :::tip 混流素材的访问地址需与混流素材的类型保持对应关系。 :::
	URL string `json:"Url"`
}

// CreateCloudMixTaskBodyMixedRulesOutput - 混流输出视频质量参数配置。
type CreateCloudMixTaskBodyMixedRulesOutput struct {

	// REQUIRED; 混流音频参数设置。
	Audio CreateCloudMixTaskBodyMixedRulesOutputAudio `json:"Audio"`

	// REQUIRED; 混流视频的推流地址，支持最多配置 8 个推流地址。
	URL []string `json:"Url"`

	// REQUIRED; 混流视频参数设置。
	Video CreateCloudMixTaskBodyMixedRulesOutputVideo `json:"Video"`

	// 任务状态回调地址配置。
	Callback *CreateCloudMixTaskBodyMixedRulesOutputCallback `json:"Callback,omitempty"`
}

// CreateCloudMixTaskBodyMixedRulesOutputAudio - 混流音频参数设置。
type CreateCloudMixTaskBodyMixedRulesOutputAudio struct {

	// REQUIRED; 混流输出流的音频码率，单位为 bps，取值范围为 [128000,320000]，常见取值及含义如下所示。
	// * 128000：128 kbps；
	// * 144000：144 kbps；
	// * 256000：256 kbps；
	// * 320000：320 kbps。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出流的音频声道设置，取值及含义如下所示。
	// * mono：单声道；
	// * stereo：立体声。
	ChannelLayout string `json:"ChannelLayout"`

	// REQUIRED; 混流输出流的音频采样率，单位为 HZ，常见取值及含义如下所示。
	// * 32000：32 kHZ；
	// * 44100：44.1 kHZ；
	// * 48000：48 kHZ。
	SampleRate int32 `json:"SampleRate"`
}

// CreateCloudMixTaskBodyMixedRulesOutputCallback - 任务状态回调地址配置。
type CreateCloudMixTaskBodyMixedRulesOutputCallback struct {

	// REQUIRED; 接收云端混流任务状态回调的 HTTP 地址。
	URL string `json:"Url"`
}

// CreateCloudMixTaskBodyMixedRulesOutputVideo - 混流视频参数设置。
type CreateCloudMixTaskBodyMixedRulesOutputVideo struct {

	// REQUIRED; 混流输出视频码率，单位为 bps，取值范围为 [1000000,20000000]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式。
	Codec string `json:"Codec"`

	// REQUIRED; 混流输出视频帧率，单位为 fps，取值范围为 [10,60]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	FrameRate int32 `json:"FrameRate"`

	// REQUIRED; IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,10]。
	GOP int32 `json:"GOP"`
}

type CreateCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateCloudMixTaskResResult `json:"Result,omitempty"`
}

type CreateCloudMixTaskResResponseMetadata struct {

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

// CreateCloudMixTaskResResult - 视请求的接口而定
type CreateCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data CreateCloudMixTaskResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// CreateCloudMixTaskResResultData - 返回的数据。
type CreateCloudMixTaskResResultData struct {

	// REQUIRED; 混流任务 ID。
	TaskID string `json:"TaskID"`
}

type CreateDomainBody struct {

	// REQUIRED; 待添加到视频直播服务进行加速的域名，域名只能由数字（0 - 9）、字母（A - Z、a -z）和连字符（-） 组成。
	Domain string `json:"Domain"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// 国内可传入：
	// * cn 国内
	// * cn-oversea 海外及港澳台
	// * cn-global 全球 byteplus可传入：
	// * cn 中国
	// * oversea 非中国区
	// * global 全球
	Region *string `json:"Region,omitempty"`
}

type CreateDomainRes struct {

	// REQUIRED
	ResponseMetadata CreateDomainResResponseMetadata `json:"ResponseMetadata"`
}

type CreateDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *CreateDomainResResponseMetadataError `json:"Error,omitempty"`
}

type CreateDomainResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type CreateDomainV2Body struct {

	// REQUIRED; 待添加到视频直播服务进行加速的域名列表信息。 :::tip 一个域名空间下最多包含 10 个域名。 :::
	Domains []CreateDomainV2BodyDomainsItem `json:"Domains"`

	// REQUIRED; 域名空间，是一组关联域名的集合，由字母（A - Z、a -z）、数字（0 - 9）和连字符（-） 组成。您可以自定义新的域名空间或调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口获取已有域名使用的域名空间。
	Vhost string `json:"Vhost"`

	// 为域名空间设置所属项目，默认为空表示归属到 default 项目， 新建域名空间时需要为域名空间绑定项目，您可以在访问控制-项目列表 [https://console.volcengine.com/iam/resourcemanage/project]查看已有项目并对项目进行管理。
	// 项目是火山引擎提供的一种资源管理方式，您可以对不同业务或项目使用的云资源进行分组管理，以实现基于项目的账单查看、子账号授权等功能。
	// * 新建域名空间时，需为域名空间设置所属项目。每个域名空间只能属于一个项目，选择已有的域名空间时，项目不可配置；
	// * 使用基于项目的账单查看需提前开通分账服务，请前往账单管理-分账账单 [https://console.volcengine.com/finance/bill/split-bill/]进行服务开通；
	// * 使用基于项目的子账号授权请参见使用 IAM 用户进行项目权限划分 [https://www.volcengine.com/docs/6469/1166573]。
	ProjectName *string `json:"ProjectName,omitempty"`

	// 为域名空间设置资源标签。您可以通过资源标签从不同维度对云资源进行分类和聚合管理，如资源分账等场景。 :::tip 如需使用标签进行资源分账，可以在可以在账单管理-费用标签 [https://console.volcengine.com/finance/bill/tag/]处管理启用标签，将对应标签运用到账单明细等数据中。
	// :::
	Tags []*CreateDomainV2BodyTagsItem `json:"Tags,omitempty"`
}

type CreateDomainV2BodyDomainsItem struct {

	// REQUIRED; 域名名称，域名由字母（A - Z、a -z）、数字（0 - 9）和连字符（-） 组成，长度为 1 到 60 个字符。
	DomainName string `json:"DomainName"`

	// REQUIRED; 国内可传入：
	// * cn 国内
	// * cn-oversea 海外及港澳台
	// * cn-global 全球 byteplus可传入：
	// * cn 中国
	// * oversea 非中国区
	// * global 全球
	Region string `json:"Region"`

	// REQUIRED; 域名类型，取值及含义如下所示。
	// * push：推流域名；
	// * pull-flv：拉流域名。
	Type string `json:"Type"`

	// HTTPS 证书链 ID，默认为空表示不为域名绑定 HTTPS 证书。您可以调用 ListCertV2 [https://www.volcengine.com/docs/6469/1126823] 接口或在视频直播控制台的证书管理 [https://console.volcengine.com/live/main/certificate]页面，获取待绑定的证书链
	// ID。
	ChainID *string `json:"ChainID,omitempty"`
}

type CreateDomainV2BodyTagsItem struct {

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type CreateDomainV2Res struct {

	// REQUIRED
	ResponseMetadata CreateDomainV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateDomainV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                  `json:"Version"`
	Error   *CreateDomainV2ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateDomainV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateHighLightTaskBody struct {

	// REQUIRED; 算法模型类型。0:足球体育；1：文娱短剧，仅适用于视频源类型为点播视频的场景，即 "Type":1；2:电商，仅适用于视频源类型为点播视频的场景，即 "Type":1。
	Model int32 `json:"Model"`

	// REQUIRED; 任务名称。
	Name string `json:"Name"`

	// REQUIRED; 数据源。直播：拉流地址支持：FLV、HLS、RTMP 、RTM
	Sources []CreateHighLightTaskBodySourcesItem `json:"Sources"`

	// REQUIRED; 任务类型。0:直播
	Type int32 `json:"Type"`

	// 输出高光信息回调参数。
	CallbackParam *CreateHighLightTaskBodyCallbackParam `json:"CallbackParam,omitempty"`

	// 是否需要输出【高光片段】以及相关配置参数设置
	HLClipsParam *CreateHighLightTaskBodyHLClipsParam `json:"HLClipsParam,omitempty"`

	// 是否需要输出【高光混剪】以及相关配置参数设置
	HLMixParam *CreateHighLightTaskBodyHLMixParam `json:"HLMixParam,omitempty"`

	// 直播高光剪辑任务配置参数设置，仅当“数据源类型”选择为【直播】时可设置生效。
	LiveParam *CreateHighLightTaskBodyLiveParam `json:"LiveParam,omitempty"`

	// 电商场景生效，填写后会加卖点效果，不填写则不加
	SellPointParam *CreateHighLightTaskBodySellPointParam `json:"SellPointParam,omitempty"`

	// 算法模型子类型。足球体育：0表示默认类型剪辑；文娱短剧：0表示默认类型剪辑；电商：0表示默认类型剪辑，1表示电商视频提取分镜转场素材。
	SubModel *int32 `json:"SubModel,omitempty"`

	// 音频生成字幕配置
	SubtitleParam *CreateHighLightTaskBodySubtitleParam `json:"SubtitleParam,omitempty"`

	// 输出高光视频上传VOD系统参数
	VodParam *CreateHighLightTaskBodyVodParam `json:"VodParam,omitempty"`
}

// CreateHighLightTaskBodyCallbackParam - 输出高光信息回调参数。
type CreateHighLightTaskBodyCallbackParam struct {

	// 业务自定义设置的参数信息，通过回调直接透传，便于业务自定义标识，默认为空
	CallbackExtra *string `json:"CallbackExtra,omitempty"`

	// 回调类型。0:表示http回调，200:表示不需要回调
	CallbackType *int32 `json:"CallbackType,omitempty"`

	// HTTP回调参数。
	HTTPParams *CreateHighLightTaskBodyCallbackParamHTTPParams `json:"HttpParams,omitempty"`
}

// CreateHighLightTaskBodyCallbackParamHTTPParams - HTTP回调参数。
type CreateHighLightTaskBodyCallbackParamHTTPParams struct {

	// REQUIRED; HTTP回调地址。
	CallbackAdr string `json:"CallbackAdr"`
}

// CreateHighLightTaskBodyHLClipsParam - 是否需要输出【高光片段】以及相关配置参数设置
type CreateHighLightTaskBodyHLClipsParam struct {

	// 针对算法检测出的高光片段前后分别增加的冗余时间，精确到s。默认为0，取值范围为[0,60]。主要适用场景：对算法检出的高光片段内容信任度不高，期望获取更多冗余素材，以便自行剪辑
	BufferDuration *int32 `json:"BufferDuration,omitempty"`

	// 是否禁止高光片段时间戳返回，默认为【否】
	DisableTimestamp *bool `json:"DisableTimestamp,omitempty"`

	// 启用状态。
	Enable *bool `json:"Enable,omitempty"`

	// 输出高光片段的编码格式，取值范围如下：0. H264（默认）, 1. H265。
	EncCodec *int32 `json:"EncCodec,omitempty"`

	// 累计生成高光片段数量上限，仅针对数据源类型为【视频】时生效。默认值为0，表示不限制数量。
	NumLimit *int32 `json:"NumLimit,omitempty"`

	// 是否返回高光片段视频素材，默认为【否】
	OutputHLClips *bool `json:"OutputHLClips,omitempty"`

	// 输出高光片段的视频封装格式，取值范围如下：0. MP4（默认）, 1. HLS（M3U8 + TS）, 2. FLV
	VideoFormat *int32 `json:"VideoFormat,omitempty"`
}

// CreateHighLightTaskBodyHLMixParam - 是否需要输出【高光混剪】以及相关配置参数设置
type CreateHighLightTaskBodyHLMixParam struct {

	// 相对任务开始时，生成并返回【高光混剪】时间点，支持设置多个时间节点。精确到【分钟】，取值范围[5,1440]。仅针对数据源类型为【直播】时生效。默认为【任务结束后返回】
	CreateTimestamps []*int32 `json:"CreateTimestamps,omitempty"`

	// 高光混剪参数：每个高光混剪的时长上限，精确到s。若不填或者填0，服务内部默认值为210
	DurationMax *int32 `json:"DurationMax,omitempty"`

	// 高光混剪参数：每个高光混剪的时长下限，精确到s。若不填或者填0，服务内部默认值为180。
	DurationMin *int32 `json:"DurationMin,omitempty"`

	// 启用状态。
	Enable *bool `json:"Enable,omitempty"`

	// 输出高光片段的编码格式，取值范围如下：0. H264（默认）, 1. H265。
	EncCodec *int32 `json:"EncCodec,omitempty"`

	// 累计生成高光混剪数量上限。默认值为1，取值范围[1,100]。
	NumLimit *int32 `json:"NumLimit,omitempty"`

	// 是否开启卖点贴纸
	SellPointSticker *bool `json:"SellPointSticker,omitempty"`

	// 是否开启声音生成字幕
	Subtitle *bool `json:"Subtitle,omitempty"`

	// 输出高光片段的视频封装格式，取值范围如下：0. MP4（默认）, 1. HLS（M3U8 + TS）, 2. FLV
	VideoFormat *int32 `json:"VideoFormat,omitempty"`
}

// CreateHighLightTaskBodyLiveParam - 直播高光剪辑任务配置参数设置，仅当“数据源类型”选择为【直播】时可设置生效。
type CreateHighLightTaskBodyLiveParam struct {

	// 切片策略，即针对直播流剪辑送检的切片时长。 默认为300s，取值范围[60,10800]
	ClipsDuration *int32 `json:"ClipsDuration,omitempty"`

	// 判断断流时长，即断流超过该时间范围，则默认直播流结束。仅当【TaskEndTime】为空时，该规则作为直播流结束的判断标志生效。默认120s，取值范围[60,900]
	StreamEndTime *int32 `json:"StreamEndTime,omitempty"`

	// 指定高光提取任务的结束时间，RFC3339 格式的时间戳，精度为秒。默认为空，表示高光提取任务执行到直播流结束。
	TaskEndTime *string `json:"TaskEndTime,omitempty"`

	// 指定高光提取任务的开始时间，RFC3339 格式的时间戳，精度为秒。默认为空，表示立即开始。
	TaskStartTime *string `json:"TaskStartTime,omitempty"`
}

// CreateHighLightTaskBodySellPointParam - 电商场景生效，填写后会加卖点效果，不填写则不加
type CreateHighLightTaskBodySellPointParam struct {

	// REQUIRED; 电商场景下使用的卖点效果配置，为 JSON Map 格式。
	// * Key：输入视频的链接或索引
	// * Value：对应视频的商家卖点信息 ProductInfo
	ECommerceInfo CreateHighLightTaskBodySellPointParamECommerceInfo `json:"ECommerceInfo"`

	// 使用的卖点效果模版，默认为default，非必填。当前仅支持default
	EffectType *string `json:"EffectType,omitempty"`
}

// CreateHighLightTaskBodySellPointParamECommerceInfo - 电商场景下使用的卖点效果配置，为 JSON Map 格式。
// * Key：输入视频的链接或索引
// * Value：对应视频的商家卖点信息 ProductInfo
type CreateHighLightTaskBodySellPointParamECommerceInfo struct {

	// REQUIRED
	Field124 string `json:"Field124"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]*Components130L268SchemasCreatehighlighttaskbodyPropertiesSellpointparamPropertiesEcommerceinfoAdditionalproperties
}

type CreateHighLightTaskBodySellPointParamECommerceInfoPropertiesItemsItem struct {

	// REQUIRED; 商家卖点信息
	Desc string `json:"Desc"`
}

type CreateHighLightTaskBodySourcesItem struct {

	// REQUIRED; 视频源路径。直播：拉流地址支持：FLV、HLS、RTMP 、RTM
	Path string `json:"Path"`

	// REQUIRED; 视频源类型。0:表示网络源。
	SourceType int32 `json:"SourceType"`
}

// CreateHighLightTaskBodySubtitleParam - 音频生成字幕配置
type CreateHighLightTaskBodySubtitleParam struct {

	// 字幕描边参数
	Border *CreateHighLightTaskBodySubtitleParamBorder `json:"Border,omitempty"`

	// 字幕字体参数
	Font *CreateHighLightTaskBodySubtitleParamFont `json:"Font,omitempty"`

	// 字幕位置参数
	Position *CreateHighLightTaskBodySubtitleParamPosition `json:"Position,omitempty"`
}

// CreateHighLightTaskBodySubtitleParamBorder - 字幕描边参数
type CreateHighLightTaskBodySubtitleParamBorder struct {

	// 描边颜色：定义方式和【字体颜色】一致，默认为【黑色】
	Color *string `json:"Color,omitempty"`

	// 描边宽度：单位为 px，默认为2
	W *int32 `json:"W,omitempty"`
}

// CreateHighLightTaskBodySubtitleParamFont - 字幕字体参数
type CreateHighLightTaskBodySubtitleParamFont struct {

	// 字幕字体，支持以下字体取值。 songticu: 宋体粗;（默认值）;songtixi: 宋体细; arial: Arial; heitifan: 黑体繁; inter: Inter; kaiti: 楷体; montserrat: Montserrat;
	// notosans: Noto Sans; notosansar: Noto Sans 阿拉伯文; notosansja: Noto Sans
	// 日文; notosansko: Noto Sans 韩文; notosansth: Noto Sans 泰文; opposans: Opposans; roboto: Roboto; simhei: 黑体; siyuanheiti: 思源黑体;
	// siyuansongti: 思源宋体
	Font *string `json:"Font,omitempty"`

	// 字体颜色，支持以下几种方法进行定义。默认为【白色】。支持以 0x 或 # 开头，后面跟着十六进制颜色 RGB 值，再跟着 @+十六进制/百分比来表示的透明度值，来定义字幕的字体颜色。例如，设置 RGB 值为 FF0000，透明度为 5%的颜色时，您可以传入
	// 0xFF0000@0x80、0xFF0000@0.5、#FF0000@0x80 或 #FF0000@0.5。支持使用前端框架 FFmpeg
	// 规定的颜色关键字，来定义字幕的字体颜色。例如，AliceBlue 表示 0xF0F8FF、AntiqueWhite 表示 0xFAEBD7、Black 表示 0x000000 等。
	FontColor *string `json:"FontColor,omitempty"`

	// 字号，默认为55px，推荐取值范围[40,60]
	FontSize *int32 `json:"FontSize,omitempty"`
}

// CreateHighLightTaskBodySubtitleParamPosition - 字幕位置参数
type CreateHighLightTaskBodySubtitleParamPosition struct {

	// 水平边距：字幕距离画面两侧的边距与画面宽度的占比，使用归一化百分表示，取值范围为 [0,0.2]
	MarginLr *float32 `json:"MarginLr,omitempty"`

	// 垂直边距：字幕距离画面底部的边距与画面高度的占比，使用归一化百分表示，取值范围为 [0,0.5]
	MarginTb *float32 `json:"MarginTb,omitempty"`
}

// CreateHighLightTaskBodyVodParam - 输出高光视频上传VOD系统参数
type CreateHighLightTaskBodyVodParam struct {

	// REQUIRED; VOD空间名称。
	Space string `json:"Space"`

	// VOD空间工作流
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type CreateHighLightTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateHighLightTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result CreateHighLightTaskResResult `json:"Result"`
}

type CreateHighLightTaskResResponseMetadata struct {

	// REQUIRED; 请求的唯一标识符。
	RequestID string `json:"RequestID"`

	// 请求的接口名，属于请求的公共参数。
	Action *string `json:"Action,omitempty"`

	// 请求的Region，例如：cn-north-1
	Region *string `json:"Region,omitempty"`

	// 请求的服务，属于请求的公共参数。
	Service *string `json:"Service,omitempty"`

	// 请求的版本号，属于请求的公共参数。
	Version *string `json:"Version,omitempty"`
}

// CreateHighLightTaskResResult - 视请求的接口而定
type CreateHighLightTaskResResult struct {

	// REQUIRED; 参数数据。
	Data CreateHighLightTaskResResultData `json:"Data"`
}

// CreateHighLightTaskResResultData - 参数数据。
type CreateHighLightTaskResResultData struct {

	// REQUIRED; 任务ID。
	TaskID string `json:"TaskID"`
}

type CreateLivePadPresetBody struct {

	// REQUIRED; 垫片时长。取值范围：>=1000。单位ms
	MaxDuration int64 `json:"MaxDuration"`

	// REQUIRED; 垫片类型，1: 图片、2: 视频、3: 源流最后一帧
	PadType int32 `json:"PadType"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// REQUIRED; 断流等待时间。断流等待时间。 取值范围：0-6000。 单位：ms。
	WaitDuration int64 `json:"WaitDuration"`

	// app
	App *string `json:"App,omitempty"`

	// 模板描述，长度上限：1024字节。
	Description *string `json:"Description,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 垫片素材地址。对 源流最后一帧类型无效。
	URL *string `json:"Url,omitempty"`
}

type CreateLivePadPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateLivePadPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateLivePadPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type CreateLiveStreamRecordIndexFilesBody struct {

	// REQUIRED; app名字
	App string `json:"App"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 结束时间
	EndTimeUTC string `json:"EndTimeUTC"`

	// REQUIRED; 开始时间
	StartTimeUTC string `json:"StartTimeUTC"`

	// REQUIRED; stream名称
	Stream string `json:"Stream"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`

	// 是否需要进行FFProbe嗅探,默认开启
	NeedFFProbe *bool `json:"NeedFFProbe,omitempty"`

	// 输出Bucket,默认为对应录制模板下参数
	OutputBucket *string `json:"OutputBucket,omitempty"`

	// 输出文件名Format,默认为对应录制模板下参数
	OutputObject *string `json:"OutputObject,omitempty"`

	// 关联的TS文件是否是独立的
	SeparatedTS *bool `json:"SeparatedTS,omitempty"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty"`

	// 访问网络协议,可选值 http,https
	TSScheme *string `json:"TSScheme,omitempty"`

	// 自定义上传工作流
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type CreateLiveStreamRecordIndexFilesRes struct {

	// REQUIRED
	ResponseMetadata CreateLiveStreamRecordIndexFilesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateLiveStreamRecordIndexFilesResResult `json:"Result"`
}

type CreateLiveStreamRecordIndexFilesResResponseMetadata struct {

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

type CreateLiveStreamRecordIndexFilesResResult struct {

	// REQUIRED; app名字
	App string `json:"App"`

	// REQUIRED; 任务创建时间
	CreateTimeUTC string `json:"CreateTimeUTC"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 录制文件长度,单位 秒
	Duration float32 `json:"Duration"`

	// REQUIRED; 结束时间
	EndTimeUTC string `json:"EndTimeUTC"`

	// REQUIRED; 流嗅探高度
	Height int32 `json:"Height"`

	// REQUIRED; 录制m3u8文件访问路径
	RecordURL string `json:"RecordURL"`

	// REQUIRED; 开始时间
	StartTimeUTC string `json:"StartTimeUTC"`

	// REQUIRED; stream名称
	Stream string `json:"Stream"`

	// REQUIRED; 流嗅探宽度
	Width int32 `json:"Width"`

	// tos上传bucket
	OutputBucket *string `json:"OutputBucket,omitempty"`

	// tos上传文件名
	OutputObject *string `json:"OutputObject,omitempty"`

	// 点播URI
	URI *string `json:"URI,omitempty"`

	// 点播VID
	VID *string `json:"VID,omitempty"`

	// 点播命名空间
	VodNamespace *string `json:"VodNamespace,omitempty"`
}

type CreateLiveVideoQualityAnalysisTaskBody struct {

	// REQUIRED; 源流地址，支持flv、hls、rtmp地址
	SrcURL string `json:"SrcURL"`

	// 任务运行时常，支持60-300，单位s，默认180
	Duration *int32 `json:"Duration,omitempty"`

	// 截图间隔，支持3-10s，单位s，不填默认为10s
	Interval *int32 `json:"Interval,omitempty"`

	// 任务名称，用来查询使用。
	Name *string `json:"Name,omitempty"`
}

type CreateLiveVideoQualityAnalysisTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateLiveVideoQualityAnalysisTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateLiveVideoQualityAnalysisTaskResResult `json:"Result,omitempty"`
}

type CreateLiveVideoQualityAnalysisTaskResResponseMetadata struct {

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

// CreateLiveVideoQualityAnalysisTaskResResult - 视请求的接口而定
type CreateLiveVideoQualityAnalysisTaskResResult struct {

	// REQUIRED; 任务ID
	ID int64 `json:"ID"`

	// REQUIRED; 任务名称
	Name string `json:"Name"`
}

type CreatePullRecordTaskBody struct {

	// REQUIRED; app名字
	App string `json:"App"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; stream名称
	Stream string `json:"Stream"`

	// REQUIRED; 拉流地址
	StreamURL string `json:"StreamURL"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type CreatePullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata CreatePullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result CreatePullRecordTaskResResult `json:"Result"`
}

type CreatePullRecordTaskResResponseMetadata struct {

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

// CreatePullRecordTaskResResult - 视请求的接口而定
type CreatePullRecordTaskResResult struct {

	// REQUIRED; 任务id
	TaskID string `json:"TaskId"`
}

type CreatePullToPushGroupBody struct {

	// REQUIRED; 群组名称，支持有中文、大小写字母和数字组成，最大长度为 20 个字符。
	Name string `json:"Name"`

	// REQUIRED; 为任务群组设置所属项目，您可以在访问控制-项目列表 [https://console.volcengine.com/iam/resourcemanage/project]查看已有项目并对项目进行管理。 项目是火山引擎提供的一种资源管理方式，您可以对不同业务或项目使用的云资源进行分组管理，以实现基于项目的账单查看、子账号授权等功能。
	ProjectName string `json:"ProjectName"`

	// 为任务群组设置资源标签。您可以通过资源标签从不同维度对云资源进行分类和聚合管理，如资源分账等场景。 :::tip 如需使用标签进行资源分账，可以在可以在账单管理-费用标签 [https://console.volcengine.com/finance/bill/tag/]处管理启用标签，将对应标签运用到账单明细等数据中。
	// :::
	Tags []*CreatePullToPushGroupBodyTagsItem `json:"Tags,omitempty"`
}

type CreatePullToPushGroupBodyTagsItem struct {

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type CreatePullToPushGroupRes struct {

	// REQUIRED
	ResponseMetadata CreatePullToPushGroupResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreatePullToPushGroupResResponseMetadata struct {

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

type CreatePullToPushTaskBody struct {

	// REQUIRED; 任务的结束时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 任务的开始时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	StartTime int32 `json:"StartTime"`

	// REQUIRED; 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type int32 `json:"Type"`

	// 推流应用名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	App *string `json:"App,omitempty"`

	// 接收拉流转推任务状态回调的地址，最大长度为 512 个字符，默认为空。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频时为必选参数，参数取值及含义如下所示。
	// * -1：无限次循环，至任务结束；
	// * 0：有限次循环，循环次数以 PlayTimes 取值为准；
	// * >0：有限次循环，循环次数以 CycleMode 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 群组所属名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 获取可用的群组。 :::tip
	// * 使用主账号调用时，为非必填，默认加入 default 群组，default 群组不存在时会默认创建，并绑定 default 项目。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下为空表示不进行偏移。 :::tip 此字段为旧版本配置，请使用 VodSrcAddrs 配置点播视频地址和播放偏移时间。
	// :::
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。拉流来源类型为点播视频时，参数生效。 :::tip 此字段为旧版本配置，请使用 VodSrcAddrs 配置点播视频地址和播放偏移时间。 :::
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。 :::tip PlayTimes 为冗余参数，您可以将 PlayTimes 置 0 后直接使用 CycleMode 指定点播视频文件循环播放次数。
	// :::
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
	// * 0：不开启；
	// * 1：开启（默认值）。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源时，为必传参数，最大长度为 1000 个字符。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频时，为必传参数，最多支持传入 30 个点播视频播放地址，每个地址最大长度为 1000 个字符。
	// :::tip 此字段为旧版本配置，请使用 VodSrcAddrs 配置点播视频地址和播放偏移时间。 :::
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Stream *string `json:"Stream,omitempty"`

	// 拉流转推任务的名称，默认为空表示不配置任务名称。支持由中文、大小写字母（A - Z、a - z）和数字（0 - 9）组成，长度为 1 到 20 各字符。
	Title *string `json:"Title,omitempty"`

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*CreatePullToPushTaskBodyVodSrcAddrsItem `json:"VodSrcAddrs,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *CreatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
}

type CreatePullToPushTaskBodyVodSrcAddrsItem struct {

	// REQUIRED; 点播文件地址。
	SrcAddr string `json:"SrcAddr"`

	// 当前点播文件结束播放的时间偏移值，单位为秒，默认为空时表示结束播放时间不进行偏移。
	EndOffset *float32 `json:"EndOffset,omitempty"`

	// 当前点播文件开始播放的时间偏移值，单位为秒。默认为空时表示开始播放时间不进行偏移。
	StartOffset *float32 `json:"StartOffset,omitempty"`
}

// CreatePullToPushTaskBodyWatermark - 为拉流转推视频添加的水印配置信息。
type CreatePullToPushTaskBodyWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture string `json:"Picture"`

	// REQUIRED; 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
	Ratio float32 `json:"Ratio"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosX float32 `json:"RelativePosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosY float32 `json:"RelativePosY"`
}

type CreatePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata CreatePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreatePullToPushTaskResResult          `json:"Result,omitempty"`
}

type CreatePullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *CreatePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type CreatePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreatePullToPushTaskResResult struct {

	// REQUIRED; 任务 ID，任务的唯一标识。
	TaskID string `json:"TaskId"`
}

type CreateRecordPresetV2Body struct {

	// REQUIRED; 直播流录制配置的详细配置。
	RecordPresetConfig CreateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，由 1 到 30 位数字（0 - 9）、大写小字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，默认为空。 :::tip
	// * App 取值为空时，Stream 取值也需为空，表示录制配置为 Vhost 级别的全局配置。
	// * App 取值不为空时，Stream 取值含义请参见 Stream 参数说明。 :::
	App *string `json:"App,omitempty"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相同，由 1 到 100 位数字（0 - 9）、大写小字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成。
	// :::tip
	// * App 取值不为空、Stream 取值为空时，表示录制配置为 Vhost + App 级别的配置。
	// * App 取值不为空、Stream 取值不为空时，表示录制为 Vhost + App + Stream 的配置。 :::
	Stream *string `json:"Stream,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfig - 直播流录制配置的详细配置。
type CreateRecordPresetV2BodyRecordPresetConfig struct {

	// 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	FlvParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	HlsParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	Mp4Param *CreateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 是否录制源流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tip 转码流和源流需至少选一个进行录制，即 TranscodeRecord 和 OriginRecord 的取值不能同时为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制为 HLS 格式时，单个 TS 切片时长，单位为秒，默认值为 10，取值范围为 [2,100]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 是否录制转码流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制全部转码流；
	// * 2：录制指定转码流，根据转码后缀列表 TranscodeSuffixList决定录制哪些转码流。如果这个列表为空，则效果和设置为 1 一样，即录制所有转码流。
	// :::tip 转码流和源流需至少选一个进行录制，即 TranscodeRecord 和 OriginRecord 的取值不能同时为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码后缀列表。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParam - 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

	// 实时录制场景下，断流等待时长，单位为秒，默认值为 180，取值范围为 [0,3600]。如果实际断流时间小于断流等待时长，录制任务不会停止；如果实际断流时间大于断流等待时长，录制任务会停止，断流恢复后重新开始一个新的录制任务。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制场景下，单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1 和 [300,86400]。
	// * 取值为 -1 时，表示不限制录制时长，录制结束后生成一个完整的录制文件。
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时长，到达时长立即生成录制文件，完成录制后一起上传。
	// :::tip
	// * 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。
	// * 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。
	// :::
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认值为 false，支持的取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制场景下，单文件录制时长，单位为秒，默认值为 1800，取值范围为 [300,21600]。录制时间到达设置的单文件录制时长时，会立即生成录制文件实时上传存储。 :::tip 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。
	// :::
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流录制场景下，断流拼接时长，单位为秒，默认值为 0，支持的取值及含义如下所示。
	// * -1：一直拼接，表示每次断流都不会影响录制任务，录制完成后生成一个完整的录制文件；
	// * 0：不拼接，表示每次断流结束录制任务生成一个录制文件，断流恢复重新开始一个新的录制任务；
	// * 大于 0：拼接容错时间，表示如果断流时间小于拼接容错时间时，则录制任务不会停止，不会生成新的录制文件；如果断流时间大于拼接容错时间，则录制任务停止，断流恢复后重新开始一个新的录制任务。
	// :::tip 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。 :::
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	TOSParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	VODParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam `json:"VODParam,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam - TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam struct {

	// TOS 存储对应的 Bucket。例如，存储位置为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 如果使用 TOS 存储，即 TOSParam
	// 中 Enable 取值为 true 时，Bucket 为必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件存储到 TOS 时的存储路径和文件名规则。支持输入字母（A - Z、a - z）、数字（0 - 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。最大长度为 200 个字符，
	// 支持以下字段作为占位符：
	// * record：自定义字段，可遵照支持字符进行自定义。
	// * {PubDomain}：当前配置中的 vhost 值。
	// * {App}：当前配置中的 AppName 值。
	// * {Stream}：当前配置中的 StreamName 值。
	// * {StartTime}：录制开始的 Unix 时间戳，精度为 s。
	// * {EndTime}：录制结束的 Unix 时间戳，精度为 s。
	// 存储路径必须至少包含两级目录。例如：live/{App}/{Stream}
	// 合法示例：
	// record/{PubDomain}/{App}/{Stream}/{StartTime}-{EndTime}
	// {App}/archive/{Stream}/recording_{StartTime}
	// vod/{Stream}/!highlight_{EndTime}
	// a/b/custom_record
	// 错误示例：
	// single_level # 错误：路径层级不足两级
	// invalid_/{S@ream}/file # 错误：含非法字符@
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储对应 Bucket 下的存储目录，默认为空。例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam - VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息，默认为空。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，支持的取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储规则，最大长度为 200 个字符，支持以record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime} 样式设置存储规则，支持输入字母（A - Z、a - z）、数字（0 -
	// 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。
	// 存储规则设置注意事项如下。
	// * 目录层级至少包含2级及以上，如live/{App}/{Stream}。
	// * record 为自定义字段；
	// * {PubDomain} 取值为当前配置的 vhost 值；
	// * {App} 取值为当前配置的 AppName 值；
	// * {Stream} 取值为当前配置的 StreamName 值；
	// * {StartTime} 取值为录制的开始时间戳；
	// * {EndTime} 取值为录制的结束时间戳。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型，存储类型介绍请参考媒资存储管理 [https://www.volcengine.com/docs/4/73629#媒资存储]。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果使用 VOD 存储，即 VODParam 中 Enable 取值为 true 时，VodNamespace
	// 为必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 视频点播工作流模板 ID，对于存储在点播的录制文件，会使用该工作流模版对录制的视频进行处理，可登录视频点播控制台 [https://console.volcengine.com/vod/]获取工作流模板 ID，默认为空。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParam - 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParam struct {

	// 断流等待时长，取值范围[0,3600]
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1，[300,86400]，-1 表示一直录制，目前只对 HLS生效.
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam `json:"VODParam,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam - TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam - VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace 必填。
	// :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigMp4Param - 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigMp4Param struct {

	// 断流等待时长，取值范围[0,3600]
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为 s，默认值为 7200，取值范围为 -1，[300,86400]，-1表示一直录制，目前只对HLS生效
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *CreateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *CreateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam `json:"VODParam,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam - TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam - VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace 必填。
	// :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type CreateRecordPresetV2Res struct {

	// REQUIRED
	ResponseMetadata CreateRecordPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateRecordPresetV2ResResult `json:"Result,omitempty"`
}

type CreateRecordPresetV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                        `json:"Version"`
	Error   *CreateRecordPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateRecordPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// CreateRecordPresetV2ResResult - 视请求的接口而定
type CreateRecordPresetV2ResResult struct {

	// REQUIRED; 录制配置 ID。
	PresetName string `json:"PresetName"`
}

type CreateRelaySourceV4Body struct {

	// REQUIRED; 应用名称，即直播流地址的AppName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 回源地址列表，支持输入 RTMP、FLV、HLS 和 SRT 协议的直播推流地址。 :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 流名称，即直播流地址的StreamName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、字母、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 回源结束时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	EndTime *int32 `json:"EndTime,omitempty"`

	// 回源开始时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	StartTime *int32 `json:"StartTime,omitempty"`
}

type CreateRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata CreateRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateRelaySourceV4ResResult          `json:"Result,omitempty"`
}

type CreateRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                       `json:"Version"`
	Error   *CreateRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateRelaySourceV4ResResult struct {

	// REQUIRED; 固定回源配置的 ID。
	TaskID string `json:"TaskId"`
}

type CreateSnapshotAuditPresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 截图间隔时间，单位为秒，取值范围为 [0.1,10]，支持保留两位小数。
	Interval float32 `json:"Interval"`

	// REQUIRED; 存储策略，支持的取值及含义如下。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储所有图片。
	StorageStrategy int32 `json:"StorageStrategy"`

	// TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 参数 Bucket 和 ServiceID
	// 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图审核配置的描述。
	Description *string `json:"Description,omitempty"`

	// 推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	// :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 审核标签，缺省情况下取值为 301、302、303、305 和 306，支持的取值及含义如下。
	// * 301：涉黄；
	// * 302：涉敏1；
	// * 303：涉敏2；
	// * 304：广告；
	// * 305：引人不适；
	// * 306：违禁；
	// * 307：二维码；
	// * 308：诈骗；
	// * 309：不良画面；
	// * 310：未成年相关；
	// * 320：文字违规。
	Label []*string `json:"Label,omitempty"`

	// veImageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符，最大长度为 180 个字符，默认值为 {audit}/{PushDomain}/{App}/{Stream}/{UnixTimestamp}。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// ToS 存储对应的 bucket 下的存储目录，默认为空。 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	// :::tip 参数 Domain
	// 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type CreateSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateSnapshotAuditPresetResResult `json:"Result,omitempty"`
}

type CreateSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *CreateSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type CreateSnapshotAuditPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// CreateSnapshotAuditPresetResResult - 视请求的接口而定
type CreateSnapshotAuditPresetResResult struct {

	// REQUIRED; 模板名称。
	PresetName string `json:"PresetName"`
}

type CreateSnapshotPresetV2Body struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 截图配置的详细参数配置。
	SnapshotPresetConfig CreateSnapshotPresetV2BodySnapshotPresetConfig `json:"SnapshotPresetConfig"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 截图配置的生效状态，取值及含义如下所示。
	// * 1：（默认值）生效；
	// * 0：不生效。
	Status *int32 `json:"Status,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfig - 截图配置的详细参数配置。
type CreateSnapshotPresetV2BodySnapshotPresetConfig struct {

	// 截图间隔时间，单位为秒，默认值为 10，取值范围为正整数。
	Interval *int32 `json:"Interval,omitempty"`

	// 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JPEGParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam `json:"JpegParam,omitempty"`

	// 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JpgParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParam `json:"JpgParam,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam - 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam struct {

	// 当前格式的截图是否开启，取值及含义如下所示。
	// * false：（默认值）不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	ImageXParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam `json:"ImageXParam,omitempty"`

	// 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	TOSParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam `json:"TOSParam,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam - 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam
// 配置且配置其中一个。 :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam struct {

	// 截图是否使用 veImageX 存储，取值及含义如下所示。
	// * false：（默认值）不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// 使用 veImageX 存储截图时，对应的 veImageX 的服务 ID。 :::tip 使用 veImageX 存储时 ServiceID 为必填项。 :::
	ServiceID *string `json:"ServiceID,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam - 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。
// :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam struct {

	// TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 使用 TOS 存储时 Bucket
	// 为必填项。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图是否使用 TOS 存储，取值及含义如下所示。
	// * false：（默认值）不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// ToS 存储对应的 Bucket 下的存储目录，默认为空。 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParam - 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParam struct {
	Enable      *bool                                                              `json:"Enable,omitempty"`
	ImageXParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam `json:"ImageXParam,omitempty"`
	TOSParam    *CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam    `json:"TOSParam,omitempty"`
}

type CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam struct {
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	ServiceID       *string `json:"ServiceID,omitempty"`
}

type CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam struct {
	Bucket          *string `json:"Bucket,omitempty"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	StorageDir      *string `json:"StorageDir,omitempty"`
}

type CreateSnapshotPresetV2Res struct {

	// REQUIRED
	ResponseMetadata CreateSnapshotPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateSnapshotPresetV2ResResult `json:"Result,omitempty"`
}

type CreateSnapshotPresetV2ResResponseMetadata struct {

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

// CreateSnapshotPresetV2ResResult - 视请求的接口而定
type CreateSnapshotPresetV2ResResult struct {

	// REQUIRED; 截图配置名称。
	PresetName string `json:"PresetName"`
}

type CreateSubtitleTranscodePresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 源语言参数
	SourceLanguage CreateSubtitleTranscodePresetBodySourceLanguage `json:"SourceLanguage"`

	// REQUIRED; 关联的转码配置后缀，一个字幕配置支持关联多个转码配置后缀。
	Suffixes []string `json:"Suffixes"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 字幕配置的描述信息。
	Description *string `json:"Description,omitempty"`

	// 预设配置，使用预设配置是系统将自动对字体大小、字幕行数、每行最大字符数和边距参数（MarginVertical 和 MarginHorizontal）进行智能化适配。默认为空，表示不使用预设配置，支持的预设配置如下所示。
	// * small ：小字幕。
	// * medium：中字幕。
	// * large：大字幕。 :::tip 使用预设配置时，字幕行数、每行最大字符数、左右边距和底部边距参数不生效，系统将使用预设配置自动进行计算。 :::
	DisplayPreset *string `json:"DisplayPreset,omitempty"`

	// 原文翻译成译文时使用的热词词库，总长度不超过 10000 个字符，默认为空。
	GlossaryWordList []*string `json:"GlossaryWordList,omitempty"`

	// 原文字幕识别时使用的热词词库，总长度不超过为 10000 个字符，默认为空。
	HotWordList []*string `json:"HotWordList,omitempty"`

	// 设置在 16:9 分辨率场景下，每行字幕展示的最大字符数。 :::tip
	// * 使用预设配置时，字幕每行最大字符数设置不生效。
	// * 不使用预设配置时，字幕每行最大字符数必填。
	// * 每个文字、字母、符号或数字均为一个字符。
	// * 当屏幕分辨率改变时，屏幕上显示的每行文字数量会相应调整，以适应新的分辨率，确保文字的显示效果和阅读体验。 :::
	MaxCharNumber *int32 `json:"MaxCharNumber,omitempty"`

	// 字幕展示的行数，同时适用于原文字幕和译文字幕，支持的取值及含义如下所示。
	// * 0：（默认值）根据字幕字数自动进行分行展示；
	// * 1：每种字幕展示一行；
	// * 2：每种字幕展示两行。 :::tip
	// * 使用预设配置时，字幕行数为自动分行展示。
	// * 超出行内字数限制时表示字幕将超过显示范围，此时字幕内容将被截断。 :::
	MaxRowNumber *int32 `json:"MaxRowNumber,omitempty"`

	// 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
	// :::tip
	// * 使用预设配置时，字幕位置设置不生效。
	// * 不使用预设配置时，字幕位置设置必填。 :::
	Position *CreateSubtitleTranscodePresetBodyPosition `json:"Position,omitempty"`

	// 字幕配置的名称，不可以与其他已有的配置名称重复。默认为空，表示由系统将自动分配配置名称。
	PresetName *string `json:"PresetName,omitempty"`

	// 译文字幕展示参数配置列表，当前最多支持配置一种译文。
	TargetLanguage []*CreateSubtitleTranscodePresetBodyTargetLanguageItem `json:"TargetLanguage,omitempty"`
}

// CreateSubtitleTranscodePresetBodyPosition - 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
// :::tip
// * 使用预设配置时，字幕位置设置不生效。
// * 不使用预设配置时，字幕位置设置必填。 :::
type CreateSubtitleTranscodePresetBodyPosition struct {

	// 左右边距，[0,0.2]
	MarginHorizontal *float32 `json:"MarginHorizontal,omitempty"`

	// 上下边距
	MarginVertical *float32 `json:"MarginVertical,omitempty"`
}

// CreateSubtitleTranscodePresetBodySourceLanguage - 源语言参数
type CreateSubtitleTranscodePresetBodySourceLanguage struct {

	// REQUIRED; 是否展示源语言
	Display bool `json:"Display"`

	// REQUIRED; 字体
	Font string `json:"Font"`

	// REQUIRED; 字体颜色
	FontColor string `json:"FontColor"`

	// REQUIRED; 原文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语。
	Language string `json:"Language"`

	// 字幕阴影配置
	Border *CreateSubtitleTranscodePresetBodySourceLanguageBorder `json:"Border,omitempty"`
}

// CreateSubtitleTranscodePresetBodySourceLanguageBorder - 字幕阴影配置
type CreateSubtitleTranscodePresetBodySourceLanguageBorder struct {

	// REQUIRED; 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Width float32 `json:"Width"`

	// 阴影颜色
	Color *string `json:"Color,omitempty"`
}

type CreateSubtitleTranscodePresetBodyTargetLanguageItem struct {

	// REQUIRED; 字体，覆盖全局参数
	Font string `json:"Font"`

	// REQUIRED; 字体颜色，覆盖全局参数
	FontColor string `json:"FontColor"`

	// REQUIRED; 译文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * zh-Hant：繁体中文；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语；
	// * ar：阿拉伯语；
	// * de：德语；
	// * es：西班牙语；
	// * fr：法语；
	// * hi：印地语；
	// * pt：葡萄牙语；
	// * ru：俄语；
	// * vi：越南语；
	// * th：泰语。
	Language string `json:"Language"`

	// 字幕阴影配置
	Border *CreateSubtitleTranscodePresetBodyTargetLanguageItemBorder `json:"Border,omitempty"`
}

// CreateSubtitleTranscodePresetBodyTargetLanguageItemBorder - 字幕阴影配置
type CreateSubtitleTranscodePresetBodyTargetLanguageItemBorder struct {

	// REQUIRED; 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Width float32 `json:"Width"`

	// 阴影颜色
	Color *string `json:"Color,omitempty"`
}

type CreateSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata CreateSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateSubtitleTranscodePresetResResponseMetadata struct {

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

type CreateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即允许用户回看的最长时间，单位为秒，支持的取值如下所示。
	// * 86400：1 天；
	// * 259200：3 天；
	// * 604800：7 天；
	// * 1296000：15 天。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 时移拉流域名。 :::tip 录制到 TOS 时，需配置直播流对应的拉流域名。录制到 VOD 时，需配置 VOD 的空间对应的分发域名。 :::
	PullDomain string `json:"PullDomain"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要时移的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示当前应用下的所有流都进行时移。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100
	// 个字符。 :::tip 流名称不为空时，表示只为此条流开启时移。通过流名称配置时移时，同一个 App 最多可指定 20 路流开启时移。 :::
	Stream *string `json:"Stream,omitempty"`

	// 时移类型。支持的取值如下所示。
	// * 0：录制时移，即时移复用录制配置，需提前创建对应级别的录制配置；
	// * 1：独立时移，即时移不复用录制配置。
	TimeShiftType *int32 `json:"TimeShiftType,omitempty"`
}

type CreateTimeShiftPresetV3Res struct {

	// REQUIRED
	ResponseMetadata CreateTimeShiftPresetV3ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateTimeShiftPresetV3ResResult `json:"Result,omitempty"`
}

type CreateTimeShiftPresetV3ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *CreateTimeShiftPresetV3ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateTimeShiftPresetV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// CreateTimeShiftPresetV3ResResult - 视请求的接口而定
type CreateTimeShiftPresetV3ResResult struct {

	// REQUIRED; 模板名称。
	PresetName string `json:"PresetName"`
}

type CreateTranscodePresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 转码后缀，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 的视频编码格式；
	// * h265：使用 H.265 的视频编码格式；
	// * h266：使用 H.266 的视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP 和 BFrames
	// 等。
	Vcodec string `json:"Vcodec"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要转码的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 音频编码格式，默认值为 aac，支持的取值及含义如下所示。
	// * aac：使用 AAC 音频编码格式；
	// * opus：使用 Opus 音频编码格式。
	// * copy：不进行音频转码，所有音频编码参数不生效，音频编码参数包括音频码率（AudioBitrate）等。
	Acodec *string `json:"Acodec,omitempty"`

	// 视频分辨率自适应模式开关，默认值为 0。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As 取值为 0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As 取值为 1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps，默认值为 128，取值范围为 [0,1000]；取值为 0 时，表示与源流的音频码率相同。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码视频分辨率不超过源流分辨率，默认值为 1 表示开启。开启后，当源流分辨率低于转码配置分辨率时（即源流宽低于转码配置宽且源流高低于转码配置高时），将按源流视频分辨率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启转码视频码率不超过源流码率，默认值为 1 表示开启。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启转码视频帧率不超过源流帧率，默认值为 1 表示开启。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，默认值为 3，取值为 0 时表示去除 B 帧。
	// 最大 B 帧数量的取值范围根据视频编码格式（Vcodec）的不同有所差异，取值范围如下所示。
	// * 视频编码格式为 H.264 （Vcodec 取值为 h264）时取值范围为 [0,7]；
	// * 视频编码格式为 H.265 或 H.266 （Vcodec 取值为 h265 或 h266）时取值范围为 [0,3]、7、15。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，取值为 0 时表示与源流视频帧率相同。
	// 视频帧率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频帧率取值范围为 [0,60]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频帧率取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,30]。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度，默认值为 0。
	// 视频高度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频高度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,7680]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,1280]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。
	//
	//
	// :::tip
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,2160]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,720]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。 :::tip
	//
	//
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码（TransType 取值为 Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码触发方式，默认值为 Pull，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频码率，单位为 bps，默认值为 1000000；取值为 0 时，表示与源流的视频码率相同。
	// 视频码率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频码率取值范围为 [0,30000000]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频码率取值范围为 [0,6000000]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度，单位为 px，默认值为 0。
	// 视频宽度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频宽度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type CreateTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata CreateTranscodePresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateTranscodePresetResResult          `json:"Result,omitempty"`
}

type CreateTranscodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                         `json:"Version"`
	Error     *CreateTranscodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                        `json:"RequestID,omitempty"`
}

type CreateTranscodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateTranscodePresetResResult struct {

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`
}

type CreateWatermarkPresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX float32 `json:"PosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY float32 `json:"PosY"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 需要添加水印的直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片编码字符串字段二选一传入，同时传入时，以水印图片编码字符串参数为准。
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。与水印相对宽度字段冲突，请选择其中一个传参。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。与水印相对高度字段冲突，请选择其中一个传参。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	// :::tip
	// * 默认为空，表示对指定的 AppName 下所有转码流均使用当前水印配置。
	// * 指定流名称时，表示仅对 AppName 下指定流名称的转码流使用当前水印配置。 :::
	Stream *string `json:"Stream,omitempty"`
}

type CreateWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateWatermarkPresetResResult          `json:"Result,omitempty"`
}

type CreateWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *CreateWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type CreateWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateWatermarkPresetResResult struct {

	// REQUIRED; id
	ID int32 `json:"ID"`

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`
}

type CreateWatermarkPresetV2Body struct {

	// 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// :::tip Picture 与 PictureUrl 字段二选一必传，同时传入时，以 Picture 参数为准。 :::
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。 :::tip Picture 与 PictureUrl 字段二选一必传，同时传入时，以 Picture 参数为准。 :::
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]，默认为 0。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]，默认为 0。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印模板的名称标识，不可重复，默认传空将由系统为您自动生成。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_） 和短横线（-）组成，长度为 1 到 20 个字符。
	PresetName *string `json:"PresetName,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。 :::tip RelativeWidth 与 RelativeHeight 二选一必传。 :::
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。 :::tip RelativeWidth 与 RelativeHeight 二选一必传。 :::
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`
}

type CreateWatermarkPresetV2Res struct {

	// REQUIRED
	ResponseMetadata CreateWatermarkPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateWatermarkPresetV2ResResult `json:"Result,omitempty"`
}

type CreateWatermarkPresetV2ResResponseMetadata struct {

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

// CreateWatermarkPresetV2ResResult - 视请求的接口而定
type CreateWatermarkPresetV2ResResult struct {

	// REQUIRED; 水印模板的 ID。
	ID int32 `json:"ID"`

	// REQUIRED; 水印模板的名称标识。
	PresetName string `json:"PresetName"`
}

type DeleteCallbackBody struct {

	// 应用名称，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 推流域名，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 Domain 取值。
	Domain *string `json:"Domain,omitempty"`

	// 消息类型，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 MessageType 取值。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制任务状态回调；
	// * audit_snapshot：截图审核结果回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 Vhost 取值。
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteCallbackRes struct {

	// REQUIRED
	ResponseMetadata DeleteCallbackResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteCallbackResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                  `json:"Version"`
	Error   *DeleteCallbackResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteCallbackResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteCarouselTaskBody struct {

	// REQUIRED; 待删除的轮播任务 ID，任务的唯一标识。调用 CreateCarouselTask 接口创建轮播任务时返回。
	TaskID string `json:"TaskID"`
}

type DeleteCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata DeleteCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DeleteCarouselTaskResResult `json:"Result"`
}

type DeleteCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type DeleteCarouselTaskResResult struct {

	// REQUIRED; 删除任务的操作结果信息，返回任务是否成功删除以及相关的 Mesos ID 和操作影响记录数。
	Data string `json:"Data"`
}

type DeleteCertBody struct {

	// REQUIRED; 待删除的 HTTPS 证书的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/1126822]接口获取。
	ChainID string `json:"ChainID"`
}

type DeleteCertRes struct {

	// REQUIRED
	ResponseMetadata DeleteCertResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *DeleteCertResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteCloudMixTaskBody struct {

	// REQUIRED; 混流任务 ID，您可以通过 ListCloudMixTask [https://www.volcengine.com/docs/6469/1271157] 接口获取待结束的混流任务 ID。
	TaskID string `json:"TaskID"`
}

type DeleteCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata DeleteCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DeleteCloudMixTaskResResult `json:"Result"`
}

type DeleteCloudMixTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type DeleteCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data string `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

type DeleteDomainBody struct {

	// REQUIRED; 待删除域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要待删除域名的信息。
	Domain string `json:"Domain"`
}

type DeleteDomainRes struct {

	// REQUIRED
	ResponseMetadata DeleteDomainResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *DeleteDomainResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteHTTPHeaderConfigBody struct {

	// REQUIRED; HTTP Header 类型，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header
	// 配置的 Phase 取值。
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Vhost
	// 取值。
	Vhost string `json:"Vhost"`

	// 拉流域名，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Domain 取值。
	Domain *string `json:"Domain,omitempty"`
}

type DeleteHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteHTTPHeaderConfigResResponseMetadata struct {

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

type DeleteIPAccessRuleBody struct {

	// REQUIRED; 推流域名或拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要删除
	// IP 访问限制的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要删除
	// IP 访问限制的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DeleteIPAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DeleteIPAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteIPAccessRuleResResponseMetadata struct {

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

type DeleteLivePadPresetBody struct {

	// REQUIRED; 预设名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DeleteLivePadPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteLivePadPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteLivePadPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteLiveVideoQualityAnalysisTaskBody struct {

	// 任务ID，和任务名二选一
	ID *int64 `json:"ID,omitempty"`

	// 任务名，和任务ID二选一
	Name *string `json:"Name,omitempty"`
}

type DeleteLiveVideoQualityAnalysisTaskRes struct {

	// REQUIRED
	ResponseMetadata DeleteLiveVideoQualityAnalysisTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteLiveVideoQualityAnalysisTaskResResponseMetadata struct {

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

type DeletePullToPushGroupBody struct {

	// REQUIRED; 拉流转推群组名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 接口获取群组名称。
	Name string `json:"Name"`
}

type DeletePullToPushGroupRes struct {

	// REQUIRED
	ResponseMetadata DeletePullToPushGroupResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeletePullToPushGroupResResponseMetadata struct {

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

type DeletePullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
	TaskID string `json:"TaskId"`

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
}

type DeletePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata DeletePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type DeletePullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *DeletePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type DeletePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRecordPresetBody struct {

	// REQUIRED; 录制配置的名称。可调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待删除录制配置的名称。
	Preset string `json:"Preset"`

	// 应用名称，您可以调用ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]接口查看待删除的录制配置 App 取值。
	App *string `json:"App,omitempty"`

	// 流名称。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待删除录制配置的 Stream 取值。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待删除录制配置的 Vhost 取值。
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteRecordPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteRecordPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRecordPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *DeleteRecordPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRecordPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRefererBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// Referer 的拉流域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// * 如创建时传了 App，删除时需要传该参数；
	// * 如创建时未传 App，删除时不传该参数。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// Referer 的拉流域名。 :::tip
	// * 如创建时传了 Domain，删除时需要传该参数；
	// * 如创建时未传 Domain，删除时不传该参数。 :::
	Domain *string `json:"Domain,omitempty"`
}

type DeleteRefererRes struct {

	// REQUIRED
	ResponseMetadata DeleteRefererResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRefererResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                 `json:"Version"`
	Error   *DeleteRefererResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRefererResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRelaySourceV3Body struct {

	// REQUIRED; 直播流使用的域名所属的域名空间，您可以调用DescribeRelaySourceV3 [https://www.volcengine.com/docs/6469/1126874]接口获取待删除配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 应用名称，您可以调用DescribeRelaySourceV3 [https://www.volcengine.com/docs/6469/1126874]接口获取待删除配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 回源组名称。
	Group *string `json:"Group,omitempty"`
}

type DeleteRelaySourceV3Res struct {

	// REQUIRED
	ResponseMetadata DeleteRelaySourceV3ResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteRelaySourceV3ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *DeleteRelaySourceV3ResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRelaySourceV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRelaySourceV4Body struct {

	// REQUIRED; 应用名称，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待删除固定回源配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待删除固定回源配置的 Domain 取值。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待删除固定回源配置的 Stream 取值。
	Stream string `json:"Stream"`
}

type DeleteRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata DeleteRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                       `json:"Version"`
	Error   *DeleteRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteSnapshotAuditPresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostSnapshotAuditPreset [https://www.volcengine.com/docs/6469/1126870]接口，获取待删除截图配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 截图审核配置名称，您可以调用ListVhostSnapshotAuditPreset [https://www.volcengine.com/docs/6469/1126870]接口，获取待删除截图配置的 PresetName
	// 取值。
	PresetName string `json:"PresetName"`

	// REQUIRED; 域名空间，您可以调用ListVhostSnapshotAuditPreset [https://www.volcengine.com/docs/6469/1126870]接口，获取待删除截图配置的 Vhost 取值。
	Vhost string `json:"Vhost"`
}

type DeleteSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *DeleteSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteSnapshotAuditPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteSnapshotPresetBody struct {

	// REQUIRED; 截图配置的名称，您可以调用 ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858] 接口获取，取值与 Name 字段取值相同。
	Preset string `json:"Preset"`

	// 应用名称，您可以调用ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858]接口，获取待更新截图配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 域名空间，您可以调用 ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858] 接口，获取待删除截图配置的 Vhost 取值。
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteSnapshotPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *DeleteSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteStreamQuotaConfigBody struct {

	// REQUIRED; 待删除限额配置的推流域名或拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名信息。
	Domain string `json:"Domain"`
}

type DeleteStreamQuotaConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteStreamQuotaConfigResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteStreamQuotaConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *DeleteStreamQuotaConfigResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteStreamQuotaConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteSubtitleTranscodePresetBody struct {

	// REQUIRED; 火山必填
	App string `json:"App"`

	// REQUIRED; 截图配置的名称，您可以调用 ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712] 接口，获取待删除字幕配置的 PresetName
	// 取值。
	PresetName string `json:"PresetName"`

	// REQUIRED; 火山必填
	Vhost string `json:"Vhost"`
}

type DeleteSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteSubtitleTranscodePresetResResponseMetadata struct {

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

type DeleteTaskByAccountIDBody struct {

	// REQUIRED; 任务ID。
	TaskID string `json:"TaskID"`
}

type DeleteTaskByAccountIDRes struct {

	// REQUIRED
	ResponseMetadata DeleteTaskByAccountIDResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 操作结果。
	Result DeleteTaskByAccountIDResResult `json:"Result"`
}

type DeleteTaskByAccountIDResResponseMetadata struct {

	// REQUIRED; 请求的唯一标识符。
	RequestID string `json:"RequestID"`
}

// DeleteTaskByAccountIDResResult - 操作结果。
type DeleteTaskByAccountIDResResult struct {

	// REQUIRED; 任务删除操作的详细信息。
	Data string `json:"Data"`
}

type DeleteTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待删除时移配置的App取值。
	App string `json:"App"`

	// REQUIRED; 域名空间名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待删除时移配置的Vhost取值。
	Vhost string `json:"Vhost"`
}

type DeleteTimeShiftPresetV3Res struct {

	// REQUIRED
	ResponseMetadata DeleteTimeShiftPresetV3ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteTimeShiftPresetV3ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *DeleteTimeShiftPresetV3ResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteTimeShiftPresetV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteTranscodePresetBody struct {

	// REQUIRED; 应用名称，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待删除转码配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 转码配置名称，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待删除转码配置的 Preset 取值。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待删除转码配置的 Vhost 取值。
	Vhost string `json:"Vhost"`
}

type DeleteTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteTranscodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                         `json:"Version"`
	Error     *DeleteTranscodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                        `json:"RequestID,omitempty"`
}

type DeleteTranscodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteWatermarkPresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待删除水印配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 域名空间，您可以调用 ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889] 接口，查看待删除水印配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 流名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待删除水印配置的 Stream 取值。
	Stream *string `json:"Stream,omitempty"`
}

type DeleteWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *DeleteWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteWatermarkPresetV2Body struct {

	// 水印模板的 ID，您可以调用 ListWatermarkPresetDetail [https://www.volcengine.com/docs/6469/1323353] 接口获取。 :::tip PresetName 和 ID 二选一必填。
	// :::
	ID *int32 `json:"ID,omitempty"`

	// 水印模板的名称，您可以调用 ListWatermarkPresetDetail [https://www.volcengine.com/docs/6469/1323353] 接口获取。 :::tip PresetName 和 ID 二选一必填。
	// :::
	PresetName *string `json:"PresetName,omitempty"`
}

type DeleteWatermarkPresetV2Res struct {

	// REQUIRED
	ResponseMetadata DeleteWatermarkPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteWatermarkPresetV2ResResponseMetadata struct {

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

type DescribeAuthBody struct {

	// REQUIRED; 鉴权场景类型，取值及含义如下所示。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty"`

	// 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看待配置鉴权的推拉流域名。
	// :::tip 参数 Domain 和 Vhost 传且仅传一个。
	// :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	// :::tip 参数
	// Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeAuthRes struct {

	// REQUIRED
	ResponseMetadata DescribeAuthResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeAuthResResult          `json:"Result,omitempty"`
}

type DescribeAuthResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *DescribeAuthResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeAuthResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeAuthResResult struct {

	// 推/拉流鉴权列表。
	AuthList []*DescribeAuthResResultAuthListItem `json:"AuthList,omitempty"`
}

type DescribeAuthResResultAuthListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 是否开启 URL 地址鉴权，取值及含义如下所示。
	// * false：关闭；
	// * true：开启。
	AuthStatus bool `json:"AuthStatus"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 鉴权场景类型，取值及含义如下所示。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// REQUIRED; 时间戳进制。取值如下：
	// * 2：二进制
	// * 8：八进制
	// * 10：十进制
	// * 16：十六进制
	// :::tipSceneType 取值为 push 时，该参数取值固定为 10。 :::
	TimeStampBase int32 `json:"TimeStampBase"`

	// REQUIRED; 鉴权生效时长，单位为秒。
	ValidDuration int32 `json:"ValidDuration"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 鉴权详情。
	AuthDetailList []*DescribeAuthResResultAuthListPropertiesItemsItem `json:"AuthDetailList,omitempty"`
}

// DescribeAuthResResultAuthListPropertiesItemsItem - 鉴权详情。
type DescribeAuthResResultAuthListPropertiesItemsItem struct {

	// 自定义推拉流地址中，鉴权参数volcSecret和volcTime的名称。
	AuthField map[string]*string `json:"AuthField,omitempty"`

	// 鉴权模版类型
	AuthType *string `json:"AuthType,omitempty"`

	// 生成加密字符串使用的加密字段。
	EncryptField []*string `json:"EncryptField,omitempty"`

	// 对称加密算法。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty"`

	// 自定义鉴权密钥。
	SecretKey *string `json:"SecretKey,omitempty"`
}

type DescribeCDNSnapshotHistoryBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip
	// * 当您查询指定截图任务详情时，DateFrom 应设置为推流开始时间之前的任意时间。
	// * 查询的最大时间跨度为 7 天。 :::
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	DateTo string `json:"DateTo"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认为 10，最大值为 1000。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 截图文件保存位置，取值及含义如下所示。
	// * tos：（默认值）TOS 对象存储服务；
	// * imageX：veImageX 图片服务。
	Type *string `json:"Type,omitempty"`
}

type DescribeCDNSnapshotHistoryRes struct {

	// REQUIRED
	ResponseMetadata DescribeCDNSnapshotHistoryResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeCDNSnapshotHistoryResResult          `json:"Result,omitempty"`
}

type DescribeCDNSnapshotHistoryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                              `json:"Version"`
	Error   *DescribeCDNSnapshotHistoryResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeCDNSnapshotHistoryResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeCDNSnapshotHistoryResResult struct {

	// REQUIRED; 分页信息。
	Pagination DescribeCDNSnapshotHistoryResResultPagination `json:"Pagination"`

	// 截图文件信息。
	Data []*DescribeCDNSnapshotHistoryResResultDataItem `json:"Data,omitempty"`
}

type DescribeCDNSnapshotHistoryResResultDataItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 截图高度，单位为 px。
	Height int32 `json:"Height"`

	// REQUIRED; 截图文件保存的路径。
	Path string `json:"Path"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 截图时间戳，精度为毫秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// REQUIRED; 截图宽度，单位为 px。
	Width int32 `json:"Width"`
}

// DescribeCDNSnapshotHistoryResResultPagination - 分页信息。
type DescribeCDNSnapshotHistoryResResultPagination struct {

	// REQUIRED; 查询数据的页码。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页显示的数据量条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeCallbackBody struct {

	// domain, app二选一必传
	App *string `json:"App,omitempty"`

	// domain, app二选一必传
	Domain *string `json:"Domain,omitempty"`

	// 回调类型。默认为空，表示查询全部回调类型，取值及含义如下所示。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	// :::tipVhost和
	// Domain传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeCallbackRes struct {

	// REQUIRED
	ResponseMetadata DescribeCallbackResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeCallbackResResult          `json:"Result,omitempty"`
}

type DescribeCallbackResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                    `json:"Version"`
	Error   *DescribeCallbackResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeCallbackResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeCallbackResResult struct {

	// 回调列表。
	CallbackList []*DescribeCallbackResResultCallbackListItem `json:"CallbackList,omitempty"`
}

type DescribeCallbackResResultCallbackListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 回调消息发送是否开启鉴权，默认为false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	AuthEnable bool `json:"AuthEnable"`

	// REQUIRED; 回调消息发送鉴权密钥，开启回调消息鉴权时生效。
	AuthKeyPrimary string `json:"AuthKeyPrimary"`

	// REQUIRED; 回调创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 格式为rfc3339，时区为utc的回调创建时间，
	CreateTimeUTC string `json:"CreateTimeUTC"`

	// REQUIRED; 回调的消息类型，取值及含义如下所示。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType string `json:"MessageType"`

	// REQUIRED; 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：不开启；
	// * 1：开启。
	TranscodeCallback int32 `json:"TranscodeCallback"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 回调数据列表。
	CallbackDetailList []*DescribeCallbackResResultCallbackListPropertiesItemsItem `json:"CallbackDetailList,omitempty"`
}

type DescribeCallbackResResultCallbackListPropertiesItemsItem struct {

	// REQUIRED; 回调类型，返回 HTTP，表示可以使用 HTTP 和 HTTPS 地址接收回调消息。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调消息接收地址。
	URL string `json:"URL"`
}

type DescribeCertDRMQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost" query:"Vhost"`
}

type DescribeCertDRMRes struct {

	// REQUIRED
	ResponseMetadata DescribeCertDRMResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DescribeCertDRMResResponseMetadata struct {

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

type DescribeCertDetailSecretV2Body struct {

	// 证书 ID，您可以通过ListCertV2 [https://www.volcengine.com/docs/6469/1126823]接口获取证书 ID。 :::tip 参数ChainID与CertID传且仅传一个。 :::
	CertID *string `json:"CertID,omitempty"`

	// 证书链 ID，您可以通过ListcCertV2 [https://www.volcengine.com/docs/6469/1126823]接口获取 证书链 ID。 :::tip 参数ChainID与CertID传且仅传一个。 :::
	ChainID *string `json:"ChainID,omitempty"`
}

type DescribeCertDetailSecretV2Res struct {

	// REQUIRED
	ResponseMetadata DescribeCertDetailSecretV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeCertDetailSecretV2ResResult `json:"Result,omitempty"`
}

type DescribeCertDetailSecretV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                              `json:"Version"`
	Error   *DescribeCertDetailSecretV2ResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeCertDetailSecretV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// DescribeCertDetailSecretV2ResResult - 视请求的接口而定
type DescribeCertDetailSecretV2ResResult struct {

	// REQUIRED; 与证书绑定的域名列表。
	CertDomainList []string `json:"CertDomainList"`

	// REQUIRED; 证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 证书链 ID。
	ChainID string `json:"ChainID"`

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为秒。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为秒。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 证书状态，取值及含义如下所示。
	// * OK：正常；
	// * Expire：过期；
	// * 30days：有效期剩余 30 天；
	// * 15days：有效期剩余 15 天；
	// * 7days：有效期剩余 7 天；
	// * 1days：有效期剩余 1 天。
	Status string `json:"Status"`

	// 证书详细信息。
	SSL *DescribeCertDetailSecretV2ResResultSSL `json:"SSL,omitempty"`
}

// DescribeCertDetailSecretV2ResResultSSL - 证书详细信息。
type DescribeCertDetailSecretV2ResResultSSL struct {

	// REQUIRED; 证书链，包括叶子证书（服务器证书）、中间证书（中间 CA 证书）以及根证书（根 CA 证书）。证书链中的证书使用 PEM 编码格式。
	Chain []string `json:"Chain"`

	// REQUIRED; 密钥类型，默认为rsa。
	KeyType string `json:"KeyType"`
}

type DescribeClosedStreamInfoByPageQuery struct {

	// REQUIRED; 查询的起始时间，RFC3339 格式的时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeFrom string `json:"EndTimeFrom" query:"EndTimeFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式表示的时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeTo string `json:"EndTimeTo" query:"EndTimeTo"`

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的历史直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 使用流名称进行查询的方式，默认值为 strict，支持的取值即含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 排列方式，根据直播流结束时间排序，默认值为 desc，支持的取值及含义如下所示。
	// * asc：从时间最远到最近排序；
	// * desc：从时间最近到最远排序。
	Sort *string `json:"Sort,omitempty" query:"Sort"`

	// 历史直播流的来源类型，默认为空，表示查询所有来源类型，支持的取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType *string `json:"SourceType,omitempty" query:"SourceType"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 域名空间，即直播流地址的域名所属的域名空间，默认为空，表示查询所有域名空间下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的历史直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeClosedStreamInfoByPageRes struct {

	// REQUIRED
	ResponseMetadata DescribeClosedStreamInfoByPageResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeClosedStreamInfoByPageResResult          `json:"Result,omitempty"`
}

type DescribeClosedStreamInfoByPageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                  `json:"Version"`
	Error   *DescribeClosedStreamInfoByPageResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeClosedStreamInfoByPageResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeClosedStreamInfoByPageResResult struct {

	// REQUIRED; 查询结果中历史流的数量。
	RoughCount int32 `json:"RoughCount"`

	// 历史直播流信息列表。
	StreamInfoList []*DescribeClosedStreamInfoByPageResResultStreamInfoListItem `json:"StreamInfoList,omitempty"`
}

type DescribeClosedStreamInfoByPageResResultStreamInfoListItem struct {

	// REQUIRED; 历史直播流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 历史直播流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 直播流的结束时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 历史直播流的来源类型，取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType string `json:"SourceType"`

	// REQUIRED; 直播流的开始时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 历史直播流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 历史直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeDomainBody struct {

	// REQUIRED; 待查询域名信息的域名列表。
	DomainList []string `json:"DomainList"`
}

type DescribeDomainRes struct {

	// REQUIRED
	ResponseMetadata DescribeDomainResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeDomainResResult          `json:"Result,omitempty"`
}

type DescribeDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *DescribeDomainResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeDomainResResult struct {

	// 域名详细信息列表。
	DomainList []*DescribeDomainResResultDomainListItem `json:"DomainList,omitempty"`
}

type DescribeDomainResResultDomainListItem struct {

	// REQUIRED; CNAME 信息。
	CNAME string `json:"CNAME"`

	// REQUIRED; 绑定的 HTTPS 证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的 HTTPS 证书的证书链 ID 信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 域名添加时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; 是否开启 HTTP/2 协议。取值如下：
	// * false: 关闭 HTTP/2 协议。
	// * true: 开启 HTTP/2 协议。
	HTTP2 bool `json:"HTTP2"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	// * 1：备案正常，未过期；
	// * 2：查存不到备案信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 当前域名所属的域名空间下的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 域名加速区域，包含以下类型。
	// * cn：中国内地；
	// * cn-global：全球加速；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态。状态说明如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时域名加速不生效；
	// * 3：删除；
	// * 4：审核被驳回，审核不通过，需要重新创建并审核；
	// * 5：欠费关停；
	// * 6：域名未备案被封禁。
	Status int32 `json:"Status"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata DescribeEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeEncryptDRMResResult `json:"Result"`
}

type DescribeEncryptDRMResResponseMetadata struct {

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

type DescribeEncryptDRMResResult struct {

	// REQUIRED; DRM 加密配置列表。
	DRMItem DescribeEncryptDRMResResultDRMItem `json:"DRMItem"`
}

// DescribeEncryptDRMResResultDRMItem - DRM 加密配置列表。
type DescribeEncryptDRMResResultDRMItem struct {

	// REQUIRED; DRM 证书管理平台 API 访问密钥。
	APIKey string `json:"APIKey"`

	// REQUIRED; 申请 FairPlay 证书过程中 Apple 返回的 ASk（Application Secret Key）字符串。
	ApplicationSecretKey string `json:"ApplicationSecretKey"`

	// REQUIRED; FairPlay 证书文件的名称。
	CertificateFileName string `json:"CertificateFileName"`

	// REQUIRED; 自定义的 FairPlay 证书名称。
	CertificateName string `json:"CertificateName"`

	// REQUIRED; 申请 FairPlay 证书时创建的私钥文件密钥。
	PrivateKey string `json:"PrivateKey"`

	// REQUIRED; 申请 FairPlay 证书时创建的私钥文件名称。
	PrivateKeyFileName string `json:"PrivateKeyFileName"`
}

type DescribeEncryptHLSRes struct {

	// REQUIRED
	ResponseMetadata DescribeEncryptHLSResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeEncryptHLSResResult `json:"Result,omitempty"`
}

type DescribeEncryptHLSResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeEncryptHLSResResult - 视请求的接口而定
type DescribeEncryptHLSResResult struct {

	// REQUIRED; 视频直播服务端生成密钥的更新周期，单位为秒，取值范围为 [60,604800]。
	CycleTime float32 `json:"CycleTime"`

	// REQUIRED; 客户自建密钥管理服务后，客户端向密钥管理服务请求获取密钥的地址。
	URL string `json:"URL"`
}

type DescribeForbiddenStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询当前域名空间下所有的禁推流。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询所有当前域名空间下的禁推直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的禁推直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 指定是否模糊匹配流名称。缺省情况为精准匹配，支持的取值及含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 排列方式，根据推流结束时间排序，默认值为 desc，支持的取值及含义如下所示。
	// * asc：从时间最远到最近排序；
	// * desc：从时间最近到最远排序。
	Sort *string `json:"Sort,omitempty" query:"Sort"`

	// 流名称，取值与禁推直播流时设置的流名称相同，默认为空，表示查询所有流名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 域名空间，取值与禁推直播流时设置的域名空间相同，默认为空，表示查询所有域名空间下的禁推流。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的禁推流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeForbiddenStreamInfoByPageRes struct {

	// REQUIRED
	ResponseMetadata DescribeForbiddenStreamInfoByPageResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeForbiddenStreamInfoByPageResResult          `json:"Result,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                     `json:"Version"`
	Error   *DescribeForbiddenStreamInfoByPageResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResult struct {

	// REQUIRED; 查询结果中禁推流数量。
	RoughCount int32 `json:"RoughCount"`

	// 禁推流的信息列表。
	StreamInfoList []*DescribeForbiddenStreamInfoByPageResResultStreamInfoListItem `json:"StreamInfoList,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResultStreamInfoListItem struct {

	// REQUIRED; 禁推流的应用名称。
	App string `json:"App"`

	// REQUIRED; 禁推流被禁推的开始时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 禁推流的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 禁推流结束禁推的时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前流的禁推配置是否启用。
	// * true：启用；
	// * false：禁用。
	Status bool `json:"Status"`

	// REQUIRED; 禁推流的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 禁推流的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeHTTPHeaderConfigBody struct {

	// REQUIRED; 0: response 1: request
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 拉流域名。默认为空，表示查询 Vhost 下的全部拉流域名的 HTTP Header 配置。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看待查询的拉流域名。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result DescribeHTTPHeaderConfigResResult `json:"Result"`
}

type DescribeHTTPHeaderConfigResResponseMetadata struct {

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

// DescribeHTTPHeaderConfigResResult - 视请求的接口而定
type DescribeHTTPHeaderConfigResResult struct {

	// REQUIRED; HTTP Header 配置信息。
	HeaderConfigList []DescribeHTTPHeaderConfigResResultHeaderConfigListItem `json:"HeaderConfigList"`
}

type DescribeHTTPHeaderConfigResResultHeaderConfigListItem struct {

	// REQUIRED; 是否保留原 Header 配置，取值及含义如下所示。
	// * 0：保留；
	// * 1：不保留。
	BlockOriginal int32 `json:"BlockOriginal"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 配置是否启用，取值及含义如下所示。
	// * true：启用；
	// * false：禁用。
	Enable bool `json:"Enable"`

	// REQUIRED; 域名的 HTTP Header 具体字段配置。
	HeaderDetailList []DescribeHTTPHeaderConfigResResultHeaderConfigListPropertiesItemsItem `json:"HeaderDetailList"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeHTTPHeaderConfigResResultHeaderConfigListPropertiesItemsItem struct {

	// REQUIRED; Header 配置中字段 Value 值的类型，取值及含义如下所示。
	// * 0：常量；
	// * 1：变量。
	HeaderFieldType int32 `json:"HeaderFieldType"`

	// REQUIRED; Header 配置中字段的 Key 值。
	HeaderKey string `json:"HeaderKey"`

	// REQUIRED; Header 配置中字段的 Value 值。
	HeaderValue string `json:"HeaderValue"`
}

type DescribeHighLightTaskByAccountIDBody struct {

	// REQUIRED; 创建高光剪辑任务时返回的taskid
	TaskID string `json:"TaskID"`

	// 是否返回高光剪辑生产结果
	NeedResult *bool `json:"NeedResult,omitempty"`
}

type DescribeHighLightTaskByAccountIDRes struct {

	// REQUIRED
	ResponseMetadata DescribeHighLightTaskByAccountIDResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeHighLightTaskByAccountIDResResult `json:"Result,omitempty"`
}

type DescribeHighLightTaskByAccountIDResResponseMetadata struct {

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestID"`

	// 请求的接口名，属于请求的公共参数。
	Action *string `json:"Action,omitempty"`

	// 请求的Region，例如：cn-north-1
	Region *string `json:"Region,omitempty"`

	// 请求的服务，属于请求的公共参数。
	Service *string `json:"Service,omitempty"`

	// 请求的版本号，属于请求的公共参数。
	Version *string `json:"Version,omitempty"`
}

// DescribeHighLightTaskByAccountIDResResult - 视请求的接口而定
type DescribeHighLightTaskByAccountIDResResult struct {

	// 高光剪辑任务的详情数据，包含任务所属账号、任务参数配置、任务状态、任务唯一标识符以及生产结果列表等信息。
	Data *DescribeHighLightTaskByAccountIDResResultData `json:"Data,omitempty"`
}

// DescribeHighLightTaskByAccountIDResResultData - 高光剪辑任务的详情数据，包含任务所属账号、任务参数配置、任务状态、任务唯一标识符以及生产结果列表等信息。
type DescribeHighLightTaskByAccountIDResResultData struct {

	// 高光剪辑任务所属的账号
	AccountID *string `json:"AccountID,omitempty"`

	// 高光剪辑任务的生产结果列表，包含多个生产回调结果。每个结果包含以下信息：
	// * code：生产回调的状态码；
	// * status：生产状态码；
	// * message：生产描述信息；
	// * exception：异常信息；
	// * TaskID：高光剪辑任务的唯一标识符；
	// * Output：高光剪辑生产结果的详细信息；
	// * ExtraData：自定义的其他信息。
	ProducedResults []*DescribeHighLightTaskByAccountIDResResultDataProducedResultsItem `json:"ProducedResults,omitempty"`

	// 创建高光剪辑任务的body参数，为json字符串
	Rule *string `json:"Rule,omitempty"`

	// 高光剪辑任务状态，init：任务初始化状态，pending：任务等待启动阶段，prepare：任务启动状态，running：任务生产状态，prestop：任务准备结束状态，done：任务完成状态，error：任务报错状态
	Status *string `json:"Status,omitempty"`

	// 创建高光剪辑任务时返回的taskid
	TaskID *string `json:"TaskID,omitempty"`
}

type DescribeHighLightTaskByAccountIDResResultDataProducedResultsItem struct {

	// REQUIRED; 高光剪辑任务生产回调的状态码，用于标识生产过程的当前状态。
	// * 100：表示生产成功，任务将持续生产；
	// * 200：表示生产完成，任务退出；
	// * 500：表示当前生产回调失败，但任务仍会持续生产；
	// * 1000：表示高光剪辑生产失败，任务停止。
	Code int64 `json:"code"`

	// REQUIRED; 高光剪辑任务生产过程中的异常信息，用于描述生产失败的具体原因。
	Exception string `json:"exception"`

	// REQUIRED; 高光剪辑任务创建时的自定义其他信息，用于传递额外的任务相关数据。
	ExtraData string `json:"ExtraData"`

	// REQUIRED; 高光剪辑任务生产过程中的描述信息。
	Message string `json:"message"`

	// REQUIRED
	Output DescribeHighLightTaskByAccountIDResResultDataProducedResultsItemOutput `json:"Output"`

	// REQUIRED; 高光剪辑任务生产状态码，可能的取值如下所示。
	// * Continue：表示生产成功，持续生产；
	// * OK：表示生产完成，任务退出；
	// * Error：表示当前生产回调失败，并且任务还会持续生产；
	// * Fatal：表示高光剪辑生产失败，任务停止。
	Status string `json:"status"`

	// REQUIRED; 高光剪辑任务的唯一标识符，用于标识特定的高光剪辑生产任务。
	TaskID string `json:"TaskID"`
}

type DescribeHighLightTaskByAccountIDResResultDataProducedResultsItemOutput struct {

	// REQUIRED
	HLClipsOutput Components1Pvao98SchemasDescribehighlighttaskbyaccountidresPropertiesResultPropertiesDataPropertiesProducedresultsItemsPropertiesOutputPropertiesHlclipsoutput `json:"HLClipsOutput"`

	// REQUIRED
	HLMixOutput DescribeHighLightTaskByAccountIDResResultDataProducedResultsProperties `json:"HLMixOutput"`
}

type DescribeHighLightTaskByAccountIDResResultDataProducedResultsProperties struct {

	// REQUIRED; 高光混剪结果信息，包含从多个高光片段中混剪生成的最终视频的详细信息。
	HLMixCuts []DescribeHighLightTaskByAccountIDResResultDataProducedResultsPropertiesItemsItem `json:"HLMixCuts"`
}

type DescribeHighLightTaskByAccountIDResResultDataProducedResultsPropertiesItemsHLClipsItem struct {

	// REQUIRED; 高光片段类型标识（即 Label）的置信度，用于表示该类型标识的可靠程度。仅适用于体育足球场景，取值范围为 [0,1]，数值越大表示置信度越高。
	Confidence float32 `json:"Confidence"`

	// REQUIRED; 高光片段内容描述。仅适用于体育足球。
	Description string `json:"Description"`

	// REQUIRED; 高光剪辑片段时长，单位为毫秒。
	Duration int64 `json:"Duration"`

	// REQUIRED; 高光剪辑片段相对高光片段或者高光混剪的截止时间位置，单位为毫秒。
	HLEnd int64 `json:"HLEnd"`

	// REQUIRED; 高光剪辑片段相对高光片段或者高光混剪的起始时间位置，单位为毫秒。
	HLStart int64 `json:"HLStart"`

	// REQUIRED; 高光剪辑片段所属原始视频的序号，用于标识该片段在原始视频中的顺序位置。
	Index int32 `json:"Index"`

	// REQUIRED; 高光片段的类型标识。当前仅用于体育足球场景，可能的取值如下所示。
	// * 0：无标签；
	// * 1：背景；
	// * 2：进球；
	// * 3：角球；
	// * 4：点球；
	// * 5：黄牌；
	// * 6：红牌；
	// * 7：犯规；
	// * 8：射门；
	// * 9：任意球。
	Label int32 `json:"Label"`

	// REQUIRED; 高光剪辑片段相对所属原始视频的截止时间位置，单位为毫秒。
	SEnd int64 `json:"SEnd"`

	// REQUIRED; 高光剪辑片段相对所属原始视频的起始时间位置，单位为毫秒。
	SStart int64 `json:"SStart"`

	// REQUIRED; 高光剪辑片段的高光分数，用于评估该片段在原始视频中的精彩程度。分数越高表示该片段越精彩，分数区间范围为 [0, 10]。
	Score float32 `json:"Score"`

	// REQUIRED; 高光剪辑片段的卖点信息结果。
	SellPointsRes []DescribeHighLightTaskByAccountIDResResultDataProducedResultsPropertiesItemsHLClipsPropertiesItemsItem `json:"SellPointsRes"`

	// REQUIRED; 高光剪辑片段所属的原始视频 URL 地址，用于标识该片段是从哪个原始视频中提取的。
	SourceURL string `json:"SourceUrl"`

	// 高光剪辑片段上传至视频点播（VOD）服务后的视频唯一标识符 vid（Video ID），用于唯一标识该视频文件。
	ClipURL *string `json:"ClipUrl,omitempty"`
}

type DescribeHighLightTaskByAccountIDResResultDataProducedResultsPropertiesItemsHLClipsPropertiesItemsItem struct {

	// REQUIRED; 卖点相对于混剪结果的结束时间，单位为毫秒。
	EndTime int64 `json:"EndTime"`

	// REQUIRED; 卖点添加过程中的错误信息，用于描述添加卖点失败的具体原因。
	Message string `json:"Message"`

	// REQUIRED; 视频卖点信息，用于描述高光剪辑片段中的关键亮点内容。
	SellPoints string `json:"SellPoints"`

	// REQUIRED; 卖点相对于混剪结果的开始时间，单位为毫秒。
	StartTime int64 `json:"StartTime"`

	// REQUIRED; 卖点添加的状态，取值如下所示。
	// * ok：卖点添加成功；
	// * fail：卖点添加失败。
	Status string `json:"Status"`

	// 使用的卖点效果模版，当前仅支持 default 值。
	EffectType *string `json:"EffectType,omitempty"`
}

type DescribeHighLightTaskByAccountIDResResultDataProducedResultsPropertiesItemsItem struct {

	// REQUIRED; 高光剪辑混剪结果视频的时长，单位为毫秒。
	Duration int64 `json:"Duration"`

	// REQUIRED; 高光剪辑混剪结果视频对应的片段详情信息，包含从原始视频中提取的高光片段信息。
	HLClips []DescribeHighLightTaskByAccountIDResResultDataProducedResultsPropertiesItemsHLClipsItem `json:"HLClips"`

	// REQUIRED; 高光剪辑混剪结果上传至视频点播（VOD）服务后的视频唯一标识符 vid（Video ID），用于唯一标识该视频文件。
	HighlightURL string `json:"HighlightUrl"`

	// REQUIRED; 高光剪辑混剪结果序号，用于标识混剪结果在结果列表中的顺序位置。
	Index int32 `json:"Index"`
}

type DescribeIPAccessRuleBody struct {

	// REQUIRED; 推流域名或拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要查询
	// IP 访问限制的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// IP 访问限制的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeIPAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DescribeIPAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeIPAccessRuleResResult `json:"Result"`
}

type DescribeIPAccessRuleResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Error string `json:"Error"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeIPAccessRuleResResult struct {

	// REQUIRED; IP 访问限制规则列表。
	AccessRuleLists []DescribeIPAccessRuleResResultAccessRuleListsItem `json:"AccessRuleLists"`
}

type DescribeIPAccessRuleResResultAccessRuleListsItem struct {

	// 推/拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// IP 访问限制规则。
	IPAccessRule *DescribeIPAccessRuleResResultAccessRuleListsItemIPAccessRule `json:"IPAccessRule,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

// DescribeIPAccessRuleResResultAccessRuleListsItemIPAccessRule - IP 访问限制规则。
type DescribeIPAccessRuleResResultAccessRuleListsItemIPAccessRule struct {

	// REQUIRED; 对于黑名单中或非白名单中的 IP 地址，系统返回的 HTTP 状态码。
	DenyHTTPStatusCode int32 `json:"DenyHTTPStatusCode"`

	// REQUIRED; 是否开启当前限制。取值如下：
	// * true: 开启。
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 校验是否包含 TS 文件。取值如下：
	// * true: 包含。
	// * false: 不包含。
	// 使用 HLS 协议拉流时，直播服务器会将完整的直播流切割成小的 TS（Transport Stream）文件。观众观看直播时，播放器会不断请求 TS 文件并拼接成连续的画面。参数取值为true时，直播服务器会严格验证请求 TS 文件的 IP
	// 地址，防止黑名单中或不在白名单中的 IP 地址访问 TS 文件。 :::tip 该配置仅对拉流域名生效。 :::
	EnableTS bool `json:"EnableTS"`

	// REQUIRED; 名单中的 IP 信息。 例如，Type取值为deny、Domain为推流域名、该参数取值为 ["192.168.1.100","192.168.1.0/24","2001:db8:85a3::8a2e:370:7334","2001:db8::/32"]
	// 时，则表示不允许以下 IP 地址推流：
	// * IP 地址192.168.1.100
	// * IP 地址2001:db8:85a3::8a2e:370:7334
	// * 192.168.1.0- 192.168.1.255范围内的所有 IP 地址
	// * 2001:db8:0000:0000:0000:0000:0000:0000-2001:db8:ffff:ffff:ffff:ffff:ffff:ffff范围内的所有 IP 地址
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型。取值如下：
	// * allow: 白名单。 * 如果 Domain 为推流域名，则只有符合规则的 IP 地址才可以推流。
	// * 如果 Domain 为拉流域名，则只有符合规则的 IP 地址才可以拉流。
	//
	//
	// * deny: 黑名单。 * 如果 Domain 为推流域名，则符合规则的 IP 地址无法推流。
	// * 如果 Domain 为拉流域名，则符合规则的 IP 地址无法拉流。
	Type string `json:"Type"`
}

type DescribeIPInfoBody struct {

	// REQUIRED; 待查询的 IP 地址列表。支持 IPv4 和 IPv6 地址，一次最多查询 50 个 IP 地址。
	IPs []string `json:"Ips"`
}

type DescribeIPInfoRes struct {

	// REQUIRED
	ResponseMetadata DescribeIPInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeIPInfoResResult          `json:"Result,omitempty"`
}

type DescribeIPInfoResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                  `json:"Version"`
	Error   *DescribeIPInfoResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeIPInfoResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeIPInfoResResult struct {

	// REQUIRED; IP 详情列表。
	List []DescribeIPInfoResResultListItem `json:"List"`
}

type DescribeIPInfoResResultListItem struct {

	// REQUIRED; CDN 节点所属城市信息。非归属火山引擎的 CDN 节点 IP 时，返回“-”。
	City string `json:"City"`

	// REQUIRED; IP 地址。
	IP string `json:"Ip"`

	// REQUIRED; CDN 节点所属网络运营商。非归属火山引擎 CDN 节点的 IP 时，返回”-“。您可以通过DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974]接口查看运营商标识符对应的运营商名称。
	Isp string `json:"Isp"`

	// REQUIRED; 当前 IP 地址是否是归属于火山引擎的 CDN 节点 IP。
	// * true：属于；
	// * false：不属于。
	LiveCdnIP bool `json:"LiveCdnIp"`

	// REQUIRED; CDN 节点所属国家或地区信息。非归属火山引擎的 CDN 节点 IP 时，返回“-”。
	Location string `json:"Location"`

	// REQUIRED; CDN 节点所属省份信息。非归属火山引擎 CDN 节点的 IP 时，返回“-”。
	Province string `json:"Province"`
}

type DescribeLicenseDRMQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段相同，由大写小字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; DRM 加密的类型，取值及含义如下所示。
	// * fp：FairPlay 加密；
	// * wv：Widevine 加密；
	// * pr：PlayReady 加密。
	DRMType string `json:"DRMType" query:"DRMType"`

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	Domain string `json:"Domain" query:"Domain"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	StreamName string `json:"StreamName" query:"StreamName"`

	// REQUIRED; 拉取加密流时使用的拉流域名所在的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名所属的域名空间。
	Vhost string `json:"Vhost" query:"Vhost"`
}

type DescribeLicenseDRMRes struct {

	// REQUIRED
	ResponseMetadata DescribeLicenseDRMResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DescribeLicenseDRMResResponseMetadata struct {

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

type DescribeLiveAuditDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 86400 秒进行数据聚合。
	// :::tip 时间粒度为 86400 秒时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain
	// 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的截图审核张数用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLiveAuditDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveAuditDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveAuditDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveAuditDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                         `json:"Version"`
	Error   *DescribeLiveAuditDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveAuditDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveAuditDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	AuditDataList []DescribeLiveAuditDataResResultAuditDataListItem `json:"AuditDataList"`

	// REQUIRED; 按维度拆分后的数据。
	AuditDetailDataList []DescribeLiveAuditDataResResultAuditDetailDataListItem `json:"AuditDetailDataList"`

	// REQUIRED; 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []string `json:"DetailField"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图审核总张数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveAuditDataResResultAuditDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的截图审核总张数。
	Count int32 `json:"Count"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveAuditDataResResultAuditDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	AuditDataList []DescribeLiveAuditDataResResultAuditDetailDataListPropertiesItemsItem `json:"AuditDataList"`

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的截图审核总张数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveAuditDataResResultAuditDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 截图审核总张数
	Count int32 `json:"Count"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain
	// 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的带宽用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBandwidthDataResResult `json:"Result"`
}

type DescribeLiveBandwidthDataResResponseMetadata struct {

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

type DescribeLiveBandwidthDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
	BandwidthDetailDataList []*DescribeLiveBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的上行峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLiveBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行带宽，单位为 Mbps
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataResResultUserRegionListItem struct {

	// 大区
	Area *string `json:"Area,omitempty"`

	// 国家
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchPushStreamAvgMetricsBody struct {

	// REQUIRED; 推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	App *string `json:"App,omitempty"`

	// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamAvgMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchPushStreamAvgMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchPushStreamAvgMetricsResResult `json:"Result"`
}

type DescribeLiveBatchPushStreamAvgMetricsResResponseMetadata struct {

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

type DescribeLiveBatchPushStreamAvgMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流的信息，包含域名、应用名称、流名称和监控数据。
	StreamMetricList []DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率平均值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的平均值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率平均值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率平均值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率平均值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的平均值，即所有 AudioPts 与 VideoPts 差值的平均值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间， RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的平均值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveBatchPushStreamMetricsBody struct {

	// REQUIRED; 推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 请控制查询的数据量，如果查询速度较慢请缩短查询时间范围
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 指标聚合算法，支持max:峰值聚合，avg：平均值，默认max
	AggType *string `json:"AggType,omitempty"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	App *string `json:"App,omitempty"`

	// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchPushStreamMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchPushStreamMetricsResResult `json:"Result"`
}

type DescribeLiveBatchPushStreamMetricsResResponseMetadata struct {

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

type DescribeLiveBatchPushStreamMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 直推流的信息，包含域名、应用名称、流名称和监控数据。
	StreamMetricList []DescribeLiveBatchPushStreamMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 数据聚合时间粒度内，动态指标的聚合算法，取值及含义如下所示。
	// * max：（默认值）计算聚合时间粒度内的最大值；
	// * avg：计算聚合时间粒度内的平均值。
	AggType *string `json:"AggType,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchPushStreamMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 标记一路推流的唯一id
	SessionID string `json:"SessionID"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchPushStreamMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 音频编码格式。 Eg. ACC
	ACodec string `json:"ACodec"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的音频码率，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的相邻音频帧显示时间戳差值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的音频帧率（每秒传输的音频数据包个数）。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的视频码率，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 客户端ip
	ClientIP string `json:"ClientIp"`

	// REQUIRED; 收到首帧的时间，，单位毫秒
	FirstFrameTime int32 `json:"FirstFrameTime"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的视频帧率，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的音视频帧显示时间戳差值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 分辨率
	Resolution string `json:"Resolution"`

	// REQUIRED; 服务端ip
	ServerIP string `json:"ServerIp"`

	// REQUIRED; 推流开始时间，单位毫秒
	StreamBeginTime int64 `json:"StreamBeginTime"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 视频编码格式。 Eg. H264
	VCodec string `json:"VCodec"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的相邻视频帧显示时间戳差值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveBatchSourceStreamAvgMetricsBody struct {

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看回源流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定回源流。 :::
	App *string `json:"App,omitempty"`

	// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定回源流。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamAvgMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchSourceStreamAvgMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchSourceStreamAvgMetricsResResult `json:"Result"`
}

type DescribeLiveBatchSourceStreamAvgMetricsResResponseMetadata struct {

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

type DescribeLiveBatchSourceStreamAvgMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流的信息，包含域名、应用名称、流名称和监控数据。
	StreamMetricList []DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率平均值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的平均值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率平均值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率平均值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率平均值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的平均值，即所有 AudioPts 与 VideoPts 差值的平均值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的平均值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveBatchSourceStreamMetricsBody struct {

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看回源流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 聚合类型。
	AggType *string `json:"AggType,omitempty"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定回源流。 :::
	App *string `json:"App,omitempty"`

	// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定回源流。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchSourceStreamMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchSourceStreamMetricsResResult `json:"Result"`
}

type DescribeLiveBatchSourceStreamMetricsResResponseMetadata struct {

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

type DescribeLiveBatchSourceStreamMetricsResResult struct {

	// REQUIRED; 聚合类型。
	AggType string `json:"AggType"`

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 回源流信息，包含域名、应用名称、流名称和监控数据。
	StreamMetricList []DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 当前数据聚合时间粒度内，按聚合算法得出的音频码率，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，按聚合算法得出的相邻音频帧显示时间戳差值的最大值的平均值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，按聚合算法得出的音频帧率，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内，按聚合算法得出的视频码率，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，按聚合算法得出的视频帧率，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，按聚合算法得出的音视频帧显示时间戳差值的最大值或平均值，即所有 AudioPts 与 VideoPts 差值的最大值或平均值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，按聚合算法得出的相邻视频帧显示时间戳差值的最大值或平均值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveBatchStreamSessionDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 24 小时，查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 拉流域名列表，默认为空，表示查询所有域名的请求数和在线人数。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 在线人数类型，支持如下枚举值： Online：瞬时session链接数； Viewer（默认值）：一分钟session链接总数。
	OnlineUserType *string `json:"OnlineUserType,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 1000，取值范围为 [100,1000]。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip
	// * 如果查询推拉流协议为 QUIC，不能同时查询其他协议。
	// * 缺省情况下，查询的总流量数据为实际产生的上下行流量。
	// * 如果传入单个协议进行查询，并对各协议的流量求和，结果将大于实际总流量。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。
	RegionList []*DescribeLiveBatchStreamSessionDataBodyRegionListItem `json:"RegionList,omitempty"`
}

type DescribeLiveBatchStreamSessionDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchStreamSessionDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchStreamSessionDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchStreamSessionDataResResult `json:"Result"`
}

type DescribeLiveBatchStreamSessionDataResResponseMetadata struct {

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

type DescribeLiveBatchStreamSessionDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamSessionDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询时间范围内的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内所有流的请求数和在线人数峰值信息。
	StreamInfoList []DescribeLiveBatchStreamSessionDataResResultStreamInfoListItem `json:"StreamInfoList"`

	// REQUIRED; 查询时间范围内的请求数总数。
	TotalRequest int32 `json:"TotalRequest"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveBatchStreamSessionDataResResultRegionListItem `json:"RegionList,omitempty"`
}

// DescribeLiveBatchStreamSessionDataResResultPagination - 数据分页的信息。
type DescribeLiveBatchStreamSessionDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveBatchStreamSessionDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchStreamSessionDataResResultStreamInfoListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 当前流的在线人数峰值
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 当前流的请求数。
	TotalRequest int32 `json:"TotalRequest"`
}

type DescribeLiveBatchStreamTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 24 小时，查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 域名列表，默认为空，表示查询所有域名下的流量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 1000，取值范围为 [100,1000]。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip
	// * 如果查询推拉流协议为 QUIC，不能同时查询其他协议。
	// * 缺省情况下，查询的总流量数据为实际产生的上下行流量。
	// * 如果传入单个协议进行查询，并对各协议的流量求和，结果将大于实际总流量。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveBatchStreamTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveBatchStreamTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBatchStreamTrafficDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchStreamTrafficDataBodyUserRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchStreamTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchStreamTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchStreamTrafficDataResResult `json:"Result"`
}

type DescribeLiveBatchStreamTrafficDataResResponseMetadata struct {

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

type DescribeLiveBatchStreamTrafficDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamTrafficDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流维度的流量用量信息详情。
	StreamInfoList []DescribeLiveBatchStreamTrafficDataResResultStreamInfoListItem `json:"StreamInfoList"`

	// REQUIRED; 当前查询条件下，所有流的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 当前查询条件下，所有流的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveBatchStreamTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveBatchStreamTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

// DescribeLiveBatchStreamTrafficDataResResultPagination - 数据分页的信息。
type DescribeLiveBatchStreamTrafficDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveBatchStreamTrafficDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchStreamTrafficDataResResultStreamInfoListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 域名。
	Domain string `json:"Domain"`

	// REQUIRED; 当前流的下行流量，单位为 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 当前流的上行流量，单位为 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveBatchStreamTrafficDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchStreamTranscodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 域名列表，默认为空表示全部域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取直播流使用的域名信息。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 1000，取值范围为 [100,1000]。
	PageSize *int32 `json:"PageSize,omitempty"`
}

type DescribeLiveBatchStreamTranscodeDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchStreamTranscodeDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchStreamTranscodeDataResResult `json:"Result"`
}

type DescribeLiveBatchStreamTranscodeDataResResponseMetadata struct {

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

type DescribeLiveBatchStreamTranscodeDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamTranscodeDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流维度的转码用量信息详情。
	StreamInfoList []DescribeLiveBatchStreamTranscodeDataResResultStreamInfoListItem `json:"StreamInfoList"`

	// REQUIRED; 当前查询条件下，所有流的转码总时长，单位为分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`
}

// DescribeLiveBatchStreamTranscodeDataResResultPagination - 数据分页的信息。
type DescribeLiveBatchStreamTranscodeDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页展示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveBatchStreamTranscodeDataResResultStreamInfoListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 当前流的转码码率，单位为 kbps。
	Coderate int32 `json:"Coderate"`

	// REQUIRED; 域名。
	Domain string `json:"Domain"`

	// REQUIRED; 当前流在查询时间内的转码总时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	VCodec string `json:"VCodec"`
}

type DescribeLiveCallbackDataBody struct {

	// REQUIRED; 结束时间，单次查询31天，历史366天
	EndTime string `json:"EndTime"`

	// REQUIRED; 开始时间
	StartTime string `json:"StartTime"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 回调事件类型 推流开始-push_start 推流结束-push_end 截图回调-snapshot_event 录制回调-record_event
	CallbackEventType []*string `json:"CallbackEventType,omitempty"`

	// 回调状态 成功-success 失败-fail
	CallbackStatus []*string `json:"CallbackStatus,omitempty"`

	// 需查询的域名列表，缺省情况下表示当前账号下的所有域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分页页码，默认是1，取值范围[1，10000]
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页的大小，默认100，取值范围[1, 1000]
	PageSize *int32 `json:"PageSize,omitempty"`

	// 流名称，用于精确定位某一路直播流。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveCallbackDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveCallbackDataResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLiveCallbackDataResResult `json:"Result,omitempty"`
}

type DescribeLiveCallbackDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeLiveCallbackDataResResult - 视请求的接口而定
type DescribeLiveCallbackDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 回调事件类型
	CallbackEventType []*string `json:"CallbackEventType,omitempty"`

	// 回调事件详情
	CallbackInfoList []*DescribeLiveCallbackDataResResultCallbackInfoListItem `json:"CallbackInfoList,omitempty"`

	// 回调状态
	CallbackStatus []*string `json:"CallbackStatus,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询结果的分页信息。
	Pagination *DescribeLiveCallbackDataResResultPagination `json:"Pagination,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveCallbackDataResResultCallbackInfoListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 回调地址。
	CallbackAddress string `json:"CallbackAddress"`

	// REQUIRED; 回调请求体。
	CallbackBody string `json:"CallbackBody"`

	// REQUIRED; 回调错误说明。
	CallbackErrorCode string `json:"CallbackErrorCode"`

	// REQUIRED; 错误信息，当回调失败时返回。
	CallbackErrorMessage string `json:"CallbackErrorMessage"`

	// REQUIRED; 回调事件类型。
	CallbackEventType string `json:"CallbackEventType"`

	// REQUIRED; 回调请求方式。
	CallbackMethod string `json:"CallbackMethod"`

	// REQUIRED; 回调响应体。
	CallbackResponseBody string `json:"CallbackResponseBody"`

	// REQUIRED; 回调响应码。
	CallbackResponseCode string `json:"CallbackResponseCode"`

	// REQUIRED; 回调响应头信息。
	CallbackResponseHeader string `json:"CallbackResponseHeader"`

	// REQUIRED; 回调响应时间。
	CallbackResponseTime string `json:"CallbackResponseTime"`

	// REQUIRED; 回调状态
	CallbackStatus string `json:"CallbackStatus"`

	// REQUIRED; 回调发生时间。
	CallbackTime string `json:"CallbackTime"`

	// REQUIRED; 域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

// DescribeLiveCallbackDataResResultPagination - 查询结果的分页信息。
type DescribeLiveCallbackDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveEdgeStatDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天，单次查询最大时间跨度与数据拆分维度和数据聚合时间粒度有关，详细如下。
	// * 当不进行维度拆分或只使用一个维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 31 天；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 31 天。
	//
	//
	// * 当使用两个或两个以上维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 3 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 7 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持以下取值。
	// * 60：1 分钟。
	// * 300：（默认值）5 分钟。
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用程序列表。
	AppList []*string `json:"AppList,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商；
	// * Area：CDN 节点 IP 所属大区。
	// * Country：CDN 节点 IP 所属国家。
	// * Province：CDN 节点 IP 所属省份。
	// * UserArea：客户端 IP 所属大区。
	// * UserCountry：客户端 IP 所属国家。
	// * UserProvince：客户端 IP 所属省份。
	// :::tip 中国（Country 或 UserCountry 为 CN）以外区域无 Province 字段，如果按 Province 或 UserProvince 字段拆分数据时，默认只返回 Country 为 CN 时的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示您添加到视频直播中的所有域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数RegionList和UserRegionList不支持同时传入。 :::
	RegionList []*DescribeLiveEdgeStatDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 流列表。
	StreamList []*string `json:"StreamList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数RegionList和UserRegionList不支持同时传入。 :::
	UserRegionList []*DescribeLiveEdgeStatDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveEdgeStatDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveEdgeStatDataBodyUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveEdgeStatDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveEdgeStatDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveEdgeStatDataResResult `json:"Result"`
}

type DescribeLiveEdgeStatDataResResponseMetadata struct {

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

type DescribeLiveEdgeStatDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 边缘统计数据列表。
	EdgeStatDataList []DescribeLiveEdgeStatDataResResultEdgeStatDataListItem `json:"EdgeStatDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 带宽峰值。取值范围为 [ ]，单位为，默认值为``。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 请求数。
	TotalRequest float32 `json:"TotalRequest"`

	// REQUIRED; 上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// 应用列表。
	AppList []*string `json:"AppList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商；
	// * Area：CDN 节点 IP 所属大区。
	// * Country：CDN 节点 IP 所属国家。
	// * Province：CDN 节点 IP 所属省份。
	// * UserArea：客户端 IP 所属大区。
	// * UserCountry：客户端 IP 所属国家。
	// * UserProvince：客户端 IP 所属省份。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 边缘统计详细数据列表。
	EdgeStatDetailDataList []*DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListItem `json:"EdgeStatDetailDataList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveEdgeStatDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流列表。
	StreamList []*string `json:"StreamList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveEdgeStatDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveEdgeStatDataResResultEdgeStatDataListItem struct {

	// REQUIRED; 下行带宽。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 请求参数。
	Request float32 `json:"Request"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC+8 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽。
	UpBandwidth float32 `json:"UpBandwidth"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总上行流量，单位 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListItem struct {

	// REQUIRED; 边缘统计数据列表。
	EdgeStatDataList []DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListPropertiesItemsItem `json:"EdgeStatDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的请求数。
	TotalRequest float32 `json:"TotalRequest"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// 区域。
	Area *string `json:"Area,omitempty"`

	// 国家。
	Country *string `json:"Country,omitempty"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`

	// 省份。
	Province *string `json:"Province,omitempty"`

	// 用户区域。
	UserArea *string `json:"UserArea,omitempty"`

	// 用户所在国家。
	UserCountry *string `json:"UserCountry,omitempty"`

	// 用户所在省份。
	UserProvince *string `json:"UserProvince,omitempty"`
}

type DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListPropertiesItemsItem struct {

	// REQUIRED
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 下行流量，单位 GB
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED
	Request float32 `json:"Request"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED
	UpBandwidth float32 `json:"UpBandwidth"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveEdgeStatDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveEdgeStatDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveISPDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveISPDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveISPDataResResult `json:"Result"`
}

type DescribeLiveISPDataResResponseMetadata struct {

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

type DescribeLiveISPDataResResult struct {

	// REQUIRED; 运营商信息，视频直播提供的网络运营商标识，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []DescribeLiveISPDataResResultISPListItem `json:"ISPList"`
}

type DescribeLiveISPDataResResultISPListItem struct {

	// REQUIRED; 运营商标识符。
	Code string `json:"Code"`

	// REQUIRED; 运营商名称。
	Name string `json:"Name"`
}

type DescribeLiveLogDataBody struct {

	// REQUIRED; 仅支持查询最近31天的数据
	EndTime string `json:"EndTime"`

	// REQUIRED; 仅支持查询最近31天的数据
	StartTime string `json:"StartTime"`

	// REQUIRED; 日志类型，支持的类型如下所示。
	// * pull：拉流日志；
	// * push：推流日志；
	// * source：回源日志；
	// * relay：拉流转推日志。
	Type string `json:"Type"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的日志文件信息。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。 :::tip
	// 日志类型为拉流转推日志（Type 取值为 relay）时，该参数无效。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认为 20，最大值为 1000。
	PageSize *int32 `json:"PageSize,omitempty"`
}

type DescribeLiveLogDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveLogDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveLogDataResResult `json:"Result"`
}

type DescribeLiveLogDataResResponseMetadata struct {

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

type DescribeLiveLogDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 日志文件的信息列表。
	LogInfoList []DescribeLiveLogDataResResultLogInfoListItem `json:"LogInfoList"`

	// REQUIRED; 数据分页信息。
	Pagination DescribeLiveLogDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 日志类型，类型说明如下所示。
	// * pull：拉流日志
	// * push：推流日志
	// * source：回源日志
	// * relay：拉流转推日志
	Type string `json:"Type"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLiveLogDataResResultLogInfoListItem struct {

	// REQUIRED; 日志文件对应的小时区间，RFC3339 格式的时间戳，精度为秒。
	DateTime string `json:"DateTime"`

	// REQUIRED; 日志文件下载链接。
	DownloadURL string `json:"DownloadUrl"`

	// REQUIRED; 日志文件名称，日志文件命名规则如下。
	// * 与域名相关时：加速域名年月日时间开始时间结束文件拆分序号。例如，push.example.com_2023_08_11_000000_010000_0.gz；
	// * 与域名无关时：年月日时间开始时间结束_文件拆分序号。例如，2023_08_11_000000_010000_0.gz； :::tip 如果某个小时内，当前事件产生的日志大于 150 万条，则会生成为多个日志文件，用文件名最后的序号标注日志文件顺序，例如，2023_08_11_000000_010000_0.gz、2023_08_11_000000_010000_1.gz。
	// :::
	LogName string `json:"LogName"`

	// REQUIRED; 日志文件大小，单位为 byte。
	LogSize int32 `json:"LogSize"`

	// 域名。 :::tip 查询拉流转推日志（Type 取值为 relay）时，该字段为空。 :::
	Domain *string `json:"Domain,omitempty"`
}

// DescribeLiveLogDataResResultPagination - 数据分页信息。
type DescribeLiveLogDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveMetricBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的带宽监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的带宽监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveMetricBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的带宽监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveMetricBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveMetricBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveMetricBandwidthDataResResult `json:"Result"`
}

type DescribeLiveMetricBandwidthDataResResponseMetadata struct {

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

type DescribeLiveMetricBandwidthDataResResult struct {

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 按维度拆分后的数据。
	BandwidthDetailDataList []*DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveMetricBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveMetricBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的上行峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行带宽，单位为 Mbps
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveMetricBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的身份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的流量监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的流量监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveMetricTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 流名称，取值与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的流量监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveMetricTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricTrafficDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveMetricTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveMetricTrafficDataResResult `json:"Result"`
}

type DescribeLiveMetricTrafficDataResResponseMetadata struct {

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

type DescribeLiveMetricTrafficDataResResult struct {

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 查询时间范围内的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveMetricTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveMetricTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按此维度进行拆分的数据。 :::
	TrafficDetailDataList []*DescribeLiveMetricTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveMetricTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总上行流量，单位 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveMetricTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	TrafficDataList []DescribeLiveMetricTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行流量，单位 GB
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveMetricTrafficDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 300 秒进行数据拆分。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的 95 峰值带宽用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveP95PeakBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveP95PeakBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveP95PeakBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveP95PeakBandwidthDataResResult `json:"Result"`
}

type DescribeLiveP95PeakBandwidthDataResResponseMetadata struct {

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

type DescribeLiveP95PeakBandwidthDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 时间范围内的上下行 95 峰值带宽总和。 :::tip 如果请求时，Regionlist中传入多个 region，则返回这些 region 的上下行带宽 95 峰值总和。 :::
	P95PeakBandwidth float32 `json:"P95PeakBandwidth"`

	// REQUIRED; 95 峰值带宽的时间戳，RFC3339 格式的时间戳，精度为秒。
	P95PeakTimestamp string `json:"P95PeakTimestamp"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表。
	RegionList []*DescribeLiveP95PeakBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表。
	UserRegionList []*DescribeLiveP95PeakBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePadPresetDetailBody struct {

	// 应用名称
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeLivePadPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePadPresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLivePadPresetDetailResResult `json:"Result,omitempty"`
}

type DescribeLivePadPresetDetailResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeLivePadPresetDetailResResult - 视请求的接口而定
type DescribeLivePadPresetDetailResResult struct {

	// REQUIRED; 模板列表
	PresetList []DescribeLivePadPresetDetailResResultPresetListItem `json:"PresetList"`
}

type DescribeLivePadPresetDetailResResultPresetListItem struct {

	// REQUIRED; 应用名称
	App string `json:"App"`

	// REQUIRED; 模板详情。
	PresetDetail DescribeLivePadPresetDetailResResultPresetListItemPresetDetail `json:"PresetDetail"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

// DescribeLivePadPresetDetailResResultPresetListItemPresetDetail - 模板详情。
type DescribeLivePadPresetDetailResResultPresetListItemPresetDetail struct {

	// REQUIRED; 描述
	Description string `json:"Description"`

	// REQUIRED; 最大持续时间。
	MaxDuration int64 `json:"MaxDuration"`

	// REQUIRED; 垫片类型，1: 图片、2: 视频、3: 源流最后一帧
	PadType int32 `json:"PadType"`

	// REQUIRED; 预设名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; 垫片内容URL。
	URL string `json:"Url"`

	// REQUIRED; 等待持续时间。
	WaitDuration int64 `json:"WaitDuration"`
}

type DescribeLivePadStreamListBody struct {

	// REQUIRED; 页码。
	PageNum int64 `json:"PageNum"`

	// REQUIRED; 分页大小。
	PageSize int64 `json:"PageSize"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称。
	App *string `json:"App,omitempty"`
}

type DescribeLivePadStreamListRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePadStreamListResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLivePadStreamListResResult `json:"Result,omitempty"`
}

type DescribeLivePadStreamListResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeLivePadStreamListResResult - 视请求的接口而定
type DescribeLivePadStreamListResResult struct {

	// REQUIRED; 流列表。
	StreamList []DescribeLivePadStreamListResResultStreamListItem `json:"StreamList"`
}

type DescribeLivePadStreamListResResultStreamListItem struct {

	// REQUIRED; 应用程序的名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeLivePlayStatusCodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：（默认值）1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空时表示查询所有域名下产生的请求状态码占比数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询请求状态码占比数据的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLivePlayStatusCodeDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 请求类型，取值及含义如下所示。
	// * Access：（默认值）推流请求和拉流请求；
	// * Source：回源请求。
	Type *string `json:"Type,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLivePlayStatusCodeDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePlayStatusCodeDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePlayStatusCodeDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLivePlayStatusCodeDataResResult `json:"Result"`
}

type DescribeLivePlayStatusCodeDataResResponseMetadata struct {

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

type DescribeLivePlayStatusCodeDataResResult struct {

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	StatusDataList []DescribeLivePlayStatusCodeDataResResultStatusDataListItem `json:"StatusDataList"`

	// REQUIRED; 当前查询条件下的状态码占比数据。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusSummaryDataListItem `json:"StatusSummaryDataList"`

	// REQUIRED; 请求类型，取值及含义如下所示。
	// * Access：推流请求和拉流请求；
	// * Source：回源请求。
	Type string `json:"Type"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLivePlayStatusCodeDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分后的数据。
	StatusDetailDataList []*DescribeLivePlayStatusCodeDataResResultStatusDetailDataListItem `json:"StatusDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLivePlayStatusCodeDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的状态码详细数据。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusDataListPropertiesItemsItem `json:"StatusSummaryDataList"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDataListPropertiesItemsItem struct {

	// 占比。
	Percent *float32 `json:"Percent,omitempty"`

	// 状态码。
	StatusCode *int32 `json:"StatusCode,omitempty"`

	// 出现次数。
	Value *int32 `json:"Value,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDetailDataListItem struct {

	// 拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// 运营商。
	ISP *string `json:"ISP,omitempty"`

	// 每个时间点的粒度数据。
	StatusDataList []*DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsItem `json:"StatusDataList,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 按状态码区分的数据列表。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsStatusSummaryDataListItem `json:"StatusSummaryDataList"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsStatusSummaryDataListItem struct {

	// 占比
	Percent *float32 `json:"Percent,omitempty"`

	// 状态码
	StatusCode *int32 `json:"StatusCode,omitempty"`

	// 出现次数
	Value *int32 `json:"Value,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusSummaryDataListItem struct {

	// 当前状态码出现次数在总状态码次数中的占比。
	Percent *float32 `json:"Percent,omitempty"`

	// 请求的状态码。
	StatusCode *int32 `json:"StatusCode,omitempty"`

	// 当前状态码出现的次数。
	Value *int32 `json:"Value,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePullToPushBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 支持域名拆分
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流转推任务群组列表，默认为空，表示查询所有拉流转推任务群组的带宽用量。
	GroupList []*string `json:"GroupList,omitempty"`
}

type DescribeLivePullToPushBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePullToPushBandwidthDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLivePullToPushBandwidthDataResResult          `json:"Result,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                       `json:"Version"`
	Error   *DescribeLivePullToPushBandwidthDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
	BandwidthDetailDataList []*DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度。
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流转推任务群组列表。
	GroupList []*string `json:"GroupList,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 查询时间范围内的维度下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按推流地址类型维度进行数据拆分时的地址类型信息。
	DstAddrType *string `json:"DstAddrType,omitempty"`

	// 按任务群组维度进行数据拆分时的群组信息。
	Group *string `json:"Group,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 转推带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 1 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 支持群组拆分
	DetailField []*string `json:"DetailField,omitempty"`

	// 群组
	GroupList []*string `json:"GroupList,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入
	// App 和 Stream。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePullToPushDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePullToPushDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLivePullToPushDataResResult          `json:"Result,omitempty"`
}

type DescribeLivePullToPushDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                              `json:"Version"`
	Error   *DescribeLivePullToPushDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLivePullToPushDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLivePullToPushDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDataListItem `json:"PullToPushDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的拉流转推总时长，单位为分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Group 即拉流转推任务群组维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流转推任务群组。
	GroupList []*string `json:"GroupList,omitempty"`

	// 按维度拆分后的数据。
	PullToPushDetailDataList []*DescribeLivePullToPushDataResResultPullToPushDetailDataListItem `json:"PullToPushDetailDataList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推总时长，单位为分钟。
	Value float32 `json:"Value"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem `json:"PullToPushDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的拉流转推总时长，单位分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 按任务群组维度进行数据拆分时的群组信息。
	Group *string `json:"Group,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 该时间片内的拉流转推总时长，单位分钟，保留小数点后 2 位
	Value float32 `json:"Value"`
}

type DescribeLivePushStreamCountDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain
	// 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 推流域名列表，默认为空，表示查询所有全部域名下的推流峰值流数。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。
	UserRegionList []*DescribeLivePushStreamCountDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePushStreamCountDataBodyUserRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePushStreamCountDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePushStreamCountDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLivePushStreamCountDataResResult `json:"Result"`
}

type DescribeLivePushStreamCountDataResResponseMetadata struct {

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

type DescribeLivePushStreamCountDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的推流数量最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TotalStreamDataList []DescribeLivePushStreamCountDataResResultTotalStreamDataListItem `json:"TotalStreamDataList"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 推流域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// 按维度拆分后的数据。
	StreamDetailDataList []*DescribeLivePushStreamCountDataResResultStreamDetailDataListItem `json:"StreamDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLivePushStreamCountDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePushStreamCountDataResResultStreamDetailDataListItem struct {

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下的所有时间粒度数据。
	TotalStreamDataList []DescribeLivePushStreamCountDataResResultStreamDetailDataListPropertiesItemsItem `json:"TotalStreamDataList"`
}

type DescribeLivePushStreamCountDataResResultStreamDetailDataListPropertiesItemsItem struct {

	// REQUIRED
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePushStreamCountDataResResultTotalStreamDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的推流数量最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePushStreamCountDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePushStreamInfoDataBody struct {

	// REQUIRED; 查询的结束时间。只能查询93d以内的数据
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间。
	StartTime string `json:"StartTime"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名列表，缺省情况表示该用户的所有域名
	DomainList []*string `json:"DomainList,omitempty"`

	// 分页页数，默认1
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页大小， 默认20
	PageSize *int32 `json:"PageSize,omitempty"`

	// 直播流名称
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePushStreamInfoDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePushStreamInfoDataResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLivePushStreamInfoDataResResult `json:"Result,omitempty"`
}

type DescribeLivePushStreamInfoDataResResponseMetadata struct {

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

// DescribeLivePushStreamInfoDataResResult - 视请求的接口而定
type DescribeLivePushStreamInfoDataResResult struct {

	// REQUIRED; 结束时间。格式rfc3339
	EndTime string `json:"EndTime"`

	// REQUIRED; 分页信息
	Pagination DescribeLivePushStreamInfoDataResResultPagination `json:"Pagination"`

	// REQUIRED; 按照搜索过滤字段和时间粒度聚合的数据
	PushStreamInfoDataList []DescribeLivePushStreamInfoDataResResultPushStreamInfoDataListItem `json:"PushStreamInfoDataList"`

	// REQUIRED; 起始时间。格式rfc3339
	StartTime string `json:"StartTime"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 直播流名称
	Stream *string `json:"Stream,omitempty"`
}

// DescribeLivePushStreamInfoDataResResultPagination - 分页信息
type DescribeLivePushStreamInfoDataResResultPagination struct {

	// REQUIRED; 当前页数
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页大小
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总共推流个数
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLivePushStreamInfoDataResResultPushStreamInfoDataListItem struct {

	// REQUIRED; 应用名称
	App string `json:"App"`

	// REQUIRED; 推流开始到结束的时长，单位s
	Duration int32 `json:"Duration"`

	// REQUIRED; 推流结束时间，格式rfc3339
	EndTime string `json:"EndTime"`

	// REQUIRED; 显示推流客户端的IP地址，如没有IP信息，返回空
	IP string `json:"IP"`

	// REQUIRED; 推流开始时间，格式rfc3339
	StartTime string `json:"StartTime"`

	// REQUIRED; 直播流名称
	Stream string `json:"Stream"`

	// REQUIRED; 推流断开原因
	StreamBreakReason string `json:"StreamBreakReason"`
}

type DescribeLivePushStreamMetricsBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：（默认值）30 秒。
	Aggregation *int32 `json:"Aggregation,omitempty"`
}

type DescribeLivePushStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePushStreamMetricsResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLivePushStreamMetricsResResult          `json:"Result,omitempty"`
}

type DescribeLivePushStreamMetricsResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                 `json:"Version"`
	Error   *DescribeLivePushStreamMetricsResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLivePushStreamMetricsResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLivePushStreamMetricsResResult struct {

	// 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 推流域名。
	Domain *string `json:"Domain,omitempty"`

	// 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 所有时间粒度的数据。
	MetricList []*DescribeLivePushStreamMetricsResResultMetricListItem `json:"MetricList,omitempty"`

	// 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePushStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveRecordDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain
	// 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的录制用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveRecordDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveRecordDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveRecordDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveRecordDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *DescribeLiveRecordDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveRecordDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveRecordDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	RecordDataList []DescribeLiveRecordDataResResultRecordDataListItem `json:"RecordDataList"`

	// REQUIRED; 当前查询条件下的录制并发路数最大值。
	RecordPeak int32 `json:"RecordPeak"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	RecordDetailDataList []*DescribeLiveRecordDataResResultRecordDetailDataListItem `json:"RecordDetailDataList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveRecordDataResResultRecordDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内录制并发路数最大值。
	Value int32 `json:"Value"`
}

type DescribeLiveRecordDataResResultRecordDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	RecordDataList []DescribeLiveRecordDataResResultRecordDetailDataListPropertiesItemsItem `json:"RecordDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的录制并发路数最大值。
	RecordPeak int32 `json:"RecordPeak"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeLiveRecordDataResResultRecordDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 录制峰值
	Value int32 `json:"Value"`
}

type DescribeLiveRegionDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveRegionDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveRegionDataResResult `json:"Result"`
}

type DescribeLiveRegionDataResResponseMetadata struct {

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

type DescribeLiveRegionDataResResult struct {

	// REQUIRED; 区域信息。
	Areas []DescribeLiveRegionDataResResultAreasItem `json:"Areas"`
}

type DescribeLiveRegionDataResResultAreasItem struct {

	// REQUIRED; 大区标识符。
	Code string `json:"Code"`

	// REQUIRED; 国家信息。
	Countries []DescribeLiveRegionDataResResultAreasPropertiesItemsItem `json:"Countries"`

	// REQUIRED; 大区名称。
	Name string `json:"Name"`
}

type DescribeLiveRegionDataResResultAreasPropertiesItemsItem struct {

	// REQUIRED; 国家标识符。
	Code string `json:"Code"`

	// REQUIRED; 国家名称。
	Name string `json:"Name"`

	// REQUIRED; 省份信息，国外暂不支持该参数。
	Provinces []DescribeLiveRegionDataResResultAreasPropertiesItemsProvincesItem `json:"Provinces"`
}

type DescribeLiveRegionDataResResultAreasPropertiesItemsProvincesItem struct {

	// REQUIRED; 省份标识符。
	Code string `json:"Code"`

	// REQUIRED; 省份名称。
	Name string `json:"Name"`
}

type DescribeLiveSnapshotDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的截图用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入
	// App 和 Stream。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveSnapshotDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSnapshotDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveSnapshotDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveSnapshotDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                            `json:"Version"`
	Error   *DescribeLiveSnapshotDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveSnapshotDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveSnapshotDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	SnapshotDataList []DescribeLiveSnapshotDataResResultSnapshotDataListItem `json:"SnapshotDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图总张数。
	Total int32 `json:"Total"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	SnapshotDetailData []*DescribeLiveSnapshotDataResResultSnapshotDetailDataItem `json:"SnapshotDetailData,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveSnapshotDataResResultSnapshotDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的截图总张数。
	Value int32 `json:"Value"`
}

type DescribeLiveSnapshotDataResResultSnapshotDetailDataItem struct {

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	SnapshotDataList []DescribeLiveSnapshotDataResResultSnapshotDetailDataPropertiesItemsItem `json:"SnapshotDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的截图总张数。
	Total int32 `json:"Total"`
}

type DescribeLiveSnapshotDataResResultSnapshotDetailDataPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s；例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 截图总张数
	Value int32 `json:"Value"`
}

type DescribeLiveSourceBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 回源流的应用名称，查询流粒度数据时必传，且需同时传入 Domain 和 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 查询流粒度的回源带宽监控数据时，需同时指定 Domain 、App 和 Stream 来指定回源流。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度的回源带宽监控数据时，需同时指定 Domain 、App
	// 和 Stream 来指定回源流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的回源带宽监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList
	// 和 Domain 传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 回源流的流名称，查询流粒度数据时必传，且需同时传入 Domain 和 App。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度的回源带宽监控数据时，需同时指定
	// Domain 、App 和 Stream 来指定回源流。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。
	UserRegionList []*DescribeLiveSourceBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceBandwidthDataBodyUserRegionListItem struct {

	// 区域信息的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSourceBandwidthDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveSourceBandwidthDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                   `json:"Version"`
	Error   *DescribeLiveSourceBandwidthDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResult struct {

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 按维度拆分后的数据。
	BandwidthDetailDataList []*DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询时间范围内的回源峰值带宽，单位为 Mbps。
	PeakBandwidth *float32 `json:"PeakBandwidth,omitempty"`

	// 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveSourceBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的回源峰值带宽，单位为 Mbps。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem struct {

	// 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []*DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList,omitempty"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按维度进行数据拆分后，当前维度的回源峰值带宽，单位为 Mbps。
	PeakBandwidth *float32 `json:"PeakBandwidth,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// 时间片内回源带宽峰值，单位 Mbps
	Bandwidth *float32 `json:"Bandwidth,omitempty"`

	// 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp *string `json:"TimeStamp,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceStreamMetricsBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看回源流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 30 秒进行数据聚合。
	Aggregation *int32 `json:"Aggregation,omitempty"`
}

type DescribeLiveSourceStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSourceStreamMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveSourceStreamMetricsResResult `json:"Result"`
}

type DescribeLiveSourceStreamMetricsResResponseMetadata struct {

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

type DescribeLiveSourceStreamMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	MetricList []DescribeLiveSourceStreamMetricsResResultMetricListItem `json:"MetricList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveSourceStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveSourceTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天，单次查询最大时间跨度与数据拆分维度和数据聚合时间粒度有关，详细如下。
	// * 当不进行维度拆分或只使用一个维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 31 天；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 31 天。
	//
	//
	// * 当使用两个或两个以上维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 3 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 7 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。
	// * 300：（默认值）5 分钟。
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 回源流的应用名称，查询流粒度数据时必传，且需同时传入 Domain 和 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 查询流粒度数据时，需同时指定 Domain 、App 和 Stream 来指定回源流。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度数据时，需同时指定 Domain 、App 和
	// Stream 来指定回源流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的回源流量带宽监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList 和 Domain 传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 回源流的流名称，查询流粒度数据时必传，且需同时传入 Domain 和 App。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时指定
	// Domain 、App 和 Stream 来指定回源流。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。
	UserRegionList []*DescribeLiveSourceTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceTrafficDataBodyUserRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSourceTrafficDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveSourceTrafficDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveSourceTrafficDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                 `json:"Version"`
	Error   *DescribeLiveSourceTrafficDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveSourceTrafficDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveSourceTrafficDataResResult struct {

	// REQUIRED; 峰值带宽。
	PeakBandwidth float32 `json:"PeakBandwidth"`

	// REQUIRED; 查询时间范围内的回源总流量，单位为 GB。
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。
	TrafficDetailDataList []*DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveSourceTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 带宽。取值范围为 [ ]，单位为，默认值为``。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的回源流量，单位 GB。
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 峰值带宽。
	PeakBandwidth float32 `json:"PeakBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的回源总流量，单位为 GB。
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 带宽。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 回源流量，单位 GB
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamCountDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天，单次查询最大时间跨度与数据拆分维度和数据聚合时间粒度有关，详细如下。
	// * 当不进行维度拆分或只使用一个维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 31 天；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 31 天。
	//
	//
	// * 当使用两个或两个以上维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 3 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 7 天。 :::
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟；
	// * 300：（默认值）5 分钟；
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 直播流使用的域名列表，默认为空，表示查询所有全部域名下的峰值流数。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// ISP 列表。
	ISPList []*string `json:"ISPList,omitempty"`

	// 流类型，缺省情况下表示全部类型，支持的流类型取值如下。
	// * push：推流；
	// * relay-source：回源流；
	// * transcode：转码流。
	StreamType []*string `json:"StreamType,omitempty"`

	// 用户区域列表。
	UserRegionList []*DescribeLiveStreamCountDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveStreamCountDataBodyUserRegionListItem struct {

	// 区域名称。
	Area *string `json:"Area,omitempty"`

	// 国家代码。
	Country *string `json:"Country,omitempty"`

	// 省份。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamCountDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamCountDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveStreamCountDataResResult `json:"Result"`
}

type DescribeLiveStreamCountDataResResponseMetadata struct {

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

type DescribeLiveStreamCountDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下流数的最大值，表示在指定的时间范围内推流、回源流或转码流的最高并发数。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TotalStreamDataList []DescribeLiveStreamCountDataResResultTotalStreamDataListItem `json:"TotalStreamDataList"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// ISP 列表。
	ISPList []*string `json:"ISPList,omitempty"`

	// 按维度拆分后的数据。
	StreamDetailDataList []*DescribeLiveStreamCountDataResResultStreamDetailDataListItem `json:"StreamDetailDataList,omitempty"`

	// 流类型，流类型说明如下。
	// * push：拉流；
	// * relay-source：回源流；
	// * transcode：转码流。
	StreamType []*string `json:"StreamType,omitempty"`

	// 用户区域列表。
	UserRegionList []*DescribeLiveStreamCountDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveStreamCountDataResResultStreamDetailDataListItem struct {

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下的所有时间粒度数据。
	TotalStreamDataList []DescribeLiveStreamCountDataResResultStreamDetailDataListPropertiesItemsItem `json:"TotalStreamDataList"`
}

type DescribeLiveStreamCountDataResResultStreamDetailDataListPropertiesItemsItem struct {

	// REQUIRED
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamCountDataResResultTotalStreamDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的流数最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamCountDataResResultUserRegionListItem struct {

	// 区域。
	Area *string `json:"Area,omitempty"`

	// 国家。
	Country *string `json:"Country,omitempty"`

	// 省份。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 使用流名称进行查询的方式，默认值为 strict，支持的取值即含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 在线流的来源类型，默认为空，表示查询所有来源类型，支持的取值即含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType *string `json:"SourceType,omitempty" query:"SourceType"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 在线流的流类型，默认为空，表示查询所有类型，支持的取值即含义如下所示。
	// * origin：原始流；
	// * trans：转码流。
	StreamType *string `json:"StreamType,omitempty" query:"StreamType"`

	// 域名空间，即直播流地址的域名所属的域名空间，默认为空，表示查询所有域名空间下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeLiveStreamInfoByPageRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamInfoByPageResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveStreamInfoByPageResResult          `json:"Result,omitempty"`
}

type DescribeLiveStreamInfoByPageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                `json:"Version"`
	Error   *DescribeLiveStreamInfoByPageResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveStreamInfoByPageResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveStreamInfoByPageResResult struct {

	// REQUIRED; 查询结果中在线流的数量。
	RoughCount int32 `json:"RoughCount"`

	// 在线流信息列表。
	StreamInfoList []*DescribeLiveStreamInfoByPageResResultStreamInfoListItem `json:"StreamInfoList,omitempty"`
}

type DescribeLiveStreamInfoByPageResResultStreamInfoListItem struct {

	// REQUIRED; 在线流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 在线流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 在线流的 ID。
	ID int64 `json:"ID"`

	// REQUIRED; 在线流的开始时间，RFC3339 格式的 UTC 时间戳，精度为毫秒。
	SessionStartTime string `json:"SessionStartTime"`

	// REQUIRED; 在线流的来源类型，取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType string `json:"SourceType"`

	// REQUIRED; 在线流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 在线流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeLiveStreamSessionDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天，单次查询最大时间跨度与数据拆分维度和数据聚合时间粒度有关，详细如下。
	// * 当不进行维度拆分或只使用一个维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 31 天；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 31 天。
	//
	//
	// * 当使用两个或两个以上维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 3 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 7 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。
	// * 300：（默认值）5 分钟。
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip 查询流粒度的请求数和在线人数数据时，需同时指定
	// Domain 、App 和 Stream 来指定直播流。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度的请求数和在线人数数据时，需同时指定 Domain 、
	// App 和 Stream 来指定直播流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的请求数和在线人数。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList 和 Domain 传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 在线人数类型，支持如下枚举值： Online（默认值）：瞬时session链接数； Viewer：一分钟session链接总数； ClientIP：一分钟拉流IP总数。
	OnlineUserType *string `json:"OnlineUserType,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。
	RegionList []*DescribeLiveStreamSessionDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 数据流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveStreamSessionDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamSessionDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamSessionDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveStreamSessionDataResResult `json:"Result"`
}

type DescribeLiveStreamSessionDataResResponseMetadata struct {

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

type DescribeLiveStreamSessionDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDataListItem `json:"SessionDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的请求数。
	TotalRequest int32 `json:"TotalRequest"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议；
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属的区域列表，缺省情况下表示所有区域。
	RegionList []*DescribeLiveStreamSessionDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分的数据。
	SessionDetailDataList []*DescribeLiveStreamSessionDataResResultSessionDetailDataListItem `json:"SessionDetailDataList,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveStreamSessionDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamSessionDataResResultSessionDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的在线人数最大值。
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 当前数据聚合时间粒度内的请求数。
	Request int32 `json:"Request"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamSessionDataResResultSessionDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDetailDataListPropertiesItemsItem `json:"SessionDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的请求数。
	TotalRequest int32 `json:"TotalRequest"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveStreamSessionDataResResultSessionDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 在线人数
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 请求数
	Request int32 `json:"Request"`

	// REQUIRED; 时间片起始时刻。RFC3339 时间，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamStateQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; 流名称，取值与直播流地址的 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream" query:"Stream"`

	// 填写直播流使用的域名，默认为空，表示查询所有直推流和回源流的状态和类型。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名。
	// :::tipVhost 和 Domain 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的直播流使用的域名所属的域名空间。 :::tipVhost 和 Domain 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeLiveStreamStateRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamStateResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveStreamStateResResult          `json:"Result,omitempty"`
}

type DescribeLiveStreamStateResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                           `json:"Version"`
	Error   *DescribeLiveStreamStateResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveStreamStateResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveStreamStateResResult struct {

	// REQUIRED; 直播流状态，取值及含义如下所示。
	// * online：在线流；
	// * offline：历史流；
	// * forbidden：禁推流。
	StreamState string `json:"stream_state"`

	// REQUIRED; 直播流类型，取值及含义如下所示。
	// * push：直推流；
	// * pull：回源流。
	Type string `json:"type"`
}

type DescribeLiveTimeShiftDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 86400 秒进行数据聚合。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名空间列表，默认为空，表示查询您视频直播产品下所有域名产生的时移存储用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取域名所在的域名空间信息。
	Vhosts []*string `json:"Vhosts,omitempty"`
}

type DescribeLiveTimeShiftDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTimeShiftDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveTimeShiftDataResResult `json:"Result"`
}

type DescribeLiveTimeShiftDataResResponseMetadata struct {

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

type DescribeLiveTimeShiftDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TimeShiftDataList []DescribeLiveTimeShiftDataResResultTimeShiftDataListItem `json:"TimeShiftDataList"`

	// 域名空间列表。
	Vhosts []*string `json:"Vhosts,omitempty"`
}

type DescribeLiveTimeShiftDataResResultTimeShiftDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的时移文件存储用量，单位为 GB。
	Storage float32 `json:"Storage"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveTopPlayDataBody struct {

	// REQUIRED; 结束时间，RFC3339 格式，例如：2021-04-14T00:00:00+08:00 单次最长跨度是31天 历史查询范围是366天
	EndTime string `json:"EndTime"`

	// REQUIRED; 起始时间，RFC3339 格式，例如：2021-04-13T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// VhostList和DomainList 2选1，支持为空，都为空查询全部
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。 最多展示排名前1000的数据。
	PageNum *string `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 10，取值范围为 [1,1000]。
	PageSize *string `json:"PageSize,omitempty"`

	// 查询类型，枚举值：
	// * Domain
	// * Stream（默认值）
	QueryType *string `json:"QueryType,omitempty"`

	// 排序指标，枚举值：
	// * PeakBandwidth 带宽峰值（默认值）
	// * AvgBandwidth 平均带宽
	// * TotalTraffic 流量加和
	SortBy *string `json:"SortBy,omitempty"`

	// VhostList和DomainList 2选1，支持为空，都为空查询全部
	VhostList []*string `json:"VhostList,omitempty"`
}

type DescribeLiveTopPlayDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTopPlayDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result DescribeLiveTopPlayDataResResult `json:"Result"`
}

type DescribeLiveTopPlayDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeLiveTopPlayDataResResult - 视请求的接口而定
type DescribeLiveTopPlayDataResResult struct {

	// REQUIRED; 带宽和流量详细数据。
	DataItemList []DescribeLiveTopPlayDataResResultDataItemListItem `json:"DataItemList"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询结果的分页信息。
	Pagination DescribeLiveTopPlayDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询类型，取值及含义如下所示。
	// * Domain ：查询 TOPN 域名的的流量带宽信息。
	// * Stream（默认值）查询 TOPN 直播流的流量带宽信息。
	QueryType string `json:"QueryType"`

	// REQUIRED; TOPN 结果的排序指标，取值及含义如下所示。
	// * PeakBandwidth（默认值）：以峰值带宽值降序展示查询结果。
	// * AvgBandwidth：以平均带宽值降序展示查询结果。
	// * TotalTraffic：以流量加值降序展示查询结果。
	SortBy string `json:"SortBy"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 域名空间列表。
	VhostList []string `json:"VhostList"`
}

type DescribeLiveTopPlayDataResResultDataItemListItem struct {

	// REQUIRED; 平均带宽，单位bps
	AvgBandwidth float32 `json:"AvgBandwidth"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 带宽峰值，单位bps
	PeakBandwidth float32 `json:"PeakBandwidth"`

	// REQUIRED; 流量加和，单位Byte
	TotalTraffic float32 `json:"TotalTraffic"`

	// AppName，根据Domain查询时为空
	App *string `json:"App,omitempty"`

	// 流名，根据Domain查询时为空
	Stream *string `json:"Stream,omitempty"`
}

// DescribeLiveTopPlayDataResResultPagination - 查询结果的分页信息。
type DescribeLiveTopPlayDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum string `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize string `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount string `json:"TotalCount"`
}

type DescribeLiveTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。聚合粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。聚合粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。聚合粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain
	// 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的流量用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveTrafficDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射 [https://www.volcengine.com/docs/6469/114196]
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射 [https://www.volcengine.com/docs/6469/114196]
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射 [https://www.volcengine.com/docs/6469/114196]
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveTrafficDataResResult `json:"Result"`
}

type DescribeLiveTrafficDataResResponseMetadata struct {

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

type DescribeLiveTrafficDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 查询时间范围内的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
	TrafficDetailDataList []*DescribeLiveTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveTrafficDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总上行流量，单位 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	TrafficDataList []DescribeLiveTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行流量，单位 GB
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 时间片起始时刻。RFC3339 时间，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveTrafficDataResResultUserRegionListItem struct {

	// 大区
	Area *string `json:"Area,omitempty"`

	// 国家
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTranscodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 86400 秒进行数据聚合。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的转码用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入
	// App 和 Stream。 :::
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，默认为空表示不指定编码格式，支持的取值和含义如下所示。
	// * Normal_H264：H.264 标准转码；
	// * Normal_H265：H.265 标准转码；
	// * Normal_H266：H.266 标准转码；
	// * ByteHD_H264：H.264 极智超清；
	// * ByteHD_H265：H.265 极智超清；
	// * ByteHD_H266：H.266 极智超清；
	// * ByteQE：画质增强；
	// * Audio：纯音频流。
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

type DescribeLiveTranscodeDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTranscodeDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveTranscodeDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveTranscodeDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                             `json:"Version"`
	Error   *DescribeLiveTranscodeDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveTranscodeDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveTranscodeDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询时间范围内的转码总时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TranscodeDataList []DescribeLiveTranscodeDataResResultTranscodeDataListItem `json:"TranscodeDataList"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264 极智超清；
	// - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

type DescribeLiveTranscodeDataResResultTranscodeDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的转码时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0P：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType string `json:"TranscodeType"`
}

type DescribeLiveTranscodeInfoDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip 查询流粒度数据时，需同时传入App和Stream。
	// :::
	App *string `json:"App,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的转码用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分页页数，默认1
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页大小， 默认20，取值范围[1,100000]
	PageSize *int32 `json:"PageSize,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入App和Stream。
	// :::
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264 极智超清；
	// - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

type DescribeLiveTranscodeInfoDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTranscodeInfoDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveTranscodeInfoDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveTranscodeInfoDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                 `json:"Version"`
	Error   *DescribeLiveTranscodeInfoDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveTranscodeInfoDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveTranscodeInfoDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 分页信息
	Pagination DescribeLiveTranscodeInfoDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的起始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TranscodeInfoDataList []DescribeLiveTranscodeInfoDataResResultTranscodeInfoDataListItem `json:"TranscodeInfoDataList"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264 极智超清；
	// - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

// DescribeLiveTranscodeInfoDataResResultPagination - 分页信息
type DescribeLiveTranscodeInfoDataResResultPagination struct {

	// REQUIRED; 当前页数
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页大小
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总个数
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveTranscodeInfoDataResResultTranscodeInfoDataListItem struct {

	// REQUIRED; 转码时长，单位分钟
	Duration float32 `json:"Duration"`

	// REQUIRED; 结束转码时间
	EndTime string `json:"EndTime"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0P：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; 开始转码时间
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名
	Stream string `json:"Stream"`

	// REQUIRED; 转码任务ID
	TaskID string `json:"TaskID"`

	// REQUIRED; 转码后缀
	TranscodeSuffix string `json:"TranscodeSuffix"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType string `json:"TranscodeType"`
}

type DescribeRecordTaskFileHistoryBody struct {

	// REQUIRED; 开始录制时间，RFC3339 格式的时间戳，精度为秒。当您查询指定录制任务详情时，DateFrom 应设置为开始时间之前的任意时间。
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 结束录制时间，RFC3339 格式的时间戳，精度为秒。结束时间需晚于 DateFrom，且与 DateFrom 间隔不超过 7 天。
	DateTo string `json:"DateTo"`

	// REQUIRED; 查询数据的页码，默认为 1，表示查询第一页的数据，取值范围为正整数。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为正整数。
	PageSize int32 `json:"PageSize"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，默认为空表示查询 vhost 下的所有录制历史。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。
	App *string `json:"App,omitempty"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相同，默认为空表示查询 App 下的所有录制历史。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100
	// 个字符。 :::tip 如果指定 Stream，必须同时指定 App 的值。 :::
	Stream *string `json:"Stream,omitempty"`

	// 录制文件保存位置，支持的取值及含义如下所示。
	// * tos：存储到 TOS（默认值）；
	// * vod：存储到 VOD。
	Type *string `json:"Type,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间，默认为空表示查询所有录制历史。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeRecordTaskFileHistoryRes struct {

	// REQUIRED
	ResponseMetadata DescribeRecordTaskFileHistoryResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeRecordTaskFileHistoryResResult          `json:"Result,omitempty"`
}

type DescribeRecordTaskFileHistoryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                 `json:"Version"`
	Error   *DescribeRecordTaskFileHistoryResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeRecordTaskFileHistoryResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeRecordTaskFileHistoryResResult struct {

	// REQUIRED; 录制文件列表。
	Data []DescribeRecordTaskFileHistoryResResultDataItem `json:"Data"`

	// REQUIRED; 查询结果的分页信息。
	Pagination DescribeRecordTaskFileHistoryResResultPagination `json:"Pagination"`
}

type DescribeRecordTaskFileHistoryResResultDataItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 存储位置为 TOS 时的 Bucket。
	Bucket string `json:"Bucket"`

	// REQUIRED; 录制时长。
	Duration string `json:"Duration"`

	// REQUIRED; 结束录制时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 结束录制时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTimeUTC string `json:"EndTimeUTC"`

	// REQUIRED; 录制文件的文件名。
	FileName string `json:"FileName"`

	// REQUIRED; 录制文件存储格式。
	Format string `json:"Format"`

	// REQUIRED; 存储位置为 TOS 时，在 Bucket 中的存储路径。
	Path string `json:"Path"`

	// REQUIRED; 开始录制时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 开始录制时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTimeUTC string `json:"StartTimeUTC"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// REQUIRED; 录制文件保存在 VOD 时，录制文件的 ID。
	Vid string `json:"Vid"`
}

// DescribeRecordTaskFileHistoryResResultPagination - 查询结果的分页信息。
type DescribeRecordTaskFileHistoryResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeRefererBody struct {

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 参数 Domain 和 App 至少传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	// :::tip
	// * 参数 Domain 和 Vhost 传且仅传一个。
	// * 参数 Domain 和 App 至少传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看拉流域名所属的域名空间。 :::tip
	// 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeRefererRes struct {

	// REQUIRED
	ResponseMetadata DescribeRefererResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeRefererResResult          `json:"Result,omitempty"`
}

type DescribeRefererResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *DescribeRefererResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeRefererResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeRefererResResult struct {

	// Referer 防盗链信息列表。
	RefererList []*DescribeRefererResResultRefererListItem `json:"RefererList,omitempty"`
}

type DescribeRefererResResultRefererListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; Referer 防盗链详情列表。
	RefererInfoList []DescribeRefererResResultRefererListPropertiesItemsItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeRefererResResultRefererListPropertiesItemsItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词，返回值为 referer。
	Key string `json:"Key"`

	// REQUIRED; 优先级，当前默认返回值为 0。当多域名返回值一致时，按照域名输入顺序区分，越早加入列表的域名优先级越高。
	Priority int32 `json:"Priority"`

	// REQUIRED; referer 防盗链黑白名单类型，取值即含义如下所示。
	// * deny：黑名单；
	// * allow：白名单。
	Type string `json:"Type"`

	// REQUIRED; Referer 字段规则，即设置的黑名单或白名单的域名。
	Value string `json:"Value"`
}

type DescribeRelaySourceV3Body struct {

	// REQUIRED; 直播流使用的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，即直播流地址的AppName字段取值，默认为空，表示查询当前域名空间下所有播放触发回源配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty"`

	// 回源组名称。
	Group *string `json:"Group,omitempty"`
}

type DescribeRelaySourceV3Res struct {

	// REQUIRED
	ResponseMetadata DescribeRelaySourceV3ResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeRelaySourceV3ResResult          `json:"Result,omitempty"`
}

type DescribeRelaySourceV3ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *DescribeRelaySourceV3ResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeRelaySourceV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeRelaySourceV3ResResult struct {

	// 回源配置列表。
	RelaySourceConfigList []*DescribeRelaySourceV3ResResultRelaySourceConfigListItem `json:"RelaySourceConfigList,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 回源组配置详情。
	GroupDetails []DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsItem `json:"GroupDetails"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsItem struct {

	// REQUIRED; 回源组名称。
	Group string `json:"Group"`

	// REQUIRED; 回源服务器配置列表。
	Servers []DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersItem `json:"Servers"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersItem struct {

	// REQUIRED; 回源地址。
	RelaySourceDomain string `json:"RelaySourceDomain"`

	// REQUIRED; 自定义回源参数。
	RelaySourceParams map[string]string `json:"RelaySourceParams"`

	// REQUIRED; 回源协议。
	RelaySourceProtocol string `json:"RelaySourceProtocol"`
}

type DescribeStreamQuotaConfigBody struct {

	// REQUIRED; 待查询限额配置的推流域名或拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名信息。
	Domain string `json:"Domain"`
}

type DescribeStreamQuotaConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeStreamQuotaConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeStreamQuotaConfigResResult `json:"Result,omitempty"`
}

type DescribeStreamQuotaConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *DescribeStreamQuotaConfigResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeStreamQuotaConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// DescribeStreamQuotaConfigResResult - 视请求的接口而定
type DescribeStreamQuotaConfigResResult struct {

	// REQUIRED; 限额配置列表。
	QuotaList []DescribeStreamQuotaConfigResResultQuotaListItem `json:"QuotaList"`
}

type DescribeStreamQuotaConfigResResultQuotaListItem struct {

	// REQUIRED; 推流域名或拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 配置详情列表。
	QuotaDetailList []DescribeStreamQuotaConfigResResultQuotaListPropertiesItemsItem `json:"QuotaDetailList"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

// DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig - 拉流域名的带宽限额配置信息。 :::tipDomain 为拉流域名时返回。
// :::
type DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig struct {

	// REQUIRED; 拉流域名的带宽限额。
	Quota int32 `json:"Quota"`

	// REQUIRED; 带宽限额的计量单位。
	QuotaUnit string `json:"QuotaUnit"`

	// 带宽限额的告警阈值，缺省情况表示未设置告警。
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`

	// 带宽限额告警的计量单位，缺省情况表示不未设置告警。
	AlarmThresholdUnit *string `json:"AlarmThresholdUnit,omitempty"`
}

// DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig - 推流域名的推流路数限额配置信息。 :::tipDomain 为推流域名时返回。
// :::
type DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig struct {

	// REQUIRED; 推流域名的推流路数限额。
	Quota int32 `json:"Quota"`

	// 推流路数限额告警阈值，缺省情况表示未设置告警。
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`
}

type DescribeStreamQuotaConfigResResultQuotaListPropertiesItemsItem struct {

	// 拉流域名的带宽限额配置信息。 :::tipDomain 为拉流域名时返回。 :::
	BandwidthConfig *DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig `json:"BandwidthConfig,omitempty"`

	// 推流域名的推流路数限额配置信息。 :::tipDomain 为推流域名时返回。 :::
	StreamConfig *DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig `json:"StreamConfig,omitempty"`
}

type DisableDomainBody struct {

	// REQUIRED; 待禁用域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要待禁用域名的信息。
	Domain string `json:"Domain"`
}

type DisableDomainRes struct {

	// REQUIRED
	ResponseMetadata DisableDomainResResponseMetadata `json:"ResponseMetadata"`
}

type DisableDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *DisableDomainResResponseMetadataError `json:"Error,omitempty"`
}

type DisableDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type EnableDomainBody struct {

	// REQUIRED; 待启用域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要待启用域名的信息。
	Domain string `json:"Domain"`
}

type EnableDomainRes struct {

	// REQUIRED
	ResponseMetadata EnableDomainResResponseMetadata `json:"ResponseMetadata"`
}

type EnableDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *EnableDomainResResponseMetadataError `json:"Error,omitempty"`
}

type EnableDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type EnableHTTPHeaderConfigBody struct {

	// REQUIRED; 启用或禁用配置，取值及含义如下所示。
	// * true：启用；
	// * false：禁用。
	Enable bool `json:"Enable"`

	// REQUIRED; 0: response 1: request
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Vhost
	// 取值。
	Vhost string `json:"Vhost"`

	// 拉流域名，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Domain 取值。
	Domain *string `json:"Domain,omitempty"`
}

type EnableHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata EnableHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type EnableHTTPHeaderConfigResResponseMetadata struct {

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

type ForbidStreamBody struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 流名称，取值与直播流地址的 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 直播流使用的域名，您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看待禁推的直播流使用的域名。
	// :::tip 参数 Domain 和 Vhost
	// 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 禁推的结束时间，RFC3339 格式的 UTC 时间，精度为毫秒，禁推有效期最长为 90 天，默认为当前时间加 90 天。
	EndTime *string `json:"EndTime,omitempty"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看待禁推的直播流使用的域名所属的域名空间。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type ForbidStreamRes struct {

	// REQUIRED
	ResponseMetadata ForbidStreamResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type ForbidStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                `json:"Version"`
	Error   *ForbidStreamResResponseMetadataError `json:"Error,omitempty"`
}

type ForbidStreamResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type GeneratePlayURLBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// REQUIRED; 生成地址为源流地址/转码流地址还是abr地址
	StreamType string `json:"StreamType"`

	// 拉流地址的过期时间，RFC3339 格式的 UTC 时间，精度为秒，过期后需要重新生成。缺省情况下表示当前时间往后的 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime
	// 的时间为准。 :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 转码流后缀，默认为空，表示生成源流的拉流地址。配置不为空时表示生成转码流的拉流地址，可通过调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853]
	// 接口查询对应流的转码流后缀。
	Suffix *string `json:"Suffix,omitempty"`

	// CDN 类型，默认值为 fcdn，支持如下取值。
	// * fcdn：火山引擎流媒体直播 CDN；
	// * 3rd：第三方 CDN。
	Type *string `json:"Type,omitempty"`

	// 拉流地址的有效时长，单位为秒，超过有效时长后需要重新生成。取值范围为正整数，缺省值为 604800，即 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ValidDuration *int32 `json:"ValidDuration,omitempty"`
}

type GeneratePlayURLRes struct {

	// REQUIRED
	ResponseMetadata GeneratePlayURLResResponseMetadata `json:"ResponseMetadata"`
	Result           *GeneratePlayURLResResult          `json:"Result,omitempty"`
}

type GeneratePlayURLResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *GeneratePlayURLResResponseMetadataError `json:"Error,omitempty"`
}

type GeneratePlayURLResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type GeneratePlayURLResResult struct {

	// REQUIRED; 生成地址对应匹配到的鉴权模版类型
	AuthType string `json:"AuthType"`

	// REQUIRED; 拉流地址信息列表。
	URLList []GeneratePlayURLResResultURLListItem `json:"URLList"`
}

type GeneratePlayURLResResultURLListItem struct {

	// REQUIRED; CDN 类型。
	// * fcdn：火山引擎流媒体直播 CDN；
	// * 3rd：第三方 CDN。
	CDN string `json:"CDN"`

	// REQUIRED; 协议类型，包括 hls、flv、rtmp、udp 和 cmaf。
	Protocol string `json:"Protocol"`

	// REQUIRED; 子流地址。仅当 StreamType 为 abr 时返回。
	SubStreamURL []GeneratePlayURLResResultURLListPropertiesItemsItem `json:"SubStreamURL"`

	// REQUIRED; 地址类型，取值及含义如下所示。
	// * pull：拉流地址；
	// * 3rd_play(relay_source)：第三方回源地址，当配置了回源且 CDN 类型为第三方 CDN 时返回；
	// * 3rd_play(relay_sink)：第三方转推地址，当配置了拉流转推且 CDN 类型为第三方 CDN 时返回。
	Type string `json:"Type"`

	// REQUIRED; 生成的拉流地址。
	URL string `json:"URL"`
}

type GeneratePlayURLResResultURLListPropertiesItemsItem struct {

	// REQUIRED; 子流转码后缀。
	Suffix string `json:"Suffix"`

	// REQUIRED; 地址标签。包括 drm、hls加密等。
	Tag string `json:"Tag"`
}

type GeneratePushURLBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间，即推流域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看推流域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 推流域名，默认为空，表示生成域名空间下所有推流域名的推流地址。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要的推流域名。
	Domain *string `json:"Domain,omitempty"`

	// 推流地址的过期时间，RFC3339 格式的时间字符串，精度为秒，过期后需要重新生成。缺省情况下表示当前时间往后的 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 推流地址的有效时长，单位为秒，超过有效时长后需要重新生成。取值范围为正整数，默认值为 604800，即 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ValidDuration *int32 `json:"ValidDuration,omitempty"`
}

type GeneratePushURLRes struct {

	// REQUIRED
	ResponseMetadata GeneratePushURLResResponseMetadata `json:"ResponseMetadata"`
	Result           *GeneratePushURLResResult          `json:"Result,omitempty"`
}

type GeneratePushURLResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *GeneratePushURLResResponseMetadataError `json:"Error,omitempty"`
}

type GeneratePushURLResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type GeneratePushURLResResult struct {

	// REQUIRED; 生成地址对应匹配到的鉴权模版类型
	AuthType string `json:"AuthType"`

	// REQUIRED; RTMP 推流地址。
	PushURLList []string `json:"PushURLList"`

	// REQUIRED; 推流地址详情。
	PushURLListDetail []GeneratePushURLResResultPushURLListDetailItem `json:"PushURLListDetail"`

	// REQUIRED; RTM 推流地址。
	RtmURLList []string `json:"RtmURLList"`

	// REQUIRED; RTMP over SRT 推流地址。
	RtmpOverSrtURLList []string `json:"RtmpOverSrtURLList"`

	// REQUIRED; TS over SRT 推流地址。
	TsOverSrtURLList []string `json:"TsOverSrtURLList"`

	// REQUIRED; 网络传输推流地址。
	WebTransportURLList []string `json:"WebTransportURLList"`
}

type GeneratePushURLResResultPushURLListDetailItem struct {

	// REQUIRED; OBS 推流地址，例如，rtmp://push.example.com/live/。
	DomainApp string `json:"DomainApp"`

	// REQUIRED; OBS 推流名称，例如，streamname1?volcTime=1675652376&volcSecret=c57d247c2f19b395b6ec9b182******7。
	StreamSign string `json:"StreamSign"`

	// REQUIRED; 推流地址。
	URL string `json:"URL"`
}

type GetCarouselDetailBody struct {

	// REQUIRED; 待查询的轮播任务 ID，任务的唯一标识。调用 CreateCarouselTask 接口创建轮播任务时返回。
	TaskID string `json:"TaskID"`
}

type GetCarouselDetailRes struct {

	// REQUIRED
	ResponseMetadata GetCarouselDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetCarouselDetailResResult `json:"Result"`
}

type GetCarouselDetailResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type GetCarouselDetailResResult struct {

	// REQUIRED; 包含轮播任务相关信息的数据对象。
	Data GetCarouselDetailResResultData `json:"Data"`
}

// GetCarouselDetailResResultData - 包含轮播任务相关信息的数据对象。
type GetCarouselDetailResResultData struct {

	// REQUIRED; 最新的播放列表序列号
	LastOperationIndex int32 `json:"LastOperationIndex"`

	// REQUIRED; 当前播放列表序列号
	LastSuccessOperationIndex int32 `json:"LastSuccessOperationIndex"`

	// REQUIRED; 当前的播放信息，json字符串
	PlayInfo string `json:"PlayInfo"`

	// REQUIRED; 当前的播单信息
	Rule string `json:"Rule"`

	// REQUIRED; 任务状态： pending：任务等待调度中 prepare：任务初始化中 running：任务运行中 prestop：任务停止中 done：任务已经停止
	Status string `json:"Status"`
}

type GetCloudMixTaskDetailBody struct {

	// REQUIRED; 混流任务 ID，您可以通过 ListCloudMixTask [https://www.volcengine.com/docs/6469/1271157] 接口获取混流任务 ID。
	TaskID string `json:"TaskID"`
}

type GetCloudMixTaskDetailRes struct {

	// REQUIRED
	ResponseMetadata GetCloudMixTaskDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetCloudMixTaskDetailResResult `json:"Result"`
}

type GetCloudMixTaskDetailResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type GetCloudMixTaskDetailResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data GetCloudMixTaskDetailResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// GetCloudMixTaskDetailResResultData - 返回的数据。
type GetCloudMixTaskDetailResResultData struct {

	// REQUIRED; 任务最近一次更新的版本标识。
	LastOperationIndex int32 `json:"LastOperationIndex"`

	// REQUIRED; 任务最近一次成功更新的版本标识。
	LastSuccessOperationIndex int32 `json:"LastSuccessOperationIndex"`

	// REQUIRED; 混流任务详细配置的 Json 字符串。
	Rule string `json:"Rule"`

	// REQUIRED; 混流任务状态，取值及含义如下所示。
	// * pending：任务调度被阻塞。
	// * prepare：正在准备任务资源。
	// * runing：任务进行中。
	// * prestop：正在清理任务资源。
	// * done：任务已结束。
	Status string `json:"Status"`

	// REQUIRED; 混流任务 ID。
	TaskID string `json:"TaskID"`
}

type GetHLSEncryptDataKeyQuery struct {

	// REQUIRED; 视频直播服务端生成的 M3U8 文件中写入的每个 TS 分片的密钥 ID。
	KeyID string `json:"KeyID" query:"KeyID"`
}

type GetHLSEncryptDataKeyRes struct {

	// REQUIRED
	ResponseMetadata GetHLSEncryptDataKeyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetHLSEncryptDataKeyResResult `json:"Result,omitempty"`
}

type GetHLSEncryptDataKeyResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// GetHLSEncryptDataKeyResResult - 视请求的接口而定
type GetHLSEncryptDataKeyResResult struct {

	// REQUIRED; 密钥。
	DataKey string `json:"DataKey"`
}

type GetLiveVideoQualityAnalysisTaskDetailBody struct {

	// 查询的任务 ID。 :::tipName 和 ID 二选一必填。 :::
	ID *int64 `json:"ID,omitempty"`

	// 查询的任务名称。 :::tipName 和 ID 二选一必填。 :::
	Name *string `json:"Name,omitempty"`
}

type GetLiveVideoQualityAnalysisTaskDetailRes struct {

	// REQUIRED
	ResponseMetadata GetLiveVideoQualityAnalysisTaskDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetLiveVideoQualityAnalysisTaskDetailResResult `json:"Result"`
}

type GetLiveVideoQualityAnalysisTaskDetailResResponseMetadata struct {

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

type GetLiveVideoQualityAnalysisTaskDetailResResult struct {

	// REQUIRED; 测评任务详细信息。
	Task GetLiveVideoQualityAnalysisTaskDetailResResultTask `json:"Task"`
}

// GetLiveVideoQualityAnalysisTaskDetailResResultTask - 测评任务详细信息。
type GetLiveVideoQualityAnalysisTaskDetailResResultTask struct {

	// REQUIRED; 测试任务的持续时长，单位为秒。
	Duration int32 `json:"Duration"`

	// REQUIRED; 请提供具体的参数ID和类型string，以便我为您生成参数描述。
	ID int64 `json:"ID"`

	// REQUIRED; 画质测评的打点间隔。
	Interval int32 `json:"Interval"`

	// REQUIRED; 任务名称。
	Name string `json:"Name"`

	// REQUIRED; 画质测评结果。
	ScoringResult GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResult `json:"ScoringResult"`

	// REQUIRED; 画质测评视频流的播放地址。
	SrcURL string `json:"SrcURL"`
}

// GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResult - 画质测评结果。
type GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResult struct {

	// REQUIRED; 画质测评结果详细信息。
	VQScoreLive []GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResultVQScoreLiveItem `json:"VQScoreLive"`
}

type GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResultVQScoreLiveItem struct {

	// REQUIRED; 测评打点的时间，Unix 时间戳，精度为秒。
	Timestamp int32 `json:"Timestamp"`

	// REQUIRED; 测评点的画质得分，画质评分范围为 0 到 100，评分越高表示视频画面色彩越好。不同的评分段对应不同的视频质量感受：
	// * 0～60：主观感受较差。
	// * 60～70：主观感受良好。
	// * 70～100：主观感受清晰。
	Value float32 `json:"Value"`
}

type GetPullRecordTaskBody struct {

	// REQUIRED; 任务 ID，录制任务的唯一标识。您可以调用 ListPullRecordTask [https://www.volcengine.com/docs/6469/1111480] 获取任务 ID。
	TaskID string `json:"TaskID"`
}

type GetPullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result GetPullRecordTaskResResult `json:"Result"`
}

type GetPullRecordTaskResResponseMetadata struct {

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

// GetPullRecordTaskResResult - 视请求的接口而定
type GetPullRecordTaskResResult struct {

	// REQUIRED; 任务详情。
	TaskDetail GetPullRecordTaskResResultTaskDetail `json:"TaskDetail"`
}

// GetPullRecordTaskResResultTaskDetail - 任务详情。
type GetPullRecordTaskResResultTaskDetail struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 域名。
	Domain string `json:"Domain"`

	// REQUIRED; 结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 状态。
	Status string `json:"Status"`

	// REQUIRED; 任务ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type KillStreamBody struct {

	// REQUIRED; 直播流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 直播流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间，您可以调用 DescribeLiveStreamInfoByPage [https://www.volcengine.com/docs/6469/1126841] 接口，查看待断开的在线流的信息，包括 Vhost、APP
	// 和 Stream。
	Vhost string `json:"Vhost"`
}

type KillStreamRes struct {

	// REQUIRED
	ResponseMetadata KillStreamResResponseMetadata `json:"ResponseMetadata"`
}

type KillStreamResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *KillStreamResResponseMetadataError `json:"Error,omitempty"`
}

type KillStreamResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListBindEncryptDRMBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询符合域名空间取值的所有的 DRM 加密配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为
	// 1 到 30 个字符。
	App *string `json:"App,omitempty"`
}

type ListBindEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata ListBindEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result ListBindEncryptDRMResResult `json:"Result"`
}

type ListBindEncryptDRMResResponseMetadata struct {

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

// ListBindEncryptDRMResResult - 视请求的接口而定
type ListBindEncryptDRMResResult struct {

	// DRM 加密配置列表。
	DRMBindingList []*ListBindEncryptDRMResResultDRMBindingListItem `json:"DRMBindingList,omitempty"`
}

type ListBindEncryptDRMResResultDRMBindingListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 加密类型，支持的取值及含义如下所示。
	// * FairPlay：使用 FairPlay 技术的商业 DRM 加密；
	// * Widevine：使用 Widevine 技术的商业 DRM 加密；
	// * PlayReady：使用 PlayReady 技术的商业 DRM 加密；
	// * ClearKey：HLS 标准加密。
	// :::tip DRM 加密与 HLS 标准加密不可同时配置。 :::
	DRMSystems []string `json:"DRMSystems"`

	// REQUIRED; 是否开启源流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。
	EncryptOriginStream bool `json:"EncryptOriginStream"`

	// REQUIRED; 是否开启转码流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。
	EncryptTranscodeStream bool `json:"EncryptTranscodeStream"`

	// REQUIRED; 进行 DRM 加密的转码流对应的转码流后缀配置。
	EncryptTranscodeSuffix []string `json:"EncryptTranscodeSuffix"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type ListCarouselTaskBody struct {

	// REQUIRED; 分页功能，展示第几页
	Page int32 `json:"Page"`

	// REQUIRED; 分页功能，页大小
	PageSize int32 `json:"PageSize"`
}

type ListCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata ListCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListCarouselTaskResResult `json:"Result"`
}

type ListCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type ListCarouselTaskResResult struct {

	// REQUIRED; 轮播任务数据对象。
	Data ListCarouselTaskResResultData `json:"Data"`
}

// ListCarouselTaskResResultData - 轮播任务数据对象。
type ListCarouselTaskResResultData struct {

	// REQUIRED; 满足查询条件的轮播任务总数。
	Count int32 `json:"Count"`

	// REQUIRED; 轮播任务的数组，每个元素表示一个任务的详细信息。
	Result []ListCarouselTaskResResultDataResultItem `json:"Result"`
}

type ListCarouselTaskResResultDataResultItem struct {

	// REQUIRED; 任务的创建时间，RFC3339 格式的时间戳，精度为秒。
	CreatedAt ListCarouselTaskResResultDataResultItemCreatedAt `json:"CreatedAt"`

	// REQUIRED; 轮播任务名称。
	Name string `json:"Name"`

	// REQUIRED; 轮播任务的当前状态。取值和含义如下：
	// * pending：任务等待调度中；
	// * prepare：任务初始化中；
	// * running：任务运行中；
	// * prestop：任务停止中；
	// * done：任务已经停止。
	Status string `json:"Status"`

	// REQUIRED; 任务的结束时间，RFC3339 格式的时间戳，精度为秒。
	StoppedAt ListCarouselTaskResResultDataResultItemStoppedAt `json:"StoppedAt"`

	// REQUIRED; 轮播任务的唯一标识。
	TaskID string `json:"TaskID"`

	// REQUIRED; 任务的更新时间，RFC3339 格式的时间戳，精度为秒。
	UpdatedAt ListCarouselTaskResResultDataResultItemUpdatedAt `json:"UpdatedAt"`
}

// ListCarouselTaskResResultDataResultItemCreatedAt - 任务的创建时间，RFC3339 格式的时间戳，精度为秒。
type ListCarouselTaskResResultDataResultItemCreatedAt struct {

	// REQUIRED; 任务的创建时间，RFC3339 格式的时间戳，精度为秒。
	Time string `json:"Time"`
}

// ListCarouselTaskResResultDataResultItemStoppedAt - 任务的结束时间，RFC3339 格式的时间戳，精度为秒。
type ListCarouselTaskResResultDataResultItemStoppedAt struct {

	// REQUIRED; 任务的结束时间，RFC3339 格式的时间戳，精度为秒。
	Time string `json:"Time"`
}

// ListCarouselTaskResResultDataResultItemUpdatedAt - 任务的更新时间，RFC3339 格式的时间戳，精度为秒。
type ListCarouselTaskResResultDataResultItemUpdatedAt struct {

	// REQUIRED; 任务的更新时间，RFC3339 格式的时间戳，精度为秒。
	Time string `json:"Time"`
}

type ListCertV2Body struct {

	// 证书是否启用，默认值为 true，支持的取值及含义如下所示。
	// * true：启用证书；
	// * false：禁用证书。
	Available *bool `json:"Available,omitempty"`

	// 证书名称，支持输入证书名称中的关键字，进行模糊查询.
	CertName *string `json:"CertName,omitempty"`

	// 域名，查询该域名对应的证书，支持精确查询。默认为空，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 只有填了Available，这个字段才生效。
	Expiring *bool `json:"Expiring,omitempty"`

	// 页码。不填默认返回全部。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 分页大小。不填默认返回所有。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 项目名称。
	ProjectName *string `json:"ProjectName,omitempty"`
}

type ListCertV2Res struct {

	// REQUIRED
	ResponseMetadata ListCertV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListCertV2ResResult          `json:"Result,omitempty"`
}

type ListCertV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *ListCertV2ResResponseMetadataError `json:"Error,omitempty"`
}

type ListCertV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListCertV2ResResult struct {

	// REQUIRED; 本次查询所有证书的过期信息。
	ExpirationInfo ListCertV2ResResultExpirationInfo `json:"ExpirationInfo"`

	// 证书列表。
	CertList []*ListCertV2ResResultCertListItem `json:"CertList,omitempty"`

	// 总数。
	Total *int32 `json:"Total,omitempty"`
}

type ListCertV2ResResultCertListItem struct {

	// REQUIRED; 与证书绑定的域名列表。
	CertDomainList []string `json:"CertDomainList"`

	// REQUIRED; 证书 ID。
	CertID string `json:"CertID"`

	// REQUIRED; 证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 证书链 ID。
	ChainID string `json:"ChainID"`

	// REQUIRED; 火山引擎证书中心证书链 ID。
	ChainIDVolc string `json:"ChainIDVolc"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为秒。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为秒。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 证书状态，由证书管理平台返回，支持的取值如下所示。
	// * OK：正常；
	// * Expire：过期；
	// * 30days：有效期剩余 30 天；
	// * 15days：有效期剩余 15 天；
	// * 7days：有效期剩余 7 天；
	// * 1days：有效期剩余 1 天。
	Status string `json:"Status"`
}

// ListCertV2ResResultExpirationInfo - 本次查询所有证书的过期信息。
type ListCertV2ResResultExpirationInfo struct {

	// REQUIRED; 生效数量。
	ActiveNum int32 `json:"ActiveNum"`

	// REQUIRED; 快要过期数量，一个月之内
	ClosingExpireNum int32 `json:"ClosingExpireNum"`

	// REQUIRED; 过期数量。
	ExpireNum int32 `json:"ExpireNum"`
}

type ListCloudMixTaskBody struct {

	// REQUIRED; 查询数据的页码。
	Page int32 `json:"Page"`

	// REQUIRED; 每页显示的数据条数，最大值为 100。
	PageSize int32 `json:"PageSize"`
}

type ListCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata ListCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListCloudMixTaskResResult `json:"Result"`
}

type ListCloudMixTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type ListCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data ListCloudMixTaskResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// ListCloudMixTaskResResultData - 返回的数据。
type ListCloudMixTaskResResultData struct {

	// REQUIRED; 查询结果的数据总条数。
	Count int32 `json:"Count"`

	// REQUIRED; 查询结果数据详细信息。
	Result []ListCloudMixTaskResResultDataResultItem `json:"Result"`
}

type ListCloudMixTaskResResultDataResultItem struct {

	// REQUIRED; 混流任务创建时间。
	CreatedAt ListCloudMixTaskResResultDataResultItemCreatedAt `json:"CreatedAt"`

	// REQUIRED; 混流任务名称。
	Name string `json:"Name"`

	// REQUIRED; 混流任务状态，取值及含义如下所示。
	// * pending：任务调度被阻塞。
	// * prepare：正在准备任务资源。
	// * runing：任务进行中。
	// * prestop：正在清理任务资源。
	// * done：任务已结束。
	Status string `json:"Status"`

	// REQUIRED; 混流任务结束时间。
	StoppedAt ListCloudMixTaskResResultDataResultItemStoppedAt `json:"StoppedAt"`

	// REQUIRED; 混流任务 ID。
	TaskID string `json:"TaskID"`

	// REQUIRED; 混流任务更新时间。
	UpdatedAt ListCloudMixTaskResResultDataResultItemUpdatedAt `json:"UpdatedAt"`
}

// ListCloudMixTaskResResultDataResultItemCreatedAt - 混流任务创建时间。
type ListCloudMixTaskResResultDataResultItemCreatedAt struct {

	// REQUIRED; 时间。
	Time string `json:"Time"`
}

// ListCloudMixTaskResResultDataResultItemStoppedAt - 混流任务结束时间。
type ListCloudMixTaskResResultDataResultItemStoppedAt struct {

	// REQUIRED
	Time string `json:"Time"`
}

// ListCloudMixTaskResResultDataResultItemUpdatedAt - 混流任务更新时间。
type ListCloudMixTaskResResultDataResultItemUpdatedAt struct {

	// REQUIRED
	Time string `json:"Time"`
}

type ListCommonTransPresetDetailBody struct {

	// 模板名称列表，缺省情况下，表示查询所有系统内置转码档位。
	PresetList []*string `json:"PresetList,omitempty"`
}

type ListCommonTransPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata ListCommonTransPresetDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListCommonTransPresetDetailResResult          `json:"Result,omitempty"`
}

type ListCommonTransPresetDetailResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                               `json:"Version"`
	Error     *ListCommonTransPresetDetailResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                              `json:"RequestID,omitempty"`
}

type ListCommonTransPresetDetailResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListCommonTransPresetDetailResResult struct {

	// 极智超清转码配置。
	NarrowBandHDPresetDetail []*ListCommonTransPresetDetailResResultNarrowBandHDPresetDetailItem `json:"NarrowBandHDPresetDetail,omitempty"`

	// 标准转码配置。
	StandardPresetDetail []*ListCommonTransPresetDetailResResultStandardPresetDetailItem `json:"StandardPresetDetail,omitempty"`
}

type ListCommonTransPresetDetailResResultNarrowBandHDPresetDetailItem struct {

	// 音频编码格式，支持的取值及含义如下。
	// * aac：使用 AAC 编码格式；
	// * opus：使用 Opus 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效。
	Acodec *string `json:"Acodec,omitempty"`

	// 视频分辨率自适应模式开关。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As 取值为 0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As 取值为 1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，默认值为 3，取值为 0 时表示去除 B 帧。
	// 最大 B 帧数量的取值范围根据视频编码格式（Vcodec）的不同有所差异，取值范围如下所示。
	// * 视频编码格式为 H.264 （Vcodec 取值为 h264）时取值范围为 [0,7]；
	// * 视频编码格式为 H.265 或 H.266 （Vcodec 取值为 h265 或 h266）时取值范围为 [0,3]、7、15。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 帧率，单位为 fps。帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度，单位为 px。 :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，单位为 px。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度，单位为 px。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码后缀，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度，单位为 px。 :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListCommonTransPresetDetailResResultStandardPresetDetailItem struct {

	// 音频编码格式。包括以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应，按照Width和Height的取值进行拉伸；
	// * 1：开启宽高自适应，按照ShortSide或LongSide等比缩放。
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 帧率，单位为 fps。帧率越大，画面越流畅
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。 :::tip 当 As 的取值为 0 时，Width 和 Height 中任意参数取 0，表示保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 都取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 是否极智超清转码。
	// * true：极智超清；
	// * false：标准转码。
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度。 :::tip 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 都取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码流后缀名
	SuffixName *string `json:"SuffixName,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。 :::tip 当 As 的取值为 0 时，如果 Width 和 Height 中任意参数取 0，表示保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListDomainDetailBody struct {

	// REQUIRED; 查询数据的页码，取值为 1 时表示查询第一页的数据，取值范围为 [1,1000]。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值为 10 时表示每页展示 10 条域名信息，取值范围为 [1, 1000]。
	PageSize int32 `json:"PageSize"`

	// 域名名称列表，缺省情况下表示全部。
	DomainNameList []*string `json:"DomainNameList,omitempty"`

	// 域名加速区域列表，缺省情况下表示查看全部。支持的取值如下所示。
	// * cn：中国内地；
	// * cn-global：全球加速；
	// * cn-oversea：海外及港澳台。
	DomainRegionList []*string `json:"DomainRegionList,omitempty"`

	// 域名状态列表，缺省情况下表示查询全部状态的域名。支持的取值如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停；
	// * 6：域名未备案被封禁。
	DomainStatusList []*int32 `json:"DomainStatusList,omitempty"`

	// 域名类型列表，缺省情况下表示全部类型的域名。支持的取值如下所示。
	// * push：推流域名；
	// * pull-flv：拉流域名。
	DomainTypeList []*string `json:"DomainTypeList,omitempty"`

	// 域名空间列表，缺省情况下表示查询全部域名空间下的域名。
	VhostList []*string `json:"VhostList,omitempty"`
}

type ListDomainDetailRes struct {

	// REQUIRED
	ResponseMetadata ListDomainDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListDomainDetailResResult          `json:"Result,omitempty"`
}

type ListDomainDetailResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                    `json:"Version"`
	Error   *ListDomainDetailResResponseMetadataError `json:"Error,omitempty"`
}

type ListDomainDetailResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListDomainDetailResResult struct {

	// REQUIRED; 总记录数。
	Total int32 `json:"Total"`

	// 域名详细信息列表。
	DomainList []*ListDomainDetailResResultDomainListItem `json:"DomainList,omitempty"`
}

type ListDomainDetailResResultDomainListItem struct {

	// REQUIRED; CNAME 信息。
	CNAME string `json:"CNAME"`

	// REQUIRED; 绑定的 HTTPS 证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的 HTTPS 证书的证书链 ID 信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态，取值及含义如下所示。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 域名添加时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态，取值及含义如下所示。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; HTTP/2协议。
	HTTP2 bool `json:"HTTP2"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	// * 1：备案正常，未过期；
	// * 2：查存不到备案信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 域名空间所属的项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 绑定的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 域名加速区域，取值及含义如下所示。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态，取值及含义如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停；
	// * 6：域名未备案被封禁。
	Status int32 `json:"Status"`

	// REQUIRED; 域名空间的标签信息。
	Tags []ListDomainDetailResResultDomainListPropertiesItemsItem `json:"Tags"`

	// REQUIRED; 域名类型，取值及含义如下所示。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type ListDomainDetailResResultDomainListPropertiesItemsItem struct {

	// REQUIRED; 标签类型，取值及含义如下所示。
	// * System：系统内置标签；
	// * Custom：自定义标签。
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type ListHighLightTaskBody struct {

	// REQUIRED; 遍历任务分页数
	Page int32 `json:"Page"`

	// REQUIRED; 遍历任务单页任务数
	PageSize int32 `json:"PageSize"`

	// 遍历在某个时间范围内创建的高光剪辑任务，时间范围的截止时间，RFC3339 格式的时间戳，精度为秒
	CreateTimeEd *string `json:"CreateTimeEd,omitempty"`

	// 遍历在某个时间范围内创建的高光剪辑任务，时间范围的起始时间，RFC3339 格式的时间戳，精度为秒
	CreateTimeSt *string `json:"CreateTimeSt,omitempty"`

	// 便利结果任务排序，asc：升序, desc：降序。默认为desc
	Order *string `json:"Order,omitempty"`

	// 遍历在某个时间范围内停止的高光剪辑任务，时间范围的截止时间，RFC3339 格式的时间戳，精度为秒
	StopTimeEd *string `json:"StopTimeEd,omitempty"`

	// 遍历在某个时间范围内停止的高光剪辑任务，时间范围的起始时间，RFC3339 格式的时间戳，精度为秒
	StopTimeSt *string `json:"StopTimeSt,omitempty"`

	// 若干高光剪辑任务taskid
	TaskIDs []*string `json:"TaskIds,omitempty"`

	// 若干高光剪辑任务状态，init：任务初始化状态，pending：任务等待启动阶段，prepare：任务启动状态，running：任务生产状态，prestop：任务准备结束状态，done：任务完成状态，error：任务报错状态。默认则查询所有状态的任务
	TaskStatus []*string `json:"TaskStatus,omitempty"`
}

type ListHighLightTaskRes struct {

	// REQUIRED
	ResponseMetadata ListHighLightTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListHighLightTaskResResult `json:"Result,omitempty"`
}

type ListHighLightTaskResResponseMetadata struct {

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestID"`

	// 请求的接口名，属于请求的公共参数。
	Action *string `json:"Action,omitempty"`

	// 请求的Region，例如：cn-north-1
	Region *string `json:"Region,omitempty"`

	// 请求的服务，属于请求的公共参数。
	Service *string `json:"Service,omitempty"`

	// 请求的版本号，属于请求的公共参数。
	Version *string `json:"Version,omitempty"`
}

// ListHighLightTaskResResult - 视请求的接口而定
type ListHighLightTaskResResult struct {

	// REQUIRED; 返回的高光剪辑任务数据，包含任务总数和任务详情列表。
	Data ListHighLightTaskResResultData `json:"Data"`
}

// ListHighLightTaskResResultData - 返回的高光剪辑任务数据，包含任务总数和任务详情列表。
type ListHighLightTaskResResultData struct {

	// REQUIRED; 遍历任务总数
	Count int32 `json:"Count"`

	// 高光剪辑任务的详情列表，包含每个任务的基本信息和状态。
	Result []*ListHighLightTaskResResultDataResultItem `json:"Result,omitempty"`
}

type ListHighLightTaskResResultDataResultItem struct {

	// 高光剪辑任务所属的账号
	AccountID *string `json:"AccountID,omitempty"`

	// 任务名称
	Name *string `json:"Name,omitempty"`

	// 高光剪辑任务状态，init：任务初始化状态，pending：任务等待启动阶段，prepare：任务启动状态，running：任务生产状态，prestop：任务准备结束状态，done：任务完成状态，error：任务报错状态
	Status *string `json:"Status,omitempty"`

	// 创建高光剪辑任务时返回的taskid
	TaskID *string `json:"TaskID,omitempty"`
}

type ListLiveVideoQualityAnalysisTasksBody struct {

	// REQUIRED; 分页参数
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 分页参数
	PageSize int32 `json:"PageSize"`

	// 查询的任务ID列表， 和Names二选一
	IDs []*int64 `json:"IDs,omitempty"`

	// 查询的任务名称列表， 和TaskIDs二选一
	Names []*string `json:"Names,omitempty"`
}

type ListLiveVideoQualityAnalysisTasksRes struct {

	// REQUIRED
	ResponseMetadata ListLiveVideoQualityAnalysisTasksResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListLiveVideoQualityAnalysisTasksResResult `json:"Result"`
}

type ListLiveVideoQualityAnalysisTasksResResponseMetadata struct {

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

type ListLiveVideoQualityAnalysisTasksResResult struct {

	// REQUIRED; 查询的数据的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 画质测评任务列表。
	Tasks []ListLiveVideoQualityAnalysisTasksResResultTasksItem `json:"Tasks"`
}

type ListLiveVideoQualityAnalysisTasksResResultTasksItem struct {

	// 测评任务持续时长。
	Duration *int32 `json:"Duration,omitempty"`

	// 任务 ID。
	ID *int64 `json:"ID,omitempty"`

	// 画质测评的打点间隔。
	Interval *int32 `json:"Interval,omitempty"`

	// 任务名称。
	Name *string `json:"Name,omitempty"`

	// 进行画质测评的直播流地址。
	SrcURL *string `json:"SrcURL,omitempty"`
}

type ListPullRecordTaskBody struct {

	// REQUIRED; 分页数
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 分页的大小
	PageSize int32 `json:"PageSize"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名
	Domain *string `json:"Domain,omitempty"`

	// 流名称
	Stream *string `json:"Stream,omitempty"`

	// 域名空间名称
	Vhost *string `json:"Vhost,omitempty"`
}

type ListPullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata ListPullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListPullRecordTaskResResult `json:"Result,omitempty"`
}

type ListPullRecordTaskResResponseMetadata struct {

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

// ListPullRecordTaskResResult - 视请求的接口而定
type ListPullRecordTaskResResult struct {

	// REQUIRED; 直播录制任务列表。
	List []ListPullRecordTaskResResultListItem `json:"List"`

	// REQUIRED; 查询结果的分页信息。
	Pagination ListPullRecordTaskResResultPagination `json:"Pagination"`
}

type ListPullRecordTaskResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名或拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 录制的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 录制的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 4种状态: 停用，未开始，生效中，已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 任务 ID，任务的唯一标识。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListPullRecordTaskResResultPagination - 查询结果的分页信息。
type ListPullRecordTaskResResultPagination struct {

	// REQUIRED; 目前页数
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 目前分页页大小
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总共页数
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 任务数量
	TotalCount int32 `json:"TotalCount"`
}

type ListPullToPushGroupBody struct {

	// REQUIRED; 查询数据的页码，取值范围为 [1,1000]。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页现实的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize"`

	// 群组的状态，取值及含义如下所示。
	// * 0: （默认值）可用;
	// * 1: 已删除，不可用。
	StatusList []*int32 `json:"StatusList,omitempty"`
}

type ListPullToPushGroupRes struct {

	// REQUIRED
	ResponseMetadata ListPullToPushGroupResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListPullToPushGroupResResult `json:"Result,omitempty"`
}

type ListPullToPushGroupResResponseMetadata struct {

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

// ListPullToPushGroupResResult - 视请求的接口而定
type ListPullToPushGroupResResult struct {

	// REQUIRED; 拉流转推群组列表。
	List []ListPullToPushGroupResResultListItem `json:"List"`

	// REQUIRED; 查询结果的数据条数。
	Total int32 `json:"Total"`
}

type ListPullToPushGroupResResultListItem struct {

	// REQUIRED; 群组名称。
	Name string `json:"Name"`

	// REQUIRED; 群组所属的项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 群组的状态，取值及含义如下所示。
	// * 0: 可用;
	// * 1: 已删除，不可用。
	Status int32 `json:"Status"`

	// REQUIRED; 群组的标签信息。
	Tags []ListPullToPushGroupResResultListPropertiesItemsItem `json:"Tags"`
}

type ListPullToPushGroupResResultListPropertiesItemsItem struct {

	// REQUIRED; 标签类型，支持以下取值。
	// * System：系统内置标签；
	// * Custom：自定义标签。
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type ListPullToPushTaskQuery struct {

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	Page *int32 `json:"Page,omitempty" query:"Page"`

	// 每页显示的数据条数，默认为 20，最大值为 500。
	Size *int32 `json:"Size,omitempty" query:"Size"`

	// 拉流转推任务的名称，不区分大小写，支持模糊查询。 例如，title取值为doc时，则返回任务名称为docspace、docs、DOC等 title 中包含doc关键词的所有任务列表。
	Title *string `json:"Title,omitempty" query:"Title"`
}

type ListPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata ListPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListPullToPushTaskResResult `json:"Result"`
}

type ListPullToPushTaskResResponseMetadata struct {
	Action    *string                                     `json:"Action,omitempty"`
	Error     *ListPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                                     `json:"Region,omitempty"`
	RequestID *string                                     `json:"RequestId,omitempty"`
	Service   *string                                     `json:"Service,omitempty"`
	Version   *string                                     `json:"Version,omitempty"`
}

type ListPullToPushTaskResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListPullToPushTaskResResult struct {

	// 任务列表。
	List []*ListPullToPushTaskResResultListItem `json:"List,omitempty"`

	// 分页数量信息。
	Pagination *ListPullToPushTaskResResultPagination `json:"Pagination,omitempty"`
}

type ListPullToPushTaskResResultListItem struct {

	// 接收拉流转推任务状态回调的地址。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频时配置生效，参数取值及含义如下所示。
	// * -1：无限次循环，至任务结束；
	// * 0：有限次循环，循环次数以 PlayTimes 取值为准；
	// * >0：有限次循环，循环次数以 CycleMode 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 推流地址类型。
	// * 1：非第三方，即推流地址域名已添加到视频直播。
	// * 2：第三方，即推流地址域名未添加到视频直播。
	DstAddrType *int32 `json:"DstAddrType,omitempty"`

	// 任务的结束时间，RFC3339 格式的 UTC 时间，单位为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。拉流来源类型为点播视频时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
	// * 0：不开启；
	// * 1：开启。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源时返回此值。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频时返回此值。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 任务的开始时间，RFC3339 格式的 UTC 时间，单位为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 拉流转推任务的状态，支持如下取值。
	// * 停用；
	// * 未开始；
	// * 生效中；
	// * 已结束。
	Status *string `json:"Status,omitempty"`

	// 任务 ID，任务的唯一标识。
	TaskID *string `json:"TaskId,omitempty"`

	// 拉流转推任务的名称。
	Title *string `json:"Title,omitempty"`

	// 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type *int32 `json:"Type,omitempty"`

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*ComponentsGg7M1TSchemasListpulltopushtaskresPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems `json:"VodSrcAddrs,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *ListPullToPushTaskResResultListItemWatermark `json:"Watermark,omitempty"`
}

// ListPullToPushTaskResResultListItemWatermark - 为拉流转推视频添加的水印配置信息。
type ListPullToPushTaskResResultListItemWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture string `json:"Picture"`

	// REQUIRED; 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
	Ratio float32 `json:"Ratio"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosX float32 `json:"RelativePosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosY float32 `json:"RelativePosY"`
}

// ListPullToPushTaskResResultPagination - 分页数量信息。
type ListPullToPushTaskResResultPagination struct {

	// 当前任务所在分页。
	PageCur *int32 `json:"PageCur,omitempty"`

	// 每页显示的数据条数。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 查询结果的数据总页数。
	PageTotal *int32 `json:"PageTotal,omitempty"`

	// 查询结果的数据总条数。
	TotalCount *int32 `json:"TotalCount,omitempty"`
}

type ListPullToPushTaskV2Body struct {

	// 群组名称。
	GroupNames []*string `json:"GroupNames,omitempty"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	Page *int32 `json:"Page,omitempty"`

	// 每页显示的数据条数，默认为 20，最大值为 500。
	Size *int32 `json:"Size,omitempty"`

	// 拉流转推任务的名称，不区分大小写，支持模糊查询。 例如，title取值为doc时，则返回任务名称为docspace、docs、DOC等 title 中包含doc关键词的所有任务列表。
	Title *string `json:"Title,omitempty"`
}

type ListPullToPushTaskV2Res struct {

	// REQUIRED
	ResponseMetadata ListPullToPushTaskV2ResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListPullToPushTaskV2ResResult `json:"Result"`
}

type ListPullToPushTaskV2ResResponseMetadata struct {
	Action    *string                                       `json:"Action,omitempty"`
	Error     *ListPullToPushTaskV2ResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                                       `json:"Region,omitempty"`
	RequestID *string                                       `json:"RequestId,omitempty"`
	Service   *string                                       `json:"Service,omitempty"`
	Version   *string                                       `json:"Version,omitempty"`
}

type ListPullToPushTaskV2ResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListPullToPushTaskV2ResResult struct {

	// 任务列表。
	List []*ListPullToPushTaskV2ResResultListItem `json:"List,omitempty"`

	// 分页数量信息。
	Pagination *ListPullToPushTaskV2ResResultPagination `json:"Pagination,omitempty"`
}

type ListPullToPushTaskV2ResResultListItem struct {

	// 接收拉流转推任务状态回调的地址。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频（Type 为 1）时配置生效，参数取值及含义如下所示。
	// * -1：无限循环，至任务结束；
	// * 0：有限次循环，循环次数为 PlayTimes 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 推流地址类型。
	// * 1：非第三方，即推流地址域名已添加到视频直播。
	// * 2：第三方，即推流地址域名未添加到视频直播。
	DstAddrType *int32 `json:"DstAddrType,omitempty"`

	// 任务的结束时间，RFC3339 格式的 UTC 时间，单位为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 任务所属的群组名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 获取可用的群组。 :::tip
	// * 使用主账号调用时，为非必填，默认为空表示查询所有群组的任务列表。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。拉流来源类型为点播视频时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时配置生效。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
	// * 0：不开启；
	// * 1：开启。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源（Type 为 0）时返回此值。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频（type 为 1）时返回此值。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 任务的开始时间，RFC3339 格式的 UTC 时间，单位为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 拉流转推任务的状态，支持如下取值。
	// * 停用；
	// * 未开始；
	// * 生效中；
	// * 已结束。
	Status *string `json:"Status,omitempty"`

	// 任务 ID，任务的唯一标识。
	TaskID *string `json:"TaskId,omitempty"`

	// 拉流转推任务的名称。
	Title *string `json:"Title,omitempty"`

	// 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type *int32 `json:"Type,omitempty"`

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*Components1Nf1A8CSchemasListpulltopushtaskv2ResPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems `json:"VodSrcAddrs,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *ListPullToPushTaskV2ResResultListItemWatermark `json:"Watermark,omitempty"`
}

// ListPullToPushTaskV2ResResultListItemWatermark - 为拉流转推视频添加的水印配置信息。
type ListPullToPushTaskV2ResResultListItemWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture string `json:"Picture"`

	// REQUIRED; 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
	Ratio float32 `json:"Ratio"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosX float32 `json:"RelativePosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosY float32 `json:"RelativePosY"`
}

// ListPullToPushTaskV2ResResultPagination - 分页数量信息。
type ListPullToPushTaskV2ResResultPagination struct {

	// 当前任务所在分页。
	PageCur *int32 `json:"PageCur,omitempty"`

	// 每页显示的数据条数。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 查询结果的数据总页数。
	PageTotal *int32 `json:"PageTotal,omitempty"`

	// 查询结果的数据总条数。
	TotalCount *int32 `json:"TotalCount,omitempty"`
}

type ListRelaySourceV4Body struct {

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain string `json:"Domain"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	Page *int32 `json:"Page,omitempty"`

	// 每页显示的数据条数，默认为 20，最大值为 500。
	Size *int32 `json:"Size,omitempty"`
}

type ListRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata ListRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRelaySourceV4ResResult          `json:"Result,omitempty"`
}

type ListRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                     `json:"Version"`
	Error   *ListRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type ListRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListRelaySourceV4ResResult struct {

	// REQUIRED; 配置列表。
	List []ListRelaySourceV4ResResultListItem `json:"List"`

	// REQUIRED; 页码信息。
	Pagination ListRelaySourceV4ResResultPagination `json:"Pagination"`
}

type ListRelaySourceV4ResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 直播流的使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 回源结束时间，StartTime 和 EndTime 同时缺省的情况下，表示永久回源。
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 回源地址列表。
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 回源开始时间，StartTime 和 EndTime 同时缺省的情况下，表示永久回源。
	StartTime int32 `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

// ListRelaySourceV4ResResultPagination - 页码信息。
type ListRelaySourceV4ResResultPagination struct {

	// REQUIRED; 当前查询的页码。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type ListTimeShiftPresetV2Body struct {

	// REQUIRED; 时移类型，默认类型为 vod。
	// * vod：点播时移，表示查询时移录制存储在 VOD 中的时移配置；
	// * tos：直播时移，表示查询时移录制存储在 TOS 以及 fcdn-tos 中的时移配置。
	Type string `json:"Type"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要时移的直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListTimeShiftPresetV2Res struct {

	// REQUIRED
	ResponseMetadata ListTimeShiftPresetV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListTimeShiftPresetV2ResResult          `json:"Result,omitempty"`
}

type ListTimeShiftPresetV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *ListTimeShiftPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type ListTimeShiftPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListTimeShiftPresetV2ResResult struct {

	// 时移配置列表。
	List []*ListTimeShiftPresetV2ResResultListItem `json:"List,omitempty"`
}

type ListTimeShiftPresetV2ResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; TOS 存储对应的 Bucket。
	Bucket string `json:"Bucket"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 最大时移时长，即允许用户回看的最长时间，单位为秒。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 时移配置名称。
	Name string `json:"Name"`

	// REQUIRED; 直播时移配置启用状态。
	// * 0：配置中；
	// * 1：已启用。
	Status int32 `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 类型。默认类型为 vod。
	// * vod：录制类型为录制时移时，录制配置中存储位置为 VOD。
	// * tos：录制类型为录制时移时，录制配置中存储喂食为 TOS。
	// * fcdn-toS：独立时移。
	Type string `json:"Type"`

	// REQUIRED; 视频点播（VOD）空间名称。
	VODNamespace string `json:"VODNamespace"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListVhostRecordPresetV2Body struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 直播录制的存储类型，默认值为 tos，支持的取值及含义如下所示。
	// * vod：录制文件存在 VOD；
	// * tos：录制文件存在 TOS。
	Type *string `json:"Type,omitempty"`
}

type ListVhostRecordPresetV2Res struct {

	// REQUIRED
	ResponseMetadata ListVhostRecordPresetV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostRecordPresetV2ResResult          `json:"Result,omitempty"`
}

type ListVhostRecordPresetV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *ListVhostRecordPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostRecordPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostRecordPresetV2ResResult struct {

	// REQUIRED; 录制配置列表。
	PresetList []ListVhostRecordPresetV2ResResultPresetListItem `json:"PresetList"`
}

type ListVhostRecordPresetV2ResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 录制参数详细信息。
	SlicePresetV2 *ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 `json:"SlicePresetV2,omitempty"`
}

// ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 - 录制参数详细信息。
type ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 struct {

	// 录制配置 ID。
	ID *int32 `json:"ID,omitempty"`

	// 录制配置名称。
	Name *string `json:"Name,omitempty"`

	// 录制模板详细配置。
	RecordPresetConfig *ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig `json:"RecordPresetConfig,omitempty"`
}

type ListVhostSnapshotAuditPresetBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type ListVhostSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostSnapshotAuditPresetResResult          `json:"Result,omitempty"`
}

type ListVhostSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                `json:"Version"`
	Error   *ListVhostSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostSnapshotAuditPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostSnapshotAuditPresetResResult struct {

	// REQUIRED; 截图审核配置列表。
	PresetList []ListVhostSnapshotAuditPresetResResultPresetListItem `json:"PresetList"`
}

type ListVhostSnapshotAuditPresetResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 截图审核配置详细信息。
	AuditPreset ListVhostSnapshotAuditPresetResResultPresetListItemAuditPreset `json:"AuditPreset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListVhostSnapshotAuditPresetResResultPresetListItemAuditPreset - 截图审核配置详细信息。
type ListVhostSnapshotAuditPresetResResultPresetListItemAuditPreset struct {

	// REQUIRED; ToS 存储对应的 Bucket。
	Bucket string `json:"Bucket"`

	// REQUIRED; 截图审核结果回调地址配置。
	CallbackDetailList []ListVhostSnapshotAuditPresetResResultPresetListPropertiesItemsItem `json:"CallbackDetailList"`

	// REQUIRED; 截图审核配置的描述。
	Description string `json:"Description"`

	// REQUIRED; 截图间隔时间，单位为秒。
	Interval float32 `json:"Interval"`

	// REQUIRED; 审核标签名称，取值及含义如下。
	// * 301：涉黄；
	// * 302：涉敏1；
	// * 303：涉敏2；
	// * 304：广告；
	// * 305：引人不适；
	// * 306：违禁；
	// * 307：二维码；
	// * 308：诈骗；
	// * 309：不良画面；
	// * 310：未成年相关；
	// * 320：文字违规。
	Label []string `json:"Label"`

	// REQUIRED; 截图审核配置的名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; veimageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID string `json:"ServiceID"`

	// REQUIRED; 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject string `json:"SnapshotObject"`

	// REQUIRED; ToS 存储对应的 Bucket 下的存储目录。
	StorageDir string `json:"StorageDir"`

	// REQUIRED; 存储策略，取值及含义如下。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储全部图片。
	StorageStrategy int32 `json:"StorageStrategy"`

	// REQUIRED; 配置信息的更新时间，RFC3339 格式的 UTC 时间，精度为秒。
	UpdateTime string `json:"UpdateTime"`
}

type ListVhostSnapshotAuditPresetResResultPresetListPropertiesItemsItem struct {

	// REQUIRED; 回调的类型 HTTP。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
}

type ListVhostSnapshotPresetV2Body struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 截图配置中截图文件的存储位置，缺省情况下表示不对存储位置进行过滤，取值及含义如下所示。
	// * tos：TOS 对象存储服务；
	// * imageX：veImageX 图片服务。
	Type *string `json:"Type,omitempty"`
}

type ListVhostSnapshotPresetV2Res struct {

	// REQUIRED
	ResponseMetadata ListVhostSnapshotPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListVhostSnapshotPresetV2ResResult `json:"Result,omitempty"`
}

type ListVhostSnapshotPresetV2ResResponseMetadata struct {

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

// ListVhostSnapshotPresetV2ResResult - 视请求的接口而定
type ListVhostSnapshotPresetV2ResResult struct {

	// REQUIRED; 截图配置列表。
	PresetList []ListVhostSnapshotPresetV2ResResultPresetListItem `json:"PresetList"`
}

type ListVhostSnapshotPresetV2ResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 截图配置基础信息。
	SlicePresetV2 ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 `json:"SlicePresetV2"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

// ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 - 截图配置基础信息。
type ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 struct {

	// REQUIRED; 截图配置名称。
	Name string `json:"Name"`

	// REQUIRED; 截图配置详细信息。
	SnapshotPresetConfig ListVhostSnapshotPresetV2ResResultPresetListProperties `json:"SnapshotPresetConfig"`

	// REQUIRED; 截图配置生效状态，取值及含义如下所示。
	// * 1：生效；
	// * 0：不生效。
	Status int32 `json:"Status"`
}

// ListVhostSnapshotPresetV2ResResultPresetListProperties - 截图配置详细信息。
type ListVhostSnapshotPresetV2ResResultPresetListProperties struct {

	// REQUIRED; 截图间隔时间，单位为秒。
	Interval int32 `json:"Interval"`

	// 图片格式为 JPEG 时的截图参数。
	JPEGParam *ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesProperties `json:"JpegParam,omitempty"`

	// 截图格式为 JPG 时的截图参数。
	JpgParam *ComponentsSlabtaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparam `json:"JpgParam,omitempty"`
}

// ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesProperties - 图片格式为 JPEG 时的截图参数。
type ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesProperties struct {

	// REQUIRED; 当前格式的截图是否开启，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable bool `json:"Enable"`

	// 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	ImageXParam *ComponentsK46Cw0SchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpegparamPropertiesImagexparam `json:"ImageXParam,omitempty"`

	// 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	TOSParam *ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesPropertiesProperties `json:"TOSParam,omitempty"`
}

// ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesPropertiesProperties - 截图存储到 TOS 时的配置。 :::tip TOSParam
// 和 ImageXParam 配置且配置其中一个。 :::
type ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesPropertiesProperties struct {

	// REQUIRED; TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。
	Bucket string `json:"Bucket"`

	// REQUIRED; 截图是否使用 TOS 存储，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable bool `json:"Enable"`

	// REQUIRED; 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。
	ExactObject string `json:"ExactObject"`

	// REQUIRED; 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。
	OverwriteObject string `json:"OverwriteObject"`

	// REQUIRED; Bucket 目录。 例如，存储路径为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir string `json:"StorageDir"`
}

type ListVhostSubtitleTranscodePresetBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type ListVhostSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListVhostSubtitleTranscodePresetResResult `json:"Result"`
}

type ListVhostSubtitleTranscodePresetResResponseMetadata struct {

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

type ListVhostSubtitleTranscodePresetResResult struct {

	// REQUIRED; 字幕配置列表。
	PresetList []ListVhostSubtitleTranscodePresetResResultPresetListItem `json:"PresetList"`
}

type ListVhostSubtitleTranscodePresetResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 转码后缀标识。
	Suffixes []string `json:"Suffixes"`

	// REQUIRED; 字幕配置详细参数。
	TranscodePreset ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePreset `json:"TranscodePreset"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

// ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePreset - 字幕配置详细参数。
type ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePreset struct {

	// REQUIRED; 字幕配置的描述信息。
	Description string `json:"Description"`

	// REQUIRED; 预设配置，使用预设配置是系统将自动对字体大小、字幕行数、每行最大字符数和边距参数（MarginVertical 和 MarginHorizontal）进行智能化适配。默认为空，表示不使用预设配置，支持的预设配置如下所示。
	// * small ：小字幕。
	// * medium：中字幕。
	// * large：大字幕。 :::tip 使用预设配置时，字幕行数、每行最大字符数、左右边距和底部边距参数不生效，系统将使用预设配置自动进行计算。 :::
	DisplayPreset string `json:"DisplayPreset"`

	// REQUIRED; 原文翻译成译文时使用的热词词库。
	GlossaryWordList []string `json:"GlossaryWordList"`

	// REQUIRED; 原文字幕识别时使用的热词词库。
	HotWordList []string `json:"HotWordList"`

	// REQUIRED; 设置在 16:9 分辨率场景下，每行字幕展示的最大字符数。 :::tip
	// * 使用预设配置时，字幕每行最大字符数设置不生效。
	// * 不使用预设配置时，字幕每行最大字符数必填。
	// * 每个文字、字母、符号或数字均为一个字符。
	// * 当屏幕分辨率改变时，屏幕上显示的每行文字数量会相应调整，以适应新的分辨率，确保文字的显示效果和阅读体验。 :::
	MaxCharNumber int32 `json:"MaxCharNumber"`

	// REQUIRED; 字幕展示的行数，同时适用于原文字幕和译文字幕，支持的取值及含义如下所示。
	// * 0：（默认值）根据字幕字数自动进行分行展示；
	// * 1：每种字幕展示一行；
	// * 2：每种字幕展示两行。 :::tip
	// * 使用预设配置时，字幕行数为自动分行展示。
	// * 超出行内字数限制时表示字幕将超过显示范围，此时字幕内容将被截断。 :::
	MaxRowNumber int32 `json:"MaxRowNumber"`

	// REQUIRED; 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
	// :::tip
	// * 使用预设配置时，字幕位置设置不生效。
	// * 不使用预设配置时，字幕位置设置必填。 :::
	Position ComponentsJ1MbxoSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesPosition `json:"Position"`

	// REQUIRED; 字幕配置的名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; 原文字幕展示参数配置。
	SourceLanguage Components1523StvSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguage `json:"SourceLanguage"`

	// REQUIRED; 译文字幕展示参数配置列表。
	TargetLanguage []Components1C398ShSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesTargetlanguageItems `json:"TargetLanguage"`
}

// ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePresetTargetLanguageItemBorder - 译文字幕的字体描边配置。
type ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePresetTargetLanguageItemBorder struct {

	// REQUIRED
	Color string `json:"Color"`

	// REQUIRED
	Width int32 `json:"Width"`
}

type ListVhostTransCodePresetBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type ListVhostTransCodePresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostTransCodePresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostTransCodePresetResResult          `json:"Result,omitempty"`
}

type ListVhostTransCodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                            `json:"Version"`
	Error     *ListVhostTransCodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                           `json:"RequestID,omitempty"`
}

type ListVhostTransCodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostTransCodePresetResResult struct {

	// REQUIRED; 全部转码配置列表。
	AllPresetList []ListVhostTransCodePresetResResultAllPresetListItem `json:"AllPresetList"`

	// REQUIRED; 使用内置参数的转码配置列表。
	CommonPresetList []ListVhostTransCodePresetResResultCommonPresetListItem `json:"CommonPresetList"`

	// REQUIRED; 使用自定义配置的转码配置列表。
	CustomizePresetList []ListVhostTransCodePresetResResultCustomizePresetListItem `json:"CustomizePresetList"`
}

type ListVhostTransCodePresetResResultAllPresetListItem struct {

	// REQUIRED; 直播流地址的 AppName 字段。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset struct {

	// 音频编码格式。包括以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应，按照Width和Height的取值进行拉伸；
	// * 1：开启宽高自适应，按照ShortSide或LongSide等比缩放。
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码不超过源流分辨率。开启后，当源流分辨率低于转码配置分辨率时(即源流宽低于转码配置宽且源流高低于转码配置高时)，将按源流视频分辨率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启不超过源流码率。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启不超过源流帧率。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 动态范围，画质增强类型生效
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效
	// * 0：不开启
	// * 1：开启
	FISwitch *int64 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为 s。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码模板参数的类型
	// * va：表示使用画质增强
	ParamType *string `json:"ParamType,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 是否极智超清转码。
	// * true：极智超清；
	// * false：标准转码。
	Roi *bool `json:"Roi,omitempty"`

	// 使用场景，画质增强时生效 football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为 s，取值范围为 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，默认为拉流转码，支持以下取值。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * h266：使用 H.266 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostTransCodePresetResResultCommonPresetListItem struct {

	// REQUIRED; 直播流地址的 AppName 字段。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset struct {

	// 音频编码格式，默认值为aac，支持的取值及含义如下所示。
	// * aac：使用 AAC 音频编码格式；
	// * opus：使用 Opus 音频编码格式。
	// * copy：不进行音频转码，所有音频编码参数不生效，音频编码参数包括音频码率（AudioBitrate）等。
	Acodec *string `json:"Acodec,omitempty"`

	// 视频分辨率自适应模式开关，默认值为0。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As取值为0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As取值为1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码视频分辨率不超过源流分辨率，默认值为1表示开启。开启后，当源流分辨率低于转码配置分辨率时（即源流宽低于转码配置宽且源流高低于转码配置高时），将按源流视频分辨率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启转码视频码率不超过源流码率，默认值为1表示开启。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启转码视频帧率不超过源流帧率，默认值为1表示开启。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，取值为0时表示去除 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。 :::tip
	// * 当关闭视频分辨率自适应（As取值为0）时，转码分辨率将取Width和Height的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As取值为0）时，Width和Height任一取值为0时，转码视频将保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip
	// * 当开启视频分辨率自适应模式时（As取值为1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式时（As取值为1）时，如果LongSide、ShortSide、Width、Height同时取0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 转码类型是否为极智超清转码，默认值为false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec 取值为 h266）时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码（TransType取值为Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为-1和 [0,300]，-1表示不停止转码，默认值为60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 视频编码格式；
	// * h265：使用 H.265 视频编码格式；
	// * h266：使用 H.266 视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP和BFrames等。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。 :::tip
	// * 当关闭视频分辨率自适应（As取值为0）时，转码分辨率将取Width和Height的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As取值为0）时，Width和Height任一取值为0时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostTransCodePresetResResultCustomizePresetListItem struct {

	// REQUIRED; 直播流地址的 AppName 字段。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset struct {

	// 音频编码格式。包括以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应，按照Width和Height的取值进行拉伸；
	// * 1：开启宽高自适应，按照ShortSide或LongSide等比缩放。
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码不超过源流分辨率。开启后，当源流分辨率低于转码配置分辨率时(即源流宽低于转码配置宽且源流高低于转码配置高时)，将按源流视频分辨率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启不超过源流码率。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启不超过源流帧率。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 动态范围，画质增强类型生效
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效
	// * 0：不开启
	// * 1：开启
	FISwitch *int64 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为 s。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码模板参数的类型
	// * va：表示使用画质增强
	ParamType *string `json:"ParamType,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 是否极智超清转码。
	// * true：极智超清；
	// * false：标准转码。
	Roi *bool `json:"Roi,omitempty"`

	// 使用场景，画质增强时生效 football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为 s，取值范围为 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，默认为拉流转码，支持以下取值。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * h266：使用 H.266 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostWatermarkPresetBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type ListVhostWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostWatermarkPresetResResult          `json:"Result,omitempty"`
}

type ListVhostWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                            `json:"Version"`
	Error   *ListVhostWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostWatermarkPresetResResult struct {

	// 统计消息，提供可用配置和不可用配置的数量。
	StaticsMsg *string `json:"StaticsMsg,omitempty"`

	// 不可正常使用的水印配置列表，如水印图片获取失败等原因导致的配置不可用。返回不可正常使用的水印配置信息及配置不可用的原因。
	WatermarkErrMsgList []*ListVhostWatermarkPresetResResultWatermarkErrMsgListItem `json:"WatermarkErrMsgList,omitempty"`

	// 可正常使用的水印配置列表。
	WatermarkPresetList []*ListVhostWatermarkPresetResResultWatermarkPresetListItem `json:"WatermarkPresetList,omitempty"`
}

type ListVhostWatermarkPresetResResultWatermarkErrMsgListItem struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 获取水印模板失败的具体错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty"`

	// 域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListVhostWatermarkPresetResResultWatermarkPresetListItem struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 水印模版 ID。
	ID *int32 `json:"ID,omitempty"`

	// 需要添加水印的直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片文件名。
	PictureKey *string `json:"PictureKey,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片字符串字段二选一传入，同时传入时，以水印图片字符串参数为准。
	PictureURL *string `json:"PictureURL,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListWatermarkPresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	// :::tip
	// * 默认为空，表示查询的 AppName 级别对所有转码流生效的配置。
	// * 指定流名称时，表示查询仅对 AppName 下指定流名称的转码流生效的配置。 :::
	Stream *string `json:"Stream,omitempty"`
}

type ListWatermarkPresetDetailBody struct {

	// 水印模板 ID 列表，默认为空，表示查询结果不对模板 ID 进行筛选。
	PresetIDList []*int64 `json:"PresetIDList,omitempty"`

	// 水印模板名称列表，默认为空，表示查询结果不对模板名称进行筛选。
	PresetNameList []*string `json:"PresetNameList,omitempty"`
}

type ListWatermarkPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata ListWatermarkPresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListWatermarkPresetDetailResResult `json:"Result"`
}

type ListWatermarkPresetDetailResResponseMetadata struct {

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

type ListWatermarkPresetDetailResResult struct {

	// 统计消息，提供可用模板和不可用模板的数量。
	StaticsMsg *string `json:"StaticsMsg,omitempty"`

	// 不可正常使用的水印模板，如水印图片获取失败等原因导致的不可用。
	WatermarkErrMsgList []*ListWatermarkPresetDetailResResultWatermarkErrMsgListItem `json:"WatermarkErrMsgList,omitempty"`

	// 可正常使用的水印模板列表。
	WatermarkPresetList []*ListWatermarkPresetDetailResResultWatermarkPresetListItem `json:"WatermarkPresetList,omitempty"`
}

type ListWatermarkPresetDetailResResultWatermarkErrMsgListItem struct {

	// 获取水印模板失败的具体错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty"`

	// 水印模板的 ID。
	ID *int32 `json:"ID,omitempty"`

	// 水印模板的名称。
	PresetName *string `json:"PresetName,omitempty"`
}

type ListWatermarkPresetDetailResResultWatermarkPresetListItem struct {

	// 水印模板的 ID。
	ID *int32 `json:"ID,omitempty"`

	// 水印模板的名称。
	Name *string `json:"Name,omitempty"`

	// 水印图片编码字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片的文件名。
	PictureKey *string `json:"PictureKey,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片字符串字段二选一传入，同时传入时，以水印图片字符串参数为准。
	PictureURL *string `json:"PictureURL,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *int32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *int32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`
}

type ListWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata ListWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListWatermarkPresetResResult          `json:"Result,omitempty"`
}

type ListWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                       `json:"Version"`
	Error     *ListWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                      `json:"RequestID,omitempty"`
}

type ListWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListWatermarkPresetResResult struct {

	// REQUIRED; 水印模板。
	Preset ListWatermarkPresetResResultPreset `json:"Preset"`
}

// ListWatermarkPresetResResultPreset - 水印模板。
type ListWatermarkPresetResResultPreset struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 水印配置的 ID。
	ID *int32 `json:"ID,omitempty"`

	// 需要添加水印的直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片文件名。
	PictureKey *string `json:"PictureKey,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片编码字符串字段二选一传入，同时传入时，以水印图片编码字符串参数为准。
	PictureURL *string `json:"PictureURL,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type RelaunchPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取状态为停用的任务 ID。
	TaskID string `json:"TaskId"`

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
}

type RelaunchPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata RelaunchPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type RelaunchPullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *RelaunchPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type RelaunchPullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type RestartTranscodingJobQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream" query:"Stream"`

	// REQUIRED; 转码配置的后缀，需去除转码后缀前的下划线（_）。如您配置的转码后缀为_hd，此处应传入hd。
	TranscodingTemplate string `json:"TranscodingTemplate" query:"TranscodingTemplate"`

	// 推流域名。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的推理域名。
	// :::tip Vhost 和 PushDomain 二选一必填。 :::
	PushDomain *string `json:"PushDomain,omitempty" query:"PushDomain"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	// :::tip
	// Vhost 和 PushDomain 二选一必填。 :::
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type RestartTranscodingJobRes struct {

	// REQUIRED
	ResponseMetadata RestartTranscodingJobResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type RestartTranscodingJobResResponseMetadata struct {

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

type ResumeStreamBody struct {

	// REQUIRED; 直播流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 直播流使用的流名称。
	Stream string `json:"Stream"`

	// 直播流使用的域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间。您可以调用 DescribeForbiddenStreamInfoByPage [https://www.volcengine.com/docs/6469/1126843] 接口，查看禁推直播流的信息，包括 Vhost、Domain、App
	// 和 Stream。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type ResumeStreamRes struct {

	// REQUIRED
	ResponseMetadata ResumeStreamResResponseMetadata `json:"ResponseMetadata"`
}

type ResumeStreamResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *ResumeStreamResResponseMetadataError `json:"Error,omitempty"`
}

type ResumeStreamResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type StopLivePadStreamBody struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type StopLivePadStreamRes struct {

	// REQUIRED
	ResponseMetadata StopLivePadStreamResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopLivePadStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type StopPullRecordTaskBody struct {

	// REQUIRED; 停止任务的TaskId
	TaskID string `json:"TaskId"`
}

type StopPullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata StopPullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopPullRecordTaskResResponseMetadata struct {

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

type StopPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取状态为未开始或生效中的任务 ID。
	TaskID string `json:"TaskId"`

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
}

type StopPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata StopPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type StopPullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *StopPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type StopPullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type TranscodingJobStatusQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream" query:"Stream"`

	// REQUIRED; 转码配置的后缀，需去除转码后缀前的下划线（_）。如您配置的转码后缀为_hd，此处应传入hd。
	TranscodingTemplate string `json:"TranscodingTemplate" query:"TranscodingTemplate"`

	// 推流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的推流域名。
	// :::tip Vhost 和 PushDomain 二选一必填。 :::
	PushDomain *string `json:"PushDomain,omitempty" query:"PushDomain"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	// :::tip Vhost
	// 和 PushDomain 二选一必填。 :::
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type TranscodingJobStatusRes struct {

	// REQUIRED
	ResponseMetadata TranscodingJobStatusResResponseMetadata `json:"ResponseMetadata"`
	Result           *TranscodingJobStatusResResult          `json:"Result,omitempty"`
}

type TranscodingJobStatusResResponseMetadata struct {

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

type TranscodingJobStatusResResult struct {

	// REQUIRED; 转码任务记录信息。
	Data []TranscodingJobStatusResResultDataItem `json:"Data"`
}

type TranscodingJobStatusResResultDataItem struct {

	// 本次转码任务的开始时间，Unix 时间戳，单位为秒。
	StartTime *int32 `json:"StartTime,omitempty"`

	// 本次转码任务的状态，取值及含义如下所示。
	// * init：转码任务初始化中；
	// * fetched：转码任务调度中；
	// * pending：转码任务即将开始；
	// * doing：转码任务进行中；
	// * prestop：转码任务即将结束；
	// * done：转码任务已结束。
	Status *string `json:"Status,omitempty"`
}

type UnBindEncryptDRMBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type UnBindEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata UnBindEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UnBindEncryptDRMResResponseMetadata struct {

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

type UnbindCertBody struct {

	// REQUIRED; 填写需要解绑 HTTPS 证书的域名。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要解绑证书的域名。
	Domain string `json:"Domain"`
}

type UnbindCertRes struct {

	// REQUIRED
	ResponseMetadata UnbindCertResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UnbindCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *UnbindCertResResponseMetadataError `json:"Error,omitempty"`
}

type UnbindCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateAuthKeyBody struct {

	// REQUIRED; 鉴权配置参数，包括鉴权密钥、鉴权参数、加密字符串生成算法等。
	AuthDetailList []UpdateAuthKeyBodyAuthDetailListItem `json:"AuthDetailList"`

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 鉴权场景类型，取值及含义如下所示。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// 是否开启 URL 地址鉴权，取值及含义如下所示。
	// * false：关闭（默认值）；
	// * true：开启。
	PushPullEnable *bool `json:"PushPullEnable,omitempty"`

	// 时间戳进制。默认值为 10，但当 AuthType 取值为 TypeC 时，默认值为 16。取值如下：
	// * 2：二进制
	// * 8：八进制
	// * 10：十进制
	// * 16：十六进制
	// :::tipSceneType 取值为 push 时，该参数取值固定为 10。 :::
	TimeStampBase *int32 `json:"TimeStampBase,omitempty"`

	// 鉴权生效时长，单位为秒，默认值为 0，取值范围为 [0,2592000]，超出生效时长后 URL 无法使用。
	ValidDuration *int32 `json:"ValidDuration,omitempty"`
}

type UpdateAuthKeyBodyAuthDetailListItem struct {

	// REQUIRED; 推/拉流鉴权时必选
	AuthType string `json:"AuthType"`

	// REQUIRED; 推/拉流鉴权时必选
	EncryptionAlgorithm string `json:"EncryptionAlgorithm"`

	// REQUIRED; 推/拉流鉴权时必选
	SecretKey string `json:"SecretKey"`

	// 推/拉流鉴权时必选
	AuthField map[string]*string `json:"AuthField,omitempty"`

	// 推/拉流鉴权时必选
	EncryptField []*string `json:"EncryptField,omitempty"`
}

type UpdateAuthKeyRes struct {

	// REQUIRED
	ResponseMetadata UpdateAuthKeyResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateAuthKeyResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *UpdateAuthKeyResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateAuthKeyResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateCallbackBody struct {

	// REQUIRED; 回调信息列表。
	CallbackDetailList []UpdateCallbackBodyCallbackDetailListItem `json:"CallbackDetailList"`

	// REQUIRED; 回调的消息类型，取值及含义如下所示。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制任务状态回调；
	// * audit_snapshot：截图审核结果回调。
	MessageType string `json:"MessageType"`

	// domain / app 二选一必传
	App *string `json:"App,omitempty"`

	// 回调消息发送是否开启鉴权，默认为 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	AuthEnable *bool `json:"AuthEnable,omitempty"`

	// 回调消息发送鉴权密钥。 :::tip 如果 AuthEnable 为 true，则密钥必填。 :::
	AuthKeyPrimary *string `json:"AuthKeyPrimary,omitempty"`

	// 直播流使用的推流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	// :::tipVhost和Domain传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：false，不开启；
	// * 1：true，开启。
	// :::tip 回调类型为推流开始或推流结束时生效。 :::
	TranscodeCallback *int32 `json:"TranscodeCallback,omitempty"`

	// domain / app 二选一必传
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateCallbackBodyCallbackDetailListItem struct {

	// REQUIRED; 回调类型，返回 HTTP，表示可以使用 HTTP 和 HTTPS 地址接收回调消息。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调消息接收地址。
	URL string `json:"URL"`
}

type UpdateCallbackRes struct {

	// REQUIRED
	ResponseMetadata UpdateCallbackResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateCallbackResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *UpdateCallbackResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateCallbackResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateCarouselTaskBody struct {

	// REQUIRED; 轮播规则，用于指定轮播播放的素材和行为等。
	Rule UpdateCarouselTaskBodyRule `json:"Rule"`

	// REQUIRED; 待更新的轮播任务 ID，任务的唯一标识。调用 CreateCarouselTask 接口创建轮播任务时返回。
	TaskID string `json:"TaskID"`
}

// UpdateCarouselTaskBodyRule - 轮播规则，用于指定轮播播放的素材和行为等。
type UpdateCarouselTaskBodyRule struct {

	// REQUIRED; 轮播素材列表，用于指定在轮播过程中播放的素材资源。
	Source []UpdateCarouselTaskBodyRuleSourceItem `json:"Source"`

	// 循环次数。取值范围为 [ ]，单位为，默认值为``。
	Loop *int32 `json:"Loop,omitempty"`

	// 对素材更新后的播放行为进行控制
	SeekInfo *UpdateCarouselTaskBodyRuleSeekInfo `json:"SeekInfo,omitempty"`
}

// UpdateCarouselTaskBodyRuleSeekInfo - 对素材更新后的播放行为进行控制
type UpdateCarouselTaskBodyRuleSeekInfo struct {

	// 0 表示推完当前播放的素材后再进行素材切换；1 表示立刻切换到指定的素材、指定的进度
	Immediate *int64 `json:"Immediate,omitempty"`

	// 更新后播放的素材ID，为空代表不指定。
	SourceID *string `json:"SourceID,omitempty"`

	// 切换素材后，素材播放的位置。
	SourceSeek *int64 `json:"SourceSeek,omitempty"`
}

type UpdateCarouselTaskBodyRuleSourceItem struct {

	// REQUIRED; 注意，如果ID相同，此结构的其余字段也需要保证相同
	ID string `json:"ID"`

	// REQUIRED; 轮播素材的文件类型，用于指定素材的文件来源类型。支持以下取值：
	// * vod：点播 MP4 或 FLV 文件；
	// * m3u8：点播 M3U8 文件。
	// :::tip 如果素材的 ID 没有变化（即更新的 ID 与原素材的 ID 相同），Type 取值要和元素材保持一致。 :::
	Type string `json:"Type"`

	// REQUIRED; 轮播素材的公网可访问地址。确保提供的地址能够被公网正常访问，以便正确加载轮播素材内容。 :::tip 如果素材的 ID 没有变化（即更新的 ID 与原素材的 ID 相同），Url 取值要和元素材保持一致。 :::
	URL string `json:"Url"`

	// 指定此素材连续播放的次数。该字段值必须大于等于 0，不传时，将保持原有轮播配置。支持的取值及含义如下：
	// * 0：不循环播放；
	// * 其他正整数：按照指定次数循环播放。
	Loop *int32 `json:"Loop,omitempty"`

	// 用于控制当前素材播放时跳过开头的一段时间，例如，跳过片头，单位为秒。该字段仅在素材类型为视频点播（type=vod）时有效。以下是该字段的使用规则：
	// * 如果 Seek 的取值小于等于 0 或大于视频的实际时长，则该字段不生效。
	// * 确保根据点播素材的实际长度设置合适的值，以实现跳过片头的效果。
	Seek *int32 `json:"Seek,omitempty"`
}

type UpdateCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata UpdateCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result UpdateCarouselTaskResResult `json:"Result"`
}

type UpdateCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type UpdateCarouselTaskResResult struct {

	// REQUIRED; 包含任务更新相关信息的数据对象。
	Data UpdateCarouselTaskResResultData `json:"Data"`
}

// UpdateCarouselTaskResResultData - 包含任务更新相关信息的数据对象。
type UpdateCarouselTaskResResultData struct {

	// REQUIRED; 当前生效的序列号
	OptID int32 `json:"OptID"`
}

type UpdateCloudMixTaskBody struct {

	// REQUIRED; 混流任务详细配置。
	MixedRules UpdateCloudMixTaskBodyMixedRules `json:"MixedRules"`

	// REQUIRED; 混流任务 ID，您可以通过 ListCloudMixTask [https://www.volcengine.com/docs/6469/1271157] 接口获取运行中的混流任务 ID。
	TaskID string `json:"TaskID"`
}

// UpdateCloudMixTaskBodyMixedRules - 混流任务详细配置。
type UpdateCloudMixTaskBodyMixedRules struct {

	// REQUIRED; 混流输出布局配置。
	InputLayout UpdateCloudMixTaskBodyMixedRulesInputLayout `json:"InputLayout"`

	// REQUIRED; 混流素材列表，最多支持配置 8 路输入源。
	InputSource []UpdateCloudMixTaskBodyMixedRulesInputSourceItem `json:"InputSource"`

	// REQUIRED; 混流输出流参数配置。 :::warning 更新云端混流任务时，Output 参数不支持变更，且必须传入与原混流任务一致的配置。 :::
	Output UpdateCloudMixTaskBodyMixedRulesOutput `json:"Output"`
}

// UpdateCloudMixTaskBodyMixedRulesInputLayout - 混流输出布局配置。
type UpdateCloudMixTaskBodyMixedRulesInputLayout struct {

	// REQUIRED; 混流输出画布配置和素材布局配置。
	Scene UpdateCloudMixTaskBodyMixedRulesInputLayoutScene `json:"Scene"`

	// 混流输出视频中 Logo 布局配置。 :::tip 支持最多配置 4 个 Logo，展示层级以添加顺序为准。 :::
	Logo []*UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItem `json:"Logo,omitempty"`
}

type UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItem struct {

	// Logo 图片在混流输出整体画面中的布局配置。
	Layout *UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout `json:"Layout,omitempty"`

	// Logo 图片访问地址。
	URL *string `json:"Url,omitempty"`
}

// UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout - Logo 图片在混流输出整体画面中的布局配置。
type UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout struct {

	// REQUIRED
	H int32 `json:"H"`

	// REQUIRED
	W int32 `json:"W"`

	// REQUIRED
	X int32 `json:"X"`

	// REQUIRED
	Y int32 `json:"Y"`
}

// UpdateCloudMixTaskBodyMixedRulesInputLayoutScene - 混流输出画布配置和素材布局配置。
type UpdateCloudMixTaskBodyMixedRulesInputLayoutScene struct {

	// REQUIRED; 混流输出整体画布高度，单位为 px，取值范围为 [10,2160]。
	Height int32 `json:"Height"`

	// REQUIRED; 混流素材在混流输出整体画面中的布局配置。 :::tip 混流素材布局中需包含所有素材的配置，且需与通过 Layer 参数与混流素材一一匹配。 :::
	Layout []UpdateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem `json:"Layout"`

	// REQUIRED; 混流输出画布整体宽度，单位为 px，取值范围为 [10,2160]。
	Width int32 `json:"Width"`
}

type UpdateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem struct {

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制高度，单位为 px，取值范围为 [10,2160]。
	// :::tip 限制宽度和限制高度指定了素材展示的限制范围，当素材尺寸和限制值不一致时，素材将在限制范围内根据长边进行等比缩放，其余区域透明展示。 :::
	H int32 `json:"H"`

	// REQUIRED; 当配置素材布局时需要通过 Layer 参数与素材进行一一对应。 :::tip 配置 Logo 图片的布局时此参数不生效。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制宽度，单位为 px，取值范围为 [10,2160]。
	W int32 `json:"W"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 X 偏移位置，单位为 px，取值范围为 0 到设置的画面宽度。
	X int32 `json:"X"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 Y 偏移位置，单位为 px，取值范围为 0 到设置的画面高度。
	Y int32 `json:"Y"`
}

type UpdateCloudMixTaskBodyMixedRulesInputSourceItem struct {

	// REQUIRED; 混流素材 ID，一个任务中素材 ID 不能重复，此 ID 用于任务状态回调消息中标识素材。
	ID string `json:"ID"`

	// REQUIRED; 混流素材的叠放顺序，1 为最底层，2 层在 1 层之上，以此类推，取值范围为[1,9999]。 :::tip 当前准备使用某个素材作为布局背景时，其叠放顺序应设置为所有素材中的最小值。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 混流素材类型，支持的取值及含义如下所示。
	// * vod：视频点播中的素材，支持 MP4、FLV 格式素材；
	// * live：直播源素材，支持 RTMP、FLV 协议拉流地址；
	// * pic：图片素材，支持 png、jpg 格式图片。
	Type string `json:"Type"`

	// REQUIRED; 混流素材访问地址。
	URL string `json:"Url"`
}

// UpdateCloudMixTaskBodyMixedRulesOutput - 混流输出流参数配置。 :::warning 更新云端混流任务时，Output 参数不支持变更，且必须传入与原混流任务一致的配置。 :::
type UpdateCloudMixTaskBodyMixedRulesOutput struct {

	// REQUIRED; 混流声音参数设置。
	Audio UpdateCloudMixTaskBodyMixedRulesOutputAudio `json:"Audio"`

	// REQUIRED; 混流视频的推流地址。
	URL []string `json:"Url"`

	// REQUIRED; 混流画面参数设置。
	Video UpdateCloudMixTaskBodyMixedRulesOutputVideo `json:"Video"`
}

// UpdateCloudMixTaskBodyMixedRulesOutputAudio - 混流声音参数设置。
type UpdateCloudMixTaskBodyMixedRulesOutputAudio struct {

	// REQUIRED; 混流输出流的音频码率，单位为 bps，取值范围为 [128000,320000]，常见取值及含义如下所示。
	// * 128000：128 kbps；
	// * 144000：144 kbps；
	// * 256000：256 kbps；
	// * 320000：320 kbps。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出流的音频声道设置，取值及含义如下所示。
	// * mono：单声道；
	// * stereo：立体声。
	ChannelLayout string `json:"ChannelLayout"`

	// REQUIRED; 混流输出流的音频采样率，单位为 HZ，常见取值及含义如下所示。
	// * 32000：32 kHZ；
	// * 44100：44.1 kHZ；
	// * 48000：48 kHZ。
	SampleRate int32 `json:"SampleRate"`
}

// UpdateCloudMixTaskBodyMixedRulesOutputVideo - 混流画面参数设置。
type UpdateCloudMixTaskBodyMixedRulesOutputVideo struct {

	// REQUIRED; 混流输出视频码率，单位为 bps，取值范围为 [1000000,20000000]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式。
	Codec string `json:"Codec"`

	// REQUIRED; 混流输出视频帧率，单位为 fps，取值范围为 [10,60]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	FrameRate int32 `json:"FrameRate"`

	// REQUIRED; IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,10]。
	GOP int32 `json:"GOP"`
}

type UpdateCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata UpdateCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result UpdateCloudMixTaskResResult `json:"Result"`
}

type UpdateCloudMixTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type UpdateCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data UpdateCloudMixTaskResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// UpdateCloudMixTaskResResultData - 返回的数据。
type UpdateCloudMixTaskResResultData struct {

	// REQUIRED; 任务版本标识符，用来标识更新后的任务版本。
	OptID int32 `json:"OptID"`
}

type UpdateDomainVhostBody struct {

	// REQUIRED; 待修改所属域名空间的的拉流/推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 目的域名空间，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名所属的域名空间信息。
	Vhost string `json:"Vhost"`
}

type UpdateDomainVhostRes struct {

	// REQUIRED
	ResponseMetadata UpdateDomainVhostResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateDomainVhostResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                     `json:"Version"`
	Error   *UpdateDomainVhostResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateDomainVhostResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateEncryptDRMBody struct {

	// DRM 证书管理平台 API 访问密钥，获取方法请参见最佳实践-直播 DRM 加密 [https://www.volcengine.com/docs/6469/1219836#在-intertrust-平台创建访问密钥]。
	APIKey *string `json:"APIKey,omitempty"`

	// 申请 FairPlay 证书过程中 Apple 返回的 ASk（Application Secret Key）字符串。
	ApplicationSecretKey *string `json:"ApplicationSecretKey,omitempty"`

	// FairPlay 证书文件内容。
	CertificateFile *string `json:"CertificateFile,omitempty"`

	// FairPlay 证书文件名称。
	CertificateFileName *string `json:"CertificateFileName,omitempty"`

	// 自定义 FairPlay 证书名称，支持由小写字母（a - z）、数字（0 - 9）和短横线（-）组成，最小长度为 2个字符，最大长度为 128 个字符。FairPlay 证书相关参数的获取方法请参见最佳实践-直播 DRM 加密 [https://www.volcengine.com/docs/6469/1219836#在-apple-官网获取-fairplay-证书]。
	CertificateName *string `json:"CertificateName,omitempty"`

	// 申请 FairPlay 证书时创建的私钥文件密钥。
	PrivateKey *string `json:"PrivateKey,omitempty"`

	// 申请 FairPlay 证书时创建的私钥文件内容。
	PrivateKeyFile *string `json:"PrivateKeyFile,omitempty"`

	// 申请 FairPlay 证书时创建的私钥文件名称。
	PrivateKeyFileName *string `json:"PrivateKeyFileName,omitempty"`
}

type UpdateEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata UpdateEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result interface{} `json:"Result"`
}

type UpdateEncryptDRMResResponseMetadata struct {

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

type UpdateEncryptHLSBody struct {

	// REQUIRED; 视频直播服务端生成密钥的更新周期，单位为秒，取值范围为 [60,604800]。
	CycleTime string `json:"CycleTime"`

	// REQUIRED; 客户自建密钥管理服务后，客户端向密钥管理服务请求获取密钥的地址。
	URL string `json:"URL"`
}

type UpdateEncryptHLSRes struct {

	// REQUIRED
	ResponseMetadata UpdateEncryptHLSResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateEncryptHLSResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateHTTPHeaderConfigBody struct {

	// REQUIRED; 配置完成后是否启用，取值及含义如下所示。
	// * true：启用；
	// * false：禁用。
	Enable bool `json:"Enable"`

	// REQUIRED; Header 具体字段配置。
	HeaderConfigList []UpdateHTTPHeaderConfigBodyHeaderConfigListItem `json:"HeaderConfigList"`

	// REQUIRED; 0: response 1: request
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 是否保留原 Header 配置，取值及含义如下所示。
	// * 0：保留（默认值）；
	// * 1：不保留。
	BlockOriginal *int32 `json:"BlockOriginal,omitempty"`

	// 拉流域名。默认为空，表示 Vhost 下的全部拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看配置的拉流域名。
	Domain *string `json:"Domain,omitempty"`
}

type UpdateHTTPHeaderConfigBodyHeaderConfigListItem struct {

	// REQUIRED; Header 配置中字段 Value 值的类型，取值及含义如下所示。
	// * 0：常量；
	// * 1：变量。
	HeaderFieldType int32 `json:"HeaderFieldType"`

	// Header 配置中字段的 Key 值，最大长度为 1024 个字符，多个 Header 不可重名。
	HeaderKey *string `json:"HeaderKey,omitempty"`

	// Header 配置中字段的 Value 值，最大长度为 1024 个字符，支持使用常量和变量作为 Value 值。
	// HTTP Header 类型为回源请求头时，支持使用以下变量为 Value 赋值：
	// * ${domain}：客户端拉流请求中使用的域名。
	// * ${uri}：客户端拉流请求中不包括查询参数的路径。如果请求被重写，则表示重写后的路径。
	// * ${args}：客户端拉流请求中的查询参数。如果请求被重写，则表示重写后的参数。
	// * ${remote_addr}：发送拉流请求的客户端 IP 地址。
	// * ${server_addr}：响应客户端拉流请求的 CDN 节点 IP 地址。
	// HTTP Header 类型为请求响应头时，支持使用以下变量为 Value 赋值：
	// * ${upstream_host}：客户端拉流请求中使用的域名。
	// * ${upstream_uri}：客户端拉流请求中不包括查询参数的路径。如果请求被重写，则表示重写后的路径。
	// * ${upstream_args}：客户端拉流请求中的查询参数。如果请求被重写，则表示重写后的参数。
	// * ${remote_addr}：发送拉流请求的客户端 IP 地址。
	HeaderValue *string `json:"HeaderValue,omitempty"`
}

type UpdateHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateHTTPHeaderConfigResResponseMetadata struct {

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

type UpdateIPAccessRuleBody struct {

	// REQUIRED; 推流域名或拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// IP 访问限制的域名。
	Domain string `json:"Domain"`

	// REQUIRED; IP 访问限制规则。
	IPAccessRule UpdateIPAccessRuleBodyIPAccessRule `json:"IPAccessRule"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// IP 访问限制的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

// UpdateIPAccessRuleBodyIPAccessRule - IP 访问限制规则。
type UpdateIPAccessRuleBodyIPAccessRule struct {

	// REQUIRED; 是否开启当前限制。取值如下：
	// * true: 开启。
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 名单中的 IP 信息。支持 IPv4 和 IPv6 格式的 IP 地址和 IP 网段。最多支持配置 500 个 IP 地址和网段。 例如，Type 取值为 deny、Domain 传入的是推流域名、该参数取值为 ["192.168.1.100","192.168.1.0/24","2001:db8:85a3::8a2e:370:7334","2001:db8::/32"]
	// 时，则表示不允许以下
	// IP 地址推流：
	// * IP 地址 192.168.1.100
	// * IP 地址 2001:db8:85a3::8a2e:370:7334
	// * 192.168.1.0 - 192.168.1.255范围内的所有 IP 地址
	// * 2001:db8:0000:0000:0000:0000:0000:0000- 2001:db8:ffff:ffff:ffff:ffff:ffff:ffff 范围内的所有 IP 地址
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型。取值如下：
	// * allow: 白名单。 * 如果 Domain 传入的是推流域名，则只有符合规则的 IP 地址才可以推流。
	// * 如果 Domain 传入的是拉流域名，则只有符合规则的 IP 地址才可以拉流。
	//
	//
	// * deny: 黑名单。 * 如果 Domain 传入的是推流域名，则符合规则的 IP 地址无法推流。
	// * 如果 Domain 传入的是拉流域名，则符合规则的 IP 地址无法拉流。
	Type string `json:"Type"`

	// 对于黑名单中或非白名单中的 IP 地址，系统默认返回 403 错误，表示禁止访问。 通过该参数，您可以自定义返回状态码，以便在查看日志时能快速区分禁止访问的原因，方便定位问题。 参数取值范围为 [200,999]。
	DenyHTTPStatusCode *int32 `json:"DenyHTTPStatusCode,omitempty"`

	// 校验是否包含 TS 文件。默认值为 false。取值如下：
	// * true: 包含。
	// * false: 不包含。
	// 使用 HLS 协议拉流时，直播服务器会将完整的直播流切割成小的 TS（Transport Stream）文件。观众观看直播时，播放器会不断请求 TS 文件并拼接成连续的画面。参数取值为 true 时，直播服务器会严格验证请求 TS 文件的
	// IP 地址，防止黑名单中或不在白名单中的 IP 地址访问 TS 文件。 :::tip
	// * 该配置仅对拉流域名生效。
	// * 如果无法确保能获取所有客户端请求 TS 文件的出口 IP 地址，建议将参数取值设置为 false，以免误拦截观众。 :::
	EnableTS *bool `json:"EnableTS,omitempty"`
}

type UpdateIPAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata UpdateIPAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateIPAccessRuleResResponseMetadata struct {

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

type UpdateLivePadPresetBody struct {

	// REQUIRED; 垫片时长。取值范围：>=1000。单位ms
	MaxDuration int64 `json:"MaxDuration"`

	// REQUIRED; 垫片类型，1: 图片、2: 视频、3: 源流最后一帧
	PadType int32 `json:"PadType"`

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`

	// REQUIRED; 断流等待时间。断流等待时间。 取值范围：0-6000。 单位：ms。
	WaitDuration int64 `json:"WaitDuration"`

	// 模板描述，长度上限：1024字节。
	Description *string `json:"Description,omitempty"`

	// 垫片素材地址。对 源流最后一帧类型无效。
	URL *string `json:"Url,omitempty"`
}

type UpdateLivePadPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateLivePadPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateLivePadPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdatePullToPushGroupBody struct {

	// REQUIRED; 拉流转推群组名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 接口获取群组名称。
	Name string `json:"Name"`

	// 任务群组的标签信息。
	Tags []*UpdatePullToPushGroupBodyTagsItem `json:"Tags,omitempty"`
}

type UpdatePullToPushGroupBodyTagsItem struct {

	// REQUIRED; 标签类型，支持以下取值。
	// * System：系统内置标签；
	// * Custom：自定义标签。
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type UpdatePullToPushGroupRes struct {

	// REQUIRED
	ResponseMetadata UpdatePullToPushGroupResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdatePullToPushGroupResResponseMetadata struct {

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

type UpdatePullToPushTaskBody struct {

	// REQUIRED; 任务等结束时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 任务的开始时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	StartTime int32 `json:"StartTime"`

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
	TaskID string `json:"TaskId"`

	// REQUIRED; 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type int32 `json:"Type"`

	// 推流应用名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	App *string `json:"App,omitempty"`

	// 接收拉流转推任务状态回调的地址，最大长度为 512 个字符。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频时为必选参数，参数取值及含义如下所示。
	// * -1：无限次循环，至任务结束；
	// * 0：有限次循环，循环次数以 PlayTimes 取值为准；
	// * >0：有限次循环，循环次数以 CycleMode 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 群组名称不支持更新，仅做校验使用。
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下表示不进行偏移。
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。 :::tip PlayTimes 为冗余参数，您可以将 PlayTimes 置 0 后直接使用 CycleMode 指定点播视频文件循环播放次数。
	// :::
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	// * 0：不开启；
	// * 1：开启（默认值）。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源（Type 为 0）时，为必选参数，最大长度为 1000 个字符。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频（Type 为 1）时，为必选参数，最多支持传入 30 个点播视频播放地址，每个地址最大长度为 1000 个字符。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	Stream *string `json:"Stream,omitempty"`

	// 拉流转推任务的名称，默认为空表示不配置任务名称。支持由中文、大小写字母（A - Z、a - z）和数字（0 - 9）组成，长度为 1 到 20 各字符。
	Title *string `json:"Title,omitempty"`

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*UpdatePullToPushTaskBodyVodSrcAddrsItem `json:"VodSrcAddrs,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *UpdatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
}

type UpdatePullToPushTaskBodyVodSrcAddrsItem struct {

	// REQUIRED; 点播文件地址。
	SrcAddr string `json:"SrcAddr"`

	// 当前点播文件结束播放的时间偏移值，单位为秒，默认为空时表示结束播放时间不进行偏移。
	EndOffset *float32 `json:"EndOffset,omitempty"`

	// 当前点播文件开始播放的时间偏移值，单位为秒。默认为空时表示开始播放时间不进行偏移。
	StartOffset *float32 `json:"StartOffset,omitempty"`
}

// UpdatePullToPushTaskBodyWatermark - 为拉流转推视频添加的水印配置信息。
type UpdatePullToPushTaskBodyWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	Picture string `json:"Picture"`

	// REQUIRED; 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
	Ratio float32 `json:"Ratio"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosX float32 `json:"RelativePosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosY float32 `json:"RelativePosY"`
}

type UpdatePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata UpdatePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type UpdatePullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *UpdatePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type UpdatePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRecordPresetV2Body struct {

	// REQUIRED; 录制配置的名称。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待更新录制配置的 Name 取值。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待更新录制配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，用来指定待更新的录制配置，默认为空。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]
	// 接口查看待更新录制配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 录制配置的详细参数配置。
	// :::tip 以下录制参数，未传入值时表示与更新前的配置相同。 :::
	RecordPresetConfig *UpdateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig,omitempty"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相应，用来指定待更新的录制配置，默认为空。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]
	// 接口查看待更新录制配置的 Stream 取值。
	Stream *string `json:"Stream,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfig - 录制配置的详细参数配置。
// :::tip 以下录制参数，未传入值时表示与更新前的配置相同。 :::
type UpdateRecordPresetV2BodyRecordPresetConfig struct {

	// 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	FlvParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	HlsParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	Mp4Param *UpdateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 是否录制源流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制为 HLS 格式时，单个 TS 切片时长，单位为秒，默认值为 10，取值范围为 [5,30]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 是否录制转码流，默认值为 0。支持的取值如下所示。
	// * 0：不录制；
	// * 1：录制全部转码流；
	// * 2：录制指定转码流，即通过转码后缀列表 TranscodeSuffixList匹配转码流进行录制，如果转码流后缀列表为空仍表示录制全部转码流。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码流后缀列表，是否录制转码设置为根据转码流列表匹配（TranscodeRecord 取值为 2）时生效，TranscodeSuffixList 默认配置为空，效果等同于录制全部转码流。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParam - 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

	// 实时录制场景下，断流等待时长，单位为秒，默认值为 180，取值范围为 [0,3600]。如果实际断流时间小于断流等待时长，录制任务不会停止；如果实际断流时间大于断流等待时长，录制任务会停止，断流恢复后重新开始一个新的录制任务。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制场景下，单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1 和 [300,86400]。
	// * 取值为 -1 时，表示不限制录制时长，录制结束后生成一个完整的录制文件。
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时长，到达时长立即生成录制文件，完成录制后一起上传。
	// :::tip
	// * 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。
	// * 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。 :::
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制场景下，单文件录制时长，单位为秒，默认值为 1800，取值范围为 [300,21600]。录制时间到达设置的单文件录制时长时，会立即生成录制文件实时上传存储。 :::tip 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。
	// :::
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流录制场景下，断流拼接时长，单位为秒，默认值为 0，支持的取值及含义如下所示。
	// * -1：一直拼接，表示每次断流都不会影响录制任务，录制完成后生成一个完整的录制文件；
	// * 0：不拼接，表示每次断流结束录制任务生成一个录制文件，断流恢复重新开始一个新的录制任务；
	// * 大于 0：拼接容错时间，表示如果断流时间小于拼接容错时间时，则录制任务不会停止，不会生成新的录制文件；如果断流时间大于拼接容错时间，则录制任务停止，断流恢复后重新开始一个新的录制任务。
	// :::tip 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。 :::
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam - TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam struct {

	// TOS 存储对应的 Bucket。例如，存储位置为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 如果使用 TOS 存储，即 TOSParam
	// 中 Enable 取值为 true 时，Bucket 为必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件存储到 TOS 时的存储路径和文件名规则。支持输入字母（A - Z、a - z）、数字（0 - 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。最大长度为 200 个字符，
	// 支持以下字段作为占位符：
	// * record：自定义字段，可遵照支持字符进行自定义。
	// * {PubDomain}：当前配置中的 vhost 值。
	// * {App}：当前配置中的 AppName 值。
	// * {Stream}：当前配置中的 StreamName 值。
	// * {StartTime}：录制开始的 Unix 时间戳，精度为 s。
	// * {EndTime}：录制结束的 Unix 时间戳，精度为 s。
	// 存储路径必须至少包含两级目录。例如：live/{App}/{Stream}
	// 合法示例：
	// record/{PubDomain}/{App}/{Stream}/{StartTime}-{EndTime}
	// {App}/archive/{Stream}/recording_{StartTime}
	// vod/{Stream}/!highlight_{EndTime}
	// a/b/custom_record
	// 错误示例：
	// single_level # 错误：路径层级不足两级
	// invalid_/{S@ream}/file # 错误：含非法字符@
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储对应 Bucket 下的存储目录，默认为空。例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam - VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息，默认为空。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储规则，最大长度为 200 个字符，支持以record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime} 样式设置存储规则，支持输入字母（A - Z、a - z）、数字（0 -
	// 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。
	// 存储规则设置注意事项如下。
	// * 目录层级至少包含2级及以上，如live/{App}/{Stream}。
	// * record 为自定义字段；
	// * {PubDomain} 取值为当前配置的 vhost 值；
	// * {App} 取值为当前配置的 AppName 值；
	// * {Stream} 取值为当前配置的 StreamName 值；
	// * {StartTime} 取值为录制的开始时间戳；
	// * {EndTime} 取值为录制的结束时间戳。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型，存储类型介绍请参考媒资存储管理 [https://www.volcengine.com/docs/4/73629#媒资存储]。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果使用 VOD 存储，即 VODParam 中 Enable 取值为 true 时，VodNamespace
	// 为必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 视频点播工作流模板 ID，对于存储在点播的录制文件，会使用该工作流模版对录制的视频进行处理，可登录视频点播控制台 [https://console.volcengine.com/vod/]获取工作流模板 ID，默认为空。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParam - 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigHlsParam struct {

	// 断流等待时长，取值范围[0, 3600]。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为 s，默认值为 7200，取值范围为 -1，[300,86400]，-1表示一直录制，目前只对 HLS 生效。
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]。
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam - TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket。 :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam - VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace
	// 必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigMp4Param - 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigMp4Param struct {

	// 断流等待时长，取值范围[0, 3600]。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为 s，默认值为 7200，取值范围为 -1，[300,86400]，-1表示一直录制，目前只对 HLS 生效。
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]。
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam - TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket。 :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam - VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject     *string `json:"ExactObject,omitempty"`
	ExactObjectName *string `json:"ExactObjectName,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace
	// 必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type UpdateRecordPresetV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateRecordPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRecordPresetV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                        `json:"Version"`
	Error   *UpdateRecordPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRecordPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRefererBody struct {

	// REQUIRED; Referer 防盗链规则列表。 :::tip
	// * 同一个 Vhost 下，默认支持配置不超过 100 个 Referer 规则，如需提升限额请创建工单 [https://console.volcengine.com/workorder/create?step=2&SubProductID=P00000076]获取技术支持；
	// * 单次请求最多支持配置 100 个 Referer 规则。 :::
	RefererInfoList []UpdateRefererBodyRefererInfoListItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 参数 Domain 和 App 传且仅传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	// :::tip 参数 Domain 和 App 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`
}

type UpdateRefererBodyRefererInfoListItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词默认取值为 referer。
	Key string `json:"Key"`

	// REQUIRED; Referer 字段规则类型，取值即含义如下所示。
	// * deny：拒绝，即黑名单；
	// * allow：通过，即白名单。
	Type string `json:"Type"`

	// Referer 字段规则的匹配优先级，默认为 0，取值范围为 [0,100]，数值越大，优先级越高。如果优先级相同，则越早加入列表的域名优先级越高。
	Priority *int32 `json:"Priority,omitempty"`

	// Referer 字段规则，即设置的黑名单或白名单的域名，最大长度限制 300 个字符。
	Value *string `json:"Value,omitempty"`
}

type UpdateRefererRes struct {

	// REQUIRED
	ResponseMetadata UpdateRefererResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRefererResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *UpdateRefererResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRefererResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRelaySourceV3Body struct {

	// REQUIRED; 回源组配置详情。
	GroupDetails []UpdateRelaySourceV3BodyGroupDetailsItem `json:"GroupDetails"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，即直播流地址的AppName字段取值，默认为空，表示为当前域名空间的全局播放触发回源配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty"`
}

type UpdateRelaySourceV3BodyGroupDetailsItem struct {

	// REQUIRED; 回源组名称。
	Group string `json:"Group"`

	// REQUIRED; 回源服务器配置列表。
	Servers []UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsItem `json:"Servers"`
}

type UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsItem struct {

	// REQUIRED; 直播源服务器的地址，支持填写回源服务的域名或 IP 地址。 :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	RelaySourceDomain string `json:"RelaySourceDomain"`

	// REQUIRED; volcengine可以传入rtmp/flv, byteplus可以传入rtmp\flv\dash\hls
	RelaySourceProtocol string `json:"RelaySourceProtocol"`

	// 自定义回源参数，缺省情况下为空。格式为"Key":"Value"，例如，"domain":"live.push.net"。
	RelaySourceParams map[string]*string `json:"RelaySourceParams,omitempty"`
}

type UpdateRelaySourceV3Res struct {

	// REQUIRED
	ResponseMetadata UpdateRelaySourceV3ResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateRelaySourceV3ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *UpdateRelaySourceV3ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRelaySourceV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRelaySourceV4Body struct {

	// REQUIRED; 应用名称，拉流域名，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待更新固定回源配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待更新固定回源配置的 Domain 取值。
	Domain string `json:"Domain"`

	// REQUIRED; 回源地址列表，支持 RTMP、FLV、HLS 和 SRT 回源协议。
	// :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 流名称，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待更新固定回源配置的 Domain 取值。
	Stream string `json:"Stream"`

	// 结束时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	EndTime *int32 `json:"EndTime,omitempty"`

	// 开始时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	StartTime *int32 `json:"StartTime,omitempty"`
}

type UpdateRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata UpdateRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                       `json:"Version"`
	Error   *UpdateRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateSnapshotAuditPresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostSnapshotAuditPreset [https://www.volcengine.com/docs/6469/1126870]接口，获取待删除截图配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 截图审核配置名称，您可以调用ListVhostSnapshotAuditPreset [https://www.volcengine.com/docs/6469/1126870]接口，获取待删除截图配置的 PresetName
	// 取值。
	PresetName string `json:"PresetName"`

	// ToS 存储对应的 Bucket。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图审核配置的描述。
	Description *string `json:"Description,omitempty"`

	// 推流域名，您可以调用ListVhostSnapshotAuditPreset [https://www.volcengine.com/docs/6469/1126870]接口，获取待删除截图配置的 Domain 取值。 :::tip 参数
	// Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 截图间隔时间，单位为秒，取值范围为 [0.1,10]，支持保留两位小数。
	Interval *float32 `json:"Interval,omitempty"`

	// 审核标签，缺省情况下取值为 301、302、302、305 和 306，支持的取值及含义如下。
	// * 301：涉黄；
	// * 302：涉敏1；
	// * 303：涉敏2；
	// * 304：广告；
	// * 305：引人不适；
	// * 306：违禁；
	// * 307：二维码；
	// * 308：诈骗；
	// * 309：不良画面；
	// * 310：未成年相关；
	// * 320：文字违规。
	Label []*string `json:"Label,omitempty"`

	// veimageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符，最大长度为 180 个字符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// ToS 存储对应 Bucket 下的存储目录，默认为空。 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`

	// 存储策略。支持以下取值。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储全部图片。
	StorageStrategy *int32 `json:"StorageStrategy,omitempty"`

	// 域名空间，您可以调用ListVhostSnapshotAuditPreset [https://www.volcengine.com/docs/6469/1126870]接口，获取待删除截图配置的 Vhost 取值。 :::tip 参数
	// Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *UpdateSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateSnapshotAuditPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateSnapshotPresetV2Body struct {

	// REQUIRED; 应用名称，您可以调用ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858]接口，获取待更新截图配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 截图配置的名称，您可以调用 ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858] 接口获取，取值与 Name 字段取值相同。
	Preset string `json:"Preset"`

	// REQUIRED; 截图配置的详细参数配置。
	SnapshotPresetConfig UpdateSnapshotPresetV2BodySnapshotPresetConfig `json:"SnapshotPresetConfig"`

	// REQUIRED; 域名空间，您可以调用 ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858] 接口，获取待更新截图配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 截图配置生效状态，默认及含义如下所示。
	// * 1：生效；
	// * 0：不生效。
	Status *int32 `json:"Status,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfig - 截图配置的详细参数配置。
type UpdateSnapshotPresetV2BodySnapshotPresetConfig struct {

	// 截图间隔时间，单位为秒，默认值为 10，取值范围为正整数。
	Interval *int32 `json:"Interval,omitempty"`

	// 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JPEGParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam `json:"JpegParam,omitempty"`

	// 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JpgParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParam `json:"JpgParam,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam - 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam struct {

	// 当前格式的截图配置是否开启，取值及含义如下所示。
	// * false：（默认值）不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	ImageXParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam `json:"ImageXParam,omitempty"`

	// 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	TOSParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam `json:"TOSParam,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam - 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam
// 配置且配置其中一个。 :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam struct {

	// 截图是否使用 veImageX 存储，取值及含义如下所示。
	// * false：（默认值）不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、*" 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// 使用 veImageX 存储截图时，对应的 veImageX 的服务 ID。 :::tip 使用 veImageX 存储时 ServiceID 为必填项。 :::
	ServiceID *string `json:"ServiceID,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam - 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。
// :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam struct {

	// TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 使用 TOS 存储时 Bucket
	// 为必填项。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图是否使用 TOS 存储，取值及含义如下所示。
	// * false：（默认值）不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// Bucket 目录，默认为空。 例如，存储路径为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParam - 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParam struct {
	Enable      *bool                                                              `json:"Enable,omitempty"`
	ImageXParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam `json:"ImageXParam,omitempty"`
	TOSParam    *UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam    `json:"TOSParam,omitempty"`
}

type UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam struct {
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	ServiceID       *string `json:"ServiceID,omitempty"`
}

type UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam struct {
	Bucket          *string `json:"Bucket,omitempty"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	StorageDir      *string `json:"StorageDir,omitempty"`
}

type UpdateSnapshotPresetV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSnapshotPresetV2ResResponseMetadata struct {

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

type UpdateStreamQuotaConfigBody struct {

	// REQUIRED; 需要配置限额的推流域名或拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 限额配置详情。
	QuotaDetailList []UpdateStreamQuotaConfigBodyQuotaDetailListItem `json:"QuotaDetailList"`
}

type UpdateStreamQuotaConfigBodyQuotaDetailListItem struct {

	// 拉流域名的带宽限额配置。 :::tipDomain 为拉流域名时，本参数为必选参数。 :::
	BandwidthConfig *UpdateStreamQuotaConfigBodyQuotaDetailListItemBandwidthConfig `json:"BandwidthConfig,omitempty"`

	// 超过限额时返回的错误码，默认值为403。
	ErrCode *int32 `json:"ErrCode,omitempty"`

	// 超过限额时返回的错误信息，默认值为forbid。
	ErrMsg *string `json:"ErrMsg,omitempty"`

	// 推流域名的推流路数限额配置。 :::tipDomain 为推流域名时，本参数为必选参数。 :::
	StreamConfig *UpdateStreamQuotaConfigBodyQuotaDetailListItemStreamConfig `json:"StreamConfig,omitempty"`
}

// UpdateStreamQuotaConfigBodyQuotaDetailListItemBandwidthConfig - 拉流域名的带宽限额配置。 :::tipDomain 为拉流域名时，本参数为必选参数。 :::
type UpdateStreamQuotaConfigBodyQuotaDetailListItemBandwidthConfig struct {

	// REQUIRED; 带宽限额，下行带宽峰值超过此限额的额外访问将会被拒绝，取值范围为 [1,10000]。
	Quota int32 `json:"Quota"`

	// REQUIRED; 拉流带宽限额的计量单位，支持的取值如下所示。
	// * Mbps
	// * Gbps
	// * Tbps
	QuotaUnit string `json:"QuotaUnit"`

	// 拉流带宽限额告警阈值，取值范围为 [1,10000]，缺省情况表示不设置告警。 :::tip 该参数的取值需要小于等于拉流带宽限额Quota，否则会报错。 :::
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`

	// 拉流带宽限额告警的计量单位，缺省情况表示不设置告警。支持的取值如下所示。
	// * Mbps
	// * Gbps
	// * Tbps
	AlarmThresholdUnit *string `json:"AlarmThresholdUnit,omitempty"`
}

// UpdateStreamQuotaConfigBodyQuotaDetailListItemStreamConfig - 推流域名的推流路数限额配置。 :::tipDomain 为推流域名时，本参数为必选参数。 :::
type UpdateStreamQuotaConfigBodyQuotaDetailListItemStreamConfig struct {

	// REQUIRED; 推流路数限额，取值[10~200000]。
	Quota int32 `json:"Quota"`

	// 推流路数限额告警阈值，缺省情况表示不设置告警。取值范围为 [10,200000]。 :::tip 该参数的取值需要小于等于推流路数限额Quota，否则会报错。 :::
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`
}

type UpdateStreamQuotaConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateStreamQuotaConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateStreamQuotaConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *UpdateStreamQuotaConfigResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateStreamQuotaConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateSubtitleTranscodePresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712]接口，获取待更新字幕配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 截图配置的名称，您可以调用ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712]接口，获取待更新字幕配置的 PresetName
	// 取值。
	PresetName string `json:"PresetName"`

	// REQUIRED; 原文字幕展示参数配置。
	SourceLanguage UpdateSubtitleTranscodePresetBodySourceLanguage `json:"SourceLanguage"`

	// REQUIRED; 关联转码配置后缀，一个字幕配置支持关联多个转码配置后缀。
	Suffixes []string `json:"Suffixes"`

	// REQUIRED; 域名空间，您可以调用 ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712] 接口，获取待更新字幕配置的 Vhost
	// 取值。
	Vhost string `json:"Vhost"`

	// 字幕配置的描述信息。
	Description *string `json:"Description,omitempty"`

	// 预设配置，使用预设配置是系统将自动对字体大小、字幕行数、每行最大字符数和边距参数（MarginVertical 和 MarginHorizontal）进行智能化适配。默认为空，表示不使用预设配置，支持的预设配置如下所示。
	// * small ：小字幕。
	// * medium：中字幕。
	// * large：大字幕。 :::tip 使用预设配置时，字幕行数、每行最大字符数、左右边距和底部边距参数不生效，系统将使用预设配置自动进行计算。 :::
	DisplayPreset *string `json:"DisplayPreset,omitempty"`

	// 原文翻译成译文时使用的热词词库，总长度不超过 10000 个字符，默认为空。
	GlossaryWordList []*string `json:"GlossaryWordList,omitempty"`

	// 原文字幕识别时使用的热词词库，总长度不超过为 10000 个字符，默认为空。
	HotWordList []*string `json:"HotWordList,omitempty"`

	// 设置在 16:9 分辨率场景下，每行字幕展示的最大字符数。 :::tip
	// * 使用预设配置时，字幕每行最大字符数设置不生效。
	// * 不使用预设配置时，字幕每行最大字符数必填。
	// * 每个文字、字母、符号或数字均为一个字符。
	// * 当屏幕分辨率改变时，屏幕上显示的每行文字数量会相应调整，以适应新的分辨率，确保文字的显示效果和阅读体验。 :::
	MaxCharNumber *int32 `json:"MaxCharNumber,omitempty"`

	// 字幕展示的行数，同时适用于原文字幕和译文字幕，支持的取值及含义如下所示。
	// * 0：（默认值）根据字幕字数自动进行分行展示；
	// * 1：每种字幕展示一行；
	// * 2：每种字幕展示两行。 :::tip
	// * 使用预设配置时，字幕行数为自动分行展示。
	// * 超出行内字数限制时表示字幕将超过显示范围，此时字幕内容将被截断。 :::
	MaxRowNumber *int32 `json:"MaxRowNumber,omitempty"`

	// 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
	// :::tip
	// * 使用预设配置时，字幕位置设置不生效。
	// * 不使用预设配置时，字幕位置设置必填。 :::
	Position *UpdateSubtitleTranscodePresetBodyPosition `json:"Position,omitempty"`

	// 译文字幕展示参数配置列表，当前最多支持配置一种译文。
	TargetLanguage []*UpdateSubtitleTranscodePresetBodyTargetLanguageItem `json:"TargetLanguage,omitempty"`
}

// UpdateSubtitleTranscodePresetBodyPosition - 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
// :::tip
// * 使用预设配置时，字幕位置设置不生效。
// * 不使用预设配置时，字幕位置设置必填。 :::
type UpdateSubtitleTranscodePresetBodyPosition struct {

	// 字幕距离画面两侧的边距与画面宽度的占比，使用归一化百分表示，取值范围为 [0,0.2]。
	MarginHorizontal *float32 `json:"MarginHorizontal,omitempty"`

	// 字幕距离画面底部的边距与画面高度的占比，使用归一化百分表示，取值范围为 [0,0.5]。
	MarginVertical *float32 `json:"MarginVertical,omitempty"`
}

// UpdateSubtitleTranscodePresetBodySourceLanguage - 原文字幕展示参数配置。
type UpdateSubtitleTranscodePresetBodySourceLanguage struct {

	// REQUIRED; 是否展示原文字幕，取值及含义如下所示。
	// * true：展示，此时将展示原文和译文双语字幕
	// * false：不展示，此时将只展示译文字幕。
	// :::tip 原文字幕语言和译文字幕语言相同时，仅展示译文字幕。 :::
	Display bool `json:"Display"`

	// REQUIRED; 原文字幕的字体，原文字幕字体根据原文字幕语言取值不同而不同，取值及含义如下所示。
	// * 当原文字幕的语言是 zh 时，支持以下字体取值。 * siyuanheiti：思源黑体；
	// * songtixi：宋体细；
	// * songticu：宋体粗；
	// * heitifan：黑体繁；
	// * kaiti：楷体。
	//
	//
	// * 当原文字幕的语言是 en 时，支持以下字体取值。 * inter：Inter；
	// * roboto：Roboto；
	// * opposans：OPPOSans；
	// * siyuansongti：思源宋体；
	// * montserrat：Montserrat。
	//
	//
	// * 当原文字幕的语言是 ko 和 ja 时，支持 notosans(Noto Sans) 字体。
	Font string `json:"Font"`

	// REQUIRED; 原文字幕的字体颜色，支持以下几种方法进行定义。
	// * 支持以 0x 或 # 开头，后面跟着十六进制颜色 RGB 值，再跟着 @+十六进制/百分比来表示的透明度值，来定义字幕的字体颜色。例如，设置 RGB 值为 FF0000，透明度为 5%的颜色时，您可以传入 0xFF0000@0x80、0xFF0000@0.5、#FF0000@0x80
	// 或 #FF0000@0.5。
	// * 支持使用前端框架 FFmpeg 规定的颜色关键字，来定义字幕的字体颜色。例如，AliceBlue 表示 0xF0F8FF、AntiqueWhite 表示 0xFAEBD7、Black 表示 0x000000 等。 :::tip 查看详细颜色定义方法及更多颜色关键字，请参考
	// FFmpeg 的颜色定义语法
	// [https://ffmpeg.org/ffmpeg-utils.html#color-syntax]。 :::
	FontColor string `json:"FontColor"`

	// REQUIRED; 原文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语。
	Language string `json:"Language"`

	// 原文字幕的阴影配置。
	Border *UpdateSubtitleTranscodePresetBodySourceLanguageBorder `json:"Border,omitempty"`
}

// UpdateSubtitleTranscodePresetBodySourceLanguageBorder - 原文字幕的阴影配置。
type UpdateSubtitleTranscodePresetBodySourceLanguageBorder struct {

	// REQUIRED; 描边的颜色，支持以下几种方法进行定义。
	// * 支持以 0x 或 # 开头，后面跟着十六进制颜色 RGB 值，再跟着 @+十六进制/百分比来表示的透明度值，来定义字幕的字体颜色。例如，设置 RGB 值为 FF0000，透明度为 5%的颜色时，您可以传入 0xFF0000@0x80、0xFF0000@0.5、#FF0000@0x80
	// 或 #FF0000@0.5。
	// * 支持使用前端框架 FFmpeg 规定的颜色关键字，来定义字幕的字体颜色。例如，AliceBlue 表示 0xF0F8FF、AntiqueWhite 表示 0xFAEBD7、Black 表示 0x000000 等。 :::tip 查看详细颜色定义方法及更多颜色关键字，请参考
	// FFmpeg 的颜色定义语法
	// [https://ffmpeg.org/ffmpeg-utils.html#color-syntax]。 :::
	Color string `json:"Color"`

	// 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Width *float32 `json:"Width,omitempty"`
}

type UpdateSubtitleTranscodePresetBodyTargetLanguageItem struct {

	// REQUIRED; 译文字幕的字体，译文字幕字体根据译文字幕语言取值不同而不同，取值及含义如下所示。
	// * 当译文字幕的语言是 zh 时，支持以下字体取值。 * siyuanheiti：思源黑体；
	// * songtixi：宋体细；
	// * songticu：宋体粗；
	// * heitifan：黑体繁；
	// * kaiti：楷体。
	//
	//
	// * 当译文字幕的语言是 zh-Hant 时，支持 siyuanheiti （思源黑体）字体。
	// * 当译文字幕的语言是 en 时，支持以下字体取值。 * inter：Inter；
	// * roboto：Roboto；
	// * opposans：OPPOSans；
	// * siyuansongti：思源宋体；
	// * montserrat：Montserrat。
	//
	//
	// * 当译文字幕的语言是 ko、ja、ar、de、es、fr、hi、pt、 ru、 vi、 th 时，支持 notosans(Noto Sans) 字体。
	Font string `json:"Font"`

	// REQUIRED; 译文字幕的字体颜色，支持以下几种方法进行定义。
	// * 支持以 0x 或 # 开头，后面跟着十六进制颜色 RGB 值，再跟着 @+十六进制/百分比来表示的透明度值，来定义字幕的字体颜色。例如，设置 RGB 值为 FF0000，透明度为 5%的颜色时，您可以传入 0xFF0000@0x80、0xFF0000@0.5、#FF0000@0x80
	// 或 #FF0000@0.5。
	// * 支持使用前端框架 FFmpeg 规定的颜色关键字，来定义字幕的字体颜色。例如，AliceBlue 表示 0xF0F8FF、AntiqueWhite 表示 0xFAEBD7、Black 表示 0x000000 等。 :::tip 查看详细颜色定义方法及更多颜色关键字，请参考
	// FFmpeg 的颜色定义语法
	// [https://ffmpeg.org/ffmpeg-utils.html#color-syntax]。 :::
	FontColor string `json:"FontColor"`

	// REQUIRED; 译文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * zh-Hant：繁体中文；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语；
	// * ar：阿拉伯语；
	// * de：德语；
	// * es：西班牙语；
	// * fr：法语；
	// * hi：印地语；
	// * pt：葡萄牙语；
	// * ru：俄语；
	// * vi：越南语；
	// * th：泰语。
	Language string `json:"Language"`

	// 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Border *UpdateSubtitleTranscodePresetBodyTargetLanguageItemBorder `json:"Border,omitempty"`
}

// UpdateSubtitleTranscodePresetBodyTargetLanguageItemBorder - 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
type UpdateSubtitleTranscodePresetBodyTargetLanguageItemBorder struct {

	// REQUIRED
	Color string   `json:"Color"`
	Width *float32 `json:"Width,omitempty"`
}

type UpdateSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSubtitleTranscodePresetResResponseMetadata struct {

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

type UpdateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待更新时移配置的App取值。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即允许用户回看的最长时间，单位为秒，支持的取值如下所示。
	// * 86400：1 天；
	// * 259200：3 天；
	// * 604800：7 天；
	// * 1296000：15 天。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 域名空间名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待更新时移配置的Vhost取值。
	Vhost string `json:"Vhost"`

	// 开启时移的流名称，默认为空表示更新 App 级别的时移配置，不为空时表示更新 Stream 级别的时移配置。您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待更新时移配置的Stream取值并进行更新。
	Stream *string `json:"Stream,omitempty"`
}

type UpdateTimeShiftPresetV3Res struct {

	// REQUIRED
	ResponseMetadata UpdateTimeShiftPresetV3ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateTimeShiftPresetV3ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *UpdateTimeShiftPresetV3ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateTimeShiftPresetV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateTranscodePresetBody struct {

	// REQUIRED; 转码配置的名称，您可以调用ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853]接口查看待更新转码配置的Preset取值。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待更新转码配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 音频编码格式，默认值为aac，支持的取值及含义如下所示。
	// * aac：使用 AAC 音频编码格式；
	// * opus：使用 Opus 音频编码格式。
	// * copy：不进行音频转码，所有音频编码参数不生效，音频编码参数包括音频码率（AudioBitrate）等。
	Acodec *string `json:"Acodec,omitempty"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待更新转码配置的
	// App 取值。
	App *string `json:"App,omitempty"`

	// 视频分辨率自适应模式开关，默认值为 0。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As 取值为 0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As 取值为 1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps，默认值为128，取值范围为 [0,1000]；取值为0时，表示与源流的音频码率相同。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码视频分辨率不超过源流分辨率，默认值为 1 表示开启。开启后，当源流分辨率低于转码配置分辨率时（即源流宽低于转码配置宽且源流高低于转码配置高时），将按源流视频分辨率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启转码视频码率不超过源流码率，默认值为 1 表示开启。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启转码视频帧率不超过源流帧率，默认值为 1 表示开启。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，默认值为 3，取值为 0 时表示去除 B 帧。
	// 最大 B 帧数量的取值范围根据视频编码格式（Vcodec）的不同有所差异，取值范围如下所示。
	// * 视频编码格式为 H.264 （Vcodec 取值为 h264）时取值范围为 [0,7]；
	// * 视频编码格式为 H.265 或 H.266 （Vcodec 取值为 h265 或 h266）时取值范围为 [0,3]、7、15。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，取值为 0 时表示与源流视频帧率相同。
	// 视频帧率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频帧率取值范围为 [0,60]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频帧率取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,30]。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度，默认值为 0。
	// 视频高度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频高度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,7680]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,1280]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。
	//
	//
	// :::tip
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,2160]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,720]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。 :::tip
	//
	//
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码（TransType 取值为 Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码后缀，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，默认值为 Pull，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 的视频编码格式；
	// * h265：使用 H.265 的视频编码格式；
	// * h266：使用 H.266 的视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP 和 BFrames
	// 等。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 bps，默认值为 1000000；取值为 0 时，表示与源流的视频码率相同。
	// 视频码率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频码率取值范围为 [0,30000000]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频码率取值范围为 [0,6000000]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度，单位为 px，默认值为 0。
	// 视频宽度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频宽度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type UpdateTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateTranscodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                         `json:"Version"`
	Error     *UpdateTranscodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                        `json:"RequestID,omitempty"`
}

type UpdateTranscodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateWatermarkPresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待更新水印配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 域名空间，您可以调用 ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889] 接口，查看待更新水印配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片使用 data URI 协议，格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片编码字符串字段二选一传入。同时传入时，以水印图片编码字符串参数为准。 :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。与水印相对宽度字段冲突，请选择其中一个传参。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。与水印相对高度字段冲突，请选择其中一个传参。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 流名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待更新水印配置的 Stream 取值。
	Stream *string `json:"Stream,omitempty"`
}

type UpdateWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *UpdateWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateWatermarkPresetV2Body struct {

	// 水印模板的 ID，您可以调用 ListWatermarkPresetDetail [https://www.volcengine.com/docs/6469/1323353] 接口获取。 :::tip PresetName 和 ID 二选一必填。
	// :::
	ID *int32 `json:"ID,omitempty"`

	// 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// :::tip Picture 与 PictureUrl 字段二选一必传，同时传入时，以 Picture 参数为准。 :::
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。 :::tip Picture 与 PictureUrl 字段二选一必传，同时传入时，以 Picture 参数为准。 :::
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印模板的名称，您可以调用 ListWatermarkPresetDetail [https://www.volcengine.com/docs/6469/1323353] 接口获取。 :::tip PresetName 和 ID 二选一必填。
	// :::
	PresetName *string `json:"PresetName,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。 :::tip RelativeWidth 与 RelativeHeight 二选一必传。 :::
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。 :::tip RelativeWidth 与 RelativeHeight 二选一必传。 :::
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`
}

type UpdateWatermarkPresetV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateWatermarkPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateWatermarkPresetV2ResResponseMetadata struct {

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
type BindCert struct{}
type BindCertQuery struct{}
type BindEncryptDRM struct{}
type BindEncryptDRMQuery struct{}
type ContinuePullToPushTask struct{}
type ContinuePullToPushTaskQuery struct{}
type CreateCarouselTask struct{}
type CreateCarouselTaskQuery struct{}
type CreateCert struct{}
type CreateCertQuery struct{}
type CreateCloudMixTask struct{}
type CreateCloudMixTaskQuery struct{}
type CreateDomain struct{}
type CreateDomainQuery struct{}
type CreateDomainV2 struct{}
type CreateDomainV2Query struct{}
type CreateHighLightTask struct{}
type CreateHighLightTaskQuery struct{}
type CreateLivePadPreset struct{}
type CreateLivePadPresetQuery struct{}
type CreateLiveStreamRecordIndexFiles struct{}
type CreateLiveStreamRecordIndexFilesQuery struct{}
type CreateLiveVideoQualityAnalysisTask struct{}
type CreateLiveVideoQualityAnalysisTaskQuery struct{}
type CreatePullRecordTask struct{}
type CreatePullRecordTaskQuery struct{}
type CreatePullToPushGroup struct{}
type CreatePullToPushGroupQuery struct{}
type CreatePullToPushTask struct{}
type CreatePullToPushTaskQuery struct{}
type CreateRecordPresetV2 struct{}
type CreateRecordPresetV2Query struct{}
type CreateRelaySourceV4 struct{}
type CreateRelaySourceV4Query struct{}
type CreateSnapshotAuditPreset struct{}
type CreateSnapshotAuditPresetQuery struct{}
type CreateSnapshotPresetV2 struct{}
type CreateSnapshotPresetV2Query struct{}
type CreateSubtitleTranscodePreset struct{}
type CreateSubtitleTranscodePresetQuery struct{}
type CreateTimeShiftPresetV3 struct{}
type CreateTimeShiftPresetV3Query struct{}
type CreateTranscodePreset struct{}
type CreateTranscodePresetQuery struct{}
type CreateWatermarkPreset struct{}
type CreateWatermarkPresetQuery struct{}
type CreateWatermarkPresetV2 struct{}
type CreateWatermarkPresetV2Query struct{}
type DeleteCallback struct{}
type DeleteCallbackQuery struct{}
type DeleteCarouselTask struct{}
type DeleteCarouselTaskQuery struct{}
type DeleteCert struct{}
type DeleteCertQuery struct{}
type DeleteCloudMixTask struct{}
type DeleteCloudMixTaskQuery struct{}
type DeleteDomain struct{}
type DeleteDomainQuery struct{}
type DeleteHTTPHeaderConfig struct{}
type DeleteHTTPHeaderConfigQuery struct{}
type DeleteIPAccessRule struct{}
type DeleteIPAccessRuleQuery struct{}
type DeleteLivePadPreset struct{}
type DeleteLivePadPresetQuery struct{}
type DeleteLiveVideoQualityAnalysisTask struct{}
type DeleteLiveVideoQualityAnalysisTaskQuery struct{}
type DeletePullToPushGroup struct{}
type DeletePullToPushGroupQuery struct{}
type DeletePullToPushTask struct{}
type DeletePullToPushTaskQuery struct{}
type DeleteRecordPreset struct{}
type DeleteRecordPresetQuery struct{}
type DeleteReferer struct{}
type DeleteRefererQuery struct{}
type DeleteRelaySourceV3 struct{}
type DeleteRelaySourceV3Query struct{}
type DeleteRelaySourceV4 struct{}
type DeleteRelaySourceV4Query struct{}
type DeleteSnapshotAuditPreset struct{}
type DeleteSnapshotAuditPresetQuery struct{}
type DeleteSnapshotPreset struct{}
type DeleteSnapshotPresetQuery struct{}
type DeleteStreamQuotaConfig struct{}
type DeleteStreamQuotaConfigQuery struct{}
type DeleteSubtitleTranscodePreset struct{}
type DeleteSubtitleTranscodePresetQuery struct{}
type DeleteTaskByAccountID struct{}
type DeleteTaskByAccountIDQuery struct{}
type DeleteTimeShiftPresetV3 struct{}
type DeleteTimeShiftPresetV3Query struct{}
type DeleteTranscodePreset struct{}
type DeleteTranscodePresetQuery struct{}
type DeleteWatermarkPreset struct{}
type DeleteWatermarkPresetQuery struct{}
type DeleteWatermarkPresetV2 struct{}
type DeleteWatermarkPresetV2Query struct{}
type DescribeAuth struct{}
type DescribeAuthQuery struct{}
type DescribeCDNSnapshotHistory struct{}
type DescribeCDNSnapshotHistoryQuery struct{}
type DescribeCallback struct{}
type DescribeCallbackQuery struct{}
type DescribeCertDRM struct{}
type DescribeCertDRMBody struct{}
type DescribeCertDetailSecretV2 struct{}
type DescribeCertDetailSecretV2Query struct{}
type DescribeClosedStreamInfoByPage struct{}
type DescribeClosedStreamInfoByPageBody struct{}
type DescribeDomain struct{}
type DescribeDomainQuery struct{}
type DescribeEncryptDRM struct{}
type DescribeEncryptDRMBody struct{}
type DescribeEncryptDRMQuery struct{}
type DescribeEncryptHLS struct{}
type DescribeEncryptHLSBody struct{}
type DescribeEncryptHLSQuery struct{}
type DescribeForbiddenStreamInfoByPage struct{}
type DescribeForbiddenStreamInfoByPageBody struct{}
type DescribeHTTPHeaderConfig struct{}
type DescribeHTTPHeaderConfigQuery struct{}
type DescribeHighLightTaskByAccountID struct{}
type DescribeHighLightTaskByAccountIDQuery struct{}
type DescribeIPAccessRule struct{}
type DescribeIPAccessRuleQuery struct{}
type DescribeIPInfo struct{}
type DescribeIPInfoQuery struct{}
type DescribeLicenseDRM struct{}
type DescribeLicenseDRMBody struct{}
type DescribeLiveAuditData struct{}
type DescribeLiveAuditDataQuery struct{}
type DescribeLiveBandwidthData struct{}
type DescribeLiveBandwidthDataQuery struct{}
type DescribeLiveBatchPushStreamAvgMetrics struct{}
type DescribeLiveBatchPushStreamAvgMetricsQuery struct{}
type DescribeLiveBatchPushStreamMetrics struct{}
type DescribeLiveBatchPushStreamMetricsQuery struct{}
type DescribeLiveBatchSourceStreamAvgMetrics struct{}
type DescribeLiveBatchSourceStreamAvgMetricsQuery struct{}
type DescribeLiveBatchSourceStreamMetrics struct{}
type DescribeLiveBatchSourceStreamMetricsQuery struct{}
type DescribeLiveBatchStreamSessionData struct{}
type DescribeLiveBatchStreamSessionDataQuery struct{}
type DescribeLiveBatchStreamTrafficData struct{}
type DescribeLiveBatchStreamTrafficDataQuery struct{}
type DescribeLiveBatchStreamTranscodeData struct{}
type DescribeLiveBatchStreamTranscodeDataQuery struct{}
type DescribeLiveCallbackData struct{}
type DescribeLiveCallbackDataQuery struct{}
type DescribeLiveEdgeStatData struct{}
type DescribeLiveEdgeStatDataQuery struct{}
type DescribeLiveISPData struct{}
type DescribeLiveISPDataBody struct{}
type DescribeLiveISPDataQuery struct{}
type DescribeLiveLogData struct{}
type DescribeLiveLogDataQuery struct{}
type DescribeLiveMetricBandwidthData struct{}
type DescribeLiveMetricBandwidthDataQuery struct{}
type DescribeLiveMetricTrafficData struct{}
type DescribeLiveMetricTrafficDataQuery struct{}
type DescribeLiveP95PeakBandwidthData struct{}
type DescribeLiveP95PeakBandwidthDataQuery struct{}
type DescribeLivePadPresetDetail struct{}
type DescribeLivePadPresetDetailQuery struct{}
type DescribeLivePadStreamList struct{}
type DescribeLivePadStreamListQuery struct{}
type DescribeLivePlayStatusCodeData struct{}
type DescribeLivePlayStatusCodeDataQuery struct{}
type DescribeLivePullToPushBandwidthData struct{}
type DescribeLivePullToPushBandwidthDataQuery struct{}
type DescribeLivePullToPushData struct{}
type DescribeLivePullToPushDataQuery struct{}
type DescribeLivePushStreamCountData struct{}
type DescribeLivePushStreamCountDataQuery struct{}
type DescribeLivePushStreamInfoData struct{}
type DescribeLivePushStreamInfoDataQuery struct{}
type DescribeLivePushStreamMetrics struct{}
type DescribeLivePushStreamMetricsQuery struct{}
type DescribeLiveRecordData struct{}
type DescribeLiveRecordDataQuery struct{}
type DescribeLiveRegionData struct{}
type DescribeLiveRegionDataBody struct{}
type DescribeLiveRegionDataQuery struct{}
type DescribeLiveSnapshotData struct{}
type DescribeLiveSnapshotDataQuery struct{}
type DescribeLiveSourceBandwidthData struct{}
type DescribeLiveSourceBandwidthDataQuery struct{}
type DescribeLiveSourceStreamMetrics struct{}
type DescribeLiveSourceStreamMetricsQuery struct{}
type DescribeLiveSourceTrafficData struct{}
type DescribeLiveSourceTrafficDataQuery struct{}
type DescribeLiveStreamCountData struct{}
type DescribeLiveStreamCountDataQuery struct{}
type DescribeLiveStreamInfoByPage struct{}
type DescribeLiveStreamInfoByPageBody struct{}
type DescribeLiveStreamSessionData struct{}
type DescribeLiveStreamSessionDataQuery struct{}
type DescribeLiveStreamState struct{}
type DescribeLiveStreamStateBody struct{}
type DescribeLiveTimeShiftData struct{}
type DescribeLiveTimeShiftDataQuery struct{}
type DescribeLiveTopPlayData struct{}
type DescribeLiveTopPlayDataQuery struct{}
type DescribeLiveTrafficData struct{}
type DescribeLiveTrafficDataQuery struct{}
type DescribeLiveTranscodeData struct{}
type DescribeLiveTranscodeDataQuery struct{}
type DescribeLiveTranscodeInfoData struct{}
type DescribeLiveTranscodeInfoDataQuery struct{}
type DescribeRecordTaskFileHistory struct{}
type DescribeRecordTaskFileHistoryQuery struct{}
type DescribeReferer struct{}
type DescribeRefererQuery struct{}
type DescribeRelaySourceV3 struct{}
type DescribeRelaySourceV3Query struct{}
type DescribeStreamQuotaConfig struct{}
type DescribeStreamQuotaConfigQuery struct{}
type DisableDomain struct{}
type DisableDomainQuery struct{}
type EnableDomain struct{}
type EnableDomainQuery struct{}
type EnableHTTPHeaderConfig struct{}
type EnableHTTPHeaderConfigQuery struct{}
type ForbidStream struct{}
type ForbidStreamQuery struct{}
type GeneratePlayURL struct{}
type GeneratePlayURLQuery struct{}
type GeneratePushURL struct{}
type GeneratePushURLQuery struct{}
type GetCarouselDetail struct{}
type GetCarouselDetailQuery struct{}
type GetCloudMixTaskDetail struct{}
type GetCloudMixTaskDetailQuery struct{}
type GetHLSEncryptDataKey struct{}
type GetHLSEncryptDataKeyBody struct{}
type GetLiveVideoQualityAnalysisTaskDetail struct{}
type GetLiveVideoQualityAnalysisTaskDetailQuery struct{}
type GetPullRecordTask struct{}
type GetPullRecordTaskQuery struct{}
type KillStream struct{}
type KillStreamQuery struct{}
type ListBindEncryptDRM struct{}
type ListBindEncryptDRMQuery struct{}
type ListCarouselTask struct{}
type ListCarouselTaskQuery struct{}
type ListCertV2 struct{}
type ListCertV2Query struct{}
type ListCloudMixTask struct{}
type ListCloudMixTaskQuery struct{}
type ListCommonTransPresetDetail struct{}
type ListCommonTransPresetDetailQuery struct{}
type ListDomainDetail struct{}
type ListDomainDetailQuery struct{}
type ListHighLightTask struct{}
type ListHighLightTaskQuery struct{}
type ListLiveVideoQualityAnalysisTasks struct{}
type ListLiveVideoQualityAnalysisTasksQuery struct{}
type ListPullRecordTask struct{}
type ListPullRecordTaskQuery struct{}
type ListPullToPushGroup struct{}
type ListPullToPushGroupQuery struct{}
type ListPullToPushTask struct{}
type ListPullToPushTaskBody struct{}
type ListPullToPushTaskV2 struct{}
type ListPullToPushTaskV2Query struct{}
type ListRelaySourceV4 struct{}
type ListRelaySourceV4Query struct{}
type ListTimeShiftPresetV2 struct{}
type ListTimeShiftPresetV2Query struct{}
type ListVhostRecordPresetV2 struct{}
type ListVhostRecordPresetV2Query struct{}
type ListVhostSnapshotAuditPreset struct{}
type ListVhostSnapshotAuditPresetQuery struct{}
type ListVhostSnapshotPresetV2 struct{}
type ListVhostSnapshotPresetV2Query struct{}
type ListVhostSubtitleTranscodePreset struct{}
type ListVhostSubtitleTranscodePresetQuery struct{}
type ListVhostTransCodePreset struct{}
type ListVhostTransCodePresetQuery struct{}
type ListVhostWatermarkPreset struct{}
type ListVhostWatermarkPresetQuery struct{}
type ListWatermarkPreset struct{}
type ListWatermarkPresetDetail struct{}
type ListWatermarkPresetDetailQuery struct{}
type ListWatermarkPresetQuery struct{}
type RelaunchPullToPushTask struct{}
type RelaunchPullToPushTaskQuery struct{}
type RestartTranscodingJob struct{}
type RestartTranscodingJobBody struct{}
type ResumeStream struct{}
type ResumeStreamQuery struct{}
type StopLivePadStream struct{}
type StopLivePadStreamQuery struct{}
type StopPullRecordTask struct{}
type StopPullRecordTaskQuery struct{}
type StopPullToPushTask struct{}
type StopPullToPushTaskQuery struct{}
type TranscodingJobStatus struct{}
type TranscodingJobStatusBody struct{}
type UnBindEncryptDRM struct{}
type UnBindEncryptDRMQuery struct{}
type UnbindCert struct{}
type UnbindCertQuery struct{}
type UpdateAuthKey struct{}
type UpdateAuthKeyQuery struct{}
type UpdateCallback struct{}
type UpdateCallbackQuery struct{}
type UpdateCarouselTask struct{}
type UpdateCarouselTaskQuery struct{}
type UpdateCloudMixTask struct{}
type UpdateCloudMixTaskQuery struct{}
type UpdateDomainVhost struct{}
type UpdateDomainVhostQuery struct{}
type UpdateEncryptDRM struct{}
type UpdateEncryptDRMQuery struct{}
type UpdateEncryptHLS struct{}
type UpdateEncryptHLSQuery struct{}
type UpdateHTTPHeaderConfig struct{}
type UpdateHTTPHeaderConfigQuery struct{}
type UpdateIPAccessRule struct{}
type UpdateIPAccessRuleQuery struct{}
type UpdateLivePadPreset struct{}
type UpdateLivePadPresetQuery struct{}
type UpdatePullToPushGroup struct{}
type UpdatePullToPushGroupQuery struct{}
type UpdatePullToPushTask struct{}
type UpdatePullToPushTaskQuery struct{}
type UpdateRecordPresetV2 struct{}
type UpdateRecordPresetV2Query struct{}
type UpdateReferer struct{}
type UpdateRefererQuery struct{}
type UpdateRelaySourceV3 struct{}
type UpdateRelaySourceV3Query struct{}
type UpdateRelaySourceV4 struct{}
type UpdateRelaySourceV4Query struct{}
type UpdateSnapshotAuditPreset struct{}
type UpdateSnapshotAuditPresetQuery struct{}
type UpdateSnapshotPresetV2 struct{}
type UpdateSnapshotPresetV2Query struct{}
type UpdateStreamQuotaConfig struct{}
type UpdateStreamQuotaConfigQuery struct{}
type UpdateSubtitleTranscodePreset struct{}
type UpdateSubtitleTranscodePresetQuery struct{}
type UpdateTimeShiftPresetV3 struct{}
type UpdateTimeShiftPresetV3Query struct{}
type UpdateTranscodePreset struct{}
type UpdateTranscodePresetQuery struct{}
type UpdateWatermarkPreset struct{}
type UpdateWatermarkPresetQuery struct{}
type UpdateWatermarkPresetV2 struct{}
type UpdateWatermarkPresetV2Query struct{}
type BindCertReq struct {
	*BindCertQuery
	*BindCertBody
}
type BindEncryptDRMReq struct {
	*BindEncryptDRMQuery
	*BindEncryptDRMBody
}
type ContinuePullToPushTaskReq struct {
	*ContinuePullToPushTaskQuery
	*ContinuePullToPushTaskBody
}
type CreateCarouselTaskReq struct {
	*CreateCarouselTaskQuery
	*CreateCarouselTaskBody
}
type CreateCertReq struct {
	*CreateCertQuery
	*CreateCertBody
}
type CreateCloudMixTaskReq struct {
	*CreateCloudMixTaskQuery
	*CreateCloudMixTaskBody
}
type CreateDomainReq struct {
	*CreateDomainQuery
	*CreateDomainBody
}
type CreateDomainV2Req struct {
	*CreateDomainV2Query
	*CreateDomainV2Body
}
type CreateHighLightTaskReq struct {
	*CreateHighLightTaskQuery
	*CreateHighLightTaskBody
}
type CreateLivePadPresetReq struct {
	*CreateLivePadPresetQuery
	*CreateLivePadPresetBody
}
type CreateLiveStreamRecordIndexFilesReq struct {
	*CreateLiveStreamRecordIndexFilesQuery
	*CreateLiveStreamRecordIndexFilesBody
}
type CreateLiveVideoQualityAnalysisTaskReq struct {
	*CreateLiveVideoQualityAnalysisTaskQuery
	*CreateLiveVideoQualityAnalysisTaskBody
}
type CreatePullRecordTaskReq struct {
	*CreatePullRecordTaskQuery
	*CreatePullRecordTaskBody
}
type CreatePullToPushGroupReq struct {
	*CreatePullToPushGroupQuery
	*CreatePullToPushGroupBody
}
type CreatePullToPushTaskReq struct {
	*CreatePullToPushTaskQuery
	*CreatePullToPushTaskBody
}
type CreateRecordPresetV2Req struct {
	*CreateRecordPresetV2Query
	*CreateRecordPresetV2Body
}
type CreateRelaySourceV4Req struct {
	*CreateRelaySourceV4Query
	*CreateRelaySourceV4Body
}
type CreateSnapshotAuditPresetReq struct {
	*CreateSnapshotAuditPresetQuery
	*CreateSnapshotAuditPresetBody
}
type CreateSnapshotPresetV2Req struct {
	*CreateSnapshotPresetV2Query
	*CreateSnapshotPresetV2Body
}
type CreateSubtitleTranscodePresetReq struct {
	*CreateSubtitleTranscodePresetQuery
	*CreateSubtitleTranscodePresetBody
}
type CreateTimeShiftPresetV3Req struct {
	*CreateTimeShiftPresetV3Query
	*CreateTimeShiftPresetV3Body
}
type CreateTranscodePresetReq struct {
	*CreateTranscodePresetQuery
	*CreateTranscodePresetBody
}
type CreateWatermarkPresetReq struct {
	*CreateWatermarkPresetQuery
	*CreateWatermarkPresetBody
}
type CreateWatermarkPresetV2Req struct {
	*CreateWatermarkPresetV2Query
	*CreateWatermarkPresetV2Body
}
type DeleteCallbackReq struct {
	*DeleteCallbackQuery
	*DeleteCallbackBody
}
type DeleteCarouselTaskReq struct {
	*DeleteCarouselTaskQuery
	*DeleteCarouselTaskBody
}
type DeleteCertReq struct {
	*DeleteCertQuery
	*DeleteCertBody
}
type DeleteCloudMixTaskReq struct {
	*DeleteCloudMixTaskQuery
	*DeleteCloudMixTaskBody
}
type DeleteDomainReq struct {
	*DeleteDomainQuery
	*DeleteDomainBody
}
type DeleteHTTPHeaderConfigReq struct {
	*DeleteHTTPHeaderConfigQuery
	*DeleteHTTPHeaderConfigBody
}
type DeleteIPAccessRuleReq struct {
	*DeleteIPAccessRuleQuery
	*DeleteIPAccessRuleBody
}
type DeleteLivePadPresetReq struct {
	*DeleteLivePadPresetQuery
	*DeleteLivePadPresetBody
}
type DeleteLiveVideoQualityAnalysisTaskReq struct {
	*DeleteLiveVideoQualityAnalysisTaskQuery
	*DeleteLiveVideoQualityAnalysisTaskBody
}
type DeletePullToPushGroupReq struct {
	*DeletePullToPushGroupQuery
	*DeletePullToPushGroupBody
}
type DeletePullToPushTaskReq struct {
	*DeletePullToPushTaskQuery
	*DeletePullToPushTaskBody
}
type DeleteRecordPresetReq struct {
	*DeleteRecordPresetQuery
	*DeleteRecordPresetBody
}
type DeleteRefererReq struct {
	*DeleteRefererQuery
	*DeleteRefererBody
}
type DeleteRelaySourceV3Req struct {
	*DeleteRelaySourceV3Query
	*DeleteRelaySourceV3Body
}
type DeleteRelaySourceV4Req struct {
	*DeleteRelaySourceV4Query
	*DeleteRelaySourceV4Body
}
type DeleteSnapshotAuditPresetReq struct {
	*DeleteSnapshotAuditPresetQuery
	*DeleteSnapshotAuditPresetBody
}
type DeleteSnapshotPresetReq struct {
	*DeleteSnapshotPresetQuery
	*DeleteSnapshotPresetBody
}
type DeleteStreamQuotaConfigReq struct {
	*DeleteStreamQuotaConfigQuery
	*DeleteStreamQuotaConfigBody
}
type DeleteSubtitleTranscodePresetReq struct {
	*DeleteSubtitleTranscodePresetQuery
	*DeleteSubtitleTranscodePresetBody
}
type DeleteTaskByAccountIDReq struct {
	*DeleteTaskByAccountIDQuery
	*DeleteTaskByAccountIDBody
}
type DeleteTimeShiftPresetV3Req struct {
	*DeleteTimeShiftPresetV3Query
	*DeleteTimeShiftPresetV3Body
}
type DeleteTranscodePresetReq struct {
	*DeleteTranscodePresetQuery
	*DeleteTranscodePresetBody
}
type DeleteWatermarkPresetReq struct {
	*DeleteWatermarkPresetQuery
	*DeleteWatermarkPresetBody
}
type DeleteWatermarkPresetV2Req struct {
	*DeleteWatermarkPresetV2Query
	*DeleteWatermarkPresetV2Body
}
type DescribeAuthReq struct {
	*DescribeAuthQuery
	*DescribeAuthBody
}
type DescribeCDNSnapshotHistoryReq struct {
	*DescribeCDNSnapshotHistoryQuery
	*DescribeCDNSnapshotHistoryBody
}
type DescribeCallbackReq struct {
	*DescribeCallbackQuery
	*DescribeCallbackBody
}
type DescribeCertDRMReq struct {
	*DescribeCertDRMQuery
	*DescribeCertDRMBody
}
type DescribeCertDetailSecretV2Req struct {
	*DescribeCertDetailSecretV2Query
	*DescribeCertDetailSecretV2Body
}
type DescribeClosedStreamInfoByPageReq struct {
	*DescribeClosedStreamInfoByPageQuery
	*DescribeClosedStreamInfoByPageBody
}
type DescribeDomainReq struct {
	*DescribeDomainQuery
	*DescribeDomainBody
}
type DescribeEncryptDRMReq struct {
	*DescribeEncryptDRMQuery
	*DescribeEncryptDRMBody
}
type DescribeEncryptHLSReq struct {
	*DescribeEncryptHLSQuery
	*DescribeEncryptHLSBody
}
type DescribeForbiddenStreamInfoByPageReq struct {
	*DescribeForbiddenStreamInfoByPageQuery
	*DescribeForbiddenStreamInfoByPageBody
}
type DescribeHTTPHeaderConfigReq struct {
	*DescribeHTTPHeaderConfigQuery
	*DescribeHTTPHeaderConfigBody
}
type DescribeHighLightTaskByAccountIDReq struct {
	*DescribeHighLightTaskByAccountIDQuery
	*DescribeHighLightTaskByAccountIDBody
}
type DescribeIPAccessRuleReq struct {
	*DescribeIPAccessRuleQuery
	*DescribeIPAccessRuleBody
}
type DescribeIPInfoReq struct {
	*DescribeIPInfoQuery
	*DescribeIPInfoBody
}
type DescribeLicenseDRMReq struct {
	*DescribeLicenseDRMQuery
	*DescribeLicenseDRMBody
}
type DescribeLiveAuditDataReq struct {
	*DescribeLiveAuditDataQuery
	*DescribeLiveAuditDataBody
}
type DescribeLiveBandwidthDataReq struct {
	*DescribeLiveBandwidthDataQuery
	*DescribeLiveBandwidthDataBody
}
type DescribeLiveBatchPushStreamAvgMetricsReq struct {
	*DescribeLiveBatchPushStreamAvgMetricsQuery
	*DescribeLiveBatchPushStreamAvgMetricsBody
}
type DescribeLiveBatchPushStreamMetricsReq struct {
	*DescribeLiveBatchPushStreamMetricsQuery
	*DescribeLiveBatchPushStreamMetricsBody
}
type DescribeLiveBatchSourceStreamAvgMetricsReq struct {
	*DescribeLiveBatchSourceStreamAvgMetricsQuery
	*DescribeLiveBatchSourceStreamAvgMetricsBody
}
type DescribeLiveBatchSourceStreamMetricsReq struct {
	*DescribeLiveBatchSourceStreamMetricsQuery
	*DescribeLiveBatchSourceStreamMetricsBody
}
type DescribeLiveBatchStreamSessionDataReq struct {
	*DescribeLiveBatchStreamSessionDataQuery
	*DescribeLiveBatchStreamSessionDataBody
}
type DescribeLiveBatchStreamTrafficDataReq struct {
	*DescribeLiveBatchStreamTrafficDataQuery
	*DescribeLiveBatchStreamTrafficDataBody
}
type DescribeLiveBatchStreamTranscodeDataReq struct {
	*DescribeLiveBatchStreamTranscodeDataQuery
	*DescribeLiveBatchStreamTranscodeDataBody
}
type DescribeLiveCallbackDataReq struct {
	*DescribeLiveCallbackDataQuery
	*DescribeLiveCallbackDataBody
}
type DescribeLiveEdgeStatDataReq struct {
	*DescribeLiveEdgeStatDataQuery
	*DescribeLiveEdgeStatDataBody
}
type DescribeLiveISPDataReq struct {
	*DescribeLiveISPDataQuery
	*DescribeLiveISPDataBody
}
type DescribeLiveLogDataReq struct {
	*DescribeLiveLogDataQuery
	*DescribeLiveLogDataBody
}
type DescribeLiveMetricBandwidthDataReq struct {
	*DescribeLiveMetricBandwidthDataQuery
	*DescribeLiveMetricBandwidthDataBody
}
type DescribeLiveMetricTrafficDataReq struct {
	*DescribeLiveMetricTrafficDataQuery
	*DescribeLiveMetricTrafficDataBody
}
type DescribeLiveP95PeakBandwidthDataReq struct {
	*DescribeLiveP95PeakBandwidthDataQuery
	*DescribeLiveP95PeakBandwidthDataBody
}
type DescribeLivePadPresetDetailReq struct {
	*DescribeLivePadPresetDetailQuery
	*DescribeLivePadPresetDetailBody
}
type DescribeLivePadStreamListReq struct {
	*DescribeLivePadStreamListQuery
	*DescribeLivePadStreamListBody
}
type DescribeLivePlayStatusCodeDataReq struct {
	*DescribeLivePlayStatusCodeDataQuery
	*DescribeLivePlayStatusCodeDataBody
}
type DescribeLivePullToPushBandwidthDataReq struct {
	*DescribeLivePullToPushBandwidthDataQuery
	*DescribeLivePullToPushBandwidthDataBody
}
type DescribeLivePullToPushDataReq struct {
	*DescribeLivePullToPushDataQuery
	*DescribeLivePullToPushDataBody
}
type DescribeLivePushStreamCountDataReq struct {
	*DescribeLivePushStreamCountDataQuery
	*DescribeLivePushStreamCountDataBody
}
type DescribeLivePushStreamInfoDataReq struct {
	*DescribeLivePushStreamInfoDataQuery
	*DescribeLivePushStreamInfoDataBody
}
type DescribeLivePushStreamMetricsReq struct {
	*DescribeLivePushStreamMetricsQuery
	*DescribeLivePushStreamMetricsBody
}
type DescribeLiveRecordDataReq struct {
	*DescribeLiveRecordDataQuery
	*DescribeLiveRecordDataBody
}
type DescribeLiveRegionDataReq struct {
	*DescribeLiveRegionDataQuery
	*DescribeLiveRegionDataBody
}
type DescribeLiveSnapshotDataReq struct {
	*DescribeLiveSnapshotDataQuery
	*DescribeLiveSnapshotDataBody
}
type DescribeLiveSourceBandwidthDataReq struct {
	*DescribeLiveSourceBandwidthDataQuery
	*DescribeLiveSourceBandwidthDataBody
}
type DescribeLiveSourceStreamMetricsReq struct {
	*DescribeLiveSourceStreamMetricsQuery
	*DescribeLiveSourceStreamMetricsBody
}
type DescribeLiveSourceTrafficDataReq struct {
	*DescribeLiveSourceTrafficDataQuery
	*DescribeLiveSourceTrafficDataBody
}
type DescribeLiveStreamCountDataReq struct {
	*DescribeLiveStreamCountDataQuery
	*DescribeLiveStreamCountDataBody
}
type DescribeLiveStreamInfoByPageReq struct {
	*DescribeLiveStreamInfoByPageQuery
	*DescribeLiveStreamInfoByPageBody
}
type DescribeLiveStreamSessionDataReq struct {
	*DescribeLiveStreamSessionDataQuery
	*DescribeLiveStreamSessionDataBody
}
type DescribeLiveStreamStateReq struct {
	*DescribeLiveStreamStateQuery
	*DescribeLiveStreamStateBody
}
type DescribeLiveTimeShiftDataReq struct {
	*DescribeLiveTimeShiftDataQuery
	*DescribeLiveTimeShiftDataBody
}
type DescribeLiveTopPlayDataReq struct {
	*DescribeLiveTopPlayDataQuery
	*DescribeLiveTopPlayDataBody
}
type DescribeLiveTrafficDataReq struct {
	*DescribeLiveTrafficDataQuery
	*DescribeLiveTrafficDataBody
}
type DescribeLiveTranscodeDataReq struct {
	*DescribeLiveTranscodeDataQuery
	*DescribeLiveTranscodeDataBody
}
type DescribeLiveTranscodeInfoDataReq struct {
	*DescribeLiveTranscodeInfoDataQuery
	*DescribeLiveTranscodeInfoDataBody
}
type DescribeRecordTaskFileHistoryReq struct {
	*DescribeRecordTaskFileHistoryQuery
	*DescribeRecordTaskFileHistoryBody
}
type DescribeRefererReq struct {
	*DescribeRefererQuery
	*DescribeRefererBody
}
type DescribeRelaySourceV3Req struct {
	*DescribeRelaySourceV3Query
	*DescribeRelaySourceV3Body
}
type DescribeStreamQuotaConfigReq struct {
	*DescribeStreamQuotaConfigQuery
	*DescribeStreamQuotaConfigBody
}
type DisableDomainReq struct {
	*DisableDomainQuery
	*DisableDomainBody
}
type EnableDomainReq struct {
	*EnableDomainQuery
	*EnableDomainBody
}
type EnableHTTPHeaderConfigReq struct {
	*EnableHTTPHeaderConfigQuery
	*EnableHTTPHeaderConfigBody
}
type ForbidStreamReq struct {
	*ForbidStreamQuery
	*ForbidStreamBody
}
type GeneratePlayURLReq struct {
	*GeneratePlayURLQuery
	*GeneratePlayURLBody
}
type GeneratePushURLReq struct {
	*GeneratePushURLQuery
	*GeneratePushURLBody
}
type GetCarouselDetailReq struct {
	*GetCarouselDetailQuery
	*GetCarouselDetailBody
}
type GetCloudMixTaskDetailReq struct {
	*GetCloudMixTaskDetailQuery
	*GetCloudMixTaskDetailBody
}
type GetHLSEncryptDataKeyReq struct {
	*GetHLSEncryptDataKeyQuery
	*GetHLSEncryptDataKeyBody
}
type GetLiveVideoQualityAnalysisTaskDetailReq struct {
	*GetLiveVideoQualityAnalysisTaskDetailQuery
	*GetLiveVideoQualityAnalysisTaskDetailBody
}
type GetPullRecordTaskReq struct {
	*GetPullRecordTaskQuery
	*GetPullRecordTaskBody
}
type KillStreamReq struct {
	*KillStreamQuery
	*KillStreamBody
}
type ListBindEncryptDRMReq struct {
	*ListBindEncryptDRMQuery
	*ListBindEncryptDRMBody
}
type ListCarouselTaskReq struct {
	*ListCarouselTaskQuery
	*ListCarouselTaskBody
}
type ListCertV2Req struct {
	*ListCertV2Query
	*ListCertV2Body
}
type ListCloudMixTaskReq struct {
	*ListCloudMixTaskQuery
	*ListCloudMixTaskBody
}
type ListCommonTransPresetDetailReq struct {
	*ListCommonTransPresetDetailQuery
	*ListCommonTransPresetDetailBody
}
type ListDomainDetailReq struct {
	*ListDomainDetailQuery
	*ListDomainDetailBody
}
type ListHighLightTaskReq struct {
	*ListHighLightTaskQuery
	*ListHighLightTaskBody
}
type ListLiveVideoQualityAnalysisTasksReq struct {
	*ListLiveVideoQualityAnalysisTasksQuery
	*ListLiveVideoQualityAnalysisTasksBody
}
type ListPullRecordTaskReq struct {
	*ListPullRecordTaskQuery
	*ListPullRecordTaskBody
}
type ListPullToPushGroupReq struct {
	*ListPullToPushGroupQuery
	*ListPullToPushGroupBody
}
type ListPullToPushTaskReq struct {
	*ListPullToPushTaskQuery
	*ListPullToPushTaskBody
}
type ListPullToPushTaskV2Req struct {
	*ListPullToPushTaskV2Query
	*ListPullToPushTaskV2Body
}
type ListRelaySourceV4Req struct {
	*ListRelaySourceV4Query
	*ListRelaySourceV4Body
}
type ListTimeShiftPresetV2Req struct {
	*ListTimeShiftPresetV2Query
	*ListTimeShiftPresetV2Body
}
type ListVhostRecordPresetV2Req struct {
	*ListVhostRecordPresetV2Query
	*ListVhostRecordPresetV2Body
}
type ListVhostSnapshotAuditPresetReq struct {
	*ListVhostSnapshotAuditPresetQuery
	*ListVhostSnapshotAuditPresetBody
}
type ListVhostSnapshotPresetV2Req struct {
	*ListVhostSnapshotPresetV2Query
	*ListVhostSnapshotPresetV2Body
}
type ListVhostSubtitleTranscodePresetReq struct {
	*ListVhostSubtitleTranscodePresetQuery
	*ListVhostSubtitleTranscodePresetBody
}
type ListVhostTransCodePresetReq struct {
	*ListVhostTransCodePresetQuery
	*ListVhostTransCodePresetBody
}
type ListVhostWatermarkPresetReq struct {
	*ListVhostWatermarkPresetQuery
	*ListVhostWatermarkPresetBody
}
type ListWatermarkPresetReq struct {
	*ListWatermarkPresetQuery
	*ListWatermarkPresetBody
}
type ListWatermarkPresetDetailReq struct {
	*ListWatermarkPresetDetailQuery
	*ListWatermarkPresetDetailBody
}
type RelaunchPullToPushTaskReq struct {
	*RelaunchPullToPushTaskQuery
	*RelaunchPullToPushTaskBody
}
type RestartTranscodingJobReq struct {
	*RestartTranscodingJobQuery
	*RestartTranscodingJobBody
}
type ResumeStreamReq struct {
	*ResumeStreamQuery
	*ResumeStreamBody
}
type StopLivePadStreamReq struct {
	*StopLivePadStreamQuery
	*StopLivePadStreamBody
}
type StopPullRecordTaskReq struct {
	*StopPullRecordTaskQuery
	*StopPullRecordTaskBody
}
type StopPullToPushTaskReq struct {
	*StopPullToPushTaskQuery
	*StopPullToPushTaskBody
}
type TranscodingJobStatusReq struct {
	*TranscodingJobStatusQuery
	*TranscodingJobStatusBody
}
type UnBindEncryptDRMReq struct {
	*UnBindEncryptDRMQuery
	*UnBindEncryptDRMBody
}
type UnbindCertReq struct {
	*UnbindCertQuery
	*UnbindCertBody
}
type UpdateAuthKeyReq struct {
	*UpdateAuthKeyQuery
	*UpdateAuthKeyBody
}
type UpdateCallbackReq struct {
	*UpdateCallbackQuery
	*UpdateCallbackBody
}
type UpdateCarouselTaskReq struct {
	*UpdateCarouselTaskQuery
	*UpdateCarouselTaskBody
}
type UpdateCloudMixTaskReq struct {
	*UpdateCloudMixTaskQuery
	*UpdateCloudMixTaskBody
}
type UpdateDomainVhostReq struct {
	*UpdateDomainVhostQuery
	*UpdateDomainVhostBody
}
type UpdateEncryptDRMReq struct {
	*UpdateEncryptDRMQuery
	*UpdateEncryptDRMBody
}
type UpdateEncryptHLSReq struct {
	*UpdateEncryptHLSQuery
	*UpdateEncryptHLSBody
}
type UpdateHTTPHeaderConfigReq struct {
	*UpdateHTTPHeaderConfigQuery
	*UpdateHTTPHeaderConfigBody
}
type UpdateIPAccessRuleReq struct {
	*UpdateIPAccessRuleQuery
	*UpdateIPAccessRuleBody
}
type UpdateLivePadPresetReq struct {
	*UpdateLivePadPresetQuery
	*UpdateLivePadPresetBody
}
type UpdatePullToPushGroupReq struct {
	*UpdatePullToPushGroupQuery
	*UpdatePullToPushGroupBody
}
type UpdatePullToPushTaskReq struct {
	*UpdatePullToPushTaskQuery
	*UpdatePullToPushTaskBody
}
type UpdateRecordPresetV2Req struct {
	*UpdateRecordPresetV2Query
	*UpdateRecordPresetV2Body
}
type UpdateRefererReq struct {
	*UpdateRefererQuery
	*UpdateRefererBody
}
type UpdateRelaySourceV3Req struct {
	*UpdateRelaySourceV3Query
	*UpdateRelaySourceV3Body
}
type UpdateRelaySourceV4Req struct {
	*UpdateRelaySourceV4Query
	*UpdateRelaySourceV4Body
}
type UpdateSnapshotAuditPresetReq struct {
	*UpdateSnapshotAuditPresetQuery
	*UpdateSnapshotAuditPresetBody
}
type UpdateSnapshotPresetV2Req struct {
	*UpdateSnapshotPresetV2Query
	*UpdateSnapshotPresetV2Body
}
type UpdateStreamQuotaConfigReq struct {
	*UpdateStreamQuotaConfigQuery
	*UpdateStreamQuotaConfigBody
}
type UpdateSubtitleTranscodePresetReq struct {
	*UpdateSubtitleTranscodePresetQuery
	*UpdateSubtitleTranscodePresetBody
}
type UpdateTimeShiftPresetV3Req struct {
	*UpdateTimeShiftPresetV3Query
	*UpdateTimeShiftPresetV3Body
}
type UpdateTranscodePresetReq struct {
	*UpdateTranscodePresetQuery
	*UpdateTranscodePresetBody
}
type UpdateWatermarkPresetReq struct {
	*UpdateWatermarkPresetQuery
	*UpdateWatermarkPresetBody
}
type UpdateWatermarkPresetV2Req struct {
	*UpdateWatermarkPresetV2Query
	*UpdateWatermarkPresetV2Body
}

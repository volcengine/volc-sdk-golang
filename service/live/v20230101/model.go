package live_v20230101

type BindCertBody struct {

	// REQUIRED; 需要绑定的 HTTPS 证书的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/1126822]接口获取。
	ChainID string `json:"ChainID"`

	// REQUIRED; 填写需要配置 HTTPS 证书的域名。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要绑定证书的域名。
	Domain string `json:"Domain"`

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

type Components1Tzc8QlSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesTosparam struct {
	Bucket          *string `json:"Bucket,omitempty"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	StorageDir      *string `json:"StorageDir,omitempty"`
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
// - FLV 录制参数，开启 FLV 录制时设置。
type Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam `json:"VODParam,omitempty"`
}

// ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam
// - HLS 录制参数，开启 HLS 录制时设置。
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

type ComponentsFceumsSchemasListvqosmetricsdimensionsresPropertiesResultItemsPropertiesDimensionsItems struct {
	Alias string `json:"Alias"`

	Attribute string `json:"Attribute"`

	Name string `json:"Name"`
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
// - MP4 录制参数，开启 MP4 录制时设置。
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

type CreateCertBody struct {

	// REQUIRED; 密钥信息
	Rsa CreateCertBodyRsa `json:"Rsa"`

	// REQUIRED; 证书用途，默认为 https，支持的取值包括：
	// * https：https 认证；
	// * sign：签名校验。
	UseWay string `json:"UseWay"`

	// 证书名称
	CertName *string `json:"CertName,omitempty"`

	// 证书链 ID，用于标识整个证书链，包括叶子证书（服务器证书）、中间证书（中间 CA 证书）以及根证书（根 CA 证书）
	ChainID *string `json:"ChainID,omitempty"`
}

// CreateCertBodyRsa - 密钥信息
type CreateCertBodyRsa struct {

	// REQUIRED; 证书私钥。
	Prikey string `json:"Prikey"`

	// REQUIRED; 证书公钥。
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

	// 证书 ID。
	ChainID *string `json:"ChainID,omitempty"`

	// 使用该证书的域名。
	Domain *string `json:"Domain,omitempty"`

	// 证书用途，包括两种取值。
	// * https：HTTPS 认证；
	// * sign：签名校验。
	UseWay *string `json:"UseWay,omitempty"`
}

type CreateDomainBody struct {

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// 区域，默认指为 cn，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
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

	// REQUIRED; 域名列表，总和最多十个。
	Domains []CreateDomainV2BodyDomainsItem `json:"Domains"`

	// REQUIRED; 区域，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 项目名称，vhost 将归类这个项目里，仅在新创建 vhost 时需要设置。
	ProjectName *string `json:"ProjectName,omitempty"`

	// 标签列表，vhost 将归类这个 tag 里。
	Tags []*CreateDomainV2BodyTagsItem `json:"Tags,omitempty"`
}

type CreateDomainV2BodyDomainsItem struct {

	// REQUIRED; 域名名称。
	DomainName string `json:"DomainName"`

	// REQUIRED; 域名类型，支持以下取值。
	// * push：推流域名
	// * pull-flv：拉流域名
	Type string `json:"Type"`

	// 证书 ID。
	ChainID *string `json:"ChainID,omitempty"`
}

type CreateDomainV2BodyTagsItem struct {

	// REQUIRED; 标签类型，支持以下取值。
	// * System：系统内置标签
	// * Custom：自定义标签
	Category string `json:"Category"`

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

type CreateLiveStreamRecordIndexFilesBody struct {

	// REQUIRED; app名字
	App string `json:"App"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 结束时间
	EndTime string `json:"EndTime"`

	// REQUIRED; 开始时间
	StartTime string `json:"StartTime"`

	// REQUIRED; stream名称
	Stream string `json:"Stream"`

	// 是否需要进行FFProbe嗅探,默认开启
	NeedFFProbe *bool `json:"NeedFFProbe,omitempty"`

	// 输出Bucket,默认为对应录制模板下参数
	OutputBucket *string `json:"OutputBucket,omitempty"`

	// 输出文件名Format,默认为对应录制模板下参数
	OutputObject *string `json:"OutputObject,omitempty"`

	// 关联的TS文件是否是独立的
	SeparatedTS *bool `json:"SeparatedTS,omitempty"`

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
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 录制文件长度,单位 秒
	Duration float32 `json:"Duration"`

	// REQUIRED; 结束时间
	EndTime string `json:"EndTime"`

	// REQUIRED; 流嗅探高度
	Height int32 `json:"Height"`

	// REQUIRED; 录制m3u8文件访问路径
	RecordURL string `json:"RecordURL"`

	// REQUIRED; 开始时间
	StartTime string `json:"StartTime"`

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

type CreatePullCDNSnapshotTaskBody struct {

	// REQUIRED; app名字
	App string `json:"App"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; stream名称
	Stream string `json:"Stream"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty"`

	// 拉流地址
	StreamURL *string `json:"StreamURL,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type CreatePullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata CreatePullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreatePullCDNSnapshotTaskResResult `json:"Result"`
}

type CreatePullCDNSnapshotTaskResResponseMetadata struct {

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

type CreatePullCDNSnapshotTaskResResult struct {

	// REQUIRED; 任务id
	TaskID string `json:"TaskId"`
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

	// 接收拉流转推任务状态回调的地址，最大长度为 2000 个字符，默认为空。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频（Type 为 1）时为必选参数，参数取值及含义如下所示。
	// * -1：无限循环，至任务结束；
	// * 0：有限次循环，循环次数为 PlayTimes 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下为空表示不进行偏移。
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时为必选参数。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	// * 0：不开启；
	// * 1：开启（默认值）。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源（Type 为 0）时，为必选参数，最大长度为 1000 个字符。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频（Type 为 1）时，为必选参数，最多支持传入 30 个点播视频播放地址，每个地址最大长度为 1000 个字符。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Stream *string `json:"Stream,omitempty"`

	// 拉流转推任务的名称，由 1 到 20 位中文、大小写字母和数字组成，默认为空，表示不配置任务名称。
	Title *string `json:"Title,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *CreatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
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

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同，支持填写星号（*）或由 1 到 30 位数字（0 - 9）、大写小字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成。
	// * 应用名称填写星号，即 App 取值为 *时， 流名称也需填写星号，此时表示录制配置为域名空间级别的配置，即直播流使用的域名属于此域名空间时，就会使用此配置进行录制。
	// * 应用名称填写非星号时，表示录制配置为域名空间 + 应用名称 + 流名称级别的配置，即直播流使用的域名属于此域名空间，且 AppName 和 StreamName 字段也同时与 App 和 Stream 的取值匹配时，就会使用此配置进行录制。
	// :::warning 当 App 取值为 * 时，Stream 取值必须为 *。 :::
	App string `json:"App"`

	// REQUIRED; 直播流录制配置的详细配置。
	RecordPresetConfig CreateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相同，支持填写星号（*）或由 1 到 100 位数字（0 - 9）、字母、下划线（_）、短横线（-）和句点（.）组成。
	// * 流名称填写星号时，表示录制配置为应用名称级别的配置，即直播流使用的域名属于此域名空间，且 AppName 字段也与 App 取值同时匹配时，就会使用此配置进行录制。
	// * 流名称填写非星号时，表示录制配置为域名空间 + 应用名称 + 流名称级别的配置，即直播流使用的域名属于此域名空间，且 AppName 和 StreamName 字段也同时与 App 和 Stream 的取值匹配时，就会使用此配置进行录制。
	// :::warning 当 App 取值为 * 时，Stream 取值必须为 *。 :::
	Stream *string `json:"Stream,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfig - 直播流录制配置的详细配置。
type CreateRecordPresetV2BodyRecordPresetConfig struct {

	// 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	FlvParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// 录制为 HLS 合适时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	HlsParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	Mp4Param *CreateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 是否源流录制，默认值为 0，支持的取值即含义如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制为 HLS 格式时，单个 TS 切片时长，单位为秒，默认值为 10，取值范围为 [5,30]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 是否录制转码流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制全部转码流；
	// * 2：通过转码流后缀列表，即TranscodeSuffixList 字段取值匹配转码流。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码流后缀列表，转码流录制配置为根据转码流列表匹配（TranscodeRecord 取值为 2）时生效，TranscodeSuffixList 默认配置为空，效果等同于录制全部转码流。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParam - 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

	// 实时录制场景下，断流等待时长，单位为秒，默认值为 180，取值范围为 [0,3600]。如果实际断流时间小于断流等待时长，录制任务不会停止；如果实际断流时间大于断流等待时长，录制任务会停止，断流恢复后重新开始一个新的录制任务。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制场景下，单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1 和 [300,86400]。
	// * 取值为 -1 时，表示不限制录制时长，录制结束后生成一个完整的录制文件。
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时长生成录制文件，完成录制后一起上传。
	// :::tip 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。 :::
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制场景下，单文件录制时长，单位为秒，默认值为 1800，取值范围为 [300,21600]。录制时间到达设置的单文件录制时长时，会立即生成录制文件实时上传存储。
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

	// TOS 存储对应的 Bucket。例如，存储位置为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 如果启用 TOSParam 配置（Enable
	// 取值为 true），则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
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

	// TOS 存储对应 Bucket 下的存储目录，默认为空。例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam - VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam struct {

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

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果启用 VODParam 配置（Enable 取值为 true），则 VodNamespace
	// 必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 视频点播工作流模板 ID，对于存储在点播的录制文件，会使用该工作流模版对录制的视频进行处理，可登录视频点播控制台 [https://console.volcengine.com/vod/]获取工作流模板 ID，默认为空。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParam - 录制为 HLS 合适时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParam struct {

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
	Result interface{} `json:"Result,omitempty"`
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

type CreateRelaySourceV4Body struct {

	// REQUIRED; 应用名称，即直播流地址的AppName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图审核结果回调地址配置。
	CallbackDetailList []CreateSnapshotAuditPresetBodyCallbackDetailListItem `json:"CallbackDetailList"`

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

	// 推流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
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

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type CreateSnapshotAuditPresetBodyCallbackDetailListItem struct {

	// REQUIRED; 回调地址的类型，当前仅支持 http。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
}

type CreateSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
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

type CreateSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图间隔时间，单位为 s，默认值为 10，取值范围为正整数。
	Interval int32 `json:"Interval"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// ToS 存储的 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 回调详情。
	CallbackDetailList []*CreateSnapshotPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图格式。默认值为 jpeg，支持如下取值。
	// * jpeg
	// * jpg
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// 截图模版状态状态。默认开启。
	// * 1：开启。
	// * 0：关闭。
	Status *int32 `json:"Status,omitempty"`

	// ToS 存储目录，不传为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

type CreateSnapshotPresetBodyCallbackDetailListItem struct {

	// 回调类型，默认值为 http。
	CallbackType *string `json:"CallbackType,omitempty"`

	// 回调地址。
	URL *string `json:"URL,omitempty"`
}

type CreateSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type CreateSnapshotPresetResResponseMetadata struct {

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
	Error   *CreateSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type CreateSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateSnapshotPresetV2Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图配置的详细参数配置。
	SnapshotPresetConfig CreateSnapshotPresetV2BodySnapshotPresetConfig `json:"SnapshotPresetConfig"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图配置生效状态，默认为生效。
	// * 1：生效；
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

	// 当前格式的截图是否开启，默认为 false，取值及含义如下所示。
	// * false：不开启；
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

	// 截图是否使用 veImageX 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
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

	// 截图是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// ToS 存储对应的 bucket 下的存储目录，默认为空。 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
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
	Result interface{} `json:"Result,omitempty"`
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

type CreateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即观看时移的最长时间，单位为 s。支持的取值如下所示。
	// * 86400
	// * 259200
	// * 604800
	// * 1296000
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 时移拉流域名
	PullDomain string `json:"PullDomain"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 开启时移的流名称，同一个 App 最多可指定 20 路。
	Stream *string `json:"Stream,omitempty"`

	// 时移类型。支持的取值如下所示。
	// * 0：录制时移，即时移复用录制模板；
	// * 1：独立时移，即时移不复用录制模板。
	TimeShiftType *int32 `json:"TimeShiftType,omitempty"`
}

type CreateTimeShiftPresetV3Res struct {

	// REQUIRED
	ResponseMetadata CreateTimeShiftPresetV3ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
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

type CreateTranscodePresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 转码后缀，支持由大小写字母（A - Z、a - z）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 视频编码格式；
	// * h265：使用 H.265 视频编码格式；
	// * h266：使用 H.266 视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP 和 BFrames
	// 等。
	Vcodec string `json:"Vcodec"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要转码的直播流使用的域名所属的域名空间。
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
	// * 视频编码格式为 H.265 （Vcodec 取值为 h265）时取值范围为 [0,3]、7、15；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时取值范围为 [0,3]、7、15。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 动态范围，画质增强类型生效，ParamType=hvq时必填，并且h264只支持SDR
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效，不填默认为不开启
	// * 0：不开启
	// * 1：开启
	FISwitch *int32 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，取值为 0 时表示与源流视频帧率相同。
	// 视频帧率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频帧率取值范围为 [0,60]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频帧率取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔时间，单位为秒，默认值为 4，取值范围为 [1,20]。
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

	// 转码模板参数的类型
	// * hvq：表示使用画质增强
	ParamType *string `json:"ParamType,omitempty"`

	// 转码类型是否为极智超清转码，当画质增强开启时，该参数默认为 true，其余默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 (Vcodec取值为h266)时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 使用场景，画质增强时生效，不填不生效
	// * football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * ParamType 取 hvq 时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）取值范围为 0 和 [150,1280]；
	// * 视频编码方式为 H.265 （Vcodec取值为h265）取值范围为 0 和 [150,1280]；
	//
	//
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

	// Anything
	Result interface{} `json:"Result,omitempty"`
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

type CreateVerifyContentBody struct {

	// REQUIRED; 推拉流域名
	Domain string `json:"Domain"`
}

type CreateVerifyContentRes struct {

	// REQUIRED
	ResponseMetadata CreateVerifyContentResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateVerifyContentResResult `json:"Result,omitempty"`
}

type CreateVerifyContentResResponseMetadata struct {

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

// CreateVerifyContentResResult - 视请求的接口而定
type CreateVerifyContentResResult struct {

	// 校验内容记录值
	Content *string `json:"Content,omitempty"`

	// 主机记录
	SubDomain *string `json:"SubDomain,omitempty"`
}

type CreateWatermarkPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX float32 `json:"PosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY float32 `json:"PosY"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// 需要添加水印的直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片字符串字段二选一传入，同时传入时，以水印图片字符串参数为准。
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。与水印相对宽度字段冲突，请选择其中一个传参。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。与水印相对高度字段冲突，请选择其中一个传参。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`
}

type CreateWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
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

type DeleteCallbackBody struct {

	// 应用名称。缺省情况下表示删除 Vhost 下的所有回调配置。如果入参选择 Domain，则不可同时传 App。
	App *string `json:"App,omitempty"`

	// 推流域名。如创建回调 UpdateCallback [https://www.volcengine.com/docs/6469/78553] 时传了参数 Domain，删除时需要传 Domain。
	Domain *string `json:"Domain,omitempty"`

	// 消息类型。缺省情况下表示删除所有消息类型。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间名称。如创建回调 UpdateCallback [https://www.volcengine.com/docs/6469/78553] 时传了参数 Vhost，删除时需要传 Vhost。
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

type DeleteCertBody struct {

	// REQUIRED; 待删除的证书链 ID。
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

type DeleteDomainBody struct {

	// REQUIRED; 待删除域名。
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

type DeleteIPAccessRuleBody struct {

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间名称。
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

type DeletePullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
	TaskID string `json:"TaskId"`
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

	// REQUIRED; 模版名称。可调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口，查询模版名称。
	Preset string `json:"Preset"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 域名空间名称。
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

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 应用名称
	// * 如创建时传了 App，删除时需要传该参数；
	// * 如创建时未传 App，删除时不传该参数。
	App *string `json:"App,omitempty"`

	// 拉流域名。
	// * 如创建时传了 Domain，删除时需要传该参数；
	// * 如创建时未传 Domain，删除时不传该参数。
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

	// REQUIRED; 直播流使用的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，即直播流地址的AppName字段取值，默认为空，表示删除当前域名空间的全局播放触发回源配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图审核配置的名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 截图配置名称。
	Preset string `json:"Preset"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 域名空间名称。
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

	// REQUIRED; 待删除限额配置的推流域名或拉流域名。
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

type DeleteTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 转码配置名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`
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

type DescribeAuthBody struct {

	// REQUIRED; 鉴权场景类型。
	// * push：推流鉴权；
	// * pull：拉流鉴权；
	SceneType string `json:"SceneType"`

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 推/拉流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
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

	// REQUIRED; 鉴权状态。
	// * false：关闭推拉流鉴权；
	// * true：开启推拉流鉴权。
	AuthStatus bool `json:"AuthStatus"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 鉴权场景类型。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// REQUIRED; 有效时长，单位为 s。
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

	// 加密字段。
	EncryptField []*string `json:"EncryptField,omitempty"`

	// 对称加密算法。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty"`

	// 自定义鉴权密钥。
	SecretKey *string `json:"SecretKey,omitempty"`
}

type DescribeCDNSnapshotHistoryBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip
	// * 当您查询指定截图任务详情时，DateFrom 应设置为推流开始时间之前的任意时间。
	// * 查询的最大时间跨度为 7 天。 :::
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	DateTo string `json:"DateTo"`

	// REQUIRED; 流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认为 10，最大值为 1000。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 截图文件保存位置，默认取值为 tos。
	// * tos：TOS 对象存储服务；
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

	// REQUIRED; 截图高度。
	Height int32 `json:"Height"`

	// REQUIRED; 截图文件保存的路径。
	Path string `json:"Path"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 截图时间戳，精度为毫秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// REQUIRED; 截图宽度。
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

	// 消息类型，缺省情况下表示查询全部。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间名称，Vhost和Domain传且仅传一个。
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

	// REQUIRED; 开启鉴权。
	AuthEnable bool `json:"AuthEnable"`

	// REQUIRED; 密钥。
	AuthKeyPrimary string `json:"AuthKeyPrimary"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 消息类型。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType string `json:"MessageType"`

	// REQUIRED; 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：false，不开启；
	// * 1：true，开启。
	TranscodeCallback int32 `json:"TranscodeCallback"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 回调数据列表。
	CallbackDetailList []*DescribeCallbackResResultCallbackListPropertiesItemsItem `json:"CallbackDetailList,omitempty"`
}

type DescribeCallbackResResultCallbackListPropertiesItemsItem struct {

	// REQUIRED; 回调类型，返回 HTTP，表示可以使用 HTTP 和 HTTPS 接收回调事件。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调的 URL。
	URL string `json:"URL"`
}

type DescribeCertDetailSecretV2Body struct {

	// 证书实例 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/81242]接口获取。 :::tip 参数ChainID与CertID传且仅传一个。 :::
	CertID *string `json:"CertID,omitempty"`

	// 证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/81242]接口获取。 :::tip 参数ChainID与CertID传且仅传一个。 :::
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

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为 s。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为 s。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 证书状态，取值与含义的对应关系如下所示。
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

	// REQUIRED; 查询的起始时间，RFC3339 格式的 UTC 时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeFrom string `json:"EndTimeFrom" query:"EndTimeFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式表示的 UTC 时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeTo string `json:"EndTimeTo" query:"EndTimeTo"`

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间（Vhost）下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
	// 页面，查看需要查询的历史直播流使用的域名。
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

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost），默认为空，表示查询所有域名空间（Vhost）下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的历史直播流使用的域名所属的域名空间。
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

	// REQUIRED; 直播流的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 历史直播流的来源类型，取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType string `json:"SourceType"`

	// REQUIRED; 直播流的开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 历史直播流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 历史直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeDenyConfigBody struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 推/拉流域名。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeDenyConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeDenyConfigResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeDenyConfigResResult          `json:"Result,omitempty"`
}

type DescribeDenyConfigResResponseMetadata struct {

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
	Error   *DescribeDenyConfigResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeDenyConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeDenyConfigResResult struct {

	// 配置列表。
	DenyList []*DescribeDenyConfigResResultDenyListItem `json:"DenyList,omitempty"`
}

type DescribeDenyConfigResResultDenyListItem struct {

	// 配置详情列表。
	DenyConfig []*DescribeDenyConfigResResultDenyListPropertiesItemsItem `json:"DenyConfig,omitempty"`

	// 推拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeDenyConfigResResultDenyListPropertiesItemsItem struct {

	// 白名单。
	AllowList []*string `json:"AllowList,omitempty"`

	// 黑名单。
	DenyList []*string `json:"DenyList,omitempty"`
}

type DescribeDomainBody struct {

	// REQUIRED; 域名列表。
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

	// REQUIRED; 所绑定证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的证书信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 绑定的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 区域，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态。状态说明如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	Status int32 `json:"Status"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type DescribeForbiddenStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询当前域名空间（Vhost）下所有的禁推流。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询所有当前域名空间（Vhost）下的禁推直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的禁推直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 指定是否模糊匹配流名称。缺省情况为精准匹配，支持的取值及含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 排列方式，根据推流结束时间排序，默认值为 desc，支持的取值及含义如下所示。
	// * asc：从时间最远到最近排序；
	// * desc：从时间最近到最远排序。
	Sort *string `json:"Sort,omitempty" query:"Sort"`

	// 流名称，取值与禁推直播流时设置的流名称相同，默认查询所有流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成，如果指定 Stream，必须同时指定 App 的值。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 域名空间，取值与禁推直播流时设置的域名空间相同，默认为空，表示查询所有域名空间（Vhost）下的禁推流。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的禁推流使用的域名所属的域名空间。
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

	// REQUIRED; 禁推流被禁推的开始时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 禁推流的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 禁推流结束禁推的时间。
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

type DescribeIPAccessRuleBody struct {

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 是否开启当前限制。
	// * true: 开启；
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 名单中的 IP 信息。
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型，取值及含义如下。
	// * allow: 白名单；
	// * deny: 黑名单。
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

	// REQUIRED; IP 所属地区。非归属火山引擎视频直播的 IP，返回“-”。
	City string `json:"City"`

	// REQUIRED; IP 地址
	IP string `json:"Ip"`

	// REQUIRED; IP 所属运营商。非归属火山引擎视频直播的 IP，返回”-“。 您可以通过DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974]接口查看运营商标识符对应的运营商名称。
	Isp string `json:"Isp"`

	// REQUIRED; 是否归属于火山引擎 CDN 节点。
	// * true：属于；
	// * false：不属于。
	LiveCdnIP bool `json:"LiveCdnIp"`

	// REQUIRED; IP 所属国家或地区。非归属火山引擎视频直播的 IP，返回“-”。
	Location string `json:"Location"`

	// REQUIRED; IP 所属省。非归属火山引擎视频直播的 IP，返回“-”。
	Province string `json:"Province"`
}

type DescribeLiveActivityBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveActivityBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveActivityBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveActivityBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveActivityBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveActivityBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveActivityBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveActivityBandwidthDataResResult `json:"Result"`
}

type DescribeLiveActivityBandwidthDataResResponseMetadata struct {

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

type DescribeLiveActivityBandwidthDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveActivityBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下的峰值带宽，单位为 Mbps。
	PeakBandwidth int32 `json:"PeakBandwidth"`

	// REQUIRED; 峰值带宽的时间戳，RFC3339 格式的 UTC 时间，精度为秒。
	PeakTimestamp string `json:"PeakTimestamp"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

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
	// * huashu：华数；
	// * other：其他。
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
	RegionList []*DescribeLiveActivityBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveActivityBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveActivityBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的峰值带宽，单位为 Mbps。
	Bandwidth int32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveActivityBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveActivityBandwidthDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveAuditDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
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
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	AuditDataList []DescribeLiveAuditDataResResultAuditDataListItem `json:"AuditDataList"`

	// REQUIRED; 按维度拆分后的数据。
	AuditDetailDataList []DescribeLiveAuditDataResResultAuditDetailDataListItem `json:"AuditDetailDataList"`

	// REQUIRED; 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []string `json:"DetailField"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图审核总张数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveAuditDataResResultAuditDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的截图审核总张数。
	Count int32 `json:"Count"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 按维度拆分的数据。 :::tip 请求时，DomainList、ProtocolList和ISPList至少有一个参数传入了多个值时，会返回该参数；否则不返回该参数。
	// 优化：当配置了数据拆分的维度，且对应的维度参数传入多个值时会返回按维度拆分的数据。 :::
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
	// * huashu：华数；
	// * other：其他。
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
	RegionList []*DescribeLiveBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率平均值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率平均值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率平均值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的平均值，即所有 AudioPts 与 VideoPts 差值的平均值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间， RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的平均值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchPushStreamMetricsBody struct {

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	StreamMetricList []DescribeLiveBatchPushStreamMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

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

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchPushStreamMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchSourceStreamAvgMetricsBody struct {

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率平均值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率平均值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率平均值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的平均值，即所有 AudioPts 与 VideoPts 差值的平均值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的平均值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchSourceStreamMetricsBody struct {

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流的监控数据。
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

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchStreamTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 域名列表，缺省情况下表示当前账号下的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 1000，取值范围为 100～1000。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamTrafficDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流维度的流量用量信息详情。
	StreamInfoList []DescribeLiveBatchStreamTrafficDataResResultStreamInfoListItem `json:"StreamInfoList"`

	// REQUIRED; 当前查询条件下，所有流的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 当前查询条件下，所有流的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`
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

type DescribeLiveBatchStreamTranscodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 域名列表，缺省情况下表示当前账号下的所有推拉流域名。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamTranscodeDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 当前流的转码码率。
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

type DescribeLiveCustomizedLogDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip
	// * 当前仅支持查询最近 31 天的日志数据。
	// * 日志文件按设置的拆分粒度进行拆分（以按 1 小时为拆分粒度为例），请确保查询的开始和结束时间包含所查询时段的所有粒度时间。例如，日志文件按 1 小时拆分时，如需查询某日 07:30:00-08:10:00 日志数据，日志查询开始时间应早于
	// 07:00:00，结束时间应晚于 09:00:00。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 日志类型，请联系技术支持同学获取参数值。
	Type string `json:"Type"`

	// 域名列表，缺省情况下表示当前用户的所有推流域名和拉流域名。 :::tip 日志类型为拉流转推日志（Type=relay）时，该参数无效。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认为 20，最大值为 1000。
	PageSize *int32 `json:"PageSize,omitempty"`
}

type DescribeLiveCustomizedLogDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveCustomizedLogDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveCustomizedLogDataResResult `json:"Result"`
}

type DescribeLiveCustomizedLogDataResResponseMetadata struct {

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

type DescribeLiveCustomizedLogDataResResult struct {

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 日志文件的信息列表。
	LogInfoList []DescribeLiveCustomizedLogDataResResultLogInfoListItem `json:"LogInfoList"`

	// REQUIRED; 数据分页信息。
	Pagination DescribeLiveCustomizedLogDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询的日志类型。
	Type string `json:"Type"`
}

type DescribeLiveCustomizedLogDataResResultLogInfoListItem struct {

	// 日志文件的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	DateTime *string `json:"DateTime,omitempty"`

	// 域名。 :::tip 查询拉流转推日志（Type=relay）时，该字段为空。 :::
	Domain *string `json:"Domain,omitempty"`

	// 日志文件下载链接。
	DownloadURL *string `json:"DownloadUrl,omitempty"`

	// 日志文件名称。
	LogName *string `json:"LogName,omitempty"`

	// 日志文件大小，单位为 byte。
	LogSize *int32 `json:"LogSize,omitempty"`
}

// DescribeLiveCustomizedLogDataResResultPagination - 数据分页信息。
type DescribeLiveCustomizedLogDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
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
	// * huashu：华数；
	// * other：其他。
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
	// * pull：拉流日志
	// * push：推流日志
	// * source：回源日志
	// * relay：拉流转推日志
	Type string `json:"Type"`

	// 域名列表，缺省情况下表示当前用户的所有推流域名和拉流域名。 :::tip 日志类型为拉流转推日志（Type=relay）时，该参数无效。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 日志文件的信息列表。
	LogInfoList []DescribeLiveLogDataResResultLogInfoListItem `json:"LogInfoList"`

	// REQUIRED; 数据分页信息。
	Pagination DescribeLiveLogDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 日志文件对应的小时区间，RFC3339 格式的 UTC 时间，精度为秒。
	DateTime string `json:"DateTime"`

	// REQUIRED; 日志文件下载链接。
	DownloadURL string `json:"DownloadUrl"`

	// REQUIRED; 日志文件名称，日志文件命名规则如下。
	// * 与域名相关时：加速域名年月日时间开始时间结束文件拆分序号。例如，www.example.com_2023_08_11_000000_010000_0.gz；
	// * 与域名无关时：年月日时间开始时间结束_文件拆分序号。例如，2023_08_11_000000_010000_0.gz；
	// * 如果某个小时内，当前事件产生的日志大于 150 万条，则会生成为多个日志文件，用文件名最后的序号标注日志文件顺序，例如，2023_08_11_000000_010000_0.gz、2023_08_11_000000_010000_1.gz。
	LogName string `json:"LogName"`

	// REQUIRED; 日志文件大小，单位为 byte。
	LogSize int32 `json:"LogSize"`

	// 域名。 :::tip 查询拉流转推日志（Type=relay）时，该字段为空。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用App构造请求时，需同时定义Stream参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveMetricBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 查询流粒度数据时的流名称参数。 :::tip 使用Stream构造请求时，需同时定义App参数，不可缺省。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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
	// * huashu：华数；
	// * other：其他。
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
	RegionList []*DescribeLiveMetricBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveMetricBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveMetricTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 查询时间范围内的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveMetricTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 应用名称。
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
	// * huashu：华数；
	// * other：其他。
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
	RegionList []*DescribeLiveMetricTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。 :::tip 请求时，DomainList、ProtocolList和ISPList至少有一个参数传入了多个值时，会返回该参数；否则不返回该参数。 :::
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名列表，缺省情况下表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 时间范围内的上下行 95 峰值带宽总和。 :::tip 如果请求时，Regionlist中传入多个 region，则返回这些 region 的上下行带宽 95 峰值总和。 :::
	P95PeakBandwidth float32 `json:"P95PeakBandwidth"`

	// REQUIRED; 95 峰值带宽的时间戳，RFC3339 格式的 UTC 时间，精度为秒。
	P95PeakTimestamp string `json:"P95PeakTimestamp"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
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

type DescribeLivePlayStatusCodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：（默认值）1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况下表示当前用户的所有推拉流域名。
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLivePlayStatusCodeDataBodyRegionListItem `json:"RegionList,omitempty"`

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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	StatusDataList []DescribeLivePlayStatusCodeDataResResultStatusDataListItem `json:"StatusDataList"`

	// REQUIRED; 当前查询条件下的状态码占比数据。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusSummaryDataListItem `json:"StatusSummaryDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 运营商。
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// 域名。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * DstAddrType：推流地址类型。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 推流域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推流地址类型，可选值如下所示。
	// * Live：非第三方；
	// * Third：（默认值）第三方。
	DstAddrTypeList []*string `json:"DstAddrTypeList,omitempty"`
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 按维度拆分的数据。
	BandwidthDetailDataList []*DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名；
	// * DstAddrType：推流地址类型。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推流地址类型。
	// * Live：非第三方；
	// * Third：第三方。
	DstAddrTypeList []*string `json:"DstAddrTypeList,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按推流地址类型维度进行数据拆分时的地址类型信息。
	// * Live：非第三方；
	// * Third：第三方。
	DstAddrType *string `json:"DstAddrType,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 转推带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 1 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需同时定义 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDataListItem `json:"PullToPushDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的拉流转推总时长，单位为分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	PullToPushDetailDataList []*DescribeLivePullToPushDataResResultPullToPushDetailDataListItem `json:"PullToPushDetailDataList,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推总时长，单位为分钟。
	Value float32 `json:"Value"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem `json:"PullToPushDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的拉流转推总时长，单位分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 该时间片内的拉流转推总时长，单位分钟，保留小数点后 2 位
	Value float32 `json:"Value"`
}

type DescribeLivePushStreamCountDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。
	// :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推流域名和拉流域名。
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
	// * huashu：华数；
	// * other：其他。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的推流数量最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TotalStreamDataList []DescribeLivePushStreamCountDataResResultTotalStreamDataListItem `json:"TotalStreamDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名。
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
	// * huashu：华数；
	// * other：其他。
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

type DescribeLivePushStreamMetricsBody struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	MetricList []DescribeLivePushStreamMetricsResResultMetricListItem `json:"MetricList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLivePushStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveRecordDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需同时定义 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称， :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	RecordDataList []DescribeLiveRecordDataResResultRecordDataListItem `json:"RecordDataList"`

	// REQUIRED; 当前查询条件下的录制并发路数最大值。
	RecordPeak int32 `json:"RecordPeak"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	RecordDetailDataList []*DescribeLiveRecordDataResResultRecordDetailDataListItem `json:"RecordDetailDataList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveRecordDataResResultRecordDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需要同时定义 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用 Stream 构造请求时，需要同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	SnapshotDataList []DescribeLiveSnapshotDataResResultSnapshotDataListItem `json:"SnapshotDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图总张数。
	Total int32 `json:"Total"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	SnapshotDetailData []*DescribeLiveSnapshotDataResResultSnapshotDetailDataItem `json:"SnapshotDetailData,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveSnapshotDataResResultSnapshotDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需要同时定义 Domain 和 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名，支持填写拉流域名。 :::tip 使用 Domain 构造请求时，需要同时定义 App 和 Stream 参数，不可缺省。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，缺省情况表示当前用户的所有推拉流域名。 :::tipDomainList 和 Domain 传且仅传一个。 :::
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用 Stream 构造请求时，需要同时定义 Domain 和 App 参数，不可缺省。 :::
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

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询流粒度数据时的应用名称。
	App string `json:"App"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度拆分后的数据。
	BandwidthDetailDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList"`

	// REQUIRED; 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []string `json:"DetailField"`

	// REQUIRED; 查询流粒度数据时的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	// * huashu：华数；
	// * other：其他。
	ISPList []string `json:"ISPList"`

	// REQUIRED; 查询时间范围内的回源峰值带宽，单位为 Mbps。
	PeakBandwidth float32 `json:"PeakBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询流粒度数据时的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 客户端 IP 所属区域列表。
	UserRegionList []DescribeLiveSourceBandwidthDataResResultUserRegionListItem `json:"UserRegionList"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的回源峰值带宽，单位为 Mbps。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按运营商维度进行数据拆分时的运营商信息。
	ISP string `json:"ISP"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的回源峰值带宽，单位为 Mbps。
	PeakBandwidth float32 `json:"PeakBandwidth"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片内回源带宽峰值，单位 Mbps
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveSourceBandwidthDataResResultUserRegionListItem struct {

	// REQUIRED; 区域信息中的大区标识符。
	Area string `json:"Area"`

	// REQUIRED; 区域信息中的国际标识符。
	Country string `json:"Country"`

	// REQUIRED; 区域信息中的省份标识符。
	Province string `json:"Province"`
}

type DescribeLiveSourceStreamMetricsBody struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：（默认值）30 秒。
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
	// * 30：30 秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	MetricList []DescribeLiveSourceStreamMetricsResResultMetricListItem `json:"MetricList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveSourceStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveSourceTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需要同时定义 Domain 和 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名，支持拉流域名。 :::tip 使用 Domain 构造请求时，需要同时定义 App 和 Stream 参数，不可缺省。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，缺省情况表示当前用户的所有推拉流域名。 :::tipDomainList 和 Domain 传且仅传一个。 :::
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用 Stream 构造请求时，需要同时定义 Domain 和 App 参数，不可缺省。 :::
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

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询流粒度数据时的应用名称。
	App string `json:"App"`

	// REQUIRED; 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []string `json:"DetailField"`

	// REQUIRED; 查询流粒度数据时的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	// * huashu：华数；
	// * other：其他。
	ISPList []string `json:"ISPList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询流粒度数据时的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 查询时间范围内的回源总流量，单位为 GB。
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// REQUIRED; 按维度拆分后的数据。
	TrafficDetailDataList []DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList"`

	// REQUIRED; 客户端 IP 所属区域列表。
	UserRegionList []DescribeLiveSourceTrafficDataResResultUserRegionListItem `json:"UserRegionList"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的回源流量，单位 GB。
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按运营商维度进行数据拆分时的运营商信息。
	ISP string `json:"ISP"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的回源中总流量，单位为 GB。
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 回源流量，单位 GB
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultUserRegionListItem struct {

	// REQUIRED; 区域信息中的大区标识符。
	Area string `json:"Area"`

	// REQUIRED; 区域信息中的国家标识符。
	Country string `json:"Country"`

	// REQUIRED; 区域信息中的省份标识符。
	Province string `json:"Province"`
}

type DescribeLiveStreamCountDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。
	// :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 流类型，缺省情况下表示全部类型，支持的流类型取值如下。
	// * push：推流；
	// * relay-source：回源流；
	// * transcode：转码流。
	StreamType []*string `json:"StreamType,omitempty"`
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
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下流数最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TotalStreamDataList []DescribeLiveStreamCountDataResResultTotalStreamDataListItem `json:"TotalStreamDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	StreamDetailDataList []*DescribeLiveStreamCountDataResResultStreamDetailDataListItem `json:"StreamDetailDataList,omitempty"`

	// 流类型，流类型说明如下。
	// * push：拉流；
	// * relay-source：回源流；
	// * transcode：转码流。
	StreamType []*string `json:"StreamType,omitempty"`
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间（Vhost）下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
	// 页面，查看直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 使用流名称进行查询的方式，默认值为 strict，支持的取值即含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 在线流的来源类型，默认为空，表示查询所有来源类型，支持的取值即含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType *string `json:"SourceType,omitempty" query:"SourceType"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 在线流的流类型，默认为空，表示查询所有类型，支持的取值即含义如下所示。
	// * origin：原始流；
	// * trans：转码流。
	StreamType *string `json:"StreamType,omitempty" query:"StreamType"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost），默认为空，表示查询所有域名空间（Vhost）下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
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

	// REQUIRED; 在线流的开始时间。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名。 :::tip 使用App构造请求时，需要同时定义Domain和Stream参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议；
	// * Referer：请求的 Referer 信息。
	// :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名参数。 :::tip 使用Domain构造请求时，需要同时定义App和Stream参数，不可缺省。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名列表，缺省情况表示该用户的所有推拉流域名。 :::tipDomainList和Domain传且仅传一个。 :::
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 请求的 Referer 信息。
	RefererList []*string `json:"RefererList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。
	RegionList []*DescribeLiveStreamSessionDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用Stream构造请求时，需要同时定义Domain和App参数，不可缺省。 :::
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
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDataListItem `json:"SessionDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的请求数。
	TotalRequest int32 `json:"TotalRequest"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议；
	// * Referer：请求的 Referer 信息。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名。
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
	// * huashu：华数；
	// * other：其他。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 请求的 Referer 信息。
	RefererList []*string `json:"RefererList,omitempty"`

	// 区域列表。
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

	// REQUIRED; 数据按时间粒度聚合时，诶个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// 按请求的 Referer 信息进行数据拆分时的 Referer 信息。
	Referer *string `json:"Referer,omitempty"`
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
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
	// 页面，查看需要查询的直播流使用的域名。 :::tipVhost 和 Domain 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
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

type DescribeLiveStreamUsageDataBody struct {

	// 域名。
	Domain *string `json:"Domain,omitempty"`

	// 查询时间，格式为 yyyy-mm-dd HH:MM。 例如：查询时间 2023-01-01 10:00 时表示查询时间范围为 10:00 到 10:01。 :::tip
	// * 未填写查询时间时，默认查询时间为当前时间减 5 分钟。
	// * 最长支持查询的历史时间范围为 31 天。 :::
	QueryTime *string `json:"QueryTime,omitempty"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty"`
}

type DescribeLiveStreamUsageDataRes struct {

	// REQUIRED; 响应数据。
	Response DescribeLiveStreamUsageDataResResponse `json:"Response"`

	// REQUIRED; 请求失败原因：
	// * 请求成功时 Result 为空
	// * 请求失败时 Result 展示失败的原因
	Result string `json:"Result"`

	// REQUIRED; 请求状态：
	// * 1：请求成功
	// * 0：请求失败
	Status int32 `json:"Status"`
}

// DescribeLiveStreamUsageDataResResponse - 响应数据。
type DescribeLiveStreamUsageDataResResponse struct {

	// REQUIRED; 详细数据信息。
	DataInfoList []DescribeLiveStreamUsageDataResResponseDataInfoListItem `json:"DataInfoList"`

	// REQUIRED; 查询时间。
	QueryTime string `json:"QueryTime"`

	// REQUIRED; 请求 ID。
	RequestID string `json:"RequestId"`
}

type DescribeLiveStreamUsageDataResResponseDataInfoListItem struct {

	// REQUIRED; 带宽，单位 kbps。
	Bandwidth int32 `json:"Bandwidth"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 在线人数。
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 协议。
	Protocol string `json:"Protocol"`

	// REQUIRED; 请求数。
	Request int32 `json:"Request"`

	// REQUIRED; 流名称。
	StreamName string `json:"StreamName"`
}

type DescribeLiveTimeShiftDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 86400：（默认值）1 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名空间列表，缺省情况表示查询当前用户的所有域名空间。
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
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TimeShiftDataList []DescribeLiveTimeShiftDataResResultTimeShiftDataListItem `json:"TimeShiftDataList"`

	// 域名空间列表。
	Vhosts []*string `json:"Vhosts,omitempty"`
}

type DescribeLiveTimeShiftDataResResultTimeShiftDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的时移文件存储用量，单位为 GB。
	Storage float32 `json:"Storage"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。聚合粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。聚合粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。聚合粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
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
	// * huashu：华数；
	// * other：其他。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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
	// * huashu：华数；
	// * other：其他。
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
	RegionList []*DescribeLiveTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分后的数据。 :::tip 请求时，DomainList、ProtocolList和ISPList至少有一个参数传入了多个值时，会返回该参数；否则不返回该参数。 :::
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 86400：（默认值）1 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用App构造请求时，需要同时定义Stream参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用Stream构造请求时，需要同时定义App参数，不可缺省。 :::
	Stream *string `json:"Stream,omitempty"`
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
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询时间范围内的转码总时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TranscodeDataList []DescribeLiveTranscodeDataResResultTranscodeDataListItem `json:"TranscodeDataList"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveTranscodeDataResResultTranscodeDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的转码时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType string `json:"TranscodeType"`
}

type DescribeRecordTaskFileHistoryBody struct {

	// REQUIRED; 开始录制时间，RFC3339 格式的 UTC 时间，精度为 s。当您查询指定录制任务详情时，DateFrom 应设置为开始时间之前的任意时间。
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 结束录制时间，结束时间需晚于 DateFrom，且与 DateFrom 间隔不超过 7天，RFC3339 格式的 UTC 时间，精度为 s。
	DateTo string `json:"DateTo"`

	// REQUIRED; 分页查询页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 单个分页中，查询的结果数量。
	PageSize int32 `json:"PageSize"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 流名称，默认查询所有流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成，如果指定 Stream，必须同时指定 App 的值。
	Stream *string `json:"Stream,omitempty"`

	// 录制文件保存位置。默认取值为 tos。
	// * tos
	// * vod
	Type *string `json:"Type,omitempty"`

	// 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
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

	// REQUIRED; 分页信息。
	Pagination DescribeRecordTaskFileHistoryResResultPagination `json:"Pagination"`
}

type DescribeRecordTaskFileHistoryResResultDataItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; ToS 存储空间。
	Bucket string `json:"Bucket"`

	// REQUIRED; 录制时长。
	Duration string `json:"Duration"`

	// REQUIRED; 结束录制时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 录制文件的文件名。
	FileName string `json:"FileName"`

	// REQUIRED; 录制文件存储格式。
	Format string `json:"Format"`

	// REQUIRED; ToS 中的保存路径。
	Path string `json:"Path"`

	// REQUIRED; 开始录制时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// REQUIRED; 录制文件保存在 VoD 时，录制视频的 ID。
	Vid string `json:"Vid"`
}

// DescribeRecordTaskFileHistoryResResultPagination - 分页信息。
type DescribeRecordTaskFileHistoryResResultPagination struct {

	// REQUIRED; 当前页。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 当前页的大小。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 当前页的数据量。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 数据总量。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeRefererBody struct {

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。 :::tip 参数 Domain 和 App 至少传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名。 :::tip
	// * 参数 Domain 和 Vhost 传且仅传一个。
	// * 参数 Domain 和 App 至少传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; Referer 防盗链详情列表。
	RefererInfoList []DescribeRefererResResultRefererListPropertiesItemsItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type DescribeRefererResResultRefererListPropertiesItemsItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词，返回值为 referer。
	Key string `json:"Key"`

	// REQUIRED; 优先级，当前默认返回值为 0。当多域名返回值一致时，按照域名输入顺序区分，越早加入列表的域名优先级越高。
	Priority int32 `json:"Priority"`

	// REQUIRED; 防盗链类型。
	// * deny：黑名单；
	// * allow：白名单。
	Type string `json:"Type"`

	// REQUIRED; 防盗链规则。
	Value string `json:"Value"`
}

type DescribeRelaySourceV3Body struct {

	// REQUIRED; 直播流使用的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。所属的域名空间。
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

	// REQUIRED; 域名空间名称。
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

type DescribeSnapshotAuditPresetDetailBody struct {
	PresetList []*string `json:"PresetList,omitempty"`
}

type DescribeSnapshotAuditPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata DescribeSnapshotAuditPresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeSnapshotAuditPresetDetailResResult `json:"Result,omitempty"`
}

type DescribeSnapshotAuditPresetDetailResResponseMetadata struct {

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

// DescribeSnapshotAuditPresetDetailResResult - 视请求的接口而定
type DescribeSnapshotAuditPresetDetailResResult struct {

	// REQUIRED
	PresetDetailList []DescribeSnapshotAuditPresetDetailResResultPresetDetailListItem `json:"PresetDetailList"`
}

// DescribeSnapshotAuditPresetDetailResResultPresetDetailListItem - 审核模版详细信息。
type DescribeSnapshotAuditPresetDetailResResultPresetDetailListItem struct {

	// REQUIRED
	AccountID string `json:"AccountID"`

	// REQUIRED
	AshePresetName string `json:"AshePresetName"`

	// REQUIRED
	AuditType string `json:"AuditType"`

	// REQUIRED; ToS 存储空间 bucket。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	Bucket string `json:"Bucket"`

	// REQUIRED; 审核结果回调配置。
	CallbackDetailList []DescribeSnapshotAuditPresetDetailResResultPresetDetailListPropertiesItemsItem `json:"CallbackDetailList"`

	// REQUIRED
	CreatedAt int32 `json:"CreatedAt"`

	// REQUIRED; 审核模板描述。
	Description string `json:"Description"`

	// REQUIRED; 审核标签名称，若为空，则默认请求所有基础模型。支持以下取值。
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

	// REQUIRED; 审核模板名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; veimageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID string `json:"ServiceID"`

	// REQUIRED
	SnapshotConfig DescribeSnapshotAuditPresetDetailResResultPresetDetailListItemSnapshotConfig `json:"SnapshotConfig"`

	// REQUIRED; 存储策略。支持以下取值。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储全部图片。
	StorageStrategy int32 `json:"StorageStrategy"`

	// REQUIRED
	UpdatedAt int32 `json:"UpdatedAt"`
}

type DescribeSnapshotAuditPresetDetailResResultPresetDetailListItemSnapshotConfig struct {

	// REQUIRED
	Format string `json:"Format"`

	// REQUIRED
	Interval float32 `json:"Interval"`

	// REQUIRED
	SnapshotObject string `json:"SnapshotObject"`

	// REQUIRED
	StorageDir string `json:"StorageDir"`
}

type DescribeSnapshotAuditPresetDetailResResultPresetDetailListPropertiesItemsItem struct {

	// REQUIRED; 回调的类型 HTTP。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
}

type DescribeStreamQuotaConfigBody struct {

	// REQUIRED; 待查询限额配置的推流域名或拉流域名。
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

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig - 推流域名的推流路数限额配置信息。 :::tipDomain 为拉流域名时返回。
// :::
type DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig struct {

	// REQUIRED; 拉流带宽限额。
	Quota int32 `json:"Quota"`

	// REQUIRED; 拉流带宽限额的计量单位。
	QuotaUnit string `json:"QuotaUnit"`

	// 拉流带宽限额告警阈值，缺省情况表示未设置告警。
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`

	// 拉流带宽限额告警的计量单位，缺省情况表示不未设置告警。
	AlarmThresholdUnit *string `json:"AlarmThresholdUnit,omitempty"`
}

// DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig - 推流域名的推流路数限额配置信息。 :::tipDomain 为推流域名时返回。
// :::
type DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig struct {

	// REQUIRED; 推流路数限额。
	Quota int32 `json:"Quota"`

	// 推流路数限额告警阈值，缺省情况表示未设置告警。
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`
}

type DescribeStreamQuotaConfigResResultQuotaListPropertiesItemsItem struct {

	// 推流域名的推流路数限额配置信息。 :::tipDomain 为拉流域名时返回。 :::
	BandwidthConfig *DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig `json:"BandwidthConfig,omitempty"`

	// 推流域名的推流路数限额配置信息。 :::tipDomain 为推流域名时返回。 :::
	StreamConfig *DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig `json:"StreamConfig,omitempty"`
}

type DisableDomainBody struct {

	// REQUIRED; 待禁用域名。
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

	// REQUIRED; 待启用域名。
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

type ForbidStreamBody struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 流名称，取值与直播流地址的 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 直播流使用的域名，您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看待禁推的直播流使用的域名。
	// :::tip 参数 Domain 和
	// Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 禁推的结束时间，RFC3339 格式的 UTC 时间，精度为毫秒，禁推有效期最长为 90 天，默认为当前时间加 90 天。
	EndTime *string `json:"EndTime,omitempty"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
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

	// REQUIRED; 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// 过期时间，拉流地址的有效时间，过期后需要重新生成。RFC3339 格式的 UTC 时间，精度为秒，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 转码流后缀，不传默认为空，可通过调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查询。
	Suffix *string `json:"Suffix,omitempty"`

	// CDN 类型，默认值为 fcdn，支持如下取值。
	// * fcdn：火山引擎流媒体直播 CDN；
	// * 3rd：第三方 CDN。
	Type *string `json:"Type,omitempty"`

	// 有效时长，拉流地址的有效时间，过期后需要重新生成。单位为秒，取值 ﹥0，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。 :::
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

	// REQUIRED; 拉流地址信息列表。
	URLList []GeneratePlayURLResResultURLListItem `json:"URLList"`
}

type GeneratePlayURLResResultURLListItem struct {

	// REQUIRED; CDN 类型。
	// * fcdn：火山引擎流媒体直播 CDN；
	// * 3rd：第三方 CDN。
	CDN string `json:"CDN"`

	// REQUIRED; 协议类型，包括 hls、flv 和 rtmp。
	Protocol string `json:"Protocol"`

	// REQUIRED; 地址类型，可能的值为：
	// * push：推流；
	// * pull：拉流；
	// * 3rdplay(relaysource)：第三方回源；
	// * 3rdplay(relaysink)：第三方转推。
	Type string `json:"Type"`

	// REQUIRED; 生成的拉流地址。
	URL string `json:"URL"`
}

type GeneratePushURLBody struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 推流域名名称，需要推流地址的域名，不填返回Vhost下所有推流域名生成的地址。
	Domain *string `json:"Domain,omitempty"`

	// 过期时间，推流地址的有效时间，过期后需要重新生成。RFC3339 格式的 UTC 时间，精度为秒，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 有效时长，推流地址的有效时间，过期后需要重新生成。单位为秒，取值 ﹥0，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。 :::
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

type GetPullCDNSnapshotTaskBody struct {

	// REQUIRED; 直播截图任务 ID，任务的唯一标识。
	TaskID string `json:"TaskID"`
}

type GetPullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result GetPullCDNSnapshotTaskResResult `json:"Result"`
}

type GetPullCDNSnapshotTaskResResponseMetadata struct {

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

// GetPullCDNSnapshotTaskResResult - 视请求的接口而定
type GetPullCDNSnapshotTaskResResult struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 截图任务的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 截图任务的开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 任务状态：
	// * 停用
	// * 未开始
	// * 生效中
	// * 已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 直播截图任务 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type GetPullRecordTaskBody struct {

	// REQUIRED; 直播录制任务的 ID，任务的唯一标识。
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

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流/拉流域名
	Domain string `json:"Domain"`

	// REQUIRED; 任务结束时间
	EndTime string `json:"EndTime"`

	// REQUIRED; 任务开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 任务状态，有以下几种状态：
	// * 停用
	// * 未开始
	// * 生效中
	// * 已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 直播录制任务的 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type KillStreamBody struct {

	// 直播流使用的应用名称。
	App *string `json:"App,omitempty"`

	// 直播流使用的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间，您可以调用 DescribeLiveStreamInfoByPage [https://www.volcengine.com/docs/6469/1126841] 接口，查看待断开的在线流的信息，包括 Vhost、APP 和 Stream。
	Vhost *string `json:"Vhost,omitempty"`
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

type ListCertV2Body struct {

	// 是否启用证书，默认值为 true，支持的取值包括：
	// * true：启用证书；
	// * false：禁用证书。
	Available *bool `json:"Available,omitempty"`

	// 证书名称，支持输入证书名称中的关键字，进行模糊查询
	CertName *string `json:"CertName,omitempty"`

	// 域名，查询该域名对应的证书，支持精确查询
	Domain *string `json:"Domain,omitempty"`

	// 只有填了Available，这个字段才生效。
	Expiring *bool `json:"Expiring,omitempty"`
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

	// 证书列表。
	CertList []*ListCertV2ResResultCertListItem `json:"CertList,omitempty"`
}

type ListCertV2ResResultCertListItem struct {

	// REQUIRED; 与证书绑定的域名列表。
	CertDomainList []string `json:"CertDomainList"`

	// REQUIRED; 证书实例ID。
	CertID string `json:"CertID"`

	// REQUIRED; 证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 证书链 ID。
	ChainID string `json:"ChainID"`

	// REQUIRED; 火山证书中心证书链 ID。
	ChainIDVolc string `json:"ChainIDVolc"`

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为 s。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为 s。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 证书状态，由证书管理平台返回，支持的取值如下所示。
	// * OK：正常；
	// * Expire：过期；
	// * 30days：有效期剩余 30 天；
	// * 15days：有效期剩余 15 天；
	// * 7days：有效期剩余 7 天；
	// * 1days：有效期剩余 1 天。
	Status string `json:"Status"`
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
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 Opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 宽高自适应模式开关，支持的取值及含义如下。
	// * 0：关闭宽高自适应
	// * 1：开启宽高自适应 :::tip
	// * 关闭宽高自适应时，转码配置分辨率取 Width 和 Height 的值对转码视频进行拉伸；
	// * 开启宽高自适应时，转码配置分辨率按照 ShortSide 、 LongSide 、Width 、Height 的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 帧率，单位为 fps。帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。 :::tip
	// * 当 As 的取值为 0 即关闭宽高自适应时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * Width 和 Height 任一配置为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 是否极智超清转码，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。 :::tip
	// * 当 As 的取值为 0 即关闭宽高自适应时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * Width 和 Height 任一配置为 0 时，转码视频将保持源流尺寸；
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

	// REQUIRED; 当前页码，取值范围为 [1,1000]。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 分页大小，取值范围为 [1, 1000]
	PageSize int32 `json:"PageSize"`

	// 域名名称列表，缺省情况下表示全部。
	DomainNameList []*string `json:"DomainNameList,omitempty"`

	// 域名区域列表。缺省情况下表示全部，区域类型支持以下取值。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	DomainRegionList []*string `json:"DomainRegionList,omitempty"`

	// 域名状态列表。缺省情况下表示全部。支持的取值如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	DomainStatusList []*int32 `json:"DomainStatusList,omitempty"`

	// 域名类型列表。缺省情况下表示全部。域名类型支持以下取值。
	// * push：推流域名；
	// * pull-flv：拉流域名。
	DomainTypeList []*string `json:"DomainTypeList,omitempty"`

	// 域名空间列表，缺省情况下表示全部。
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

	// REQUIRED; 所绑定证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的证书信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 绑定的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 区域，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态。状态说明如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	Status int32 `json:"Status"`

	// REQUIRED; 标签信息。
	Tags []ListDomainDetailResResultDomainListPropertiesItemsItem `json:"Tags"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListDomainDetailResResultDomainListPropertiesItemsItem struct {

	// REQUIRED; 标签类型。
	// * System：系统内置标签
	// * Custom：自定义标签
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type ListPullCDNSnapshotTaskBody struct {

	// REQUIRED; 分页数
	PageNum string `json:"PageNum"`

	// REQUIRED; 分页的大小
	PageSize string `json:"PageSize"`
}

type ListPullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata ListPullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListPullCDNSnapshotTaskResResult `json:"Result,omitempty"`
}

type ListPullCDNSnapshotTaskResResponseMetadata struct {

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

// ListPullCDNSnapshotTaskResResult - 视请求的接口而定
type ListPullCDNSnapshotTaskResResult struct {

	// REQUIRED; 直播截图列表记录。
	List []ListPullCDNSnapshotTaskResResultListItem `json:"List"`

	// REQUIRED; 分页信息。
	Pagination ListPullCDNSnapshotTaskResResultPagination `json:"Pagination"`
}

type ListPullCDNSnapshotTaskResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 截图任务的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 截图任务的开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 4种状态:停用，未开始，生效中，已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 任务ID，任务的唯一标识。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListPullCDNSnapshotTaskResResultPagination - 分页信息。
type ListPullCDNSnapshotTaskResResultPagination struct {

	// REQUIRED; 当前页。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 当前页的数据量。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 总数据量。
	TotalCount int32 `json:"TotalCount"`
}

type ListPullRecordTaskBody struct {

	// REQUIRED; 分页数
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 分页的大小
	PageSize int32 `json:"PageSize"`
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

	// REQUIRED; 直播录制列表记录。
	List []ListPullRecordTaskResResultListItem `json:"List"`

	// REQUIRED; 分页信息。
	Pagination ListPullRecordTaskResResultPagination `json:"Pagination"`
}

type ListPullRecordTaskResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名或拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 录制的结束时间，RFC3339 格式表示的 UTC 时间戳，精度为 s。
	EndTime string `json:"EndTime"`

	// REQUIRED; 录制的开始时间，RFC3339 格式表示的 UTC 时间戳，精度为 s。
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

// ListPullRecordTaskResResultPagination - 分页信息。
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

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// 0：从断流处续播（默认值）； 1：从断流处+自然流逝时长处续播。
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

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时配置生效。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
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

type ListRelaySourceV4Body struct {

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。
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

	// 时移类型。默认类型为 VoD。
	// * vod：点播时移，获取 vod 类型的时移模板配置信息；
	// * tos：直播时移，获取 tos 以及 fcdn-tos 类型的时移模板配置信息。
	Type *string `json:"Type,omitempty"`

	// 域名空间名称。
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

	// 模板列表。
	List []*ListTimeShiftPresetV2ResResultListItem `json:"List,omitempty"`
}

type ListTimeShiftPresetV2ResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; ToS 存储目录。
	Bucket string `json:"Bucket"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 最大时移时长，即观看时移的最长时间，单位为 s。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 模板名称。
	Name string `json:"Name"`

	// REQUIRED; 直播时移配置模版状态。
	// * 0：配置中；
	// * 1：已启用。
	Status int32 `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 类型。默认类型为 VoD。
	// * VoD
	// * ToS
	// * fcdn-ToS
	Type string `json:"Type"`

	// REQUIRED; 点播空间。
	VODNamespace string `json:"VODNamespace"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListVhostRecordPresetV2Body struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 直播录制的存储类型。默认值为 tos，支持以下取值。
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

	// REQUIRED; 录制模板列表。
	PresetList []ListVhostRecordPresetV2ResResultPresetListItem `json:"PresetList"`
}

type ListVhostRecordPresetV2ResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 模板详细信息。
	SlicePresetV2 *ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 `json:"SlicePresetV2,omitempty"`
}

// ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 - 模板详细信息。
type ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 struct {

	// 模板 ID。
	ID *int32 `json:"ID,omitempty"`

	// 模板名称。
	Name *string `json:"Name,omitempty"`

	// 录制模板详细配置。
	RecordPresetConfig *ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig `json:"RecordPresetConfig,omitempty"`
}

type ListVhostSnapshotAuditPresetBody struct {

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
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

	// REQUIRED; 截图间隔时间，单位秒，取值范围为[0.1,10]，支持保留两位小数。
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

type ListVhostSnapshotPresetBody struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图存储类型。
	// * tos；
	// * imageX。
	Type *string `json:"Type,omitempty"`
}

type ListVhostSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostSnapshotPresetResResult          `json:"Result,omitempty"`
}

type ListVhostSnapshotPresetResResponseMetadata struct {

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
	Error   *ListVhostSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostSnapshotPresetResResult struct {

	// 模版列表。
	PresetList []*ListVhostSnapshotPresetResResultPresetListItem `json:"PresetList,omitempty"`
}

type ListVhostSnapshotPresetResResultPresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图模板名称。
	SlicePreset *ListVhostSnapshotPresetResResultPresetListItemSlicePreset `json:"SlicePreset,omitempty"`
}

// ListVhostSnapshotPresetResResultPresetListItemSlicePreset - 截图模板名称。
type ListVhostSnapshotPresetResResultPresetListItemSlicePreset struct {

	// 截图在 ToS 中的存储位置。
	Bucket *string `json:"Bucket,omitempty"`

	// 回调信息。
	CallbackDetail []*ListVhostSnapshotPresetResResultPresetListPropertiesItemsItem `json:"CallbackDetail,omitempty"`

	// 截图间隔时间。
	Interval *int32 `json:"Interval,omitempty"`

	// 截图模版名称。
	Preset *string `json:"Preset,omitempty"`

	// veImageX 的服务 ID。
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图模版状态。
	// * 1：开启
	// * 0：关闭
	Status *int32 `json:"Status,omitempty"`
}

// ListVhostSnapshotPresetResResultPresetListPropertiesItemsItem - 回调信息
type ListVhostSnapshotPresetResResultPresetListPropertiesItemsItem struct {

	// REQUIRED; 回调地址。
	URL string `json:"URL"`

	// 回调类型。
	// * http
	// * nsq
	// * kafka
	// * rpc
	CallbackType *string `json:"CallbackType,omitempty"`
}

type ListVhostSnapshotPresetV2Body struct {

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 - 截图配置基础信息。
type ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 struct {

	// REQUIRED; 截图配置名称。
	Name string `json:"Name"`

	// REQUIRED; 截图配置详细信息。
	SnapshotPresetConfig ListVhostSnapshotPresetV2ResResultPresetListProperties `json:"SnapshotPresetConfig"`

	// REQUIRED; 截图配置生效状态。
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

	// REQUIRED; 当前格式的截图是否开启，默认为 false，取值及含义如下所示。
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

type ListVhostTransCodePresetBody struct {

	// REQUIRED; 域名空间。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
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
	FISwitch *string `json:"FISwitch,omitempty"`

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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset struct {

	// 音频编码格式，取值含义如下。
	// * aac：使用 AAC 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 Opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应；
	// * 1：开启宽高自适应。 :::tip
	// * 关闭宽高自适应时，转码配置分辨率取 Width 和 Height 的值对转码视频进行拉伸；
	// * 开启宽高自适应时，转码配置分辨率按照 ShortSide 、 LongSide 、Width 、Height 的优先级取值，另一边等比缩放。 :::
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

	// 2 个参考帧之间的最大 B 帧数。取值为 0 时，表示去除 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 动态范围，画质增强类型生效
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效
	// * 0：不开启
	// * 1：开启
	FISwitch *string `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅，开启智能插帧时该值为 0。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码模板参数的类型
	// * hvq：表示使用画质增强
	ParamType *string `json:"ParamType,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 是否极智超清转码，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	Roi *bool `json:"Roi,omitempty"`

	// 使用场景，画质增强时生效 football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为 s，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，取值及含义如下。
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

type ListVhostTransCodePresetResResultCustomizePresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
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
	FISwitch *string `json:"FISwitch,omitempty"`

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

	// REQUIRED; 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
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

	// 统计消息，提供查询成功和失败的数量。
	StaticsMsg *string `json:"StaticsMsg,omitempty"`

	// 获取模板失败的列表，返回获取失败的模版及获取失败的原因。
	WatermarkErrMsgList []*ListVhostWatermarkPresetResResultWatermarkErrMsgListItem `json:"WatermarkErrMsgList,omitempty"`

	// 水印模版列表。
	WatermarkPresetList []*ListVhostWatermarkPresetResResultWatermarkPresetListItem `json:"WatermarkPresetList,omitempty"`
}

type ListVhostWatermarkPresetResResultWatermarkErrMsgListItem struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 获取水印模板失败的具体错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListVhostWatermarkPresetResResultWatermarkPresetListItem struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 水印模版 ID。
	ID *int32 `json:"ID,omitempty"`

	// 直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片链接。
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

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListVqosMetricsDimensionsQuery struct {

	// REQUIRED
	VqosService string `json:"VqosService" query:"VqosService"`
}

type ListVqosMetricsDimensionsRes struct {

	// REQUIRED
	ResponseMetadata ListVqosMetricsDimensionsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result []ListVqosMetricsDimensionsResResultItem `json:"Result"`
}

type ListVqosMetricsDimensionsResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestId"`
}

type ListVqosMetricsDimensionsResResultItem struct {
	Dimensions []*ComponentsFceumsSchemasListvqosmetricsdimensionsresPropertiesResultItemsPropertiesDimensionsItems `json:"Dimensions,omitempty"`
	Metrics    []*ListVqosMetricsDimensionsResResultPropertiesItemsItem                                             `json:"Metrics,omitempty"`
	Service    *string                                                                                              `json:"Service,omitempty"`
}

type ListVqosMetricsDimensionsResResultPropertiesItemsItem struct {

	// REQUIRED
	Alias string `json:"Alias"`

	// REQUIRED
	Attached string `json:"Attached"`

	// REQUIRED
	Attribute string `json:"Attribute"`

	// REQUIRED
	Desc string `json:"Desc"`

	// REQUIRED
	Name string `json:"Name"`

	// REQUIRED
	Type string `json:"Type"`
}

type ListWatermarkPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`
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

	// 水印模版 ID。
	ID *int32 `json:"ID,omitempty"`

	// 直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片链接。
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

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type RestartPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
	TaskID string `json:"TaskId"`
}

type RestartPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata RestartPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type RestartPullToPushTaskResResponseMetadata struct {

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
	Error   *RestartPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type RestartPullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
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

type StopPullCDNSnapshotTaskBody struct {

	// REQUIRED; 任务id
	TaskID string `json:"TaskId"`
}

type StopPullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata StopPullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopPullCDNSnapshotTaskResResponseMetadata struct {

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

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
	TaskID string `json:"TaskId"`
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

type UnbindCertBody struct {

	// REQUIRED; 需要解绑证书的域名。
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

	// REQUIRED; 鉴权详情，数量阈值为 100。
	AuthDetailList []UpdateAuthKeyBodyAuthDetailListItem `json:"AuthDetailList"`

	// REQUIRED; 鉴权场景类型。
	// * push：推流鉴权；
	// * pull：拉流鉴权；
	SceneType string `json:"SceneType"`

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 推/拉流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 鉴权状态。创建推拉流鉴权时，默认值为 false；更新推拉流鉴权时，缺省情况表示不修改推拉流鉴权状态。
	// * false：关闭推拉流鉴权；
	// * true：开启推拉流鉴权。
	PushPullEnable *bool `json:"PushPullEnable,omitempty"`

	// 有效时长，单位为 s，默认值为 604800，取值范围为 [60,2592000]。
	ValidDuration *int32 `json:"ValidDuration,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateAuthKeyBodyAuthDetailListItem struct {

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

	// REQUIRED; 消息类型。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType string `json:"MessageType"`

	// domain / app 二选一必传
	App *string `json:"App,omitempty"`

	// 是否开启鉴权，默认为 false。取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	AuthEnable *bool `json:"AuthEnable,omitempty"`

	// 密钥。 :::tip 如果 AuthEnable 为 true，则密钥必填。 :::
	AuthKeyPrimary *string `json:"AuthKeyPrimary,omitempty"`

	// 推流域名。Vhost和Domain传且仅传一个。
	Domain *string `json:"Domain,omitempty"`

	// 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：false，不开启；
	// * 1：true，开启。
	TranscodeCallback *int32 `json:"TranscodeCallback,omitempty"`

	// domain / app 二选一必传
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateCallbackBodyCallbackDetailListItem struct {

	// REQUIRED; 回调类型，支持设置为 HTTP，表示可以使用 HTTP 和 HTTPS 接收回调事件。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调的 URL。
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

type UpdateDenyConfigBody struct {

	// REQUIRED; 黑白名单配置列表。
	DenyConfigList []UpdateDenyConfigBodyDenyConfigListItem `json:"DenyConfigList"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type UpdateDenyConfigBodyDenyConfigListItem struct {

	// 白名单。
	AllowList []*string `json:"AllowList,omitempty"`

	// 黑名单。
	DenyList []*string `json:"DenyList,omitempty"`
}

type UpdateDenyConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateDenyConfigResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateDenyConfigResResponseMetadata struct {

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
	Error   *UpdateDenyConfigResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateDenyConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateDomainVhostBody struct {

	// REQUIRED; 需要切换的拉流/推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 目的域名空间。
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

type UpdateIPAccessRuleBody struct {

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; IP 访问限制规则。
	IPAccessRule UpdateIPAccessRuleBodyIPAccessRule `json:"IPAccessRule"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// UpdateIPAccessRuleBodyIPAccessRule - IP 访问限制规则。
type UpdateIPAccessRuleBodyIPAccessRule struct {

	// REQUIRED; 是否开启当前限制，取值及含义如下。
	// * true: 开启；
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 名单中的 IP 信息。
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型，取值及含义如下。
	// * allow: 白名单；
	// * deny: 黑名单。
	Type string `json:"Type"`
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

	// 接收拉流转推任务状态回调的地址，最大长度为 2000 个字符。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频（Type 为 1）时为必选参数，参数取值及含义如下所示。
	// * -1：无限循环，至任务结束；
	// * 0：有限次循环，循环次数为 PlayTimes 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下表示不进行偏移。
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时为必选参数。
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

	// 拉流转推任务的名称，由 1 到 20 位中文、大小写字母和数字组成，默认为空，表示不修改任务名称。
	Title *string `json:"Title,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *UpdatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
}

// UpdatePullToPushTaskBodyWatermark - 为拉流转推视频添加的水印配置信息。
type UpdatePullToPushTaskBodyWatermark struct {

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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 模版名称，由 1 到 50 位数字、字母、下划线及"-"和"."组成。可调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]
	// 接口，查询模版名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 录制模板详细配置。
	RecordPresetConfig *UpdateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfig - 录制模板详细配置。
type UpdateRecordPresetV2BodyRecordPresetConfig struct {

	// FLV 录制参数，开启 FLV 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
	FlvParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// HLS 录制参数，开启 HLS 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
	HlsParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// MP4 录制参数，开启 MP4 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
	Mp4Param *UpdateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 源流录制，默认值为 0。支持的取值如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tipTranscodeRecord和OriginRecord的取值至少一个不为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制 HLS 格式时，单个 TS 切片时长，单位为 s，默认值为 5，取值范围为 [5,30]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 转码流录制，默认值为 0。支持的取值如下所示。
	// * 0：不录制；
	// * 1：录制。
	// * 2：全部录制，如果录制转码流后缀列表（TranscodeSuffixList）为空则全部录制，不为空则录制 TranscodeSuffixList 命中的转码后缀。
	// :::tipTranscodeRecord 和 OriginRecord 的取值至少一个不为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 录制转码流后缀列表，转码流录制配置为全部录制时（TranscodeRecord 配置等于 2）生效。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParam - FLV 录制参数，开启 FLV 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

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
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam - TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam struct {

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

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam - VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam struct {

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

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParam - HLS 录制参数，开启 HLS 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
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

// UpdateRecordPresetV2BodyRecordPresetConfigMp4Param - MP4 录制参数，开启 MP4 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
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

	// REQUIRED; Referer 防盗链信息列表。
	RefererInfoList []UpdateRefererBodyRefererInfoListItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。 :::tip 参数 Domain 和 App 传且仅传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名。 :::tip 参数 Domain 和 App 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`
}

type UpdateRefererBodyRefererInfoListItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词默认取值为 referer。
	Key string `json:"Key"`

	// REQUIRED; 防盗链类型，支持如下取值。
	// * deny：黑名单；
	// * allow：白名单。
	Type string `json:"Type"`

	// 指定域名的优先级。默认值为 0，取值范围为 [0,100]，数值越大，优先级越高。如果优先级相同，则越早加入列表的域名优先级越高。
	Priority *int32 `json:"Priority,omitempty"`

	// 防盗链规则，即设置的黑名单或白名单的域名，最大长度限制 300 个字符。
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

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
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

	// REQUIRED; 回源协议，支持两种回源协议。
	// * rtmp：RTMP 回源协议；
	// * flv：FLV 回源协议。
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

	// REQUIRED; 应用名称，即直播流地址的AppName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 回源地址列表，支持 RTMP、FLV、HLS 和 SRT 回源协议。
	// :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 流名称，即直播流地址的StreamName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、字母、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图审核配置的名称，您可以通过调用查询截图审核配置列表 [https://www.volcengine.com/docs/6469/1126870]接口获取。
	PresetName string `json:"PresetName"`

	// ToS 存储对应的 Bucket。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图审核结果回调地址配置。
	CallbackDetailList []*UpdateSnapshotAuditPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`

	// 截图审核配置的描述。
	Description *string `json:"Description,omitempty"`

	// 推流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 截图间隔时间，单位秒，取值范围为[0.1,10]，支持保留两位小数。
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

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateSnapshotAuditPresetBodyCallbackDetailListItem struct {

	// REQUIRED; 回调地址的类型，当前仅支持 http。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
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

type UpdateSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图模板名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// ToS 的存储 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 回调详情。
	CallbackDetailList []*UpdateSnapshotPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`

	// 截图间隔时间，单位为 s，默认值为 10，取值范围为正整数
	Interval *int32 `json:"Interval,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图格式。支持如下取值。- jpeg - jpg
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// 截图模版状态。
	// * 1：开启
	// * 0：关闭
	Status *int32 `json:"Status,omitempty"`

	// ToS 的存储目录，不传为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

type UpdateSnapshotPresetBodyCallbackDetailListItem struct {

	// 回调类型，默认值为 http。
	CallbackType *string `json:"CallbackType,omitempty"`

	// 回调地址。
	URL *string `json:"URL,omitempty"`
}

type UpdateSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSnapshotPresetResResponseMetadata struct {

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
	Error   *UpdateSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateSnapshotPresetV2Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图配置的名称。
	Preset string `json:"Preset"`

	// REQUIRED; 截图配置的详细参数配置。
	SnapshotPresetConfig UpdateSnapshotPresetV2BodySnapshotPresetConfig `json:"SnapshotPresetConfig"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图配置生效状态，默认为生效。
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

	// 当前格式的截图配置是否开启，默认为 false，取值及含义如下所示。
	// * false：不开启；
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

	// 截图是否使用 veImageX 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
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

	// 截图是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
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

	// REQUIRED; 需要配置限额的推流域名或拉流域名。
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

	// REQUIRED; 带宽限额，取值[1~1000]。
	Quota int32 `json:"Quota"`

	// REQUIRED; 拉流带宽限额的计量单位，支持的取值如下所示。
	// * Mbps
	// * Gbps
	// * Tbps
	QuotaUnit string `json:"QuotaUnit"`

	// 拉流带宽限额告警阈值，取值范围为 [1,1000]，缺省情况表示不设置告警。 :::tip 该参数的取值需要小于等于拉流带宽限额Quota，否则会报错。 :::
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

type UpdateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即观看时移的最长时间，单位为 s。支持的取值如下所示。
	// * 86400
	// * 259200
	// * 604800
	// * 1296000
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 开启时移的流名称，同一个 App 最多可指定 20 路。
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

	// REQUIRED; 转码配置的名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 音频编码格式，支持以下 3 种类型。
	// * aac：使用 AAC 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 Opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 宽高自适应模式开关，默认值为 0。支持的取值及含义如下。
	// * 0：关闭宽高自适应
	// * 1：开启宽高自适应 :::tip
	// * 关闭宽高自适应时，转码配置分辨率取 Width 和 Height 的值对转码视频进行拉伸；
	// * 开启宽高自适应时，转码配置分辨率按照 ShortSide 、 LongSide 、Width 、Height 的优先级取值，另一边等比缩放。
	// * 修改 As 为 0 时，请确认 Width 和 Height 的取值是否超出阈值；
	// * 修改 As 为 1 时，请确认 ShortSide 和 LongSide 的取值是否超出阈值。 :::
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

	// 2 个参考帧之间的最大 B 帧数，默认值为 3。配置不同的视频编码格式时，最大 B 帧数的取值存在如下差异。
	// * H.264：取值范围为 [0,7]；
	// * H.265：取值范围为 [0,3]、7、15；
	// * H.266：取值范围为 [0,3]、7、15。
	// 取值为 0 时，表示去除 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 动态范围，画质增强类型生效
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效
	// * 0：不开启
	// * 1：开启
	FISwitch *int32 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，帧率越大，画面越流畅。 配置不同视频编码格式时，视频帧率的取值存在如下差异。
	// * H.264：取值范围为 [0,60]；
	// * H.265：取值范围为 [0,60]；
	// * H.266：取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒，默认值为 4，取值范围为 [1,20]。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度，默认值为 0。配置不同视频编码格式时，视频高度的取值存在如下差异。
	// * H.264：取值范围为 [150,1920]；
	// * H.265：取值范围为 [150,1920]。
	// :::tip
	// * 当 As 的取值为 0 时，参数生效；反之则不生效；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，配置不同的视频编码方式和转码类型时，长边长度的取值范围存在如下差异。
	// * Roi 取 false 时： * H.264：取值范围为 0 和 [150,4096]；
	// * H.265：取值范围为 0 和 [150,7680]；
	// * H.266：取值范围为 0 和 [150,1280]。
	//
	//
	// * Roi 取 true 时： * H.264：取值范围为 0 和 [150,1920]；
	// * H.265：取值范围为 0 和 [150,1920]。 :::tip
	//
	//
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码模板参数的类型
	// * hvq：表示使用画质增强
	ParamType *string `json:"ParamType,omitempty"`

	// 是否极智超清转码，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。 :::tip
	// * 修改 Roi 为 true，且 As 为 1 时，请确认 ShortSide 和 LongSide 的取值是否超出阈值。
	// * 视频编码格式为 H.266 时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 使用场景，画质增强时生效
	// football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度，配置不同的视频编码方式和转码类型时，短边长度的取值范围存在如下差异。
	// * ParamType取 hvq 时： * H.264：取值范围为 0 和 [150,1280]；
	// * H.265：取值范围为 0 和 [150,1280]；
	//
	//
	// * Roi 取 false 时： * H.264：取值范围为 0 和 [150,2160]；
	// * H.265：取值范围为 0 和 [150,4096]；
	// * H.266：取值范围为 0 和 [150,720]。
	//
	//
	// * Roi 取 true 时： * H.264：取值范围为 0 和 [150,1920]；
	// * H.265：取值范围为 0 和 [150,1920]。 :::tip
	//
	//
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。支持 10 个字符以内的大小写字母、下划线与中划线，常见后缀包括：sd、hd、uhd。 例如，配置的转码流后缀名为hd，则拉转码流时转码的流名为 stream-123456789_hd。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，默认值为 Pull，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * h266：使用 H.266 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 bps，默认值为 1000000；取 0 时，表示使用源流码率。 配置不同的视频编码格式时，视频码率的取值范围存在如下差异。
	// * H.264：取值范围为 [0,30000000]；
	// * H.265：取值范围为 [0,30000000]；
	// * H.266：取值范围为 [0,6000000]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度，默认值为 0。配置不同视频编码格式时，视频宽度的取值存在如下差异。
	// * H.264：取值范围为 [150,1920]；
	// * H.265：取值范围为 [150,1920]。
	// :::tip
	// * 当 As 的取值为 0 即关闭宽高自适应时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * Width 和 Height 任一配置为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// 直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片链接，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片使用 data URI 协议，格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片字符串字段二选一传入。同时传入时，以水印图片字符串参数为准。 :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
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

type VerifyDomainOwnerBody struct {

	// REQUIRED; 推拉流域名
	Domain string `json:"Domain"`

	// REQUIRED; 校验方式，取值：
	// * dnsCheck：DNS验证。
	// * fileCheck：文件验证。
	VerifyType string `json:"VerifyType"`
}

type VerifyDomainOwnerRes struct {

	// REQUIRED
	ResponseMetadata VerifyDomainOwnerResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *VerifyDomainOwnerResResult `json:"Result,omitempty"`
}

type VerifyDomainOwnerResResponseMetadata struct {

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

// VerifyDomainOwnerResResult - 视请求的接口而定
type VerifyDomainOwnerResResult struct {

	// REQUIRED; 检查结果
	CheckResult bool `json:"CheckResult"`
}
type BindCert struct{}
type BindCertQuery struct{}
type CreateCert struct{}
type CreateCertQuery struct{}
type CreateDomain struct{}
type CreateDomainQuery struct{}
type CreateDomainV2 struct{}
type CreateDomainV2Query struct{}
type CreateLiveStreamRecordIndexFiles struct{}
type CreateLiveStreamRecordIndexFilesQuery struct{}
type CreatePullCDNSnapshotTask struct{}
type CreatePullCDNSnapshotTaskQuery struct{}
type CreatePullRecordTask struct{}
type CreatePullRecordTaskQuery struct{}
type CreatePullToPushTask struct{}
type CreatePullToPushTaskQuery struct{}
type CreateRecordPresetV2 struct{}
type CreateRecordPresetV2Query struct{}
type CreateRelaySourceV4 struct{}
type CreateRelaySourceV4Query struct{}
type CreateSnapshotAuditPreset struct{}
type CreateSnapshotAuditPresetQuery struct{}
type CreateSnapshotPreset struct{}
type CreateSnapshotPresetQuery struct{}
type CreateSnapshotPresetV2 struct{}
type CreateSnapshotPresetV2Query struct{}
type CreateTimeShiftPresetV3 struct{}
type CreateTimeShiftPresetV3Query struct{}
type CreateTranscodePreset struct{}
type CreateTranscodePresetQuery struct{}
type CreateVerifyContent struct{}
type CreateVerifyContentQuery struct{}
type CreateWatermarkPreset struct{}
type CreateWatermarkPresetQuery struct{}
type DeleteCallback struct{}
type DeleteCallbackQuery struct{}
type DeleteCert struct{}
type DeleteCertQuery struct{}
type DeleteDomain struct{}
type DeleteDomainQuery struct{}
type DeleteIPAccessRule struct{}
type DeleteIPAccessRuleQuery struct{}
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
type DeleteTimeShiftPresetV3 struct{}
type DeleteTimeShiftPresetV3Query struct{}
type DeleteTranscodePreset struct{}
type DeleteTranscodePresetQuery struct{}
type DeleteWatermarkPreset struct{}
type DeleteWatermarkPresetQuery struct{}
type DescribeAuth struct{}
type DescribeAuthQuery struct{}
type DescribeCDNSnapshotHistory struct{}
type DescribeCDNSnapshotHistoryQuery struct{}
type DescribeCallback struct{}
type DescribeCallbackQuery struct{}
type DescribeCertDetailSecretV2 struct{}
type DescribeCertDetailSecretV2Query struct{}
type DescribeClosedStreamInfoByPage struct{}
type DescribeClosedStreamInfoByPageBody struct{}
type DescribeDenyConfig struct{}
type DescribeDenyConfigQuery struct{}
type DescribeDomain struct{}
type DescribeDomainQuery struct{}
type DescribeForbiddenStreamInfoByPage struct{}
type DescribeForbiddenStreamInfoByPageBody struct{}
type DescribeIPAccessRule struct{}
type DescribeIPAccessRuleQuery struct{}
type DescribeIPInfo struct{}
type DescribeIPInfoQuery struct{}
type DescribeLiveActivityBandwidthData struct{}
type DescribeLiveActivityBandwidthDataQuery struct{}
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
type DescribeLiveBatchStreamTrafficData struct{}
type DescribeLiveBatchStreamTrafficDataQuery struct{}
type DescribeLiveBatchStreamTranscodeData struct{}
type DescribeLiveBatchStreamTranscodeDataQuery struct{}
type DescribeLiveCustomizedLogData struct{}
type DescribeLiveCustomizedLogDataQuery struct{}
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
type DescribeLivePlayStatusCodeData struct{}
type DescribeLivePlayStatusCodeDataQuery struct{}
type DescribeLivePullToPushBandwidthData struct{}
type DescribeLivePullToPushBandwidthDataQuery struct{}
type DescribeLivePullToPushData struct{}
type DescribeLivePullToPushDataQuery struct{}
type DescribeLivePushStreamCountData struct{}
type DescribeLivePushStreamCountDataQuery struct{}
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
type DescribeLiveStreamUsageData struct{}
type DescribeLiveStreamUsageDataQuery struct{}
type DescribeLiveTimeShiftData struct{}
type DescribeLiveTimeShiftDataQuery struct{}
type DescribeLiveTrafficData struct{}
type DescribeLiveTrafficDataQuery struct{}
type DescribeLiveTranscodeData struct{}
type DescribeLiveTranscodeDataQuery struct{}
type DescribeRecordTaskFileHistory struct{}
type DescribeRecordTaskFileHistoryQuery struct{}
type DescribeReferer struct{}
type DescribeRefererQuery struct{}
type DescribeRelaySourceV3 struct{}
type DescribeRelaySourceV3Query struct{}
type DescribeSnapshotAuditPresetDetail struct{}
type DescribeSnapshotAuditPresetDetailQuery struct{}
type DescribeStreamQuotaConfig struct{}
type DescribeStreamQuotaConfigQuery struct{}
type DisableDomain struct{}
type DisableDomainQuery struct{}
type EnableDomain struct{}
type EnableDomainQuery struct{}
type ForbidStream struct{}
type ForbidStreamQuery struct{}
type GeneratePlayURL struct{}
type GeneratePlayURLQuery struct{}
type GeneratePushURL struct{}
type GeneratePushURLQuery struct{}
type GetPullCDNSnapshotTask struct{}
type GetPullCDNSnapshotTaskQuery struct{}
type GetPullRecordTask struct{}
type GetPullRecordTaskQuery struct{}
type KillStream struct{}
type KillStreamQuery struct{}
type ListCertV2 struct{}
type ListCertV2Query struct{}
type ListCommonTransPresetDetail struct{}
type ListCommonTransPresetDetailQuery struct{}
type ListDomainDetail struct{}
type ListDomainDetailQuery struct{}
type ListPullCDNSnapshotTask struct{}
type ListPullCDNSnapshotTaskQuery struct{}
type ListPullRecordTask struct{}
type ListPullRecordTaskQuery struct{}
type ListPullToPushTask struct{}
type ListPullToPushTaskBody struct{}
type ListRelaySourceV4 struct{}
type ListRelaySourceV4Query struct{}
type ListTimeShiftPresetV2 struct{}
type ListTimeShiftPresetV2Query struct{}
type ListVhostRecordPresetV2 struct{}
type ListVhostRecordPresetV2Query struct{}
type ListVhostSnapshotAuditPreset struct{}
type ListVhostSnapshotAuditPresetQuery struct{}
type ListVhostSnapshotPreset struct{}
type ListVhostSnapshotPresetQuery struct{}
type ListVhostSnapshotPresetV2 struct{}
type ListVhostSnapshotPresetV2Query struct{}
type ListVhostTransCodePreset struct{}
type ListVhostTransCodePresetQuery struct{}
type ListVhostWatermarkPreset struct{}
type ListVhostWatermarkPresetQuery struct{}
type ListVqosMetricsDimensions struct{}
type ListVqosMetricsDimensionsBody struct{}
type ListWatermarkPreset struct{}
type ListWatermarkPresetQuery struct{}
type RestartPullToPushTask struct{}
type RestartPullToPushTaskQuery struct{}
type ResumeStream struct{}
type ResumeStreamQuery struct{}
type StopPullCDNSnapshotTask struct{}
type StopPullCDNSnapshotTaskQuery struct{}
type StopPullRecordTask struct{}
type StopPullRecordTaskQuery struct{}
type StopPullToPushTask struct{}
type StopPullToPushTaskQuery struct{}
type UnbindCert struct{}
type UnbindCertQuery struct{}
type UpdateAuthKey struct{}
type UpdateAuthKeyQuery struct{}
type UpdateCallback struct{}
type UpdateCallbackQuery struct{}
type UpdateDenyConfig struct{}
type UpdateDenyConfigQuery struct{}
type UpdateDomainVhost struct{}
type UpdateDomainVhostQuery struct{}
type UpdateIPAccessRule struct{}
type UpdateIPAccessRuleQuery struct{}
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
type UpdateSnapshotPreset struct{}
type UpdateSnapshotPresetQuery struct{}
type UpdateSnapshotPresetV2 struct{}
type UpdateSnapshotPresetV2Query struct{}
type UpdateStreamQuotaConfig struct{}
type UpdateStreamQuotaConfigQuery struct{}
type UpdateTimeShiftPresetV3 struct{}
type UpdateTimeShiftPresetV3Query struct{}
type UpdateTranscodePreset struct{}
type UpdateTranscodePresetQuery struct{}
type UpdateWatermarkPreset struct{}
type UpdateWatermarkPresetQuery struct{}
type VerifyDomainOwner struct{}
type VerifyDomainOwnerQuery struct{}
type BindCertReq struct {
	*BindCertQuery
	*BindCertBody
}
type CreateCertReq struct {
	*CreateCertQuery
	*CreateCertBody
}
type CreateDomainReq struct {
	*CreateDomainQuery
	*CreateDomainBody
}
type CreateDomainV2Req struct {
	*CreateDomainV2Query
	*CreateDomainV2Body
}
type CreateLiveStreamRecordIndexFilesReq struct {
	*CreateLiveStreamRecordIndexFilesQuery
	*CreateLiveStreamRecordIndexFilesBody
}
type CreatePullCDNSnapshotTaskReq struct {
	*CreatePullCDNSnapshotTaskQuery
	*CreatePullCDNSnapshotTaskBody
}
type CreatePullRecordTaskReq struct {
	*CreatePullRecordTaskQuery
	*CreatePullRecordTaskBody
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
type CreateSnapshotPresetReq struct {
	*CreateSnapshotPresetQuery
	*CreateSnapshotPresetBody
}
type CreateSnapshotPresetV2Req struct {
	*CreateSnapshotPresetV2Query
	*CreateSnapshotPresetV2Body
}
type CreateTimeShiftPresetV3Req struct {
	*CreateTimeShiftPresetV3Query
	*CreateTimeShiftPresetV3Body
}
type CreateTranscodePresetReq struct {
	*CreateTranscodePresetQuery
	*CreateTranscodePresetBody
}
type CreateVerifyContentReq struct {
	*CreateVerifyContentQuery
	*CreateVerifyContentBody
}
type CreateWatermarkPresetReq struct {
	*CreateWatermarkPresetQuery
	*CreateWatermarkPresetBody
}
type DeleteCallbackReq struct {
	*DeleteCallbackQuery
	*DeleteCallbackBody
}
type DeleteCertReq struct {
	*DeleteCertQuery
	*DeleteCertBody
}
type DeleteDomainReq struct {
	*DeleteDomainQuery
	*DeleteDomainBody
}
type DeleteIPAccessRuleReq struct {
	*DeleteIPAccessRuleQuery
	*DeleteIPAccessRuleBody
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
type DescribeCertDetailSecretV2Req struct {
	*DescribeCertDetailSecretV2Query
	*DescribeCertDetailSecretV2Body
}
type DescribeClosedStreamInfoByPageReq struct {
	*DescribeClosedStreamInfoByPageQuery
	*DescribeClosedStreamInfoByPageBody
}
type DescribeDenyConfigReq struct {
	*DescribeDenyConfigQuery
	*DescribeDenyConfigBody
}
type DescribeDomainReq struct {
	*DescribeDomainQuery
	*DescribeDomainBody
}
type DescribeForbiddenStreamInfoByPageReq struct {
	*DescribeForbiddenStreamInfoByPageQuery
	*DescribeForbiddenStreamInfoByPageBody
}
type DescribeIPAccessRuleReq struct {
	*DescribeIPAccessRuleQuery
	*DescribeIPAccessRuleBody
}
type DescribeIPInfoReq struct {
	*DescribeIPInfoQuery
	*DescribeIPInfoBody
}
type DescribeLiveActivityBandwidthDataReq struct {
	*DescribeLiveActivityBandwidthDataQuery
	*DescribeLiveActivityBandwidthDataBody
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
type DescribeLiveBatchStreamTrafficDataReq struct {
	*DescribeLiveBatchStreamTrafficDataQuery
	*DescribeLiveBatchStreamTrafficDataBody
}
type DescribeLiveBatchStreamTranscodeDataReq struct {
	*DescribeLiveBatchStreamTranscodeDataQuery
	*DescribeLiveBatchStreamTranscodeDataBody
}
type DescribeLiveCustomizedLogDataReq struct {
	*DescribeLiveCustomizedLogDataQuery
	*DescribeLiveCustomizedLogDataBody
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
type DescribeLiveStreamUsageDataReq struct {
	*DescribeLiveStreamUsageDataQuery
	*DescribeLiveStreamUsageDataBody
}
type DescribeLiveTimeShiftDataReq struct {
	*DescribeLiveTimeShiftDataQuery
	*DescribeLiveTimeShiftDataBody
}
type DescribeLiveTrafficDataReq struct {
	*DescribeLiveTrafficDataQuery
	*DescribeLiveTrafficDataBody
}
type DescribeLiveTranscodeDataReq struct {
	*DescribeLiveTranscodeDataQuery
	*DescribeLiveTranscodeDataBody
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
type DescribeSnapshotAuditPresetDetailReq struct {
	*DescribeSnapshotAuditPresetDetailQuery
	*DescribeSnapshotAuditPresetDetailBody
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
type GetPullCDNSnapshotTaskReq struct {
	*GetPullCDNSnapshotTaskQuery
	*GetPullCDNSnapshotTaskBody
}
type GetPullRecordTaskReq struct {
	*GetPullRecordTaskQuery
	*GetPullRecordTaskBody
}
type KillStreamReq struct {
	*KillStreamQuery
	*KillStreamBody
}
type ListCertV2Req struct {
	*ListCertV2Query
	*ListCertV2Body
}
type ListCommonTransPresetDetailReq struct {
	*ListCommonTransPresetDetailQuery
	*ListCommonTransPresetDetailBody
}
type ListDomainDetailReq struct {
	*ListDomainDetailQuery
	*ListDomainDetailBody
}
type ListPullCDNSnapshotTaskReq struct {
	*ListPullCDNSnapshotTaskQuery
	*ListPullCDNSnapshotTaskBody
}
type ListPullRecordTaskReq struct {
	*ListPullRecordTaskQuery
	*ListPullRecordTaskBody
}
type ListPullToPushTaskReq struct {
	*ListPullToPushTaskQuery
	*ListPullToPushTaskBody
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
type ListVhostSnapshotPresetReq struct {
	*ListVhostSnapshotPresetQuery
	*ListVhostSnapshotPresetBody
}
type ListVhostSnapshotPresetV2Req struct {
	*ListVhostSnapshotPresetV2Query
	*ListVhostSnapshotPresetV2Body
}
type ListVhostTransCodePresetReq struct {
	*ListVhostTransCodePresetQuery
	*ListVhostTransCodePresetBody
}
type ListVhostWatermarkPresetReq struct {
	*ListVhostWatermarkPresetQuery
	*ListVhostWatermarkPresetBody
}
type ListVqosMetricsDimensionsReq struct {
	*ListVqosMetricsDimensionsQuery
	*ListVqosMetricsDimensionsBody
}
type ListWatermarkPresetReq struct {
	*ListWatermarkPresetQuery
	*ListWatermarkPresetBody
}
type RestartPullToPushTaskReq struct {
	*RestartPullToPushTaskQuery
	*RestartPullToPushTaskBody
}
type ResumeStreamReq struct {
	*ResumeStreamQuery
	*ResumeStreamBody
}
type StopPullCDNSnapshotTaskReq struct {
	*StopPullCDNSnapshotTaskQuery
	*StopPullCDNSnapshotTaskBody
}
type StopPullRecordTaskReq struct {
	*StopPullRecordTaskQuery
	*StopPullRecordTaskBody
}
type StopPullToPushTaskReq struct {
	*StopPullToPushTaskQuery
	*StopPullToPushTaskBody
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
type UpdateDenyConfigReq struct {
	*UpdateDenyConfigQuery
	*UpdateDenyConfigBody
}
type UpdateDomainVhostReq struct {
	*UpdateDomainVhostQuery
	*UpdateDomainVhostBody
}
type UpdateIPAccessRuleReq struct {
	*UpdateIPAccessRuleQuery
	*UpdateIPAccessRuleBody
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
type UpdateSnapshotPresetReq struct {
	*UpdateSnapshotPresetQuery
	*UpdateSnapshotPresetBody
}
type UpdateSnapshotPresetV2Req struct {
	*UpdateSnapshotPresetV2Query
	*UpdateSnapshotPresetV2Body
}
type UpdateStreamQuotaConfigReq struct {
	*UpdateStreamQuotaConfigQuery
	*UpdateStreamQuotaConfigBody
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
type VerifyDomainOwnerReq struct {
	*VerifyDomainOwnerQuery
	*VerifyDomainOwnerBody
}

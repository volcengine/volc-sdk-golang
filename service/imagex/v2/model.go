package imagex

type AddImageElementsBodyType string

type CreateImageStyleBodyUnit string

type DeleteImageElementsBodyType string

type Enum0 string

type Enum1 string

type Enum2 string

type Enum4 string

type GetImageStylesResResultStylesItemUnit string

type GetImageTranscodeQueuesResResultQueuesItemStatus string

type GetImageTranscodeQueuesResResultQueuesItemType string
type AIProcessBody struct {
	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; AI 图像处理模板参数，需要将 JSON 压缩并转义为字符串。根据您需要的图像处理功能，参看 AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板
	// ID 和参数信息。
	WorkflowParameter string `json:"WorkflowParameter"`

	// REQUIRED; AI 图像处理模板 ID。根据您需要的图像处理功能，参看 AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板 ID 和参数信息。
	WorkflowTemplateID string `json:"WorkflowTemplateId"`
}

type AIProcessRes struct {
	// REQUIRED
	ResponseMetadata *AIProcessResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *AIProcessResResult `json:"Result"`
}

type AIProcessResResponseMetadata struct {
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

// AIProcessResResult - 视请求的接口而定
type AIProcessResResult struct {
	// REQUIRED; AI 图像处理结果，是 JSON 压缩并转义后的字符串。参看 AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板 ID 和参数信息，根据具体的工作流的说明进行解析。
	Output string `json:"Output"`
}

type AddDomainV1Body struct {

	// REQUIRED; 新增域名。
	Domain string `json:"domain"`

	// 访问控制配置。
	AccessControl []*AddDomainV1BodyAccessControlItem `json:"access_control,omitempty"`

	// 证书配置，海外加速或者全球加速为必选，否则审核不通过。
	HTTPS []*AddDomainV1BodyHTTPSItem `json:"https,omitempty"`

	// 请求需要添加的响应头。
	RespHdrs []*AddDomainV1BodyRespHdrsItem `json:"resp_hdrs,omitempty"`
}

type AddDomainV1BodyAccessControlItem struct {

	// Referer 配置。
	ReferLink []*AddDomainV1BodyAccessControlPropertiesItemsItem `json:"refer_link,omitempty"`
}

type AddDomainV1BodyAccessControlPropertiesItemsItem struct {

	// REQUIRED; 是否开启 Referer 防盗链，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"enabled"`

	// 是否允许空 Referer，取值如下所示：
	// * true：允许空 Referer
	// * false：不允许空 Referer
	AllowEmptyRefer bool `json:"allow_empty_refer,omitempty"`

	// 是否配置白名单，取值如下所示：
	// * true：配置白名单
	// * false：配置黑名单
	IsWhiteMode bool `json:"is_white_mode,omitempty"`

	// 根据是否为白名单，为对应的白/黑名单的值。
	Values []string `json:"values,omitempty"`
}

type AddDomainV1BodyHTTPSItem struct {

	// 证书 ID，若enable_https为true，则为必选。
	CertID string `json:"cert_id,omitempty"`

	// 是否开启 Https，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableHTTPS bool `json:"enable_https,omitempty"`

	// 是否强制使用 Https，取值如下所示：
	// * true：强制
	// * false：不强制
	ForceHTTPS bool `json:"force_https,omitempty"`
}

type AddDomainV1BodyRespHdrsItem struct {

	// REQUIRED; Header Key，请见支持配置的响应头 [https://www.volcengine.com/docs/508/196704#%E6%94%AF%E6%8C%81%E9%85%8D%E7%BD%AE%E7%9A%84%E5%93%8D%E5%BA%94%E5%A4%B4]。
	Key string `json:"key"`

	// Header Value，设置该响应头字段的值。字段值不能超过 1,024 个字符，可以包含除美元符号（$），Delete（ASCII code 127）外的可打印 ASCII 字符。
	Value string `json:"value,omitempty"`
}

type AddDomainV1Query struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type AddDomainV1Res struct {

	// REQUIRED
	ResponseMetadata *AddDomainV1ResResponseMetadata `json:"ResponseMetadata"`
	Result           *AddDomainV1ResResult           `json:"Result,omitempty"`
}

type AddDomainV1ResResponseMetadata struct {

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

type AddDomainV1ResResult struct {

	// 新增的域名。
	Domain string `json:"Domain,omitempty"`
}

type AddImageBackgroundColorsBody struct {

	// REQUIRED; 待添加的颜色列表
	Colors []string `json:"Colors"`
}

type AddImageBackgroundColorsRes struct {

	// REQUIRED
	ResponseMetadata *AddImageBackgroundColorsResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *AddImageBackgroundColorsResResult `json:"Result,omitempty"`
}

type AddImageBackgroundColorsResResponseMetadata struct {

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

// AddImageBackgroundColorsResResult - title
type AddImageBackgroundColorsResResult struct {

	// REQUIRED; 若全部添加失败，则接口将返回失败
	FailedList []string `json:"FailedList"`
}

type AddImageElementsBody struct {

	// REQUIRED; 待添加的图片 URI 列表。
	Images []string `json:"Images"`

	// REQUIRED; 取值image表示图片要素，background表示背景要素
	Type AddImageElementsBodyType `json:"Type"`
}

type AddImageElementsRes struct {

	// REQUIRED
	ResponseMetadata *AddImageElementsResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *AddImageElementsResResult `json:"Result,omitempty"`
}

type AddImageElementsResResponseMetadata struct {

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

// AddImageElementsResResult - title
type AddImageElementsResResult struct {

	// REQUIRED; 若全部添加失败，则接口将返回失败
	FailedList []string `json:"FailedList"`
}

type ApplyImageUploadQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 文件扩展名(如：.java, .txt, .go 等)，最大长度限制为 8 个字节。 :::tip 仅当未指定StoreKeys时生效。 :::
	FileExtension string `json:"FileExtension,omitempty" query:"FileExtension"`

	// 是否开启重名文件覆盖上传，取值如下所示：
	// * true：开启
	// * false：（默认）关闭
	// :::warning 在指定Overwrite为true前，请确保您指定的ServiceId对应服务已开启了覆盖上传 [https://www.volcengine.com/docs/508/1119912]能力。 :::
	Overwrite bool `json:"Overwrite,omitempty" query:"Overwrite"`

	// 指定的上传文件路径。
	// * 指定Prefix时，下发的存储 Key 为：Prefix/{随机Key}{FileExtension}，其中Prefix + FileExtension最大长度限制为 145 个字节。
	// * 不支持以/开头或结尾，不支持/连续出现。 :::tip 仅当未指定StoreKeys时生效。 :::
	Prefix string `json:"Prefix,omitempty" query:"Prefix"`

	// 一次上传会话 Key。 :::tip 本接口上一次请求的SessionKey，可在重试时带上，作为服务端的再次选路时的一个参考。 :::
	SessionKey string `json:"SessionKey,omitempty" query:"SessionKey"`

	// 上传文件的存储 Key。默认使用随机生成的字符串作为存储 Key。
	// * 数组长度和UploadNum保持一致。
	// * 存储 Key 详细命名规范请参看 veImageX 存储 Key 通用字符规则 [https://www.volcengine.com/docs/508/1458331]。
	StoreKeys []string `json:"StoreKeys,omitempty" query:"StoreKeys"`

	// 上传文件的数量，将决定下发的 StoreUri 的数量，取值范围为[1,10]，默认为 1。
	UploadNum int `json:"UploadNum,omitempty" query:"UploadNum"`
}

type ApplyImageUploadRes struct {

	// REQUIRED
	ResponseMetadata *ApplyImageUploadResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *ApplyImageUploadResResult `json:"Result"`
}

type ApplyImageUploadResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                    `json:"Version"`
	Error   *ApplyImageUploadResResponseMetadataError `json:"Error,omitempty"`
}

type ApplyImageUploadResResponseMetadataError struct {

	// REQUIRED; 错误码
	Code string `json:"Code"`

	// REQUIRED; 错误信息
	Message string `json:"Message"`
}

type ApplyImageUploadResResult struct {

	// REQUIRED; 请求的唯一标识 ID。
	RequestID string `json:"RequestId"`

	// REQUIRED; 上传地址
	UploadAddress *ApplyImageUploadResResultUploadAddress `json:"UploadAddress"`
}

// ApplyImageUploadResResultUploadAddress - 上传地址
type ApplyImageUploadResResultUploadAddress struct {

	// REQUIRED; 一次上传会话 Key。 上传完成上报时使用该值，该 Key 可以在解码后提取信息及参数校验。
	SessionKey string `json:"SessionKey"`

	// REQUIRED; 上传 Uri
	// :::tip 数量由请求参数中的 UploadNum 决定。 :::
	StoreInfos []*ApplyImageUploadResResultUploadAddressStoreInfosItem `json:"StoreInfos"`

	// REQUIRED; 上传域名列表，可以用于客户端容灾，并行上传等。
	UploadHosts []string `json:"UploadHosts"`
}

type ApplyImageUploadResResultUploadAddressStoreInfosItem struct {

	// REQUIRED; 上传凭证。
	Auth string `json:"Auth"`

	// REQUIRED; 资源 ID。
	StoreURI string `json:"StoreUri"`
}

type ApplyVpcUploadInfoQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 上传文件的 Content-Type 值。
	// 需确保指定值在服务维度的白名单内，否则无法成功上传，参看上传 Content-Type 限制 [https://www.volcengine.com/docs/508/1319948]。
	ContentType string `json:"ContentType,omitempty" query:"ContentType"`

	// 文件扩展名，最大长度限制为 20 个字节。
	// :::tip 仅当未指定 StoreKeys 时生效。 :::
	FileExtension string `json:"FileExtension,omitempty" query:"FileExtension"`

	// 文件大小。
	FileSize int `json:"FileSize,omitempty" query:"FileSize"`

	// 是否开启重名文件覆盖上传，取值如下所示：
	// * true：开启
	// * false：（默认）关闭
	// :::tip 在指定Overwrite为true前，请确保您指定的ServiceId对应服务已开启了覆盖上传 [https://www.volcengine.com/docs/508/1119912]能力。 :::
	Overwrite bool `json:"Overwrite,omitempty" query:"Overwrite"`

	// 分片大小，单位为字节，默认值为 200 MB。 当 FileSize 大于 PartSize 时，下发分片上传的 URL。
	PartSize int `json:"PartSize,omitempty" query:"PartSize"`

	// 指定的上传文件路径。指定 Prefix 时，下发的存储 Key 为：Prefix/{随机Key}.{FileExtension}，拼接形成的存储 Key 需满足veImageX 存储 Key 通用字符规则 [https://www.volcengine.com/docs/508/1458331]。
	// :::tip 仅当未指定 StoreKeys 时生效。 :::
	Prefix string `json:"Prefix,omitempty" query:"Prefix"`

	// 存储类型。
	// * STANDARD：标准存储
	// * IA：低频存储
	// * ARCHIVE_FR：归档闪回存储
	// * ARCHIVE：归档存储
	// * COLD_ARCHIVE：冷归档存储
	StorageClass string `json:"StorageClass,omitempty" query:"StorageClass"`

	// 上传文件的存储 Key。默认使用随机生成的字符串作为存储 Key。
	// 存储 Key 详细命名规范请参看veImageX 存储 Key 通用字符规则 [https://www.volcengine.com/docs/508/1458331]。
	StoreKey string `json:"StoreKey,omitempty" query:"StoreKey"`
}

type ApplyVpcUploadInfoRes struct {

	// REQUIRED
	ResponseMetadata *ApplyVpcUploadInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *ApplyVpcUploadInfoResResult           `json:"Result,omitempty"`
}

type ApplyVpcUploadInfoResResponseMetadata struct {

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

type ApplyVpcUploadInfoResResult struct {

	// REQUIRED; 参数的唯一标识符
	Oid string `json:"Oid"`

	// REQUIRED; 分片上传信息。
	PartUploadInfo *ApplyVpcUploadInfoResResultPartUploadInfo `json:"PartUploadInfo"`

	// REQUIRED; 上传URL。
	PutURL string `json:"PutURL"`

	// REQUIRED; 自定义请求头。
	PutURLHeaders []*ApplyVpcUploadInfoResResultPutURLHeadersItem `json:"PutURLHeaders"`

	// REQUIRED; 上传会话
	SessionKey string `json:"SessionKey"`

	// REQUIRED; 上传模式。
	// * direct：直接上传
	// * part：分片上传
	UploadMode string `json:"UploadMode"`
}

// ApplyVpcUploadInfoResResultPartUploadInfo - 分片上传信息。
type ApplyVpcUploadInfoResResultPartUploadInfo struct {

	// REQUIRED; 完整部分的URL。
	CompletePartURL string `json:"CompletePartURL"`

	// REQUIRED; 合并分片请求头信息，map结构
	CompletePartURLHeaders []*ApplyVpcUploadInfoResResultPartUploadInfoCompletePartURLHeadersItem `json:"CompletePartURLHeaders"`

	// REQUIRED; 分块上传URL列表。
	PartPutURLs []string `json:"PartPutURLs"`

	// REQUIRED; 上传分片大小，默认为 200MB。
	PartSize int `json:"PartSize"`
}

type ApplyVpcUploadInfoResResultPartUploadInfoCompletePartURLHeadersItem struct {

	// REQUIRED; 请提供具体的Key和string值，我将根据这些信息生成参数的一句话描述。
	Key string `json:"Key"`

	// REQUIRED; 请提供参数的名字和类型。
	Value string `json:"Value"`
}

type ApplyVpcUploadInfoResResultPutURLHeadersItem struct {

	// REQUIRED; 请提供具体的Key和string值，我将根据这些信息生成参数描述。
	Key string `json:"Key"`

	// REQUIRED; 参数的值。
	Value string `json:"Value"`
}

type CommitImageUploadBody struct {

	// REQUIRED; 一次上传会话 Key。您可参考获取文件上传地址 [https://www.volcengine.com/docs/508/9397]获取。
	SessionKey  string   `json:"SessionKey"`
	DecryptKeys []string `json:"DecryptKeys,omitempty"`

	// 上传成功的资源 ID。
	SuccessOids []string `json:"SuccessOids,omitempty"`
}

type CommitImageUploadQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 是否返回图片 meta 信息。
	// * true：不返回图片 meta 信息。
	// * false：（默认）获取图片 meta 信息并返回对应 meta 信息。
	// :::tip
	// * 其中若 meta 获取超时或失败，接口返回成功，但对应 meta 信息将为空。
	// * 如果您的业务要求必须获取 meta 信息，请您参考图片Meta信息 [https://www.volcengine.com/docs/508/64085]获取。 :::
	SkipMeta bool `json:"SkipMeta,omitempty" query:"SkipMeta"`
}

type CommitImageUploadRes struct {

	// REQUIRED
	ResponseMetadata *CommitImageUploadResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CommitImageUploadResResult `json:"Result"`
}

type CommitImageUploadResResponseMetadata struct {

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

type CommitImageUploadResResult struct {

	// REQUIRED; 请求的唯一标识 ID。
	RequestID string `json:"RequestId"`

	// REQUIRED; 运行结果，数组长度对应上传的数量。
	Results []*CommitImageUploadResResultResultsItem `json:"Results"`

	// JSON 序列化之后的图片信息。
	PluginResult []*CommitImageUploadResResultPluginResultItem `json:"PluginResult,omitempty"`
}

type CommitImageUploadResResultPluginResultItem struct {

	// REQUIRED; 图片 Uri。
	ImageURI string `json:"ImageUri"`

	// 图片时长，单位为 ms。仅当图片为动图时有值
	Duration int `json:"Duration,omitempty"`

	// 图片文件名。
	FileName string `json:"FileName,omitempty"`

	// 图片的帧数量。
	FrameCnt int `json:"FrameCnt,omitempty"`

	// 图片格式。
	ImageFormat string `json:"ImageFormat,omitempty"`

	// 图片的高。
	ImageHeight int `json:"ImageHeight,omitempty"`

	// 图片的 MD5 哈希值。
	ImageMD5 string `json:"ImageMd5,omitempty"`

	// 图片的大小。
	ImageSize int `json:"ImageSize,omitempty"`

	// 图片的宽。
	ImageWidth int `json:"ImageWidth,omitempty"`

	// 源文件 URI
	SourceURI string `json:"SourceUri,omitempty"`
}

type CommitImageUploadResResultResultsItem struct {

	// REQUIRED; 源图片的 Uri。
	URI string `json:"Uri"`

	// REQUIRED; 上传结果。
	// * 传入 SuccessOids 时，无论上传图片/非图片，成功返回 2000，失败返回 2001；
	// * 未传入 SuccessOids 时，对于非图片返回 0；对于图片，成功返回 2000，失败返回 2001。
	URIStatus int `json:"UriStatus"`

	// 加密结果。 :::tip 指定了 Encryption Function 时有值 。 :::
	Encryption *CommitImageUploadResResultResultsItemEncryption `json:"Encryption,omitempty"`
	ImageMeta  *CommitImageUploadResResultResultsItemImageMeta  `json:"ImageMeta,omitempty"`
}

// CommitImageUploadResResultResultsItemEncryption - 加密结果。 :::tip 指定了 Encryption Function 时有值 。 :::
type CommitImageUploadResResultResultsItemEncryption struct {

	// REQUIRED; 加密算法。
	Algorithm string `json:"Algorithm"`

	// REQUIRED; 额外信息，键值均为 String。
	Extra map[string]string `json:"Extra"`

	// REQUIRED; 加密私钥。
	SecretKey string `json:"SecretKey"`

	// REQUIRED; 加密源 MD5。
	SourceMD5 string `json:"SourceMd5"`

	// REQUIRED; 加密图片的 Uri。
	URI string `json:"Uri"`

	// REQUIRED; 加密版本。
	Version string `json:"Version"`
}

type CommitImageUploadResResultResultsItemImageMeta struct {

	// REQUIRED
	Field75 int `json:"Field75"`

	// REQUIRED
	Format string `json:"Format"`

	// REQUIRED
	Height int `json:"Height"`

	// REQUIRED
	MD5 string `json:"Md5"`

	// REQUIRED
	Size int `json:"Size"`

	// REQUIRED
	URI string `json:"Uri"`
}

type Components10W6HmvSchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesFillptnPropertiesViewpoint struct {
	Height int `json:"height"`

	Width int `json:"width"`

	X int `json:"x"`

	Y int `json:"y"`
}

type Components11LotgnSchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesFillptn struct {
	Name string `json:"name"`

	Param     *ComponentsEyl12ZSchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesFillptnPropertiesParam      `json:"param"`
	BgColor   string                                                                                                                                                  `json:"bgColor,omitempty"`
	Ptn       int                                                                                                                                                     `json:"ptn,omitempty"`
	ViewLoc   int                                                                                                                                                     `json:"viewLoc,omitempty"`
	Viewpoint *Components10W6HmvSchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesFillptnPropertiesViewpoint `json:"viewpoint,omitempty"`
}

type Components16Kv6ElSchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesEventrulesItems struct {
	CallbackURL string `json:"CallbackUrl"`

	Enable bool `json:"Enable"`

	EventType []string `json:"EventType"`

	ID string `json:"Id"`

	MatchRule string `json:"MatchRule"`
}

type Components18K1KvdSchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesResourcetagsItems struct {
	Key string `json:"Key"`

	Value string `json:"Value"`
}

type Components1AkatydSchemasGetimageserviceresPropertiesResultPropertiesStoragerulesv2ItemsPropertiesNoncurrentversiontransitionsItems struct {
	StorageClass string `json:"StorageClass"`

	NonCurrentDate string `json:"NonCurrentDate,omitempty"`

	NonCurrentDays int `json:"NonCurrentDays,omitempty"`
}

type Components1D40MkcSchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesStoragerulesItems struct {
	Enable bool `json:"Enable"`

	Event string `json:"Event"`

	Prefix string `json:"Prefix"`

	Action string `json:"Action,omitempty"`

	Day int `json:"Day,omitempty"`

	NonCurrentAction string `json:"NonCurrentAction,omitempty"`

	NonCurrentDay int `json:"NonCurrentDay,omitempty"`
}

// Components1Muxqr1SchemasGetimagealertrecordsresPropertiesResultPropertiesAlertrecordsItemsPropertiesAlertcondPropertiesAlertcontent
// - 各指标告警信息
type Components1Muxqr1SchemasGetimagealertrecordsresPropertiesResultPropertiesAlertrecordsItemsPropertiesAlertcondPropertiesAlertcontent struct {
	AggrInterval int `json:"AggrInterval"`

	Dim string `json:"Dim"`

	Func string `json:"Func"`

	Item string `json:"Item"`

	Op string `json:"Op"`

	RepeatCnt int `json:"RepeatCnt"`

	Threshold float64 `json:"Threshold"`

	Vals []*GetImageAlertRecordsResResultAlertRecordsPropertiesItemsItem `json:"Vals"`
}

type Components1T23IneSchemasGetallimagetemplatesresPropertiesResultPropertiesTemplatesItemsPropertiesFiltersItems struct {
	Name string `json:"Name"`

	Param map[string]interface{} `json:"Param"`
}

type Components1WypgicSchemasGetimagemonitorrulesresPropertiesResultPropertiesMonitorrulesItemsPropertiesFilterPropertiesDimfilterItems struct {
	Dim string `json:"Dim"`

	Vals []string `json:"Vals"`

	Not bool `json:"Not,omitempty"`
}

type Components1Xh7Lz4SchemasDescribeimagexcompressusageresPropertiesResultPropertiesCompressdataItemsPropertiesOutsizeItems struct {
	TimeStamp string `json:"TimeStamp"`

	Value float64 `json:"Value"`
}

// Components3Ic6Z9SchemasGetimagemonitorrulesresPropertiesResultPropertiesMonitorrulesItemsPropertiesCondPropertiesItemcondItems
// - 监控规则配置
type Components3Ic6Z9SchemasGetimagemonitorrulesresPropertiesResultPropertiesMonitorrulesItemsPropertiesCondPropertiesItemcondItems struct {
	AggrInterval int `json:"AggrInterval"`

	Func string `json:"Func"`

	Item string `json:"Item"`

	Op string `json:"Op"`

	RepeatCnt int `json:"RepeatCnt"`

	Threshold float64 `json:"Threshold"`

	CntThreshold int `json:"CntThreshold,omitempty"`
}

type Components6Rlnt3SchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesBorder struct {
	Color string `json:"color"`

	Dash int `json:"dash"`

	PaddingBottom int `json:"paddingBottom"`

	PaddingLeft int `json:"paddingLeft"`

	PaddingRight int `json:"paddingRight"`

	PaddingTop int `json:"paddingTop"`

	Radius int `json:"radius"`

	Weight int `json:"weight"`
}

type Components8GbgysSchemasGettemplatesfrombinresPropertiesResultPropertiesTemplatesItemsPropertiesFiltersItems struct {
	Name string `json:"Name"`

	Param map[string]interface{} `json:"Param"`
}

type ComponentsDgbrljSchemasGetimageaudittasksresPropertiesResultPropertiesTasksItemsPropertiesAudittaskPropertiesImageinfosItems struct {
	DataID string `json:"DataId"`

	ImageURI string `json:"ImageUri"`
}

type ComponentsDrd6S5SchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesStorageItems struct {
	AllTypes bool `json:"AllTypes"`

	BktName string `json:"BktName"`

	TTL int `json:"TTL"`
}

type ComponentsEyl12ZSchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesFillptnPropertiesParam struct {
	Color string `json:"color"`
}

type ComponentsK7Ou2VSchemasGetimageocrv2ResPropertiesResultPropertiesLicenseresultAdditionalproperties struct {
	Content string `json:"Content"`

	Location []int `json:"Location"`
}

type ComponentsPqmsj3SchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesMirrorItems struct {
	Headers map[string]string `json:"Headers"`

	Host string `json:"Host"`

	Hosts map[string]int `json:"Hosts"`

	Origin *GetAllImageServicesResResultServicesItemMirrorItemOrigin `json:"Origin"`

	Schema string `json:"Schema"`

	Source string `json:"Source"`
}

type ComponentsQpab6RSchemasUpdatestoragerulesv2BodyPropertiesStoragerulesItemsPropertiesNoncurrentversiontransitionsItems struct {
	StorageClass string `json:"StorageClass"`

	NonCurrentDate string `json:"NonCurrentDate,omitempty"`

	NonCurrentDays int `json:"NonCurrentDays,omitempty"`
}

type CreateBatchProcessTaskBody struct {

	// REQUIRED; 指定服务下待批量处理的资源链接信息
	BatchingInfo []*CreateBatchProcessTaskBodyBatchingInfoItem `json:"BatchingInfo"`

	// 回调地址，用于接收返回的回调信息。
	Callback string `json:"Callback,omitempty"`

	// 自定义回调内容，取值需要符合CallbackBodyType指定格式。
	CallbackBody string `json:"CallbackBody,omitempty"`

	// 回调内容格式。默认为空，若需指定CallbackBody时，也需同时指定CallbackBodyType的值。取值如下所示：
	// * application/json
	// * application/x-www-form-urlencoded
	CallbackBodyType string `json:"CallbackBodyType,omitempty"`
}

type CreateBatchProcessTaskBodyBatchingInfoItem struct {

	// 批处理能力，取值如下所示：
	// * meta：获取资源元信息
	// * preload：源站图片预热 :::warning 如需批量预热源站图片，请提交工单 [https://console.volcengine.com/ticket/createTicketV2/?step=3&Service=rtc&FlowKey=NGnOHeWkbeCrEAkrNvjT]联系技术支持开启。
	// :::
	Action string `json:"Action,omitempty"`

	// 指定服务下待批处理资源的可访问 URL
	URL string `json:"Url,omitempty"`
}

type CreateBatchProcessTaskQuery struct {

	// REQUIRED; 待执行异步批处理的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateBatchProcessTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateBatchProcessTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateBatchProcessTaskResResult           `json:"Result,omitempty"`
}

type CreateBatchProcessTaskResResponseMetadata struct {

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

type CreateBatchProcessTaskResResult struct {

	// REQUIRED; 任务 ID
	TaskID string `json:"TaskId"`
}

type CreateCVImageGenerateTaskBody struct {

	// REQUIRED; 服务下绑定的域名，域名状态需正常可用。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取绑定的域名信息。
	// * 您也可以通过 OpenAPI 的方式获取域名，具体请参考获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]。
	Domain string `json:"Domain"`

	// REQUIRED; 模型接口action
	ModelAction string `json:"ModelAction"`

	// REQUIRED; 文生图系列模型的接口 Version 名称。
	// 例如，使用通用 2.0S-文生图异步 [https://www.volcengine.com/docs/6791/1347773]，则 ModelVersion 需要取值为 2022-08-31。
	ModelVersion string `json:"ModelVersion"`

	// REQUIRED; 参数输出。
	Outputs []string `json:"Outputs"`

	// REQUIRED; 请求的JSON字符串。
	ReqJSON map[string]interface{} `json:"ReqJson"`

	// REQUIRED; 服务下创建的图片处理模板名称，指定后，将按照模板中的处理配置对生成的原始图片进行图片处理。
	// 您可在 veImageX 控制台的处理配置页面，参考新建模板 [https://www.volcengine.com/docs/508/8087]配置模板并获取模版名称，例如 tplv-f0****5k-test。
	Template string `json:"Template"`

	// 是否覆盖现有内容。 * false：不覆盖；
	// * true：覆盖。
	// 默认值为false。
	Overwrite bool `json:"Overwrite,omitempty"`
}

type CreateCVImageGenerateTaskQuery struct {

	// REQUIRED; 指定存储结果图并计量计费的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateCVImageGenerateTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateCVImageGenerateTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateCVImageGenerateTaskResResult `json:"Result,omitempty"`
}

type CreateCVImageGenerateTaskResResponseMetadata struct {

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

// CreateCVImageGenerateTaskResResult - 视请求的接口而定
type CreateCVImageGenerateTaskResResult struct {

	// REQUIRED; 响应的JSON数据。
	RespJSON map[string]interface{} `json:"RespJson"`

	// REQUIRED; 任务 ID，用于查询结果。
	TaskID string `json:"TaskId"`
}

type CreateFileRestoreBody struct {

	// REQUIRED; 恢复时长，取值范围为[1,365]，单位为天。
	Duration int `json:"Duration"`

	// REQUIRED; 文件存储 URI。
	// * 您可以在 veImageX 控制台资源管理 [https://console.volcengine.com/imagex/resource_manage/]页面，在已上传文件的名称列获取资源 URI。
	// * 您也可以通过 OpenAPI 的方式获取 URI，具体请参考获取服务下全部上传文件 [https://www.volcengine.com/docs/508/9393]。
	StoreURI string `json:"StoreUri"`

	// 取回方式： Expedited：快速取回 Standard：标准取回 Bulk：批量取回；不设置默认standard
	Tier string `json:"Tier,omitempty"`
}

type CreateFileRestoreQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateFileRestoreRes struct {

	// REQUIRED
	ResponseMetadata *CreateFileRestoreResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateFileRestoreResResponseMetadata struct {

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

type CreateHiddenWatermarkImageBody struct {

	// REQUIRED; 盲水印模型，取值如下所示：
	// * tracev1：前景图层水印模型（纯色背景适用）。
	//
	// 该模型可以生成带有水印的透明图像，但仅适用于纯色网页泄露溯源场景。该模型可有效抵抗常见的社交软件传播。然而，该算法对页面背景色的影响相对较大，因此不适合用于保护多彩页面或图片，例如商品页面。
	//
	//
	// * tracev2：前景图层水印模型（彩色背景通用）
	//
	// 该模型可以生成含水印的透明图像，主要应用在前端页面截图泄露溯源场景。该模型生成的水印纹理密集，在正常界面添加后肉眼基本不可见（截图放大后存在肉眼可见的水印纹理），可抵抗常见的社交软件传播。
	//
	// :::tip 此模型建议在 PC 端使用，移动端使用视觉效果较差。 :::
	//
	//
	// * tracev2-app：前景图层水印模型（移动端）
	//
	// 该模型可以生成含水印的透明图像，主要应用在前端页面截图泄露溯源场景。该模型生成的水印纹理密集，在正常界面添加后肉眼基本不可见（截图放大后存在肉眼可见的水印纹理），可抵抗常见的社交软件传播。
	//
	// :::tip 此模型建议在移动端使用，PC 端使用视觉效果较差。 :::
	Algorithm string `json:"Algorithm"`

	// REQUIRED; 自定义盲水印文本信息，最多支持 128 个字符。
	Info string `json:"Info"`

	// REQUIRED; 盲水印强度，取值如下所示：
	// * low：低强度，适用于对图像质量要求高。
	// * medium：中强度
	// * strong：高强度，适合图像纹理丰富时使用。
	Strength string `json:"Strength"`
}

type CreateHiddenWatermarkImageQuery struct {

	// REQUIRED; 用于存储结果图和计量计费的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateHiddenWatermarkImageRes struct {

	// REQUIRED
	ResponseMetadata *CreateHiddenWatermarkImageResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateHiddenWatermarkImageResResult `json:"Result,omitempty"`
}

type CreateHiddenWatermarkImageResResponseMetadata struct {

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

// CreateHiddenWatermarkImageResResult - 视请求的接口而定
type CreateHiddenWatermarkImageResResult struct {

	// REQUIRED; 盲水印图片 Uri，当前仅支持输出 png 格式。
	StoreURI string `json:"StoreUri"`
}

type CreateHmExtractTaskBody struct {

	// REQUIRED; 取值支持tracev2-app, tracev2
	Algorithm string `json:"Algorithm"`

	// REQUIRED; 水印强度取值low，medium，strong。推荐取值strong
	Strength string `json:"Strength"`

	// 任务回调地址，回调内容详见盲水印提取回调 [https://www.volcengine.com/docs/508/1554763]。
	Callback string `json:"Callback,omitempty"`

	// 资源key
	ImageURI string `json:"ImageUri,omitempty"`

	// 资源URL，与ImageUri必填其中一个。两者都填以ImageUri为准
	ImageURL string `json:"ImageUrl,omitempty"`
}

type CreateHmExtractTaskQuery struct {

	// REQUIRED; 待提取水印图对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateHmExtractTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateHmExtractTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateHmExtractTaskResResult `json:"Result,omitempty"`
}

type CreateHmExtractTaskResResponseMetadata struct {

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

// CreateHmExtractTaskResResult - 视请求的接口而定
type CreateHmExtractTaskResResult struct {

	// REQUIRED; 异步任务ID
	TaskID string `json:"TaskId"`
}

type CreateImageAIProcessCallbackBody struct {

	// REQUIRED; 条目id
	EntryID string `json:"EntryId"`
}

type CreateImageAIProcessCallbackRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageAIProcessCallbackResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateImageAIProcessCallbackResResponseMetadata struct {

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

type CreateImageAIProcessQueueBody struct {

	// REQUIRED; 队列属性
	Attribute string `json:"Attribute"`

	// REQUIRED; 队列名
	Name string `json:"Name"`

	// 任务队列回调设置。
	CallbackConf *CreateImageAIProcessQueueBodyCallbackConf `json:"CallbackConf,omitempty"`

	// 队列描述
	Desc string `json:"Desc,omitempty"`

	// 是否开启队列
	IsStart bool `json:"IsStart,omitempty"`
}

// CreateImageAIProcessQueueBodyCallbackConf - 任务队列回调设置。
type CreateImageAIProcessQueueBodyCallbackConf struct {

	// REQUIRED; 回调数据格式，JSON或者XML
	DataFormat string `json:"DataFormat"`

	// REQUIRED; 回调url
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 回到url协议，默认http
	Method string `json:"Method"`

	// REQUIRED; 回调类型，支持task 或entry
	Type string `json:"Type"`

	// 参数
	Args string `json:"Args,omitempty"`
}

type CreateImageAIProcessQueueRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageAIProcessQueueResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateImageAIProcessQueueResResult `json:"Result,omitempty"`
}

type CreateImageAIProcessQueueResResponseMetadata struct {

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

// CreateImageAIProcessQueueResResult - 视请求的接口而定
type CreateImageAIProcessQueueResResult struct {

	// REQUIRED; 队列id
	QueueID string `json:"QueueId"`
}

type CreateImageAITaskBody struct {

	// REQUIRED; 待进行 AI 处理的图片 URI 或 URL 列表，其中 URI 不需要带 tos-cn-i-*** 前缀，最多传入 10,000 张图片。传入图片的宽高、大小等要求请参看对应的附加组件使用限制 [https://www.volcengine.com/docs/508/1270839]。
	// :::warning 若 DataType 取值 uri，则待转码图片 URI 必须为指定服务 ID 下的存储 URI。您可通过调用GetImageUploadFiles [https://www.volcengine.com/docs/508/9392]获取指定服务下全部的上传文件存储
	// URI。 :::
	DataList []string `json:"DataList"`

	// REQUIRED; 需要提交的图片数据类型，取值如下所示：
	// * uri：指定 ServiceId 下存储 URI。
	// * url：公网可访问的 URL。
	DataType string `json:"DataType"`

	// REQUIRED
	QueueID string `json:"QueueId"`

	// REQUIRED; 服务 ID。若 DataType 取值 uri，则提交的图片 URI 列表需在该服务内。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; AI 图像处理模板参数，需要将 JSON 压缩并转义为字符串。根据您需要的图像处理功能，参看AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板 ID
	// 和参数信息。
	WorkflowParameter string `json:"WorkflowParameter"`

	// REQUIRED; AI 图像处理模板 ID。根据您需要的图像处理功能，参看 AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板 ID 和参数信息。
	WorkflowTemplateID string `json:"WorkflowTemplateId"`

	// 任务回调配置。
	CallbackConf *CreateImageAITaskBodyCallbackConf `json:"CallbackConf,omitempty"`
}

// CreateImageAITaskBodyCallbackConf - 任务回调配置。
type CreateImageAITaskBodyCallbackConf struct {

	// REQUIRED; 回调 HTTP 请求地址，用于接收转码结果详情。支持使用 https 和 http 协议。
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 回调方式，仅支持取值 HTTP。
	Method string `json:"Method"`

	// 业务自定义回调参数，将在回调消息的 callback_args 中透传。具体回调参数请参考 AI 图像处理回调 [https://www.volcengine.com/docs/508/1526662]。
	Args string `json:"Args,omitempty"`

	// 回调数据格式，仅支持取值 JSON。
	DataFormat string `json:"DataFormat,omitempty"`

	// 回调的维度类型，缺省情况下按照条目级别进行回调。取值如下所示：
	// * task：将按照任务级别进行回调。可分批回调，一个批次内最多一次性可回调 5000 条图片转码条目执行信息。
	// * entry：将按照条目级别进行回调。当该条目执行完毕，将立即产生回调。
	Type string `json:"Type,omitempty"`
}

type CreateImageAITaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageAITaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageAITaskResResult `json:"Result"`
}

type CreateImageAITaskResResponseMetadata struct {

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

type CreateImageAITaskResResult struct {

	// REQUIRED; 队列 ID。查询接口需要使用，请注意保存。
	QueueID string `json:"QueueId"`

	// REQUIRED; 任务 ID。查询接口需要使用，请注意保存。
	TaskID string `json:"TaskId"`
}

type CreateImageAnalyzeTaskBody struct {

	// REQUIRED; 自定义离线评估任务名称
	Name string `json:"Name"`

	// REQUIRED; 任务地区。当前仅支持取值 cn，表示国内。
	Region string `json:"Region"`

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 任务类型。取值如下所示：
	// * OnlineUrl（暂不支持）
	// * SdkUrl（暂不支持）
	// * UrlFile：在线提交 URL 离线评估，即在.txt 文件（评估文件）内填写了待评估图片文件 URL，并将该 txt 文件上传至指定服务后获取并传入该文件的 StoreUri。
	// * UriFile：在线提交 URI 离线评估，即在.txt 文件（评估文件）内填写了待评估图片文件 URI，并将该 txt 文件上传至指定服务后获取并传入该文件的 StoreUri。
	Type string `json:"Type"`

	// 任务描述，可作为该条任务的备注信息。
	Desc string `json:"Desc,omitempty"`

	// txt 评估文件的 Store URI，该文件需上传至指定服务对应存储中。
	// * Type 取值 UrlFile 时，填写合法 URL
	// * Type 取值 UriFile 时，填写指定服务的存储 URI
	ResURI []string `json:"ResUri,omitempty"`

	// 仅当Type 取值 UriFile 时，配置有效。 模板名称，您可通过调用 GetAllImageTemplates [https://www.volcengine.com/docs/508/9386] 获取服务下所有已创建的 TemplateName。
	Tpl string `json:"Tpl,omitempty"`
}

type CreateImageAnalyzeTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageAnalyzeTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageAnalyzeTaskResResult `json:"Result"`
}

type CreateImageAnalyzeTaskResResponseMetadata struct {

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

type CreateImageAnalyzeTaskResResult struct {

	// REQUIRED; 任务 ID
	TaskID string `json:"TaskId"`
}

type CreateImageAuditTaskBody struct {

	// REQUIRED; 审核能力，取值如下所示：
	// * 0：基础审核能力
	// * 1：智能审核能力
	AuditAbility int `json:"AuditAbility"`

	// REQUIRED; "porn" // 涉黄 "govern" // 涉政 "terror" // 涉恐 "illegal" // 违法违规 "sensitive1" // 涉敏1 "sensitive2" // 涉敏2 "forbidden"
	// // 违禁 "uncomfortable" // 引人不适 "qrcode" // 二维码 "badpicture" // 不良画面 "sexy" // 性感低俗 "age"
	// // 年龄 "underage" // 未成年 "quality" // 图片质量
	AuditDimensions []string `json:"AuditDimensions"`

	// REQUIRED; 任务地区。当前仅支持取值 cn，表示国内。
	Region string `json:"Region"`

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 任务类型，当前仅支持取值为 audit。
	TaskType string `json:"TaskType"`

	// REQUIRED; 图片审核任务场景。取值如下所示：
	// * UrlFile：存量图片处理，进针对已有存储内的图片请求获取审核结果。传入方式是 ResUri方式，即在.txt 文件（审核文件）内填写了待审核图片文件 URL，并将该 txt 文件上传至指定服务后获取并传入该文件的 StoreUri。
	// * Url：URL 直传场景。传入方式为 ImageInfos 方式，即可直接传入待审核图片的 URL 及区分标识。
	// * Upload：图片上传场景，针对上传图片到指定服务下的场景。可在 EnableAuditRange下指定审核的范围，例如对指定上传到某目录下的图片进行审核。
	Type string `json:"Type"`

	// 仅当 EnableAuditRange 取值 1 时，配置生效。 指定前缀审核，若你希望对某个目录进行审核，请设置路径为对应的目录名，以/结尾。例如123/test/
	AuditPrefix []string `json:"AuditPrefix,omitempty"`

	// "ad" // 广告 "defraud" // 诈骗 "charillegal" // 文字违规
	AuditTextDimensions []string `json:"AuditTextDimensions,omitempty"`

	// 回调类型，取值需要与 AuditDimensions 审核维度保持一致或少于 AuditDimensions。
	// 例如，AuditDimensions 取值 ["pron","sexy"]，AuditTextDimensions 取值 ["ad"]，支持您将 FreezeDimensions 取值 ["pron","sexy","ad"] 、 ["pron","sexy"]、["pron","ad"]
	// 和 ["sexy","ad"] 任意一种。
	CallbackDimensions []string `json:"CallbackDimensions,omitempty"`

	// 回调图片类型，取值如下所示：
	// * normal：正常图片
	//
	//
	// * problem：问题图片
	//
	//
	// * frozen：冻结图片
	//
	//
	// * fail：审核失败图片
	CallbackImageTypes []string `json:"CallbackImageTypes,omitempty"`

	// 回调 URL，veImageX 以 Post 方式向业务服务器发送 JSON 格式回调数据。具体回调参数请参考回调内容 [https://www.volcengine.com/docs/508/134676#%E5%9B%9E%E8%B0%83%E5%86%85%E5%AE%B9]。
	CallbackURL string `json:"CallbackUrl,omitempty"`

	// 默认0
	EnableAuditRange int `json:"EnableAuditRange,omitempty"`

	// 是否开启回调，取值如下所示：
	// * true：开启
	// * false：（默认）不开启
	EnableCallback bool `json:"EnableCallback,omitempty"`

	// 是否开启冻结，取值如下所示：
	// * true：开启
	// * false：（默认）不开启
	EnableFreeze bool `json:"EnableFreeze,omitempty"`

	// 图片审核仅支持审核 5MB 以下的图片，若您的图片大小在 5MB~32MB，您可以开启大图审核功能，veImageX 会对图片压缩后再进行审核。开启后，将对压缩能力按照基础图片处理
	// [https://www.volcengine.com/docs/508/65935#%E5%9F%BA%E7%A1%80%E5%9B%BE%E5%83%8F%E5%A4%84%E7%90%86%E6%9C%8D%E5%8A%A1]进行计费。（您每月可使用
	// 20TB 的免费额度） 取值如下所示：
	// * true：开启
	// * false：（默认）不开启
	// :::tip
	// * 若未开启大图审核且图片大小 ≥ 5 MB，可能会导致系统超时报错；
	// * 若已开启大图审核但图片大小 ≥ 32 MB，可能会导致系统超时报错。 :::
	EnableLargeImageDetect bool `json:"EnableLargeImageDetect,omitempty"`

	// 若开启冻结，则不可为空
	FreezeDimensions []string `json:"FreezeDimensions,omitempty"`

	// 若开启冻结，则不可为空
	FreezeStrategy int `json:"FreezeStrategy,omitempty"`

	// 若开启冻结，则不可为空
	FreezeType []string `json:"FreezeType,omitempty"`

	// 仅当 Type 为 Url 时，配置生效。
	// 批量提交图片 URL 列表
	ImageInfos []*CreateImageAuditTaskBodyImageInfosItem `json:"ImageInfos,omitempty"`

	// 仅当 EnableAuditRange 取值 1 时，配置生效。 指定前缀不审核，若你希望对某个目录不进行审核，请设置路径为对应的目录名，以/结尾。例如123/test/
	NoAuditPrefix []string `json:"NoAuditPrefix,omitempty"`

	// 仅当 Type 为 UrlFile 时，配置生效。
	// 审核文件的 StoreUri，为 .txt 文件，该文件需上传至指定服务对应存储中。该 txt 文件内需填写待审核图片文件的 URL，每行填写一个，最多可填写 10000 行。
	ResURI []string `json:"ResUri,omitempty"`
}

type CreateImageAuditTaskBodyImageInfosItem struct {

	// 建议您根据实际业务情况，将该参数作为可区分审核图片 ImageUri 的自定义标识。
	DataID string `json:"DataId,omitempty"`

	// 待审核图片 URL，需满足公网可访问。
	ImageURI string `json:"ImageUri,omitempty"`
}

type CreateImageAuditTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageAuditTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageAuditTaskResResult `json:"Result"`
}

type CreateImageAuditTaskResResponseMetadata struct {

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

type CreateImageAuditTaskResResult struct {

	// REQUIRED; 任务 ID
	TaskID string `json:"TaskId"`
}

type CreateImageCompressTaskBody struct {

	// REQUIRED; 文件的处理方式，取值如下所示：
	// * 0：使用 ZIP DEFLATE 压缩方法，将文件进行压缩并打包为 ZIP 包。该方式适用于需要长期存储大量未经压缩的文件的场景，以节省存储空间。但需注意的是，若文件本身已经过压缩处理，将会因为文件的可压缩空间有限，导致该方式的压缩效果不明显。
	// * 1：仅将文件打包为 ZIP 包，但不执行压缩操作。该方式适用于快速整理文件而无需节省存储空间的场景。例如已经过压缩的文件的打包存储，以便于传输或归档。
	ZipMode int `json:"ZipMode"`

	// POST 类型的回调 URL，用于接收相关回调 JSON 类型数据。回调参数请参考回调内容。
	Callback string `json:"Callback,omitempty"`

	// 与IndexFile为二选一，必填。
	// 压缩文件列表配置，最多可支持 3000 个文件，压缩前文件的总大小不得超过 45000 MB。若超出限制，则取消压缩打包操作，致使打包任务失败。
	FileList []*CreateImageCompressTaskBodyFileListItem `json:"FileList,omitempty"`

	// 与FileList为二选一，必填。
	// 索引文件的 StoreUri，为 .txt 文件，该索引文件需上传至指定服务对应存储中。
	// 该索引文件内需填写待压缩文件的地址，每行填写一个，最多可填写 3000 行。压缩前文件的总大小不超过 45000 MB。若超出限制，则取消压缩打包操作，致使打包任务失败。
	IndexFile string `json:"IndexFile,omitempty"`

	// 压缩后文件压缩包的指定名称，默认为随机 Key。
	Output string `json:"Output,omitempty"`
}

type CreateImageCompressTaskBodyFileListItem struct {

	// REQUIRED; 公网可访问的待压缩文件 URL
	URL string `json:"Url"`

	// * 指定时，为 URL 文件所在压缩包中的别名。输入规则如下所示： * 支持汉字、字母、数字及符号-、_和.；
	// * 不能以-、_和.开头；
	// * 长度不得超过 100 个字符。
	//
	//
	// * 不指定时，若能提取原文件名称时，则以 URL 原文件名称；否则，使用随机名称。
	Alias string `json:"Alias,omitempty"`

	// 指定源文件在压缩包中的文件夹，不传时则将该资源存储至一级目录下。输入规则如下所示：
	// * 支持汉字、字母、数字及符号-、_和.；
	// * 不能以-、_和.开头；
	// * 不能以/结尾。
	// :::warning
	// * 建议命名长度不超过 256 个字节。
	// * 建议您在命名中仅使用-、_和.这三种符号。如果您使用了其他符号，可能因操作系统不支持导致处理异常。 :::
	Folder string `json:"Folder,omitempty"`
}

type CreateImageCompressTaskQuery struct {

	// REQUIRED; 压缩文件存储的目标服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateImageCompressTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageCompressTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateImageCompressTaskResResult `json:"Result,omitempty"`
}

type CreateImageCompressTaskResResponseMetadata struct {

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

// CreateImageCompressTaskResResult - 视请求的接口而定
type CreateImageCompressTaskResResult struct {

	// REQUIRED; 压缩任务 ID
	TaskID string `json:"TaskId"`
}

type CreateImageContentTaskBody struct {

	// REQUIRED; 操作类型。取值如下所示：
	// * refresh_url：刷新 URL
	// * refresh_dir：目录刷新（支持根目录刷新）
	// * preload_url：预热 URL
	// * block_url：禁用 URL
	// * unblock_url：解禁 URL
	TaskType string `json:"TaskType"`

	// REQUIRED; 资源 URL 或者目录。
	// * 当TaskType取值refresh_url，一次可提交的最大限额为 2000 个；
	// * 当TaskType取值refresh_dir，一次可提交的最大限额为 50 个；
	// * 当TaskType取值preload_url，一次可提交的最大限额为 2000 个；
	// * 当TaskType取值block_url，一次可提交的最大限额为 2000 个；
	// * 当TaskType取值unblock_url，一次可提交的最大限额为 2000 个。
	Urls []string `json:"Urls"`

	// 仅当 TaskType 为 refresh_dir 使用目录刷新时，可通过此配置开启前缀刷新。取值如下所示： true：开启前缀刷新 false：（默认）关闭前缀刷新，进行标准的目录匹配刷新。
	PrefixRefreshDir bool `json:"PrefixRefreshDir,omitempty"`
}

type CreateImageContentTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageContentTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateImageContentTaskResResult `json:"Result,omitempty"`
}

type CreateImageContentTaskResResponseMetadata struct {

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

// CreateImageContentTaskResResult - 视请求的接口而定
type CreateImageContentTaskResResult struct {

	// REQUIRED; 完成结果提示信息
	Msg string `json:"Msg"`

	// REQUIRED; 对应的任务 ID
	TaskID string `json:"TaskId"`
}

type CreateImageFromURIBody struct {

	// REQUIRED; 待复制资源对应的存储 URI。您可在控制台的资源管理页面，获取上传文件的存储 URI [https://www.volcengine.com/docs/508/1205054]；或者调用 GetImageStorageFiles
	// [https://www.volcengine.com/docs/508/1158308] 获取指定服务下的存储 URI。
	StoreURI string `json:"StoreUri"`

	// 复制后资源的存储 Key。缺省情况下与源存储的资源存储 Key 相同。自定义规则如下所示：
	// * 不支持空格。
	// * 不支持以/开头或结尾，不支持/连续出现，Key 值最大长度限制为 180 个字节。
	DstKey string `json:"DstKey,omitempty"`

	// 待复制资源对应的服务 ID
	OriServiceID string `json:"OriServiceId,omitempty"`

	// 是否保留目标存储中的同名文件。False：不保留目标存储中的同名文件直接覆盖。True：保留目标存储中的同名文件，不覆盖。
	SkipDuplicate bool `json:"SkipDuplicate,omitempty"`
}

type CreateImageFromURIQuery struct {

	// REQUIRED; 复制目标对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateImageFromURIRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageFromURIResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateImageFromURIResResult `json:"Result,omitempty"`
}

type CreateImageFromURIResResponseMetadata struct {

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

// CreateImageFromURIResResult - 视请求的接口而定
type CreateImageFromURIResResult struct {

	// REQUIRED; 资源复制到目标服务后的存储 URI
	DstURI string `json:"DstUri"`

	// REQUIRED; 待复制资源的源存储 URI
	StoreURI string `json:"StoreUri"`
}

type CreateImageHmEmbedBody struct {

	// REQUIRED; 算法模型，取值如下所示：
	// * default：文本嵌入模型，默认文本嵌入模型；
	// * adapt_resize：画质自适应文本嵌入模型。
	Algorithm string `json:"Algorithm"`

	// REQUIRED; 自定义盲水印文本内容。
	// * 文本嵌入模型支持最长可嵌入115个水印内容字符。
	// * 画质自适应文本嵌入模型无水印内容长度限制。
	Info string `json:"Info"`

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待添加盲水印的原图 Uri。
	StoreURI string `json:"StoreUri"`

	// 待添加盲水印的可公网访问原图 Url。当 StoreUri 和 ImageUrl 均不为空，以 StoreUri 取值为准。
	ImageURL string `json:"ImageUrl,omitempty"`

	// 输出图片格式，默认 png，支持图片格式有：png、jpeg、webp。
	OutFormat string `json:"OutFormat,omitempty"`

	// 输出图片质量参数。取值范围为 [1,100]，默认为 75。 对于 PNG 无损压缩，其他格式下其值越小，压缩率越高，画质越差。
	OutQuality int `json:"OutQuality,omitempty"`

	// 算法强度，强度越高，图像抵抗攻击性能越强。
	// 取值如下所示：
	// * low：低强度，适用于纯色图场景以及对图像质量要求高；
	// * medium：中强度，默认中强度；
	// * strong：高强度，适合图像纹理丰富时使用。
	StrengthLevel string `json:"StrengthLevel,omitempty"`
}

type CreateImageHmEmbedRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageHmEmbedResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateImageHmEmbedResResult `json:"Result,omitempty"`
}

type CreateImageHmEmbedResResponseMetadata struct {

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

// CreateImageHmEmbedResResult - 视请求的接口而定
type CreateImageHmEmbedResResult struct {

	// REQUIRED; 添加盲水印后的结果图 Uri。
	StoreURI string `json:"StoreUri"`
}

type CreateImageHmExtractQuery struct {

	// REQUIRED; 算法模型，取值如下所示：
	// * default：文本嵌入基础模型
	// * adapt_resize：画质自适应文本嵌入模型。
	// * adapt: 文本嵌入自适应模型（AIGC 适用）
	// * natural：文本嵌入基础模型（彩色图片通用）
	// * tracev1：前景图层水印模型（纯色背景适用）
	// * tracev2：前景图层水印模型（彩色背景通用）
	// * tracev2-app：前景图层水印模型（移动端）
	// :::warning 指定tracev1、tracev2、tracev2-app模型时，请传入已添加对应模型水印的背景网页的截图。若图片错误，则无法提取水印。 :::
	Algorithm string `json:"Algorithm" query:"Algorithm"`

	// REQUIRED; 待提取盲水印图片的 URL。StoreUri和ImageUrl都不为空时，以StoreUri为准。
	ImageURL string `json:"ImageUrl" query:"ImageUrl"`

	// REQUIRED; 待提取水印图对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 待提取盲水印的图片的 URI。StoreUri和ImageUrl都不为空时，以StoreUri为准。
	StoreURI string `json:"StoreUri" query:"StoreUri"`
}

type CreateImageHmExtractRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageHmExtractResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageHmExtractResResult `json:"Result"`
}

type CreateImageHmExtractResResponseMetadata struct {

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

type CreateImageHmExtractResResult struct {

	// REQUIRED; 提取出的盲水印文本信息。
	Info string `json:"Info"`

	// REQUIRED; 响应码，具体取值如下所示：
	// * -1：盲水印为空；
	// * 0：info不为空时表示盲水印提取成功。 :::tip 提取失败时显示接口错误。 :::
	StatusCode int `json:"StatusCode"`

	// 仅当 Algorithm 取值为 tracev2 时，有返回值。
	// 编码附加信息。
	AdditionInfo *CreateImageHmExtractResResultAdditionInfo `json:"AdditionInfo,omitempty"`
}

// CreateImageHmExtractResResultAdditionInfo - 仅当 Algorithm 取值为 tracev2 时，有返回值。
// 编码附加信息。
type CreateImageHmExtractResResultAdditionInfo struct {

	// 所提取的水印背景图层的生成周期，从 0 开始，表示处于生成的第一周内。
	HmCode int `json:"HmCode,omitempty"`

	// 生成周期所对应的起始与结束时间段，固定为 7 天。
	HmDateInfo *CreateImageHmExtractResResultAdditionInfoHmDateInfo `json:"HmDateInfo,omitempty"`

	// 水印类型
	HmType string `json:"HmType,omitempty"`
}

// CreateImageHmExtractResResultAdditionInfoHmDateInfo - 生成周期所对应的起始与结束时间段，固定为 7 天。
type CreateImageHmExtractResResultAdditionInfoHmDateInfo struct {

	// 结束时间戳
	EndDate int `json:"EndDate,omitempty"`

	// 开始时间戳
	StartDate int `json:"StartDate,omitempty"`
}

type CreateImageMigrateTaskBody struct {

	// REQUIRED; 任务信息
	Task *CreateImageMigrateTaskBodyTask `json:"Task"`
}

// CreateImageMigrateTaskBodyTask - 任务信息
type CreateImageMigrateTaskBodyTask struct {

	// REQUIRED; 目的信息
	Dst *CreateImageMigrateTaskBodyTaskDst `json:"Dst"`

	// REQUIRED; 自定义迁移任务名称
	Name string `json:"Name"`

	// REQUIRED; 迁移源信息
	Source *CreateImageMigrateTaskBodyTaskSource `json:"Source"`

	// 回调信息。配置后，当任务执行完成时，将往该回调配置地址发送任务回调信息。
	CallbackCfg *CreateImageMigrateTaskBodyTaskCallbackCfg `json:"CallbackCfg,omitempty"`

	// 迁移策略
	RunStrategy *CreateImageMigrateTaskBodyTaskRunStrategy `json:"RunStrategy,omitempty"`

	// 转码配置
	Transcode *CreateImageMigrateTaskBodyTaskTranscode `json:"Transcode,omitempty"`
}

// CreateImageMigrateTaskBodyTaskCallbackCfg - 回调信息。配置后，当任务执行完成时，将往该回调配置地址发送任务回调信息。
type CreateImageMigrateTaskBodyTaskCallbackCfg struct {

	// REQUIRED; 回调地址。Method取值http时，填写公网可访问的 URL 地址，任务结束将向该地址发送 HTTP POST 请求。具体回调参数请参考回调内容。
	Addr string `json:"Addr"`

	// REQUIRED; 回调方法。仅支持取值为 http。
	Method string `json:"Method"`

	// 任务维度自定义回调参数，最大长度为1024
	CallbackArgs string `json:"CallbackArgs,omitempty"`

	// 回调信息中是否包含具体迁移任务条目信息。取值如下所示：
	// * true：包含。仅包含迁移成功的任务条目信息，迁移失败的任务列表请在迁移完成后调用 ExportFailedMigrateTask [https://www.volcengine.com/docs/508/1261309] 接口获取。
	// * false：（默认）不包含。 :::warning 若任务中包含的条目数量过多，会导致回调消息体过大，增加回调失败的风险。因此建议仅在任务中条目量级不超过十万时使用该参数。 :::
	IncludeEntry bool `json:"IncludeEntry,omitempty"`
}

// CreateImageMigrateTaskBodyTaskDst - 目的信息
type CreateImageMigrateTaskBodyTaskDst struct {

	// REQUIRED; 迁移目标服务 ID，请提前新建服务 [https://www.volcengine.com/docs/508/357114#%E6%96%B0%E5%BB%BA%E6%9C%8D%E5%8A%A1]。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 目标 key 前缀，即保存到到指定目录下。如需多重目录，请使用/分割，并以/结尾。 默认为空，表示迁移到根目录。
	// * 使用非 URL 方式迁移到根目录时：迁移后 存储 Key 与源存储 Bucket 的文件存储 Key 相同。
	// * 使用 Url 方式迁移到根目录时：迁移后存储 Key 与源 URL 中 Path 值相同。
	Prefix string `json:"Prefix,omitempty"`

	// 源 Bucket 名称保留规则。取值如下所示：
	// * true：不保留，迁移后资源访问 URI 中，不保留迁移源的 Bucket 名称。
	// * false：（默认）保留，迁移后资源访问 URI 中，会保留迁移源的 Bucket 名称。
	SkipBucket bool `json:"SkipBucket,omitempty"`

	// 同名文件覆盖规则配置。取值如下所示：
	// * 0：（默认）直接覆盖同名文件
	// * 1：增加文件名后缀，后缀为任务 ID
	// * 2：跳过同名文件，即不做迁移
	// :::tip 同名文件指文件在对象存储中的访问 Key 相同的文件，调用 veImageX 服务时会用到文件访问 Key。 :::
	UploadConf int `json:"UploadConf,omitempty"`
}

// CreateImageMigrateTaskBodyTaskRunStrategy - 迁移策略
type CreateImageMigrateTaskBodyTaskRunStrategy struct {

	// 源下载 QPS 限制。如取值不为空，则长度必须为 24，表示一天 24 小时内各小时的 QPS 限制值。默认无限制。
	// * 取值为负值时，表示无限制
	// * 取值为 0 时，表示对应时间不允许迁移
	ReadQPS []int `json:"ReadQps,omitempty"`

	// 源下载流量限制。单位为 Byte。如取值不为空，则长度必须为24，表示一天 24 小时内各小时的流量限制值。默认无限制。
	// * 取值为负值时，表示无限制
	// * 取值为 0 时，表示对应时间不允许迁移
	ReadRate []int `json:"ReadRate,omitempty"`
}

// CreateImageMigrateTaskBodyTaskSource - 迁移源信息
type CreateImageMigrateTaskBodyTaskSource struct {

	// REQUIRED; * 仅当Vendor为URL时，需填写 URL 列表文件地址（公网 URL 地址）。 :::tip 支持指定迁移文件和转码后迁移文件进行重命名，详见URL 列表迁移文件说明 [https://www.volcengine.com/docs/508/1263268]。
	// :::
	// * 当Vendor为其他时，需填写对应云服务商所需迁移数据的 Bucket 名称。您可参考云数据迁移准备 [https://www.volcengine.com/docs/508/129213]获取对应阿里云OSS、腾讯云COS、七牛云KODO、百度云BOS、华为云OBS、
	// 优刻得（Ucloud File)、AWS国际站的 Bucket 名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 迁移云服务商。取值如下所示：
	// * OSS：阿里云
	// * COS：腾讯云
	// * KODO：七牛云
	// * BOS：百度云
	// * OBS：华为云
	// * Ucloud：Ucloud file
	// * AWS：AWS 国际站
	// * S3：其他 S3 协议存储
	// * URL：以上传 URL 列表的方式迁移
	Vendor string `json:"Vendor"`

	// 仅当Vendor非URL时，为必填。
	// Access Key，与 Secret Key 同时填写，为了保证有访问源数据桶的权限。
	// * 请参考云数据迁移准备 [https://www.volcengine.com/docs/508/129213]获取对应阿里云OSS、腾讯云COS、七牛云KODO、百度云BOS、华为云OBS、 优刻得（Ucloud File)、AWS国际站的账号
	// AK/SK。
	// * 对于其他 S3 协议存储的AK/SK，请根据其具体源站信息填写。
	AK string `json:"AK,omitempty"`

	// Vendor != URL 时选填，表示桶清单文件在Bucket中的存储目录路径。若不为空，则使用桶清单获取桶内待迁移文件。否则，会调用桶的 ListObject 接口获取桶内待迁移文件。
	BucketInventoryDir string `json:"BucketInventoryDir,omitempty"`

	// BucketInventoryDir != "" 时必填，表示桶清单CSV文件的表头信息，数组元素顺序需与桶清单CSV文件的列顺序保持一致。ImageX会根据该参数解析桶清单CSV文件，参数规范为：
	// * 必须包含 Key ：表示待迁移的资源Key
	// * 强烈推荐包含 Size ：表示待迁移的资源大小
	// * 推荐包含 ETag ：表示待迁移资源的ETag值
	BucketInventorySchema []string `json:"BucketInventorySchema,omitempty"`

	// 仅当Vendor非URL时，为选填。
	// 迁移源云服务商 CDN 域名，若不为空将使用该 CDN 域名下载三方云厂商的资源。
	CdnHost string `json:"CdnHost,omitempty"`

	// 仅当Vendor为S3时，为必填。
	// S3 协议 Endpoint，需以http://或https://开头。请根据源站信息填写。
	Endpoint string `json:"Endpoint,omitempty"`

	// 仅迁移匹配的前缀列表文件。文件路径前缀无需包含桶名称，但需要完整路径。 默认为空，表示对该存储 Bucket 内资源执行全量迁移。若不为空，表示仅做部分迁移，即指定迁移的文件路径前缀。
	Prefix []string `json:"Prefix,omitempty"`

	// 仅迁移匹配的正则表达式列表的文件。默认为空，表示对该存储 Bucket 内资源执行全量迁移。
	// :::tip
	// * 多条正则表达式之间是"或"的关系，即源文件匹配任何一条正则表达式即视为符合迁移条件。
	// * 正则过滤规则需要遍历源桶中的全部文件，如果源桶中文件数量较多会降低迁移速度。 :::
	Regex []string `json:"Regex,omitempty"`

	// 仅当Vendor非URL/OSS/KODO/AWS时，为必填。
	// Bucket 所在地区。
	// * 请参考云数据迁移准备 [https://www.volcengine.com/docs/508/129213]获取对应阿里云OSS、腾讯云COS、七牛云KODO、百度云BOS、华为云OBS、 优刻得（Ucloud File)、AWS国际站的
	// Bucket 地区。
	// * 对于其他 S3 协议存储的 Bucket 地区，请根据其具体源站信息填写。
	Region string `json:"Region,omitempty"`

	// 仅当Vendor非URL时，为必填。
	// Secret Key，与 Access Key 同时填写，为了保证有访问源数据桶的权限。
	// * 请参考云数据迁移准备 [https://www.volcengine.com/docs/508/129213]获取对应阿里云OSS、腾讯云COS、七牛云KODO、百度云BOS、华为云OBS、 优刻得（Ucloud File)、AWS国际站的账号
	// AK/SK。
	// * 对于其他 S3 协议存储的AK/SK，请根据其具体源站信息填写。
	SK string `json:"SK,omitempty"`

	// 是否丢弃源 Header。取值如下所示：
	// * true：丢弃源 Header
	// * false：（默认）保留源 Header
	SkipHeader bool `json:"SkipHeader,omitempty"`

	// 迁移文件结束时间点。默认为空。仅迁移该查询时间段内新增或变更的文件。
	// 日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeEnd string `json:"TimeEnd,omitempty"`

	// 迁移文件起始时间点。仅迁移该查询时间段内新增或变更的文件。默认为空。
	// 日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStart string `json:"TimeStart,omitempty"`
}

// CreateImageMigrateTaskBodyTaskTranscode - 转码配置
type CreateImageMigrateTaskBodyTaskTranscode struct {

	// REQUIRED; 目标转码格式，仅针对静图执行转码策略。支持的格式有 png、jpeg、heic、avif、webp、vvic。
	Format string `json:"Format"`

	// REQUIRED; 转码质量参数，取值范围为 [1,100]。对于 PNG 为无损压缩，其他格式下其值越小，压缩率越高，画质越差。
	Quality int `json:"Quality"`

	// 是否开启自适应编码。仅编码/降级格式为heic/webp/jpeg时生效
	Adapt bool `json:"Adapt,omitempty"`

	// 包含透明通道的图片是否编码为降级格式。取值如下所示：
	// * true：降级
	// * false：（默认）不降级
	AlphaDemotion bool `json:"AlphaDemotion,omitempty"`

	// 降级编码格式，仅当AlphaDemotion为true时必填。支持的格式有 png、jpeg、heic、avif、webp、vvic。
	DemotionFmt string `json:"DemotionFmt,omitempty"`

	// 转码是否保留 exif 信息。取值如下所示：
	// * true：保留
	// * false：（默认）不保留
	EnableExif bool `json:"EnableExif,omitempty"`

	// 当 jpeg 原图在迁移中指定转码为 heic 图时， heic 图是否需要存储原图大小的数据。
	// * true：是
	// * false：（默认）否
	ReserveJPEGSize bool `json:"ReserveJpegSize,omitempty"`

	// 对带有 CMYK 色彩空间的图片，是否跳过转码处理直接存储原图。取值如下所示：
	// * true：是
	// * false：（默认）否
	SkipCMYK bool `json:"SkipCMYK,omitempty"`
}

type CreateImageMigrateTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageMigrateTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageMigrateTaskResResult `json:"Result"`
}

type CreateImageMigrateTaskResResponseMetadata struct {

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

type CreateImageMigrateTaskResResult struct {

	// REQUIRED; 创建成功的迁移任务 ID
	TaskID string `json:"TaskId"`
}

type CreateImageMonitorRuleBody struct {

	// REQUIRED; 告警规则
	MonitorRule *CreateImageMonitorRuleBodyMonitorRule `json:"MonitorRule"`
}

// CreateImageMonitorRuleBodyMonitorRule - 告警规则
type CreateImageMonitorRuleBodyMonitorRule struct {

	// REQUIRED; 监控的应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppID。
	Appid string `json:"Appid"`

	// REQUIRED; 监测规则。
	Cond *CreateImageMonitorRuleBodyMonitorRuleCond `json:"Cond"`

	// REQUIRED; 创建后是否立即开启告警，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"Enabled"`

	// REQUIRED; 监控频率，单位为分钟。取值如下所示：
	// * 5
	// * 10
	// * 20
	// * 30
	// * 40
	// * 50
	Frequency int `json:"Frequency"`

	// REQUIRED; 告警级别，取值如下所示：
	// * warn：警告
	// * error：错误
	// * fatal：致命
	Level string `json:"Level"`

	// REQUIRED; 自定义告警规则名称
	Name string `json:"Name"`

	// REQUIRED; 告警通知配置。
	Notification *CreateImageMonitorRuleBodyMonitorRuleNotification `json:"Notification"`

	// REQUIRED; 监控阶段，取值如下所示：
	// * upload：图片上传-上传 1.0
	// * uploadv2：图片上传-上传 2.0
	// * cdn：图片加载-下行网络监控
	// * client：图片加载-客户端传状态监控
	// * sensible：图片加载-感知指标监控
	Phase string `json:"Phase"`

	// 维度过滤条件，具体参数请见Filter。用于指定需要告警提示的维度配置。
	Filter *CreateImageMonitorRuleBodyMonitorRuleFilter `json:"Filter,omitempty"`

	// 拆分维度，由公共拆分维度 [https://www.volcengine.com/docs/508/1113944]和自定义拆分维度 [https://www.volcengine.com/docs/508/34554]组合而成。
	GroupBy string `json:"GroupBy,omitempty"`

	// 监控平台，取值如下所示：
	// * iOS
	// * Android
	// * WEB
	OS string `json:"OS,omitempty"`
}

// CreateImageMonitorRuleBodyMonitorRuleCond - 监测规则。
type CreateImageMonitorRuleBodyMonitorRuleCond struct {

	// REQUIRED
	ItemCond []*CreateImageMonitorRuleBodyMonitorRuleCondItem `json:"ItemCond"`

	// REQUIRED; 多条监控规则之间的逻辑关系，取值如下所示：
	// * and：且。表示有多条监控规则时，需满足所有监控规则才会触发告警通知。
	// * or：或。表示有多条监控规则时，满足其中一条监控规则就会触发告警通知。
	LogicOp string `json:"LogicOp"`
}

// CreateImageMonitorRuleBodyMonitorRuleCondItem - 监控规则配置
type CreateImageMonitorRuleBodyMonitorRuleCondItem struct {

	// REQUIRED; 聚合周期，单位为分钟。被监控指标在该指定周期内满足指标比较阈值且上报量满足样本量阈值时，才会触发告警。取值如下所示：
	// * 5
	// * 10
	AggrInterval int `json:"AggrInterval"`

	// REQUIRED; 指标取值函数，取值如下所示：
	// * max：最大值
	// * min：最小值
	// * avg：平均值
	// * pct25：25峰值
	// * pct50：50峰值
	// * pct90：90峰值
	// * pct99：99峰值
	// * sum：总和
	// :::tip 各指标支持的函数参考veImageX 告警指标定义 [https://www.volcengine.com/docs/508/1113944]。 :::
	Func string `json:"Func"`

	// REQUIRED; 指标名称，取值参考veImageX 告警指标定义 [https://www.volcengine.com/docs/508/1113944]。
	Item string `json:"Item"`

	// REQUIRED; 指标比较方法，取值如下所示：
	// * LE：小于等于
	// * GE：大于等于
	// * INC：环比上升大于等于
	// * INC_LE：环比上升小于等于
	// * DEC：环比下降小于等于
	// * DEC_GE：环比下降大于等于
	// * HOH_INC：与上小时同比上升大于等于
	// * HOH_INC_LE：与上小时同比上升小于等于
	// * HOH_DEC：与上小时同比下降小于等于
	// * HOH_DEC_GE：与上小时同比下降大于等于
	// * DOD_INC：与昨天同比上升大于等于
	// * DOD_INC_LE：与昨天同比上升小于等于
	// * DOD_DEC：与昨天同比下降小于等于
	// * DOD_DEC_GE：与昨天同比下降大于等于
	Op string `json:"Op"`

	// REQUIRED; 持续周期，当监控指标在聚合周期内，连续RepeatCnt次满足指标比较阈值且上报量满足样本量阈值时，才会触发告警。取值如下所示：
	// * 1
	// * 3
	// * 5
	RepeatCnt int `json:"RepeatCnt"`

	// REQUIRED; 指标比较阈值，需要与CntThreshold同时被满足才会触发告警。
	Threshold float64 `json:"Threshold"`

	// 样本量阈值。被监控指标超过该值时触发告警。
	CntThreshold int `json:"CntThreshold,omitempty"`
}

// CreateImageMonitorRuleBodyMonitorRuleFilter - 维度过滤条件，具体参数请见Filter。用于指定需要告警提示的维度配置。
type CreateImageMonitorRuleBodyMonitorRuleFilter struct {

	// REQUIRED; 过滤条件
	DimFilter []*CreateImageMonitorRuleBodyMonitorRuleFilterDimFilterItem `json:"DimFilter"`

	// REQUIRED; 过滤条件之间的逻辑关系，取值如下所示：
	// * and：和
	// * or：或
	LogicOp string `json:"LogicOp"`
}

type CreateImageMonitorRuleBodyMonitorRuleFilterDimFilterItem struct {

	// REQUIRED; 维度名称，由公共过滤维度 [https://www.volcengine.com/docs/508/1113944]和自定义过滤维度 [https://www.volcengine.com/docs/508/34554]组合而成。
	Dim string `json:"Dim"`

	// REQUIRED; 维度取值，您可以通过调用获取自定义维度值 [https://www.volcengine.com/docs/508/34555]来获取。
	Vals []string `json:"Vals"`

	// 纬度值是否取反，取值如下所示：
	// * true：指定维度的实际值不得满足Vals所有指定值
	// * false：（默认）维度值等于Vals中之一即可
	Not bool `json:"Not,omitempty"`
}

// CreateImageMonitorRuleBodyMonitorRuleNotification - 告警通知配置。
type CreateImageMonitorRuleBodyMonitorRuleNotification struct {

	// REQUIRED; 通知内容模板，模板中变量格式为$Name$。Name 取值如下所示：
	// * 报警名称
	// * 报警级别
	// * 报警App
	// * 报警平台
	// * 报警时间
	// * 报警内容
	Content string `json:"Content"`

	// REQUIRED; 通知方式，仅支持取值http_callback，表示回调。
	Mode []string `json:"Mode"`

	// REQUIRED; 沉默周期，单位为分钟。告警发生后，若未恢复正常，则会间隔一个沉默周期后再次重复发送一次告警通知。取值如下所示：
	// * 0
	// * 30
	// * 60
	// * 360
	SilentDur int `json:"SilentDur"`

	// REQUIRED; 告警通知标题
	Title string `json:"Title"`

	// 回调地址，Mode包含http_callback时，为必填。
	CallbackURL string `json:"CallbackUrl,omitempty"`
}

type CreateImageMonitorRuleRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageMonitorRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateImageMonitorRuleResResult `json:"Result,omitempty"`
}

type CreateImageMonitorRuleResResponseMetadata struct {

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

// CreateImageMonitorRuleResResult - 视请求的接口而定
type CreateImageMonitorRuleResResult struct {

	// REQUIRED; 告警规则 ID
	RuleID string `json:"RuleId"`
}

type CreateImageRetryAuditTaskBody struct {

	// REQUIRED; 失败图片 ID，您可通过调用获取审核任务结果 [https://www.volcengine.com/docs/508/1160410]查看待更新状态的图片条目 ID。
	Entry string `json:"Entry"`

	// REQUIRED; 任务 ID，您可通过调用查询所有审核任务 [https://www.volcengine.com/docs/508/1160409]获取所需的任务 ID。
	TaskID string `json:"TaskId"`
}

type CreateImageRetryAuditTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageRetryAuditTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageRetryAuditTaskResResult `json:"Result"`
}

type CreateImageRetryAuditTaskResResponseMetadata struct {

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

type CreateImageRetryAuditTaskResResult struct {

	// REQUIRED
	Result string `json:"Result"`
}

type CreateImageServiceBody struct {

	// REQUIRED; 服务名称，最多不超过32个字符
	ServiceName string `json:"ServiceName"`

	// REQUIRED; 服务所在区域，cn或va或sg
	ServiceRegion string `json:"ServiceRegion"`

	// 创建服务时绑定的域名列表
	Domains []*CreateImageServiceBodyDomainsItem `json:"Domains,omitempty"`

	// 服务绑定的项目。仅对ToB账号请求生效，默认default
	ProjectName string `json:"ProjectName,omitempty"`

	// 服务绑定的标签。仅对ToB账号请求生效，默认空，不绑定标签。
	ResourceTags []*CreateImageServiceBodyResourceTagsItem `json:"ResourceTags,omitempty"`

	// 取值为StaticRc时表示创建「静态资源托管服务」，取值为Image时表示创建「图片处理服务」。默认创建「图片处理服务」
	ServiceType string `json:"ServiceType,omitempty"`
}

type CreateImageServiceBodyDomainsItem struct {

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// 证书ID
	CertID string `json:"CertID,omitempty"`
}

type CreateImageServiceBodyResourceTagsItem struct {

	// REQUIRED; 标签键
	Key string `json:"Key"`

	// REQUIRED; 标签值
	Value string `json:"Value"`
}

type CreateImageServiceRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageServiceResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageServiceResResult `json:"Result"`
}

type CreateImageServiceResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error *CreateImageServiceResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type CreateImageServiceResResponseMetadataError struct {

	// REQUIRED; 错误码
	Code string `json:"Code"`

	// REQUIRED; 错误信息
	Message string `json:"Message"`
}

type CreateImageServiceResResult struct {

	// REQUIRED; 创建的服务 ID。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 服务的名称。
	ServiceName string `json:"ServiceName"`
}

type CreateImageSettingRuleBody struct {

	// REQUIRED; 应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppId。
	AppID string `json:"AppId"`

	// REQUIRED; 规则内容
	Rule *CreateImageSettingRuleBodyRule `json:"Rule"`

	// REQUIRED; 配置项 ID，您可以通过调用获取配置项列表 [https://www.volcengine.com/docs/508/1324617]的方式获取所需的配置项 ID。
	SettingID string `json:"SettingId"`
}

// CreateImageSettingRuleBodyRule - 规则内容
type CreateImageSettingRuleBodyRule struct {

	// REQUIRED; 规则名称，仅支持字母、数字、下划线，最多输入 32 个字符。
	Name string `json:"Name"`

	// REQUIRED; 类型由对应配置项决定，此处是为了方便生成 SDK
	Value interface{} `json:"Value"`

	// 匹配条件，仅当条件匹配后规则才会生效。
	Cond *CreateImageSettingRuleBodyRuleCond `json:"Cond,omitempty"`
}

// CreateImageSettingRuleBodyRuleCond - 匹配条件，仅当条件匹配后规则才会生效。
type CreateImageSettingRuleBodyRuleCond struct {

	// 规则条件
	Conds []*CreateImageSettingRuleBodyRuleCondCondsItem `json:"Conds,omitempty"`

	// 匹配条件，取值如下所示：
	// * AND：表示与
	// * OR：表示或
	Type string `json:"Type,omitempty"`
}

type CreateImageSettingRuleBodyRuleCondCondsItem struct {

	// 过滤维度，取值请参考规则配置条件 [https://www.volcengine.com/docs/508/65940#%E8%A7%84%E5%88%99%E9%85%8D%E7%BD%AE%E6%9D%A1%E4%BB%B6]。
	Key string `json:"Key,omitempty"`

	// 操作符。支持取值：==、!=、>、>=、<、<=、in
	Op string `json:"Op,omitempty"`

	// 类型由Op决定，此处类型是为了方便生成 SDK
	Value interface{} `json:"Value,omitempty"`
}

type CreateImageSettingRuleRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageSettingRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateImageSettingRuleResResult `json:"Result,omitempty"`
}

type CreateImageSettingRuleResResponseMetadata struct {

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

// CreateImageSettingRuleResResult - 视请求的接口而定
type CreateImageSettingRuleResResult struct {

	// REQUIRED; 规则 ID。
	RuleID string `json:"RuleId"`
}

type CreateImageStyleBody struct {

	// REQUIRED; 样式画布的高度，取值范围为[0,1000]。
	Height int `json:"Height"`

	// REQUIRED; 样式名称，当前对字符长度及支持字符暂无限制。
	Name string `json:"Name"`

	// REQUIRED; 绑定的服务 ID，用于计量计费和样式渲染结果图的存储。
	// * 您可以在veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 样式画布的宽度，取值范围为[0,1000]。
	Width int `json:"Width"`

	// 当前仅支持取值px表示像素
	Unit CreateImageStyleBodyUnit `json:"Unit,omitempty"`
}

type CreateImageStyleRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageStyleResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *CreateImageStyleResResult `json:"Result,omitempty"`
}

type CreateImageStyleResResponseMetadata struct {

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

// CreateImageStyleResResult - title
type CreateImageStyleResResult struct {

	// REQUIRED; 创建好的样式 ID。
	StyleID string `json:"StyleId"`
}

type CreateImageTemplateBody struct {

	// REQUIRED; 模板名称，必须使用该服务的图片模板固定前缀，具体参考GetImageService接口的返回参数TemplatePrefix。模板名称能包含的字符正则集合为[a-zA-Z0-9_-]
	TemplateName string `json:"TemplateName"`

	// 图像格式自适应配置
	AdaptiveFmt *CreateImageTemplateBodyAdaptiveFmt `json:"AdaptiveFmt,omitempty"`

	// 仅当指定输出格式为静图时，配置有效。 动图截帧配置。
	AnimExtract *CreateImageTemplateBodyAnimExtract `json:"AnimExtract,omitempty"`

	// 视频转动图配置
	Animation *CreateImageTemplateBodyAnimation `json:"Animation,omitempty"`

	// 模板计划使用的降级格式，仅对heic静图有效
	DemotionFormat string `json:"DemotionFormat,omitempty"`

	// 是否直接更新模板。如果为true，已有的线上模板会同步更新，直接生效；如果为false，会新增一个模板，已有模板不受影响。
	DoUpdate bool `json:"DoUpdate,omitempty"`

	// 对结果图片执行的画质评估配置
	Evals []*CreateImageTemplateBodyEvalsItem `json:"Evals,omitempty"`

	// 仅当指定输出格式非动图时，配置有效。 保留 EXIF 信息配置。
	Exif *CreateImageTemplateBodyExif `json:"Exif,omitempty"`

	// 对图片的编辑操作
	Filters []*CreateImageTemplateBodyFiltersItem `json:"Filters,omitempty"`

	// 对图片编码使用的质量参数，默认为0
	OuputQuality int `json:"OuputQuality,omitempty"`

	// 用于图片服务输出时的图片编码
	OutputExtra *CreateImageTemplateBodyOutputExtra `json:"OutputExtra,omitempty"`

	// 该模板计划使用的输出格式
	OutputFormat string `json:"OutputFormat,omitempty"`

	// 图片模板使用的参数列表，URL中下发参数的顺序需要跟列表中的保持一致
	Parameters []string `json:"Parameters,omitempty"`

	// 处理结果持久化
	Persistence string `json:"Persistence,omitempty"`

	// 绝对质量/相对质量
	QualityMode string `json:"QualityMode,omitempty"`

	// URL的失效期，为Unix时间戳，一般配置为通过模板参数下发
	ReqDeadline string `json:"ReqDeadline,omitempty"`

	// 视频截帧配置
	Snapshot *CreateImageTemplateBodySnapshot `json:"Snapshot,omitempty"`

	// 是否同步处理，仅对heic图有效
	Sync bool `json:"Sync,omitempty"`

	// 是否为临时使用，取值如下所示：
	// * true：是
	// * false：否
	Temporary bool `json:"Temporary,omitempty"`

	// 是否开启鉴权，一般当通过模板参数下发敏感信息时，比如文字水印内容、URL失效期，需要对图片URL鉴权保护，防止内容被篡改
	WithSig bool `json:"WithSig,omitempty"`
}

// CreateImageTemplateBodyAdaptiveFmt - 图像格式自适应配置
type CreateImageTemplateBodyAdaptiveFmt struct {

	// 动图自适应，可选"webp"/"heic"/"avif"/"dynamic"
	Animated string `json:"Animated,omitempty"`

	// 静图自适应，可选"webp"/"heic"/"avif"/"dynamic"
	Static string `json:"Static,omitempty"`
}

// CreateImageTemplateBodyAnimExtract - 仅当指定输出格式为静图时，配置有效。 动图截帧配置。
type CreateImageTemplateBodyAnimExtract struct {

	// 动图截帧策略，取值如下所示：
	// * 0：智能模式，从动图首帧开始逐帧检测当前帧亮度是否大于 80，并最终返回第一个亮度大于 80 的帧。
	// * 1：全局最优，从动图首帧开始逐帧检测并返回亮度最大的一帧。
	Strategy int `json:"Strategy,omitempty"`

	// 动图异步处理超时时间，单位为 ms。默认为 1500，取值负数时表示无超时时间。若在指定时间范围内处理未完成则返回失败。
	Timeout int `json:"Timeout,omitempty"`
}

// CreateImageTemplateBodyAnimation - 视频转动图配置
type CreateImageTemplateBodyAnimation struct {

	// REQUIRED; 降级类型： image：抽取一帧降级返回 video：直接返回原视频降级
	DemotionType string `json:"DemotionType"`

	// REQUIRED; 动图时长(ms)
	Duration int `json:"Duration"`

	// REQUIRED; 帧率，1秒X帧
	FramePerSecond int `json:"FramePerSecond"`

	// REQUIRED; X秒1帧
	SecondPerFrame int `json:"SecondPerFrame"`

	// REQUIRED; 抽帧策略： fps：frame per second，1秒X帧 spf：second per frame，X秒1帧 key：抽取关键帧
	SelectFrameMode string `json:"SelectFrameMode"`

	// REQUIRED; 动图起始时间戳(ms)
	StartTime int `json:"StartTime"`

	// REQUIRED; 同步等待时长(s)，超时未完成则根据DemotionType降级
	WaitTime int `json:"WaitTime"`
}

type CreateImageTemplateBodyEvalsItem struct {

	// REQUIRED; 评估名，画质评估固定取值为 quality。
	Name string `json:"Name"`

	// REQUIRED; 画质评估参数配置内容，Key 为 参数名称，Value 为 参数配置。具体详情请见图片编辑数据结构 [https://www.volcengine.com/docs/508/127820]。
	Param map[string]interface{} `json:"Param"`
}

// CreateImageTemplateBodyExif - 仅当指定输出格式非动图时，配置有效。 保留 EXIF 信息配置。
type CreateImageTemplateBodyExif struct {

	// 是否开启保留全部 EXIF 信息。取值如下所示：
	// * true：是
	// * false：否
	AutoOrientOff bool `json:"AutoOrientOff,omitempty"`

	// 是否开启保留全部 EXIF 信息，取值如下所示：
	// * true：是
	// * false：否
	ExifReserve bool `json:"ExifReserve,omitempty"`

	// 保留部分 EXIF 信息的具体内容，多个之间用,分隔。更多信息请参考标准 EXIF 标签 [https://exiv2.org/tags.html]。
	ExifRetainNames []string `json:"ExifRetainNames,omitempty"`
}

type CreateImageTemplateBodyFiltersItem struct {

	// REQUIRED; 编辑操作的名称
	Name string `json:"Name"`

	// REQUIRED; 编辑操作的参数
	Param map[string]interface{} `json:"Param"`
}

// CreateImageTemplateBodyOutputExtra - 用于图片服务输出时的图片编码
type CreateImageTemplateBodyOutputExtra struct {
	AvifDemfmt string `json:"avif.demfmt,omitempty"`

	// 仅当OutputFormat取值为heic时配置有效 是否带透明通道编码，取值如下所示：
	// * true：是
	// * false：否
	HeicAlphaReserve string `json:"heic.alpha.reserve,omitempty"`
	HeicAqMode       string `json:"heic.aq.mode,omitempty"`
	HeicDemfmt       string `json:"heic.demfmt,omitempty"`

	// 仅当OutputFormat取值为heic时配置有效 色位深度，值越大则提供的图像色彩范围越多，使图像颜色变化的更细腻，但图像体积也会增大。取值如下所示：
	// * 8：8bit
	// * 10：10bit
	HeicEncodeDepth          string `json:"heic.encode.depth,omitempty"`
	HeicJPEGSizeReserve      string `json:"heic.jpeg.size.reserve,omitempty"`
	HeicQualityAdaptPixlimit string `json:"heic.quality.adapt.pixlimit,omitempty"`
	HeicQualityAdaptVersion  string `json:"heic.quality.adapt.version,omitempty"`

	// 仅当OutputFormat取值为heic时配置有效 是否开启 ROI 编码，取值如下所示：
	// * true：是
	// * false：否
	HeicRoi string `json:"heic.roi,omitempty"`

	// 仅当OutputFormat取值为heic时配置有效 缩略图比例。在原图基础上，编码缩小一定倍数的小分辨率图片，跟大图一起封装在同一张图片中，缩小倍数不建议过大，一般为 5~10 之间相对合理。
	HeicThumbRatio string `json:"heic.thumb.ratio,omitempty"`

	// jpeg 的 alpha 图片是否降级为 png，指定为 png 时表示降级为 png 格式。缺省情况下默认为空，表示不降级。
	JPEGAlphaDemotion string `json:"jpeg.alpha.demotion,omitempty"`

	// 是否采用 jpeg 渐进编码格式，取值如下所示：
	// * true：是
	// * false：否
	JPEGProgressive          string `json:"jpeg.progressive,omitempty"`
	JPEGQualityAdaptPixlimit string `json:"jpeg.quality.adapt.pixlimit,omitempty"`
	JPEGQualityAdaptVersion  string `json:"jpeg.quality.adapt.version,omitempty"`

	// 指定 jpeg 体积的输出大小，需同时设置 jpeg.size.fixed.padding，二者缺一不可。 指定输出体积大小，单位为 Byte。
	JPEGSizeFixed string `json:"jpeg.size.fixed,omitempty"`

	// 指定 jpeg 体积的输出大小，需同时指定 jpeg.size.fixed，二者缺一不可。 体积填充方式，取值固定为 append。
	JPEGSizeFixedPadding string `json:"jpeg.size.fixed.padding,omitempty"`
	JPEGSizeRecover      string `json:"jpeg.size.recover,omitempty"`

	// 是否压缩颜色空间，取值如下所示：
	// * true：是
	// * false：否
	PNGUseQuant              string `json:"png.use_quant,omitempty"`
	VvicAqMode               string `json:"vvic.aq.mode,omitempty"`
	VvicQualityAdaptPixlimit string `json:"vvic.quality.adapt.pixlimit,omitempty"`
	VvicQualityAdaptVersion  string `json:"vvic.quality.adapt.version,omitempty"`
	VvicRoi                  string `json:"vvic.roi,omitempty"`
	WebpQualityAdaptPixlimit string `json:"webp.quality.adapt.pixlimit,omitempty"`
	WebpQualityAdaptVersion  string `json:"webp.quality.adapt.version,omitempty"`
}

// CreateImageTemplateBodySnapshot - 视频截帧配置
type CreateImageTemplateBodySnapshot struct {

	// REQUIRED; 截图类型，可选"default"/"offset"，智能截图和指定时间戳截图
	Type string `json:"Type"`

	// 截图的时间戳(ms)
	TimeOffsetMs int `json:"TimeOffsetMs,omitempty"`

	// 当 Type 为 offset 时，在TimeOffsetMs 和 TimeOffsetMsStr 之间二选一。 指定截图时间为使用模板参数动态下发的方式，取值固定为${snapshot_time}。
	TimeOffsetMsStr string `json:"TimeOffsetMsStr,omitempty"`
}

type CreateImageTemplateQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateImageTemplateRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageTemplateResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *CreateImageTemplateResResult `json:"Result,omitempty"`
}

type CreateImageTemplateResResponseMetadata struct {

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

// CreateImageTemplateResResult - title
type CreateImageTemplateResResult struct {

	// REQUIRED; 模板的创建时间
	CreateAt string `json:"CreateAt"`

	// REQUIRED; 服务id
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 创建的模板名称，如果基于原模板更新，则会生成一个新模板
	TemplateName string `json:"TemplateName"`

	// REQUIRED; 不支持的filter参数
	Unsupported []*CreateImageTemplateResResultUnsupportedItem `json:"Unsupported"`
}

type CreateImageTemplateResResultUnsupportedItem struct {

	// REQUIRED; filter名称
	Name string `json:"Name"`

	// REQUIRED; 参数对象
	Param map[string]interface{} `json:"Param"`
}

type CreateImageTemplatesByImportBody struct {

	// REQUIRED; 模板导入目标服务id
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待导入的模板JSON内容列表
	Templates []string `json:"Templates"`

	// 模板名称冲突时是否重命名(增加版本号)。默认否，忽略重名模板，不做导入
	Rename bool `json:"Rename,omitempty"`
}

type CreateImageTemplatesByImportRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageTemplatesByImportResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *CreateImageTemplatesByImportResResult `json:"Result,omitempty"`
}

type CreateImageTemplatesByImportResResponseMetadata struct {

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

// CreateImageTemplatesByImportResResult - title
type CreateImageTemplatesByImportResResult struct {

	// REQUIRED; 导入结果
	ImportResults []*CreateImageTemplatesByImportResResultImportResultsItem `json:"ImportResults"`
}

type CreateImageTemplatesByImportResResultImportResultsItem struct {

	// REQUIRED; 导入后模版名称。Success=true时有值
	ImportedName string `json:"ImportedName"`

	// REQUIRED; 导入失败原因。Success=false时有值
	Msg string `json:"Msg"`

	// REQUIRED; 导入是否成功
	Success bool `json:"Success"`

	// REQUIRED; 源模板名称
	TemplateName string `json:"TemplateName"`
}

type CreateImageTranscodeCallbackBody struct {

	// REQUIRED; 任务条目 ID
	EntryID string `json:"EntryId"`

	// 地域。
	Region string `json:"Region,omitempty"`
}

type CreateImageTranscodeCallbackRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageTranscodeCallbackResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateImageTranscodeCallbackResResponseMetadata struct {

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

type CreateImageTranscodeQueueBody struct {

	// REQUIRED; 是否启动队列，开始执行离线转码操作。取值如下所示：
	// * true：启动
	// * false：不启动
	IsStart bool `json:"IsStart"`

	// REQUIRED; 自定义任务队列名称
	Name string `json:"Name"`

	// 队列回调设置
	CallbackConf *CreateImageTranscodeQueueBodyCallbackConf `json:"CallbackConf,omitempty"`

	// 自定义任务队列描述，可用于备注该队列的用途。
	Desc string `json:"Desc,omitempty"`

	// 队列区域。默认当前区域。ToB支持取值：cn、va、sg。
	Region string `json:"Region,omitempty"`
}

// CreateImageTranscodeQueueBodyCallbackConf - 队列回调设置
type CreateImageTranscodeQueueBodyCallbackConf struct {

	// REQUIRED; 若配置CallbackConf则不为空
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 若配置CallbackConf则不为空
	Method string `json:"Method"`

	// 业务自定义回调参数，将在回调消息的callback_args中透传出去。具体回调参数请参考回调内容 [https://www.volcengine.com/docs/508/1104726#%E5%9B%9E%E8%B0%83%E5%86%85%E5%AE%B9]。
	Args string `json:"Args,omitempty"`

	// 回调数据格式。取值如下所示：
	// * XML
	// * JSON（默认）
	DataFormat string `json:"DataFormat,omitempty"`
}

type CreateImageTranscodeQueueRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageTranscodeQueueResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageTranscodeQueueResResult `json:"Result"`
}

type CreateImageTranscodeQueueResResponseMetadata struct {

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

type CreateImageTranscodeQueueResResult struct {

	// REQUIRED; 任务 ID
	QueueID string `json:"QueueId"`
}

type CreateImageTranscodeTaskBody struct {

	// REQUIRED; 数据类型，取值如下所示：
	// * uri：指定 ServiceId 下存储 URI。
	// * url：公网可访问的 URL。
	DataType string `json:"DataType"`

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 转码模板。您可通过调用GetAllImageTemplates [https://www.volcengine.com/docs/508/9386]获取指定服务下全部模版信息。
	Template string `json:"Template"`

	// 任务回调配置。缺省情况下默认使用队列回调配置。
	CallbackConf *CreateImageTranscodeTaskBodyCallbackConf `json:"CallbackConf,omitempty"`

	// 数据列表。
	DataList []string `json:"DataList,omitempty"`

	// 转码是否保留 exif。取值如下所示：
	// * true：保留
	// * false：（默认）不保留
	EnableExif bool `json:"EnableExif,omitempty"`

	// 文件列表。
	FileList []string `json:"FileList,omitempty"`

	// 任务队列名称 ID。缺省情况下提交至账号默认任务队列。您可通过调用GetImageTranscodeQueues [https://www.volcengine.com/docs/508/1160404]获取该账号下全部任务队列 ID。
	QueueID string `json:"QueueId,omitempty"`

	// 转码产物的存储 Key 列表，仅当 DataList 不为空时有效，长度需与DataList长度一致。不传时默认使用固定规则生成产物的存储 Key。
	// 存储 Key 详细命名规范请参看 veImageX 存储 Key 通用字符规则 [https://www.volcengine.com/docs/508/1458331]。
	ResKeyList []string `json:"ResKeyList,omitempty"`
}

// CreateImageTranscodeTaskBodyCallbackConf - 任务回调配置。缺省情况下默认使用队列回调配置。
type CreateImageTranscodeTaskBodyCallbackConf struct {

	// REQUIRED; Method=MQ时取值rmq:{topic}/{cluster}
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 还支持取值 MQ
	Method string `json:"Method"`

	// 业务自定义回调参数，将在回调消息的callback_args中透传出去。具体回调参数请参考回调内容 [https://www.volcengine.com/docs/508/1104726#%E5%9B%9E%E8%B0%83%E5%86%85%E5%AE%B9]。
	Args string `json:"Args,omitempty"`

	// 回调数据格式。取值如下所示：
	// * XML
	// * JSON（默认）
	DataFormat string `json:"DataFormat,omitempty"`

	// 参数类型。
	Type string `json:"Type,omitempty"`
}

type CreateImageTranscodeTaskRes struct {

	// REQUIRED
	ResponseMetadata *CreateImageTranscodeTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *CreateImageTranscodeTaskResResult `json:"Result"`
}

type CreateImageTranscodeTaskResResponseMetadata struct {

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

type CreateImageTranscodeTaskResResult struct {

	// REQUIRED; 任务 ID
	TaskID string `json:"TaskId"`
}

type CreateTemplatesFromBinBody struct {

	// REQUIRED; 待恢复模板的名称。您可以通过调用获取回收站中所有模板 [https://www.volcengine.com/docs/508/132241]获取所需的模板名称。
	TemplateNames []string `json:"TemplateNames"`
}

type CreateTemplatesFromBinQuery struct {

	// REQUIRED; 待恢复模板对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type CreateTemplatesFromBinRes struct {

	// REQUIRED
	ResponseMetadata *CreateTemplatesFromBinResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateTemplatesFromBinResResult           `json:"Result,omitempty"`
}

type CreateTemplatesFromBinResResponseMetadata struct {

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

type CreateTemplatesFromBinResResult struct {

	// REQUIRED; 返回各模板恢复的结果。
	Results []*CreateTemplatesFromBinResResultResultsItem `json:"Results"`
}

type CreateTemplatesFromBinResResultResultsItem struct {

	// REQUIRED; 回收站中的模板名。
	BinName string `json:"BinName"`

	// REQUIRED; 恢复后的模板名。
	NewName string `json:"NewName"`

	// REQUIRED; 是否恢复成功，取值如下所示：
	// * true：恢复成功
	// * false：恢复不成功
	Success bool `json:"Success"`
}

type DelDomainBody struct {

	// REQUIRED; 待删除的域名，您可以通过 获取服务下全部域名 [https://www.volcengine.com/docs/508/9379] 获取服务下域名信息。
	Domain string `json:"Domain"`
}

type DelDomainQuery struct {

	// REQUIRED; 待删除域名的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type DelDomainRes struct {

	// REQUIRED
	ResponseMetadata *DelDomainResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DelDomainResResponseMetadata struct {

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

type DeleteImageAIProcessDetailBody struct {

	// 批量删除的条目id
	Entries []string `json:"Entries,omitempty"`

	// 条目id
	EntryID string `json:"EntryId,omitempty"`
}

type DeleteImageAIProcessDetailRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageAIProcessDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteImageAIProcessDetailResResponseMetadata struct {

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

type DeleteImageAIProcessQueueBody struct {

	// REQUIRED; 队列id
	QueueID string `json:"QueueId"`
}

type DeleteImageAIProcessQueueRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageAIProcessQueueResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteImageAIProcessQueueResResponseMetadata struct {

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

type DeleteImageAnalyzeTaskBody struct {

	// REQUIRED; 待删除的任务 ID，您可以通过调用 GetImageAnalyzeTasks [https://www.volcengine.com/docs/508/1160417] 获取指定地区全部离线评估任务 ID。
	TaskID string `json:"TaskId"`
}

type DeleteImageAnalyzeTaskRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageAnalyzeTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type DeleteImageAnalyzeTaskResResponseMetadata struct {

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

type DeleteImageAnalyzeTaskRunBody struct {

	// REQUIRED; 执行 ID，您可以通过调用 GetImageAnalyzeResult [https://www.volcengine.com/docs/508/1160419] 获取该任务全部的执行记录查看执行 ID。
	RunID string `json:"RunId"`

	// REQUIRED; 待删除执行记录的任务 ID，您可以通过调用 GetImageAnalyzeTasks [https://www.volcengine.com/docs/508/1160417] 获取指定地区全部离线评估任务 ID。
	TaskID string `json:"TaskId"`
}

type DeleteImageAnalyzeTaskRunRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageAnalyzeTaskRunResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type DeleteImageAnalyzeTaskRunResResponseMetadata struct {

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

type DeleteImageAuditResultBody struct {

	// REQUIRED; 审核结果 ID，您可通过调用获取审核任务结果 [https://www.volcengine.com/docs/508/1160410]查看待更新状态的图片条目 ID。
	EntryID string `json:"EntryId"`

	// REQUIRED; 任务 ID，您可通过调用查询所有审核任务 [https://www.volcengine.com/docs/508/1160409]获取所需的任务 ID。
	TaskID string `json:"TaskId"`
}

type DeleteImageAuditResultRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageAuditResultResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type DeleteImageAuditResultResResponseMetadata struct {

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

type DeleteImageBackgroundColorsBody struct {

	// REQUIRED; 待删除的颜色列表
	Colors []string `json:"Colors"`
}

type DeleteImageBackgroundColorsRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageBackgroundColorsResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *DeleteImageBackgroundColorsResResult `json:"Result,omitempty"`
}

type DeleteImageBackgroundColorsResResponseMetadata struct {

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

// DeleteImageBackgroundColorsResResult - title
type DeleteImageBackgroundColorsResResult struct {

	// REQUIRED; 若全部删除失败，则接口将返回失败
	FailedList []string `json:"FailedList"`
}

type DeleteImageElementsBody struct {

	// REQUIRED; 待删除的 StoreUri 列表。
	ImageList []string `json:"ImageList"`

	// REQUIRED; 取值image表示图片要素，background表示背景要素
	Type DeleteImageElementsBodyType `json:"Type"`
}

type DeleteImageElementsRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageElementsResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *DeleteImageElementsResResult `json:"Result,omitempty"`
}

type DeleteImageElementsResResponseMetadata struct {

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

// DeleteImageElementsResResult - title
type DeleteImageElementsResResult struct {

	// REQUIRED; 若全部删除失败，则接口将返回失败
	FailedList []string `json:"FailedList"`
}

type DeleteImageMigrateTaskQuery struct {

	// REQUIRED; 仅当任务状态为非Running时生效。 任务 ID，请参考GetImageMigrateTasks [https://www.volcengine.com/docs/508/1108670]获取返回的任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`

	// 任务地区（即任务目标服务的地区），默认空，返回国内任务。
	// * cn：国内
	// * sg：新加坡
	Region string `json:"Region,omitempty" query:"Region"`
}

type DeleteImageMigrateTaskRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageMigrateTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteImageMigrateTaskResResponseMetadata struct {

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

type DeleteImageMonitorRecordsBody struct {

	// REQUIRED; 待删除的报警记录 Marker 列表，您可通过调用GetImageAlertRecords [https://www.volcengine.com/docs/508/1112187]获取所需值。
	Markers []string `json:"Markers"`
}

type DeleteImageMonitorRecordsRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageMonitorRecordsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DeleteImageMonitorRecordsResResult `json:"Result,omitempty"`
}

type DeleteImageMonitorRecordsResResponseMetadata struct {

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

// DeleteImageMonitorRecordsResResult - 视请求的接口而定
type DeleteImageMonitorRecordsResResult struct {

	// REQUIRED; 成功删除的报警记录列表
	DeletedRecords []string `json:"DeletedRecords"`
}

type DeleteImageMonitorRulesBody struct {

	// REQUIRED; 待删除的告警规则 ID 列表，您可以调用GetImageMonitorRules [https://www.volcengine.com/docs/508/1112186]获取所需的告警规则 ID。
	RuleIDs []string `json:"RuleIds"`
}

type DeleteImageMonitorRulesRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageMonitorRulesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DeleteImageMonitorRulesResResult `json:"Result,omitempty"`
}

type DeleteImageMonitorRulesResResponseMetadata struct {

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

// DeleteImageMonitorRulesResResult - 视请求的接口而定
type DeleteImageMonitorRulesResResult struct {

	// REQUIRED; 成功删除的告警规则 ID 列表
	DeletedRules []string `json:"DeletedRules"`
}

type DeleteImageServiceQuery struct {

	// REQUIRED; 待删除的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type DeleteImageServiceRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageServiceResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteImageServiceResResponseMetadata struct {

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

type DeleteImageSettingRuleBody struct {

	// REQUIRED; 应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppId。
	AppID string `json:"AppId"`

	// REQUIRED; 待删除的规则 ID，您可以通过调用获取规则列表 [https://www.volcengine.com/docs/508/1324618]的方式获取所需的规则 ID。
	RuleID string `json:"RuleId"`

	// REQUIRED; 配置项 ID，您可以通过调用获取配置项列表 [https://www.volcengine.com/docs/508/1324617]的方式获取所需的配置项 ID。
	SettingID string `json:"SettingId"`
}

type DeleteImageSettingRuleRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageSettingRuleResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteImageSettingRuleResResponseMetadata struct {

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

type DeleteImageStyleBody struct {

	// REQUIRED; 待删除的样式 ID。
	StyleID string `json:"StyleId"`
}

type DeleteImageStyleRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageStyleResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result interface{} `json:"Result,omitempty"`
}

type DeleteImageStyleResResponseMetadata struct {

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

type DeleteImageTemplateBody struct {

	// REQUIRED; 待删除模板名称，不能超过100个。模板需要从属于参数中的service。
	TemplateNames []string `json:"TemplateNames"`
}

type DeleteImageTemplateQuery struct {

	// REQUIRED; 待删除模板对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type DeleteImageTemplateRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageTemplateResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *DeleteImageTemplateResResult `json:"Result,omitempty"`
}

type DeleteImageTemplateResResponseMetadata struct {

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

// DeleteImageTemplateResResult - title
type DeleteImageTemplateResResult struct {

	// REQUIRED; 成功删除的模板名称。如果入参为单个模板，删除失败直接返回错误。如果为多个模板，全部删除失败，返回错误；否则返回成功删除的模板。
	TemplateNames []string `json:"TemplateNames"`
}

type DeleteImageTranscodeDetailBody struct {

	// 待删除的entry ID列表
	Entries []string `json:"Entries,omitempty"`

	// 待删除的entry ID
	EntryID string `json:"EntryId,omitempty"`
}

type DeleteImageTranscodeDetailRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageTranscodeDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type DeleteImageTranscodeDetailResResponseMetadata struct {

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

type DeleteImageTranscodeQueueBody struct {

	// REQUIRED; 待删除的队列 ID，您可通过调用GetImageTranscodeQueues [https://www.volcengine.com/docs/508/1107341]获取该账号下全部任务队列 ID。 账号内置默认任务队列不允许被删除。
	QueueID string `json:"QueueId"`
}

type DeleteImageTranscodeQueueRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageTranscodeQueueResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type DeleteImageTranscodeQueueResResponseMetadata struct {

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

type DeleteImageUploadFilesBody struct {

	// REQUIRED; 待删除文件的存储 URI 列表，最多传 1000 个。您可以通过调用获取服务下的上传文件 [https://www.volcengine.com/docs/508/9392]来获取所需的文件 URI。
	StoreUris []string `json:"StoreUris"`

	// 待删除文件的存储版本 ID。传值时需要和 StoreUris 一一对应。您可在 veImageX 控制台资源管理查看文件版本号，或调用 GetImageStorageVersionFiles [https://www.volcengine.com/docs/508/1359366]
	// 查询服务下所有文件的版本信息。
	// :::warning 当删除文件未指定 StoreVersions，那么删除逻辑将根据版本控制的状态产生以下差异。
	// * 若此时版本控制为未开启，则 StoreUris 对应文件将被永久删除，不可恢复。
	// * 若此时版本控制为开启，则 StoreUris 对应文件未被真正删除，该文件可以被恢复，同时将增加一个删除标记 [https://www.volcengine.com/docs/508/1359410]用于标识该文件为删除状态。
	// * 若此时版本控制为暂停，则根据版本 ID 是否为 null 而有以下差异： * 若文件的版本 ID 为 null，则 StoreUris 对应文件被真正删除，不可恢复，同时将增加一个删除标记 [https://www.volcengine.com/docs/508/1359410]用于标识该文件为删除状态。
	// * 若文件的版本 ID 不为 null，则 StoreUris 对应文件未被真正删除，而是转换为历史版本保留。该文件可以被恢复，同时将增加一个删除标记 [https://www.volcengine.com/docs/508/1359410]。
	// :::
	StoreVersions []string `json:"StoreVersions,omitempty"`
}

type DeleteImageUploadFilesQuery struct {

	// REQUIRED; 待删除文件所在的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type DeleteImageUploadFilesRes struct {

	// REQUIRED
	ResponseMetadata *DeleteImageUploadFilesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DeleteImageUploadFilesResResult `json:"Result,omitempty"`
}

type DeleteImageUploadFilesResResponseMetadata struct {

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

// DeleteImageUploadFilesResResult - 视请求的接口而定
type DeleteImageUploadFilesResResult struct {

	// REQUIRED; 文件成功删除的 URI 列表
	DeletedFiles []string `json:"DeletedFiles"`

	// REQUIRED; 文件不存在的无效 URI 列表
	InvaildFiles []string `json:"InvaildFiles"`

	// 已删除文件的版本列表。
	DeletedFilesVersion []string `json:"DeletedFilesVersion,omitempty"`

	// 无效文件版本列表。
	InvaildFilesVersion []string `json:"InvaildFilesVersion,omitempty"`
}

type DeleteTemplatesFromBinBody struct {

	// REQUIRED; 待删除模板名称。您可以通过调用获取回收站中所有模板 [https://www.volcengine.com/docs/508/132241]获取所需的模板名称。
	TemplateNames []string `json:"TemplateNames"`
}

type DeleteTemplatesFromBinQuery struct {

	// REQUIRED; 待删除模板对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type DeleteTemplatesFromBinRes struct {

	// REQUIRED
	ResponseMetadata *DeleteTemplatesFromBinResResponseMetadata `json:"ResponseMetadata"`
	Result           *DeleteTemplatesFromBinResResult           `json:"Result,omitempty"`
}

type DeleteTemplatesFromBinResResponseMetadata struct {

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

type DeleteTemplatesFromBinResResult struct {

	// REQUIRED; 成功删除的模板名称。
	// * 如果指定为单个模板，删除失败直接返回错误。
	// * 如果指定多个模板，如果全部删除失败则返回错误；否则返回成功删除的模板。
	TemplateNames []string `json:"TemplateNames"`
}

type DescribeImageVolcCdnAccessLogBody struct {

	// REQUIRED; 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 指定一个结束时间对日志进行过滤，时间格式是 Unix 时间戳。 例子1：您指定 EndTime 为 1641953589，代表 2022-01-12 10:13:09 UTC。在查询返回的日志列表中，最后一个统计在内的日志包名称是
	// domain2022011210000020220112110000.gz。该日志包中包含该域名在 10:00:00 和 10:59:59 之间的访问日志。
	EndTime int `json:"EndTime"`

	// REQUIRED; 指定一个页码。系统只返回该页面上的那些日志包。 默认值为 1。如果指定的页码不存在，则返回空值。建议第一次提交查询时使用默认值。您可以根据响应正文中的 TotalCount 和 PageSize 参数的值计算结果页数。然后再查询中指定一个
	// PageNum 来获取该页面上的那些日志包。
	PageNum int `json:"PageNum"`

	// REQUIRED; 指定查询结果中每页包含的日志包数量。 默认值是 10，最大值是 100。
	PageSize int `json:"PageSize"`

	// REQUIRED; 指定一个加速区域对日志进行过滤。该参数仅当您通过提交工单启用了 中国大陆 以外的加速区域时才有效。 该参数有以下取值： chinesemainland：表示获取访问请求是来自 中国大陆 的日志。查询返回的日志包的名称不包含
	// outsideChineseMainland。 global：表示获取访问请求是来自 全球 的日志。 outsidechinese_mainland：表示获取访问请求是来自
	// 全球（除中国大陆） 的日志。查询返回的日志包的名称包含 outsideChineseMainland。 该参数的默认值是 global
	Region string `json:"Region"`

	// REQUIRED; 指定一个开始时间对日志进行过滤，时间格式是 Unix 时间戳。 例子1：您指定 StartTime 为 1641844915，代表 2022-01-11 04:01:55 UTC。在查询返回的日志列表中，第一个统计在内的日志包名称是domain2022011105000020220111060000.gz。该日志包中包含该域名在
	// 05:00:00 和 05:59:59 之间的访问日志。
	// StartTime 和 EndTime 之间的时间范围不能大于 30 天。
	StartTime int `json:"StartTime"`
}

type DescribeImageVolcCdnAccessLogQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type DescribeImageVolcCdnAccessLogRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageVolcCdnAccessLogResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageVolcCdnAccessLogResResult `json:"Result"`
}

type DescribeImageVolcCdnAccessLogResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED
	Error *DescribeImageVolcCdnAccessLogResResponseMetadataError `json:"Error"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DescribeImageVolcCdnAccessLogResResponseMetadataError struct {

	// REQUIRED; 错误码
	Code string `json:"Code"`

	// REQUIRED; 错误信息
	Message string `json:"Message"`
}

type DescribeImageVolcCdnAccessLogResResult struct {

	// REQUIRED; 加速域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的日志数据。
	Logs []*DescribeImageVolcCdnAccessLogResResultLogsItem `json:"Logs"`

	// REQUIRED; 指定的页码数。
	PageNum int `json:"PageNum"`

	// REQUIRED; 指定的每页日志包数量。
	PageSize int `json:"PageSize"`

	// REQUIRED; 日志包总数。
	TotalCount int `json:"TotalCount"`
}

type DescribeImageVolcCdnAccessLogResResultLogsItem struct {

	// REQUIRED; 日志结束时间。
	EndTime int `json:"EndTime"`

	// REQUIRED; 日志包名称。
	LogName string `json:"LogName"`

	// REQUIRED; 日志包下载地址，下载链接有效时间为 10 分钟。
	LogPath string `json:"LogPath"`

	// REQUIRED; 日志包大小，单位为 byte。
	LogSize int `json:"LogSize"`

	// REQUIRED; 日志起始时间。
	StartTime int `json:"StartTime"`
}

type DescribeImageXAIRequestCntUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 组件名称。支持查询多个组件，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有组件。您可通过调用 GetImageAddOnConf [https://www.volcengine.com/docs/508/66145] 查看Key返回值。
	AdvFeats string `json:"AdvFeats,omitempty" query:"AdvFeats"`

	// 是否叠加计费倍率。默认为false。
	EnableBillingRate bool `json:"EnableBillingRate,omitempty" query:"EnableBillingRate"`

	// 维度拆分的维度值。不传表示不拆分维度，传入多个时用英文逗号“,”分割。支持取值如下：
	// * ServiceId：服务
	// * AdvFeat：组件
	// * Model：模型
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒。缺省时查询 StartTime 和 EndTime 时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 图片处理模板。支持查询多个模板，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有模板。您可通过调用 GetAllImageTemplates [https://www.volcengine.com/docs/508/9386] 获取服务下全部模版信息。
	Templates string `json:"Templates,omitempty" query:"Templates"`
}

type DescribeImageXAIRequestCntUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXAIRequestCntUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXAIRequestCntUsageResResult `json:"Result"`
}

type DescribeImageXAIRequestCntUsageResResponseMetadata struct {

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

type DescribeImageXAIRequestCntUsageResResult struct {

	// REQUIRED; 请求次数据。
	RequestCntData []*DescribeImageXAIRequestCntUsageResResultRequestCntDataItem `json:"RequestCntData"`
}

type DescribeImageXAIRequestCntUsageResResultRequestCntDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXAIRequestCntUsageResResultRequestCntDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 附加组件模型，GroupBy包含Model时有返回值。
	Model string `json:"Model"`

	// 附加组件类型，GroupBy包含AdvFeat时有返回值。如：enhance，smartcut。取值为total，表示包含所选AdvFeat总请求次数。
	AdvFeat string `json:"AdvFeat,omitempty"`

	// 服务 ID，GroupBy包含ServiceId时有返回值。
	ServiceID string `json:"ServiceId,omitempty"`
}

type DescribeImageXAIRequestCntUsageResResultRequestCntDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求次
	Value int `json:"Value"`
}

type DescribeImageXAddOnQPSUsageQuery struct {

	// REQUIRED; 附加组件名称。仅传单个。您可通过调用 GetImageAddOnConf [https://www.volcengine.com/docs/508/66145] 查看Key返回值。
	AdvFeat string `json:"AdvFeat" query:"AdvFeat"`

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 查询数据的时间粒度。单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 1: 单次查询最大时间跨度为 6 小时
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXAddOnQPSUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXAddOnQPSUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *DescribeImageXAddOnQPSUsageResResult `json:"Result"`
}

type DescribeImageXAddOnQPSUsageResResponseMetadata struct {

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

// DescribeImageXAddOnQPSUsageResResult - 视请求的接口而定
type DescribeImageXAddOnQPSUsageResResult struct {

	// REQUIRED; QPS 用量
	QPSData []*DescribeImageXAddOnQPSUsageResResultQPSDataItem `json:"QPSData"`
}

type DescribeImageXAddOnQPSUsageResResultQPSDataItem struct {

	// REQUIRED; 时序数据。
	Data []*DescribeImageXAddOnQPSUsageResResultQPSDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 附加组件模型。
	// * 取值为total：表示总请求QPS。
	// * 取值不为total：表示附加组件使用模型的请求QPS。
	Model string `json:"Model"`
}

type DescribeImageXAddOnQPSUsageResResultQPSDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求QPS
	Value int `json:"Value"`
}

type DescribeImageXBaseOpUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 需要分组查询的参数，当前仅支持取值 ServiceId，表示按服务 ID 进行分组。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒。缺省时查询 StartTime 和 EndTime 时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXBaseOpUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXBaseOpUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXBaseOpUsageResResult `json:"Result"`
}

type DescribeImageXBaseOpUsageResResponseMetadata struct {

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

type DescribeImageXBaseOpUsageResResult struct {

	// REQUIRED; 计量数据列表
	BaseOpData []*DescribeImageXBaseOpUsageResResultBaseOpDataItem `json:"BaseOpData"`
}

type DescribeImageXBaseOpUsageResResultBaseOpDataItem struct {

	// REQUIRED; 计量数据
	Data []*DescribeImageXBaseOpUsageResResultBaseOpDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy中包含ServiceId时出现
	ServiceID string `json:"ServiceId,omitempty"`
}

type DescribeImageXBaseOpUsageResResultBaseOpDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 基础处理量，单位Byte
	Value float64 `json:"Value"`
}

type DescribeImageXBillingRequestCntUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 固定值，仅支持AdvFeat即附加组件。
	GroupBy string `json:"GroupBy" query:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 组件名称。支持查询多个组件，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有组件。您可通过调用GetImageAddOnConf [https://www.volcengine.com/docs/508/66145]查看Key返回值。
	AdvFeats string `json:"AdvFeats,omitempty" query:"AdvFeats"`

	// 查询数据的时间粒度。单位为秒。缺省时查询 StartTime 和 EndTime 时间段全部数据，此时单次查询最大时间跨度为 93 天。支持取值如下：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXBillingRequestCntUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXBillingRequestCntUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXBillingRequestCntUsageResResult `json:"Result"`
}

type DescribeImageXBillingRequestCntUsageResResponseMetadata struct {

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

type DescribeImageXBillingRequestCntUsageResResult struct {

	// REQUIRED; 附加组件通用请求次数据。
	RequestCntData []*DescribeImageXBillingRequestCntUsageResResultRequestCntDataItem `json:"RequestCntData"`
}

type DescribeImageXBillingRequestCntUsageResResultRequestCntDataItem struct {

	// REQUIRED; 附加组件类型，如：enhance，smartcut。 取值为total，表示包含所选AdvFeat请求次数相加。
	AdvFeat string `json:"AdvFeat"`

	// REQUIRED; 时序数据列表。
	Data []*DescribeImageXBillingRequestCntUsageResResultRequestCntDataPropertiesItemsItem `json:"Data"`
}

type DescribeImageXBillingRequestCntUsageResResultRequestCntDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求次
	Value int `json:"Value"`
}

type DescribeImageXBucketRetrievalUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// Bucket 名称。支持同时查询多个 BucketName，不同的 BucketNmae 使用逗号分隔。 您可以通过调用 GetAllImageServices [https://www.volcengine.com/docs/508/9360]
	// 获取所需的 Bucket 名称。
	BucketNames string `json:"BucketNames,omitempty" query:"BucketNames"`

	// 需要分组查询的参数，多个数据用逗号分隔。支持取值如下：
	// * ServiceId：服务 ID
	// * BucketName：Bucket 名称
	// * StorageType：存储类型
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 是否查询数据取回量。
	// * true：查询取回量。
	// * false：查询存储量。
	IsRetrieval bool `json:"IsRetrieval,omitempty" query:"IsRetrieval"`

	// 服务 ID。为空时表示不筛选，支持查询多个服务，使用逗号分隔不同的服务。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXBucketRetrievalUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXBucketRetrievalUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXBucketRetrievalUsageResResult `json:"Result"`
}

type DescribeImageXBucketRetrievalUsageResResponseMetadata struct {

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

type DescribeImageXBucketRetrievalUsageResResult struct {

	// REQUIRED; 计量数据列表
	StorageData []*DescribeImageXBucketRetrievalUsageResResultStorageDataItem `json:"StorageData"`
}

type DescribeImageXBucketRetrievalUsageResResultStorageDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXBucketRetrievalUsageResResultStorageDataPropertiesItemsItem `json:"Data"`

	// Bucket 名称，GroupBy包含BucketName时有返回值。
	BucketName string `json:"BucketName,omitempty"`

	// 服务 ID，GroupBy包含ServiceId时有返回值。
	ServiceID string `json:"ServiceId,omitempty"`

	// 存储类型，GroupBy包含StorageType时有返回值。 当传参IsRetrieval=false时，表示存储值，取值：
	// * STANDARD：标准存储
	// * IA：低频存储
	// * ARCHIVE：归档存储
	// * COLD_ARCHIVE：冷归档存储
	// * ARCHIVE_FR：归档闪回存储
	// 当传参IsRetrieval=true时，表示取回值，取值：
	// * IA：低频取回
	// * ARCHIVE：归档标准取回
	// * COLD_ARCHIVE：冷归档快速取回
	// * ARCHIVE_FR：归档闪回取回
	// * COLD_ARCHIVE_STANDARD：冷归档标准取回
	// * COLD_ARCHIVE_BULK：冷归档批量取回
	StorageType string `json:"StorageType,omitempty"`
}

type DescribeImageXBucketRetrievalUsageResResultStorageDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 单位：Byte
	Value float64 `json:"Value"`
}

type DescribeImageXBucketUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// Bucket 名称。支持同时查询多个 BucketName，不同的 BucketNmae 使用逗号分隔。 您可以通过调用 GetAllImageServices [https://www.volcengine.com/docs/508/9360]
	// 获取所需的 Bucket 名称。
	BucketNames string `json:"BucketNames,omitempty" query:"BucketNames"`

	// 需要分组查询的参数，多个数据用逗号分隔。支持取值如下：
	// * ServiceId：服务 ID
	// * BucketName：Bucket 名称
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 服务 ID。为空时表示不筛选，支持查询多个服务，使用逗号分隔不同的服务。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXBucketUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXBucketUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXBucketUsageResResult `json:"Result"`
}

type DescribeImageXBucketUsageResResponseMetadata struct {

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

type DescribeImageXBucketUsageResResult struct {

	// REQUIRED; 计量数据列表
	StorageData []*DescribeImageXBucketUsageResResultStorageDataItem `json:"StorageData"`
}

type DescribeImageXBucketUsageResResultStorageDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXBucketUsageResResultStorageDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy中包含BucketName时出现
	BucketName string `json:"BucketName,omitempty"`

	// 当GroupBy中包含ServiceId时出现
	ServiceID string `json:"ServiceId,omitempty"`
}

type DescribeImageXBucketUsageResResultStorageDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 单位：Byte
	Value float64 `json:"Value"`
}

type DescribeImageXCDNTopRequestDataQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如 2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 排序依据，取值如下所示：
	// * URL：生成的图片访问 URL
	// * UserAgent：用户代理
	// * Refer：请求 Refer
	// * ClientIP：客户端 IP
	// * Region：访问区域
	// * Domain：域名
	// * Isp：运营商
	KeyType string `json:"KeyType" query:"KeyType"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如 2019-06-02T00:00:00+08:00。
	StartTime string `json:"StartTime" query:"StartTime"`

	// REQUIRED; 排序依据，即获取按ValueType值排序的KeyType列表。取值如下所示：
	// * Traffic：按流量排序
	// * RequestCnt：按请求次数排序
	// :::tip 当KeyType取值为Domain时，仅支持将ValueType取值为Traffic，即按照流量排序获取域名列表。 :::
	ValueType string `json:"ValueType" query:"ValueType"`

	// 数据访问区域。仅在KeyType取值为Region或Isp时生效，取值如下所示：
	// * China：中国。
	// * Other：中国境外的区域。
	Country string `json:"Country,omitempty" query:"Country"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用 GetServiceDomains [https://www.volcengine.com/docs/508/9379] 获取所需的域名。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 网络协议。缺省情况下则表示不限制网络协议，取值如下所示：
	// * IPv4
	// * IPv6
	// :::tipKeyType取值为Domain时，IPVersion的取值无效。 :::
	IPVersion string `json:"IPVersion,omitempty" query:"IPVersion"`

	// 每页查询数据量，默认为0，即返回所有数据。
	Limit string `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，默认取值为0。取值为10时，表示跳过前 10 条数据，从第 11 条数据开始取值。
	Offset string `json:"Offset,omitempty" query:"Offset"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXCDNTopRequestDataRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCDNTopRequestDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCDNTopRequestDataResResult `json:"Result"`
}

type DescribeImageXCDNTopRequestDataResResponseMetadata struct {

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

type DescribeImageXCDNTopRequestDataResResult struct {

	// REQUIRED; 可展示的数据总数量。
	Count int `json:"Count"`

	// REQUIRED; 数据列表，按Value降序排列。
	TopValue []*DescribeImageXCDNTopRequestDataResResultTopValueItem `json:"TopValue"`

	// REQUIRED; 全量Traffic/RequestCnt的总量，用于计算占比。Traffic单位：Byte。RequestCnt单位：次。
	TotalValue float64 `json:"TotalValue"`
}

type DescribeImageXCDNTopRequestDataResResultTopValueItem struct {

	// REQUIRED; 对应于参数KeyType，维度。
	Key string `json:"Key"`

	// REQUIRED; Traffic/RequestCnt，指标值。Traffic单位：Byte。RequestCnt单位：次。
	Value float64 `json:"Value"`
}

type DescribeImageXCdnDurationAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 聚合维度。 Domain：域名； Region：地区； Isp：运营商。
	GroupBy string `json:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnDurationAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc bool `json:"OrderAsc,omitempty"`

	// 排序依据。 Duration：按耗时排序。 Count：（默认）按上报量排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnDurationAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnDurationAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnDurationAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnDurationAllResResult `json:"Result"`
}

type DescribeImageXCdnDurationAllResResponseMetadata struct {

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

type DescribeImageXCdnDurationAllResResult struct {

	// REQUIRED; 网络耗时数据。
	DurationData []*DescribeImageXCdnDurationAllResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXCdnDurationAllResResultDurationDataItem struct {

	// REQUIRED; 上报数据量。
	Count int `json:"Count"`

	// REQUIRED; 网络耗时，单位毫秒
	Value float64 `json:"Value"`

	// 当GroupBy取值Domain时出现，则表示域名信息。
	Domain string `json:"Domain,omitempty"`

	// 当GroupBy取值Isp时出现，则表示运营商信息。
	Isp string `json:"Isp,omitempty"`

	// 当GroupBy取值Region时出现，表示地区信息。
	Region string `json:"Region,omitempty"`

	// 当GroupBy取值Region时出现，表示地区类型。 Country表示国家。 Province表示省份。
	RegionType string `json:"RegionType,omitempty"`
}

type DescribeImageXCdnDurationDetailByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnDurationDetailByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：Duration（拆分网络耗时分位数据）、Phase（拆分网络耗时分布数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 指定查询特定阶段的耗时数据。默认空，返回总耗时数据。 dns： DNS 耗时，即网络阶段的 DNS 平均耗时。 ssl： SSL 耗时，即网络阶段的 SSL 握手平均耗时。 connect：连接耗时，即网络阶段的 TCP 建立连接平均耗时。
	// send：发送耗时，即网络阶段的发送数据平均耗时。 wait：等待耗时，即网络阶段发送完数据后等待首个回包字节的耗时。
	// receive：接收耗时，即网络阶段的接收数据耗时。 proxy：代理耗时，即网络阶段的建立代理连接的耗时。
	Phase string `json:"Phase,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnDurationDetailByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnDurationDetailByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnDurationDetailByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnDurationDetailByTimeResResult `json:"Result"`
}

type DescribeImageXCdnDurationDetailByTimeResResponseMetadata struct {

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

type DescribeImageXCdnDurationDetailByTimeResResult struct {

	// REQUIRED; 网络耗时数据
	DurationData []*DescribeImageXCdnDurationDetailByTimeResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXCdnDurationDetailByTimeResResultDurationDataItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的网络耗时数据列表。
	Data []*DescribeImageXCdnDurationDetailByTimeResResultDurationDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空时，取值为：Total。 当GroupBy为Duration时，取值为：pct25、pct50、pct90、pct99、avg。 当GroupBy为Phase时，取值为：dns、ssl、connect、send、wait、receive、proxy。
	// 除上述外取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXCdnDurationDetailByTimeResResultDurationDataPropertiesItemsItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 平均耗时，单位毫秒
	Value float64 `json:"Value"`
}

type DescribeImageXCdnErrorCodeAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 聚合维度。 Domain：域名； ErrorCode：错误码； Region：地区； Isp：运营商。
	GroupBy string `json:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnErrorCodeAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc bool `json:"OrderAsc,omitempty"`

	// 目前仅支持传入Count按错误码数量排序。不传或者传空默认按错误码数量排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnErrorCodeAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnErrorCodeAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnErrorCodeAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnErrorCodeAllResResult `json:"Result"`
}

type DescribeImageXCdnErrorCodeAllResResponseMetadata struct {

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

type DescribeImageXCdnErrorCodeAllResResult struct {

	// REQUIRED; 错误码数据
	ErrorCodeData []*DescribeImageXCdnErrorCodeAllResResultErrorCodeDataItem `json:"ErrorCodeData"`
}

type DescribeImageXCdnErrorCodeAllResResultErrorCodeDataItem struct {

	// REQUIRED; 错误码数量
	Value int `json:"Value"`

	// 当GroupBy取值非ErrorCode时出现，表示错误码详细信息。
	Details []*DescribeImageXCdnErrorCodeAllResResultErrorCodeDataPropertiesItemsItem `json:"Details,omitempty"`

	// 当GroupBy取值Domain时出现，则表示域名信息。
	Domain string `json:"Domain,omitempty"`

	// 当GroupBy取值ErrorCode时出现，表示错误码内容。
	ErrorCode string `json:"ErrorCode,omitempty"`

	// 当GroupBy取值ErrorCode时出现，表示错误码的描述信息。
	ErrorCodeDesc string `json:"ErrorCodeDesc,omitempty"`

	// 当GroupBy取值Isp时出现，则表示运营商信息。
	Isp string `json:"Isp,omitempty"`

	// 当GroupBy取值Region时出现，表示地区信息。
	Region string `json:"Region,omitempty"`

	// 当GroupBy取值Region时出现，表示地区类型。 取值Country时，表示国家； 取值Province时，表示省份。
	RegionType string `json:"RegionType,omitempty"`
}

type DescribeImageXCdnErrorCodeAllResResultErrorCodeDataPropertiesItemsItem struct {

	// REQUIRED; 错误码内容
	ErrorCode string `json:"ErrorCode"`

	// REQUIRED; 错误码的描述信息
	ErrorCodeDesc string `json:"ErrorCodeDesc"`

	// REQUIRED; 错误码数量
	Value int `json:"Value"`
}

type DescribeImageXCdnErrorCodeByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnErrorCodeByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnErrorCodeByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnErrorCodeByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnErrorCodeByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnErrorCodeByTimeResResult `json:"Result"`
}

type DescribeImageXCdnErrorCodeByTimeResResponseMetadata struct {

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

type DescribeImageXCdnErrorCodeByTimeResResult struct {

	// REQUIRED; 所有错误码数据。
	ErrorCodeData []*DescribeImageXCdnErrorCodeByTimeResResultErrorCodeDataItem `json:"ErrorCodeData"`
}

type DescribeImageXCdnErrorCodeByTimeResResultErrorCodeDataItem struct {

	// REQUIRED; 错误码数量
	Count int `json:"Count"`

	// REQUIRED; 错误码对应的时序数据。
	Data []*DescribeImageXCdnErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 错误码
	ErrorCode string `json:"ErrorCode"`
}

type DescribeImageXCdnErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem struct {

	// REQUIRED; 错误码数量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 错误码数量
	Value int `json:"Value"`
}

type DescribeImageXCdnProtocolRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnProtocolRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnProtocolRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnProtocolRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnProtocolRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnProtocolRateByTimeResResult `json:"Result"`
}

type DescribeImageXCdnProtocolRateByTimeResResponseMetadata struct {

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

type DescribeImageXCdnProtocolRateByTimeResResult struct {

	// REQUIRED; 各协议占比数据
	ProtocolRateData []*DescribeImageXCdnProtocolRateByTimeResResultProtocolRateDataItem `json:"ProtocolRateData"`
}

type DescribeImageXCdnProtocolRateByTimeResResultProtocolRateDataItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的协议占比数据列表。
	Data []*DescribeImageXCdnProtocolRateByTimeResResultProtocolRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 网络协议类型。 取值为：http、https
	Type string `json:"Type"`
}

type DescribeImageXCdnProtocolRateByTimeResResultProtocolRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 网络协议占比
	Value float64 `json:"Value"`
}

type DescribeImageXCdnReuseRateAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 聚合维度 Domain：域名； Region：地区； Isp：运营商。
	GroupBy string `json:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnReuseRateAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc bool `json:"OrderAsc,omitempty"`

	// 维度区分，不传或者传空默认按上报量排序。 ReuseRate：按连接复用率排序； Count：按上报量排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnReuseRateAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnReuseRateAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnReuseRateAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnReuseRateAllResResult `json:"Result"`
}

type DescribeImageXCdnReuseRateAllResResponseMetadata struct {

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

type DescribeImageXCdnReuseRateAllResResult struct {

	// REQUIRED; 连接复用率数据。
	ReuseRateData []*DescribeImageXCdnReuseRateAllResResultReuseRateDataItem `json:"ReuseRateData"`
}

type DescribeImageXCdnReuseRateAllResResultReuseRateDataItem struct {

	// REQUIRED; 上报数据量。
	Count int `json:"Count"`

	// REQUIRED; 连接复用率。
	Value float64 `json:"Value"`

	// 当GroupBy取值Domain时出现，表示域名信息。
	Domain string `json:"Domain,omitempty"`

	// 当GroupBy取值Isp时出现，表示运营商信息。
	Isp string `json:"Isp,omitempty"`

	// 当GroupBy取值Region时出现，表示地区信息。
	Region string `json:"Region,omitempty"`

	// 当GroupBy取值Region时出现，表示地区类型。 取值Country，表示国家。 取值Province，表示省份。
	RegionType string `json:"RegionType,omitempty"`
}

type DescribeImageXCdnReuseRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnReuseRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnReuseRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnReuseRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnReuseRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnReuseRateByTimeResResult `json:"Result"`
}

type DescribeImageXCdnReuseRateByTimeResResponseMetadata struct {

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

type DescribeImageXCdnReuseRateByTimeResResult struct {

	// REQUIRED; 连接复用率数据
	ReuseRateData []*DescribeImageXCdnReuseRateByTimeResResultReuseRateDataItem `json:"ReuseRateData"`
}

type DescribeImageXCdnReuseRateByTimeResResultReuseRateDataItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的连接复用率数据列表。
	Data []*DescribeImageXCdnReuseRateByTimeResResultReuseRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXCdnReuseRateByTimeResResultReuseRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 网络连接复用率。
	Value float64 `json:"Value"`
}

type DescribeImageXCdnSuccessRateAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 聚合维度。 Domain：域名； Region：地区； Isp：运营商。
	GroupBy string `json:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnSuccessRateAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc bool `json:"OrderAsc,omitempty"`

	// 取值为SuccessRate时，表示按网络成功率排序； 取值为Count时，表示按上报量排序； 不传或者传空默认按上报量排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnSuccessRateAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnSuccessRateAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnSuccessRateAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnSuccessRateAllResResult `json:"Result"`
}

type DescribeImageXCdnSuccessRateAllResResponseMetadata struct {

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

type DescribeImageXCdnSuccessRateAllResResult struct {

	// REQUIRED; 网络成功率数据。
	SuccessRateData []*DescribeImageXCdnSuccessRateAllResResultSuccessRateDataItem `json:"SuccessRateData"`
}

type DescribeImageXCdnSuccessRateAllResResultSuccessRateDataItem struct {

	// REQUIRED; 上报数据量。
	Count int `json:"Count"`

	// REQUIRED; 网络成功率
	Value float64 `json:"Value"`

	// 当GroupBy取值Domain时出现，则表示域名信息。
	Domain string `json:"Domain,omitempty"`

	// 当GroupBy取值Isp时出现，则表示运营商信息。
	Isp string `json:"Isp,omitempty"`

	// 当GroupBy取值Region时出现，表示地区信息。
	Region string `json:"Region,omitempty"`

	// 当GroupBy取值Region时出现，表示地区类型。 取值Country，表示国家； 取值Province，表示省份。
	RegionType string `json:"RegionType,omitempty"`
}

type DescribeImageXCdnSuccessRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXCdnSuccessRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXCdnSuccessRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXCdnSuccessRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCdnSuccessRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCdnSuccessRateByTimeResResult `json:"Result"`
}

type DescribeImageXCdnSuccessRateByTimeResResponseMetadata struct {

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

type DescribeImageXCdnSuccessRateByTimeResResult struct {

	// REQUIRED; 网络成功率数据
	SuccessRateData []*DescribeImageXCdnSuccessRateByTimeResResultSuccessRateDataItem `json:"SuccessRateData"`
}

type DescribeImageXCdnSuccessRateByTimeResResultSuccessRateDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 具体数据
	Data []*DescribeImageXCdnSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXCdnSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 网络成功率
	Value float64 `json:"Value"`
}

type DescribeImageXClientCountByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientCountByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientCountByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientCountByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientCountByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientCountByTimeResResult `json:"Result"`
}

type DescribeImageXClientCountByTimeResResponseMetadata struct {

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

type DescribeImageXClientCountByTimeResResult struct {

	// REQUIRED; 客户端上报量数据。
	ClientCountData []*DescribeImageXClientCountByTimeResResultClientCountDataItem `json:"ClientCountData"`
}

type DescribeImageXClientCountByTimeResResultClientCountDataItem struct {

	// REQUIRED; 该数据类型对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的客户端上报量数据列表。
	Data []*DescribeImageXClientCountByTimeResResultClientCountDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXClientCountByTimeResResultClientCountDataPropertiesItemsItem struct {

	// REQUIRED; 上报量数据
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 上报量数据
	Value int `json:"Value"`
}

type DescribeImageXClientDecodeDurationByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientDecodeDurationByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientDecodeDurationByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientDecodeDurationByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientDecodeDurationByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientDecodeDurationByTimeResResult `json:"Result"`
}

type DescribeImageXClientDecodeDurationByTimeResResponseMetadata struct {

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

type DescribeImageXClientDecodeDurationByTimeResResult struct {

	// REQUIRED; 解码耗时数据
	DurationData []*DescribeImageXClientDecodeDurationByTimeResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXClientDecodeDurationByTimeResResultDurationDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的解码耗时数据列表。
	Data []*DescribeImageXClientDecodeDurationByTimeResResultDurationDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。当GroupBy为空或Duration时，取值avg/pct25/pct50/pct90/pct99/min/max，否则取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXClientDecodeDurationByTimeResResultDurationDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 平均耗时，单位毫秒
	Value float64 `json:"Value"`
}

type DescribeImageXClientDecodeSuccessRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientDecodeSuccessRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientDecodeSuccessRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientDecodeSuccessRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientDecodeSuccessRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientDecodeSuccessRateByTimeResResult `json:"Result"`
}

type DescribeImageXClientDecodeSuccessRateByTimeResResponseMetadata struct {

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

type DescribeImageXClientDecodeSuccessRateByTimeResResult struct {

	// REQUIRED; 解码成功率数据
	SuccessRateData []*DescribeImageXClientDecodeSuccessRateByTimeResResultSuccessRateDataItem `json:"SuccessRateData"`
}

type DescribeImageXClientDecodeSuccessRateByTimeResResultSuccessRateDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的解码成功率数据列表。
	Data []*DescribeImageXClientDecodeSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXClientDecodeSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 解码成功率
	Value float64 `json:"Value"`
}

type DescribeImageXClientDemotionRateByTimeBody struct {

	// REQUIRED; 降级类型。取值如下所示： heic：查询 heic 降级率 heif：查询 heif 降级率 avif：查询 avif 降级率
	DemotionType string `json:"DemotionType"`

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientDemotionRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientDemotionRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientDemotionRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientDemotionRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientDemotionRateByTimeResResult `json:"Result"`
}

type DescribeImageXClientDemotionRateByTimeResResponseMetadata struct {

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

type DescribeImageXClientDemotionRateByTimeResResult struct {

	// REQUIRED; 降级率数据
	DemotionRateData []*DescribeImageXClientDemotionRateByTimeResResultDemotionRateDataItem `json:"DemotionRateData"`
}

type DescribeImageXClientDemotionRateByTimeResResultDemotionRateDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 具体数据
	Data []*DescribeImageXClientDemotionRateByTimeResResultDemotionRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXClientDemotionRateByTimeResResultDemotionRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 降级率
	Value float64 `json:"Value"`
}

type DescribeImageXClientErrorCodeAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 聚合维度。 Domain：域名； ErrorCode：错误码； Region：地区； Isp：运营商。
	GroupBy string `json:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientErrorCodeAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc bool `json:"OrderAsc,omitempty"`

	// 目前仅支持传入Count按错误码数量排序。不传或者传空默认按错误码数量排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientErrorCodeAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientErrorCodeAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientErrorCodeAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientErrorCodeAllResResult `json:"Result"`
}

type DescribeImageXClientErrorCodeAllResResponseMetadata struct {

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

type DescribeImageXClientErrorCodeAllResResult struct {

	// REQUIRED; 错误码数据
	ErrorCodeData []*DescribeImageXClientErrorCodeAllResResultErrorCodeDataItem `json:"ErrorCodeData"`
}

type DescribeImageXClientErrorCodeAllResResultErrorCodeDataItem struct {

	// REQUIRED; 错误码数量
	Value int `json:"Value"`

	// 当GroupBy取值非ErrorCode时出现，表示错误码详细信息。
	Details []*DescribeImageXClientErrorCodeAllResResultErrorCodeDataPropertiesItemsItem `json:"Details,omitempty"`

	// 当GroupBy取值Domain时出现，则表示域名信息。
	Domain string `json:"Domain,omitempty"`

	// 当GroupBy取值ErrorCode时出现，表示错误码内容。
	ErrorCode string `json:"ErrorCode,omitempty"`

	// 当GroupBy取值ErrorCode时出现，表示错误码的描述信息。
	ErrorCodeDesc string `json:"ErrorCodeDesc,omitempty"`

	// 当GroupBy取值Isp时出现，则表示运营商信息。
	Isp string `json:"Isp,omitempty"`

	// 当GroupBy取值Region时出现，表示地区信息。
	Region string `json:"Region,omitempty"`

	// 当GroupBy取值Region时出现，表示地区类型。 取值Country，表示国家。 取值Province，表示省份。
	RegionType string `json:"RegionType,omitempty"`
}

type DescribeImageXClientErrorCodeAllResResultErrorCodeDataPropertiesItemsItem struct {

	// REQUIRED; 错误码内容
	ErrorCode string `json:"ErrorCode"`

	// REQUIRED; 错误码的描述信息
	ErrorCodeDesc string `json:"ErrorCodeDesc"`

	// REQUIRED; 错误码数量
	Value int `json:"Value"`
}

type DescribeImageXClientErrorCodeByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientErrorCodeByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientErrorCodeByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientErrorCodeByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientErrorCodeByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientErrorCodeByTimeResResult `json:"Result"`
}

type DescribeImageXClientErrorCodeByTimeResResponseMetadata struct {

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

type DescribeImageXClientErrorCodeByTimeResResult struct {

	// REQUIRED; 所有错误码数据
	ErrorCodeData []*DescribeImageXClientErrorCodeByTimeResResultErrorCodeDataItem `json:"ErrorCodeData"`
}

type DescribeImageXClientErrorCodeByTimeResResultErrorCodeDataItem struct {

	// REQUIRED; 错误码数量
	Count int `json:"Count"`

	// REQUIRED; 错误码对应的时序数据。
	Data []*DescribeImageXClientErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 错误码
	ErrorCode string `json:"ErrorCode"`
}

type DescribeImageXClientErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem struct {

	// REQUIRED; 错误码数量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 错误码数量
	Value int `json:"Value"`
}

type DescribeImageXClientFailureRateBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientFailureRateBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientFailureRateBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientFailureRateRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientFailureRateResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientFailureRateResResult `json:"Result"`
}

type DescribeImageXClientFailureRateResResponseMetadata struct {

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

type DescribeImageXClientFailureRateResResult struct {

	// REQUIRED; 用户感知失败率数据
	FailureRateData []*DescribeImageXClientFailureRateResResultFailureRateDataItem `json:"FailureRateData"`
}

type DescribeImageXClientFailureRateResResultFailureRateDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的用户感知失败率数据列表。
	Data []*DescribeImageXClientFailureRateResResultFailureRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXClientFailureRateResResultFailureRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 用户感知失败率
	Value float64 `json:"Value"`
}

type DescribeImageXClientFileSizeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientFileSizeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientFileSizeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientFileSizeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientFileSizeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientFileSizeResResult `json:"Result"`
}

type DescribeImageXClientFileSizeResResponseMetadata struct {

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

type DescribeImageXClientFileSizeResResult struct {

	// REQUIRED; 客户端文件大小分布数据
	FSizeData []*DescribeImageXClientFileSizeResResultFSizeDataItem `json:"FSizeData"`
}

type DescribeImageXClientFileSizeResResultFSizeDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的文件大小数据列表。
	Data []*DescribeImageXClientFileSizeResResultFSizeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空或Duration时，取值avg/pct25/pct50/pct90/pct99/min/max，否则取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXClientFileSizeResResultFSizeDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 文件大小，单位byte
	Value float64 `json:"Value"`
}

type DescribeImageXClientLoadDurationAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 聚合维度。 Domain：域名 Region：地区 Isp：运营商
	GroupBy string `json:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientLoadDurationAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc bool `json:"OrderAsc,omitempty"`

	// 目前仅支持传入Count按错误码数量排序。不传或者传空默认按错误码数量排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientLoadDurationAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientLoadDurationAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientLoadDurationAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientLoadDurationAllResResult `json:"Result"`
}

type DescribeImageXClientLoadDurationAllResResponseMetadata struct {

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

type DescribeImageXClientLoadDurationAllResResult struct {

	// REQUIRED; 网络耗时数据
	DurationData []*DescribeImageXClientLoadDurationAllResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXClientLoadDurationAllResResultDurationDataItem struct {

	// REQUIRED; 上报数据量
	Count int `json:"Count"`

	// REQUIRED; 加载耗时，单位为毫秒
	Value float64 `json:"Value"`

	// 当GroupBy取值Domain时出现，则表示域名信息。
	Domain string `json:"Domain,omitempty"`

	// 当GroupBy取值Isp时出现，则表示运营商信息。
	Isp string `json:"Isp,omitempty"`

	// 当GroupBy取值Region时出现，表示地区信息。
	Region string `json:"Region,omitempty"`

	// 当GroupBy取值Region时出现，表示地区类型。 Country：国家。 Province：省份。
	RegionType string `json:"RegionType,omitempty"`
}

type DescribeImageXClientLoadDurationBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientLoadDurationBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientLoadDurationBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientLoadDurationRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientLoadDurationResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientLoadDurationResResult `json:"Result"`
}

type DescribeImageXClientLoadDurationResResponseMetadata struct {

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

type DescribeImageXClientLoadDurationResResult struct {

	// REQUIRED; 加载耗时数据
	DurationData []*DescribeImageXClientLoadDurationResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXClientLoadDurationResResultDurationDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的加载耗时数据列表
	Data []*DescribeImageXClientLoadDurationResResultDurationDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。当GroupBy为空或Duration时，取值avg/pct25/pct50/pct75/pct90/pct95/pct99/min/max，否则取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXClientLoadDurationResResultDurationDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 平均耗时，单位毫秒
	Value float64 `json:"Value"`
}

type DescribeImageXClientQualityRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 质量类型。取值如下所示： whitesuspected：查询白屏率 blacksuspected：查询黑屏率 transparent_suspected：查询透明图率
	QualityType string `json:"QualityType"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientQualityRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientQualityRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientQualityRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientQualityRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientQualityRateByTimeResResult `json:"Result"`
}

type DescribeImageXClientQualityRateByTimeResResponseMetadata struct {

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

type DescribeImageXClientQualityRateByTimeResResult struct {

	// REQUIRED; 黑白屏率数据
	QualityRateData []*DescribeImageXClientQualityRateByTimeResResultQualityRateDataItem `json:"QualityRateData"`
}

type DescribeImageXClientQualityRateByTimeResResultQualityRateDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 具体数据
	Data []*DescribeImageXClientQualityRateByTimeResResultQualityRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXClientQualityRateByTimeResResultQualityRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 黑白屏率
	Value float64 `json:"Value"`
}

type DescribeImageXClientQueueDurationByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientQueueDurationByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientQueueDurationByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientQueueDurationByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientQueueDurationByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientQueueDurationByTimeResResult `json:"Result"`
}

type DescribeImageXClientQueueDurationByTimeResResponseMetadata struct {

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

type DescribeImageXClientQueueDurationByTimeResResult struct {

	// REQUIRED; 排队耗时数据
	DurationData []*DescribeImageXClientQueueDurationByTimeResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXClientQueueDurationByTimeResResultDurationDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的排队耗时数据列表
	Data []*DescribeImageXClientQueueDurationByTimeResResultDurationDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空或Duration时，取值avg/pct25/pct50/pct90/pct99/min/max，否则取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXClientQueueDurationByTimeResResultDurationDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 平均耗时，单位毫秒
	Value float64 `json:"Value"`
}

type DescribeImageXClientScoreByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 打分类型。取值如下所示： vq：查询vqscore指标 aes：查询美学指标 noi：查询噪声指标 psnr：查询psnr指标 ssim：查询ssim指标 vmaf：查询vmaf指标
	ScoreType string `json:"ScoreType"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientScoreByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientScoreByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientScoreByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientScoreByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientScoreByTimeResResult `json:"Result"`
}

type DescribeImageXClientScoreByTimeResResponseMetadata struct {

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

type DescribeImageXClientScoreByTimeResResult struct {

	// REQUIRED; 画质打分数据
	ScoreData []*DescribeImageXClientScoreByTimeResResultScoreDataItem `json:"ScoreData"`
}

type DescribeImageXClientScoreByTimeResResultScoreDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 具体数据
	Data []*DescribeImageXClientScoreByTimeResResultScoreDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXClientScoreByTimeResResultScoreDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 画质打分
	Value float64 `json:"Value"`
}

type DescribeImageXClientSdkVerByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的域名，不传则匹配所有域名
	Domain []string `json:"Domain,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXClientSdkVerByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXClientSdkVerByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXClientSdkVerByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientSdkVerByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientSdkVerByTimeResResult `json:"Result"`
}

type DescribeImageXClientSdkVerByTimeResResponseMetadata struct {

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

type DescribeImageXClientSdkVerByTimeResResult struct {

	// REQUIRED; SDK版本数据
	SdkVerData []*DescribeImageXClientSdkVerByTimeResResultSdkVerDataItem `json:"SdkVerData"`
}

type DescribeImageXClientSdkVerByTimeResResultSdkVerDataItem struct {

	// REQUIRED; 该 SDK 版本对应的总上报量。
	Count int `json:"Count"`

	// REQUIRED; 对应的版本数据列表。
	Data []*DescribeImageXClientSdkVerByTimeResResultSdkVerDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; SDK版本号
	SdkVer string `json:"SdkVer"`
}

type DescribeImageXClientSdkVerByTimeResResultSdkVerDataPropertiesItemsItem struct {

	// REQUIRED; 版本数量。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 版本数量。
	Value int `json:"Value"`
}

type DescribeImageXClientTopDemotionURLBody struct {

	// REQUIRED; 降级类型。取值如下所示： heic：HEIC 降级 heif：HEIF 降级 avif：AVIF 降级
	DemotionType string `json:"DemotionType"`

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 应用 ID。默认为空，不传则匹配账号下的所有的 App ID。
	Appid string `json:"Appid,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 查询 Top URL 条数，取值范围为 [0,1000]，默认 1000。
	Top int `json:"Top,omitempty"`
}

type DescribeImageXClientTopDemotionURLRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientTopDemotionURLResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientTopDemotionURLResResult `json:"Result"`
}

type DescribeImageXClientTopDemotionURLResResponseMetadata struct {

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

type DescribeImageXClientTopDemotionURLResResult struct {

	// REQUIRED; Top URL数据
	TopURLData []*DescribeImageXClientTopDemotionURLResResultTopURLDataItem `json:"TopUrlData"`
}

type DescribeImageXClientTopDemotionURLResResultTopURLDataItem struct {

	// REQUIRED; 上报数据量
	Count int `json:"Count"`

	// REQUIRED; 文件URL
	URL string `json:"Url"`
}

type DescribeImageXClientTopFileSizeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 应用 ID。默认为空，不传则匹配账号下的所有的 App ID。
	Appid string `json:"Appid,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 查询 Top URL 条数，取值范围为[0,1000]，默认值为 1000。
	Top int `json:"Top,omitempty"`
}

type DescribeImageXClientTopFileSizeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientTopFileSizeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientTopFileSizeResResult `json:"Result"`
}

type DescribeImageXClientTopFileSizeResResponseMetadata struct {

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

type DescribeImageXClientTopFileSizeResResult struct {

	// REQUIRED; Top URL数据
	TopURLData []*DescribeImageXClientTopFileSizeResResultTopURLDataItem `json:"TopUrlData"`
}

type DescribeImageXClientTopFileSizeResResultTopURLDataItem struct {

	// REQUIRED; 上报数据量
	Count int `json:"Count"`

	// REQUIRED; 文件URL
	URL string `json:"Url"`

	// REQUIRED; 文件大小，单位byte
	Value float64 `json:"Value"`
}

type DescribeImageXClientTopQualityURLBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 类型。取值如下所示： transparentsuspected：透明图 whitesuspected：白屏 black_suspected：黑屏
	QualityType string `json:"QualityType"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 应用 ID。默认为空，不传则匹配账号下的所有的 App ID。
	Appid string `json:"Appid,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp,Domain,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 查询 Top URL 条数，取值范围为 [0,1000]，默认 1000。
	Top int `json:"Top,omitempty"`
}

type DescribeImageXClientTopQualityURLRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXClientTopQualityURLResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXClientTopQualityURLResResult `json:"Result"`
}

type DescribeImageXClientTopQualityURLResResponseMetadata struct {

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

type DescribeImageXClientTopQualityURLResResult struct {

	// REQUIRED; Top URL数据
	TopURLData []*DescribeImageXClientTopQualityURLResResultTopURLDataItem `json:"TopUrlData"`
}

type DescribeImageXClientTopQualityURLResResultTopURLDataItem struct {

	// REQUIRED; 上报数据量
	Count int `json:"Count"`

	// REQUIRED; 文件URL
	URL string `json:"Url"`
}

type DescribeImageXCompressUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 需要分组查询的参数，当前仅支持取值 ServiceId，表示按服务 ID 进行分组。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用
	// GetAllImageServices [https://www.volcengine.com/docs/508/9360] 接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXCompressUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCompressUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCompressUsageResResult `json:"Result"`
}

type DescribeImageXCompressUsageResResponseMetadata struct {

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

type DescribeImageXCompressUsageResResult struct {

	// REQUIRED; 压缩量数据
	CompressData []*DescribeImageXCompressUsageResResultCompressDataItem `json:"CompressData"`
}

type DescribeImageXCompressUsageResResultCompressDataItem struct {

	// REQUIRED; 压缩前大小，单位为 byte。
	InSize []*DescribeImageXCompressUsageResResultCompressDataPropertiesItemsItem `json:"InSize"`

	// REQUIRED; 压缩后大小，单位为 byte。
	OutSize []*Components1Xh7Lz4SchemasDescribeimagexcompressusageresPropertiesResultPropertiesCompressdataItemsPropertiesOutsizeItems `json:"OutSize"`

	// 当GroupBy中包含ServiceId时出现
	ServiceID string `json:"ServiceId,omitempty"`
}

type DescribeImageXCompressUsageResResultCompressDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 压缩前大小，单位Byte
	Value float64 `json:"Value"`
}

type DescribeImageXCubeUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 查询数据的时间粒度。单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXCubeUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXCubeUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXCubeUsageResResult `json:"Result"`
}

type DescribeImageXCubeUsageResResponseMetadata struct {

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

type DescribeImageXCubeUsageResResult struct {

	// REQUIRED; 创意魔方数据。
	CubeData []*DescribeImageXCubeUsageResResultCubeDataItem `json:"CubeData"`
}

type DescribeImageXCubeUsageResResultCubeDataItem struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXCubeUsageResResultCubeDataPropertiesItemsItem `json:"Data"`
}

type DescribeImageXCubeUsageResResultCubeDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 创意魔方请求次数
	Value int `json:"Value"`
}

type DescribeImageXDomainBandwidthDataQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 计费区域。支持查询多个区域，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有区域。取值如下所示：
	// * CHN：中国内地
	// * AP1：亚太一区
	// * AP2：亚太二区
	// * AP3：亚太三区
	// * EU：欧洲
	// * ME：中东和非洲
	// * SA：南美
	// * NA：北美
	// * OTHER：其他
	BillingRegion string `json:"BillingRegion,omitempty" query:"BillingRegion"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 需要分组查询的参数。不传表示不拆分维度，传入多个用英文逗号分隔。取值如下所示：
	// * ServiceId：服务 ID
	// * DomainName：域名
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒。缺省时查询 StartTime 和 EndTime 时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval int `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXDomainBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXDomainBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXDomainBandwidthDataResResult `json:"Result"`
}

type DescribeImageXDomainBandwidthDataResResponseMetadata struct {

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

type DescribeImageXDomainBandwidthDataResResult struct {

	// REQUIRED; 计量数据列表
	BpsData []*DescribeImageXDomainBandwidthDataResResultBpsDataItem `json:"BpsData"`
}

type DescribeImageXDomainBandwidthDataResResultBpsDataItem struct {

	// REQUIRED; 具体数据。
	Data []*DescribeImageXDomainBandwidthDataResResultBpsDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy中包含DomainName时出现
	DomainName string `json:"DomainName,omitempty"`

	// 当GroupBy中包含ServiceId时出现
	ServiceID string `json:"ServiceId,omitempty"`
}

type DescribeImageXDomainBandwidthDataResResultBpsDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 带宽单位：bps
	Value float64 `json:"Value"`
}

type DescribeImageXDomainBandwidthNinetyFiveDataQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 计费区域。支持查询多个区域，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有区域。取值如下所示：
	// * CHN：中国内地
	// * AP1：亚太一区
	// * AP2：亚太二区
	// * AP3：亚太三区
	// * EU：欧洲
	// * ME：中东和非洲
	// * SA：南美
	// * NA：北美
	// * OTHER：其他
	BillingRegion string `json:"BillingRegion,omitempty" query:"BillingRegion"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXDomainBandwidthNinetyFiveDataRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXDomainBandwidthNinetyFiveDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXDomainBandwidthNinetyFiveDataResResult `json:"Result"`
}

type DescribeImageXDomainBandwidthNinetyFiveDataResResponseMetadata struct {

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

type DescribeImageXDomainBandwidthNinetyFiveDataResResult struct {

	// REQUIRED; 带宽95值，单位bps
	BpsData float64 `json:"BpsData"`

	// REQUIRED; 统计时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeImageXDomainTrafficDataQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 计费区域。支持查询多个区域，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有区域。取值如下所示：
	// * CHN：中国内地
	// * AP1：亚太一区
	// * AP2：亚太二区
	// * AP3：亚太三区
	// * EU：欧洲
	// * ME：中东和非洲
	// * SA：南美
	// * NA：北美
	// * OTHER：其他
	BillingRegion string `json:"BillingRegion,omitempty" query:"BillingRegion"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 需要分组查询的参数。不传表示不拆分维度，传入多个用英文逗号分隔。取值如下所示：
	// * ServiceId：服务 ID
	// * DomainName：域名
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒。缺省时查询 StartTime 和 EndTime 时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省时表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务 ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXDomainTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXDomainTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXDomainTrafficDataResResult `json:"Result"`
}

type DescribeImageXDomainTrafficDataResResponseMetadata struct {

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

type DescribeImageXDomainTrafficDataResResult struct {

	// REQUIRED; 计量数据列表
	TrafficData []*DescribeImageXDomainTrafficDataResResultTrafficDataItem `json:"TrafficData"`
}

type DescribeImageXDomainTrafficDataResResultTrafficDataItem struct {

	// REQUIRED; 具体数据。
	Data []*DescribeImageXDomainTrafficDataResResultTrafficDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy中包含DomainName时出现
	DomainName string `json:"DomainName,omitempty"`

	// 当GroupBy中包含ServiceId时出现
	ServiceID string `json:"ServiceId,omitempty"`
}

type DescribeImageXDomainTrafficDataResResultTrafficDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 流量单位：Byte。
	Value float64 `json:"Value"`
}

type DescribeImageXEdgeRequestBandwidthQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 分组依据，取值仅支持DomainName。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 过滤运营商。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * 电信
	// * 联通
	// * 移动
	// * 鹏博士
	// * 教育网
	// * 广电网
	// * 其它
	Isp string `json:"Isp,omitempty" query:"Isp"`

	// 过滤网络协议。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * HTTP
	// * HTTPS
	Protocols string `json:"Protocols,omitempty" query:"Protocols"`

	// 计费区域。支持查询多个区域，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有区域。取值如下所示：
	// * 中国大陆
	// * 亚太一区
	// * 亚太二区
	// * 亚太三区
	// * 欧洲区
	// * 北美区
	// * 南美区
	// * 中东区
	Regions string `json:"Regions,omitempty" query:"Regions"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 客户端国家。传入多个时用英文逗号作为分割符，缺省情况下表示不过滤。可调用 DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574] 获取
	// IP 解析后的客户端国家。取值如下所示：
	// * 海外，除中国外全部国家。
	// * 具体国家取值，如中国、美国。
	UserCountry string `json:"UserCountry,omitempty" query:"UserCountry"`

	// 客户端省份。传入多个用英文逗号分割，缺省情况下表示不过滤。可调用 DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574] 获取 IP 解析后的客户端省份。
	UserProvince string `json:"UserProvince,omitempty" query:"UserProvince"`
}

type DescribeImageXEdgeRequestBandwidthRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXEdgeRequestBandwidthResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXEdgeRequestBandwidthResResult `json:"Result"`
}

type DescribeImageXEdgeRequestBandwidthResResponseMetadata struct {

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

type DescribeImageXEdgeRequestBandwidthResResult struct {

	// REQUIRED; 带宽数据。
	BpsData []*DescribeImageXEdgeRequestBandwidthResResultBpsDataItem `json:"BpsData"`
}

type DescribeImageXEdgeRequestBandwidthResResultBpsDataItem struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXEdgeRequestBandwidthResResultBpsDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy带有DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXEdgeRequestBandwidthResResultBpsDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 带宽用量，单位为 bps。
	Value float64 `json:"Value"`
}

type DescribeImageXEdgeRequestQuery struct {

	// REQUIRED; 状态码，传入多个时用英文逗号分隔。取值如下所示：
	// * req_cnt：返回所有状态码数据
	// * 2xx：返回 2xx 状态码汇总数据
	// * 3xx：返回 3xx 状态码汇总数据
	// * 4xx：返回 4xx 状态码汇总数据
	// * 5xx：返回 5xx 状态码汇总数据。
	DataTypes string `json:"DataTypes" query:"DataTypes"`

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 是否拆分状态码，取值如下所示：
	// * true：拆分
	// * false：（默认）不拆分
	DetailedCode bool `json:"DetailedCode,omitempty" query:"DetailedCode"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 需要分组查询的参数，传入多个用英文逗号分割。取值如下所示：
	// * DomainName：域名
	// * DataType：状态码
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 过滤运营商。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * 电信
	// * 联通
	// * 移动
	// * 鹏博士
	// * 教育网
	// * 广电网
	// * 其它
	Isp string `json:"Isp,omitempty" query:"Isp"`

	// 过滤网络协议。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * HTTP
	// * HTTPS
	Protocols string `json:"Protocols,omitempty" query:"Protocols"`

	// 计费区域。支持查询多个区域，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有区域。取值如下所示：
	// * 中国大陆
	// * 亚太一区
	// * 亚太二区
	// * 亚太三区
	// * 欧洲区
	// * 北美区
	// * 南美区
	// * 中东区
	Regions string `json:"Regions,omitempty" query:"Regions"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 客户端国家。支持查询多个国家，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有国家。您可通过调用 DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]
	// 获取 IP 解析后的客户端国家。取值如下所示：
	// * 海外，除中国外全部国家。
	// * 具体国家取值，如中国、美国。
	UserCountry string `json:"UserCountry,omitempty" query:"UserCountry"`

	// 客户端省份。支持查询多个省份，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有省份。可调用 DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]
	// 获取 IP 解析后的客户端省份。
	UserProvince string `json:"UserProvince,omitempty" query:"UserProvince"`
}

type DescribeImageXEdgeRequestRegionsQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。
	// 起始时间与结束时间间隔小于等于 90 天。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。
	// 起始时间与结束时间间隔小于等于 90 天。
	StartTime string `json:"StartTime" query:"StartTime"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXEdgeRequestRegionsRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXEdgeRequestRegionsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXEdgeRequestRegionsResResult `json:"Result"`
}

type DescribeImageXEdgeRequestRegionsResResponseMetadata struct {

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

type DescribeImageXEdgeRequestRegionsResResult struct {

	// REQUIRED; 客户端国家
	UserCountry []string `json:"UserCountry"`

	// REQUIRED; 客户端省份
	UserProvince []string `json:"UserProvince"`
}

type DescribeImageXEdgeRequestRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXEdgeRequestResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXEdgeRequestResResult `json:"Result"`
}

type DescribeImageXEdgeRequestResResponseMetadata struct {

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

type DescribeImageXEdgeRequestResResult struct {

	// REQUIRED; 数据列表
	RequestCnt []*DescribeImageXEdgeRequestResResultRequestCntItem `json:"RequestCnt"`
}

type DescribeImageXEdgeRequestResResultRequestCntItem struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXEdgeRequestResResultRequestCntPropertiesItemsItem `json:"Data"`

	// 当GroupBy带有DataType时出现
	DataType string `json:"DataType,omitempty"`

	// 当GroupBy带有DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXEdgeRequestResResultRequestCntPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求次数，单位为次。
	Value int `json:"Value"`
}

type DescribeImageXEdgeRequestTrafficQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 分组依据，取值仅支持DomainName。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 过滤运营商。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * 电信
	// * 联通
	// * 移动
	// * 鹏博士
	// * 教育网
	// * 广电网
	// * 其它
	Isp string `json:"Isp,omitempty" query:"Isp"`

	// 过滤网络协议。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * HTTP
	// * HTTPS
	Protocols string `json:"Protocols,omitempty" query:"Protocols"`

	// 区域。支持查询多个区域，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有区域。取值如下所示：
	// * 中国大陆
	// * 亚太一区
	// * 亚太二区
	// * 亚太三区
	// * 欧洲区
	// * 北美区
	// * 南美区
	// * 中东区
	Regions string `json:"Regions,omitempty" query:"Regions"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 客户端国家。传入多个时用英文逗号作为分割符，缺省情况下表示不过滤。您可以通过调用 DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]
	// 获取 IP 解析后的客户端国家。取值如下所示：
	// * 海外，除中国外全部国家。
	// * 具体国家取值，如中国、美国。
	UserCountry string `json:"UserCountry,omitempty" query:"UserCountry"`

	// 客户端省份。传入多个用英文逗号分割，缺省情况下表示不过滤。可调用 DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574] 获取 IP 解析后的客户端省份。
	UserProvince string `json:"UserProvince,omitempty" query:"UserProvince"`
}

type DescribeImageXEdgeRequestTrafficRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXEdgeRequestTrafficResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXEdgeRequestTrafficResResult `json:"Result"`
}

type DescribeImageXEdgeRequestTrafficResResponseMetadata struct {

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

type DescribeImageXEdgeRequestTrafficResResult struct {

	// REQUIRED; 查询数据
	TrafficData []*DescribeImageXEdgeRequestTrafficResResultTrafficDataItem `json:"TrafficData"`
}

type DescribeImageXEdgeRequestTrafficResResultTrafficDataItem struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXEdgeRequestTrafficResResultTrafficDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy带有DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXEdgeRequestTrafficResResultTrafficDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 流量，单位为 byte。
	Value float64 `json:"Value"`
}

type DescribeImageXExceedCountByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXExceedCountByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXExceedCountByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXExceedCountByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXExceedCountByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXExceedCountByTimeResResult `json:"Result"`
}

type DescribeImageXExceedCountByTimeResResponseMetadata struct {

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

type DescribeImageXExceedCountByTimeResResult struct {

	// REQUIRED; 上报量数据。
	ExceedCountData []*DescribeImageXExceedCountByTimeResResultExceedCountDataItem `json:"ExceedCountData"`
}

type DescribeImageXExceedCountByTimeResResultExceedCountDataItem struct {

	// REQUIRED; 该数据类型对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的客户端上报量数据列表。
	Data []*DescribeImageXExceedCountByTimeResResultExceedCountDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXExceedCountByTimeResResultExceedCountDataPropertiesItemsItem struct {

	// REQUIRED; 上报量数据
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 上报量数据
	Value int `json:"Value"`
}

type DescribeImageXExceedFileSizeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXExceedFileSizeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,Country,Province,Isp,Domain），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXExceedFileSizeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXExceedFileSizeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXExceedFileSizeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXExceedFileSizeResResult `json:"Result"`
}

type DescribeImageXExceedFileSizeResResponseMetadata struct {

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

type DescribeImageXExceedFileSizeResResult struct {

	// REQUIRED; 文件大小分布数据
	FSizeData []*DescribeImageXExceedFileSizeResResultFSizeDataItem `json:"FSizeData"`
}

type DescribeImageXExceedFileSizeResResultFSizeDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的文件大小数据列表。
	Data []*DescribeImageXExceedFileSizeResResultFSizeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空或Duration时，取值avg/pct25/pct50/pct90/pct99/min/max，否则取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXExceedFileSizeResResultFSizeDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 文件大小，单位byte
	Value float64 `json:"Value"`
}

type DescribeImageXExceedResolutionRatioAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXExceedResolutionRatioAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc string `json:"OrderAsc,omitempty"`

	// * 不传或者传空或者取值为Count时，表示按上报量排序。
	// * 取值为WidthRatio时，表示按宽比排序。
	// * 取值为HeightRatio时，表示按高比排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXExceedResolutionRatioAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXExceedResolutionRatioAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXExceedResolutionRatioAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXExceedResolutionRatioAllResResult `json:"Result"`
}

type DescribeImageXExceedResolutionRatioAllResResponseMetadata struct {

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

type DescribeImageXExceedResolutionRatioAllResResult struct {

	// REQUIRED; 文件大小分布数据
	ResolutionRatioData []*DescribeImageXExceedResolutionRatioAllResResultResolutionRatioDataItem `json:"ResolutionRatioData"`
}

type DescribeImageXExceedResolutionRatioAllResResultResolutionRatioDataItem struct {

	// REQUIRED; 大图样本量
	Count int `json:"Count"`

	// REQUIRED; 高比，即为图片高/view高取整
	HeightRatio int `json:"HeightRatio"`

	// REQUIRED; 格式为：宽比-高比
	Ratio string `json:"Ratio"`

	// REQUIRED; 宽比，即为图片宽/view宽取整
	WidthRatio int `json:"WidthRatio"`
}

type DescribeImageXHeifEncodeDurationByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXHeifEncodeDurationByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片分辨率，不传则匹配所有图片分辨率。
	ImageResolution []string `json:"ImageResolution,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配所有系统。取值如下所示： iOS Android
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXHeifEncodeDurationByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXHeifEncodeDurationByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHeifEncodeDurationByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHeifEncodeDurationByTimeResResult `json:"Result"`
}

type DescribeImageXHeifEncodeDurationByTimeResResponseMetadata struct {

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

type DescribeImageXHeifEncodeDurationByTimeResResult struct {

	// REQUIRED; 编码耗时数据
	DurationData []*DescribeImageXHeifEncodeDurationByTimeResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXHeifEncodeDurationByTimeResResultDurationDataItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的编码耗时数据列表。
	Data []*DescribeImageXHeifEncodeDurationByTimeResResultDurationDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空时，取值为：Total。 当GroupBy为Duration时，取值为：min、max、pct25、pct50、pct90、pct99、avg。 除上述外取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXHeifEncodeDurationByTimeResResultDurationDataPropertiesItemsItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 平均耗时，单位毫秒
	Value float64 `json:"Value"`
}

type DescribeImageXHeifEncodeErrorCodeByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXHeifEncodeErrorCodeByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片分辨率，不传则匹配所有图片分辨率。
	ImageResolution []string `json:"ImageResolution,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配所有系统。取值如下所示： iOS Android
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXHeifEncodeErrorCodeByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXHeifEncodeErrorCodeByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHeifEncodeErrorCodeByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHeifEncodeErrorCodeByTimeResResult `json:"Result"`
}

type DescribeImageXHeifEncodeErrorCodeByTimeResResponseMetadata struct {

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

type DescribeImageXHeifEncodeErrorCodeByTimeResResult struct {

	// REQUIRED; 编码错误码数据。
	ErrorCodeData []*DescribeImageXHeifEncodeErrorCodeByTimeResResultErrorCodeDataItem `json:"ErrorCodeData"`
}

type DescribeImageXHeifEncodeErrorCodeByTimeResResultErrorCodeDataItem struct {

	// REQUIRED; 错误码总量。
	Count int `json:"Count"`

	// REQUIRED; 错误码对应的时序数据。
	Data []*DescribeImageXHeifEncodeErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 错误码。
	ErrorCode string `json:"ErrorCode"`
}

type DescribeImageXHeifEncodeErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem struct {

	// REQUIRED; 错误码数量。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 错误码数量。
	Value int `json:"Value"`
}

type DescribeImageXHeifEncodeFileInSizeByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXHeifEncodeFileInSizeByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片分辨率，不传则匹配所有图片分辨率。
	ImageResolution []string `json:"ImageResolution,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配所有系统。取值如下所示： iOS Android
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXHeifEncodeFileInSizeByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXHeifEncodeFileInSizeByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHeifEncodeFileInSizeByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHeifEncodeFileInSizeByTimeResResult `json:"Result"`
}

type DescribeImageXHeifEncodeFileInSizeByTimeResResponseMetadata struct {

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

type DescribeImageXHeifEncodeFileInSizeByTimeResResult struct {

	// REQUIRED; 文件大小数据
	FileSizeData []*DescribeImageXHeifEncodeFileInSizeByTimeResResultFileSizeDataItem `json:"FileSizeData"`
}

type DescribeImageXHeifEncodeFileInSizeByTimeResResultFileSizeDataItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的文件大小数据列表。
	Data []*DescribeImageXHeifEncodeFileInSizeByTimeResResultFileSizeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空时，取值为：Total。 当GroupBy为Duration时，取值为：min、max、pct25、pct50、pct90、pct99、avg。 除上述外取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXHeifEncodeFileInSizeByTimeResResultFileSizeDataPropertiesItemsItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 文件大小，单位byte
	Value float64 `json:"Value"`
}

type DescribeImageXHeifEncodeFileOutSizeByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXHeifEncodeFileOutSizeByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：Duration（拆分分位数据）、公共维度（Appid,OS,AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片分辨率，不传则匹配所有图片分辨率。
	ImageResolution []string `json:"ImageResolution,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配所有系统。取值如下所示： iOS Android
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXHeifEncodeFileOutSizeByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXHeifEncodeFileOutSizeByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHeifEncodeFileOutSizeByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHeifEncodeFileOutSizeByTimeResResult `json:"Result"`
}

type DescribeImageXHeifEncodeFileOutSizeByTimeResResponseMetadata struct {

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

type DescribeImageXHeifEncodeFileOutSizeByTimeResResult struct {

	// REQUIRED; 文件大小数据
	FileSizeData []*DescribeImageXHeifEncodeFileOutSizeByTimeResResultFileSizeDataItem `json:"FileSizeData"`
}

type DescribeImageXHeifEncodeFileOutSizeByTimeResResultFileSizeDataItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 对应的文件大小数据列表。
	Data []*DescribeImageXHeifEncodeFileOutSizeByTimeResResultFileSizeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空时，取值为：Total。 当GroupBy为Duration时，取值为：avg、min、max、pct25、pct50、pct90、pct99。 除上述外取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXHeifEncodeFileOutSizeByTimeResResultFileSizeDataPropertiesItemsItem struct {

	// REQUIRED; 数据上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 文件大小，单位byte
	Value float64 `json:"Value"`
}

type DescribeImageXHeifEncodeSuccessCountByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXHeifEncodeSuccessCountByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片分辨率，不传则匹配所有图片分辨率。
	ImageResolution []string `json:"ImageResolution,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配所有系统。取值如下所示： iOS Android
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXHeifEncodeSuccessCountByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXHeifEncodeSuccessCountByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHeifEncodeSuccessCountByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHeifEncodeSuccessCountByTimeResResult `json:"Result"`
}

type DescribeImageXHeifEncodeSuccessCountByTimeResResponseMetadata struct {

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

type DescribeImageXHeifEncodeSuccessCountByTimeResResult struct {

	// REQUIRED; 编码成功次数数据
	SuccessCountData []*DescribeImageXHeifEncodeSuccessCountByTimeResResultSuccessCountDataItem `json:"SuccessCountData"`
}

type DescribeImageXHeifEncodeSuccessCountByTimeResResultSuccessCountDataItem struct {

	// REQUIRED; 编码成功次数
	Count int `json:"Count"`

	// REQUIRED; 对应的编码成功次数数据列表。
	Data []*DescribeImageXHeifEncodeSuccessCountByTimeResResultSuccessCountDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空时，取值为：Total。 除上述外取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXHeifEncodeSuccessCountByTimeResResultSuccessCountDataPropertiesItemsItem struct {

	// REQUIRED; 编码成功次数
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 编码成功次数
	Value int `json:"Value"`
}

type DescribeImageXHeifEncodeSuccessRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXHeifEncodeSuccessRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType,ImageResolution），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片分辨率，不传则匹配所有图片分辨率。
	ImageResolution []string `json:"ImageResolution,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。
	ImageType []string `json:"ImageType,omitempty"`

	// 需要匹配的系统类型，不传则匹配所有系统。取值如下所示： iOS Android
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXHeifEncodeSuccessRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXHeifEncodeSuccessRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHeifEncodeSuccessRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHeifEncodeSuccessRateByTimeResResult `json:"Result"`
}

type DescribeImageXHeifEncodeSuccessRateByTimeResResponseMetadata struct {

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

type DescribeImageXHeifEncodeSuccessRateByTimeResResult struct {

	// REQUIRED; 编码成功率数据
	SuccessRateData []*DescribeImageXHeifEncodeSuccessRateByTimeResResultSuccessRateDataItem `json:"SuccessRateData"`
}

type DescribeImageXHeifEncodeSuccessRateByTimeResResultSuccessRateDataItem struct {

	// REQUIRED; 数据上报次数。
	Count int `json:"Count"`

	// REQUIRED; 对应的编码成功次数数据列表。
	Data []*DescribeImageXHeifEncodeSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空时，取值为：Total。 除上述外取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXHeifEncodeSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据上报次数。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 编码成功率。
	Value float64 `json:"Value"`
}

type DescribeImageXHitRateRequestDataQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 需要分组查询的参数。取值仅支持DomainName。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度，单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXHitRateRequestDataRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHitRateRequestDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHitRateRequestDataResResult `json:"Result"`
}

type DescribeImageXHitRateRequestDataResResponseMetadata struct {

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

type DescribeImageXHitRateRequestDataResResult struct {

	// REQUIRED; 命中率数据
	HitRateData []*DescribeImageXHitRateRequestDataResResultHitRateDataItem `json:"HitRateData"`
}

type DescribeImageXHitRateRequestDataResResultHitRateDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXHitRateRequestDataResResultHitRateDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy=DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXHitRateRequestDataResResultHitRateDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求命中率
	Value float64 `json:"Value"`
}

type DescribeImageXHitRateTrafficDataQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 域名。支持查询多个域名，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有域名。您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 需要分组查询的参数。取值仅支持DomainName。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度，单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXHitRateTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXHitRateTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXHitRateTrafficDataResResult `json:"Result"`
}

type DescribeImageXHitRateTrafficDataResResponseMetadata struct {

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

type DescribeImageXHitRateTrafficDataResResult struct {

	// REQUIRED; 命中率数据
	HitRateData []*DescribeImageXHitRateTrafficDataResResultHitRateDataItem `json:"HitRateData"`
}

type DescribeImageXHitRateTrafficDataResResultHitRateDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXHitRateTrafficDataResResultHitRateDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy=DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXHitRateTrafficDataResResultHitRateDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 流量命中率。
	Value float64 `json:"Value"`
}

type DescribeImageXMirrorRequestBandwidthBody struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 限制查询的域名，传入多个用英文逗号分割。不传为不限制。
	DomainNames []string `json:"DomainNames,omitempty"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty"`

	// 限制查询的服务id，传入多个用英文逗号分割。不传为不限制。
	ServiceIDs []string `json:"ServiceIds,omitempty"`
}

type DescribeImageXMirrorRequestBandwidthRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXMirrorRequestBandwidthResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXMirrorRequestBandwidthResResult `json:"Result"`
}

type DescribeImageXMirrorRequestBandwidthResResponseMetadata struct {

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

type DescribeImageXMirrorRequestBandwidthResResult struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXMirrorRequestBandwidthResResultDataItem `json:"Data"`
}

type DescribeImageXMirrorRequestBandwidthResResultDataItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 带宽，单位bps
	Value float64 `json:"Value"`
}

type DescribeImageXMirrorRequestHTTPCodeByTimeBody struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	StartTime string `json:"StartTime"`

	// 状态码是否聚合。取值：true/false。默认为false
	AggregateCode string `json:"AggregateCode,omitempty"`

	// 限制查询的域名，传入多个用英文逗号分割。不传为不限制。
	DomainNames []string `json:"DomainNames,omitempty"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty"`

	// 限制查询的服务id，传入多个用英文逗号分割。不传为不限制。
	ServiceIDs []string `json:"ServiceIds,omitempty"`
}

type DescribeImageXMirrorRequestHTTPCodeByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXMirrorRequestHTTPCodeByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXMirrorRequestHTTPCodeByTimeResResult `json:"Result"`
}

type DescribeImageXMirrorRequestHTTPCodeByTimeResResponseMetadata struct {

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

type DescribeImageXMirrorRequestHTTPCodeByTimeResResult struct {

	// REQUIRED; 数据列表
	Data []*DescribeImageXMirrorRequestHTTPCodeByTimeResResultDataItem `json:"Data"`
}

type DescribeImageXMirrorRequestHTTPCodeByTimeResResultDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXMirrorRequestHTTPCodeByTimeResResultDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 状态码
	HTTPCode string `json:"HttpCode"`
}

type DescribeImageXMirrorRequestHTTPCodeByTimeResResultDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求次数
	Value int `json:"Value"`
}

type DescribeImageXMirrorRequestHTTPCodeOverviewBody struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	StartTime string `json:"StartTime"`

	// 状态码是否聚合。取值：true/false。默认为false
	AggregateCode string `json:"AggregateCode,omitempty"`

	// 限制查询的域名，传入多个用英文逗号分割。不传为不限制。
	DomainNames []string `json:"DomainNames,omitempty"`

	// 限制查询的服务id，传入多个用英文逗号分割。不传为不限制。
	ServiceIDs []string `json:"ServiceIds,omitempty"`
}

type DescribeImageXMirrorRequestHTTPCodeOverviewRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXMirrorRequestHTTPCodeOverviewResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXMirrorRequestHTTPCodeOverviewResResult `json:"Result"`
}

type DescribeImageXMirrorRequestHTTPCodeOverviewResResponseMetadata struct {

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

type DescribeImageXMirrorRequestHTTPCodeOverviewResResult struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXMirrorRequestHTTPCodeOverviewResResultDataItem `json:"Data"`
}

type DescribeImageXMirrorRequestHTTPCodeOverviewResResultDataItem struct {

	// REQUIRED; 总请求次数
	Count int `json:"Count"`

	// REQUIRED; 详细数据
	Details []*DescribeImageXMirrorRequestHTTPCodeOverviewResResultDataPropertiesItemsItem `json:"Details"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`
}

type DescribeImageXMirrorRequestHTTPCodeOverviewResResultDataPropertiesItemsItem struct {

	// REQUIRED; Http 状态码
	HTTPCode string `json:"HttpCode"`

	// REQUIRED; http状态码对应的请求次数
	Value int `json:"Value"`
}

type DescribeImageXMirrorRequestTrafficBody struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	StartTime string `json:"StartTime"`

	// 限制查询的域名，传入多个用英文逗号分割。不传为不限制。
	DomainNames []string `json:"DomainNames,omitempty"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty"`

	// 限制查询的服务id，传入多个用英文逗号分割。不传为不限制。
	ServiceIDs []string `json:"ServiceIds,omitempty"`
}

type DescribeImageXMirrorRequestTrafficRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXMirrorRequestTrafficResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXMirrorRequestTrafficResResult `json:"Result"`
}

type DescribeImageXMirrorRequestTrafficResResponseMetadata struct {

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

type DescribeImageXMirrorRequestTrafficResResult struct {

	// REQUIRED; 数据列表
	Data []*DescribeImageXMirrorRequestTrafficResResultDataItem `json:"Data"`
}

type DescribeImageXMirrorRequestTrafficResResultDataItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 流量，单位byte
	Value int `json:"Value"`
}

type DescribeImageXMultiCompressUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXMultiCompressUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXMultiCompressUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXMultiCompressUsageResResult `json:"Result"`
}

type DescribeImageXMultiCompressUsageResResponseMetadata struct {

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

type DescribeImageXMultiCompressUsageResResult struct {

	// REQUIRED; 多文件压缩用量数据
	CompressData []*DescribeImageXMultiCompressUsageResResultCompressDataItem `json:"CompressData"`
}

type DescribeImageXMultiCompressUsageResResultCompressDataItem struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXMultiCompressUsageResResultCompressDataPropertiesItemsItem `json:"Data"`
}

type DescribeImageXMultiCompressUsageResResultCompressDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 压缩量，单位Byte。
	Value float64 `json:"Value"`
}

type DescribeImageXRequestCntUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 组件名称。支持查询多个组件，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有组件。您可通过调用 GetImageAddOnConf [https://www.volcengine.com/docs/508/66145] 查看Key返回值。
	AdvFeats string `json:"AdvFeats,omitempty" query:"AdvFeats"`

	// 维度拆分的维度值。不传表示不拆分维度，传入多个时用英文逗号“,”分割。支持取值如下：
	// * ServiceId：服务
	// * AdvFeat：组件
	// * Model：模型
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒。缺省时查询 StartTime 和 EndTime 时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 图片处理模板。支持查询多个模板，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有模板。您可通过调用 GetAllImageTemplates [https://www.volcengine.com/docs/508/9386] 获取服务下全部模版信息。
	Templates string `json:"Templates,omitempty" query:"Templates"`
}

type DescribeImageXRequestCntUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXRequestCntUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXRequestCntUsageResResult `json:"Result"`
}

type DescribeImageXRequestCntUsageResResponseMetadata struct {

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

type DescribeImageXRequestCntUsageResResult struct {

	// REQUIRED; 请求次数据。
	RequestCntData []*DescribeImageXRequestCntUsageResResultRequestCntDataItem `json:"RequestCntData"`
}

type DescribeImageXRequestCntUsageResResultRequestCntDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXRequestCntUsageResResultRequestCntDataPropertiesItemsItem `json:"Data"`

	// 附加组件类型，GroupBy包含AdvFeat时有返回值。如：enhance，smartcut。取值为total，表示包含所选AdvFeat实际请求次数。
	AdvFeat string `json:"AdvFeat,omitempty"`

	// 模型名称，GroupBy包含Model时有返回值。
	Model string `json:"Model,omitempty"`

	// 服务 ID，GroupBy包含ServiceId时有返回值。
	ServiceID string `json:"ServiceId,omitempty"`
}

type DescribeImageXRequestCntUsageResResultRequestCntDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片结束时刻。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求次
	Value int `json:"Value"`
}

type DescribeImageXScreenshotUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 查询数据的时间粒度。单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXScreenshotUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXScreenshotUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXScreenshotUsageResResult `json:"Result"`
}

type DescribeImageXScreenshotUsageResResponseMetadata struct {

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

type DescribeImageXScreenshotUsageResResult struct {

	// REQUIRED; 截帧用量
	ScreenshotData []*DescribeImageXScreenshotUsageResResultScreenshotDataItem `json:"ScreenshotData"`
}

type DescribeImageXScreenshotUsageResResultScreenshotDataItem struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXScreenshotUsageResResultScreenshotDataPropertiesItemsItem `json:"Data"`
}

type DescribeImageXScreenshotUsageResResultScreenshotDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 截帧使用次数
	Value int `json:"Value"`
}

type DescribeImageXSensibleCacheHitRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// REQUIRED; 缓存类型。支持以下取值： 1：查询内存命中率数据； 2：查询磁盘命中率数据。
	Type string `json:"Type"`

	// 需要匹配 App 版本，缺省情况下则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXSensibleCacheHitRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXSensibleCacheHitRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXSensibleCacheHitRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSensibleCacheHitRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSensibleCacheHitRateByTimeResResult `json:"Result"`
}

type DescribeImageXSensibleCacheHitRateByTimeResResponseMetadata struct {

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

type DescribeImageXSensibleCacheHitRateByTimeResResult struct {

	// REQUIRED; 命中率数据
	CacheHitData []*DescribeImageXSensibleCacheHitRateByTimeResResultCacheHitDataItem `json:"CacheHitData"`
}

type DescribeImageXSensibleCacheHitRateByTimeResResultCacheHitDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 具体数据
	Data []*DescribeImageXSensibleCacheHitRateByTimeResResultCacheHitDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXSensibleCacheHitRateByTimeResResultCacheHitDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应的上报量
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 内存/磁盘命中率
	Value float64 `json:"Value"`
}

type DescribeImageXSensibleCountByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配 App 版本，缺省情况下则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXSensibleCountByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXSensibleCountByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXSensibleCountByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSensibleCountByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSensibleCountByTimeResResult `json:"Result"`
}

type DescribeImageXSensibleCountByTimeResResponseMetadata struct {

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

type DescribeImageXSensibleCountByTimeResResult struct {

	// REQUIRED; 采集样本量数据
	SensibleCountData []*DescribeImageXSensibleCountByTimeResResultSensibleCountDataItem `json:"SensibleCountData"`
}

type DescribeImageXSensibleCountByTimeResResultSensibleCountDataItem struct {

	// REQUIRED; 该数据类型对应的总上报量
	Count int `json:"Count"`

	// REQUIRED; 具体数据
	Data []*DescribeImageXSensibleCountByTimeResResultSensibleCountDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXSensibleCountByTimeResResultSensibleCountDataPropertiesItemsItem struct {

	// REQUIRED; 上报量数据
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 上报量数据，与Count数值相同。
	Value int `json:"Value"`
}

type DescribeImageXSensibleTopRAMURLBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 支持以下取值： 1：按上报次数排序； 2：按内存大小排序。
	OrderByIdx int `json:"OrderByIdx"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配 App 版本，缺省情况下则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXSensibleTopRAMURLBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 查询 Top URL 条数，取值范围为(0,1000]。缺省情况下默认为 1000。
	Top int `json:"Top,omitempty"`
}

type DescribeImageXSensibleTopRAMURLBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXSensibleTopRAMURLRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSensibleTopRAMURLResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSensibleTopRAMURLResResult `json:"Result"`
}

type DescribeImageXSensibleTopRAMURLResResponseMetadata struct {

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

type DescribeImageXSensibleTopRAMURLResResult struct {

	// REQUIRED; Top URL 数据
	TopURLData []*DescribeImageXSensibleTopRAMURLResResultTopURLDataItem `json:"TopUrlData"`
}

type DescribeImageXSensibleTopRAMURLResResultTopURLDataItem struct {

	// REQUIRED; Activity+View 树，控件信息。
	ActivityViewTree string `json:"ActivityViewTree"`

	// REQUIRED; 业务标识
	BizTag string `json:"BizTag"`

	// REQUIRED; 上报次数
	Count int `json:"Count"`

	// REQUIRED; 图片内存大小
	RAMSize int `json:"RamSize"`

	// REQUIRED; 图片URL
	URL string `json:"URL"`
}

type DescribeImageXSensibleTopResolutionURLBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 支持以下取值： 1：按上报次数排序； 2：按图片分辨率排序； 3：按 view 分辨率排序。
	OrderByIdx int `json:"OrderByIdx"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配 App 版本，缺省情况下则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXSensibleTopResolutionURLBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 查询 Top URL 条数，取值范围为(0,1000]。缺省情况下默认为 1000。
	Top int `json:"Top,omitempty"`
}

type DescribeImageXSensibleTopResolutionURLBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXSensibleTopResolutionURLRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSensibleTopResolutionURLResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSensibleTopResolutionURLResResult `json:"Result"`
}

type DescribeImageXSensibleTopResolutionURLResResponseMetadata struct {

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

type DescribeImageXSensibleTopResolutionURLResResult struct {

	// REQUIRED; 具体数据
	TopURLData []*DescribeImageXSensibleTopResolutionURLResResultTopURLDataItem `json:"TopUrlData"`
}

type DescribeImageXSensibleTopResolutionURLResResultTopURLDataItem struct {

	// REQUIRED; Activity+View 树，控件信息。
	ActivityViewTree string `json:"ActivityViewTree"`

	// REQUIRED; 业务标识
	BizTag string `json:"BizTag"`

	// REQUIRED; 上报次数
	Count int `json:"Count"`

	// REQUIRED; 图片高
	ImageHeight int `json:"ImageHeight"`

	// REQUIRED; 图片宽
	ImageWidth int `json:"ImageWidth"`

	// REQUIRED; 图片URL
	URL string `json:"URL"`

	// REQUIRED; 图片展示背景 view 高。
	ViewHeight int `json:"ViewHeight"`

	// REQUIRED; 图片展示背景 view 宽。
	ViewWidth int `json:"ViewWidth"`
}

type DescribeImageXSensibleTopSizeURLBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 支持以下取值： 1：按上报次数排序； 2：按文件体积排序。
	OrderByIdx int `json:"OrderByIdx"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配 App 版本，缺省情况下则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXSensibleTopSizeURLBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 查询 Top URL 条数，取值范围为(0,1000]。缺省情况下默认为 1000。
	Top int `json:"Top,omitempty"`
}

type DescribeImageXSensibleTopSizeURLBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXSensibleTopSizeURLRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSensibleTopSizeURLResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSensibleTopSizeURLResResult `json:"Result"`
}

type DescribeImageXSensibleTopSizeURLResResponseMetadata struct {

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

type DescribeImageXSensibleTopSizeURLResResult struct {

	// REQUIRED; Top URL数据
	TopURLData []*DescribeImageXSensibleTopSizeURLResResultTopURLDataItem `json:"TopUrlData"`
}

type DescribeImageXSensibleTopSizeURLResResultTopURLDataItem struct {

	// REQUIRED; Activity+View 树，控件信息。
	ActivityViewTree string `json:"ActivityViewTree"`

	// REQUIRED; 业务标识
	BizTag string `json:"BizTag"`

	// REQUIRED; 上报次数
	Count int `json:"Count"`

	// REQUIRED; 文件URL
	URL string `json:"URL"`

	// REQUIRED; 文件体积
	Value int `json:"Value"`
}

type DescribeImageXSensibleTopUnknownURLBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 支持以下取值： 1：按上报量排序 2：按内存大小排序 3：按文件大小排序 4：按图片分辨率排序 5：按 view 分辨率排序
	OrderByIdx int `json:"OrderByIdx"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配 App 版本，缺省情况下则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXSensibleTopUnknownURLBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的图片类型，不传则匹配所有图片类型。 GIF PNG JPEG HEIF HEIC WEBP AWEBP VVIC 其他
	ImageType []string `json:"ImageType,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,ImageType），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型，不传则匹配非WEB端所有系统。取值如下所示： iOS Android WEB
	OS string `json:"OS,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 查询 Top URL 条数，取值范围为(0,1000]。缺省情况下默认为 1000。
	Top int `json:"Top,omitempty"`
}

type DescribeImageXSensibleTopUnknownURLBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXSensibleTopUnknownURLRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSensibleTopUnknownURLResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSensibleTopUnknownURLResResult `json:"Result"`
}

type DescribeImageXSensibleTopUnknownURLResResponseMetadata struct {

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

type DescribeImageXSensibleTopUnknownURLResResult struct {

	// REQUIRED; Top URL 详情
	TopURLData []*DescribeImageXSensibleTopUnknownURLResResultTopURLDataItem `json:"TopUrlData"`
}

type DescribeImageXSensibleTopUnknownURLResResultTopURLDataItem struct {

	// Activity+View 树，控件信息
	ActivityViewTree string `json:"ActivityViewTree,omitempty"`

	// 业务标识
	BizTag string `json:"BizTag,omitempty"`

	// 上报次数
	Count int `json:"Count,omitempty"`

	// 文件大小
	FileSize int `json:"FileSize,omitempty"`

	// 图片高
	ImageHeight int `json:"ImageHeight,omitempty"`

	// 图片宽
	ImageWidth int `json:"ImageWidth,omitempty"`

	// 图片内存大小
	RAMSize int `json:"RamSize,omitempty"`

	// 图片 URL
	URL string `json:"URL,omitempty"`

	// 展示 view 高
	ViewHeight int `json:"ViewHeight,omitempty"`

	// 展示 view 宽
	ViewWidth int `json:"ViewWidth,omitempty"`
}

type DescribeImageXServerQPSUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 查询数据的时间粒度。单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 1: 单次查询最大时间跨度为 6 小时
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXServerQPSUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXServerQPSUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *DescribeImageXServerQPSUsageResResult `json:"Result"`
}

type DescribeImageXServerQPSUsageResResponseMetadata struct {

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

// DescribeImageXServerQPSUsageResResult - 视请求的接口而定
type DescribeImageXServerQPSUsageResResult struct {

	// REQUIRED; QPS 用量
	QPSData []*DescribeImageXServerQPSUsageResResultQPSDataItem `json:"QPSData"`
}

type DescribeImageXServerQPSUsageResResultQPSDataItem struct {

	// REQUIRED; 时序数据。
	Data []*DescribeImageXServerQPSUsageResResultQPSDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; QPS 类型，取值如下所示：
	// * Request：专有图像处理和高效压缩
	// * Mirror：资源下载与镜像代理
	QPSType string `json:"QPSType"`
}

type DescribeImageXServerQPSUsageResResultQPSDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; QPS 用量的值。单位为 "次/秒"，表示每秒处理的请求数量。
	Value int `json:"Value"`
}

type DescribeImageXServiceQualityQuery struct {

	// 获取指定地区的数据。默认为空，表示获取国内数据。
	// * cn：国内。
	// * sg：新加坡。
	Region string `json:"Region,omitempty" query:"Region"`
}

type DescribeImageXServiceQualityRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXServiceQualityResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *DescribeImageXServiceQualityResResult `json:"Result"`
}

type DescribeImageXServiceQualityResResponseMetadata struct {

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

// DescribeImageXServiceQualityResResult - 视请求的接口而定
type DescribeImageXServiceQualityResResult struct {

	// REQUIRED; 下行概览数据，为当日 0 点至今整体数据。
	CdnData *DescribeImageXServiceQualityResResultCdnData `json:"CdnData"`

	// REQUIRED; 客户端概览数据，为当日 0 点查询时间的整体数据。
	ClientData *DescribeImageXServiceQualityResResultClientData `json:"ClientData"`

	// REQUIRED; 用户感知失败率，为当日 0 点至查询时间的整体数据。
	FailureRate float64 `json:"FailureRate"`

	// REQUIRED; 上传概览数据，为当日 0 点查询时间的整体数据。
	UploadData *DescribeImageXServiceQualityResResultUploadData `json:"UploadData"`
}

// DescribeImageXServiceQualityResResultCdnData - 下行概览数据，为当日 0 点至今整体数据。
type DescribeImageXServiceQualityResResultCdnData struct {

	// REQUIRED; 网络平均耗时，单位为 ms。
	AvgDuration float64 `json:"AvgDuration"`

	// REQUIRED; 网络成功率
	SuccessRate float64 `json:"SuccessRate"`

	// REQUIRED; 下行总上报数据量。
	TotalCount int `json:"TotalCount"`
}

// DescribeImageXServiceQualityResResultClientData - 客户端概览数据，为当日 0 点查询时间的整体数据。
type DescribeImageXServiceQualityResResultClientData struct {

	// REQUIRED; 平均解码耗时，单位为 ms。
	AvgDecodeDuration float64 `json:"AvgDecodeDuration"`

	// REQUIRED; 平均加载耗时，单位为 ms。
	AvgLoadDuration float64 `json:"AvgLoadDuration"`

	// REQUIRED; 平均排队耗时，单位为 ms。
	AvgQueueDuration float64 `json:"AvgQueueDuration"`

	// REQUIRED; 解码成功率
	SuccessRate float64 `json:"SuccessRate"`

	// REQUIRED; 客户端总上报数据量。
	TotalCount int `json:"TotalCount"`
}

// DescribeImageXServiceQualityResResultUploadData - 上传概览数据，为当日 0 点查询时间的整体数据。
type DescribeImageXServiceQualityResResultUploadData struct {

	// REQUIRED; 上传平均耗时，单位为 ms。
	AvgDuration float64 `json:"AvgDuration"`

	// REQUIRED; 上传文件平均大小，单位字节。
	AvgSize float64 `json:"AvgSize"`

	// REQUIRED; 上传成功率
	SuccessRate float64 `json:"SuccessRate"`

	// REQUIRED; 上传总上报数据量。
	TotalCount int `json:"TotalCount"`

	// REQUIRED; 上传有效次数
	UploadCount int `json:"UploadCount"`
}

type DescribeImageXSourceRequestBandwidthQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 域名。为空时表示不筛选，支持查询多个域名，不同的域名使用逗号分隔。 您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 分组依据，仅支持取值DomainName。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 过滤运营商。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * 电信
	// * 联通
	// * 移动
	// * 鹏博士
	// * 教育网
	// * 广电网
	// * 其它
	Isp string `json:"Isp,omitempty" query:"Isp"`

	// 过滤网络协议。传入多个用英文逗号分割。缺省情况下表示不过滤。取值如下所示：
	// * HTTP
	// * HTTPS
	Protocols string `json:"Protocols,omitempty" query:"Protocols"`

	// 区域。传入多个时用英文逗号作为分割符，缺省情况下表示查询所有区域。取值如下所示：
	// * 中国大陆
	// * 亚太一区
	// * 亚太二区
	// * 亚太三区
	// * 欧洲区
	// * 北美区
	// * 南美区
	// * 中东区
	Regions string `json:"Regions,omitempty" query:"Regions"`

	// 服务 ID。为空时表示不筛选，支持查询多个服务，使用逗号分隔不同的服务。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 客户端国家。传入多个时用英文逗号作为分割符，缺省情况下表示不过滤。可调用DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]获取 IP
	// 解析后的客户端国家。取值如下所示：
	// * 海外，除中国外全部国家。
	// * 具体国家取值，如中国、美国。
	UserCountry string `json:"UserCountry,omitempty" query:"UserCountry"`

	// 客户端省份。传入多个用英文逗号分割。缺省情况下表示不过滤。可调用DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]获取 IP 解析后的客户端省份。
	UserProvince string `json:"UserProvince,omitempty" query:"UserProvince"`
}

type DescribeImageXSourceRequestBandwidthRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSourceRequestBandwidthResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSourceRequestBandwidthResResult `json:"Result"`
}

type DescribeImageXSourceRequestBandwidthResResponseMetadata struct {

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

type DescribeImageXSourceRequestBandwidthResResult struct {

	// REQUIRED; 带宽数据。
	BpsData []*DescribeImageXSourceRequestBandwidthResResultBpsDataItem `json:"BpsData"`
}

type DescribeImageXSourceRequestBandwidthResResultBpsDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXSourceRequestBandwidthResResultBpsDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy带有DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXSourceRequestBandwidthResResultBpsDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 带宽，单位为 bps。
	Value float64 `json:"Value"`
}

type DescribeImageXSourceRequestQuery struct {

	// REQUIRED; 状态码。传入多个时用英文逗号分隔。支持以下取值：
	// * req_cnt：返回所有状态码数据
	// * 2xx：返回 2xx 状态码汇总数据
	// * 3xx：返回 3xx 状态码汇总数据
	// * 4xx：返回 4xx 状态码汇总数据
	// * 5xx：返回 5xx 状态码汇总数据。
	DataTypes string `json:"DataTypes" query:"DataTypes"`

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 是否拆分状态码。取值如下所示：
	// * true：拆分
	// * false：（默认）不拆分
	DetailedCode bool `json:"DetailedCode,omitempty" query:"DetailedCode"`

	// 域名。为空时表示不筛选，支持查询多个域名，不同的域名使用逗号分隔。 您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 需要分组查询的参数。传入多个用英文逗号分割。支持以下取值：
	// * DomainName：域名
	// * DataType：状态码
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 过滤运营商。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * 电信
	// * 联通
	// * 移动
	// * 鹏博士
	// * 教育网
	// * 广电网
	// * 其它
	Isp string `json:"Isp,omitempty" query:"Isp"`

	// 过滤网络协议。传入多个用英文逗号分割。缺省情况下表示不过滤。取值如下所示：
	// * HTTP
	// * HTTPS
	Protocols string `json:"Protocols,omitempty" query:"Protocols"`

	// 区域。传入多个时用英文逗号作为分割符，缺省情况下表示查询所有区域。取值如下所示：
	// * 中国大陆
	// * 亚太一区
	// * 亚太二区
	// * 亚太三区
	// * 欧洲区
	// * 北美区
	// * 南美区
	// * 中东区
	Regions string `json:"Regions,omitempty" query:"Regions"`

	// 服务 ID。为空时表示不筛选，支持查询多个服务，使用逗号分隔不同的服务。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 客户端国家。传入多个时用英文逗号作为分割符，缺省情况下表示不过滤。可调用DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]获取 IP
	// 解析后的客户端国家。取值如下所示：
	// * 海外，除中国外全部国家。
	// * 具体国家取值，如中国、美国。
	UserCountry string `json:"UserCountry,omitempty" query:"UserCountry"`

	// 客户端省份。传入多个用英文逗号分割。缺省情况下表示不过滤。可调用DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]获取 IP 解析后的客户端省份。
	UserProvince string `json:"UserProvince,omitempty" query:"UserProvince"`
}

type DescribeImageXSourceRequestRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSourceRequestResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSourceRequestResResult `json:"Result"`
}

type DescribeImageXSourceRequestResResponseMetadata struct {

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

type DescribeImageXSourceRequestResResult struct {

	// REQUIRED; 数据列表 :::tip
	// * 若未设置Interval，则上报一条StartTime与EndTime时间段内查询的全部请求次数据；
	// * 若设置了Interval，则按Interval粒度分段上报查询的时间粒度内请求次数据； :::
	RequestCnt []*DescribeImageXSourceRequestResResultRequestCntItem `json:"RequestCnt"`
}

type DescribeImageXSourceRequestResResultRequestCntItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXSourceRequestResResultRequestCntPropertiesItemsItem `json:"Data"`

	// 当GroupBy带有DataType时出现
	DataType string `json:"DataType,omitempty"`

	// 当GroupBy带有DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXSourceRequestResResultRequestCntPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 请求次数，单位为次。
	Value int `json:"Value"`
}

type DescribeImageXSourceRequestTrafficQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 93 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 域名。为空时表示不筛选，支持查询多个域名，不同的域名使用逗号分隔。 您可以通过调用GetServiceDomains [https://www.volcengine.com/docs/508/9379]获取服务下所有域名信息。
	DomainNames string `json:"DomainNames,omitempty" query:"DomainNames"`

	// 分组依据，仅支持取值DomainName。
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 60：单次查询最大时间跨度为 6 小时
	// * 120：单次查询最大时间跨度为 6 小时
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 过滤运营商。传入多个用英文逗号分割，缺省情况下表示不过滤。取值如下所示：
	// * 电信
	// * 联通
	// * 移动
	// * 鹏博士
	// * 教育网
	// * 广电网
	// * 其它
	Isp string `json:"Isp,omitempty" query:"Isp"`

	// 过滤网络协议。传入多个用英文逗号分割。缺省情况下表示不过滤。取值如下所示：
	// * HTTP
	// * HTTPS
	Protocols string `json:"Protocols,omitempty" query:"Protocols"`

	// 区域。传入多个时用英文逗号作为分割符，缺省情况下表示查询所有区域。取值如下所示：
	// * 中国大陆
	// * 亚太一区
	// * 亚太二区
	// * 亚太三区
	// * 欧洲区
	// * 北美区
	// * 南美区
	// * 中东区
	Regions string `json:"Regions,omitempty" query:"Regions"`

	// 服务 ID。为空时表示不筛选，支持查询多个服务，使用逗号分隔不同的服务。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`

	// 客户端国家。传入多个时用英文逗号作为分割符，缺省情况下表示不过滤。可调用DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]获取 IP
	// 解析后的客户端国家。取值如下所示：
	// * 海外，除中国外全部国家。
	// * 具体国家取值，如中国、美国。
	UserCountry string `json:"UserCountry,omitempty" query:"UserCountry"`

	// 客户端省份。传入多个用英文逗号分割。缺省情况下表示不过滤。可调用DescribeImageXEdgeRequestRegions [https://www.volcengine.com/docs/508/1209574]获取 IP 解析后的客户端省份。
	UserProvince string `json:"UserProvince,omitempty" query:"UserProvince"`
}

type DescribeImageXSourceRequestTrafficRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSourceRequestTrafficResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSourceRequestTrafficResResult `json:"Result"`
}

type DescribeImageXSourceRequestTrafficResResponseMetadata struct {

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

type DescribeImageXSourceRequestTrafficResResult struct {

	// REQUIRED; 查询数据
	TrafficData []*DescribeImageXSourceRequestTrafficResResultTrafficDataItem `json:"TrafficData"`
}

type DescribeImageXSourceRequestTrafficResResultTrafficDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXSourceRequestTrafficResResultTrafficDataPropertiesItemsItem `json:"Data"`

	// 当GroupBy带有DomainName时出现
	DomainName string `json:"DomainName,omitempty"`
}

type DescribeImageXSourceRequestTrafficResResultTrafficDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 流量，单位为 byte。
	Value float64 `json:"Value"`
}

type DescribeImageXStorageUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 查询数据的时间粒度。单位为秒，缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。支持以下取值：
	// * 300：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval" query:"Interval"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。例如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// Bucket 名称。支持同时查询多个 BucketName，不同的 BucketNmae 使用逗号分隔。 您可以通过调用 GetAllImageServices [https://www.volcengine.com/docs/508/9360]
	// 获取所需的 Bucket 名称。
	BucketNames string `json:"BucketNames,omitempty" query:"BucketNames"`

	// 需要分组查询的参数，多个数据用逗号分隔。支持取值如下：
	// * ServiceId：服务 ID
	// * BucketName：Bucket 名称
	// * StorageType：存储类型
	GroupBy string `json:"GroupBy,omitempty" query:"GroupBy"`

	// 是否查询数据取回量。
	// * true：查询取回量。
	// * false：查询存储量。
	IsRetrieval bool `json:"IsRetrieval,omitempty" query:"IsRetrieval"`

	// 服务 ID。为空时表示不筛选，支持查询多个服务，使用逗号分隔不同的服务。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXStorageUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXStorageUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXStorageUsageResResult `json:"Result"`
}

type DescribeImageXStorageUsageResResponseMetadata struct {

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

type DescribeImageXStorageUsageResResult struct {

	// REQUIRED; 计量数据列表
	StorageData []*DescribeImageXStorageUsageResResultStorageDataItem `json:"StorageData"`
}

type DescribeImageXStorageUsageResResultStorageDataItem struct {

	// REQUIRED; 具体数据
	Data []*DescribeImageXStorageUsageResResultStorageDataPropertiesItemsItem `json:"Data"`

	// Bucket 名称，GroupBy包含BucketName时有返回值。
	BucketName string `json:"BucketName,omitempty"`

	// 服务 ID，GroupBy包含ServiceId时有返回值。
	ServiceID string `json:"ServiceId,omitempty"`

	// 存储类型，GroupBy包含StorageType时有返回值。 当传参IsRetrieval=false时，表示存储值，取值：
	// * STANDARD：标准存储
	// * IA：低频存储
	// * ARCHIVE：归档存储
	// * COLD_ARCHIVE：冷归档存储
	// * ARCHIVE_FR：归档闪回存储
	// 当传参IsRetrieval=true时，表示取回值，取值：
	// * IA：低频取回
	// * ARCHIVE：归档标准取回
	// * COLD_ARCHIVE：冷归档快速取回
	// * ARCHIVE_FR：归档闪回取回
	// * COLD_ARCHIVE_STANDARD：冷归档标准取回
	// * COLD_ARCHIVE_BULK：冷归档批量取回
	StorageType string `json:"StorageType,omitempty"`
}

type DescribeImageXStorageUsageResResultStorageDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点，时间片开始时刻，格式为：YYYY-MM-DDThh:mm:ss±hh:mm
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 单位：Byte
	Value int `json:"Value"`
}

type DescribeImageXSummaryQuery struct {

	// REQUIRED; 数据查询时间段，即Timestamp所在月份的 1 日 0 时起至传入时间Timestamp的时间范围。 格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	Timestamp string `json:"Timestamp" query:"Timestamp"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXSummaryRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXSummaryResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXSummaryResResult `json:"Result"`
}

type DescribeImageXSummaryResResponseMetadata struct {

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

type DescribeImageXSummaryResResult struct {

	// REQUIRED; 当月图像处理量
	BaseOpData *DescribeImageXSummaryResResultBaseOpData `json:"BaseOpData"`

	// REQUIRED; 当月带宽用量
	CdnBandwidthData *DescribeImageXSummaryResResultCdnBandwidthData `json:"CdnBandwidthData"`

	// REQUIRED; 当月流量用量
	CdnTrafficData *DescribeImageXSummaryResResultCdnTrafficData `json:"CdnTrafficData"`

	// REQUIRED; 当月源站请求次数
	RequestCntData *DescribeImageXSummaryResResultRequestCntData `json:"RequestCntData"`

	// REQUIRED; 当月资源占用量
	StorageData *DescribeImageXSummaryResResultStorageData `json:"StorageData"`
}

// DescribeImageXSummaryResResultBaseOpData - 当月图像处理量
type DescribeImageXSummaryResResultBaseOpData struct {

	// REQUIRED; 当月图像处理量，单位为：Byte
	Value int `json:"Value"`
}

// DescribeImageXSummaryResResultCdnBandwidthData - 当月带宽用量
type DescribeImageXSummaryResResultCdnBandwidthData struct {

	// REQUIRED; 当月带宽用量，单位为：bps
	Value float64 `json:"Value"`
}

// DescribeImageXSummaryResResultCdnTrafficData - 当月流量用量
type DescribeImageXSummaryResResultCdnTrafficData struct {

	// REQUIRED; 当月流量用量，单位为：Byte
	Value int `json:"Value"`
}

// DescribeImageXSummaryResResultRequestCntData - 当月源站请求次数
type DescribeImageXSummaryResResultRequestCntData struct {

	// REQUIRED; 当月源站请求次数，单位为：次
	Value int `json:"Value"`
}

// DescribeImageXSummaryResResultStorageData - 当月资源占用量
type DescribeImageXSummaryResResultStorageData struct {

	// REQUIRED; 当月最新一日资源占用量，单位为：Byte
	Value float64 `json:"Value"`
}

type DescribeImageXUploadCountByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadCountByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 上传类型，默认为空，返回上传 1.0 数据。 1：上传 1.0。 2：上传 2.0。
	UploadType int `json:"UploadType,omitempty"`
}

type DescribeImageXUploadCountByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadCountByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadCountByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadCountByTimeResResult `json:"Result"`
}

type DescribeImageXUploadCountByTimeResResponseMetadata struct {

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

type DescribeImageXUploadCountByTimeResResult struct {

	// REQUIRED; 上传有效次数数据。
	UploadCountData []*DescribeImageXUploadCountByTimeResResultUploadCountDataItem `json:"UploadCountData"`
}

type DescribeImageXUploadCountByTimeResResultUploadCountDataItem struct {

	// REQUIRED; 对应的总上传有效次数。
	Count int `json:"Count"`

	// REQUIRED; 对应的上传有效次数数据列表。
	Data []*DescribeImageXUploadCountByTimeResResultUploadCountDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXUploadCountByTimeResResultUploadCountDataPropertiesItemsItem struct {

	// REQUIRED; 上传有效次数
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 上传有效次数
	Value int `json:"Value"`
}

type DescribeImageXUploadDurationBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadDurationBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（分位数据）、Phase（分阶段数据，只有当UploadType=2时才能取该值）、公共维度（Appid,OS,AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 上传类型，默认为空，返回上传 1.0 数据。 1：上传 1.0。 2：上传 2.0。
	UploadType int `json:"UploadType,omitempty"`
}

type DescribeImageXUploadDurationBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadDurationRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadDurationResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadDurationResResult `json:"Result"`
}

type DescribeImageXUploadDurationResResponseMetadata struct {

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

type DescribeImageXUploadDurationResResult struct {

	// REQUIRED; 上传耗时数据
	DurationData []*DescribeImageXUploadDurationResResultDurationDataItem `json:"DurationData"`
}

type DescribeImageXUploadDurationResResultDurationDataItem struct {

	// REQUIRED; 对应的总次数。
	Count int `json:"Count"`

	// REQUIRED; 上传耗时数据列表。
	Data []*DescribeImageXUploadDurationResResultDurationDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 当GroupBy为空或Duration时，取值为：pct25、pct50、pct90、pct99、avg。 当GroupBy取值为Phase时，取值为：1001、1002、1003、1004、1005。 除上述外取值为指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXUploadDurationResResultDurationDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应上传次数。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 上传耗时，单位为 ms。
	Value float64 `json:"Value"`
}

type DescribeImageXUploadErrorCodeAllBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 聚合维度。 ErrorCode：错误码 Region：地区 Isp：运营商
	GroupBy string `json:"GroupBy"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadErrorCodeAllBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 是否升序排序。不传则默认降序排序。
	OrderAsc bool `json:"OrderAsc,omitempty"`

	// 目前仅支持传入Count按错误码数量排序。不传或者传空默认按错误码数量排序。
	OrderBy string `json:"OrderBy,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 上传类型，默认为空，返回上传 1.0 数据。 1：上传 1.0。 2：上传 2.0。
	UploadType int `json:"UploadType,omitempty"`
}

type DescribeImageXUploadErrorCodeAllBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadErrorCodeAllRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadErrorCodeAllResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadErrorCodeAllResResult `json:"Result"`
}

type DescribeImageXUploadErrorCodeAllResResponseMetadata struct {

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

type DescribeImageXUploadErrorCodeAllResResult struct {

	// REQUIRED; 错误码数据。
	ErrorCodeData []*DescribeImageXUploadErrorCodeAllResResultErrorCodeDataItem `json:"ErrorCodeData"`
}

type DescribeImageXUploadErrorCodeAllResResultErrorCodeDataItem struct {

	// REQUIRED; 错误码数量。
	Value int `json:"Value"`

	// 当GroupBy取值非ErrorCode时出现，表示错误码详细信息。
	Details []*DescribeImageXUploadErrorCodeAllResResultErrorCodeDataPropertiesItemsItem `json:"Details,omitempty"`

	// 当GroupBy取值ErrorCode时出现，表示错误码内容。
	ErrorCode string `json:"ErrorCode,omitempty"`

	// 当GroupBy取值ErrorCode时出现，表示错误码的描述信息。
	ErrorCodeDesc string `json:"ErrorCodeDesc,omitempty"`

	// 当GroupBy取值Isp时出现，则表示运营商信息。
	Isp string `json:"Isp,omitempty"`

	// 当GroupBy取值Region时出现，表示地区信息。
	Region string `json:"Region,omitempty"`

	// 当GroupBy取值Region时出现，表示地区类型。 Country：国家； Province：省份。
	RegionType string `json:"RegionType,omitempty"`
}

type DescribeImageXUploadErrorCodeAllResResultErrorCodeDataPropertiesItemsItem struct {

	// REQUIRED; 错误码内容
	ErrorCode string `json:"ErrorCode"`

	// REQUIRED; 错误码的描述信息
	ErrorCodeDesc string `json:"ErrorCodeDesc"`

	// REQUIRED; 错误码数量
	Value int `json:"Value"`
}

type DescribeImageXUploadErrorCodeByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadErrorCodeByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 上传类型，默认为空，返回上传 1.0 数据。 1：上传 1.0。 2：上传 2.0。
	UploadType int `json:"UploadType,omitempty"`
}

type DescribeImageXUploadErrorCodeByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadErrorCodeByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadErrorCodeByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadErrorCodeByTimeResResult `json:"Result"`
}

type DescribeImageXUploadErrorCodeByTimeResResponseMetadata struct {

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

type DescribeImageXUploadErrorCodeByTimeResResult struct {

	// REQUIRED; 所有错误码数据。
	ErrorCodeData []*DescribeImageXUploadErrorCodeByTimeResResultErrorCodeDataItem `json:"ErrorCodeData"`
}

type DescribeImageXUploadErrorCodeByTimeResResultErrorCodeDataItem struct {

	// REQUIRED; 错误码数量。
	Count int `json:"Count"`

	// REQUIRED; 错误码对应的时序数据。
	Data []*DescribeImageXUploadErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 错误码。
	ErrorCode string `json:"ErrorCode"`
}

type DescribeImageXUploadErrorCodeByTimeResResultErrorCodeDataPropertiesItemsItem struct {

	// REQUIRED; 错误码数量。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 错误码数量。
	Value int `json:"Value"`
}

type DescribeImageXUploadFileSizeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadFileSizeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（分位数据）、公共维度（Appid,OS,AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 上传类型，默认为空，返回上传 1.0 数据。 1：上传 1.0。 2：上传 2.0。
	UploadType int `json:"UploadType,omitempty"`
}

type DescribeImageXUploadFileSizeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadFileSizeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadFileSizeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadFileSizeResResult `json:"Result"`
}

type DescribeImageXUploadFileSizeResResponseMetadata struct {

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

type DescribeImageXUploadFileSizeResResult struct {

	// REQUIRED; 上传文件大小分布数据。
	FSizeData []*DescribeImageXUploadFileSizeResResultFSizeDataItem `json:"FSizeData"`
}

type DescribeImageXUploadFileSizeResResultFSizeDataItem struct {

	// REQUIRED; 对应的总次数。
	Count int `json:"Count"`

	// REQUIRED; 对应的文件大小数据列表。
	Data []*DescribeImageXUploadFileSizeResResultFSizeDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 GroupBy为空或Duration时，取值为：pct25、pct50、pct90、pct99、avg。 GroupBy取值为其他时，则指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXUploadFileSizeResResultFSizeDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应次数。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 对应文件大小。
	Value float64 `json:"Value"`
}

type DescribeImageXUploadSegmentSpeedByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// REQUIRED; 上传类型，固定值传入2，表示上传 2.0 数据。
	UploadType int `json:"UploadType"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadSegmentSpeedByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（分位数据）、公共维度（Appid,OS,AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`
}

type DescribeImageXUploadSegmentSpeedByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadSegmentSpeedByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadSegmentSpeedByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadSegmentSpeedByTimeResResult `json:"Result"`
}

type DescribeImageXUploadSegmentSpeedByTimeResResponseMetadata struct {

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

type DescribeImageXUploadSegmentSpeedByTimeResResult struct {

	// REQUIRED; 上传速度数据
	SpeedData []*DescribeImageXUploadSegmentSpeedByTimeResResultSpeedDataItem `json:"SpeedData"`
}

type DescribeImageXUploadSegmentSpeedByTimeResResultSpeedDataItem struct {

	// REQUIRED; 总上传次数
	Count int `json:"Count"`

	// REQUIRED; 具体数据
	Data []*DescribeImageXUploadSegmentSpeedByTimeResResultSpeedDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 GroupBy为空或Duration时，取值为：pct25、pct50、pct90、pct99、avg。 GroupBy取值为其他时，则指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXUploadSegmentSpeedByTimeResResultSpeedDataPropertiesItemsItem struct {

	// REQUIRED; 上传次数
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 上传分片速度，单位KB/s
	Value float64 `json:"Value"`
}

type DescribeImageXUploadSpeedBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadSpeedBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，表示拆分分位数据。支持取值：Duration（分位数据）、公共维度（Appid,OS,AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 上传类型，默认为空，返回上传 1.0 数据。 1：上传 1.0。 2：上传 2.0。
	UploadType int `json:"UploadType,omitempty"`
}

type DescribeImageXUploadSpeedBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadSpeedRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadSpeedResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadSpeedResResult `json:"Result"`
}

type DescribeImageXUploadSpeedResResponseMetadata struct {

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

type DescribeImageXUploadSpeedResResult struct {

	// REQUIRED; 上传速度数据
	SpeedData []*DescribeImageXUploadSpeedResResultSpeedDataItem `json:"SpeedData"`
}

type DescribeImageXUploadSpeedResResultSpeedDataItem struct {

	// REQUIRED; 对应的总次数。
	Count int `json:"Count"`

	// REQUIRED; 对应的上传速度数据列表。
	Data []*DescribeImageXUploadSpeedResResultSpeedDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型。 GroupBy为空或Duration时，取值为：pct25、pct50、pct90、pct99、avg。 GroupBy取值为其他时，则指定拆分维度的各个值。
	Type string `json:"Type"`
}

type DescribeImageXUploadSpeedResResultSpeedDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应上传次数。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 平均速度，单位为 KB/s。
	Value float64 `json:"Value"`
}

type DescribeImageXUploadSuccessRateByTimeBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	EndTime string `json:"EndTime"`

	// REQUIRED; 返回数据的时间粒度。 5m：为 5 分钟； 1h：为 1 小时； 1d：为 1 天。
	Granularity string `json:"Granularity"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// 需要匹配的 App 版本，不传则匹配 App 的所有版本。
	AppVer []string `json:"AppVer,omitempty"`

	// 应用 ID。默认为空，匹配账号下的所有的App ID。
	Appid string `json:"Appid,omitempty"`

	// 需要匹配的国家名称。 不传则匹配所有国家。 取值为海外时，匹配海外所有国家。
	Country string `json:"Country,omitempty"`

	// 需要匹配的自定义维度项
	ExtraDims []*DescribeImageXUploadSuccessRateByTimeBodyExtraDimsItem `json:"ExtraDims,omitempty"`

	// 拆分维度。默认为空，不拆分。支持取值：公共维度（Appid,OS,AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	GroupBy string `json:"GroupBy,omitempty"`

	// 需要匹配的运营商名称，不传则匹配所有运营商。支持取值如下： 电信 联通 移动 铁通 鹏博士 教育网 其他
	Isp []string `json:"Isp,omitempty"`

	// 取值为不等于的维度（默认为等于）。支持取值：公共维度（AppVer,SdkVer,Country,Province,Isp），自定义维度（通过"获取自定义维度列表"接口获取）
	Not []string `json:"Not,omitempty"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web。
	// * Imp：小程序。
	OS string `json:"OS,omitempty"`

	// 需要匹配的省份名称，不传则匹配所有省份
	Province string `json:"Province,omitempty"`

	// 需要匹配的SDK版本，不传则匹配所有版本
	SdkVer []string `json:"SdkVer,omitempty"`

	// 上传类型，默认为空，返回上传 1.0 数据。 1：上传 1.0。 2：上传 2.0。
	UploadType int `json:"UploadType,omitempty"`
}

type DescribeImageXUploadSuccessRateByTimeBodyExtraDimsItem struct {

	// REQUIRED; 自定义维度名称。
	Dim string `json:"Dim"`

	// REQUIRED; 需要匹配的对应维度值
	Vals []string `json:"Vals"`
}

type DescribeImageXUploadSuccessRateByTimeRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXUploadSuccessRateByTimeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXUploadSuccessRateByTimeResResult `json:"Result"`
}

type DescribeImageXUploadSuccessRateByTimeResResponseMetadata struct {

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

type DescribeImageXUploadSuccessRateByTimeResResult struct {

	// REQUIRED; 上传成功率数据。
	SuccessRateData []*DescribeImageXUploadSuccessRateByTimeResResultSuccessRateDataItem `json:"SuccessRateData"`
}

type DescribeImageXUploadSuccessRateByTimeResResultSuccessRateDataItem struct {

	// REQUIRED; 对应的总上传有效次数。
	Count int `json:"Count"`

	// REQUIRED; 对应的上传成功率数据列表。
	Data []*DescribeImageXUploadSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem `json:"Data"`

	// REQUIRED; 数据类型，不拆分时值为Total，拆分时为具体的维度值
	Type string `json:"Type"`
}

type DescribeImageXUploadSuccessRateByTimeResResultSuccessRateDataPropertiesItemsItem struct {

	// REQUIRED; 数据对应上传有效次数。
	Count int `json:"Count"`

	// REQUIRED; 数据对应时间点，按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	Timestamp string `json:"Timestamp"`

	// REQUIRED; 上传成功率
	Value float64 `json:"Value"`
}

type DescribeImageXVideoClipDurationUsageQuery struct {

	// REQUIRED; 获取数据结束时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照 ISO8601 表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近一年历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2010-11-21T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime" query:"StartTime"`

	// 查询数据的时间粒度。单位为秒。缺省时查询StartTime和EndTime时间段全部数据，此时单次查询最大时间跨度为 93 天。取值如下所示：
	// * 300：单次查询最大时间跨度为 31 天
	// * 600：单次查询最大时间跨度为 31 天
	// * 1200：单次查询最大时间跨度为 31 天
	// * 1800：单次查询最大时间跨度为 31 天
	// * 3600：单次查询最大时间跨度为 93 天
	// * 86400：单次查询最大时间跨度为 93 天
	// * 604800：单次查询最大时间跨度为 93 天
	Interval string `json:"Interval,omitempty" query:"Interval"`

	// 服务 ID。支持查询多个服务，传入多个时用英文逗号“,”分割，缺省情况下表示查询所有服务。您可以在 veImageX 控制台的服务管理 [https://console.volcengine.com/imagex/service_manage/]模块或者调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口获取服务
	// ID。
	ServiceIDs string `json:"ServiceIds,omitempty" query:"ServiceIds"`
}

type DescribeImageXVideoClipDurationUsageRes struct {

	// REQUIRED
	ResponseMetadata *DescribeImageXVideoClipDurationUsageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *DescribeImageXVideoClipDurationUsageResResult `json:"Result"`
}

type DescribeImageXVideoClipDurationUsageResResponseMetadata struct {

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

type DescribeImageXVideoClipDurationUsageResResult struct {

	// REQUIRED; 小视频转动图转换时长
	VideoClipDurationData []*DescribeImageXVideoClipDurationUsageResResultVideoClipDurationDataItem `json:"VideoClipDurationData"`
}

type DescribeImageXVideoClipDurationUsageResResultVideoClipDurationDataItem struct {

	// REQUIRED; 时序数据
	Data []*DescribeImageXVideoClipDurationUsageResResultVideoClipDurationDataPropertiesItemsItem `json:"Data"`
}

type DescribeImageXVideoClipDurationUsageResResultVideoClipDurationDataPropertiesItemsItem struct {

	// REQUIRED; 统计时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 转换时长，单位ms
	Value float64 `json:"Value"`
}

type DownloadCertQuery struct {

	// REQUIRED; 证书 ID，您可以通过调用 获取账号下全部证书 [https://www.volcengine.com/docs/508/66017] 获取账号下所有证书信息。
	CertID string `json:"CertID" query:"CertID"`
}

type DownloadCertRes struct {

	// REQUIRED
	ResponseMetadata *DownloadCertResResponseMetadata `json:"ResponseMetadata"`
	Result           *DownloadCertResResult           `json:"Result,omitempty"`
}

type DownloadCertResResponseMetadata struct {

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

type DownloadCertResResult struct {

	// 证书压缩数据base64后的字符串
	ZipData string `json:"ZipData,omitempty"`

	// 证书zip压缩文件名
	ZipName string `json:"ZipName,omitempty"`
}

type ExportFailedMigrateTaskQuery struct {

	// REQUIRED; 任务地区（即任务目标服务的地区），默认空，返回国内任务。
	// * cn：国内
	// * sg：新加坡
	Region string `json:"Region" query:"Region"`

	// REQUIRED; 任务 ID，请参考CreateImageMigrateTask [https://www.volcengine.com/docs/508/1009929]获取返回的任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type ExportFailedMigrateTaskRes struct {

	// REQUIRED
	ResponseMetadata *ExportFailedMigrateTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ExportFailedMigrateTaskResResult `json:"Result,omitempty"`
}

type ExportFailedMigrateTaskResResponseMetadata struct {

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

// ExportFailedMigrateTaskResResult - 视请求的接口而定
type ExportFailedMigrateTaskResResult struct {

	// REQUIRED; 迁移失败文件列表。仅当任务状态为Partial时有值。
	Entries []*ExportFailedMigrateTaskResResultEntriesItem `json:"Entries"`
}

type ExportFailedMigrateTaskResResultEntriesItem struct {

	// REQUIRED; 失败错误码
	ErrCode int `json:"ErrCode"`

	// REQUIRED; 失败原因
	ErrMsg string `json:"ErrMsg"`

	// REQUIRED; 文件地址
	Key string `json:"Key"`
}

type FetchImageURLBody struct {

	// REQUIRED; 目标服务 ID，迁移后的文件将上传至该服务绑定的存储。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待抓取上传的文件 URL。
	URL string `json:"Url"`

	// 是否采用异步，取值如下所示：
	// * true：采用异步
	// * false：（默认）不采用异步
	// :::tip 若您的资源大小小于 5 G，但预估资源迁移超时时间超过 20 s，建议您选择异步处理。 :::
	Async bool `json:"Async,omitempty"`

	// 回调 URL，veImageX 以 Post 方式向业务服务器发送 JSON 格式回调数据。当Async取值为true，即采用异步处理时，为必填。
	Callback string `json:"Callback,omitempty"`

	// 透传给业务的回调内容，当Callback不为空时为必填，取值需要符合CallbackBodyType指定格式。
	CallbackBody string `json:"CallbackBody,omitempty"`

	// 透传给业务的回调内容格式。当CallbackBody不为空时为必填。取值如下所示：
	// * application/json
	// * application/x-www-form-urlencoded
	CallbackBodyType string `json:"CallbackBodyType,omitempty"`

	// 回调时使用的 IP 地址
	CallbackHost string `json:"CallbackHost,omitempty"`

	// 是否仅迁移文件，取值如下所示：
	// * true：仅将文件迁移至目标服务对应的存储。适用于文件快速迁移且无需获取图片元信息场景，例如对时间敏感度极高的文件传输任务。
	// * false：（默认）迁移文件的同时，对图片类文件进行解码处理。适用于需要获取图片元信息而对迁移时间要求不高的场景。解码图片资源后，您可在返回参数获取图片的元信息，包括图片宽高、图片类型、动图的时间和帧数等，便于后续的图片分析和管理。
	FetchOnly bool `json:"FetchOnly,omitempty"`

	// 迁移资源的 IP 地址
	Host string `json:"Host,omitempty"`

	// 服务存储中存在相同的存储 key 时，是否忽略本次迁移操作。取值如下所示：
	// * true：忽略本次迁移操作。
	// * false：（默认）继续迁移并覆盖相同 key 的资源。
	IgnoreSameKey bool `json:"IgnoreSameKey,omitempty"`

	// 校验下载文件的 MD5，若校验不一致则停止文件的上传。
	MD5 string `json:"MD5,omitempty"`

	// 请求 header
	RequestHeader map[string][]string `json:"RequestHeader,omitempty"`

	// 指定抓取成功后的文件存储 key，不包含 bucket 部分。默认使用随机生成的 key。 :::tip 若指定 key 已存在，则会覆盖服务中 StoreKey 对应的已有文件，此时服务中保存文件的数量未发生变化。 :::
	StoreKey string `json:"StoreKey,omitempty"`

	// 资源下载超时时间。
	// * 同步处理下最大超时为 20 秒；
	// * 异步处理下最大超时为 90 秒。
	TimeOut int `json:"TimeOut,omitempty"`
}

type FetchImageURLRes struct {

	// REQUIRED
	ResponseMetadata *FetchImageURLResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *FetchImageURLResResult `json:"Result,omitempty"`
}

type FetchImageURLResResponseMetadata struct {

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

// FetchImageURLResResult - 视请求的接口而定
type FetchImageURLResResult struct {

	// REQUIRED; 源文件 URL
	URL string `json:"Url"`

	// 动图持续时间，在同步处理情况下、迁移至图像处理服务且为图片资源时有返回值。
	Duration int `json:"Duration,omitempty"`

	// 任务结束执行时间戳，UTC 时间，单位为 ns。
	EndTime int `json:"EndTime,omitempty"`

	// 文件大小，单位为 byte。同步处理情况下有返回值。
	FSize int `json:"FSize,omitempty"`

	// 图片帧数，在同步处理情况下、迁移至图像处理服务且为图片资源时有返回值。
	FrameCnt int `json:"FrameCnt,omitempty"`

	// 图片类型，在同步处理情况下、迁移至图像处理服务且为图片资源时有返回值。
	ImageFormat string `json:"ImageFormat,omitempty"`

	// 图片高，在同步处理情况下、迁移至图像处理服务且为图片资源时有返回值。
	ImageHeight int `json:"ImageHeight,omitempty"`

	// 图片宽，在同步处理情况下、迁移至图像处理服务且为图片资源时有返回值。
	ImageWidth int `json:"ImageWidth,omitempty"`

	// 任务开始执行时间戳，UTC 时间，单位为 ns。
	StartTime int `json:"StartTime,omitempty"`

	// 迁移后的文件 URI，包含 bucket/key 两部分。
	StoreURI string `json:"StoreUri,omitempty"`

	// 异步任务 ID，仅当Async取值true，即采用异步时有返回值。
	TaskID string `json:"TaskId,omitempty"`

	// 完成任务总耗时，单位为毫秒。
	TimeCost int `json:"TimeCost,omitempty"`
}

type GetAllCertsRes struct {

	// REQUIRED
	ResponseMetadata *GetAllCertsResResponseMetadata `json:"ResponseMetadata"`
	Result           []*GetAllCertsResResultItem     `json:"Result,omitempty"`
}

type GetAllCertsResResponseMetadata struct {

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

type GetAllCertsResResultItem struct {

	// REQUIRED; 证书ID
	CertID string `json:"CertID"`

	// REQUIRED; 证书名
	CertName string `json:"CertName"`

	// REQUIRED; 允许https使用的域名
	CommonName string `json:"CommonName"`

	// REQUIRED; 创建时间戳
	CreateTime float64 `json:"CreateTime"`

	// REQUIRED; 签发机构
	Issuer string `json:"Issuer"`

	// REQUIRED; 过期时间
	NotAfter float64 `json:"NotAfter"`

	// REQUIRED
	San []string `json:"San"`
}

type GetAllImageServicesQuery struct {

	// 筛选服务的参数，当该值为空时返回所有服务，指定后返回服务名或者 ID 中包含该字符串的服务。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`
}

type GetAllImageServicesRes struct {

	// REQUIRED
	ResponseMetadata *GetAllImageServicesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetAllImageServicesResResult `json:"Result,omitempty"`
}

type GetAllImageServicesResResponseMetadata struct {

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

// GetAllImageServicesResResult - 视请求的接口而定
type GetAllImageServicesResResult struct {

	// REQUIRED; 所有的服务信息。
	Services []*GetAllImageServicesResResultServicesItem `json:"Services"`
}

type GetAllImageServicesResResultServicesItem struct {

	// REQUIRED; 服务的授权 Bucket 列表。
	AllowBkts []string `json:"AllowBkts"`

	// REQUIRED; 是否允许配置其他镜像站类型，取值如下所示：
	// * true：是
	// * false：否
	AllowMirrorTypes bool `json:"AllowMirrorTypes"`

	// REQUIRED; 是否开启精简 URL，取值如下所示：
	// * true：是
	// * false：否
	CompactURL bool `json:"CompactURL"`

	// REQUIRED; 服务创建时间，即创建时当地服务器时间。
	CreateAt string `json:"CreateAt"`

	// REQUIRED; 绑定域名的相关信息。
	DomainInfos []*GetAllImageServicesResResultServicesPropertiesItemsItem `json:"DomainInfos"`

	// REQUIRED; 是否配置鉴权 key，取值如下所示：
	// * true：是
	// * false：否
	HasSigkey bool `json:"HasSigkey"`

	// REQUIRED
	ImageY bool `json:"ImageY"`

	// REQUIRED; 自定义处理样式具体配置
	ImageYAttribute *GetAllImageServicesResResultServicesItemImageYAttribute `json:"ImageYAttribute"`

	// REQUIRED; 镜像回源配置。
	Mirror *GetAllImageServicesResResultServicesItemMirror `json:"Mirror"`

	// REQUIRED; 是否开启源地址访问，取值如下所示：
	// * true：是
	// * false：否
	ObjectAccess bool `json:"ObjectAccess"`

	// REQUIRED; 主鉴权 Key。
	PrimaryKey string `json:"PrimaryKey"`

	// REQUIRED; 仅tob账号有值
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 资源封禁配置
	ResourceLimitedVisit *GetAllImageServicesResResultServicesItemResourceLimitedVisit `json:"ResourceLimitedVisit"`

	// REQUIRED; 仅tob账号有值
	ResourceTags []*Components18K1KvdSchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesResourcetagsItems `json:"ResourceTags"`

	// REQUIRED; 备鉴权 Key。
	SecondaryKey string `json:"SecondaryKey"`

	// REQUIRED; 服务 ID。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 服务名称。
	ServiceName string `json:"ServiceName"`

	// REQUIRED; 服务地域，取值如下所示：
	// * cn：中国
	// * sg：新加坡
	ServiceRegion string `json:"ServiceRegion"`

	// REQUIRED; 服务状态。状态分为未审核、审核未通过、正常、禁用。 :::tip
	// * 只有服务状态为正常时，该服务才可用。
	// * 如果是其他异常状态，请参考服务管理 [https://www.volcengine.com/docs/508/8086]进行处理。 :::
	ServiceStatus string `json:"ServiceStatus"`

	// REQUIRED; 服务类型，取值如下所示：
	// * StaticRc：素材托管服务
	// * Image：图片处理服务
	ServiceType string `json:"ServiceType"`

	// REQUIRED; 资源配置。
	Storage *GetAllImageServicesResResultServicesItemStorage `json:"Storage"`

	// REQUIRED; 该服务的图片模板固定前缀。
	TemplatePrefix string `json:"TemplatePrefix"`

	// REQUIRED; 是否开启覆盖上传，取值如下所示：
	// * true：是
	// * false：否
	UploadOverwrite bool `json:"UploadOverwrite"`

	// REQUIRED; 绑定点播空间配置
	VodSpace *GetAllImageServicesResResultServicesItemVodSpace `json:"VodSpace"`

	// 事件通知规则
	EventRules []*Components16Kv6ElSchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesEventrulesItems `json:"EventRules,omitempty"`

	// 降冷存储配置
	StorageRules []*Components1D40MkcSchemasGetallimageservicesresPropertiesResultPropertiesServicesItemsPropertiesStoragerulesItems `json:"StorageRules,omitempty"`

	// 版本控制启用状态，取值如下所示：
	// * 0：未开启
	// * 1：已开启
	// * 2：暂停
	StorageVersioning int `json:"StorageVersioning,omitempty"`
}

// GetAllImageServicesResResultServicesItemImageYAttribute - 自定义处理样式具体配置
type GetAllImageServicesResResultServicesItemImageYAttribute struct {

	// REQUIRED; 是否开启原图保护，取值如下所示：
	// * true：开启
	// * false：关闭
	ImageProtect bool `json:"ImageProtect"`

	// REQUIRED; 图像样式分隔符。
	ImageStyleSeparators []string `json:"ImageStyleSeparators"`

	// REQUIRED
	QnCosPreference string `json:"QnCosPreference"`

	// REQUIRED
	QueryStyleCombine bool `json:"QueryStyleCombine"`
}

// GetAllImageServicesResResultServicesItemMirror - 镜像回源配置。
type GetAllImageServicesResResultServicesItemMirror struct {

	// REQUIRED; 镜像回源下载原图时，携带的 HTTP 头部，键值都为 String 类型。
	Headers map[string]string `json:"Headers"`

	// REQUIRED; 镜像回源域名。
	Host string `json:"Host"`

	// REQUIRED; 带权重回源域名，key 为 String 类型时，代表镜像回源域名；value 为 Integer 类型时，代表域名权重。
	Hosts map[string]int `json:"Hosts"`

	// REQUIRED; 下载图片的协议，取值如下所示：
	// * http
	// * https
	Schema string `json:"Schema"`

	// REQUIRED; 镜像源 URI，其中图片名用 %s 占位符替代，比如/obj/%s。
	Source string `json:"Source"`
}

// GetAllImageServicesResResultServicesItemMirrorItemOrigin - 镜像源站
type GetAllImageServicesResResultServicesItemMirrorItemOrigin struct {

	// REQUIRED; 镜像源站参数
	Param interface{} `json:"Param"`

	// REQUIRED; 镜像源站类型
	Type string `json:"Type"`
}

// GetAllImageServicesResResultServicesItemResourceLimitedVisit - 资源封禁配置
type GetAllImageServicesResResultServicesItemResourceLimitedVisit struct {

	// REQUIRED; 域名白名单列表，封禁资源仅可被白名单的域名访问。
	AllowDomains []string `json:"AllowDomains"`

	// REQUIRED; 资源封禁开关，取值如下所示：
	// * true：开启
	// * false：关闭
	Enable bool `json:"Enable"`
}

// GetAllImageServicesResResultServicesItemStorage - 资源配置。
type GetAllImageServicesResResultServicesItemStorage struct {

	// REQUIRED; 是否支持任意文件格式上传，取值如下所示：
	// * true：是
	// * false：否
	AllTypes bool `json:"AllTypes"`

	// REQUIRED; 存储 Bucket 名称。
	BktName string `json:"BktName"`

	// REQUIRED; 保存时间，单位为秒。
	TTL int `json:"TTL"`
}

// GetAllImageServicesResResultServicesItemVodSpace - 绑定点播空间配置
type GetAllImageServicesResResultServicesItemVodSpace struct {

	// REQUIRED; 点播空间存储桶名称
	Bucket string `json:"Bucket"`

	// REQUIRED; 空间所在地区
	Region string `json:"Region"`

	// REQUIRED; 点播空间名
	SpaceName string `json:"SpaceName"`
}

type GetAllImageServicesResResultServicesPropertiesItemsItem struct {

	// REQUIRED; 域名解析到的 cname。
	CNAME string `json:"CNAME"`

	// REQUIRED; 绑定的域名。
	DomainName string `json:"DomainName"`

	// REQUIRED; 是否为默认域名，取值如下所示：
	// * true：是
	// * false：否
	IsDefault bool `json:"IsDefault"`

	// REQUIRED; 域名状态。
	Status string `json:"Status"`

	// REQUIRED; 是否开启鉴权
	URLAuth bool `json:"UrlAuth"`
}

type GetAllImageTemplatesQuery struct {

	// REQUIRED; 服务ID
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 是否按照模板创建时间升序查询，默认为false。
	// * 取值为true时，表示按升序查询。
	// * 取值为false时，表示降序查询。
	Asc string `json:"Asc,omitempty" query:"Asc"`

	// 分页获取条数，默认 10。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，默认 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 支持的字符正则集合为[a-zA-Z0-9_-]。指定时返回模板名称包含该字符串的图片模板，不填或者为空则返回所有模板。
	TemplateNamePattern string `json:"TemplateNamePattern,omitempty" query:"TemplateNamePattern"`
}

type GetAllImageTemplatesRes struct {

	// REQUIRED
	ResponseMetadata *GetAllImageTemplatesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetAllImageTemplatesResResult `json:"Result"`
}

type GetAllImageTemplatesResResponseMetadata struct {

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

type GetAllImageTemplatesResResult struct {

	// REQUIRED
	ServiceID string `json:"ServiceId"`

	// REQUIRED
	Templates []*GetAllImageTemplatesResResultTemplatesItem `json:"Templates"`

	// REQUIRED
	Total int `json:"Total"`
}

type GetAllImageTemplatesResResultTemplatesItem struct {

	// REQUIRED
	Abstract []string `json:"Abstract"`

	// REQUIRED; Anything
	AdaptiveFmt interface{} `json:"AdaptiveFmt"`

	// REQUIRED; Anything
	AnimExtract interface{} `json:"AnimExtract"`

	// REQUIRED; Anything
	Animation interface{} `json:"Animation"`

	// REQUIRED
	AutoQuality bool `json:"AutoQuality"`

	// REQUIRED
	Content string `json:"Content"`

	// REQUIRED
	CreateAt string `json:"CreateAt"`

	// REQUIRED
	Creator string `json:"Creator"`

	// REQUIRED
	DemotionFormat string `json:"DemotionFormat"`

	// REQUIRED
	Encrypted bool `json:"encrypted"`

	// REQUIRED
	Evals []interface{} `json:"Evals"`

	// REQUIRED; Anything
	Exif interface{} `json:"Exif"`

	// REQUIRED
	Filters []*Components1T23IneSchemasGetallimagetemplatesresPropertiesResultPropertiesTemplatesItemsPropertiesFiltersItems `json:"Filters"`

	// REQUIRED
	LoopCount int `json:"LoopCount"`

	// REQUIRED
	MaxAge int `json:"MaxAge"`

	// REQUIRED
	OuputQuality int `json:"OuputQuality"`

	// REQUIRED; Dictionary of
	OutputExtra map[string]string `json:"OutputExtra"`

	// REQUIRED
	OutputFormat string `json:"OutputFormat"`

	// REQUIRED; Anything
	OutputVideo interface{} `json:"OutputVideo"`

	// REQUIRED
	Parameters []string `json:"Parameters"`

	// REQUIRED
	Persistence string `json:"Persistence"`

	// REQUIRED
	QualityMode string `json:"QualityMode"`

	// REQUIRED
	QualityStr string `json:"QualityStr"`

	// REQUIRED
	ReqDeadline string `json:"ReqDeadline"`

	// REQUIRED
	ServiceID string `json:"ServiceID"`

	// REQUIRED; Anything
	Snapshot interface{} `json:"Snapshot"`

	// REQUIRED
	Sync bool `json:"Sync"`

	// REQUIRED
	TemplateName string `json:"TemplateName"`

	// REQUIRED
	TemplateType string `json:"TemplateType"`

	// REQUIRED
	Timeout int `json:"Timeout"`

	// REQUIRED
	Updater string `json:"Updater"`

	// REQUIRED
	Usage string `json:"Usage"`

	// REQUIRED
	WithSig bool `json:"WithSig"`
}

type GetAuditEntrysCountQuery struct {

	// 任务 ID，您可通过调用查询所有审核任务 [https://www.volcengine.com/docs/508/1160409]获取所需的任务 ID。
	TaskID string `json:"TaskId,omitempty" query:"TaskId"`
}

type GetAuditEntrysCountRes struct {

	// REQUIRED
	ResponseMetadata *GetAuditEntrysCountResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetAuditEntrysCountResResult `json:"Result"`
}

type GetAuditEntrysCountResResponseMetadata struct {

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

type GetAuditEntrysCountResResult struct {

	// REQUIRED; 异常审核的总次数，即审核失败、建议不通过和建议复审的审核次数之和。
	AuditExceptionCount int `json:"AuditExceptionCount"`

	// REQUIRED; 审核失败的审核次数
	AuditFailCount int `json:"AuditFailCount"`

	// REQUIRED; 建议不通过的审核次数
	AuditNopassCount int `json:"AuditNopassCount"`

	// REQUIRED; 建议复审的审核次数
	AuditRecheckCount int `json:"AuditRecheckCount"`

	// REQUIRED; 审核成功的审核次数。
	AuditSuccessCount int `json:"AuditSuccessCount"`

	// REQUIRED; 该任务的审核总次数，为审核成功和审核失败的图片数量之和。
	AuditTotal int `json:"AuditTotal"`

	// REQUIRED; 审核异常的图片数量，即审核失败、建议不通过和建议复审的图片数量之和。
	ExceptionCount int `json:"ExceptionCount"`

	// REQUIRED; 审核失败的图片数量
	FailedCount int `json:"FailedCount"`

	// REQUIRED; 建议不通过的图片数量
	NopassCount int `json:"NopassCount"`

	// REQUIRED; 建议复审的图片数量
	RecheckCount int `json:"RecheckCount"`

	// REQUIRED; 审核成功的图片数量
	SuccessCount int `json:"SuccessCount"`

	// REQUIRED; 累计审核图片数量，为审核成功和审核失败的图片数量之和。
	Total int `json:"Total"`
}

type GetBatchProcessResultBody struct {

	// REQUIRED; 待批量处理的资源链接信息
	BatchingInfo []*GetBatchProcessResultBodyBatchingInfoItem `json:"BatchingInfo"`
}

type GetBatchProcessResultBodyBatchingInfoItem struct {

	// 批处理能力，取值如下所示：
	// * meta：获取资源元信息
	// * preload：源站图片预热 :::warning 如需批量预热源站图片，请提交工单 [https://console.volcengine.com/ticket/createTicketV2/?step=3&Service=rtc&FlowKey=NGnOHeWkbeCrEAkrNvjT]联系技术支持开启。
	// :::
	Action string `json:"Action,omitempty"`

	// 指定服务下待批处理资源的可访问 URL
	URL string `json:"Url,omitempty"`
}

type GetBatchProcessResultQuery struct {

	// REQUIRED; 指定同步批处理的服务 ID。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetBatchProcessResultRes struct {

	// REQUIRED
	ResponseMetadata *GetBatchProcessResultResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetBatchProcessResultResResult           `json:"Result,omitempty"`
}

type GetBatchProcessResultResResponseMetadata struct {

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

type GetBatchProcessResultResResult struct {

	// REQUIRED; 火车模式信息
	Data []*GetBatchProcessResultResResultDataItem `json:"Data"`
}

type GetBatchProcessResultResResultDataItem struct {

	// REQUIRED; 该资源使用的批处理能力
	Action string `json:"Action"`

	// REQUIRED; 资源 URL
	URL string `json:"Url"`

	// 该资源执行批处理操作时的错误描述
	Err string `json:"Err,omitempty"`

	// 该资源的文件大小，单位为 byte。
	Size int `json:"Size,omitempty"`

	// 访问该资源时返回的 HTTP 状态码 [https://developer.mozilla.org/en-US/docs/Web/HTTP/Status]
	StatusCode int `json:"StatusCode,omitempty"`
}

type GetBatchTaskInfoQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 异步任务 ID，传入 CreateBatchProcessTask [https://www.volcengine.com/docs/508/1185525] 获取的异步任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetBatchTaskInfoRes struct {

	// REQUIRED
	ResponseMetadata *GetBatchTaskInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetBatchTaskInfoResResult           `json:"Result,omitempty"`
}

type GetBatchTaskInfoResResponseMetadata struct {

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

type GetBatchTaskInfoResResult struct {

	// REQUIRED; 异步任务 ID
	TaskID string `json:"TaskId"`

	// 传入的回调地址
	Callback string `json:"Callback,omitempty"`

	// 传入的回调内容
	CallbackBody string `json:"CallbackBody,omitempty"`

	// 传入的回调内容格式
	CallbackBodyType string `json:"CallbackBodyType,omitempty"`

	// 错误码。仅当Status取值Failed时，有返回值。
	Code int `json:"Code,omitempty"`

	// 异步批处理任务详情
	Data []*GetBatchTaskInfoResResultDataItem `json:"Data,omitempty"`

	// 错误信息。仅当Status取值Failed时，有返回值。
	Err string `json:"Err,omitempty"`

	// 任务状态，取值如下所示：
	// * Running：进行中
	// * Pending：排队中
	// * Failed：失败
	// * Success：成功
	Status string `json:"Status,omitempty"`
}

type GetBatchTaskInfoResResultDataItem struct {

	// REQUIRED; 该资源使用的批处理能力，取值如下所示：
	// * meta：获取资源元信息
	// * preload：源站图片预热
	Action string `json:"Action"`

	// REQUIRED; 资源 URL
	URL string `json:"Url"`

	// 该资源执行批处理操作时的错误描述
	Err string `json:"Err,omitempty"`

	// 该资源对应的文件大小，单位为 byte。
	Size int `json:"Size,omitempty"`

	// 访问该资源时返回的 HTTP 状态码 [https://developer.mozilla.org/en-US/docs/Web/HTTP/Status]
	StatusCode int `json:"StatusCode,omitempty"`
}

type GetCVAnimeGenerateImageBody struct {

	// REQUIRED; 服务下绑定的域名，域名状态需正常可用。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取绑定的域名信息。
	// * 您也可以通过 OpenAPI 的方式获取域名，具体请参考获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]。
	Domain string `json:"Domain"`

	// REQUIRED; 模型操作action
	ModelAction string `json:"ModelAction"`

	// REQUIRED; 模型版本。
	ModelVersion string `json:"ModelVersion"`

	// REQUIRED; 输出结果。
	Outputs []string `json:"Outputs"`

	// REQUIRED; 请求的JSON数据。
	ReqJSON map[string]interface{} `json:"ReqJson"`

	// REQUIRED; 服务下创建的图片处理模板名称，指定后，将按照模板中的处理配置对豆包大模型生成的图片进行图片处理。
	// 您可在 veImageX 控制台的处理配置页面，参考新建模板 [https://www.volcengine.com/docs/508/8087]配置模板并获取模版名称，例如 tplv-f0****5k-test。
	Template string `json:"Template"`

	// 是否覆盖现有文件。
	Overwrite bool `json:"Overwrite,omitempty"`
}

type GetCVAnimeGenerateImageQuery struct {

	// REQUIRED; 指定存储结果图并计量计费的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetCVAnimeGenerateImageRes struct {

	// REQUIRED
	ResponseMetadata *GetCVAnimeGenerateImageResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetCVAnimeGenerateImageResResult `json:"Result,omitempty"`
}

type GetCVAnimeGenerateImageResResponseMetadata struct {

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

// GetCVAnimeGenerateImageResResult - 视请求的接口而定
type GetCVAnimeGenerateImageResResult struct {

	// REQUIRED; 存储URI列表。
	ImageUrls []string `json:"ImageUrls"`

	// REQUIRED; 响应的JSON数据。
	RespJSON map[string]interface{} `json:"RespJson"`

	// REQUIRED; 存储URI。
	StoreUris []string `json:"StoreUris"`
}

type GetCVImageGenerateResultBody struct {

	// REQUIRED; 服务下绑定的域名，域名状态需正常可用。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取绑定的域名信息。
	// * 您也可以通过 OpenAPI 的方式获取域名，具体请参考获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]。
	Domain string `json:"Domain"`

	// REQUIRED; 模型操作action
	ModelAction string `json:"ModelAction"`

	// REQUIRED; 模型版本。
	ModelVersion string `json:"ModelVersion"`

	// REQUIRED; 输出结果。
	Outputs []string `json:"Outputs"`

	// REQUIRED; 请求的JSON字符串。
	ReqJSON map[string]interface{} `json:"ReqJson"`

	// REQUIRED; 服务下创建的图片处理模板名称，指定后，将按照模板中的处理配置对豆包大模型生成的图片进行图片处理。
	// 您可在 veImageX 控制台的处理配置页面，参考新建模板 [https://www.volcengine.com/docs/508/8087]配置模板并获取模版名称，例如 tplv-f0****5k-test。
	Template string `json:"Template"`

	// 图片URL
	ImageURL string `json:"ImageUrl,omitempty"`

	// 是否覆盖现有文件。
	Overwrite bool `json:"Overwrite,omitempty"`
}

type GetCVImageGenerateResultQuery struct {

	// REQUIRED; 指定存储结果图并计量计费的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetCVImageGenerateResultRes struct {

	// REQUIRED
	ResponseMetadata *GetCVImageGenerateResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetCVImageGenerateResultResResult `json:"Result,omitempty"`
}

type GetCVImageGenerateResultResResponseMetadata struct {

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

// GetCVImageGenerateResultResResult - 视请求的接口而定
type GetCVImageGenerateResultResResult struct {

	// REQUIRED; 图片URL列表。
	ImageUrls []string `json:"ImageUrls"`

	// REQUIRED; 响应的JSON内容。
	RespJSON map[string]interface{} `json:"RespJson"`

	// REQUIRED; 存储URI。
	StoreUris []string `json:"StoreUris"`
}

type GetCVImageGenerateTaskBody struct {

	// REQUIRED; 操作模型。
	ModelAction string `json:"ModelAction"`

	// REQUIRED; 模型版本。
	ModelVersion string `json:"ModelVersion"`

	// REQUIRED; 请求的JSON数据。
	ReqJSON map[string]interface{} `json:"ReqJson"`

	// REQUIRED; 指定文生图异步任务的任务 ID。
	TaskID string `json:"TaskId"`
}

type GetCVImageGenerateTaskQuery struct {

	// REQUIRED; 指定要查询的服务 ID。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetCVImageGenerateTaskRes struct {

	// REQUIRED
	ResponseMetadata *GetCVImageGenerateTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetCVImageGenerateTaskResResult `json:"Result,omitempty"`
}

type GetCVImageGenerateTaskResResponseMetadata struct {

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

// GetCVImageGenerateTaskResResult - 视请求的接口而定
type GetCVImageGenerateTaskResResult struct {

	// REQUIRED; 最终上传至 veImageX 服务的结果图访问 URL，数量为 1。
	ImageUrls []string `json:"ImageUrls"`

	// REQUIRED; 响应的 JSON 数据。
	RespJSON map[string]interface{} `json:"RespJson"`

	// REQUIRED; 任务状态，取值如下所示：
	// * in_queue：任务已提交
	// * generating：任务处理中
	// * done：任务处理完成
	// * not_found：任务未找到，可能原因是无此任务或任务已过期（12小时）
	Status string `json:"Status"`

	// REQUIRED; 存储URI。
	StoreUris []string `json:"StoreUris"`

	// REQUIRED; 任务 ID
	TaskID string `json:"TaskId"`
}

type GetCVTextGenerateImageBody struct {

	// REQUIRED; 服务下绑定的域名，域名状态需正常可用。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取绑定的域名信息。
	// * 您也可以通过 OpenAPI 的方式获取域名，具体请参考获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]。
	Domain string `json:"Domain"`

	// REQUIRED; 操作模型action
	ModelAction string `json:"ModelAction"`

	// REQUIRED; 模型版本。
	ModelVersion string `json:"ModelVersion"`

	// REQUIRED; 输出结果。
	Outputs []string `json:"Outputs"`

	// REQUIRED; 请求的JSON数据。
	ReqJSON map[string]interface{} `json:"ReqJson"`

	// REQUIRED; 服务下创建的图片处理模板配置（不带~），指定后，将按照模板中的处理配置对豆包大模型生成的图片进行图片处理。 您可在 veImageX 控制台的处理配置页面，参考新建模板 [https://www.volcengine.com/docs/508/8087]配置模板并获取模版信息，例如tplv-f0****5k-test.image。
	Template string `json:"Template"`

	// 覆盖现有内容。
	Overwrite bool `json:"Overwrite,omitempty"`
}

type GetCVTextGenerateImageQuery struct {

	// REQUIRED; 指定存储结果图并计量计费的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetCVTextGenerateImageRes struct {

	// REQUIRED
	ResponseMetadata *GetCVTextGenerateImageResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetCVTextGenerateImageResResult `json:"Result,omitempty"`
}

type GetCVTextGenerateImageResResponseMetadata struct {

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

// GetCVTextGenerateImageResResult - 视请求的接口而定
type GetCVTextGenerateImageResResult struct {

	// REQUIRED; 存储URI列表。
	ImageUrls []string `json:"ImageUrls"`

	// REQUIRED; 响应的JSON字符串。
	RespJSON map[string]interface{} `json:"RespJson"`

	// REQUIRED; 存储URI。
	StoreUris []string `json:"StoreUris"`
}

type GetCertInfoQuery struct {

	// REQUIRED; 证书 ID，您可以通过调用获取账号下全部证书 [https://www.volcengine.com/docs/508/66017]获取账号下所有证书信息。
	CertID string `json:"CertID" query:"CertID"`
}

type GetCertInfoRes struct {

	// REQUIRED
	ResponseMetadata *GetCertInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetCertInfoResResult           `json:"Result,omitempty"`
}

type GetCertInfoResResponseMetadata struct {

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

type GetCertInfoResResult struct {

	// REQUIRED; 证书id
	CertID string `json:"CertID"`

	// REQUIRED; 证书名
	CertName string `json:"CertName"`

	// REQUIRED; 允许https使用的域名
	CommonName string `json:"CommonName"`

	// REQUIRED; 创建时间戳
	CreateTime float64 `json:"CreateTime"`

	// REQUIRED; 签发机构
	Issuer string `json:"Issuer"`

	// REQUIRED; 过期时间
	NotAfter float64 `json:"NotAfter"`

	// REQUIRED
	San []string `json:"San"`
}

type GetComprehensiveEnhanceImageBody struct {

	// REQUIRED; 待增强图片的存储 URI 或访问 URL（公网可访问）。您可在控制台资源管理获取图片的存储 URI [https://www.volcengine.com/docs/508/1205057]以及访问 URL [https://www.volcengine.com/docs/508/1205054]。
	ImageURI string `json:"ImageUri"`

	// REQUIRED; 优化策略，取值如下所示：
	// * 0：通用画质提升
	// * 1：显著画质提升 :::tip 推荐优先使用通用方案，显著画质提升方案画质分提高 10% 左右，但体积会有一定浮动提升。以上幅度变化基于测试集总结，具体以使用场景为准。 :::
	Mode int `json:"Mode"`

	// REQUIRED; 服务 ID。
	ServiceID string `json:"ServiceId"`

	// 去压缩失真强度，取值范围为[0,1]。取值为0时表示不处理，取值越大去压缩失真强度越大。
	ArStrength float64 `json:"ArStrength,omitempty"`

	// EnableConfig 取值为 true 时，为必填。
	// 亮度，取值范围为[90,100]。取值越小，亮度提升越明显。
	Brightness int `json:"Brightness,omitempty"`

	// 去模糊强度，取值范围为[0,1]。取值为0时表示不处理，取值越大去模糊强度越大。
	DeblurStrength float64 `json:"DeblurStrength,omitempty"`

	// 降噪强度，取值范围为[0,1]。取值为0时表示不降噪，取值越大降噪强度越大。
	DenoiseStrength float64 `json:"DenoiseStrength,omitempty"`

	// 下采样模式，取值如下所示：
	// * 0: 仅缩小，图片大于设置的“宽/高”时，将缩小图片
	// * 1: 仅放大，图片小于设置的“宽/高”时，将放大图片
	// * 2: 既缩小又放大，即按照自定义“宽/高”输出结果图，该操作可能导致图片变形
	DownsampleMode int `json:"DownsampleMode,omitempty"`

	// 下采样输出图片高度，图片将适配对应大小。单位为 px。
	DownsampleOutHeight int `json:"DownsampleOutHeight,omitempty"`

	// 下采样输出图片宽度，图片将适配对应大小。单位为 px。
	DownsampleOutWidth int `json:"DownsampleOutWidth,omitempty"`

	// 是否启用高级配置，取值如下所示：
	// * true：开启。开启后，下述高级配置才会生效。
	// * false：（默认）关闭。 :::tip 适用于 8000 x 8000 以分辨率图像的画质增强。 :::
	EnableConfig bool `json:"EnableConfig,omitempty"`

	// 是否开启下采样，即图片在执行增强效果的同时可自定义输出图片分辨率大小。取值如下所示：
	// * true：开启。仅当开启后，下述下采样配置才会生效。
	// * false：（默认）关闭。 :::tip 适用于 8000 x 8000 以内分辨率图像的画质增强。 :::
	EnableDownsample bool `json:"EnableDownsample,omitempty"`

	// 是否开启超分配置，仅满足图像输入边界的图像执行超分处理。取值如下所示：
	// * true：开启。仅当开启后，下述超分配置才会生效。
	// * false：（默认）关闭。 :::tip 适用于 8000 x 8000 以内分辨率图像的画质增强。 :::
	EnableSuperResolution bool `json:"EnableSuperResolution,omitempty"`

	// 默认值为false
	EnableTextEnhance bool `json:"EnableTextEnhance,omitempty"`

	// 执行超分处理的长边范围最大值，仅当满足图像边界输入的图像执行超分处理。单位为 px。
	LongMax int `json:"LongMax,omitempty"`

	// 执行超分处理的长边范围最小值，仅当满足图像边界输入的图像执行超分处理。单位为 px。
	LongMin int `json:"LongMin,omitempty"`

	// 超分倍率，仅支持 2 倍和 4 倍。 :::tip 4 倍超分辨率只适用于 4000 x 4000 以内分辨率图像的画质增强。 :::
	Multiple int `json:"Multiple,omitempty"`

	// 饱和度，取值范围为[0,2]。取值大于 1 表示提升饱和度，取值小于 1 表示降低饱和度。
	Saturation float64 `json:"Saturation,omitempty"`

	// EnableSuperResolution 取值为 true 时，为必填。
	// 执行超分处理的短边范围最大值，仅当满足图像边界输入的图像执行超分处理。单位为 px。
	ShortMax int `json:"ShortMax,omitempty"`

	// EnableSuperResolution 取值为 true 时，为必填。
	// 执行超分处理的短边范围最小值，仅当满足图像边界输入的图像执行超分处理。单位为 px。
	ShortMin int `json:"ShortMin,omitempty"`

	// 取值范围[0,1]
	TextEnhanceStrength float64 `json:"TextEnhanceStrength,omitempty"`
}

type GetComprehensiveEnhanceImageRes struct {

	// REQUIRED
	ResponseMetadata *GetComprehensiveEnhanceImageResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetComprehensiveEnhanceImageResResult `json:"Result,omitempty"`
}

type GetComprehensiveEnhanceImageResResponseMetadata struct {

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

// GetComprehensiveEnhanceImageResResult - 视请求的接口而定
type GetComprehensiveEnhanceImageResResult struct {

	// REQUIRED; 结果图 URI。您可使用结果图 URI（即ResUri）拼接完整访问 URL [https://www.volcengine.com/docs/508/105405#%E9%A2%84%E8%A7%88%E8%BF%94%E5%9B%9E%E7%9A%84%E7%BB%93%E6%9E%9C%E5%9B%BE]后，在浏览器查看图像增强效果。
	ResURI string `json:"ResUri"`
}

type GetCompressTaskInfoQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 异步任务 ID
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetCompressTaskInfoRes struct {

	// REQUIRED
	ResponseMetadata *GetCompressTaskInfoResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetCompressTaskInfoResResult `json:"Result,omitempty"`
}

type GetCompressTaskInfoResResponseMetadata struct {

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

// GetCompressTaskInfoResResult - 视请求的接口而定
type GetCompressTaskInfoResResult struct {

	// REQUIRED; 压缩文件包 StoreUri
	OutputURI string `json:"OutputUri"`

	// REQUIRED; 异步任务状态，取值如下所示：
	// * Pending：排队中
	//
	//
	// * Running：进行中
	//
	//
	// * Failed：执行失败
	//
	//
	// * Success：执行成功
	Status string `json:"Status"`

	// REQUIRED; 异步任务 ID
	TaskID string `json:"TaskId"`

	// 错误码。仅当Status为Failed时，该值不为 0。
	ErrCode int `json:"ErrCode,omitempty"`

	// 错误信息。仅当Status为Failed时，该值不为空。
	ErrMsg string `json:"ErrMsg,omitempty"`

	// 任务处理进度，即已处理的文件数量。
	ProcessCount int    `json:"ProcessCount,omitempty"`
	ResURI       string `json:"ResUri,omitempty"`
}

type GetDedupTaskStatusQuery struct {

	// REQUIRED; 任务 ID，您可以通过调用使用图片去重获取结果值 [https://www.volcengine.com/docs/508/138658]接口获取异步去重TaskId返回值。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetDedupTaskStatusRes struct {

	// REQUIRED
	ResponseMetadata *GetDedupTaskStatusResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetDedupTaskStatusResResult `json:"Result,omitempty"`
}

type GetDedupTaskStatusResResponseMetadata struct {

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

// GetDedupTaskStatusResResult - 视请求的接口而定
type GetDedupTaskStatusResResult struct {

	// REQUIRED; 分组结果，若输入 2 个以上原图时，将按组返回内容相同的图片，每组的图片个数非固定值。
	Groups map[string][]string `json:"Groups"`

	// REQUIRED; 图片评分，仅在两张图片对比才返回 Score 字段。
	Score string `json:"Score"`

	// REQUIRED; 异步任务状态，取值如下所示：
	// * 0：任务进行中
	// * 1：任务执行成功
	// * 2：任务执行失败
	Status int `json:"Status"`

	// REQUIRED; 任务 ID。
	TaskID string `json:"TaskId"`
}

type GetDenoisingImageBody struct {

	// REQUIRED; 是否支持降级，即发生错误时返回原图地址。 传入 StoreUri 则返回 StoreUri，传入 ImageUrl 则返回 ImageUrl。 取值如下所示：
	// * true：支持降级
	// * false：不支持降级
	CanDemotion bool `json:"CanDemotion"`

	// REQUIRED; 降噪强度，取值范围为[0,1]。取值为0时表示不降噪，取值越大降噪强度越大。
	Intensity float64 `json:"Intensity"`

	// REQUIRED; 输出格式，支持格式有：png、jpeg、webp。
	OutFormat string `json:"OutFormat"`

	// 公网可访问的待降噪的原图 URL。
	ImageURL string `json:"ImageUrl,omitempty"`

	// 待降噪的原图 URI。 若同时传入 StoreUri 和 ImageUrl，仅取值 StoreUri，ImageUrl 将会被忽略。
	StoreURI string `json:"StoreUri,omitempty"`
}

type GetDenoisingImageQuery struct {

	// REQUIRED; 服务 ID。
	// 您可以在 veImageX 控制台 服务管理页面，在创建好的图片服务中获取服务 ID。 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetDenoisingImageRes struct {

	// REQUIRED
	ResponseMetadata *GetDenoisingImageResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetDenoisingImageResResult `json:"Result,omitempty"`
}

type GetDenoisingImageResResponseMetadata struct {

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

// GetDenoisingImageResResult - 视请求的接口而定
type GetDenoisingImageResResult struct {

	// REQUIRED; 是否发生降级，取值如下所示：
	// * true：降级
	// * false：未降级
	Demotion bool `json:"Demotion"`

	// REQUIRED; 降噪后的图片地址，根据输入时的地址决定返回值。
	ResURI string `json:"ResUri"`
}

type GetDomainConfigQuery struct {

	// REQUIRED; 域名，您可以通过调用 GetServiceDomains [https://www.volcengine.com/docs/508/9379] 获取当前服务下所有域名。
	DomainName string `json:"DomainName" query:"DomainName"`

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetDomainConfigRes struct {

	// REQUIRED
	ResponseMetadata *GetDomainConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetDomainConfigResResult `json:"Result,omitempty"`
}

type GetDomainConfigResResponseMetadata struct {

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

// GetDomainConfigResResult - 视请求的接口而定
type GetDomainConfigResResult struct {

	// REQUIRED; 访问控制配置
	AccessControl *GetDomainConfigResResultAccessControl `json:"access_control"`

	// REQUIRED; 是否开启自适应格式，取值如下所示：
	// * true：开启自适应
	// * false：关闭自适应
	Adaptfmt bool `json:"adaptfmt"`

	// REQUIRED; 是否开启集智瘦身，取值如下所示：
	// * true：开启集智瘦身
	// * false：关闭集智瘦身
	DoSlim bool `json:"do_slim"`

	// REQUIRED; 域名
	Domain string `json:"domain"`

	// REQUIRED; 是否开启全球加速，取值如下所示：
	// * true：开启全球加速
	// * false：关闭全球加速
	GlobalAcceleration bool `json:"global_acceleration"`

	// REQUIRED; HTTPS 配置
	HTTPSConfig *GetDomainConfigResResultHTTPSConfig `json:"https_config"`

	// REQUIRED; 域名锁定状态
	LockStatus *GetDomainConfigResResultLockStatus `json:"lock_status"`

	// REQUIRED; HTTP Header 配置
	RespHdrs []*GetDomainConfigResResultRespHdrsItem `json:"resp_hdrs"`

	// REQUIRED; 域名状态
	Status string `json:"status"`

	// 页面优化设置
	PageOptimization *GetDomainConfigResResultPageOptimization `json:"page_optimization,omitempty"`
}

// GetDomainConfigResResultAccessControl - 访问控制配置
type GetDomainConfigResResultAccessControl struct {

	// REQUIRED; IP 访问控制配置
	IPAuth *GetDomainConfigResResultAccessControlIPAuth `json:"ip_auth"`

	// REQUIRED; Referer 防盗链配置
	ReferLink *GetDomainConfigResResultAccessControlReferLink `json:"refer_link"`

	// REQUIRED; 远程鉴权设置
	RemoteAuth *GetDomainConfigResResultAccessControlRemoteAuth `json:"remote_auth"`

	// REQUIRED; URL 鉴权配置
	URLAuth *GetDomainConfigResResultAccessControlURLAuth `json:"url_auth"`

	// REQUIRED; UA 访问控制配置
	UaList *GetDomainConfigResResultAccessControlUaList `json:"ua_list"`
}

// GetDomainConfigResResultAccessControlIPAuth - IP 访问控制配置
type GetDomainConfigResResultAccessControlIPAuth struct {

	// REQUIRED; 是否开启 IP 访问控制，取值如下所示：
	// * true：开启 IP 访问控制
	// * false：关闭 IP 访问控制
	Enabled bool `json:"enabled"`

	// REQUIRED; 是否是白名单模式，取值如下所示：
	// * true：白名单模式
	// * false：黑名单模式
	IsWhiteMode bool `json:"is_white_mode"`

	// REQUIRED; IP 列表
	Values []string `json:"values"`
}

// GetDomainConfigResResultAccessControlReferLink - Referer 防盗链配置
type GetDomainConfigResResultAccessControlReferLink struct {

	// REQUIRED; 是否允许空 Refer，取值如下所示：
	// * true：允许空 Refer
	// * false：不允许空 Refer
	AllowEmptyRefer bool `json:"allow_empty_refer"`

	// REQUIRED; 是否开启 Referer 防盗链，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"enabled"`

	// REQUIRED; Referers 列表在匹配时是否是大小写敏感的。取值如下所示：
	// * true: 表示大小写不敏感。
	// * false: 表示大小写敏感。
	IgnoreCase bool `json:"ignore_case"`

	// REQUIRED; Referers 列表的 Referer 头部值是否必须以 HTTP 或者 HTTPS 开头。取值如下所示：
	// * true: 表示不以 HTTP 或者 HTTPS 开头的 Referer 头部值是合法的。
	// * false: 表示不以 HTTP 或者 HTTPS 开头 Referer 头部值是非法的。
	IgnoreScheme bool `json:"ignore_scheme"`

	// REQUIRED; 是否选择白名单，取值如下所示：
	// * true：选择白名单
	// * false：不选择白名单
	IsWhiteMode bool `json:"is_white_mode"`

	// REQUIRED; 正则表达式规则列表
	RegexValues []string `json:"regex_values"`

	// REQUIRED; 根据是否为白名单，为对应的白/黑名单的值。
	Values []string `json:"values"`
}

// GetDomainConfigResResultAccessControlRemoteAuth - 远程鉴权设置
type GetDomainConfigResResultAccessControlRemoteAuth struct {

	// REQUIRED; 鉴权请求头设置
	AuthRequestHeader *GetDomainConfigResResultAccessControlRemoteAuthRequestHeader `json:"auth_request_header"`

	// REQUIRED; 鉴权请求参数设置
	AuthRequestQuery *GetDomainConfigResResultAccessControlRemoteAuthRequestQuery `json:"auth_request_query"`

	// REQUIRED; 鉴权响应设置
	AuthResponse *GetDomainConfigResResultAccessControlRemoteAuthResponse `json:"auth_response"`

	// REQUIRED; 鉴权服务器设置
	AuthServer *GetDomainConfigResResultAccessControlRemoteAuthServer `json:"auth_server"`

	// REQUIRED; 是否开启远程鉴权，取值如下所示：
	// * true：是
	// * false：否
	Enabled bool `json:"enabled"`

	// REQUIRED; 生效对象
	MatchRule []*GetDomainConfigResResultAccessControlRemoteAuthMatchRuleItem `json:"match_rule"`
}

// GetDomainConfigResResultAccessControlRemoteAuthAuthResponseResponse - 响应设置
type GetDomainConfigResResultAccessControlRemoteAuthAuthResponseResponse struct {

	// REQUIRED; 鉴权失败时，响应用户的状态码
	FailCode string `json:"fail_code"`
}

type GetDomainConfigResResultAccessControlRemoteAuthMatchRuleItem struct {

	// REQUIRED; 匹配方式，取值如下所示：
	// * match：object 匹配 Value。
	// * not_match：object 不匹配 Value。
	MatchOperator string `json:"match_operator"`

	// REQUIRED; 对哪些对象类型进行规则匹配，取值如下所示：
	// * filetype：表示特定后缀的文件。
	// * directory：表示特定文件目录下的所有文件。
	// * path：表示特定的文件。
	Object string `json:"object"`

	// REQUIRED; Object 对应的具体对象
	Value string `json:"value"`
}

// GetDomainConfigResResultAccessControlRemoteAuthRequestHeader - 鉴权请求头设置
type GetDomainConfigResResultAccessControlRemoteAuthRequestHeader struct {

	// REQUIRED; 鉴权请求头是否包含用户请求头，取值如下所示：
	// * exclude：表示鉴权请求头中不包含任何用户请求头。
	// * include：表示鉴权请求头中包含所有用户请求头。
	// * includePart：表示鉴权请求头包含指定的用户请求头。
	Action string `json:"action"`

	// REQUIRED; 额外请求头
	Header []*GetDomainConfigResResultAccessControlRemoteAuthRequestHeaderItem `json:"header"`

	// REQUIRED; 鉴权请求中 HOST 头部的值，固定为 default，即 HOST 头部的值与您的加速域名相同。
	Host string `json:"host"`

	// REQUIRED; Action 参数所对应的参数值
	Value string `json:"value"`
}

type GetDomainConfigResResultAccessControlRemoteAuthRequestHeaderItem struct {

	// REQUIRED; 您设置的请求头
	Key string `json:"key"`

	// REQUIRED; 请求头的值
	Value string `json:"value"`

	// REQUIRED; 请求头的类型
	ValueType string `json:"value_type"`
}

// GetDomainConfigResResultAccessControlRemoteAuthRequestQuery - 鉴权请求参数设置
type GetDomainConfigResResultAccessControlRemoteAuthRequestQuery struct {

	// REQUIRED; 鉴权请求是否包含用户请求 URL 中的查询参数。取值如下所示：
	// * exclude：表示鉴权请求不包含任何查询参数。
	// * include：表示鉴权请求包含所有查询参数。
	// * includePart：表示鉴权请求包含指定的查询参数。
	Action string `json:"action"`

	// REQUIRED; 额外参数
	Query []*GetDomainConfigResResultAccessControlRemoteAuthRequestQueryItem `json:"query"`

	// REQUIRED; Action 参数所对应的参数值
	Value string `json:"value"`
}

type GetDomainConfigResResultAccessControlRemoteAuthRequestQueryItem struct {

	// REQUIRED; 您设置的鉴权请求参数
	Key string `json:"key"`

	// REQUIRED; 鉴权请求参数的值
	Value string `json:"value"`

	// REQUIRED; 您在 Key 中设置的鉴权请求参数的类型
	ValueType string `json:"value_type"`
}

// GetDomainConfigResResultAccessControlRemoteAuthResponse - 鉴权响应设置
type GetDomainConfigResResultAccessControlRemoteAuthResponse struct {

	// REQUIRED; 鉴权结果缓存设置
	AuthResultCache *GetDomainConfigResResultAccessControlRemoteAuthResponseAuthResultCache `json:"auth_result_cache"`

	// REQUIRED; 鉴权服务器状态码设置
	AuthServerStatusCode *GetDomainConfigResResultAccessControlRemoteAuthResponseAuthServerStatusCode `json:"auth_server_status_code"`

	// REQUIRED; 鉴权服务超时时间
	AuthServerTimeout *GetDomainConfigResResultAccessControlRemoteAuthResponseAuthServerTimeout `json:"auth_server_timeout"`

	// REQUIRED; 响应设置
	Response *GetDomainConfigResResultAccessControlRemoteAuthAuthResponseResponse `json:"response"`
}

// GetDomainConfigResResultAccessControlRemoteAuthResponseAuthResultCache - 鉴权结果缓存设置
type GetDomainConfigResResultAccessControlRemoteAuthResponseAuthResultCache struct {

	// REQUIRED; 是否缓存鉴权状态码，取值如下所示：
	// * nocache：veImageX 不缓存鉴权状态码。
	// * cache：veImageX 缓存鉴权状态码。
	Action string `json:"action"`

	// REQUIRED; 缓存 key 指定了用于区分不同请求 URI 的查询参数
	CacheKey []string `json:"cache_key"`

	// REQUIRED; 鉴权状态码的缓存时间，单位是秒
	TTL int `json:"ttl"`
}

// GetDomainConfigResResultAccessControlRemoteAuthResponseAuthServerStatusCode - 鉴权服务器状态码设置
type GetDomainConfigResResultAccessControlRemoteAuthResponseAuthServerStatusCode struct {

	// REQUIRED; 如果鉴权状态码既不是 FailCode，又不是 SuccessCode 时，处理鉴权请求的方式
	DefaultAction string `json:"default_action"`

	// REQUIRED; 鉴权失败时的鉴权状态码
	FailCode string `json:"fail_code"`

	// REQUIRED; 鉴权成功时的鉴权状态码
	SuccessCode string `json:"success_code"`
}

// GetDomainConfigResResultAccessControlRemoteAuthResponseAuthServerTimeout - 鉴权服务超时时间
type GetDomainConfigResResultAccessControlRemoteAuthResponseAuthServerTimeout struct {

	// REQUIRED; 鉴权超时后处理鉴权请求的策略，取值如下所示：
	// * reject：veImageX 认为鉴权失败。
	// * pass：veImageX 认为鉴权成功。
	Action string `json:"action"`

	// REQUIRED; 鉴权超时的时间，单位是毫秒。
	Time int `json:"time"`
}

// GetDomainConfigResResultAccessControlRemoteAuthServer - 鉴权服务器设置
type GetDomainConfigResResultAccessControlRemoteAuthServer struct {

	// REQUIRED; 鉴权服务器的主地址
	Address string `json:"address"`

	// REQUIRED; 鉴权服务器的备地址
	BackupAddress string `json:"backup_address"`

	// REQUIRED; 鉴权请求的路径，取值如下所示：
	// * constant：表示鉴权请求中的路径与用户请求中的路径相同。
	// * variable：表示您需要在 pathValue 参数中指定一个鉴权请求中的路径。
	PathType string `json:"path_type"`

	// REQUIRED; 一个鉴权请求的路径
	PathValue string `json:"path_value"`

	// REQUIRED; 在发送鉴权请求时所使用的请求方法，取值如下所示：
	// * default：鉴权请求所使用的方法与用户的请求相同。
	// * get：鉴权请求使用 GET 方法。
	// * post：鉴权请求使用 POST 方法。
	// * head：鉴权请求使用 HEAD 方法。
	RequestMethod string `json:"request_method"`
}

// GetDomainConfigResResultAccessControlURLAuth - URL 鉴权配置
type GetDomainConfigResResultAccessControlURLAuth struct {

	// REQUIRED; 是否开启 URL 鉴权 ，取值如下所示：
	// * true：开启 URL 鉴权
	// * false：关闭 URL 鉴权
	Enabled bool `json:"enabled"`

	// REQUIRED; A 鉴权配置
	TypeA *GetDomainConfigResResultAccessControlURLAuthTypeA `json:"type_a"`

	// REQUIRED; B 鉴权配置
	TypeB *GetDomainConfigResResultAccessControlURLAuthTypeB `json:"type_b"`

	// REQUIRED; C 鉴权配置
	TypeC *GetDomainConfigResResultAccessControlURLAuthTypeC `json:"type_c"`

	// REQUIRED; D 鉴权配置
	TypeD *GetDomainConfigResResultAccessControlURLAuthTypeD `json:"type_d"`
}

// GetDomainConfigResResultAccessControlURLAuthTypeA - A 鉴权配置
type GetDomainConfigResResultAccessControlURLAuthTypeA struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`

	// REQUIRED; 签名参数名
	SignParam string `json:"sign_param"`
}

// GetDomainConfigResResultAccessControlURLAuthTypeB - B 鉴权配置
type GetDomainConfigResResultAccessControlURLAuthTypeB struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`
}

// GetDomainConfigResResultAccessControlURLAuthTypeC - C 鉴权配置
type GetDomainConfigResResultAccessControlURLAuthTypeC struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`
}

// GetDomainConfigResResultAccessControlURLAuthTypeD - D 鉴权配置
type GetDomainConfigResResultAccessControlURLAuthTypeD struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`

	// REQUIRED; md5hash 参数名
	SignParam string `json:"sign_param"`

	// REQUIRED; 时间戳格式，取值如下所示：
	// * decimal：十进制（Unix 时间戳）
	// * heximal：十六进制（Unix 时间戳）
	TimeFormat string `json:"time_format"`

	// REQUIRED; TimeStamp 参数名
	TimeParam string `json:"time_param"`
}

// GetDomainConfigResResultAccessControlUaList - UA 访问控制配置
type GetDomainConfigResResultAccessControlUaList struct {

	// REQUIRED; 是否开启 UA 访问控制，取值如下所示：
	// * true：开启 UA 访问控制
	// * false：关闭 UA 访问控制
	Enabled bool `json:"enabled"`

	// REQUIRED; 是否是白名单模式，取值如下所示：
	// * true：白名单模式
	// * false：黑名单模式
	IsWhiteMode bool `json:"is_white_mode"`

	// REQUIRED; UA 列表
	Values []string `json:"values"`
}

// GetDomainConfigResResultHTTPSConfig - HTTPS 配置
type GetDomainConfigResResultHTTPSConfig struct {

	// REQUIRED
	Cert *GetDomainConfigResResultHTTPSConfigCert `json:"cert"`

	// REQUIRED; 证书 ID，若enable_https为true，则为必选。
	CertID string `json:"cert_id"`

	// REQUIRED; 是否开启强制跳转，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableForceRedirect bool `json:"enable_force_redirect"`

	// REQUIRED; 是否开启 HTTP2，取值如下所示：
	// * true：开启 HTTP2
	// * false：关闭 HTTP2
	EnableHTTP2 bool `json:"enable_http2"`

	// REQUIRED; 是否开启 HTTPS，取值如下所示：
	// * true：开启 HTTPS
	// * false：关闭 HTTPS
	EnableHTTPS bool `json:"enable_https"`

	// REQUIRED; 是否启用 OCSP 装订配置，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableOcsp bool `json:"enable_ocsp"`

	// REQUIRED; 是否强制使用 HTTPS，取值如下所示：
	// * true：强制 HTTPS
	// * false：不强制 HTTPS
	ForceHTTPS bool `json:"force_https"`

	// REQUIRED; 强制跳转状态码，取值如下所示：
	// * 301：返回给用户 301 状态码进行重定向。
	// * 302：返回给用户 302 状态码进行重定向。
	ForceRedirectCode string `json:"force_redirect_code"`

	// REQUIRED; 强制跳转类型，取值如下所示：
	// * http2https：HTTP 到 HTTPS
	// * https2http：HTTPS 到 HTTP
	ForceRedirectType string `json:"force_redirect_type"`

	// REQUIRED; HSTS 配置
	Hsts *GetDomainConfigResResultHTTPSConfigHsts `json:"hsts"`

	// REQUIRED; 支持的 tls 版本。取值如下所示：
	// * tlsv1.0
	// * tlsv1.1
	// * tlsv1.2
	// * tlsv1.3
	TLSVersions []string `json:"tls_versions"`
}

type GetDomainConfigResResultHTTPSConfigCert struct {

	// REQUIRED; 证书通用名
	CommonName string `json:"common_name"`

	// REQUIRED; 证书id
	ID string `json:"id"`

	// REQUIRED; 证书过期时间，unix时间戳，单位秒
	NotAfter int `json:"not_after"`
}

// GetDomainConfigResResultHTTPSConfigHsts - HSTS 配置
type GetDomainConfigResResultHTTPSConfigHsts struct {

	// 是否启用 HSTS 配置，取值如下所示：
	// * true：启用
	// * false：关闭
	Enabled bool `json:"enabled,omitempty"`

	// HSTS 配置是否也应用于加速域名的子域名。取值如下所示：
	// * include：应用于子域名站点。
	// * exclude：不应用于子域名站点。
	Subdomain string `json:"subdomain,omitempty"`

	// Strict-Transport-Security 响应头在浏览器中的缓存过期时间，单位是秒。取值范围是 0 - 31,536,000。31,536,000 秒表示 365 天。如果该参数值为 0，其效果等同于禁用 HSTS 设置。
	TTL int `json:"ttl,omitempty"`
}

// GetDomainConfigResResultLockStatus - 域名锁定状态
type GetDomainConfigResResultLockStatus struct {

	// REQUIRED; 智能压缩是否锁定，取值如下所示：
	// * true：是
	// * false：否
	CompressionLocked bool `json:"compression_locked"`

	// REQUIRED; 整个域名是否锁定，取值如下所示：
	// * true：是
	// * false：否
	DomainLocked bool `json:"domain_locked"`

	// REQUIRED; IP 访问限制是否锁定，取值如下所示：
	// * true：是
	// * false：否
	IPAccessRuleLocked bool `json:"ip_access_rule_locked"`

	// REQUIRED; 锁定原因
	Reason string `json:"reason"`

	// REQUIRED; Referer 防盗链配置是否锁定，取值如下所示：
	// * true：是
	// * false：否
	RefererAccessRuleLocked bool `json:"referer_access_rule_locked"`

	// REQUIRED; 远程鉴权是否锁定，取值如下所示：
	// * true：是
	// * false：否
	RemoteAuthLocked bool `json:"remote_auth_locked"`

	// REQUIRED; 响应头配置是否锁定，取值如下所示：
	// * true：是
	// * false：否
	ResponseHeaderLocked bool `json:"response_header_locked"`

	// REQUIRED; URL 鉴权签算配置是否锁定，取值如下所示：
	// * true：是
	// * false：否
	SignURLAuthLocked bool `json:"sign_url_auth_locked"`

	// REQUIRED; UA 访问限制配置是否锁定，取值如下所示：
	// * true：是
	// * false：否
	UaAccessRuleLocked bool `json:"ua_access_rule_locked"`
}

// GetDomainConfigResResultPageOptimization - 页面优化设置
type GetDomainConfigResResultPageOptimization struct {

	// REQUIRED; 是否开启页面优化，取值如下所示：
	// * true：启用
	// * false：关闭
	Enabled bool `json:"enabled"`

	// REQUIRED; 需要优化的对象列表，取值如下所示：
	// * html: 表示 HTML 页面。
	// * js: 表示 Javascript 代码。
	// * css: 表示 CSS 代码。
	OptimizationType []string `json:"optimization_type"`
}

type GetDomainConfigResResultRespHdrsItem struct {

	// REQUIRED
	AccessOriginControl bool `json:"access_origin_control"`

	// REQUIRED; 头部操作动作，取值如下所示：
	// * set：设置。
	// * delete：删除。
	Action string `json:"action"`

	// REQUIRED; header key
	Key string `json:"key"`

	// REQUIRED; header value
	Value string `json:"value"`
}

type GetDomainOwnerVerifyContentQuery struct {

	// REQUIRED; 待校验的域名。仅支持单域名校验。
	Domain string `json:"Domain" query:"Domain"`
}

type GetDomainOwnerVerifyContentRes struct {

	// REQUIRED
	ResponseMetadata *GetDomainOwnerVerifyContentResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetDomainOwnerVerifyContentResResult `json:"Result,omitempty"`
}

type GetDomainOwnerVerifyContentResResponseMetadata struct {

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

// GetDomainOwnerVerifyContentResResult - 视请求的接口而定
type GetDomainOwnerVerifyContentResResult struct {

	// REQUIRED; DNS 解析校验信息，仅当 NeedVerify 为 true 时返回。
	DNSVerifyInfo *GetDomainOwnerVerifyContentResResultDNSVerifyInfo `json:"DNSVerifyInfo"`

	// REQUIRED; 域名是否需要归属权校验： true: 需要校验 false: 无需校验
	NeedVerify bool `json:"NeedVerify"`
}

// GetDomainOwnerVerifyContentResResultDNSVerifyInfo - DNS 解析校验信息，仅当 NeedVerify 为 true 时返回。
type GetDomainOwnerVerifyContentResResultDNSVerifyInfo struct {

	// REQUIRED; 主机记录。
	Host string `json:"Host"`

	// REQUIRED; 记录类型。
	RecordType string `json:"RecordType"`

	// REQUIRED; 记录值。
	RecordValue string `json:"RecordValue"`
}

type GetImageAIDetailsQuery struct {

	// REQUIRED; 查询的结束 Unix 时间戳，StartTime 与 EndTime 时间间隔最大不超过 7 天。
	EndTime int `json:"EndTime" query:"EndTime"`

	// REQUIRED; 分页条数，取值范围为 (0, 100]。
	Limit int `json:"Limit" query:"Limit"`

	// REQUIRED; 队列 ID，通过 CreateImageAITask 接口返回。
	QueueID string `json:"QueueId" query:"QueueId"`

	// REQUIRED; 服务 ID。若 DataType 取值 uri，则提交的图片 URI 列表需在该服务内。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 查询的起始 Unix 时间戳，StartTime与EndTime时间间隔最大不超过 7 天。
	StartTime int `json:"StartTime" query:"StartTime"`

	// REQUIRED; 任务 ID，通过 CreateImageAITask 接口返回，缺省时查询指定队列下全部的任务。
	TaskID string `json:"TaskId" query:"TaskId"`

	// 分页偏移量，默认为 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 返回图片 URL 或 URI 中包含该值的任务。默认为空，不传则返回所有任务。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`

	// 执行状态，填入多个时使用英文逗号分隔。取值如下所示：
	// * Pending：排队中
	// * Dispatched：执行中
	// * Success：执行成功
	// * Fail：执行失败
	Status string `json:"Status,omitempty" query:"Status"`
}

type GetImageAIDetailsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAIDetailsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageAIDetailsResResult `json:"Result"`
}

type GetImageAIDetailsResResponseMetadata struct {

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

type GetImageAIDetailsResResult struct {

	// REQUIRED; 任务中每个条目的执行详情。
	ExecInfo []*GetImageAIDetailsResResultExecInfoItem `json:"ExecInfo"`

	// REQUIRED; 任务中包含的条目数。
	Total int `json:"Total"`
}

type GetImageAIDetailsResResultExecInfoItem struct {

	// REQUIRED; 使用的 AI 图像处理模板 ID，参看AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板 ID 对应的模板信息。
	WorkflowTemplateID string `json:"WorkflowTemplateId"`

	// 结束时间。
	EndAt string `json:"EndAt,omitempty"`

	// 条目 ID。
	EntryID string `json:"EntryId,omitempty"`

	// 执行输入。
	ExecInput *GetImageAIDetailsResResultExecInfoItemExecInput `json:"ExecInput,omitempty"`

	// 执行输出。
	ExecOutput *GetImageAIDetailsResResultExecInfoItemExecOutput `json:"ExecOutput,omitempty"`

	// 开始时间。
	StartAt string `json:"StartAt,omitempty"`

	// 执行状态。取值如下所示：
	// * Pending：排队中
	// * Dispatched：执行中
	// * Success：执行成功
	// * Fail：执行失败
	Status string `json:"Status,omitempty"`

	// 提交时间。
	SubmitAt string `json:"SubmitAt,omitempty"`
}

// GetImageAIDetailsResResultExecInfoItemExecInput - 执行输入。
type GetImageAIDetailsResResultExecInfoItemExecInput struct {

	// REQUIRED; 数据类型，取值如下所示：
	// * uri：指定 ServiceId 下存储 URI。
	// * url：公网可访问的 URL。
	DataType string `json:"DataType"`

	// REQUIRED; 图片 URL 或 URI。
	ObjectKey string `json:"ObjectKey"`
}

// GetImageAIDetailsResResultExecInfoItemExecOutput - 执行输出。
type GetImageAIDetailsResResultExecInfoItemExecOutput struct {

	// REQUIRED; AI 图像处理失败错误码 [https://www.volcengine.com/docs/508/1104726#%E9%94%99%E8%AF%AF%E7%A0%81]。仅当 Status 值为 Fail 时，ErrCode
	// 有值。
	ErrCode string `json:"ErrCode"`

	// REQUIRED; AI 图像处理失败错误信息。
	ErrMsg string `json:"ErrMsg"`

	// REQUIRED; AI 图像处理结果，是 JSON 压缩并转义后的字符串，仅当 Status 值为 Success 时，Output 有值。参看AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板
	// ID 和参数信息，根据具体的工作流的说明进行解析。
	Output string `json:"Output"`
}

type GetImageAIProcessQueuesQuery struct {

	// REQUIRED; 分页条数，取值范围为 (0, 100]。
	Limit int `json:"Limit" query:"Limit"`

	// 分页偏移量，默认为 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 返回队列名称或队列描述中包含该值的队列。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`
}

type GetImageAIProcessQueuesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAIProcessQueuesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *GetImageAIProcessQueuesResResult `json:"Result"`
}

type GetImageAIProcessQueuesResResponseMetadata struct {

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

// GetImageAIProcessQueuesResResult - 视请求的接口而定
type GetImageAIProcessQueuesResResult struct {

	// REQUIRED; 当前分页队列详细信息。
	Queues []*GetImageAIProcessQueuesResResultQueuesItem `json:"Queues"`

	// REQUIRED; 总数
	Total int `json:"Total"`
}

type GetImageAIProcessQueuesResResultQueuesItem struct {

	// REQUIRED; 队列ID
	ID string `json:"Id"`

	// REQUIRED; 队列名称
	Name string `json:"Name"`

	// 队列属性
	Attribute string `json:"Attribute,omitempty"`

	// 任务队列回调设置。
	CallbackConf *GetImageAIProcessQueuesResResultQueuesItemCallbackConf `json:"CallbackConf,omitempty"`

	// 队列创建时间
	CreateAt string `json:"CreateAt,omitempty"`

	// 队列描述
	Desc string `json:"Desc,omitempty"`

	// 队列是否开启回调配置
	EnableCallback bool `json:"EnableCallback,omitempty"`

	// 队列状态
	Status string `json:"Status,omitempty"`

	// 队列类型
	Type string `json:"Type,omitempty"`
}

// GetImageAIProcessQueuesResResultQueuesItemCallbackConf - 任务队列回调设置。
type GetImageAIProcessQueuesResResultQueuesItemCallbackConf struct {

	// REQUIRED; 回调地址
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 回调协议
	Method string `json:"Method"`

	// 回调附带参数
	Args string `json:"Args,omitempty"`

	// 回调数据类型，JSON或者XML
	DataFormat string `json:"DataFormat,omitempty"`

	// 回调触发类型，整个任务回调或者单个条目回调
	Type string `json:"Type,omitempty"`
}

type GetImageAITasksQuery struct {

	// REQUIRED; 查询的结束 Unix 时间戳，StartTime 与 EndTime 时间间隔最大不超过 7 天。
	EndTime int `json:"EndTime" query:"EndTime"`

	// REQUIRED; 队列 ID，通过 CreateImageAITask 接口返回。
	QueueID string `json:"QueueId" query:"QueueId"`

	// REQUIRED; 查询的起始 Unix 时间戳，StartTime 与 EndTime 时间间隔最大不超过 7 天。
	StartTime int `json:"StartTime" query:"StartTime"`

	// 单次查询列出的任务的个数，取值范围为 (0,1000]，默认值为 1000。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 上一次查询返回的位置标记，作为本次查询的起点信息，默认值为空。
	Marker string `json:"Marker,omitempty" query:"Marker"`

	// 搜索url、uri
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`

	// 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId,omitempty" query:"ServiceId"`

	// 指定查询的任务状态，缺省时将查询全部状态的任务。取值如下所示：
	// * Running：任务运行中
	// * Suspend：任务中断
	// * Done：任务已完成
	// * Cancel：任务取消
	// * Failed：任务失败
	Status string `json:"Status,omitempty" query:"Status"`

	// 任务 ID，通过 CreateImageAITask 接口返回，缺省时查询指定队列下全部的任务。
	TaskID string `json:"TaskId,omitempty" query:"TaskId"`
}

type GetImageAITasksRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAITasksResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *GetImageAITasksResResult `json:"Result"`
}

type GetImageAITasksResResponseMetadata struct {

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

// GetImageAITasksResResult - 视请求的接口而定
type GetImageAITasksResResult struct {

	// REQUIRED; 是否还有更多任务，取值如下所示：
	// * true：是，还有任务未列出。
	// * false：否，已列出所有任务。
	HasMore bool `json:"HasMore"`

	// REQUIRED; 指定的队列 ID。
	QueueID string `json:"QueueId"`

	// REQUIRED; AI 图像处理任务的各类信息。
	TaskInfo []*GetImageAITasksResResultTaskInfoItem `json:"TaskInfo"`

	// HasMore取值为true时（即本次查询还有未列举到的任务时），Marker应作为查询起始位置标记，您需要在下一次查询时传入该值。
	Marker string `json:"Marker,omitempty"`
}

type GetImageAITasksResResultTaskInfoItem struct {

	// REQUIRED; 数据类型，取值如下所示：
	// * uri：指定 ServiceId 下存储 URI。
	// * url：公网可访问的 URL。
	DataType string `json:"DataType"`

	// REQUIRED; 任务的结束执行时间。
	EndAt string `json:"EndAt"`

	// REQUIRED; 任务中执行失败的条目数。
	Fail int `json:"Fail"`

	// REQUIRED; 任务中重试的条目数。
	// :::tip 当因系统内部原因导致的条目转码失败，系统将自动重试该条目，最大重试次数为 5。 :::
	Retry int `json:"Retry"`

	// REQUIRED; 任务的开始执行时间。
	StartAt string `json:"StartAt"`

	// REQUIRED; 任务的执行状态，取值如下所示：
	// * Running：任务运行中
	// * Suspend：任务中断
	// * Done：任务已完成
	// * Cancel：任务取消
	// * Failed：任务失败
	Status string `json:"Status"`

	// REQUIRED; 任务的提交时间。
	SubmitAt string `json:"SubmitAt"`

	// REQUIRED; 任务中执行成功的条目数。
	Success int `json:"Success"`

	// REQUIRED; 任务 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 任务中包含的条目数。
	Total int `json:"Total"`

	// REQUIRED; AI 图像处理模板参数，参看AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取参数信息。
	WorkflowParameter string `json:"WorkflowParameter"`

	// REQUIRED; 使用的 AI 图像处理模板 ID，参看AI 图像处理模板 [https://www.volcengine.com/docs/508/1515840]页面获取模板 ID 对应的模板信息。
	WorkflowTemplateID string `json:"WorkflowTemplateId"`
}

type GetImageAddOnTagQuery struct {

	// REQUIRED; 组件标签 key。取值固定为功能属性，返回相关标签值。
	Key string `json:"Key" query:"Key"`

	// 组件类型，默认获取所有类型的标签信息。取值如下所示：
	// * AI：智能处理类型
	// * Other：其他增值类型
	Type string `json:"Type,omitempty" query:"Type"`
}

type GetImageAddOnTagRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAddOnTagResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetImageAddOnTagResResult           `json:"Result,omitempty"`
}

type GetImageAddOnTagResResponseMetadata struct {

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

type GetImageAddOnTagResResult struct {

	// REQUIRED; 账号内的可见的附加组件标签名称，一个标签下可以包含多个组件功能。
	Tag []string `json:"Tag"`
}

type GetImageAiGenerateTaskQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; AIGC 任务 ID，您可通过调用创建 AIGC 异步任务 [https://www.volcengine.com/docs/508/1134013]获取。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetImageAiGenerateTaskRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAiGenerateTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageAiGenerateTaskResResult `json:"Result,omitempty"`
}

type GetImageAiGenerateTaskResResponseMetadata struct {

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

// GetImageAiGenerateTaskResResult - 视请求的接口而定
type GetImageAiGenerateTaskResResult struct {

	// REQUIRED
	AestheticScores []float64 `json:"AestheticScores"`

	// REQUIRED
	SDScores []float64 `json:"SDScores"`

	// REQUIRED; 任务状态，取值如下所示：
	// * Success：处理成功
	// * Running：进行中
	// * Failed：处理失败
	Status string `json:"Status"`

	// REQUIRED; 生成的处理图 URI，固定为 4 个长宽比为 1:1 的方图。
	Uris []string `json:"Uris"`
}

type GetImageAlertRecordsBody struct {

	// REQUIRED; 获取数据结束时间点，需在起始时间点之后。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。
	EndTime string `json:"EndTime"`

	// REQUIRED; 获取数据起始时间点。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm，比如2019-06-02T00:00:00+08:00。 :::tip 由于仅支持查询近 90 天的历史数据，则若此刻时间为2011-11-21T16:14:00+08:00，那么您可输入最早的开始时间为2011-08-18T00:00:00+08:00。
	// :::
	StartTime string `json:"StartTime"`

	// 应用 ID。您可以通过调用GetImageXQueryApps [https://www.volcengine.com/docs/508/19511]的方式获取账号下全部的 AppId。
	AppID string `json:"AppId,omitempty"`

	// 获取个数限制。默认值为 10，取值范围为 (0,100]。
	Limit int `json:"Limit,omitempty"`

	// 分页偏移量。默认值为 0，表示从最新一个开始获取。
	Marker string `json:"Marker,omitempty"`

	// 告警名称，正则匹配。不填则查询所有告警记录。
	Name string `json:"Name,omitempty"`

	// 告警规则 ID，完全匹配。不填则查询所有告警记录。
	RuleID string `json:"RuleId,omitempty"`
}

type GetImageAlertRecordsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAlertRecordsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *GetImageAlertRecordsResResult `json:"Result"`
}

type GetImageAlertRecordsResResponseMetadata struct {

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

// GetImageAlertRecordsResResult - 视请求的接口而定
type GetImageAlertRecordsResResult struct {

	// REQUIRED; 告警记录列表
	AlertRecords []*GetImageAlertRecordsResResultAlertRecordsItem `json:"AlertRecords"`

	// REQUIRED; 是否有更多记录，取值如下所示：
	// * true：是
	// * false：否
	HasMore bool `json:"HasMore"`

	// REQUIRED; 记录总数
	Total int `json:"Total"`
}

type GetImageAlertRecordsResResultAlertRecordsItem struct {

	// REQUIRED; 告警时间。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	AlertAt string `json:"AlertAt"`

	// REQUIRED; 告警条件
	AlertCond *GetImageAlertRecordsResResultAlertRecordsItemAlertCond `json:"AlertCond"`

	// REQUIRED; 告警级别，取值如下所示：
	// * warn：警告
	// * error：错误
	// * fatal：致命
	Level string `json:"Level"`

	// REQUIRED; 记录标识，用于获取下一页数据。
	Marker string `json:"Marker"`

	// REQUIRED; 告警规则名称
	Name string `json:"Name"`

	// REQUIRED; 告警平台
	OS string `json:"OS"`

	// REQUIRED; 告警阶段，取值如下所示：
	// * upload：图片上传-上传 1.0
	// * uploadv2：图片上传-上传 2.0
	// * cdn：图片加载-下行网络监控
	// * client：图片加载-客户端传状态监控
	// * sensible：图片加载-感知指标监控
	Phase string `json:"Phase"`

	// REQUIRED; 告警规则 ID
	RuleID string `json:"RuleId"`
}

// GetImageAlertRecordsResResultAlertRecordsItemAlertCond - 告警条件
type GetImageAlertRecordsResResultAlertRecordsItemAlertCond struct {

	// REQUIRED; 各指标告警信息
	AlertContent *Components1Muxqr1SchemasGetimagealertrecordsresPropertiesResultPropertiesAlertrecordsItemsPropertiesAlertcondPropertiesAlertcontent `json:"AlertContent"`

	// REQUIRED; 规则之间的逻辑关系，取值如下所示：
	// * and：和
	// * or：或
	LogicOp string `json:"LogicOp"`
}

type GetImageAlertRecordsResResultAlertRecordsPropertiesItemsItem struct {

	// REQUIRED; 指标上一周期值对应的时间。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。仅在Op为同比/环比相关方法时有值。
	BaseTime string `json:"BaseTime"`

	// REQUIRED; 触发告警的拆分维度值，仅当告警配置了维度拆分时有值。
	DimVal string `json:"DimVal"`

	// REQUIRED; 前序指标值列表。长度为 RepeatCnt-1
	PrevVals []*GetImageAlertRecordsResResultAlertRecordsPropertiesItemsPrevValsItem `json:"PrevVals"`

	// REQUIRED; 指标值
	Val float64 `json:"Val"`

	// REQUIRED; 指标上一周期值，仅在Op为同比/环比相关方法时有值。
	ValBase float64 `json:"ValBase"`

	// REQUIRED; 指标值对应的时间。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	ValTime string `json:"ValTime"`
}

type GetImageAlertRecordsResResultAlertRecordsPropertiesItemsPrevValsItem struct {

	// REQUIRED; 指标上一周期值对应的时间。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。仅在Op为同比/环比相关方法时有值。
	BaseTime string `json:"BaseTime"`

	// REQUIRED; 指标值
	Val float64 `json:"Val"`

	// REQUIRED; 指标上一周期值，仅在Op为同比/环比相关方法时有值。
	ValBase string `json:"ValBase"`

	// REQUIRED; 指标值对应的时间。日期格式按照ISO8601表示法，格式为：YYYY-MM-DDThh:mm:ss±hh:mm。
	ValTime string `json:"ValTime"`
}

type GetImageAllDomainCertRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAllDomainCertResResponseMetadata `json:"ResponseMetadata"`
	Result           []*GetImageAllDomainCertResResultItem     `json:"Result,omitempty"`
}

type GetImageAllDomainCertResResponseMetadata struct {

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

type GetImageAllDomainCertResResultItem struct {

	// REQUIRED; 证书通用名称
	CertCommonName string `json:"CertCommonName"`

	// REQUIRED; 证书id
	CertID string `json:"CertID"`

	// REQUIRED; 证书备注
	CertName string `json:"CertName"`

	// REQUIRED; 证书过期时间
	CertNotAfter float64 `json:"CertNotAfter"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 域名状态
	DomainStatus string `json:"DomainStatus"`

	// REQUIRED; 可通过san与certcommonname来匹配证书可用的域名
	San []string `json:"San"`
}

type GetImageAnalyzeResultQuery struct {

	// REQUIRED; 任务运行结束时间，Unix 时间戳。
	EndTime int `json:"EndTime" query:"EndTime"`

	// REQUIRED; 任务运行开始时间，Unix 时间戳。
	StartTime int `json:"StartTime" query:"StartTime"`

	// REQUIRED; 任务 ID，您可以通过调用GetImageAnalyzeTasks [https://www.volcengine.com/docs/508/1160417]获取指定地区全部离线评估任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`

	// 文件名
	File string `json:"File,omitempty" query:"File"`

	// 分页条数。默认值为 10。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，默认为 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 任务执行 ID
	RunID string `json:"RunId,omitempty" query:"RunId"`
}

type GetImageAnalyzeResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAnalyzeResultResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageAnalyzeResultResResult `json:"Result"`
}

type GetImageAnalyzeResultResResponseMetadata struct {

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

type GetImageAnalyzeResultResResult struct {

	// REQUIRED; 任务结果
	Results []*GetImageAnalyzeResultResResultResultsItem `json:"Results"`

	// REQUIRED; 查询到的总量
	Total int `json:"Total"`
}

type GetImageAnalyzeResultResResultResultsItem struct {

	// 评估结果存储 URI
	ResultURI string `json:"ResultUri,omitempty"`

	// 任务指定 ID
	RunID   string `json:"RunId,omitempty"`
	StartAt string `json:"StartAt,omitempty"`

	// 任务状态，取值如下所示：
	// * Running：进行中
	// * Suspend：暂停
	// * Done：结束
	Status string `json:"Status,omitempty"`
}

type GetImageAnalyzeTasksQuery struct {

	// 分页条数。取值范围为 (0,100]，默认值为 100。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，默认为 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 任务地区。默认为 cn，表示国内。
	Region string `json:"Region,omitempty" query:"Region"`

	// 返回任务名称或描述中包含该值的任务。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`
}

type GetImageAnalyzeTasksRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAnalyzeTasksResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageAnalyzeTasksResResult `json:"Result"`
}

type GetImageAnalyzeTasksResResponseMetadata struct {

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

type GetImageAnalyzeTasksResResult struct {

	// REQUIRED; 任务
	Tasks []*GetImageAnalyzeTasksResResultTasksItem `json:"Tasks"`

	// REQUIRED; 查询到任务数量
	Total int `json:"Total"`
}

type GetImageAnalyzeTasksResResultTasksItem struct {

	// REQUIRED
	EvalStages []string `json:"EvalStages"`

	// REQUIRED
	ID string `json:"Id"`

	// REQUIRED; 评估任务更新时间
	UpdateAt string `json:"UpdateAt"`

	// REQUIRED
	VqTypes []string `json:"VqTypes"`

	// 评估任务创建时间
	CreateAt string `json:"CreateAt,omitempty"`

	// 任务描述
	Desc string `json:"Desc,omitempty"`

	// 是否模拟模板每阶段输出，取值如下所示：
	// * true：是
	// * false：否
	EvalPerStage bool `json:"EvalPerStage,omitempty"`

	// 离线评估任务名称
	Name string `json:"Name,omitempty"`

	// txt 评估文件的 Store URI。
	ResURI     []string `json:"ResUri,omitempty"`
	SampleRate float64  `json:"SampleRate,omitempty"`

	// 所在服务的服务 ID
	ServiceID string `json:"ServiceId,omitempty"`

	// 任务状态，取值如下所示：
	// * Pending：排队中
	// * Done：已完成
	// * Running：进行中
	Status string `json:"Status,omitempty"`

	// 该条评估任务的任务 ID
	TaskID string `json:"TaskId,omitempty"`

	// Type 取值 UriFile 时，指定的模版名称。
	Tpl string `json:"Tpl,omitempty"`

	// 任务类型，取值如下所示：
	// * UrlFile：在线提交 URL 离线评估
	// * UriFile：在线提交 URI 离线评估
	Type string `json:"Type,omitempty"`
}

type GetImageAuditResultQuery struct {

	// REQUIRED; 任务 ID，您可通过调用查询所有审核任务 [https://www.volcengine.com/docs/508/1160409]获取所需的任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`

	// 审核建议，缺省情况下返回全部任务。取值如下所示：
	// * nopass：建议不通过
	// * recheck：建议复审
	AuditSuggestion string `json:"AuditSuggestion,omitempty" query:"AuditSuggestion"`

	// 图片类型，缺省情况下返回全部类型任务。取值如下所示：
	// * problem：问题图片
	// * frozen：冻结图片
	// * fail：审核失败图片
	ImageType string `json:"ImageType,omitempty" query:"ImageType"`

	// 分页条数。取值范围为 (0,100]，默认值为 10。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 上一次查询返回的位置标记，作为本次列举的起点信息。默认值为 0。
	Marker int `json:"Marker,omitempty" query:"Marker"`

	// 问题类型，取值根据审核类型的不同其取值不同。缺省情况下返回全部类型任务。
	// * 基础安全审核 * govern：涉政
	// * porn：涉黄
	// * illegal：违法违规
	// * terror：涉暴
	//
	//
	// * 智能安全审核 * 图像风险识别 * porn：涉黄，主要适用于通用色情、色情动作、性行为、性暗示、性分泌物、色情动漫、色情裸露等涉黄场景的风险识别
	// * sensitive1：涉敏1，具体指涉及暴恐风险
	// * sensitive2：涉敏2，具体值涉及政治内容风险
	// * forbidden：违禁，主要适用于打架斗殴、爆炸、劣迹艺人等场景的风险识别
	// * uncomfortable：引人不适，主要适用于恶心、恐怖、尸体、血腥等引人不适场景的风险识别
	// * qrcode：二维码，主要适用于识别常见二维码（QQ、微信、其他二维码等）
	// * badpicture：不良画面，主要适用于自我伤害、丧葬、不当车播、吸烟/纹身/竖中指等不良社会风气的风险识别
	// * sexy：性感低俗，主要适用于舌吻、穿衣性行为、擦边裸露等多种性感低俗场景的风险识别
	// * age：年龄，主要适用于图中人物对应的年龄段识别
	// * underage：未成年相关，主要适用于儿童色情、儿童邪典等风险识别
	// * quality：图片质量，主要适用于图片模糊、纯色边框、纯色屏等风险识别
	//
	//
	// * 图文风险识别 * ad：广告，综合图像及文字内容智能识别广告
	// * defraud：诈骗，综合图像及文字内容智能识别诈骗
	// * charillegal：文字违规，图片上存在涉黄、涉敏、违禁等违规文字
	Problem string `json:"Problem,omitempty" query:"Problem"`

	// 审核场景，缺省情况下查询全部场景的任务。取值如下所示：
	// * UrlFile：上传 txt 审核文件处理场景
	// * Url：上传审核图片 URL 处理场景
	// * Upload：图片上传场景
	Type string `json:"Type,omitempty" query:"Type"`
}

type GetImageAuditResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAuditResultResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageAuditResultResResult `json:"Result"`
}

type GetImageAuditResultResResponseMetadata struct {

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

type GetImageAuditResultResResult struct {

	// REQUIRED; 是否还有更多任务，取值如下所示：
	// * true：是，还有其他任务未列举
	// * false：否，已列举所有任务
	HaveMore bool `json:"HaveMore"`

	// REQUIRED; 任务结果
	Results []*GetImageAuditResultResResultResultsItem `json:"Results"`
}

type GetImageAuditResultResResultResultsItem struct {

	// 该任务的审核能力。取值如下所示：
	// * 0：基础审核能力
	// * 1：智能审核能力
	Ability int `json:"Ability,omitempty"`

	// 审核建议，取值如下所示：
	// * nopass：建议不通过
	// * recheck：建议复审
	AuditSuggestion string `json:"AuditSuggestion,omitempty"`

	// 审核结束时间
	EndAt string `json:"EndAt,omitempty"`

	// 条目 ID
	EntryID string `json:"EntryId,omitempty"`

	// 错误信息
	ErrMsg string `json:"ErrMsg,omitempty"`

	// 审核结果，取值如下所示：
	// * problem：问题图片
	// * frozen：冻结图片
	// * fail：审核失败图片
	ImageType string `json:"ImageType,omitempty"`

	// 表示 txt 审核文件的存储 URI。
	ImageURI string `json:"ImageUri,omitempty"`

	// HaveMore取值true时，即本次查询还有未列举到的任务时。Marker作为起始条目位置标记，您需要在下一次查询时传入该值。
	Marker string `json:"Marker,omitempty"`

	// 审核发现图片问题类型
	Problems []string `json:"Problems,omitempty"`

	// 任务 ID
	TaskID string `json:"TaskId,omitempty"`

	// 该任务被指定的审核场景，取值如下所示：
	// * UrlFile：上传 txt 审核文件处理场景
	// * Url：上传审核图片 URL 处理场景
	// * Upload：图片上传场景
	Type string `json:"Type,omitempty"`
}

type GetImageAuditTasksQuery struct {

	// REQUIRED; 审核任务类型，当前仅支持取值为audit。
	TaskType string `json:"TaskType" query:"TaskType"`

	// 审核能力，缺省情况下查询全部审核类型的任务。取值如下所示：
	// * 0：基础审核能力
	// * 1：智能审核能力
	AuditAbility int `json:"AuditAbility,omitempty" query:"AuditAbility"`

	// 分页条数。取值范围为 (0,100]，默认值为 100。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，默认为 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 任务地区。仅支持默认取值cn，表示国内。
	Region string `json:"Region,omitempty" query:"Region"`

	// 审核状态，缺省情况下查询全部状态的任务。取值如下所示：
	// * Running：审核中
	// * Suspend：已暂停
	// * Done：已完成
	// * Failed：审核失败
	// * Cancel：已取消
	Status string `json:"Status,omitempty" query:"Status"`

	// 审核场景，缺省情况下查询全部场景的任务。取值如下所示：
	// * UrlFile：上传 txt 审核文件处理场景
	// * Url：上传审核图片 URL 处理场景
	// * Upload：图片上传场景
	Type string `json:"Type,omitempty" query:"Type"`
}

type GetImageAuditTasksRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAuditTasksResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageAuditTasksResResult `json:"Result"`
}

type GetImageAuditTasksResResponseMetadata struct {

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

type GetImageAuditTasksResResult struct {

	// REQUIRED; 审核任务详情
	Tasks []*GetImageAuditTasksResResultTasksItem `json:"Tasks"`

	// REQUIRED; 符合请求的审核任务数量
	Total int `json:"Total"`
}

type GetImageAuditTasksResResultTasksItem struct {

	// REQUIRED; 该审核任务的更新时间
	UpdateAt string `json:"UpdateAt"`

	// 审核任务
	AuditTask *GetImageAuditTasksResResultTasksItemAuditTask `json:"AuditTask,omitempty"`

	// 该审核任务的创建时间
	CreateAt string `json:"CreateAt,omitempty"`

	// 审核任务 ID
	ID string `json:"Id,omitempty"`

	// 仅当Type值为UrlFile时，有返回值。表示 txt 审核文件的存储 URI。
	ResURI []string `json:"ResUri,omitempty"`

	// 该审核任务所在的服务 ID
	ServiceID string `json:"ServiceId,omitempty"`

	// 该条审核任务的状态，取值如下所示：
	// * Running：审核中
	// * Suspend：已暂停
	// * Done：已完成
	// * Failed：审核失败
	// * Cancel：已取消
	Status string `json:"Status,omitempty"`

	// 审核任务类型，当前仅支持取值为 audit。
	TaskType string `json:"TaskType,omitempty"`

	// 审核场景，取值如下所示：
	// * UrlFile：上传 txt 审核文件处理场景
	// * Url：上传审核图片 URL 处理场景
	// * Upload：图片上传场景
	Type string `json:"Type,omitempty"`
}

// GetImageAuditTasksResResultTasksItemAuditTask - 审核任务
type GetImageAuditTasksResResultTasksItemAuditTask struct {

	// REQUIRED; 审核能力，取值如下所示：
	// * 0：基础审核能力
	// * 1：智能审核能力
	AuditAbility int `json:"AuditAbility"`

	// REQUIRED; 指定的审核维度
	AuditDimensions []string `json:"AuditDimensions"`

	// REQUIRED; 仅当EnableAuditRange取值1时，有返回值。表示指定审核的目录前缀。
	AuditPrefix []string `json:"AuditPrefix"`

	// REQUIRED; 指定的智能安全审核类型下图片文本审核维度。
	AuditTextDimensions []string `json:"AuditTextDimensions"`

	// REQUIRED; 指定的回调类型
	CallbackDimensions []string `json:"CallbackDimensions"`

	// REQUIRED; 回调图片类型，取值如下所示：
	// * normal：正常图片
	// * problem：问题图片
	// * frozen：冻结图片
	// * fail：审核失败图片
	CallbackImageTypes []string `json:"CallbackImageTypes"`

	// REQUIRED; 指定的回调 URL
	CallbackURL string `json:"CallbackUrl"`

	// REQUIRED; 仅当Type取值Upload时，有返回值。 审核范围，取值如下所示：
	// * 0：不限范围
	// * 1：指定范围
	EnableAuditRange int `json:"EnableAuditRange"`

	// REQUIRED; 是否开启回调，取值如下所示：
	// * true：开启
	// * false：不开启
	EnableCallback bool `json:"EnableCallback"`

	// REQUIRED; 是否开启冻结，取值如下所示：
	// * true：开启
	// * false：不开启
	EnableFreeze bool `json:"EnableFreeze"`

	// REQUIRED; 是否开启大图审核，取值如下所示：
	// * true：开启
	// * false：不开启
	EnableLargeImageDetect bool `json:"EnableLargeImageDetect"`

	// REQUIRED; 审核失败的图片数量
	FailedCount int `json:"FailedCount"`

	// REQUIRED; 指定的冻结维度
	FreezeDimensions []string `json:"FreezeDimensions"`

	// REQUIRED; 冻结策略，当前仅支持取0，表示禁用图片。
	FreezeStrategy int `json:"FreezeStrategy"`

	// REQUIRED; 指定的冻结措施，取值如下所示：
	// * recheck：建议复审
	// * nopass：审核不通过
	FreezeType []string `json:"FreezeType"`

	// REQUIRED; 仅当 Type 取值 Url 时，有返回值。表示上传的待审核图片 Url 列表详情。
	ImageInfos []*ComponentsDgbrljSchemasGetimageaudittasksresPropertiesResultPropertiesTasksItemsPropertiesAudittaskPropertiesImageinfosItems `json:"ImageInfos"`

	// REQUIRED; 仅当EnableAuditRange取值1时，有返回值。表示指定不审核的目录前缀。
	NoAuditPrefix []string `json:"NoAuditPrefix"`

	// REQUIRED; 审核中的图片数量
	ProcessedNumber int `json:"ProcessedNumber"`

	// REQUIRED; 审核成功的图片数量
	SuccessCount int `json:"SuccessCount"`
}

type GetImageAuthKeyQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageAuthKeyRes struct {

	// REQUIRED
	ResponseMetadata *GetImageAuthKeyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageAuthKeyResResult `json:"Result,omitempty"`
}

type GetImageAuthKeyResResponseMetadata struct {

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

// GetImageAuthKeyResResult - 视请求的接口而定
type GetImageAuthKeyResResult struct {

	// REQUIRED; 该服务的主鉴权 Key。
	PrimaryKey string `json:"PrimaryKey"`

	// REQUIRED; 该服务的备鉴权 Key。
	SecondaryKey string `json:"SecondaryKey"`

	// REQUIRED; 指定的服务 ID。
	ServiceID string `json:"ServiceId"`
}

type GetImageBackgroundColorsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageBackgroundColorsResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *GetImageBackgroundColorsResResult `json:"Result,omitempty"`
}

type GetImageBackgroundColorsResResponseMetadata struct {

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

// GetImageBackgroundColorsResResult - title
type GetImageBackgroundColorsResResult struct {

	// REQUIRED; 颜色列表。
	Colors []string `json:"Colors"`
}

type GetImageBgFillResultBody struct {

	// REQUIRED; 填充模型，取值如下所示：
	// * 0：国漫风格模型
	// * 1：通用模型
	// *
	Model int `json:"Model"`

	// REQUIRED; 服务ID。
	// * 您可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待填充原图的存储 URI 和 URL（公网可访问的 URL）。
	StoreURI string `json:"StoreUri"`

	// 向下填充比例，取值范围为 [0, 0.4]。
	Bottom float64 `json:"Bottom,omitempty"`

	// 向左填充比例，取值范围为 [0, 0.4]。
	Left float64 `json:"Left,omitempty"`

	// 向右填充比例，取值范围为 [0, 0.4]。
	Right float64 `json:"Right,omitempty"`

	// 向上填充比例，取值范围为 [0, 0.4]。
	Top float64 `json:"Top,omitempty"`
}

type GetImageBgFillResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageBgFillResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageBgFillResultResResult `json:"Result,omitempty"`
}

type GetImageBgFillResultResResponseMetadata struct {

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

// GetImageBgFillResultResResult - 视请求的接口而定
type GetImageBgFillResultResResult struct {

	// REQUIRED; 填充结果图的 URI。
	ResURI string `json:"ResUri"`
}

type GetImageComicResultBody struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待处理原图的存储 URI 和 URL（可公网访问的 URL）。
	StoreURI string `json:"StoreUri"`
}

type GetImageComicResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageComicResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageComicResultResResult `json:"Result,omitempty"`
}

type GetImageComicResultResResponseMetadata struct {

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

// GetImageComicResultResResult - 视请求的接口而定
type GetImageComicResultResResult struct {

	// REQUIRED; 结果图存储 URI。
	ResURI string `json:"ResUri"`
}

type GetImageContentBlockListBody struct {

	// REQUIRED; 结束查询时间，unix 时间戳，单位为秒。
	EndTime int `json:"EndTime"`

	// REQUIRED; 开始查询时间，unix 时间戳，单位为秒。
	StartTime int `json:"StartTime"`

	// 域名，指定后将返回包含该域名的 URL 禁用任务。
	Domain string `json:"Domain,omitempty"`

	// 按时间排序，取值如下所示：
	// * asc：正序
	// * desc：逆序
	Order string `json:"Order,omitempty"`

	// 页码，仅返回该页码上的任务。默认值为 1。
	PageNum int `json:"PageNum,omitempty"`

	// 每页条数，取值范围是[10,1000]。默认值为 100。
	PageSize int `json:"PageSize,omitempty"`

	// 资源更新状态，取值如下所示：
	// * submitting：提交中
	// * running：执行中
	// * succeed：成功
	// * failed：失败
	State string `json:"State,omitempty"`

	// 指定要查询的 URL，缺省情况下查询当前服务所有禁用任务列表。
	URL string `json:"Url,omitempty"`
}

type GetImageContentBlockListRes struct {

	// REQUIRED
	ResponseMetadata *GetImageContentBlockListResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *GetImageContentBlockListResResult `json:"Result"`
}

type GetImageContentBlockListResResponseMetadata struct {

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

// GetImageContentBlockListResResult - 视请求的接口而定
type GetImageContentBlockListResResult struct {

	// REQUIRED; 具体数据
	Data []*GetImageContentBlockListResResultDataItem `json:"Data"`

	// REQUIRED; 当前页码
	PageNum int `json:"PageNum"`

	// REQUIRED; 每页最大记录数
	PageSize int `json:"PageSize"`

	// REQUIRED; 符合查询条件的总记录数
	Total int `json:"Total"`
}

type GetImageContentBlockListResResultDataItem struct {

	// REQUIRED; 任务的创建时间
	CreateTime int `json:"CreateTime"`

	// REQUIRED; 完成结果提示信息
	Msg string `json:"Msg"`

	// REQUIRED; 任务进度
	Process string `json:"Process"`

	// REQUIRED; URL 状态
	State string `json:"State"`

	// REQUIRED; 任务类型，取值如下所示：
	// * block_url：禁用任务。
	// * unblock_url：解禁任务，表示此时解禁未完成。
	TaskType string `json:"TaskType"`

	// REQUIRED; 指定的 URL
	URL string `json:"Url"`

	// REQUIRED; 任务的更新时间
	UpdateTime int `json:"UpdateTime"`
}

type GetImageContentTaskDetailBody struct {

	// REQUIRED; 查询结束时间，unix 时间戳，单位为秒。
	EndTime int `json:"EndTime"`

	// REQUIRED; 查询开始时间，unix 时间戳，单位为秒。
	StartTime int `json:"StartTime"`

	// 域名，指定后返回包含该域名的 URL 任务。
	Domain string `json:"Domain,omitempty"`

	// 按时间排序，取值如下所示：
	// * asc：正序
	// * desc：逆序
	Order string `json:"Order,omitempty"`

	// 页码，系统将仅返回该页面上的任务。默认值为 1。
	PageNum int `json:"PageNum,omitempty"`

	// 每页条数，取值范围为 [10,1000]。默认值为 100。
	PageSize int `json:"PageSize,omitempty"`

	// 内容管理资源状态，取值如下所示：
	// * submitting：提交中
	// * running：执行中
	// * succeed：成功
	// * failed：失败
	State string `json:"State,omitempty"`

	// 待查询任务 ID
	TaskID string `json:"TaskId,omitempty"`

	// 内容管理任务类型，缺省情况下表示查询全部任务。取值如下所示：
	// * refresh：刷新任务，包含刷新 URL 和刷新目录。
	// * refresh_url：刷新 URL
	// * block_url：禁用 URL
	// * unblock_url：解禁 URL
	// * preload_url：预热 URL
	// * refresh_dir：目录刷新（支持根目录刷新）
	TaskType string `json:"TaskType,omitempty"`

	// 资源 URL 或者目录，可精确匹配，取值为空时表示查询全部任务。
	URL string `json:"Url,omitempty"`
}

type GetImageContentTaskDetailRes struct {

	// REQUIRED
	ResponseMetadata *GetImageContentTaskDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *GetImageContentTaskDetailResResult `json:"Result"`
}

type GetImageContentTaskDetailResResponseMetadata struct {

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

// GetImageContentTaskDetailResResult - 视请求的接口而定
type GetImageContentTaskDetailResResult struct {

	// REQUIRED; 具体数据
	Data []*GetImageContentTaskDetailResResultDataItem `json:"Data"`

	// REQUIRED; 当前页码
	PageNum int `json:"PageNum"`

	// REQUIRED; 每页最大记录数
	PageSize int `json:"PageSize"`

	// REQUIRED; 总记录数
	Total int `json:"Total"`
}

type GetImageContentTaskDetailResResultDataItem struct {

	// REQUIRED; 任务的创建时间
	CreateTime int `json:"CreateTime"`

	// REQUIRED; 完成结果提示信息
	Msg string `json:"Msg"`

	// REQUIRED; 任务进度
	Process string `json:"Process"`

	// REQUIRED; 资源任务状态，取值如下所示：
	// * submitting：提交中
	// * running：执行中
	// * succeed：成功
	// * failed：失败
	State string `json:"State"`

	// REQUIRED; 任务 ID
	TaskID string `json:"TaskId"`

	// REQUIRED; 资源任务类型，取值如下所示：
	// * refresh：刷新任务
	// * refresh_url：刷新 URL
	// * block_url：禁用 URL
	// * unblock_url：解禁 URL
	// * preload_url：预热 URL
	// * refresh_dir：目录刷新
	TaskType string `json:"TaskType"`

	// REQUIRED; 任务指定的 URL 或目录
	URL string `json:"Url"`

	// REQUIRED; 任务的更新时间
	UpdateTime int `json:"UpdateTime"`
}

type GetImageDetectResultBody struct {

	// REQUIRED; 检测类型，取值仅支持 face，表示检测图片中人脸所在坐标。
	DetectType string `json:"DetectType"`

	// REQUIRED; 指定服务下的待检测图片的 StoreUri 或者公网可访问 Url。
	ImageURI string `json:"ImageUri"`

	// 选择人脸检测时必填，人脸检测阈值，取值范围（0,1）
	FaceDetectThresh float64 `json:"FaceDetectThresh,omitempty"`
}

type GetImageDetectResultQuery struct {

	// REQUIRED; 待检测图片对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageDetectResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageDetectResultResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetImageDetectResultResResult           `json:"Result,omitempty"`
}

type GetImageDetectResultResResponseMetadata struct {

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

type GetImageDetectResultResResult struct {

	// REQUIRED; 检测类型
	DetectType string `json:"DetectType"`

	// REQUIRED; 人脸识别结果
	Faces []*GetImageDetectResultResResultFacesItem `json:"Faces"`
}

type GetImageDetectResultResResultFacesItem struct {

	// REQUIRED; 人脸坐标位置。 :::tip 例如[11,22,33,44]，表示人脸的左上角和右下角坐标。其中11为左上角横坐标，22为左上角纵坐标，33为右下角横坐标，44为右下角纵坐标。 :::
	Box []int `json:"Box"`

	// REQUIRED; 坐标置信度，表示识别内容可信程度。 值越大，代表可信程度越高。置信度高于 90% 表示高可信，60%～90% 建议人工二次确认，低于 60% 表示不可信。
	Score float64 `json:"Score"`
}

type GetImageDuplicateDetectionBody struct {

	// REQUIRED; 需要去重的图片 URL，多个地址以英文逗号分割。图片格式仅支持 JPEG（.jpeg 或 .jpg）和 PNG，支持格式混合输入。
	Urls []string `json:"Urls"`

	// 是否使用异步，取值如下所示：
	// * true：使用异步去重
	// * false：（默认）不使用异步去重
	Async bool `json:"Async,omitempty"`

	// 当Async取值为true即启用异步时需要添加回调地址，地址使用 POST 请求方式。回调内容详见回调参数。
	Callback string `json:"Callback,omitempty"`

	// 图像特征提取类型，取值如下所示：
	// * hash：（默认）基于图像的像素值、颜色分布、纹理等特征生成哈希码，并通过哈希码进行比较来判定图片的相似度。该方式处理速度快，但对图片的旋转和颜色的敏感度较高，适用于大规模且纹理相对简单的图片的快速去重。
	// * cnn：通过深度学习技术来提取图像的高级语义特征，如对象、场景和纹理等，这些特征用于比较不同图像之间的相似性。该提取方式鲁棒性较好，对旋转、缩放和变形的敏感度较低，适用于纹理复杂、细节丰富的图片去重。
	ExtractorType string `json:"ExtractorType,omitempty"`

	// 相似度阈值。上传图片数量超过 2 张并执行去重分组时，系统将对原图中满足该阈值的图片进行分组。取值范围为 [0,1]，默认值为 0.84。最多支持两位小数。
	Similarity float64 `json:"Similarity,omitempty"`
}

type GetImageDuplicateDetectionQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageDuplicateDetectionRes struct {

	// REQUIRED
	ResponseMetadata *GetImageDuplicateDetectionResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageDuplicateDetectionResResult `json:"Result,omitempty"`
}

type GetImageDuplicateDetectionResResponseMetadata struct {

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

// GetImageDuplicateDetectionResResult - 视请求的接口而定
type GetImageDuplicateDetectionResResult struct {

	// 回调地址，与请求参数中的Callback相同。具体异步去重信息请参考GetDedupTaskStatus [https://www.volcengine.com/docs/508/138909]接口。具体回调内容请参考回调参数 [https://www.volcengine.com/docs/508/138658#%E5%9B%9E%E8%B0%83%E5%8F%82%E6%95%B0]。
	Callback string `json:"Callback,omitempty"`

	// 分组结果，若输入 2 个以上原图时，将按组返回内容相同的图片，每组的图片个数非固定值。
	Groups map[string][]string `json:"Groups,omitempty"`

	// 图片评分，仅在只上传两张图片并对比时才返回 Score 字段。
	Score string `json:"Score,omitempty"`

	// 异步任务 ID。
	TaskID string `json:"TaskId,omitempty"`
}

type GetImageElementsQuery struct {

	// REQUIRED; 要素类型，取值如下所示：
	// * image：图片要素；
	// * background：背景要素；
	// * mask：蒙版要素。
	Type string `json:"Type" query:"Type"`

	// 分页返回条数。默认 10，最大限制为 100。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移，默认 0，取值为 1 时，表示跳过一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 返回图片 URI 中包含该值的要素列表。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`
}

type GetImageElementsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageElementsResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *GetImageElementsResResult `json:"Result,omitempty"`
}

type GetImageElementsResResponseMetadata struct {

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

// GetImageElementsResResult - title
type GetImageElementsResResult struct {

	// REQUIRED; 要素列表。
	ImageList []*GetImageElementsResResultImageListItem `json:"ImageList"`

	// REQUIRED; 要素总个数。
	Total float64 `json:"Total"`
}

// GetImageElementsResResultImageListItem - 要素
type GetImageElementsResResultImageListItem struct {

	// REQUIRED; 要素添加时间，添加要素时的服务器当地时间。
	AddAt string `json:"AddAt"`

	// REQUIRED; 图片 URI。
	StoreURI string `json:"StoreUri"`
}

type GetImageEnhanceResultBody struct {

	// REQUIRED; 增强模型，取值如下所示：
	// * 0：通用模型
	// * 1：低质专清模型
	Model int `json:"Model"`

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待增强的原图地址的存储 URI 和 URL（公网可访问的 URL）。
	StoreURI string `json:"StoreUri"`

	// 是否不去压缩失真。Model取值为0时选填，支持以下取值：
	// * true：不进行去压缩失真处理
	// * false：（默认）进行去压缩失真处理
	DisableAr bool `json:"DisableAr,omitempty"`

	// 是否不自适应锐化。Model取值为0时选填，支持以下取值：
	// * true：不进行锐化处理
	// * false：（默认）进行锐化处理
	DisableSharp bool `json:"DisableSharp,omitempty"`
}

type GetImageEnhanceResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageEnhanceResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageEnhanceResultResResult `json:"Result,omitempty"`
}

type GetImageEnhanceResultResResponseMetadata struct {

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

// GetImageEnhanceResultResResult - 视请求的接口而定
type GetImageEnhanceResultResResult struct {

	// REQUIRED; 实际执行的增强方法。
	Method string `json:"Method"`

	// REQUIRED; 增强图像的 URI。
	ResURI string `json:"ResUri"`
}

type GetImageEraseModelsQuery struct {

	// 模型。默认取值为0。
	// * 0：自动检测并擦除模型列表。
	// * 1：指定区域擦除模型列表。
	Type int `json:"Type,omitempty" query:"Type"`
}

type GetImageEraseModelsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageEraseModelsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageEraseModelsResResult `json:"Result,omitempty"`
}

type GetImageEraseModelsResResponseMetadata struct {

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

// GetImageEraseModelsResResult - 视请求的接口而定
type GetImageEraseModelsResResult struct {

	// REQUIRED; 模型列表。
	Models []string `json:"Models"`
}

type GetImageEraseResultBody struct {

	// REQUIRED; 修复模型，支持取值如下所示：
	// * 自动检测并擦除类型模型：erasermodelimagex_0.1.0
	// * 指定区域擦除模型： * erasermodelimagex_0.1.0 （推荐）
	// * erasermodelaliab
	Model string `json:"Model"`

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待修复原图的存储 URI 或 URL（公网可访问的 URL）。
	StoreURI string `json:"StoreUri"`

	// 指定区域擦除时图片待修复区域。
	// 不填表示自动检测内容并修复，若所选模型不支持自动检测，则直接返回原图。
	BBox []*GetImageEraseResultBodyBBoxItem `json:"BBox,omitempty"`
}

type GetImageEraseResultBodyBBoxItem struct {

	// REQUIRED; 待修复区域左上角的 X 坐标，取值范围为[0,1]。
	X1 float64 `json:"X1"`

	// REQUIRED; 待修复区域右下角的 X 坐标，取值范围为[0,1]。
	X2 float64 `json:"X2"`

	// REQUIRED; 待修复区域左上角的 Y 坐标，取值范围为[0,1]。
	Y1 float64 `json:"Y1"`

	// REQUIRED; 待修复区域右下角的 Y 坐标，取值范围为[0,1]。
	Y2 float64 `json:"Y2"`

	// 是否模糊匹配，开启文字匹配后必选
	OCRMode int `json:"OCRMode,omitempty"`

	// 匹配的文本
	Text string `json:"Text,omitempty"`

	// 是否开启文字匹配
	UseOCR int `json:"UseOCR,omitempty"`
}

type GetImageEraseResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageEraseResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageEraseResultResResult `json:"Result,omitempty"`
}

type GetImageEraseResultResResponseMetadata struct {

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

// GetImageEraseResultResResult - 视请求的接口而定
type GetImageEraseResultResResult struct {

	// REQUIRED; 修复后结果图的 URI。
	ResURI string `json:"ResUri"`
}

type GetImageFontsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageFontsResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetImageFontsResResult           `json:"Result,omitempty"`
}

type GetImageFontsResResponseMetadata struct {

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

type GetImageFontsResResult struct {

	// REQUIRED; 字体列表。
	Fonts []*GetImageFontsResResultFontsItem `json:"Fonts"`
}

type GetImageFontsResResultFontsItem struct {

	// REQUIRED; 字体值，即字体库字体名称。
	Font string `json:"Font"`

	// REQUIRED; 字体中文名称，如思源黑体。
	Name string `json:"Name"`

	// REQUIRED; 字体资源 URI。
	URI string `json:"Uri"`
}

type GetImageHmExtractTaskInfoBody struct {

	// REQUIRED; 任务id
	TaskID string `json:"TaskId"`
}

type GetImageHmExtractTaskInfoQuery struct {

	// REQUIRED; 提取水印图任务对应的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageHmExtractTaskInfoRes struct {

	// REQUIRED
	ResponseMetadata *GetImageHmExtractTaskInfoResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageHmExtractTaskInfoResResult `json:"Result,omitempty"`
}

type GetImageHmExtractTaskInfoResResponseMetadata struct {

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

// GetImageHmExtractTaskInfoResResult - 视请求的接口而定
type GetImageHmExtractTaskInfoResResult struct {

	// 错误信息
	ErrMsg string `json:"ErrMsg,omitempty"`

	// 水印信息
	HmInfo string `json:"HmInfo,omitempty"`

	// 任务状态
	TaskStatus string `json:"TaskStatus,omitempty"`
}

type GetImageMigrateTasksQuery struct {

	// 分页查询时，显示的每页数据的最大条数。默认值为 10，最大值为 1000。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，用于控制分页查询返回结果的起始位置，以便对数据进行分页展示和浏览。默认值为 0。 :::tip 例如，指定分页条数 Limit = 10，分页偏移量 Offset = 10，表示从查询结果的第 11 条记录开始返回数据，共展示
	// 10 条数据。 :::
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 任务地区（即任务目标服务的地区），缺省时将返回国内列表。取值如下所示：
	// * cn：国内
	// * sg：新加坡
	Region string `json:"Region,omitempty" query:"Region"`

	// 迁移的目标服务 ID。
	ServiceID string `json:"ServiceId,omitempty" query:"ServiceId"`

	// 任务状态，填入多个时使用英文逗号分隔。取值如下所示：
	// * Initial：创建中
	// * Running：运行中
	// * Done：全部迁移完成
	// * Partial：部分迁移完成
	// * Failed：迁移失败
	// * Terminated：中止
	Status string `json:"Status,omitempty" query:"Status"`

	// 任务 ID。
	TaskID string `json:"TaskId,omitempty" query:"TaskId"`

	// 返回任务名称中包含该值的迁移任务信息。
	TaskNamePtn string `json:"TaskNamePtn,omitempty" query:"TaskNamePtn"`
}

type GetImageMigrateTasksRes struct {

	// REQUIRED
	ResponseMetadata *GetImageMigrateTasksResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageMigrateTasksResResult `json:"Result,omitempty"`
}

type GetImageMigrateTasksResResponseMetadata struct {

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

// GetImageMigrateTasksResResult - 视请求的接口而定
type GetImageMigrateTasksResResult struct {

	// REQUIRED; 迁移任务列表
	Tasks []*GetImageMigrateTasksResResultTasksItem `json:"Tasks"`

	// REQUIRED; 总任务数
	Total int `json:"Total"`
}

// GetImageMigrateTasksResResultTasksItem - 迁移任务列表
type GetImageMigrateTasksResResultTasksItem struct {

	// REQUIRED; 目标信息
	Dst *GetImageMigrateTasksResResultTasksItemDst `json:"Dst"`

	// REQUIRED; 任务 ID
	ID string `json:"ID"`

	// REQUIRED; 任务名称
	Name string `json:"Name"`

	// REQUIRED; 迁移进度信息
	Progress *GetImageMigrateTasksResResultTasksItemProgress `json:"Progress"`

	// REQUIRED; 运行时长信息
	Run []*GetImageMigrateTasksResResultTasksPropertiesItemsItem `json:"Run"`

	// REQUIRED; 迁移策略
	RunStrategy *GetImageMigrateTasksResResultTasksItemRunStrategy `json:"RunStrategy"`

	// REQUIRED; 源信息
	Source *GetImageMigrateTasksResResultTasksItemSource `json:"Source"`

	// REQUIRED; 任务状态。取值如下所示：
	// * Initial：创建中
	// * Running：运行中
	// * Done：全部迁移完成
	// * Partial：部分迁移完成
	// * Failed：迁移失败
	// * Terminated：中止
	Status string `json:"Status"`

	// REQUIRED; 转码配置
	Transcode *GetImageMigrateTasksResResultTasksItemTranscode `json:"Transcode"`
}

// GetImageMigrateTasksResResultTasksItemDst - 目标信息
type GetImageMigrateTasksResResultTasksItemDst struct {

	// REQUIRED; 目标 key 前缀
	Prefix string `json:"Prefix"`

	// REQUIRED; 服务 ID
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 服务名称
	ServiceName string `json:"ServiceName"`

	// REQUIRED; 源 Bucket 名称保留规则
	SkipBucket bool `json:"SkipBucket"`

	// REQUIRED; 上传配置。取值如下所示：
	// * 0：直接覆盖同名文件
	// * 1：增加文件名后缀，后缀为 任务 ID
	// * 2：跳过同名文件，即不做迁移
	UploadConf int `json:"UploadConf"`
}

// GetImageMigrateTasksResResultTasksItemProgress - 迁移进度信息
type GetImageMigrateTasksResResultTasksItemProgress struct {

	// REQUIRED; 失败错误码。仅当 Status=Failed 时有值。
	ErrCode int `json:"ErrCode"`

	// REQUIRED; 失败原因。仅当 Status=Failed 时有值。
	ErrMsg string `json:"ErrMsg"`

	// REQUIRED; 迁移失败文件数
	FailCnt int `json:"FailCnt"`

	// REQUIRED; 迁移成功文件量，单位为 byte
	SuccessAmount int `json:"SuccessAmount"`

	// REQUIRED; 迁移成功文件数
	SuccessCnt int `json:"SuccessCnt"`

	// REQUIRED; 迁移文件总量，单位为 byte
	TotalAmount int `json:"TotalAmount"`

	// REQUIRED; 总文件数
	TotalCnt int `json:"TotalCnt"`
}

// GetImageMigrateTasksResResultTasksItemRunStrategy - 迁移策略
type GetImageMigrateTasksResResultTasksItemRunStrategy struct {

	// REQUIRED; 源下载 QPS 限制。如值不为空，则长度必须为 24，表示一天 24 小时内各小时的 QPS 限制值。
	// * 取值为负值时，表示无限制
	// * 取值为 0 时，表示对应时间不允许迁移
	ReadQPS []int `json:"ReadQps"`

	// REQUIRED; 源下载流量限制。单位为 Byte。如值不为空，则长度必须为24，表示一天 24 小时内各小时的流量限制值。
	// * 取值为负值时，表示无限制
	// * 取值为 0 时，表示对应时间不允许迁移
	ReadRate []int `json:"ReadRate"`
}

// GetImageMigrateTasksResResultTasksItemSource - 源信息
type GetImageMigrateTasksResResultTasksItemSource struct {

	// REQUIRED; Access Key
	AK string `json:"AK"`

	// REQUIRED; 源 bucket
	Bucket string `json:"Bucket"`

	// REQUIRED; 迁移源云服务商 CDN 域名
	CdnHost string `json:"CdnHost"`

	// REQUIRED; 源 Endpoint
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 迁移前缀列表
	Prefix []string `json:"Prefix"`

	// REQUIRED; 迁移正则表达式列表
	Regex []string `json:"Regex"`

	// REQUIRED; 源 bucket 地区
	Region string `json:"Region"`

	// REQUIRED; Secret Key
	SK string `json:"SK"`

	// REQUIRED; 是否丢弃源 Header，取值如下所示：
	// * true：丢弃源 Header
	// * false：不丢弃源 Header
	SkipHeader bool `json:"SkipHeader"`

	// REQUIRED; 指定迁移结束时间点，为迁移结束当地服务器时间。表示仅迁移该时间段内新增或变更的文件。
	TimeEnd string `json:"TimeEnd"`

	// REQUIRED; 指定迁移开始时间点，为迁移开始当地服务器时间。表示仅迁移该时间段内新增或变更的文件。
	TimeStart string `json:"TimeStart"`

	// REQUIRED; 源服务商
	Vendor string `json:"Vendor"`
}

// GetImageMigrateTasksResResultTasksItemTranscode - 转码配置
type GetImageMigrateTasksResResultTasksItemTranscode struct {

	// REQUIRED; 包含透明通道的图片是否编码为降级格式。取值如下所示：
	// * true：降级
	// * false：不降级
	AlphaDemotion bool `json:"AlphaDemotion"`

	// REQUIRED; 降级编码格式。支持的格式有 png、jpeg、heic、avif、webp、vvic。
	DemotionFmt string `json:"DemotionFmt"`

	// REQUIRED; 目标转码格式。支持的格式有 png、jpeg、heic、avif、webp、vvic。
	Format string `json:"Format"`

	// REQUIRED; 转码质量参数。对于 PNG 为无损压缩，其他格式下其值越小，压缩率越高，画质越差。
	Quality int `json:"Quality"`
}

type GetImageMigrateTasksResResultTasksPropertiesItemsItem struct {

	// REQUIRED; 迁移任务结束时间
	DoneAt string `json:"DoneAt"`

	// REQUIRED; 迁移任务开始时间
	StartAt string `json:"StartAt"`
}

type GetImageMonitorRulesQuery struct {
	AppID string `json:"AppId,omitempty" query:"AppId"`

	// 分页条数。默认值为 10，取值范围为（0，100]。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 告警名称，以正则表达式进行筛选匹配。缺省时默认获取所有报警规则。
	NamePtn string `json:"NamePtn,omitempty" query:"NamePtn"`

	// 分页偏移量。默认值为 0，表示从最新一个开始获取。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 报警规则 ID，按照指定 ID 返回对应报警规则。缺省时默认获取所有报警规则。
	RuleID string `json:"RuleId,omitempty" query:"RuleId"`
}

type GetImageMonitorRulesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageMonitorRulesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageMonitorRulesResResult `json:"Result,omitempty"`
}

type GetImageMonitorRulesResResponseMetadata struct {

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

// GetImageMonitorRulesResResult - 视请求的接口而定
type GetImageMonitorRulesResResult struct {

	// REQUIRED
	MonitorRules []*GetImageMonitorRulesResResultMonitorRulesItem `json:"MonitorRules"`

	// REQUIRED; 规则总数
	Total int `json:"Total"`
}

// GetImageMonitorRulesResResultMonitorRulesItem - 告警规则
type GetImageMonitorRulesResResultMonitorRulesItem struct {

	// REQUIRED; 监控的应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppID。
	Appid string `json:"Appid"`

	// REQUIRED; 监测规则。
	Cond *GetImageMonitorRulesResResultMonitorRulesItemCond `json:"Cond"`

	// REQUIRED; 规则创建时间，ISO 8601 格式时间戳。
	CreateAt string `json:"CreateAt"`

	// REQUIRED; 创建后是否立即开启告警，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"Enabled"`

	// REQUIRED; 监控频率，单位为分钟。取值如下所示：
	// * 5
	// * 10
	// * 20
	// * 30
	// * 40
	// * 50
	Frequency int `json:"Frequency"`

	// REQUIRED; 告警级别，取值如下所示：
	// * warn：警告
	// * error：错误
	// * fatal：致命
	Level string `json:"Level"`

	// REQUIRED; 自定义告警规则名称
	Name string `json:"Name"`

	// REQUIRED; 告警通知配置。
	Notification *GetImageMonitorRulesResResultMonitorRulesItemNotification `json:"Notification"`

	// REQUIRED; 监控阶段，取值如下所示：
	// * upload：图片上传-上传 1.0
	// * uploadv2：图片上传-上传 2.0
	// * cdn：图片加载-下行网络监控
	// * client：图片加载-客户端传状态监控
	// * sensible：图片加载-感知指标监控
	Phase string `json:"Phase"`

	// REQUIRED; 报警规则 ID
	RuleID string `json:"RuleId"`

	// REQUIRED; 规则更新时间，ISO 8601 格式时间戳。
	UpdateAt string `json:"UpdateAt"`

	// 维度过滤条件，具体参数请见Filter [https://www.volcengine.com/docs/508/1112182#filter]。用于指定需要告警提示的维度配置。
	Filter *GetImageMonitorRulesResResultMonitorRulesItemFilter `json:"Filter,omitempty"`

	// 拆分维度，由公共拆分维度 [https://www.volcengine.com/docs/508/1113944]和自定义拆分维度 [https://www.volcengine.com/docs/508/34554]组合而成。
	GroupBy string `json:"GroupBy,omitempty"`

	// 监控平台，取值如下所示：
	// * iOS
	// * Android
	// * WEB
	OS string `json:"OS,omitempty"`
}

// GetImageMonitorRulesResResultMonitorRulesItemCond - 监测规则。
type GetImageMonitorRulesResResultMonitorRulesItemCond struct {

	// REQUIRED
	ItemCond []*Components3Ic6Z9SchemasGetimagemonitorrulesresPropertiesResultPropertiesMonitorrulesItemsPropertiesCondPropertiesItemcondItems `json:"ItemCond"`

	// REQUIRED; 多条监控规则之间的逻辑关系，取值如下所示：
	// * and：且。表示有多条监控规则时，需满足所有监控规则才会触发告警通知。
	// * or：或。表示有多条监控规则时，满足其中一条监控规则就会触发告警通知。
	LogicOp string `json:"LogicOp"`
}

// GetImageMonitorRulesResResultMonitorRulesItemFilter - 维度过滤条件，具体参数请见Filter [https://www.volcengine.com/docs/508/1112182#filter]。用于指定需要告警提示的维度配置。
type GetImageMonitorRulesResResultMonitorRulesItemFilter struct {

	// REQUIRED; 过滤条件
	DimFilter []*Components1WypgicSchemasGetimagemonitorrulesresPropertiesResultPropertiesMonitorrulesItemsPropertiesFilterPropertiesDimfilterItems `json:"DimFilter"`

	// REQUIRED; 过滤条件之间的逻辑关系，取值如下所示：
	// * and：和
	// * or：或
	LogicOp string `json:"LogicOp"`
}

// GetImageMonitorRulesResResultMonitorRulesItemNotification - 告警通知配置。
type GetImageMonitorRulesResResultMonitorRulesItemNotification struct {

	// REQUIRED; 通知内容模板，模板中变量格式为$Name$。Name 取值如下所示：
	// * 报警名称
	// * 报警级别
	// * 报警App
	// * 报警平台
	// * 报警时间
	// * 报警内容
	Content string `json:"Content"`

	// REQUIRED; 通知方式，仅支持取值http_callback，表示回调。
	Mode []string `json:"Mode"`

	// REQUIRED; 沉默周期，单位为分钟。告警发生后，若未恢复正常，则会间隔一个沉默周期后再次重复发送一次告警通知。取值如下所示：
	// * 0
	// * 30
	// * 60
	// * 360
	SilentDur int `json:"SilentDur"`

	// REQUIRED; 告警通知标题
	Title string `json:"Title"`

	// 回调地址，Mode包含http_callback时，为必填。
	CallbackURL string `json:"CallbackUrl,omitempty"`
}

type GetImageOCRV2Body struct {

	// REQUIRED; 图片OCR识别场景，取值如下所示：
	// * general：通用场景，用于通用印刷体场景识别文本信息。
	// * license：营业执照场景，用于识别营业执照中社会信用代码等文本信息。
	// * instrument：设备识别场景，用于一些设备显示文字识别。
	// * defect：缺陷检测场景
	// :::warning 当前仅支持识别图片中简体中文和简体英文这两种文本信息。 :::
	Scene string `json:"Scene"`

	// REQUIRED; 待识别图片文件的存储 URI。
	StoreURI string `json:"StoreUri"`

	// 定制化保留字段，如果是正常调用忽略该字段，若为定制化需求则需要和算法开发者对齐调用方式
	Extra *GetImageOCRV2BodyExtra `json:"Extra,omitempty"`

	// 待识别的图片 URL，满足公网可访问。仅当 StoreUri 为空时取值有效，两者都为空时报错。
	ImageURL string `json:"ImageUrl,omitempty"`

	// 待识别的设备名称，仅当 Scene 为 Instrument 时，配置有效。取值如下所示：
	// * freezing-point-tester：冰点仪
	// * brake-fluid-tester：制动液测试仪
	// * thermometer： 温度计
	// * oil-tester：机油仪
	InstrumentName string `json:"InstrumentName,omitempty"`
}

// GetImageOCRV2BodyExtra - 定制化保留字段，如果是正常调用忽略该字段，若为定制化需求则需要和算法开发者对齐调用方式
type GetImageOCRV2BodyExtra struct {

	// 默认为False，不开启Extra
	Enable bool `json:"Enable,omitempty"`

	// 算子的输入参数
	MMServiceInput string `json:"MMServiceInput,omitempty"`

	// 算子名称
	MMServiceName string `json:"MMServiceName,omitempty"`

	// 算子版本
	MMServiceVersion string `json:"MMServiceVersion,omitempty"`
}

type GetImageOCRV2Query struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageOCRV2Res struct {

	// REQUIRED
	ResponseMetadata *GetImageOCRV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageOCRV2ResResult `json:"Result,omitempty"`
}

type GetImageOCRV2ResResponseMetadata struct {

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

// GetImageOCRV2ResResult - 视请求的接口而定
type GetImageOCRV2ResResult struct {

	// REQUIRED; 指定的识别场景。
	Scene string `json:"Scene"`

	// 缺陷检测场景识别结果。
	DefectResult []*GetImageOCRV2ResResultDefectResultItem `json:"DefectResult,omitempty"`

	// 定制化保留字段，默认为空 如果输入参数内Extra.Enable设置为True，该字段返回的结果依赖于算子结果
	Extra *GetImageOCRV2ResResultExtra `json:"Extra,omitempty"`

	// 定制化保留字段，默认为空 如果输入参数内Extra.Enable设置为True，该字段返回的结果依赖于算子结果
	ExtraOutput *GetImageOCRV2ResResultExtraOutput `json:"ExtraOutput,omitempty"`

	// 通用场景识别结果。
	GeneralResult []*GetImageOCRV2ResResultGeneralResultItem `json:"GeneralResult,omitempty"`

	// 设备识别场景识别的文本信息。根据您识别设备的不同取值含义分别为：
	// * 当设备为冰点仪时，表示液体冰点温度值，例如44.89。
	//
	//
	// * 当设备为制动液测试仪时，表示制动液含水量程度。例如Middle表示中等。
	//
	//
	// * 当设备为温度计时，表示测量的温度值。例如37.8。
	//
	//
	// * 当设备为机油仪时，表示机油的剩余使用寿命的估计值，例如93.8。
	InstructmentResult string `json:"InstructmentResult,omitempty"`

	// 营业执照场景识别结果。
	LicenseResult map[string]*ComponentsK7Ou2VSchemasGetimageocrv2ResPropertiesResultPropertiesLicenseresultAdditionalproperties `json:"LicenseResult,omitempty"`
}

type GetImageOCRV2ResResultDefectResultItem struct {

	// 识别出的缺陷类别编号，当前仅在取值为 0 时，表示漏油。 :::tip 其他缺陷类别编号识别还在训练增加中。 :::
	ClassLabel int `json:"ClassLabel,omitempty"`

	// 识别出的文本块的坐标位置。
	Location []int `json:"Location,omitempty"`
}

// GetImageOCRV2ResResultExtra - 定制化保留字段，默认为空 如果输入参数内Extra.Enable设置为True，该字段返回的结果依赖于算子结果
type GetImageOCRV2ResResultExtra struct {

	// 算子服务返回的json序列化结果
	Info string `json:"Info,omitempty"`

	// 状态对应信息， 0: 成功 -1: 图片载入失败 -2: 图片检测失败
	Message string `json:"Message,omitempty"`

	// 算子服务处理状态，0: 成功 -1: 图片载入失败 -2: 图片检测失败
	Status int `json:"Status,omitempty"`
}

// GetImageOCRV2ResResultExtraOutput - 定制化保留字段，默认为空 如果输入参数内Extra.Enable设置为True，该字段返回的结果依赖于算子结果
type GetImageOCRV2ResResultExtraOutput struct {

	// 算子服务返回的结果
	Info interface{} `json:"Info,omitempty"`

	// 状态对应信息， 0: 成功 -1: 图片载入失败 -2: 图片检测失败
	Message string `json:"Message,omitempty"`

	// 算子服务处理状态，0: 成功 -1: 图片载入失败 -2: 图片检测失败
	Status int `json:"Status,omitempty"`
}

type GetImageOCRV2ResResultGeneralResultItem struct {

	// REQUIRED; 文本置信度。识别出的内容可信程度，值越大内容越准确。
	Confidence string `json:"Confidence"`

	// REQUIRED; 识别的通用文本信息。
	Content string `json:"Content"`

	// REQUIRED; 文本块的坐标位置。
	Location [][]int `json:"Location"`
}

type GetImagePSDetectionBody struct {

	// REQUIRED; 原图的存储 URI。
	ImageURI string `json:"ImageUri"`

	// 调用方法，默认为调用全部方法，若参数不为all，则其他方法仅可选一种。取值如下所示：
	// * all：调用所有方法
	// * ela：误差水平分析方法
	// * na：噪声分析方法
	// * he：基于颜色分布直方图均化的图像增强分析方法
	Method string `json:"Method,omitempty"`
}

type GetImagePSDetectionQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImagePSDetectionRes struct {

	// REQUIRED
	ResponseMetadata *GetImagePSDetectionResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImagePSDetectionResResult `json:"Result,omitempty"`
}

type GetImagePSDetectionResResponseMetadata struct {

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

// GetImagePSDetectionResResult - 视请求的接口而定
type GetImagePSDetectionResResult struct {

	// REQUIRED; 误差水平分析方法对应的图片。
	ElaImage string `json:"ElaImage"`

	// REQUIRED; 基于颜色分布直方图均化的图像增强方法对应的图片。
	HeImage string `json:"HeImage"`

	// REQUIRED; 置信度标签，取值如下所示：
	// * 1：正常
	// * 0：高风险
	Label int `json:"Label"`

	// REQUIRED; 噪声分析方法对应的图片。
	NaImage string `json:"NaImage"`

	// REQUIRED; 置信度，取值范围为[0,1]。
	// * 小于 0.3 为 PS 处理图；
	// * 大于 0.7 为真图；
	// * 取值范围为[0.3,0.7]之间的图片则为模型无法判断。
	Score float64 `json:"Score"`
}

type GetImageQualityBody struct {

	// REQUIRED; 图片存储 Uri 或访问 URL。
	// * 图片 Uri 格式，例如：tos-example/7a7979974.jpeg
	// * 图片 URL 格式，例如：https://example.org/tos-example/7a7979974.jpeg~tplv.png :::tip 若传 URL，必须保证 URL 公网可访问。 :::
	ImageURL string `json:"ImageUrl"`

	// REQUIRED; 评估工具。指定多个评估工具时使用英文逗号分隔，当前支持以下工具：
	// * nr_index
	// * vqscore
	// * advcolor
	// * blockiness
	// * noise
	// * aesmod
	// * blur
	// * cg
	// * contrast
	// * texture
	// * brightness
	// * overexposure
	// * hue
	// * saturation
	// * psnr
	// * ssim
	// * vmaf
	// * green
	// * cmartifacts :::tip nr_index 工具支持评估 contrast、brightness 等多个维度。您也可以单独指定各维度，获取指定维度估值。 :::
	VqType string `json:"VqType"`

	// 指定服务下的评估参照图片存储 Uri 或访问 URL，用于和 ImageUrl 图片进行特定维度的对比。 说明：当 VqType 中包含 psnr、ssim、vmaf等任一字段时，该字段为必填，否则上述评估指标无法正常输出结果。
	ImageURLRef string `json:"ImageUrlRef,omitempty"`
}

type GetImageQualityQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageQualityRes struct {

	// REQUIRED
	ResponseMetadata *GetImageQualityResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageQualityResResult `json:"Result"`
}

type GetImageQualityResResponseMetadata struct {

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

type GetImageQualityResResult struct {

	// REQUIRED
	FrScoreResult *GetImageQualityResResultFrScoreResult `json:"FrScoreResult"`

	// REQUIRED
	FrScores *GetImageQualityResResultFrScores `json:"FrScores"`

	// REQUIRED
	NrScoreResult *GetImageQualityResResultNrScoreResult `json:"NrScoreResult"`

	// REQUIRED
	NrScores *GetImageQualityResResultNrScores `json:"NrScores"`

	// REQUIRED
	VqType string `json:"VqType"`
}

type GetImageQualityResResultFrScoreResult struct {

	// REQUIRED
	Psnr float64 `json:"Psnr"`

	// REQUIRED
	Ssim float64 `json:"Ssim"`

	// REQUIRED
	Vmaf float64 `json:"Vmaf"`
}

type GetImageQualityResResultFrScores struct {

	// REQUIRED
	Psnr float64 `json:"psnr"`

	// REQUIRED
	Ssim float64 `json:"ssim"`

	// REQUIRED
	Vmaf float64 `json:"vmaf"`
}

type GetImageQualityResResultNrScoreResult struct {

	// REQUIRED
	AdvColor float64 `json:"AdvColor"`

	// REQUIRED
	Aesthetic float64 `json:"Aesthetic"`

	// REQUIRED
	Blockiness float64 `json:"Blockiness"`

	// REQUIRED
	Blur float64 `json:"Blur"`

	// REQUIRED
	Brightness float64 `json:"Brightness"`

	// REQUIRED
	Cg float64 `json:"Cg"`

	// REQUIRED
	CmArtifact float64 `json:"CmArtifact"`

	// REQUIRED
	Contrast float64 `json:"Contrast"`

	// REQUIRED
	Green float64 `json:"Green"`

	// REQUIRED
	Hue float64 `json:"Hue"`

	// REQUIRED
	Noise float64 `json:"Noise"`

	// REQUIRED
	OverExposure float64 `json:"OverExposure"`

	// REQUIRED
	Saturation float64 `json:"Saturation"`

	// REQUIRED; 文字质量分数
	TextQualityScore float64 `json:"TextQualityScore"`

	// REQUIRED
	Texture float64 `json:"Texture"`

	// REQUIRED
	VqScore float64 `json:"VqScore"`
}

type GetImageQualityResResultNrScores struct {

	// REQUIRED
	Advcolor float64 `json:"advcolor"`

	// REQUIRED
	Aesthetic float64 `json:"aesthetic"`

	// REQUIRED
	Blockiness float64 `json:"blockiness"`

	// REQUIRED
	Blur float64 `json:"blur"`

	// REQUIRED
	Brightness float64 `json:"brightness"`

	// REQUIRED
	Cg float64 `json:"cg"`

	// REQUIRED
	Cmartifact float64 `json:"cmartifact"`

	// REQUIRED
	Colorfulness float64 `json:"colorfulness"`

	// REQUIRED
	Contrast float64 `json:"contrast"`

	// REQUIRED
	Green float64 `json:"green"`

	// REQUIRED
	Hue float64 `json:"hue"`

	// REQUIRED
	Noise float64 `json:"noise"`

	// REQUIRED
	Overexposure float64 `json:"overexposure"`

	// REQUIRED
	Saturation float64 `json:"saturation"`

	// REQUIRED
	Texture float64 `json:"texture"`

	// REQUIRED
	Vqscore float64 `json:"vqscore"`
}

type GetImageServiceQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageServiceRes struct {

	// REQUIRED
	ResponseMetadata *GetImageServiceResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetImageServiceResResult           `json:"Result,omitempty"`
}

type GetImageServiceResResponseMetadata struct {

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

type GetImageServiceResResult struct {

	// REQUIRED; 服务的授权 Bucket 列表。
	AllowBkts []string `json:"AllowBkts"`

	// REQUIRED; 是否允许配置其他镜像站类型，取值如下所示：
	// * true：是
	// * false：否
	AllowMirrorTypes bool `json:"AllowMirrorTypes"`

	// REQUIRED; 是否开启精简 URL，取值如下所示：
	// * true：开启
	// * false：关闭
	CompactURL bool `json:"CompactURL"`

	// REQUIRED; 服务创建时间，即创建时当地服务器时间。
	CreateAt string `json:"CreateAt"`

	// REQUIRED; 绑定域名的相关信息。
	DomainInfos []*GetImageServiceResResultDomainInfosItem `json:"DomainInfos"`

	// REQUIRED; 服务是否已经配置鉴权 key，取值如下所示：
	// * true：已配置
	// * false：未配置
	HasSigkey bool `json:"HasSigkey"`

	// REQUIRED; 是否开启自定义处理样式，取值如下所示：
	// * true：是
	// * false：否
	ImageY bool `json:"ImageY"`

	// REQUIRED; 自定义处理相关配置
	ImageYAttribute *GetImageServiceResResultImageYAttribute `json:"ImageYAttribute"`

	// REQUIRED; 镜像回源配置，默认关闭。
	Mirror *GetImageServiceResResultMirror `json:"Mirror"`

	// REQUIRED; 是否开启源地址访问，取值如下所示：
	// * true：开启
	// * false：关闭
	ObjectAccess bool `json:"ObjectAccess"`

	// REQUIRED; 主鉴权 Key。
	PrimaryKey string `json:"PrimaryKey"`

	// REQUIRED
	ResourceLimitedVisit *GetImageServiceResResultResourceLimitedVisit `json:"ResourceLimitedVisit"`

	// REQUIRED; 备鉴权 Key。
	SecondaryKey string `json:"SecondaryKey"`

	// REQUIRED; 指定的服务 ID。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 服务名称。
	ServiceName string `json:"ServiceName"`

	// REQUIRED; 服务地域，取值如下所示：
	// * cn：中国
	// * va：美东
	// * sg：新加坡
	ServiceRegion string `json:"ServiceRegion"`

	// REQUIRED; 服务状态。状态分为未审核、审核未通过、正常、禁用。 :::tip
	// * 只有服务状态为正常时，该服务才可用。
	// * 如果是其他异常状态，请参考服务管理 [https://www.volcengine.com/docs/508/8086]进行处理。 :::
	ServiceStatus string `json:"ServiceStatus"`

	// REQUIRED; 服务类型，取值如下所示：
	// * StaticRc：素材托管服务
	// * Image：图片处理服务
	ServiceType string `json:"ServiceType"`

	// REQUIRED; 资源配置。
	Storage *GetImageServiceResResultStorage `json:"Storage"`

	// REQUIRED; 该服务的图片模板固定前缀。
	TemplatePrefix string `json:"TemplatePrefix"`

	// 事件通知规则
	EventRules []*GetImageServiceResResultEventRulesItem `json:"EventRules,omitempty"`

	// 服务绑定的项目，默认为 default。
	ProjectName string `json:"ProjectName,omitempty"`

	// 服务绑定的标签。
	ResourceTags []*GetImageServiceResResultResourceTagsItem `json:"ResourceTags,omitempty"`

	// 用于保护「数据加密密钥」的密钥，只有加密上传的图片需要做处理时需要申请。
	RsaPublicKey string `json:"RsaPublicKey,omitempty"`

	// 存储降冷策略
	StorageRules []*GetImageServiceResResultStorageRulesItem `json:"StorageRules,omitempty"`

	// 存储规则。
	StorageRulesV2    []*GetImageServiceResResultStorageRulesV2Item `json:"StorageRulesV2,omitempty"`
	StorageVersioning int                                           `json:"StorageVersioning,omitempty"`

	// 是否开启覆盖上传，取值如下所示：
	// * true：开启
	// * false：关闭
	UploadOverwrite bool `json:"UploadOverwrite,omitempty"`

	// 绑定的点播空间信息
	VodSpace *GetImageServiceResResultVodSpace `json:"VodSpace,omitempty"`
}

type GetImageServiceResResultDomainInfosItem struct {

	// REQUIRED; 域名解析到的 cname。
	CNAME string `json:"CNAME"`

	// REQUIRED; 绑定的域名。
	DomainName string `json:"DomainName"`

	// REQUIRED; 是否是默认域名，取值如下所示：
	// * true：默认域名
	// * false：非默认域名
	IsDefault bool `json:"IsDefault"`

	// REQUIRED; 域名状态。
	Status string `json:"Status"`

	// REQUIRED; 是否开启鉴权
	URLAuth bool `json:"UrlAuth"`
}

type GetImageServiceResResultEventRulesItem struct {

	// REQUIRED; 事件触发时接收回调的回调 URL。
	CallbackURL string `json:"CallbackUrl"`

	// REQUIRED; 规则是否被启用，取值如下所示：
	// * true：是
	// * false：否
	Enable bool `json:"Enable"`

	// REQUIRED; 事件类型。取值如下所示：
	// * Upload：上传文件
	// * Delete：删除文件
	// * Mirror：镜像回源
	// * Migrate：数据迁移
	// * OffTrans：离线转码（仅图像处理服务可配置）
	// * TplStore：模板持久化存储（仅图像处理服务可配置）
	EventType []string `json:"EventType"`

	// REQUIRED; 规则 ID
	ID string `json:"Id"`

	// REQUIRED; 匹配规则的正则表达式。
	MatchRule string `json:"MatchRule"`
}

// GetImageServiceResResultImageYAttribute - 自定义处理相关配置
type GetImageServiceResResultImageYAttribute struct {

	// REQUIRED; 是否开启原图保护，取值如下所示：
	// * true：开启
	// * false：关闭
	ImageProtect bool `json:"ImageProtect"`

	// REQUIRED; 样式分割符
	ImageStyleSeparators []string `json:"ImageStyleSeparators"`

	// REQUIRED
	QnCosPreference string `json:"QnCosPreference"`

	// REQUIRED
	QueryStyleCombine bool `json:"QueryStyleCombine"`
}

// GetImageServiceResResultMirror - 镜像回源配置，默认关闭。
type GetImageServiceResResultMirror struct {

	// REQUIRED; 镜像回源下载原图时，携带的 HTTP 头部，键值都为 String 类型。
	Headers map[string]string `json:"Headers"`

	// REQUIRED; 镜像回源域名。
	Host string `json:"Host"`

	// REQUIRED; 带权重回源域名，key 为 String 类型时，代表镜像回源域名；value 为 Integer 类型时，代表域名权重。
	Hosts map[string]int `json:"Hosts"`

	// REQUIRED; 下载图片的协议，支持取值：http、https。
	Schema string `json:"Schema"`

	// REQUIRED; 镜像源 URI，其中图片名用 %s 占位符替代，比如/obj/%s。
	Source string `json:"Source"`
}

type GetImageServiceResResultResourceLimitedVisit struct {

	// REQUIRED
	AllowDomains []string `json:"AllowDomains"`

	// REQUIRED
	Enable bool `json:"Enable"`
}

type GetImageServiceResResultResourceTagsItem struct {

	// REQUIRED; 标签键
	Key string `json:"Key"`

	// REQUIRED; 标签值
	Value string `json:"Value"`
}

// GetImageServiceResResultStorage - 资源配置。
type GetImageServiceResResultStorage struct {

	// REQUIRED; 是否支持任意文件格式上传，取值如下所示：
	// * true：支持
	// * false：不支持
	AllTypes bool `json:"AllTypes"`

	// REQUIRED; 存储 Bucket 名称。
	BktName string `json:"BktName"`

	// REQUIRED; 保存时间，单位为秒。
	TTL int `json:"TTL"`
}

type GetImageServiceResResultStorageRulesItem struct {

	// REQUIRED; 是否启用策略，取值如下所示：
	// * true：是
	// * false：否
	Enable bool `json:"Enable"`

	// REQUIRED; 策略类型，固定取值Upload，表示按上传时间。
	Event string `json:"Event"`

	// * DELETE，删除
	// * IA，转为低频存储
	// * ARCHIVE_FR，转为归档闪回存储
	// * COLD_ARCHIVE，转为冷归档存储
	// * ARCHIVE，转为归档存储
	Action string `json:"Action,omitempty"`

	// 天数，与Date冲突
	Day              int    `json:"Day,omitempty"`
	NonCurrentAction string `json:"NonCurrentAction,omitempty"`

	// 非当前天数。取值范围为 [ ]，单位为，默认值为``。
	NonCurrentDay int `json:"NonCurrentDay,omitempty"`

	// 前缀。
	Prefix string `json:"Prefix,omitempty"`
}

type GetImageServiceResResultStorageRulesV2Item struct {

	// REQUIRED; 是否启用。
	Enable bool `json:"Enable"`

	// REQUIRED; 事件类型。
	Event string `json:"Event"`

	// REQUIRED; 请提供参数的名字（Id）和类型（string），我会根据输入生成参数的一句话描述。
	ID string `json:"Id"`

	// REQUIRED; 参数的前缀和类型。
	Prefix string `json:"Prefix"`

	// 中止未完成的分段上传配置。
	AbortIncompleteMultipartUpload *GetImageServiceResResultStorageRulesV2ItemAbortIncompleteMultipartUpload `json:"AbortIncompleteMultipartUpload,omitempty"`

	// 过期时间。
	Expiration *GetImageServiceResResultStorageRulesV2ItemExpiration `json:"Expiration,omitempty"`

	// 过滤条件。
	Filter *GetImageServiceResResultStorageRulesV2ItemFilter `json:"Filter,omitempty"`

	// 非当前版本的过期时间。
	NonCurrentVersionExpiration *GetImageServiceResResultStorageRulesV2ItemNonCurrentVersionExpiration `json:"NonCurrentVersionExpiration,omitempty"`

	// 非当前版本转换规则。
	NonCurrentVersionTransitions []*Components1AkatydSchemasGetimageserviceresPropertiesResultPropertiesStoragerulesv2ItemsPropertiesNoncurrentversiontransitionsItems `json:"NonCurrentVersionTransitions,omitempty"`

	// 状态转换。
	Transitions []*GetImageServiceResResultStorageRulesV2PropertiesItemsItem `json:"Transitions,omitempty"`
}

// GetImageServiceResResultStorageRulesV2ItemAbortIncompleteMultipartUpload - 中止未完成的分段上传配置。
type GetImageServiceResResultStorageRulesV2ItemAbortIncompleteMultipartUpload struct {

	// 距启动后的天数。
	DaysAfterInitiation int `json:"DaysAfterInitiation,omitempty"`
}

// GetImageServiceResResultStorageRulesV2ItemExpiration - 过期时间。
type GetImageServiceResResultStorageRulesV2ItemExpiration struct {

	// 日期。
	Date string `json:"Date,omitempty"`

	// 天数。
	Days int `json:"Days,omitempty"`
}

// GetImageServiceResResultStorageRulesV2ItemFilter - 过滤条件。
type GetImageServiceResResultStorageRulesV2ItemFilter struct {

	// 是否包含大于等于。
	GreaterThanIncludeEqual string `json:"GreaterThanIncludeEqual,omitempty"`

	// 小于等于条件。
	LessThanIncludeEqual string `json:"LessThanIncludeEqual,omitempty"`

	// 对象大小大于。取值范围为 [ ]，单位为，默认值为``。
	ObjectSizeGreaterThan int `json:"ObjectSizeGreaterThan,omitempty"`

	// 对象大小上限。
	ObjectSizeLessThan int `json:"ObjectSizeLessThan,omitempty"`
}

// GetImageServiceResResultStorageRulesV2ItemNonCurrentVersionExpiration - 非当前版本的过期时间。
type GetImageServiceResResultStorageRulesV2ItemNonCurrentVersionExpiration struct {

	// 非当前日期。
	NonCurrentDate string `json:"NonCurrentDate,omitempty"`

	// 非当前天数。
	NonCurrentDays int `json:"NonCurrentDays,omitempty"`
}

type GetImageServiceResResultStorageRulesV2PropertiesItemsItem struct {

	// REQUIRED; 存储类型。
	StorageClass string `json:"StorageClass"`

	// 日期字符串。
	Date string `json:"Date,omitempty"`

	// 天数。
	Days int `json:"Days,omitempty"`
}

// GetImageServiceResResultVodSpace - 绑定的点播空间信息
type GetImageServiceResResultVodSpace struct {

	// REQUIRED; 点播空间存储桶名称
	Bucket string `json:"Bucket"`

	// REQUIRED; 空间所在地区
	Region string `json:"Region"`

	// REQUIRED; 点播空间名
	SpaceName string `json:"SpaceName"`
}

type GetImageServiceSubscriptionQuery struct {

	// 附加组件ID，获取指定附加组件的开通情况。默认返回ImageX服务开通情况
	AddOnID string `json:"AddOnId,omitempty" query:"AddOnId"`

	// 附加组件英文标识，获取指定附加组件的开通情况。默认返回ImageX服务开通情况。
	AddOnKey string `json:"AddOnKey,omitempty" query:"AddOnKey"`

	// 附加组件类型，取值AI获取服务端智能处理开通情况。默认返回ImageX服务开通情况
	AddOnType string `json:"AddOnType,omitempty" query:"AddOnType"`
}

type GetImageServiceSubscriptionRes struct {

	// REQUIRED
	ResponseMetadata *GetImageServiceSubscriptionResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageServiceSubscriptionResResult `json:"Result"`
}

type GetImageServiceSubscriptionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                               `json:"Version"`
	Error   *GetImageServiceSubscriptionResResponseMetadataError `json:"Error,omitempty"`
}

type GetImageServiceSubscriptionResResponseMetadataError struct {

	// REQUIRED; 错误码
	Code string `json:"Code"`

	// REQUIRED; 错误信息
	Message string `json:"Message"`
}

type GetImageServiceSubscriptionResResult struct {

	// REQUIRED; 账号ID
	AccountID string `json:"AccountId"`

	// REQUIRED; 生效开始时间
	BeginTime string `json:"BeginTime"`

	// REQUIRED; 0 未运行 1 运行中 2 创建中 3 更配中 4 续费中 5 退订中
	BusinessStatus int `json:"BusinessStatus"`

	// REQUIRED; 购买的商品配置
	Configuration string `json:"Configuration"`

	// REQUIRED; 生效结束时间
	ExpiredTime string `json:"ExpiredTime"`

	// REQUIRED; 开通后的实例ID
	InstanceID string `json:"InstanceId"`

	// REQUIRED; 实例类型：1：正式，2：试用
	InstanceType int `json:"InstanceType"`

	// REQUIRED; 购买的商品
	Product string `json:"Product"`

	// REQUIRED; 实例状态：0 创建中；1 运行中；2 创建失败；3 已退订；4 到期关停；5 到期回收
	Status int `json:"Status"`
}

type GetImageSettingRuleHistoryQuery struct {

	// REQUIRED; 应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppId。
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 配置项 ID，您可以通过调用获取配置项列表 [https://www.volcengine.com/docs/508/1324617]的方式获取所需的配置项 ID。
	SettingID string `json:"SettingId" query:"SettingId"`

	// 分页查询时，显示的每页数据的最大条数。取值范围为 [1,100]，默认值为 10。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，用于控制分页查询返回结果的起始位置，以便对数据进行分页展示和浏览。默认值为 0。 :::tip 例如，指定分页条数 Limit = 10，分页偏移量 Offset = 10，表示从查询结果的第 11 条记录开始返回数据，共展示
	// 10 条数据。 :::
	Offset int `json:"Offset,omitempty" query:"Offset"`
}

type GetImageSettingRuleHistoryRes struct {

	// REQUIRED
	ResponseMetadata *GetImageSettingRuleHistoryResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageSettingRuleHistoryResResult `json:"Result,omitempty"`
}

type GetImageSettingRuleHistoryResResponseMetadata struct {

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

// GetImageSettingRuleHistoryResResult - 视请求的接口而定
type GetImageSettingRuleHistoryResResult struct {

	// REQUIRED; 修改记录列表
	Records []*GetImageSettingRuleHistoryResResultRecordsItem `json:"Records"`

	// REQUIRED; 总记录条数
	Total int `json:"Total"`
}

type GetImageSettingRuleHistoryResResultRecordsItem struct {

	// REQUIRED; 修改时间，即修改规则的服务器当地时间。
	ModifiedAt string `json:"ModifiedAt"`

	// REQUIRED; 修改人，即修改规则的当前账号。
	Modifier string `json:"Modifier"`

	// REQUIRED; 修改信息，如：新增规则、删除规则、调整优先级。
	ModifyMsg string `json:"ModifyMsg"`

	// REQUIRED; 新规则内容
	NewInfo *GetImageSettingRuleHistoryResResultRecordsItemNewInfo `json:"NewInfo"`

	// REQUIRED; 旧规则内容
	OldInfo *GetImageSettingRuleHistoryResResultRecordsItemOldInfo `json:"OldInfo"`
}

// GetImageSettingRuleHistoryResResultRecordsItemNewInfo - 新规则内容
type GetImageSettingRuleHistoryResResultRecordsItemNewInfo struct {

	// REQUIRED; 规则配置信息，即匹配条件。
	Content string `json:"Content"`

	// REQUIRED; 类型由对应配置项类型决定，此处是为了方便生成 SDK
	Value interface{} `json:"Value"`
}

// GetImageSettingRuleHistoryResResultRecordsItemOldInfo - 旧规则内容
type GetImageSettingRuleHistoryResResultRecordsItemOldInfo struct {

	// REQUIRED; 规则配置信息，即匹配条件。
	Content string `json:"Content"`

	// REQUIRED; 类型由对应配置项类型决定，此处是为了方便生成 SDK
	Value interface{} `json:"Value"`
}

type GetImageSettingRulesQuery struct {

	// REQUIRED; 应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppId。
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 配置项 ID，您可以通过调用获取配置项列表 [https://www.volcengine.com/docs/508/1324617]的方式获取所需的配置项 ID。
	SettingID string `json:"SettingId" query:"SettingId"`
}

type GetImageSettingRulesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageSettingRulesResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetImageSettingRulesResResult           `json:"Result,omitempty"`
}

type GetImageSettingRulesResResponseMetadata struct {

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

type GetImageSettingRulesResResult struct {

	// REQUIRED; 应用 ID
	AppID string `json:"AppId"`

	// REQUIRED; 应用地域
	AppRegion string `json:"AppRegion"`

	// REQUIRED; 所属组件，取值如下所示
	// * HEIF：表示 HEIF 编解码库
	// * SR：表示客户端加载 SDK
	Category string `json:"Category"`

	// REQUIRED; 规则所属配置项备注信息
	Comment string `json:"Comment"`

	// REQUIRED; 类型由Type决定，此处改为object是为了方便生成 SDK
	DefaultValue interface{} `json:"DefaultValue"`

	// REQUIRED; 规则所属配置项名称
	Name string `json:"Name"`

	// REQUIRED; 父节点 ID，若本身为父节点则该值为空。
	ParentID string `json:"ParentId"`

	// REQUIRED; 规则列表
	Rules []*GetImageSettingRulesResResultRulesItem `json:"Rules"`

	// REQUIRED; 配置项 ID
	SettingID string `json:"SettingId"`

	// REQUIRED; 配置项状态。当前仅支持取值为 0，表示状态正常。
	Status int `json:"Status"`

	// REQUIRED; 配置项类型，取值如下所示：
	// * sample：采样率类型
	// * integer：整数类型
	// * float：浮点数类型
	// * string：字符串类型
	// * strarr：字符串数组类型
	// * bool：布尔值类型
	// * parent：父节点类型
	// * object：对象类型
	Type string `json:"Type"`

	// REQUIRED; 修改时间，修改时的服务器当地时间。
	UpdateAt string `json:"UpdateAt"`

	// REQUIRED; 取值限制范围。仅当Type取值为integer/float/sample /object时有效。Type 取值为 object 时，表示 ValueType 的取值范围。
	ValueRange *GetImageSettingRulesResResultValueRange `json:"ValueRange"`

	// REQUIRED; 仅当Type取值为object时有值，表示 value 类型，key 类型统一为 String。
	ValueType string `json:"ValueType"`
}

type GetImageSettingRulesResResultRulesItem struct {

	// REQUIRED; 规则配置信息
	Content string `json:"Content"`

	// REQUIRED; 规则创建账号
	Creator string `json:"Creator"`

	// REQUIRED; 规则名称
	Name string `json:"Name"`

	// REQUIRED; 规则优先级，值越小优先级越高。
	Priority int `json:"Priority"`

	// REQUIRED; 规则 ID
	RuleID string `json:"RuleId"`

	// REQUIRED; 类型由配置项决定，此处改为object是为了方便生成 SDK
	Value interface{} `json:"Value"`

	// 规则条件
	Cond *GetImageSettingRulesResResultRulesItemCond `json:"Cond,omitempty"`
}

// GetImageSettingRulesResResultRulesItemCond - 规则条件
type GetImageSettingRulesResResultRulesItemCond struct {

	// 条件列表
	Conds []*GetImageSettingRulesResResultRulesPropertiesItemsItem `json:"Conds,omitempty"`

	// 匹配条件，取值如下所示：
	// * AND：表示与
	// * OR：表示或
	Type string `json:"Type,omitempty"`
}

type GetImageSettingRulesResResultRulesPropertiesItemsItem struct {

	// 过滤维度。
	Key string `json:"Key,omitempty"`

	// 操作符。支持取值：==、!=、>、>=、<、<=、in
	Op string `json:"Op,omitempty"`

	// 类型由Op决定，此处改为object是为了方便生成 SDK
	Value interface{} `json:"Value,omitempty"`
}

// GetImageSettingRulesResResultValueRange - 取值限制范围。仅当Type取值为integer/float/sample /object时有效。Type 取值为 object 时，表示 ValueType
// 的取值范围。
type GetImageSettingRulesResResultValueRange struct {

	// 类型由Type决定，此处改为object是为了方便生成 SDK
	Lower interface{} `json:"Lower,omitempty"`

	// 类型由Type决定，此处改为object是为了方便生成 SDK
	Upper interface{} `json:"Upper,omitempty"`
}

type GetImageSettingsQuery struct {

	// REQUIRED; 应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppId。
	AppID string `json:"AppId" query:"AppId"`

	// 所属组件，缺省情况下表示获取基础配置列表。
	// * 取值为HEIF时，表示获取 HEIF 解码库下配置列表；
	// * 取值为SR时，表示获取客户端下配置列表。
	Category string `json:"Category,omitempty" query:"Category"`
}

type GetImageSettingsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageSettingsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageSettingsResResult `json:"Result,omitempty"`
}

type GetImageSettingsResResponseMetadata struct {

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

// GetImageSettingsResResult - 视请求的接口而定
type GetImageSettingsResResult struct {

	// REQUIRED; 配置项列表。
	Settings []*GetImageSettingsResResultSettingsItem `json:"Settings"`
}

type GetImageSettingsResResultSettingsItem struct {

	// REQUIRED; 所属组件，即请求时的组件名称。
	Category string `json:"Category"`

	// REQUIRED; 配置子节点列表，仅当Type取值为parent时有意义。其参数结构与Settings相同。
	ChildSettings []interface{} `json:"ChildSettings"`

	// REQUIRED; 备注信息
	Comment string `json:"Comment"`

	// REQUIRED; 类型由Type决定，此处改为object是为了方便生成 SDK
	DefaultValue interface{} `json:"DefaultValue"`

	// REQUIRED; 配置项名称
	Name string `json:"Name"`

	// REQUIRED; 父节点 ID，若本身为父节点则该值为空。
	ParentID string `json:"ParentId"`

	// REQUIRED; 配置项 ID
	SettingID string `json:"SettingId"`

	// REQUIRED; 配置项状态。当前仅支持取值为 0，表示状态正常。
	Status int `json:"Status"`

	// REQUIRED; 配置项类型，取值如下所示：
	// * sample：采样率类型
	// * integer：整数类型
	// * float：浮点数类型
	// * string：字符串类型
	// * strarr：字符串数组类型
	// * bool：布尔值类型
	// * parent：父节点类型
	// * object：对象类型
	Type string `json:"Type"`

	// REQUIRED; 配置修改时间，修改时的服务器当地时间。
	UpdateAt string `json:"UpdateAt"`

	// REQUIRED; 取值限制范围。仅当Type取值为integer/float/sample/object时有效。Type取值为object时，表示ValueType的取值范围。
	ValueRange *GetImageSettingsResResultSettingsItemValueRange `json:"ValueRange"`

	// REQUIRED; 仅当 Type 取值为 object 时有值，表示 value 类型，key 类型统一为 String。
	ValueType string `json:"ValueType"`
}

// GetImageSettingsResResultSettingsItemValueRange - 取值限制范围。仅当Type取值为integer/float/sample/object时有效。Type取值为object时，表示ValueType的取值范围。
type GetImageSettingsResResultSettingsItemValueRange struct {

	// REQUIRED; 类型由Type决定，此处改为object是为了方便生成 SDK
	Lower interface{} `json:"Lower"`

	// REQUIRED; 类型由Type决定，此处改为object是为了方便生成 SDK
	Upper interface{} `json:"Upper"`
}

type GetImageSmartCropResultBody struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待处理原图的存储 URI 和 URL（公网可访问的 URL）。
	StoreURI string `json:"StoreUri"`

	// 图片裁剪后的高度设置，单位为 px。当图片小于设置的宽高时，将不被裁剪。
	Height int `json:"Height,omitempty"`

	// 降级策略，即当智能裁剪失败时的操作，默认居中裁剪。
	// 支持取值如下：
	// * center：居中裁剪，默认裁剪中间图片；
	// * top：居上裁剪，默认裁剪上方图片；
	// * fglass：高斯模糊，将按设置宽高自动适配图片，结果图多出原图部分以原图背景高斯模糊效果填充。
	Policy string `json:"Policy,omitempty"`

	// 裁剪场景 (normal,cartoon)，默认普通人脸裁剪。支持取值如下：
	// * normal：普通人脸裁剪；
	// * cartoon：动漫人脸裁剪。 :::tip 当前仅支持了智能人脸裁剪能力，其他裁剪能力在持续开放中，敬请期待。 :::
	Scene string `json:"Scene,omitempty"`

	// 当Policy取值为fglass时的高斯模糊参数，取值为大于 0 的整数，值越大越模糊。
	Sigma float64 `json:"Sigma,omitempty"`

	// 图片裁剪后的宽度设置，单位为 px。当图片小于设置的宽高时，将不被裁剪。
	Width int `json:"Width,omitempty"`
}

type GetImageSmartCropResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageSmartCropResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageSmartCropResultResResult `json:"Result,omitempty"`
}

type GetImageSmartCropResultResResponseMetadata struct {

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

// GetImageSmartCropResultResResult - 视请求的接口而定
type GetImageSmartCropResultResResult struct {

	// REQUIRED; 是否降级。支持取值如下：
	// * true：降级；
	// * flase：未降级。
	Demotioned bool `json:"Demotioned"`

	// REQUIRED; 结果图存储 URI。
	ResURI string `json:"ResUri"`
}

type GetImageStorageFilesQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 指定目录分隔符，默认值为空。所有文件名字包含指定的前缀，第一次出现 Delimiter 字符之间的文件作为一组元素（即 CommonPrefixe）。
	Delimiter string `json:"Delimiter,omitempty" query:"Delimiter"`

	// 一次查询列出的文件信息条目数，取值范围为[1,1000]。默认值为 10。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 上一次列举返回的位置标记，作为本次列举的起点信息。默认值为空。
	Marker string `json:"Marker,omitempty" query:"Marker"`

	// 指定需要查询文件的前缀，只有资源名匹配该前缀的文件会被列出。缺省时将返回所有文件信息。
	// 例如，一个存储服务中有三个文件，分别为 Example/imagex.png、Example/mov/a.avis 和 Example/mov/b.avis。若指定 Prefix 取值 Example/且指定 Delimiter 取值 /：则返回
	// Example/imagex.png，并在 CommonPrefix 里返回 Example/mov/。
	Prefix string `json:"Prefix,omitempty" query:"Prefix"`
}

type GetImageStorageFilesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageStorageFilesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageStorageFilesResResult `json:"Result"`
}

type GetImageStorageFilesResResponseMetadata struct {

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

type GetImageStorageFilesResResult struct {

	// REQUIRED; 返回目录名称的数组集合。
	CommonPrefix []string `json:"CommonPrefix"`

	// REQUIRED; 是否还有更多文件，取值如下所示：
	// * true：是，还有文件信息未列出
	// * false：否，已列出所有文件信息
	HasMore bool `json:"HasMore"`

	// REQUIRED; 文件列表
	Items []*GetImageStorageFilesResResultItemsItem `json:"Items"`

	// REQUIRED; HasMore取值true时，即本次查询还有未列举到的文件信息时。Marker作为起始条目位置标记，您需要在下一次列举时传入该值。
	Marker string `json:"Marker"`
}

type GetImageStorageFilesResResultItemsItem struct {

	// 文件大小，单位为 byte。
	FileSize int `json:"FileSize,omitempty"`

	// 文件存储 Key
	Key string `json:"Key,omitempty"`

	// 文件最后修改时间，RFC 时间格式。
	LastModified string `json:"LastModified,omitempty"`

	// 文件的存储类型，取值如下所示：
	// * STANDARD：标准存储
	// * IA：低频存储
	// * ARCHIVE：归档存储
	// * COLD_ARCHIVE：冷归档存储
	StorageClass string `json:"StorageClass,omitempty"`

	// 文件存储 URI
	StoreURI string `json:"StoreUri,omitempty"`
}

type GetImageStyleDetailQuery struct {

	// REQUIRED; 样式 ID。
	StyleID string `json:"StyleId" query:"StyleId"`
}

type GetImageStyleDetailRes struct {

	// REQUIRED
	ResponseMetadata *GetImageStyleDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageStyleDetailResResult `json:"Result"`
}

type GetImageStyleDetailResResponseMetadata struct {

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

type GetImageStyleDetailResResult struct {

	// REQUIRED
	Style *GetImageStyleDetailResResultStyle `json:"Style"`
}

type GetImageStyleDetailResResultStyle struct {

	// REQUIRED
	Background *GetImageStyleDetailResResultStyleBackground `json:"background"`

	// REQUIRED
	Elements []*GetImageStyleDetailResResultStyleElementsItem `json:"elements"`

	// REQUIRED
	Height int `json:"height"`

	// REQUIRED
	ID string `json:"id"`

	// REQUIRED
	Name string `json:"name"`

	// REQUIRED
	Output *GetImageStyleDetailResResultStyleOutput `json:"output"`

	// REQUIRED
	Service string `json:"service"`

	// REQUIRED
	Unit string `json:"unit"`

	// REQUIRED
	Width int `json:"width"`
}

type GetImageStyleDetailResResultStyleBackground struct {

	// REQUIRED
	Fill string `json:"fill"`

	// REQUIRED
	FillSrc string `json:"fillSrc"`

	// REQUIRED
	Viewpoint *GetImageStyleDetailResResultStyleBackgroundViewpoint `json:"viewpoint"`
}

type GetImageStyleDetailResResultStyleBackgroundViewpoint struct {

	// REQUIRED
	Height int `json:"height"`

	// REQUIRED
	Width float64 `json:"width"`

	// REQUIRED
	X float64 `json:"x"`

	// REQUIRED
	Y int `json:"y"`
}

type GetImageStyleDetailResResultStyleElementsItem struct {

	// REQUIRED
	Angle int `json:"angle"`

	// REQUIRED
	Attr *GetImageStyleDetailResResultStyleElementsItemAttr `json:"attr"`

	// REQUIRED
	Content string `json:"content"`

	// REQUIRED
	FlipX bool `json:"flipX"`

	// REQUIRED
	FlipY bool `json:"flipY"`

	// REQUIRED
	Height int `json:"height"`

	// REQUIRED
	ID string `json:"id"`

	// REQUIRED
	Left int `json:"left"`

	// REQUIRED
	Name string `json:"name"`

	// REQUIRED
	Opacity int `json:"opacity"`

	// REQUIRED
	Top int `json:"top"`

	// REQUIRED
	Type string `json:"type"`

	// REQUIRED
	Width int `json:"width"`
}

type GetImageStyleDetailResResultStyleElementsItemAttr struct {

	// REQUIRED
	Type        int                                                                                                                                  `json:"type"`
	Adapt       bool                                                                                                                                 `json:"adapt,omitempty"`
	Bold        bool                                                                                                                                 `json:"bold,omitempty"`
	Border      *Components6Rlnt3SchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesBorder   `json:"border,omitempty"`
	Fillptn     *Components11LotgnSchemasGetimagestyledetailresPropertiesResultPropertiesStylePropertiesElementsItemsPropertiesAttrPropertiesFillptn `json:"fillptn,omitempty"`
	Filter      []string                                                                                                                             `json:"filter,omitempty"`
	Font        string                                                                                                                               `json:"font,omitempty"`
	FontAdapt   bool                                                                                                                                 `json:"fontAdapt,omitempty"`
	FontSize    int                                                                                                                                  `json:"fontSize,omitempty"`
	LineSpace   float64                                                                                                                              `json:"lineSpace,omitempty"`
	Linethrough bool                                                                                                                                 `json:"linethrough,omitempty"`
	Oblique     bool                                                                                                                                 `json:"oblique,omitempty"`
	TextAlign   int                                                                                                                                  `json:"textAlign,omitempty"`
	TextAlignH  int                                                                                                                                  `json:"textAlignH,omitempty"`
	Underline   bool                                                                                                                                 `json:"underline,omitempty"`
	Viewpoint   *GetImageStyleDetailResResultStyleElementsProperties                                                                                 `json:"viewpoint,omitempty"`
	WordSpace   int                                                                                                                                  `json:"wordSpace,omitempty"`
	WritingMode int                                                                                                                                  `json:"writingMode,omitempty"`
}

type GetImageStyleDetailResResultStyleElementsProperties struct {

	// REQUIRED
	Height int `json:"height"`

	// REQUIRED
	Width int `json:"width"`

	// REQUIRED
	X int `json:"x"`

	// REQUIRED
	Y int `json:"y"`
}

type GetImageStyleDetailResResultStyleOutput struct {

	// REQUIRED
	Format string `json:"format"`

	// REQUIRED
	Quality int `json:"quality"`
}

type GetImageStyleResultBody struct {

	// REQUIRED; 要素属性，结构请参考样式定义 [https://www.volcengine.com/docs/508/127402#%E6%A0%B7%E5%BC%8F%E5%AE%9A%E4%B9%89]。
	// 此参数不为空则使用自定义参数内容替换样式定义中对应 elements [https://www.volcengine.com/docs/508/127402#____elements] 的相关属性值。 :::tip
	// * 要素类型不允许更改；
	// * 若elements和params两个参数同时指定了某个要素的内容，则以elements中指定的为准。 :::
	Elements []interface{} `json:"elements"`

	// REQUIRED; 图片渲染所用样式的样式 ID。获取方法请参见如何获取样式 ID [https://www.volcengine.com/docs/508/105396#如何获取创意魔方的样式-id？]。
	StyleID string `json:"StyleId"`

	// 样式背景，结构请参考样式定义 [https://www.volcengine.com/docs/508/127402#%E6%A0%B7%E5%BC%8F%E5%AE%9A%E4%B9%89]。
	// 此参数不为空则使用自定义参数内容替换样式定义中的 background [https://www.volcengine.com/docs/508/127402#background] 属性值。
	Background interface{} `json:"background,omitempty"`

	// 渲染结果图的编码格式，默认值为 webp。取值如下所示：
	// * jpeg
	// * webp
	// * png
	// * heic
	OutputFormat string `json:"OutputFormat,omitempty"`

	// 渲染结果图的编码质量。默认为 75，取值范围为 [1,100]，值越大，结果图的质量越高。
	OutputQuality int `json:"OutputQuality,omitempty"`

	// 样式中的动态要素和要素取值。格式为 "Key":"Value"，例如，"el17fbb3a2134*":"Hello,World",
	// * Key：表示要素 ID，String 类型。获取方式请参见如何获取要素 ID [https://www.volcengine.com/docs/508/105396#如何获取创意魔方样式中的动态要素-id？]；
	// * Value：表示要素的取值，String 类型。需要您根据实际需求自定义，例如，自定义图片地址、文本及二维码等内容。
	Params map[string]string `json:"Params,omitempty"`
}

type GetImageStyleResultQuery struct {

	// REQUIRED; 图片渲染所用样式关联的服务的 ID，用于计量计费和渲染结果的存储。获取方式请参见如何获取调用参数 [https://www.volcengine.com/docs/508/105396]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetImageStyleResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageStyleResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageStyleResultResResult `json:"Result,omitempty"`
}

type GetImageStyleResultResResponseMetadata struct {

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

// GetImageStyleResultResResult - 视请求的接口而定
type GetImageStyleResultResResult struct {

	// REQUIRED; 渲染详情
	RenderDetail []*GetImageStyleResultResResultRenderDetailItem `json:"RenderDetail"`

	// REQUIRED; 渲染结果图的 URI
	ResURI string `json:"ResUri"`
}

type GetImageStyleResultResResultRenderDetailItem struct {

	// REQUIRED; 渲染失败的要素 ID
	Element string `json:"Element"`

	// REQUIRED; 渲染失败的原因
	ErrMsg string `json:"ErrMsg"`
}

type GetImageStylesQuery struct {

	// REQUIRED; 样式类型。取值 user 表示用户样式。
	Type string `json:"Type" query:"Type"`

	// 分页返回条数，取值范围为[0,100]，默认 10 条。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移，默认 0，取值为 1 时，表示跳过一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 需要返回的样式列表标识。
	// * 返回的样式名称、样式 ID 包含了该值的样式列表。
	// * 返回的样式宽度或样式高度为该值的样式列表。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`
}

type GetImageStylesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageStylesResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *GetImageStylesResResult `json:"Result,omitempty"`
}

type GetImageStylesResResponseMetadata struct {

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

// GetImageStylesResResult - title
type GetImageStylesResResult struct {

	// REQUIRED; 当前分页样式数据
	Styles []*GetImageStylesResResultStylesItem `json:"Styles"`

	// REQUIRED; 总样式个数。
	Total int `json:"Total"`
}

type GetImageStylesResResultStylesItem struct {

	// REQUIRED; 样式创建时间
	CreateAt string `json:"CreateAt"`

	// REQUIRED; 样式高
	Height int `json:"Height"`

	// REQUIRED; 样式ID
	ID string `json:"Id"`

	// REQUIRED; 样式名称
	Name string `json:"Name"`

	// REQUIRED; 样式渲染结果图的TOS URI
	ResURI string `json:"ResUri"`

	// REQUIRED; 样式绑定的服务ID
	Service string `json:"Service"`

	// REQUIRED; 尺寸单位。当前仅支持取值px表示像素。
	Unit GetImageStylesResResultStylesItemUnit `json:"Unit"`

	// REQUIRED; 样式更新时间
	UpdateAt string `json:"UpdateAt"`

	// REQUIRED; 样式宽
	Width int `json:"Width"`
}

type GetImageSuperResolutionResultBody struct {

	// REQUIRED; 用于存储结果图和计量计费的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待处理原图的存储 URI 或访问 URL（可公网访问）。您可在控制台资源管理获取图片的存储 URI [https://www.volcengine.com/docs/508/1205057]以及访问 URL [https://www.volcengine.com/docs/508/1205054]。
	StoreURI string `json:"StoreUri"`

	// 超分倍率，默认值为2，支持取值为：2、3、4、5、6、7、8。
	Multiple float64 `json:"Multiple,omitempty"`
}

type GetImageSuperResolutionResultRes struct {

	// REQUIRED
	ResponseMetadata *GetImageSuperResolutionResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageSuperResolutionResultResResult `json:"Result,omitempty"`
}

type GetImageSuperResolutionResultResResponseMetadata struct {

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

// GetImageSuperResolutionResultResResult - 视请求的接口而定
type GetImageSuperResolutionResultResResult struct {

	// REQUIRED; 结果图 URI。您可使用结果图 URI（即ResUri）拼接完整访问 URL [https://www.volcengine.com/docs/508/105405#预览返回的结果图]后，在浏览器查看图像超分辨率效果。
	ResURI string `json:"ResUri"`
}

type GetImageTemplateQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 模板名称。
	// * 您可以通过调用获取服务下所有模板 [https://www.volcengine.com/docs/508/9386]获取所需的模板名称。
	TemplateName string `json:"TemplateName" query:"TemplateName"`
}

type GetImageTemplateRes struct {

	// REQUIRED
	ResponseMetadata *GetImageTemplateResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetImageTemplateResResult           `json:"Result,omitempty"`
}

type GetImageTemplateResResponseMetadata struct {

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

type GetImageTemplateResResult struct {

	// REQUIRED; 模板计划使用的降级格式，仅对heic静图有效
	DemotionFormat string `json:"DemotionFormat"`

	// REQUIRED; 该模板计划使用的输出格式
	OutputFormat string `json:"OutputFormat"`

	// REQUIRED; 绝对质量/相对质量
	QualityMode string `json:"QualityMode"`

	// REQUIRED; 服务id
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 是否同步处理，仅对heic图有效
	Sync bool `json:"Sync"`

	// REQUIRED; 模板名称，必须使用该服务的图片模板固定前缀，具体参考GetImageService接口的返回参数TemplatePrefix。模板名称能包含的字符正则集合为[a-zA-Z0-9_-]
	TemplateName string `json:"TemplateName"`

	// 图像格式自适应配置
	AdaptiveFmt *GetImageTemplateResResultAdaptiveFmt `json:"AdaptiveFmt,omitempty"`

	// 视频转动图配置
	Animation *GetImageTemplateResResultAnimation `json:"Animation,omitempty"`

	// 是否直接更新模板。如果为true，已有的线上模板会同步更新，直接生效；如果为false，会新增一个模板，已有模板不受影响。
	DoUpdate bool `json:"DoUpdate,omitempty"`

	// 对图片的编辑操作
	Filters []*GetImageTemplateResResultFiltersItem `json:"Filters,omitempty"`

	// 对图片编码使用的质量参数，默认为0
	OuputQuality float64 `json:"OuputQuality,omitempty"`

	// 用于图片服务输出时的图片编码
	OutputExtra *GetImageTemplateResResultOutputExtra `json:"OutputExtra,omitempty"`

	// 图片模板使用的参数列表，URL中下发参数的顺序需要跟列表中的保持一致
	Parameters []string `json:"Parameters,omitempty"`

	// URL的失效期，为Unix时间戳，一般配置为通过模板参数下发
	ReqDeadline string `json:"ReqDeadline,omitempty"`

	// 视频截帧配置
	Snapshot *GetImageTemplateResResultSnapshot `json:"Snapshot,omitempty"`

	// 是否开启鉴权，一般当通过模板参数下发敏感信息时，比如文字水印内容、URL失效期，需要对图片URL鉴权保护，防止内容被篡改
	WithSig bool `json:"WithSig,omitempty"`
}

// GetImageTemplateResResultAdaptiveFmt - 图像格式自适应配置
type GetImageTemplateResResultAdaptiveFmt struct {

	// REQUIRED; 动图自适应，可选"webp"/"heic"/"avif"/"dynamic"
	Animated string `json:"Animated"`

	// REQUIRED; 静图自适应，可选"webp"/"heic"/"avif"/"dynamic"
	Static string `json:"Static"`
}

// GetImageTemplateResResultAnimation - 视频转动图配置
type GetImageTemplateResResultAnimation struct {

	// REQUIRED; 降级类型： image：抽取一帧降级返回 video：直接返回原视频降级
	DemotionType string `json:"DemotionType"`

	// REQUIRED; 动图时长(ms)
	Duration int `json:"Duration"`

	// REQUIRED; 帧率，1秒X帧
	FramePerSecond int `json:"FramePerSecond"`

	// REQUIRED; X秒1帧
	SecondPerFrame int `json:"SecondPerFrame"`

	// REQUIRED; 抽帧策略： fps：frame per second，1秒X帧 spf：second per frame，X秒1帧 key：抽取关键帧
	SelectFrameMode string `json:"SelectFrameMode"`

	// REQUIRED; 动图起始时间戳(ms)
	StartTime int `json:"StartTime"`

	// REQUIRED; 同步等待时长(s)，超时未完成则根据DemotionType降级
	WaitTime int `json:"WaitTime"`
}

type GetImageTemplateResResultFiltersItem struct {

	// REQUIRED; 编辑操作的名称
	Name string `json:"Name"`

	// REQUIRED; 编辑操作的参数
	Param interface{} `json:"Param"`
}

// GetImageTemplateResResultOutputExtra - 用于图片服务输出时的图片编码
type GetImageTemplateResResultOutputExtra struct {

	// REQUIRED; 是否带透明通道编码，"false"/"true"
	HeicAlphaReserve string `json:"heic.alpha.reserve"`

	// REQUIRED; heic位深，"8"/"10"
	HeicEncodeDepth string `json:"heic.encode.depth"`

	// REQUIRED; heic格式是否开启ROI编码"true"/"false"
	HeicRoi string `json:"heic.roi"`

	// REQUIRED; heic缩略图比例
	HeicThumbRatio string `json:"heic.thumb.ratio"`

	// REQUIRED; 是否采用jpeg渐进编码格式,取值为"true" / "false"
	JPEGProgressive string `json:"jpeg.progressive"`

	// REQUIRED; 是否压缩颜色空间,取值为"true" / "false"
	PNGUseQuant string `json:"png.use_quant"`
}

// GetImageTemplateResResultSnapshot - 视频截帧配置
type GetImageTemplateResResultSnapshot struct {

	// REQUIRED; 截图的时间戳(ms)
	TimeOffsetMs string `json:"TimeOffsetMs"`

	// REQUIRED; 截图类型，可选"default"/"offset"，智能截图和指定时间戳截图
	Type string `json:"Type"`
}

type GetImageTranscodeDetailsQuery struct {

	// REQUIRED; 任务提交的截止 Unix 时间戳StartTime与EndTime时间间隔最大不超过 7 天。
	EndTime int `json:"EndTime" query:"EndTime"`

	// REQUIRED; 分页条数，取值范围为(0, 100]。
	Limit int `json:"Limit" query:"Limit"`

	// REQUIRED; 队列 ID，您可通过调用GetImageTranscodeQueues [https://www.volcengine.com/docs/508/1107341]获取该账号下全部任务队列 ID。
	QueueID string `json:"QueueId" query:"QueueId"`

	// REQUIRED; 任务提交的起始 Unix 时间戳StartTime与EndTime时间间隔最大不超过 7 天。
	StartTime int `json:"StartTime" query:"StartTime"`

	// 分页偏移量，默认为 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 队列所在地区。默认当前地区为 cn。
	Region string `json:"Region,omitempty" query:"Region"`

	// 返回图片 url 或 uri 中包含该值的任务。默认为空，不传则返回所有任务。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`

	// 执行状态，填入多个时使用英文逗号分隔。取值如下所示：
	// * Pending：排队中
	// * Running：执行中
	// * Success：执行成功
	// * Fail：执行失败
	Status string `json:"Status,omitempty" query:"Status"`

	// 任务 ID，缺省情况下查询指定队列下所有任务详情。您可通过调用 GetImageTranscodeTasks [https://www.volcengine.com/docs/508/1356555]获取指定队列的全部任务 ID。
	TaskID string `json:"TaskId,omitempty" query:"TaskId"`
}

type GetImageTranscodeDetailsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageTranscodeDetailsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageTranscodeDetailsResResult `json:"Result"`
}

type GetImageTranscodeDetailsResResponseMetadata struct {

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

type GetImageTranscodeDetailsResResult struct {

	// REQUIRED; 执行任务详情
	ExecInfo []*GetImageTranscodeDetailsResResultExecInfoItem `json:"ExecInfo"`

	// REQUIRED; 总数
	Total int `json:"Total"`
}

type GetImageTranscodeDetailsResResultExecInfoItem struct {

	// 结束时间
	EndAt string `json:"EndAt,omitempty"`

	// 条目 ID
	EntryID string `json:"EntryId,omitempty"`

	// 执行输入
	ExecInput *GetImageTranscodeDetailsResResultExecInfoItemExecInput `json:"ExecInput,omitempty"`

	// 执行输出
	ExecOutput *GetImageTranscodeDetailsResResultExecInfoItemExecOutput `json:"ExecOutput,omitempty"`

	// 开始时间
	StartAt string `json:"StartAt,omitempty"`

	// 执行状态。取值如下所示：
	// * Pending：排队中
	// * Running：执行中
	// * Success：执行成功
	// * Fail：执行失败
	Status string `json:"Status,omitempty"`

	// 提交时间
	SubmitAt string `json:"SubmitAt,omitempty"`
}

// GetImageTranscodeDetailsResResultExecInfoItemExecInput - 执行输入
type GetImageTranscodeDetailsResResultExecInfoItemExecInput struct {

	// REQUIRED; 原图图片的 URL 或存储 URI。
	Image string `json:"Image"`

	// REQUIRED; 转码模板
	Template string `json:"Template"`
}

// GetImageTranscodeDetailsResResultExecInfoItemExecOutput - 执行输出
type GetImageTranscodeDetailsResResultExecInfoItemExecOutput struct {

	// REQUIRED; 转码失败错误码 [https://www.volcengine.com/docs/508/1104726#%E9%94%99%E8%AF%AF%E7%A0%81]。仅当Status值为Fail时，ErrCode有值。
	ErrCode string `json:"ErrCode"`

	// REQUIRED; 转码失败错误信息。仅当Status值为Fail时，ErrMsg有值。
	ErrMsg string `json:"ErrMsg"`

	// REQUIRED; 请提供具体的参数名字和类型。
	Evals []*GetImageTranscodeDetailsResResultExecInfoPropertiesItemsItem `json:"Evals"`

	// REQUIRED; 转码结果图格式
	Format string `json:"Format"`

	// REQUIRED; 转码结果图的存储 URI。仅当Status值为Success时，Output有值。
	Output string `json:"Output"`

	// REQUIRED; 尺寸。取值范围为 [ ]，单位为，默认值为``。
	Size int `json:"Size"`
}

// GetImageTranscodeDetailsResResultExecInfoItemExecOutputEvalsItemValue - 参数值。
type GetImageTranscodeDetailsResResultExecInfoItemExecOutputEvalsItemValue struct {

	// REQUIRED; 美学配置。
	Aesthetic float64 `json:"Aesthetic"`

	// REQUIRED; 噪声类型。
	Noise float64 `json:"Noise"`

	// REQUIRED; 峰值信噪比。
	PSNR float64 `json:"PSNR"`

	// REQUIRED; 结构相似性指数。
	SSIM float64 `json:"SSIM"`

	// REQUIRED; 视频多重评估函数。
	VMAF float64 `json:"VMAF"`

	// REQUIRED; 视频质量评分。
	VQScore float64 `json:"VQScore"`
}

type GetImageTranscodeDetailsResResultExecInfoPropertiesItemsItem struct {

	// REQUIRED; 参数格式。
	Format string `json:"Format"`

	// REQUIRED; 请提供具体的参数名字（Index）和类型（string），以便我为您生成参数描述。
	Index int `json:"Index"`

	// REQUIRED; 参数名称。
	Name string `json:"Name"`

	// REQUIRED; 当前阶段。
	Phase string `json:"Phase"`

	// REQUIRED; 实例规格。
	Size int `json:"Size"`

	// REQUIRED; 参数值。
	Value *GetImageTranscodeDetailsResResultExecInfoItemExecOutputEvalsItemValue `json:"Value"`
}

type GetImageTranscodeQueuesQuery struct {

	// REQUIRED; 分页条数，取值范围为(0,100]。
	Limit int `json:"Limit" query:"Limit"`

	// 分页偏移量，默认为 0。取值为 1 时，表示跳过第一条数据，从第二条数据取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 队列所在地区。默认当前地区。ToB取值枚举：cn、va、sg。
	Region string `json:"Region,omitempty" query:"Region"`

	// 返回队列名称或队列描述中包含该值的队列。默认为空，不传则返回所有队列。
	SearchPtn string `json:"SearchPtn,omitempty" query:"SearchPtn"`
}

type GetImageTranscodeQueuesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageTranscodeQueuesResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result *GetImageTranscodeQueuesResResult `json:"Result,omitempty"`
}

type GetImageTranscodeQueuesResResponseMetadata struct {

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

// GetImageTranscodeQueuesResResult - title
type GetImageTranscodeQueuesResResult struct {

	// REQUIRED; 当前分页队列详细信息
	Queues []*GetImageTranscodeQueuesResResultQueuesItem `json:"Queues"`

	// REQUIRED; 符合条件的队列总数
	Total int `json:"Total"`
}

type GetImageTranscodeQueuesResResultQueuesItem struct {

	// REQUIRED; 队列回调设置
	CallbackConf *GetImageTranscodeQueuesResResultQueuesItemCallbackConf `json:"CallbackConf"`

	// REQUIRED; 队列创建时间
	CreateAt string `json:"CreateAt"`

	// REQUIRED; 队列描述
	Desc string `json:"Desc"`

	// REQUIRED; 是否启用回调，取值如下所示：
	// * true：启用
	// * false：不启用
	EnableCallback bool `json:"EnableCallback"`

	// REQUIRED; 队列 ID
	ID string `json:"Id"`

	// REQUIRED; 队列名称
	Name string `json:"Name"`

	// REQUIRED; 队列状态。取值如下所示：
	// * Pending：排队中
	// * Running：执行中
	Status GetImageTranscodeQueuesResResultQueuesItemStatus `json:"Status"`

	// REQUIRED; 队列类型。取值如下所示：
	// * default：表示账号默认队列，每个账号一个
	// * user：表示用户创建队列，个数将有配额限制
	Type GetImageTranscodeQueuesResResultQueuesItemType `json:"Type"`
}

// GetImageTranscodeQueuesResResultQueuesItemCallbackConf - 队列回调设置
type GetImageTranscodeQueuesResResultQueuesItemCallbackConf struct {

	// REQUIRED; 回调数据格式。取值如下所示：
	// * XML
	// * JSON
	DataFormat string `json:"DataFormat"`

	// REQUIRED; Method=MQ时取值rmq:{topic}/{cluster}
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 枚举值还有 MQ
	Method string `json:"Method"`

	// 业务自定义回调参数，将在回调消息的callback_args中透传出去。具体回调参数请参考回调内容 [https://www.volcengine.com/docs/508/1104726#%E5%9B%9E%E8%B0%83%E5%86%85%E5%AE%B9]。
	Args string `json:"Args,omitempty"`
}

type GetImageUpdateFilesQuery struct {

	// REQUIRED; 需要查询的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 分页查询时，显示的每页数据的最大条数。最大值为 100。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移量，用于控制分页查询返回结果的起始位置，以便对数据进行分页展示和浏览。默认值为 0。 :::tip 例如，指定分页条数 Limit = 10，分页偏移量 Offset = 10，表示从查询结果的第 11 条记录开始返回数据，共展示
	// 10 条数据。 :::
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 获取类型，取值如下所示：
	// * 0：获取刷新 URL 列表
	// * 1：获取禁用 URL 列表
	Type int `json:"Type,omitempty" query:"Type"`

	// URL 格式，若指定 URL 格式则仅返回 URL 中包含该字符串的 URL 列表。默认为空，缺省情况下返回所有 URL 列表。
	URLPattern string `json:"UrlPattern,omitempty" query:"UrlPattern"`
}

type GetImageUpdateFilesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageUpdateFilesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageUpdateFilesResResult `json:"Result,omitempty"`
}

type GetImageUpdateFilesResResponseMetadata struct {

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

// GetImageUpdateFilesResResult - 视请求的接口而定
type GetImageUpdateFilesResResult struct {

	// REQUIRED; 服务 ID。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 当前存储状态，取值为：正常、未审核、禁用。
	Status string `json:"Status"`

	// REQUIRED; 符合条件的 URL 总数
	Total int `json:"Total"`

	// REQUIRED; 符合条件的 URL 列表。
	UpdateUrls []*GetImageUpdateFilesResResultUpdateUrlsItem `json:"UpdateUrls"`
}

type GetImageUpdateFilesResResultUpdateUrlsItem struct {

	// REQUIRED; 图片 URL。
	ImageURL string `json:"ImageUrl"`

	// REQUIRED; URL 更新时间，即图像更新时的服务器当地时间。
	UpdateAt string `json:"UpdateAt"`
}

type GetImageUploadFileQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 文件 Uri。
	// * 您可以在 veImageX 控制台资源管理 [https://console.volcengine.com/imagex/resource_manage/]页面，在已上传文件的名称列获取资源 Uri。
	// * 您也可以通过 OpenAPI 的方式获取Uri，具体请参考 GetImageUploadFiles [https://www.volcengine.com/docs/508/9392]。
	StoreURI string `json:"StoreUri" query:"StoreUri"`
}

type GetImageUploadFileRes struct {

	// REQUIRED
	ResponseMetadata *GetImageUploadFileResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetImageUploadFileResResult `json:"Result,omitempty"`
}

type GetImageUploadFileResResponseMetadata struct {

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

// GetImageUploadFileResResult - 视请求的接口而定
type GetImageUploadFileResResult struct {

	// 是否被禁用
	Disabled bool `json:"Disabled,omitempty"`

	// 文件字节数。
	FileSize int `json:"FileSize,omitempty"`

	// 文件修改时间，即文件修改时的服务器当地时间。
	LastModified string `json:"LastModified,omitempty"`

	// 文件恢复副本的到期时间。仅当文件执行过恢复操作时有值
	RestoreExpiryDate string `json:"RestoreExpiryDate,omitempty"`

	// 恢复请求日期。
	RestoreRequestDate string `json:"RestoreRequestDate,omitempty"`

	// 恢复取回方式
	RestoreTier string `json:"RestoreTier,omitempty"`

	// 文件是否处于恢复中状态。
	Restoring bool `json:"Restoring,omitempty"`

	// 服务 ID。
	ServiceID string `json:"ServiceId,omitempty"`

	// 底层存储类型。 STANDARD：标准存储 IA：低频存储 ARCHIVE：归档存储 COLD_ARCHIVE：冷归档存储
	StorageClass string `json:"StorageClass,omitempty"`

	// 底层存储的 content-type 值。
	StorageContentType string `json:"StorageContentType,omitempty"`

	// 文件 Uri。
	StoreURI string `json:"StoreUri,omitempty"`
}

type GetImageUploadFilesQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 获取文件个数，最大值为 100。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页标志。
	Marker int `json:"Marker,omitempty" query:"Marker"`
}

type GetImageUploadFilesRes struct {

	// REQUIRED
	ResponseMetadata *GetImageUploadFilesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageUploadFilesResResult `json:"Result"`
}

type GetImageUploadFilesResResponseMetadata struct {

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

type GetImageUploadFilesResResult struct {

	// REQUIRED; 上传文件信息。
	FileObjects []*GetImageUploadFilesResResultFileObjectsItem `json:"FileObjects"`

	// REQUIRED; 是否还有下页数据。
	HasMore bool `json:"hasMore"`

	// REQUIRED; 服务 ID。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 当前存储状态，取值为：正常、未审核、禁用。
	Status string `json:"Status"`
}

type GetImageUploadFilesResResultFileObjectsItem struct {

	// 文件字节数。
	FileSize int `json:"FileSize,omitempty"`

	// 文件修改时间，文件修改时的服务器当地时间。
	LastModified string `json:"LastModified,omitempty"`

	// 分页标志。
	Marker int `json:"Marker,omitempty"`

	// 文件 Uri。
	StoreURI string `json:"StoreUri,omitempty"`
}

type GetImageXQueryAppsQuery struct {

	// 项目名。仅子用户使用。
	Project string `json:"Project,omitempty" query:"Project"`

	// 数据来源，账号下近 60 天内有数据上报的应用 ID，缺省情况下返回账号对应的全部应用 ID。取值如下所示：
	// * upload：上传 1.0 数据。
	// * cdn：下行网络数据。
	// * client：客户端数据。
	// * sensible：感知数据。
	// * uploadv2：上传 2.0 数据。
	// * exceed：大图监控数据。
	Source string `json:"Source,omitempty" query:"Source"`
}

type GetImageXQueryAppsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageXQueryAppsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageXQueryAppsResResult `json:"Result"`
}

type GetImageXQueryAppsResResponseMetadata struct {

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

type GetImageXQueryAppsResResult struct {

	// REQUIRED; 火山引擎账号内的所有 App 信息。
	Apps []*GetImageXQueryAppsResResultAppsItem `json:"Apps"`
}

type GetImageXQueryAppsResResultAppsItem struct {

	// REQUIRED; 应用 ID。
	AppID string `json:"AppId"`

	// REQUIRED; 应用名称。
	AppName string `json:"AppName"`

	// REQUIRED; App 所属地区。 cn：国内。 va：美东。 sg：新加坡。
	AppRegion string `json:"AppRegion"`

	// REQUIRED; App 绑定的 iOS 包名。
	BundleID string `json:"BundleId"`

	// REQUIRED; App 绑定的安卓包名。
	PackageName string `json:"PackageName"`
}

type GetImageXQueryDimsQuery struct {

	// REQUIRED; 数据来源，取值如下所示：
	// * upload：上传 1.0 数据。
	// * cdn：下行网络数据。
	// * client：客户端数据。
	// * sensible：感知数据。
	// * uploadv2：上传 2.0 数据。
	// * exceed：大图监控数据，包含大图样本量和大图明细。
	// * exceed_all：大图分布数据。
	Source string `json:"Source" query:"Source"`

	// 应用 ID。传入多个用英文逗号分割。默认为空，匹配账号下所有的 AppID。 :::tip 您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/1213042]的方式获取所需的应用 ID。
	// :::
	Appid string `json:"Appid,omitempty" query:"Appid"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web，仅当Source为upload或uploadv2时可传。
	// * Imp：小程序，仅当Source为upload或uploadv2时可传。
	OS string `json:"OS,omitempty" query:"OS"`
}

type GetImageXQueryDimsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageXQueryDimsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageXQueryDimsResResult `json:"Result"`
}

type GetImageXQueryDimsResResponseMetadata struct {

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

type GetImageXQueryDimsResResult struct {

	// REQUIRED; 上报数据中出现的维度信息。
	Dim []string `json:"Dim"`
}

type GetImageXQueryRegionsQuery struct {

	// REQUIRED; 数据来源，取值如下所示：
	// * upload：上传 1.0 数据。
	// * cdn：下行网络数据。
	// * client：客户端数据。
	// * uploadv2：上传 2.0 数据。
	Source string `json:"Source" query:"Source"`

	// 应用 ID。传入多个用英文逗号分割。默认为空，匹配账号下所有的 AppID。 :::tip 您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/1213042]的方式获取所需的应用 ID。
	// :::
	Appid string `json:"Appid,omitempty" query:"Appid"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web，仅当Source为upload或uploadv2时可传。
	// * Imp：小程序，仅当Source为upload或uploadv2时可传。
	OS string `json:"OS,omitempty" query:"OS"`
}

type GetImageXQueryRegionsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageXQueryRegionsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageXQueryRegionsResResult `json:"Result"`
}

type GetImageXQueryRegionsResResponseMetadata struct {

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

type GetImageXQueryRegionsResResult struct {

	// REQUIRED; 60 天内上报数据中出现的海外国家，按出现次数降序排列。
	Country []string `json:"Country"`

	// REQUIRED; 60 天内上报数据中出现的国内省份，按出现次数降序排列。
	Province []string `json:"Province"`
}

type GetImageXQueryValsQuery struct {

	// REQUIRED; 自定义维度名称。 :::tip 您可以通过调用获取自定义维度列表 [https://www.volcengine.com/docs/508/1213048]获取所需的维度名称。 :::
	Dim string `json:"Dim" query:"Dim"`

	// REQUIRED; 数据来源。
	// * upload：上传 1.0 数据。
	// * cdn：下行网络数据。
	// * client：客户端数据。
	// * sensible：感知数据。
	// * uploadv2：上传 2.0 数据。
	// * exceed：大图监控数据，包含大图样本量和大图明细。
	// * exceed_all：大图分布数据。
	Source string `json:"Source" query:"Source"`

	// 应用 ID。传入多个用英文逗号分割。默认为空，匹配中账号下所有的 AppID。 :::tip 您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/1213042]的方式获取所需的 AppID。
	// :::
	Appid string `json:"Appid,omitempty" query:"Appid"`

	// 需要过滤的关键词（包含），不传则不过滤关键词。
	Keyword string `json:"Keyword,omitempty" query:"Keyword"`

	// 需要匹配的系统类型。取值如下所示：
	// * 不传或传空字符串：Android+iOS。
	// * iOS：iOS。
	// * Android：Android。
	// * WEB：web+小程序。
	// * Web：web，仅当Source为upload或uploadv2时可传。
	// * Imp：小程序，仅当Source为upload或uploadv2时可传。
	OS string `json:"OS,omitempty" query:"OS"`
}

type GetImageXQueryValsRes struct {

	// REQUIRED
	ResponseMetadata *GetImageXQueryValsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result *GetImageXQueryValsResResult `json:"Result"`
}

type GetImageXQueryValsResResponseMetadata struct {

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

type GetImageXQueryValsResResult struct {

	// REQUIRED; 90 天内上报数据中出现的维度值列表，按上报次数降序排列，仅返回前1000条数据。
	DimVal []string `json:"DimVal"`
}

type GetLicensePlateDetectionBody struct {

	// REQUIRED; 原图的存储 URI。
	ImageURI string `json:"ImageUri"`
}

type GetLicensePlateDetectionQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetLicensePlateDetectionRes struct {

	// REQUIRED
	ResponseMetadata *GetLicensePlateDetectionResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetLicensePlateDetectionResResult `json:"Result,omitempty"`
}

type GetLicensePlateDetectionResResponseMetadata struct {

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

// GetLicensePlateDetectionResResult - 视请求的接口而定
type GetLicensePlateDetectionResResult struct {

	// REQUIRED; 检测到的车牌位置坐标
	Locations []*GetLicensePlateDetectionResResultLocationsItem `json:"Locations"`
}

type GetLicensePlateDetectionResResultLocationsItem struct {

	// REQUIRED; 车牌区域左上角 X 轴坐标。
	X1 int `json:"X1"`

	// REQUIRED; 车牌区域右下角 X 轴坐标。
	X2 int `json:"X2"`

	// REQUIRED; 车牌区域左上角 Y 轴坐标。
	Y1 int `json:"Y1"`

	// REQUIRED; 车牌区域右下角 Y 轴坐标。
	Y2 int `json:"Y2"`
}

type GetPrivateImageTypeBody struct {

	// REQUIRED; 原图的存储 URI。
	ImageURI string `json:"ImageUri"`

	// 检测内容，默认为all，取值如下所示：
	// * face：图片内人脸检测
	// * cloth：图片内衣物检测
	// * all：图片内人脸和衣物均检测
	Method string `json:"Method,omitempty"`

	// 衣物置信度，取值范围为[0, 1], 默认值为 0.8，表示高于 0.8 即为有效检测。
	ThresCloth float64 `json:"ThresCloth,omitempty"`

	// 人脸置信度，取值范围为[0, 1], 默认值为 0.52，表示高于 0.52 即为有效检测。
	ThresFace float64 `json:"ThresFace,omitempty"`
}

type GetPrivateImageTypeQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetPrivateImageTypeRes struct {

	// REQUIRED
	ResponseMetadata *GetPrivateImageTypeResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetPrivateImageTypeResResult `json:"Result,omitempty"`
}

type GetPrivateImageTypeResResponseMetadata struct {

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

// GetPrivateImageTypeResResult - 视请求的接口而定
type GetPrivateImageTypeResResult struct {

	// REQUIRED; 请求参数Method中包含cloth则返回该参数，取值如下所示：
	// * 1： 包含衣物
	// * 0： 不包含衣物
	Cloth int `json:"Cloth"`

	// REQUIRED; 请求参数Method中包含face则返回该参数，取值如下所示：
	// * 1： 包含人脸
	// * 0： 不包含人脸
	Face int `json:"Face"`
}

type GetProductAIGCResultBody struct {

	// REQUIRED; 默认为True
	BackgroundOnly bool `json:"BackgroundOnly"`

	// REQUIRED; 输出图片的长度，512~1024
	OutputSize int `json:"OutputSize"`

	// REQUIRED; 是否返回场景图默认为True
	ReturnBackgroundImage bool `json:"ReturnBackgroundImage"`

	// REQUIRED; 商品主体图的访问 URL（公网可访问）。建议为包含完整商品主体的白底图或透底图，尽量避免复杂背景的影响，以确保最终生成效果的质量。
	URL string `json:"Url"`

	// 每次生成的图片数量，[0,4], 默认4
	BatchSize int `json:"BatchSize,omitempty"`

	// 设置商品放置的安全区中心坐标和宽高：设为默认值-1时，商品自动居中，安全区为全图；否则用户需同时指定四个参数的值，取值范围大于等于-1
	CX int `json:"CX,omitempty"`

	// 设置商品放置的安全区中心坐标和宽高：设为默认值-1时，商品自动居中，安全区为全图；否则用户需同时指定四个参数的值，取值范围大于等于-1
	CY int `json:"CY,omitempty"`

	// 额外参数
	Extra string `json:"Extra,omitempty"`

	// Lora 配置。
	LoraConfig string `json:"LoraConfig,omitempty"`

	// 使用分割处理图片，False则不做分割，从输入图像读取alpha通道作为mask
	NeedSeg bool `json:"NeedSeg,omitempty"`

	// 输入到AIGC模型的正向提示词,长度限制为300个字符内
	NegativePrompt string `json:"NegativePrompt,omitempty"`

	// 场景图输出高度
	OutputHeight int `json:"OutputHeight,omitempty"`

	// 场景图输出宽度
	OutputWidth int `json:"OutputWidth,omitempty"`

	// 输入到AIGC模型的负向提示词,长度限制为300个字符内
	PositivePrompt string `json:"PositivePrompt,omitempty"`

	// 商品mask长宽与output_size比值的上限，（0，1）
	ProductRatio float64 `json:"ProductRatio,omitempty"`

	// 商品比例
	ProductRatios [][]float64 `json:"ProductRatios,omitempty"`

	// 根据业务需要，取值为True时，只返回最高分的生成图及其得分，否则返回所有生成图及其得分
	ReturnTop1 bool `json:"ReturnTop1,omitempty"`

	// 设置商品放置的安全区中心坐标和宽高：设为默认值-1时，商品自动居中，安全区为全图；否则用户需同时指定四个参数的值，取值范围大于等于-1
	SafeH int `json:"SafeH,omitempty"`

	// 设置商品放置的安全区中心坐标和宽高：设为默认值-1时，商品自动居中，安全区为全图；否则用户需同时指定四个参数的值，取值范围大于等于-1
	SafeW int `json:"SafeW,omitempty"`

	// aigc场景
	Scene string `json:"Scene,omitempty"`

	// 卖点图配置信息。
	SellingPointConfig *GetProductAIGCResultBodySellingPointConfig `json:"SellingPointConfig,omitempty"`

	// 打分策略规则。
	StrategyRules map[string]interface{} `json:"StrategyRules,omitempty"`

	// 取值为True时，使用blip模型提取对商品的描述，和positive_prompt共同作为输入到AIGC模型的正向提示词
	UseCaption bool `json:"UseCaption,omitempty"`

	// 使用默认的背景图
	UseDefaultBg bool `json:"UseDefaultBg,omitempty"`
}

// GetProductAIGCResultBodySellingPointConfig - 卖点图配置信息。
type GetProductAIGCResultBodySellingPointConfig struct {

	// 用户指定的场景图
	BackgroundURL string `json:"BackgroundUrl,omitempty"`

	// 场景图对应图层ID
	ProdUniqueID string `json:"ProdUniqueId,omitempty"`

	// 商品卖点描述,不超过430字符
	ProductDescription string `json:"ProductDescription,omitempty"`

	// 商详图url列表,最大支持三张
	ProductDetailImages       []string                                                                   `json:"ProductDetailImages,omitempty"`
	SellingPointExtractConfig []*GetProductAIGCResultBodySellingPointConfigSellingPointExtractConfigItem `json:"SellingPointExtractConfig,omitempty"`

	// 卖点渲染时的图层设置
	SellingPointRenderStyle map[string]interface{} `json:"SellingPointRenderStyle,omitempty"`

	// 卖点渲染模板，使用默认卖点格式使用default，自定义卖点提取类型使用custom
	SellingPointRenderTemplate string `json:"SellingPointRenderTemplate,omitempty"`

	// 卖点文本.形如"maidian1\nmaidian2\nmaidian3"或者"0.mdian1\n1.maindian2\n3.maiian".每个卖点不超过5个字.
	SellingPointText string `json:"SellingPointText,omitempty"`
}

type GetProductAIGCResultBodySellingPointConfigSellingPointExtractConfigItem struct {

	// 选择需要提取的卖点类型、取值：coresp（核心卖点提取）, shorttitle(短标题), productdescsp(商品信息类卖点)，productpromotionsp(导购激发类卖点)
	Category string `json:"Category,omitempty"`

	// 卖点渲染图层索引列表
	UniqueIDs []string `json:"UniqueIds,omitempty"`

	// 卖点信息
	Value string `json:"Value,omitempty"`
}

type GetProductAIGCResultQuery struct {

	// REQUIRED; 用于存储结果图和计量计费的服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	//
	//
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetProductAIGCResultRes struct {

	// REQUIRED
	ResponseMetadata *GetProductAIGCResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetProductAIGCResultResResult `json:"Result,omitempty"`
}

type GetProductAIGCResultResResponseMetadata struct {

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

// GetProductAIGCResultResResult - 视请求的接口而定
type GetProductAIGCResultResResult struct {

	// 商品场景图对应的美学得分，值越高表示图片越符合美学测评。
	AestheticScores []float64 `json:"AestheticScores,omitempty"`

	// AIGC图像评分。
	AigcImageScores map[string][]float64 `json:"AigcImageScores,omitempty"`

	// 商品场景图 URI 列表，未采用文字卖点。
	BackgroundImages []string                 `json:"BackgroundImages,omitempty"`
	ComposedJsons    []map[string]interface{} `json:"ComposedJsons,omitempty"`

	// 指定的商品信息
	Extra string `json:"Extra,omitempty"`

	// 商品场景图与正向提示词的匹配度得分，值越高表示匹配度越高。
	SDScores []float64 `json:"SDScores,omitempty"`

	// 指定的卖点渲染模板
	SellingPointRenderTemplate string `json:"SellingPointRenderTemplate,omitempty"`

	// 卖点文本信息
	SellingPointText string `json:"SellingPointText,omitempty"`

	// 生成的商品卖点图 URI 列表。排序规则为，当存在SDScores > 0的返回结果时，首个返回结果为SDScores > 0且总分(SDScores+AestheticScores)最高的结果图 URI，否则首个返回结果为SDScores
	// < 0且总分最高的结果；其余结果按总分降序排序。
	SellpointImages []string `json:"SellpointImages,omitempty"`
}

type GetResourceURLQuery struct {

	// REQUIRED; 域名。您可以通过调用 OpenAPI 获取服务下所有域名 [https://www.volcengine.com/docs/508/9379]获取。
	Domain string `json:"Domain" query:"Domain"`

	// REQUIRED; 资源所在的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 文件存储 Uri。您可以通过调用 OpenAPI 获取服务下的上传文件 [https://www.volcengine.com/docs/508/9392]获取。
	URI string `json:"URI" query:"URI"`

	// 创建模板时设置的图片输出格式，默认为 image，支持取值有：
	// * image：表示输出原格式；
	// * 静图格式：png、jpeg、heic、avif、webp;
	// * 动图格式：awebp、heif、avis。
	Format string `json:"Format,omitempty" query:"Format"`

	// 协议，默认为 http，隐私图片使用 https，公开图片支持取值 http 以及 https。
	Proto string `json:"Proto,omitempty" query:"Proto"`

	// 过期时长，最大限制为 1 年，默认为 1800s。 :::tip 仅当开启URL 鉴权 [https://www.volcengine.com/docs/508/128828]配置后，Timestamp配置生效。 :::
	Timestamp int `json:"Timestamp,omitempty" query:"Timestamp"`

	// 模板名称，缺省情况下表示无模板处理图片。您可以通过调用 OpenAPI 获取服务下所有图片模板 [https://www.volcengine.com/docs/508/9386]获取。
	Tpl string `json:"Tpl,omitempty" query:"Tpl"`
}

type GetResourceURLRes struct {

	// REQUIRED
	ResponseMetadata *GetResourceURLResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetResourceURLResResult           `json:"Result,omitempty"`
}

type GetResourceURLResResponseMetadata struct {

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

type GetResourceURLResResult struct {

	// REQUIRED; 结果图访问精简地址，与默认地址相比缺少 Bucket 部分。
	CompactURL string `json:"CompactURL"`

	// REQUIRED; 精简源文件地址，与默认地址相比缺少 Bucket 部分。
	ObjCompactURL string `json:"ObjCompactURL"`

	// REQUIRED; 默认源文件访问地址。
	ObjURL string `json:"ObjURL"`

	// REQUIRED; 结果图访问默认地址。
	URL string `json:"URL"`
}

type GetResponseHeaderValidateKeysRes struct {

	// REQUIRED
	ResponseMetadata *GetResponseHeaderValidateKeysResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetResponseHeaderValidateKeysResResult `json:"Result,omitempty"`
}

type GetResponseHeaderValidateKeysResResponseMetadata struct {

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

// GetResponseHeaderValidateKeysResResult - 视请求的接口而定
type GetResponseHeaderValidateKeysResResult struct {

	// REQUIRED; 合法的响应头 key 列表
	ValidateRespHdrKeys []string `json:"ValidateRespHdrKeys"`
}

type GetSegmentImageBody struct {

	// REQUIRED; 图片类型，支持可选择的模型如下。
	// * general：通用模型v1
	// * human：人脸模型v1
	// * product：商品模型v1
	// * humanv2：人脸模型v2
	// * productv2：商品模型v2
	Class string `json:"Class"`

	// REQUIRED; 输出图片格式，取值如下所示：
	// * png
	// * jpeg
	// * webp
	OutFormat string `json:"OutFormat"`

	// REQUIRED; 处理效果，当Class取值为humanv2或productv2时，默认为true。
	// * false：保留的图像主体边缘以粗线条处理，图像处理的效率更高。
	// * true：保留的图像主体边缘以发丝级细线条处理，图像处理后的效果更好。
	Refine bool `json:"Refine"`

	// REQUIRED; 待擦除原图的存储 URI 和 URL（公网可访问的 URL）。
	StoreURI string `json:"StoreUri"`

	// 描边设置，仅当Class取值humanv2或productv2时有效。 默认关闭，如果开启，抠出的结果图中保留的图像主体会外包一圈描边效果。
	Contour *GetSegmentImageBodyContour `json:"Contour,omitempty"`

	// 是否开启透明背景裁剪设置。默认false，关闭状态。如果开启，抠出的结果图会被裁剪到刚好包住保留的图像主体。
	TransBg bool `json:"TransBg,omitempty"`
}

// GetSegmentImageBodyContour - 描边设置，仅当Class取值humanv2或productv2时有效。 默认关闭，如果开启，抠出的结果图中保留的图像主体会外包一圈描边效果。
type GetSegmentImageBodyContour struct {

	// REQUIRED; 描边颜色，支持以 HEX、HSL、RGP 表示。例如HEX中白色为#FFFFFF。
	Color string `json:"Color"`

	// REQUIRED; 描边宽度，单位为 px。取值范围为 0 到正整数，默认 10px。
	Size int `json:"Size"`
}

type GetSegmentImageQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetSegmentImageRes struct {

	// REQUIRED
	ResponseMetadata *GetSegmentImageResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetSegmentImageResResult `json:"Result,omitempty"`
}

type GetSegmentImageResResponseMetadata struct {

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

// GetSegmentImageResResult - 视请求的接口而定
type GetSegmentImageResResult struct {

	// REQUIRED; 背景去除后的图片 URI。
	ResURI string `json:"ResUri"`
}

type GetServiceDomainsQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetServiceDomainsRes struct {
	ResponseMetadata *GetServiceDomainsResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           []*GetServiceDomainsResResultItem     `json:"Result,omitempty"`
}

type GetServiceDomainsResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type GetServiceDomainsResResultItem struct {

	// REQUIRED; 域名锁定状态
	LockStatus *GetServiceDomainsResResultItemLockStatus `json:"lock_status"`

	// 指向的 cname 值
	Cname string `json:"cname,omitempty"`

	// 创建时间
	CreateTime float64 `json:"create_time,omitempty"`

	// 域名
	Domain string `json:"domain,omitempty"`

	// HTTPS 配置
	HTTPSConfig *GetServiceDomainsResResultItemHTTPSConfig `json:"https_config,omitempty"`

	// 是否是默认域名，取值如下所示：
	// * true：默认域名
	// * false：非默认域名
	IsDefault bool `json:"is_default,omitempty"`

	// 域名锁定原因
	Reason string `json:"reason,omitempty"`

	// 地域
	Region string `json:"region,omitempty"`

	// 域名状态
	Status string `json:"status,omitempty"`

	// 域名申请人
	User string `json:"user,omitempty"`
}

// GetServiceDomainsResResultItemHTTPSConfig - HTTPS 配置
type GetServiceDomainsResResultItemHTTPSConfig struct {

	// REQUIRED; 证书 ID
	CertID string `json:"cert_id"`

	// 是否开启 HTTP2
	EnableHTTP2 bool `json:"enable_http2,omitempty"`
}

// GetServiceDomainsResResultItemLockStatus - 域名锁定状态
type GetServiceDomainsResResultItemLockStatus struct {

	// REQUIRED; 智能压缩是否锁定，取值如下所示：
	// * true：是
	// * false：否
	CompressionLocked bool `json:"compression_locked"`

	// REQUIRED; 整个域名是否锁定，取值如下所示：
	// * true：是，您无法手动修改其配置。如需修改，请提交工单 [https://console.volcengine.com/workorder/create?step=2&SubProductID=P00000080]联系技术支持。
	// * false：否
	DomainLocked bool `json:"domain_locked"`

	// REQUIRED; IP 访问限制是否锁定，取值如下所示：
	// * true：是
	// * false：否
	IPAccessRuleLocked bool `json:"ip_access_rule_locked"`

	// REQUIRED; Referer 防盗链是否锁定，取值如下所示：
	// * true：是
	// * false：否
	RefererAccessRuleLocked bool `json:"referer_access_rule_locked"`

	// REQUIRED; 远程鉴权是否锁定，取值如下所示：
	// * true：是
	// * false：否
	RemoteAuthLocked bool `json:"remote_auth_locked"`

	// REQUIRED; 响应头是否锁定，取值如下所示：
	// * true：是
	// * false：否
	ResponseHeaderLocked bool `json:"response_header_locked"`

	// REQUIRED; URL 鉴权是否锁定，取值如下所示：
	// * true：是
	// * false：否
	SignURLAuthLocked bool `json:"sign_url_auth_locked"`

	// REQUIRED; UA 访问限制是否锁定，取值如下所示：
	// * true：是
	// * false：否
	UaAccessRuleLocked bool `json:"ua_access_rule_locked"`

	// 域名被锁定原因
	Reason string `json:"reason,omitempty"`
}

type GetSyncAuditResultBody struct {

	// REQUIRED
	AuditAbility int `json:"AuditAbility"`

	// REQUIRED
	AuditDimensions []string `json:"AuditDimensions"`

	// REQUIRED
	AuditTextDimensions []string `json:"AuditTextDimensions"`

	// REQUIRED
	DataID string `json:"DataId"`

	// REQUIRED
	EnableLargeImageDetect bool `json:"EnableLargeImageDetect"`

	// REQUIRED
	ImageURI string `json:"ImageUri"`
}

type GetSyncAuditResultQuery struct {

	// REQUIRED; 服务唯一id
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetSyncAuditResultRes struct {

	// REQUIRED
	ResponseMetadata *GetSyncAuditResultResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetSyncAuditResultResResult `json:"Result,omitempty"`
}

type GetSyncAuditResultResResponseMetadata struct {

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

// GetSyncAuditResultResResult - 视请求的接口而定
type GetSyncAuditResultResResult struct {

	// REQUIRED
	Advice string `json:"Advice"`

	// REQUIRED
	DataID string `json:"DataId"`

	// REQUIRED
	ImageType string `json:"ImageType"`

	// REQUIRED
	ImageURI string   `json:"ImageUri"`
	Label    []string `json:"Label,omitempty"`
	SubLabel []string `json:"SubLabel,omitempty"`
}

type GetTemplatesFromBinQuery struct {

	// REQUIRED; 服务 ID
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// 是否按照模板创建时间升序查询，支持取值：true、false，默认为false。
	Asc string `json:"Asc,omitempty" query:"Asc"`

	// 分页获取条数，默认 10。
	Limit int `json:"Limit,omitempty" query:"Limit"`

	// 分页偏移。默认 0。取值为1，表示跳过第一条数据，从第二条数据开始取值。
	Offset int `json:"Offset,omitempty" query:"Offset"`

	// 仅返回模板名称包含该字符串的图片模板，不填或者为空则返回所有模板。
	TemplateNamePattern string `json:"TemplateNamePattern,omitempty" query:"TemplateNamePattern"`
}

type GetTemplatesFromBinRes struct {

	// REQUIRED
	ResponseMetadata *GetTemplatesFromBinResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetTemplatesFromBinResResult           `json:"Result,omitempty"`
}

type GetTemplatesFromBinResResponseMetadata struct {

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

type GetTemplatesFromBinResResult struct {

	// REQUIRED; 所有模板信息
	Templates []*GetTemplatesFromBinResResultTemplatesItem `json:"Templates"`

	// REQUIRED; 模板总数。
	Total int `json:"Total"`
}

type GetTemplatesFromBinResResultTemplatesItem struct {

	// REQUIRED; 模板的摘要信息。
	Abstract []string `json:"Abstract"`

	// REQUIRED; 指定图像自适应配置。
	AdaptiveFmt interface{} `json:"AdaptiveFmt"`

	// REQUIRED; 视频转动图配置。
	Animation interface{} `json:"Animation"`

	// REQUIRED; 模版 JSON 内容。
	Content string `json:"Content"`

	// REQUIRED; 模板创建时间。
	CreateAt string `json:"CreateAt"`

	// REQUIRED; 模版降级格式，仅针对 heic 格式有效。
	DemotionFormat string `json:"DemotionFormat"`

	// REQUIRED; 对图片的编辑操作。
	Filters []*Components8GbgysSchemasGettemplatesfrombinresPropertiesResultPropertiesTemplatesItemsPropertiesFiltersItems `json:"Filters"`

	// REQUIRED; 对图片编码使用的质量参数。
	OuputQuality int `json:"OuputQuality"`

	// REQUIRED; 用于图片服务输出时的图片编码。
	// * 取值png.use_quant表示是否开启 png quant 压缩，取值为true表示开启，取值为false表示关闭；
	// * 取值heic.sync表示使用 heic 同步编码，取值为true表示同步；
	// * 取值heic.timeout表示 heic 同步编码的超时时间，比如 20。
	OutputExtra interface{} `json:"OutputExtra"`

	// REQUIRED; 该模板计划使用的输出图片格式。
	// * 取值为image，表示输出原格式。
	// * 支持输出的静图格式：png、jpeg、heic、avif、webp。
	// * 支持输出的动图格式：awebp、heif、avis。
	OutputFormat string `json:"OutputFormat"`

	// REQUIRED; 图片模板的参数列表，URL 中下发参数的顺序需要跟列表中的保持一致。
	Parameters []string `json:"Parameters"`

	// REQUIRED; URL 的失效期，为 Unix 时间戳。
	ReqDeadline string `json:"ReqDeadline"`

	// REQUIRED; 服务 ID。
	ServiceID string `json:"ServiceID"`

	// REQUIRED; 视频截帧配置。
	Snapshot interface{} `json:"Snapshot"`

	// REQUIRED; 是否同步处理，仅针对 heic 格式有效。
	Sync bool `json:"Sync"`

	// REQUIRED; 模板名称。
	TemplateName string `json:"TemplateName"`

	// REQUIRED; 该模板的使用用例。
	Usage string `json:"Usage"`

	// REQUIRED; 模板是否开启鉴权。
	// * 取值为true，表示开启鉴权。
	// * 取值为false，表示关闭鉴权。
	WithSig bool `json:"WithSig"`
}

type GetURLFetchTaskQuery struct {

	// REQUIRED; 异步任务 ID，您可通过调用FetchImageUrl [https://www.volcengine.com/docs/508/1261301]接口获取该 ID。
	ID string `json:"Id" query:"Id"`

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type GetURLFetchTaskRes struct {

	// REQUIRED
	ResponseMetadata *GetURLFetchTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetURLFetchTaskResResult `json:"Result,omitempty"`
}

type GetURLFetchTaskResResponseMetadata struct {

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

// GetURLFetchTaskResResult - 视请求的接口而定
type GetURLFetchTaskResResult struct {

	// REQUIRED; 异步任务 ID
	ID string `json:"Id"`

	// REQUIRED; 任务状态，取值如下所示：
	// * Running：进行中
	// * Pending：排队中
	// * Failed：失败
	// * Success：成功
	Status string `json:"Status"`

	// REQUIRED; 资源下载链接
	URL string `json:"Url"`

	// 传入的回调地址，仅当Status取值Failed时有返回值。
	Callback string `json:"Callback,omitempty"`

	// 传入的回调内容，仅当Status取值Failed时有返回值。
	CallbackBody string `json:"CallbackBody,omitempty"`

	// 传入的回调内容类型，仅当Status取值Failed时有返回值。
	CallbackBodyType string `json:"CallbackBodyType,omitempty"`

	// 错误码，仅当Status取值Failed时有返回值。
	Code int `json:"Code,omitempty"`

	// 错误信息，仅当Status取值Failed时有返回值。
	Err string `json:"Err,omitempty"`

	// 资源上传后的资源 URI
	StoreURI string `json:"StoreUri,omitempty"`
}

type GetVendorBucketsBody struct {

	// REQUIRED; Access Key。是与 Secret Key 同时填写的，为了保证有访问源数据桶的权限。
	AK string `json:"AK"`

	// REQUIRED; S3 协议 Endpoint，需以http://或https://开头。仅当Vendor为S3时必填。
	Endpoint string `json:"Endpoint"`

	// REQUIRED; Secret Key。是与 Access Key 同时填写的，为了保证有访问源数据桶的权限。
	SK string `json:"SK"`

	// REQUIRED; 服务商。取值如下所示：
	// * OSS：阿里云
	// * COS：腾讯云
	// * KODO：七牛云
	// * BOS：百度云
	// * OBS：华为云
	// * Ucloud：Ucloud file
	// * AWS：AWS 国际站
	// * S3：其他 S3 协议存储
	// * URL：以上传 URL 列表的方式迁移
	Vendor string `json:"Vendor"`

	// Bucket 所在地区。仅当Vendor非URL/OSS/KODO/AWS时为必填。
	Region string `json:"Region,omitempty"`
}

type GetVendorBucketsRes struct {

	// REQUIRED
	ResponseMetadata *GetVendorBucketsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetVendorBucketsResResult `json:"Result,omitempty"`
}

type GetVendorBucketsResResponseMetadata struct {

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

// GetVendorBucketsResResult - 视请求的接口而定
type GetVendorBucketsResResult struct {

	// REQUIRED; bucket 列表
	Buckets []string `json:"Buckets"`
}

type PreviewImageUploadFileQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`

	// REQUIRED; 文件存储 URI。
	// * 您可以在 veImageX 控制台资源管理 [https://console.volcengine.com/imagex/resource_manage/]页面，在已上传文件的名称列获取。
	// * 您也可以通过 OpenAPI 的方式获取，具体请参考获取服务下的上传文件 [https://www.volcengine.com/docs/508/9392]。
	StoreURI string `json:"StoreUri" query:"StoreUri"`
}

type PreviewImageUploadFileRes struct {

	// REQUIRED
	ResponseMetadata *PreviewImageUploadFileResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *PreviewImageUploadFileResResult `json:"Result,omitempty"`
}

type PreviewImageUploadFileResResponseMetadata struct {

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

// PreviewImageUploadFileResResult - 视请求的接口而定
type PreviewImageUploadFileResResult struct {

	// REQUIRED; 图片格式。
	ImageFormat string `json:"ImageFormat"`

	// REQUIRED; 图片帧数。
	ImageFrames int `json:"ImageFrames"`

	// REQUIRED; 图片高度。
	ImageHeight int `json:"ImageHeight"`

	// REQUIRED; 图片大小，单位为字节。
	ImageSize int `json:"ImageSize"`

	// REQUIRED; 图片宽度。
	ImageWidth int `json:"ImageWidth"`

	// REQUIRED; 服务 ID。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 图片的 Content-Type 值。
	StorageContentType string `json:"StorageContentType"`

	// REQUIRED; 文件存储 URI。
	StoreURI string `json:"StoreUri"`

	// 图片播放时间，单位为毫秒，仅当图片为动态图时返回。
	ImageDuration int `json:"ImageDuration,omitempty"`
}

type ReportEventRes struct {

	// REQUIRED
	ResponseMetadata *ReportEventResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type ReportEventResResponseMetadata struct {

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

type RerunImageMigrateTaskQuery struct {

	// REQUIRED; 仅当任务状态为Partial时生效。 任务 ID，请参考GetImageMigrateTasks [https://www.volcengine.com/docs/508/1108670]获取返回的任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`

	// 任务地区（即任务目标服务的地区），默认空，返回国内任务。
	// * cn：国内
	// * sg：新加坡
	Region string `json:"Region,omitempty" query:"Region"`
}

type RerunImageMigrateTaskRes struct {

	// REQUIRED
	ResponseMetadata *RerunImageMigrateTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type RerunImageMigrateTaskResResponseMetadata struct {

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

type SetDefaultDomainBody struct {

	// REQUIRED; 指定新的默认域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"Domain"`
}

type SetDefaultDomainRes struct {

	// REQUIRED
	ResponseMetadata *SetDefaultDomainResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type SetDefaultDomainResResponseMetadata struct {

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

type TerminateImageMigrateTaskQuery struct {

	// REQUIRED; 任务 ID，请参考GetImageMigrateTasks [https://www.volcengine.com/docs/508/1108670]获取返回的任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`

	// 任务地区（即任务目标服务的地区），默认空，返回国内任务。
	// * cn：国内
	// * sg：新加坡
	Region string `json:"Region,omitempty" query:"Region"`
}

type TerminateImageMigrateTaskRes struct {

	// REQUIRED
	ResponseMetadata *TerminateImageMigrateTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type TerminateImageMigrateTaskResResponseMetadata struct {

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

type UpdateAdvanceBody struct {

	// REQUIRED; 高级配置
	Advance *UpdateAdvanceBodyAdvance `json:"advance"`

	// REQUIRED
	Domain string `json:"domain"`
}

// UpdateAdvanceBodyAdvance - 高级配置
type UpdateAdvanceBodyAdvance struct {

	// REQUIRED; 是否开启 Brotli 压缩，取值如下所示：
	// * true：开启
	// * false：关闭 :::tip 支持同时配置 Gzip 压缩和 Brotli 压缩，详细内容请参考智能压缩 [https://www.volcengine.com/docs/508/75858]。 :::
	EnableBr bool `json:"enable_br"`

	// REQUIRED; 是否开启 Gzip 压缩，取值如下所示：
	// * true：开启
	// * false：关闭 :::tip 支持同时配置 Gzip 压缩和 Brotli 压缩，详细内容请参考智能压缩 [https://www.volcengine.com/docs/508/75858]。 :::
	EnableGzip bool `json:"enable_gzip"`

	// REQUIRED; 是否开启 IPV6，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableIPv6 bool `json:"enable_ipv6"`
}

type UpdateAdvanceQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过调用 GetAllImageServices [https://www.volcengine.com/docs/508/9360] 接口方式获取服务 ID。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateAdvanceRes struct {

	// REQUIRED
	ResponseMetadata *UpdateAdvanceResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateAdvanceResResponseMetadata struct {

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

type UpdateAuditImageStatusBody struct {

	// REQUIRED; 审核图片 ID。您可通过调用获取审核任务结果 [https://www.volcengine.com/docs/508/1160410]查看待更新状态的图片条目 ID。
	EntryID string `json:"EntryId"`

	// REQUIRED; original->frozen执行冻结操作。frozen->original执行恢复操作
	Status string `json:"Status"`

	// REQUIRED; 任务 ID，您可通过调用查询所有审核任务 [https://www.volcengine.com/docs/508/1160409]获取所需的任务 ID。
	TaskID string `json:"TaskId"`
}

type UpdateAuditImageStatusRes struct {

	// REQUIRED
	ResponseMetadata *UpdateAuditImageStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateAuditImageStatusResResponseMetadata struct {

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

type UpdateDomainAdaptiveFmtBody struct {

	// REQUIRED; 是否开启自适应，取值如下所示：
	// * true：开启自适应
	// * false：关闭自适应
	AdaptFmt bool `json:"AdaptFmt"`

	// REQUIRED; 自适应格式列表，取值如下所示：
	// * webp：WEBP 自适应
	// * heic：HEIC 自适应
	// * avif：AVIF 自适应
	AdaptFormats []string `json:"AdaptFormats"`

	// REQUIRED; 是否开启体积校验，取值如下所示：
	// * true：开启。开启后会对经自适应编码后的图片体积和编码前原图体积进行对比，若编码后体积更小则输出编码后图片；否则输出原图。
	// * false：关闭
	CheckAdaptFsize bool `json:"CheckAdaptFsize"`

	// REQUIRED; 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"Domain"`
}

type UpdateDomainAdaptiveFmtQuery struct {

	// REQUIRED; 待修改配置的域名的所属服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateDomainAdaptiveFmtRes struct {
	ResponseMetadata *UpdateDomainAdaptiveFmtResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                                      `json:"Result,omitempty"`
}

type UpdateDomainAdaptiveFmtResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateFileStorageClassBody struct {

	// REQUIRED; * STANDARD 标准存储
	// * IA 低频存储
	// * ARCHIVE_FR 归档闪回存储
	// * COLD_ARCHIVE 冷归档存储
	// * ARCHIVE 归档存储
	StorageClass string `json:"StorageClass"`

	// REQUIRED; 文件存储 URI。
	// * 您可以在 veImageX 控制台资源管理 [https://console.volcengine.com/imagex/resource_manage/]页面，在已上传文件的名称列获取资源 URI。
	// * 您也可以通过 OpenAPI 的方式获取 URI，具体请参考获取服务下全部上传文件 [https://www.volcengine.com/docs/508/9393]。
	StoreURI string `json:"StoreUri"`
}

type UpdateFileStorageClassQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateFileStorageClassRes struct {

	// REQUIRED
	ResponseMetadata *UpdateFileStorageClassResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateFileStorageClassResResponseMetadata struct {

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

type UpdateHTTPSBody struct {

	// REQUIRED; 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"domain"`

	// REQUIRED; https 配置
	HTTPS *UpdateHTTPSBodyHTTPS `json:"https"`
}

// UpdateHTTPSBodyHTTPS - https 配置
type UpdateHTTPSBodyHTTPS struct {

	// REQUIRED; ["tlsv1.0","tlsv1.1","tlsv1.2","tlsv1.3"]
	TLSVersions []string `json:"tls_versions"`

	// 需要关联的证书 ID，若enable_https为true，则为必填。
	CertID string `json:"cert_id,omitempty"`

	// 是否开启强制跳转，支持取值如下所示：
	// * true：开启
	// * false：关闭
	EnableForceRedirect bool `json:"enable_force_redirect,omitempty"`

	// 是否开启 http2，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableHTTP2 bool `json:"enable_http2,omitempty"`

	// 是否开启 https，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableHTTPS bool `json:"enable_https,omitempty"`

	// 是否开启 ocsp 装订，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableOcsp bool `json:"enable_ocsp,omitempty"`

	// 是否开启 quic 协议支持，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableQuic bool `json:"enable_quic,omitempty"`

	// 301、302HTTP 强制跳转 HTTPS只支持301
	ForceRedirectCode string `json:"force_redirect_code,omitempty"`

	// http2https、https2http
	ForceRedirectType string `json:"force_redirect_type,omitempty"`

	// 配置hsts
	Hsts *UpdateHTTPSBodyHTTPSHsts `json:"hsts,omitempty"`
}

// UpdateHTTPSBodyHTTPSHsts - 配置hsts
type UpdateHTTPSBodyHTTPSHsts struct {

	// 是否开启hsts
	Enable bool `json:"enable,omitempty"`

	// include、exclude
	Subdomain string `json:"subdomain,omitempty"`

	// 表示 Strict-Transport-Security 响应头在浏览器中的缓存过期时间，单位是秒。该参数的取值范围是 0 - 31,536,000。31,536,000 秒表示 365 天。如果该参数值为 0，其效果等同于禁用 HSTS
	// 设置。该参数的默认值是 0。
	TTL int `json:"ttl,omitempty"`
}

type UpdateHTTPSQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateHTTPSRes struct {
	ResponseMetadata *UpdateHTTPSResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                          `json:"Result,omitempty"`
}

type UpdateHTTPSResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateImageAIProcessQueueBody struct {

	// REQUIRED; aisuperresolution、bgfill、super_resolution、InfiniteCreations
	Attribute string `json:"Attribute"`

	// REQUIRED; 队列ID
	ID string `json:"Id"`

	// REQUIRED; 队列名
	Name string `json:"Name"`

	// 任务队列回调设置。
	CallbackConf *UpdateImageAIProcessQueueBodyCallbackConf `json:"CallbackConf,omitempty"`

	// 队列描述
	Desc string `json:"Desc,omitempty"`

	// 是否开启回调配置
	EnableCallback bool `json:"EnableCallback,omitempty"`
}

// UpdateImageAIProcessQueueBodyCallbackConf - 任务队列回调设置。
type UpdateImageAIProcessQueueBodyCallbackConf struct {

	// 回调附带参数
	Args string `json:"Args,omitempty"`

	// 回调参数类型，JSON或者XML
	DataFormat string `json:"DataFormat,omitempty"`

	// 回调地址
	Endpoint string `json:"Endpoint,omitempty"`

	// 回调请求协议
	Method string `json:"Method,omitempty"`

	// 回调触发类型，任务级别或者单个条目
	Type string `json:"Type,omitempty"`
}

type UpdateImageAIProcessQueueRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageAIProcessQueueResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageAIProcessQueueResResponseMetadata struct {

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

type UpdateImageAIProcessQueueStatusBody struct {

	// REQUIRED; 队列id
	ID string `json:"Id"`

	// REQUIRED; 状态Pending,Running
	Status string `json:"Status"`
}

type UpdateImageAIProcessQueueStatusRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageAIProcessQueueStatusResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageAIProcessQueueStatusResResponseMetadata struct {

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

type UpdateImageAnalyzeTaskBody struct {

	// REQUIRED; 任务名称
	Name string `json:"Name"`

	// REQUIRED; 服务 ID
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 待更新的任务 ID，您可以通过调用 GetImageAnalyzeTasks [https://www.volcengine.com/docs/508/1160417] 获取指定地区全部离线评估任务详情。
	TaskID string `json:"TaskId"`

	// 任务描述
	Desc string `json:"Desc,omitempty"`

	// 仅当Type取值UriFile时，配置有效。 是否模拟模板每阶段输出，取值如下所示：
	// * true：是，一个模版中可以选择多种图像处理, 模拟输出时会将所有的处理逐步叠加并编码为最终图片格式运行并输出评估结果。
	// * false：否。
	EvalPerStage bool `json:"EvalPerStage,omitempty"`

	// txt 评估文件的 Store URI，该文件需上传至指定服务对应存储中。
	// * Type取值UrlFile时，填写合法 URL
	// * Type取值UriFile时，填写指定服务的存储 URI
	ResURI []string `json:"ResUri,omitempty"`

	// 仅当Type取值UriFile时，配置有效。 模板名称，您可通过调用GetAllImageTemplates [https://www.volcengine.com/docs/508/9386]获取服务下所有已创建的TemplateName。
	Tpl string `json:"Tpl,omitempty"`
}

type UpdateImageAnalyzeTaskRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageAnalyzeTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateImageAnalyzeTaskResResponseMetadata struct {

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

type UpdateImageAnalyzeTaskStatusBody struct {

	// REQUIRED; 更新后的任务状态。取值如下所示：
	// * Running：进行中
	// * Suspend：暂停
	// * Done：结束
	Status string `json:"Status"`

	// REQUIRED; 待更新的任务 ID，您可以通过调用GetImageAnalyzeTasks [https://www.volcengine.com/docs/508/1160417]获取指定地区全部离线评估任务 ID。
	TaskID string `json:"TaskId"`
}

type UpdateImageAnalyzeTaskStatusRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageAnalyzeTaskStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateImageAnalyzeTaskStatusResResponseMetadata struct {

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

type UpdateImageAuditTaskBody struct {

	// REQUIRED; 审核维度，根据审核能力的不同，其具体取值不同。基础审核与智能审核之间不支持混用。
	// * 基础安全审核，仅当 AuditAbility 取值为 0 时，配置生效。
	//
	// * govern：涉政
	// * porn ：涉黄
	// * illegal：违法违规
	// * terror：涉暴
	//
	//
	// * 智能安全审核，仅当 AuditAbility 取值为 1 时，配置生效。
	//
	// * 图像风险识别 * porn ：涉黄，主要适用于通用色情、色情动作、性行为、性暗示、性分泌物、色情动漫、色情裸露等涉黄场景的风险识别
	// * sensitive1 ：涉敏1，具体指涉及暴恐风险
	// * sensitive2：涉敏2，具体值涉及政治内容风险
	// * forbidden：违禁，主要适用于打架斗殴、爆炸、劣迹艺人等场景的风险识别
	// * uncomfortable：引人不适，主要适用于恶心、恐怖、尸体、血腥等引人不适场景的风险识别
	// * qrcode：二维码，主要适用于识别常见二维码（QQ、微信、其他二维码等）
	// * badpicture：不良画面，主要适用于自我伤害、丧葬、不当车播、吸烟/纹身/竖中指等不良社会风气的风险识别
	// * sexy：性感低俗，主要适用于舌吻、穿衣性行为、擦边裸露等多种性感低俗场景的风险识别
	// * age：年龄，主要适用于图中人物对应的年龄段识别
	// * underage：未成年相关，主要适用于儿童色情、儿童邪典等风险识别
	// * quality：图片质量，主要适用于图片模糊、纯色边框、纯色屏等风险识别
	//
	//
	// * 图文风险识别，您可在 AuditTextDimensions 配置文字审核的维度。
	//
	// :::tip 您可将智能安全审核的图像风险识别和图文风险识别搭配使用。 :::
	AuditDimensions []string `json:"AuditDimensions"`

	// REQUIRED; 指定待更新审核任务所在的服务 ID，您可通过调用查询所有审核任务 [https://www.volcengine.com/docs/508/1158717]获取待更新任务对应的服务 ID。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 任务 ID，您可通过调用 查询所有审核任务 [https://www.volcengine.com/docs/508/1158717] 获取所需的任务 ID。
	TaskID string `json:"TaskId"`

	// 仅当EnableAuditRange取值1时，配置生效。 指定前缀审核，若你希望对某个目录进行审核，请设置路径为对应的目录名，以/结尾。例如123/
	AuditPrefix []string `json:"AuditPrefix,omitempty"`

	// 智能安全审核类型下图片文本审核的具体维度，取值如下所示：
	// * ad：广告，综合图像及文字内容智能识别广告
	// * defraud：诈骗，综合图像及文字内容智能识别诈骗
	// * charillegal：文字违规，图片上存在涉黄、涉敏、违禁等违规文字
	// :::tip 仅当 AuditDimensions 取值为智能安全审核模型时，您可将 AuditTextDimensions 与 AuditDimensions 搭配使用。 :::
	AuditTextDimensions []string `json:"AuditTextDimensions,omitempty"`

	// 回调类型，取值需要与 AuditDimensions 审核维度保持一致或少于 AuditDimensions。
	// 例如，AuditDimensions 取值 ["pron","sexy"]，AuditTextDimensions 取值 ["ad"]，支持您将 FreezeDimensions 取值 ["pron","sexy","ad"] 、 ["pron","sexy"]、["pron","ad"]
	// 和 ["sexy","ad"] 任意一种。
	CallbackDimensions []string `json:"CallbackDimensions,omitempty"`

	// 回调图片类型，取值如下所示：
	// * normal：正常图片
	//
	//
	// * problem：问题图片
	//
	//
	// * frozen：冻结图片
	//
	//
	// * fail：审核失败图片
	CallbackImageTypes []string `json:"CallbackImageTypes,omitempty"`

	// 回调 URL，veImageX 以 Post 方式向业务服务器发送 JSON 格式回调数据。具体回调参数请参考回调内容 [https://www.volcengine.com/docs/508/134676#%E5%9B%9E%E8%B0%83%E5%86%85%E5%AE%B9]。
	CallbackURL string `json:"CallbackUrl,omitempty"`

	// 仅当Type取值Upload时，配置生效。 审核范围，取值如下所示：
	// * 0：（默认）不限范围
	// * 1：指定范围
	EnableAuditRange int `json:"EnableAuditRange,omitempty"`

	// 是否开启回调，取值如下所示：
	// * true：开启
	// * false：（默认）不开启
	EnableCallback bool `json:"EnableCallback,omitempty"`

	// 是否开启冻结，取值如下所示：
	// * true：开启
	// * false：（默认）不开启
	EnableFreeze bool `json:"EnableFreeze,omitempty"`

	// 图片审核仅支持审核 5MB 以下的图片，若您的图片大小在 5MB~32MB，您可以开启大图审核功能，veImageX 会对图片压缩后再进行审核。开启后，将对压缩能力按照基础图片处理
	// [https://www.volcengine.com/docs/508/65935#%E5%9F%BA%E7%A1%80%E5%9B%BE%E5%83%8F%E5%A4%84%E7%90%86%E6%9C%8D%E5%8A%A1]进行计费。（您每月可使用
	// 20TB 的免费额度） 取值如下所示：
	// * true：开启
	// * false：（默认）不开启
	// :::tip
	// * 若未开启大图审核且图片大小 ≥ 5 MB，可能会导致系统超时报错；
	// * 若已开启大图审核但图片大小 ≥ 32 MB，可能会导致系统超时报错。 :::
	EnableLargeImageDetect bool `json:"EnableLargeImageDetect,omitempty"`

	// 若开启冻结，则不可为空
	FreezeDimensions []string `json:"FreezeDimensions,omitempty"`

	// 若开启冻结，则不可为空
	FreezeStrategy int `json:"FreezeStrategy,omitempty"`

	// 若开启冻结，则不可为空
	FreezeType []string `json:"FreezeType,omitempty"`

	// 仅当 EnableAuditRange 取值 1 时，配置生效。 指定前缀不审核，若你希望对某个目录不进行审核，请设置路径为对应的目录名，以/结尾。例如123/
	NoAuditPrefix []string `json:"NoAuditPrefix,omitempty"`

	// 任务地区。当前仅支持取值 cn，表示国内。
	Region string `json:"Region,omitempty"`

	// 仅当 Type 为 UrlFile 时，配置生效。
	// 审核文件的 StoreUri，为 .txt 文件，该文件需上传至指定服务对应存储中。该 txt 文件内需填写待审核图片文件的 URL，每行填写一个，最多可填写 10000 行。
	ResURI []string `json:"ResUri,omitempty"`
}

type UpdateImageAuditTaskRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageAuditTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateImageAuditTaskResResponseMetadata struct {

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

type UpdateImageAuditTaskStatusBody struct {

	// REQUIRED; 更新后的任务状态。取值如下所示：
	// * Running：审核中
	// * Suspend：审核暂停
	// * Done：审核完成
	// * Cancel：审核取消
	Status string `json:"Status"`

	// REQUIRED; 待更新的任务 ID，您可通过调用查询所有审核任务 [https://www.volcengine.com/docs/508/1160409]获取所需的任务 ID。
	TaskID string `json:"TaskId"`
}

type UpdateImageAuditTaskStatusRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageAuditTaskStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateImageAuditTaskStatusResResponseMetadata struct {

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

type UpdateImageAuthKeyBody struct {

	// 主鉴权 key，长度为 8-32 字节，为空则不更新主鉴权值。
	PrimaryKey string `json:"PrimaryKey,omitempty"`

	// 备鉴权 key，长度为 8-32 字节，为空则不更新备鉴权值。
	SecondaryKey string `json:"SecondaryKey,omitempty"`
}

type UpdateImageAuthKeyQuery struct {

	// REQUIRED; 待更新配置的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageAuthKeyRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageAuthKeyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageAuthKeyResResponseMetadata struct {

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

type UpdateImageDomainAreaAccessBody struct {

	// REQUIRED; 区域访问限制配置
	AreaAccess *UpdateImageDomainAreaAccessBodyAreaAccess `json:"area_access"`

	// REQUIRED; 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"domain"`
}

// UpdateImageDomainAreaAccessBodyAreaAccess - 区域访问限制配置
type UpdateImageDomainAreaAccessBodyAreaAccess struct {

	// REQUIRED; 对照名称表 https://www.volcengine.com/docs/6454/131750
	Areas []string `json:"areas"`

	// REQUIRED; 是否开启区域限制
	Enable bool `json:"enable"`

	// REQUIRED; 黑白名单设置类型，取值如下所示：
	// * deny：黑名单
	// * allow：白名单
	RuleType string `json:"rule_type"`
}

type UpdateImageDomainAreaAccessQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageDomainAreaAccessRes struct {
	ResponseMetadata *UpdateImageDomainAreaAccessResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                                          `json:"Result,omitempty"`
}

type UpdateImageDomainAreaAccessResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateImageDomainBandwidthLimitBody struct {

	// REQUIRED; 带宽限速配置
	BandwidthLimit *UpdateImageDomainBandwidthLimitBodyBandwidthLimit `json:"bandwidth_limit"`

	// REQUIRED; 域名。您可以通过调用 GetServiceDomains [https://www.volcengine.com/docs/508/9379] 接口获取域名。
	Domain string `json:"domain"`
}

// UpdateImageDomainBandwidthLimitBodyBandwidthLimit - 带宽限速配置
type UpdateImageDomainBandwidthLimitBodyBandwidthLimit struct {

	// REQUIRED; 是否开启带宽限制功能，取值如下所示：
	// * true：开启
	// * false：关闭 :::tip 仅当 enabled 为 true 时，threshold、limit_type 等配置项有效。 :::
	Enabled bool `json:"enabled"`

	// REQUIRED; 全局带宽阈值，指定加速域名的带宽阈值。单位为 bps，取值范围为 [1, 1000000000000000] 的整数。单位换算：1 Gbps = 1000 Mbps。
	LimitType string `json:"limit_type"`

	// REQUIRED; 设置节点响应访问请求的速度下限，在 veImageX 逐步降低最大速度的过程中，最大速度不会低于该配置。 单位：B/S，取值范围为 [1,1073741824000]。默认值为 + 4096 KB/S。 单位换算：1
	// KB/S = 1024 B/S。
	// * 当 limit_type 为 downloadspeedlimit 时，表示每个请求的最低速度。
	// * 当 limit_type 为 speedlimit 时，表示每个 IP 地址的最低速度。 :::tip 当 limit_type 为 randomreject 时，不支持自定义该配置。 :::
	SpeedLimitRate int `json:"speed_limit_rate"`

	// REQUIRED; 初始速率，即初始最大速度。限速发生时， veImageX 会从该速度开始，逐步降低最大速度。 单位：B/S，取值范围为[1,1073741824000]。 单位换算：1 KB/S = 1024 B/S。
	// * 当 limit_type 为 downloadspeedlimit 时，表示每个请求的初始最大速度。
	// * 当 limit_type 为 speedlimit 时，表示每个 IP 地址的初始最大速度。 :::tip 当 limit_type 为 randomreject 时，不支持自定义该配置。 :::
	SpeedLimitRateMax int `json:"speed_limit_rate_max"`

	// REQUIRED; 全局带宽阈值，指定加速域名的带宽阈值。单位为 bps，取值范围为 [1, 1000000000000000] 的整数。 单位换算：1 Gbps = 1000 Mbps。
	Threshold int `json:"threshold"`
}

type UpdateImageDomainBandwidthLimitQuery struct {

	// REQUIRED; 服务 ID。您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。您也可以通过调用
	// GetAllImageServices [https://www.volcengine.com/docs/508/9360] 接口方式获取服务 ID。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageDomainBandwidthLimitRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageDomainBandwidthLimitResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *UpdateImageDomainBandwidthLimitResResult `json:"Result,omitempty"`
}

type UpdateImageDomainBandwidthLimitResResponseMetadata struct {

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

// UpdateImageDomainBandwidthLimitResResult - 视请求的接口而定
type UpdateImageDomainBandwidthLimitResResult struct {

	// REQUIRED; 通知描述
	Msg string `json:"msg"`
}

type UpdateImageDomainConfigBody struct {

	// REQUIRED; 域名列表，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domains []string `json:"domains"`

	// 自适应格式配置
	Adaptfmt *UpdateImageDomainConfigBodyAdaptfmt `json:"adaptfmt,omitempty"`

	// 高级配置
	Advanced *UpdateImageDomainConfigBodyAdvanced `json:"advanced,omitempty"`

	// 区域访问限制，不传不更新
	AreaACL *UpdateImageDomainConfigBodyAreaACL `json:"area_acl,omitempty"`

	// 集智瘦身配置
	DoSlim *UpdateImageDomainConfigBodyDoSlim `json:"do_slim,omitempty"`

	// 全球加速配置
	GlobalAcceleration *UpdateImageDomainConfigBodyGlobalAcceleration `json:"global_acceleration,omitempty"`

	// HTTPS 配置
	HTTPS *UpdateImageDomainConfigBodyHTTPS `json:"https,omitempty"`

	// IP 黑白名单配置
	IPAuth *UpdateImageDomainConfigBodyIPAuth `json:"ip_auth,omitempty"`

	// 页面优化设置，仅素材托管服务下域名支持配置。
	PageOptimization *UpdateImageDomainConfigBodyPageOptimization `json:"page_optimization,omitempty"`

	// Referer 防盗链配置
	RefererLink *UpdateImageDomainConfigBodyRefererLink `json:"referer_link,omitempty"`

	// 远程鉴权设置
	RemoteAuth *UpdateImageDomainConfigBodyRemoteAuth `json:"remote_auth,omitempty"`

	// HTTP 响应头配置
	RespHdrs []*UpdateImageDomainConfigBodyRespHdrsItem `json:"resp_hdrs,omitempty"`

	// 共享缓存配置。共享缓存允许同账号下多个加速域名共享同一份节点上的缓存。在 veImageX 中，您可以通过设置共享缓存配置，使各个子站点之间可共享相同的公共资源，以减少带宽的使用，提高资源命中率。详细功能说明参看共享缓存 [https://www.volcengine.com/docs/508/196769]。
	// :::tip 共享缓存为白名单功能，请提交工单 [https://console.volcengine.com/workorder/create?step=2&SubProductID=P00000080]联系技术支持为您的账号开启配置能力。
	// :::
	ShareCache *UpdateImageDomainConfigBodyShareCache `json:"share_cache,omitempty"`

	// URL 鉴权配置
	URLAuth *UpdateImageDomainConfigBodyURLAuth `json:"url_auth,omitempty"`

	// UA 访问限制配置
	UserAgentACL *UpdateImageDomainConfigBodyUserAgentACL `json:"user_agent_acl,omitempty"`
}

// UpdateImageDomainConfigBodyAdaptfmt - 自适应格式配置
type UpdateImageDomainConfigBodyAdaptfmt struct {

	// REQUIRED; 自适应格式列表，支持以下取值：
	// * webp：WEBP 自适应
	// * heic：HEIC 自适应
	// * avif：AVIF 自适应
	AdaptFormats []string `json:"adapt_formats"`

	// REQUIRED; 是否开启体积校验，取值如下所示：
	// * true：开启
	// * false：关闭
	CheckAdaptFsize bool `json:"check_adapt_fsize"`

	// REQUIRED; 是否开启自适应，取值如下所示：
	// * true：开启自适应
	// * false：关闭自适应
	Enabled bool `json:"enabled"`
}

// UpdateImageDomainConfigBodyAdvanced - 高级配置
type UpdateImageDomainConfigBodyAdvanced struct {

	// 是否开启 Brotli 压缩，取值如下所示：
	// * true：开启
	// * false：关闭
	// :::tip 支持同时配置 Gzip 压缩和 Brotli 压缩，详细内容请参考智能压缩 [https://www.volcengine.com/docs/508/75858]。 :::
	EnableBr bool `json:"enable_br,omitempty"`

	// 是否开启 Gzip 压缩，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableGzip bool `json:"enable_gzip,omitempty"`

	// 是否开启 IPV6，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableIPv6 bool `json:"enable_ipv6,omitempty"`
}

// UpdateImageDomainConfigBodyAreaACL - 区域访问限制，不传不更新
type UpdateImageDomainConfigBodyAreaACL struct {

	// REQUIRED; 对照名称表 https://www.volcengine.com/docs/6454/131750
	Areas []string `json:"areas"`

	// REQUIRED; 是否开启区域限制，取值如下所示：
	// * true：开启
	// * false：未开启
	Enabled bool `json:"enabled"`

	// REQUIRED; 黑白名单设置类型，取值如下所示：
	// * deny：黑名单
	// * allow：白名单
	RuleType string `json:"rule_type"`
}

// UpdateImageDomainConfigBodyDoSlim - 集智瘦身配置
type UpdateImageDomainConfigBodyDoSlim struct {

	// REQUIRED; 是否关闭持久化。取值如下所示：
	// * true：关闭
	// * false：开启
	DiscardSlimedFile bool `json:"discard_slimed_file"`

	// REQUIRED; 是否开启集智瘦身，取值如下所示：
	// * true：开启集智瘦身
	// * false：关闭集智瘦身
	Enabled bool `json:"enabled"`
}

// UpdateImageDomainConfigBodyGlobalAcceleration - 全球加速配置
type UpdateImageDomainConfigBodyGlobalAcceleration struct {

	// REQUIRED; 是否开启全球加速，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"enabled"`
}

// UpdateImageDomainConfigBodyHTTPS - HTTPS 配置
type UpdateImageDomainConfigBodyHTTPS struct {

	// REQUIRED; 是否开启 http2，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableHTTP2 bool `json:"enable_http2"`

	// REQUIRED; 是否开启 https，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableHTTPS bool `json:"enable_https"`

	// REQUIRED; ["tlsv1.0","tlsv1.1","tlsv1.2","tlsv1.3"]
	TLSVersions []string `json:"tls_versions"`

	// 需要关联的证书 ID，若enable_https为true，则为必填。
	CertID string `json:"cert_id,omitempty"`

	// 是否开启强制跳转，取值如下所示：
	// * true：开启
	// * false：关闭
	EnableForceRedirect bool `json:"enable_force_redirect,omitempty"`

	// 301、302HTTP 强制跳转 HTTPS只支持301
	ForceRedirectCode string `json:"force_redirect_code,omitempty"`

	// http2https、https2http
	ForceRedirectType string `json:"force_redirect_type,omitempty"`

	// HSTS 配置
	Hsts *UpdateImageDomainConfigBodyHTTPSHsts `json:"hsts,omitempty"`
}

// UpdateImageDomainConfigBodyHTTPSHsts - HSTS 配置
type UpdateImageDomainConfigBodyHTTPSHsts struct {

	// 是否开启 HSTS 配置，取值如下所示：
	// * true：是
	// * false：否
	Enabled bool `json:"enabled,omitempty"`

	// HSTS 配置是否也应用于加速域名的子域名。取值如下所示：
	// * include：应用于子域名站点。
	// * exclude：（默认）不应用于子域名站点。
	Subdomain string `json:"subdomain,omitempty"`

	// 如果 enable_https 是 true，该参数为必填。 Strict-Transport-Security 响应头在浏览器中的缓存过期时间，单位是秒。取值范围是 [0,31,536,000]。31,536,000 秒表示 365 天。如果该参数值指定为
	// 0，其效果等同于禁用 HSTS 设置。
	TTL int `json:"ttl,omitempty"`
}

// UpdateImageDomainConfigBodyIPAuth - IP 黑白名单配置
type UpdateImageDomainConfigBodyIPAuth struct {

	// REQUIRED; 是否开启黑白名单配置，取值如下所示：
	// * true：开启黑白名单配置
	// * false：关闭黑白名单配置
	Enabled bool `json:"enabled"`

	// REQUIRED; 是否是 IP 白名单，取值如下所示：
	// * true：配置白名单
	// * false：配置黑名单
	IsWhiteMode bool `json:"is_white_mode"`

	// REQUIRED; 黑白名单 IP 地址，最大限制为 1000。
	Values []string `json:"values"`
}

// UpdateImageDomainConfigBodyPageOptimization - 页面优化设置，仅素材托管服务下域名支持配置。
type UpdateImageDomainConfigBodyPageOptimization struct {

	// REQUIRED; 是否开启页面优化，取值如下所示：
	// * true：是
	// * false：否
	Enabled bool `json:"enabled"`

	// REQUIRED; 表示需要优化的对象列表。该参数有以下取值：
	// * html: （默认）表示 HTML 页面。
	// * js: 表示 Javascript 代码。
	// * css: 表示 CSS 代码。 :::tip 如果对象列表包含 js 或者 js，html 也必须被包含。 :::
	OptimizationType []string `json:"optimization_type"`
}

// UpdateImageDomainConfigBodyRefererLink - Referer 防盗链配置
type UpdateImageDomainConfigBodyRefererLink struct {

	// REQUIRED; 是否开启 Referer 防盗链，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"enabled"`

	// REQUIRED; 是否选择白名单，取值如下所示：
	// * true：配置白名单
	// * false：配置黑名单
	IsWhiteMode bool `json:"is_white_mode"`

	// 是否允许空 Referer，取值如下所示：
	// * true：允许空 Referer
	// * false：不允许空 Referer
	AllowEmptyRefer bool `json:"allow_empty_refer,omitempty"`

	// 是否忽略大小写。取值如下所示：
	// * true: （默认）大小写不敏感。
	// * false: 大小写敏感。
	IgnoreCase bool `json:"ignore_case,omitempty"`

	// Referer 头部值是否必须以 HTTP 或者 HTTPS 开头。取值如下所示：
	// * true: 表示不以 HTTP 或者 HTTPS 开头的 Referer 头部值是合法的。在这个情况下，veImagex 会尝试将其与 Referers 列表匹配。
	// * false: （默认）表示不以 HTTP 或者 HTTPS 开头 Referer 头部值是非法的。在这个情况下，veImagex 判定为不匹配 CommonType 下的这个 Referers 列表。
	IgnoreScheme bool `json:"ignore_scheme,omitempty"`

	// 请提供具体的参数名字和类型。
	RegexValues []string `json:"regex_values,omitempty"`

	// 请提供具体的参数名字（values）和参数类型（array），以便我为您生成参数描述。
	Values []string `json:"values,omitempty"`
}

// UpdateImageDomainConfigBodyRemoteAuth - 远程鉴权设置
type UpdateImageDomainConfigBodyRemoteAuth struct {

	// REQUIRED; 鉴权请求头设置
	AuthRequestHeader *UpdateImageDomainConfigBodyRemoteAuthRequestHeader `json:"auth_request_header"`

	// REQUIRED; 鉴权请求参数设置
	AuthRequestQuery *UpdateImageDomainConfigBodyRemoteAuthRequestQuery `json:"auth_request_query"`

	// REQUIRED; 鉴权响应设置
	AuthResponse *UpdateImageDomainConfigBodyRemoteAuthResponse `json:"auth_response"`

	// REQUIRED; 鉴权服务器设置
	AuthServer *UpdateImageDomainConfigBodyRemoteAuthServer `json:"auth_server"`

	// REQUIRED; 是否开启远程鉴权，取值如下所示：
	// * true：是
	// * false：否
	Enabled bool `json:"enabled"`

	// REQUIRED; 生效对象
	MatchRule []*UpdateImageDomainConfigBodyRemoteAuthMatchRuleItem `json:"match_rule"`
}

// UpdateImageDomainConfigBodyRemoteAuthAuthResponseResponse - 响应设置
type UpdateImageDomainConfigBodyRemoteAuthAuthResponseResponse struct {

	// REQUIRED; 鉴权失败时 veImageX 响应用户的状态码。取值范围为 [400,499] 。默认值是 403。
	FailCode string `json:"fail_code"`
}

type UpdateImageDomainConfigBodyRemoteAuthMatchRuleItem struct {

	// REQUIRED; 匹配方式，取值如下所示：
	// * match：（默认）表示 object 匹配 Value。
	// * not_match：表示 object 不匹配 Value。 :::tip 如果您创建了多个生效对象配置，每个配置中该参数的值必须相同。 :::
	MatchOperator string `json:"match_operator"`

	// REQUIRED; 表示 veImageX 对哪些对象类型进行规则匹配。取值如下所示：
	// * filetype：表示特定后缀的文件。
	// * directory：表示特定文件目录下的所有文件。
	// * path：表示特定的文件。
	Object string `json:"object"`

	// REQUIRED; 表示 Object 对应的具体对象，并且是大小写敏感的。参数值的长度不能超过 1,024 个字符。您可以指定一个或者多个对象。多个对象之间使用英文分号（;）分隔。该参数有以下说明：
	// * 如果 Object 是 filetype，您需要指定一个或者多个文件后缀。文件后缀可以包含英文字母和数字。多个文件后缀使用分号（;）分隔。例如 xlsx 或者 png;txt。
	// * 如果 Object 是 directory，您需要指定一个或者多个目录路径。多个目录路径使用分号（;）分隔。每个目录路径必须以斜杠（/）开头和结尾， :::tip 例如 /www/img/volc/;/www/doc/。您可以使用 /
	// 表示域名下的所有目录。同时，目录路径可以包含除了以下字符的可打印 ASCII 字符：
	// 连续斜杠（//）、百分号（%）、美元符号（$）、空格、问号（?）、Delete（ASCII code 127） :::
	// * 如果 Object 是 path，您需要指定一个或者多个文件路径。文件路径支持使用通配符（*）表示一个或者多个字符。多个文件路径使用分号（;）分隔。 :::tip 例如 /www/img/volcano.png;/doc/study.docx。文件路径必须以
	// / 开头。同时，文件路径可以包含除了以下字符的可打印 ASCII 字符：
	// 连续斜杠（//）、百分号（%）、美元符号（$）、空格、问号（?）、Delete（ASCII code 127） :::
	Value string `json:"value"`
}

// UpdateImageDomainConfigBodyRemoteAuthRequestHeader - 鉴权请求头设置
type UpdateImageDomainConfigBodyRemoteAuthRequestHeader struct {

	// REQUIRED; 鉴权请求头是否包含用户请求头。取值如下所示：
	// * exclude：表示鉴权请求头中不包含任何用户请求头。
	// * include：表示鉴权请求头中包含所有用户请求头。
	// * includePart：表示鉴权请求头包含指定的用户请求头。
	Action string `json:"action"`

	// REQUIRED; 表示鉴权请求中额外的请求头设置。您最多可以设置 50 个请求头。
	Header []*UpdateImageDomainConfigBodyRemoteAuthRequestHeaderItem `json:"header"`

	// REQUIRED; 鉴权请求中 HOST 头部的值。该参数的默认值是 default，表示 HOST 头部的值与您的加速域名相同。
	Host string `json:"host"`

	// REQUIRED; 表示 Action 参数所对应的参数值，长度不能超过 1,024 个字符。该参数有以下说明：
	// * 如果 Action 是 exclude 或 include，Value 必须是 *。
	// * 如果 Action 是 includePart，Value 参数的取值是用户请求中的一个或者多个头部。多个头部使用英文分号（;）分隔。其取值不能只是 *，可以包含除了以下字符的可打印 ASCII 字符： 下划线（_）、空格、双引号（"），Delete（ASCII
	// code 127） 该参数的默认值是 *。
	Value string `json:"value"`
}

type UpdateImageDomainConfigBodyRemoteAuthRequestHeaderItem struct {

	// REQUIRED; 您需要设置的请求头。请求头不能是 host，长度不能超过 1,024 个字符，并且不区分大小写。请求头可以包含除了以下字符的可打印 ASCII 字符： 下划线（_）、双引号（"）、空格、Delete（ASCII code
	// 127）
	Key string `json:"key"`

	// REQUIRED; 表示请求头的值。取值如下所示：
	// * 当 ValueType 是 constant 时，您需要指定一个常量值。该常量值的长度不能超过 1,024 个字符，并且区分大小写。同时，该常量值不能以美元符号（$）开头，可以包含除了以下字符的可打印 ASCII 字符： 双引号（"）、Delete（ASCII
	// code 127）
	// * 当 ValueType 是 variable 时，表示请求头的值来自一个变量。您可以指定该变量列表中的变量。
	Value string `json:"value"`

	// REQUIRED; 请求头的类型。取值如下所示：
	// * constant：表示请求头的值是一个常量。您需要在 Value 参数中指定该常量的值。
	// * variable：表示请求头的值来自一个变量。参见 Value 的说明。
	ValueType string `json:"value_type"`
}

// UpdateImageDomainConfigBodyRemoteAuthRequestQuery - 鉴权请求参数设置
type UpdateImageDomainConfigBodyRemoteAuthRequestQuery struct {

	// REQUIRED; 表示鉴权请求是否包含用户请求 URL 中的查询参数。取值如下所示：
	// * exclude：表示鉴权请求不包含任何查询参数。
	// * include：表示鉴权请求包含所有查询参数。
	// * includePart：表示鉴权请求包含指定的查询参数。
	Action string `json:"action"`

	// REQUIRED; 表示鉴权请求中额外的参数设置。您最多可以设置 50 个参数。
	Query []*UpdateImageDomainConfigBodyRemoteAuthRequestQueryItem `json:"query"`

	// REQUIRED; 表示 Action 参数所对应的参数值，长度不能超过1,024 个字符。该参数有以下说明：
	// * 如果 Action 是 exclude 或 include，Value 必须是 *。
	// * 如果 Action 是 includePart，您需要在 Value 参数中指定用户请求 URL 中的一个或者多个查询参数，多个查询参数使用英文分号（;）分隔。您不能指定 *。查询参数是区分大小写的，可以包含除了以下字符的可打印 ASCII
	// 字符： 双引号（"）、空格、Delete（ASCII code 127） 该参数的默认值是 *。
	Value string `json:"value"`
}

type UpdateImageDomainConfigBodyRemoteAuthRequestQueryItem struct {

	// REQUIRED; 您需要设置的鉴权请求参数，长度不能超过 1,024 个字符。鉴权请求参数可以包含除了以下字符的可打印 ASCII 字符： 双引号（"）、空格、Delete（ASCII code 127）
	Key string `json:"key"`

	// REQUIRED; 鉴权请求参数的值，长度不能超过 1,024 个字符，并且区分大小写。Value 有以下取值：
	// * 当 ValueType 是 constant 时，表示鉴权请求参数的值是一个常量。您需要指定该常量值。常量值不能以美元符号（$）开头，可以包含除了以下字符的可打印 ASCII 字符： 双引号（"）、Delete（ASCII code
	// 127）
	// * 当 ValueType 是 variable 时，表示鉴权请求参数的值来自一个变量。您可以指定该变量列表中的变量。
	Value string `json:"value"`

	// REQUIRED; 您在 Key 中设置的鉴权请求参数的类型。ValueType 有以下取值：
	// * constant：表示鉴权请求参数是一个常量。此时，您需要在 Value 中指定该常量的值。
	// * variable：表示鉴权请求参数的值来自一个变量。参见 Value 的说明。
	ValueType string `json:"value_type"`
}

// UpdateImageDomainConfigBodyRemoteAuthResponse - 鉴权响应设置
type UpdateImageDomainConfigBodyRemoteAuthResponse struct {

	// REQUIRED; 鉴权结果缓存设置
	AuthResultCache *UpdateImageDomainConfigBodyRemoteAuthResponseAuthResultCache `json:"auth_result_cache"`

	// REQUIRED; 鉴权服务器状态码设置
	AuthServerStatusCode *UpdateImageDomainConfigBodyRemoteAuthResponseAuthServerStatusCode `json:"auth_server_status_code"`

	// REQUIRED; 鉴权服务超时时间
	AuthServerTimeout *UpdateImageDomainConfigBodyRemoteAuthResponseAuthServerTimeout `json:"auth_server_timeout"`

	// REQUIRED; 响应设置
	Response *UpdateImageDomainConfigBodyRemoteAuthAuthResponseResponse `json:"response"`
}

// UpdateImageDomainConfigBodyRemoteAuthResponseAuthResultCache - 鉴权结果缓存设置
type UpdateImageDomainConfigBodyRemoteAuthResponseAuthResultCache struct {

	// REQUIRED; veImageX 是否缓存鉴权状态码。取值如下所示：
	// * nocache：veImageX 不缓存鉴权状态码。
	// * cache：veImageX 缓存鉴权状态码。
	Action string `json:"action"`

	// REQUIRED; 缓存 key 指定了用于区分不同请求 URI 的查询参数。可以指定变量字段说明 [https://www.volcengine.com/docs/508/1171078]中的参数, 必须包含 URI。
	CacheKey []string `json:"cache_key"`

	// REQUIRED; 鉴权状态码的缓存时间。单位是秒。取值范围是 [1,86400]。86400 秒表示 24 小时。
	TTL int `json:"ttl"`
}

// UpdateImageDomainConfigBodyRemoteAuthResponseAuthServerStatusCode - 鉴权服务器状态码设置
type UpdateImageDomainConfigBodyRemoteAuthResponseAuthServerStatusCode struct {

	// REQUIRED; 如果鉴权状态码既不是 FailCode，又不是 SuccessCode 时，veImageX 处理鉴权请求的方式。取值如下所示：
	// * reject：veImageX 认为鉴权失败。
	// * pass：veImageX 认为鉴权成功。
	DefaultAction string `json:"default_action"`

	// REQUIRED; 指定鉴权失败时的鉴权状态码。默认值是 401。
	// * 您可以指定范围在 400-499 中的一个或者多个状态码。多个状态码使用英文分号（;）分隔。
	// * 您也可以指定 4xx 表示 400-499 中的任意一个状态码。
	FailCode string `json:"fail_code"`

	// REQUIRED; 指定鉴权成功时的鉴权状态码。默认值是 200。
	// * 您可以指定范围在 200-299 中的一个或者多个状态码。多个状态码使用英文分号（;）分隔。
	// * 您也可以指定 2xx 表示 200-299 中的任意一个状态码。
	SuccessCode string `json:"success_code"`
}

// UpdateImageDomainConfigBodyRemoteAuthResponseAuthServerTimeout - 鉴权服务超时时间
type UpdateImageDomainConfigBodyRemoteAuthResponseAuthServerTimeout struct {

	// REQUIRED; 鉴权超时后 veImageX 处理鉴权请求的策略。取值如下所示：
	// * reject：veImageX 认为鉴权失败。
	// * pass：veImageX 认为鉴权成功。
	Action string `json:"action"`

	// REQUIRED; 鉴权超时的时间，单位是毫秒。默认值为 200，取值范围是 [200,3600]。
	Time int `json:"time"`
}

// UpdateImageDomainConfigBodyRemoteAuthServer - 鉴权服务器设置
type UpdateImageDomainConfigBodyRemoteAuthServer struct {

	// REQUIRED; 鉴权服务器的主地址。主地址的格式是 \://\ 或 \://\。该参数值的长度不能超过 100 个字符。
	// * \ 的值是 http 或者 https。
	// * \ 的值不能是 localhost。
	// * \ 的值不能是 127.0.0.1。
	Address string `json:"address"`

	// REQUIRED; 鉴权服务器的备地址。地址格式和要求与主地址 address 相同。
	BackupAddress string `json:"backup_address"`

	// REQUIRED; 鉴权请求的路径。鉴权地址和请求路径组成了完整的鉴权 URL。veImageX 会把用户的请求转发到该鉴权 URL。取值如下所示：
	// * constant：表示鉴权请求中的路径与用户请求中的路径相同。
	// * variable：表示您需要在 pathValue 参数中指定一个鉴权请求中的路径。
	PathType string `json:"path_type"`

	// REQUIRED; 表示一个鉴权请求的路径，长度不能超过 100 个字符。路径必须以斜杠（/）开头，可以包含除了以下字符的可打印 ASCII 字符： 连续斜杠（//）、百分号（%）、美元符号（$）、空格、问号（?）、Delete（ASCII
	// code 127）
	PathValue string `json:"path_value"`

	// REQUIRED; 在发送鉴权请求时，veImageX 所使用的请求方法。取值如下所示：
	// * default：鉴权请求所使用的方法与用户的请求相同。
	// * get：鉴权请求使用 GET 方法。
	// * post：鉴权请求使用 POST 方法。
	// * head：鉴权请求使用 HEAD 方法。
	RequestMethod string `json:"request_method"`
}

type UpdateImageDomainConfigBodyRespHdrsItem struct {

	// REQUIRED; Header Key，请见支持配置的响应头 [https://www.volcengine.com/docs/508/196704#%E6%94%AF%E6%8C%81%E9%85%8D%E7%BD%AE%E7%9A%84%E5%93%8D%E5%BA%94%E5%A4%B4]。
	Key string `json:"key"`

	// REQUIRED; Header Value，设置该响应头字段的值。字段值不能超过 1,024 个字符，可以包含除美元符号（$），Delete（ASCII code 127）外的可打印 ASCII 字符。
	Value string `json:"value"`
}

// UpdateImageDomainConfigBodyShareCache - 共享缓存配置。共享缓存允许同账号下多个加速域名共享同一份节点上的缓存。在 veImageX 中，您可以通过设置共享缓存配置，使各个子站点之间可共享相同的公共资源，以减少带宽的使用，提高资源命中率。详细功能说明参看共享缓存
// [https://www.volcengine.com/docs/508/196769]。
// :::tip 共享缓存为白名单功能，请提交工单 [https://console.volcengine.com/workorder/create?step=2&SubProductID=P00000080]联系技术支持为您的账号开启配置能力。
// :::
type UpdateImageDomainConfigBodyShareCache struct {

	// REQUIRED; 共享域名。
	Domains []string `json:"domains"`
}

// UpdateImageDomainConfigBodyURLAuth - URL 鉴权配置
type UpdateImageDomainConfigBodyURLAuth struct {

	// REQUIRED; 是否开启 URL 鉴权配置，取值如下所示：
	// * true：是
	// * false：否
	Enabled bool `json:"enabled"`

	// REQUIRED; A 鉴权配置
	TypeA *UpdateImageDomainConfigBodyURLAuthTypeA `json:"type_a"`

	// REQUIRED; B 鉴权配置
	TypeB *UpdateImageDomainConfigBodyURLAuthTypeB `json:"type_b"`

	// REQUIRED; C 鉴权配置
	TypeC *UpdateImageDomainConfigBodyURLAuthTypeC `json:"type_c"`

	// REQUIRED; D 鉴权配置
	TypeD *UpdateImageDomainConfigBodyURLAuthTypeD `json:"type_d"`
}

// UpdateImageDomainConfigBodyURLAuthTypeA - A 鉴权配置
type UpdateImageDomainConfigBodyURLAuthTypeA struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`

	// REQUIRED; md5hash 参数名
	SignParam string `json:"sign_param"`
}

// UpdateImageDomainConfigBodyURLAuthTypeB - B 鉴权配置
type UpdateImageDomainConfigBodyURLAuthTypeB struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`
}

// UpdateImageDomainConfigBodyURLAuthTypeC - C 鉴权配置
type UpdateImageDomainConfigBodyURLAuthTypeC struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`
}

// UpdateImageDomainConfigBodyURLAuthTypeD - D 鉴权配置
type UpdateImageDomainConfigBodyURLAuthTypeD struct {

	// REQUIRED; 备用鉴权密钥
	BackupSk string `json:"backup_sk"`

	// REQUIRED; 有效时间，单位为秒。取值范围为[1, 630720000]内的正整数，默认为 1800 秒。
	ExpireTime int `json:"expire_time"`

	// REQUIRED; 主鉴权密钥
	MainSk string `json:"main_sk"`

	// REQUIRED; md5hash 参数名
	SignParam string `json:"sign_param"`

	// REQUIRED; 时间戳格式，取值如下所示：
	// * decimal：十进制（Unix 时间戳）
	// * heximal：十六进制（Unix 时间戳）
	TimeFormat string `json:"time_format"`

	// REQUIRED; TimeStamp 参数名
	TimeParam string `json:"time_param"`
}

// UpdateImageDomainConfigBodyUserAgentACL - UA 访问限制配置
type UpdateImageDomainConfigBodyUserAgentACL struct {

	// REQUIRED; 表示是否允许 UA 为空或者不包含 UA 字段的请求访问加速域名。取值如下所示：
	// * true：允许
	// * false：不允许
	AllowEmpty bool `json:"allow_empty"`

	// REQUIRED; 是否开启 UA 访问限制，取值如下所示：
	// * true：开启
	// * false：未开启
	Enabled bool `json:"enabled"`

	// REQUIRED; deny黑名单，allow白名单
	RuleType string `json:"rule_type"`

	// REQUIRED; Agent 列表，最多可支持输入 1000 个，支持通配符 * 匹配任意字符串。
	UserAgents []string `json:"user_agents"`
}

type UpdateImageDomainConfigQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageDomainConfigRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageDomainConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageDomainConfigResResponseMetadata struct {

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

type UpdateImageDomainDownloadSpeedLimitBody struct {

	// REQUIRED; 域名。您可以通过调用 GetServiceDomains [https://www.volcengine.com/docs/508/9379]接口获取域名。
	Domain string `json:"domain"`

	// REQUIRED; 下载限速配置
	DownloadSpeedLimit *UpdateImageDomainDownloadSpeedLimitBodyDownloadSpeedLimit `json:"download_speed_limit"`
}

// UpdateImageDomainDownloadSpeedLimitBodyDownloadSpeedLimit - 下载限速配置
type UpdateImageDomainDownloadSpeedLimitBodyDownloadSpeedLimit struct {

	// REQUIRED; 是否启用。
	Enabled bool `json:"enabled"`

	// REQUIRED; 参数规则。
	Rules []*UpdateImageDomainDownloadSpeedLimitBodyDownloadSpeedLimitRulesItem `json:"rules"`
}

type UpdateImageDomainDownloadSpeedLimitBodyDownloadSpeedLimitRulesItem struct {

	// REQUIRED; 开始时间。
	BeginTime string `json:"begin_time"`

	// REQUIRED; 星期几。
	DayWeek string `json:"day_week"`

	// REQUIRED; 结束时间。
	EndTime string `json:"end_time"`

	// REQUIRED; 限速配置。
	LimitRate int `json:"limit_rate"`

	// REQUIRED; 限制速率的起始点。
	LimitRateAfter int `json:"limit_rate_after"`

	// REQUIRED; 匹配类型。
	MatchType string `json:"match_type"`

	// REQUIRED; 匹配值。
	MatchValue string `json:"match_value"`
}

type UpdateImageDomainDownloadSpeedLimitQuery struct {

	// REQUIRED; 服务 ID。您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。您也可以通过调用GetAllImageServices
	// [https://www.volcengine.com/docs/508/9360]接口方式获取服务 ID。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageDomainDownloadSpeedLimitRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageDomainDownloadSpeedLimitResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *UpdateImageDomainDownloadSpeedLimitResResult `json:"Result,omitempty"`
}

type UpdateImageDomainDownloadSpeedLimitResResponseMetadata struct {

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

// UpdateImageDomainDownloadSpeedLimitResResult - 视请求的接口而定
type UpdateImageDomainDownloadSpeedLimitResResult struct {

	// REQUIRED; 通知描述
	Msg string `json:"msg"`
}

type UpdateImageDomainIPAuthBody struct {

	// REQUIRED; 待修改配置的域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"domain"`

	// REQUIRED; 黑白名单配置
	IPAuth *UpdateImageDomainIPAuthBodyIPAuth `json:"ip_auth"`
}

// UpdateImageDomainIPAuthBodyIPAuth - 黑白名单配置
type UpdateImageDomainIPAuthBodyIPAuth struct {

	// REQUIRED; 是否开启黑白名单配置，取值如下所示：
	// * true：开启黑白名单配置
	// * false：关闭黑白名单配置
	Enabled bool `json:"enabled"`

	// REQUIRED; 是否是 IP 白名单，取值如下所示：
	// * true：配置白名单
	// * false：配置黑名单
	IsWhiteMode bool `json:"is_white_mode"`

	// REQUIRED; 黑白名单 IP 地址，您可以指定一个或者多个 IP 地址（如 192.0.2.0）和 IP 地址网段（192.0.2.0/24）。IP 地址和网段可以是 IPv4 或 IPv6 格式，可混合填写，最多可输入 1000
	// 个地址。
	// :::warning 若您需要对同类型名单内已设定的 values 地址进行增删处理，那么您可调用 获取域名配置 [https://www.volcengine.com/docs/508/9366#ip-auth] 接口获取已配置的全部地址列表后，在此基础上添加或删除您期望变更的地址，最后重新传入
	// values。 :::
	Values []string `json:"values"`
}

type UpdateImageDomainIPAuthQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageDomainIPAuthRes struct {
	ResponseMetadata *UpdateImageDomainIPAuthResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                                      `json:"Result,omitempty"`
}

type UpdateImageDomainIPAuthResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateImageDomainUaAccessBody struct {

	// REQUIRED; 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"domain"`

	// REQUIRED; UA 访问限制配置。
	UaAuth *UpdateImageDomainUaAccessBodyUaAuth `json:"ua_auth"`
}

// UpdateImageDomainUaAccessBodyUaAuth - UA 访问限制配置。
type UpdateImageDomainUaAccessBodyUaAuth struct {

	// REQUIRED; 是否允许 UA 为空或者不包含 UA 字段的请求访问加速域名。取值如下所示：
	// * true：允许
	// * false：不允许
	AllowEmpty bool `json:"allow_empty"`

	// REQUIRED; 是否开启 UA 访问限制，取值如下所示：
	// * true：开启
	// * false：未开启（默认）
	Enable bool `json:"enable"`

	// REQUIRED; deny黑名单，allow白名单
	RuleType string `json:"rule_type"`

	// REQUIRED; Agent 列表，最多可支持输入 1000 个，支持通配符 * 匹配任意字符串。输入多个时以 | 分割，或者一行仅输入一个。
	// :::warning 若您需要对同类型名单内已设定的 Agent 列表进行增删处理，那么您可调用获取域名配置 [https://www.volcengine.com/docs/508/9366#ua-list]接口获取已配置的全部列表后，在此基础上添加或删除您期望变更的值，最后重新传入user_agents。
	// :::
	UserAgents []string `json:"user_agents"`
}

type UpdateImageDomainUaAccessQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageDomainUaAccessRes struct {
	ResponseMetadata *UpdateImageDomainUaAccessResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                                        `json:"Result,omitempty"`
}

type UpdateImageDomainUaAccessResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateImageDomainVolcOriginBody struct {

	// REQUIRED; 域名
	Domain string `json:"domain"`

	// REQUIRED
	OriginConfig *UpdateImageDomainVolcOriginBodyOriginConfig `json:"origin_config"`

	// REQUIRED; 回源host
	OriginHost string `json:"origin_host"`

	// REQUIRED; 回源协议配置，存在以下三种情况： http：用户侧发起 HTTP 及 HTTPS 请求均会使用 HTTP 回源； https：用户侧发起的 HTTP 及 HTTPS 请求均会使用 HTTPS 回源； followclient：用户侧发起的
	// HTTP 请求使用 HTTP 回源，发起的 HTTPS 请求使用 HTTPS 回源。
	OriginProtocol string `json:"origin_protocol"`

	// REQUIRED; 分片回源
	OriginRange bool `json:"origin_range"`

	// REQUIRED; 是否使用ImageX源站
	UseImagex bool `json:"use_imagex"`
}

type UpdateImageDomainVolcOriginBodyOriginConfig struct {

	// REQUIRED
	Origins []*UpdateImageDomainVolcOriginBodyOriginConfigOriginsItem `json:"origins"`
}

type UpdateImageDomainVolcOriginBodyOriginConfigOriginsItem struct {

	// REQUIRED; 与 instancetype 填充的内容对应： instancetype 为 ip ，则仅可填充一条 IPv4 记录； instance_type 为 domain，则可填充一个域名，域名长度不超过 1024 个字符。
	Address string `json:"address"`

	// REQUIRED; HTTP 请求回源至对应 Address 的端口
	HTTPPort string `json:"http_port"`

	// REQUIRED; HTTPS 请求回源至对应 Address 的端口，修改时需要指定，取值范围为 1 ~ 65535
	HTTPSPort string `json:"https_port"`

	// REQUIRED; ip：IP 类型源站； domain：域名类型
	InstanceType string `json:"instance_type"`

	// REQUIRED; 指定源站的Address维度的回源 Host 的值; 若不为空，则优先级高于域名维度的OriginHost。 若为空，则遵循域名维度的OriginHost
	OriginHost string `json:"origin_host"`

	// REQUIRED; primary：主源站； backup：备源站
	OriginType string `json:"origin_type"`

	// REQUIRED; 多源站配置场景下，权重决定了回源至对应源站的概率。 InstanceType 为 ip 时，指定当前 Address 对应的权重，取值范围为 1~1000。 InstanceType 为 domain 或 tos 时，权重默认为
	// 1。
	Weight string `json:"weight"`
}

type UpdateImageDomainVolcOriginQuery struct {

	// REQUIRED; 服务ID
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageDomainVolcOriginRes struct {
	ResponseMetadata *UpdateImageDomainVolcOriginResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                                          `json:"Result,omitempty"`
}

type UpdateImageDomainVolcOriginResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateImageExifDataBody struct {

	// REQUIRED; 修改操作
	Actions []*UpdateImageExifDataBodyActionsItem `json:"Actions"`

	// REQUIRED; 原图URI
	StoreURI string `json:"StoreUri"`

	// 目标Key
	DstKey string `json:"DstKey,omitempty"`
}

type UpdateImageExifDataBodyActionsItem struct {

	// REQUIRED; 修改类型
	Type string `json:"Type"`

	// Tag名
	TagName string `json:"TagName,omitempty"`

	// Tag值
	TagValue string `json:"TagValue,omitempty"`
}

type UpdateImageExifDataQuery struct {

	// REQUIRED; 待修改图片所在的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageExifDataRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageExifDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result *UpdateImageExifDataResResult `json:"Result"`
}

type UpdateImageExifDataResResponseMetadata struct {

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

// UpdateImageExifDataResResult - 视请求的接口而定
type UpdateImageExifDataResResult struct {

	// REQUIRED; 存储URI
	DstURI string `json:"DstUri"`
}

type UpdateImageFileCTBody struct {

	// REQUIRED; 设置更新后的 Content-Type 值。 :::tip 请确保更新后的 Content-Type，在服务维度设置的 Content-Type 白名单内。 :::
	StorageContentType string `json:"StorageContentType"`

	// REQUIRED; 待更新文件的存储 URI，您可以通过调用获取服务下的上传文件 [https://www.volcengine.com/docs/508/9392]来获取所需的文件 URI。
	StoreURI string `json:"StoreUri"`
}

type UpdateImageFileCTQuery struct {

	// REQUIRED; 待更新文件所在的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageFileCTRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageFileCTResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageFileCTResResponseMetadata struct {

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

type UpdateImageFileKeyBody struct {

	// REQUIRED; 重命名后的文件名。输入限制如下所示：
	// * 不支持空格，如果中间有空格将会导致重命名失败。
	// * 不支持以/开头或结尾，不支持/连续出现，Key 值最大长度限制为 180 个字节。
	DstKey string `json:"DstKey"`

	// REQUIRED; 源文件名，即上传文件的 storekey。
	OriKey string `json:"OriKey"`
}

type UpdateImageFileKeyQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageFileKeyRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageFileKeyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageFileKeyResResponseMetadata struct {

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

type UpdateImageMirrorConfBody struct {

	// REQUIRED; 镜像回源下载原图时，携带的 HTTP 头部，键值都为 String 类型。
	Headers map[string]string `json:"Headers"`

	// REQUIRED; 镜像回源域名。
	Host string `json:"Host"`

	// REQUIRED; 下载图片的协议，取值为：http、https。
	Schema string `json:"Schema"`

	// 带权重回源域名。key 为 String 类型，代表镜像回源域名；value 为 Integer 类型，代表域名权重。
	Hosts map[string]int `json:"Hosts,omitempty"`

	// 镜像站
	Origin *UpdateImageMirrorConfBodyOrigin `json:"Origin,omitempty"`

	// 镜像源 URI，其中图片名用 %s 占位符替代，比如/obj/%s。
	Source string `json:"Source,omitempty"`
}

// UpdateImageMirrorConfBodyOrigin - 镜像站
type UpdateImageMirrorConfBodyOrigin struct {

	// REQUIRED; 镜像站配置
	Param map[string]interface{} `json:"Param"`

	// REQUIRED; 镜像站类型
	Type string `json:"Type"`
}

type UpdateImageMirrorConfQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageMirrorConfRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageMirrorConfResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageMirrorConfResResponseMetadata struct {

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

type UpdateImageMonitorRuleBody struct {

	// REQUIRED; 更新后的报警规则，具体请见MonitorRule [https://www.volcengine.com/docs/508/1112183#monitorrule]。
	MonitorRule *UpdateImageMonitorRuleBodyMonitorRule `json:"MonitorRule"`
}

// UpdateImageMonitorRuleBodyMonitorRule - 更新后的报警规则，具体请见MonitorRule [https://www.volcengine.com/docs/508/1112183#monitorrule]。
type UpdateImageMonitorRuleBodyMonitorRule struct {

	// REQUIRED; 监控的应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppID。
	Appid string `json:"Appid"`

	// REQUIRED; 监测规则。
	Cond *UpdateImageMonitorRuleBodyMonitorRuleCond `json:"Cond"`

	// REQUIRED; 创建后是否立即开启告警，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"Enabled"`

	// REQUIRED; 监控频率，单位为分钟。取值如下所示：
	// * 5
	// * 10
	// * 20
	// * 30
	// * 40
	// * 50
	Frequency int `json:"Frequency"`

	// REQUIRED; 告警级别，取值如下所示：
	// * warn：警告
	// * error：错误
	// * fatal：致命
	Level string `json:"Level"`

	// REQUIRED; 自定义告警规则名称
	Name string `json:"Name"`

	// REQUIRED; 告警通知配置。
	Notification *UpdateImageMonitorRuleBodyMonitorRuleNotification `json:"Notification"`

	// REQUIRED; 监控阶段，取值如下所示：
	// * upload：图片上传-上传 1.0
	// * uploadv2：图片上传-上传 2.0
	// * cdn：图片加载-下行网络监控
	// * client：图片加载-客户端传状态监控
	// * sensible：图片加载-感知指标监控
	Phase string `json:"Phase"`

	// REQUIRED; 待更新的报警规则 ID，您可以调用GetImageMonitorRules [https://www.volcengine.com/docs/508/1112186]获取所需的告警规则 ID。
	RuleID string `json:"RuleId"`

	// 维度过滤条件，具体参数请见Filter [https://www.volcengine.com/docs/508/1112182#filter]。用于指定需要告警提示的维度配置。
	Filter *UpdateImageMonitorRuleBodyMonitorRuleFilter `json:"Filter,omitempty"`

	// 拆分维度，由公共拆分维度 [https://www.volcengine.com/docs/508/1113944]和自定义拆分维度 [https://www.volcengine.com/docs/508/34554]组合而成。
	GroupBy string `json:"GroupBy,omitempty"`

	// 监控平台，取值如下所示：
	// * iOS
	// * Android
	// * WEB
	OS string `json:"OS,omitempty"`
}

// UpdateImageMonitorRuleBodyMonitorRuleCond - 监测规则。
type UpdateImageMonitorRuleBodyMonitorRuleCond struct {

	// REQUIRED
	ItemCond []*UpdateImageMonitorRuleBodyMonitorRuleCondItem `json:"ItemCond"`

	// REQUIRED; 多条监控规则之间的逻辑关系，取值如下所示：
	// * and：且。表示有多条监控规则时，需满足所有监控规则才会触发告警通知。
	// * or：或。表示有多条监控规则时，满足其中一条监控规则就会触发告警通知。
	LogicOp string `json:"LogicOp"`
}

// UpdateImageMonitorRuleBodyMonitorRuleCondItem - 监控规则配置
type UpdateImageMonitorRuleBodyMonitorRuleCondItem struct {

	// REQUIRED; 聚合周期，单位为分钟。被监控指标在该指定周期内满足指标比较阈值且上报量满足样本量阈值时，才会触发告警。取值如下所示：
	// * 5
	// * 10
	AggrInterval int `json:"AggrInterval"`

	// REQUIRED; 指标取值函数，取值如下所示：
	// * max：最大值
	// * min：最小值
	// * avg：平均值
	// * pct25：25峰值
	// * pct50：50峰值
	// * pct90：90峰值
	// * pct99：99峰值
	// * sum：总和
	// :::tip 各指标支持的函数参考veImageX 告警指标定义 [https://www.volcengine.com/docs/508/1113944]。 :::
	Func string `json:"Func"`

	// REQUIRED; 指标名称，取值参考veImageX 告警指标定义 [https://www.volcengine.com/docs/508/1113944]。
	Item string `json:"Item"`

	// REQUIRED; 指标比较方法，取值如下所示：
	// * LE：小于等于
	// * GE：大于等于
	// * INC：环比上升
	// * DEC：环比下降
	// * HOH_INC：与上小时同比上升
	// * HOH_DEC：与上小时同比下降
	// * DOD_INC：与昨天同比上升
	// * DOD_DEC：与昨天同比下降
	Op string `json:"Op"`

	// REQUIRED; 持续周期，当监控指标在聚合周期内，连续RepeatCnt次满足指标比较阈值且上报量满足样本量阈值时，才会触发告警。取值如下所示：
	// * 1
	// * 3
	// * 5
	RepeatCnt int `json:"RepeatCnt"`

	// REQUIRED; 指标比较阈值，需要与CntThreshold同时被满足才会触发告警。
	Threshold float64 `json:"Threshold"`

	// 样本量阈值。被监控指标超过该值时触发告警。
	CntThreshold int `json:"CntThreshold,omitempty"`
}

// UpdateImageMonitorRuleBodyMonitorRuleFilter - 维度过滤条件，具体参数请见Filter [https://www.volcengine.com/docs/508/1112182#filter]。用于指定需要告警提示的维度配置。
type UpdateImageMonitorRuleBodyMonitorRuleFilter struct {

	// REQUIRED; 过滤条件
	DimFilter []*UpdateImageMonitorRuleBodyMonitorRuleFilterDimFilterItem `json:"DimFilter"`

	// REQUIRED; 过滤条件之间的逻辑关系，取值如下所示：
	// * and：和
	// * or：或
	LogicOp string `json:"LogicOp"`
}

type UpdateImageMonitorRuleBodyMonitorRuleFilterDimFilterItem struct {

	// REQUIRED; 维度名称，由公共过滤维度 [https://www.volcengine.com/docs/508/1113944]和自定义过滤维度 [https://www.volcengine.com/docs/508/34554]组合而成。
	Dim string `json:"Dim"`

	// REQUIRED; 维度取值，您可以通过调用获取自定义维度值 [https://www.volcengine.com/docs/508/34555]来获取。
	Vals []string `json:"Vals"`

	// 纬度值是否取反，取值如下所示：
	// * true：指定维度的实际值不得满足Vals所有指定值
	// * false：（默认）维度值等于Vals中之一即可
	Not bool `json:"Not,omitempty"`
}

// UpdateImageMonitorRuleBodyMonitorRuleNotification - 告警通知配置。
type UpdateImageMonitorRuleBodyMonitorRuleNotification struct {

	// REQUIRED; 通知内容模板，模板中变量格式为$Name$。Name 取值如下所示：
	// * 报警名称
	// * 报警级别
	// * 报警App
	// * 报警平台
	// * 报警时间
	// * 报警内容
	Content string `json:"Content"`

	// REQUIRED; 通知方式，仅支持取值http_callback，表示回调。
	Mode []string `json:"Mode"`

	// REQUIRED; 沉默周期，单位为分钟。告警发生后，若未恢复正常，则会间隔一个沉默周期后再次重复发送一次告警通知。取值如下所示：
	// * 0
	// * 30
	// * 60
	// * 360
	SilentDur int `json:"SilentDur"`

	// REQUIRED; 告警通知标题
	Title string `json:"Title"`

	// 回调地址，Mode包含http_callback时，为必填。
	CallbackURL string `json:"CallbackUrl,omitempty"`
}

type UpdateImageMonitorRuleRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageMonitorRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageMonitorRuleResResponseMetadata struct {

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

type UpdateImageMonitorRuleStatusBody struct {

	// REQUIRED; 是否开启告警监测，取值如下所示：
	// * true：开启
	// * false：不开启
	Enabled bool `json:"Enabled"`

	// REQUIRED; 待更新的告警规则 ID，您可以调用GetImageMonitorRules [https://www.volcengine.com/docs/508/1112186]获取所需的告警规则 ID。
	RuleID string `json:"RuleId"`
}

type UpdateImageMonitorRuleStatusRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageMonitorRuleStatusResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageMonitorRuleStatusResResponseMetadata struct {

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

type UpdateImageObjectAccessBody struct {

	// 是否开启源地址访问，取值如下所示：
	// * true：开启源地址访问
	// * false：关闭源地址访问
	ObjectAccess bool `json:"ObjectAccess,omitempty"`
}

type UpdateImageObjectAccessQuery struct {

	// REQUIRED; 待更新配置的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageObjectAccessRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageObjectAccessResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageObjectAccessResResponseMetadata struct {

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

type UpdateImageResourceStatusBody struct {

	// REQUIRED; 待修改封禁状态的资源存储 Key（不携带 Bucket 信息），可在控制台资源管理页面查看。
	ObjectName string `json:"ObjectName"`

	// REQUIRED; 资源的封禁状态，取值如下所示：
	// * disable：禁用。禁用状态，您无法访问该资源。
	// * enable：启用。启用状态，您可正常访问该资源。
	Status string `json:"Status"`
}

type UpdateImageResourceStatusQuery struct {

	// REQUIRED; 指定配置资源封禁的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageResourceStatusRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageResourceStatusResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageResourceStatusResResponseMetadata struct {

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

type UpdateImageSettingRuleBody struct {

	// REQUIRED; 应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppId。
	AppID string `json:"AppId"`

	// REQUIRED; 规则内容
	Rule *UpdateImageSettingRuleBodyRule `json:"Rule"`

	// REQUIRED; 待更新的规则 ID，您可以通过调用获取规则列表 [https://www.volcengine.com/docs/508/1324618]的方式获取所需的规则 ID。
	RuleID string `json:"RuleId"`

	// REQUIRED; 配置项 ID，您可以通过调用获取配置项列表 [https://www.volcengine.com/docs/508/1324617]的方式获取所需的配置项 ID。
	SettingID string `json:"SettingId"`
}

// UpdateImageSettingRuleBodyRule - 规则内容
type UpdateImageSettingRuleBodyRule struct {

	// REQUIRED; 规则名称，仅支持字母、数字、下划线，最多输入 32 个字符。
	Name string `json:"Name"`

	// REQUIRED; 由对应配置项的类型决定，此处类型是为了方便生成 SDK
	Value interface{} `json:"Value"`

	// 匹配条件，仅当条件匹配后规则才会生效。
	Cond *UpdateImageSettingRuleBodyRuleCond `json:"Cond,omitempty"`
}

// UpdateImageSettingRuleBodyRuleCond - 匹配条件，仅当条件匹配后规则才会生效。
type UpdateImageSettingRuleBodyRuleCond struct {

	// 规则条件
	Conds []*UpdateImageSettingRuleBodyRuleCondCondsItem `json:"Conds,omitempty"`

	// 匹配条件，取值如下所示：
	// * AND：表示与
	// * OR：表示或
	Type string `json:"Type,omitempty"`
}

type UpdateImageSettingRuleBodyRuleCondCondsItem struct {

	// 过滤维度，取值请参考规则配置条件 [https://www.volcengine.com/docs/508/65940#%E8%A7%84%E5%88%99%E9%85%8D%E7%BD%AE%E6%9D%A1%E4%BB%B6]。
	Key string `json:"Key,omitempty"`

	// 操作符。支持取值：==、!=、>、>=、<、<=、in
	Op string `json:"Op,omitempty"`

	// 类型由 Op 决定，此处是为了方便生成 SDK
	Value interface{} `json:"Value,omitempty"`
}

type UpdateImageSettingRulePriorityBody struct {

	// REQUIRED; 应用 ID，您可以通过调用获取应用列表 [https://www.volcengine.com/docs/508/19511]的方式获取所需的 AppId。
	AppID string `json:"AppId"`

	// REQUIRED; 更新后的优先级信息。
	Priorities []*UpdateImageSettingRulePriorityBodyPrioritiesItem `json:"Priorities"`

	// REQUIRED; 配置项 ID，您可以通过调用获取配置项列表 [https://www.volcengine.com/docs/508/1324617]的方式获取所需的配置项 ID。
	SettingID string `json:"SettingId"`
}

type UpdateImageSettingRulePriorityBodyPrioritiesItem struct {

	// REQUIRED; 规则优先级。 :::tip 如果配置项下创建了多个规则，需要填写全部规则更新后的优先级。 :::
	Priority int `json:"Priority"`

	// REQUIRED; 待更新优先级的规则 ID，您可以通过调用获取规则列表 [https://www.volcengine.com/docs/508/1324618]的方式获取所需的规则 ID。
	RuleID string `json:"RuleId"`
}

type UpdateImageSettingRulePriorityRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageSettingRulePriorityResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageSettingRulePriorityResResponseMetadata struct {

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

type UpdateImageSettingRuleRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageSettingRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageSettingRuleResResponseMetadata struct {

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

type UpdateImageStorageTTLBody struct {

	// REQUIRED; 服务 ID 。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// REQUIRED; 更新后的有效期时间。单位为秒，取值范围为 [86400, 7776000]。 若取值超过整天，则默认向上取整，例如当TTL取值为86401时，实际有效期被设置为 2 天。
	TTL int `json:"TTL"`
}

type UpdateImageStorageTTLRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageStorageTTLResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageStorageTTLResResponseMetadata struct {

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

type UpdateImageStyleBody struct {

	// REQUIRED; 更新的样式结构，包含图片编辑、文字编辑、背景等自定义参数配置，具体请参考样式结构 [https://www.volcengine.com/docs/508/127402]。
	Style interface{} `json:"Style"`

	// 是否执行对上传图像的样式渲染和渲染结果图上传操作，默认为``。取值如下所示：
	// * true：将所有已上传至该样式的图像以更新后的样式数据进行重新处理，并将结果图上传至该样式所绑定服务的存储中。其更新后的结果图 Uri 请在获取样式详情 [https://www.volcengine.com/docs/508/127401]中获取。
	// * false：（默认）不执行上述操作。 :::tip 建议您仅在手动保存样式或关闭当前页面时指定DoUpload为TRUE，可节省后端渲染成本。 :::
	DoUpload bool `json:"DoUpload,omitempty"`
}

type UpdateImageStyleMetaBody struct {

	// REQUIRED; 更新后的样式名称。
	Name string `json:"Name"`

	// REQUIRED; 待更新的样式 ID。
	StyleID string `json:"StyleId"`
}

type UpdateImageStyleMetaRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageStyleMetaResResponseMetadata `json:"ResponseMetadata"`

	// title
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageStyleMetaResResponseMetadata struct {

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

type UpdateImageStyleRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageStyleResResponseMetadata `json:"ResponseMetadata"`
	Result           *UpdateImageStyleResResult           `json:"Result,omitempty"`
}

type UpdateImageStyleResResponseMetadata struct {

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

type UpdateImageStyleResResult struct {

	// 保留字段
	BaseResp      *UpdateImageStyleResResultBaseResp `json:"BaseResp,omitempty"`
	CollectResult string                             `json:"CollectResult,omitempty"`
}

// UpdateImageStyleResResultBaseResp - 保留字段
type UpdateImageStyleResResultBaseResp struct {
	Extra         string `json:"Extra,omitempty"`
	StatusCode    int    `json:"StatusCode,omitempty"`
	StatusMessage string `json:"StatusMessage,omitempty"`
}

type UpdateImageTaskStrategyBody struct {

	// REQUIRED; 调整后的迁移策略
	RunStrategy *UpdateImageTaskStrategyBodyRunStrategy `json:"RunStrategy"`

	// REQUIRED; 任务ID，请参考CreateImageMigrateTask [https://www.volcengine.com/docs/508/1009929]获取返回的任务 ID。
	TaskID string `json:"TaskId"`
}

// UpdateImageTaskStrategyBodyRunStrategy - 调整后的迁移策略
type UpdateImageTaskStrategyBodyRunStrategy struct {

	// 源下载 QPS 限制。如取值不为空，则长度必须为 24，表示一天 24 小时内各小时的 QPS 限制值。默认无限制。
	// * 取值为负值时，表示无限制
	// * 取值为 0 时，表示对应时间不允许迁移
	ReadQPS []int `json:"ReadQps,omitempty"`

	// 源下载流量限制。单位为 Byte。如取值不为空，则长度必须为24，表示一天 24 小时内各小时的流量限制值。默认无限制。
	// * 取值为负值时，表示无限制
	// * 取值为 0 时，表示对应时间不允许迁移
	ReadRate []int `json:"ReadRate,omitempty"`
}

type UpdateImageTaskStrategyRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageTaskStrategyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageTaskStrategyResResponseMetadata struct {

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

type UpdateImageTranscodeQueueBody struct {

	// REQUIRED; 是否启用回调。取值如下所示：
	// * true：启用
	// * false：不启用
	EnableCallback bool `json:"EnableCallback"`

	// REQUIRED; 待更新的队列 ID，您可通过调用GetImageTranscodeQueues [https://www.volcengine.com/docs/508/1107341]获取该账号下全部任务队列 ID。
	ID string `json:"Id"`

	// REQUIRED; 更新后的队列名称
	Name string `json:"Name"`

	// 更新后的队列回调配置
	CallbackConf *UpdateImageTranscodeQueueBodyCallbackConf `json:"CallbackConf,omitempty"`

	// 更新后的队列描述
	Desc string `json:"Desc,omitempty"`
}

// UpdateImageTranscodeQueueBodyCallbackConf - 更新后的队列回调配置
type UpdateImageTranscodeQueueBodyCallbackConf struct {

	// REQUIRED; 回调 HTTP 请求地址，用于接收转码结果详情。支持使用 https 和 http 协议。
	Endpoint string `json:"Endpoint"`

	// REQUIRED; 回调方式。仅支持取值 HTTP。
	Method string `json:"Method"`

	// 业务自定义回调参数，将在回调消息的callback_args中透传出去。具体回调参数请参考回调内容 [https://www.volcengine.com/docs/508/1104726#%E5%9B%9E%E8%B0%83%E5%86%85%E5%AE%B9]。
	Args string `json:"Args,omitempty"`

	// 回调数据格式。取值如下所示：
	// * XML
	// * JSON（默认）
	DataFormat string `json:"DataFormat,omitempty"`
}

type UpdateImageTranscodeQueueRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageTranscodeQueueResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateImageTranscodeQueueResResponseMetadata struct {

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

type UpdateImageTranscodeQueueStatusBody struct {

	// REQUIRED; 待更新的队列 ID，您可通过调用GetImageTranscodeQueues [https://www.volcengine.com/docs/508/1107341]获取该账号下全部任务队列 ID。
	ID string `json:"Id"`

	// REQUIRED; 更新后的队列状态。取值如下所示：
	// * Pending：排队中
	// * Running：执行中
	Status string `json:"Status"`
}

type UpdateImageTranscodeQueueStatusRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageTranscodeQueueStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateImageTranscodeQueueStatusResResponseMetadata struct {

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

type UpdateImageUploadFilesBody struct {

	// REQUIRED; 更新操作。
	// * 取值为 0 表示刷新 URL
	// * 取值为 1 表示禁用 URL
	// * 取值为 2 表示解禁 URL
	// * 取值为 4 表示预热 URL
	// * 取值为 5 表示目录刷新
	// * 取值为 6 表示样式刷新
	Action int `json:"Action"`

	// REQUIRED; 待操作的图片 URL 列表，最多传 100 个。
	// * 当 [Action] 取值为 5 时，表示待刷新的目录列表，目录必须以 / 结尾，不支持刷新根目录，也不支持海外域名;
	// * 当 [Action] 取值为 6 时，表示样式刷新，一次性最多传入 5 个 url，单个 url 最多匹配10万个样式，暂不支持海外使用。
	ImageUrls []string `json:"ImageUrls"`
}

type UpdateImageUploadFilesQuery struct {

	// REQUIRED; 服务 ID。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageUploadFilesRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageUploadFilesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *UpdateImageUploadFilesResResult `json:"Result,omitempty"`
}

type UpdateImageUploadFilesResResponseMetadata struct {

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

// UpdateImageUploadFilesResResult - 视请求的接口而定
type UpdateImageUploadFilesResResult struct {

	// REQUIRED
	FailUrls []string `json:"FailUrls"`

	// REQUIRED; 操作成功的图片 URL 列表。
	ImageUrls []string `json:"ImageUrls"`
}

type UpdateImageUploadOverwriteBody struct {

	// REQUIRED; 是否开启重名覆盖上传，取值如下所示：
	// * true：开启
	// * false：关闭
	UploadOverwrite bool `json:"UploadOverwrite"`
}

type UpdateImageUploadOverwriteQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考GetAllImageServices [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateImageUploadOverwriteRes struct {

	// REQUIRED
	ResponseMetadata *UpdateImageUploadOverwriteResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateImageUploadOverwriteResResponseMetadata struct {

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

type UpdateReferBody struct {

	// REQUIRED; 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"domain"`

	// REQUIRED; Referer 配置
	ReferLink *UpdateReferBodyReferLink `json:"refer_link"`
}

// UpdateReferBodyReferLink - Referer 配置
type UpdateReferBodyReferLink struct {

	// 是否允许空 Referer 访问，取值如下所示：
	// * true：允许
	// * false：不允许
	AllowEmptyRefer bool `json:"allow_empty_refer,omitempty"`

	// 是否开启 Referer 访问限制，取值如下所示：
	// * true：开启
	// * false：关闭
	Enabled bool `json:"enabled,omitempty"`

	// 是否选择白名单模式，取值如下所示：
	// * true：选择白名单
	// * false：选择黑名单
	IsWhiteMode bool `json:"is_white_mode,omitempty"`

	// Referer 的正则表达式的列表，仅支持填写 IPv4 和 IPv6 格式的 IP 地址，参数长度范围为（1，1024）。不支持域名、泛域名、CIDR 网段。最多支持设置 100 条规则。
	RegexValues []string `json:"regex_values,omitempty"`

	// 是否启用正则表达列表，取值如下所示：
	// * true：启用
	// * false：不启用
	RegexValuesEnabled bool `json:"regex_values_enabled,omitempty"`

	// Referer 通用规则列表，根据是否为白名单，为对应的白/黑名单的值。您可以指定一个或者多个 IP 地址，域名和泛域名。支持填写二级域名，支持混合输入。
	// * IP 地址格式支持 IPv4 和 IPv6，最多可输入 1000 个 IP 地址。
	// * 域名无需包含http://或https://。
	// :::tip
	// * values 和 regex_valses 均存在时，两者同时生效。
	// * 若您需要对同类型名单内已设定的values地址进行增删处理，那么您可调用获取域名配置 [https://www.volcengine.com/docs/508/9366#refer-link]接口获取已配置的全部地址列表后，在此基础上添加或删除您期望变更的值，最后重新传入values。
	// :::
	Values []string `json:"values,omitempty"`
}

type UpdateReferQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateReferRes struct {
	ResponseMetadata *UpdateReferResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                          `json:"Result,omitempty"`
}

type UpdateReferResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateResEventRuleBody struct {

	// REQUIRED; 事件通知规则
	EventRules []*UpdateResEventRuleBodyEventRulesItem `json:"EventRules"`
}

type UpdateResEventRuleBodyEventRulesItem struct {

	// REQUIRED; 回调 URL，以 http:// 或 https:// 开头，需满足公网可访问。当事件触发时，会向该 URL 发送 HTTP POST 请求，body 为具体的事件信息。具体回调参数详见回调内容。
	CallbackURL string `json:"CallbackUrl"`

	// REQUIRED; 规则启用状态，取值如下所示：
	// * true：开启
	// * false：关闭
	Enable bool `json:"Enable"`

	// REQUIRED; 事件类型。取值如下所示：
	// * Upload：上传文件
	// * Delete：删除文件
	// * Mirror：镜像回源
	// * Migrate：数据迁移
	// * OffTrans：离线转码（素材托管服务配置无效）
	// * TplStore：模板持久化存储（素材托管服务配置无效）
	EventType []string `json:"EventType"`

	// 匹配规则的正则表达式。仅当资源的 StoreKey 匹配该正则表达式时触发事件通知。缺省情况下表示匹配所有资源。
	MatchRule string `json:"MatchRule,omitempty"`
}

type UpdateResEventRuleQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateResEventRuleRes struct {

	// REQUIRED
	ResponseMetadata *UpdateResEventRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateResEventRuleResResponseMetadata struct {

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

type UpdateResponseHeaderBody struct {

	// 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"domain,omitempty"`

	// Resp Headers 配置
	RespHdrs []*UpdateResponseHeaderBodyRespHdrsItem `json:"resp_hdrs,omitempty"`
}

type UpdateResponseHeaderBodyRespHdrsItem struct {

	// REQUIRED; 在 veImageX 响应用户请求时，是否校验请求头中的 Origin 字段。仅对响应头部Access-Control-Allow-Origin有效，取值如下所示：
	// * true：开启校验，veImageX 会校验 Origin 字段。 * 如果校验成功，响应头中会包含 Access-Control-Allow-Origin 字段。字段值与 Origin 字段值相同。
	// * 如果校验失败，响应头中不会包含 Access-Control-Allow-Origin 字段。
	//
	//
	// * false：关闭校验，veImageX 不会校验 Origin 字段。响应头中将始终包含 Access-Control-Allow-Origin 头部和您配置的值。
	AccessOriginControl bool `json:"access_origin_control"`

	// Header Key，请见支持配置的响应头 [https://www.volcengine.com/docs/508/196704#%E6%94%AF%E6%8C%81%E9%85%8D%E7%BD%AE%E7%9A%84%E5%93%8D%E5%BA%94%E5%A4%B4]。
	Key string `json:"key,omitempty"`

	// Header Value，设置该响应头字段的值。字段值不能超过 1,024 个字符，可以包含除美元符号（$），Delete（ASCII code 127）外的可打印 ASCII 字符。
	Value string `json:"value,omitempty"`
}

type UpdateResponseHeaderQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateResponseHeaderRes struct {
	ResponseMetadata *UpdateResponseHeaderResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           string                                   `json:"Result,omitempty"`
}

type UpdateResponseHeaderResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateServiceNameBody struct {

	// REQUIRED; 服务名称，最多不超过 32 个字符。
	ServiceName string `json:"ServiceName"`
}

type UpdateServiceNameQuery struct {

	// REQUIRED; 待修改名称的服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateServiceNameRes struct {

	// REQUIRED
	ResponseMetadata *UpdateServiceNameResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateServiceNameResResponseMetadata struct {

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

type UpdateSlimConfigBody struct {

	// REQUIRED; 是否关闭持久化。取值如下所示：
	// * true：关闭
	// * false：开启
	DiscardSlimedFile bool `json:"DiscardSlimedFile"`

	// REQUIRED; 是否开启集智瘦身，取值如下所示：
	// * true：开启集智瘦身
	// * false：关闭集智瘦身
	DoSlim bool `json:"DoSlim"`

	// REQUIRED; 域名，您可以通过获取服务下全部域名 [https://www.volcengine.com/docs/508/9379]获取服务下域名信息。
	Domain string `json:"Domain"`
}

type UpdateSlimConfigQuery struct {

	// REQUIRED; 待修改配置的域名的所属服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateSlimConfigRes struct {
	ResponseMetadata *UpdateSlimConfigResResponseMetadata `json:"ResponseMetadata,omitempty"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSlimConfigResResponseMetadata struct {
	Action    string `json:"Action,omitempty"`
	Region    string `json:"Region,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
	Service   string `json:"Service,omitempty"`
	Version   string `json:"Version,omitempty"`
}

type UpdateStorageRulesBody struct {

	// 更新后的存储降冷策略
	StorageRules []*UpdateStorageRulesBodyStorageRulesItem `json:"StorageRules,omitempty"`
}

type UpdateStorageRulesBodyStorageRulesItem struct {

	// REQUIRED; 最新版本文件在策略命中后需要执行的操作，取值如下所示：
	// * DELETE：删除文件
	// * IA：文件转低频存储
	// * ARCHIVE：文件转归档存储
	// * COLD_ARCHIVE：文件转冷归档存储
	Action string `json:"Action"`

	// REQUIRED; 最新版本文件的策略天数，取值范围为 [1,365]，单位为天。按照 Event 事件 Day 天后执行 Action 事件，即当匹配文件的上传时间符合指定天数后，自动按照处理策略对资源进行处理。
	Day int `json:"Day"`

	// REQUIRED; 是否启用。 * false： ；
	// * true： 。
	// *
	// * false
	// *
	Enable bool `json:"Enable"`

	// REQUIRED; 策略类型，固定取值 Upload，表示按上传时间。
	Event string `json:"Event"`

	// REQUIRED; 非当前操作类型。
	NonCurrentAction string `json:"NonCurrentAction"`

	// REQUIRED; 非当前日期。
	NonCurrentDay int `json:"NonCurrentDay"`

	// 文件前缀。例如设置为prefix后，规则将只对名称以prefix开头的存储资源生效。输入规则如下：
	// * 不能以正斜线（/）或者反斜线（\）开头；
	// * 不支持使用正则表达式匹配前缀；
	// * 长度为 1～1024 个字符。
	Prefix string `json:"Prefix,omitempty"`
}

type UpdateStorageRulesQuery struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateStorageRulesRes struct {

	// REQUIRED
	ResponseMetadata *UpdateStorageRulesResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateStorageRulesResResponseMetadata struct {

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

type UpdateStorageRulesV2Body struct {

	// 存储规则。
	StorageRules []*UpdateStorageRulesV2BodyStorageRulesItem `json:"StorageRules,omitempty"`
}

type UpdateStorageRulesV2BodyStorageRulesItem struct {

	// REQUIRED; 是否启用。
	Enable bool `json:"Enable"`

	// REQUIRED; 事件名称，目前仅支持upload
	Event string `json:"Event"`

	// REQUIRED; 规则名称，长度256以内
	ID string `json:"Id"`

	// REQUIRED; 文件前缀
	Prefix string `json:"Prefix"`

	// 中止未完成的分段上传配置。
	AbortIncompleteMultipartUpload *UpdateStorageRulesV2BodyStorageRulesItemAbortIncompleteMultipartUpload `json:"AbortIncompleteMultipartUpload,omitempty"`

	// 过期时间。过期后会删除文件。过期时间要小于Transitions中天数，另外不能与Transitions里混用Day与Date
	Expiration *UpdateStorageRulesV2BodyStorageRulesItemExpiration `json:"Expiration,omitempty"`

	// 过滤条件。
	Filter *UpdateStorageRulesV2BodyStorageRulesItemFilter `json:"Filter,omitempty"`

	// 多版本过期配置。过期后会删除旧版本文件。
	NonCurrentVersionExpiration *UpdateStorageRulesV2BodyStorageRulesItemNonCurrentVersionExpiration `json:"NonCurrentVersionExpiration,omitempty"`

	// 非当前版本转换规则。要么全部days要么全部date，不能混合配置
	NonCurrentVersionTransitions []*ComponentsQpab6RSchemasUpdatestoragerulesv2BodyPropertiesStoragerulesItemsPropertiesNoncurrentversiontransitionsItems `json:"NonCurrentVersionTransitions,omitempty"`

	// 存储沉降配置，要么全部days要么全部date，不能混合配置
	Transitions []*UpdateStorageRulesV2BodyStorageRulesPropertiesItemsItem `json:"Transitions,omitempty"`
}

// UpdateStorageRulesV2BodyStorageRulesItemAbortIncompleteMultipartUpload - 中止未完成的分段上传配置。
type UpdateStorageRulesV2BodyStorageRulesItemAbortIncompleteMultipartUpload struct {

	// 分片过期时间
	DaysAfterInitiation int `json:"DaysAfterInitiation,omitempty"`
}

// UpdateStorageRulesV2BodyStorageRulesItemExpiration - 过期时间。过期后会删除文件。过期时间要小于Transitions中天数，另外不能与Transitions里混用Day与Date
type UpdateStorageRulesV2BodyStorageRulesItemExpiration struct {

	// 过期日期。与day冲突
	Date string `json:"Date,omitempty"`

	// 过期天数。与date冲突
	Days int `json:"Days,omitempty"`
}

// UpdateStorageRulesV2BodyStorageRulesItemFilter - 过滤条件。
type UpdateStorageRulesV2BodyStorageRulesItemFilter struct {

	// 是否包含大于等于。
	GreaterThanIncludeEqual string `json:"GreaterThanIncludeEqual,omitempty"`

	// 小于等于条件。
	LessThanIncludeEqual string `json:"LessThanIncludeEqual,omitempty"`

	// 对象大小大于。
	ObjectSizeGreaterThan int `json:"ObjectSizeGreaterThan,omitempty"`

	// 对象大小上限。
	ObjectSizeLessThan int `json:"ObjectSizeLessThan,omitempty"`
}

// UpdateStorageRulesV2BodyStorageRulesItemNonCurrentVersionExpiration - 多版本过期配置。过期后会删除旧版本文件。
type UpdateStorageRulesV2BodyStorageRulesItemNonCurrentVersionExpiration struct {

	// 日期，只填写年月日，无法指定时分秒
	NonCurrentDate string `json:"NonCurrentDate,omitempty"`

	// 多版本过期天数。
	NonCurrentDays int `json:"NonCurrentDays,omitempty"`
}

type UpdateStorageRulesV2BodyStorageRulesPropertiesItemsItem struct {

	// REQUIRED; 沉降存储类别。
	// * IA，转为低频存储
	// * ARCHIVE_FR，转为归档闪回存储
	// * COLD_ARCHIVE，转为冷归档存储
	// * ARCHIVE，转为归档存储
	StorageClass string `json:"StorageClass"`

	// 日期。
	Date string `json:"Date,omitempty"`

	// 天数。
	Days int `json:"Days,omitempty"`
}

type UpdateStorageRulesV2Query struct {

	// REQUIRED; 服务 ID。
	// * 您可以在 veImageX 控制台服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，在创建好的图片服务中获取服务 ID。
	// * 您也可以通过 OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId" query:"ServiceId"`
}

type UpdateStorageRulesV2Res struct {

	// REQUIRED
	ResponseMetadata *UpdateStorageRulesV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateStorageRulesV2ResResponseMetadata struct {

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

type VerifyDomainOwnerBody struct {

	// REQUIRED; 待校验的域名。仅支持单域名校验。
	Domain string `json:"Domain"`

	// REQUIRED; 校验方式，取值如下： dns: DNS 解析验证。
	VerifyType string `json:"VerifyType"`
}

type VerifyDomainOwnerRes struct {

	// REQUIRED
	ResponseMetadata *VerifyDomainOwnerResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *VerifyDomainOwnerResResult `json:"Result,omitempty"`
}

type VerifyDomainOwnerResResponseMetadata struct {

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

// VerifyDomainOwnerResResult - 视请求的接口而定
type VerifyDomainOwnerResResult struct {

	// REQUIRED; 校验失败的原因，当 VerifyResult 为 false 时返回。 verify domain owner by dns error: 域名归属权 DNS 验证错误，DNS 解析失败。 verify domain
	// owner by dns failed, record value not match: 域名归属权 DNS 验证错误，TXT 记录值不匹配。
	ErrorMessage string `json:"ErrorMessage"`

	// REQUIRED; 校验是否成功： true: 校验成功 false: 校验失败
	VerifyResult bool `json:"VerifyResult"`
}
type AIProcess struct{}
type AIProcessQuery struct{}
type AddDomainV1 struct{}
type AddImageBackgroundColors struct{}
type AddImageBackgroundColorsQuery struct{}
type AddImageElements struct{}
type AddImageElementsQuery struct{}
type ApplyImageUpload struct{}
type ApplyImageUploadBody struct{}
type ApplyVpcUploadInfo struct{}
type ApplyVpcUploadInfoBody struct{}
type CommitImageUpload struct{}
type CreateBatchProcessTask struct{}
type CreateCVImageGenerateTask struct{}
type CreateFileRestore struct{}
type CreateHiddenWatermarkImage struct{}
type CreateHmExtractTask struct{}
type CreateImageAIProcessCallback struct{}
type CreateImageAIProcessCallbackQuery struct{}
type CreateImageAIProcessQueue struct{}
type CreateImageAIProcessQueueQuery struct{}
type CreateImageAITask struct{}
type CreateImageAITaskQuery struct{}
type CreateImageAnalyzeTask struct{}
type CreateImageAnalyzeTaskQuery struct{}
type CreateImageAuditTask struct{}
type CreateImageAuditTaskQuery struct{}
type CreateImageCompressTask struct{}
type CreateImageContentTask struct{}
type CreateImageContentTaskQuery struct{}
type CreateImageFromURI struct{}
type CreateImageHmEmbed struct{}
type CreateImageHmEmbedQuery struct{}
type CreateImageHmExtract struct{}
type CreateImageHmExtractBody struct{}
type CreateImageMigrateTask struct{}
type CreateImageMigrateTaskQuery struct{}
type CreateImageMonitorRule struct{}
type CreateImageMonitorRuleQuery struct{}
type CreateImageRetryAuditTask struct{}
type CreateImageRetryAuditTaskQuery struct{}
type CreateImageService struct{}
type CreateImageServiceQuery struct{}
type CreateImageSettingRule struct{}
type CreateImageSettingRuleQuery struct{}
type CreateImageStyle struct{}
type CreateImageStyleQuery struct{}
type CreateImageTemplate struct{}
type CreateImageTemplatesByImport struct{}
type CreateImageTemplatesByImportQuery struct{}
type CreateImageTranscodeCallback struct{}
type CreateImageTranscodeCallbackQuery struct{}
type CreateImageTranscodeQueue struct{}
type CreateImageTranscodeQueueQuery struct{}
type CreateImageTranscodeTask struct{}
type CreateImageTranscodeTaskQuery struct{}
type CreateTemplatesFromBin struct{}
type DelDomain struct{}
type DeleteImageAIProcessDetail struct{}
type DeleteImageAIProcessDetailQuery struct{}
type DeleteImageAIProcessQueue struct{}
type DeleteImageAIProcessQueueQuery struct{}
type DeleteImageAnalyzeTask struct{}
type DeleteImageAnalyzeTaskQuery struct{}
type DeleteImageAnalyzeTaskRun struct{}
type DeleteImageAnalyzeTaskRunQuery struct{}
type DeleteImageAuditResult struct{}
type DeleteImageAuditResultQuery struct{}
type DeleteImageBackgroundColors struct{}
type DeleteImageBackgroundColorsQuery struct{}
type DeleteImageElements struct{}
type DeleteImageElementsQuery struct{}
type DeleteImageMigrateTask struct{}
type DeleteImageMigrateTaskBody struct{}
type DeleteImageMonitorRecords struct{}
type DeleteImageMonitorRecordsQuery struct{}
type DeleteImageMonitorRules struct{}
type DeleteImageMonitorRulesQuery struct{}
type DeleteImageService struct{}
type DeleteImageServiceBody struct{}
type DeleteImageSettingRule struct{}
type DeleteImageSettingRuleQuery struct{}
type DeleteImageStyle struct{}
type DeleteImageStyleQuery struct{}
type DeleteImageTemplate struct{}
type DeleteImageTranscodeDetail struct{}
type DeleteImageTranscodeDetailQuery struct{}
type DeleteImageTranscodeQueue struct{}
type DeleteImageTranscodeQueueQuery struct{}
type DeleteImageUploadFiles struct{}
type DeleteTemplatesFromBin struct{}
type DescribeImageVolcCdnAccessLog struct{}
type DescribeImageXAIRequestCntUsage struct{}
type DescribeImageXAIRequestCntUsageBody struct{}
type DescribeImageXAddOnQPSUsage struct{}
type DescribeImageXAddOnQPSUsageBody struct{}
type DescribeImageXBaseOpUsage struct{}
type DescribeImageXBaseOpUsageBody struct{}
type DescribeImageXBillingRequestCntUsage struct{}
type DescribeImageXBillingRequestCntUsageBody struct{}
type DescribeImageXBucketRetrievalUsage struct{}
type DescribeImageXBucketRetrievalUsageBody struct{}
type DescribeImageXBucketUsage struct{}
type DescribeImageXBucketUsageBody struct{}
type DescribeImageXCDNTopRequestData struct{}
type DescribeImageXCDNTopRequestDataBody struct{}
type DescribeImageXCdnDurationAll struct{}
type DescribeImageXCdnDurationAllQuery struct{}
type DescribeImageXCdnDurationDetailByTime struct{}
type DescribeImageXCdnDurationDetailByTimeQuery struct{}
type DescribeImageXCdnErrorCodeAll struct{}
type DescribeImageXCdnErrorCodeAllQuery struct{}
type DescribeImageXCdnErrorCodeByTime struct{}
type DescribeImageXCdnErrorCodeByTimeQuery struct{}
type DescribeImageXCdnProtocolRateByTime struct{}
type DescribeImageXCdnProtocolRateByTimeQuery struct{}
type DescribeImageXCdnReuseRateAll struct{}
type DescribeImageXCdnReuseRateAllQuery struct{}
type DescribeImageXCdnReuseRateByTime struct{}
type DescribeImageXCdnReuseRateByTimeQuery struct{}
type DescribeImageXCdnSuccessRateAll struct{}
type DescribeImageXCdnSuccessRateAllQuery struct{}
type DescribeImageXCdnSuccessRateByTime struct{}
type DescribeImageXCdnSuccessRateByTimeQuery struct{}
type DescribeImageXClientCountByTime struct{}
type DescribeImageXClientCountByTimeQuery struct{}
type DescribeImageXClientDecodeDurationByTime struct{}
type DescribeImageXClientDecodeDurationByTimeQuery struct{}
type DescribeImageXClientDecodeSuccessRateByTime struct{}
type DescribeImageXClientDecodeSuccessRateByTimeQuery struct{}
type DescribeImageXClientDemotionRateByTime struct{}
type DescribeImageXClientDemotionRateByTimeQuery struct{}
type DescribeImageXClientErrorCodeAll struct{}
type DescribeImageXClientErrorCodeAllQuery struct{}
type DescribeImageXClientErrorCodeByTime struct{}
type DescribeImageXClientErrorCodeByTimeQuery struct{}
type DescribeImageXClientFailureRate struct{}
type DescribeImageXClientFailureRateQuery struct{}
type DescribeImageXClientFileSize struct{}
type DescribeImageXClientFileSizeQuery struct{}
type DescribeImageXClientLoadDuration struct{}
type DescribeImageXClientLoadDurationAll struct{}
type DescribeImageXClientLoadDurationAllQuery struct{}
type DescribeImageXClientLoadDurationQuery struct{}
type DescribeImageXClientQualityRateByTime struct{}
type DescribeImageXClientQualityRateByTimeQuery struct{}
type DescribeImageXClientQueueDurationByTime struct{}
type DescribeImageXClientQueueDurationByTimeQuery struct{}
type DescribeImageXClientScoreByTime struct{}
type DescribeImageXClientScoreByTimeQuery struct{}
type DescribeImageXClientSdkVerByTime struct{}
type DescribeImageXClientSdkVerByTimeQuery struct{}
type DescribeImageXClientTopDemotionURL struct{}
type DescribeImageXClientTopDemotionURLQuery struct{}
type DescribeImageXClientTopFileSize struct{}
type DescribeImageXClientTopFileSizeQuery struct{}
type DescribeImageXClientTopQualityURL struct{}
type DescribeImageXClientTopQualityURLQuery struct{}
type DescribeImageXCompressUsage struct{}
type DescribeImageXCompressUsageBody struct{}
type DescribeImageXCubeUsage struct{}
type DescribeImageXCubeUsageBody struct{}
type DescribeImageXDomainBandwidthData struct{}
type DescribeImageXDomainBandwidthDataBody struct{}
type DescribeImageXDomainBandwidthNinetyFiveData struct{}
type DescribeImageXDomainBandwidthNinetyFiveDataBody struct{}
type DescribeImageXDomainTrafficData struct{}
type DescribeImageXDomainTrafficDataBody struct{}
type DescribeImageXEdgeRequest struct{}
type DescribeImageXEdgeRequestBandwidth struct{}
type DescribeImageXEdgeRequestBandwidthBody struct{}
type DescribeImageXEdgeRequestBody struct{}
type DescribeImageXEdgeRequestRegions struct{}
type DescribeImageXEdgeRequestRegionsBody struct{}
type DescribeImageXEdgeRequestTraffic struct{}
type DescribeImageXEdgeRequestTrafficBody struct{}
type DescribeImageXExceedCountByTime struct{}
type DescribeImageXExceedCountByTimeQuery struct{}
type DescribeImageXExceedFileSize struct{}
type DescribeImageXExceedFileSizeQuery struct{}
type DescribeImageXExceedResolutionRatioAll struct{}
type DescribeImageXExceedResolutionRatioAllQuery struct{}
type DescribeImageXHeifEncodeDurationByTime struct{}
type DescribeImageXHeifEncodeDurationByTimeQuery struct{}
type DescribeImageXHeifEncodeErrorCodeByTime struct{}
type DescribeImageXHeifEncodeErrorCodeByTimeQuery struct{}
type DescribeImageXHeifEncodeFileInSizeByTime struct{}
type DescribeImageXHeifEncodeFileInSizeByTimeQuery struct{}
type DescribeImageXHeifEncodeFileOutSizeByTime struct{}
type DescribeImageXHeifEncodeFileOutSizeByTimeQuery struct{}
type DescribeImageXHeifEncodeSuccessCountByTime struct{}
type DescribeImageXHeifEncodeSuccessCountByTimeQuery struct{}
type DescribeImageXHeifEncodeSuccessRateByTime struct{}
type DescribeImageXHeifEncodeSuccessRateByTimeQuery struct{}
type DescribeImageXHitRateRequestData struct{}
type DescribeImageXHitRateRequestDataBody struct{}
type DescribeImageXHitRateTrafficData struct{}
type DescribeImageXHitRateTrafficDataBody struct{}
type DescribeImageXMirrorRequestBandwidth struct{}
type DescribeImageXMirrorRequestBandwidthQuery struct{}
type DescribeImageXMirrorRequestHTTPCodeByTime struct{}
type DescribeImageXMirrorRequestHTTPCodeByTimeQuery struct{}
type DescribeImageXMirrorRequestHTTPCodeOverview struct{}
type DescribeImageXMirrorRequestHTTPCodeOverviewQuery struct{}
type DescribeImageXMirrorRequestTraffic struct{}
type DescribeImageXMirrorRequestTrafficQuery struct{}
type DescribeImageXMultiCompressUsage struct{}
type DescribeImageXMultiCompressUsageBody struct{}
type DescribeImageXRequestCntUsage struct{}
type DescribeImageXRequestCntUsageBody struct{}
type DescribeImageXScreenshotUsage struct{}
type DescribeImageXScreenshotUsageBody struct{}
type DescribeImageXSensibleCacheHitRateByTime struct{}
type DescribeImageXSensibleCacheHitRateByTimeQuery struct{}
type DescribeImageXSensibleCountByTime struct{}
type DescribeImageXSensibleCountByTimeQuery struct{}
type DescribeImageXSensibleTopRAMURL struct{}
type DescribeImageXSensibleTopRAMURLQuery struct{}
type DescribeImageXSensibleTopResolutionURL struct{}
type DescribeImageXSensibleTopResolutionURLQuery struct{}
type DescribeImageXSensibleTopSizeURL struct{}
type DescribeImageXSensibleTopSizeURLQuery struct{}
type DescribeImageXSensibleTopUnknownURL struct{}
type DescribeImageXSensibleTopUnknownURLQuery struct{}
type DescribeImageXServerQPSUsage struct{}
type DescribeImageXServerQPSUsageBody struct{}
type DescribeImageXServiceQuality struct{}
type DescribeImageXServiceQualityBody struct{}
type DescribeImageXSourceRequest struct{}
type DescribeImageXSourceRequestBandwidth struct{}
type DescribeImageXSourceRequestBandwidthBody struct{}
type DescribeImageXSourceRequestBody struct{}
type DescribeImageXSourceRequestTraffic struct{}
type DescribeImageXSourceRequestTrafficBody struct{}
type DescribeImageXStorageUsage struct{}
type DescribeImageXStorageUsageBody struct{}
type DescribeImageXSummary struct{}
type DescribeImageXSummaryBody struct{}
type DescribeImageXUploadCountByTime struct{}
type DescribeImageXUploadCountByTimeQuery struct{}
type DescribeImageXUploadDuration struct{}
type DescribeImageXUploadDurationQuery struct{}
type DescribeImageXUploadErrorCodeAll struct{}
type DescribeImageXUploadErrorCodeAllQuery struct{}
type DescribeImageXUploadErrorCodeByTime struct{}
type DescribeImageXUploadErrorCodeByTimeQuery struct{}
type DescribeImageXUploadFileSize struct{}
type DescribeImageXUploadFileSizeQuery struct{}
type DescribeImageXUploadSegmentSpeedByTime struct{}
type DescribeImageXUploadSegmentSpeedByTimeQuery struct{}
type DescribeImageXUploadSpeed struct{}
type DescribeImageXUploadSpeedQuery struct{}
type DescribeImageXUploadSuccessRateByTime struct{}
type DescribeImageXUploadSuccessRateByTimeQuery struct{}
type DescribeImageXVideoClipDurationUsage struct{}
type DescribeImageXVideoClipDurationUsageBody struct{}
type DownloadCert struct{}
type DownloadCertBody struct{}
type ExportFailedMigrateTask struct{}
type ExportFailedMigrateTaskBody struct{}
type FetchImageURL struct{}
type FetchImageURLQuery struct{}
type GetAllCerts struct{}
type GetAllCertsBody struct{}
type GetAllCertsQuery struct{}
type GetAllImageServices struct{}
type GetAllImageServicesBody struct{}
type GetAllImageTemplates struct{}
type GetAllImageTemplatesBody struct{}
type GetAuditEntrysCount struct{}
type GetAuditEntrysCountBody struct{}
type GetBatchProcessResult struct{}
type GetBatchTaskInfo struct{}
type GetBatchTaskInfoBody struct{}
type GetCVAnimeGenerateImage struct{}
type GetCVImageGenerateResult struct{}
type GetCVImageGenerateTask struct{}
type GetCVTextGenerateImage struct{}
type GetCertInfo struct{}
type GetCertInfoBody struct{}
type GetComprehensiveEnhanceImage struct{}
type GetComprehensiveEnhanceImageQuery struct{}
type GetCompressTaskInfo struct{}
type GetCompressTaskInfoBody struct{}
type GetDedupTaskStatus struct{}
type GetDedupTaskStatusBody struct{}
type GetDenoisingImage struct{}
type GetDomainConfig struct{}
type GetDomainConfigBody struct{}
type GetDomainOwnerVerifyContent struct{}
type GetDomainOwnerVerifyContentBody struct{}
type GetImageAIDetails struct{}
type GetImageAIDetailsBody struct{}
type GetImageAIProcessQueues struct{}
type GetImageAIProcessQueuesBody struct{}
type GetImageAITasks struct{}
type GetImageAITasksBody struct{}
type GetImageAddOnTag struct{}
type GetImageAddOnTagBody struct{}
type GetImageAiGenerateTask struct{}
type GetImageAiGenerateTaskBody struct{}
type GetImageAlertRecords struct{}
type GetImageAlertRecordsQuery struct{}
type GetImageAllDomainCert struct{}
type GetImageAllDomainCertBody struct{}
type GetImageAllDomainCertQuery struct{}
type GetImageAnalyzeResult struct{}
type GetImageAnalyzeResultBody struct{}
type GetImageAnalyzeTasks struct{}
type GetImageAnalyzeTasksBody struct{}
type GetImageAuditResult struct{}
type GetImageAuditResultBody struct{}
type GetImageAuditTasks struct{}
type GetImageAuditTasksBody struct{}
type GetImageAuthKey struct{}
type GetImageAuthKeyBody struct{}
type GetImageBackgroundColors struct{}
type GetImageBackgroundColorsBody struct{}
type GetImageBackgroundColorsQuery struct{}
type GetImageBgFillResult struct{}
type GetImageBgFillResultQuery struct{}
type GetImageComicResult struct{}
type GetImageComicResultQuery struct{}
type GetImageContentBlockList struct{}
type GetImageContentBlockListQuery struct{}
type GetImageContentTaskDetail struct{}
type GetImageContentTaskDetailQuery struct{}
type GetImageDetectResult struct{}
type GetImageDuplicateDetection struct{}
type GetImageElements struct{}
type GetImageElementsBody struct{}
type GetImageEnhanceResult struct{}
type GetImageEnhanceResultQuery struct{}
type GetImageEraseModels struct{}
type GetImageEraseModelsBody struct{}
type GetImageEraseResult struct{}
type GetImageEraseResultQuery struct{}
type GetImageFonts struct{}
type GetImageFontsBody struct{}
type GetImageFontsQuery struct{}
type GetImageHmExtractTaskInfo struct{}
type GetImageMigrateTasks struct{}
type GetImageMigrateTasksBody struct{}
type GetImageMonitorRules struct{}
type GetImageMonitorRulesBody struct{}
type GetImageOCRV2 struct{}
type GetImagePSDetection struct{}
type GetImageQuality struct{}
type GetImageService struct{}
type GetImageServiceBody struct{}
type GetImageServiceSubscription struct{}
type GetImageServiceSubscriptionBody struct{}
type GetImageSettingRuleHistory struct{}
type GetImageSettingRuleHistoryBody struct{}
type GetImageSettingRules struct{}
type GetImageSettingRulesBody struct{}
type GetImageSettings struct{}
type GetImageSettingsBody struct{}
type GetImageSmartCropResult struct{}
type GetImageSmartCropResultQuery struct{}
type GetImageStorageFiles struct{}
type GetImageStorageFilesBody struct{}
type GetImageStyleDetail struct{}
type GetImageStyleDetailBody struct{}
type GetImageStyleResult struct{}
type GetImageStyles struct{}
type GetImageStylesBody struct{}
type GetImageSuperResolutionResult struct{}
type GetImageSuperResolutionResultQuery struct{}
type GetImageTemplate struct{}
type GetImageTemplateBody struct{}
type GetImageTranscodeDetails struct{}
type GetImageTranscodeDetailsBody struct{}
type GetImageTranscodeQueues struct{}
type GetImageTranscodeQueuesBody struct{}
type GetImageUpdateFiles struct{}
type GetImageUpdateFilesBody struct{}
type GetImageUploadFile struct{}
type GetImageUploadFileBody struct{}
type GetImageUploadFiles struct{}
type GetImageUploadFilesBody struct{}
type GetImageXQueryApps struct{}
type GetImageXQueryAppsBody struct{}
type GetImageXQueryDims struct{}
type GetImageXQueryDimsBody struct{}
type GetImageXQueryRegions struct{}
type GetImageXQueryRegionsBody struct{}
type GetImageXQueryVals struct{}
type GetImageXQueryValsBody struct{}
type GetLicensePlateDetection struct{}
type GetPrivateImageType struct{}
type GetProductAIGCResult struct{}
type GetResourceURL struct{}
type GetResourceURLBody struct{}
type GetResponseHeaderValidateKeys struct{}
type GetResponseHeaderValidateKeysBody struct{}
type GetResponseHeaderValidateKeysQuery struct{}
type GetSegmentImage struct{}
type GetServiceDomains struct{}
type GetServiceDomainsBody struct{}
type GetSyncAuditResult struct{}
type GetTemplatesFromBin struct{}
type GetTemplatesFromBinBody struct{}
type GetURLFetchTask struct{}
type GetURLFetchTaskBody struct{}
type GetVendorBuckets struct{}
type GetVendorBucketsQuery struct{}
type PreviewImageUploadFile struct{}
type PreviewImageUploadFileBody struct{}
type ReportEvent struct{}
type ReportEventBody struct{}
type ReportEventQuery struct{}
type RerunImageMigrateTask struct{}
type RerunImageMigrateTaskBody struct{}
type SetDefaultDomain struct{}
type SetDefaultDomainQuery struct{}
type TerminateImageMigrateTask struct{}
type TerminateImageMigrateTaskBody struct{}
type UpdateAdvance struct{}
type UpdateAuditImageStatus struct{}
type UpdateAuditImageStatusQuery struct{}
type UpdateDomainAdaptiveFmt struct{}
type UpdateFileStorageClass struct{}
type UpdateHTTPS struct{}
type UpdateImageAIProcessQueue struct{}
type UpdateImageAIProcessQueueQuery struct{}
type UpdateImageAIProcessQueueStatus struct{}
type UpdateImageAIProcessQueueStatusQuery struct{}
type UpdateImageAnalyzeTask struct{}
type UpdateImageAnalyzeTaskQuery struct{}
type UpdateImageAnalyzeTaskStatus struct{}
type UpdateImageAnalyzeTaskStatusQuery struct{}
type UpdateImageAuditTask struct{}
type UpdateImageAuditTaskQuery struct{}
type UpdateImageAuditTaskStatus struct{}
type UpdateImageAuditTaskStatusQuery struct{}
type UpdateImageAuthKey struct{}
type UpdateImageDomainAreaAccess struct{}
type UpdateImageDomainBandwidthLimit struct{}
type UpdateImageDomainConfig struct{}
type UpdateImageDomainDownloadSpeedLimit struct{}
type UpdateImageDomainIPAuth struct{}
type UpdateImageDomainUaAccess struct{}
type UpdateImageDomainVolcOrigin struct{}
type UpdateImageExifData struct{}
type UpdateImageFileCT struct{}
type UpdateImageFileKey struct{}
type UpdateImageMirrorConf struct{}
type UpdateImageMonitorRule struct{}
type UpdateImageMonitorRuleQuery struct{}
type UpdateImageMonitorRuleStatus struct{}
type UpdateImageMonitorRuleStatusQuery struct{}
type UpdateImageObjectAccess struct{}
type UpdateImageResourceStatus struct{}
type UpdateImageSettingRule struct{}
type UpdateImageSettingRulePriority struct{}
type UpdateImageSettingRulePriorityQuery struct{}
type UpdateImageSettingRuleQuery struct{}
type UpdateImageStorageTTL struct{}
type UpdateImageStorageTTLQuery struct{}
type UpdateImageStyle struct{}
type UpdateImageStyleMeta struct{}
type UpdateImageStyleMetaQuery struct{}
type UpdateImageStyleQuery struct{}
type UpdateImageTaskStrategy struct{}
type UpdateImageTaskStrategyQuery struct{}
type UpdateImageTranscodeQueue struct{}
type UpdateImageTranscodeQueueQuery struct{}
type UpdateImageTranscodeQueueStatus struct{}
type UpdateImageTranscodeQueueStatusQuery struct{}
type UpdateImageUploadFiles struct{}
type UpdateImageUploadOverwrite struct{}
type UpdateRefer struct{}
type UpdateResEventRule struct{}
type UpdateResponseHeader struct{}
type UpdateServiceName struct{}
type UpdateSlimConfig struct{}
type UpdateStorageRules struct{}
type UpdateStorageRulesV2 struct{}
type VerifyDomainOwner struct{}
type VerifyDomainOwnerQuery struct{}
type AIProcessReq struct {
	*AIProcessQuery
	*AIProcessBody
}
type AddDomainV1Req struct {
	*AddDomainV1Query
	*AddDomainV1Body
}
type AddImageBackgroundColorsReq struct {
	*AddImageBackgroundColorsQuery
	*AddImageBackgroundColorsBody
}
type AddImageElementsReq struct {
	*AddImageElementsQuery
	*AddImageElementsBody
}
type ApplyImageUploadReq struct {
	*ApplyImageUploadQuery
	*ApplyImageUploadBody
}
type ApplyVpcUploadInfoReq struct {
	*ApplyVpcUploadInfoQuery
	*ApplyVpcUploadInfoBody
}
type CommitImageUploadReq struct {
	*CommitImageUploadQuery
	*CommitImageUploadBody
}
type CreateBatchProcessTaskReq struct {
	*CreateBatchProcessTaskQuery
	*CreateBatchProcessTaskBody
}
type CreateCVImageGenerateTaskReq struct {
	*CreateCVImageGenerateTaskQuery
	*CreateCVImageGenerateTaskBody
}
type CreateFileRestoreReq struct {
	*CreateFileRestoreQuery
	*CreateFileRestoreBody
}
type CreateHiddenWatermarkImageReq struct {
	*CreateHiddenWatermarkImageQuery
	*CreateHiddenWatermarkImageBody
}
type CreateHmExtractTaskReq struct {
	*CreateHmExtractTaskQuery
	*CreateHmExtractTaskBody
}
type CreateImageAIProcessCallbackReq struct {
	*CreateImageAIProcessCallbackQuery
	*CreateImageAIProcessCallbackBody
}
type CreateImageAIProcessQueueReq struct {
	*CreateImageAIProcessQueueQuery
	*CreateImageAIProcessQueueBody
}
type CreateImageAITaskReq struct {
	*CreateImageAITaskQuery
	*CreateImageAITaskBody
}
type CreateImageAnalyzeTaskReq struct {
	*CreateImageAnalyzeTaskQuery
	*CreateImageAnalyzeTaskBody
}
type CreateImageAuditTaskReq struct {
	*CreateImageAuditTaskQuery
	*CreateImageAuditTaskBody
}
type CreateImageCompressTaskReq struct {
	*CreateImageCompressTaskQuery
	*CreateImageCompressTaskBody
}
type CreateImageContentTaskReq struct {
	*CreateImageContentTaskQuery
	*CreateImageContentTaskBody
}
type CreateImageFromURIReq struct {
	*CreateImageFromURIQuery
	*CreateImageFromURIBody
}
type CreateImageHmEmbedReq struct {
	*CreateImageHmEmbedQuery
	*CreateImageHmEmbedBody
}
type CreateImageHmExtractReq struct {
	*CreateImageHmExtractQuery
	*CreateImageHmExtractBody
}
type CreateImageMigrateTaskReq struct {
	*CreateImageMigrateTaskQuery
	*CreateImageMigrateTaskBody
}
type CreateImageMonitorRuleReq struct {
	*CreateImageMonitorRuleQuery
	*CreateImageMonitorRuleBody
}
type CreateImageRetryAuditTaskReq struct {
	*CreateImageRetryAuditTaskQuery
	*CreateImageRetryAuditTaskBody
}
type CreateImageServiceReq struct {
	*CreateImageServiceQuery
	*CreateImageServiceBody
}
type CreateImageSettingRuleReq struct {
	*CreateImageSettingRuleQuery
	*CreateImageSettingRuleBody
}
type CreateImageStyleReq struct {
	*CreateImageStyleQuery
	*CreateImageStyleBody
}
type CreateImageTemplateReq struct {
	*CreateImageTemplateQuery
	*CreateImageTemplateBody
}
type CreateImageTemplatesByImportReq struct {
	*CreateImageTemplatesByImportQuery
	*CreateImageTemplatesByImportBody
}
type CreateImageTranscodeCallbackReq struct {
	*CreateImageTranscodeCallbackQuery
	*CreateImageTranscodeCallbackBody
}
type CreateImageTranscodeQueueReq struct {
	*CreateImageTranscodeQueueQuery
	*CreateImageTranscodeQueueBody
}
type CreateImageTranscodeTaskReq struct {
	*CreateImageTranscodeTaskQuery
	*CreateImageTranscodeTaskBody
}
type CreateTemplatesFromBinReq struct {
	*CreateTemplatesFromBinQuery
	*CreateTemplatesFromBinBody
}
type DelDomainReq struct {
	*DelDomainQuery
	*DelDomainBody
}
type DeleteImageAIProcessDetailReq struct {
	*DeleteImageAIProcessDetailQuery
	*DeleteImageAIProcessDetailBody
}
type DeleteImageAIProcessQueueReq struct {
	*DeleteImageAIProcessQueueQuery
	*DeleteImageAIProcessQueueBody
}
type DeleteImageAnalyzeTaskReq struct {
	*DeleteImageAnalyzeTaskQuery
	*DeleteImageAnalyzeTaskBody
}
type DeleteImageAnalyzeTaskRunReq struct {
	*DeleteImageAnalyzeTaskRunQuery
	*DeleteImageAnalyzeTaskRunBody
}
type DeleteImageAuditResultReq struct {
	*DeleteImageAuditResultQuery
	*DeleteImageAuditResultBody
}
type DeleteImageBackgroundColorsReq struct {
	*DeleteImageBackgroundColorsQuery
	*DeleteImageBackgroundColorsBody
}
type DeleteImageElementsReq struct {
	*DeleteImageElementsQuery
	*DeleteImageElementsBody
}
type DeleteImageMigrateTaskReq struct {
	*DeleteImageMigrateTaskQuery
	*DeleteImageMigrateTaskBody
}
type DeleteImageMonitorRecordsReq struct {
	*DeleteImageMonitorRecordsQuery
	*DeleteImageMonitorRecordsBody
}
type DeleteImageMonitorRulesReq struct {
	*DeleteImageMonitorRulesQuery
	*DeleteImageMonitorRulesBody
}
type DeleteImageServiceReq struct {
	*DeleteImageServiceQuery
	*DeleteImageServiceBody
}
type DeleteImageSettingRuleReq struct {
	*DeleteImageSettingRuleQuery
	*DeleteImageSettingRuleBody
}
type DeleteImageStyleReq struct {
	*DeleteImageStyleQuery
	*DeleteImageStyleBody
}
type DeleteImageTemplateReq struct {
	*DeleteImageTemplateQuery
	*DeleteImageTemplateBody
}
type DeleteImageTranscodeDetailReq struct {
	*DeleteImageTranscodeDetailQuery
	*DeleteImageTranscodeDetailBody
}
type DeleteImageTranscodeQueueReq struct {
	*DeleteImageTranscodeQueueQuery
	*DeleteImageTranscodeQueueBody
}
type DeleteImageUploadFilesReq struct {
	*DeleteImageUploadFilesQuery
	*DeleteImageUploadFilesBody
}
type DeleteTemplatesFromBinReq struct {
	*DeleteTemplatesFromBinQuery
	*DeleteTemplatesFromBinBody
}
type DescribeImageVolcCdnAccessLogReq struct {
	*DescribeImageVolcCdnAccessLogQuery
	*DescribeImageVolcCdnAccessLogBody
}
type DescribeImageXAIRequestCntUsageReq struct {
	*DescribeImageXAIRequestCntUsageQuery
	*DescribeImageXAIRequestCntUsageBody
}
type DescribeImageXAddOnQPSUsageReq struct {
	*DescribeImageXAddOnQPSUsageQuery
	*DescribeImageXAddOnQPSUsageBody
}
type DescribeImageXBaseOpUsageReq struct {
	*DescribeImageXBaseOpUsageQuery
	*DescribeImageXBaseOpUsageBody
}
type DescribeImageXBillingRequestCntUsageReq struct {
	*DescribeImageXBillingRequestCntUsageQuery
	*DescribeImageXBillingRequestCntUsageBody
}
type DescribeImageXBucketRetrievalUsageReq struct {
	*DescribeImageXBucketRetrievalUsageQuery
	*DescribeImageXBucketRetrievalUsageBody
}
type DescribeImageXBucketUsageReq struct {
	*DescribeImageXBucketUsageQuery
	*DescribeImageXBucketUsageBody
}
type DescribeImageXCDNTopRequestDataReq struct {
	*DescribeImageXCDNTopRequestDataQuery
	*DescribeImageXCDNTopRequestDataBody
}
type DescribeImageXCdnDurationAllReq struct {
	*DescribeImageXCdnDurationAllQuery
	*DescribeImageXCdnDurationAllBody
}
type DescribeImageXCdnDurationDetailByTimeReq struct {
	*DescribeImageXCdnDurationDetailByTimeQuery
	*DescribeImageXCdnDurationDetailByTimeBody
}
type DescribeImageXCdnErrorCodeAllReq struct {
	*DescribeImageXCdnErrorCodeAllQuery
	*DescribeImageXCdnErrorCodeAllBody
}
type DescribeImageXCdnErrorCodeByTimeReq struct {
	*DescribeImageXCdnErrorCodeByTimeQuery
	*DescribeImageXCdnErrorCodeByTimeBody
}
type DescribeImageXCdnProtocolRateByTimeReq struct {
	*DescribeImageXCdnProtocolRateByTimeQuery
	*DescribeImageXCdnProtocolRateByTimeBody
}
type DescribeImageXCdnReuseRateAllReq struct {
	*DescribeImageXCdnReuseRateAllQuery
	*DescribeImageXCdnReuseRateAllBody
}
type DescribeImageXCdnReuseRateByTimeReq struct {
	*DescribeImageXCdnReuseRateByTimeQuery
	*DescribeImageXCdnReuseRateByTimeBody
}
type DescribeImageXCdnSuccessRateAllReq struct {
	*DescribeImageXCdnSuccessRateAllQuery
	*DescribeImageXCdnSuccessRateAllBody
}
type DescribeImageXCdnSuccessRateByTimeReq struct {
	*DescribeImageXCdnSuccessRateByTimeQuery
	*DescribeImageXCdnSuccessRateByTimeBody
}
type DescribeImageXClientCountByTimeReq struct {
	*DescribeImageXClientCountByTimeQuery
	*DescribeImageXClientCountByTimeBody
}
type DescribeImageXClientDecodeDurationByTimeReq struct {
	*DescribeImageXClientDecodeDurationByTimeQuery
	*DescribeImageXClientDecodeDurationByTimeBody
}
type DescribeImageXClientDecodeSuccessRateByTimeReq struct {
	*DescribeImageXClientDecodeSuccessRateByTimeQuery
	*DescribeImageXClientDecodeSuccessRateByTimeBody
}
type DescribeImageXClientDemotionRateByTimeReq struct {
	*DescribeImageXClientDemotionRateByTimeQuery
	*DescribeImageXClientDemotionRateByTimeBody
}
type DescribeImageXClientErrorCodeAllReq struct {
	*DescribeImageXClientErrorCodeAllQuery
	*DescribeImageXClientErrorCodeAllBody
}
type DescribeImageXClientErrorCodeByTimeReq struct {
	*DescribeImageXClientErrorCodeByTimeQuery
	*DescribeImageXClientErrorCodeByTimeBody
}
type DescribeImageXClientFailureRateReq struct {
	*DescribeImageXClientFailureRateQuery
	*DescribeImageXClientFailureRateBody
}
type DescribeImageXClientFileSizeReq struct {
	*DescribeImageXClientFileSizeQuery
	*DescribeImageXClientFileSizeBody
}
type DescribeImageXClientLoadDurationReq struct {
	*DescribeImageXClientLoadDurationQuery
	*DescribeImageXClientLoadDurationBody
}
type DescribeImageXClientLoadDurationAllReq struct {
	*DescribeImageXClientLoadDurationAllQuery
	*DescribeImageXClientLoadDurationAllBody
}
type DescribeImageXClientQualityRateByTimeReq struct {
	*DescribeImageXClientQualityRateByTimeQuery
	*DescribeImageXClientQualityRateByTimeBody
}
type DescribeImageXClientQueueDurationByTimeReq struct {
	*DescribeImageXClientQueueDurationByTimeQuery
	*DescribeImageXClientQueueDurationByTimeBody
}
type DescribeImageXClientScoreByTimeReq struct {
	*DescribeImageXClientScoreByTimeQuery
	*DescribeImageXClientScoreByTimeBody
}
type DescribeImageXClientSdkVerByTimeReq struct {
	*DescribeImageXClientSdkVerByTimeQuery
	*DescribeImageXClientSdkVerByTimeBody
}
type DescribeImageXClientTopDemotionURLReq struct {
	*DescribeImageXClientTopDemotionURLQuery
	*DescribeImageXClientTopDemotionURLBody
}
type DescribeImageXClientTopFileSizeReq struct {
	*DescribeImageXClientTopFileSizeQuery
	*DescribeImageXClientTopFileSizeBody
}
type DescribeImageXClientTopQualityURLReq struct {
	*DescribeImageXClientTopQualityURLQuery
	*DescribeImageXClientTopQualityURLBody
}
type DescribeImageXCompressUsageReq struct {
	*DescribeImageXCompressUsageQuery
	*DescribeImageXCompressUsageBody
}
type DescribeImageXCubeUsageReq struct {
	*DescribeImageXCubeUsageQuery
	*DescribeImageXCubeUsageBody
}
type DescribeImageXDomainBandwidthDataReq struct {
	*DescribeImageXDomainBandwidthDataQuery
	*DescribeImageXDomainBandwidthDataBody
}
type DescribeImageXDomainBandwidthNinetyFiveDataReq struct {
	*DescribeImageXDomainBandwidthNinetyFiveDataQuery
	*DescribeImageXDomainBandwidthNinetyFiveDataBody
}
type DescribeImageXDomainTrafficDataReq struct {
	*DescribeImageXDomainTrafficDataQuery
	*DescribeImageXDomainTrafficDataBody
}
type DescribeImageXEdgeRequestReq struct {
	*DescribeImageXEdgeRequestQuery
	*DescribeImageXEdgeRequestBody
}
type DescribeImageXEdgeRequestBandwidthReq struct {
	*DescribeImageXEdgeRequestBandwidthQuery
	*DescribeImageXEdgeRequestBandwidthBody
}
type DescribeImageXEdgeRequestRegionsReq struct {
	*DescribeImageXEdgeRequestRegionsQuery
	*DescribeImageXEdgeRequestRegionsBody
}
type DescribeImageXEdgeRequestTrafficReq struct {
	*DescribeImageXEdgeRequestTrafficQuery
	*DescribeImageXEdgeRequestTrafficBody
}
type DescribeImageXExceedCountByTimeReq struct {
	*DescribeImageXExceedCountByTimeQuery
	*DescribeImageXExceedCountByTimeBody
}
type DescribeImageXExceedFileSizeReq struct {
	*DescribeImageXExceedFileSizeQuery
	*DescribeImageXExceedFileSizeBody
}
type DescribeImageXExceedResolutionRatioAllReq struct {
	*DescribeImageXExceedResolutionRatioAllQuery
	*DescribeImageXExceedResolutionRatioAllBody
}
type DescribeImageXHeifEncodeDurationByTimeReq struct {
	*DescribeImageXHeifEncodeDurationByTimeQuery
	*DescribeImageXHeifEncodeDurationByTimeBody
}
type DescribeImageXHeifEncodeErrorCodeByTimeReq struct {
	*DescribeImageXHeifEncodeErrorCodeByTimeQuery
	*DescribeImageXHeifEncodeErrorCodeByTimeBody
}
type DescribeImageXHeifEncodeFileInSizeByTimeReq struct {
	*DescribeImageXHeifEncodeFileInSizeByTimeQuery
	*DescribeImageXHeifEncodeFileInSizeByTimeBody
}
type DescribeImageXHeifEncodeFileOutSizeByTimeReq struct {
	*DescribeImageXHeifEncodeFileOutSizeByTimeQuery
	*DescribeImageXHeifEncodeFileOutSizeByTimeBody
}
type DescribeImageXHeifEncodeSuccessCountByTimeReq struct {
	*DescribeImageXHeifEncodeSuccessCountByTimeQuery
	*DescribeImageXHeifEncodeSuccessCountByTimeBody
}
type DescribeImageXHeifEncodeSuccessRateByTimeReq struct {
	*DescribeImageXHeifEncodeSuccessRateByTimeQuery
	*DescribeImageXHeifEncodeSuccessRateByTimeBody
}
type DescribeImageXHitRateRequestDataReq struct {
	*DescribeImageXHitRateRequestDataQuery
	*DescribeImageXHitRateRequestDataBody
}
type DescribeImageXHitRateTrafficDataReq struct {
	*DescribeImageXHitRateTrafficDataQuery
	*DescribeImageXHitRateTrafficDataBody
}
type DescribeImageXMirrorRequestBandwidthReq struct {
	*DescribeImageXMirrorRequestBandwidthQuery
	*DescribeImageXMirrorRequestBandwidthBody
}
type DescribeImageXMirrorRequestHTTPCodeByTimeReq struct {
	*DescribeImageXMirrorRequestHTTPCodeByTimeQuery
	*DescribeImageXMirrorRequestHTTPCodeByTimeBody
}
type DescribeImageXMirrorRequestHTTPCodeOverviewReq struct {
	*DescribeImageXMirrorRequestHTTPCodeOverviewQuery
	*DescribeImageXMirrorRequestHTTPCodeOverviewBody
}
type DescribeImageXMirrorRequestTrafficReq struct {
	*DescribeImageXMirrorRequestTrafficQuery
	*DescribeImageXMirrorRequestTrafficBody
}
type DescribeImageXMultiCompressUsageReq struct {
	*DescribeImageXMultiCompressUsageQuery
	*DescribeImageXMultiCompressUsageBody
}
type DescribeImageXRequestCntUsageReq struct {
	*DescribeImageXRequestCntUsageQuery
	*DescribeImageXRequestCntUsageBody
}
type DescribeImageXScreenshotUsageReq struct {
	*DescribeImageXScreenshotUsageQuery
	*DescribeImageXScreenshotUsageBody
}
type DescribeImageXSensibleCacheHitRateByTimeReq struct {
	*DescribeImageXSensibleCacheHitRateByTimeQuery
	*DescribeImageXSensibleCacheHitRateByTimeBody
}
type DescribeImageXSensibleCountByTimeReq struct {
	*DescribeImageXSensibleCountByTimeQuery
	*DescribeImageXSensibleCountByTimeBody
}
type DescribeImageXSensibleTopRAMURLReq struct {
	*DescribeImageXSensibleTopRAMURLQuery
	*DescribeImageXSensibleTopRAMURLBody
}
type DescribeImageXSensibleTopResolutionURLReq struct {
	*DescribeImageXSensibleTopResolutionURLQuery
	*DescribeImageXSensibleTopResolutionURLBody
}
type DescribeImageXSensibleTopSizeURLReq struct {
	*DescribeImageXSensibleTopSizeURLQuery
	*DescribeImageXSensibleTopSizeURLBody
}
type DescribeImageXSensibleTopUnknownURLReq struct {
	*DescribeImageXSensibleTopUnknownURLQuery
	*DescribeImageXSensibleTopUnknownURLBody
}
type DescribeImageXServerQPSUsageReq struct {
	*DescribeImageXServerQPSUsageQuery
	*DescribeImageXServerQPSUsageBody
}
type DescribeImageXServiceQualityReq struct {
	*DescribeImageXServiceQualityQuery
	*DescribeImageXServiceQualityBody
}
type DescribeImageXSourceRequestReq struct {
	*DescribeImageXSourceRequestQuery
	*DescribeImageXSourceRequestBody
}
type DescribeImageXSourceRequestBandwidthReq struct {
	*DescribeImageXSourceRequestBandwidthQuery
	*DescribeImageXSourceRequestBandwidthBody
}
type DescribeImageXSourceRequestTrafficReq struct {
	*DescribeImageXSourceRequestTrafficQuery
	*DescribeImageXSourceRequestTrafficBody
}
type DescribeImageXStorageUsageReq struct {
	*DescribeImageXStorageUsageQuery
	*DescribeImageXStorageUsageBody
}
type DescribeImageXSummaryReq struct {
	*DescribeImageXSummaryQuery
	*DescribeImageXSummaryBody
}
type DescribeImageXUploadCountByTimeReq struct {
	*DescribeImageXUploadCountByTimeQuery
	*DescribeImageXUploadCountByTimeBody
}
type DescribeImageXUploadDurationReq struct {
	*DescribeImageXUploadDurationQuery
	*DescribeImageXUploadDurationBody
}
type DescribeImageXUploadErrorCodeAllReq struct {
	*DescribeImageXUploadErrorCodeAllQuery
	*DescribeImageXUploadErrorCodeAllBody
}
type DescribeImageXUploadErrorCodeByTimeReq struct {
	*DescribeImageXUploadErrorCodeByTimeQuery
	*DescribeImageXUploadErrorCodeByTimeBody
}
type DescribeImageXUploadFileSizeReq struct {
	*DescribeImageXUploadFileSizeQuery
	*DescribeImageXUploadFileSizeBody
}
type DescribeImageXUploadSegmentSpeedByTimeReq struct {
	*DescribeImageXUploadSegmentSpeedByTimeQuery
	*DescribeImageXUploadSegmentSpeedByTimeBody
}
type DescribeImageXUploadSpeedReq struct {
	*DescribeImageXUploadSpeedQuery
	*DescribeImageXUploadSpeedBody
}
type DescribeImageXUploadSuccessRateByTimeReq struct {
	*DescribeImageXUploadSuccessRateByTimeQuery
	*DescribeImageXUploadSuccessRateByTimeBody
}
type DescribeImageXVideoClipDurationUsageReq struct {
	*DescribeImageXVideoClipDurationUsageQuery
	*DescribeImageXVideoClipDurationUsageBody
}
type DownloadCertReq struct {
	*DownloadCertQuery
	*DownloadCertBody
}
type ExportFailedMigrateTaskReq struct {
	*ExportFailedMigrateTaskQuery
	*ExportFailedMigrateTaskBody
}
type FetchImageURLReq struct {
	*FetchImageURLQuery
	*FetchImageURLBody
}
type GetAllCertsReq struct {
	*GetAllCertsQuery
	*GetAllCertsBody
}
type GetAllImageServicesReq struct {
	*GetAllImageServicesQuery
	*GetAllImageServicesBody
}
type GetAllImageTemplatesReq struct {
	*GetAllImageTemplatesQuery
	*GetAllImageTemplatesBody
}
type GetAuditEntrysCountReq struct {
	*GetAuditEntrysCountQuery
	*GetAuditEntrysCountBody
}
type GetBatchProcessResultReq struct {
	*GetBatchProcessResultQuery
	*GetBatchProcessResultBody
}
type GetBatchTaskInfoReq struct {
	*GetBatchTaskInfoQuery
	*GetBatchTaskInfoBody
}
type GetCVAnimeGenerateImageReq struct {
	*GetCVAnimeGenerateImageQuery
	*GetCVAnimeGenerateImageBody
}
type GetCVImageGenerateResultReq struct {
	*GetCVImageGenerateResultQuery
	*GetCVImageGenerateResultBody
}
type GetCVImageGenerateTaskReq struct {
	*GetCVImageGenerateTaskQuery
	*GetCVImageGenerateTaskBody
}
type GetCVTextGenerateImageReq struct {
	*GetCVTextGenerateImageQuery
	*GetCVTextGenerateImageBody
}
type GetCertInfoReq struct {
	*GetCertInfoQuery
	*GetCertInfoBody
}
type GetComprehensiveEnhanceImageReq struct {
	*GetComprehensiveEnhanceImageQuery
	*GetComprehensiveEnhanceImageBody
}
type GetCompressTaskInfoReq struct {
	*GetCompressTaskInfoQuery
	*GetCompressTaskInfoBody
}
type GetDedupTaskStatusReq struct {
	*GetDedupTaskStatusQuery
	*GetDedupTaskStatusBody
}
type GetDenoisingImageReq struct {
	*GetDenoisingImageQuery
	*GetDenoisingImageBody
}
type GetDomainConfigReq struct {
	*GetDomainConfigQuery
	*GetDomainConfigBody
}
type GetDomainOwnerVerifyContentReq struct {
	*GetDomainOwnerVerifyContentQuery
	*GetDomainOwnerVerifyContentBody
}
type GetImageAIDetailsReq struct {
	*GetImageAIDetailsQuery
	*GetImageAIDetailsBody
}
type GetImageAIProcessQueuesReq struct {
	*GetImageAIProcessQueuesQuery
	*GetImageAIProcessQueuesBody
}
type GetImageAITasksReq struct {
	*GetImageAITasksQuery
	*GetImageAITasksBody
}
type GetImageAddOnTagReq struct {
	*GetImageAddOnTagQuery
	*GetImageAddOnTagBody
}
type GetImageAiGenerateTaskReq struct {
	*GetImageAiGenerateTaskQuery
	*GetImageAiGenerateTaskBody
}
type GetImageAlertRecordsReq struct {
	*GetImageAlertRecordsQuery
	*GetImageAlertRecordsBody
}
type GetImageAllDomainCertReq struct {
	*GetImageAllDomainCertQuery
	*GetImageAllDomainCertBody
}
type GetImageAnalyzeResultReq struct {
	*GetImageAnalyzeResultQuery
	*GetImageAnalyzeResultBody
}
type GetImageAnalyzeTasksReq struct {
	*GetImageAnalyzeTasksQuery
	*GetImageAnalyzeTasksBody
}
type GetImageAuditResultReq struct {
	*GetImageAuditResultQuery
	*GetImageAuditResultBody
}
type GetImageAuditTasksReq struct {
	*GetImageAuditTasksQuery
	*GetImageAuditTasksBody
}
type GetImageAuthKeyReq struct {
	*GetImageAuthKeyQuery
	*GetImageAuthKeyBody
}
type GetImageBackgroundColorsReq struct {
	*GetImageBackgroundColorsQuery
	*GetImageBackgroundColorsBody
}
type GetImageBgFillResultReq struct {
	*GetImageBgFillResultQuery
	*GetImageBgFillResultBody
}
type GetImageComicResultReq struct {
	*GetImageComicResultQuery
	*GetImageComicResultBody
}
type GetImageContentBlockListReq struct {
	*GetImageContentBlockListQuery
	*GetImageContentBlockListBody
}
type GetImageContentTaskDetailReq struct {
	*GetImageContentTaskDetailQuery
	*GetImageContentTaskDetailBody
}
type GetImageDetectResultReq struct {
	*GetImageDetectResultQuery
	*GetImageDetectResultBody
}
type GetImageDuplicateDetectionReq struct {
	*GetImageDuplicateDetectionQuery
	*GetImageDuplicateDetectionBody
}
type GetImageElementsReq struct {
	*GetImageElementsQuery
	*GetImageElementsBody
}
type GetImageEnhanceResultReq struct {
	*GetImageEnhanceResultQuery
	*GetImageEnhanceResultBody
}
type GetImageEraseModelsReq struct {
	*GetImageEraseModelsQuery
	*GetImageEraseModelsBody
}
type GetImageEraseResultReq struct {
	*GetImageEraseResultQuery
	*GetImageEraseResultBody
}
type GetImageFontsReq struct {
	*GetImageFontsQuery
	*GetImageFontsBody
}
type GetImageHmExtractTaskInfoReq struct {
	*GetImageHmExtractTaskInfoQuery
	*GetImageHmExtractTaskInfoBody
}
type GetImageMigrateTasksReq struct {
	*GetImageMigrateTasksQuery
	*GetImageMigrateTasksBody
}
type GetImageMonitorRulesReq struct {
	*GetImageMonitorRulesQuery
	*GetImageMonitorRulesBody
}
type GetImageOCRV2Req struct {
	*GetImageOCRV2Query
	*GetImageOCRV2Body
}
type GetImagePSDetectionReq struct {
	*GetImagePSDetectionQuery
	*GetImagePSDetectionBody
}
type GetImageQualityReq struct {
	*GetImageQualityQuery
	*GetImageQualityBody
}
type GetImageServiceReq struct {
	*GetImageServiceQuery
	*GetImageServiceBody
}
type GetImageServiceSubscriptionReq struct {
	*GetImageServiceSubscriptionQuery
	*GetImageServiceSubscriptionBody
}
type GetImageSettingRuleHistoryReq struct {
	*GetImageSettingRuleHistoryQuery
	*GetImageSettingRuleHistoryBody
}
type GetImageSettingRulesReq struct {
	*GetImageSettingRulesQuery
	*GetImageSettingRulesBody
}
type GetImageSettingsReq struct {
	*GetImageSettingsQuery
	*GetImageSettingsBody
}
type GetImageSmartCropResultReq struct {
	*GetImageSmartCropResultQuery
	*GetImageSmartCropResultBody
}
type GetImageStorageFilesReq struct {
	*GetImageStorageFilesQuery
	*GetImageStorageFilesBody
}
type GetImageStyleDetailReq struct {
	*GetImageStyleDetailQuery
	*GetImageStyleDetailBody
}
type GetImageStyleResultReq struct {
	*GetImageStyleResultQuery
	*GetImageStyleResultBody
}
type GetImageStylesReq struct {
	*GetImageStylesQuery
	*GetImageStylesBody
}
type GetImageSuperResolutionResultReq struct {
	*GetImageSuperResolutionResultQuery
	*GetImageSuperResolutionResultBody
}
type GetImageTemplateReq struct {
	*GetImageTemplateQuery
	*GetImageTemplateBody
}
type GetImageTranscodeDetailsReq struct {
	*GetImageTranscodeDetailsQuery
	*GetImageTranscodeDetailsBody
}
type GetImageTranscodeQueuesReq struct {
	*GetImageTranscodeQueuesQuery
	*GetImageTranscodeQueuesBody
}
type GetImageUpdateFilesReq struct {
	*GetImageUpdateFilesQuery
	*GetImageUpdateFilesBody
}
type GetImageUploadFileReq struct {
	*GetImageUploadFileQuery
	*GetImageUploadFileBody
}
type GetImageUploadFilesReq struct {
	*GetImageUploadFilesQuery
	*GetImageUploadFilesBody
}
type GetImageXQueryAppsReq struct {
	*GetImageXQueryAppsQuery
	*GetImageXQueryAppsBody
}
type GetImageXQueryDimsReq struct {
	*GetImageXQueryDimsQuery
	*GetImageXQueryDimsBody
}
type GetImageXQueryRegionsReq struct {
	*GetImageXQueryRegionsQuery
	*GetImageXQueryRegionsBody
}
type GetImageXQueryValsReq struct {
	*GetImageXQueryValsQuery
	*GetImageXQueryValsBody
}
type GetLicensePlateDetectionReq struct {
	*GetLicensePlateDetectionQuery
	*GetLicensePlateDetectionBody
}
type GetPrivateImageTypeReq struct {
	*GetPrivateImageTypeQuery
	*GetPrivateImageTypeBody
}
type GetProductAIGCResultReq struct {
	*GetProductAIGCResultQuery
	*GetProductAIGCResultBody
}
type GetResourceURLReq struct {
	*GetResourceURLQuery
	*GetResourceURLBody
}
type GetResponseHeaderValidateKeysReq struct {
	*GetResponseHeaderValidateKeysQuery
	*GetResponseHeaderValidateKeysBody
}
type GetSegmentImageReq struct {
	*GetSegmentImageQuery
	*GetSegmentImageBody
}
type GetServiceDomainsReq struct {
	*GetServiceDomainsQuery
	*GetServiceDomainsBody
}
type GetSyncAuditResultReq struct {
	*GetSyncAuditResultQuery
	*GetSyncAuditResultBody
}
type GetTemplatesFromBinReq struct {
	*GetTemplatesFromBinQuery
	*GetTemplatesFromBinBody
}
type GetURLFetchTaskReq struct {
	*GetURLFetchTaskQuery
	*GetURLFetchTaskBody
}
type GetVendorBucketsReq struct {
	*GetVendorBucketsQuery
	*GetVendorBucketsBody
}
type PreviewImageUploadFileReq struct {
	*PreviewImageUploadFileQuery
	*PreviewImageUploadFileBody
}
type ReportEventReq struct {
	*ReportEventQuery
	*ReportEventBody
}
type RerunImageMigrateTaskReq struct {
	*RerunImageMigrateTaskQuery
	*RerunImageMigrateTaskBody
}
type SetDefaultDomainReq struct {
	*SetDefaultDomainQuery
	*SetDefaultDomainBody
}
type TerminateImageMigrateTaskReq struct {
	*TerminateImageMigrateTaskQuery
	*TerminateImageMigrateTaskBody
}
type UpdateAdvanceReq struct {
	*UpdateAdvanceQuery
	*UpdateAdvanceBody
}
type UpdateAuditImageStatusReq struct {
	*UpdateAuditImageStatusQuery
	*UpdateAuditImageStatusBody
}
type UpdateDomainAdaptiveFmtReq struct {
	*UpdateDomainAdaptiveFmtQuery
	*UpdateDomainAdaptiveFmtBody
}
type UpdateFileStorageClassReq struct {
	*UpdateFileStorageClassQuery
	*UpdateFileStorageClassBody
}
type UpdateHTTPSReq struct {
	*UpdateHTTPSQuery
	*UpdateHTTPSBody
}
type UpdateImageAIProcessQueueReq struct {
	*UpdateImageAIProcessQueueQuery
	*UpdateImageAIProcessQueueBody
}
type UpdateImageAIProcessQueueStatusReq struct {
	*UpdateImageAIProcessQueueStatusQuery
	*UpdateImageAIProcessQueueStatusBody
}
type UpdateImageAnalyzeTaskReq struct {
	*UpdateImageAnalyzeTaskQuery
	*UpdateImageAnalyzeTaskBody
}
type UpdateImageAnalyzeTaskStatusReq struct {
	*UpdateImageAnalyzeTaskStatusQuery
	*UpdateImageAnalyzeTaskStatusBody
}
type UpdateImageAuditTaskReq struct {
	*UpdateImageAuditTaskQuery
	*UpdateImageAuditTaskBody
}
type UpdateImageAuditTaskStatusReq struct {
	*UpdateImageAuditTaskStatusQuery
	*UpdateImageAuditTaskStatusBody
}
type UpdateImageAuthKeyReq struct {
	*UpdateImageAuthKeyQuery
	*UpdateImageAuthKeyBody
}
type UpdateImageDomainAreaAccessReq struct {
	*UpdateImageDomainAreaAccessQuery
	*UpdateImageDomainAreaAccessBody
}
type UpdateImageDomainBandwidthLimitReq struct {
	*UpdateImageDomainBandwidthLimitQuery
	*UpdateImageDomainBandwidthLimitBody
}
type UpdateImageDomainConfigReq struct {
	*UpdateImageDomainConfigQuery
	*UpdateImageDomainConfigBody
}
type UpdateImageDomainDownloadSpeedLimitReq struct {
	*UpdateImageDomainDownloadSpeedLimitQuery
	*UpdateImageDomainDownloadSpeedLimitBody
}
type UpdateImageDomainIPAuthReq struct {
	*UpdateImageDomainIPAuthQuery
	*UpdateImageDomainIPAuthBody
}
type UpdateImageDomainUaAccessReq struct {
	*UpdateImageDomainUaAccessQuery
	*UpdateImageDomainUaAccessBody
}
type UpdateImageDomainVolcOriginReq struct {
	*UpdateImageDomainVolcOriginQuery
	*UpdateImageDomainVolcOriginBody
}
type UpdateImageExifDataReq struct {
	*UpdateImageExifDataQuery
	*UpdateImageExifDataBody
}
type UpdateImageFileCTReq struct {
	*UpdateImageFileCTQuery
	*UpdateImageFileCTBody
}
type UpdateImageFileKeyReq struct {
	*UpdateImageFileKeyQuery
	*UpdateImageFileKeyBody
}
type UpdateImageMirrorConfReq struct {
	*UpdateImageMirrorConfQuery
	*UpdateImageMirrorConfBody
}
type UpdateImageMonitorRuleReq struct {
	*UpdateImageMonitorRuleQuery
	*UpdateImageMonitorRuleBody
}
type UpdateImageMonitorRuleStatusReq struct {
	*UpdateImageMonitorRuleStatusQuery
	*UpdateImageMonitorRuleStatusBody
}
type UpdateImageObjectAccessReq struct {
	*UpdateImageObjectAccessQuery
	*UpdateImageObjectAccessBody
}
type UpdateImageResourceStatusReq struct {
	*UpdateImageResourceStatusQuery
	*UpdateImageResourceStatusBody
}
type UpdateImageSettingRuleReq struct {
	*UpdateImageSettingRuleQuery
	*UpdateImageSettingRuleBody
}
type UpdateImageSettingRulePriorityReq struct {
	*UpdateImageSettingRulePriorityQuery
	*UpdateImageSettingRulePriorityBody
}
type UpdateImageStorageTTLReq struct {
	*UpdateImageStorageTTLQuery
	*UpdateImageStorageTTLBody
}
type UpdateImageStyleReq struct {
	*UpdateImageStyleQuery
	*UpdateImageStyleBody
}
type UpdateImageStyleMetaReq struct {
	*UpdateImageStyleMetaQuery
	*UpdateImageStyleMetaBody
}
type UpdateImageTaskStrategyReq struct {
	*UpdateImageTaskStrategyQuery
	*UpdateImageTaskStrategyBody
}
type UpdateImageTranscodeQueueReq struct {
	*UpdateImageTranscodeQueueQuery
	*UpdateImageTranscodeQueueBody
}
type UpdateImageTranscodeQueueStatusReq struct {
	*UpdateImageTranscodeQueueStatusQuery
	*UpdateImageTranscodeQueueStatusBody
}
type UpdateImageUploadFilesReq struct {
	*UpdateImageUploadFilesQuery
	*UpdateImageUploadFilesBody
}
type UpdateImageUploadOverwriteReq struct {
	*UpdateImageUploadOverwriteQuery
	*UpdateImageUploadOverwriteBody
}
type UpdateReferReq struct {
	*UpdateReferQuery
	*UpdateReferBody
}
type UpdateResEventRuleReq struct {
	*UpdateResEventRuleQuery
	*UpdateResEventRuleBody
}
type UpdateResponseHeaderReq struct {
	*UpdateResponseHeaderQuery
	*UpdateResponseHeaderBody
}
type UpdateServiceNameReq struct {
	*UpdateServiceNameQuery
	*UpdateServiceNameBody
}
type UpdateSlimConfigReq struct {
	*UpdateSlimConfigQuery
	*UpdateSlimConfigBody
}
type UpdateStorageRulesReq struct {
	*UpdateStorageRulesQuery
	*UpdateStorageRulesBody
}
type UpdateStorageRulesV2Req struct {
	*UpdateStorageRulesV2Query
	*UpdateStorageRulesV2Body
}
type VerifyDomainOwnerReq struct {
	*VerifyDomainOwnerQuery
	*VerifyDomainOwnerBody
}

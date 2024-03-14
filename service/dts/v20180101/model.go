package dts_v20180101

type Enum0 string

type Enum5 string

type GetAsyncPreCheckResultResResultPreChecksItemLevel string

type GetAsyncPreCheckResultResResultStatus string

type ObjectMapping20221001MappingListItemObjectType string

type ObjectMapping20221001MappingListPropertiesItemsItem string

type ObjectMapping20221001ObjectMappingSettingObjectTransTypesItem string

type ObjectMapping20221001ObjectMappingSettingPolicyForKeyConflict string
type ObjectMapping20221001ObjectType string

type Components8K6CbSchemasObjectmapping20221001PropertiesMappinglistItemsPropertiesObjectmappingsettingPropertiesEsmetamappingsetting struct {
	EnableRouting *bool     `json:"EnableRouting,omitempty"`
	PidCol        []*string `json:"PidCol,omitempty"`
	RoutingCol    []*string `json:"RoutingCol,omitempty"`
}

type GetAsyncPreCheckResultBody struct {
	// REQUIRED; 预检查任务ID
	ID string `json:"ID"`
}

type GetAsyncPreCheckResultRes struct {
	ResponseMetadata *GetAsyncPreCheckResultResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *GetAsyncPreCheckResultResResult           `json:"Result,omitempty"`
}

type GetAsyncPreCheckResultResResponseMetadata struct {
	// REQUIRED; 请求的接口名，属于请求的公共参数。
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

type GetAsyncPreCheckResultResResult struct {
	// 是否通过
	Pass *bool `json:"Pass,omitempty"`

	// 预检查结果
	PreChecks []*GetAsyncPreCheckResultResResultPreChecksItem `json:"PreChecks,omitempty"`

	// 进度，如10
	Progress *int64                                 `json:"Progress,omitempty"`
	Status   *GetAsyncPreCheckResultResResultStatus `json:"Status,omitempty"`
}

type GetAsyncPreCheckResultResResultPreChecksItem struct {
	// 描述
	Desc *string `json:"Desc,omitempty"`

	// 详情
	Details *string                                            `json:"Details,omitempty"`
	Level   *GetAsyncPreCheckResultResResultPreChecksItemLevel `json:"Level,omitempty"`

	// 名称
	Name *string `json:"Name,omitempty"`

	// 是否通过
	Pass *bool `json:"Pass,omitempty"`
}

type ObjectMapping20221001 struct {
	DestObjName          *string                                    `json:"DestObjName,omitempty"`
	MappingList          []*ObjectMapping20221001MappingListItem    `json:"MappingList,omitempty"`
	ObjectMappingSetting *ObjectMapping20221001ObjectMappingSetting `json:"ObjectMappingSetting,omitempty"`
	ObjectType           *ObjectMapping20221001ObjectType           `json:"ObjectType,omitempty"`
	SrcObjName           *string                                    `json:"SrcObjName,omitempty"`
}

type ObjectMapping20221001MappingListItem struct {
	DestObjName          *string                                                   `json:"DestObjName,omitempty"`
	MappingList          []*ObjectMapping20221001                                  `json:"MappingList,omitempty"`
	ObjectMappingSetting *ObjectMapping20221001MappingListItemObjectMappingSetting `json:"ObjectMappingSetting,omitempty"`
	ObjectType           *ObjectMapping20221001MappingListItemObjectType           `json:"ObjectType,omitempty"`
	SrcObjName           *string                                                   `json:"SrcObjName,omitempty"`
}

type ObjectMapping20221001MappingListItemObjectMappingSetting struct {
	ESMetaMappingSetting *Components8K6CbSchemasObjectmapping20221001PropertiesMappinglistItemsPropertiesObjectmappingsettingPropertiesEsmetamappingsetting `json:"ESMetaMappingSetting,omitempty"`
	ObjectTransTypes     []*ObjectMapping20221001MappingListPropertiesItemsItem                                                                             `json:"ObjectTransTypes,omitempty"`
	PolicyForKeyConflict *Enum5                                                                                                                             `json:"PolicyForKeyConflict,omitempty"`

	// 向前兼容，默认false，表示所有操作类型
	SetObjectTransType *bool `json:"SetObjectTransType,omitempty"`
}

type ObjectMapping20221001ObjectMappingSetting struct {
	ESMetaMappingSetting *ObjectMapping20221001ObjectMappingSettingESMetaMappingSetting   `json:"ESMetaMappingSetting,omitempty"`
	ObjectTransTypes     []*ObjectMapping20221001ObjectMappingSettingObjectTransTypesItem `json:"ObjectTransTypes,omitempty"`
	PolicyForKeyConflict *ObjectMapping20221001ObjectMappingSettingPolicyForKeyConflict   `json:"PolicyForKeyConflict,omitempty"`

	// 向前兼容，默认false，表示所有操作类型
	SetObjectTransType *bool `json:"SetObjectTransType,omitempty"`
}

type ObjectMapping20221001ObjectMappingSettingESMetaMappingSetting struct {
	EnableRouting *bool     `json:"EnableRouting,omitempty"`
	PidCol        []*string `json:"PidCol,omitempty"`
	RoutingCol    []*string `json:"RoutingCol,omitempty"`
}

type PreCheckAsyncBody struct {
	// REQUIRED; 迁移ID
	TaskID string `json:"TaskID"`
}

type PreCheckAsyncRes struct {
	ResponseMetadata *PreCheckAsyncResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *PreCheckAsyncResResult           `json:"Result,omitempty"`
}

type PreCheckAsyncResResponseMetadata struct {
	// REQUIRED; 请求的接口名，属于请求的公共参数。
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

type PreCheckAsyncResResult struct {

	// 预检查结果id
	ID *string `json:"ID,omitempty"`
}
type GetAsyncPreCheckResult struct{}
type GetAsyncPreCheckResultQuery struct{}
type PreCheckAsync struct{}
type PreCheckAsyncQuery struct{}
type GetAsyncPreCheckResultReq struct {
	*GetAsyncPreCheckResultQuery
	*GetAsyncPreCheckResultBody
}
type PreCheckAsyncReq struct {
	*PreCheckAsyncQuery
	*PreCheckAsyncBody
}

package sms

import "github.com/volcengine/volc-sdk-golang/base"

type DocType = int

const (
	DocTypeThreeInOne                        DocType = 0  // 三证合一
	DocTypeBusinessLicense                   DocType = 1  // 企业营业执照
	DocTypeOrganizationCodeCertificate       DocType = 2  // 组织机构代码证
	DocTypeTaxRegistrationCertificate        DocType = 3  // 税务登记证
	DocTypeSocialCreditCodeCertificate       DocType = 4  // 社会信用代码证书
	DocTypePowerOfAttorney                   DocType = 5  // 授权委托书
	DocTypeOthers                            DocType = 6  // 其他/更多
	DocTypeInstitutionLegalPersonCertificate DocType = 7  // 事业单位法人证书
	DocTypeRepresentativeIDCardFront         DocType = 8  // 经办人身份证人像面
	DocTypeRepresentativeIDCardBack          DocType = 9  // 经办人身份证国徽面
	DocTypeResponsiblePersonIDCardFront      DocType = 10 // 责任人身份证人像面
	DocTypeResponsiblePersonIDCardBack       DocType = 11 // 责任人身份证国徽面
	DocTypePassportCard                      DocType = 12 // 护照照片
	DocTypeHKMPassportCard                   DocType = 13 // 港澳居民来往内地通行证照片
	DocTypeTWPassportCard                    DocType = 14 // 台湾居民来往大陆通行证照片
	DocTypeHMTResidenceCard                  DocType = 15 // 港澳台居民居住证照片
	DocTypeAppIcpCertificate                 DocType = 16 // APPICP证书
	DocTypeTrademarkCertificate              DocType = 17 // 商标证书
)

type SignSourceType = string

const (
	SignSourceTypeCompany          SignSourceType = "公司全称/简称"
	SignSourceTypeSite             SignSourceType = "工信部备案网站全称/简称" // deprecated
	SignSourceTypeApp              SignSourceType = "APP全称/简称"
	SignSourceTypeOfficialAccounts SignSourceType = "公众号、小程序全称/简称" // deprecated
	SignSourceTypeBrand            SignSourceType = "商标全称/简称"
	SignSourceTypeStore            SignSourceType = "电商平台店铺名的全称/简称" // deprecated
)

// 签名来源类型
const (
	SignSourceCompany = 1 // 公司全称/简称
	SignSourceApp     = 2 // APP全称
	SignSourceBrand   = 3 // 商标全称
)

type SignPurpose = int

const (
	SignPurposeForOwn   SignPurpose = 1
	SignPurposeForOther SignPurpose = 2
)

type CertificateType int32

const (
	IDCard CertificateType = 0 // 身份证
)

type GetSignatureAndOrderListRequest struct {
	SubAccounts []string `url:"subAccounts,omitempty"`
	Signature   string   `url:"signature,omitempty"`
	PageIndex   int      `url:"pageIndex"`
	PageSize    int      `url:"pageSize"`
	Status      int      `url:"status,omitempty"`
}

type GetSignatureAndOrderListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *struct {
		List  []*SmsSignatureInfo `json:"list"`
		Total int                 `json:"total"`
	}
}

type SmsSignatureInfo struct {
	Id      string `json:"id"`
	ApplyId string `json:"applyId"`

	Content            string         `json:"content"`
	CreatedTime        int64          `json:"createdTime"`
	Status             SmsOrderStatus `json:"status"`
	Source             string         `json:"source"`
	Reason             string         `json:"reason"`
	IsOrder            bool           `json:"isOrder"`
	Application        string         `json:"application"`
	OpenId             string         `json:"openId"`
	IdentificationID   int64          `json:"identificationID"`
	IdentificationName string         `json:"identificationName"`
	Purpose            int32          `json:"purpose"`
	UpdateStatus       int32          `json:"updateStatus"`
	UpdateReason       string         `json:"updateReason"`

	IsCommonSign bool         `json:"isCommonSign"`
	SubAccounts  []string     `json:"subAccounts"`
	ChannelTypes []string     `json:"channelTypes"`
	Industry     string       `json:"industry"`
	IndustryCN   string       `json:"industryCN"`
	Usable       bool         `json:"usable"`
	ReportStatus ReportStatus `json:"reportStatus"`
	AuditedAt    int64        `json:"auditedAt"`
}

type ReportStatus struct {
	Status int32 `json:"status"`
	CMCC   struct {
		Status int32  `json:"status"`
		Reason string `json:"reason"`
	} `json:"cmcc"`
	CT struct {
		Status int32  `json:"status"`
		Reason string `json:"reason"`
	} `json:"ct"`
	UniCom struct {
		Status int32  `json:"status"`
		Reason string `json:"reason"`
	} `json:"unicom"`
}

type ApplySmsSignatureRequest struct {
	SubAccount                string         `json:"subAccount"`
	Content                   string         `json:"content"`
	Source                    string         `json:"source"`
	Domain                    string         `json:"domain"`
	Desc                      string         `json:"desc,omitempty"`
	UploadFileList            []SignAuthFile `json:"uploadFileList"` // Deprecated
	Purpose                   int            `json:"purpose"`
	SignatureIdentificationID int64          `json:"signatureIdentificationID"` // 绑定的资质id
}

type ApplySmsSignatureRequestV2 struct {
	SubAccounts               []string       `json:"subAccounts"`
	Content                   string         `json:"content"`
	Source                    int            `json:"source"`
	Domain                    string         `json:"domain"`
	Desc                      string         `json:"desc,omitempty"`
	UploadFileList            []SignAuthFile `json:"uploadFileList"`
	Purpose                   int            `json:"purpose"`
	SignatureIdentificationID int64          `json:"signatureIdentificationID"` // 绑定的资质id
	AppIcp                    AppIcp         `json:"appIcp"`                    // app icp信息
	Trademark                 Trademark      `json:"trademark"`                 // 商标信息
	Scene                     string         `json:"scene"`                     // 业务场景
	ChannelTypes              []string       `json:"channelTypes"`
}

type UpdateSmsSignatureRequestV2 struct {
	SubAccounts               []string       `json:"subAccounts"`
	Content                   string         `json:"content"`
	Source                    int            `json:"source"`
	Domain                    string         `json:"domain"`
	Desc                      string         `json:"desc,omitempty"`
	UploadFileList            []SignAuthFile `json:"uploadFileList"`
	Purpose                   int            `json:"purpose"`
	SignatureIdentificationID int64          `json:"signatureIdentificationID"` // 绑定的资质id
	AppIcp                    AppIcp         `json:"appIcp"`                    // app icp信息
	Trademark                 Trademark      `json:"trademark"`                 // 商标信息
	Scene                     string         `json:"scene"`                     // 业务场景
	ChannelTypes              []string       `json:"channelTypes"`
	UpdateType                int            `json:"updateType"`
}

type AppIcp struct {
	AppIcpFilling  string         `json:"appIcpFilling"`
	AppIcpFileList []SignAuthFile `json:"appIcpFileList"` // AppIcp相关的文件信息
}

type Trademark struct {
	TrademarkCn       string         `json:"trademarkCn"`
	TrademarkEn       string         `json:"trademarkEn"`
	TrademarkNumber   string         `json:"trademarkNumber"`
	TrademarkFileList []SignAuthFile `json:"trademarkFileList"` // 商标相关的文件信息
}

type SignAuthFile struct {
	FileType    DocType `json:"fileType"`
	FileContent string  `json:"fileContent"` // 文件base64
	FileSuffix  string  `json:"fileSuffix"`
	FileUrl     string  `json:"fileUrl"` // 文件下载url
}

type ApplySmsSignatureResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SmsSignatureInfo
}

type UpdateSmsSignatureResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SmsSignatureInfo
}

type DeleteSignatureRequest struct {
	SubAccounts []string `json:"subAccounts"`
	Id          int64    `json:"id"`
	IsOrder     bool     `json:"isOrder"`
}

type DeleteSignatureResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

// 资质相关
type ApplySignatureIdentRequest struct {
	Id                    int64          `json:"id"`                    // 资质id，重新编辑需要提供之前的id
	Purpose               int32          `json:"purpose"`               // 资质用途； 1.自用，2.他用
	MaterialName          string         `json:"materialName"`          // 资质名称
	BusinessInfo          BusinessInfo   `json:"businessInfo"`          // 企业信息
	OperatorPersonInfo    PersonInfo     `json:"operatorPerson"`        // 经办人信息
	ResponsiblePersonInfo PersonInfo     `json:"responsiblePersonInfo"` // 责任人信息
	LegalPersonInfo       PersonInfo     `json:"legalPerson"`           // 法人信息
	PowerOfAttorney       []SignAuthFile `json:"powerOfAttorney"`       // 授权书
	OtherMaterials        []SignAuthFile `json:"otherMaterials"`        // 其他材料
	EffectSignatures      []string       `json:"effectSignatures"`      // 生效签名范围
}

type BusinessInfo struct {
	BusinessCertificateType                DocType      `json:"businessCertificateType"`                // 营业证件类型
	BusinessCertificate                    SignAuthFile `json:"businessCertificate"`                    // 营业证件
	BusinessCertificateName                string       `json:"businessCertificateName"`                // 营业证件名称
	UnifiedSocialCreditIdentifier          string       `json:"unifiedSocialCreditIdentifier"`          // 统一社会信用代码
	BusinessCertificateValidityPeriodStart string       `json:"businessCertificateValidityPeriodStart"` // 营业证件有效期开始
	BusinessCertificateValidityPeriodEnd   string       `json:"businessCertificateValidityPeriodEnd"`   // 营业证件有效期结束
	LegalPersonName                        string       `json:"legalPersonName"`                        // 法人名称
}

type PersonInfo struct {
	CertificateType   CertificateType `json:"certificateType"`   // 证件类型 0.身份证
	PersonCertificate []SignAuthFile  `json:"personCertificate"` // 证件信息
	PersonName        string          `json:"personName"`        // 名字
	PersonIDCard      string          `json:"personIDCard"`      // 证件号码
	PersonMobile      string          `json:"personMobile"`      // 手机号
}

type ApplySignatureIdentResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           struct {
		Id int64 `json:"id"`
	}
}

type GetSignatureIdentListRequest struct {
	Ids       []int64 `query:"ids" json:"ids"` // 资质id列表
	PageIndex int32   `query:"pageIndex" json:"pageIndex"`
	PageSize  int32   `query:"pageSize" json:"pageSize"`
}

type GetSignatureIdentListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SignatureIdentList
}

type SignatureIdentList struct {
	List []struct {
		Id                      int64    `json:"id"`
		Purpose                 int32    `json:"purpose"`                 // 资质用途； 1.自用，2.他用
		MaterialName            string   `json:"materialName"`            // 资质名称
		BusinessCertificateName string   `json:"businessCertificateName"` // 营业证件名称
		OperatorPersonName      string   `json:"operatorPersonName"`      // 经办人名字
		ResponsiblePersonName   string   `json:"responsiblePersonName"`   // 责任人名字
		LegalPersonName         string   `json:"legalPersonName"`         // 法人名字
		EffectSignatures        []string `json:"effectSignatures"`        // 生效签名范围
	} `json:"list"`
	Total int64 `json:"total"`
}

type BatchBindSignatureIdentRequest struct {
	SubAccount string   `json:"subAccount"` // 子账号
	Signatures []string `json:"signatures"` // 签名,可多个
	Id         int64    `json:"id"`         // 资质id
}

type BatchBindSignatureIdentResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           struct {
		Msg string `json:"msg"`
	}
}

type UpdateSignatureIdentificationRequest struct {
	Id                    int64      `json:"id"`
	OperatorPersonInfo    PersonInfo `json:"operatorPerson"`        // 经办人信息
	ResponsiblePersonInfo PersonInfo `json:"responsiblePersonInfo"` // 责任人信息
	LegalPersonInfo       PersonInfo `json:"legalPerson"`           // 法人信息
}

type UpdateSignatureIdentificationResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           struct {
		Msg string `json:"msg"`
	}
}

type TobTrafficDrivingPhoneNumberType int32

const (
	MobilePhone    TobTrafficDrivingPhoneNumberType = 1
	Number400      TobTrafficDrivingPhoneNumberType = 2
	Number700      TobTrafficDrivingPhoneNumberType = 3
	Number800      TobTrafficDrivingPhoneNumberType = 4
	Number95Or96   TobTrafficDrivingPhoneNumberType = 5
	Number12       TobTrafficDrivingPhoneNumberType = 6
	LandlineNumber TobTrafficDrivingPhoneNumberType = 7
)

type TobTrafficDrivingPhonePersonType int32

const (
	IdNumber TobTrafficDrivingPhonePersonType = 1
)

type CreateTobTrafficDrivingPhoneItem struct {
	NumberType     TobTrafficDrivingPhoneNumberType `json:"numberType"`
	Number         string                           `json:"number"`
	Company        string                           `json:"company"`
	NumberPerson   string                           `json:"numberPerson"`
	PersonType     TobTrafficDrivingPhonePersonType `json:"personType"`
	PersonId       string                           `json:"personId"`
	NumberProvince string                           `json:"numberProvince"`
	NumberCity     string                           `json:"numberCity"`
}

type BulkCreateTobTrafficDrivingPhoneRequest struct {
	Data []CreateTobTrafficDrivingPhoneItem `json:"data"`
}

type BulkCreateTobTrafficDrivingPhoneResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type GetTobTrafficDrivingPhoneListRequest struct {
	Number string `query:"number" json:"number"`
}

type GetTobTrafficDrivingPhoneListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           struct {
		List  []TobTrafficDrivingPhoneJson `json:"list"`
		Total int64                        `json:"total"`
	}
}

type TobTrafficDrivingPhoneJson struct {
	Account        string `json:"account"`
	NumberType     int32  `json:"numberType"`
	Number         string `json:"number"`
	Company        string `json:"company"`
	NumberPerson   string `json:"numberPerson"`
	PersonType     int32  `json:"personType"`
	PersonId       string `json:"personId"`
	NumberProvince string `json:"numberProvince"`
	NumberCity     string `json:"numberCity"`
	NumberEvidence string `json:"numberEvidence"`
	SoundInfo      string `json:"soundInfo"`
	ReviewStatus   int32  `json:"reviewStatus"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
}

type UpdateTobTrafficDrivingPhoneRequest struct {
	NumberType     TobTrafficDrivingPhoneNumberType `json:"numberType"`
	Number         string                           `json:"number"`
	Company        string                           `json:"company"`
	NumberPerson   string                           `json:"numberPerson"`
	PersonType     TobTrafficDrivingPhonePersonType `json:"personType"`
	PersonId       string                           `json:"personId"`
	NumberProvince string                           `json:"numberProvince"`
	NumberCity     string                           `json:"numberCity"`
	UpdatePersonId bool                             `json:"updatePersonId"`
}

type UpdateTobTrafficDrivingPhoneResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type DeleteTobTrafficDrivingPhoneRequest struct {
	Number string `json:"number"`
}

type DeleteTobTrafficDrivingPhoneResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type BulkCreateTobTrafficDrivingLinkRequest struct {
	Data []BulkCreateTobTrafficDrivingLinkItem `json:"data"`
}

type BulkCreateTobTrafficDrivingLinkItem struct {
	UseVolcLink           *bool  `json:"useVolcLink"`
	Protocol              string `json:"protocol"`
	Link                  string `json:"link"`
	LinkDomainIcp         string `json:"linkDomainIcp"`
	LinkDomainIcpBody     string `json:"linkDomainIcpBody"`
	JumpLink              string `json:"jumpLink"`
	JumpLinkDomainIcp     string `json:"jumpLinkDomainIcp"`
	JumpLinkDomainIcpBody string `json:"jumpLinkDomainIcpBody"`
}

type BulkCreateTobTrafficDrivingLinkResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type GetTobTrafficDrivingLinkListRequest struct {
	Link     string `query:"link"`
	JumpLink string `query:"jumpLink"`
}

type GetTobTrafficDrivingLinkListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           struct {
		List  []TobTrafficDrivingLinkJson `json:"list"`
		Total int64                       `json:"total"`
	}
}

type TobTrafficDrivingLinkJson struct {
	Account               string `json:"account"`
	UseVolcLink           bool   `json:"useVolcLink"`
	Protocol              string `json:"protocol"`
	Link                  string `json:"link"`
	LinkDomainIcp         string `json:"linkDomainIcp"`
	LinkDomainIcpBody     string `json:"linkDomainIcpBody"`
	LinkDomain            string `json:"linkDomain"`
	LinkPath              string `json:"linkPath"`
	JumpLink              string `json:"jumpLink"`
	JumpLinkDomain        string `json:"jumpLinkDomain"`
	JumpLinkPath          string `json:"jumpLinkPath"`
	JumpLinkDomainIcp     string `json:"jumpLinkDomainIcp"`
	JumpLinkDomainIcpBody string `json:"jumpLinkDomainIcpBody"`
	CreatedAt             int64  `json:"createdAt"`
	UpdatedAt             int64  `json:"updatedAt"`
}

type DeleteTobTrafficDrivingLinkRequest struct {
	Link     string `json:"link"`
	JumpLink string `json:"jumpLink"`
}

type DeleteTobTrafficDrivingLinkResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type ListRelationTemplateRequest struct {
	Signature string `query:"signature" json:"signature"`
	Number    string `query:"number" json:"number"`
	Link      string `query:"link" json:"link"`
	JumpLink  string `query:"jumpLink" json:"jumpLink"`
}

type ListRelationTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           struct {
		List  []TemplateV2 `json:"list"`
		Total int64        `json:"total"`
	}
}

type TemplateV2 struct {
	TemplateId        string             `json:"templateId"`
	Content           string             `json:"content"`
	ChannelType       string             `json:"channelType"`
	Signatures        []string           `json:"signatures"`
	SubAccounts       []string           `json:"subAccounts"`
	SubAccountNames   []string           `json:"subAccountNames"`
	Name              string             `json:"name"`
	TemplateParams    []TemplateParamsV2 `json:"templateParams"`
	FixTrafficDriving []TemplateParamsV2 `json:"fixTrafficDriving"`
	CreatedAt         int64              `json:"createdAt"`
	Status            int32              `json:"status"`
	CanDelete         bool               `json:"canDelete"`
	Account           string             `json:"account"`
	UseVolcLink       bool               `json:"useVolcLink"`
	ShortUrlConfig    ShortUrlConfig     `json:"shortUrlConfig"`
}

type TemplateParamsV2 struct {
	Name      string `query:"name" json:"name"`
	ParamType int32  `query:"paramType" json:"paramType"`
	Location  string `query:"location" json:"location"`
	Id        int64  `query:"id" json:"id"`
	Content   string `query:"content" json:"content"`
}
type ApplyTemplateV2Request struct {
	SubAccounts            []string                            `json:"subAccounts"`
	TemplateId             string                              `json:"templateId"`
	SecondTemplateId       string                              `json:"secondTemplateId"` // 二级文案ID
	Content                string                              `json:"content"`
	ChannelType            string                              `json:"channelType"`
	Area                   string                              `json:"area"`
	Name                   string                              `json:"name"`
	ShortUrlConfig         ShortUrlConfig                      `json:"shortUrlConfig"`
	Desc                   string                              `json:"desc"`
	Signatures             []string                            `json:"signatures"`
	TemplateParams         []TemplateParamsV2                  `json:"templateParams"`
	TemplateTrafficDriving [][]TemplateParamWithTrafficDriving `json:"templateTrafficDriving"`
}

type TemplateParamWithTrafficDriving struct {
	Name      string `query:"name" json:"name"`
	ParamType int    `query:"paramType" json:"paramType"`
	Content   string `query:"content" json:"content"`
}

type ApplyTemplateV2Response struct {
	ResponseMetadata base.ResponseMetadata
	Result           BindSignaturesResp `json:"Result"`
}

type BindSignaturesResp struct {
	TemplateId      string               `json:"templateId"`
	SecondTemplates []SecondTemplateResp `json:"secondTemplates"`
}

type SecondTemplateResp struct {
	SecondTemplateId string `json:"secondTemplateId"`
	Signature        string `json:"signature"`
	Industry         string `json:"industry"`
	IndustryCn       string `json:"industryCn"`
	ChannelType      string `json:"channelType"`
	Content          string `json:"content"`
}

type ListSmsTemplateV2Request struct {
	SubAccounts []string `query:"subAccounts" json:"subAccounts"`
	Page        int32    `query:"page" json:"page"`
	PageSize    int32    `query:"pageSize" json:"pageSize"`
	TemplateId  string   `query:"templateId" json:"templateId"`
	Signature   string   `query:"signature" json:"signature"`
	Name        string   `query:"name" json:"name"`
	Content     string   `query:"content" json:"content"`
}
type ListSmsTemplateV2Response struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListTemplateV2Result `json:"Result"`
}

type ListTemplateV2Result struct {
	Total int32        `json:"total"`
	List  []TemplateV2 `json:"list"`
}

type BindSignatureRequest struct {
	SubAccounts            []string                            `json:"subAccounts"`
	TemplateId             string                              `json:"templateId"`
	Signatures             []string                            `json:"signatures"`
	TemplateTrafficDriving [][]TemplateParamWithTrafficDriving `form:"templateTrafficDriving" json:"templateTrafficDriving"`
}

type BindSignatureResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           BindSignaturesResp `json:"Result"`
}

type ListSecondTemplateRequest struct {
	SubAccounts          []string `query:"subAccounts" json:"subAccounts"`
	TemplateId           string   `query:"templateId" json:"templateId"`
	TemplateIdList       []string `query:"templateIdList" json:"templateIdList"`
	SignatureList        []string `query:"signatureList" json:"signatureList"`
	SecondTemplateId     string   `query:"secondTemplateId" json:"secondTemplateId"`         // 二级模版名称-精确查询
	SecondTemplateIdList []string `query:"secondTemplateIdList" json:"secondTemplateIdList"` // 二级模版ID列表
	StatusList           []int32  `query:"statusList" json:"statusList"`                     // 状态列表
}

type ListSecondTemplateResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListSecondTemplateResult `json:"Result"`
}

type ListSecondTemplateResult struct {
	Total int32            `json:"total"`
	List  []SecondTemplate `json:"list"`
}

type SecondTemplate struct {
	TemplateId        string             `json:"templateId"`
	SecondTemplateId  string             `json:"secondTemplateId"`
	Content           string             `json:"content"`
	ChannelType       string             `json:"channelType"`
	Signature         string             `json:"signature"`
	SubAccounts       []string           `json:"subAccounts"`
	TemplateParams    []TemplateParamsV2 `json:"templateParams"`
	FixTrafficDriving []TemplateParamsV2 `json:"fixTrafficDriving"`
	CreatedAt         int64              `json:"createdAt"`
	UpdatedAt         int64              `json:"updatedAt"`
	ReviewStatus      int32              `json:"reviewStatus"` // 0 未送审；1. 审核中；2.审核失败；3.审核通过
	Industry          string             `json:"industry"`
	IndustryCn        string             `json:"industryCn"` // 行业中文名
	Name              string             `json:"name"`
	Account           string             `json:"account"`
	ShortUrlConfig    ShortUrlConfig     `json:"shortUrlConfig"`
}

type BindTrafficDrivingParamsRequest struct {
	SecondTemplateId       string                              `json:"secondTemplateId"`
	SubContentId           string                              `json:"subContentId"`
	TemplateTrafficDriving [][]TemplateParamWithTrafficDriving `json:"templateTrafficDriving" vd:"len($)>0"`
}

type BindTrafficDrivingParamsResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string `json:"Result"`
}

type ListSubContentRequest struct {
	SecondTemplateId string `query:"secondTemplateId" json:"secondTemplateId"`
	Page             int32  `query:"page" json:"page"`
	PageSize         int32  `query:"pageSize" json:"pageSize"`
	SubContentId     string `query:"subContentId" json:"subContentId"`
}

type ListSubContentResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListSubContentResult `json:"Result"`
}

type ListSubContentResult struct {
	Total int32        `json:"total"`
	List  []SubContent `json:"list"`
}

type SubContent struct {
	Content           string             `json:"content"`
	CreatedAt         int64              `json:"createdAt"`
	FixTrafficDriving []TemplateParamsV2 `json:"fixTrafficDriving"`
	TemplateParams    []TemplateParamsV2 `json:"templateParams"`
	SubContentId      string             `json:"subContentId"`
	Signature         string             `json:"signatures"`
	TemplateId        string             `json:"templateId"`
	SecondTemplateId  string             `json:"secondTemplateId"`
}

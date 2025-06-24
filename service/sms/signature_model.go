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
	SubAccount string `url:"subAccount"`
	Signature  string `url:"signature,omitempty"`
	PageIndex  int    `url:"pageIndex"`
	PageSize   int    `url:"pageSize"`
}

type GetSignatureAndOrderListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *struct {
		List  []*SmsSignatureInfo `json:"list"`
		Total int                 `json:"total"`
	}
}

type SmsSignatureInfo struct {
	Id          string         `json:"id"`
	ApplyId     string         `json:"applyId"`
	Content     string         `json:"content"`
	Source      string         `json:"source"`
	Application string         `json:"application"`
	CreatedTime int64          `json:"createdTime"`
	IsOrder     bool           `json:"isOrder"`
	Status      SmsOrderStatus `json:"status"`
	Reason      string         `json:"reason"`
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
	SubAccount                string         `json:"subAccount"`
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

type DeleteSignatureRequest struct {
	SubAccount string `json:"subAccount"`
	Id         string `json:"id"`
	IsOrder    bool   `json:"isOrder"`
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

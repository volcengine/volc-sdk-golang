package sms

import "github.com/volcengine/volc-sdk-golang/base"

type DocType = int

const (
	DocTypeThreeInOne                  DocType = 0 // 三证合一
	DocTypeBusinessLicense             DocType = 1 // 企业营业执照
	DocTypeOrganizationCodeCertificate DocType = 2 // 组织机构代码证
	DocTypeTaxRegistrationCertificate  DocType = 3 // 税务登记证
	DocTypeSocialCreditCodeCertificate DocType = 4 // 社会信用代码证书
	DocTypePowerOfAttorney             DocType = 5 // 授权委托书
	DocTypeOthers                      DocType = 6 // 其他/更多

)

type SignSourceType = string

const (
	SignSourceTypeCompany          SignSourceType = "公司全称/简称"
	SignSourceTypeSite             SignSourceType = "工信部备案网站全称/简称"
	SignSourceTypeApp              SignSourceType = "APP全称/简称"
	SignSourceTypeOfficialAccounts SignSourceType = "公众号、小程序全称/简称"
	SignSourceTypeBrand            SignSourceType = "商标全称/简称"
	SignSourceTypeStore            SignSourceType = "电商平台店铺名的全称/简称"
)

type SignPurpose = int

const (
	SignPurposeForOwn   SignPurpose = 1
	SignPurposeForOther SignPurpose = 2
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
	SubAccount     string         `json:"subAccount"`
	Content        string         `json:"content"`
	Source         string         `json:"source"`
	Domain         string         `json:"domain"`
	Desc           string         `json:"desc,omitempty"`
	UploadFileList []SignAuthFile `json:"uploadFileList"`
	Purpose        int            `json:"purpose"`
}

type SignAuthFile struct {
	FileType    DocType `json:"fileType"`
	FileContent string  `json:"fileContent"`
	FileSuffix  string  `json:"fileSuffix"`
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

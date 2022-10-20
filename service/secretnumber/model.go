package secretnumber

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type BindAXBRequest struct {
	PhoneNoA        string
	PhoneNoB        string
	PhoneNoX        string
	NumberPoolNo    string
	ExpireTime      int64
	AudioRecordFlag int32
	UserData        string
	OutId           string
}

type SelectNumberAndBindAXBRequest struct {
	PhoneNoA          string
	PhoneNoB          string
	NumberPoolNo      string
	ExpireTime        int64
	AudioRecordFlag   int32
	CityCode          string
	CityCodeByPhoneNo string
	DegradeCityList   string
	UserData          string
	RandomFlag        int32
	IdempotentId      string
	OutId             string
}

type SecretBindResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SecretBindResult
}

type SecretBindResult struct {
	PhoneNoX string
	PhoneNoE string
	SubId    string
	Status   int32
}

type OperationResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           bool
}

type SpecificSubIdRequest struct {
	NumberPoolNo string
	SubId        string
}

type QuerySubscriptionResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           Subscription
}

type QuerySubscriptionForListRequest struct {
	NumberPoolNo string
	PhoneNoX     string
	PhoneNoA     string
	PhoneNoB     string
	Status       int32
	SubId        string
	Offset       int32
	Limit        int32
	OutId        string
}

type QuerySubscriptionForListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           SubscriptionList
}

type SubscriptionList struct {
	Records []Subscription
	Offset  int32
	Limit   int32
	Total   int32
}

type Subscription struct {
	SubId      string
	PhoneNoA   string
	PhoneNoB   string
	PhoneNoX   string
	Status     int32
	RecordFlag int32
	EnableTime int64
	ExpireTime int64
	OutId      string
}

type UpgradeAXToAXBRequest struct {
	NumberPoolNo string
	SubId        string
	PhoneNoB     string
	UserData     string
	OutId        string
}

type UpdateAXBRequest struct {
	UpdateType   string
	NumberPoolNo string
	SubId        string
	ExpireTime   int64
	PhoneNoB     string
	UserData     string
	OutId        string
}

type BindAXNRequest struct {
	PhoneNoA        string
	PhoneNoB        string
	PhoneNoX        string
	NumberPoolNo    string
	ExpireTime      int64
	AudioRecordFlag int32
	UserData        string
}

type SelectNumberAndBindAXNRequest struct {
	PhoneNoA          string
	PhoneNoB          string
	NumberPoolNo      string
	ExpireTime        int64
	AudioRecordFlag   int32
	CityCode          string
	CityCodeByPhoneNo string
	DegradeCityList   string
	UserData          string
}

type UpdateAXNRequest struct {
	UpdateType   string
	NumberPoolNo string
	SubId        string
	ExpireTime   int64
	PhoneNoB     string
	PhoneNoA     string
}

type BindAXNERequest struct {
	PhoneNoA          string
	PhoneNoB          string
	PhoneNoX          string
	PhoneNoE          string
	NumberPoolNo      string
	ExpireTime        int64
	AudioRecordFlag   int32
	CityCode          string
	CityCodeByPhoneNo string
	DegradeCityList   string
	UserData          string
}

type UpdateAXNERequest struct {
	UpdateType   string
	NumberPoolNo string
	SubId        string
	ExpireTime   int64
	PhoneNoB     string
}

type Click2CallRequest struct {
	Caller                      string
	CallerNumber                string
	CallerNumberPoolNo          string
	CallerNumberCityCode        string
	CallerNumberDegradeCityList string
	Callee                      string
	CalleeNumber                string
	CalleeNumberPoolNo          string
	CalleeNumberCityCode        string
	CalleeNumberDegradeCityList string
	MaxTime                     int32
	PreVoice                    string
	PreVoicePlay                bool
	LastMinutes                 int32
	LastVoice                   string
	LastVoiceTo                 string
	UserData                    string
	CityCodeByPhoneNo           string
}

type Click2CallResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           Click2CallResult
}

type Click2CallResult struct {
	CallId string
}

type Click2CallLiteRequest struct {
	Caller          string
	Callee          string
	NumberPoolNo    string
	AudioRecordFlag int32
	UserData        string
}

type Click2CallLiteResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           Click2CallLiteResult
}

type Click2CallLiteResult struct {
	CallId string
}

type QueryAudioRecordFileUrlRequest struct {
	CallId string
}

type QueryAudioRecordFileUrlData struct {
	AudioRecordFileUrl      string
	AudioRecordLeftFileUrl  string
	AudioRecordRightFileUrl string
}

type QueryAudioRecordFileUrlResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           QueryAudioRecordFileUrlData
}

type QueryAudioRecordToTextFileRequest struct {
	CallIdList string
}

type QueryAudioRecordToTextFile struct {
	CallId                   string
	AudioRecordToTextFileUrl string
}

type QueryAudioRecordToTextFileResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []QueryAudioRecordToTextFile
}

type QueryCallRecordMsgRequest struct {
	CallIdList string
}

type QueryCallRecordMsg struct {
	AccountId              string
	CallId                 string
	ServiceType            int32
	SubServiceType         int32
	Caller                 string
	CallerCountryIsoCode   string
	CallerProvinceCode     string
	CallerCityCode         string
	Callee                 string
	CalleeCountryIsoCode   string
	CalleeProvinceCode     string
	CalleeCityCode         string
	BeginCallTime          string
	EndTime                string
	ReleaseType            int32
	CallDuration           int32
	CallResult             int32
	AudioRecordFlag        int32
	CdrCreateTime          string
	UserData               string
	CallType               int32
	CallerShowNumber       string
	CallerShowNumberPoolNo string
	CalleeShowNumber       string
	CalleeShowNumberPoolNo string
	CallerCallingTime      string
	CallerRingingTime      string
	CallerDuration         int32
}

type QueryCallRecordMsgResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           []QueryCallRecordMsg
}

type CreateNumberPoolRequest struct {
	Name           string
	ServiceType    string
	SubServiceType string
}

type CreateNumberPoolData struct {
	Name           string
	NumberPoolNo   string
	ServiceType    int32
	SubServiceType int32
}

type CreateNumberPoolResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           CreateNumberPoolData
}

type UpdateNumberPoolRequest struct {
	NumberPoolNo   string
	Name           string
	ServiceType    string
	SubServiceType string
}

type UpdateNumberPoolResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           bool
}

type NumberPoolListRequest struct {
	ServiceType    string
	SubServiceType string
	Limit          string
	Offset         string
}

type NumberPoolData struct {
	NumberPoolName     string
	NumberPoolNo       string
	ServiceType        int32
	ServiceTypeName    string
	SubServiceType     int32
	SubServiceTypeName string
	NumberCount        int32
}

type NumberPoolListPagedResponse struct {
	Records []NumberPoolData
	Limit   int32
	Offset  int32
	Total   int32
}

type NumberPoolListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           NumberPoolListPagedResponse
}

type NumberListRequest struct {
	Number             string
	NumberPoolNo       string
	NumberPoolTypeCode string
	NumberStatusCode   string
	NumberTypeCode     string
	Limit              string
	Offset             string
}

type NumberData struct {
	Number              string
	NumberStatusCode    int32
	NumberStatusDesc    string
	NumberTypeCode      int32
	NumberTypeDesc      string
	NumberLocation      string
	NumberPurchaseTime  string
	NumberPoolNo        string
	NumberPoolName      string
	NumberPoolTypeCode  int32
	NumberPoolTypeDesc  string
	ServiceTypeCode     int32
	ServiceTypeDesc     string
	QualificationNo     string
	QualificationEntity string
}

type NumberListPagedResponse struct {
	Records []NumberData
	Limit   int32
	Offset  int32
	Total   int32
}

type NumberListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           NumberListPagedResponse
}

type EnableOrDisableNumberRequest struct {
	NumberList string
	EnableCode string
}

type EnableOrDisableNumberResponse struct {
	ResponseMetadata base.ResponseMetadata
}

type QueryNumberApplyRecordListRequest struct {
	ApplyBillId          string
	QueryAccountId       string
	ApplyStatusCode      string
	ApplyTimeLowerBound  string
	ApplyTimeUpperBound  string
	SubServiceType       string
	NumberType           string
	UpdateTimeLowerBound string
	UpdateTimeUpperBound string
	Limit                string
	Offset               string
}

type NumberApplyRecordDetail struct {
	NumberLocation    string
	ApplyNumberCount  int32
	ImportNumberCount int32
}

type NumberApplyRecordData struct {
	Id                  int32
	ApplyTime           string
	ApplyStatusCode     int32
	ApplyStatusDesc     string
	SubServiceTypeCode  int32
	SubServiceTypeDesc  string
	NumberPoolNo        string
	NumberPoolName      string
	NumberTypeCode      int32
	NumberTypeDesc      string
	ApplyUserId         string
	ApplyUserName       string
	UpdateTime          string
	Notes               string
	DetailList          []NumberApplyRecordDetail
	ApplyBillId         string
	QualificationId     int32
	QualificationEntity string
}

type QueryNumberApplyRecordListPagedResponse struct {
	Records []NumberApplyRecordData
	Limit   int32
	Offset  int32
	Total   int32
}

type QueryNumberApplyRecordListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           QueryNumberApplyRecordListPagedResponse
}

type NumberApplicationCityItem struct {
	NumberCount    string
	CountryIsoCode string
	ProvinceCode   string
	CityCode       string
}

type CreateNumberApplicationRequest struct {
	QualificationId               string
	QualificationNo               string
	NumberPoolNo                  string
	NumberPurpose                 string
	NumberType                    string
	SubServiceType                string
	Remark                        string
	NumberApplicationCityItemList []NumberApplicationCityItem
}

type CreateNumberApplicationData struct {
	ApplyBillId string
}

type CreateNumberApplicationResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           CreateNumberApplicationData
}

type AddQualificationRequest struct {
	QualificationMainInfoFormDO         QualificationMainInfoFormDO
	QualificationAdminInfoFormDO        QualificationAdminInfoFormDO
	QualificationScenarioInfoFormDOList []QualificationScenarioInfoFormDO
}

type QualificationMainInfoFormDO struct {
	QualificationEntity                            string
	CertificateThreeInOne                          int32
	EnterpriseAddress                              string
	LegalRepresentativeName                        string
	LegalRepresentativeId                          string
	LegalRepresentativeFrontIdPhotoFileCode        string
	DocOfNumberApplyPhotoFileCode                  string
	CommitmentLetterOfNetAccessPhotoFileCode       string
	UnitSocialCreditCode                           string
	ThreeInOneBusinessLicensePhotoFileCode         string
	CodeOfOrganizationCertificate                  string
	BusinessLicensePhotoFileCode                   string
	CertificateOfOrganizationCodesPhotoFileCode    string
	CertificateOfTaxationRegistrationPhotoFileCode string
	QualificationNo                                string
}

type QualificationAdminInfoFormDO struct {
	Name                              string
	ContactNumber                     string
	IdCardNumber                      string
	IdCardFrontPhotoFileCode          string
	IdCardBackPhotoWithPeopleFileCode string
}

type QualificationScenarioInfoFormDO struct {
	QualificationNo   string
	SceneType         int32
	Description       string
	ScenarioOfCalling string
}

type AddQualificationResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           AddQualificationData
}

type AddQualificationData struct {
	QualificationNo string
}

type UpdateQualificationRequest struct {
	QualificationMainInfoFormDO         QualificationMainInfoFormDO
	QualificationAdminInfoFormDO        QualificationAdminInfoFormDO
	QualificationScenarioInfoFormDOList []QualificationScenarioInfoFormDO
}

type UpdateQualificationResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type AddQualificationSceneRequest struct {
	QualificationScenarioInfoFormDOList []QualificationScenarioInfoFormDO
}

type AddQualificationSceneResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type UpdateQualificationSceneRequest struct {
	QualificationScenarioInfoFormDOList []QualificationScenarioInfoFormDO
}

type UpdateQualificationSceneResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}

type QueryQualificationRequest struct {
	QualificationNoList []string
	ApprovalStatus      string
	Offset              int32
	Limit               int32
}

type QualificationMainInfoVO struct {
	QualificationId                           int32
	QualificationNo                           string
	QualificationEntity                       string
	AccountId                                 string
	CertificateThreeInOne                     int32
	EnterpriseAddress                         string
	LegalRepresentativeName                   string
	LegalRepresentativeId                     string
	LegalRepresentativeFrontIDPhotoURL        string
	DocOfNumberApplyPhotoURL                  string
	CommitmentLetterOfNetAccessPhotoURL       string
	UnitSocialCreditCode                      string
	ThreeInOneBusinessLicensePhotoURL         string
	CodeOfOrganizationCertificateURL          string
	BusinessLicensePhotoURL                   string
	CertificateOfOrganizationCodesPhotoURL    string
	CertificateOfTaxationRegistrationPhotoURL string
	ApprovalStatus                            int32
	CreateTime                                string
	UpdateTime                                string
}

type QualificationAdminInfoVO struct {
	Name                         string
	ContactNumber                string
	IdCardNumber                 string
	IdCardFrontPhotoURL          string
	IdCardBackPhotoWithPeopleURL string
}

type QualificationScenarioInfoVO struct {
	SceneType         int32
	Description       string
	ScenarioOfCalling string
}

type QueryQualificationVO struct {
	QualificationMainInfoVO         QualificationMainInfoVO
	QualificationAdminInfoVO        QualificationAdminInfoVO
	QualificationScenarioInfoVOList []QualificationScenarioInfoVO
}

type QueryQualificationData struct {
	Records []QueryQualificationVO
	Limit   int32
	Offset  int32
	Total   int32
}

type QueryQualificationResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           QueryQualificationData
}

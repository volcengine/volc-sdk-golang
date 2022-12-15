package iam

import "github.com/volcengine/volc-sdk-golang/base"

type NullResultResp struct {
	ResponseMetadata *base.ResponseMetadata
}

// User

type ThirdPartyPersonalIdentity struct {
	IDPType                 string
	ThirdPartyIdentityID    string
	ThirdPartyIdentityName  string
	ThirdPartyIdentityExtra string
}

type ThirdPartyEnterpriseIdentity struct {
	IDPType                string
	EnterpriseID           string
	EnterpriseName         string
	ThirdPartyIdentityID   string
	ThirdPartyIdentityName string
}

type ThirdPartyIdentity struct {
	Personal   []*ThirdPartyPersonalIdentity
	Enterprise []*ThirdPartyEnterpriseIdentity
}

type UserStructure struct {
	Id                  uint `json:",omitempty"`
	CreateDate          string
	UpdateDate          string
	Status              string
	AccountId           uint
	UserName            string
	Description         string
	DisplayName         string
	Email               string
	EmailIsVerify       bool
	MobilePhone         string
	MobilePhoneIsVerify bool
	Trn                 string
	Source              string
	ThirdPartyIdentity  *ThirdPartyIdentity `json:",omitempty"`
}

type UserList struct {
	UserMetadata []*UserStructure
	Limit        int
	Offset       int
	Total        int
}

type UserListResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *UserList `json:",omitempty"`
}

type LoginProfile struct {
	UserName              string
	LoginAllowed          bool
	PasswordResetRequired bool
	LastLoginDate         string
	LastLoginIp           string
}

type CreateLoginProfileResult struct {
	LoginProfile *LoginProfile
}

type CreateLoginProfileResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CreateLoginProfileResult `json:",omitempty"`
}

type GetLoginProfileResult struct {
	LoginProfile *LoginProfile
}

type GetLoginProfileResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetLoginProfileResult `json:",omitempty"`
}

type UpdateLoginProfileResult struct {
	LoginProfile *LoginProfile
}

type UpdateLoginProfileResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *UpdateLoginProfileResult `json:",omitempty"`
}

// Access Key

type AccessKeyMetadata struct {
	AccessKeyId string
	UserName    string
	CreateDate  string
	UpdateDate  string
	Status      string
}

type AccessKeyStructure struct {
	AccessKeyId     string
	SecretAccessKey string `json:",omitempty"`
	UserName        string
	CreateDate      string
	UpdateDate      string
	Status          string
}

type CreateAccessKeyResult struct {
	AccessKey AccessKeyStructure
}

type CreateAccessKeyResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CreateAccessKeyResult `json:",omitempty"`
}

type ListAccessKeysResult struct {
	AccessKeyMetadata []*AccessKeyStructure
}

type ListAccessKeysResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ListAccessKeysResult `json:",omitempty"`
}

// Policy

type PolicyStructure struct {
	PolicyTrn      string
	PolicyName     string
	PolicyType     string
	CreateDate     string
	UpdateDate     string
	Description    string
	PolicyDocument string
}

type Policy struct {
	Policy PolicyStructure
}

type PolicyResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *Policy `json:",omitempty"`
}

type PolicyList struct {
	PolicyMetadata []*PolicyStructure
	Limit          int
	Offset         int
	Total          int
}
type PolicyListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *PolicyList `json:",omitempty"`
}

type AttachedPolicyStructure struct {
	PolicyTrn   string
	PolicyName  string
	PolicyType  string
	CreateDate  string
	Description string
}

type AttachedPolicyList struct {
	AttachedPolicyMetadata []*AttachedPolicyStructure
	Limit                  int
	Offset                 int
	Total                  int
}

type AttachedPolicy struct {
	PolicyTrn   string
	PolicyName  string
	PolicyType  string
	AttachDate  string
	Description string
}

type ListAttachedPoliciesResult struct {
	AttachedPolicyMetadata []*AttachedPolicy
}

type ListAttachedUserPolicies struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ListAttachedPoliciesResult `json:",omitempty"`
}

type ListAttachedRolePolicies struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ListAttachedPoliciesResult `json:",omitempty"`
}

type AttachedPolicyListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *AttachedPolicyList `json:",omitempty"`
}

type ListEntitiesForPolicyPolicyUser struct {
	UserName    string
	DisplayName string
	AttachDate  string
	Description string
}

type ListEntitiesForPolicyPolicyRole struct {
	RoleName    string
	AttachDate  string
	Description string
}

type ListEntitiesForPolicyResult struct {
	PolicyUsers []*ListEntitiesForPolicyPolicyUser `json:",omitempty"`
	PolicyRoles []*ListEntitiesForPolicyPolicyRole `json:",omitempty"`
}

type ListEntitiesForPolicyResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ListEntitiesForPolicyResult `json:",omitempty"`
}

// Role

type RoleStructure struct {
	CreateDate          string
	UpdateDate          string
	Trn                 string
	RoleName            string
	DisplayName         string
	TrustPolicyDocument string `json:",omitempty"`
	Description         string
	IsServiceLinkedRole int
	MaxSessionDuration  int
}

type Role struct {
	Role RoleStructure
}

type RoleResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *Role `json:",omitempty"`
}

type RoLeList struct {
	RoleMetadata []*RoleStructure
	Limit        int
	Offset       int
	Total        int
}

type RoleListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *RoLeList `json:",omitempty"`
}

type User struct {
	User UserStructure
}

type UserResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *User `json:",omitempty"`
}

type AddedUserStructure struct {
	UserName   string
	CreateDate string
}

type AddedUserResp struct {
	AddedUser AddedUserStructure
}

type AddedUserList struct {
	AddedUserMetadata []*AddedUserStructure
	Limit             int
	Offset            int
	Total             int
}

type AddedUserListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *AddedUserList `json:",omitempty"`
}

// Identity Provider

type CreateSAMLProviderSAMLProvider struct {
	Trn              string
	SAMLProviderName string
	SSOType          int
	Status           int
	Description      string
	CreateDate       string
	UpdateDate       string
}

type GetSAMLProviderSAMLProvider struct {
	Trn                         string
	ProviderName                string
	EncodedSAMLMetadataDocument string
	SSOType                     int
	Status                      int
	Description                 string
	CreateDate                  string
	UpdateDate                  string
}

type ListSAMLProvidersSAMLProvider struct {
	Trn          string
	ProviderName string
	SSOType      int
	Status       int
	Description  string
	CreateDate   string
	UpdateDate   string
}

type UpdateSAMLProviderSAMLProvider struct {
	Trn              string
	SAMLProviderName string
	SSOType          int
	Status           int
	Description      string
	CreateDate       string
	UpdateDate       string
}

type CreateSAMLProviderResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CreateSAMLProviderSAMLProvider `json:",omitempty"`
}

type GetSAMLProviderResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetSAMLProviderSAMLProvider `json:",omitempty"`
}

type UpdateSAMLProviderResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *UpdateSAMLProviderSAMLProvider `json:",omitempty"`
}

type ListSAMLProviderResult struct {
	SAMLProviders []*ListSAMLProvidersSAMLProvider
	Limit         int
	Offset        int
	Total         int
}

type ListSAMLProvidersResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ListSAMLProviderResult `json:",omitempty"`
}

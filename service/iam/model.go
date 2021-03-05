package iam

import "github.com/volcengine/volc-sdk-golang/base"

// user
type UserStructure struct {
	base.BaseResp
	Id                  uint `json:",omitempty"`
	Trn                 string
	UserName            string
	Description         string
	DisplayName         string
	Email               string
	EmailIsVerify       bool
	MobilePhone         string
	MobilePhoneIsVerify bool
	Source              string
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

type NullResultResp struct {
	ResponseMetadata *base.ResponseMetadata
}

// accessKey
type AccessKeyStructure struct {
	base.BaseResp
	AccessKeyId     string
	SecretAccessKey string `json:",omitempty"`
	UserName        string
}

type AccessKey struct {
	AccessKey AccessKeyStructure
}

type AccessKeyResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *AccessKey `json:",omitempty"`
}

type AccessKeyList struct {
	AccessKeyMetadata []*AccessKeyStructure
}

type AccessKeyListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *AccessKeyList `json:",omitempty"`
}

// policy
type PolicyStructure struct {
	base.BaseResp
	PolicyTrn      string
	PolicyName     string
	PolicyType     string
	CreateDate     string
	UpdateDate     string
	Description    string
	PolicyId       string
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

type AttachedPolicyListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *AttachedPolicyList `json:",omitempty"`
}

// role
type RoleStructure struct {
	base.BaseResp
	Trn                 string
	RoleName            string
	RoleId              uint
	TrustPolicyDocument string `json:",omitempty"`
	Description         string
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

// service
type ServiceStructure struct {
	base.BaseResp
	ServiceName string
}

type Service struct {
	Service ServiceStructure `json:",omitempty"`
}
type ServiceResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *Service `json:",omitempty"`
}

// federation
type AppStructure struct {
	base.BaseResp
	AppId       string
	AppSecret   string
	AccountId   uint
	AppName     string
	IDP         string `json:",omitempty"`
	IDPType     string `json:",omitempty"`
	Description string
}

type App struct {
	App AppStructure
}

type AppResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *App `json:",omitempty"`
}

type AppList struct {
	AppMetadata []*AppStructure
	Limit       int
	Offset      int
	Total       int
}

type AppListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *AppList `json:",omitempty"`
}

// deprecated
type ListRolesAfterActorFilterResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *RoLeList `json:",omitempty"`
}

type IdentityProviderStruct struct {
	IdentityProviderName string
	IdentityProviderType string `json:",omitempty"`
	CreateDate           string `json:",omitempty"`
	Description          string
}

type IdentityProviderResp struct {
	IdentityProvider IdentityProviderStruct
}

type IdentityProviderList struct {
	IdentityProviderMetadata []*IdentityProviderStruct
	Limit                    int
	Offset                   int
	Total                    int
}

type IdentityProviderListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *IdentityProviderList `json:",omitempty"`
}

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

package iam

// ErrorObj 错误结构
type ErrorObj struct {
	Code    string
	Message string
}

// ResponseMetadata metadata结构
type ResponseMetadata struct {
	RequestID string
	Action    string
	Version   string
	Service   string    `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}

type BaseResp struct {
	Status     string
	CreateDate string
	UpdateDate string `json:",omitempty"`
}

// user
type UserStructure struct {
	BaseResp
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
	ResponseMetadata ResponseMetadata
	Result           *UserList `json:",omitempty"`
}

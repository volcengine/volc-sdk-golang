package iam

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20180101 = "2018-01-01"
	ServiceName            = "iam"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{

		// accessKey
		"ListAccessKeys": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAccessKeys"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateAccessKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateAccessKey"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateAccessKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAccessKey"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeleteAccessKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteAccessKey"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":       []string{"CreateService"},
				"Version":      []string{"2018-01-01"},
				"X-Account-Id": []string{"1"},
			},
		},
		"ListAccessKeysForService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":       []string{"ListAccessKeysForService"},
				"Version":      []string{"2018-01-01"},
				"X-Account-Id": []string{"1"},
			},
		},
		"CreateAccessKeyForService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":       []string{"CreateAccessKeyForService"},
				"Version":      []string{"2018-01-01"},
				"X-Account-Id": []string{"1"},
			},
		},
		"UpdateAccessKeyForService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":       []string{"UpdateAccessKeyForService"},
				"Version":      []string{"2018-01-01"},
				"X-Account-Id": []string{"1"},
			},
		},
		"DeleteAccessKeyForService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":       []string{"DeleteAccessKeyForService"},
				"Version":      []string{"2018-01-01"},
				"X-Account-Id": []string{"1"},
			},
		},

		// federation
		"AddAppIDToOAuthProvider": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddAppIDToOAuthProvider"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"RemoveAppIDFromOAuthProvider": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemoveAppIDFromOAuthProvider"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"GetAppIDofOAuthProvider": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAppIDofOAuthProvider"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateAppIDName": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAppIDName"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListAppIDsofOAuthProvider": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAppIDsofOAuthProvider"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateActorFilter": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateActorFilter"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListRolesAfterActorFilter": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRolesAfterActorFilter"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListIdentityProviders": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListIdentityProviders"},
				"Version": []string{ServiceVersion20180101},
			},
		},

		// policy
		"CreatePolicy": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePolicy"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"GetPolicy": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPolicy"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListPolicies": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPolicies"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeletePolicy": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePolicy"},
				"Version": []string{ServiceVersion20180101},
			},
		},

		// role
		"CreateRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"GetRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeleteRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListRoles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRoles"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"AttachRolePolicy": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AttachRolePolicy"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DetachRolePolicy": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetachRolePolicy"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListAttachedRolePolicies": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAttachedRolePolicies"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListUsersForRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUsersForRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"AddIdpToRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddIdpToRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"RemoveIDPFromRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemoveIDPFromRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListIDPsForRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListIDPsForRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},

		// User
		"CreateUser": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateUser"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"GetUser": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetUser"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListUsers": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUsers"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeleteUser": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteUser"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"AddUserToRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddUserToRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"RemoveUserFromRole": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemoveUserFromRole"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListRolesForUser": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRolesForUser"},
				"Version": []string{ServiceVersion20180101},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type IAM struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *IAM {
	instance := &IAM{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *IAM) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *IAM) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *IAM) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *IAM) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *IAM) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}

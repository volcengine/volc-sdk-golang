package iam

import (
	"encoding/json"
	"net/url"
)

// helper func
func (p *IAM) commonHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *IAM) ListAccessKeys(query url.Values) (*AccessKeyListResp, int, error) {
	respBody, status, err := p.Client.Query("ListAccessKeys", query)
	if err != nil {
		return nil, status, err
	}

	output := new(AccessKeyListResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *IAM) CreateAccessKey(query url.Values) (*AccessKeyResp, int, error) {
	respBody, status, err := p.Client.Query("CreateAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(AccessKeyResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *IAM) DeleteAccessKey(query url.Values) (*AccessKeyResp, int, error) {
	respBody, status, err := p.Client.Query("DeleteAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(AccessKeyResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *IAM) UpdateAccessKey(query url.Values) (*AccessKeyResp, int, error) {
	respBody, status, err := p.Client.Query("UpdateAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(AccessKeyResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

// role
func (p *IAM) CreateRole(query url.Values) (*RoleResp, int, error) {
	resp := new(RoleResp)
	statusCode, err := p.commonHandler("CreateRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) GetRole(query url.Values) (*RoleResp, int, error) {
	resp := new(RoleResp)
	statusCode, err := p.commonHandler("GetRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) DeleteRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DeleteRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListRoles(query url.Values) (*RoleListResp, int, error) {
	resp := new(RoleListResp)
	statusCode, err := p.commonHandler("ListRoles", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) UpdateRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("UpdateRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) AttachRolePolicy(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("AttachRolePolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) DetachRolePolicy(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DetachRolePolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListAttachedRolePolicies(query url.Values) (*AttachedPolicyListResp, int, error) {
	resp := new(AttachedPolicyListResp)
	statusCode, err := p.commonHandler("ListAttachedRolePolicies", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListUsersForRole(query url.Values) (*AddedUserListResp, int, error) {
	resp := new(AddedUserListResp)
	statusCode, err := p.commonHandler("ListUsersForRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// user
func (p *IAM) CreateUser(query url.Values) (*UserResp, int, error) {
	resp := new(UserResp)
	statusCode, err := p.commonHandler("CreateUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) GetUser(query url.Values) (*UserResp, int, error) {
	resp := new(UserResp)
	statusCode, err := p.commonHandler("GetUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) DeleteUser(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DeleteUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListUsers(query url.Values) (*UserListResp, int, error) {
	resp := new(UserListResp)
	statusCode, err := p.commonHandler("ListUsers", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) AddUserToRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("AddUserToRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) RemoveUserFromRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("RemoveUserFromRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListRolesForUser(query url.Values) (*RoleListResp, int, error) {
	resp := new(RoleListResp)
	statusCode, err := p.commonHandler("ListRolesForUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// policy
func (p *IAM) ListPolicies(query url.Values) (*PolicyListResp, int, error) {
	resp := new(PolicyListResp)
	statusCode, err := p.commonHandler("ListPolicies", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

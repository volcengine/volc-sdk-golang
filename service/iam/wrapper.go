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

// AccessKey

func (p *IAM) ListAccessKeys(query url.Values) (*ListAccessKeysResp, int, error) {
	respBody, status, err := p.Client.Query("ListAccessKeys", query)
	if err != nil {
		return nil, status, err
	}

	output := new(ListAccessKeysResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *IAM) CreateAccessKey(query url.Values) (*CreateAccessKeyResp, int, error) {
	respBody, status, err := p.Client.Query("CreateAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(CreateAccessKeyResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *IAM) DeleteAccessKey(query url.Values) (*NullResultResp, int, error) {
	respBody, status, err := p.Client.Query("DeleteAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(NullResultResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *IAM) UpdateAccessKey(query url.Values) (*NullResultResp, int, error) {
	respBody, status, err := p.Client.Query("UpdateAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(NullResultResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

// Role

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

func (p *IAM) CreateServiceLinkedRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("CreateServiceLinkedRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// Policy

func (p *IAM) CreatePolicy(query url.Values) (*PolicyResp, int, error) {
	resp := new(PolicyResp)
	statusCode, err := p.commonHandler("CreatePolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) DeletePolicy(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DeletePolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) AttachUserPolicy(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("AttachUserPolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) DetachUserPolicy(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DetachUserPolicy", query, resp)
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

func (p *IAM) ListAttachedUserPolicies(query url.Values) (*ListAttachedUserPolicies, int, error) {
	resp := new(ListAttachedUserPolicies)
	statusCode, err := p.commonHandler("ListAttachedUserPolicies", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListAttachedRolePolicies(query url.Values) (*ListAttachedRolePolicies, int, error) {
	resp := new(ListAttachedRolePolicies)
	statusCode, err := p.commonHandler("ListAttachedRolePolicies", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListPolicies(query url.Values) (*PolicyListResp, int, error) {
	resp := new(PolicyListResp)
	statusCode, err := p.commonHandler("ListPolicies", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) GetPolicy(query url.Values) (*PolicyResp, int, error) {
	resp := new(PolicyResp)
	statusCode, err := p.commonHandler("GetPolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) UpdatePolicy(query url.Values) (*PolicyResp, int, error) {
	resp := new(PolicyResp)
	statusCode, err := p.commonHandler("UpdatePolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) ListEntitiesForPolicy(query url.Values) (*ListEntitiesForPolicyResp, int, error) {
	resp := new(ListEntitiesForPolicyResp)
	statusCode, err := p.commonHandler("ListEntitiesForPolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// User

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

func (p *IAM) UpdateUser(query url.Values) (*UserResp, int, error) {
	resp := new(UserResp)
	statusCode, err := p.commonHandler("UpdateUser", query, resp)
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

func (p *IAM) CreateLoginProfile(query url.Values) (*CreateLoginProfileResp, int, error) {
	resp := new(CreateLoginProfileResp)
	statusCode, err := p.commonHandler("CreateLoginProfile", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) GetLoginProfile(query url.Values) (*GetLoginProfileResp, int, error) {
	resp := new(GetLoginProfileResp)
	statusCode, err := p.commonHandler("GetLoginProfile", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) UpdateLoginProfile(query url.Values) (*UpdateLoginProfileResp, int, error) {
	resp := new(UpdateLoginProfileResp)
	statusCode, err := p.commonHandler("UpdateLoginProfile", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) DeleteLoginProfile(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DeleteLoginProfile", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// Identity Provider

func (p *IAM) CreateSAMLProvider(query url.Values) (*CreateSAMLProviderResp, int, error) {
	resp := new(CreateSAMLProviderResp)
	respBody, statusCode, err := p.Client.Post("CreateSAMLProvider", url.Values{}, query)
	if err != nil {
		return nil, statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *IAM) GetSAMLProvider(query url.Values) (*GetSAMLProviderResp, int, error) {
	resp := new(GetSAMLProviderResp)
	statusCode, err := p.commonHandler("GetSAMLProvider", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) UpdateSAMLProvider(query url.Values) (*UpdateSAMLProviderResp, int, error) {
	resp := new(UpdateSAMLProviderResp)
	respBody, statusCode, err := p.Client.Post("UpdateSAMLProvider", url.Values{}, query)
	if err != nil {
		return nil, statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *IAM) ListSAMLProviders(query url.Values) (*ListSAMLProvidersResp, int, error) {
	resp := new(ListSAMLProvidersResp)
	statusCode, err := p.commonHandler("ListSAMLProviders", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *IAM) DeleteSAMLProvider(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DeleteSAMLProvider", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

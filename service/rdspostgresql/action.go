package rdspostgresql

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ActionCreateInstance      = "CreateDBInstance"
	ActionCreateIpWhiteList   = "CreateDBInstanceIPList"
	ActionCreateAccount       = "CreateAccount"
	ActionCreateDatabase      = "CreateDatabase"
	ActionModifyDatabaseOwner = "ModifyDatabaseOwner"
)

const (
	APIVersion20180101 = "2018-01-01"
)

// Create instance
func (p *RdsPostgresql) CreateInstance(req *CreateInstanceReq) (*CreateInstanceResp, error) {
	setCreateDefaultValues(req, false)
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateInstance: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionCreateInstance, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("CreateDBInstance: fail to do request, %v", err)
	}
	result := new(CreateInstanceResp)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Create instance
func (p *RdsPostgresql) CreateROInstance(req *CreateInstanceReq) (*CreateInstanceResp, error) {
	setCreateDefaultValues(req, true)
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateInstance: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionCreateInstance, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("CreateDBInstance: fail to do request, %v", err)
	}
	result := new(CreateInstanceResp)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func setCreateDefaultValues(req *CreateInstanceReq, isReadonly bool) {
	if isReadonly {
		req.InstanceCategory = "ReadOnly"
		req.InstanceType = "Basic"
	} else {
		req.InstanceCategory = "Primary"
		req.InstanceType = "HA"
	}
}

// Create instance IP white list
func (p *RdsPostgresql) CreateDBInstanceIPList(req *CreateDBInstanceIPListReq) (*base.CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateDBInstanceIPList: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionCreateIpWhiteList, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("CreateDBInstanceIPList: fail to do request, %v", err)
	}

	resp := new(base.CommonResponse)
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, fmt.Errorf("CreateDBInstanceIPList fail to unmarshal response, %v", err)
	}
	return resp, nil
}

// Create account
func (p *RdsPostgresql) CreateAccount(req *CreateAccountReq) (*base.CommonResponse, error) {
	fmt.Println(req)
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount: fail to marshal request, %v", err)
	}
	fmt.Println(string(reqData))
	respBody, _, err := p.Client.Json(ActionCreateAccount, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("CreateAccount: fail to do request, %v", err)
	}
	fmt.Println(string(respBody))
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, fmt.Errorf("CreateAccount fail to unmarshal response, %v", err)
	}
	return resp, nil
}

// Create account
func (p *RdsPostgresql) CreateDatabase(req *CreateDatabaseReq) (*base.CommonResponse, error) {
	fmt.Println(req)
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateDatabase: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionCreateDatabase, nil, string(reqData))
	if err != nil {
		fmt.Println(string(respBody))
		return nil, fmt.Errorf("CreateDatabase: fail to do request, %v", err)
	}
	fmt.Println(string(respBody))
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, fmt.Errorf("CreateDatabase fail to unmarshal response, %v", err)
	}
	return resp, nil
}

// Modify database owner
func (p *RdsPostgresql) ModifyDatabaseOwner(req *ModifyDatabaseOwnerReq) (*base.CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ModifyDatabaseOwner: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionModifyDatabaseOwner, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("ModifyDatabaseOwner: fail to do request, %v", err)
	}
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, fmt.Errorf("ModifyDatabaseOwner fail to unmarshal response, %v", err)
	}
	return resp, nil
}

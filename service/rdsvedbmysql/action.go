package rdsvedbmysql

import (
	"encoding/json"
	"fmt"
)

// Create instance
func (p *RdsVedbMysql) CreateDBInstance(req *CreateDBInstanceReq) (*CreateDBInstanceResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateDBInstance: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionCreateDBInstance, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("CreateDBInstance: fail to do request, %v", err)
	}
	result := new(CreateDBInstanceResp)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Describe Instance Detail
func (p *RdsVedbMysql) DescribeDBInstanceDetail(req *DescribeDBInstanceDetailReq) (*DescribeDBInstanceDetailResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DescribeDBInstanceDetail: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionDescribeDBInstanceDetail, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("DescribeDBInstanceDetail: fail to do request, %v", err)
	}
	result := new(DescribeDBInstanceDetailResp)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Describe Databases
func (p *RdsVedbMysql) DescribeDatabases(req *DescribeDatabasesReq) (*DescribeDatabasesResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DescribeDatabases: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionDescribeDatabases, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("DescribeDatabases: fail to do request, %v", err)
	}
	result := new(DescribeDatabasesResp)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Create Database
func (p *RdsVedbMysql) CreateDatabase(req *CreateDatabaseReq) (*CreateDatabaseReq, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateDatabase: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionCreateDatabase, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("CreateDatabase: fail to do request, %v", err)
	}
	result := new(CreateDatabaseReq)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Delete Database
func (p *RdsVedbMysql) DeleteDatabase(req *DeleteDatabaseReq) (*DeleteDatabaseResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteDatabase: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionDeleteDatabase, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("DeleteDatabase: fail to do request, %v", err)
	}
	result := new(DeleteDatabaseResp)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Create DB Account
func (p *RdsVedbMysql) CreateDBAccount(req *CreateDBAccountReq) (*CreateDBAccountResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateDBAccount: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionCreateDBAccount, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("CreateDBAccount: fail to do request, %v", err)
	}
	result := new(CreateDBAccountResp)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Delete DB Account
func (p *RdsVedbMysql) DeleteDBAccount(req *DeleteDBAccountReq) (*DeleteDBAccountResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteDBAccount: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionDeleteDBAccount, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("DeleteDBAccount: fail to do request, %v", err)
	}
	result := new(DeleteDBAccountResp)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Delete Database
func (p *RdsVedbMysql) DescribeDBAccounts(req *DescribeDBAccountsReq) (*DescribeDBAccountsResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DescribeDBAccounts: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json(ActionDescribeDBAccounts, nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("DescribeDBAccounts: fail to do request, %v", err)
	}
	result := new(DescribeDBAccountsResp)
	if err = UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

package rdsvedbmysql

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"
)

const (
	testAk = ""
	testSk = ""

	VpcId                = ""
	SubnetId             = ""
	SuperAccountPassword = ""
	InstanceId           = ""
	DBName               = ""
	PageNumber           = 1
	AccountName          = ""
	AccountPassword      = ""
)

const (
	DBEngineVersion           = "MySQL_8_0"
	NodeSpec                  = "vedb.mysql.x4.xlarge"
	NodeNumber                = 2
	SuperAccountName          = "vedbuser"
	ChargeType                = "PostPaid"
	ZoneIds                   = "cn-beijing-b"
	AccountPrivilegeReadWrite = "ReadWrite"
	AccountType               = "Normal"
)

func init() {
	_ = DefaultInstance.SetRegion(base.RegionCnNorth1)
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestCreateDBInstance(t *testing.T) {
	req := &CreateDBInstanceReq{
		DBEngineVersion:      DBEngineVersion,
		ZoneIds:              ZoneIds,
		NodeSpec:             NodeSpec,
		NodeNumber:           NodeNumber,
		VpcId:                VpcId,
		SubnetId:             SubnetId,
		SuperAccountName:     SuperAccountName,
		SuperAccountPassword: SuperAccountPassword,
		ChargeType:           ChargeType,
	}
	result, err := DefaultInstance.CreateDBInstance(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestDescribeDBInstanceDetail(t *testing.T) {
	req := &DescribeDBInstanceDetailReq{
		InstanceId: InstanceId,
	}
	result, err := DefaultInstance.DescribeDBInstanceDetail(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateDatabase(t *testing.T) {
	req := &CreateDatabaseReq{
		InstanceId: InstanceId,
		DBName:     DBName,
	}
	result, err := DefaultInstance.CreateDatabase(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestDescribeDatabases(t *testing.T) {
	req := &DescribeDatabasesReq{
		InstanceId: InstanceId,
		PageNumber: PageNumber,
	}
	result, err := DefaultInstance.DescribeDatabases(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestDeleteDatabase(t *testing.T) {
	req := &DeleteDatabaseReq{
		InstanceId: InstanceId,
		DBName:     DBName,
	}
	result, err := DefaultInstance.DeleteDatabase(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateDBAccount(t *testing.T) {
	req := &CreateDBAccountReq{
		InstanceId:      InstanceId,
		AccountName:     AccountName,
		AccountPassword: AccountPassword,
		AccountType:     AccountType,
		AccountPrivileges: []AccountPrivilegeObject{
			{
				DBName:           DBName,
				AccountPrivilege: AccountPrivilegeReadWrite,
			},
		},
	}
	result, err := DefaultInstance.CreateDBAccount(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestDescribeDBAccounts(t *testing.T) {
	req := &DescribeDBAccountsReq{
		InstanceId: InstanceId,
		PageNumber: PageNumber,
	}
	result, err := DefaultInstance.DescribeDBAccounts(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestDeleteDBAccount(t *testing.T) {
	req := &DeleteDBAccountReq{
		InstanceId:  InstanceId,
		AccountName: AccountName,
	}
	result, err := DefaultInstance.DeleteDBAccount(req)
	t.Logf("req = %+v\n", req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

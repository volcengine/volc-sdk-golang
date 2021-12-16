package rdspostgresql

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"
)

const (
	testAk = ""
	testSk = ""
)

var instanceID = "postgres-xxx"

func init() {
	_ = DefaultInstance.SetRegion(base.RegionCnNorth1)
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestCreateInstance(t *testing.T) {
	req := &CreateInstanceReq{
		Region:           "cn-north-x",
		Zone:             "cn-north-x",
		DBEngine:         "PostgreSQL",
		DBEngineVersion:  "PostgreSQL_12",     // PostgreSQL version: "PostgreSQL_12" | "PostgreSQL_11"
		InstanceType:     "HA",                // Primary instance type should be "HA"
		InstanceSpecName: "rds.postgres.1c2g", // 1c2g means 1CPU and 2Gi Memory
		StorageSpaceGB:   100,                 // G
		Number:           1,                   // create instance number
		NetworkType:      "VPC",               // Network type
		StorageType:      "LocalSSD",          // storage type: "LocalSSD"
		InstanceCategory: "Primary",           // Primary instance type should be "Primary"
		VpcID:            "vpc-xxx",
	}
	result, err := DefaultInstance.CreateInstance(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateROInstance(t *testing.T) {
	req := &CreateInstanceReq{
		Region:           "cn-north-x",
		Zone:             "cn-north-x-a",
		DBEngine:         "PostgreSQL",
		DBEngineVersion:  "PostgreSQL_12",     // PostgreSQL version: "PostgreSQL_12" | "PostgreSQL_11"
		InstanceType:     "Basic",             // readonly instance type should be "Basic"
		InstanceSpecName: "rds.postgres.1c2g", // 1c2g means 1CPU and 2Gi Memory
		StorageSpaceGB:   100,                 // G
		Number:           1,                   // create instance number
		NetworkType:      "VPC",               // Network type
		StorageType:      "LocalSSD",          // storage type: LocalSSD
		InstanceCategory: "ReadOnly",          // readonly instance type should be "ReadOnly"
		MasterInstanceId: instanceID,          // Primary instanceID
		VpcID:            "vpc-xxx",
	}
	result, err := DefaultInstance.CreateInstance(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateDBInstanceIPList(t *testing.T) {
	req := &CreateDBInstanceIPListReq{
		InstanceId: instanceID,
		GroupName:  "haha",
		IPList:     []string{"1.1.1.1"},
	}
	result, err := DefaultInstance.CreateDBInstanceIPList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateAccount(t *testing.T) {
	req := &CreateAccountReq{
		InstanceId:      instanceID,
		AccountName:     "user003",
		AccountPassword: "Admin@123",
		AccountType:     AccountTypeNormal,
	}
	result, err := DefaultInstance.CreateAccount(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateDatabase(t *testing.T) {
	req := &CreateDatabaseReq{
		InstanceId:       instanceID,
		DBName:           "database001",
		CharacterSetName: Charset_UTF8,
		Collate:          CollateType_EnUsUtf8,
		Ctype:            CType_C,
		Owner:            "user001",
	}
	result, err := DefaultInstance.CreateDatabase(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestModifyDatabaseOwner(t *testing.T) {
	req := &ModifyDatabaseOwnerReq{
		InstanceId: instanceID,
		DBName:     "database001",
		Owner:      "user002",
	}
	result, err := DefaultInstance.ModifyDatabaseOwner(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

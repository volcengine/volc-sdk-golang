package main

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"

	pg "github.com/volcengine/volc-sdk-golang/service/rdspostgresql"
)

const (
	testAk = ""
	testSk = ""
)

func TestCreateInstance(t *testing.T) {
	_ = pg.DefaultInstance.SetRegion(base.RegionCnNorth1)
	pg.DefaultInstance.Client.SetAccessKey(testAk)
	pg.DefaultInstance.Client.SetSecretKey(testSk)

	req := &pg.CreateInstanceReq{
		Region:           "cn-north-x",
		Zone:             "cn-north-xx",
		DBEngine:         "PostgreSQL",
		DBEngineVersion:  "PostgreSQL_12",     // PostgreSQL version: "PostgreSQL_12" | "PostgreSQL_11"
		InstanceType:     "HA",                // Primary instance type should be "HA"
		InstanceSpecName: "rds.postgres.1c2g", // 1c2g means 1CPU and 2Gi Memory
		StorageSpaceGB:   100,                 // G
		Number:           1,                   // create instance number
		StorageType:      "LocalSSD",          // storage type: "LocalSSD"
		InstanceCategory: "Primary",           // Primary instance type should be "Primary"
		VpcID:            "vpc-xxx",
		ChargeType:       "PostPaid",          // ChargeType should be "PostPaid"
		RequestSource:    "OpenAPI",           // RequestSource should be "OpenAPI"

	}
	result, err := pg.DefaultInstance.CreateInstance(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateROInstance(t *testing.T) {
	_ = pg.DefaultInstance.SetRegion(base.RegionCnNorth1)
	pg.DefaultInstance.Client.SetAccessKey(testAk)
	pg.DefaultInstance.Client.SetSecretKey(testSk)

	req := &pg.CreateInstanceReq{
		Region:           "cn-north-x",
		Zone:             "cn-north-x-x",
		DBEngine:         "PostgreSQL",
		DBEngineVersion:  "PostgreSQL_12",     // PostgreSQL version: "PostgreSQL_12" | "PostgreSQL_11"
		InstanceType:     "Basic",             // readonly instance type should be "Basic"
		InstanceSpecName: "rds.postgres.1c2g", // 1c2g means 1CPU and 2Gi Memory
		StorageSpaceGB:   100,                 // G
		Number:           1,                   // create instance number
		StorageType:      "LocalSSD",          // storage type: LocalSSD
		InstanceCategory: "ReadOnly",          // readonly instance type should be "ReadOnly"
		MasterInstanceId: "postgres-xxx",      // Primary instanceID
		VpcID:            "vpc-xxx",
		ChargeType:       "PostPaid",          // ChargeType should be "PostPaid"
		RequestSource:    "OpenAPI",           // RequestSource should be "OpenAPI"
	}
	result, err := pg.DefaultInstance.CreateInstance(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateDBInstanceIPList(t *testing.T) {
	_ = pg.DefaultInstance.SetRegion(base.RegionCnNorth1)
	pg.DefaultInstance.Client.SetAccessKey(testAk)
	pg.DefaultInstance.Client.SetSecretKey(testSk)

	req := &pg.CreateDBInstanceIPListReq{
		InstanceId: "postgres-xxx",
		GroupName:  "test001",
		IPList:     []string{"1.1.1.1"},
	}
	result, err := pg.DefaultInstance.CreateDBInstanceIPList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateAccount(t *testing.T) {
	_ = pg.DefaultInstance.SetRegion(base.RegionCnNorth1)
	pg.DefaultInstance.Client.SetAccessKey(testAk)
	pg.DefaultInstance.Client.SetSecretKey(testSk)

	req := &pg.CreateAccountReq{
		InstanceId:      "postgres-xxx",
		AccountName:     "user003",
		AccountPassword: "xxx",
		AccountType:     pg.AccountTypeNormal,
	}
	result, err := pg.DefaultInstance.CreateAccount(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestCreateDatabase(t *testing.T) {
	_ = pg.DefaultInstance.SetRegion(base.RegionCnNorth1)
	pg.DefaultInstance.Client.SetAccessKey(testAk)
	pg.DefaultInstance.Client.SetSecretKey(testSk)

	req := &pg.CreateDatabaseReq{
		InstanceId:       "postgres-xxx",
		DBName:           "db001",
		CharacterSetName: pg.Charset_UTF8,
		Collate:          pg.CollateType_EnUsUtf8,
		Ctype:            pg.CType_C,
		Owner:            "user001",
	}
	result, err := pg.DefaultInstance.CreateDatabase(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

func TestModifyDatabaseOwner(t *testing.T) {
	_ = pg.DefaultInstance.SetRegion(base.RegionCnNorth1)
	pg.DefaultInstance.Client.SetAccessKey(testAk)
	pg.DefaultInstance.Client.SetSecretKey(testSk)

	req := &pg.ModifyDatabaseOwnerReq{
		InstanceId: "postgres-xxx",
		DBName:     "db001",
		Owner:      "user002",
	}
	result, err := pg.DefaultInstance.ModifyDatabaseOwner(req)
	t.Logf("result = %+v\n", result)
	t.Logf("err = %+v\n", err)
}

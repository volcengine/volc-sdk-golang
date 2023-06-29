package rdsvedbmysql

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ActionCreateDBInstance         = "CreateDBInstance"
	ActionDescribeDBInstanceDetail = "DescribeDBInstanceDetail"
	ActionCreateDatabase           = "CreateDatabase"
	ActionDescribeDatabases        = "DescribeDatabases"
	ActionDeleteDatabase           = "DeleteDatabase"
	ActionCreateDBAccount          = "CreateDBAccount"
	ActionDescribeDBAccounts       = "DescribeDBAccounts"
	ActionDeleteDBAccount          = "DeleteDBAccount"
)

const (
	APIVersion20220101 = "2022-01-01"
	ServiceName        = "vedbm"
)

type CreateDBInstanceReq struct {
	DBEngineVersion      string `json:"DBEngineVersion"`
	ZoneIds              string `json:"ZoneIds"`
	NodeSpec             string `json:"NodeSpec"`
	NodeNumber           int    `json:"NodeNumber"`
	VpcId                string `json:"VpcId"`
	SubnetId             string `json:"SubnetId"`
	SuperAccountName     string `json:"SuperAccountName"`
	SuperAccountPassword string `json:"SuperAccountPassword"`
	ChargeType           string `json:"ChargeType"`
	InstanceName         string `json:"InstanceName,omitempty"`
	DBTimeZone           string `json:"DBTimeZone,omitempty"`
	LowerCaseTableNames  string `json:"LowerCaseTableNames,omitempty"`
	ProjectName          string `json:"ProjectName,omitempty"`
}

type CreateDBInstanceResp struct {
	InstanceId string
	OrderId    string
}

type DescribeDBInstanceDetailReq struct {
	InstanceId string `json:"InstanceId"`
}

type ChargeDetailObject struct {
	ChargeType         string
	ChargeStatus       string
	OverdueReclaimTime string
	OverdueTime        string
}

type NodeObject struct {
	Memory   int
	NodeId   string
	NodeSpec string
	NodeType string
	ZoneId   string
	VCPU     int
}

type EndpointObject struct {
	EndpointId               string
	AutoAddNewNodes          bool
	EndpointType             string
	Description              string
	Addresses                []AddressesObject
	MasterAcceptReadRequests bool
	NodeIds                  []string
	ReadWriteMode            string
	DistributedTransaction   bool
	EndpointName             string
}

type AddressesObject struct {
	DNSVisibility bool
	Domain        string
	IPAddress     string
	NetworkType   string
	Port          string
	SubnetId      string
	EiPId         string
}

type InstanceDetailObject struct {
	InstanceId          string
	InstanceName        string
	InstanceStatus      string
	RegionId            string
	ZoneIds             string
	DBEngineVersion     string
	CreateTime          string
	StorageUsedGiB      float64
	VpcId               string
	SubnetId            string
	TimeZone            string
	ProjectName         string
	LowerCaseTableNames string
}

type DescribeDBInstanceDetailResp struct {
	ChargeDetail   ChargeDetailObject
	Nodes          []NodeObject
	Endpoints      []EndpointObject
	InstanceDetail InstanceDetailObject
}

type DatabasesPrivilegeObject struct {
	AccountName      string `json:"AccountName"`
	AccountPrivilege string `json:"AccountPrivilege"`
}

type CreateDatabaseReq struct {
	InstanceId          string                     `json:"InstanceId"`
	DBName              string                     `json:"DBName"`
	CharacterSetName    string                     `json:"CharacterSetName,omitempty"`
	DatabasesPrivileges []DatabasesPrivilegeObject `json:"DatabasesPrivileges,omitempty"`
}

type CreateDatabaseResp struct {
}

type DescribeDatabasesReq struct {
	InstanceId string `json:"InstanceId"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize,omitempty"`
	DBName     string `json:"DBName,omitempty"`
}

type DatabaseObject struct {
	DBName              string
	CharacterSetName    string
	DatabasesPrivileges []DatabasesPrivilegeObject
}

type DescribeDatabasesResp struct {
	Total     int
	Databases []DatabaseObject
}

type DeleteDatabaseReq struct {
	InstanceId string `json:"InstanceId"`
	DBName     string `json:"DBName"`
}

type DeleteDatabaseResp struct {
}

type AccountPrivilegeObject struct {
	DBName           string
	AccountPrivilege string
}

type CreateDBAccountReq struct {
	InstanceId        string                   `json:"InstanceId"`
	AccountName       string                   `json:"AccountName"`
	AccountPassword   string                   `json:"AccountPassword"`
	AccountType       string                   `json:"AccountType"`
	AccountPrivileges []AccountPrivilegeObject `json:"AccountPrivileges,omitempty"`
}

type CreateDBAccountResp struct {
}

type DeleteDBAccountReq struct {
	InstanceId  string `json:"InstanceId"`
	AccountName string `json:"AccountName"`
}

type DeleteDBAccountResp struct {
}

type DescribeDBAccountsReq struct {
	InstanceId  string `json:"InstanceId"`
	PageNumber  int    `json:"PageNumber"`
	PageSize    int    `json:"PageSize,omitempty"`
	AccountName string `json:"AccountName,omitempty"`
}

type AccountObject struct {
	DBName              string
	CharacterSetName    string
	DatabasesPrivileges []DatabasesPrivilegeObject
}

type DescribeDBAccountsResp struct {
	Total    int
	Accounts []AccountObject
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return fmt.Errorf("fail to marshal result, %v", err)
	}
	if err = json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

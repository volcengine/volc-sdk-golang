package rdspostgresql

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
)

// createInstance
type CreateInstanceReq struct {
	Region           string `json:"Region"`
	Zone             string `json:"Zone"`
	ZoneSlave        string `json:"ZoneSlave"`
	DBEngine         string `json:"DBEngine"`
	DBEngineVersion  string `json:"DBEngineVersion"`
	InstanceType     string `json:"InstanceType"`
	StorageType      string `json:"StorageType"`
	StorageSpaceGB   int32  `json:"StorageSpaceGB"`
	Internal         bool   `json:"Internal"`
	ClusterName      string `json:"ClusterName"`
	NodePoolName     string `json:"NodePoolName"`
	Number           int32  `json:"Number"`
	InstanceSpecName string `json:"InstanceSpecName"`
	VpcID            string `json:"VpcID"`
	InstanceName     string `json:"InstanceName"`
	InstanceCategory string `json:"InstanceCategory"`
	MasterInstanceId string `json:"MasterInstanceId"`
	ChargeType       string `json:"ChargeType"`
	RequestSource    string `json:"RequestSource"`
}

type CreateInstanceResp struct {
	InstanceID string
}

type CreateDBInstanceIPListReq struct {
	InstanceId string   `json:"InstanceId"`
	GroupName  string   `json:"GroupName"`
	IPList     []string `json:"IPList"`
}

type SchemaInfo struct {
	SchemaName string `json:"SchemaName"`
	DBName     string `json:"DBName"`
	Owner      string `json:"Owner"`
}

type AccountType string

const (
	AccountTypeSuper  AccountType = "Super"
	AccountTypeNormal AccountType = "Normal"
)

type CreateAccountReq struct {
	InstanceId      string      `json:"InstanceId"`
	AccountName     string      `json:"AccountName"`
	AccountPassword string      `json:"AccountPassword"`
	AccountDesc     string      `json:"AccountDesc"`
	AccountType     AccountType `json:"AccountType"`
}

type CollateType string

const (
	CollateType_C        CollateType = "C"
	CollateType_CUTF8    CollateType = "CUTF8"
	CollateType_EnUsUtf8 CollateType = "EnUsUtf8"
)

type CType string

const (
	CType_C        CType = "C"
	CType_CUTF8    CType = "CUTF8"
	CType_EnUsUtf8 CType = "EnUsUtf8"
)

type Charset string

const (
	Charset_UTF8   Charset = "utf8"
	Charset_LATIN1 Charset = "latin1"
	Charset_ASCII  Charset = "ascii"
)

type CreateDatabaseReq struct {
	InstanceId       string      `json:"InstanceId"`
	DBName           string      `json:"DBName"`
	CharacterSetName Charset     `json:"CharacterSetName"`
	Collate          CollateType `json:"Collate"`
	Ctype            CType       `json:"Ctype"`
	Owner            string      `json:"Owner"`
}

type ModifyDatabaseOwnerReq struct {
	InstanceId string `json:"InstanceId"`
	Owner      string `json:"Owner"`
	DBName     string `json:"DBName"`
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

package rdspostgresql

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type RdsPostgresql struct {
	*base.Client
}

var DefaultInstance = NewInstance(base.RegionCnNorth1)

func NewInstance(region string) *RdsPostgresql {
	instance := &RdsPostgresql{
		Client: base.NewClient(ServiceInfoMap[region], ApiInfoList),
	}
	return instance
}

func (p *RdsPostgresql) SetRegion(region string) error {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		return fmt.Errorf("region does not spport or unknown region")
	}
	p.ServiceInfo = serviceInfo
	p.SetScheme("http")
	return nil
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "xxx.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "rds_postgresql"},
		},
		//base.RegionCnNorth3: {
		//	Timeout: 5 * time.Second,
		//	Host:    "xxx.volcengineapi.com",
		//	Header: http.Header{
		//		"Accept": []string{"application/json"},
		//	},
		//	Credentials: base.Credentials{Region: base.RegionCnNorth3, Service: "rds_postgresql"},
		//},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		ActionCreateInstance: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionCreateInstance},
				"Version": []string{APIVersion20180101},
			},
		},
		ActionCreateIpWhiteList: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionCreateIpWhiteList},
				"Version": []string{APIVersion20180101},
			},
		},
		ActionCreateAccount: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionCreateAccount},
				"Version": []string{APIVersion20180101},
			},
		},
		ActionCreateDatabase: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionCreateDatabase},
				"Version": []string{APIVersion20180101},
			},
		},
		ActionModifyDatabaseOwner: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionModifyDatabaseOwner},
				"Version": []string{APIVersion20180101},
			},
		},
	}
)

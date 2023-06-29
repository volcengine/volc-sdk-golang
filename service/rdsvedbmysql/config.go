package rdsvedbmysql

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type RdsVedbMysql struct {
	*base.Client
}

var DefaultInstance = NewInstance(base.RegionCnNorth1)

func NewInstance(region string) *RdsVedbMysql {
	instance := &RdsVedbMysql{
		Client: base.NewClient(ServiceInfoMap[region], ApiInfoList),
	}
	return instance
}

func (p *RdsVedbMysql) SetRegion(region string) error {
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
			Host:    "open.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: ServiceName},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		ActionCreateDBInstance: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionCreateDBInstance},
				"Version": []string{APIVersion20220101},
			},
		},
		ActionDescribeDBInstanceDetail: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionDescribeDBInstanceDetail},
				"Version": []string{APIVersion20220101},
			},
		},
		ActionCreateDatabase: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionCreateDatabase},
				"Version": []string{APIVersion20220101},
			},
		},
		ActionDescribeDatabases: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionDescribeDatabases},
				"Version": []string{APIVersion20220101},
			},
		},
		ActionDeleteDatabase: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionDeleteDatabase},
				"Version": []string{APIVersion20220101},
			},
		},
		ActionCreateDBAccount: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionCreateDBAccount},
				"Version": []string{APIVersion20220101},
			},
		},
		ActionDescribeDBAccounts: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionDescribeDBAccounts},
				"Version": []string{APIVersion20220101},
			},
		},
		ActionDeleteDBAccount: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionDeleteDBAccount},
				"Version": []string{APIVersion20220101},
			},
		},
	}
)

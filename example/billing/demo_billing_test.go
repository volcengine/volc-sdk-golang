package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/billing"
)

const (
	testAk = "<your access key id>"
	testSk = "<your access key secret>"
)

func TestListBill(t *testing.T) {
	billing.DefaultInstance.Client.SetAccessKey(testAk)
	billing.DefaultInstance.Client.SetSecretKey(testSk)
	//
	query := url.Values{}
	query.Add("BillPeriod", "2022-01")
	query.Add("Limit", "10")
	query.Add("Offset", "0")
	//
	query.Add("Product", "")
	query.Add("BillingMode", "")
	query.Add("BillCategoryParent", "")
	query.Add("PayStatus", "")
	query.Add("IgnoreZero", "0")
	query.Add("NeedRecordNum", "0")
	resp, status, err := billing.DefaultInstance.ListBill(query)
	assert.NoError(t, err)
	assert.Equal(t, status, http.StatusOK)
	assert.NotNil(t, resp)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestListBillDetail(t *testing.T) {
	billing.DefaultInstance.Client.SetAccessKey(testAk)
	billing.DefaultInstance.Client.SetSecretKey(testSk)
	//
	query := url.Values{}
	query.Add("BillPeriod", "2022-01")
	query.Add("Limit", "10")
	query.Add("Offset", "0")
	query.Add("GroupTerm", "0")
	query.Add("GroupPeriod", "2")
	//
	query.Add("Product", "")
	query.Add("BillingMode", "")
	query.Add("BillCategory", "")
	query.Add("InstanceNo", "")
	query.Add("IgnoreZero", "0")
	query.Add("NeedRecordNum", "0")
	resp, status, err := billing.DefaultInstance.ListBillDetail(query)
	assert.NoError(t, err)
	assert.Equal(t, status, http.StatusOK)
	assert.NotNil(t, resp)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestListListBillOverviewByProd(t *testing.T) {
	billing.DefaultInstance.Client.SetAccessKey(testAk)
	billing.DefaultInstance.Client.SetSecretKey(testSk)
	//
	query := url.Values{}
	query.Add("BillPeriod", "2022-01")
	query.Add("Limit", "10")
	query.Add("Offset", "0")
	//
	query.Add("Product", "")
	query.Add("BillingMode", "")
	query.Add("BillCategoryParent", "")
	query.Add("IgnoreZero", "0")
	query.Add("NeedRecordNum", "0")
	resp, status, err := billing.DefaultInstance.ListBillOverviewByProd(query)
	assert.NoError(t, err)
	assert.Equal(t, status, http.StatusOK)
	assert.NotNil(t, resp)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestListSplitBillDetail(t *testing.T) {
	billing.DefaultInstance.Client.SetAccessKey(testAk)
	billing.DefaultInstance.Client.SetSecretKey(testSk)

	//
	query := url.Values{}
	query.Add("GroupPeriod", "0")
	query.Add("BillPeriod", "2022-01")
	query.Add("Limit", "10")
	query.Add("Offset", "0")
	//
	query.Add("Product", "")
	query.Add("BillingMode", "")
	query.Add("BillCategory", "")
	query.Add("InstanceNo", "")
	query.Add("SplitItemID", "")
	query.Add("IgnoreZero", "0")
	query.Add("NeedRecordNum", "0")
	resp, status, err := billing.DefaultInstance.ListSplitBillDetail(query)
	assert.NoError(t, err)
	assert.Equal(t, status, http.StatusOK)
	assert.NotNil(t, resp)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

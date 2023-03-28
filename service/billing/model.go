package billing

import "github.com/volcengine/volc-sdk-golang/base"

type Bill struct {
	BillPeriod             string
	PayerID                string
	PayerUserName          string
	PayerCustomerName      string
	SellerID               string
	SellerUserName         string
	SellerCustomerName     string
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	Product                string
	ProductZh              string
	BusinessMode           string
	BillingMode            string
	ExpenseBeginTime       string
	ExpenseEndTime         string
	TradeTime              string
	BillID                 string
	BillCategoryParent     string
	OriginalBillAmount     string
	PreferentialBillAmount string
	RoundBillAmount        string
	DiscountBillAmount     string
	CouponAmount           string
	PayableAmount          string
	PaidAmount             string
	UnpaidAmount           string
	Currency               string
	PayStatus              string
	SettlementType         string
}

type BillList struct {
	List   []*Bill
	Total  int
	Limit  int
	Offset int
}

type BillListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BillList `json:",omitempty"`
}

type BillDetail struct {
	BillPeriod             string
	ExpenseDate            string
	PayerID                string
	PayerUserName          string
	PayerCustomerName      string
	SellerID               string
	SellerUserName         string
	SellerCustomerName     string
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	BusinessMode           string
	Product                string
	ProductZh              string
	BillingMode            string
	ExpenseBeginTime       string
	ExpenseEndTime         string
	UseDuration            string
	UseDurationUnit        string
	TradeTime              string
	BillID                 string
	BillCategory           string
	InstanceNo             string
	InstanceName           string
	ConfigName             string
	Element                string
	Region                 string
	Zone                   string
	Factor                 string
	ExpandField            string
	Price                  string
	PriceUnit              string
	Count                  string
	Unit                   string
	DeductionCount         string
	OriginalBillAmount     string
	PreferentialBillAmount string
	DiscountBillAmount     string
	CouponAmount           string
	PayableAmount          string
	PaidAmount             string
	UnpaidAmount           string
	Currency               string
	SettlementType         string
	Project                string
	Tag                    string
	SellingMode            string
	SolutionZh             string
}

type BillDetailList struct {
	List   []*BillDetail
	Total  int
	Limit  int
	Offset int
}

type BillDetailListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BillDetailList `json:",omitempty"`
}

type BillOverviewByProd struct {
	BillPeriod             string
	PayerID                string
	PayerUserName          string
	PayerCustomerName      string
	SellerID               string
	SellerUserName         string
	SellerCustomerName     string
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	BusinessMode           string
	Product                string
	ProductZh              string
	BillingMode            string
	BillCategoryParent     string
	OriginalBillAmount     string
	PreferentialBillAmount string
	RoundBillAmount        string
	DiscountBillAmount     string
	CouponAmount           string
	PayableAmount          string
	PaidAmount             string
	UnpaidAmount           string
	SettlementType         string
}

type BillOverviewByProdList struct {
	List   []*BillOverviewByProd
	Total  int
	Limit  int
	Offset int
}

type BillOverviewByProdListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BillOverviewByProdList `json:",omitempty"`
}

type SplitBillDetail struct {
	BillPeriod             string
	ExpenseTime            string
	PayerUserName          string
	SellerUserName         string
	OwnerUserName          string
	Product                string
	ProductZh              string
	BusinessMode           string
	BillingMode            string
	UseDuration            string
	UseDurationUnit        string
	TradeTime              string
	BillID                 string
	BillCategory           string
	SettlementType         string
	InstanceNo             string
	InstanceName           string
	ConfigName             string
	Element                string
	Region                 string
	Zone                   string
	Factor                 string
	ExpandField            string
	SplitItemID            string
	SplitItemName          string
	Price                  string
	PriceUnit              string
	SplitItemAmount        string
	Unit                   string
	SplitItemRatio         string
	DeductionCount         string
	SolutionZh             string
	OriginalBillAmount     string
	PreferentialBillAmount string
	DiscountBillAmount     string
	CouponDeductionAmount  string
	PayableAmount          string
	PaidAmount             string
	UnpaidAmount           string
	Currency               string
	Project                string
	Tag                    string
	SellingMode            string
	SubjectName            string
}

type SplitBillDetailList struct {
	List   []*SplitBillDetail
	Total  int
	Limit  int
	Offset int
}

type SplitBillDetailListResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *SplitBillDetailList `json:",omitempty"`
}

type UnsubscribeInstanceReq struct {
	Product                    string `json:"Product"`
	InstanceID                 string `json:"InstanceID"`
	UnsubscribeRelatedInstance bool   `json:"UnsubscribeRelatedInstance"`
	ClientToken                string `json:"ClientToken"`
}

type UnsubscribeInstanceResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *UnsubscribeInstanceResult `json:",omitempty"`
}

type UnsubscribeInstanceResult struct {
	OrderID              string                 `json:"OrderID"`
	SuccessInstanceInfos []*SuccessInstanceInfo `json:"SuccessInstanceInfos"`
}

type SuccessInstanceInfo struct {
	Product    string `json:"Product"`
	InstanceID string `json:"InstanceID"`
}

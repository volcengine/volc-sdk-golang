package billing

import "github.com/volcengine/volc-sdk-golang/base"

type Bill struct {
	BillPeriod             string
	PayerID                string
	PayerUserName          string
	PayerCustomerName      string
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	Product                string
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
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	BusinessMode           string
	Product                string
	BillingMode            string
	ExpenseBeginTime       string
	ExpenseEndTime         string
	UseDuration            string
	UseDurationUnit        string
	TradeTime              string
	BillID                 string
	BillCategory           string
	InstanceNo             string
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
	OwnerID                string
	OwnerUserName          string
	OwnerCustomerName      string
	BusinessMode           string
	Product                string
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

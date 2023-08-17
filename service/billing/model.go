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
	ReservationInstance    string
	BillDetailId           string
	ElementCode            string
	RegionCode             string
	ZoneCode               string
	FactorCode             string
	ConfigurationCode      string
	DeductionUseDuration   string
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
	ReservationInstance    string
	SplitBillDetailId      string
	ElementCode            string
	RegionCode             string
	ZoneCode               string
	FactorCode             string
	ConfigurationCode      string
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

type ListAmortizedCostBillDetailReq struct {
	BillPeriod     string `json:"BillPeriod"`
	AmortizedMonth string `json:"AmortizedMonth"`
	AmortizedDay   string `json:"AmortizedDay"`
	Product        string `json:"Product"`
	InstanceNo     string `json:"InstanceNo"`
	BillingMode    string `json:"BillingMode"`
	BillCategory   string `json:"BillCategory"`
	AmortizedType  string `json:"AmortizedType"`
	IgnoreZero     string `json:"IgnoreZero"`
	NeedRecordNum  string `json:"NeedRecordNum"`
	Offset         string `json:"Offset"`
	Limit          string `json:"Limit"`
}

type CostBillDetail struct {
	CostID                      string
	AmortizedMonth              string
	AmortizedDay                string
	BillPeriod                  string
	BusiPeriod                  string
	PayerID                     string
	PayerUserName               string
	PayerCustomerName           string
	SellerID                    string
	SellerUserName              string
	SellerCustomerName          string
	OwnerID                     string
	OwnerUserName               string
	OwnerCustomerName           string
	Product                     string
	ProductZh                   string
	BusinessMode                string
	BillingMode                 string
	BillCategory                string
	AmortizedType               string
	AmortizedBeginTime          string
	AmortizedEndTime            string
	BillID                      string
	InstanceNo                  string
	InstanceName                string
	Element                     string
	Region                      string
	Zone                        string
	Factor                      string
	ExpandField                 string
	ExpenseBeginTime            string
	ExpenseEndTime              string
	TradeTime                   string
	Price                       string
	PriceUnit                   string
	Count                       string
	Unit                        string
	UseDuration                 string
	UseDurationUnit             string
	CouponAmount                string
	PayableAmount               string
	DailyAmortizedCouponAmount  string
	DailyAmortizedPayableAmount string
	Currency                    string
	Project                     string
	Tag                         string
	SubjectName                 string
}

type CostBillDetailList struct {
	List   []*CostBillDetail
	Total  int
	Limit  int
	Offset int
}

type ListAmortizedCostBillDetailResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CostBillDetailList `json:",omitempty"`
}

type ListAmortizedCostBillMonthlyReq struct {
	BillPeriod     string `json:"BillPeriod"`
	AmortizedMonth string `json:"AmortizedMonth"`
	Product        string `json:"Product"`
	InstanceNo     string `json:"InstanceNo"`
	BillingMode    string `json:"BillingMode"`
	BillCategory   string `json:"BillCategory"`
	AmortizedType  string `json:"AmortizedType"`
	IgnoreZero     string `json:"IgnoreZero"`
	NeedRecordNum  string `json:"NeedRecordNum"`
	Offset         string `json:"Offset"`
	Limit          string `json:"Limit"`
}

type CostBillMonthly struct {
	AmortizedMonth               string
	BillPeriod                   string
	BusiPeriod                   string
	PayerID                      string
	PayerUserName                string
	PayerCustomerName            string
	SellerID                     string
	SellerUserName               string
	SellerCustomerName           string
	OwnerID                      string
	OwnerUserName                string
	OwnerCustomerName            string
	Product                      string
	ProductZh                    string
	BusinessMode                 string
	BillingMode                  string
	BillCategory                 string
	AmortizedType                string
	AmortizedBeginTime           string
	AmortizedEndTime             string
	AmortizedDayNum              string
	BillID                       string
	InstanceNo                   string
	InstanceName                 string
	Element                      string
	Region                       string
	Zone                         string
	Factor                       string
	ExpandField                  string
	Price                        string
	PriceUnit                    string
	Count                        string
	Unit                         string
	UseDuration                  string
	UseDurationUnit              string
	CouponAmount                 string
	PayableAmount                string
	DailyAmortizedCouponAmount   string
	DailyAmortizedPayableAmount  string
	BeforeAmortizedCouponAmount  string
	BeforeAmortizedPayableAmount string
	NowAmortizedCouponAmount     string
	NowAmortizedPayableAmount    string
	UnamortizedCouponAmount      string
	UnamortizedPayableAmount     string
	Currency                     string
	Project                      string
	Tag                          string
	SubjectName                  string
}

type CostBillMonthlyList struct {
	List   []*CostBillMonthly
	Total  int
	Limit  int
	Offset int
}

type ListAmortizedCostBillMonthlyResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CostBillMonthlyList `json:",omitempty"`
}

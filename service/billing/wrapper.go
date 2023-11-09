package billing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// ListBillDetail 分页查询账单明细
func (p *Billing) ListBillDetail(query url.Values) (*BillDetailListResp, int, error) {
	respBody, status, err := p.Client.Query("ListBillDetail", query)
	if err != nil {
		return nil, status, err
	}

	output := new(BillDetailListResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

// ListBill 分页查询账单
func (p *Billing) ListBill(query url.Values) (*BillListResp, int, error) {
	respBody, status, err := p.Client.Query("ListBill", query)
	if err != nil {
		return nil, status, err
	}

	output := new(BillListResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

// ListBillOverviewByProd 分页查询账单总览-产品汇总
func (p *Billing) ListBillOverviewByProd(query url.Values) (*BillOverviewByProdListResp, int, error) {
	respBody, status, err := p.Client.Query("ListBillOverviewByProd", query)
	if err != nil {
		return nil, status, err
	}

	output := new(BillOverviewByProdListResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

// ListBillOverviewByCategory 查询账单总览-汇总
func (p *Billing) ListBillOverviewByCategory(query url.Values) (*BillOverviewByCategoryMapResp, int, error) {
	respBody, status, err := p.Client.Query("ListBillOverviewByCategory", query)
	if err != nil {
		return nil, status, err
	}

	output := new(BillOverviewByCategoryMapResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

// ListSplitBillDetail 分页查询分账账单明细
func (p *Billing) ListSplitBillDetail(query url.Values) (*SplitBillDetailListResp, int, error) {
	respBody, status, err := p.Client.Query("ListSplitBillDetail", query)
	if err != nil {
		return nil, status, err
	}

	output := new(SplitBillDetailListResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

// UnsubscribeInstance 退订实例
func (p *Billing) UnsubscribeInstance(req *UnsubscribeInstanceReq) (*UnsubscribeInstanceResp, int, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("UnsubscribeInstance: fail to marshal request, %v", err)
	}

	respBody, status, err := p.Client.Json("UnsubscribeInstance", nil, string(reqData))
	if err != nil {
		return nil, status, err
	}

	output := new(UnsubscribeInstanceResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

// ListAmortizedCostBillDetail 查询成本账单明细
func (p *Billing) ListAmortizedCostBillDetail(req *ListAmortizedCostBillDetailReq) (*ListAmortizedCostBillDetailResp, int, error) {
	query := make(url.Values)
	reflectType := reflect.TypeOf(*req)
	reflectValue := reflect.ValueOf(*req)
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		tag := field.Tag.Get("json")
		value := reflectValue.Field(i).Interface()
		query.Add(tag, value.(string))
	}
	respBody, status, err := p.Client.Query("ListAmortizedCostBillDetail", query)
	if err != nil {
		return nil, status, err
	}

	output := new(ListAmortizedCostBillDetailResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

// ListAmortizedCostBillMonthly 查询成本账单总览
func (p *Billing) ListAmortizedCostBillMonthly(req *ListAmortizedCostBillMonthlyReq) (*ListAmortizedCostBillMonthlyResp, int, error) {
	query := make(url.Values)
	reflectType := reflect.TypeOf(*req)
	reflectValue := reflect.ValueOf(*req)
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		tag := field.Tag.Get("json")
		value := reflectValue.Field(i).Interface()
		query.Add(tag, value.(string))
	}
	respBody, status, err := p.Client.Query("ListAmortizedCostBillMonthly", query)
	if err != nil {
		return nil, status, err
	}

	output := new(ListAmortizedCostBillMonthlyResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = ServiceName
		return output, status, nil
	}
}

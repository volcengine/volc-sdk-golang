package billing

import (
	"encoding/json"
	"net/url"
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

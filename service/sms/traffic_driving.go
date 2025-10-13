package sms

func (p *SMS) BulkCreateTobTrafficDrivingLink(req *BulkCreateTobTrafficDrivingLinkRequest) (*BulkCreateTobTrafficDrivingLinkResponse, int, error) {
	resp := new(BulkCreateTobTrafficDrivingLinkResponse)
	statusCode, err := p.smsHandler("BulkCreateTobTrafficDrivingLink", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) BulkCreateTobTrafficDrivingPhone(req *BulkCreateTobTrafficDrivingPhoneRequest) (*BulkCreateTobTrafficDrivingPhoneResponse, int, error) {
	resp := new(BulkCreateTobTrafficDrivingPhoneResponse)
	statusCode, err := p.smsHandler("BulkCreateTobTrafficDrivingPhone", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) DeleteTobTrafficDrivingLink(req *DeleteTobTrafficDrivingLinkRequest) (*DeleteTobTrafficDrivingLinkResponse, int, error) {
	resp := new(DeleteTobTrafficDrivingLinkResponse)
	statusCode, err := p.smsHandler("DeleteTobTrafficDrivingLink", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) DeleteTobTrafficDrivingPhone(req *DeleteTobTrafficDrivingPhoneRequest) (*DeleteTobTrafficDrivingPhoneResponse, int, error) {
	resp := new(DeleteTobTrafficDrivingPhoneResponse)
	statusCode, err := p.smsHandler("DeleteTobTrafficDrivingPhone", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) GetRelationTemplateList(req *ListRelationTemplateRequest) (*ListRelationTemplateResponse, int, error) {
	resp := new(ListRelationTemplateResponse)
	statusCode, err := p.smsHandler("GetRelationTemplateList", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) GetTobTrafficDrivingLinkList(req *GetTobTrafficDrivingLinkListRequest) (*GetTobTrafficDrivingLinkListResponse, int, error) {
	resp := new(GetTobTrafficDrivingLinkListResponse)
	statusCode, err := p.smsHandler("GetTobTrafficDrivingLinkList", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) GetTobTrafficDrivingPhoneList(req *GetTobTrafficDrivingPhoneListRequest) (*GetTobTrafficDrivingPhoneListResponse, int, error) {
	resp := new(GetTobTrafficDrivingPhoneListResponse)
	statusCode, err := p.smsHandler("GetTobTrafficDrivingPhoneList", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) UpdateTobTrafficDrivingPhone(req *UpdateTobTrafficDrivingPhoneRequest) (*UpdateTobTrafficDrivingPhoneResponse, int, error) {
	resp := new(UpdateTobTrafficDrivingPhoneResponse)
	statusCode, err := p.smsHandler("UpdateTobTrafficDrivingPhone", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) BindTrafficDrivingParams(req *BindTrafficDrivingParamsRequest) (*BindTrafficDrivingParamsResponse, int, error) {
	resp := new(BindTrafficDrivingParamsResponse)
	statusCode, err := p.smsHandler("BindTrafficDrivingParams", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

package kms

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func (p *KMS) commonHandlerQuery(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err = json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *KMS) commonHandlerJson(api string, query url.Values, reqBody map[string]interface{}, resp interface{}) (int, error) {
	var statusCode int
	var respBody []byte
	body, err := json.Marshal(reqBody)
	if err != nil {
		return statusCode, err
	}

	respBody, statusCode, err = p.Client.Json(api, query, string(body))
	if err != nil {
		return statusCode, err
	}

	if err = json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *KMS) CreateKeyring(req *CreateKeyringRequest) (*CreateKeyringResponse, int, error) {
	query := url.Values{}
	resp := new(CreateKeyringResponse)

	query.Set("KeyringName", req.KeyringName)
	if req.KeyringType != nil {
		query.Set("KeyringType", *req.KeyringType)
	}
	if req.Description != nil {
		query.Set("Description", *req.Description)
	}

	statusCode, err := p.commonHandlerQuery("CreateKeyring", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DescribeKeyrings(req *DescribeKeyringsRequest) (*DescribeKeyringsResponse, int, error) {
	query := url.Values{}
	resp := new(DescribeKeyringsResponse)

	if req.CurrentPage != nil {
		query.Set("CurrentPage", strconv.Itoa(*req.CurrentPage))
	}
	if req.PageSize != nil {
		query.Set("PageSize", strconv.Itoa(*req.PageSize))
	}

	statusCode, err := p.commonHandlerQuery("DescribeKeyrings", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) UpdateKeyring(req *UpdateKeyringRequest) (*UpdateKeyringResponse, int, error) {
	query := url.Values{}
	resp := new(UpdateKeyringResponse)

	query.Set("KeyringName", req.KeyringName)
	if req.NewKeyringName != nil {
		query.Set("NewKeyringName", *req.NewKeyringName)
	}
	if req.Description != nil {
		query.Set("Description", *req.Description)
	}

	statusCode, err := p.commonHandlerQuery("UpdateKeyring", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) QueryKeyring(req *QueryKeyringRequest) (*QueryKeyringResponse, int, error) {
	query := url.Values{}
	resp := new(QueryKeyringResponse)

	query.Set("KeyringName", req.KeyringName)
	statusCode, err := p.commonHandlerQuery("QueryKeyring", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) CreateKey(req *CreateKeyRequest) (*CreateKeyResponse, int, error) {
	query := url.Values{}
	resp := new(CreateKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	if req.KeySpec != nil {
		query.Set("KeySpec", *req.KeySpec)
	}
	if req.KeyUsage != nil {
		query.Set("KeyUsage", *req.KeyUsage)
	}
	if req.ProtectionLevel != nil {
		query.Set("ProtectionLevel", *req.ProtectionLevel)
	}
	if req.Description != nil {
		query.Set("Description", *req.Description)
	}
	statusCode, err := p.commonHandlerQuery("CreateKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DescribeKeys(req *DescribeKeysRequest) (*DescribeKeysResponse, int, error) {
	query := url.Values{}
	resp := new(DescribeKeysResponse)

	query.Set("KeyringName", req.KeyringName)
	if req.CurrentPage != nil {
		query.Set("CurrentPage", strconv.Itoa(*req.CurrentPage))
	}
	if req.PageSize != nil {
		query.Set("PageSize", strconv.Itoa(*req.PageSize))
	}
	statusCode, err := p.commonHandlerQuery("DescribeKeys", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) UpdateKey(req *UpdateKeyRequest) (*UpdateKeyResponse, int, error) {
	query := url.Values{}
	resp := new(UpdateKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	if req.NewKeyName != nil {
		query.Set("NewKeyName", *req.NewKeyName)
	}
	if req.Description != nil {
		query.Set("Description", *req.Description)
	}

	statusCode, err := p.commonHandlerQuery("UpdateKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) GenerateDataKey(req *GenerateDataKeyRequest) (*GenerateDataKeyResponse, int, error) {
	query := url.Values{}
	resp := new(GenerateDataKeyResponse)


	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	if req.NumberOfBytes != nil {
		query.Set("NumberOfBytes ", strconv.Itoa(*req.NumberOfBytes))
	}

	reqBody := make(map[string]interface{})
	if req.EncryptionContext != nil {
		reqBody["EncryptionContext "] = req.EncryptionContext
	}

	statusCode, err := p.commonHandlerJson("GenerateDataKey", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) Encrypt(req *EncryptRequest) (*EncryptResponse, int, error) {
	query := url.Values{}
	resp := new(EncryptResponse)


	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	reqBody := make(map[string]interface{})
	if req.EncryptionContext != nil {
		reqBody["EncryptionContext "] = req.EncryptionContext
	}
	reqBody["Plaintext"] = req.Plaintext

	statusCode, err := p.commonHandlerJson("Encrypt", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) Decrypt(req *DecryptRequest) (*DecryptResponse, int, error) {
	query := url.Values{}
	resp := new(DecryptResponse)

	reqBody := make(map[string]interface{})
	if req.EncryptionContext != nil {
		reqBody["EncryptionContext "] = req.EncryptionContext
	}
	reqBody["CiphertextBlob"] = req.CiphertextBlob

	statusCode, err := p.commonHandlerJson("Decrypt", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) EnableKey(req *EnableKeyRequest) (*EnableKeyResponse, int, error) {
	query := url.Values{}
	resp := new(EnableKeyResponse)


	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("EnableKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DisableKey(req *DisableKeyRequest) (*DisableKeyResponse, int, error) {
	query := url.Values{}
	resp := new(DisableKeyResponse)


	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("DisableKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) ScheduleKeyDeletion(req *ScheduleKeyDeletionRequest) (*ScheduleKeyDeletionResponse, int, error) {
	query := url.Values{}
	resp := new(ScheduleKeyDeletionResponse)


	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	if req.PendingWindowInDays != nil {
		query.Set("PendingWindowInDays", strconv.Itoa(*req.PendingWindowInDays))
	}

	statusCode, err := p.commonHandlerQuery("ScheduleKeyDeletion", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) CancelKeyDeletion(req *CancelKeyDeletionRequest) (*CancelKeyDeletionResponse, int, error) {
	query := url.Values{}
	resp := new(CancelKeyDeletionResponse)


	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("CancelKeyDeletion", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

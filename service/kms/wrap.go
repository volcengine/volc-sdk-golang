package kms

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

func (p *KMS) commonHandlerQuery(ctx context.Context, api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.CtxQuery(ctx, api, query)
	if err != nil {
		return statusCode, err
	}

	if err = json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *KMS) commonHandlerJson(ctx context.Context, api string, query url.Values, reqBody map[string]interface{}, resp interface{}) (int, error) {
	var statusCode int
	var respBody []byte
	body, err := json.Marshal(reqBody)
	if err != nil {
		return statusCode, err
	}

	respBody, statusCode, err = p.Client.CtxJson(ctx, api, query, string(body))
	if err != nil {
		return statusCode, err
	}

	if err = json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *KMS) CtxCreateKeyring(ctx context.Context, req *CreateKeyringRequest) (*CreateKeyringResponse, int, error) {
	query := url.Values{}
	resp := new(CreateKeyringResponse)

	query.Set("KeyringName", req.KeyringName)
	if req.KeyringType != nil {
		query.Set("KeyringType", *req.KeyringType)
	}
	if req.Description != nil {
		query.Set("Description", *req.Description)
	}

	statusCode, err := p.commonHandlerQuery(ctx, "CreateKeyring", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) CreateKeyring(req *CreateKeyringRequest) (*CreateKeyringResponse, int, error) {
	return p.CtxCreateKeyring(context.Background(), req)
}

func (p *KMS) CtxDescribeKeyrings(ctx context.Context, req *DescribeKeyringsRequest) (*DescribeKeyringsResponse, int, error) {
	query := url.Values{}
	resp := new(DescribeKeyringsResponse)

	if req.CurrentPage != nil {
		query.Set("CurrentPage", strconv.Itoa(*req.CurrentPage))
	}
	if req.PageSize != nil {
		query.Set("PageSize", strconv.Itoa(*req.PageSize))
	}

	statusCode, err := p.commonHandlerQuery(ctx, "DescribeKeyrings", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DescribeKeyrings(req *DescribeKeyringsRequest) (*DescribeKeyringsResponse, int, error) {
	return p.CtxDescribeKeyrings(context.Background(), req)
}

func (p *KMS) CtxUpdateKeyring(ctx context.Context, req *UpdateKeyringRequest) (*UpdateKeyringResponse, int, error) {
	query := url.Values{}
	resp := new(UpdateKeyringResponse)

	query.Set("KeyringName", req.KeyringName)
	if req.NewKeyringName != nil {
		query.Set("NewKeyringName", *req.NewKeyringName)
	}
	if req.Description != nil {
		query.Set("Description", *req.Description)
	}

	statusCode, err := p.commonHandlerQuery(ctx, "UpdateKeyring", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) UpdateKeyring(req *UpdateKeyringRequest) (*UpdateKeyringResponse, int, error) {
	return p.CtxUpdateKeyring(context.Background(), req)
}

func (p *KMS) CtxQueryKeyring(ctx context.Context, req *QueryKeyringRequest) (*QueryKeyringResponse, int, error) {
	query := url.Values{}
	resp := new(QueryKeyringResponse)

	query.Set("KeyringName", req.KeyringName)
	statusCode, err := p.commonHandlerQuery(ctx, "QueryKeyring", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) QueryKeyring(req *QueryKeyringRequest) (*QueryKeyringResponse, int, error) {
	return p.CtxQueryKeyring(context.Background(), req)
}

func (p *KMS) CtxCreateKey(ctx context.Context, req *CreateKeyRequest) (*CreateKeyResponse, int, error) {
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
	statusCode, err := p.commonHandlerQuery(ctx, "CreateKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}
func (p *KMS) CreateKey(req *CreateKeyRequest) (*CreateKeyResponse, int, error) {
	return p.CtxCreateKey(context.Background(), req)
}

func (p *KMS) CtxDescribeKeys(ctx context.Context, req *DescribeKeysRequest) (*DescribeKeysResponse, int, error) {
	query := url.Values{}
	resp := new(DescribeKeysResponse)

	query.Set("KeyringName", req.KeyringName)
	if req.CurrentPage != nil {
		query.Set("CurrentPage", strconv.Itoa(*req.CurrentPage))
	}
	if req.PageSize != nil {
		query.Set("PageSize", strconv.Itoa(*req.PageSize))
	}
	statusCode, err := p.commonHandlerQuery(ctx, "DescribeKeys", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DescribeKey(req *DescribeKeyRequest) (*DescribeKeyResponse, int, error) {
	query := url.Values{}
	resp := new(DescribeKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("DescribeKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DescribeKeys(req *DescribeKeysRequest) (*DescribeKeysResponse, int, error) {
	return p.CtxDescribeKeys(context.Background(), req)
}

func (p *KMS) CtxUpdateKey(ctx context.Context, req *UpdateKeyRequest) (*UpdateKeyResponse, int, error) {
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

	statusCode, err := p.commonHandlerQuery(ctx, "UpdateKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) UpdateKey(req *UpdateKeyRequest) (*UpdateKeyResponse, int, error) {
	return p.CtxUpdateKey(context.Background(), req)
}

func (p *KMS) CtxGenerateDataKey(ctx context.Context, req *GenerateDataKeyRequest) (*GenerateDataKeyResponse, int, error) {
	query := url.Values{}
	resp := new(GenerateDataKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	if req.NumberOfBytes != nil {
		query.Set("NumberOfBytes", strconv.Itoa(*req.NumberOfBytes))
	}

	reqBody := make(map[string]interface{})
	if req.EncryptionContext != nil {
		reqBody["EncryptionContext"] = req.EncryptionContext
	}

	statusCode, err := p.commonHandlerJson(ctx, "GenerateDataKey", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) GenerateDataKey(req *GenerateDataKeyRequest) (*GenerateDataKeyResponse, int, error) {
	return p.CtxGenerateDataKey(context.Background(), req)
}

func (p *KMS) CtxEncrypt(ctx context.Context, req *EncryptRequest) (*EncryptResponse, int, error) {
	query := url.Values{}
	resp := new(EncryptResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	reqBody := make(map[string]interface{})
	if req.EncryptionContext != nil {
		reqBody["EncryptionContext"] = req.EncryptionContext
	}
	reqBody["Plaintext"] = req.Plaintext

	statusCode, err := p.commonHandlerJson(ctx, "Encrypt", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) Encrypt(req *EncryptRequest) (*EncryptResponse, int, error) {
	return p.CtxEncrypt(context.Background(), req)
}

func (p *KMS) CtxDecrypt(ctx context.Context, req *DecryptRequest) (*DecryptResponse, int, error) {
	query := url.Values{}
	resp := new(DecryptResponse)

	reqBody := make(map[string]interface{})
	if req.EncryptionContext != nil {
		reqBody["EncryptionContext"] = req.EncryptionContext
	}
	reqBody["CiphertextBlob"] = req.CiphertextBlob

	statusCode, err := p.commonHandlerJson(ctx, "Decrypt", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) Decrypt(req *DecryptRequest) (*DecryptResponse, int, error) {
	return p.CtxDecrypt(context.Background(), req)
}

func (p *KMS) CtxEnableKey(ctx context.Context, req *EnableKeyRequest) (*EnableKeyResponse, int, error) {
	query := url.Values{}
	resp := new(EnableKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery(ctx, "EnableKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) EnableKey(req *EnableKeyRequest) (*EnableKeyResponse, int, error) {
	return p.CtxEnableKey(context.Background(), req)
}

func (p *KMS) CtxDisableKey(ctx context.Context, req *DisableKeyRequest) (*DisableKeyResponse, int, error) {
	query := url.Values{}
	resp := new(DisableKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery(ctx, "DisableKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DisableKey(req *DisableKeyRequest) (*DisableKeyResponse, int, error) {
	return p.CtxDisableKey(context.Background(), req)
}

func (p *KMS) CtxScheduleKeyDeletion(ctx context.Context, req *ScheduleKeyDeletionRequest) (*ScheduleKeyDeletionResponse, int, error) {
	query := url.Values{}
	resp := new(ScheduleKeyDeletionResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)
	if req.PendingWindowInDays != nil {
		query.Set("PendingWindowInDays", strconv.Itoa(*req.PendingWindowInDays))
	}

	statusCode, err := p.commonHandlerQuery(ctx, "ScheduleKeyDeletion", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil

}

func (p *KMS) ScheduleKeyDeletion(req *ScheduleKeyDeletionRequest) (*ScheduleKeyDeletionResponse, int, error) {
	return p.CtxScheduleKeyDeletion(context.Background(), req)
}

func (p *KMS) CtxCancelKeyDeletion(ctx context.Context, req *CancelKeyDeletionRequest) (*CancelKeyDeletionResponse, int, error) {
	query := url.Values{}
	resp := new(CancelKeyDeletionResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery(ctx, "CancelKeyDeletion", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) ArchiveKey(req *ArchiveKeyRequest) (*ArchiveKeyResponse, int, error) {
	query := url.Values{}
	resp := new(ArchiveKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("ArchiveKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) CancelArchiveKey(req *CancelArchiveKeyRequest) (*CancelArchiveKeyResponse, int, error) {
	query := url.Values{}
	resp := new(CancelArchiveKeyResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("CancelArchiveKey", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) EnableKeyRotation(req *EnableKeyRotationRequest) (*EnableKeyRotationResponse, int, error) {
	query := url.Values{}
	resp := new(EnableKeyRotationResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("EnableKeyRotation", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) DisableKeyRotation(req *DisableKeyRotationRequest) (*DisableKeyRotationResponse, int, error) {
	query := url.Values{}
	resp := new(DisableKeyRotationResponse)

	query.Set("KeyringName", req.KeyringName)
	query.Set("KeyName", req.KeyName)

	statusCode, err := p.commonHandlerQuery("DisableKeyRotation", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) ReEncrypt(req *ReEncryptRequest) (*ReEncryptResponse, int, error) {
	query := url.Values{}
	resp := new(ReEncryptResponse)

	query.Set("NewKeyringName", req.NewKeyringName)
	query.Set("NewKeyName", req.NewKeyName)
	reqBody := make(map[string]interface{})
	if req.OldEncryptionContext != nil {
		reqBody["OldEncryptionContext"] = req.OldEncryptionContext
	}
	if req.NewEncryptionContext != nil {
		reqBody["NewEncryptionContext"] = req.NewEncryptionContext
	}
	reqBody["CiphertextBlob"] = req.CiphertextBlob

	statusCode, err := p.commonHandlerJson("ReEncrypt", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *KMS) CancelKeyDeletion(req *CancelKeyDeletionRequest) (*CancelKeyDeletionResponse, int, error) {
	return p.CtxCancelKeyDeletion(context.Background(), req)
}

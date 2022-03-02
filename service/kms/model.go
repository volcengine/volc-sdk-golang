package kms

import (
	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/base"
	"time"
)

// Base ...
type Base struct {
	ID           uuid.UUID
	CreationDate int64
	UpdateDate   int64
}

// PaginationInfo ...
type PaginationInfo struct {
	TotalCount  int
	PageSize    int
	CurrentPage int
	Count       int
}

// Keyring ...
type Keyring struct {
	Base
	KeyringName string
	KeyringType string
	Description string
	UID         string
}

// CustomerMasterKey ...
type CustomerMasterKey struct {
	Base
	KeyName            string
	KeySpec            string
	Description        string
	KeyState           string
	KeyUsage           string
	ProtectionLevel    string
	ScheduleDeleteTime *time.Time `json:",omitempty"`
}

// CreateKeyringRequest ...
type CreateKeyringRequest struct {
	KeyringName string
	KeyringType *string
	Description *string
}

// CreateKeyringResult ...
type CreateKeyringResult struct {
	Keyring *Keyring
}

// CreateKeyringResponse ...
type CreateKeyringResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CreateKeyringResult `json:",omitempty"`
}

// DescribeKeyringsRequest ...
type DescribeKeyringsRequest struct {
	CurrentPage *int
	PageSize    *int
}

// DescribeKeyringsResult ...
type DescribeKeyringsResult struct {
	Keyrings []*Keyring
}

// DescribeKeyringsResponse ...
type DescribeKeyringsResponse struct {
	PageInfo         *PaginationInfo `json:",omitempty"`
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeKeyringsResult `json:",omitempty"`
}

// UpdateKeyringRequest ...
type UpdateKeyringRequest struct {
	KeyringName    string
	NewKeyringName *string
	Description    *string
}

// UpdateKeyringResponse ...
type UpdateKeyringResponse struct {
	ResponseMetadata base.ResponseMetadata
}

// QueryKeyringRequest ...
type QueryKeyringRequest struct {
	KeyringName string
}

// QueryKeyringResult ...
type QueryKeyringResult struct {
	Keyring *Keyring
}

// QueryKeyringResponse ...
type QueryKeyringResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *QueryKeyringResult `json:",omitempty"`
}

// CreateKeyRequest ...
type CreateKeyRequest struct {
	KeyringName     string
	KeyName         string
	KeySpec         *string
	Description     *string
	KeyUsage        *string
	ProtectionLevel *string
}

// CreateKeyResult ...
type CreateKeyResult struct {
	Key *CustomerMasterKey
}

// CreateKeyResponse ...
type CreateKeyResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CreateKeyResult `json:",omitempty"`
}

// DescribeKeysRequest ...
type DescribeKeysRequest struct {
	KeyringName string
	CurrentPage *int
	PageSize    *int
}

// DescribeKeysResult ...
type DescribeKeysResult struct {
	Keys []*CustomerMasterKey
}

// DescribeKeysResponse ...
type DescribeKeysResponse struct {
	PageInfo         *PaginationInfo `json:",omitempty"`
	ResponseMetadata base.ResponseMetadata
	Result           *DescribeKeysResult `json:",omitempty"`
}

// UpdateKeyRequest ...
type UpdateKeyRequest struct {
	KeyringName string
	KeyName     string
	NewKeyName  *string
	Description *string
}

// UpdateKeyResponse ...
type UpdateKeyResponse struct {
	ResponseMetadata base.ResponseMetadata
}

// GenerateDataKeyRequest ...
type GenerateDataKeyRequest struct {
	KeyringName       string
	KeyName           string
	EncryptionContext map[string]string
	NumberOfBytes     *int
}

// GenerateDataKeyResult ...
type GenerateDataKeyResult struct {
	Plaintext      []byte
	CiphertextBlob []byte
}

// GenerateDataKeyResponse ...
type GenerateDataKeyResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GenerateDataKeyResult `json:",omitempty"`
}

// EncryptRequest ...
type EncryptRequest struct {
	KeyringName       string
	KeyName           string
	EncryptionContext map[string]string
	Plaintext         []byte
}

// EncryptResult ...
type EncryptResult struct {
	CiphertextBlob []byte
}

// EncryptResponse ...
type EncryptResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *EncryptResult `json:",omitempty"`
}

// DecryptRequest ...
type DecryptRequest struct {
	EncryptionContext map[string]string
	CiphertextBlob    []byte
}

// DecryptResult ...
type DecryptResult struct {
	Plaintext []byte
}

// DecryptResponse ...
type DecryptResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *DecryptResult `json:",omitempty"`
}

// EnableKeyRequest ...
type EnableKeyRequest struct {
	KeyringName string
	KeyName     string
}

// EnableKeyResponse ...
type EnableKeyResponse struct {
	ResponseMetadata base.ResponseMetadata
}

// DisableKeyRequest ...
type DisableKeyRequest struct {
	KeyringName string
	KeyName     string
}

// DisableKeyResponse ...
type DisableKeyResponse struct {
	ResponseMetadata base.ResponseMetadata
}

// ScheduleKeyDeletionRequest ...
type ScheduleKeyDeletionRequest struct {
	KeyringName         string
	KeyName             string
	PendingWindowInDays *int
}

// ScheduleKeyDeletionResponse ...
type ScheduleKeyDeletionResponse struct {
	ResponseMetadata base.ResponseMetadata
}

// CancelKeyDeletionRequest ...
type CancelKeyDeletionRequest struct {
	KeyringName string
	KeyName     string
}

// CancelKeyDeletionResponse ...
type CancelKeyDeletionResponse struct {
	ResponseMetadata base.ResponseMetadata
}

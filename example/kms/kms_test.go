package kms

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/kms"
	"net/http"
	"os"
	"testing"
)

const (
	testAk = "<your access key id>"
	testSk = "<your access key secret>"
	region = "<region>"
)

var (
	keyringName       = "TOP_SDK_EXAMPLE_KEYRING-" + uuid.New().String()[20:25]
	keyName           = "TOP_SDK_EXAMPLE_KEY-" + uuid.New().String()[20:30]
	encryptionContext = map[string]string{
		"key": "value",
	}
)

func init() {
	kms.DefaultInstance.SetRegion(region)
	kms.DefaultInstance.Client.SetAccessKey(testAk)
	kms.DefaultInstance.Client.SetSecretKey(testSk)
	kms.DefaultInstance.Client.SetRetrySettings(&base.RetrySettings{AutoRetry: true})
}

func analyze(resp interface{}, status int, err error) {
	if err != nil {
		fmt.Printf("Perform action error: %s", err)
		os.Exit(1)
	}
	if status != http.StatusOK {
		fmt.Printf("Perform action status code[%d] invalid.", status)
		os.Exit(1)
	}
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestKMS_CreateKeyring(t *testing.T) {
	analyze(kms.DefaultInstance.CreateKeyring(&kms.CreateKeyringRequest{KeyringName: keyringName}))
}

func TestKMS_DescribeKeyrings(t *testing.T) {
	currentPage, pageSize := 1, 2
	analyze(kms.DefaultInstance.DescribeKeyrings(&kms.DescribeKeyringsRequest{
		CurrentPage: &currentPage,
		PageSize:    &pageSize,
	}))
}

func TestKMS_UpdateKeyring(t *testing.T) {
	desc := "test"
	analyze(kms.DefaultInstance.UpdateKeyring(&kms.UpdateKeyringRequest{
		KeyringName: keyringName,
		Description: &desc,
	}))
}

func TestKMS_QueryKeyring(t *testing.T) {
	analyze(kms.DefaultInstance.QueryKeyring(&kms.QueryKeyringRequest{KeyringName: keyringName}))
}

func TestKMS_CreateKey(t *testing.T) {
	analyze(kms.DefaultInstance.CreateKey(&kms.CreateKeyRequest{KeyringName: keyringName, KeyName: keyName}))
}

func TestKMS_DescribeKeys(t *testing.T) {
	currentPage, pageSize := 1, 2
	analyze(kms.DefaultInstance.DescribeKeys(&kms.DescribeKeysRequest{
		KeyringName: keyringName,
		CurrentPage: &currentPage,
		PageSize:    &pageSize,
	}))
}

func TestKMS_UpdateKey(t *testing.T) {
	desc := "test"
	analyze(kms.DefaultInstance.UpdateKey(&kms.UpdateKeyRequest{
		KeyringName: keyringName,
		KeyName:     keyName,
		Description: &desc,
	}))
}

func TestKMS_GenerateDataKey_And_Decrypt(t *testing.T) {
	numberOfBytes := 32
	resp, status, err := kms.DefaultInstance.GenerateDataKey(&kms.GenerateDataKeyRequest{
		KeyringName:       keyringName,
		KeyName:           keyName,
		EncryptionContext: encryptionContext,
		NumberOfBytes:     &numberOfBytes,
	})
	analyze(resp, status, err)
	analyze(kms.DefaultInstance.Decrypt(&kms.DecryptRequest{
		EncryptionContext: encryptionContext,
		CiphertextBlob:    resp.Result.CiphertextBlob,
	}))
}

func TestKMS_Encrypt(t *testing.T) {
	analyze(kms.DefaultInstance.Encrypt(&kms.EncryptRequest{
		KeyringName:       keyringName,
		KeyName:           keyName,
		EncryptionContext: encryptionContext,
		Plaintext:         []byte(base64.StdEncoding.EncodeToString([]byte("hello world"))),
	}))
}

func TestKMS_DisableKey(t *testing.T) {
	analyze(kms.DefaultInstance.DisableKey(&kms.DisableKeyRequest{
		KeyringName: keyringName,
		KeyName:     keyName,
	}))
}

func TestKMS_EnableKey(t *testing.T) {
	analyze(kms.DefaultInstance.EnableKey(&kms.EnableKeyRequest{
		KeyringName: keyringName,
		KeyName:     keyName,
	}))
}

func TestKMS_ScheduleKeyDeletion(t *testing.T) {
	pendingWindowInDays := 20
	analyze(kms.DefaultInstance.ScheduleKeyDeletion(&kms.ScheduleKeyDeletionRequest{
		KeyringName:         keyringName,
		KeyName:             keyName,
		PendingWindowInDays: &pendingWindowInDays,
	}))
}

func TestKMS_CancelKeyDeletion(t *testing.T) {
	analyze(kms.DefaultInstance.CancelKeyDeletion(&kms.CancelKeyDeletionRequest{
		KeyringName: keyringName,
		KeyName:     keyName,
	}))
}

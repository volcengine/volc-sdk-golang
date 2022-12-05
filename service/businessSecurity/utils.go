package businessSecurity

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func ToUrlValues(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	return
}

// pkcs7Padding padding
func pkcs7Padding(ciphertext []byte) []byte {
	bs := aes.BlockSize
	padding := bs - len(ciphertext)%bs
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

// AesCBCEncryptWithBase64 encrypt
func AesCBCEncryptWithBase64(secretKey, origData string) (string, error) {
	if len(secretKey) < aes.BlockSize {
		return "", fmt.Errorf("key len must be more than 16")
	}
	newKey := []byte(secretKey)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	newOrigData := []byte(origData)
	newOrigData = pkcs7Padding(newOrigData)
	blockMode := cipher.NewCBCEncrypter(block, newKey[:block.BlockSize()])
	encrypted := make([]byte, len(newOrigData))
	blockMode.CryptBlocks(encrypted, newOrigData)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

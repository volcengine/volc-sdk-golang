package businessSecurity

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
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

func makeRequest(apiInfo *base.ApiInfo, serviceInfo *base.ServiceInfo, query url.Values, ct string) (*http.Request, error) {
	header := mergeHeader(serviceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)

	u := url.URL{
		Scheme:   serviceInfo.Scheme,
		Host:     serviceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request")
	}
	req.Header = header
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}

	// Because service info could be changed by SetRegion, so set UA header for every request here.
	req.Header.Set("User-Agent", strings.Join([]string{base.SDKName, base.SDKVersion}, "/"))

	return req, nil
}

func mergeQuery(query1, query2 url.Values) (query url.Values) {
	query = url.Values{}

	for k, vv := range query1 {
		for _, v := range vv {
			query.Add(k, v)
		}
	}

	for k, vv := range query2 {
		for _, v := range vv {
			query.Add(k, v)
		}
	}
	return
}

func mergeHeader(header1, header2 http.Header) (header http.Header) {
	header = http.Header{}

	for k, v := range header1 {
		header.Set(k, strings.Join(v, ";"))
	}

	for k, v := range header2 {
		header.Set(k, strings.Join(v, ";"))
	}

	return
}

func getTimeout(serviceTimeout, apiTimeout time.Duration) time.Duration {
	timeout := time.Second
	if serviceTimeout != time.Duration(0) {
		timeout = serviceTimeout
	}
	if apiTimeout != time.Duration(0) {
		timeout = apiTimeout
	}
	return timeout
}

package vod

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	DSAHmacSha1   = "HMAC-SHA1"
	DSAHmacSha256 = "HMAC-SHA256"

	Version1 = "1.0"

	SprAuth = ":"
	SprSign = "&"

	DateYYMMDD = "20060102"
)

var (
	ServiceVOD = str2Bytes("vod")

	ErrAccessKeyInvalid = errors.New("access key invalid")
	ErrSecretKeyInvalid = errors.New("secret key invalid")
)

func createAuth(dsa, version, accessKey, secretKey string, expireSeconds int64) (string, error) {
	if err := validate(accessKey, secretKey); err != nil {
		return "", err
	}

	deadline := time.Now().Add(time.Duration(expireSeconds) * time.Second)
	timestamp := strconv.FormatInt(deadline.Unix(), 10)
	sign := buildSign(dsa, version, timestamp, parseKey(secretKey, deadline))
	tokens := []string{dsa, version, timestamp, accessKey, sign}
	return strings.Join(tokens, SprAuth), nil
}

func validate(accessKey, secretKey string) error {
	if accessKey == "" {
		return ErrAccessKeyInvalid
	}
	if secretKey == "" {
		return ErrSecretKeyInvalid
	}
	return nil
}

func buildSign(dsa, version, timestamp string, key []byte) string {
	data := str2Bytes(join(dsa, SprSign, version, SprSign, timestamp))
	switch dsa {
	case DSAHmacSha1:
		return encodeToBase64(getHmac1(key, data))
	case DSAHmacSha256:
		return encodeToBase64(getHmac256(key, data))
	}
	return ""
}

func parseKey(key string, t time.Time) []byte {
	dateKey := getHmac256(str2Bytes(key), str2Bytes(GetDate(t)))
	return str2Bytes(hex.EncodeToString(getHmac256(dateKey, ServiceVOD)))
}

func join(tokens ...string) string {
	var buf bytes.Buffer
	for idx := range tokens {
		buf.WriteString(tokens[idx])
	}
	return buf.String()
}

func GetDate(t time.Time) string {
	return t.Format(DateYYMMDD)
}

func getHmac1(key, data []byte) []byte {
	return getHmac(key, data, sha1.New)
}

func getHmac256(key, data []byte) []byte {
	return getHmac(key, data, sha256.New)
}

type hashFunc func() hash.Hash

func getHmac(key, data []byte, hf hashFunc) []byte {
	hash := hmac.New(hf, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func str2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := &reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bh))
}

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
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	DSAHmacSha1   = "HMAC-SHA1"

	Version2 = "2.0"

	SprAuth = ":"
	SprSign = "&"

	DateYYYYMMDDTHHmmss = "20060102T150405Z"
)

var (
	ServiceVOD = stringToBytes("vod")

	ErrAccessKeyInvalid = errors.New("access key invalid")
	ErrSecretKeyInvalid = errors.New("secret key invalid")
)

func createAuth(dsa, version, accessKey, secretKey string, expireSeconds int64) (string, error) {
	if err := validate(accessKey, secretKey); err != nil {
		return "", err
	}
	deadline := time.Now().Add(time.Duration(expireSeconds) * time.Second)
	timestamp := strconv.FormatInt(deadline.Unix(), 10)
	dataKey := parseKey(secretKey, deadline)
	sign := BuildSign(dsa, version, timestamp, dataKey)
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

func BuildSign(dsa, version, timestamp string, key []byte) string {
	data := stringToBytes(join(dsa, SprSign, version, SprSign, timestamp))
	switch dsa {
	case DSAHmacSha1:
		return encodeToBase64(getHmac1(key, data))
	}
	return ""
}

func parseKey(key string, t time.Time) []byte {
	dateKey := getHmac256(stringToBytes(key), stringToBytes(GetDate(t)))
	return stringToBytes(hex.EncodeToString(getHmac256(dateKey, ServiceVOD)))
	//return stringToBytes(hex.EncodeToString(getHmac256(dateKey, stringToBytes(""))))
}

func join(tokens ...string) string {
	var buf bytes.Buffer
	for idx := range tokens {
		buf.WriteString(tokens[idx])
	}
	return buf.String()
}

func GetDate(t time.Time) string {
	return t.UTC().Format(DateYYYYMMDDTHHmmss)
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

func stringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}
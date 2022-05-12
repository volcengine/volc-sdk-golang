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
)

const (
	DSAHmacSha1 = "HMAC-SHA1"

	Version2 = "2.0"

	SprAuth = ":"
	SprSign = "&"

	DateYYYYMMDDTHHmmss = "20060102T150405Z"
)

var (
	ServiceVOD = []byte("vod")

	ErrAccessKeyInvalid = errors.New("access key invalid")
	ErrSecretKeyInvalid = errors.New("secret key invalid")
)

func createAuth(dsa, version, accessKey, secretKey, region string, expireSeconds int64) (string, error) {
	if err := validate(accessKey, secretKey); err != nil {
		return "", err
	}
	deadline := time.Now().Add(time.Duration(expireSeconds) * time.Second)
	timestamp := strconv.FormatInt(deadline.Unix(), 10)
	dataKey := getSignedKey(secretKey, deadline, region)
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
	data := []byte(join(dsa, SprSign, version, SprSign, timestamp))
	switch dsa {
	case DSAHmacSha1:
		return encodeToBase64(getHmac1(key, data))
	}
	return ""
}

func getSignedKey(key string, t time.Time, region string) []byte {
	kDate := getHmac256([]byte(key), []byte(GetDate(t)))
	kRegion := getHmac256(kDate, []byte(region))
	kService := getHmac256(kRegion, ServiceVOD)
	kCredentials := getHmac256(kService, []byte("request"))
	return []byte(hex.EncodeToString(kCredentials))
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
	hashData := hmac.New(hf, key)
	hashData.Write(data)
	return hashData.Sum(nil)
}

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

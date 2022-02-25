package url_sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func getStorageAndRegion(oid string) (string, error) {
	tokens := strings.Split(oid, "/")
	if len(tokens) != lengthTokenOID {
		return "", ErrOIDInvalid
	}

	bucket := tokens[0]
	tokens = strings.Split(bucket, "-")
	if len(tokens) != lengthTokenBucket {
		return "", ErrBucketInvalid
	}
	return tokens[0] + "/" + tokens[1], nil
}

func buildAuthPath(sr, oid string) string {
	return fmt.Sprintf(authPathTemplate, sr, oid)
}

func buildURLPath(token, expiration, authPath string) string {
	return fmt.Sprintf(urlPathTemplate, token, expiration, authPath)
}

func buildExpiration(ttl time.Duration) string {
	return fmt.Sprintf(expirationTemplate, now().Add(ttl).Unix())
}

func buildMD5(parts ...string) string {
	hash := md5.New()
	for _, part := range parts {
		hash.Write(str2Bytes(part))
	}
	return hex.EncodeToString(hash.Sum(nil))
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


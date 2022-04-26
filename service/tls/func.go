package tls

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var now = func() time.Time {
	return time.Now().UTC()
}

func prepareRequestV4(request *http.Request) *http.Request {
	basicHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
		"X-Date":       timestampV4(),
	}

	for header, value := range basicHeaders {
		if len(request.Header.Get(header)) == 0 {
			request.Header.Set(header, value)
		}
	}

	if request.URL.Path == "" {
		request.URL.Path += "/"
	}

	return request
}

func timestampV4() string {
	return now().Format(timeFormatV4)
}

func hashedCanonicalRequestV4(request *http.Request, meta *metadata) string {
	payload := readAndReplaceBody(request)
	payloadHash := hashSHA256(payload)
	request.Header.Set("X-Content-Sha256", payloadHash)
	request.Header.Set("Host", request.Host)

	sortedHeaderKeys := filterHeader(request.Header)
	headersToSign := concatHeader(sortedHeaderKeys, request.Header)

	meta.signedHeaders = concat(";", sortedHeaderKeys...)
	canonicalRequest := concat("\n", request.Method, normuri(request.URL.Path), normquery(request.URL.Query()), headersToSign, meta.signedHeaders, payloadHash)

	return hashSHA256([]byte(canonicalRequest))
}

func concatHeader(headers []string, header http.Header) string {
	var headersToSign string
	for _, key := range headers {
		value := strings.TrimSpace(header.Get(key))

		if key == "host" {
			if strings.Contains(value, ":") {
				split := strings.Split(value, ":")
				port := split[1]
				if port == "80" || port == "443" {
					value = split[0]
				}
			}
		}

		headersToSign += key + ":" + value + "\n"
	}

	return headersToSign
}

func concat(delim string, str ...string) string {
	return strings.Join(str, delim)
}

func hashSHA256(content []byte) string {
	h := sha256.New()
	h.Write(content)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func readAndReplaceBody(request *http.Request) []byte {
	if request.Body == nil {
		return []byte{}
	}

	payload, _ := ioutil.ReadAll(request.Body)
	request.Body = ioutil.NopCloser(bytes.NewReader(payload))

	return payload
}

func normuri(uri string) string {
	uriParts := strings.Split(uri, "/")

	for i := range uriParts {
		uriParts[i] = encodePathFrag(uriParts[i])
	}

	return strings.Join(uriParts, "/")
}

func encodePathFrag(s string) string {
	hexCount := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			hexCount++
		}
	}

	t := make([]byte, len(s)+2*hexCount)

	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		} else {
			t[j] = c
			j++
		}
	}

	return string(t)
}

func filterHeader(header http.Header) []string {
	var sortedHeaderKeys []string

	for key := range header {
		switch key {
		case "Content-Type", "Content-Md5", "Host":
		default:
			if !strings.HasPrefix(key, "X-") {
				continue
			}
		}

		sortedHeaderKeys = append(sortedHeaderKeys, strings.ToLower(key))
	}

	sort.Strings(sortedHeaderKeys)

	return sortedHeaderKeys
}

func normquery(v url.Values) string {
	queryString := v.Encode()

	return strings.Replace(queryString, "+", "%20", -1)
}

func shouldEscape(c byte) bool {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
		return false
	}

	if '0' <= c && c <= '9' {
		return false
	}

	if c == '-' || c == '_' || c == '.' || c == '~' {
		return false
	}

	return true
}

func stringToSign(request *http.Request, hashedCanonReq string, meta *metadata) string {
	requestTs := request.Header.Get("X-Date")

	meta.algorithm = "HMAC-SHA256"
	meta.date = tsDate(requestTs)
	meta.credentialScope = concat("/", meta.date, meta.region, meta.service, "request")

	return concat("\n", meta.algorithm, requestTs, meta.credentialScope, hashedCanonReq)
}

func tsDate(timestamp string) string {
	return timestamp[:8]
}

func signingKey(secretKey, date, region, service string) []byte {
	kDate := hmacSHA256([]byte(secretKey), date)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "request")
	return kSigning
}

func signature(signingKey []byte, stringToSign string) string {
	return hex.EncodeToString(hmacSHA256(signingKey, stringToSign))
}

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func buildAuthHeader(signature string, meta *metadata, keys Credentials) string {
	credential := keys.AccessKeyID + "/" + meta.credentialScope

	return meta.algorithm +
		" Credential=" + credential +
		", SignedHeaders=" + meta.signedHeaders +
		", Signature=" + signature
}

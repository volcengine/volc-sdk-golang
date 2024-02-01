package base

import (
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
	"time"
)

const (
	RegionCnNorth1    = "cn-north-1"
	RegionUsEast1     = "us-east-1"
	RegionApSingapore = "ap-singapore-1"

	RegionApSouthEast1 = "ap-southeast-1"

	timeFormatV4 = "20060102T150405Z"
)

type ServiceInfo struct {
	Timeout     time.Duration
	Scheme      string
	Host        string
	Header      http.Header
	Credentials Credentials
	Retry       RetrySettings
}

type ApiInfo struct {
	Method  string
	Path    string
	Query   url.Values
	Form    url.Values
	Timeout time.Duration
	Header  http.Header
	Retry   RetrySettings
}

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	Service         string
	Region          string
	SessionToken    string
}

type metadata struct {
	algorithm       string
	credentialScope string
	signedHeaders   string
	date            string
	region          string
	service         string
}

// Unified JSON return results
type CommonResponse struct {
	ResponseMetadata ResponseMetadata
	Result           interface{} `json:"Result,omitempty"`
}

type BaseResp struct {
	Status      string
	CreatedTime int64
	UpdatedTime int64
}

type ErrorObj struct {
	CodeN   int
	Code    string
	Message string
}

type ResponseMetadata struct {
	RequestId string
	Service   string    `json:",omitempty"`
	Region    string    `json:",omitempty"`
	Action    string    `json:",omitempty"`
	Version   string    `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}

type Policy struct {
	Statement []*Statement
}

const (
	StatementEffectAllow = "Allow"
	StatementEffectDeny  = "Deny"
)

type Statement struct {
	Effect    string
	Action    []string
	Resource  []string
	Condition string `json:",omitempty"`
}

type SecurityToken2 struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	ExpiredTime     string
	CurrentTime     string
}

type InnerToken struct {
	LTAccessKeyId         string
	AccessKeyId           string
	SignedSecretAccessKey string
	ExpiredTime           int64
	PolicyString          string
	Signature             string
}

type RetrySettings struct {
	AutoRetry     bool
	RetryTimes    *uint64
	RetryInterval *time.Duration
}

type RequestParam struct {
	IsSignUrl bool
	Body      []byte
	Method    string
	Date      time.Time
	Path      string
	Host      string
	QueryList url.Values
	Headers   http.Header
}

type SignRequest struct {
	XDate          string
	XNotSignBody   string
	XCredential    string
	XAlgorithm     string
	XSignedHeaders string
	XSignedQueries string
	XSignature     string
	XSecurityToken string

	Host           string
	ContentType    string
	XContentSha256 string
	Authorization  string
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type MultiPartItem struct {
	header textproto.MIMEHeader
	data   io.Reader
}

func CreateMultiPartItem(header textproto.MIMEHeader, data io.Reader) *MultiPartItem {
	return &MultiPartItem{
		header: header,
		data:   data,
	}
}

func CreateMultiPartItemFormFile(fieldname, filename string, data io.Reader) *MultiPartItem {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", "application/octet-stream")
	return CreateMultiPartItem(h, data)
}

func CreateMultiPartItemFormField(fieldname string, data string) *MultiPartItem {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"`, escapeQuotes(fieldname)))
	return CreateMultiPartItem(h, strings.NewReader(data))
}

package base

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var (
	serviceInfo = &ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: Credentials{Region: RegionCnNorth1, Service: "iam"},
	}

	apiList = map[string]*ApiInfo{
		"ListUsers": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUsers"},
				"Version": []string{"2018-01-01"},
			},
		},
	}
)

func TestClient_GetSignUrl(t *testing.T) {
	client := NewClient(serviceInfo, apiList)
	// set aksk or read from ~/.volc/config
	// client.SetAccessKey("your-ak")
	// client.SetSecretKey("your-sk")

	url, _ := client.GetSignUrl("ListUsers", nil)
	fmt.Println(url)
}

func TestClient_Query(t *testing.T) {
	client := NewClient(serviceInfo, apiList)
	// set aksk or read from ~/.volc/config
	// client.SetAccessKey("your-ak")
	// client.SetSecretKey("your-sk")

	resp, _, _ := client.Query("ListUsers", nil)
	fmt.Println(string(resp))
}

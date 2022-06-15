package base

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	client.SetAccessKey("ak")
	client.SetSecretKey("sk")
	urlStr, _ := client.GetSignUrl("ListUsers", nil)
	fmt.Println(urlStr)
	if resp, err := http.Get("baseUrl?" + urlStr); err == nil {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func TestClient_Query(t *testing.T) {
	client := NewClient(serviceInfo, apiList)
	// set aksk or read from ~/.volc/config
	// client.SetAccessKey("your-ak")
	// client.SetSecretKey("your-sk")

	resp, _, _ := client.Query("ListUsers", nil)
	fmt.Println(string(resp))
}

func TestClientProxy(t *testing.T) {
	// normal case
	client := NewClient(serviceInfo, apiList)
	_, code, _ := client.Query("ListUsers", nil)
	if code > 499 {
		t.Fatalf("default client should return bad request, got %v instead", code)
	}

	// with bad proxy
	t.Setenv(httpProxy, "http://127.0.0.1:1234")
	t.Setenv(httpsProxy, "http://127.0.0.1:1234")
	client.Client.Transport = &http.Transport{Proxy: volcProxy()}
	_, code, _ = client.Query("ListUsers", nil)
	if code != 500 {
		t.Fatalf("invaid proxy client should return connection refused, got %v instead", code)
	}

	// with proxy
	go setupProxy("127.0.0.1:1234")
	_, code, _ = client.Query("ListUsers", nil)
	if code > 499 {
		t.Fatalf("proxy client should return as default client, got %v instead", code)
	}
}

func setupProxy(addr string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		outReq := request.Clone(context.Background())
		res, err := http.DefaultTransport.RoundTrip(outReq)
		if err != nil {
			writer.WriteHeader(http.StatusBadGateway)
			return
		}
		writer.WriteHeader(res.StatusCode)
		io.Copy(writer, res.Body)
		res.Body.Close()
	})
	http.ListenAndServe(addr, nil)
}

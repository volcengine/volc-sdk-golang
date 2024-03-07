package main

import (
	"errors"
	"fmt"
	api "github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	client "github.com/volcengine/volc-sdk-golang/service/maas/v2"
)

func TestCreateApiKey(r *client.MaaS, req *api.CreateOrRefreshApiKeyRequest) {
	rsp, _, err := r.CreateOrRefreshApiKey(req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v", errVal)
		}
		return
	}
	fmt.Printf("ApiKey=%v", rsp.ApiKey)
}

func main() {
	r := client.NewInstance("open.volcengineapi.com", "cn-beijing")

	endpointId1 := "{YOUR_ENDPOINT_ID1}"
	endpointId2 := "{YOUR_ENDPOINT_ID2}"

	req := &api.CreateOrRefreshApiKeyRequest{
		Ttl:            86400, // expected apiKey expired time
		EndpointIdList: []string{endpointId1, endpointId2},
	}

	TestCreateApiKey(r, req)
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	api "github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	client "github.com/volcengine/volc-sdk-golang/service/maas/v2"
)

func main() {
	r := client.NewInstance("maas-api.ml-platform-cn-beijing.volces.com", "cn-beijing")

	apikey := "{YOUR_APIKEY}"
	r.SetApikey(apikey)
	req := &api.ChatReq{
		Messages: []*api.Message{
			{
				Role:    api.ChatRoleUser,
				Content: "天为什么这么蓝？",
			},
			{
				Role:    api.ChatRoleAssistant,
				Content: "因为有你",
			},
			{
				Role:    api.ChatRoleUser,
				Content: "花儿为什么这么香？",
			},
		},
		Parameters: &api.Parameters{
			MaxNewTokens: 2000,
			Temperature:  0.8,
		},
	}

	endpointId := "{YOUR_ENDPOINT_ID}"
	testNormalChat(r, endpointId, req)
	testStreamChat(r, endpointId, req)
}

func testNormalChat(r *client.MaaS, endpointId string, req *api.ChatReq) {
	got, status, err := r.Chat(endpointId, req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v, status=%d\n", errVal, status)
		}
		return
	}
	fmt.Println("chat answer", marshalJson(got))
}

func testStreamChat(r *client.MaaS, endpointId string, req *api.ChatReq) {
	ch, err := r.StreamChat(endpointId, req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Println("meet maas error", errVal.Error())
		}
		return
	}

	for resp := range ch {
		if resp.Error != nil {
			// it is possible that error occurs during response processing
			fmt.Println(marshalJson(resp.Error))
			return
		}
		fmt.Println(marshalJson(resp))
		// last response may contain `usage`
		if resp.Usage != nil {
			// last message, will return full response including usage, role, finish_reason, etc.
			fmt.Println(marshalJson(resp.Usage))
		}
	}
}

func marshalJson(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

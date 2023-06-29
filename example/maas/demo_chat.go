package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/volcengine/volc-sdk-golang/service/maas"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api"
)

const modelName = "${YOUR_MODEL_NAME}"

func main() {
	r := maas.NewInstance("maas-api.ml-platform-cn-beijing.volces.com", "cn-beijing")

	// fetch ak&sk from environmental variables
	r.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	r.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))

	req := &api.ChatReq{
		Model: &api.Model{
			Name: modelName,
		},
		Messages: []*api.Message{
			{
				Role:    maas.ChatRoleOfUser,
				Content: "天为什么这么蓝？",
			},
			{
				Role:    maas.ChatRoleOfAssistant,
				Content: "因为有你",
			},
			{
				Role:    maas.ChatRoleOfUser,
				Content: "花儿为什么这么香？",
			},
		},
		Parameters: &api.Parameters{
			MaxNewTokens: 2000,
			Temperature:  0.8,
		},
	}
	TestChat(r, req)
	TestStreamChat(r, req)
}

func TestChat(r *maas.MaaS, req *api.ChatReq) {
	got, status, err := r.Chat(req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v, status=%d\n", errVal, status)
		}
		return
	}
	fmt.Println("chat answer", mustMarshalJson(got))
}

func TestStreamChat(r *maas.MaaS, req *api.ChatReq) {
	ch, err := r.StreamChat(req)

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
			fmt.Println(mustMarshalJson(resp.Error))
			return
		}
		fmt.Println(mustMarshalJson(resp))
		// last response may contain `usage`
		if resp.Usage != nil {
			// last message, will return full response including usage, role, finish_reason, etc.
			fmt.Println(mustMarshalJson(resp.Usage))
		}
	}
}

func mustMarshalJson(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

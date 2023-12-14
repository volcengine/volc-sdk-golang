package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	api "github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	client "github.com/volcengine/volc-sdk-golang/service/maas/v2"
)

func main() {
	r := client.NewInstance("maas-api.ml-platform-cn-beijing.volces.com", "cn-beijing")

	// fetch ak&sk from environmental variables
	r.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	r.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))

	req := &api.EmbeddingsReq{
		Input: []string{
			"天很蓝",
			"海很深",
		},
	}

	endpointId := "{YOUR_ENDPOINT_ID}"
	TestEmbeddings(r, endpointId, req)
}

func TestEmbeddings(r *client.MaaS, endpointId string, req *api.EmbeddingsReq) {
	got, status, err := r.Embeddings(endpointId, req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v, status=%d\n", errVal, status)
		}
		return
	}

	answer, _ := json.Marshal(got)
	fmt.Println("embeddings answer", string(answer))
}

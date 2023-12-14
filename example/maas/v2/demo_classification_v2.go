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

	req := &api.ClassificationReq{
		Query:  "花儿为什么这么香？",
		Labels: []string{"陈述句", "疑问句", "肯定句"},
	}

	endpointId := "{YOUR_ENDPOINT_ID}"
	TestClassification(r, endpointId, req)
}

func TestClassification(r *client.MaaS, endpointId string, req *api.ClassificationReq) {
	got, status, err := r.Classification(endpointId, req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v, status=%d\n", errVal, status)
		}
		return
	}

	answer, _ := json.Marshal(got)
	fmt.Println("classification answer", string(answer))
}

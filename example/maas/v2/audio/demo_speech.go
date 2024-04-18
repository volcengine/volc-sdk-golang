package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	api "github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	client "github.com/volcengine/volc-sdk-golang/service/maas/v2"
)

func main() {
	r := client.NewInstance("maas-api.ml-platform-cn-beijing.volces.com", "cn-beijing")

	// fetch ak&sk from environmental variables
	r.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	r.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))

	req := &api.SpeechReq{
		Input:          "你好欢迎光临",
		Voice:          "zh_male_rap",
		ResponseFormat: "mp3",
		Speed:          1.0,
	}

	endpointId := "{YOUR_ENDPOINT_ID}"
	path := "{YOUR_PATH}"
	TestSpeech(r, endpointId, req, path)
}

func TestSpeech(maas *client.MaaS, endpointId string, req *api.SpeechReq, path string) {
	resp, status, err := maas.Audio().Speech.Create(endpointId, req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error, requestId=%s, status=%d, err=%v", resp.GetRequestId(), status, errVal.Error())
		}
		return
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("create file meet error, requestId=%s, err=%v", resp.GetRequestId(), err.Error())
		return
	}

	// remember to close the file and read closer
	defer f.Close()
	defer resp.Close()

	_, err = io.Copy(f, resp)
	if err != nil {
		fmt.Printf("write file meet error, requestId=%s, err=%v", resp.GetRequestId(), err.Error())
	}

	fmt.Printf("finish create audio file, requestId=%s.", resp.GetRequestId())
}

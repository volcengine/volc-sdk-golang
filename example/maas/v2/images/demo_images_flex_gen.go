package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	client "github.com/volcengine/volc-sdk-golang/service/maas/v2"
)

func main() {
	r := client.NewInstance("maas-api.ml-platform-cn-beijing.volces.com", "cn-beijing")

	// fetch ak&sk from environmental variables
	r.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	r.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))

	controlImage, err := os.ReadFile("{YOUR_CONTROL_PICTURE_PATH}")
	if err != nil {
		panic(err)
	}

	req := &api.ImagesReq{
		Prompt:            "(sfw:1.0),(masterpiece,best quality,ultra highres),(realistic:1.15),(3D:1.0)",
		NegativePrompt:    "(embedding:EasyNegative:0.9),(embedding:badhandv4:1.3),terrible,injured,(nsfw:1.0),(nude:1.0)",
		ControlImageList:  [][]byte{controlImage},
		Strength:          0.75,
		Height:            512,
		Width:             512,
		NumInferenceSteps: 20,
	}

	endpointId := "{YOUR_ENDPOINT_ID}"
	TestImagesFlexGen(r, endpointId, req)
}

func TestImagesFlexGen(maas *client.MaaS, endpointId string, req *api.ImagesReq) {
	got, status, err := maas.Images().ImagesFlexGen(endpointId, req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v, status=%d\n", errVal, status)
		}
		return
	}

	s, _ := json.Marshal(got)
	fmt.Println("images flex gen chat answer", string(s))
}

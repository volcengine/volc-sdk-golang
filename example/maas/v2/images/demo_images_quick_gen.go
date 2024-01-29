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

	controlImage := loadImage("{YOUR_CONTROL_PICTURE_PATH}")
	req := &api.ImagesQuickGenReq{
		Prompt:         "(sfw:1.0),(masterpiece,best quality,ultra highres),(realistic:1.15),(3D:1.0),8k wallpaper,ultra detailed,beautiful and aesthetic,official art,real,(tech city background:1.3),(depth of field:1.1),(colorful:1),wind,(sky:1.25),chinese (high quality:1.3) ((golden:1.3) dragon:1.3),glowing,(1girl:1.1),portrait,(bright face:1.2),bangs,light smile,hair,(look at viewer1.2),(young:1.0),(big eyes:1),solo,Brilliant,face to viewer,(future architecture:1.1),(milk print hanfu:1.2),(liquid:1.1),(bubble:1.2),pvc texture,(building:1.2),(detailed skin:1.2),(science fiction:1.3),(machinery:1.2),anmuxi,(iridescent film coat:1.3),(iridescent (blue:0.15) film hanfu:1.30),(red lantern:1.3),",
		NegativePrompt: "(embedding:EasyNegative:0.9),(embedding:badhandv4:1.3),terrible,injured,(nsfw:1.0),(nude:1.0),naked,small eyes,Sleepy,big small eyes,(breasts:1.0),lowres,text,log,signature,symbol-shaped pupils,heterochromia,multicolored eyes,no pupils,slit pupils,asymmetrical pupils,asymmetrical eyes,asymmetrical eyebrows,streaked hair,colored inner hair,two-tone hair,multicolored hair,gradient hair,earrings,hair ornaments,asymmetrical breasts,multiple views,reference sheet,simple background,room,indoors,japan,blue arm,(old:1.2),sad,asian face,blue skin,monster,bone,europa,american,(facial tattoo:1),tattoo,(hat:1.1),blue face,blue skin,white lantern,western dargon,Cross-eyed,curls hai,flower on head,skin blemish,tongue out,(sky:1.1),(yellow skin:1.3),white lantern,centre parting,fog,(facepaint:1.2),",
		ControlImage:   controlImage,
		Parameters: &api.ImagesParameters{
			Seed:              38,
			Strength:          0.75,
			NumInferenceSteps: 20,
		},
	}

	endpointId := "{YOUR_ENDPOINT_ID}"
	TestImagesQuickGenChat(r, endpointId, req)
}

func TestImagesQuickGenChat(maas *client.MaaS, endpointId string, req *api.ImagesQuickGenReq) {
	got, status, err := maas.Images().ImagesQuickGen(endpointId, req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v, status=%d\n", errVal, status)
		}
		return
	}

	s, _ := json.Marshal(got)
	fmt.Println("images quick gen chat answer", string(s))
}

func loadImage(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}

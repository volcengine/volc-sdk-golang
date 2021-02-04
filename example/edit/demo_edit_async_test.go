package edit

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vedit"
)

func TestSubmitDirectEditTaskAsync(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vedit.NewInstance()
	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "",
	//	SecretAccessKey: "",
	//})

	// or set ak and ak as follow
	//edit.NewInstance().SetAccessKey("")
	//edit.NewInstance().SetSecretKey("")

	// your param str
	// for example
	paramStr := `
{
	"Upload": {
		"Uploader": "your uploader",
		"VideoName": "video"
	},
	"Output": {
		"Fps": 25,
		"Height": 720,
		"Quality": "medium",
		"Width": 1280
	},
	"Segments": [{
		"BackGround": "0xFFFFFFFF",
		"Duration": 3,
		"Elements": [],
		"Volume": 1
	}],
	"GlobalElements": []
}`
	var param interface{}

	err := json.Unmarshal([]byte(paramStr), &param)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request := &vedit.SubmitDirectEditTaskRequest{
		Param:        param,
		Priority:     0,
		CallbackArgs: "your callback args",
		CallbackUri:  "your callback uri",
	}

	resp, err := instance.SubmitDirectEditTaskAsync(request)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	retString, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(retString))
	return
}

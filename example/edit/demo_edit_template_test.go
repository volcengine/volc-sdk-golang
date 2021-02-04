package edit

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vedit"
)

func TestSubmitTemplateEditTaskAsync(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vedit.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//edit.NewInstance().SetAccessKey("")
	//edit.NewInstance().SetSecretKey("")

	request := &vedit.SubmitTemplateTaskRequest{
		Space:        "your space",
		VideoName:    []string{},
		TemplateId:   "your template_id",
		Params:       [][]vedit.TemplateParamItem{},
		Priority:     0,
		CallbackArgs: "your callback args",
		CallbackUri:  "your callback uri",
	}

	resp, err := instance.SubmitTemplateTaskAsync(request)

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

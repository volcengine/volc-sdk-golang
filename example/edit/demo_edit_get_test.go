package edit

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vedit"
)

func TestEdit_GetDirectEditResult(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vedit.NewInstance()
	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//edit.NewInstance().SetAccessKey("")
	//edit.NewInstance().SetSecretKey("")

	resp, err := instance.GetDirectEditResult(&vedit.GetDirectEditResultRequest{[]string{"your req id"}})

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

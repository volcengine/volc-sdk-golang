package demo_vod_workflow

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/models/vod/business"
	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_StartWorkflow(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")
	queryRequest := &request.VodStartWorkflowRequest{
		Vid:          "your vid",
		TemplateId:   "your template_id",
		Input:        &business.WorkflowParams{},
		Priority:     0,
		CallbackArgs: "your callback args",
	}
	resp, _, err := instance.StartWorkflow(queryRequest)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	fmt.Println(resp.GetResult().GetRunId())
}

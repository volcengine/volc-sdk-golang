package codePipeline

import (
	"fmt"
	"testing"

	. "github.com/volcengine/volc-sdk-golang/service/codePipeline/models"
)

const (
	Ak = "" // write your access key
	Sk = "" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func CreatePipeline(workspaceId string, item Pipeline) {
	res, err := DefaultInstance.CreatePipeline(&CreatePipelineRequest{
		WorkspaceId: workspaceId,
		Item:        item,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func ListPipelines(workspaceId string) {
	res, err := DefaultInstance.ListPipelines(&ListPipelinesRequest{
		WorkspaceId: workspaceId,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func GetPipeline(workspaceId string, pipelineId string) {
	res, err := DefaultInstance.GetPipeline(&GetPipelineRequest{
		WorkspaceId: workspaceId,
		Id:          pipelineId,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func RunPipeline(workspaceId string, pipelineId string) {
	res, err := DefaultInstance.RunPipeline(&RunPipelineRequest{
		WorkspaceId: workspaceId,
		Id:          pipelineId,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func ListPipelineRecords(workspaceId string, pipelineId string) {
	res, err := DefaultInstance.ListPipelineRecords(&ListPipelineRecordsRequest{
		WorkspaceId: workspaceId,
		PipelineId:  pipelineId,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func GetPipelineRecord(workspaceId string, pipelineId string, id string) {
	res, err := DefaultInstance.GetPipelineRecord(&GetPipelineRecordRequest{
		WorkspaceId: workspaceId,
		PipelineId:  pipelineId,
		Id:          id,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestCodePipeline_CreatePipeline(t *testing.T) {
	CreatePipeline("620f4b987c14263ff59168f0", Pipeline{
		Name: "create-pipeline-api-test",
		Scm: SCM{
			Disabled: true,
		},
		Stages: []Stage{
			{
				Id:   "99c2c7eb-ed08-44fb-a72c-57880d0fb6f3",
				Name: "Stage One",
				Tasks: []*Task{
					{
						Id:   "20704e25-b0ec-41e4-aedc-1594364fd6a6",
						Name: "Task One",
						Type: "ExecCmd",
						Steps: []Step{
							{
								Id:       "41c14899-833f-4b61-be30-1694202557f0",
								Name:     "Step One",
								Type:     "ExecCmd",
								Language: "",
								Params: []KVPair{
									{
										Key:   "k1",
										Value: "v1",
									},
								},
							},
						},
					},
				},
			},
		},
		ClusterPool: "public",
	})
}

func DeletePipeline(workspaceId string, id string) {
	res, err := DefaultInstance.DeletePipeline(&DeletePipelineRequest{
		WorkspaceId: workspaceId,
		Id:          id,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestCodePipeline_ListPipelines(t *testing.T) {
	ListPipelines("620f4b987c14263ff59168f0")
}

func TestCodePipeline_GetPipeline(t *testing.T) {
	GetPipeline("620f4b987c14263ff59168f0", "620f52eb7c14263ff59168f4")
}

func TestCodePipeline_RunPipeline(t *testing.T) {
	RunPipeline("620f4b987c14263ff59168f0", "620f52eb7c14263ff59168f4")
}

func TestCodePipeline_DeletePipeline(t *testing.T) {
	DeletePipeline("620f4b987c14263ff59168f0", "62172d1d46bea1167b7a30b1")
}

func TestCodePipeline_ListPipelineRecords(t *testing.T) {
	ListPipelineRecords("620f4b987c14263ff59168f0", "620f52eb7c14263ff59168f4")
}

func TestCodePipeline_GetPipelineRecord(t *testing.T) {
	GetPipelineRecord("620f4b987c14263ff59168f0", "620f52eb7c14263ff59168f4", "6217295146bea1167b7a30ad")
}

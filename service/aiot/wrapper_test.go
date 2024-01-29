package aiot

import "testing"

func TestSample(t *testing.T) {
	// this is just an example
	ak := "************"
	sk := "*****************=="
	host := "open.volcengineapi.com"
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
	DefaultInstance.SetHost(host)

	createSpace := &CreateSpaceRequest{
		AccessType: "gb28181",
		GB: struct {
			PullOnDemand bool `json:"PullOnDemand"`
		}{
			PullOnDemand: true,
		},
		Region:    "cn-beijing-a",
		SpaceName: "golang-sdk-create",
	}
	resp, _, err := DefaultInstance.CreateSpace(createSpace)
	if err != nil {
		t.Logf("error occur %v", err)
	}
	t.Logf("%+v", resp)

	//other action
}

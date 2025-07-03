package businessSecurity

import (
	"fmt"
	"testing"
)

const (
	Ak = "" // fill in your access key
	Sk = "" // fill in your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)

	SecuritySourceInstance.Client.SetAccessKey(Ak)
	SecuritySourceInstance.Client.SetSecretKey(Sk)

}
func TestOpenProduct(t *testing.T) {
	resp, err := DefaultInstance.OpenProduct(&OpenProductReq{
		PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestCheckProduct(t *testing.T) {
	resp, err := DefaultInstance.CheckProductStatus(&CheckProductStatusReq{
		PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestOpenService(t *testing.T) {
	resp, err := DefaultInstance.EnableService(&EnableServiceReq{
		Service: ServiceContentTextRisk,
		AppID:   0,
		Product: PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestCheckServiceStatus(t *testing.T) {
	resp, err := DefaultInstance.CheckServiceStatus(&CheckServiceStatusReq{
		Service: ServiceContentTextRisk,
		AppID:   0,
		Product: PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestBusinessSecurity_ActivateRiskBasePackage(t *testing.T) {
	req := &ActivateRiskBasePackageReq{
		PackageId:       "Ab123456",
		TotalPackageNum: 1,
		PackageSeq:      1,
		DataType:        "1",
		Data:            []string{"cfcd208495d565ef66e7dff9f98764da"},
	}
	resp, err := DefaultInstance.ActivateRiskBasePackage(req)
	fmt.Println(resp, err)
}

func TestBusinessSecurity_ActivateRiskSampleData(t *testing.T) {
	list := make([]SampleData, 0)
	list = append(list, SampleData{
		Id:           "cfcd208495d565ef66e7dff9f98764da",
		ReachType:    "X",
		LaunchStatus: "2",
	})
	req := &ActivateRiskSampleDataReq{
		PackageId:       "A11",
		TotalPackageNum: 1,
		PackageSeq:      1,
		DataType:        "1",
		BusinessType:    "A1",
		Data:            list,
	}
	resp, err := DefaultInstance.ActivateRiskSampleData(req)
	fmt.Println(resp, err)
}

func TestBusinessSecurity_ActivateRiskResult(t *testing.T) {

	req := &ActivateRiskResultReq{
		ActivateCode: "a00c7e86-7b42-494c-87a8-0c8c2aaa819c",
		PlanId:       11,
	}

	resp, err := DefaultInstance.ActivateRiskResult(req)

	fmt.Println(resp, err)

}

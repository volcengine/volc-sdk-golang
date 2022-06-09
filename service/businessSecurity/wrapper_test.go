package businessSecurity

import (
	"fmt"
	"testing"
)

const (
	Ak = "AK" // write your access key
	Sk = "SK" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func RiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.RiskDetection(&RiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TextRiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.TextRisk(&RiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func AsyncImageRiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.AsyncImageRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func AsyncRiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.AsyncRiskDetection(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func GetImageResult(dataId, service string, appId int64) (*ImageResultResponse, error) {
	return DefaultInstance.GetImageResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func ImageRisk(appId int64, service string, parameters string) (*ImageResultResponse, error) {
	return DefaultInstance.ImageContentRisk(&RiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func AsyncVideoRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncVideoRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetVideoResult(dataId, service string, appId int64) (*VideoResultResponse, error) {
	return DefaultInstance.GetVideoLiveResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func RiskResult(appId int64, service string, startTime, endTime, pageSize, pageNum int64) {
	res, err := DefaultInstance.RiskResult(&RiskResultRequest{
		AppId:     appId,   // write your app id
		Service:   service, // write business security service
		StartTime: startTime,
		EndTime:   endTime,
		Page: Page{
			PageSize: pageSize,
			PageNum:  pageNum,
		},
	})
	fmt.Println(err)
	if res != nil {
		fmt.Printf("result %+v\n", *res)
	}
}

func DataReport(appId int64, service string, parameters string) {
	res, err := DefaultInstance.DataReport(&DataReportRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func AsyncAudioRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncAudioRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetAudioResult(dataId, service string, appId int64) (*AudioResultResponse, error) {
	return DefaultInstance.GetAudioResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func AsyncVideoLiveRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncLiveVideoRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetVideoLiveResult(dataId, service string, appId int64) (*VideoResultResponse, error) {
	return DefaultInstance.GetVideoLiveResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func AsyncAudioLiveRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncLiveAudioRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetAudioLiveResult(dataId, service string, appId int64) (*AudioResultResponse, error) {
	return DefaultInstance.GetAudioLiveResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func NewCustomContents(appId int64, service, name, description, decision string, matchType int) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.CreateCustomContents(&NewCustomContentsReq{
		AppID:       appId,
		Service:     service,
		Name:        name,
		Description: description,
		Decision:    decision,
		MatchType:   matchType,
	})
}

func EnableCustomContents(appId int64, name string) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.EnableCustomContents(&UpdateContentReq{
		AppID: appId,
		Name:  name,
	})
}

func DisableCustomContents(appId int64, name string) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.DisableCustomContents(&UpdateContentReq{
		AppID: appId,
		Name:  name,
	})
}

func DeleteCustomContents(appId int64, name string) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.DeleteCustomContents(&UpdateContentReq{
		AppID: appId,
		Name:  name,
	})
}

func UploadCustomContents(appId int64, name string, contents []string, modifyType int) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.UploadCustomContents(&UpdateContentReq{
		AppID: appId,
		Name:  name,
	})
}

func TestBusinessSecurity_RiskResult(t *testing.T) {
	RiskResult(3332, "login", 1615535000, 1615540603, 10, 1)
}

func TestBusinessSecurity_RiskDetection(t *testing.T) {
	RiskDetection(3332, "login", "{\"operate_time\": 1615540603, \"uid\":123444}")
}

func ElementVerify(appId int64, service string, parameters string) {
	res, err := DefaultInstance.ElementVerify(&ElementVerifyRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}
func TestBusinessSecurity_ElementVerify(t *testing.T) {
	ElementVerify(3332, "idcard_two_element_verify", "{\"operate_time\": 1615540603, \"idcard_no\": \"\", \"idcard_name\":\"\"}")
}

func TestBusinessSecurity_ElementVerifyV2(t *testing.T) {
	appId := int64(3332)
	service := "mobile_two_element_verify"
	parameters := "{\"operate_time\": 1635321212,\"mobile\":\"\",\"idcard_name\":\"\"}"

	res, err := DefaultInstance.ElementVerifyV2(&ElementVerifyRequest{
		AppId:      appId,
		Service:    service,
		Parameters: parameters,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *res)
}

func TestBusinessSecurity_MobileStatus(t *testing.T) {
	appId := int64(3332)
	service := "mobile_empty_status"
	parameters := "{\"operate_time\":1617960951,\"mobile\":\"12312341234\"}"

	res, err := DefaultInstance.MobileStatus(&MobileStatusRequest{
		AppId:      appId,
		Service:    service,
		Parameters: parameters,
	})
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%v", *res)
}

func TestBusinessSecurity_MobileStatusV2(t *testing.T) {
	appId := int64(3332)
	service := "mobile_address"
	parameters := "{\"operate_time\":1617960951,\"mobile\":\"12312341234\"}"

	res, err := DefaultInstance.MobileStatusV2(&MobileStatusRequest{
		AppId:      appId,
		Service:    service,
		Parameters: parameters,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *res)
}

func TestNewCustomContents(t *testing.T) {
	resp, err := NewCustomContents(5461, "text_risk", "text", "text", "BLOCK", 1)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}

func TestEnableCustomContents(t *testing.T) {
	resp, err := UploadCustomContents(5461, "text", []string{"test1", "test2"}, 0)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}

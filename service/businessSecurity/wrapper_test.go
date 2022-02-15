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

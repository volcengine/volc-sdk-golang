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
		DataId: dataId,
		AppId: appId,
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
		DataId: dataId,
		AppId: appId,
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
		DataId: dataId,
		AppId: appId,
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
		DataId: dataId,
		AppId: appId,
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
		DataId: dataId,
		AppId: appId,
		Service: service,
	})
}

func TestBusinessSecurity_RiskResult(t *testing.T) {
	RiskResult(3332, "login", 1615535000, 1615540603, 10, 1)
}

func TestBusinessSecurity_RiskDetection(t *testing.T) {
	RiskDetection(3332, "login", "{\"operate_time\": 1615540603, \"uid\":123444}")
}
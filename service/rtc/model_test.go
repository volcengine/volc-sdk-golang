package rtc

import (
	"encoding/json"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"

	"github.com/stretchr/testify/assert"
)

func TestListRoomsResponse(t *testing.T) {
	expectedJson := `{
    "ResponseMetadata": {
        "RequestId": "Your_RequestId",
        "Action": "ListRooms",
        "Version": "2020-12-01",
        "Service": "rtc",
        "Region": "cn-north-1"
    },
    "Result": {
        "Total": 100,
        "ActiveNum": 20,
        "InactiveNum": 80,
        "Offset": 2,
        "Limit": 10,
        "Rooms": [
            {
            "RoomId": "lufei",
            "AppId": "test",
            "UserNum": 4,
            "StreamNum": 2,
            "State": 1,
            "CreatedAt": "2021-03-16T11:35:25+08:00",
            "UpdatedAt": "2021-03-16T11:35:27+08:00"
            }
        ]
    }
}`

	expectedModel := ListRoomsResponse{}
	err := json.Unmarshal([]byte(expectedJson), &expectedModel)
	assert.Nil(t, err)

	commonResponse := base.ResponseMetadata{
		RequestId: "Your_RequestId",
		Service:   "rtc",
		Region:    "cn-north-1",
		Action:    "ListRooms",
		Version:   "2020-12-01",
		Error:     nil,
	}
	room := []Room{Room{RoomId: "lufei", AppId: "test", UserNum: 4, StreamNum: 2, State: 1, CreatedAt: "2021-03-16T11:35:25+08:00", UpdatedAt: "2021-03-16T11:35:27+08:00"}}
	result := ListRoomsResult{
		Total:       100,
		ActiveNum:   20,
		InactiveNum: 80,
		Offset:      2,
		Limit:       10,
		Rooms:       room,
	}
	response := ListRoomsResponse{
		ResponseMetadata: &commonResponse,
		Result:           &result,
	}
	assert.Equal(t, response, expectedModel)
}

func TestListIndicatorsRequest(t *testing.T) {
	expectedJson := `{
    "AppId" : "Your_AppId",
    "StartTime" : "2021-07-24T08:00:00+08:00",
    "EndTime" : "2021-07-28T08:00:00+08:00",
    "Indicator":"NetworkTransDelay",
    "OS":"android",
    "Network":"4g"
    }`
	expectedModel := ListIndicatorsRequest{}
	err := json.Unmarshal([]byte(expectedJson), &expectedModel)
	assert.Nil(t, err)
	r := ListIndicatorsRequest{
		AppId:     "Your_AppId",
		StartTime: "2021-07-24T08:00:00+08:00",
		EndTime:   "2021-07-28T08:00:00+08:00",
		Indicator: "NetworkTransDelay",
		OS:        "android",
		Network:   "4g",
	}
	assert.Equal(t, r, expectedModel)
}

func TestListIndicatorsResponse(t *testing.T) {
	expectedJson := `{
    "ResponseMetadata":{
        "RequestId":"Your_RequestId",
        "Action":"ListIndicators",
        "Version":"2020-12-01",
        "Service":"rtc",
        "Region":"cn-north-1"
    },
    "Result":{
        "Indicators":[
            {
                "Name":"100ms传输延时达标率",
                "Unit":"%",
                "Data":[
                    {
                        "TimeStamp":"2021-07-24T00:00:00+08:00",
                        "Value":98.68
                    },
                    {
                        "TimeStamp":"2021-07-25T00:00:00+08:00",
                        "Value":98.20
                    }
                ]
            },
            {
                "Name":"200ms传输延时达标率",
                "Unit":"%",
                "Data":[
                    {
                        "TimeStamp":"2021-07-24T00:00:00+08:00",
                        "Value":99.68
                    },
                    {
                        "TimeStamp":"2021-07-25T00:00:00+08:00",
                        "Value":99.20
                    }
                ]
            }
        ]
    }
    }`
	expectedModel := ListIndicatorsResponse{}
	err := json.Unmarshal([]byte(expectedJson), &expectedModel)
	assert.Nil(t, err)
	commonResponse := base.ResponseMetadata{
		RequestId: "Your_RequestId",
		Service:   "rtc",
		Region:    "cn-north-1",
		Action:    "ListIndicators",
		Version:   "2020-12-01",
		Error:     nil,
	}
	indicators := []Indicator{
		{
			Name: "100ms传输延时达标率",
			Unit: "%",
			Datas: []Data{
				{
					TimeStamp: "2021-07-24T00:00:00+08:00",
					Value:     98.68,
				}, {
					TimeStamp: "2021-07-25T00:00:00+08:00",
					Value:     98.20,
				}},
		},
		{
			Name: "200ms传输延时达标率",
			Unit: "%",
			Datas: []Data{{
				TimeStamp: "2021-07-24T00:00:00+08:00",
				Value:     99.68,
			}, {
				TimeStamp: "2021-07-25T00:00:00+08:00",
				Value:     99.20,
			}},
		}}
	response := ListIndicatorsResponse{
		ResponseMetadata: &commonResponse,
		Result: &ListIndicatorsResult{
			Indicators: indicators,
		},
	}
	assert.Equal(t, response, expectedModel)
}


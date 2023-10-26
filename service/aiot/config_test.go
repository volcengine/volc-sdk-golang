package aiot

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"testing"
)

func TestMergeApiConfig(t *testing.T) {
	ApiInfoList = map[string]*base.ApiInfo{
		"DeleteSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteSpace"}, "Version": []string{ServiceVersion20231001}},
		},
		"GetRecordList": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetRecordList"}, "Version": []string{ServiceVersion20231001}},
		},
	}

	ApiInfoListV3 = map[string]*base.ApiInfo{
		"DeleteSpaceV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query:  url.Values{"Action": []string{"DeleteSpace"}, "Version": []string{ServiceVersion20231001}},
		},
		"GetRecordListV3": {
			Method: http.MethodGet,
			Path:   "/",
			Query:  url.Values{"Action": []string{"GetRecordList"}, "Version": []string{ServiceVersion20231001}},
		},
	}

	_, err := MergeApiConfig()
	assert.NotNil(t, err)

	noMergeApiConfig, _ := MergeApiConfig(ApiInfoList)
	assert.NotNil(t, noMergeApiConfig)
	assert.Equal(t, 2, len(noMergeApiConfig))
	assert.NotNil(t, noMergeApiConfig["DeleteSpace"])
	assert.NotNil(t, noMergeApiConfig["GetRecordList"])

	newApiConfig, _ := MergeApiConfig(ApiInfoList, ApiInfoListV3)
	assert.NotNil(t, newApiConfig)
	assert.Equal(t, 4, len(newApiConfig))
	assert.NotNil(t, newApiConfig["DeleteSpace"])
	assert.NotNil(t, newApiConfig["GetRecordList"])
	assert.NotNil(t, newApiConfig["DeleteSpaceV3"])
	assert.NotNil(t, newApiConfig["GetRecordListV3"])
}

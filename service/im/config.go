package im

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "rtc"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-north-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "rtc.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-north-1",
				Service: ServiceName,
			},
		},
		"ap-southeast-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "rtc.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "ap-southeast-1",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"GetConversationMarks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetConversationMarks"},
				"Version": []string{"2020-12-01"},
			},
		},
		"MarkConversation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MarkConversation"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetMessagesReadReceipt": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMessagesReadReceipt"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ModifyParticipantReadIndex": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyParticipantReadIndex"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ScanConversationParticipantList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ScanConversationParticipantList"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchGetBlockParticipants": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchGetBlockParticipants"},
				"Version": []string{"2020-12-01"},
			},
		},
		"IsUserInConversation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"IsUserInConversation"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchModifyConversationParticipant": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchModifyConversationParticipant"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchDeleteConversationParticipant": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchDeleteConversationParticipant"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchDeleteBlockParticipants": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchDeleteBlockParticipants"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchGetConversationParticipant": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchGetConversationParticipant"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchGetWhitelistParticipant": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchGetWhitelistParticipant"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchAddWhitelistParticipant": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchAddWhitelistParticipant"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchAddManager": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchAddManager"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchAddConversationParticipant": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchAddConversationParticipant"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchAddBlockParticipants": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchAddBlockParticipants"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchRemoveWhitelistParticipant": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchRemoveWhitelistParticipant"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchRemoveManager": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchRemoveManager"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchUpdateLiveParticipants": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchUpdateLiveParticipants"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetParticipantReadIndex": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetParticipantReadIndex"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetConversationUserCount": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetConversationUserCount"},
				"Version": []string{"2020-12-01"},
			},
		},
		"QueryLiveParticipantStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryLiveParticipantStatus"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ModifyConversation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyConversation"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ModifyConversationSetting": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyConversationSetting"},
				"Version": []string{"2020-12-01"},
			},
		},
		"CreateConversation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateConversation"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchGetConversations": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchGetConversations"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetConversationSetting": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetConversationSetting"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetUserConversations": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetUserConversations"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DestroyConversation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DestroyConversation"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ModifyMessage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyMessage"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetConversationMessages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetConversationMessages"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DeleteConversationMessage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteConversationMessage"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DeleteMessage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteMessage"},
				"Version": []string{"2020-12-01"},
			},
		},
		"SendMessage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendMessage"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetMessages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMessages"},
				"Version": []string{"2020-12-01"},
			},
		},
		"RecallMessage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RecallMessage"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DeleteFriend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteFriend"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateFriend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateFriend"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateBlackList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateBlackList"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListFriend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListFriend"},
				"Version": []string{"2020-12-01"},
			},
		},
		"QueryOnlineStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryOnlineStatus"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetBlackList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetBlackList"},
				"Version": []string{"2020-12-01"},
			},
		},
		"IsFriend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"IsFriend"},
				"Version": []string{"2020-12-01"},
			},
		},
		"IsInBlackList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"IsInBlackList"},
				"Version": []string{"2020-12-01"},
			},
		},
		"AddFriend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddFriend"},
				"Version": []string{"2020-12-01"},
			},
		},
		"AddBlackList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddBlackList"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetAppToken": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAppToken"},
				"Version": []string{"2020-12-01"},
			},
		},
		"RemoveBlackList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemoveBlackList"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UserBroadcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UserBroadcast"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchUpdateUserTags": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchUpdateUserTags"},
				"Version": []string{"2020-12-01"},
			},
		},
		"RegisterUsers": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RegisterUsers"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UnRegisterUsers": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnRegisterUsers"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchUpdateUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchUpdateUser"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchGetUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchGetUser"},
				"Version": []string{"2020-12-01"},
			},
		},
	}
)

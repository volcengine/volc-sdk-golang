package im

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"

	common "github.com/volcengine/volc-sdk-golang/base"
)

type queryMarshalFilter func(key string, value []string) (accept bool)

func skipEmptyValue() queryMarshalFilter {
	return func(_ string, value []string) (accept bool) {
		if len(value) == 0 {
			return false
		}

		for _, item := range value {
			if item == "" {
				return false
			}
		}

		return true
	}
}

func marshalToQuery(model interface{}, filters ...queryMarshalFilter) (url.Values, error) {
	ret := url.Values{}
	if model == nil {
		return ret, nil
	}

	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)
	if modelValue.IsNil() {
		return ret, nil
	}

	if modelType.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
		modelType = modelValue.Type()
	} else {
		return ret, fmt.Errorf("not struct")
	}
	fieldCount := modelType.NumField()

	for i := 0; i < fieldCount; i++ {
		field := modelType.Field(i)
		queryKey := field.Tag.Get("query")
		if queryKey == "" {
			continue
		}

		queryRawValue := modelValue.FieldByName(field.Name)
		queryValues := make([]string, 0)
		if field.Type.Kind() != reflect.Array && field.Type.Kind() != reflect.Slice {
			value := resolveQueryValue(queryRawValue)
			if value == nil {
				continue
			}
			queryValues = append(queryValues, fmt.Sprintf("%v", value))
		} else {
			length := queryRawValue.Len()
			for idx := 0; idx < length; idx++ {
				value := resolveQueryValue(queryRawValue.Index(idx))
				if value == nil {
					continue
				}
				queryValues = append(queryValues, fmt.Sprintf("%v", value))
			}
		}

		for _, fun := range filters {
			if !fun(queryKey, queryValues) {
				goto afterAddQuery
			}
		}

		for _, t := range queryValues {
			ret.Add(queryKey, t)
		}
	afterAddQuery:
	}

	return ret, nil
}

func resolveQueryValue(queryRawValue reflect.Value) interface{} {
	if queryRawValue.Kind() == reflect.Ptr {
		if queryRawValue.IsNil() {
			return nil
		}
		decayedQueryRawValue := queryRawValue.Elem()
		decayedReflectValue := decayedQueryRawValue.Interface()
		return fmt.Sprintf("%v", decayedReflectValue)
	} else {
		queryReflectValue := queryRawValue.Interface()
		return fmt.Sprintf("%v", queryReflectValue)
	}
}

func marshalToJson(model interface{}) ([]byte, error) {
	if model == nil {
		return make([]byte, 0), nil
	}

	result, err := json.Marshal(model)
	if err != nil {
		return []byte{}, fmt.Errorf("can not marshal model to json, %v", err)
	}
	return result, nil
}

func unmarshalResultInto(data []byte, result interface{}) error {
	resp := new(common.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

func (c *Im) ModifyParticipantReadIndex(ctx context.Context, arg *ModifyParticipantReadIndexBody) (*ModifyParticipantReadIndexRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ModifyParticipantReadIndex", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ModifyParticipantReadIndexRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) ScanConversationParticipantList(ctx context.Context, arg *ScanConversationParticipantListBody) (*ScanConversationParticipantListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ScanConversationParticipantList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ScanConversationParticipantListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchGetBlockParticipants(ctx context.Context, arg *BatchGetBlockParticipantsBody) (*BatchGetBlockParticipantsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchGetBlockParticipants", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchGetBlockParticipantsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) IsUserInConversation(ctx context.Context, arg *IsUserInConversationBody) (*IsUserInConversationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "IsUserInConversation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(IsUserInConversationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchModifyConversationParticipant(ctx context.Context, arg *BatchModifyConversationParticipantBody) (*BatchModifyConversationParticipantRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchModifyConversationParticipant", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchModifyConversationParticipantRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchDeleteConversationParticipant(ctx context.Context, arg *BatchDeleteConversationParticipantBody) (*BatchDeleteConversationParticipantRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchDeleteConversationParticipant", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchDeleteConversationParticipantRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchDeleteBlockParticipants(ctx context.Context, arg *BatchDeleteBlockParticipantsBody) (*BatchDeleteBlockParticipantsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchDeleteBlockParticipants", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchDeleteBlockParticipantsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchGetConversationParticipant(ctx context.Context, arg *BatchGetConversationParticipantBody) (*BatchGetConversationParticipantRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchGetConversationParticipant", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchGetConversationParticipantRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchGetWhitelistParticipant(ctx context.Context, arg *BatchGetWhitelistParticipantBody) (*BatchGetWhitelistParticipantRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchGetWhitelistParticipant", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchGetWhitelistParticipantRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchAddWhitelistParticipant(ctx context.Context, arg *BatchAddWhitelistParticipantBody) (*BatchAddWhitelistParticipantRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchAddWhitelistParticipant", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchAddWhitelistParticipantRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchAddManager(ctx context.Context, arg *BatchAddManagerBody) (*BatchAddManagerRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchAddManager", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchAddManagerRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchAddConversationParticipant(ctx context.Context, arg *BatchAddConversationParticipantBody) (*BatchAddConversationParticipantRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchAddConversationParticipant", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchAddConversationParticipantRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchAddBlockParticipants(ctx context.Context, arg *BatchAddBlockParticipantsBody) (*BatchAddBlockParticipantsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchAddBlockParticipants", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchAddBlockParticipantsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchRemoveWhitelistParticipant(ctx context.Context, arg *BatchRemoveWhitelistParticipantBody) (*BatchRemoveWhitelistParticipantRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchRemoveWhitelistParticipant", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchRemoveWhitelistParticipantRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchRemoveManager(ctx context.Context, arg *BatchRemoveManagerBody) (*BatchRemoveManagerRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchRemoveManager", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchRemoveManagerRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchUpdateLiveParticipants(ctx context.Context, arg *BatchUpdateLiveParticipantsBody) (*BatchUpdateLiveParticipantsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchUpdateLiveParticipants", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchUpdateLiveParticipantsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) GetParticipantReadIndex(ctx context.Context, arg *GetParticipantReadIndexBody) (*GetParticipantReadIndexRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetParticipantReadIndex", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetParticipantReadIndexRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) GetConversationUserCount(ctx context.Context, arg *GetConversationUserCountBody) (*GetConversationUserCountRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetConversationUserCount", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetConversationUserCountRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) QueryLiveParticipantStatus(ctx context.Context, arg *QueryLiveParticipantStatusBody) (*QueryLiveParticipantStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "QueryLiveParticipantStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(QueryLiveParticipantStatusRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) ModifyConversation(ctx context.Context, arg *ModifyConversationBody) (*ModifyConversationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ModifyConversation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ModifyConversationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) ModifyConversationSetting(ctx context.Context, arg *ModifyConversationSettingBody) (*ModifyConversationSettingRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ModifyConversationSetting", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ModifyConversationSettingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) CreateConversation(ctx context.Context, arg *CreateConversationBody) (*CreateConversationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateConversation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateConversationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) BatchGetConversations(ctx context.Context, arg *BatchGetConversationsBody) (*BatchGetConversationsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchGetConversations", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchGetConversationsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) GetConversationSetting(ctx context.Context, arg *GetConversationSettingBody) (*GetConversationSettingRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetConversationSetting", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetConversationSettingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) GetUserConversations(ctx context.Context, arg *GetUserConversationsBody) (*GetUserConversationsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetUserConversations", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetUserConversationsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) DestroyConversation(ctx context.Context, arg *DestroyConversationBody) (*DestroyConversationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DestroyConversation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DestroyConversationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) ModifyMessage(ctx context.Context, arg *ModifyMessageBody) (*ModifyMessageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ModifyMessage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ModifyMessageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) GetConversationMessages(ctx context.Context, arg *GetConversationMessagesBody) (*GetConversationMessagesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetConversationMessages", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetConversationMessagesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) DeleteConversationMessage(ctx context.Context, arg *DeleteConversationMessageBody) (*DeleteConversationMessageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteConversationMessage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteConversationMessageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) DeleteMessage(ctx context.Context, arg *DeleteMessageBody) (*DeleteMessageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteMessage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteMessageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) SendMessage(ctx context.Context, arg *SendMessageBody) (*SendMessageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SendMessage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SendMessageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) GetMessages(ctx context.Context, arg *GetMessagesBody) (*GetMessagesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetMessages", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetMessagesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) RecallMessage(ctx context.Context, arg *RecallMessageBody) (*RecallMessageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RecallMessage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RecallMessageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) DeleteFriend(ctx context.Context, arg *DeleteFriendBody) (*DeleteFriendRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteFriend", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteFriendRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) UpdateFriend(ctx context.Context, arg *UpdateFriendBody) (*UpdateFriendRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateFriend", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateFriendRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) UpdateBlackList(ctx context.Context, arg *UpdateBlackListBody) (*UpdateBlackListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateBlackList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateBlackListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) ListFriend(ctx context.Context, arg *ListFriendBody) (*ListFriendRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListFriend", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListFriendRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) QueryOnlineStatus(ctx context.Context, arg *QueryOnlineStatusBody) (*QueryOnlineStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "QueryOnlineStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(QueryOnlineStatusRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) GetBlackList(ctx context.Context, arg *GetBlackListBody) (*GetBlackListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetBlackList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetBlackListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) IsFriend(ctx context.Context, arg *IsFriendBody) (*IsFriendRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "IsFriend", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(IsFriendRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) IsInBlackList(ctx context.Context, arg *IsInBlackListBody) (*IsInBlackListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "IsInBlackList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(IsInBlackListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) AddFriend(ctx context.Context, arg *AddFriendBody) (*AddFriendRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddFriend", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddFriendRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) AddBlackList(ctx context.Context, arg *AddBlackListBody) (*AddBlackListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddBlackList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddBlackListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Im) RemoveBlackList(ctx context.Context, arg *RemoveBlackListBody) (*RemoveBlackListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RemoveBlackList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RemoveBlackListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

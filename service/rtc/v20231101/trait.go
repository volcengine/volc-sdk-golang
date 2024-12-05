package rtc_v20231101

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

type queryMarshalFilter func(key string, value []string) (accept bool)

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

func (c *Rtc) LimitTokenPrivilege(ctx context.Context, arg *LimitTokenPrivilegeBody) (*LimitTokenPrivilegeRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "LimitTokenPrivilege", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(LimitTokenPrivilegeRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) BanRoomUser(ctx context.Context, arg *BanRoomUserBody) (*BanRoomUserRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "BanRoomUser", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(BanRoomUserRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) BanUserStream(ctx context.Context, arg *BanUserStreamBody) (*BanUserStreamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "BanUserStream", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(BanUserStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UpdateBanRoomUserRule(ctx context.Context, arg *UpdateBanRoomUserRuleBody) (*UpdateBanRoomUserRuleRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdateBanRoomUserRule", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UpdateBanRoomUserRuleRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) GetRoomOnlineUsers(ctx context.Context, arg *GetRoomOnlineUsersQuery) (*GetRoomOnlineUsersRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetRoomOnlineUsers", query)

	if len(data) > 0 {
		result := new(GetRoomOnlineUsersRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UnbanUserStream(ctx context.Context, arg *UnbanUserStreamBody) (*UnbanUserStreamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UnbanUserStream", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UnbanUserStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopPushPublicStream(ctx context.Context, arg *StopPushPublicStreamBody) (*StopPushPublicStreamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopPushPublicStream", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopPushPublicStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopRelayStream(ctx context.Context, arg *StopRelayStreamBody) (*StopRelayStreamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopRelayStream", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopRelayStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopSnapshot(ctx context.Context, arg *StopSnapshotBody) (*StopSnapshotRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopSnapshot", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopSnapshotRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartPushPublicStream(ctx context.Context, arg *StartPushPublicStreamBody) (*StartPushPublicStreamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartPushPublicStream", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartPushPublicStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartWebcast(ctx context.Context, arg *StartWebcastBody) (*StartWebcastRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartWebcast", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartWebcastRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartSnapshot(ctx context.Context, arg *StartSnapshotBody) (*StartSnapshotRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartSnapshot", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartSnapshotRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartSegment(ctx context.Context, arg *StartSegmentBody) (*StartSegmentRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartSegment", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartSegmentRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartRecord(ctx context.Context, arg *StartRecordBody) (*StartRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartRecord", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartRecordRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartPushSingleStreamToCDN(ctx context.Context, arg *StartPushSingleStreamToCDNBody) (*StartPushSingleStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartPushSingleStreamToCDN", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartPushSingleStreamToCDNRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartPushMixedStreamToCDN(ctx context.Context, arg *StartPushMixedStreamToCDNBody) (*StartPushMixedStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartPushMixedStreamToCDN", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartPushMixedStreamToCDNRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartRelayStream(ctx context.Context, arg *StartRelayStreamBody) (*StartRelayStreamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartRelayStream", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartRelayStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UpdateRecord(ctx context.Context, arg *UpdateRecordBody) (*UpdateRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdateRecord", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UpdateRecordRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UpdatePublicStreamParam(ctx context.Context, arg *UpdatePublicStreamParamBody) (*UpdatePublicStreamParamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdatePublicStreamParam", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UpdatePublicStreamParamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UpdatePushMixedStreamToCDN(ctx context.Context, arg *UpdatePushMixedStreamToCDNBody) (*UpdatePushMixedStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdatePushMixedStreamToCDN", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UpdatePushMixedStreamToCDNRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UpdateRelayStream(ctx context.Context, arg *UpdateRelayStreamBody) (*UpdateRelayStreamRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdateRelayStream", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UpdateRelayStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UpdateSnapshot(ctx context.Context, arg *UpdateSnapshotBody) (*UpdateSnapshotRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdateSnapshot", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UpdateSnapshotRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) UpdateSegment(ctx context.Context, arg *UpdateSegmentBody) (*UpdateSegmentRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdateSegment", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(UpdateSegmentRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) GetWebCastTask(ctx context.Context, arg *GetWebCastTaskQuery) (*GetWebCastTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetWebCastTask", query)

	if len(data) > 0 {
		result := new(GetWebCastTaskRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) GetPushSingleStreamToCDNTask(ctx context.Context, arg *GetPushSingleStreamToCDNTaskQuery) (*GetPushSingleStreamToCDNTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetPushSingleStreamToCDNTask", query)

	if len(data) > 0 {
		result := new(GetPushSingleStreamToCDNTaskRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) GetPushMixedStreamToCDNTask(ctx context.Context, arg *GetPushMixedStreamToCDNTaskQuery) (*GetPushMixedStreamToCDNTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetPushMixedStreamToCDNTask", query)

	if len(data) > 0 {
		result := new(GetPushMixedStreamToCDNTaskRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) GetRecordTask(ctx context.Context, arg *GetRecordTaskQuery) (*GetRecordTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetRecordTask", query)

	if len(data) > 0 {
		result := new(GetRecordTaskRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) GetSnapshotTask(ctx context.Context, arg *GetSnapshotTaskQuery) (*GetSnapshotTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetSnapshotTask", query)

	if len(data) > 0 {
		result := new(GetSnapshotTaskRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListRelayStream(ctx context.Context, arg *ListRelayStreamQuery) (*ListRelayStreamRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "ListRelayStream", query)

	if len(data) > 0 {
		result := new(ListRelayStreamRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) GetSegmentTask(ctx context.Context, arg *GetSegmentTaskQuery) (*GetSegmentTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetSegmentTask", query)

	if len(data) > 0 {
		result := new(GetSegmentTaskRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopWebcast(ctx context.Context, arg *StopWebcastBody) (*StopWebcastRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopWebcast", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopWebcastRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopRecord(ctx context.Context, arg *StopRecordBody) (*StopRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopRecord", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopRecordRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopPushStreamToCDN(ctx context.Context, arg *StopPushStreamToCDNBody) (*StopPushStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopPushStreamToCDN", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopPushStreamToCDNRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopSegment(ctx context.Context, arg *StopSegmentBody) (*StopSegmentRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopSegment", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopSegmentRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) CreateApp(ctx context.Context, arg *CreateAppBody) (*CreateAppRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "CreateApp", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(CreateAppRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ModifyAppStatus(ctx context.Context, arg *ModifyAppStatusBody) (*ModifyAppStatusRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ModifyAppStatus", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ModifyAppStatusRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListApps(ctx context.Context, arg *ListAppsBody) (*ListAppsRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListApps", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListAppsRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartWBRecord(ctx context.Context, arg *StartWBRecordBody) (*StartWBRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartWBRecord", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartWBRecordRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) WbTranscodeCreate(ctx context.Context, arg *WbTranscodeCreateBody) (*WbTranscodeCreateRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "WbTranscodeCreate", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(WbTranscodeCreateRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopWBRecord(ctx context.Context, arg *StopWBRecordBody) (*StopWBRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopWBRecord", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopWBRecordRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) WbTranscodeQuery(ctx context.Context, arg *WbTranscodeQueryQuery) (*WbTranscodeQueryRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "WbTranscodeQuery", query)

	if len(data) > 0 {
		result := new(WbTranscodeQueryRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) WbTranscodeGet(ctx context.Context, arg *WbTranscodeGetQuery) (*WbTranscodeGetRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "WbTranscodeGet", query)

	if len(data) > 0 {
		result := new(WbTranscodeGetRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListRealTimePublicStreamInfo(ctx context.Context, arg *ListRealTimePublicStreamInfoBody) (*ListRealTimePublicStreamInfoRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListRealTimePublicStreamInfo", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListRealTimePublicStreamInfoRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListRealTimeQualityDistribution(ctx context.Context, arg *ListRealTimeQualityDistributionBody) (*ListRealTimeQualityDistributionRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListRealTimeQualityDistribution", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListRealTimeQualityDistributionRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListRealTimeQuality(ctx context.Context, arg *ListRealTimeQualityBody) (*ListRealTimeQualityRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListRealTimeQuality", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListRealTimeQualityRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListRealTimeOperationData(ctx context.Context, arg *ListRealTimeOperationDataBody) (*ListRealTimeOperationDataRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListRealTimeOperationData", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListRealTimeOperationDataRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListQualityDistribution(ctx context.Context, arg *ListQualityDistributionBody) (*ListQualityDistributionRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListQualityDistribution", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListQualityDistributionRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListQuality(ctx context.Context, arg *ListQualityBody) (*ListQualityRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListQuality", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListQualityRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListOperationDistribution(ctx context.Context, arg *ListOperationDistributionBody) (*ListOperationDistributionRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListOperationDistribution", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListOperationDistributionRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListOperationData(ctx context.Context, arg *ListOperationDataBody) (*ListOperationDataRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListOperationData", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListOperationDataRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListUsages(ctx context.Context, arg *ListUsagesBody) (*ListUsagesRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "ListUsages", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(ListUsagesRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListCallDetail(ctx context.Context, arg *ListCallDetailQuery) (*ListCallDetailRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "ListCallDetail", query)

	if len(data) > 0 {
		result := new(ListCallDetailRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListRoomInfo(ctx context.Context, arg *ListRoomInfoQuery) (*ListRoomInfoRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "ListRoomInfo", query)

	if len(data) > 0 {
		result := new(ListRoomInfoRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListUserInfo(ctx context.Context, arg *ListUserInfoQuery) (*ListUserInfoRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "ListUserInfo", query)

	if len(data) > 0 {
		result := new(ListUserInfoRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StopDetection(ctx context.Context, arg *StopDetectionBody) (*StopDetectionRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopDetection", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StopDetectionRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) ListDetectionTask(ctx context.Context, arg *ListDetectionTaskQuery) (*ListDetectionTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "ListDetectionTask", query)

	if len(data) > 0 {
		result := new(ListDetectionTaskRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) StartDetection(ctx context.Context, arg *StartDetectionBody) (*StartDetectionRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartDetection", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(StartDetectionRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

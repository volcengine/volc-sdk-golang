package wtn_v20231101

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

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

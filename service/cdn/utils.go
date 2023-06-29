package cdn

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

type CDNError struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	Status  int64  `json:"Status"`
}

func (e CDNError) Error() string {
	return fmt.Sprintf("status: %d, code: %s, message: %s", e.Status, e.Code, e.Message)
}

func wrapperError(errObj *ErrorObj) (err error) {
	if errObj == nil {
		return
	}
	err = CDNError{
		Message: errObj.Message,
		Status:  errObj.CodeN,
		Code:    errObj.Code,
	}
	return
}

func validateResponse(meta *ResponseMetadata) (err error) {
	if meta == nil {
		return errors.New("response meta is not found")
	}
	err = wrapperError(meta.Error)
	return
}

type OptionArg struct {
	useGet bool
}

func UseGet() OptionArg {
	return OptionArg{
		useGet: true,
	}
}

func (s *CDN) post(apiName string, requestBody interface{}, responseBody interface{}, options ...OptionArg) (err error) {
	var body []byte
	if requestBody == nil {
		requestBody = struct{}{}
	}
	body, err = json.Marshal(requestBody)
	if err != nil {
		err = fmt.Errorf("marshal request body failed, %w", err)
		return
	}
	query := url.Values{}
	option := OptionArg{useGet: false}
	if len(options) > 0 {
		option = options[0]
		if option.useGet {
			if s.SetMethod(apiName, http.MethodGet) {
				defer s.SetMethod(apiName, http.MethodPost)
			} else {
				err = fmt.Errorf("set method falied")
				return
			}
			query, err = MergeQueryArgs(requestBody, query)
			if err != nil {
				return
			}
		}
	}

	var data []byte
	if option.useGet {
		data, _, err = s.Client.Query(apiName, query)
	} else {
		data, _, err = s.Client.Json(apiName, query, string(body))
	}

	if err != nil {
		err = fmt.Errorf("request %s api failed: %w", apiName, err)
		return
	}
	if err = json.Unmarshal(data, responseBody); err != nil {
		err = fmt.Errorf("unmarshal response body failed: %w", err)
		return
	}
	return
}

func GetStrPtr(str string) *string {
	return &str
}

func MergeQueryArgs(body interface{}, query url.Values) (url.Values, error) {
	t := reflect.TypeOf(body)
	v := reflect.ValueOf(body)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		vi := v.Field(i)
		if vi.Kind() == reflect.Ptr {
			if vi.IsNil() {
				continue
			}
			vi = vi.Elem()
		}
		if !isBaseType(vi.Kind()) {
			if vi.Len() <= 0 {
				continue
			}
			return nil, fmt.Errorf("don't support struct or array")
		}
		query.Set(t.Field(i).Name, fmt.Sprintf("%v", vi.Interface()))
	}
	return query, nil
}

func isBaseType(kind reflect.Kind) bool {
	switch kind {
	case reflect.Bool, reflect.String,
		reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Interface, reflect.Ptr, reflect.Uintptr:
		return false
	case reflect.Struct, reflect.Array, reflect.Map, reflect.Slice:
		return false
	}
	return false
}

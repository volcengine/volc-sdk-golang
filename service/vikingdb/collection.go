package vikingdb

import (
	"context"
	"errors"
	"fmt"
)

type Collection struct {
	CollectionName  string
	Fields          []Field
	VikingDBService *VikingDBService
	PrimaryKey      string
	Indexes         []*Index
	Description     string
	Stat            map[string]interface{}
	CreateTime      string
	UpdateTime      string
	UpdatePerson    string
}

func (collection *Collection) UpsertData(data interface{}, opts ...ParamOption) error {
	options := ParamOptions{
		AsyncUpsert: false,
	}
	for _, opt := range opts {
		opt(&options)
	}
	if _, ok := data.(Data); ok {
		data := data.(Data)
		fieldsArr := []interface{}{data.Fields}
		params := map[string]interface{}{
			"collection_name": collection.CollectionName,
			"fields":          fieldsArr,
		}
		if data.TTL != 0 {
			params["ttl"] = data.TTL
		}
		if options.AsyncUpsert {
			params["async"] = true
		}
		res, err := collection.VikingDBService.retryRequest(context.Background(), "UpsertData", nil, collection.VikingDBService.convertMapToJson(params), MAX_RETRIES)
		_ = res
		return err
	} else if _, ok := data.([]Data); ok {
		datas := data.([]Data)
		record := map[int64][]interface{}{}
		for _, item := range datas {
			if value, exist := record[item.TTL]; exist {
				fields := value
				fields = append(fields, item.Fields)
				record[item.TTL] = fields
			} else {
				record[item.TTL] = []interface{}{item.Fields}
			}
		}
		for index, item := range record {
			params := map[string]interface{}{
				"collection_name": collection.CollectionName,
				"fields":          item,
			}
			if index != 0 {
				params["ttl"] = index
			}
			if options.AsyncUpsert {
				params["async"] = true
			}
			_, err := collection.VikingDBService.retryRequest(context.Background(), "UpsertData", nil, collection.VikingDBService.convertMapToJson(params), MAX_RETRIES)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return errors.New("invalid data")
	}
}

func (collection *Collection) UpdateData(data interface{}, opts ...ParamOption) error {
	if _, ok := data.(Data); ok {
		data := data.(Data)
		fieldsArr := []interface{}{data.Fields}
		params := map[string]interface{}{
			"collection_name": collection.CollectionName,
			"fields":          fieldsArr,
		}
		if data.TTL != 0 {
			params["ttl"] = data.TTL
		}
		_, err := collection.VikingDBService.retryRequest(context.Background(), "UpdateData", nil, collection.VikingDBService.convertMapToJson(params), MAX_RETRIES)
		return err
	} else if _, ok := data.([]Data); ok {
		datas := data.([]Data)
		record := map[int64][]interface{}{}
		for _, item := range datas {
			if value, exist := record[item.TTL]; exist {
				fields := value
				fields = append(fields, item.Fields)
				record[item.TTL] = fields
			} else {
				record[item.TTL] = []interface{}{item.Fields}
			}
		}
		for idx, item := range record {
			params := map[string]interface{}{
				"collection_name": collection.CollectionName,
				"fields":          item,
			}
			if idx != 0 {
				params["ttl"] = idx
			}
			_, err := collection.VikingDBService.retryRequest(context.Background(), "UpdateData", nil, collection.VikingDBService.convertMapToJson(params), MAX_RETRIES)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return errors.New("invalid data")
	}
}

func (collection *Collection) AsyncUpsertData(data interface{}) error {
	if _, ok := data.(Data); ok {
		data := data.(Data)
		fieldsArr := []interface{}{data.Fields}
		params := map[string]interface{}{
			"collection_name": collection.CollectionName,
			"fields":          fieldsArr,
		}
		if data.TTL != 0 {
			params["ttl"] = data.TTL
		}
		res, err := collection.VikingDBService.DoRequest(context.Background(), "AsyncUpsertData", nil, collection.VikingDBService.convertMapToJson(params))
		_ = res
		return err
	} else if _, ok := data.([]Data); ok {
		datas := data.([]Data)
		record := map[int64][]interface{}{}
		for _, item := range datas {
			if value, exist := record[item.TTL]; exist {
				fields := value
				fields = append(fields, item.Fields)
				record[item.TTL] = fields
			} else {
				record[item.TTL] = []interface{}{item.Fields}
			}
		}
		for index, item := range record {
			params := map[string]interface{}{
				"collection_name": collection.CollectionName,
				"fields":          item,
			}
			if index != 0 {
				params["ttl"] = index
			}
			res, err := collection.VikingDBService.DoRequest(context.Background(), "AsyncUpsertData", nil, collection.VikingDBService.convertMapToJson(params))
			_ = res
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return errors.New("invalid data")
	}

}

func (collection *Collection) DeleteData(id interface{}) error {
	params := map[string]interface{}{
		"collection_name": collection.CollectionName,
		"primary_keys":    id,
	}
	_, err := collection.VikingDBService.DoRequest(context.Background(), "DeleteData", nil, collection.VikingDBService.convertMapToJson(params))
	return err
}

func (collection *Collection) DeleteAllData() error {
	params := map[string]interface{}{
		"collection_name": collection.CollectionName,
		"del_all":         true,
	}
	_, err := collection.VikingDBService.DoRequest(context.Background(), "DeleteData", nil, collection.VikingDBService.convertMapToJson(params))
	return err
}

func (collection *Collection) FetchData(id interface{}) ([]*Data, error) {
	_, intType := id.(int)
	_, stringType := id.(string)
	_, ListStringType := id.([]string)
	_, ListIntType := id.([]int)
	datas := []*Data{}
	if intType || stringType {
		params := map[string]interface{}{
			"collection_name": collection.CollectionName,
			"primary_keys":    id,
		}
		res, err := collection.VikingDBService.retryRequest(context.Background(), "FetchData", nil, collection.VikingDBService.convertMapToJson(params), MAX_RETRIES)
		if err != nil {
			return nil, err
		}

		var resData []interface{}
		var fields map[string]interface{}
		if d, ok := res["data"]; !ok {
			return nil, fmt.Errorf("invalid response, data does not exist: %v", res)
		} else if resData, ok = d.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not a list: %v", res)
		} else if fields, ok = resData[0].(map[string]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not list[map]: %v", res)
		}

		data := &Data{
			Id:        id,
			Fields:    fields,
			Timestamp: nil,
		}
		datas = append(datas, data)
	} else if ListIntType || ListStringType {
		params := map[string]interface{}{
			"collection_name": collection.CollectionName,
			"primary_keys":    id,
		}
		res, err := collection.VikingDBService.retryRequest(context.Background(), "FetchData", nil, collection.VikingDBService.convertMapToJson(params), MAX_RETRIES)
		if err != nil {
			return nil, err
		}
		var resData []interface{}
		if d, ok := res["data"]; !ok {
			return nil, fmt.Errorf("invalid response, data does not exist: %v", res)
		} else if resData, ok = d.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not a list: %v", res)
		}
		for _, item := range resData {
			var itemMap map[string]interface{}
			var ok bool
			if itemMap, ok = item.(map[string]interface{}); !ok {
				return nil, fmt.Errorf("invalid response, data is not list[map]: %v", res)
			}
			data := &Data{
				Id:        itemMap[collection.PrimaryKey],
				Fields:    itemMap,
				Timestamp: nil,
			}
			datas = append(datas, data)
		}
	} else {
		return nil, errors.New("invalid id")
	}
	return datas, nil
}

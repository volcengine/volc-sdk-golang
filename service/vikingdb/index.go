package vikingdb

import (
	"context"
	"encoding/json"
	"fmt"
)

type Index struct {
	CollectionName  string
	IndexName       string
	Description     string
	VectorIndex     *VectorIndexParams
	ScalarIndex     interface{}
	Stat            string
	VikingDBService *VikingDBService
	CpuQuota        int64
	PartitionBy     string
	CreateTime      string
	UpdateTime      string
	UpdatePerson    string
	IndexCost       interface{}
	ShardCount      int64
	ShardPolicy     string
	primaryKey      string
}

type SearchOptions struct {
	filter                map[string]interface{}
	limit                 int64
	outputFields          []string
	partition             string
	denseWeight           *float64
	sparseVectors         []map[string]interface{}
	needInstruction       *bool
	primaryKeyIn          []interface{}
	primaryKeyNotIn       []interface{}
	postProcessOps        []map[string]interface{}
	postProcessInputLimit *int64
	retry                 bool
}

func NewSearchOptions() *SearchOptions {
	searchOptions := &SearchOptions{
		filter:                nil,
		limit:                 10,
		outputFields:          nil,
		partition:             "default",
		denseWeight:           nil,
		sparseVectors:         nil,
		needInstruction:       nil,
		primaryKeyIn:          nil,
		primaryKeyNotIn:       nil,
		postProcessOps:        nil,
		postProcessInputLimit: nil,
		retry:                 false,
	}
	return searchOptions
}
func (s *SearchOptions) SetFilter(filter map[string]interface{}) *SearchOptions {
	s.filter = filter
	return s
}
func (s *SearchOptions) SetLimit(limit int64) *SearchOptions {
	s.limit = limit
	return s
}
func (s *SearchOptions) SetOutputFields(outputFields []string) *SearchOptions {
	s.outputFields = outputFields
	return s
}
func (s *SearchOptions) SetPartition(partition string) *SearchOptions {
	s.partition = partition
	return s
}
func (s *SearchOptions) SetDenseWeight(denseWeight float64) *SearchOptions {
	s.denseWeight = &denseWeight
	return s
}
func (s *SearchOptions) SetSparseVectors(sparseVectors map[string]interface{}) *SearchOptions {
	s.sparseVectors = []map[string]interface{}{sparseVectors}
	return s
}
func (s *SearchOptions) SetNeedInstruction(needInstruction bool) *SearchOptions {
	s.needInstruction = &needInstruction
	return s
}
func (s *SearchOptions) SetPrimaryKeyIn(primaryKeyIn []interface{}) *SearchOptions {
	s.primaryKeyIn = primaryKeyIn
	return s
}
func (s *SearchOptions) SetPrimaryKeyNotIn(primaryKeyNotIn []interface{}) *SearchOptions {
	s.primaryKeyNotIn = primaryKeyNotIn
	return s
}
func (s *SearchOptions) SetPostProcessOps(postProcessOps []map[string]interface{}) *SearchOptions {
	s.postProcessOps = postProcessOps
	return s
}
func (s *SearchOptions) SetPostProcessInputLimit(postProcessInputLimit int64) *SearchOptions {
	s.postProcessInputLimit = &postProcessInputLimit
	return s
}
func (s *SearchOptions) SetRetry(retry bool) *SearchOptions {
	s.retry = retry
	return s
}

type SearchAggOptions struct {
	filter    map[string]interface{}
	partition string
	agg       map[string]interface{}
	retry     bool
}

func NewSearchAggOptions() *SearchAggOptions {
	searchAggOptions := &SearchAggOptions{
		filter:    nil,
		partition: "default",
		agg:       nil,
		retry:     false,
	}
	return searchAggOptions
}
func (a *SearchAggOptions) SetFilter(filter map[string]interface{}) *SearchAggOptions {
	a.filter = filter
	return a
}
func (a *SearchAggOptions) SetPartition(partition string) *SearchAggOptions {
	a.partition = partition
	return a
}
func (a *SearchAggOptions) SetAgg(agg map[string]interface{}) *SearchAggOptions {
	a.agg = agg
	return a
}
func (a *SearchAggOptions) SetRetry(retry bool) *SearchAggOptions {
	a.retry = retry
	return a
}

type SearchAggResult struct {
	AggOp        string
	GroupByField string
	AggResult    map[string]interface{}
}

type SortOptions struct {
	queryVector []float64
	primaryKeys []interface{}
	retry       bool
}

func NewSortOptions() *SortOptions {
	sortOptions := &SortOptions{
		queryVector: nil,
		primaryKeys: nil,
		retry:       false,
	}
	return sortOptions
}
func (t *SortOptions) SetQueryVector(queryVector []float64) *SortOptions {
	t.queryVector = queryVector
	return t
}
func (t *SortOptions) SetPrimaryKeys(primaryKeys []interface{}) *SortOptions {
	t.primaryKeys = primaryKeys
	return t
}
func (t *SortOptions) SetRetry(retry bool) *SortOptions {
	t.retry = retry
	return t
}

type IndexSortResult struct {
	SortResult         []*SortResultItem
	PrimaryKeyNotExist []interface{}
}
type SortResultItem struct {
	PrimaryKey interface{}
	Score      float64
}

func (index *Index) Search(order interface{}, searchOptions *SearchOptions) ([]*Data, error) {
	if _, ok := order.(VectorOrder); ok {
		order := order.(VectorOrder)
		if order.Vector != nil {
			return index.SearchByVector(order.Vector.([]float64), searchOptions)

		} else if order.Id != nil {
			return index.SearchById(order.Id, searchOptions)
		}
	} else if _, ok := order.(ScalarOrder); ok {
		order := order.(ScalarOrder)
		orderByScalar := map[string]interface{}{
			"order":      order.Order,
			"field_name": order.FieldName,
		}
		search := map[string]interface{}{
			"order_by_scalar": orderByScalar,
			"limit":           searchOptions.limit,
			"partition":       searchOptions.partition,
		}
		if searchOptions.filter != nil {
			search["filter"] = searchOptions.filter
		}
		var outputFields interface{} = nil
		if searchOptions.outputFields != nil {
			search["output_fields"] = searchOptions.outputFields
			outputFields = searchOptions.outputFields
		}
		if searchOptions.primaryKeyIn != nil {
			search["primary_key_in"] = searchOptions.primaryKeyIn
		}
		if searchOptions.primaryKeyNotIn != nil {
			search["primary_key_not_in"] = searchOptions.primaryKeyNotIn
		}
		if searchOptions.postProcessOps != nil {
			search["post_process_ops"] = searchOptions.postProcessOps
		}
		if searchOptions.postProcessInputLimit != nil {
			search["post_process_input_limit"] = searchOptions.postProcessInputLimit
		}
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"search":          search,
		}
		remainingRetries := 0
		if searchOptions.retry {
			remainingRetries = MAX_RETRIES
		}
		res, err := index.VikingDBService.retryRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params), remainingRetries)
		if err != nil {
			return nil, err
		}
		return index.getData(res, outputFields)
	} else if order == nil {
		search := map[string]interface{}{
			"limit":     searchOptions.limit,
			"partition": searchOptions.partition,
		}
		if searchOptions.filter != nil {
			search["filter"] = searchOptions.filter
		}
		var outputFields interface{} = nil
		if searchOptions.outputFields != nil {
			search["output_fields"] = searchOptions.outputFields
			outputFields = searchOptions.outputFields
		}
		if searchOptions.primaryKeyIn != nil {
			search["primary_key_in"] = searchOptions.primaryKeyIn
		}
		if searchOptions.primaryKeyNotIn != nil {
			search["primary_key_not_in"] = searchOptions.primaryKeyNotIn
		}
		if searchOptions.postProcessOps != nil {
			search["post_process_ops"] = searchOptions.postProcessOps
		}
		if searchOptions.postProcessInputLimit != nil {
			search["post_process_input_limit"] = searchOptions.postProcessInputLimit
		}
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"search":          search,
		}
		remainingRetries := 0
		if searchOptions.retry {
			remainingRetries = MAX_RETRIES
		}
		res, err := index.VikingDBService.retryRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params), remainingRetries)
		if err != nil {
			return nil, err
		}
		return index.getData(res, outputFields)
	}
	return nil, nil
}

func (index *Index) SearchById(id interface{}, searchOptions *SearchOptions) ([]*Data, error) {
	orderById := map[string]interface{}{"primary_keys": id}
	search := map[string]interface{}{
		"order_by_vector": orderById,
		"limit":           searchOptions.limit,
		"partition":       searchOptions.partition,
	}
	if searchOptions.filter != nil {
		search["filter"] = searchOptions.filter
	}
	if searchOptions.denseWeight != nil {
		search["dense_weight"] = *searchOptions.denseWeight
	}
	var outputFields interface{} = nil
	if searchOptions.outputFields != nil {
		search["output_fields"] = searchOptions.outputFields
		outputFields = searchOptions.outputFields
	}
	if searchOptions.primaryKeyIn != nil {
		search["primary_key_in"] = searchOptions.primaryKeyIn
	}
	if searchOptions.primaryKeyNotIn != nil {
		search["primary_key_not_in"] = searchOptions.primaryKeyNotIn
	}
	if searchOptions.postProcessOps != nil {
		search["post_process_ops"] = searchOptions.postProcessOps
	}
	if searchOptions.postProcessInputLimit != nil {
		search["post_process_input_limit"] = searchOptions.postProcessInputLimit
	}
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"search":          search,
	}
	remainingRetries := 0
	if searchOptions.retry {
		remainingRetries = MAX_RETRIES
	}
	res, err := index.VikingDBService.retryRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params), remainingRetries)
	if err != nil {
		return nil, err
	}
	return index.getData(res, outputFields)
}

func (index *Index) SearchByVector(vector []float64, searchOptions *SearchOptions) ([]*Data, error) {
	orderByVector := map[string]interface{}{"vectors": []interface{}{vector}}
	if searchOptions.sparseVectors != nil {
		orderByVector["sparse_vectors"] = searchOptions.sparseVectors
	}
	search := map[string]interface{}{
		"order_by_vector": orderByVector,
		"limit":           searchOptions.limit,
		"partition":       searchOptions.partition,
	}
	if searchOptions.filter != nil {
		search["filter"] = searchOptions.filter
	}
	if searchOptions.denseWeight != nil {
		search["dense_weight"] = *searchOptions.denseWeight
	}
	var outputFields interface{} = nil
	if searchOptions.outputFields != nil {
		search["output_fields"] = searchOptions.outputFields
		outputFields = searchOptions.outputFields
	}
	if searchOptions.primaryKeyIn != nil {
		search["primary_key_in"] = searchOptions.primaryKeyIn
	}
	if searchOptions.primaryKeyNotIn != nil {
		search["primary_key_not_in"] = searchOptions.primaryKeyNotIn
	}
	if searchOptions.postProcessOps != nil {
		search["post_process_ops"] = searchOptions.postProcessOps
	}
	if searchOptions.postProcessInputLimit != nil {
		search["post_process_input_limit"] = searchOptions.postProcessInputLimit
	}
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"search":          search,
	}
	remainingRetries := 0
	if searchOptions.retry {
		remainingRetries = MAX_RETRIES
	}
	res, err := index.VikingDBService.retryRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params), remainingRetries)
	if err != nil {
		return nil, err
	}
	return index.getData(res, outputFields)
}

func (index *Index) SearchByText(text TextObject, searchOptions *SearchOptions) ([]*Data, error) {
	orderByRaw := map[string]interface{}{"text": text.Text}
	search := map[string]interface{}{
		"order_by_raw": orderByRaw,
		"limit":        searchOptions.limit,
		"partition":    searchOptions.partition,
	}
	if searchOptions.filter != nil {
		search["filter"] = searchOptions.filter
	}
	if searchOptions.denseWeight != nil {
		search["dense_weight"] = *searchOptions.denseWeight
	}
	var outputFields interface{} = nil
	if searchOptions.outputFields != nil {
		search["output_fields"] = searchOptions.outputFields
		outputFields = searchOptions.outputFields
	}
	if searchOptions.needInstruction != nil {
		search["need_instruction"] = *searchOptions.needInstruction
	}
	if searchOptions.primaryKeyIn != nil {
		search["primary_key_in"] = searchOptions.primaryKeyIn
	}
	if searchOptions.primaryKeyNotIn != nil {
		search["primary_key_not_in"] = searchOptions.primaryKeyNotIn
	}
	if searchOptions.postProcessOps != nil {
		search["post_process_ops"] = searchOptions.postProcessOps
	}
	if searchOptions.postProcessInputLimit != nil {
		search["post_process_input_limit"] = searchOptions.postProcessInputLimit
	}
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"search":          search,
	}
	remainingRetries := 0
	if searchOptions.retry {
		remainingRetries = MAX_RETRIES
	}
	res, err := index.VikingDBService.retryRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params), remainingRetries)
	if err != nil {
		return nil, err
	}
	return index.getData(res, outputFields)
}

func (index *Index) SearchAgg(searchAggOptions *SearchAggOptions) (*SearchAggResult, error) {
	search := map[string]interface{}{
		"partition": searchAggOptions.partition,
	}
	if searchAggOptions.filter != nil {
		search["filter"] = searchAggOptions.filter
	}
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"search":          search,
		"agg":             searchAggOptions.agg,
	}
	remainingRetries := 0
	if searchAggOptions.retry {
		remainingRetries = MAX_RETRIES
	}
	res, err := index.VikingDBService.retryRequest(context.Background(), "SearchAgg", nil, index.VikingDBService.convertMapToJson(params), remainingRetries)
	if err != nil {
		return nil, err
	}
	return index.getAggResult(res)
}

func (index *Index) Sort(sortOptions *SortOptions) (*IndexSortResult, error) {
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"sort": map[string]interface{}{
			"query_vector": sortOptions.queryVector,
			"primary_keys": sortOptions.primaryKeys,
		},
	}
	remainingRetries := 0
	if sortOptions.retry {
		remainingRetries = MAX_RETRIES
	}
	res, err := index.VikingDBService.retryRequest(context.Background(), "IndexSort", nil, index.VikingDBService.convertMapToJson(params), remainingRetries)
	if err != nil {
		return nil, err
	}
	return index.getIndexSortResult(res)
}

func (index *Index) FetchData(id interface{}, searchOptions *SearchOptions) ([]*Data, error) {
	_, intType := id.(int)
	_, stringType := id.(string)
	datas := []*Data{}
	if intType || stringType {
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"primary_keys":    id,
			"partition":       searchOptions.partition,
		}
		if searchOptions.outputFields != nil {
			params["output_fields"] = searchOptions.outputFields
		}
		res, err := index.VikingDBService.retryRequest(context.Background(), "FetchIndexData", nil, index.VikingDBService.convertMapToJson(params), MAX_RETRIES)
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
			Id: id,
			Fields: map[string]interface{}{
				"message": "no data found",
			},
			Timestamp: nil,
		}
		if _, exists := fields["fields"]; exists {
			var dataFields map[string]interface{}
			var ok bool
			if dataFields, ok = fields["fields"].(map[string]interface{}); !ok {
				return nil, fmt.Errorf("invalid response, fields is not a map: %v", res)
			}
			data = &Data{
				Id:        id,
				Fields:    dataFields,
				Timestamp: nil,
			}
		}
		datas = append(datas, data)
	} else {
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"primary_keys":    id,
			"partition":       searchOptions.partition,
		}
		if searchOptions.outputFields != nil {
			params["output_fields"] = searchOptions.outputFields
		}
		res, err := index.VikingDBService.retryRequest(context.Background(), "FetchIndexData", nil, index.VikingDBService.convertMapToJson(params), MAX_RETRIES)
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
			primaryKey, err := index.getPrimaryKey()
			if err != nil {
				return nil, err
			}
			data := &Data{
				Id: itemMap[primaryKey],
				Fields: map[string]interface{}{
					"message": "no data found",
				},
				Timestamp: nil,
			}
			if _, exists := itemMap["fields"]; exists {
				var dataFields map[string]interface{}
				var ok bool
				if dataFields, ok = itemMap["fields"].(map[string]interface{}); !ok {
					return nil, fmt.Errorf("invalid response, fields is not a map: %v", res)
				}

				primaryKey, err := index.getPrimaryKey()
				if err != nil {
					return nil, err
				}
				data = &Data{
					Id:        itemMap[primaryKey],
					Fields:    dataFields,
					Timestamp: nil,
				}
			}

			datas = append(datas, data)
		}
	}
	return datas, nil
}

func (index *Index) getData(resData map[string]interface{}, outputField interface{}) ([]*Data, error) {
	var res []interface{}
	if d, ok := resData["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", resData)
	} else if res, ok = d.([]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", resData)
	}

	datas := []*Data{}
	for _, items := range res {
		var itemsList []interface{}
		var ok bool
		if itemsList, ok = items.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not list[list]: %v", resData)
		}
		for _, l := range itemsList {
			var item map[string]interface{}
			var ok bool
			if item, ok = l.(map[string]interface{}); !ok {
				return nil, fmt.Errorf("invalid response, data is not list[list[map]]: %v", resData)
			}
			fields := map[string]interface{}{}
			if outputField == nil || len(outputField.([]string)) != 0 {
				if f, ok := item["fields"]; !ok {
					return nil, fmt.Errorf("invalid response, fields does not exist: %v", resData)
				} else if fields, ok = f.(map[string]interface{}); !ok {
					return nil, fmt.Errorf("invalid response, fields is not a map: %v", resData)
				}
			}
			text := ""
			if value, exists := item["text"]; exists {
				if t, ok := value.(string); !ok {
					return nil, fmt.Errorf("invalid response, text is not string: %v", resData)
				} else {
					text = t
				}
			}
			primaryKey, err := index.getPrimaryKey()
			if err != nil {
				return nil, err
			}
			var scoreFloat float64
			if s, ok := item["score"]; !ok {
				return nil, fmt.Errorf("invalid response, score not exists: %v", resData)
			} else if val, err := ParseJsonFloat64Field(s); err != nil {
				return nil, fmt.Errorf("invalid response, score is %v, res is %v", s, res)
			} else {
				scoreFloat = val
			}
			data := &Data{
				Fields:    fields,
				Id:        item[primaryKey],
				Score:     scoreFloat,
				Timestamp: nil,
				Text:      text,
			}
			datas = append(datas, data)
		}
	}
	return datas, nil
}

func (index *Index) getPrimaryKey() (string, error) {
	if index.primaryKey == "" {
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
		}
		resData, err := index.VikingDBService.DoRequest(context.Background(), "GetCollection", nil, index.VikingDBService.convertMapToJson(params))
		if err != nil {
			return "", err
		}
		var resDataItem map[string]interface{}
		if d, ok := resData["data"]; !ok {
			return "", fmt.Errorf("invalid response, data does not exist: %v", resData)
		} else if resDataItem, ok = d.(map[string]interface{}); !ok {
			return "", fmt.Errorf("invalid response, data is not a map: %v", resData)
		}
		if value, exist := resDataItem["primary_key"]; exist {
			if v, ok := value.(string); !ok {
				return "", fmt.Errorf("invalid response, primary_key is invalid:%v", resDataItem)
			} else {
				index.primaryKey = v
			}
		}
		return index.primaryKey, err
	} else {
		return index.primaryKey, nil
	}
}

func (index *Index) getAggResult(response map[string]interface{}) (*SearchAggResult, error) {
	var data map[string]interface{}
	if d, ok := response["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist")
	} else if data, ok = d.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", d)
	}
	result := &SearchAggResult{}
	if raw, ok1 := data["agg_op"]; !ok1 {
		return nil, fmt.Errorf("invalid response, agg_op not found")
	} else if val, ok2 := raw.(string); !ok2 {
		return nil, fmt.Errorf("invalid response, agg_op is not string")
	} else {
		result.AggOp = val
	}
	if raw, ok1 := data["group_by_field"]; !ok1 {
		return nil, fmt.Errorf("invalid response, group_by_field not found")
	} else if val, ok2 := raw.(string); !ok2 {
		return nil, fmt.Errorf("invalid response, group_by_field is not string")
	} else {
		result.GroupByField = val
	}
	if raw, ok1 := data["agg_result"]; !ok1 {
		return nil, fmt.Errorf("invalid response, agg_result not found")
	} else if val, ok2 := raw.(map[string]interface{}); !ok2 {
		return nil, fmt.Errorf("invalid response, agg_result is not map[string]interface{}")
	} else {
		result.AggResult = val
	}
	return result, nil
}

func (index *Index) getIndexSortResult(response map[string]interface{}) (*IndexSortResult, error) {
	var data map[string]interface{}
	if d, ok := response["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist")
	} else if data, ok = d.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", d)
	}
	result := &IndexSortResult{}
	if raw, ok1 := data["sort_result"]; !ok1 {
		return nil, fmt.Errorf("invalid response, sort_result not found")
	} else if vals, ok2 := raw.([]interface{}); !ok2 {
		return nil, fmt.Errorf("invalid response, sort_result is not []interface{}")
	} else {
		sortResult := make([]*SortResultItem, 0)
		for _, val := range vals {
			if valMap, ok3 := val.(map[string]interface{}); !ok3 {
				return nil, fmt.Errorf("invalid response, sort_result item is not map[string]interface{}")
			} else {
				item := &SortResultItem{}
				if valPk, ok4 := valMap["primary_key"]; !ok4 {
					return nil, fmt.Errorf("invalid response, primary_key not found in sort_result item")
				} else {
					item.PrimaryKey = valPk
				}
				if valScore, ok4 := valMap["score"]; !ok4 {
					return nil, fmt.Errorf("invalid response, score not found in sort_result item")
				} else if valNum, ok5 := valScore.(json.Number); !ok5 {
					return nil, fmt.Errorf("invalid response, score is not a number in sort_result item")
				} else if valFloat, err := valNum.Float64(); err != nil {
					return nil, fmt.Errorf("invalid response, score is not a float in sort_result item, err: %v", err)
				} else {
					item.Score = valFloat
				}
				sortResult = append(sortResult, item)
			}
		}
		result.SortResult = sortResult
	}
	if raw, ok1 := data["primary_key_not_exist"]; !ok1 {
		return nil, fmt.Errorf("invalid response, sort_result not found")
	} else if val, ok2 := raw.([]interface{}); !ok2 {
		return nil, fmt.Errorf("invalid response, sort_result is not []interface{}")
	} else {
		result.PrimaryKeyNotExist = val
	}
	return result, nil
}

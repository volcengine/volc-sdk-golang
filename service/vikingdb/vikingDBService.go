package vikingdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

type VikingDBService struct {
	Client *base.Client
}

func NewVikingDBService(host string, region string, ak string, sk string, scheme string) *VikingDBService {
	client := base.NewClient(getServiceInfo(host, region, ak, sk, scheme), getApiInfo())
	vikingDBService := &VikingDBService{
		Client: client,
	}
	_, err := vikingDBService.DoRequest(context.Background(), "Ping", nil, "")
	if err != nil {
		panic(fmt.Errorf("host or region is incorrect: %v", err))
	}
	return vikingDBService

}
func (vikingDBService *VikingDBService) SetConnectionTimeout(connectionTimeout int64) {
	vikingDBService.Client.ServiceInfo.Timeout = time.Duration(connectionTimeout) * time.Second
}
func getApiInfo() map[string]*base.ApiInfo {
	apiInfos := map[string]*base.ApiInfo{
		"CreateCollection": {
			Method: http.MethodPost,
			Path:   "/api/collection/create",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"GetCollection": {
			Method: http.MethodGet,
			Path:   "/api/collection/info",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"DropCollection": {
			Method: http.MethodPost,
			Path:   "/api/collection/drop",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"ListCollections": {
			Method: http.MethodGet,
			Path:   "/api/collection/list",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"CreateIndex": {
			Method: http.MethodPost,
			Path:   "/api/index/create",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"GetIndex": {
			Method: http.MethodGet,
			Path:   "/api/index/info",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"DropIndex": {
			Method: http.MethodPost,
			Path:   "/api/index/drop",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"ListIndexes": {
			Method: http.MethodGet,
			Path:   "/api/index/list",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"UpsertData": {
			Method: http.MethodPost,
			Path:   "/api/collection/upsert_data",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"AsyncUpsertData": {
			Method: http.MethodPost,
			Path:   "/api/collection/async_upsert_data",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"FetchData": {
			Method: http.MethodGet,
			Path:   "/api/collection/fetch_data",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"DeleteData": {
			Method: http.MethodPost,
			Path:   "/api/collection/del_data",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"SearchIndex": {
			Method: http.MethodPost,
			Path:   "/api/index/search",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"FetchIndexData": {
			Method: http.MethodGet,
			Path:   "/api/index/fetch_data",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"Embedding": {
			Method: http.MethodPost,
			Path:   "/api/data/embedding",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"UpdateCollection": {
			Method: http.MethodPost,
			Path:   "/api/collection/update",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"UpdateIndex": {
			Method: http.MethodPost,
			Path:   "/api/index/update",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"Rerank": {
			Method: http.MethodPost,
			Path:   "/api/index/rerank",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"BatchRerank": {
			Method: http.MethodPost,
			Path:   "/api/index/batch_rerank",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"Ping": {
			Method: http.MethodGet,
			Path:   "/api/viking_db/data/ping",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
		"EmbeddingV2": {
			Method: http.MethodPost,
			Path:   "/api/data/embedding/version/2",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
	}
	return apiInfos
}
func (vikingDBService *VikingDBService) SetHeader(header map[string]string) {
	apiInfo := vikingDBService.Client.ApiInfoList
	for key, _ := range apiInfo {
		for item, _ := range header {
			apiInfo[key].Header[item] = []string{header[item]}
		}
	}
	vikingDBService.Client.ApiInfoList = apiInfo
}

func (vikingDBService *VikingDBService) DoRequest(ctx context.Context, api string, query url.Values, body string) (map[string]interface{}, error) {
	res, code, err := vikingDBService.Client.CtxJson(ctx, api, query, body)
	if err != nil {
		return nil, err
	}
	_ = code
	var data map[string]interface{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}
	//fmt.Println(data)
	return data, err
}
func getServiceInfo(host string, region string, ak string, sk string, scheme string) *base.ServiceInfo {
	serviceInfo := &base.ServiceInfo{
		Scheme: scheme,
		Host:   host,
		Credentials: base.Credentials{
			AccessKeyID:     ak,
			SecretAccessKey: sk,
			Service:         "air",
			Region:          region,
		},
		Timeout: 5 * time.Second,
		Header: http.Header{
			"Host": []string{host},
		},
	}
	return serviceInfo
}
func (vikingDBService *VikingDBService) convertMapToJson(params map[string]interface{}) string {
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(params)
	return buffer.String()
}
func (vikingDBService *VikingDBService) packageCollection(collectionName string, res map[string]interface{}) (*Collection, error) {
	var description string
	var stat map[string]interface{}
	var fields []Field
	var indexes []*Index
	var createTime string
	var updateTime string
	var updatePerson string
	if value, exist := res["description"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, description is not string: %v", res)
		} else {
			description = v
		}
	}
	if value, exist := res["stat"]; exist {
		if v, ok := value.(map[string]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, stat is not map: %v", res)
		} else {
			stat = v
		}
	}
	if value, exist := res["create_time"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, createTime is not string: %v", res)
		} else {
			createTime = v
		}
	}
	if value, exist := res["update_time"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, updateTime is not string: %v", res)
		} else {
			updateTime = v
		}
	}
	if value, exist := res["update_person"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, updatePerson is not string: %v", res)
		} else {
			updatePerson = v
		}
	}
	if value, exist := res["indexes"]; exist {
		var indexesName []interface{}
		var ok bool
		if indexesName, ok = value.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, indexes is not list: %v", res)
		}
		for _, item := range indexesName {
			var itemString string
			if itemString, ok = item.(string); !ok {
				return nil, fmt.Errorf("invalid response, indexeName is not string: %v", res)
			}
			index, err := vikingDBService.GetIndex(collectionName, itemString)
			if err != nil {
				return nil, err
			}
			indexes = append(indexes, index)
		}
	}
	if value, exist := res["fields"]; exist {
		fieldList := value.([]interface{})
		for _, item := range fieldList {
			var fieldName string
			var fieldType string
			var defaultVal interface{}
			var dim int64
			var isPrimaryKey bool
			var pipelineName string
			itemMap := item.(map[string]interface{})
			if value, exist := itemMap["field_name"]; exist {
				if v, ok := value.(string); !ok {
					return nil, fmt.Errorf("invalid response, field_name is not string: %v", res)
				} else {
					fieldName = v
				}
			}
			if value, exist := itemMap["field_type"]; exist {
				if v, ok := value.(string); !ok {
					return nil, fmt.Errorf("invalid response, field_type is not string: %v", res)
				} else {
					fieldType = v
				}
			}
			if value, exist := itemMap["default_val"]; exist {
				defaultVal = value
			}
			if value, exist := itemMap["dim"]; exist {
				if v, ok := value.(float64); !ok {
					return nil, fmt.Errorf("invalid response, dim is not a number: %v", res)
				} else {
					dim = int64(v)
				}
			}
			if value, exist := itemMap["pipeline_name"]; exist {
				if v, ok := value.(string); !ok {
					return nil, fmt.Errorf("invalid response, pipeline_name is not string: %v", res)
				} else {
					pipelineName = v
				}
			}
			if value, exist := res["primary_key"]; exist {
				if v, ok := value.(string); !ok {
					return nil, fmt.Errorf("invalid response, primary_key is invalid:%v", res)
				} else if v == fieldName {
					isPrimaryKey = true
				}
			}
			field := Field{
				FieldName:    fieldName,
				FieldType:    fieldType,
				DefaultVal:   defaultVal,
				Dim:          dim,
				IsPrimaryKey: isPrimaryKey,
				PipelineName: pipelineName,
			}
			fields = append(fields, field)
		}

	}
	collection := &Collection{
		CollectionName:  collectionName,
		Fields:          fields,
		VikingDBService: vikingDBService,
		PrimaryKey:      res["primary_key"].(string),
		Indexes:         indexes,
		Description:     description,
		Stat:            stat,
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		UpdatePerson:    updatePerson,
	}
	return collection, nil
}
func (vikingDBService *VikingDBService) packageIndex(collectionName string, indexName string, res map[string]interface{}) (*Index, error) {
	var description string
	var stat string
	var createTime string
	var updateTime string
	var updatePerson string
	var partitionBy string
	var cpuQuota int64
	var scalarIndex interface{}
	var vectorIndex *VectorIndexParams
	var indexCost interface{}
	var shardCount int64
	var shardPolicy string
	if value, exist := res["description"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, description is not string: %v", res)
		} else {
			description = v
		}
	}
	if value, exist := res["status"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, status is not string: %v", res)
		} else {
			stat = v
		}
	}
	if value, exist := res["create_time"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, create_time is not string: %v", res)
		} else {
			createTime = v
		}
	}
	if value, exist := res["update_time"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, update_time is not string: %v", res)
		} else {
			updateTime = v
		}
	}
	if value, exist := res["update_person"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, update_person is not string: %v", res)
		} else {
			updatePerson = v
		}
	}
	if value, exist := res["cpu_quota"]; exist {
		if v, ok := value.(float64); !ok {
			return nil, fmt.Errorf("invalid response, cpu_quota is not a number: %v", res)
		} else {
			cpuQuota = int64(v)
		}
	}
	if value, exist := res["partition_by"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, partition_by is not string: %v", res)
		} else {
			partitionBy = v
		}
	}
	if value, exist := res["index_cost"]; exist {
		if v, ok := value.(map[string]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, index_cost is not a map: %v", res)
		} else {
			indexCost = v
		}
	}
	if value, exist := res["shard_count"]; exist {
		if v, ok := value.(float64); !ok {
			return nil, fmt.Errorf("invalid response, shard_count is not a number: %v", res)
		} else {
			shardCount = int64(v)
		}
	}
	if value, exist := res["shard_policy"]; exist {
		if v, ok := value.(string); !ok {
			return nil, fmt.Errorf("invalid response, shard_policy is not string: %v", res)
		} else {
			shardPolicy = v
		}
	}
	if value, exist := res["vector_index"]; exist {
		var item map[string]interface{}
		var ok bool
		if item, ok = value.(map[string]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, vector_index is not a map: %v", res)
		}
		vectorIndex = &VectorIndexParams{
			IndexType: "hnsw",
			Distance:  "ip",
			Quant:     "int8",
			HnswM:     20,
			HnswSef:   800,
			HnswCef:   400,
		}
		if v, e := item["index_type"]; e {
			if vs, ok := v.(string); !ok {
				return nil, fmt.Errorf("invalid response, index_type is not string: %v", res)
			} else {
				vectorIndex.IndexType = vs
			}
		}
		if v, e := item["distance"]; e {
			if vs, ok := v.(string); !ok {
				return nil, fmt.Errorf("invalid response, distance is not string: %v", res)
			} else {
				vectorIndex.Distance = vs
			}
		}
		if v, e := item["quant"]; e {
			if vs, ok := v.(string); !ok {
				return nil, fmt.Errorf("invalid response, quant is not string: %v", res)
			} else {
				vectorIndex.Quant = vs
			}
		}
		if v, e := item["hnsw_m"]; e {
			if vf, ok := v.(float64); !ok {
				return nil, fmt.Errorf("invalid response, hnsw_m is not a number: %v", res)
			} else {
				vectorIndex.HnswM = int64(vf)
			}
		}
		if v, e := item["hnsw_sef"]; e {
			if vf, ok := v.(float64); !ok {
				return nil, fmt.Errorf("invalid response, hnsw_sef is not a number: %v", res)
			} else {
				vectorIndex.HnswSef = int64(vf)
			}
		}
		if v, e := item["hnsw_cef"]; e {
			if vf, ok := v.(float64); !ok {
				return nil, fmt.Errorf("invalid response, hnsw_cef is not a number: %v", res)
			} else {
				vectorIndex.HnswCef = int64(vf)
			}
		}
	}
	if value, exist := res["range_index"]; exist {
		if v, ok := value.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, range_index is not a map: %v", res)
		} else {
			scalarIndex = v
		}
	}
	if value, exist := res["enum_index"]; exist {
		var ok bool
		var items []interface{}
		if items, ok = value.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, enum_index is not a map: %v", res)
		}
		if scalarIndex == nil {
			scalarIndex = items
		} else {
			for _, itemEnum := range items {
				flag := 1
				for _, itemScalar := range scalarIndex.([]interface{}) {
					var enumString string
					var scalarString string
					if enumString, ok = itemEnum.(string); !ok {
						return nil, fmt.Errorf("invalid response, enum_index is not string: %v", res)
					}
					if scalarString, ok = itemScalar.(string); !ok {
						return nil, fmt.Errorf("invalid response, range_index is not string: %v", res)
					}
					if enumString == scalarString {
						flag = 0
						break
					}
				}
				if flag == 1 {
					scalarIndex = append(scalarIndex.([]interface{}), itemEnum)
				}

			}
		}
	}
	index := &Index{
		CollectionName:  collectionName,
		IndexName:       indexName,
		VectorIndex:     vectorIndex,
		ScalarIndex:     scalarIndex,
		Stat:            stat,
		VikingDBService: vikingDBService,
		Description:     description,
		CpuQuota:        cpuQuota,
		PartitionBy:     partitionBy,
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		UpdatePerson:    updatePerson,
		IndexCost:       indexCost,
		ShardCount:      shardCount,
		ShardPolicy:     shardPolicy,
	}
	if index.primaryKey == "" {
		_, err := index.getPrimaryKey()
		if err != nil {
			return nil, err
		}
	}
	return index, nil
}

func (vikingDBService *VikingDBService) CreateCollection(collectionName string, fields []Field, description string) (*Collection, error) {
	params := map[string]interface{}{
		"collection_name": collectionName,
		"description":     description,
	}
	var primaryKey string = ""
	var _fields []interface{}
	for _, field := range fields {
		if field.IsPrimaryKey {
			primaryKey = field.FieldName
		}
		_field := map[string]interface{}{
			"field_name": field.FieldName,
			"field_type": field.FieldType,
		}
		if field.DefaultVal != nil {
			_field["default_val"] = field.DefaultVal
		}
		if field.Dim != 0 {
			_field["dim"] = field.Dim
		}
		if field.PipelineName != "" {
			_field["pipeline_name"] = field.PipelineName
		}
		_fields = append(_fields, _field)
	}
	if len(primaryKey) == 0 {
		params["primary_key"] = "__AUTO_ID__"
	} else {
		params["primary_key"] = primaryKey
	}
	params["fields"] = _fields

	request, err := vikingDBService.DoRequest(context.Background(), "CreateCollection", nil, vikingDBService.convertMapToJson(params))
	_ = request
	if err != nil {
		return nil, err
	}
	collection := &Collection{
		CollectionName:  collectionName,
		Fields:          fields,
		VikingDBService: vikingDBService,
		PrimaryKey:      primaryKey,
		Description:     description,
	}
	return collection, err
}
func (vikingDBService *VikingDBService) GetCollection(collectionName string) (*Collection, error) {
	params := map[string]interface{}{
		"collection_name": collectionName,
	}
	resData, err := vikingDBService.DoRequest(context.Background(), "GetCollection", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	if d, ok := resData["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", resData)
	} else if res, ok = d.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a map: %v", resData)
	}
	return vikingDBService.packageCollection(collectionName, res)
}
func (vikingDBService *VikingDBService) DropCollection(collectionName string) error {
	params := map[string]interface{}{
		"collection_name": collectionName,
	}
	_, err := vikingDBService.DoRequest(context.Background(), "DropCollection", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return err
	}
	return nil

}
func (vikingDBService *VikingDBService) ListCollections() ([]*Collection, error) {
	resData, err := vikingDBService.DoRequest(context.Background(), "ListCollections", nil, vikingDBService.convertMapToJson(make(map[string]interface{})))
	if err != nil {
		return nil, err
	}
	//fmt.Println(res)
	var collections []*Collection

	var res []interface{}
	if d, ok := resData["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", resData)
	} else if res, ok = d.([]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", resData)
	}

	for _, itemMap := range res {
		var item map[string]interface{}
		var ok bool
		var nameString string
		if item, ok = itemMap.(map[string]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not list[map]: %v", res)
		} else if name, ok := item["collection_name"]; !ok {
			return nil, fmt.Errorf("invalid response, collection_name does not exist: %v", res)
		} else if nameString, ok = name.(string); !ok {
			return nil, fmt.Errorf("invalid response, collection_name is not string: %v", res)
		}
		collection, err := vikingDBService.packageCollection(nameString, item)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)

	}
	return collections, err
}

func (vikingDBService *VikingDBService) CreateIndex(collectionName string, indexName string, indexOptions *IndexOptions) (*Index, error) {
	params := map[string]interface{}{
		"collection_name": collectionName,
		"index_name":      indexName,
		"cpu_quota":       indexOptions.cpuQuota,
		"description":     indexOptions.description,
		"partition_by":    indexOptions.partitionBy,
	}
	if indexOptions.vectorIndex != nil {
		params["vector_index"] = indexOptions.vectorIndex.dict(indexOptions.vectorIndex)
	}
	if indexOptions.scalarIndex != nil {
		params["scalar_index"] = indexOptions.scalarIndex
	}
	if indexOptions.shardCount != nil {
		params["shard_count"] = *indexOptions.shardCount
	}
	if indexOptions.shardPolicy != "" {
		params["shard_policy"] = indexOptions.shardPolicy
	}
	// fmt.Println(vikingDBService.convertMapToJson(params))
	res, err := vikingDBService.DoRequest(context.Background(), "CreateIndex", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	//fmt.Println(res)
	_ = res
	index := &Index{
		CollectionName:  collectionName,
		IndexName:       indexName,
		VectorIndex:     indexOptions.vectorIndex,
		ScalarIndex:     indexOptions.scalarIndex,
		VikingDBService: vikingDBService,
		Description:     indexOptions.description,
		CpuQuota:        indexOptions.cpuQuota,
		PartitionBy:     indexOptions.partitionBy,
		ShardPolicy:     indexOptions.shardPolicy,
		ShardCount:      *indexOptions.shardCount,
	}
	_, err = index.getPrimaryKey()
	if err != nil {
		return nil, err
	}
	return index, err
}
func (vikingDBService *VikingDBService) GetIndex(collectionName string, indexName string) (*Index, error) {
	params := map[string]interface{}{
		"collection_name": collectionName,
		"index_name":      indexName,
	}
	resData, err := vikingDBService.DoRequest(context.Background(), "GetIndex", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	if d, ok := resData["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", resData)
	} else if res, ok = d.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a map: %v", resData)
	}
	return vikingDBService.packageIndex(collectionName, indexName, res)
}
func (vikingDBService *VikingDBService) DropIndex(collectionName string, indexName string) error {
	params := map[string]interface{}{
		"collection_name": collectionName,
		"index_name":      indexName,
	}
	_, err := vikingDBService.DoRequest(context.Background(), "DropIndex", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return err
	}
	return nil
	//fmt.Println(res)
}
func (vikingDBService *VikingDBService) ListIndexes(collectionName string) ([]*Index, error) {
	params := map[string]interface{}{
		"collection_name": collectionName,
	}
	resData, err := vikingDBService.DoRequest(context.Background(), "ListIndexes", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	var res []interface{}
	if d, ok := resData["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", resData)
	} else if res, ok = d.([]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", resData)
	}
	indexes := []*Index{}
	for _, itemMap := range res {
		var item map[string]interface{}
		var ok bool
		var nameString string
		if item, ok = itemMap.(map[string]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not a list[map]: %v", itemMap)
		} else if name, ok := item["index_name"]; !ok {
			return nil, fmt.Errorf("invalid response, collection_name does not exist: %v", res)
		} else if nameString, ok = name.(string); !ok {
			return nil, fmt.Errorf("invalid response, collection_name is not string: %v", res)
		}
		index, err := vikingDBService.packageIndex(collectionName, nameString, item)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, index)
	}
	return indexes, err
}
func (vikingDBService *VikingDBService) UpdateCollection(collectionName string, updateCollectionOptions *UpdateCollectionOptions) error {
	var _fields []interface{}
	for _, field := range updateCollectionOptions.fields {
		_field := map[string]interface{}{
			"field_name": field.FieldName,
			"field_type": field.FieldType,
		}
		if field.DefaultVal != nil {
			_field["default_val"] = field.DefaultVal
		}
		if field.Dim != 0 {
			_field["dim"] = field.Dim
		}
		if field.PipelineName != "" {
			_field["pipeline_name"] = field.PipelineName
		}
		_fields = append(_fields, _field)
	}
	params := map[string]interface{}{
		"collection_name": collectionName,
		"fields":          _fields,
	}
	if updateCollectionOptions.description != nil {
		params["description"] = *updateCollectionOptions.description
	}
	_, err := vikingDBService.DoRequest(context.Background(), "UpdateCollection", nil, vikingDBService.convertMapToJson(params))
	return err
}
func (vikingDBService *VikingDBService) Embedding(embModel EmbModel, rawData interface{}) ([][]float64, error) {
	model := map[string]interface{}{
		"model_name": embModel.ModelName,
		"params":     embModel.Params,
	}
	_, single := rawData.(RawData)
	_, list := rawData.([]RawData)
	var data []interface{}
	if single {
		param := map[string]interface{}{
			"data_type": rawData.(RawData).DataType,
			"text":      rawData.(RawData).Text,
		}
		data = append(data, param)
	} else if list {
		for _, item := range rawData.([]RawData) {
			param := map[string]interface{}{
				"data_type": item.DataType,
				"text":      item.Text,
			}
			data = append(data, param)
		}
	}
	params := map[string]interface{}{
		"model": model,
		"data":  data,
	}
	res, err := vikingDBService.DoRequest(context.Background(), "Embedding", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	//fmt.Println(res)
	var ret [][]float64
	var items []interface{}
	if d, ok := res["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", data)
	} else if items, ok = d.([]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", data)
	}

	for _, itemList := range items {
		var floatList []float64
		var l []interface{}
		var ok bool
		if l, ok = itemList.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, dataType is not list: %v", res["data"])
		}
		for _, item := range l {
			if itemFloat, ok := item.(float64); !ok {
				return nil, fmt.Errorf("invalid response, dataNum is not float64: %v", res["data"])
			} else {
				floatList = append(floatList, itemFloat)
			}
		}
		ret = append(ret, floatList)
	}
	return ret, err
}
func (vikingDBService *VikingDBService) UpdateIndex(collectionName string, indexName string, updateIndexOptions *UpdateIndexOptions) error {
	params := map[string]interface{}{
		"collection_name": collectionName,
		"index_name":      indexName,
	}
	if updateIndexOptions.description != nil {
		params["description"] = *updateIndexOptions.description
	}
	if updateIndexOptions.cpuQuota != nil {
		params["cpu_quota"] = *updateIndexOptions.cpuQuota
	}
	if updateIndexOptions.scalarIndex != nil {
		params["scalar_index"] = updateIndexOptions.scalarIndex
	}
	if updateIndexOptions.shardCount != nil {
		params["shard_count"] = *updateIndexOptions.shardCount
	}
	_, err := vikingDBService.DoRequest(context.Background(), "UpdateIndex", nil, vikingDBService.convertMapToJson(params))
	return err
}
func (vikingDBService *VikingDBService) Rerank(query string, content string, title string) (float64, error) {
	params := map[string]interface{}{
		"query":   query,
		"content": content,
		"title":   title,
	}
	data, err := vikingDBService.DoRequest(context.Background(), "Rerank", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return 0, err
	}
	if d, ok := data["data"]; !ok {
		return 0, fmt.Errorf("invalid response, data does not exist: %v", data)
	} else if score, ok := d.(float64); !ok {
		return 0, fmt.Errorf("invalid response, data is not float64: %v", data)
	} else {
		return score, nil
	}
}
func (vikingDBService *VikingDBService) BatchRerank(datas []map[string]interface{}) ([]float64, error) {
	params := map[string]interface{}{
		"datas": datas,
	}
	data, err := vikingDBService.DoRequest(context.Background(), "BatchRerank", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	var scores []interface{}
	var res []float64
	if d, ok := data["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", data)
	} else if scores, ok = d.([]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", data)
	}
	for _, score := range scores {
		if scoreFloat, ok := score.(float64); !ok {
			return nil, fmt.Errorf("invalid response, data is not list[float64]: %v", data)
		} else {
			res = append(res, scoreFloat)
		}
	}
	return res, nil
}

func (vikingDBService *VikingDBService) EmbeddingV2(embModel EmbModel, rawData interface{}) (map[string]interface{}, error) {
	model := map[string]interface{}{
		"model_name": embModel.ModelName,
		"params":     embModel.Params,
	}
	_, single := rawData.(RawData)
	_, list := rawData.([]RawData)
	var data []interface{}
	if single {
		param := map[string]interface{}{
			"data_type": rawData.(RawData).DataType,
		}
		if rawData.(RawData).Text != "" {
			param["text"] = rawData.(RawData).Text
		}
		if rawData.(RawData).Image != "" {
			param["image"] = rawData.(RawData).Image
		}
		data = append(data, param)
	} else if list {
		for _, item := range rawData.([]RawData) {
			param := map[string]interface{}{
				"data_type": item.DataType,
			}
			if item.Text != "" {
				param["text"] = item.Text
			}
			if item.Image != "" {
				param["image"] = item.Image
			}
			data = append(data, param)
		}
	}
	params := map[string]interface{}{
		"model": model,
		"data":  data,
	}
	res, err := vikingDBService.DoRequest(context.Background(), "EmbeddingV2", nil, vikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	//fmt.Println(res)
	var items map[string]interface{}
	if d, ok := res["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", data)
	} else if items, ok = d.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a map: %v", data)
	}
	return items, nil
}

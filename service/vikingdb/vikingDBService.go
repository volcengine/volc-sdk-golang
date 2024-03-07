package vikingdb

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type VikingDBService struct {
	Client *base.Client
}

func NewVikingDBService(host string, region string, ak string, sk string, scheme string) *VikingDBService {
	client := base.NewClient(getServiceInfo(host, region, ak, sk, scheme), getApiInfo())
	vikingDBService := &VikingDBService{
		Client: client,
	}
	return vikingDBService

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
		description = value.(string)
	}
	if value, exist := res["stat"]; exist {
		stat = value.(map[string]interface{})
	}
	if value, exist := res["create_time"]; exist {
		createTime = value.(string)
	}
	if value, exist := res["update_time"]; exist {
		updateTime = value.(string)
	}
	if value, exist := res["update_person"]; exist {
		updatePerson = value.(string)
	}
	if value, exist := res["indexes"]; exist {
		indexesName := value.([]interface{})
		for _, item := range indexesName {
			index, err := vikingDBService.GetIndex(collectionName, item.(string))
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
				fieldName = value.(string)
			}
			if value, exist := itemMap["field_type"]; exist {
				fieldType = value.(string)
			}
			if value, exist := itemMap["default_val"]; exist {
				defaultVal = value
			}
			if value, exist := itemMap["dim"]; exist {
				value := int64(value.(float64))
				dim = value
			}
			if value, exist := itemMap["pipeline_name"]; exist {
				pipelineName = value.(string)
			}
			if value, exist := res["primary_key"]; exist {
				if value.(string) == fieldName {
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
	var shard_count int64
	if value, exist := res["description"]; exist {
		description = value.(string)
	}
	if value, exist := res["status"]; exist {
		stat = value.(string)
	}
	if value, exist := res["create_time"]; exist {
		createTime = value.(string)
	}
	if value, exist := res["update_time"]; exist {
		updateTime = value.(string)
	}
	if value, exist := res["update_person"]; exist {
		updatePerson = value.(string)
	}
	if value, exist := res["cpu_quota"]; exist {
		value := int64(value.(float64))
		cpuQuota = value
	}
	if value, exist := res["partition_by"]; exist {
		partitionBy = value.(string)
	}
	if value, exist := res["index_cost"]; exist {
		indexCost = value.(map[string]interface{})
	}
	if value, exist := res["shard_count"]; exist {
		value := value.(float64)
		valueInt := int64(value)
		shard_count = valueInt
	}
	if value, exist := res["vector_index"]; exist {
		item := value.(map[string]interface{})
		vectorIndex = &VectorIndexParams{
			IndexType: "hnsw",
			Distance:  "ip",
			Quant:     "int8",
			HnswM:     20,
			HnswSef:   800,
			HnswCef:   400,
		}
		if v, e := item["index_type"]; e {
			vectorIndex.IndexType = v.(string)
		}
		if v, e := item["distance"]; e {
			vectorIndex.Distance = v.(string)
		}
		if v, e := item["quant"]; e {
			vectorIndex.Quant = v.(string)
		}
		if v, e := item["hnsw_m"]; e {
			v := int64(v.(float64))
			vectorIndex.HnswM = v
		}
		if v, e := item["hnsw_sef"]; e {
			v := int64(v.(float64))
			vectorIndex.HnswSef = v
		}
		if v, e := item["hnsw_cef"]; e {
			v := int64(v.(float64))
			vectorIndex.HnswCef = v
		}
	}
	if value, exist := res["range_index"]; exist {
		scalarIndex = value.([]interface{})
	}
	if value, exist := res["enum_index"]; exist {
		if scalarIndex == nil {
			scalarIndex = value.([]interface{})
		} else {
			for _, itemEnum := range value.([]interface{}) {
				flag := 1
				for _, itemScalar := range scalarIndex.([]interface{}) {
					if itemEnum.(string) == itemScalar.(string) {
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
		ShardCount:      shard_count,
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
	params["primary_key"] = primaryKey
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
	//fmt.Println(resData)
	res := resData["data"].(map[string]interface{})
	return vikingDBService.packageCollection(collectionName, res)
	//var description string
	//var stat map[string]interface{}
	//var fields []Field
	//var indexes []*Index
	//var createTime string
	//var updateTime string
	//var updatePerson string
	//if value, exist := res["description"]; exist {
	//	description = value.(string)
	//}
	//if value, exist := res["stat"]; exist {
	//	stat = value.(map[string]interface{})
	//}
	//if value, exist := res["create_time"]; exist {
	//	createTime = value.(string)
	//}
	//if value, exist := res["update_time"]; exist {
	//	updateTime = value.(string)
	//}
	//if value, exist := res["update_person"]; exist {
	//	updatePerson = value.(string)
	//}
	//if value, exist := res["indexes"]; exist {
	//	indexesName := value.([]interface{})
	//	for _, item := range indexesName {
	//		index, err := vikingDBService.GetIndex(collectionName, item.(string))
	//		if err != nil {
	//			return nil, err
	//		}
	//		indexes = append(indexes, index)
	//	}
	//}
	//if value, exist := res["fields"]; exist {
	//	fieldList := value.([]interface{})
	//	for _, item := range fieldList {
	//		var fieldName string
	//		var fieldType string
	//		var defaultVal interface{}
	//		var dim int64
	//		var isPrimaryKey bool
	//		var pipelineName string
	//		itemMap := item.(map[string]interface{})
	//		if value, exist := itemMap["field_name"]; exist {
	//			fieldName = value.(string)
	//		}
	//		if value, exist := itemMap["field_type"]; exist {
	//			fieldType = value.(string)
	//		}
	//		if value, exist := itemMap["default_val"]; exist {
	//			defaultVal = value
	//		}
	//		if value, exist := itemMap["dim"]; exist {
	//			value := int64(value.(float64))
	//			dim = value
	//		}
	//		if value, exist := itemMap["pipeline_name"]; exist {
	//			pipelineName = value.(string)
	//		}
	//		if value, exist := res["primary_key"]; exist {
	//			if value.(string) == fieldName {
	//				isPrimaryKey = true
	//			}
	//		}
	//		field := Field{
	//			FieldName:    fieldName,
	//			FieldType:    fieldType,
	//			DefaultVal:   defaultVal,
	//			Dim:          dim,
	//			IsPrimaryKey: isPrimaryKey,
	//			PipelineName: pipelineName,
	//		}
	//		fields = append(fields, field)
	//	}
	//
	//}
	//collection := &Collection{
	//	CollectionName:  collectionName,
	//	Fields:          fields,
	//	VikingDBService: vikingDBService,
	//	PrimaryKey:      res["primary_key"].(string),
	//	Indexes:         indexes,
	//	Description:     description,
	//	Stat:            stat,
	//	CreateTime:      createTime,
	//	UpdateTime:      updateTime,
	//	UpdatePerson:    updatePerson,
	//}
	//return collection, nil
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
	res := resData["data"].([]interface{})
	for _, itemMap := range res {
		//fmt.Println(item)
		item := itemMap.(map[string]interface{})
		collection, err := vikingDBService.packageCollection(item["collection_name"].(string), item)
		if err != nil {
			return nil, err
		}
		//var description string
		//var stat map[string]interface{}
		//var fields []Field
		//var indexes []*Index
		//var createTime string
		//var updateTime string
		//var updatePerson string
		//var collectionName string
		//if value, exist := item["collection_name"]; exist {
		//	collectionName = value.(string)
		//}
		//if value, exist := item["description"]; exist {
		//	description = value.(string)
		//}
		//if value, exist := item["stat"]; exist {
		//	stat = value.(map[string]interface{})
		//}
		//if value, exist := item["create_time"]; exist {
		//	createTime = value.(string)
		//}
		//if value, exist := item["update_time"]; exist {
		//	updateTime = value.(string)
		//}
		//if value, exist := item["update_person"]; exist {
		//	updatePerson = value.(string)
		//}
		//if value, exist := item["indexes"]; exist {
		//	indexesName := value.([]interface{})
		//	for _, item := range indexesName {
		//		index, err := vikingDBService.GetIndex(collectionName, item.(string))
		//		if err != nil {
		//			return nil, err
		//		}
		//		indexes = append(indexes, index)
		//	}
		//}
		//if value, exist := item["fields"]; exist {
		//	fieldList := value.([]interface{})
		//	for _, itemField := range fieldList {
		//		var fieldName string
		//		var fieldType string
		//		var defaultVal interface{}
		//		var dim int64
		//		var isPrimaryKey bool
		//		var pipelineName string
		//		itemMap := itemField.(map[string]interface{})
		//		if value, exist := itemMap["field_name"]; exist {
		//			fieldName = value.(string)
		//		}
		//		if value, exist := itemMap["field_type"]; exist {
		//			fieldType = value.(string)
		//		}
		//		if value, exist := itemMap["default_val"]; exist {
		//			defaultVal = value
		//		}
		//		if value, exist := itemMap["dim"]; exist {
		//			value := int64(value.(float64))
		//			dim = value
		//		}
		//		if value, exist := itemMap["pipeline_name"]; exist {
		//			pipelineName = value.(string)
		//		}
		//		if value, exist := item["primary_key"]; exist {
		//			if value.(string) == fieldName {
		//				isPrimaryKey = true
		//			}
		//		}
		//		field := Field{
		//			FieldName:    fieldName,
		//			FieldType:    fieldType,
		//			DefaultVal:   defaultVal,
		//			Dim:          dim,
		//			IsPrimaryKey: isPrimaryKey,
		//			PipelineName: pipelineName,
		//		}
		//		fields = append(fields, field)
		//	}
		//
		//}
		//collection := &Collection{
		//	CollectionName:  collectionName,
		//	Fields:          fields,
		//	VikingDBService: vikingDBService,
		//	PrimaryKey:      item["primary_key"].(string),
		//	Indexes:         indexes,
		//	Description:     description,
		//	Stat:            stat,
		//	CreateTime:      createTime,
		//	UpdateTime:      updateTime,
		//	UpdatePerson:    updatePerson,
		//}
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
	//fmt.Println(vikingDBService.convertMapToJson(params))
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
	res := resData["data"].(map[string]interface{})
	return vikingDBService.packageIndex(collectionName, indexName, res)
	//fmt.Println(resData)
	//var description string
	//var stat string
	//var createTime string
	//var updateTime string
	//var updatePerson string
	//var partitionBy string
	//var cpuQuota int64
	//var scalarIndex interface{}
	//var vectorIndex *VectorIndexParams
	//var indexCost interface{}
	//var shard_count int64
	//if value, exist := res["description"]; exist {
	//	description = value.(string)
	//}
	//if value, exist := res["status"]; exist {
	//	stat = value.(string)
	//}
	//if value, exist := res["create_time"]; exist {
	//	createTime = value.(string)
	//}
	//if value, exist := res["update_time"]; exist {
	//	updateTime = value.(string)
	//}
	//if value, exist := res["update_person"]; exist {
	//	updatePerson = value.(string)
	//}
	//if value, exist := res["cpu_quota"]; exist {
	//	value := int64(value.(float64))
	//	cpuQuota = value
	//}
	//if value, exist := res["partition_by"]; exist {
	//	partitionBy = value.(string)
	//}
	//if value, exist := res["index_cost"]; exist {
	//	indexCost = value.(map[string]interface{})
	//}
	//if value, exist := res["shard_count"]; exist {
	//	value := value.(float64)
	//	valueInt := int64(value)
	//	shard_count = valueInt
	//}
	//if value, exist := res["vector_index"]; exist {
	//	item := value.(map[string]interface{})
	//	vectorIndex = &VectorIndexParams{
	//		IndexType: "hnsw",
	//		Distance:  "ip",
	//		Quant:     "int8",
	//		HnswM:     20,
	//		HnswSef:   800,
	//		HnswCef:   400,
	//	}
	//	if v, e := item["index_type"]; e {
	//		vectorIndex.IndexType = v.(string)
	//	}
	//	if v, e := item["distance"]; e {
	//		vectorIndex.Distance = v.(string)
	//	}
	//	if v, e := item["quant"]; e {
	//		vectorIndex.Quant = v.(string)
	//	}
	//	if v, e := item["hnsw_m"]; e {
	//		v := int64(v.(float64))
	//		vectorIndex.HnswM = v
	//	}
	//	if v, e := item["hnsw_sef"]; e {
	//		v := int64(v.(float64))
	//		vectorIndex.HnswSef = v
	//	}
	//	if v, e := item["hnsw_cef"]; e {
	//		v := int64(v.(float64))
	//		vectorIndex.HnswCef = v
	//	}
	//}
	//if value, exist := res["range_index"]; exist {
	//	scalarIndex = value.([]interface{})
	//}
	//if value, exist := res["enum_index"]; exist {
	//	if scalarIndex == nil {
	//		scalarIndex = value.([]interface{})
	//	} else {
	//		for _, itemEnum := range value.([]interface{}) {
	//			flag := 1
	//			for _, itemScalar := range scalarIndex.([]interface{}) {
	//				if itemEnum.(string) == itemScalar.(string) {
	//					flag = 0
	//					break
	//				}
	//			}
	//			if flag == 1 {
	//				scalarIndex = append(scalarIndex.([]interface{}), itemEnum)
	//			}
	//
	//		}
	//	}
	//}
	//index := &Index{
	//	CollectionName:  collectionName,
	//	IndexName:       indexName,
	//	VectorIndex:     vectorIndex,
	//	ScalarIndex:     scalarIndex,
	//	Stat:            stat,
	//	VikingDBService: vikingDBService,
	//	Description:     description,
	//	CpuQuota:        cpuQuota,
	//	PartitionBy:     partitionBy,
	//	CreateTime:      createTime,
	//	UpdateTime:      updateTime,
	//	UpdatePerson:    updatePerson,
	//	IndexCost:       indexCost,
	//	ShardCount:      shard_count,
	//}
	//return index, err
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
	res := resData["data"].([]interface{})
	indexes := []*Index{}
	for _, itemMap := range res {
		//fmt.Println(itemMap)
		item := itemMap.(map[string]interface{})
		index, err := vikingDBService.packageIndex(collectionName, item["index_name"].(string), item)
		if err != nil {
			return nil, err
		}
		//var description string
		//var indexName string
		//var stat string
		//var createTime string
		//var updateTime string
		//var updatePerson string
		//var partitionBy string
		//var cpuQuota int64
		//var scalarIndex interface{}
		//var vectorIndex *VectorIndexParams
		//var indexCost interface{}
		//var shard_count int64
		//if value, exist := item["description"]; exist {
		//	description = value.(string)
		//}
		//if value, exist := item["index_name"]; exist {
		//	indexName = value.(string)
		//}
		//if value, exist := item["status"]; exist {
		//	stat = value.(string)
		//}
		//if value, exist := item["create_time"]; exist {
		//	createTime = value.(string)
		//}
		//if value, exist := item["update_time"]; exist {
		//	updateTime = value.(string)
		//}
		//if value, exist := item["update_person"]; exist {
		//	updatePerson = value.(string)
		//}
		//if value, exist := item["cpu_quota"]; exist {
		//	value := int64(value.(float64))
		//	cpuQuota = value
		//}
		//if value, exist := item["partition_by"]; exist {
		//	partitionBy = value.(string)
		//}
		//if value, exist := item["index_cost"]; exist {
		//	indexCost = value.(map[string]interface{})
		//}
		//if value, exist := item["shard_count"]; exist {
		//	value := value.(float64)
		//	valueInt := int64(value)
		//	shard_count = valueInt
		//}
		//if value, exist := item["vector_index"]; exist {
		//	itemVector := value.(map[string]interface{})
		//	vectorIndex = &VectorIndexParams{
		//		IndexType: "hnsw",
		//		Distance:  "ip",
		//		Quant:     "int8",
		//		HnswM:     20,
		//		HnswSef:   800,
		//		HnswCef:   400,
		//	}
		//	if v, e := itemVector["index_type"]; e {
		//		vectorIndex.IndexType = v.(string)
		//	}
		//	if v, e := itemVector["distance"]; e {
		//		vectorIndex.Distance = v.(string)
		//	}
		//	if v, e := itemVector["quant"]; e {
		//		vectorIndex.Quant = v.(string)
		//	}
		//	if v, e := itemVector["hnsw_m"]; e {
		//		v := int64(v.(float64))
		//		vectorIndex.HnswM = v
		//	}
		//	if v, e := itemVector["hnsw_sef"]; e {
		//		v := int64(v.(float64))
		//		vectorIndex.HnswSef = v
		//	}
		//	if v, e := itemVector["hnsw_cef"]; e {
		//		v := int64(v.(float64))
		//		vectorIndex.HnswCef = v
		//	}
		//}
		//if value, exist := item["range_index"]; exist {
		//	scalarIndex = value.([]interface{})
		//}
		//if value, exist := item["enum_index"]; exist {
		//	if scalarIndex == nil {
		//		scalarIndex = value.([]interface{})
		//	} else {
		//		for _, itemEnum := range value.([]interface{}) {
		//			flag := 1
		//			for _, itemScalar := range scalarIndex.([]interface{}) {
		//				if itemEnum.(string) == itemScalar.(string) {
		//					flag = 0
		//					break
		//				}
		//			}
		//			if flag == 1 {
		//				scalarIndex = append(scalarIndex.([]interface{}), itemEnum)
		//			}
		//
		//		}
		//	}
		//}
		//index := &Index{
		//	CollectionName:  collectionName,
		//	IndexName:       indexName,
		//	VectorIndex:     vectorIndex,
		//	ScalarIndex:     scalarIndex,
		//	Stat:            stat,
		//	VikingDBService: vikingDBService,
		//	Description:     description,
		//	CpuQuota:        cpuQuota,
		//	PartitionBy:     partitionBy,
		//	CreateTime:      createTime,
		//	UpdateTime:      updateTime,
		//	UpdatePerson:    updatePerson,
		//	IndexCost:       indexCost,
		//	ShardCount:      shard_count,
		//}
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
		"params":     embModel.params,
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
	for _, itemList := range res["data"].([]interface{}) {
		var floatList []float64
		itemList := itemList.([]interface{})
		for _, item := range itemList {
			floatList = append(floatList, item.(float64))
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
	return data["data"].(float64), err
}
func (vikingDBService *VikingDBService) BatchRerank(datas []map[string]interface{}) ([]float64, error) {
	params := map[string]interface{}{
		"datas": datas,
	}
	data, err := vikingDBService.DoRequest(context.Background(), "BatchRerank", nil, vikingDBService.convertMapToJson(params))
	var res []float64
	for _, score := range data["data"].([]interface{}) {
		res = append(res, score.(float64))
	}
	return res, err
}

package vikingdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

const (
	Vector        = "vector"
	Int64         = "int64"
	ListInt64     = "list<int64>"
	String        = "string"
	ListString    = "list<string>"
	Float32       = "float32"
	Bool          = "bool"
	Text          = "text"
	Sparse_Vector = "sparse_vector"

	L2     = "l2"
	IP     = "ip"
	COSINE = "cosine"

	FLAT        = "flat"
	HNSW        = "hnsw"
	IVF         = "ivf"
	DiskANN     = "DiskANN"
	HNSW_HYBRID = "hnsw_hybrid"
	Sort        = "sort"

	Float = "float"
	Int8  = "int8"
	Fix16 = "fix16"
	PQ    = "pq"

	Asc  = "asc"
	Desc = "desc"

	Custom = "custom"
	Auto   = "auto"

	Data_Import   = "data_import"
	Filter_Delete = "filter_delete"
	Data_Export   = "data_export"

	MAX_RETRIES         = 3
	INITIAL_RETRY_DELAY = 0.5
	MAX_RETRY_DELAY     = 8.0

	Init      = "init"
	Queued    = "queued"
	Running   = "running"
	Done      = "done"
	Fail      = "fail"
	Confirm   = "confirm"
	Confirmed = "confirmed"
)

type Field struct {
	FieldName    string
	FieldType    string
	DefaultVal   interface{}
	Dim          int64
	IsPrimaryKey bool
	PipelineName string
}
type VectorIndexParams struct {
	IndexType string
	Distance  string
	Quant     string
	HnswM     int64
	HnswCef   int64
	HnswSef   int64
}

func (vectorIndexParams *VectorIndexParams) dict(vectorIndex *VectorIndexParams) map[string]interface{} {
	res := map[string]interface{}{
		"distance":   IP,
		"index_type": HNSW,
		"quant":      Int8,
		"hnsw_m":     20,
		"hnsw_cef":   400,
		"hnsw_sef":   800,
	}
	if vectorIndex.Distance != "" {
		res["distance"] = vectorIndex.Distance
	}
	if vectorIndex.IndexType != "" {
		res["index_type"] = vectorIndex.IndexType
	}
	if vectorIndex.Quant != "" {
		res["quant"] = vectorIndex.Quant
	}
	if vectorIndex.HnswM != 0 {
		res["hnsw_m"] = vectorIndex.HnswM
	}
	if vectorIndex.HnswCef != 0 {
		res["hnsw_cef"] = vectorIndex.HnswCef
	}
	if vectorIndex.HnswSef != 0 {
		res["hnsw_sef"] = vectorIndex.HnswSef
	}
	return res

}

type Data struct {
	Id        interface{}
	Fields    map[string]interface{}
	Timestamp interface{}
	TTL       int64
	Score     float64
	Text      string
}
type VectorOrder struct {
	Vector interface{}
	Id     interface{}
}
type ScalarOrder struct {
	FieldName string
	Order     string
}
type TextObject struct {
	Text   string
	Url    string
	Base64 interface{}
}
type EmbModel struct {
	ModelName string
	Params    map[string]interface{}
}
type RawData struct {
	DataType string
	Text     string
	Image    string
}

type IndexOptions struct {
	vectorIndex *VectorIndexParams
	cpuQuota    int64
	description string
	partitionBy string
	scalarIndex []string
	shardCount  *int64
	shardPolicy string
}

func NewIndexOptions() *IndexOptions {
	indexOptions := &IndexOptions{
		vectorIndex: nil,
		cpuQuota:    2,
		description: "",
		partitionBy: "",
		scalarIndex: nil,
		shardCount:  nil,
		shardPolicy: "",
	}
	return indexOptions
}
func (indexOptions *IndexOptions) SetVectorIndex(vectorIndex *VectorIndexParams) *IndexOptions {
	indexOptions.vectorIndex = vectorIndex
	return indexOptions
}
func (indexOptions *IndexOptions) SetCpuQuota(cpuQuota int64) *IndexOptions {
	indexOptions.cpuQuota = cpuQuota
	return indexOptions
}
func (indexOptions *IndexOptions) SetDescription(description string) *IndexOptions {
	indexOptions.description = description
	return indexOptions
}
func (indexOptions *IndexOptions) SetPartitionBy(partitionBy string) *IndexOptions {
	indexOptions.partitionBy = partitionBy
	return indexOptions
}
func (indexOptions *IndexOptions) SetScalarIndex(scalarIndex []string) *IndexOptions {
	indexOptions.scalarIndex = scalarIndex
	return indexOptions
}
func (indexOptions *IndexOptions) SetShardCount(shardCount int64) *IndexOptions {
	indexOptions.shardCount = &shardCount
	return indexOptions
}
func (indexOptions *IndexOptions) SetShardPolicy(shardPolicy string) *IndexOptions {
	indexOptions.shardPolicy = shardPolicy
	return indexOptions
}

type UpdateCollectionOptions struct {
	fields      []Field
	description *string
}

func NewUpdateCollectionOptions() *UpdateCollectionOptions {
	updateCollectionOptions := &UpdateCollectionOptions{
		fields:      nil,
		description: nil,
	}
	return updateCollectionOptions
}
func (updateCollectionOptions *UpdateCollectionOptions) SetFields(fields []Field) *UpdateCollectionOptions {
	updateCollectionOptions.fields = fields
	return updateCollectionOptions
}
func (updateCollectionOptions *UpdateCollectionOptions) SetDescription(description string) *UpdateCollectionOptions {
	updateCollectionOptions.description = &description
	return updateCollectionOptions
}

type UpdateIndexOptions struct {
	description *string
	cpuQuota    *int64
	scalarIndex []string
	shardCount  *int64
}

func NewUpdateIndexOptions() *UpdateIndexOptions {
	updateIndexOptions := &UpdateIndexOptions{
		description: nil,
		cpuQuota:    nil,
		scalarIndex: nil,
		shardCount:  nil,
	}
	return updateIndexOptions
}
func (updateIndexOptions *UpdateIndexOptions) SetDescription(description string) *UpdateIndexOptions {
	updateIndexOptions.description = &description
	return updateIndexOptions
}
func (updateIndexOptions *UpdateIndexOptions) SetCpuQuota(cpuQuota int64) *UpdateIndexOptions {
	updateIndexOptions.cpuQuota = &cpuQuota
	return updateIndexOptions
}
func (updateIndexOptions *UpdateIndexOptions) SetScalarIndex(scalarIndex []string) *UpdateIndexOptions {
	updateIndexOptions.scalarIndex = scalarIndex
	return updateIndexOptions
}
func (updateIndexOptions *UpdateIndexOptions) SetShardCount(shardCount int64) *UpdateIndexOptions {
	updateIndexOptions.shardCount = &shardCount
	return updateIndexOptions
}

type ParamOptions struct {
	AsyncUpsert bool
}
type ParamOption func(*ParamOptions)

func WithAsyncUpsert(yes bool) ParamOption {
	return func(co *ParamOptions) {
		co.AsyncUpsert = yes
	}
}

func ParseJsonUseNumber2(input []byte, target interface{}) error {
	var d *json.Decoder
	var err error
	d = json.NewDecoder(bytes.NewBuffer(input))
	if d == nil {
		return fmt.Errorf("ParseJsonUseNumber init NewDecoder failed")
	}
	d.UseNumber()
	err = d.Decode(&target)
	if err != nil {
		return fmt.Errorf("ParseJsonUseNumber Decode failed %v", err)
	}
	return nil
}

func SerilizeToJsonBytesUseNumber(source interface{}) ([]byte, error) {
	//  buffer := make([]byte, 0)
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(source)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func HasDecimalPart(f float64) bool {
	_, decimal := math.Modf(f)
	return decimal > 1e-16
}

func ParseJsonInt64Field(default_val_i interface{}) (int64, error) {
	if default_val_i == nil {
		return 0, nil
	}
	if default_val_int64, ok := default_val_i.(int64); ok {
		return default_val_int64, nil
	} else if default_val_float64, ok := default_val_i.(float64); ok {
		default_val_int64 := int64(default_val_float64)
		if HasDecimalPart(default_val_float64) {
			return 0, fmt.Errorf("can not convert float64 to int64: not int")
		}
		return default_val_int64, nil
	} else if default_val_string, ok := default_val_i.(string); ok {
		if default_val_int64, err := strconv.ParseInt(default_val_string, 10, 64); err != nil {
			return 0, fmt.Errorf("can not convert string to int64: %v", err)
		} else {
			return default_val_int64, nil
		}
	} else if default_val_jsonNumber, ok := default_val_i.(json.Number); ok {
		if default_val_int64, err := default_val_jsonNumber.Int64(); err != nil {
			return 0, fmt.Errorf("can not convert json.Number[%v] to int64: %v", default_val_jsonNumber, err)
		} else {
			return default_val_int64, nil
		}
	} else {
		return 0, fmt.Errorf("can not convert %s to int64", reflect.TypeOf(default_val_i))
	}
}

func ParseJsonFloat64Field(default_val_i interface{}) (float64, error) {
	if default_val_i == nil {
		return 0, nil
	}
	if val, ok := default_val_i.(float64); ok {
		return val, nil
	} else if default_val_jsonNumber, ok := default_val_i.(json.Number); ok {
		if default_val_float64, err := default_val_jsonNumber.Float64(); err != nil {
			return 0, fmt.Errorf("can not convert json.Number[%v] to float64: %v", default_val_jsonNumber, err)
		} else {
			return default_val_float64, nil
		}
	} else if default_val_string, ok := default_val_i.(string); ok {
		default_val_jsonNumber := json.Number(default_val_string)
		if default_val_float64, err := default_val_jsonNumber.Float64(); err != nil {
			return 0, fmt.Errorf("can not convert json.Number[%v] to float64: %v", default_val_jsonNumber, err)
		} else {
			return default_val_float64, nil
		}
	} else {
		return 0, fmt.Errorf("invalid field for float type")
	}
}

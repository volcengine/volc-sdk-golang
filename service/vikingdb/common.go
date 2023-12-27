package vikingdb

const (
	Vector     = "vector"
	Int64      = "int64"
	ListInt64  = "list<int64>"
	String     = "string"
	ListString = "list<string>"
	Float32    = "float32"
	Bool       = "bool"
	Text       = "text"

	L2     = "l2"
	IP     = "ip"
	COSINE = "cosine"
	FLAT   = "flat"
	HNSW   = "hnsw"
	Float  = "float"
	Int8   = "int8"

	Asc  = "asc"
	Desc = "desc"
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
	params    interface{}
}
type RawData struct {
	DataType string
	Text     string
}

type IndexOptions struct {
	vectorIndex *VectorIndexParams
	cpuQuota    int64
	description string
	partitionBy string
	scalarIndex []string
	shardCount  *int64
}

func NewIndexOptions() *IndexOptions {
	indexOptions := &IndexOptions{
		vectorIndex: nil,
		cpuQuota:    2,
		description: "",
		partitionBy: "",
		scalarIndex: nil,
		shardCount:  nil,
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
}

func NewUpdateIndexOptions() *UpdateIndexOptions {
	updateIndexOptions := &UpdateIndexOptions{
		description: nil,
		cpuQuota:    nil,
		scalarIndex: nil,
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

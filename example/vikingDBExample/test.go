package main

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/vikingdb"
	"math/rand"
	"time"
)

func main() {
	service := vikingdb.NewVikingDBService("host", "region", "ak", "sk", "http")

	rerank, err := service.Rerank("退改", "如果您需要人工服务，可以拨打人工客服电话：4006660921", "转人工")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rerank)

	//fields := []vikingdb.Field{
	//	{
	//		FieldName:    "doc_id",
	//		FieldType:    vikingdb.String,
	//		IsPrimaryKey: true,
	//	},
	//	{
	//		FieldName: "text_vector",
	//		FieldType: vikingdb.Vector,
	//		Dim:       12,
	//	},
	//	{
	//		FieldName:  "like",
	//		FieldType:  vikingdb.Int64,
	//		DefaultVal: 0,
	//	},
	//	{
	//		FieldName: "price",
	//		FieldType: vikingdb.Float32,
	//		Dim:       12,
	//	},
	//	{
	//		FieldName:  "author",
	//		FieldType:  vikingdb.ListString,
	//		DefaultVal: []string{},
	//	},
	//	{
	//		FieldName:  "aim",
	//		FieldType:  vikingdb.Bool,
	//		DefaultVal: true,
	//	},
	//}
	//collection, err := service.CreateCollection("go", fields, "this is a go example")
	//if err != nil {
	//	print(err.Error())
	//}
	//print(collection)

	//collection, err := service.GetCollection("go")
	//if err != nil {
	//	print(err.Error())
	//}
	//fmt.Println(collection)

	//err := service.DropCollection("go")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//collections, err := service.ListCollections()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//for _, collection := range collections {
	//	fmt.Println(collection)
	//	fmt.Println("-------------------")
	//}

	//vectorIndex := &vikingdb.VectorIndexParams{
	//	Distance:  vikingdb.COSINE,
	//	IndexType: vikingdb.HNSW,
	//}
	//indexOptions := vikingdb.NewIndexOptions().SetVectorIndex(vectorIndex).SetCpuQuota(2).SetDescription("this is an index").SetPartitionBy("").SetScalarIndex([]string{"price", "like"}).SetShardCount(12)
	//index, err := service.CreateIndex("go", "goIndex", indexOptions)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(index)

	//index, err := service.GetIndex("go", "goIndex")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(index)

	//err := service.DropIndex("go", "goIndex")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//indexes, err := service.ListIndexes("go")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, index := range indexes {
	//	fmt.Println(index)
	//}

	//collection, _ := service.GetCollection("go")
	//field1 := map[string]interface{}{
	//	"doc_id":      "111",
	//	"text_vector": genRandomVector(12),
	//	"like":        1,
	//	"price":       1.11,
	//	"author":      []string{"gy"},
	//	"aim":         true,
	//}
	//field2 := map[string]interface{}{
	//	"doc_id":      "222",
	//	"text_vector": genRandomVector(12),
	//	"like":        2,
	//	"price":       2.22,
	//	"author":      []string{"gy", "xjq"},
	//	"aim":         false,
	//}
	//field3 := map[string]interface{}{
	//	"doc_id":      "333",
	//	"text_vector": genRandomVector(12),
	//	"like":        3,
	//	"price":       3.33,
	//	"author":      []string{"gy"},
	//	"aim":         true,
	//}
	//field4 := map[string]interface{}{
	//	"doc_id":      "1qw",
	//	"text_vector": genRandomVector(12),
	//	"like":        4,
	//	"price":       4.44,
	//	"author":      []string{"gy"},
	//	"aim":         true,
	//}
	//data1 := vikingdb.Data{
	//	Fields: field1,
	//	TTL:    100000,
	//}
	//data2 := vikingdb.Data{
	//	Fields: field2,
	//	TTL:    200000,
	//}
	//data3 := vikingdb.Data{
	//	Fields: field3,
	//	TTL:    100000,
	//}
	//data4 := vikingdb.Data{
	//	Fields: field4,
	//}
	//datas := []vikingdb.Data{
	//	data1,
	//	data2,
	//	data3,
	//	data4,
	//}
	//err := collection.UpsertData(datas)
	//if err != nil {
	//	print(err.Error())
	//}

	//collection, _ := service.GetCollection("go")
	//err := collection.DeleteData([]string{"111", "222"})
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//err := collection.DeleteData("111")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//collection, _ := service.GetCollection("go")
	//res, err := collection.FetchData([]string{"111", "222", "333", "444"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//index, _ := service.GetIndex("go", "goIndex")
	//searchOption := vikingdb.NewSearchOptions().SetFilter(map[string]interface{}{"op": "range", "field": "price", "lt": 3.5}).SetLimit(5).SetOutputFields([]string{"doc_id", "like", "text_vector", "price"})
	//res, err := index.SearchById("111", searchOption)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//index, _ := service.GetIndex("go", "goIndex")
	//searchOption := vikingdb.NewSearchOptions().SetFilter(map[string]interface{}{"op": "range", "field": "price", "lt": 3.5}).SetLimit(5).SetOutputFields([]string{"doc_id", "like", "text_vector", "price"})
	//res, err := index.SearchByVector(genRandomVector(12), searchOption)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//index, _ := service.GetIndex("go", "goIndex")
	//searchOption := vikingdb.NewSearchOptions().SetFilter(map[string]interface{}{"op": "range", "field": "price", "lt": 3.5}).SetLimit(5).SetOutputFields([]string{"doc_id", "like", "text_vector", "price"})
	//res, err := index.Search(vikingdb.VectorOrder{Id: "111"}, searchOption)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//index, _ := service.GetIndex("go", "goIndex")
	//searchOption := vikingdb.NewSearchOptions().SetFilter(map[string]interface{}{"op": "range", "field": "price", "lt": 3.5}).SetLimit(5).SetOutputFields([]string{"doc_id", "like", "text_vector", "price"})
	//res, err := index.Search(vikingdb.ScalarOrder{FieldName: "price", Order: vikingdb.Desc}, searchOption)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//index, _ := service.GetIndex("go", "goIndex")
	//searchOption := vikingdb.NewSearchOptions().SetOutputFields([]string{"doc_id", "like", "text_vector", "price"})
	//res, err := index.FetchData([]string{"111", "222"}, searchOption)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//collection1, _ := service.GetCollection("goTest")
	//fmt.Println(collection1)
	//fields := []vikingdb.Field{
	//	{
	//		FieldName: "new_field1",
	//		FieldType: vikingdb.String,
	//	},
	//	{
	//		FieldName:  "new_field2",
	//		FieldType:  vikingdb.Float32,
	//		DefaultVal: 100,
	//	},
	//}
	//updateCollectionOptions := vikingdb.NewUpdateCollectionOptions().SetFields(fields).SetDescription("new field")
	//err := service.UpdateCollection("goTest", updateCollectionOptions)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//collection2, _ := service.GetCollection("goTest")
	//fmt.Println(collection2)

	//list := []vikingdb.RawData{
	//	{
	//		DataType: "text",
	//		Text:     "hello1",
	//	},
	//	{
	//		DataType: "text",
	//		Text:     "hello2",
	//	},
	//}
	//res, err := service.Embedding(vikingdb.EmbModel{ModelName: "bge_large_zh"}, list)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//	fmt.Println("------------------------------")
	//}

	//fields := []vikingdb.Field{
	//	{
	//		FieldName:    "doc_id",
	//		FieldType:    vikingdb.String,
	//		IsPrimaryKey: true,
	//	},
	//	{
	//		FieldName:    "text_test",
	//		FieldType:    vikingdb.Text,
	//		PipelineName: "text_split_bge_large_zh",
	//	},
	//}
	//service.CreateCollection("go_text_new", fields, "test")

	//collection, _ := service.GetCollection("go_text_new")
	//field1 := map[string]interface{}{
	//	"doc_id":    "111",
	//	"text_test": map[string]interface{}{"text": "this is one"},
	//}
	//field2 := map[string]interface{}{
	//	"doc_id":    "222",
	//	"text_test": map[string]interface{}{"text": "this is two"},
	//}
	//data1 := vikingdb.Data{
	//	Fields: field1,
	//	TTL:    100000,
	//}
	//data2 := vikingdb.Data{
	//	Fields: field2,
	//	TTL:    200000,
	//}
	//datas := []vikingdb.Data{
	//	data1,
	//	data2,
	//}
	//collection.UpsertData(datas)

	//vectorIndex := &vikingdb.VectorIndexParams{
	//	Distance:  vikingdb.COSINE,
	//	IndexType: vikingdb.HNSW,
	//}
	//index, err := service.CreateIndex("go_text_new", "goIndex", vectorIndex,
	//	2, "this is an index", "", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(index)

	//index, _ := service.GetIndex("go_text_new", "goIndex")
	//searchOption := vikingdb.NewSearchOptions().SetLimit(5).SetOutputFields([]string{"doc_id"})
	//res, err := index.SearchByText(vikingdb.TextObject{Text: "example"}, searchOption)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//index, _ := service.GetIndex("example", "example_index")
	//fmt.Println(index)
	//params := vikingdb.NewUpdateIndexOptions().SetDescription("go").SetCpuQuota(5).SetScalarIndex([]string{"like", "aim"})
	//err := service.UpdateIndex("example", "example_index", params)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//index1, _ := service.GetIndex("example", "example_index")
	//fmt.Println(index1)

}
func genRandomVector(dim int) []float64 {
	res := []float64{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dim; i++ {
		res = append(res, rand.Float64()-0.5)
	}
	return res
}

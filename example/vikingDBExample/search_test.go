package vikingdbtest

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vikingdb"
)

func Test_Search_By_Raw(t *testing.T) {
	index, err := service.GetIndex("test_coll_for_sdk", "index_hnsw_hybrid")
	if err != nil {
		panic(err)
	}
	raw := vikingdb.TextObject{
		Text: "this is eeecejjdggahifbfgdbdadabbciaejgaihidbhhacifcgfbdig",
	}
	filter := map[string]interface{}{
		"op":    "range",
		"field": "f_int64",
		"gt":    90,
	}
	searchOpt := vikingdb.NewSearchOptions().SetFilter(filter).SetLimit(5).
		SetNeedInstruction(false).SetRetry(true)
	datas, err := index.SearchByText(raw, searchOpt)
	if err != nil {
		panic(err)
	}
	for _, data := range datas {
		fmt.Println(data.Fields, data.Score)
	}
}

func Test_Search_With_PrimaryKey_Filter(t *testing.T) {
	index, err := service.GetIndex("test_coll_for_sdk", "index_hnsw_hybrid")
	if err != nil {
		panic(err)
	}
	filter := map[string]interface{}{
		"op":    "range",
		"field": "f_int64",
		"gt":    10,
	}
	searchOpt := vikingdb.NewSearchOptions().SetFilter(filter).SetLimit(5).
		SetPrimaryKeyIn([]interface{}{"1", "2", "3", "4"}).
		SetPrimaryKeyNotIn([]interface{}{"1"}).
		SetRetry(true)
	datas, err := index.Search(nil, searchOpt)
	if err != nil {
		panic(err)
	}
	for _, data := range datas {
		fmt.Println(data.Fields, data.Score)
	}
}

func Test_Search_With_PostProcessOps(t *testing.T) {
	index, err := service.GetIndex("test_coll_for_sdk", "index_hnsw_hybrid")
	if err != nil {
		panic(err)
	}
	filter := map[string]interface{}{
		"op":    "range",
		"field": "f_int64",
		"gt":    10,
	}
	searchOpt := vikingdb.NewSearchOptions().SetFilter(filter).SetLimit(15).
		SetPostProcessOps([]map[string]interface{}{
			{"op": "enum_freq_limiter", "field": "f_string", "threshold": 1},
			{"op": "string_like", "field": "f_string", "pattern": "doc2%"},
		}).
		SetPostProcessInputLimit(10).
		SetRetry(true)
	datas, err := index.Search(nil, searchOpt)
	if err != nil {
		panic(err)
	}
	for _, data := range datas {
		fmt.Println(data.Fields, data.Score)
	}
}

func Test_Search_Agg(t *testing.T) {
	index, err := service.GetIndex("test_coll_for_sdk", "index_hnsw_hybrid")
	if err != nil {
		panic(err)
	}
	searchAggOpt := vikingdb.NewSearchAggOptions().
		SetFilter(map[string]interface{}{
			"op":    "range",
			"field": "f_int64",
			"gt":    10,
		}).
		SetAgg(map[string]interface{}{
			"op":    "count",
			"field": "f_string",
			"gt":    0,
		})
	aggResult, err := index.SearchAgg(searchAggOpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(aggResult.AggOp)
	fmt.Println(aggResult.GroupByField)
	fmt.Println(aggResult.AggResult)
}

func Test_Index_Sort(t *testing.T) {
	index, err := service.GetIndex("test_coll_for_sdk_1", "index_sort")
	if err != nil {
		panic(err)
	}
	sortOpt := vikingdb.NewSortOptions().
		SetQueryVector([]float64{0.0, 1.9, 2.2, -1.1}).
		SetPrimaryKeys([]interface{}{"doc-x", "doc0", "doc1", "doc2", "doc3", "doc4", "doc5", "doc6"}).
		SetRetry(true)
	indexSortResult, err := index.Sort(sortOpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(indexSortResult.PrimaryKeyNotExist)
	for _, item := range indexSortResult.SortResult {
		fmt.Printf("%v, %.5f\n", item.PrimaryKey, item.Score)
	}
}

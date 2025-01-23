package vikingdbtest

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vikingdb"
)

func Test_Index_Create_Sort(t *testing.T) {
	opt := vikingdb.NewIndexOptions().SetShardPolicy(vikingdb.Custom).SetShardCount(1).
		SetVectorIndex(&vikingdb.VectorIndexParams{IndexType: vikingdb.Sort, Distance: vikingdb.IP, Quant: vikingdb.Float})
	index, err := service.CreateIndex("test_coll_for_sdk_1", "index_sort", opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(index.CollectionName)
	fmt.Println(index.IndexName)
}

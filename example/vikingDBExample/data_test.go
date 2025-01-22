package vikingdbtest

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vikingdb"
)

func Test_Data_Upsert(t *testing.T) {
	collection, err := service.GetCollection("test_coll_for_sdk_1")
	if err != nil {
		panic(err)
	}
	var datas []vikingdb.Data
	for i := 0; i < 100; i++ {
		datas = append(datas, vikingdb.Data{
			Fields: map[string]interface{}{
				"f_id":            "doc" + strconv.Itoa(i+1),
				"f_int64":         i + 1,
				"f_vector":        []float32{float32(i), rand.Float32(), rand.Float32(), rand.Float32()},
				"f_sparse_vector": map[string]interface{}{"$": strconv.Itoa(i + 1)},
			},
		})
	}
	err = collection.UpsertData(datas)
	if err != nil {
		panic(err)
	}
}

func Test_Data_Fetch_By_Collection(t *testing.T) {
	collection, err := service.GetCollection("test_coll_for_sdk_2")
	if err != nil {
		panic(err)
	}
	datas, err := collection.FetchData([]string{"doc1", "doc2"})
	if err != nil {
		panic(err)
	}
	for _, data := range datas {
		fmt.Println(data.Id, data.Fields)
	}
}

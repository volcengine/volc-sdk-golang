package vikingdbtest

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vikingdb"
)

func Test_Collection_Create_1(t *testing.T) {
	fields := []vikingdb.Field{
		{FieldName: "f_id", FieldType: vikingdb.String, IsPrimaryKey: true},
		{FieldName: "f_int64", FieldType: vikingdb.Int64},
		{FieldName: "f_vector", FieldType: vikingdb.Vector, Dim: 4},
	}
	collection, err := service.CreateCollection("test_coll_for_sdk_1", fields, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(collection.CollectionName)
}

func Test_Collection_Create_2(t *testing.T) {
	fields := []vikingdb.Field{
		{FieldName: "f_id", FieldType: vikingdb.String, IsPrimaryKey: true},
		{FieldName: "f_int64", FieldType: vikingdb.Int64},
		{FieldName: "f_vector", FieldType: vikingdb.Vector, Dim: 4},
		{FieldName: "f_sparse_vector", FieldType: vikingdb.Sparse_Vector},
	}
	collection, err := service.CreateCollection("test_coll_for_sdk_2", fields, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(collection.CollectionName)
}

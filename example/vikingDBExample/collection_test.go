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

func Test_Collection_Create_3(t *testing.T) {
	fields := []vikingdb.Field{
		{FieldName: "f_id", FieldType: vikingdb.String, IsPrimaryKey: true},
		{FieldName: "f_int64", FieldType: vikingdb.Int64},
		{FieldName: "f_text1", FieldType: vikingdb.Text},
		{FieldName: "f_text2", FieldType: vikingdb.Text},
		{FieldName: "f_image1", FieldType: vikingdb.Image},
		{FieldName: "f_image2", FieldType: vikingdb.Image},
	}
	vectorize := []*vikingdb.VectorizeTuple{
		vikingdb.NewVectorizeTuple().
			SetDense(vikingdb.NewVectorizeModelConf().
				SetTextField("f_text1").
				SetImageField("f_image1").
				SetModelName("doubao-embedding-vision").
				SetModelVersion("241215").
				SetDim(3072)).
			SetSparse(vikingdb.NewVectorizeModelConf().
				SetTextField("f_text2").
				SetModelName("bge-m3")),
	}
	collection, err := service.CreateCollection("test_coll_for_sdk_with_vectorize", fields, "", vectorize)
	if err != nil {
		panic(err)
	}
	fmt.Println(collection.CollectionName)
}

func Test_Collection_Info(t *testing.T) {
	collection, err := service.GetCollection("test_coll_for_sdk_with_vectorize")
	if err != nil {
		panic(err)
	}
	fmt.Println(collection.CollectionName)
}

func Test_Collection_Delete(t *testing.T) {
	err := service.DropCollection("test_coll_for_sdk_with_vectorize")
	if err != nil {
		panic(err)
	}
}

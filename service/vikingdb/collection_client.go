package vikingdb

type CollectionClient struct {
	Collection //
}

func NewCollectionClient(collection string, host string, region string, ak string, sk string, schema string) *CollectionClient {
	vikingdbService := NewVikingDBService(host, region, ak, sk, schema)
	return &CollectionClient{
		Collection: Collection{
			CollectionName: collection,
			VikingDBService: vikingdbService,
			isClient: true,
			PrimaryKey: "__primary_key",
		},
	}
}
package vikingdb


type IndexClient struct {
	Index //
}

func NewIndexClient(collection string, index string, host string, region string, ak string, sk string, schema string) *IndexClient {
	vikingdbService := NewVikingDBService(host, region, ak, sk, schema)
	return &IndexClient{
		Index: Index{
			CollectionName: collection,
			IndexName:      index,
			VikingDBService: vikingdbService,
			isClient: true,
			primaryKey: "__primary_key",
		},
	}
}


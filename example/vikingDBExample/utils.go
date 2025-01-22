package vikingdbtest

import (
	"os"

	"github.com/volcengine/volc-sdk-golang/service/vikingdb"
)

var service = initVikingDBService()

func initVikingDBService() *vikingdb.VikingDBService {
	ak := os.Getenv("ak")
	sk := os.Getenv("sk")
	scheme := "http"
	host := "api-vikingdb.volces.com"
	region := "cn-beijing"
	vikingdbService := vikingdb.NewVikingDBService(host, region, ak, sk, scheme)
	vikingdbService.SetConnectionTimeout(10)
	return vikingdbService
}

package consumer

import (
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"os"
)

func NewClientWithEnv() tls.Client {
	return tls.NewClient(os.Getenv("LOG_SERVICE_ENDPOINT"), os.Getenv("LOG_SERVICE_AK"),
		os.Getenv("LOG_SERVICE_SK"), "", os.Getenv("LOG_SERVICE_REGION"))
}

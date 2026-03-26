package tls

import "os"

func hasLogServiceEnv() bool {
	return os.Getenv("LOG_SERVICE_ENDPOINT") != "" &&
		os.Getenv("LOG_SERVICE_AK") != "" &&
		os.Getenv("LOG_SERVICE_SK") != "" &&
		os.Getenv("LOG_SERVICE_REGION") != ""
}


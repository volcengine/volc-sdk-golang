package url_sign

import (
	"time"
)

const (
	lengthTokenOID    = 2
	lengthTokenBucket = 4

	urlPathTemplate    = "%s/%s%s"
	authPathTemplate   = "/video/%s/%s/"
	expirationTemplate = "%08x"

	schemeHTTP  = "http"
	schemeHTTPS = "https"

	ttlOneHour = 1 * time.Hour
	ttlOneYear = 365 * 24 * time.Hour
)

var (
	now = time.Now
)

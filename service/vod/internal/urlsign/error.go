package url_sign

import (
	"errors"
)

var (
	ErrOIDInvalid    = errors.New("[pubsdk] oid is invalid")
	ErrBucketInvalid = errors.New("[pubsdk] bucket is invalid")
	ErrNoCDN         = errors.New("[pubsdk] there is no cdn")
)

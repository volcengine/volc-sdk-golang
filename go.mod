module github.com/volcengine/volc-sdk-golang

go 1.14

require (
	github.com/Shopify/sarama v1.30.1
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/volcengine/volc-sdk-golang => ../volc-sdk-golang

module github.com/volcengine/volc-sdk-golang

go 1.14

require (
	github.com/Shopify/sarama v1.30.1
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/volcengine/volc-sdk-golang => ../volc-sdk-golang

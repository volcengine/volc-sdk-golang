module github.com/volcengine/volc-sdk-golang

go 1.14

require (
	github.com/Shopify/sarama v1.30.1
	github.com/actgardner/gogen-avro/v10 v10.2.1
	github.com/apache/rocketmq-client-go/v2 v2.1.2
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/go-kit/kit v0.12.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-querystring v1.1.0
	github.com/google/martian v2.1.0+incompatible
	github.com/google/uuid v1.6.0
	github.com/pierrec/lz4 v2.6.1+incompatible
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.9.0
	github.com/twpayne/go-geom v1.5.4
	golang.org/x/crypto v0.21.0
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225
	golang.org/x/net v0.22.0
	google.golang.org/protobuf v1.33.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/volcengine/volc-sdk-golang => ../volc-sdk-golang

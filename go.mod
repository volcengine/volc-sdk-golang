module github.com/volcengine/volc-sdk-golang

go 1.14

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/Shopify/sarama v1.30.1
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/demdxx/gocast v1.2.0 // indirect
	github.com/go-kit/kit v0.12.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/go-querystring v1.1.0
	github.com/google/martian v2.1.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.11.0
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6
	golang.org/x/net v0.12.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/volcengine/volc-sdk-golang => ../volc-sdk-golang

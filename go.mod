module github.com/volcengine/volc-sdk-golang

go 1.14

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Shopify/sarama v1.30.1
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/frankban/quicktest v1.14.0 // indirect
	github.com/go-kit/kit v0.12.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0
	github.com/google/martian v2.1.0+incompatible
	github.com/google/uuid v1.6.0
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible
	github.com/pkg/errors v0.9.1
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.21.0
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225
	golang.org/x/net v0.22.0
	google.golang.org/protobuf v1.33.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/volcengine/volc-sdk-golang => ../volc-sdk-golang

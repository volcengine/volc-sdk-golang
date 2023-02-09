package common

type LoggerConfig struct {
	LogLevel      string
	LogFileName   string
	IsJsonType    bool
	LogMaxSize    int
	LogMaxBackups int
	LogCompress   bool
}

type ClientConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	SecurityToken   string
	Region          string
}

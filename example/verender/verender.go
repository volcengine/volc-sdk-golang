package main

import (
	"github.com/volcengine/volc-sdk-golang/service/verender"
)

func getVerenderInstance() *verender.Verender {
	v := verender.NewVerenderInstance()
	v.SetFtransClientAddr("127.0.0.1:8280")   // set if transport by ftrans client
	v.SetFtransProxyAddr("10.211.55.3:30001") // set if transport by ftrans proxy

	v.Client.SetAccessKey("your ak")
	v.Client.SetSecretKey("your sk")

	return v
}

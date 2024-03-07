package maas

import (
	"io"
)

type ResponseContent interface {
	io.ReadCloser
	GetRequestId() string
}

type binaryResponseContent struct {
	io.ReadCloser
	requestId string
}

func (b *binaryResponseContent) GetRequestId() string {
	return b.requestId
}

func NewBinaryResponseContent(i io.ReadCloser, requestId string) ResponseContent {
	return &binaryResponseContent{i, requestId}
}

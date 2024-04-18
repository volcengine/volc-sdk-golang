package maas

import (
	"context"
	"io"
)

type ResponseContent interface {
	io.ReadCloser
	GetRequestId() string
}

type binaryResponseContent struct {
	io.ReadCloser
	requestId string
	cancel    context.CancelFunc
}

func (b *binaryResponseContent) GetRequestId() string {
	return b.requestId
}

func NewBinaryResponseContent(i io.ReadCloser, requestId string, cancel context.CancelFunc) ResponseContent {
	return &binaryResponseContent{i, requestId, cancel}
}

func (b *binaryResponseContent) Close() error {
	b.ReadCloser.Close()
	b.cancel()
	return nil
}

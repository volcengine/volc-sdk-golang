package main

import (
	"bytes"
	"context"
	"strings"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/volcengine/volc-sdk-golang/example/dts/data-subscription-demo/avro/avro"
)

func main() {
	nameserver := "your nameserver address"
	topic := "your topic"
	group := "your group"
	accessKey := "your access key"
	secretKey := "your secret key"

	cli, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(strings.Split(nameserver, ","))),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumerOrder(true),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}),
	)

	err := cli.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			record, err := avro.DeserializeRecord(bytes.NewReader(msg.Body))
			if err != nil {
				panic(err)
			}
			printRecord(&record)
		}
		return consumer.ConsumeSuccess, nil

	})
	if err != nil {
		panic(err.Error())
	}

	err = cli.Start()
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(time.Hour)
	cli.Shutdown()
}

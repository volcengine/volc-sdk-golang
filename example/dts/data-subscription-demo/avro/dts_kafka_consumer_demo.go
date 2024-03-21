package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/volcengine/volc-sdk-golang/example/dts/data-subscription-demo/avro/avro"
)

func main() {
	config := sarama.NewConfig()
	brokers := "your brokers addrress"
	topic := "your topic"
	group := "your group"
	username := "your username"
	password := "your password"

	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Net.SASL.Enable = true
	config.Net.SASL.User = username
	config.Net.SASL.Password = password

	cons, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		panic(err)
	}
	defer cons.Close()

	handler := &Handler{
		topic:          topic,
		partitionCount: make(map[int32]int),
	}
	for {
		err = cons.Consume(context.Background(), []string{handler.topic}, handler)
		if err != nil {
			fmt.Printf("consume failde, err=%v", err)
		}
	}
}

type Handler struct {
	format         string
	topic          string
	partitionCount map[int32]int
	totalCount     int
	mu             sync.Mutex
}

func (h *Handler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("setup")
	return nil
}

func (h *Handler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("clean up")
	return nil
}

func (h *Handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		record, err := avro.DeserializeRecord(bytes.NewReader(msg.Value))
		if err != nil {
			panic(err)
		}
		printRecord(&record)
		session.MarkMessage(msg, "")
		session.Commit()
	}
	return nil
}

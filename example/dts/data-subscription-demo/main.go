package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	proto "github.com/volcengine/volc-sdk-golang/example/dts/data-subscription-demo/proto"

	"flag"

	"github.com/Shopify/sarama"
	protobuf "google.golang.org/protobuf/proto"
)

type Handler struct {
	topic          string
	partitionCount map[int32]int
	totalCount     int
	mu             sync.Mutex
}

type Config struct {
	username string
	password string
	topic    string
	group    string
	addr     string
}

var (
	c Config
)

func init() {
	flag.StringVar(&c.username, "user", "", "data subscription user name")
	flag.StringVar(&c.password, "pwd", "", "data subscription password")
	flag.StringVar(&c.topic, "topic", "", "data subscription topic")
	flag.StringVar(&c.group, "group", "", "data subscription group")
	flag.StringVar(&c.addr, "addr", "", "data subscription kafka brokers")
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
	fmt.Println("ConsumeClaim")
	for m := range claim.Messages() {
		h.handleMsg(m)
		session.MarkMessage(m, "")
		session.Commit()
	}
	return nil
}

func getValueString(str *string) string {
	if str == nil {
		return "NULL"
	}
	return *str
}

func (h *Handler) handleMsg(msg *sarama.ConsumerMessage) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.totalCount++
	h.partitionCount[msg.Partition]++

	entry := &proto.Entry{}
	if err := protobuf.Unmarshal(msg.Value, entry); err != nil {
		panic(err)
	}

	fmt.Println("-------------- handle message --------------")
	fmt.Printf("get message EventType:%v\n", entry.EntryType.String())
	switch entry.GetEntryType() {
	case proto.EntryType_DDL:
		event := entry.GetDdlEvent()
		fmt.Printf("ddl %v\n", event.Sql)
	case proto.EntryType_DML:
		event := entry.GetDmlEvent()
		cols := event.ColumnDefs
		for i, row := range event.Rows {
			var before, after []string
			for _, col := range row.BeforeCols {
				before = append(before, fmt.Sprintf("%+v[%+v]", cols[i].GetName(), col.GetValue()))
			}
			for _, col := range row.AfterCols {
				after = append(after, fmt.Sprintf("%+v[%+v]", cols[i].GetName(), col.GetValue()))
			}
			fmt.Printf("get row before=%v after=%v\n", before, after)
		}
	}

	fmt.Printf("fetch message partition=%v key=%v\n", msg.Partition, string(msg.Key))
	fmt.Printf("count partition-count=%v total-count=%v\n", h.partitionCount, h.totalCount)
}

func main() {
	flag.Parse()
	fmt.Printf("config: %+v", c)

	config := sarama.NewConfig()
	config.Net.SASL.User = c.username
	config.Net.SASL.Password = c.password
	config.Net.SASL.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	topic := c.topic
	group := c.group

	addr := strings.Split(c.addr, ",")
	cons, err := sarama.NewConsumerGroup(addr, group, config)
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
			log.Fatalln(err)
		}
	}
}

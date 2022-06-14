package producer

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func GenerateLog(logTime int64, addLogMap map[string]string) *pb.Log {
	content := []*pb.LogContent{}
	for key, value := range addLogMap {
		content = append(content, &pb.LogContent{
			Key:   key,
			Value: value,
		})
	}
	return &pb.Log{
		Time:     logTime,
		Contents: content,
	}
}

func GetTimeMs(t int64) int64 {
	return t / 1000 / 1000
}

func GetLogSize(log *pb.Log) int {
	return log.Size()
}

func GetLogListSize(logList []*pb.Log) int {
	sizeInBytes := 0
	for _, log := range logList {
		sizeInBytes += GetLogSize(log)
	}
	return sizeInBytes
}

func WithRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fn()
}

package cache

import (
	"IMProject/pkg/logging"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"

	"log"
	"strings"
)

// MessageProducer 定义接口，用于生产消息的结构体
type MessageProducer interface {
	ToMap() map[string]interface{}
}

// CreateStreamGroup 创建一个消费者组，并在必要时创建流
func CreateStreamGroup(ctx context.Context, streamName string, groupName string) {
	err := RC.XGroupCreateMkStream(ctx, streamName, groupName, "0").Err()
	// 忽略 BUSYGROUP 错误，如果该组已存在，则不需要再次创建
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Fatalf("XGroupCreateMkStream error: %v", err)
	}
}

// ProductSendMsg 生产者发送消息
func ProductSendMsg(ctx context.Context, streamName string, producer MessageProducer) {
	msgID, err := RC.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: producer.ToMap(), // 调用 ToMap 方法将数据转换为 map
		MaxLen: 1000,             // 设置最大长度为 1000
	}).Result()
	if err != nil {
		logging.Info(err)
	}
	fmt.Printf("Message ID added: %s\n", msgID)
}

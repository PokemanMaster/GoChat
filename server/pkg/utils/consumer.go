package utils

import (
	"context"
	"fmt"
	"github.com/PokemanMaster/GoChat/server/server/app/warehouse/model"
	"github.com/PokemanMaster/GoChat/server/server/common/cache"
	"github.com/PokemanMaster/GoChat/server/server/common/db"
	"github.com/go-redis/redis/v8"
	"log"
	"math"

	"strconv"
	"sync"
	"time"
)

const (
	maxRetries    = 5 // 最大连接重试次数
	maxAckRetries = 3 // 最大确认消息重试次数
)

// ConsumerReadMsg 消费者读取消息并减少库存（使用 sync.WaitGroup 管理多个协程）
func ConsumerReadMsg(streamName string, groupName string, consumerName string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	retryCount := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer shutting down...")
			return
		default:
			msgs, err := cache.RC.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    groupName,
				Consumer: consumerName,
				Streams:  []string{streamName, ">"},
				Count:    1,
				Block:    10 * time.Second,
			}).Result()

			if err != nil {
				if err == redis.Nil {
					continue // 没有新消息，继续
				}
				retryCount++
				if retryCount >= maxRetries {
					log.Printf("Failed to read from stream after %d attempts: %v", retryCount, err)
					retryCount = 0               // 重置重试次数，继续尝试连接
					time.Sleep(10 * time.Second) // 休眠一段时间后再重试
					continue
				}
				log.Printf("Failed to read from stream, retrying (%d/%d): %v", retryCount, maxRetries, err)
				time.Sleep(time.Duration(retryCount) * time.Second)
				continue
			}
			retryCount = 0 // 成功读取消息后，重置重试计数

			// 处理消息
			for _, msg := range msgs[0].Messages {
				fmt.Printf("Received: %s %s\n", msg.ID, msg.Values)

				// 使用 fmt.Sprintf 直接将所有字段转换为字符串，避免复杂的类型检查
				warehouseIDStr := fmt.Sprintf("%v", msg.Values["warehouse_id"])
				productIDStr := fmt.Sprintf("%v", msg.Values["product_id"])
				reduceNumStr := fmt.Sprintf("%v", msg.Values["reduce_num"])

				// 转换为 uint 类型
				warehouseID, err1 := strconv.ParseUint(warehouseIDStr, 10, 32)
				productID, err2 := strconv.ParseUint(productIDStr, 10, 32)
				reduceNum, err3 := strconv.ParseUint(reduceNumStr, 10, 32)
				if err1 != nil || err2 != nil || err3 != nil {
					log.Printf("Failed to parse message values: warehouse_id=%v, product_id=%v, reduce_num=%v, error=%v", warehouseIDStr, productIDStr, reduceNumStr, err)
					continue
				}

				// 调用减少库存的函数
				err = model.ReduceStock(db.DB, uint(warehouseID), uint(productID), uint(reduceNum))
				if err != nil {
					log.Printf("Failed to reduce stock: %v", err)
					continue
				}

				// 消息确认 (XAck)
				ackRetries := 0
				for {
					_, err = cache.RC.XAck(ctx, streamName, groupName, msg.ID).Result()
					if err != nil {
						ackRetries++
						if ackRetries >= maxAckRetries {
							log.Printf("Failed to ack message %s after %d attempts: %v", msg.ID, ackRetries, err)
							break
						}
						// 使用指数退避机制增加重试时间
						time.Sleep(time.Duration(math.Pow(2, float64(ackRetries))) * time.Second)
						continue
					}
					log.Printf("Message %s acknowledged by consumer %s", msg.ID, consumerName)
					break
				}
			}
		}
	}
}

package dao

import (
	"encoding/json"
	"fmt"
	Morder "github.com/PokemanMaster/GoChat/v1/server/app/order/model"
	Mwarehouse "github.com/PokemanMaster/GoChat/v1/server/app/warehouse/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache/rabbit"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/resp"

	"log"
	"time"
)

const MaxRetryCount = 3 // 最大重试次数

func Consumer1() {
	ch := rabbit.RMQ
	constant := rabbit.Constant{
		ExchangeType:     "topic",
		NormalExchange:   "exchange.yna",
		DeadExchange:     "exchange.dead.yna",
		NormalQueue:      "queue.order.yna",
		DeadQueue:        "queue.dead.order.yna",
		NormalRoutingKey: "key.order.#",
		DeadRoutingKey:   "key.dead.order.#",
		ConsumerName:     "yna1",
	}
	// 声明交换机
	err := rabbit.DeclareExchange(ch, constant.NormalExchange, constant.ExchangeType)
	resp.FailOnError(err, "Failed to declare a Exchange")

	// 声明死信交换机
	err = rabbit.DeclareDeadExchange(ch, constant.DeadExchange, constant.ExchangeType)
	resp.FailOnError(err, "Failed to declare a deadQueue")

	// 声明队列
	_, err = rabbit.DeclareQueue(ch, constant.NormalQueue, constant.DeadExchange, constant.DeadRoutingKey)
	resp.FailOnError(err, "Failed to declare a queue")

	// 声明死信队列
	_, err = rabbit.DeclareDeadQueue(ch, constant.DeadQueue)
	resp.FailOnError(err, "Failed to declare a queue")

	// 交换机绑定队列
	err = rabbit.BindQueue(ch, constant.NormalQueue, constant.NormalRoutingKey, constant.NormalExchange)
	resp.FailOnError(err, "Failed to bind a queue")

	// 死信队列绑定死信交换机
	err = rabbit.BindQueue(ch, constant.DeadQueue, constant.DeadRoutingKey, constant.DeadExchange)
	resp.FailOnError(err, "Failed to bind a deadQueue")

	// 声明消费者
	msgs, err := rabbit.DeclareConsume(ch, constant.NormalQueue, constant.ConsumerName)
	resp.FailOnError(err, "Failed to consume messages")

	// 定义消费者规则
	err = ch.Qos(1, 0, false)
	resp.FailOnError(err, "Failed to set QoS")

	// 消息消费和重试逻辑
	var forever chan struct{}
	go func() {
		for d := range msgs {
			var msg Morder.OrderDetail
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Error decoding message: %s", err)
				continue
			}
			retryCount := 0
			for {
				err := Mwarehouse.ReduceStock(db.DB, 1, msg.ProductID, msg.Num)
				// 进行重试
				if err != nil && retryCount < MaxRetryCount {
					log.Printf("Error processing message: %s. Retrying...", err)
					retryCount++
					time.Sleep(1 * time.Second)
					d.Nack(false, true) // 重新入队等待重试
					continue
				} else if err != nil {
					// 达到最大重试次数，转发到死信队列或做其他处理
					log.Printf("Failed to process message after %d retries: %s", retryCount, err)
					d.Nack(false, false) // 不重新入队，可能会将其发送到死信队列
					break
				}
				d.Ack(false)
				break
			}
		}
	}()
	fmt.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}

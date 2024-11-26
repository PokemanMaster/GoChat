package dao

import (
	"context"
	"encoding/json"
	"fmt"
	rabbit2 "github.com/PokemanMaster/GoChat/server/common/cache/rabbit"
	"github.com/PokemanMaster/GoChat/server/resp"

	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

// Product1 发送消息
func Product1(message interface{}) error {
	ch := rabbit2.RMQ
	constant := rabbit2.Constant{
		ExchangeType:      "topic",
		NormalExchange:    "exchange.yna",
		DeadExchange:      "exchange.dead.yna",
		NormalQueue:       "queue.order.yna",
		DeadQueue:         "queue.dead.order.yna",
		ProductRoutingKey: "key.order.warehouse",
		NormalRoutingKey:  "key.order.#",
		DeadRoutingKey:    "key.dead.order.#",
		ConsumerName:      "yna1",
	}
	// 声明交换机
	err := rabbit2.DeclareExchange(ch, constant.NormalExchange, constant.ExchangeType)
	resp.FailOnError(err, "Failed to declare a Exchange")

	// 设置超时上下文.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 确保开启确认模式
	if err := ch.Confirm(false); err != nil {
		fmt.Println("Could not put channel into confirm mode")
	}

	// 异步获取确认消息
	confirmChan := ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	returnChan := ch.NotifyReturn(make(chan amqp.Return))

	// 发送消息
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx,
		constant.NormalExchange,    // exchange
		constant.ProductRoutingKey, // routing key
		true,                       // mandatory 设置为 true 可以确保消息在无法路由到任何队列时退回给生产者。
		false,                      // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // 持久化消息
		})

	// 如果消息发送失败
	if err != nil {
		return err
	}

	// 等待确认消息
	select {
	case confirm := <-confirmChan:
		if confirm.Ack {
			fmt.Println("Sent message acknowledged", confirm.Ack)
		} else {
			fmt.Println("Sent message not acknowledged", confirm.Ack)
		}
	case returnMsg := <-returnChan:
		fmt.Println("Message returned", string(returnMsg.Body))
	case <-time.After(5 * time.Second): // 超过5秒超时
		fmt.Println("Timeout waiting for message confirmation")
	}

	return nil
}

package dao

import (
	"github.com/PokemanMaster/GoChat/v1/server/common/cache/rabbit"
	"github.com/PokemanMaster/GoChat/v1/server/resp"

	"log"
)

func Consumer2() {
	ch := rabbit.RMQ
	constant := rabbit.Constant{
		ExchangeType:     "topic",
		NormalExchange:   "exchange.yna",
		DeadExchange:     "exchange.dead.yna",
		NormalQueue:      "queue.order.yna",
		DeadQueue:        "queue.dead.order.yna",
		NormalRoutingKey: "key.order.#",
		DeadRoutingKey:   "key.dead.order.#",
		ConsumerName:     "yna2",
	}

	// 声明死信交换机
	err := rabbit.DeclareExchange(ch, constant.DeadExchange, constant.ExchangeType)
	resp.FailOnError(err, "Failed to Declare a exchange")

	// 声明消费者
	msgs, err := rabbit.DeclareConsume(ch, constant.DeadQueue, constant.ConsumerName)
	resp.FailOnError(err, "Failed to consume messages")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

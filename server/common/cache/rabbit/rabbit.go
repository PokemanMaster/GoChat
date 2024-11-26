package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Constant struct {
	ExchangeType      string
	NormalExchange    string
	DeadExchange      string
	NormalQueue       string
	DeadQueue         string
	ProductRoutingKey string
	NormalRoutingKey  string
	DeadRoutingKey    string
	ConsumerName      string
	ProductName       string
}

// DeclareExchange 声明交换机
func DeclareExchange(ch *amqp.Channel, exchangeName string, exchangeType string) error {
	return ch.ExchangeDeclare(
		exchangeName, // exchange name
		exchangeType, // exchange type：topic、fan、direct
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

// DeclareDeadExchange 声明死信交换机
func DeclareDeadExchange(ch *amqp.Channel, deadExchangeName string, deadExchangeType string) error {
	return ch.ExchangeDeclare(
		deadExchangeName, // exchange name
		deadExchangeType, // exchange type：topic、fan、direct
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
}

// DeclareQueue 声明队列
func DeclareQueue(ch *amqp.Channel, queueName string, deadExchangeName string, deadRoutingKey string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{ // arguments for the queue
			//"x-message-ttl":             5000,                    // 指定过期时间
			//"x-max-length":              6,                        // 指定长度。超过这个长度的消息会发送到dead_exchange中
			"x-dead-letter-exchange":    deadExchangeName, // 指定死信交换机
			"x-dead-letter-routing-key": deadRoutingKey,   // 指定死信routing-key
		},
	)
}

// DeclareDeadQueue 声明死信队列
func DeclareDeadQueue(ch *amqp.Channel, deadQueueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		deadQueueName, // queue name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,
	)
}

// BindQueue 绑定队列和交换机
func BindQueue(ch *amqp.Channel, queueName string, routingKey string, exchangeName string) error {
	return ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
}

// DeclareConsume 声明消费者
func DeclareConsume(ch *amqp.Channel, queueName string, consumerName string) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		queueName,    // queue
		consumerName, // consumer
		false,        // 设置手动应答模式
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
}

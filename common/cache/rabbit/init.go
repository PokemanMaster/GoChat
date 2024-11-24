package rabbit

import (
	"github.com/PokemanMaster/GoChat/resp"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

var RMQ *amqp.Channel

func InitRabbitMQ() {
	driverName := viper.GetString("rabbit.driverName")
	userName := viper.GetString("rabbit.userName")
	password := viper.GetString("rabbit.password")
	address := viper.GetString("rabbit.address")
	conn, err := amqp.Dial(driverName + "://" + userName + ":" + password + "@" + address + "/")
	resp.FailOnError(err, "Failed to connect to RabbitMQ")
	RMQ, err = conn.Channel()
	resp.FailOnError(err, "Failed to open a channel")
}

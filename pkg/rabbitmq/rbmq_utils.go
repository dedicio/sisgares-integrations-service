package rbmq_utils

import (
	"github.com/streadway/amqp"
)

func DeclareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

func DeclareExchange(ch *amqp.Channel, exchangeName string) error {
	return ch.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
